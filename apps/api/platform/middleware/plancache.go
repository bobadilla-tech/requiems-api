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

// get returns the limits for planName and any Postgres lookup error. An
// unknown plan name is a successful lookup with unlimited limits; a database
// failure is returned separately so callers can preserve stale limits or
// choose their own failure policy.
func (c *PlanCache) get(ctx context.Context, planName string) (planLimits, bool, error) {
	c.mu.Lock()
	entry, ok := c.entries[planName]
	c.mu.Unlock()

	if ok && time.Now().Before(entry.expiresAt) {
		return entry.limits, true, nil
	}

	limits, err := c.fetch(ctx, planName)
	if err != nil {
		if ok {
			return entry.limits, true, err
		}
		return planLimits{}, false, err
	}

	c.mu.Lock()
	c.entries[planName] = planCacheEntry{limits: limits, expiresAt: time.Now().Add(planCacheTTL)}
	c.mu.Unlock()

	return limits, false, nil
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
