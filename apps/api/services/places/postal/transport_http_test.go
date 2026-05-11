package postal

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
