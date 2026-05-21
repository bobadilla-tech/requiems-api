package geocode

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

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
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

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
}

func TestGeocodeBatch_HappyPath(t *testing.T) {
	t.Parallel()
	mock := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode([]nominatimSearchResult{ //nolint:errcheck // test server response body writes are intentionally ignored
			{Lat: "48.8566", Lon: "2.3522", DisplayName: "Paris, France",
				Address: nominatimAddress{City: "Paris", CountryCode: "fr"}},
		})
	}))
	defer mock.Close()

	body := `{"addresses":["Paris","London"]}`
	req := httptest.NewRequest(http.MethodPost, "/geocode/batch", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	setupRouter(mock).ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "2", w.Header().Get("X-Usage-Count"))
	var resp httpx.Response[httpx.BatchResponse[BatchGeocodeItem]]
	require.NoError(t, json.NewDecoder(w.Body).Decode(&resp))
	assert.Equal(t, 2, resp.Data.Total)
	assert.NotNil(t, resp.Data.Results[0].Result)
	assert.NotNil(t, resp.Data.Results[1].Result)
}

func TestGeocodeBatch_PartialError(t *testing.T) {
	t.Parallel()
	mock := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if strings.Contains(r.URL.RawQuery, "unknown") {
			json.NewEncoder(w).Encode([]nominatimSearchResult{}) //nolint:errcheck // test server response body writes are intentionally ignored
		} else {
			json.NewEncoder(w).Encode([]nominatimSearchResult{ //nolint:errcheck // test server response body writes are intentionally ignored
				{Lat: "48.8566", Lon: "2.3522", DisplayName: "Paris, France",
					Address: nominatimAddress{City: "Paris", CountryCode: "fr"}},
			})
		}
	}))
	defer mock.Close()

	body := `{"addresses":["Paris","unknown place xyz"]}`
	req := httptest.NewRequest(http.MethodPost, "/geocode/batch", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	setupRouter(mock).ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)
	var resp httpx.Response[httpx.BatchResponse[BatchGeocodeItem]]
	require.NoError(t, json.NewDecoder(w.Body).Decode(&resp))
	require.Len(t, resp.Data.Results, 2)
	assert.NotNil(t, resp.Data.Results[0].Result)
	assert.Nil(t, resp.Data.Results[1].Result)
	assert.NotEmpty(t, resp.Data.Results[1].Error)
}

func TestGeocodeBatch_EmptyArray(t *testing.T) {
	t.Parallel()
	mock := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte(`[]`)) //nolint:errcheck // test server response body writes are intentionally ignored
	}))
	defer mock.Close()

	req := httptest.NewRequest(http.MethodPost, "/geocode/batch", strings.NewReader(`{"addresses":[]}`))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	setupRouter(mock).ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
}

func TestGeocodeBatch_ExceedsLimit(t *testing.T) {
	t.Parallel()
	mock := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte(`[]`)) //nolint:errcheck // test server response body writes are intentionally ignored
	}))
	defer mock.Close()

	addrs := make([]string, 21)
	for i := range addrs {
		addrs[i] = `"addr"`
	}
	body := `{"addresses":[` + strings.Join(addrs, ",") + `]}`
	req := httptest.NewRequest(http.MethodPost, "/geocode/batch", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	setupRouter(mock).ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
}

func TestReverseGeocodeBatch_HappyPath(t *testing.T) {
	t.Parallel()
	mock := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"display_name":"Eiffel Tower, Paris, France","address":{"city":"Paris","country_code":"fr"}}`)) //nolint:errcheck // test server response body writes are intentionally ignored
	}))
	defer mock.Close()

	body := `{"items":[{"lat":48.8584,"lon":2.2945}]}`
	req := httptest.NewRequest(http.MethodPost, "/reverse-geocode/batch", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	setupRouter(mock).ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)
	var resp httpx.Response[httpx.BatchResponse[BatchReverseGeocodeItem]]
	require.NoError(t, json.NewDecoder(w.Body).Decode(&resp))
	assert.Equal(t, 1, resp.Data.Total)
	require.Len(t, resp.Data.Results, 1)
	assert.NotNil(t, resp.Data.Results[0].Result)
}

func TestReverseGeocodeBatch_EmptyArray(t *testing.T) {
	t.Parallel()
	mock := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {}))
	defer mock.Close()

	req := httptest.NewRequest(http.MethodPost, "/reverse-geocode/batch", strings.NewReader(`{"items":[]}`))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	setupRouter(mock).ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
}

func TestReverseGeocodeBatch_ExceedsLimit(t *testing.T) {
	t.Parallel()
	mock := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {}))
	defer mock.Close()

	items := make([]string, 21)
	for i := range items {
		items[i] = `{"lat":48.8584,"lon":2.2945}`
	}
	body := `{"items":[` + strings.Join(items, ",") + `]}`
	req := httptest.NewRequest(http.MethodPost, "/reverse-geocode/batch", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	setupRouter(mock).ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
}
