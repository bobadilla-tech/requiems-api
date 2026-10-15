package middleware

import (
	"context"
	"fmt"
	"log/slog"
	"math"
	"net/http"
	"strconv"
	"strings"
	"sync/atomic"
	"time"

	"github.com/getsentry/sentry-go"
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

	// matcher is late-bound (the router is built after this middleware) and
	// nil until SetRouteMatcher is called; a nil matcher fails open.
	matcher atomic.Pointer[RouteMatcher]
}

// RouteMatcher reports whether the router has a handler for method+path.
type RouteMatcher func(method, path string) bool

// SetRouteMatcher installs the pre-dispatch route check. Requests the matcher
// rejects (unknown routes, wrong methods) skip quota, ledger and Redis work
// entirely: they are not usage. Safe to call after the middleware is mounted.
func (u *UsageQuota) SetRouteMatcher(m RouteMatcher) {
	if m == nil {
		u.matcher.Store(nil)
		return
	}

	u.matcher.Store(&m)
}

// routeMissKey carries a *routeMiss through the request context so the
// router's NotFound/MethodNotAllowed handlers can tell UsageQuota that the
// router itself answered, i.e. no real route ran.
type routeMissKey struct{}

type routeMiss struct {
	flagged atomic.Bool
	// matcher lets MethodNotAllowedHandler rebuild the Allow header, which
	// chi only hands to its own default 405 handler.
	matcher RouteMatcher
}

// allowedMethods probes the router for the methods that would serve path.
func (m *routeMiss) allowedMethods(path string) []string {
	if m == nil || m.matcher == nil {
		return nil
	}

	var allowed []string

	defer func() {
		if rec := recover(); rec != nil {
			allowed = nil
		}
	}()

	for _, method := range []string{
		http.MethodConnect, http.MethodDelete, http.MethodGet, http.MethodHead, http.MethodOptions,
		http.MethodPatch, http.MethodPost, http.MethodPut, http.MethodTrace,
	} {
		if m.matcher(method, path) {
			allowed = append(allowed, method)
		}
	}

	return allowed
}

// NotFoundHandler is the 404 handler for every /v1 router: it flags the
// request as a route miss (so a reserved credit is refunded and no ledger row
// is written when pre-matching was unavailable) and writes chi's default 404.
func NotFoundHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		flagRouteMiss(r.Context())
		http.NotFound(w, r)
	}
}

// MethodNotAllowedHandler is the 405 counterpart of NotFoundHandler.
func MethodNotAllowedHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		miss, _ := r.Context().Value(routeMissKey{}).(*routeMiss)
		if miss != nil {
			miss.flagged.Store(true)

			if allowed := miss.allowedMethods(requestRoutePath(r)); len(allowed) > 0 {
				w.Header().Set("Allow", strings.Join(allowed, ", "))
			}
		}

		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// requestRoutePath is the path chi routes on: RawPath when set, else Path.
func requestRoutePath(r *http.Request) string {
	if r.URL.RawPath != "" {
		return r.URL.RawPath
	}

	return r.URL.Path
}

// TrackRouteMiss returns a context in which NotFoundHandler and
// MethodNotAllowedHandler record that the router answered, and a function that
// reports it. UsageQuota uses it per request; router tests use it to prove a
// mount carries the handlers.
func TrackRouteMiss(ctx context.Context) (tracked context.Context, missed func() bool) {
	return trackRouteMiss(ctx, nil)
}

func trackRouteMiss(ctx context.Context, matcher RouteMatcher) (tracked context.Context, missed func() bool) {
	miss := &routeMiss{matcher: matcher}

	return context.WithValue(ctx, routeMissKey{}, miss), miss.flagged.Load
}

func flagRouteMiss(ctx context.Context) {
	if m, ok := ctx.Value(routeMissKey{}).(*routeMiss); ok {
		m.flagged.Store(true)
	}
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

			// A request the router has no handler for is not usage: no
			// reservation, no ledger row. Fails open when no matcher is set or
			// it misbehaves.
			if !u.routeExists(r) {
				u.recordRouteMiss(r.Context(), principal)

				// The router still answers; give a 405 what it needs to
				// rebuild its Allow header.
				if m := u.matcher.Load(); m != nil {
					missCtx, _ := trackRouteMiss(r.Context(), *m)
					r = r.WithContext(missCtx)
				}

				next.ServeHTTP(w, r)
				return
			}

			// Taken once so the cycle key and the ledger's used_at agree.
			now := time.Now()
			cycle := cycleStart(principal.CurrentPeriodStart, now)
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
			var matcher RouteMatcher
			if m := u.matcher.Load(); m != nil {
				matcher = *m
			}

			trackedCtx, routeMissed := trackRouteMiss(r.Context(), matcher)
			r = r.WithContext(trackedCtx)

			handlerStart := time.Now()
			recovered, panicked := u.serve(ww, r, next, principal)

			u.settle(r, principal, key, requestedCredits, now, handlerStart, ww, routeMissed(), panicked)

			if panicked && recovered == http.ErrAbortHandler {
				// An intentional abort must still reach net/http.
				panic(http.ErrAbortHandler)
			}
		})
	}
}

// settle reconciles the reservation with what actually happened and writes the
// ledger row.
func (u *UsageQuota) settle(
	r *http.Request, principal APIKeyPrincipal, key string, requestedCredits int64,
	now, handlerStart time.Time, ww chimw.WrapResponseWriter, routeMissed, panicked bool,
) {
	if routeMissed {
		// The router answered 404/405 itself (pre-matching was not
		// available): give the reservation back, write no ledger row.
		if err := u.adjust(r.Context(), key, -requestedCredits); err != nil {
			u.logger.Error("usage quota: failed to refund route miss", "error", err, "user_id", principal.UserID)
		}
		u.recordRouteMiss(r.Context(), principal)

		return
	}

	status := ww.Status()
	if status == 0 {
		// The handler never called WriteHeader: net/http sends 200.
		status = http.StatusOK
	}
	if panicked {
		status = http.StatusInternalServerError
	}

	// Client (4xx) and server (5xx) errors are never charged. The override runs
	// before the reconcile below so Redis and the ledger agree, even when
	// X-Usage-Count was set.
	creditsUsed := responseCredits(ww.Header())
	if status >= http.StatusBadRequest {
		creditsUsed = 0
	}

	if creditsUsed != requestedCredits {
		if err := u.adjust(r.Context(), key, creditsUsed-requestedCredits); err != nil {
			u.logger.Error("usage quota: failed to reconcile Redis credits", "error", err, "user_id", principal.UserID)
		}
	}

	u.recordUsage(r.Context(), principal, r, status, now, handlerStart, creditsUsed)
}

// serve runs the downstream handler and recovers a panic so the reservation
// is reconciled and the request is logged. If nothing was written yet it
// answers a JSON 500. http.ErrAbortHandler is returned (not swallowed) so the
// caller can finish accounting and re-panic.
func (u *UsageQuota) serve(ww chimw.WrapResponseWriter, r *http.Request, next http.Handler, principal APIKeyPrincipal) (recovered any, panicked bool) {
	defer func() {
		rec := recover()
		if rec == nil {
			return
		}

		recovered, panicked = rec, true

		if rec != http.ErrAbortHandler {
			u.logger.Error("usage quota: handler panicked", "panic", rec, "user_id", principal.UserID, "path", r.URL.Path)
			sentry.CurrentHub().Recover(rec)

			if ww.Status() == 0 {
				httpx.Error(ww, http.StatusInternalServerError, "internal_error", "internal server error")
			}
		}
	}()

	next.ServeHTTP(ww, r)

	return nil, false
}

// routeExists asks the late-bound matcher whether the router would handle r.
func (u *UsageQuota) routeExists(r *http.Request) (exists bool) {
	m := u.matcher.Load()
	if m == nil {
		return true
	}

	defer func() {
		if rec := recover(); rec != nil {
			exists = true
		}
	}()

	return (*m)(r.Method, requestRoutePath(r))
}

// routeMissTTL bounds the per-day route-miss buckets.
const routeMissTTL = 8 * 24 * time.Hour

// recordRouteMiss keeps a best-effort per-user, per-day count of requests for
// routes that do not exist, so scanner traffic is not lost now that it is no
// longer in usage_logs. Nothing reads the bucket yet (surfacing it to the
// admin abuse tooling is a follow-on). A Redis error never affects the request.
func (u *UsageQuota) recordRouteMiss(ctx context.Context, principal APIKeyPrincipal) {
	u.logger.Debug("usage quota: route miss", "user_id", principal.UserID)

	if u.rdb == nil {
		return
	}

	rctx, cancel := context.WithTimeout(context.WithoutCancel(ctx), redisCallTimeout)
	defer cancel()

	bucket := "route_miss:" + time.Now().UTC().Format("2006-01-02")

	pipe := u.rdb.Pipeline()
	pipe.HIncrBy(rctx, bucket, strconv.FormatInt(principal.UserID, 10), 1)
	pipe.Expire(rctx, bucket, routeMissTTL)

	if _, err := pipe.Exec(rctx); err != nil {
		u.logger.Debug("usage quota: route miss counter unavailable", "error", err, "user_id", principal.UserID)
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

// adjust applies a delta to the reserved counter. It detaches from the request
// context's cancellation: a client that disconnects mid-request must not leave
// Redis over-counted while the ledger says the request was free.
func (u *UsageQuota) adjust(ctx context.Context, key string, delta int64) error {
	if delta == 0 {
		return nil
	}

	rctx, cancel := context.WithTimeout(context.WithoutCancel(ctx), redisCallTimeout)
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
func (u *UsageQuota) recordUsage(ctx context.Context, principal APIKeyPrincipal, r *http.Request, status int, usedAt, handlerStart time.Time, creditsUsed int64) {
	writeCtx, cancel := context.WithTimeout(context.WithoutCancel(ctx), 5*time.Second)
	defer cancel()

	// used_at is the request start (the same instant the cycle key was
	// derived from); latency is the handler's alone; status is the normalised
	// one, never a raw 0.
	err := u.insertUsageRow(writeCtx, principal, r.URL.Path, r.Method, creditsUsed,
		status, time.Since(handlerStart).Milliseconds(), usedAt.UTC())

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

// headerCredits parses X-Usage-Count, bounded to [0, math.MaxInt32] so the
// usage_logs.credits_used is an integer column. An absent header is the default
// unit charge; an explicit "0" (an empty batch) charges nothing; anything else
// out of range or malformed is treated as the default rather than persisted.
func headerCredits(headers http.Header) int64 {
	if raw := headers.Get("X-Usage-Count"); raw != "" {
		if n, err := strconv.ParseInt(raw, 10, 64); err == nil && n >= 0 && n <= math.MaxInt32 {
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
