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

	count, err := uq.increment(context.Background(), key, userID, cycle, 1)
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

	count, err := uq.increment(context.Background(), key, userID, cycle, 1)
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

func TestUsageQuota_RejectedBatchReservationsAreRolledBack(t *testing.T) {
	pool := setupUsageTestDB(t)
	rdb := setupAPIKeyAuthTestRedis(t)

	userID := time.Now().UnixNano()
	planID := "usage-batch-plan-" + randomHex(t, 4)
	insertPlan(t, pool, planID, ptr(3), nil)

	uq := NewUsageQuota(pool, rdb, NewPlanCache(pool), nil)
	batchHandler := http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("X-Usage-Count", "3")
		w.WriteHeader(http.StatusOK)
	})
	handler := uq.Middleware()(batchHandler)
	principal := APIKeyPrincipal{UserID: userID, Plan: planID, APIKeyID: userID, CurrentPeriodStart: time.Now()}
	t.Cleanup(func() { _, _ = pool.Exec(context.Background(), `DELETE FROM usage_logs WHERE user_id = $1`, userID) })

	first := httptest.NewRequest(http.MethodPost, "/v1/text/words/batch", http.NoBody)
	first = first.WithContext(context.WithValue(first.Context(), apiKeyContextKey{}, principal))
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, first)
	require.Equal(t, http.StatusOK, w.Code)

	cycle := cycleStart(principal.CurrentPeriodStart, time.Now())
	usageKey := fmt.Sprintf("usage:%d:%d", userID, cycle.Unix())
	require.EqualValues(t, "3", rdb.Get(context.Background(), usageKey).Val())

	for range 3 {
		req := httptest.NewRequest(http.MethodPost, "/v1/text/words/batch", http.NoBody)
		req.Header.Set("X-Usage-Count", "999") // request headers are untrusted
		req = req.WithContext(context.WithValue(req.Context(), apiKeyContextKey{}, principal))
		w = httptest.NewRecorder()
		handler.ServeHTTP(w, req)
		require.Equal(t, http.StatusTooManyRequests, w.Code)
		require.EqualValues(t, "3", rdb.Get(context.Background(), usageKey).Val(), "a rejected batch must not consume another reservation")
	}
}

func TestUsageQuota_RedisDownFailsClosed(t *testing.T) {
	pool := setupUsageTestDB(t)

	userID := time.Now().UnixNano()
	insertUsageLog(t, pool, userID, userID, 5, time.Now())

	uq := NewUsageQuota(pool, unreachableRedis(t), NewPlanCache(pool), nil)

	now := time.Now()
	cycle := cycleStart(now, now)
	key := fmt.Sprintf("usage:%d:%d", userID, cycle.Unix())

	_, err := uq.increment(context.Background(), key, userID, cycle, 1)
	require.Error(t, err)
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

	_, err = uq.increment(ctx, key, userID, cycle, 1)
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

	_, err = uq.increment(context.Background(), key, userID, cycle, 1)
	// An OOM write rejection is a Redis outage for quota reservation and must
	// fail closed rather than race a non-atomic Postgres SUM fallback.
	require.Error(t, err)
}

func TestUsageQuota_IncrementUsesCredits(t *testing.T) {
	pool := setupUsageTestDB(t)
	rdb := setupAPIKeyAuthTestRedis(t)

	userID := time.Now().UnixNano()
	uq := NewUsageQuota(pool, rdb, NewPlanCache(pool), nil)
	now := time.Now()
	cycle := cycleStart(now, now)
	key := fmt.Sprintf("usage:%d:%d", userID, cycle.Unix())

	count, err := uq.increment(context.Background(), key, userID, cycle, 3)
	require.NoError(t, err)
	require.EqualValues(t, 3, count)

	count, err = uq.increment(context.Background(), key, userID, cycle, 2)
	require.NoError(t, err)
	require.EqualValues(t, 5, count)
}

func TestCycleStartClampsAnchorDayToMonthEnd(t *testing.T) {
	tests := []struct {
		name      string
		anchorDay int
		now       time.Time
		want      time.Time
	}{
		{"february anchor 29", 29, time.Date(2026, time.February, 20, 12, 0, 0, 0, time.UTC), time.Date(2026, time.January, 29, 0, 0, 0, 0, time.UTC)},
		{"february anchor 30", 30, time.Date(2026, time.February, 20, 12, 0, 0, 0, time.UTC), time.Date(2026, time.January, 30, 0, 0, 0, 0, time.UTC)},
		{"february anchor 31", 31, time.Date(2026, time.February, 20, 12, 0, 0, 0, time.UTC), time.Date(2026, time.January, 31, 0, 0, 0, 0, time.UTC)},
		{"early april anchor 31", 31, time.Date(2026, time.April, 2, 12, 0, 0, 0, time.UTC), time.Date(2026, time.March, 31, 0, 0, 0, 0, time.UTC)},
		{"30-day month anchor 31", 31, time.Date(2026, time.April, 30, 12, 0, 0, 0, time.UTC), time.Date(2026, time.April, 30, 0, 0, 0, 0, time.UTC)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			anchor := time.Date(2025, time.January, tt.anchorDay, 12, 0, 0, 0, time.UTC)
			require.Equal(t, tt.want, cycleStart(anchor, tt.now))
		})
	}
}

func TestRequestCreditsUsesRouteCostNotRequestHeader(t *testing.T) {
	tests := []struct {
		name   string
		method string
		path   string
		want   int64
	}{
		{"dictionary multiplier", http.MethodGet, "/v1/text/dictionary/word", 2},
		{"thesaurus multiplier", http.MethodGet, "/v1/text/thesaurus/word", 2},
		{"ordinary route", http.MethodGet, "/v1/text/advice", 1},
		{"batch route reserves one before response", http.MethodPost, "/v1/text/words/batch", 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, tt.path, http.NoBody)
			req.Header.Set("X-Usage-Count", "999")
			require.Equal(t, tt.want, requestCredits(req))
		})
	}
}

// TestHeaderCredits_BoundsUnvalidatedParseInt guards the CodeQL
// go/incorrect-integer-conversion fix: headerCredits must reject anything
// that wouldn't survive the later int64->int narrowing in recordUsage
// (usage_logs.credits_used is an int column), on top of its existing
// missing/zero/negative handling.
func TestHeaderCredits_BoundsUnvalidatedParseInt(t *testing.T) {
	tests := []struct {
		name string
		raw  string
		want int64
	}{
		{"missing header defaults to 1", "", 1},
		{"non-numeric defaults to 1", "not-a-number", 1},
		{"zero defaults to 1", "0", 1},
		{"negative defaults to 1", "-5", 1},
		{"typical batch count passes through", "50", 50},
		{"exactly MaxInt32 passes through", "2147483647", 2147483647},
		{"above MaxInt32 defaults to 1", "2147483648", 1},
		{"far above MaxInt32 (int64 max) defaults to 1", "9223372036854775807", 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			headers := http.Header{}
			if tt.raw != "" {
				headers.Set("X-Usage-Count", tt.raw)
			}
			require.Equal(t, tt.want, headerCredits(headers))
		})
	}
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
