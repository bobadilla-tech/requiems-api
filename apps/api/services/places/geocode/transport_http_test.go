package geocode

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

func setupRouter(mockServer *httptest.Server) chi.Router {
	svc := NewService(mockServer.URL, mockServer.Client(), nil)
	r := chi.NewRouter()
	RegisterRoutes(r, svc)
	return r
}

func TestGeocode_HappyPath(t *testing.T) {
	t.Parallel()
	mock := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`[{"lat":"38.8976763","lon":"-77.0365298","display_name":"White House, Washington, DC","address":{"city":"Washington","country_code":"us"}}]`)) //nolint:errcheck // test helper, write error is inconsequential
	}))
	defer mock.Close()

	req := httptest.NewRequest(http.MethodGet, "/geocode?address=1600+Pennsylvania+Ave", http.NoBody)
	w := httptest.NewRecorder()
	setupRouter(mock).ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[GeocodeResponse]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	assert.Equal(t, "Washington", resp.Data.City)
	assert.Equal(t, "US", resp.Data.Country)
	if resp.Data.Lat == 0 {
		t.Error("expected non-zero latitude")
	}
}

func TestGeocode_MissingAddress(t *testing.T) {
	t.Parallel()
	mock := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte(`[]`)) //nolint:errcheck // test helper, write error is inconsequential
	}))
	defer mock.Close()

	req := httptest.NewRequest(http.MethodGet, "/geocode", http.NoBody)
	w := httptest.NewRecorder()
	setupRouter(mock).ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGeocode_NoResults(t *testing.T) {
	t.Parallel()
	mock := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`[]`)) //nolint:errcheck // test helper, write error is inconsequential
	}))
	defer mock.Close()

	req := httptest.NewRequest(http.MethodGet, "/geocode?address=zzznoresultsxxx", http.NoBody)
	w := httptest.NewRecorder()
	setupRouter(mock).ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestGeocode_UpstreamError(t *testing.T) {
	t.Parallel()
	mock := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer mock.Close()

	req := httptest.NewRequest(http.MethodGet, "/geocode?address=anywhere", http.NoBody)
	w := httptest.NewRecorder()
	setupRouter(mock).ServeHTTP(w, req)

	assert.Equal(t, http.StatusServiceUnavailable, w.Code)
}

func TestReverseGeocode_HappyPath(t *testing.T) {
	t.Parallel()
	mock := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"display_name":"White House, Washington, DC","address":{"city":"Washington","country_code":"us"}}`)) //nolint:errcheck // test helper, write error is inconsequential
	}))
	defer mock.Close()

	req := httptest.NewRequest(http.MethodGet, "/reverse-geocode?lat=38.8977&lon=-77.0365", http.NoBody)
	w := httptest.NewRecorder()
	setupRouter(mock).ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[ReverseGeocodeResponse]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	assert.Equal(t, "Washington", resp.Data.City)
}

func TestReverseGeocode_MissingParams(t *testing.T) {
	t.Parallel()
	mock := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {}))
	defer mock.Close()

	req := httptest.NewRequest(http.MethodGet, "/reverse-geocode?lat=38.8977", http.NoBody)
	w := httptest.NewRecorder()
	setupRouter(mock).ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}
