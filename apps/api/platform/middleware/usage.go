package middleware

import (
	"context"
	"fmt"
	"log/slog"
	"math"
	"net/http"
	"strconv"
	"strings"
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
return redis.call("INCRBY", KEYS[1], ARGV[1])
`)

var quotaBootstrapScript = redis.NewScript(`
if redis.call("EXISTS", KEYS[1]) == 0 then
  redis.call("SET", KEYS[1], ARGV[1])
  redis.call("EXPIRE", KEYS[1], ARGV[2])
end
return redis.call("INCRBY", KEYS[1], ARGV[3])
`)

var quotaAdjustScript = redis.NewScript(`
if redis.call("EXISTS", KEYS[1]) == 0 then
  return -1
end
return redis.call("INCRBY", KEYS[1], ARGV[1])
`)

// UsageQuota enforces a per-user monthly request quota and writes a
// row-level usage_logs record for every request that clears it. The two
// pieces run in this order, per request:
//
//  1. Quota check (Lua bootstrap-and-increment against Redis). Redis is the
//     atomic reservation store; an unavailable reservation path fails closed.
//     The counter reserves the authoritative route cost before the handler
//     runs. If the reservation is over quota, it is rolled back because the
//     request was not admitted.
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

			limits, hasStaleLimits, planErr := u.plans.get(r.Context(), principal.Plan)
			if planErr != nil {
				u.logger.Error("usage quota: plan lookup failed", "error", planErr, "user_id", principal.UserID)
				// A stale plan entry is safe to use; without one, quota
				// enforcement cannot be established safely.
				if !hasStaleLimits {
					httpx.Error(w, http.StatusServiceUnavailable, "service_unavailable", "Usage limits temporarily unavailable")
					return
				}
			}

			requestedCredits := requestCredits(r)
			count, err := u.increment(r.Context(), key, principal.UserID, cycle, requestedCredits)
			if err != nil {
				// Quota reservation must not silently fail open. A bare
				// Postgres SUM cannot reserve quota atomically under
				// concurrency, so Redis reservation errors fail closed.
				u.logger.Error("usage quota: atomic reservation unavailable, failing closed",
					"error", err, "user_id", principal.UserID)
				httpx.Error(w, http.StatusServiceUnavailable, "service_unavailable", "Usage tracking temporarily unavailable")
				return
			}

			if limits.RequestLimit != nil && count > *limits.RequestLimit {
				if err := u.adjust(r.Context(), key, -requestedCredits); err != nil {
					u.logger.Error("usage quota: failed to roll back rejected reservation", "error", err, "user_id", principal.UserID)
				}
				httpx.Error(w, http.StatusTooManyRequests, "quota_exceeded", "Monthly request quota exceeded")
				return
			}

			ww := chimw.NewWrapResponseWriter(w, r.ProtoMajor)
			start := time.Now()

			next.ServeHTTP(ww, r)

			creditsUsed := responseCredits(ww.Header())
			if creditsUsed != requestedCredits {
				if err := u.adjust(r.Context(), key, creditsUsed-requestedCredits); err != nil {
					u.logger.Error("usage quota: failed to reconcile Redis credits", "error", err, "user_id", principal.UserID)
				}
			}

			u.recordUsage(r.Context(), principal, r, ww, start, creditsUsed)
		})
	}
}

// increment returns the post-increment quota count for key. It uses Redis for
// the atomic reservation, bootstrapping from a Postgres baseline on a cold key.
// If Redis errors, times out, or rejects the write (e.g. "OOM command not
// allowed" under noeviction at capacity), it returns an error rather than
// attempting a non-atomic Postgres SUM fallback.
func (u *UsageQuota) increment(ctx context.Context, key string, userID int64, cycle time.Time, credits int64) (int64, error) {
	rctx, cancel := context.WithTimeout(ctx, redisCallTimeout)
	val, err := quotaIncrementIfPresentScript.Run(rctx, u.rdb, []string{key}, credits).Int64()
	cancel()

	if err == nil && val != -1 {
		return val, nil
	}

	if err == nil && val == -1 {
		// Cache miss: bootstrap from Postgres, summed across all of this
		// user's API keys (not just the one on this request) — a user with
		// more than one active key must have all of them summed into the
		// baseline, or it silently under-counts.
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
		val, err = quotaBootstrapScript.Run(rctx2, u.rdb, []string{key}, baseline, int64(ttl.Seconds()), credits).Int64()
		cancel2()

		if err == nil {
			return val, nil
		}

		return 0, fmt.Errorf("redis quota reservation failed after baseline read: %w", err)
	}

	return 0, fmt.Errorf("redis quota reservation failed: %w", err)
}

func (u *UsageQuota) adjust(ctx context.Context, key string, delta int64) error {
	if delta == 0 {
		return nil
	}

	rctx, cancel := context.WithTimeout(ctx, redisCallTimeout)
	defer cancel()

	value, err := quotaAdjustScript.Run(rctx, u.rdb, []string{key}, delta).Int64()
	if err != nil {
		return err
	}
	if value == -1 {
		return fmt.Errorf("quota key %q disappeared during reconciliation", key)
	}
	return nil
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
func (u *UsageQuota) recordUsage(ctx context.Context, principal APIKeyPrincipal, r *http.Request, ww chimw.WrapResponseWriter, start time.Time, creditsUsed int64) {
	writeCtx, cancel := context.WithTimeout(context.WithoutCancel(ctx), 5*time.Second)
	defer cancel()

	err := u.insertUsageRow(writeCtx, principal, r.URL.Path, r.Method, creditsUsed,
		ww.Status(), time.Since(start).Milliseconds(), time.Now().UTC())

	if err != nil {
		u.logger.Error("usage quota: failed to write usage_logs row, dropping",
			"error", err, "user_id", principal.UserID)
	}
}

// requestCredits returns the static, authoritative cost for a route before
// dispatch. Dynamic batch costs are reconciled from the response header after
// the handler runs; the request's X-Usage-Count header is never trusted.
func requestCredits(r *http.Request) int64 {
	if r.Method != http.MethodGet {
		return 1
	}

	for _, route := range []string{"/v1/text/dictionary", "/v1/text/thesaurus"} {
		if r.URL.Path == route || strings.HasPrefix(r.URL.Path, route+"/") {
			return 2
		}
	}

	return 1
}

func responseCredits(headers http.Header) int64 {
	return headerCredits(headers)
}

// headerCredits parses X-Usage-Count, bounded to [1, math.MaxInt32] so the
// usage_logs.credits_used is an integer column, so values outside this range
// are treated as the default unit charge rather than being persisted.
func headerCredits(headers http.Header) int64 {
	if raw := headers.Get("X-Usage-Count"); raw != "" {
		if n, err := strconv.ParseInt(raw, 10, 64); err == nil && n > 0 && n <= math.MaxInt32 {
			return n
		}
	}
	return 1
}

// insertUsageRow is split out from recordUsage so tests can force a
// (api_key_id, used_at, endpoint) collision with an explicit usedAt, rather
// than racing time.Now()'s microsecond precision.
func (u *UsageQuota) insertUsageRow(
	ctx context.Context, principal APIKeyPrincipal, endpoint, method string,
	creditsUsed int64, statusCode int, responseTimeMs int64, usedAt time.Time,
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
// midnight UTC, at or before now: a billing cycle
// rolls over every month on the same day it started. This lets a static
// anchor (free-tier users fall back to their key's created_at, which never
// changes) still roll into a new monthly cycle automatically, rather than
// pinning the quota window to whatever moment the key happened to be read.
func cycleStart(anchor, now time.Time) time.Time {
	anchor = anchor.UTC()
	now = now.UTC()

	anchorDay := anchor.Day()
	monthStart := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.UTC)
	currentMonthEnd := monthStart.AddDate(0, 1, -1).Day()
	currentDay := anchorDay
	if currentDay > currentMonthEnd {
		currentDay = currentMonthEnd
	}
	candidate := time.Date(now.Year(), now.Month(), currentDay, 0, 0, 0, 0, time.UTC)
	if candidate.After(now) {
		monthStart = monthStart.AddDate(0, -1, 0)
	}

	// Derive the candidate month before clamping the original anchor day. This
	// keeps a January 31 anchor on January 31 during February instead of first
	// turning it into February 28 and then subtracting a month.
	monthEnd := monthStart.AddDate(0, 1, -1).Day()
	if anchorDay > monthEnd {
		anchorDay = monthEnd
	}

	return time.Date(monthStart.Year(), monthStart.Month(), anchorDay, 0, 0, 0, 0, time.UTC)
}
