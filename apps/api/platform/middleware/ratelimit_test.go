package middleware

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/require"
)

// setupPlansTestDB returns a pgxpool.Pool backed by DATABASE_URL, with a
// self-contained plans table (no dependency on Rails' schema having been
// loaded — mirrors setupAPIKeyAuthTestDB's convention in apikeyauth_test.go).
// Shared by ratelimit_test.go and usage_test.go.
func setupPlansTestDB(t *testing.T) *pgxpool.Pool {
	t.Helper()

	dsn := testDSN()
	if dsn == "" {
		t.Skip("TEST_DATABASE_URL/DATABASE_URL not set; skipping rate limiter/usage integration tests")
	}

	ctx := context.Background()
	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		t.Skipf("unable to create pgx pool: %v", err)
	}
	t.Cleanup(pool.Close)

	if err := pool.Ping(ctx); err != nil {
		t.Skipf("database unavailable for rate limiter/usage tests: %v", err)
	}

	_, err = pool.Exec(ctx, `
		CREATE TABLE IF NOT EXISTS plans (
			id TEXT PRIMARY KEY,
			request_limit BIGINT,
			rate_limit_per_minute BIGINT
		)
	`)
	require.NoError(t, err)

	return pool
}

// insertPlan upserts a plans row. nil limits mean "unlimited" — matches the
// enterprise row's shape in production.
func insertPlan(t *testing.T, pool *pgxpool.Pool, id string, requestLimit, rateLimitPerMinute *int64) {
	t.Helper()

	_, err := pool.Exec(context.Background(), `
		INSERT INTO plans (id, request_limit, rate_limit_per_minute) VALUES ($1, $2, $3)
		ON CONFLICT (id) DO UPDATE SET request_limit = $2, rate_limit_per_minute = $3
	`, id, requestLimit, rateLimitPerMinute)
	require.NoError(t, err)

	t.Cleanup(func() {
		_, _ = pool.Exec(context.Background(), `DELETE FROM plans WHERE id = $1`, id)
	})
}

// okHandler is the terminal handler used by both middleware test files: it
// just confirms the request reached the other side of the middleware.
func okHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
}

// requestWithPrincipal builds a request carrying p as its authenticated
// principal, exactly as APIKeyAuth would have attached it to the context.
func requestWithPrincipal(p APIKeyPrincipal) *http.Request {
	req := httptest.NewRequest(http.MethodGet, "/v1/text/advice", http.NoBody)
	return req.WithContext(context.WithValue(req.Context(), apiKeyContextKey{}, p))
}

func ptr(v int64) *int64 { return &v }

func TestRateLimiter_ConcurrentRequestsDontUndercount(t *testing.T) {
	pool := setupPlansTestDB(t)
	rdb := setupAPIKeyAuthTestRedis(t)

	planID := "rl-plan-" + randomHex(t, 4)
	insertPlan(t, pool, planID, nil, ptr(5))

	rl := NewRateLimiter(rdb, NewPlanCache(pool), nil)
	handler := rl.Middleware()(okHandler())

	principal := APIKeyPrincipal{UserID: 1, Plan: planID, APIKeyID: time.Now().UnixNano()}

	const concurrency = 20
	var wg sync.WaitGroup
	var mu sync.Mutex
	successes := 0

	for range concurrency {
		wg.Add(1)
		go func() {
			defer wg.Done()

			w := httptest.NewRecorder()
			handler.ServeHTTP(w, requestWithPrincipal(principal))

			if w.Code == http.StatusOK {
				mu.Lock()
				successes++
				mu.Unlock()
			}
		}()
	}

	wg.Wait()

	// If the atomic INCR script raced/undercounted, more
	// than 5 of these 20 concurrent requests would have slipped through.
	require.Equal(t, 5, successes)
}

func TestRateLimiter_PreviousMinuteWindowDoesNotLeak(t *testing.T) {
	pool := setupPlansTestDB(t)
	rdb := setupAPIKeyAuthTestRedis(t)

	planID := "rl-plan-" + randomHex(t, 4)
	insertPlan(t, pool, planID, nil, ptr(1))

	apiKeyID := time.Now().UnixNano()
	previousMinuteBucket := time.Now().UTC().Unix()/60 - 1
	staleKey := fmt.Sprintf("ratelimit:%d:%d", apiKeyID, previousMinuteBucket)
	require.NoError(t, rdb.Set(context.Background(), staleKey, 999, time.Minute).Err())
	t.Cleanup(func() { _ = rdb.Del(context.Background(), staleKey).Err() })

	rl := NewRateLimiter(rdb, NewPlanCache(pool), nil)
	handler := rl.Middleware()(okHandler())

	principal := APIKeyPrincipal{UserID: 1, Plan: planID, APIKeyID: apiKeyID}
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, requestWithPrincipal(principal))

	// The stale previous-minute key sits at 999 (way over the limit of 1),
	// but the current minute's key is a distinct Redis key and starts fresh.
	require.Equal(t, http.StatusOK, w.Code)
}

func TestRateLimiter_RedisDownFailsOpen(t *testing.T) {
	pool := setupPlansTestDB(t)

	planID := "rl-plan-" + randomHex(t, 4)
	insertPlan(t, pool, planID, nil, ptr(1))

	rl := NewRateLimiter(unreachableRedis(t), NewPlanCache(pool), nil)
	handler := rl.Middleware()(okHandler())

	principal := APIKeyPrincipal{UserID: 1, Plan: planID, APIKeyID: time.Now().UnixNano()}

	for range 3 {
		w := httptest.NewRecorder()
		handler.ServeHTTP(w, requestWithPrincipal(principal))
		require.Equal(t, http.StatusOK, w.Code)
	}
}

func TestRateLimiter_RedisTimeoutFailsOpenWithoutHanging(t *testing.T) {
	pool := setupPlansTestDB(t)
	rdb := setupAPIKeyAuthTestRedis(t)

	planID := "rl-plan-" + randomHex(t, 4)
	insertPlan(t, pool, planID, nil, ptr(1))

	// CLIENT PAUSE stalls command processing on the whole (shared, local)
	// test Redis instance briefly, without dropping the connection — this
	// simulates "alive but slow," a distinct failure mode from "unreachable"
	// that redisCallTimeout exists specifically to catch. No other test in
	// this package runs concurrently (none call t.Parallel), so this is safe.
	require.NoError(t, rdb.Do(context.Background(), "CLIENT", "PAUSE", "300").Err())

	rl := NewRateLimiter(rdb, NewPlanCache(pool), nil)
	handler := rl.Middleware()(okHandler())

	principal := APIKeyPrincipal{UserID: 1, Plan: planID, APIKeyID: time.Now().UnixNano()}

	start := time.Now()
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, requestWithPrincipal(principal))
	elapsed := time.Since(start)

	require.Equal(t, http.StatusOK, w.Code)
	require.Less(t, elapsed, 250*time.Millisecond, "request should fail open at redisCallTimeout, not wait out the full 300ms pause")
}

func TestRateLimiter_EnterprisePlanNeverLimited(t *testing.T) {
	pool := setupPlansTestDB(t)
	rdb := setupAPIKeyAuthTestRedis(t)

	planID := "rl-enterprise-" + randomHex(t, 4)
	insertPlan(t, pool, planID, nil, nil) // both limits null = unlimited

	rl := NewRateLimiter(rdb, NewPlanCache(pool), nil)
	handler := rl.Middleware()(okHandler())

	principal := APIKeyPrincipal{UserID: 1, Plan: planID, APIKeyID: time.Now().UnixNano()}

	for range 20 {
		w := httptest.NewRecorder()
		handler.ServeHTTP(w, requestWithPrincipal(principal))
		require.Equal(t, http.StatusOK, w.Code)
	}
}
