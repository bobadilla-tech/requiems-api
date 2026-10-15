package app

import (
	"context"
	"crypto/rand"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"

	"requiems-api/platform/config"
)

// seedAPIKeyFixture creates a self-contained api_keys row (mirroring the
// convention in platform/middleware/apikeyauth_test.go: Go's own test
// Postgres never has Rails' schema loaded, only Go's own migrations run) and
// returns the raw key to present in the requiems-api-key header. It also
// creates subscriptions/plans/usage_logs — empty otherwise, but the protected
// route group now runs APIKeyAuth's LEFT JOIN subscriptions query plus the
// rate limiter and usage/quota middleware from
// docs/plans/2026-08-21-go-auth-foundation-phase-2.md, both of which read
// plans and (on a served request) write usage_logs. A "free" plans row with
// null limits keeps this test focused on auth header wiring rather than
// rate-limit/quota specifics, which have their own dedicated tests in
// platform/middleware.
func seedAPIKeyFixture(t *testing.T, dsn string) string {
	t.Helper()
	return seedAPIKeyFixtureWithKey(t, dsn, "requiem_apphandlertestfixturekey0001")
}

// seedAPIKeyFixtureWithKey is seedAPIKeyFixture parameterized on the raw key.
// Each test that authenticates needs its own distinct key/prefix: the
// APIKeyAuth candidate cache is keyed by prefix and, once warm, verifies
// bcrypt against whatever hash it has cached rather than re-querying
// Postgres — reusing the same literal key across tests whose fixtures
// delete-and-reinsert the row (getting a new api_keys.id each time) lets a
// later test authenticate against an earlier test's stale, already-deleted
// api_key_id, silently orphaning that later test's usage_logs writes.
func seedAPIKeyFixtureWithKey(t *testing.T, dsn, fullKey string) string {
	t.Helper()

	ctx := context.Background()
	pool, err := pgxpool.New(ctx, dsn)
	require.NoError(t, err)
	t.Cleanup(pool.Close)

	_, err = pool.Exec(ctx, `
		CREATE TABLE IF NOT EXISTS api_keys (
			id BIGSERIAL PRIMARY KEY,
			key_prefix TEXT NOT NULL,
			key_hash TEXT NOT NULL,
			user_id BIGINT NOT NULL,
			active BOOLEAN NOT NULL DEFAULT TRUE,
			revoked_at TIMESTAMPTZ,
			created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
			updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
		)
	`)
	require.NoError(t, err)

	_, err = pool.Exec(ctx, `
		CREATE TABLE IF NOT EXISTS subscriptions (
			id BIGSERIAL PRIMARY KEY,
			user_id BIGINT NOT NULL,
			plan_name TEXT,
			status TEXT,
			current_period_start TIMESTAMPTZ NOT NULL DEFAULT NOW(),
			created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
			updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
		)
	`)
	require.NoError(t, err)
	_, err = pool.Exec(ctx, `ALTER TABLE subscriptions ADD COLUMN IF NOT EXISTS status TEXT`)
	require.NoError(t, err)

	_, err = pool.Exec(ctx, `
		CREATE TABLE IF NOT EXISTS plans (
			id TEXT PRIMARY KEY,
			request_limit BIGINT,
			rate_limit_per_minute BIGINT
		)
	`)
	require.NoError(t, err)

	_, err = pool.Exec(ctx, `
		INSERT INTO plans (id, request_limit, rate_limit_per_minute) VALUES ('free', NULL, NULL)
		ON CONFLICT (id) DO UPDATE SET request_limit = NULL, rate_limit_per_minute = NULL
	`)
	require.NoError(t, err)

	_, err = pool.Exec(ctx, `
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
		CREATE UNIQUE INDEX IF NOT EXISTS idx_usage_logs_apphandler_dedup
		ON usage_logs (api_key_id, used_at, endpoint)
	`)
	require.NoError(t, err)

	prefix := fullKey[:12]

	hash, err := bcrypt.GenerateFromPassword([]byte(fullKey), bcrypt.MinCost)
	require.NoError(t, err)

	_, err = pool.Exec(ctx, `DELETE FROM api_keys WHERE key_prefix = $1`, prefix)
	require.NoError(t, err)

	_, err = pool.Exec(ctx, `
		INSERT INTO api_keys (key_prefix, key_hash, user_id, active)
		VALUES ($1, $2, $3, TRUE)
	`, prefix, string(hash), time.Now().UnixNano())
	require.NoError(t, err)

	t.Cleanup(func() {
		_, _ = pool.Exec(context.Background(), `DELETE FROM api_keys WHERE key_prefix = $1`, prefix)
	})

	return fullKey
}

// TestNew_ErrorOnBadDatabaseURL verifies that New returns an error immediately
// when the DATABASE_URL cannot be used to connect (no retries for malformed URLs).
func TestNew_ErrorOnBadDatabaseURL(t *testing.T) {
	t.Chdir("..") // resolve "migrations" relative to api root, not package dir

	_, err := New(context.Background(), config.Config{
		DatabaseURL: "postgres://invalid-host-that-does-not-exist/db?sslmode=disable&connect_timeout=1",
		RedisURL:    "redis://localhost:6379/0",
	})

	require.Error(t, err)
}

// TestApp_Close verifies Close releases the pool and Redis client without
// panicking, and that both are actually unusable afterward (not just that
// Close returned).
func TestApp_Close(t *testing.T) {
	dsn := os.Getenv("TEST_DATABASE_URL")
	if dsn == "" {
		dsn = os.Getenv("DATABASE_URL")
	}

	if dsn == "" {
		t.Skip("TEST_DATABASE_URL/DATABASE_URL not set; skipping App integration test")
	}

	redisURL := os.Getenv("REDIS_URL")
	if redisURL == "" {
		redisURL = "redis://localhost:6379/0"
	}

	t.Chdir("..") // resolve "migrations" relative to api root, not package dir

	app, err := New(context.Background(), config.Config{
		DatabaseURL: dsn,
		RedisURL:    redisURL,
	})
	if err != nil {
		t.Skipf("infrastructure unavailable; skipping App integration test: %v", err)
	}

	app.Close()

	assert.Error(t, app.pool.Ping(context.Background()))
	assert.Error(t, app.rdb.Ping(context.Background()).Err())
}

// TestApp_Handler is an integration test that creates a real App and verifies
// the HTTP handler has the expected routing structure:
//   - GET /healthz is publicly accessible (no auth required)
//   - GET /v1/* routes require a valid requiems-api-key header — APIKeyAuth
//     is the sole enforcing gate now that BackendSecretAuth is retired (see
//     docs/plans/2026-08-21-go-auth-foundation-phase-3-4.md Phase 3 item 5)
//
// The test is skipped when DATABASE_URL is not set.
func TestApp_Handler(t *testing.T) {
	// TEST_DATABASE_URL (a dedicated database, see docker-compose.dev.yml's
	// db-init service) takes priority so seedAPIKeyFixture's self-contained
	// tables never collide with Rails' real migrations of the same names
	// when both suites run against the same Postgres server (agents.md).
	dsn := os.Getenv("TEST_DATABASE_URL")
	if dsn == "" {
		dsn = os.Getenv("DATABASE_URL")
	}

	if dsn == "" {
		t.Skip("TEST_DATABASE_URL/DATABASE_URL not set; skipping App integration test")
	}

	redisURL := os.Getenv("REDIS_URL")

	if redisURL == "" {
		redisURL = "redis://localhost:6379/0"
	}

	t.Chdir("..") // resolve "migrations" relative to api root, not package dir

	cfg := config.Config{
		DatabaseURL: dsn,
		RedisURL:    redisURL,
	}

	app, err := New(context.Background(), cfg)

	if err != nil {
		t.Skipf("infrastructure unavailable; skipping App integration test: %v", err)
	}

	h := app.Handler()
	require.NotNil(t, h)

	t.Run("healthz is publicly accessible", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/healthz", http.NoBody)
		w := httptest.NewRecorder()
		h.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("v1 routes require an api key", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/v1/entertainment/advice", http.NoBody)
		w := httptest.NewRecorder()
		h.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})

	t.Run("v1 routes are accessible with a valid api key and no X-Backend-Secret header at all", func(t *testing.T) {
		// This is the traffic shape Cloudflare-proxied requests carry once
		// the Worker is retired (and the shape any direct client already
		// carries): no X-Backend-Secret, just the api key. Proves item 5
		// actually fixed the 401-on-all-traffic state described in the
		// plan's Context section, not just changed why it failed.
		apiKey := seedAPIKeyFixture(t, dsn)

		req := httptest.NewRequest(http.MethodGet, "/v1/entertainment/advice", http.NoBody)
		req.Header.Set("requiems-api-key", apiKey)
		w := httptest.NewRecorder()
		h.ServeHTTP(w, req)

		// The endpoint itself may return any 2xx, but a non-2xx response means
		// this regression test did not prove the authenticated request served.
		assert.GreaterOrEqual(t, w.Code, http.StatusOK)
		assert.Less(t, w.Code, http.StatusMultipleChoices)
	})
}

// creditsUsedFor queries the credits_used column usage/quota middleware just
// wrote for the given endpoint path, ordered to the most recent row (each
// subtest below uses a distinct word so paths don't collide, but ORDER BY id
// DESC keeps this robust even if that ever changes).
func creditsUsedFor(t *testing.T, pool *pgxpool.Pool, apiKeyPrefix, endpoint string) int {
	t.Helper()

	var credits int

	err := pool.QueryRow(context.Background(), `
		SELECT ul.credits_used FROM usage_logs ul
		JOIN api_keys ak ON ak.id = ul.api_key_id
		WHERE ak.key_prefix = $1 AND ul.endpoint = $2
		ORDER BY ul.id DESC LIMIT 1
	`, apiKeyPrefix, endpoint).Scan(&credits)
	require.NoError(t, err)

	return credits
}

// TestApp_UsageMultiplier_DictionaryAndThesaurus is an integration test for
// docs/plans/2026-08-22-go-auth-foundation-phase-6-usage-multiplier-and-loose-ends.md
// Phase 6a: GET /v1/text/dictionary/{word} and GET /v1/text/thesaurus/{word}
// must bill 2 credits (X-Usage-Count response header and the usage_logs
// credits_used column), through the full APIKeyAuth -> rate limit -> usage
// quota -> handler chain. The batch routes must be left unaffected, still
// billing per item via httpx.HandleBatch.
func TestApp_UsageMultiplier_DictionaryAndThesaurus(t *testing.T) {
	dsn := os.Getenv("TEST_DATABASE_URL")
	if dsn == "" {
		dsn = os.Getenv("DATABASE_URL")
	}

	if dsn == "" {
		t.Skip("TEST_DATABASE_URL/DATABASE_URL not set; skipping usage-multiplier integration test")
	}

	redisURL := os.Getenv("REDIS_URL")
	if redisURL == "" {
		redisURL = "redis://localhost:6379/0"
	}

	t.Chdir("..")

	cfg := config.Config{DatabaseURL: dsn, RedisURL: redisURL}

	app, err := New(context.Background(), cfg)
	if err != nil {
		t.Skipf("infrastructure unavailable; skipping usage-multiplier integration test: %v", err)
	}

	h := app.Handler()

	pool, err := pgxpool.New(context.Background(), dsn)
	require.NoError(t, err)
	t.Cleanup(pool.Close)

	// A unique key per run, not a hardcoded literal: the APIKeyAuth Redis
	// cache is keyed by the 12-char prefix and survives across separate
	// `go test` process invocations (it's a real shared Redis, TTL-bound,
	// not per-process). A fixed literal risks a cache hit against a stale
	// api_key_id left over from an earlier run's now-deleted row, silently
	// orphaning every usage_logs write this test then tries to read back.
	randSuffix := make([]byte, 4)
	_, err = rand.Read(randSuffix)
	require.NoError(t, err)

	apiKey := seedAPIKeyFixtureWithKey(t, dsn, fmt.Sprintf("requiem_%x00000000000000000000", randSuffix))
	prefix := apiKey[:12]

	t.Run("GET dictionary lookup bills 2 credits", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/v1/text/dictionary/ephemeral", http.NoBody)
		req.Header.Set("requiems-api-key", apiKey)
		w := httptest.NewRecorder()
		h.ServeHTTP(w, req)

		require.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "2", w.Header().Get("X-Usage-Count"))
		assert.Equal(t, 2, creditsUsedFor(t, pool, prefix, "/v1/text/dictionary/ephemeral"))
	})

	t.Run("GET thesaurus lookup bills 2 credits", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/v1/text/thesaurus/happy", http.NoBody)
		req.Header.Set("requiems-api-key", apiKey)
		w := httptest.NewRecorder()
		h.ServeHTTP(w, req)

		require.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "2", w.Header().Get("X-Usage-Count"))
		assert.Equal(t, 2, creditsUsedFor(t, pool, prefix, "/v1/text/thesaurus/happy"))
	})

	t.Run("POST words batch still bills per item, not flattened by the GET-route fix", func(t *testing.T) {
		body := `{"items": ["ephemeral", "serendipity", "melancholy"]}`
		req := httptest.NewRequest(http.MethodPost, "/v1/text/words/batch", strings.NewReader(body))
		req.Header.Set("requiems-api-key", apiKey)
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		h.ServeHTTP(w, req)

		require.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "3", w.Header().Get("X-Usage-Count"))
		assert.Equal(t, 3, creditsUsedFor(t, pool, prefix, "/v1/text/words/batch"))
	})

	t.Run("POST thesaurus batch still bills per item, not flattened by the GET-route fix", func(t *testing.T) {
		body := `{"words": ["happy", "sad"]}`
		req := httptest.NewRequest(http.MethodPost, "/v1/text/thesaurus/batch", strings.NewReader(body))
		req.Header.Set("requiems-api-key", apiKey)
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		h.ServeHTTP(w, req)

		require.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "2", w.Header().Get("X-Usage-Count"))
		assert.Equal(t, 2, creditsUsedFor(t, pool, prefix, "/v1/text/thesaurus/batch"))
	})
}

// usageCounterSum totals the Redis quota counters (usage:<user>:<cycle>) of a
// user, which is what enforces the monthly quota.
func usageCounterSum(t *testing.T, rdb *redis.Client, userID int64) int64 {
	t.Helper()

	keys, err := rdb.Keys(context.Background(), fmt.Sprintf("usage:%d:*", userID)).Result()
	require.NoError(t, err)

	var sum int64
	for _, k := range keys {
		n, err := rdb.Get(context.Background(), k).Int64()
		require.NoError(t, err)
		sum += n
	}

	return sum
}

// TestApp_RouteMissesAndErrorsAreNotCharged is the end-to-end proof for the
// usage plan's PR-B (docs/plans/2026-10-15-dashboard-overview-usage-fixes.md):
// requests for routes that do not exist are not usage at all (no ledger row,
// no quota movement), and handler-level 4xx/5xx are logged but never charged.
func TestApp_RouteMissesAndErrorsAreNotCharged(t *testing.T) {
	dsn := os.Getenv("TEST_DATABASE_URL")
	if dsn == "" {
		dsn = os.Getenv("DATABASE_URL")
	}

	if dsn == "" {
		t.Skip("TEST_DATABASE_URL/DATABASE_URL not set; skipping route-miss integration test")
	}

	redisURL := os.Getenv("REDIS_URL")
	if redisURL == "" {
		redisURL = "redis://localhost:6379/0"
	}

	t.Chdir("..")

	app, err := New(context.Background(), config.Config{DatabaseURL: dsn, RedisURL: redisURL})
	if err != nil {
		t.Skipf("infrastructure unavailable; skipping route-miss integration test: %v", err)
	}

	t.Cleanup(app.Close)

	h := app.Handler()

	pool, err := pgxpool.New(context.Background(), dsn)
	require.NoError(t, err)
	t.Cleanup(pool.Close)

	randSuffix := make([]byte, 4)
	_, err = rand.Read(randSuffix)
	require.NoError(t, err)

	apiKey := seedAPIKeyFixtureWithKey(t, dsn, fmt.Sprintf("requiem_%x00000000000000000000", randSuffix))
	prefix := apiKey[:12]

	var userID int64
	require.NoError(t, pool.QueryRow(context.Background(), `SELECT user_id FROM api_keys WHERE key_prefix = $1`, prefix).Scan(&userID))
	t.Cleanup(func() { _, _ = pool.Exec(context.Background(), `DELETE FROM usage_logs WHERE user_id = $1`, userID) })

	do := func(method, target, body string) *httptest.ResponseRecorder {
		req := httptest.NewRequest(method, target, strings.NewReader(body))
		req.Header.Set("requiems-api-key", apiKey)
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		h.ServeHTTP(w, req)

		return w
	}

	rows := func() int {
		var n int
		require.NoError(t, pool.QueryRow(context.Background(), `SELECT COUNT(*) FROM usage_logs WHERE user_id = $1`, userID).Scan(&n))
		return n
	}

	// A valid request first: one ledger row, one credit on the counter.
	require.Equal(t, http.StatusOK, do(http.MethodGet, "/v1/entertainment/advice", "").Code)
	require.Equal(t, 1, rows(), "a valid route is recorded")
	require.EqualValues(t, 1, usageCounterSum(t, app.rdb, userID))

	misses := []struct {
		name, method, target string
		want                 int
	}{
		{"unknown top-level route", http.MethodGet, "/v1/nonexistent", http.StatusNotFound},
		{"unknown route inside a service", http.MethodGet, "/v1/text/nonexistent", http.StatusNotFound},
		{"route that never existed in the text service", http.MethodGet, "/v1/text/advice", http.StatusNotFound},
		{"wrong method", http.MethodPost, "/v1/entertainment/advice", http.StatusMethodNotAllowed},
		{"percent-encoded path", http.MethodGet, "/v1/text/non%2Fexistent", http.StatusNotFound},
		{"percent-encoded unreserved path", http.MethodGet, "/v1/text/%6Eonexistent", http.StatusNotFound},
		{"deep unknown path", http.MethodGet, "/v1/text/a/b/c/d", http.StatusNotFound},
	}

	for _, tt := range misses {
		t.Run(tt.name, func(t *testing.T) {
			beforeRows, beforeCounter := rows(), usageCounterSum(t, app.rdb, userID)

			w := do(tt.method, tt.target, "")

			assert.Equal(t, tt.want, w.Code)
			assert.Equal(t, beforeRows, rows(), "a route miss writes no ledger row")
			assert.Equal(t, beforeCounter, usageCounterSum(t, app.rdb, userID), "and moves no quota")
		})
	}

	t.Run("a valid route is still recorded after the misses", func(t *testing.T) {
		beforeRows, beforeCounter := rows(), usageCounterSum(t, app.rdb, userID)

		require.Equal(t, http.StatusOK, do(http.MethodGet, "/v1/entertainment/advice", "").Code)

		assert.Equal(t, beforeRows+1, rows())
		assert.Equal(t, beforeCounter+1, usageCounterSum(t, app.rdb, userID))
	})

	t.Run("a handler-level validation error is logged but free", func(t *testing.T) {
		beforeRows, beforeCounter := rows(), usageCounterSum(t, app.rdb, userID)

		w := do(http.MethodPost, "/v1/text/words/batch", `{}`)

		require.GreaterOrEqual(t, w.Code, http.StatusBadRequest)
		require.Less(t, w.Code, http.StatusInternalServerError)
		assert.Equal(t, beforeRows+1, rows(), "the error is still in the ledger")
		assert.Equal(t, 0, creditsUsedFor(t, pool, prefix, "/v1/text/words/batch"))
		assert.Equal(t, beforeCounter, usageCounterSum(t, app.rdb, userID), "and refunded from the counter")
	})

	t.Run("the dictionary route still bills 2 on success and nothing on a miss", func(t *testing.T) {
		beforeCounter := usageCounterSum(t, app.rdb, userID)

		w := do(http.MethodGet, "/v1/text/dictionary/ephemeral", "")
		require.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "2", w.Header().Get("X-Usage-Count"))
		assert.Equal(t, beforeCounter+2, usageCounterSum(t, app.rdb, userID))
	})
}
