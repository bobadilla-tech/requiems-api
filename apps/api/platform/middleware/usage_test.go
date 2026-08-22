package middleware

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/require"
)

// setupUsageTestDB extends setupPlansTestDB (ratelimit_test.go) with a
// self-contained usage_logs table, including the same unique index the real
// schema uses as the ON CONFLICT target in usage.go's insertUsageRow.
func setupUsageTestDB(t *testing.T) *pgxpool.Pool {
	t.Helper()

	pool := setupPlansTestDB(t)
	ctx := context.Background()

	_, err := pool.Exec(ctx, `
		CREATE TABLE IF NOT EXISTS usage_logs (
			id BIGSERIAL PRIMARY KEY,
			user_id BIGINT NOT NULL,
			api_key_id BIGINT NOT NULL,
			endpoint TEXT NOT NULL,
			credits_used INTEGER,
			request_method TEXT,
			status_code INTEGER,
			response_time_ms INTEGER,
			used_at TIMESTAMPTZ NOT NULL,
			usage_date DATE,
			created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
			updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
		)
	`)
	require.NoError(t, err)

	_, err = pool.Exec(ctx, `
		CREATE UNIQUE INDEX IF NOT EXISTS idx_usage_logs_test_dedup
		ON usage_logs (api_key_id, used_at, endpoint)
	`)
	require.NoError(t, err)

	return pool
}

func insertUsageLog(t *testing.T, pool *pgxpool.Pool, userID, apiKeyID int64, creditsUsed int, usedAt time.Time) {
	t.Helper()

	_, err := pool.Exec(context.Background(), `
		INSERT INTO usage_logs (user_id, api_key_id, endpoint, credits_used, request_method, status_code, response_time_ms, used_at, usage_date)
		VALUES ($1, $2, 'seed', $3, 'GET', 200, 10, $4, $5)
	`, userID, apiKeyID, creditsUsed, usedAt, usedAt.Format("2006-01-02"))
	require.NoError(t, err)

	t.Cleanup(func() {
		_, _ = pool.Exec(context.Background(), `DELETE FROM usage_logs WHERE user_id = $1`, userID)
	})
}

func TestUsageQuota_BootstrapSumsAcrossAllUserAPIKeys(t *testing.T) {
	pool := setupUsageTestDB(t)
	rdb := setupAPIKeyAuthTestRedis(t)

	userID := time.Now().UnixNano()
	keyA := userID + 1
	keyB := userID + 2

	insertUsageLog(t, pool, userID, keyA, 3, time.Now())
	insertUsageLog(t, pool, userID, keyB, 4, time.Now())

	uq := NewUsageQuota(pool, rdb, NewPlanCache(pool), nil)

	now := time.Now()
	cycle := cycleStart(now, now)
	key := fmt.Sprintf("usage:%d:%d", userID, cycle.Unix())

	count, err := uq.increment(context.Background(), key, userID, cycle)
	require.NoError(t, err)
	// 3 (keyA) + 4 (keyB) = 7 baseline, summed at the user level per the plan
	// doc's Context section, plus 1 for this request's own increment.
	require.EqualValues(t, 8, count)
}

func TestUsageQuota_NewBillingCycleExcludesPriorCycleUsage(t *testing.T) {
	pool := setupUsageTestDB(t)
	rdb := setupAPIKeyAuthTestRedis(t)

	userID := time.Now().UnixNano()
	now := time.Now().UTC()

	// 40 days ago is outside the current monthly cycle regardless of which
	// day-of-month the anchor falls on.
	insertUsageLog(t, pool, userID, userID, 50, now.AddDate(0, 0, -40))

	uq := NewUsageQuota(pool, rdb, NewPlanCache(pool), nil)

	cycle := cycleStart(now, now) // anchor == now: today is the cycle's own start
	key := fmt.Sprintf("usage:%d:%d", userID, cycle.Unix())

	count, err := uq.increment(context.Background(), key, userID, cycle)
	require.NoError(t, err)
	require.EqualValues(t, 1, count, "the 40-day-old row must not be counted in the new cycle's baseline")
}

func TestUsageQuota_CrossingLimitRejectsAndSkipsRowWrite(t *testing.T) {
	pool := setupUsageTestDB(t)
	rdb := setupAPIKeyAuthTestRedis(t)

	userID := time.Now().UnixNano()
	planID := "usage-plan-" + randomHex(t, 4)
	insertPlan(t, pool, planID, ptr(2), nil)

	uq := NewUsageQuota(pool, rdb, NewPlanCache(pool), nil)
	handler := uq.Middleware()(okHandler())

	principal := APIKeyPrincipal{UserID: userID, Plan: planID, APIKeyID: userID, CurrentPeriodStart: time.Now()}
	t.Cleanup(func() { _, _ = pool.Exec(context.Background(), `DELETE FROM usage_logs WHERE user_id = $1`, userID) })

	var lastCode int
	for range 3 {
		w := httptest.NewRecorder()
		handler.ServeHTTP(w, requestWithPrincipal(principal))
		lastCode = w.Code
	}

	require.Equal(t, http.StatusTooManyRequests, lastCode)

	var rowCount int
	err := pool.QueryRow(context.Background(), `SELECT COUNT(*) FROM usage_logs WHERE user_id = $1`, userID).Scan(&rowCount)
	require.NoError(t, err)
	require.EqualValues(t, 2, rowCount, "only the 2 requests that cleared quota should have written a row; the 3rd (rejected) must not")
}

func TestUsageQuota_RedisDownFallsThroughAndStillEnforces(t *testing.T) {
	pool := setupUsageTestDB(t)

	userID := time.Now().UnixNano()
	insertUsageLog(t, pool, userID, userID, 5, time.Now())

	uq := NewUsageQuota(pool, unreachableRedis(t), NewPlanCache(pool), nil)

	now := time.Now()
	cycle := cycleStart(now, now)
	key := fmt.Sprintf("usage:%d:%d", userID, cycle.Unix())

	count, err := uq.increment(context.Background(), key, userID, cycle)
	require.NoError(t, err)
	require.EqualValues(t, 6, count) // 5 baseline + 1, enforced via direct Postgres SUM since Redis is down
}

func TestUsageQuota_BothRedisAndPostgresDownFailsClosed(t *testing.T) {
	ctx := context.Background()
	badPool, err := pgxpool.New(ctx, "postgres://invalid-host-does-not-exist/db?sslmode=disable&connect_timeout=1")
	require.NoError(t, err)
	t.Cleanup(badPool.Close)

	uq := NewUsageQuota(badPool, unreachableRedis(t), NewPlanCache(badPool), nil)

	userID := time.Now().UnixNano()
	now := time.Now()
	cycle := cycleStart(now, now)
	key := fmt.Sprintf("usage:%d:%d", userID, cycle.Unix())

	_, err = uq.increment(ctx, key, userID, cycle)
	require.Error(t, err)
}

func TestUsageQuota_BothDownFailsClosedAtMiddlewareLevel(t *testing.T) {
	ctx := context.Background()
	badPool, err := pgxpool.New(ctx, "postgres://invalid-host-does-not-exist/db?sslmode=disable&connect_timeout=1")
	require.NoError(t, err)
	t.Cleanup(badPool.Close)

	uq := NewUsageQuota(badPool, unreachableRedis(t), NewPlanCache(badPool), nil)
	handler := uq.Middleware()(okHandler())

	principal := APIKeyPrincipal{UserID: time.Now().UnixNano(), Plan: "free", APIKeyID: 1, CurrentPeriodStart: time.Now()}
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, requestWithPrincipal(principal))

	require.Equal(t, http.StatusServiceUnavailable, w.Code)
}

func TestUsageQuota_RedisOOMRejectionHandledLikeConnectionError(t *testing.T) {
	pool := setupUsageTestDB(t)
	rdb := setupAPIKeyAuthTestRedis(t)

	// Snapshot current config so this shared test Redis instance is restored
	// after the test — we're about to make every write on it fail.
	origMaxMemory, err := rdb.ConfigGet(context.Background(), "maxmemory").Result()
	require.NoError(t, err)
	origPolicy, err := rdb.ConfigGet(context.Background(), "maxmemory-policy").Result()
	require.NoError(t, err)

	t.Cleanup(func() {
		if v, ok := origMaxMemory["maxmemory"]; ok {
			_ = rdb.ConfigSet(context.Background(), "maxmemory", v).Err()
		}
		if v, ok := origPolicy["maxmemory-policy"]; ok {
			_ = rdb.ConfigSet(context.Background(), "maxmemory-policy", v).Err()
		}
	})

	require.NoError(t, rdb.ConfigSet(context.Background(), "maxmemory-policy", "noeviction").Err())
	require.NoError(t, rdb.ConfigSet(context.Background(), "maxmemory", "1").Err()) // 1 byte: guarantees the next write is rejected

	userID := time.Now().UnixNano()
	insertUsageLog(t, pool, userID, userID, 2, time.Now())

	uq := NewUsageQuota(pool, rdb, NewPlanCache(pool), nil)
	now := time.Now()
	cycle := cycleStart(now, now)
	key := fmt.Sprintf("usage:%d:%d", userID, cycle.Unix())

	count, err := uq.increment(context.Background(), key, userID, cycle)
	// An OOM write rejection must be handled exactly like a connection
	// error or timeout — falling back to Postgres, not surfacing as an
	// unhandled error or 500.
	require.NoError(t, err)
	require.EqualValues(t, 3, count)
}

func TestUsageQuota_RowLevelWriteDedupsOnConflict(t *testing.T) {
	pool := setupUsageTestDB(t)
	rdb := setupAPIKeyAuthTestRedis(t)

	uq := NewUsageQuota(pool, rdb, NewPlanCache(pool), nil)

	userID := time.Now().UnixNano()
	principal := APIKeyPrincipal{UserID: userID, APIKeyID: userID}
	usedAt := time.Now().UTC()
	t.Cleanup(func() { _, _ = pool.Exec(context.Background(), `DELETE FROM usage_logs WHERE user_id = $1`, userID) })

	err1 := uq.insertUsageRow(context.Background(), principal, "/v1/text/advice", http.MethodGet, 1, 200, 5, usedAt)
	err2 := uq.insertUsageRow(context.Background(), principal, "/v1/text/advice", http.MethodGet, 1, 200, 7, usedAt)

	require.NoError(t, err1)
	require.NoError(t, err2, "a colliding (api_key_id, used_at, endpoint) must be silently absorbed, not surfaced as an error")

	var rowCount int
	err := pool.QueryRow(context.Background(), `SELECT COUNT(*) FROM usage_logs WHERE user_id = $1`, userID).Scan(&rowCount)
	require.NoError(t, err)
	require.EqualValues(t, 1, rowCount)
}

func TestUsageQuota_EnterprisePlanNeverQuotaRejected(t *testing.T) {
	pool := setupUsageTestDB(t)
	rdb := setupAPIKeyAuthTestRedis(t)

	userID := time.Now().UnixNano()
	planID := "usage-enterprise-" + randomHex(t, 4)
	insertPlan(t, pool, planID, nil, nil) // both limits null = unlimited

	uq := NewUsageQuota(pool, rdb, NewPlanCache(pool), nil)
	handler := uq.Middleware()(okHandler())

	principal := APIKeyPrincipal{UserID: userID, Plan: planID, APIKeyID: userID, CurrentPeriodStart: time.Now()}
	t.Cleanup(func() { _, _ = pool.Exec(context.Background(), `DELETE FROM usage_logs WHERE user_id = $1`, userID) })

	for range 10 {
		w := httptest.NewRecorder()
		handler.ServeHTTP(w, requestWithPrincipal(principal))
		require.Equal(t, http.StatusOK, w.Code)
	}
}
