package cities

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
		index: map[string]City{
			"london": {
				Name:       "London",
				Country:    "GB",
				Population: 7556900,
				Timezone:   "Europe/London",
				Lat:        51.5085,
				Lon:        -0.1257,
			},
			"new york city": {
				Name:       "New York City",
				Country:    "US",
				Population: 8336817,
				Timezone:   "America/New_York",
				Lat:        40.7128,
				Lon:        -74.0060,
			},
		},
	}
	r := chi.NewRouter()
	RegisterRoutes(r, svc)
	return r
}

func TestFind_HappyPath(t *testing.T) {
	t.Parallel()
	req := httptest.NewRequest(http.MethodGet, "/cities/london", http.NoBody)
	w := httptest.NewRecorder()
	setupRouter().ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[City]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	assert.Equal(t, "London", resp.Data.Name)
	assert.Equal(t, "GB", resp.Data.Country)
	assert.Equal(t, "Europe/London", resp.Data.Timezone)
}

func TestFind_CaseInsensitive(t *testing.T) {
	t.Parallel()
	req := httptest.NewRequest(http.MethodGet, "/cities/London", http.NoBody)
	w := httptest.NewRecorder()
	setupRouter().ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)
}

func TestFind_NotFound(t *testing.T) {
	t.Parallel()
	req := httptest.NewRequest(http.MethodGet, "/cities/atlantis", http.NoBody)
	w := httptest.NewRecorder()
	setupRouter().ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestFind_MultiWordCity(t *testing.T) {
	t.Parallel()
	req := httptest.NewRequest(http.MethodGet, "/cities/new%20york%20city", http.NoBody)
	w := httptest.NewRecorder()
	setupRouter().ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[City]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	if resp.Data.Population == 0 {
		t.Error("expected non-zero population")
	}
}

func TestCities_Batch_Found(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	req := httptest.NewRequest(http.MethodPost, "/cities/batch", strings.NewReader(`{"names":["london","atlantis"]}`))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[httpx.BatchResponse[BatchCityItem]]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	assert.Len(t, resp.Data.Results, 2)
	assert.Equal(t, 2, resp.Data.Total)
	assert.True(t, resp.Data.Results[0].Found, "london should be found")
	assert.NotNil(t, resp.Data.Results[0].Result)
	assert.False(t, resp.Data.Results[1].Found, "atlantis should not be found")
	assert.Nil(t, resp.Data.Results[1].Result)
}

func TestCities_Batch_EmptyNames422(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	req := httptest.NewRequest(http.MethodPost, "/cities/batch", strings.NewReader(`{"names":[]}`))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
}
