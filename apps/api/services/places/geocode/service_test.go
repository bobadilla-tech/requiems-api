package geocode

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"requiems-api/platform/svcerr"
)

func fakeSearchServer(results []nominatimSearchResult) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(results)
	}))
}

func fakeReverseServer(result nominatimReverseResult) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)
	}))
}

func newTestService(srv *httptest.Server) *Service {
	return NewService(srv.URL, srv.Client(), nil)
}

// --- Geocode ---

func TestGeocodeService_HappyPath(t *testing.T) {
	t.Parallel()
	srv := fakeSearchServer([]nominatimSearchResult{
		{
			Lat:         "48.8566",
			Lon:         "2.3522",
			DisplayName: "Paris, France",
			Address: nominatimAddress{
				City:        "Paris",
				CountryCode: "fr",
			},
		},
	})
	defer srv.Close()

	svc := newTestService(srv)
	res, err := svc.Geocode(t.Context(), "Paris, France")
	require.NoError(t, err)

	assert.InDelta(t, 48.8566, res.Lat, 0.001)
	assert.InDelta(t, 2.3522, res.Lon, 0.001)
	assert.Equal(t, "Paris, France", res.Address)
	assert.Equal(t, "Paris", res.City)
	assert.Equal(t, "FR", res.Country)
}

func TestGeocodeService_EmptyResults_NotFound(t *testing.T) {
	t.Parallel()
	srv := fakeSearchServer([]nominatimSearchResult{})
	defer srv.Close()

	svc := newTestService(srv)
	_, err := svc.Geocode(t.Context(), "Nonexistent Place XYZ")
	require.Error(t, err)

	var se *svcerr.Error
	require.ErrorAs(t, err, &se)
	assert.Equal(t, svcerr.KindNotFound, se.Kind)
	assert.Equal(t, "not_found", se.Code)
}

func TestGeocodeService_NonOKStatus_Upstream(t *testing.T) {
	t.Parallel()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "internal error", http.StatusInternalServerError)
	}))
	defer srv.Close()

	svc := newTestService(srv)
	_, err := svc.Geocode(t.Context(), "anywhere")
	require.Error(t, err)

	var se *svcerr.Error
	require.ErrorAs(t, err, &se)
	assert.Equal(t, svcerr.KindUpstream, se.Kind)
}

func TestGeocodeService_InvalidJSON_Upstream(t *testing.T) {
	t.Parallel()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("not-json"))
	}))
	defer srv.Close()

	svc := newTestService(srv)
	_, err := svc.Geocode(t.Context(), "anywhere")
	require.Error(t, err)

	var se *svcerr.Error
	require.ErrorAs(t, err, &se)
	assert.Equal(t, svcerr.KindUpstream, se.Kind)
}

// --- ReverseGeocode ---

func TestReverseGeocodeService_HappyPath(t *testing.T) {
	t.Parallel()
	srv := fakeReverseServer(nominatimReverseResult{
		DisplayName: "Eiffel Tower, Paris, France",
		Address: nominatimAddress{
			City:        "Paris",
			CountryCode: "fr",
		},
	})
	defer srv.Close()

	svc := newTestService(srv)
	res, err := svc.ReverseGeocode(t.Context(), 48.8584, 2.2945)
	require.NoError(t, err)

	assert.Equal(t, "Eiffel Tower, Paris, France", res.Address)
	assert.Equal(t, "Paris", res.City)
	assert.Equal(t, "FR", res.Country)
	assert.InDelta(t, 48.8584, res.Lat, 0.001)
	assert.InDelta(t, 2.2945, res.Lon, 0.001)
}

func TestReverseGeocodeService_EmptyDisplayName_NotFound(t *testing.T) {
	t.Parallel()
	srv := fakeReverseServer(nominatimReverseResult{DisplayName: ""})
	defer srv.Close()

	svc := newTestService(srv)
	_, err := svc.ReverseGeocode(t.Context(), 0, 0)
	require.Error(t, err)

	var se *svcerr.Error
	require.ErrorAs(t, err, &se)
	assert.Equal(t, svcerr.KindNotFound, se.Kind)
}

func TestReverseGeocodeService_NonOKStatus_Upstream(t *testing.T) {
	t.Parallel()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusServiceUnavailable)
	}))
	defer srv.Close()

	svc := newTestService(srv)
	_, err := svc.ReverseGeocode(t.Context(), 0, 0)
	require.Error(t, err)

	var se *svcerr.Error
	require.ErrorAs(t, err, &se)
	assert.Equal(t, svcerr.KindUpstream, se.Kind)
}

func TestReverseGeocodeService_InvalidJSON_Upstream(t *testing.T) {
	t.Parallel()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("{bad json"))
	}))
	defer srv.Close()

	svc := newTestService(srv)
	_, err := svc.ReverseGeocode(t.Context(), 0, 0)
	require.Error(t, err)

	var se *svcerr.Error
	require.ErrorAs(t, err, &se)
	assert.Equal(t, svcerr.KindUpstream, se.Kind)
}

// --- resolveCity ---

func TestResolveCity_City(t *testing.T) {
	t.Parallel()
	a := nominatimAddress{City: "Paris", Town: "Suburb", Village: "Village", County: "County"}
	assert.Equal(t, "Paris", a.resolveCity())
}

func TestResolveCity_Town(t *testing.T) {
	t.Parallel()
	a := nominatimAddress{Town: "Suburb", Village: "Village", County: "County"}
	assert.Equal(t, "Suburb", a.resolveCity())
}

func TestResolveCity_Village(t *testing.T) {
	t.Parallel()
	a := nominatimAddress{Village: "Village", County: "County"}
	assert.Equal(t, "Village", a.resolveCity())
}

func TestResolveCity_County(t *testing.T) {
	t.Parallel()
	a := nominatimAddress{County: "County"}
	assert.Equal(t, "County", a.resolveCity())
}

func TestGeocodeBatch_Service_HappyPath(t *testing.T) {
	t.Parallel()
	srv := fakeSearchServer([]nominatimSearchResult{
		{Lat: "48.8566", Lon: "2.3522", DisplayName: "Paris, France",
			Address: nominatimAddress{City: "Paris", CountryCode: "fr"}},
	})
	defer srv.Close()

	svc := newTestService(srv)
	results := svc.GeocodeBatch(t.Context(), []string{"Paris", "France"})

	require.Len(t, results, 2)
	for i, r := range results {
		assert.NotNil(t, r.Result, "expected result at index %d", i)
		assert.Empty(t, r.Error, "expected no error at index %d", i)
	}
}

func TestGeocodeBatch_Service_PartialNotFound(t *testing.T) {
	t.Parallel()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
	defer srv.Close()

	svc := newTestService(srv)
	results := svc.GeocodeBatch(t.Context(), []string{"Paris", "unknown place xyz"})

	require.Len(t, results, 2)
	assert.NotNil(t, results[0].Result)
	assert.Empty(t, results[0].Error)
	assert.Nil(t, results[1].Result)
	assert.NotEmpty(t, results[1].Error)
}

func TestReverseGeocodeBatch_Service_HappyPath(t *testing.T) {
	t.Parallel()
	srv := fakeReverseServer(nominatimReverseResult{
		DisplayName: "Eiffel Tower, Paris, France",
		Address:     nominatimAddress{City: "Paris", CountryCode: "fr"},
	})
	defer srv.Close()

	svc := newTestService(srv)
	items := []ReverseQuery{{Lat: 48.8584, Lon: 2.2945}, {Lat: 51.5014, Lon: -0.1419}}
	results := svc.ReverseGeocodeBatch(t.Context(), items)

	require.Len(t, results, 2)
	for i, r := range results {
		assert.NotNil(t, r.Result, "expected result at index %d", i)
		assert.Empty(t, r.Error, "expected no error at index %d", i)
	}
}

func TestReverseGeocodeBatch_Service_PartialNotFound(t *testing.T) {
	t.Parallel()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.URL.Query().Get("lat") == "0.000000" {
				json.NewEncoder(w).Encode(nominatimReverseResult{DisplayName: ""}) //nolint:errcheck // test server response body writes are intentionally ignored
		} else {
				json.NewEncoder(w).Encode(nominatimReverseResult{ //nolint:errcheck // test server response body writes are intentionally ignored
				DisplayName: "Eiffel Tower, Paris, France",
				Address:     nominatimAddress{City: "Paris", CountryCode: "fr"},
			})
		}
	}))
	defer srv.Close()

	svc := newTestService(srv)
	items := []ReverseQuery{{Lat: 48.8584, Lon: 2.2945}, {Lat: 0, Lon: 0}}
	results := svc.ReverseGeocodeBatch(t.Context(), items)

	require.Len(t, results, 2)
	assert.NotNil(t, results[0].Result)
	assert.Empty(t, results[0].Error)
	assert.Nil(t, results[1].Result)
	assert.NotEmpty(t, results[1].Error)
}
