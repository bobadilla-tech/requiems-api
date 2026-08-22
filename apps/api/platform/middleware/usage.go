package middleware

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
	"time"

	chimw "github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"

	"requiems-api/platform/httpx"
)

// quotaIncrementIfPresentScript and quotaBootstrapScript are the same
// EXISTS-then-INCR / SET-then-INCR shape as
// services/technology/counter/redis_mutations.go's
// incrementIfPresentScript/incrementWithBootstrapScript, reimplemented here
// rather than imported: counter is a leaf feature package backing a public
// product API and shouldn't gain an auth-domain dependency, and its
// namespace->total data model isn't meant to be reused as-is. The one
// addition here is the EXPIRE in the bootstrap script, set only when the key
// is actually created — this is what makes a new billing cycle "just miss
// and rebootstrap from zero" with no explicit reset job.
var quotaIncrementIfPresentScript = redis.NewScript(`
if redis.call("EXISTS", KEYS[1]) == 0 then
  return -1
end
return redis.call("INCR", KEYS[1])
`)

var quotaBootstrapScript = redis.NewScript(`
if redis.call("EXISTS", KEYS[1]) == 0 then
  redis.call("SET", KEYS[1], ARGV[1])
  redis.call("EXPIRE", KEYS[1], ARGV[2])
end
return redis.call("INCR", KEYS[1])
`)

// UsageQuota enforces a per-user monthly request quota and writes a
// row-level usage_logs record for every request that clears it. The two
// pieces run in this order, per request:
//
//  1. Quota check (Lua bootstrap-and-increment against Redis, falling back
//     to a direct Postgres SUM if Redis is unavailable). The counter
//     increments unconditionally, before the handler runs — a rejected
//     request still consumes one unit of quota, matching the legacy
//     Worker's check-then-serve-or-reject semantics. It counts *checked*
//     requests, not *served* ones, by design.
//  2. A synchronous INSERT into usage_logs after the handler runs, skipped
//     entirely for a request the quota check rejected.
type UsageQuota struct {
	pool   *pgxpool.Pool
	rdb    *redis.Client
	plans  *PlanCache
	logger *slog.Logger
}

// NewUsageQuota builds a UsageQuota. plans should be shared with RateLimiter.
func NewUsageQuota(pool *pgxpool.Pool, rdb *redis.Client, plans *PlanCache, logger *slog.Logger) *UsageQuota {
	if logger == nil {
		logger = slog.Default()
	}

	return &UsageQuota{pool: pool, rdb: rdb, plans: plans, logger: logger}
}

// Middleware returns the enforcing http middleware. Mount after APIKeyAuth
// (and, conventionally, after RateLimiter) in the same protected group — it
// reads APIKeyPrincipalFromContext and has nothing to key on without it.
func (u *UsageQuota) Middleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			principal, ok := APIKeyPrincipalFromContext(r.Context())
			if !ok {
				// No authenticated principal to key on; see the identical
				// note in ratelimit.go — shouldn't be reachable given how
				// this is mounted, but never panic over a wiring mistake.
				next.ServeHTTP(w, r)
				return
			}

			cycle := cycleStart(principal.CurrentPeriodStart, time.Now())
			key := fmt.Sprintf("usage:%d:%d", principal.UserID, cycle.Unix())

			limits := u.plans.get(r.Context(), principal.Plan)

			count, err := u.increment(r.Context(), key, principal.UserID, cycle)
			if err != nil {
				// Both Redis and Postgres failed to affirmatively check
				// quota. Unlike rate limiting, usage/quota accounting must
				// not silently fail-open here (the audit is explicit: doing
				// so would mean permanently losing those usage records) —
				// and step 2's row-level write couldn't happen either way,
				// since it needs the same Postgres connection.
				u.logger.Error("usage quota: redis and postgres both unavailable, failing closed",
					"error", err, "user_id", principal.UserID)
				httpx.Error(w, http.StatusServiceUnavailable, "service_unavailable", "Usage tracking temporarily unavailable")
				return
			}

			if limits.RequestLimit != nil && count > *limits.RequestLimit {
				httpx.Error(w, http.StatusTooManyRequests, "quota_exceeded", "Monthly request quota exceeded")
				return
			}

			ww := chimw.NewWrapResponseWriter(w, r.ProtoMajor)
			start := time.Now()

			next.ServeHTTP(ww, r)

			u.recordUsage(r.Context(), principal, r, ww, start)
		})
	}
}

// increment returns the post-increment quota count for key. It tries the
// Redis fast path first (bootstrapping from a Postgres baseline on a cold
// key), and falls back to a direct Postgres SUM — enforced as sum+1 for this
// request, since no Redis increment happened — if Redis errors, times out,
// or rejects the write (e.g. "OOM command not allowed" under noeviction at
// capacity: handled identically to a connection error, not special-cased).
// An error is returned only when that Postgres fallback itself also fails.
func (u *UsageQuota) increment(ctx context.Context, key string, userID int64, cycle time.Time) (int64, error) {
	rctx, cancel := context.WithTimeout(ctx, redisCallTimeout)
	val, err := quotaIncrementIfPresentScript.Run(rctx, u.rdb, []string{key}).Int64()
	cancel()

	if err == nil && val != -1 {
		return val, nil
	}

	if err == nil && val == -1 {
		// Cache miss: bootstrap from Postgres, summed across all of this
		// user's API keys (not just the one on this request) — a user with
		// more than one active key must have all of them summed into the
		// baseline, or it silently under-counts. Mirrors the legacy
		// Worker's own quota-cache-miss D1 read.
		baseline, sumErr := u.sumUsage(ctx, userID, cycle)
		if sumErr != nil {
			return 0, sumErr
		}

		ttl := time.Until(cycle.AddDate(0, 1, 0))
		if ttl <= 0 {
			// Shouldn't happen — cycleStart always returns a boundary at or
			// before now, so cycle+1 month is always in the future — but
			// guard against a zero/negative EXPIRE deleting the key
			// immediately if it ever does.
			ttl = time.Minute
		}

		rctx2, cancel2 := context.WithTimeout(ctx, redisCallTimeout)
		val, err = quotaBootstrapScript.Run(rctx2, u.rdb, []string{key}, baseline, int64(ttl.Seconds())).Int64()
		cancel2()

		if err == nil {
			return val, nil
		}

		u.logger.Warn("usage quota: redis bootstrap-increment failed after a successful baseline read, enforcing on the baseline alone for this request",
			"error", err, "user_id", userID)
		return baseline + 1, nil
	}

	u.logger.Warn("usage quota: redis error, falling back to postgres", "error", err, "user_id", userID)

	sum, sumErr := u.sumUsage(ctx, userID, cycle)
	if sumErr != nil {
		return 0, sumErr
	}

	return sum + 1, nil
}

func (u *UsageQuota) sumUsage(ctx context.Context, userID int64, cycle time.Time) (int64, error) {
	var sum int64

	err := u.pool.QueryRow(ctx, `
		SELECT COALESCE(SUM(credits_used), 0) FROM usage_logs WHERE user_id = $1 AND used_at >= $2
	`, userID, cycle).Scan(&sum)

	return sum, err
}

// recordUsage writes the row-level usage_logs ledger entry for a request the
// quota check let through. credits_used defaults to 1 and is only
// overridden by the X-Usage-Count response header — the same header
// httpx.Handle/HandleBatch already set for the (today, exactly one) response
// type implementing httpx.UsageCounter. On conflict with an existing row for
// the same (api_key_id, used_at, endpoint) — a known, accepted dedup
// collision under rapid same-second traffic, not addressed in this phase —
// the row is silently dropped, matching Postgres's own DO NOTHING. If the
// write itself fails (a real error, not a conflict), it's logged and
// dropped: the response has already been sent, so there's nothing left to
// reject.
func (u *UsageQuota) recordUsage(ctx context.Context, principal APIKeyPrincipal, r *http.Request, ww chimw.WrapResponseWriter, start time.Time) {
	creditsUsed := 1

	if raw := ww.Header().Get("X-Usage-Count"); raw != "" {
		if n, err := strconv.Atoi(raw); err == nil && n > 0 {
			creditsUsed = n
		}
	}

	err := u.insertUsageRow(ctx, principal, r.URL.Path, r.Method, creditsUsed,
		ww.Status(), time.Since(start).Milliseconds(), time.Now().UTC())

	if err != nil {
		u.logger.Error("usage quota: failed to write usage_logs row, dropping",
			"error", err, "user_id", principal.UserID, "api_key_id", principal.APIKeyID)
	}
}

// insertUsageRow is split out from recordUsage so tests can force a
// (api_key_id, used_at, endpoint) collision with an explicit usedAt, rather
// than racing time.Now()'s microsecond precision.
func (u *UsageQuota) insertUsageRow(
	ctx context.Context, principal APIKeyPrincipal, endpoint, method string,
	creditsUsed, statusCode int, responseTimeMs int64, usedAt time.Time,
) error {
	_, err := u.pool.Exec(ctx, `
		INSERT INTO usage_logs (
			user_id, api_key_id, endpoint, credits_used, request_method,
			status_code, response_time_ms, used_at, usage_date, created_at, updated_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $8, $8)
		ON CONFLICT (api_key_id, used_at, endpoint) DO NOTHING
	`,
		principal.UserID, principal.APIKeyID, endpoint, creditsUsed, method,
		statusCode, responseTimeMs, usedAt, usedAt.Format("2006-01-02"),
	)

	return err
}

// cycleStart returns the most recent occurrence of anchor's day-of-month, at
// midnight UTC, at or before now — mirroring the legacy Worker's
// getResetTime (apps/workers/auth-gateway/src/requests.ts): a billing cycle
// rolls over every month on the same day it started. This lets a static
// anchor (free-tier users fall back to their key's created_at, which never
// changes) still roll into a new monthly cycle automatically, rather than
// pinning the quota window to whatever moment the key happened to be read.
func cycleStart(anchor, now time.Time) time.Time {
	anchor = anchor.UTC()
	now = now.UTC()

	candidate := time.Date(now.Year(), now.Month(), anchor.Day(), 0, 0, 0, 0, time.UTC)
	if candidate.After(now) {
		candidate = candidate.AddDate(0, -1, 0)
	}

	return candidate
}
