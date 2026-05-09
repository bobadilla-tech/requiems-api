package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequestID(t *testing.T) {
	t.Parallel()

	captureHandler := func(gotID *string) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			*gotID = GetRequestID(r.Context())
			w.WriteHeader(http.StatusOK)
		})
	}

	t.Run("propagates existing X-Request-ID", func(t *testing.T) {
		t.Parallel()

		var ctxID string
		handler := RequestID(captureHandler(&ctxID))

		req := httptest.NewRequest(http.MethodGet, "/", http.NoBody)
		req.Header.Set("X-Request-ID", "trace-abc-123")
		w := httptest.NewRecorder()

		handler.ServeHTTP(w, req)

		assert.Equal(t, "trace-abc-123", ctxID)
		assert.Equal(t, "trace-abc-123", w.Header().Get("X-Request-ID"))
	})

	t.Run("generates ID when header is absent", func(t *testing.T) {
		t.Parallel()

		var ctxID string
		handler := RequestID(captureHandler(&ctxID))

		req := httptest.NewRequest(http.MethodGet, "/", http.NoBody)
		w := httptest.NewRecorder()

		handler.ServeHTTP(w, req)

		assert.NotEmpty(t, ctxID)
		assert.Equal(t, ctxID, w.Header().Get("X-Request-ID"))
	})

	t.Run("generated IDs are unique", func(t *testing.T) {
		t.Parallel()

		ids := make(map[string]struct{}, 100)
		for i := 0; i < 100; i++ {
			var ctxID string
			handler := RequestID(captureHandler(&ctxID))
			req := httptest.NewRequest(http.MethodGet, "/", http.NoBody)
			handler.ServeHTTP(httptest.NewRecorder(), req)
			ids[ctxID] = struct{}{}
		}
		assert.Len(t, ids, 100)
	})
}

func TestGetRequestID_missing(t *testing.T) {
	t.Parallel()

	req := httptest.NewRequest(http.MethodGet, "/", http.NoBody)
	assert.Equal(t, "", GetRequestID(req.Context()))
}
