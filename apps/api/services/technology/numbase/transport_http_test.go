package numbase

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

func TestNumbase_Batch_HappyPath(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	body := `{"items":[{"from":10,"to":16,"value":"255"},{"from":2,"to":10,"value":"11111111"}]}`
	req := httptest.NewRequest(http.MethodPost, "/base/batch", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)
	var resp httpx.Response[httpx.BatchResponse[BatchResult]]
	require.NoError(t, json.NewDecoder(w.Body).Decode(&resp))
	assert.Equal(t, 2, resp.Data.Total)
	assert.NotEmpty(t, resp.Data.Results[0].Result)
	assert.NotEmpty(t, resp.Data.Results[1].Result)
}

func TestNumbase_Batch_PartialError(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	body := `{"items":[{"from":10,"to":16,"value":"255"},{"from":10,"to":16,"value":"notanumber"}]}`
	req := httptest.NewRequest(http.MethodPost, "/base/batch", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)
	var resp httpx.Response[httpx.BatchResponse[BatchResult]]
	require.NoError(t, json.NewDecoder(w.Body).Decode(&resp))
	require.Len(t, resp.Data.Results, 2)
	assert.NotEmpty(t, resp.Data.Results[0].Result)
	assert.Empty(t, resp.Data.Results[0].Error)
	assert.Empty(t, resp.Data.Results[1].Result)
	assert.NotEmpty(t, resp.Data.Results[1].Error)
}

func TestNumbase_Batch_EmptyArray(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	req := httptest.NewRequest(http.MethodPost, "/base/batch", strings.NewReader(`{"items":[]}`))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
}

func TestNumbase_Batch_ExceedsLimit(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	items := make([]string, 51)
	for i := range items {
		items[i] = `{"from":10,"to":16,"value":"1"}`
	}
	body := `{"items":[` + strings.Join(items, ",") + `]}`
	req := httptest.NewRequest(http.MethodPost, "/base/batch", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
}
