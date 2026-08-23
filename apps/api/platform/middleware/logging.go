package middleware

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	chimw "github.com/go-chi/chi/v5/middleware"
)

// RequestLogger emits one structured JSON log line per request via log/slog:
// request ID, method, route pattern, status code, and latency. Sentry is
// unaffected — it continues to own uncaught-exception capture, not
// per-request logging.
func RequestLogger(logger *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			ww := chimw.NewWrapResponseWriter(w, r.ProtoMajor)

			next.ServeHTTP(ww, r)

			route := chi.RouteContext(r.Context())
			pattern := r.URL.Path

			if route != nil && route.RoutePattern() != "" {
				pattern = route.RoutePattern()
			}

			// A handler that never calls WriteHeader/Write leaves ww.Status()
			// at its zero value, but net/http itself writes a 200 to the wire
			// in that case — match that behavior instead of logging "status":0.
			status := ww.Status()
			if status == 0 {
				status = http.StatusOK
			}

			logger.Info("request",
				"request_id", GetRequestID(r.Context()),
				"method", r.Method,
				"route", pattern,
				"status", status,
				"latency_ms", time.Since(start).Milliseconds(),
			)
		})
	}
}
