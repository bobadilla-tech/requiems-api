package middleware

import (
	"bytes"
	"encoding/json"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/require"
)

func TestRequestLogger(t *testing.T) {
	t.Parallel()

	t.Run("logs request id, method, route pattern, status, and latency", func(t *testing.T) {
		t.Parallel()

		var buf bytes.Buffer
		logger := slog.New(slog.NewJSONHandler(&buf, nil))

		r := chi.NewRouter()
		r.Use(RequestID)
		r.Use(RequestLogger(logger))
		r.Get("/things/{id}", func(w http.ResponseWriter, _ *http.Request) {
			w.WriteHeader(http.StatusTeapot)
		})

		req := httptest.NewRequest(http.MethodGet, "/things/42", http.NoBody)
		req.Header.Set("X-Request-ID", "test-request-id")
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		require.Equal(t, http.StatusTeapot, w.Code)

		var logLine map[string]any
		require.NoError(t, json.Unmarshal(buf.Bytes(), &logLine))

		require.Equal(t, "test-request-id", logLine["request_id"])
		require.Equal(t, http.MethodGet, logLine["method"])
		require.Equal(t, "/things/{id}", logLine["route"])
		require.Equal(t, float64(http.StatusTeapot), logLine["status"])
		require.Contains(t, logLine, "latency_ms")
	})

	t.Run("falls back to the request path when no route matches", func(t *testing.T) {
		t.Parallel()

		var buf bytes.Buffer
		logger := slog.New(slog.NewJSONHandler(&buf, nil))

		r := chi.NewRouter()
		r.Use(RequestLogger(logger))
		r.Get("/known", func(w http.ResponseWriter, _ *http.Request) {
			w.WriteHeader(http.StatusOK)
		})

		req := httptest.NewRequest(http.MethodGet, "/does-not-exist", http.NoBody)
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)

		require.Equal(t, http.StatusNotFound, w.Code)

		var logLine map[string]any
		require.NoError(t, json.Unmarshal(buf.Bytes(), &logLine))

		require.Equal(t, "/does-not-exist", logLine["route"])
		require.Equal(t, float64(http.StatusNotFound), logLine["status"])
	})
}
