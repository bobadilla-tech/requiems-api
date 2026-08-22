package middleware

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// planCacheTTL matches the concrete TTL value the plan doc specifies for
// this cache (same convention as the geocode/crypto/exchange Redis caches in
// apps/api/services: a fixed, explicitly-chosen TTL rather than "cache
// forever" or "no cache"). This one is in-process, not Redis-backed, since
// plan rows change essentially never and don't need to survive a restart or
// be shared across replicas.
const planCacheTTL = 60 * time.Second

// planLimits is the subset of the plans table both the rate limiter and the
// usage/quota middleware need. A nil field means "unlimited" (enterprise, or
// any future plan with a null column) and the corresponding check is skipped
// entirely.
type planLimits struct {
	RequestLimit       *int64
	RateLimitPerMinute *int64
}

type planCacheEntry struct {
	limits    planLimits
	expiresAt time.Time
}

// PlanCache is a small in-process, TTL-expiring cache over the plans table,
// shared by RateLimiter and UsageQuota so both middlewares read the same
// data through a single Postgres-backed cache rather than each rolling its
// own. Construct one and pass it to both.
type PlanCache struct {
	pool *pgxpool.Pool

	mu      sync.Mutex
	entries map[string]planCacheEntry
}

func NewPlanCache(pool *pgxpool.Pool) *PlanCache {
	return &PlanCache{pool: pool, entries: make(map[string]planCacheEntry)}
}

// get returns the limits for planName, querying Postgres on a cache miss or
// expired entry. An unknown plan name (shouldn't happen in practice, since
// APIKeyPrincipal.Plan always comes from subscriptions.plan_name or the
// "free" default) is treated as fully unlimited rather than erroring the
// request over a data-consistency issue in a different table.
func (c *PlanCache) get(ctx context.Context, planName string) planLimits {
	c.mu.Lock()
	entry, ok := c.entries[planName]
	c.mu.Unlock()

	if ok && time.Now().Before(entry.expiresAt) {
		return entry.limits
	}

	limits, err := c.fetch(ctx, planName)
	if err != nil {
		// A Postgres error here degrades to "unlimited" rather than failing
		// the request: this cache backs soft abuse-prevention (rate
		// limiting) and quota accounting, and a transient plans-table read
		// failure shouldn't itself become the reason a request is rejected.
		// A stale cached entry, if one exists, is preferred over this
		// fallback, but a bootstrap-time failure has no stale entry to fall
		// back to.
		if ok {
			return entry.limits
		}
		return planLimits{}
	}

	c.mu.Lock()
	c.entries[planName] = planCacheEntry{limits: limits, expiresAt: time.Now().Add(planCacheTTL)}
	c.mu.Unlock()

	return limits
}

func (c *PlanCache) fetch(ctx context.Context, planName string) (planLimits, error) {
	var limits planLimits

	err := c.pool.QueryRow(ctx, `
		SELECT request_limit, rate_limit_per_minute FROM plans WHERE id = $1
	`, planName).Scan(&limits.RequestLimit, &limits.RateLimitPerMinute)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return planLimits{}, nil
		}
		return planLimits{}, err
	}

	return limits, nil
}
