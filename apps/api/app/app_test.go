package app

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"requiems-api/platform/config"
)

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

// TestApp_Handler is an integration test that creates a real App and verifies
// the HTTP handler has the expected routing structure:
//   - GET /healthz is publicly accessible (no auth required)
//   - GET /v1/* routes require the X-Backend-Secret header (returns 401 when absent)
//
// The test is skipped when DATABASE_URL or BACKEND_SECRET is not set.
func TestApp_Handler(t *testing.T) {
	dsn := os.Getenv("DATABASE_URL")

	if dsn == "" {
		t.Skip("DATABASE_URL not set; skipping App integration test")
	}

	backendSecret := os.Getenv("BACKEND_SECRET")

	if backendSecret == "" {
		t.Skip("BACKEND_SECRET not set; skipping App integration test")
	}

	redisURL := os.Getenv("REDIS_URL")

	if redisURL == "" {
		redisURL = "redis://localhost:6379/0"
	}

	t.Chdir("..") // resolve "migrations" relative to api root, not package dir

	cfg := config.Config{
		DatabaseURL:   dsn,
		BackendSecret: backendSecret,
		RedisURL:      redisURL,
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

	t.Run("v1 routes require backend secret", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/v1/text/advice", http.NoBody)
		w := httptest.NewRecorder()
		h.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})

	t.Run("v1 routes are accessible with valid backend secret", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/v1/text/advice", http.NoBody)
		req.Header.Set("X-Backend-Secret", backendSecret)
		w := httptest.NewRecorder()
		h.ServeHTTP(w, req)

		// The endpoint itself may return any 2xx; 401/403 would indicate auth failure.
		assert.NotEqual(t, http.StatusUnauthorized, w.Code)
		assert.NotEqual(t, http.StatusForbidden, w.Code)
	})
}
