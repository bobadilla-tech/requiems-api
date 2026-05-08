package randomuser

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"requiems-api/platform/httpx"
)

func newTestRouter(svc *Service) http.Handler {
	r := chi.NewRouter()
	RegisterRoutes(r, svc)
	return r
}

func TestRandomUserHandler(t *testing.T) {
	t.Parallel()
	t.Run("returns 200 with valid user fields", func(t *testing.T) {
		t.Parallel()
		svc := NewService()

		req := httptest.NewRequest(http.MethodGet, "/random-user", http.NoBody)
		w := httptest.NewRecorder()

		newTestRouter(svc).ServeHTTP(w, req)

		require.Equal(t, http.StatusOK, w.Code)

		var resp httpx.Response[User]
		err := json.NewDecoder(w.Body).Decode(&resp)
		require.NoError(t, err)

		u := resp.Data
		assert.NotEmpty(t, u.Name)
		assert.NotEmpty(t, u.Email)
		assert.NotEmpty(t, u.Phone)
		assert.NotEmpty(t, u.Address.Street)
		assert.NotEmpty(t, u.Avatar)
	})

	t.Run("content-type is application/json", func(t *testing.T) {
		t.Parallel()
		svc := NewService()

		req := httptest.NewRequest(http.MethodGet, "/random-user", http.NoBody)
		w := httptest.NewRecorder()

		newTestRouter(svc).ServeHTTP(w, req)

		ct := w.Header().Get("Content-Type")
		assert.Equal(t, "application/json", ct)
	})

	t.Run("returns different users on successive calls", func(t *testing.T) {
		t.Parallel()
		svc := NewService()
		router := newTestRouter(svc)

		names := make(map[string]struct{})
		for range 10 {
			req := httptest.NewRequest(http.MethodGet, "/random-user", http.NoBody)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			var resp httpx.Response[User]
			err := json.NewDecoder(w.Body).Decode(&resp)
			require.NoError(t, err)
			names[resp.Data.Name] = struct{}{}
		}

		// With 30 first × 30 last = 900 combinations, 10 calls should yield > 1 unique name.
		if len(names) <= 1 {
			t.Error("expected varied output across multiple calls")
		}
	})
}
