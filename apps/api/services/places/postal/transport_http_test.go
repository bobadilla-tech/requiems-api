package postal

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
	svc := &Service{
		index: map[string]PostalCode{
			"US:10001": {
				PostalCode: "10001",
				City:       "New York City",
				State:      "New York",
				Country:    "US",
				Lat:        40.7484,
				Lon:        -73.9967,
			},
			"GB:SW1A1AA": {
				PostalCode: "SW1A1AA",
				City:       "London",
				State:      "England",
				Country:    "GB",
				Lat:        51.5014,
				Lon:        -0.1419,
			},
		},
	}
	r := chi.NewRouter()
	RegisterRoutes(r, svc)
	return r
}

func TestLookup_HappyPath(t *testing.T) {
	t.Parallel()
	req := httptest.NewRequest(http.MethodGet, "/postal/10001?country=US", http.NoBody)
	w := httptest.NewRecorder()
	setupRouter().ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[PostalCode]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	assert.Equal(t, "New York City", resp.Data.City)
	if resp.Data.Lat == 0 {
		t.Error("expected non-zero latitude")
	}
}

func TestLookup_DefaultsToUS(t *testing.T) {
	t.Parallel()
	req := httptest.NewRequest(http.MethodGet, "/postal/10001", http.NoBody)
	w := httptest.NewRecorder()
	setupRouter().ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)
}

func TestLookup_NotFound(t *testing.T) {
	t.Parallel()
	req := httptest.NewRequest(http.MethodGet, "/postal/99999?country=US", http.NoBody)
	w := httptest.NewRecorder()
	setupRouter().ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestLookup_NonUSCountry(t *testing.T) {
	t.Parallel()
	req := httptest.NewRequest(http.MethodGet, "/postal/SW1A1AA?country=GB", http.NoBody)
	w := httptest.NewRecorder()
	setupRouter().ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[PostalCode]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	assert.Equal(t, "GB", resp.Data.Country)
}

func TestPostal_Batch_HappyPath(t *testing.T) {
	t.Parallel()
	body := `{"items":[{"code":"10001","country":"US"},{"code":"SW1A1AA","country":"GB"}]}`
	req := httptest.NewRequest(http.MethodPost, "/postal/batch", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	setupRouter().ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "2", w.Header().Get("X-Usage-Count"))
	var resp httpx.Response[httpx.BatchResponse[BatchResult]]
	require.NoError(t, json.NewDecoder(w.Body).Decode(&resp))
	assert.Equal(t, 2, resp.Data.Total)
	assert.True(t, resp.Data.Results[0].Found)
	assert.True(t, resp.Data.Results[1].Found)
}

func TestPostal_Batch_NotFound(t *testing.T) {
	t.Parallel()
	body := `{"items":[{"code":"10001","country":"US"},{"code":"99999","country":"US"}]}`
	req := httptest.NewRequest(http.MethodPost, "/postal/batch", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	setupRouter().ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "2", w.Header().Get("X-Usage-Count"))
	var resp httpx.Response[httpx.BatchResponse[BatchResult]]
	require.NoError(t, json.NewDecoder(w.Body).Decode(&resp))
	require.Len(t, resp.Data.Results, 2)
	assert.True(t, resp.Data.Results[0].Found)
	assert.False(t, resp.Data.Results[1].Found)
	assert.Nil(t, resp.Data.Results[1].Result)
}

func TestPostal_Batch_DefaultCountry(t *testing.T) {
	t.Parallel()
	body := `{"items":[{"code":"10001"}]}`
	req := httptest.NewRequest(http.MethodPost, "/postal/batch", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	setupRouter().ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "1", w.Header().Get("X-Usage-Count"))
	var resp httpx.Response[httpx.BatchResponse[BatchResult]]
	require.NoError(t, json.NewDecoder(w.Body).Decode(&resp))
	require.Len(t, resp.Data.Results, 1)
	assert.Equal(t, "US", resp.Data.Results[0].Country)
}

func TestPostal_Batch_EmptyArray(t *testing.T) {
	t.Parallel()
	req := httptest.NewRequest(http.MethodPost, "/postal/batch", strings.NewReader(`{"items":[]}`))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	setupRouter().ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
}

func TestPostal_Batch_ExceedsLimit(t *testing.T) {
	t.Parallel()
	items := make([]string, 51)
	for i := range items {
		items[i] = `{"code":"10001","country":"US"}`
	}
	body := `{"items":[` + strings.Join(items, ",") + `]}`
	req := httptest.NewRequest(http.MethodPost, "/postal/batch", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	setupRouter().ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
}
