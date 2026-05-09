package randomuser

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
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

func TestRandomUserBatchHandler(t *testing.T) {
	t.Parallel()

	t.Run("returns 200 with correct number of users", func(t *testing.T) {
		t.Parallel()

		r := newTestRouter(NewService())

		body := `{"count":5}`
		req := httptest.NewRequest(http.MethodPost, "/random-user/batch", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		require.Equal(t, http.StatusOK, w.Code, "expected 200: %s", w.Body.String())

		var resp httpx.Response[BatchGenerateResponse]
		require.NoError(t, json.NewDecoder(w.Body).Decode(&resp))

		assert.Equal(t, 5, resp.Data.Total)
		assert.Len(t, resp.Data.Results, 5)
	})

	t.Run("each user in batch has all fields populated", func(t *testing.T) {
		t.Parallel()

		r := newTestRouter(NewService())

		body := `{"count":3}`
		req := httptest.NewRequest(http.MethodPost, "/random-user/batch", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		require.Equal(t, http.StatusOK, w.Code)

		var resp httpx.Response[BatchGenerateResponse]
		require.NoError(t, json.NewDecoder(w.Body).Decode(&resp))

		for i, u := range resp.Data.Results {
			assert.NotEmpty(t, u.Name, "user[%d].Name should not be empty", i)
			assert.NotEmpty(t, u.Email, "user[%d].Email should not be empty", i)
			assert.NotEmpty(t, u.Phone, "user[%d].Phone should not be empty", i)
			assert.NotEmpty(t, u.Address.Street, "user[%d].Address.Street should not be empty", i)
			assert.NotEmpty(t, u.Avatar, "user[%d].Avatar should not be empty", i)
		}
	})

	t.Run("sets X-Usage-Count header equal to count", func(t *testing.T) {
		t.Parallel()

		r := newTestRouter(NewService())

		body := `{"count":7}`
		req := httptest.NewRequest(http.MethodPost, "/random-user/batch", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		require.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "7", w.Header().Get("X-Usage-Count"))
	})

	t.Run("returns 422 when count is zero", func(t *testing.T) {
		t.Parallel()

		r := newTestRouter(NewService())

		body := `{"count":0}`
		req := httptest.NewRequest(http.MethodPost, "/random-user/batch", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
	})

	t.Run("returns 422 when count exceeds limit", func(t *testing.T) {
		t.Parallel()

		r := newTestRouter(NewService())

		body := `{"count":51}`
		req := httptest.NewRequest(http.MethodPost, "/random-user/batch", strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
	})

	t.Run("returns 400 on malformed JSON", func(t *testing.T) {
		t.Parallel()

		r := newTestRouter(NewService())

		req := httptest.NewRequest(http.MethodPost, "/random-user/batch", strings.NewReader(`{not valid json`))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}
