package middleware

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"

	"requiems-api/platform/httpx"
)

// redisCallTimeout bounds every Redis call this file and usage.go make. A
// Redis instance that's alive but slow (under memory pressure, for example)
// is a distinct failure mode from a clean connection error, and without an
// explicit timeout it isn't caught by the same error-handling path at all —
// it just makes every request hang instead of falling through/failing open.
// Tens of milliseconds, per the plan doc's explicit guidance.
const redisCallTimeout = 50 * time.Millisecond

// rateLimitScript is a single atomic round trip: increment the per-minute
// bucket, and set it to expire in 60s the first time it's created (cleanup
// only — the minute-bucketed key name is what actually prevents a previous
// window's count from leaking into the next one, not this TTL).
var rateLimitScript = redis.NewScript(`
local count = redis.call("INCR", KEYS[1])
if count == 1 then redis.call("EXPIRE", KEYS[1], 60) end
return count
`)

// RateLimiter enforces a per-API-key, per-minute request cap read from the
// plans table. It is soft abuse-prevention, not a security boundary (per the
// architecture audit's Reliability Findings): any Redis error, including a
// call that exceeds redisCallTimeout, fails open (the request proceeds,
// logged at warn) rather than turning a Redis blip into an outage.
//
// Known, unmitigated gap: this limiter is keyed by api_key_id, which only
// exists after APIKeyAuth has already resolved an identity from a warm cache
// entry trusting a presented key_prefix alone. It does nothing to bound the
// brute-force exposure on that cache-hit trust boundary (see
// apikeyauth.go/the Phase 0-1 plan's Final notes) — a single correct-prefix
// guess succeeds before any per-key limit ever applies. Not addressed here;
// flagged as a follow-up in the Phase 2 plan doc.
type RateLimiter struct {
	rdb    *redis.Client
	plans  *PlanCache
	logger *slog.Logger
}

// NewRateLimiter builds a RateLimiter. plans should be shared with UsageQuota
// (both read the same plans table through the same in-process cache).
func NewRateLimiter(rdb *redis.Client, plans *PlanCache, logger *slog.Logger) *RateLimiter {
	if logger == nil {
		logger = slog.Default()
	}

	return &RateLimiter{rdb: rdb, plans: plans, logger: logger}
}

// Middleware returns the enforcing http middleware. Mount after APIKeyAuth
// in the same protected group — it reads APIKeyPrincipalFromContext and has
// nothing to key on without it.
func (rl *RateLimiter) Middleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			principal, ok := APIKeyPrincipalFromContext(r.Context())
			if !ok {
				// No authenticated principal to key on. In practice this
				// middleware is always mounted after APIKeyAuth, so this
				// path shouldn't be reachable in production; don't let a
				// wiring mistake turn into a panic or a false rate-limit.
				next.ServeHTTP(w, r)
				return
			}

			limits, _, planErr := rl.plans.get(r.Context(), principal.Plan)
			if planErr != nil {
				rl.logger.Warn("rate limiter: plan lookup failed, failing open", "error", planErr)
			}
			if limits.RateLimitPerMinute == nil {
				// Unlimited plan (enterprise, or any plan with a null
				// rate_limit_per_minute) skips the check entirely.
				next.ServeHTTP(w, r)
				return
			}

			now := time.Now().UTC().Unix()
			minuteBucket := now / 60
			key := fmt.Sprintf("ratelimit:%d:%d", principal.APIKeyID, minuteBucket)

			ctx, cancel := context.WithTimeout(r.Context(), redisCallTimeout)
			count, err := rateLimitScript.Run(ctx, rl.rdb, []string{key}).Int64()
			cancel()

			if err != nil {
				rl.logger.Warn("rate limiter: redis error, failing open", "error", err)
				next.ServeHTTP(w, r)
				return
			}

			if count > *limits.RateLimitPerMinute {
				retryAfter := 60 - now%60
				w.Header().Set("Retry-After", strconv.FormatInt(retryAfter, 10))
				httpx.Error(w, http.StatusTooManyRequests, "rate_limited", "Rate limit exceeded")
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
