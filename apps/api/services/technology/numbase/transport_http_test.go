package numbase

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

func setupRouter() chi.Router {
	r := chi.NewRouter()
	svc := NewService()
	RegisterRoutes(r, svc)
	return r
}

func TestNumbase_HappyPath(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	req := httptest.NewRequest(http.MethodGet, "/base?from=16&to=2&value=0xFF", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "application/json", w.Header().Get("Content-Type"))

	var resp httpx.Response[Result]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	assert.Equal(t, "11111111", resp.Data.Result)
}

func TestNumbase_MissingRequiredParam(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		url  string
	}{
		{
			name: "missing all params",
			url:  "/base",
		},
		{
			name: "missing from",
			url:  "/base?to=2&value=0xFF",
		},
		{
			name: "missing to",
			url:  "/base?from=16&value=0xFF",
		},
		{
			name: "missing value",
			url:  "/base?from=16&to=2",
		},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			t.Parallel()
			r := setupRouter()
			req := httptest.NewRequest(http.MethodGet, v.url, http.NoBody)
			w := httptest.NewRecorder()

			r.ServeHTTP(w, req)

			assert.Equal(t, http.StatusBadRequest, w.Code)
		})
	}
}

func TestNumbase_InvalidParamValue(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		url  string
	}{
		{
			name: "invalid from",
			url:  "/base?from=5&to=2&value=100",
		},
		{
			name: "invalid to",
			url:  "/base?from=10&to=7&value=100",
		},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			t.Parallel()
			r := setupRouter()

			req := httptest.NewRequest(http.MethodGet, v.url, http.NoBody)
			w := httptest.NewRecorder()

			r.ServeHTTP(w, req)

			assert.Equal(t, http.StatusBadRequest, w.Code)
		})
	}
}

func TestNumbase_ServiceError(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		url  string
	}{
		{
			name: "Value not valid for binary",
			url:  "/base?from=2&to=10&value=999",
		},
		{
			name: "value not valid for hex",
			url:  "/base?from=16&to=2&value=ZZZZ",
		},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			t.Parallel()
			r := setupRouter()

			req := httptest.NewRequest(http.MethodGet, v.url, http.NoBody)
			w := httptest.NewRecorder()

			r.ServeHTTP(w, req)

			assert.Equal(t, http.StatusBadRequest, w.Code)
		})
	}
}
