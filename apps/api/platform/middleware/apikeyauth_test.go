package middleware

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

type discardLogger struct{}

func (discardLogger) Printf(context.Context, string, ...any) {}

func init() {
	// The "Redis down" test cases below deliberately dial an unreachable
	// address; silence go-redis's internal reconnect logging so test output
	// stays readable.
	redis.SetLogger(discardLogger{})
}

// testDSN prefers TEST_DATABASE_URL (a dedicated database, see
// docker-compose.dev.yml's db-init service) so this package's
// self-contained fixture tables (api_keys/subscriptions/plans/usage_logs)
// never collide with Rails' real migrations of the same table names when
// both suites run against the same Postgres server. Falls back to
// DATABASE_URL when TEST_DATABASE_URL isn't set (documented in agents.md).
func testDSN() string {
	if dsn := os.Getenv("TEST_DATABASE_URL"); dsn != "" {
		return dsn
	}
	return os.Getenv("DATABASE_URL")
}

// setupAPIKeyAuthTestDB returns a pgxpool.Pool backed by DATABASE_URL, with
// self-contained api_keys/subscriptions tables (no FK to a users table,
// matching the CI environment where only Go's own migrations have run —
// Rails' schema is never loaded here). Mirrors the existing
// services/finance/bin integration-test convention in this repo.
func setupAPIKeyAuthTestDB(t *testing.T) *pgxpool.Pool {
	t.Helper()

	dsn := testDSN()
	if dsn == "" {
		t.Skip("TEST_DATABASE_URL/DATABASE_URL not set; skipping API key auth integration tests")
	}

	ctx := context.Background()
	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		t.Skipf("unable to create pgx pool: %v", err)
	}
	t.Cleanup(pool.Close)

	if err := pool.Ping(ctx); err != nil {
		t.Skipf("database unavailable for API key auth tests: %v", err)
	}

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

	return pool
}

func setupAPIKeyAuthTestRedis(t *testing.T) *redis.Client {
	t.Helper()

	url := os.Getenv("REDIS_URL")
	if url == "" {
		url = "redis://127.0.0.1:6379/0"
	}

	opts, err := redis.ParseURL(url)
	require.NoError(t, err)

	// Matches platform/reqredis.Connect's production client: required for
	// the rate limiter/usage-quota tests' context.WithTimeout calls to
	// actually bound a slow Redis instead of blocking on go-redis's own
	// (much longer) default ReadTimeout.
	opts.ContextTimeoutEnabled = true

	rdb := redis.NewClient(opts)
	t.Cleanup(func() { _ = rdb.Close() })

	if err := rdb.Ping(context.Background()).Err(); err != nil {
		t.Skipf("redis unavailable for API key auth tests: %v", err)
	}

	return rdb
}

// unreachableRedis returns a client pointed at a port nothing listens on, to
// simulate a Redis outage without touching the real test Redis instance.
func unreachableRedis(t *testing.T) *redis.Client {
	t.Helper()

	rdb := redis.NewClient(&redis.Options{
		Addr:        "127.0.0.1:1",
		DialTimeout: 100 * time.Millisecond,
		ReadTimeout: 100 * time.Millisecond,
		MaxRetries:  -1,
	})
	t.Cleanup(func() { _ = rdb.Close() })

	return rdb
}

func randomHex(t *testing.T, n int) string {
	t.Helper()

	b := make([]byte, n)
	_, err := rand.Read(b)
	require.NoError(t, err)

	return hex.EncodeToString(b)
}

// testKey builds a (prefix, fullKey) pair. prefix is always exactly
// keyPrefixLength characters ("requiem_" + 4 random hex chars), so each call
// gets its own isolated Redis/Postgres key.
func testKey(t *testing.T) (prefix string, fullKey string) {
	t.Helper()

	prefix = "requiem_" + randomHex(t, 2) // "requiem_" (8) + 4 hex chars = 12
	fullKey = prefix + randomHex(t, 10)

	return prefix, fullKey
}

func hashKey(t *testing.T, fullKey string) string {
	t.Helper()

	h, err := bcrypt.GenerateFromPassword([]byte(fullKey), bcrypt.MinCost)
	require.NoError(t, err)

	return string(h)
}

func insertAPIKey(t *testing.T, pool *pgxpool.Pool, prefix, fullKey string, userID int64, active bool, revoked bool) {
	t.Helper()

	var revokedAt any
	if revoked {
		revokedAt = time.Now()
	}

	_, err := pool.Exec(context.Background(), `
		INSERT INTO api_keys (key_prefix, key_hash, user_id, active, revoked_at)
		VALUES ($1, $2, $3, $4, $5)
	`, prefix, hashKey(t, fullKey), userID, active, revokedAt)
	require.NoError(t, err)

	t.Cleanup(func() {
		_, _ = pool.Exec(context.Background(), `DELETE FROM api_keys WHERE key_prefix = $1`, prefix)
	})
}

func insertSubscription(t *testing.T, pool *pgxpool.Pool, userID int64, planName string) {
	t.Helper()

	_, err := pool.Exec(context.Background(), `
		INSERT INTO subscriptions (user_id, plan_name, status) VALUES ($1, $2, 'active')
	`, userID, planName)
	require.NoError(t, err)

	t.Cleanup(func() {
		_, _ = pool.Exec(context.Background(), `DELETE FROM subscriptions WHERE user_id = $1`, userID)
	})
}

func newTestHandler() (http.Handler, *int64) {
	var calls int64

	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		calls++

		if _, ok := APIKeyPrincipalFromContext(r.Context()); !ok {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	})

	return next, &calls
}

func doRequest(t *testing.T, h http.Handler, key string) *httptest.ResponseRecorder {
	t.Helper()

	req := httptest.NewRequest(http.MethodGet, "/v1/text/advice", http.NoBody)
	if key != "" {
		req.Header.Set(apiKeyHeader, key)
	}

	w := httptest.NewRecorder()
	h.ServeHTTP(w, req)

	return w
}

func TestAPIKeyAuth_ValidKeyResolvesUserAndPlan(t *testing.T) {
	pool := setupAPIKeyAuthTestDB(t)
	rdb := setupAPIKeyAuthTestRedis(t)

	prefix, fullKey := testKey(t)
	userID := time.Now().UnixNano()
	insertAPIKey(t, pool, prefix, fullKey, userID, true, false)
	insertSubscription(t, pool, userID, "developer")

	auth := NewAPIKeyAuth(pool, rdb, time.Minute)

	var gotPlan string
	var gotPrincipal APIKeyPrincipal
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		p, ok := APIKeyPrincipalFromContext(r.Context())
		require.True(t, ok)
		require.Equal(t, userID, p.UserID)
		gotPlan = p.Plan
		gotPrincipal = p
		w.WriteHeader(http.StatusOK)
	})

	w := doRequest(t, auth.Middleware()(next), fullKey)

	require.Equal(t, http.StatusOK, w.Code)
	require.Equal(t, "developer", gotPlan)
	require.NotZero(t, gotPrincipal.APIKeyID)
	// The subscriptions test table defaults current_period_start to NOW() at
	// insert time (see setupAPIKeyAuthTestDB), so it should round-trip close
	// to "now", not the zero value.
	require.WithinDuration(t, time.Now(), gotPrincipal.CurrentPeriodStart, 10*time.Second)
}

func TestAPIKeyAuth_DefaultsToFreePlanWithoutSubscription(t *testing.T) {
	pool := setupAPIKeyAuthTestDB(t)
	rdb := setupAPIKeyAuthTestRedis(t)

	prefix, fullKey := testKey(t)
	userID := time.Now().UnixNano()
	insertAPIKey(t, pool, prefix, fullKey, userID, true, false)

	auth := NewAPIKeyAuth(pool, rdb, time.Minute)

	var gotPrincipal APIKeyPrincipal
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		p, _ := APIKeyPrincipalFromContext(r.Context())
		gotPrincipal = p
		w.WriteHeader(http.StatusOK)
	})

	w := doRequest(t, auth.Middleware()(next), fullKey)

	require.Equal(t, http.StatusOK, w.Code)
	require.Equal(t, "free", gotPrincipal.Plan)
	// No subscriptions row exists for this user, so CurrentPeriodStart must
	// fall back to the key's own created_at (defaults to NOW() at insert).
	require.WithinDuration(t, time.Now(), gotPrincipal.CurrentPeriodStart, 10*time.Second)
}

func TestAPIKeyAuth_SelectsEligibleSubscriptionDeterministically(t *testing.T) {
	pool := setupAPIKeyAuthTestDB(t)
	rdb := setupAPIKeyAuthTestRedis(t)

	prefix, fullKey := testKey(t)
	userID := time.Now().UnixNano()
	insertAPIKey(t, pool, prefix, fullKey, userID, true, false)

	_, err := pool.Exec(context.Background(), `
		INSERT INTO subscriptions (user_id, plan_name, status, updated_at)
		VALUES ($1, 'developer', 'active', NOW() - INTERVAL '1 day'),
		       ($1, 'business', 'cancelled', NOW())
	`, userID)
	require.NoError(t, err)
	t.Cleanup(func() { _, _ = pool.Exec(context.Background(), `DELETE FROM subscriptions WHERE user_id = $1`, userID) })

	auth := NewAPIKeyAuth(pool, rdb, time.Minute)
	var gotPlan string
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		principal, _ := APIKeyPrincipalFromContext(r.Context())
		gotPlan = principal.Plan
		w.WriteHeader(http.StatusOK)
	})

	w := doRequest(t, auth.Middleware()(next), fullKey)
	require.Equal(t, http.StatusOK, w.Code)
	require.Equal(t, "developer", gotPlan)
}

func TestAPIKeyAuth_RevokedKeyRejected(t *testing.T) {
	pool := setupAPIKeyAuthTestDB(t)
	rdb := setupAPIKeyAuthTestRedis(t)

	prefix, fullKey := testKey(t)
	userID := time.Now().UnixNano()
	insertAPIKey(t, pool, prefix, fullKey, userID, false, true)

	auth := NewAPIKeyAuth(pool, rdb, time.Minute)
	next, calls := newTestHandler()

	t.Run("cache-miss path rejects immediately", func(t *testing.T) {
		w := doRequest(t, auth.Middleware()(next), fullKey)
		require.Equal(t, http.StatusUnauthorized, w.Code)
		require.EqualValues(t, 0, *calls)
	})

	t.Run("still rejects with a warm cache entry", func(t *testing.T) {
		// The first request above already cached the revoked result;
		// re-requesting within TTL must still reject from cache, not just
		// from Postgres.
		w := doRequest(t, auth.Middleware()(next), fullKey)
		require.Equal(t, http.StatusUnauthorized, w.Code)
		require.EqualValues(t, 0, *calls)
	})
}

func TestAPIKeyAuth_RevokeInvalidatesWarmCache(t *testing.T) {
	// This is the test that actually catches a namespace/keyspace mismatch
	// in Rails' revocation DEL (e.g. deleting "rails_cache:{prefix}" while Go
	// reads plain "apikey:{prefix}") — without it, that bug ships silently.
	pool := setupAPIKeyAuthTestDB(t)
	rdb := setupAPIKeyAuthTestRedis(t)

	prefix, fullKey := testKey(t)
	userID := time.Now().UnixNano()
	insertAPIKey(t, pool, prefix, fullKey, userID, true, false)

	auth := NewAPIKeyAuth(pool, rdb, time.Minute)
	next, _ := newTestHandler()
	handler := auth.Middleware()(next)

	// Warm the cache with a valid, active key.
	w := doRequest(t, handler, fullKey)
	require.Equal(t, http.StatusOK, w.Code)

	// Revoke in Postgres...
	_, err := pool.Exec(context.Background(), `UPDATE api_keys SET active = FALSE, revoked_at = NOW() WHERE key_prefix = $1`, prefix)
	require.NoError(t, err)

	// ...and invalidate the same raw, unnamespaced Redis key Go reads/writes.
	// This is what Rails' ApiKey#revoke! must do in production — see
	// apiKeyCacheKeyPrefix in apikeyauth.go for the exact contract.
	require.NoError(t, rdb.Del(context.Background(), apiKeyCacheKeyPrefix+prefix).Err())

	// Re-request before TTL expiry: must be rejected, not served from a
	// stale warm cache entry.
	w = doRequest(t, handler, fullKey)
	require.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestAPIKeyAuth_CollisionCandidateSet(t *testing.T) {
	pool := setupAPIKeyAuthTestDB(t)
	rdb := setupAPIKeyAuthTestRedis(t)

	prefix := "requiem_" + randomHex(t, 2)
	fullKeyA := prefix + "aaaaaaaaaaaaaaaaaaaa"
	fullKeyB := prefix + "bbbbbbbbbbbbbbbbbbbb"

	userA := time.Now().UnixNano()
	userB := userA + 1

	insertAPIKey(t, pool, prefix, fullKeyA, userA, true, false)
	insertAPIKey(t, pool, prefix, fullKeyB, userB, true, false)
	insertSubscription(t, pool, userA, "free")
	insertSubscription(t, pool, userB, "business")

	auth := NewAPIKeyAuth(pool, rdb, time.Minute)

	var gotUserID int64
	var gotPlan string
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		p, _ := APIKeyPrincipalFromContext(r.Context())
		gotUserID = p.UserID
		gotPlan = p.Plan
		w.WriteHeader(http.StatusOK)
	})

	w := doRequest(t, auth.Middleware()(next), fullKeyB)

	require.Equal(t, http.StatusOK, w.Code)
	require.Equal(t, userB, gotUserID)
	require.Equal(t, "business", gotPlan)
}

func TestAPIKeyAuth_RedisDownFallsThroughToPostgres(t *testing.T) {
	pool := setupAPIKeyAuthTestDB(t)

	prefix, fullKey := testKey(t)
	userID := time.Now().UnixNano()
	insertAPIKey(t, pool, prefix, fullKey, userID, true, false)
	insertSubscription(t, pool, userID, "professional")

	auth := NewAPIKeyAuth(pool, unreachableRedis(t), time.Minute)

	var gotPlan string
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		p, _ := APIKeyPrincipalFromContext(r.Context())
		gotPlan = p.Plan
		w.WriteHeader(http.StatusOK)
	})

	w := doRequest(t, auth.Middleware()(next), fullKey)

	require.Equal(t, http.StatusOK, w.Code)
	require.Equal(t, "professional", gotPlan)
}

func TestAPIKeyAuth_RedisAndPostgresBothDownFailsClosed(t *testing.T) {
	ctx := context.Background()
	badPool, err := pgxpool.New(ctx, "postgres://invalid-host-does-not-exist/db?sslmode=disable&connect_timeout=1")
	require.NoError(t, err) // ParseConfig succeeds; the failure happens on first use
	t.Cleanup(badPool.Close)

	auth := NewAPIKeyAuth(badPool, unreachableRedis(t), time.Minute)
	next, calls := newTestHandler()

	_, fullKey := testKey(t)
	w := doRequest(t, auth.Middleware()(next), fullKey)

	require.Equal(t, http.StatusUnauthorized, w.Code)
	require.EqualValues(t, 0, *calls)
}

func TestAPIKeyAuth_UnknownKeyRejected(t *testing.T) {
	pool := setupAPIKeyAuthTestDB(t)
	rdb := setupAPIKeyAuthTestRedis(t)

	auth := NewAPIKeyAuth(pool, rdb, time.Minute)
	next, calls := newTestHandler()

	_, fullKey := testKey(t) // never inserted
	w := doRequest(t, auth.Middleware()(next), fullKey)

	require.Equal(t, http.StatusUnauthorized, w.Code)
	require.EqualValues(t, 0, *calls)
}

func TestAPIKeyAuth_MalformedKeyRejectedWithoutPanic(t *testing.T) {
	pool := setupAPIKeyAuthTestDB(t)
	rdb := setupAPIKeyAuthTestRedis(t)

	auth := NewAPIKeyAuth(pool, rdb, time.Minute)
	next, calls := newTestHandler()
	handler := auth.Middleware()(next)

	for _, key := range []string{"", "short", "requiem_"} {
		require.NotPanics(t, func() {
			w := doRequest(t, handler, key)
			require.Equal(t, http.StatusUnauthorized, w.Code)
		})
	}

	require.EqualValues(t, 0, *calls)
}

func TestAPIKeyAuth_WarmCacheRejectsPrefixOnlyKey(t *testing.T) {
	// Regression test: a cache entry is a candidate (it carries the matched
	// row's key_hash), not an authorization. Before this was fixed, a cache
	// HIT skipped bcrypt entirely and trusted the prefix alone — so once a
	// legitimate request had warmed the cache for a prefix, presenting just
	// that 12-character prefix (which isn't secret — it's what masked_key
	// displays) was enough to authenticate as that key's owner.
	pool := setupAPIKeyAuthTestDB(t)
	rdb := setupAPIKeyAuthTestRedis(t)

	prefix, fullKey := testKey(t)
	userID := time.Now().UnixNano()
	insertAPIKey(t, pool, prefix, fullKey, userID, true, false)

	auth := NewAPIKeyAuth(pool, rdb, time.Minute)
	next, calls := newTestHandler()
	handler := auth.Middleware()(next)

	// Warm the cache with the real key.
	w := doRequest(t, handler, fullKey)
	require.Equal(t, http.StatusOK, w.Code)

	// Now present only the prefix (exactly keyPrefixLength characters) —
	// must NOT authenticate from the warm cache entry.
	w = doRequest(t, handler, prefix)
	require.Equal(t, http.StatusUnauthorized, w.Code)
	require.EqualValues(t, 1, *calls) // only the first, legitimate request reached next
}

func TestAPIKeyAuth_WarmCacheRejectsAlteredSuffix(t *testing.T) {
	pool := setupAPIKeyAuthTestDB(t)
	rdb := setupAPIKeyAuthTestRedis(t)

	prefix, fullKey := testKey(t)
	userID := time.Now().UnixNano()
	insertAPIKey(t, pool, prefix, fullKey, userID, true, false)

	auth := NewAPIKeyAuth(pool, rdb, time.Minute)
	next, calls := newTestHandler()
	handler := auth.Middleware()(next)

	w := doRequest(t, handler, fullKey)
	require.Equal(t, http.StatusOK, w.Code)

	alteredSuffix := prefix + "zzzzzzzzzz"
	w = doRequest(t, handler, alteredSuffix)
	require.Equal(t, http.StatusUnauthorized, w.Code)
	require.EqualValues(t, 1, *calls)
}

func TestAPIKeyAuth_WarmCacheFallsThroughToCorrectCollisionCandidate(t *testing.T) {
	// A cache entry only remembers the last-verified candidate for a prefix.
	// If a second, different key shares that prefix, presenting it must
	// still resolve to *that* key's own principal — not 401, and not
	// silently authenticated as the first (cached) key.
	pool := setupAPIKeyAuthTestDB(t)
	rdb := setupAPIKeyAuthTestRedis(t)

	prefix := "requiem_" + randomHex(t, 2)
	fullKeyA := prefix + "aaaaaaaaaaaaaaaaaaaa"
	fullKeyB := prefix + "bbbbbbbbbbbbbbbbbbbb"

	userA := time.Now().UnixNano()
	userB := userA + 1

	insertAPIKey(t, pool, prefix, fullKeyA, userA, true, false)
	insertAPIKey(t, pool, prefix, fullKeyB, userB, true, false)

	auth := NewAPIKeyAuth(pool, rdb, time.Minute)

	var gotUserID int64
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		p, _ := APIKeyPrincipalFromContext(r.Context())
		gotUserID = p.UserID
		w.WriteHeader(http.StatusOK)
	})
	handler := auth.Middleware()(next)

	// Warm the cache with key A.
	w := doRequest(t, handler, fullKeyA)
	require.Equal(t, http.StatusOK, w.Code)
	require.Equal(t, userA, gotUserID)

	// Presenting key B (same prefix, different secret) must resolve to B's
	// own principal via a fresh Postgres lookup, not A's cached one.
	w = doRequest(t, handler, fullKeyB)
	require.Equal(t, http.StatusOK, w.Code)
	require.Equal(t, userB, gotUserID)
}

func TestAPIKeyAuth_ResponseDoesNotLeakWhetherPrefixMatched(t *testing.T) {
	pool := setupAPIKeyAuthTestDB(t)
	rdb := setupAPIKeyAuthTestRedis(t)

	prefix, fullKey := testKey(t)
	userID := time.Now().UnixNano()
	insertAPIKey(t, pool, prefix, fullKey, userID, true, false)

	auth := NewAPIKeyAuth(pool, rdb, time.Minute)
	next, _ := newTestHandler()
	handler := auth.Middleware()(next)

	// Right prefix, wrong secret suffix (fails bcrypt compare) vs. a prefix
	// that doesn't exist at all: both must produce byte-identical bodies.
	wrongSuffix := prefix + "00000000000000000000"
	w1 := doRequest(t, handler, wrongSuffix)

	_, unknownKey := testKey(t)
	w2 := doRequest(t, handler, unknownKey)

	require.Equal(t, http.StatusUnauthorized, w1.Code)
	require.Equal(t, http.StatusUnauthorized, w2.Code)

	var body1, body2 map[string]any
	require.NoError(t, json.Unmarshal(w1.Body.Bytes(), &body1))
	require.NoError(t, json.Unmarshal(w2.Body.Bytes(), &body2))
	require.Equal(t, body1["error"], body2["error"])
	require.Equal(t, body1["message"], body2["message"])
}
