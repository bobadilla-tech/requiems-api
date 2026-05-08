package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBackendSecretAuth(t *testing.T) {
	t.Parallel()

	validSecret := "this_is_a_valid_secret_with_32_chars_minimum"

	// Test handler that just returns 200 OK
	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("authorized"))
	})

	t.Run("allows request with valid secret", func(t *testing.T) {
		t.Parallel()

		middleware := BackendSecretAuth(validSecret)
		handler := middleware(testHandler)

		req := httptest.NewRequest(http.MethodGet, "/test", http.NoBody)
		req.Header.Set("X-Backend-Secret", validSecret)
		w := httptest.NewRecorder()

		handler.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("rejects request with missing header", func(t *testing.T) {
		t.Parallel()

		middleware := BackendSecretAuth(validSecret)
		handler := middleware(testHandler)

		req := httptest.NewRequest(http.MethodGet, "/test", http.NoBody)
		w := httptest.NewRecorder()

		handler.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})

	t.Run("rejects request with invalid secret", func(t *testing.T) {
		t.Parallel()

		middleware := BackendSecretAuth(validSecret)
		handler := middleware(testHandler)

		req := httptest.NewRequest(http.MethodGet, "/test", http.NoBody)
		req.Header.Set("X-Backend-Secret", "wrong_secret")
		w := httptest.NewRecorder()

		handler.ServeHTTP(w, req)

		assert.Equal(t, http.StatusForbidden, w.Code)
	})

	t.Run("panics if secret is empty", func(t *testing.T) {
		t.Parallel()

		defer func() {
			if r := recover(); r == nil {
				t.Error("expected panic when secret is empty")
			}
		}()

		BackendSecretAuth("")
	})

	t.Run("panics if secret is too short", func(t *testing.T) {
		t.Parallel()

		defer func() {
			if r := recover(); r == nil {
				t.Error("expected panic when secret is too short")
			}
		}()

		BackendSecretAuth("short")
	})
}
