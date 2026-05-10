package exchange

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"requiems-api/platform/svcerr"
)

// fakeFrankfurter returns a handler that serves a Frankfurter-shaped response.
func fakeFrankfurter(base, target string, rate float64) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp := frankfurterResponse{
			Base:  base,
			Date:  "2024-12-15",
			Rates: map[string]float64{target: rate},
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}
}

func newTestService(handler http.Handler) (svc *Service, cleanup func()) {
	srv := httptest.NewServer(handler)
	svc = newServiceWithClient(nil, srv.Client(), srv.URL)
	return svc, srv.Close
}

func TestGetRate_CacheMiss_FetchesFromAPI(t *testing.T) {
	t.Parallel()
	svc, cleanup := newTestService(fakeFrankfurter("USD", "EUR", 0.92))
	defer cleanup()

	rate, ts, err := svc.GetRate(t.Context(), "USD", "EUR")
	require.NoError(t, err)
	assert.Equal(t, 0.92, rate)
	assert.False(t, ts.IsZero(), "timestamp must not be zero")
}

func TestGetRate_DateParsedCorrectly(t *testing.T) {
	t.Parallel()
	svc, cleanup := newTestService(fakeFrankfurter("USD", "GBP", 0.78))
	defer cleanup()

	_, ts, err := svc.GetRate(t.Context(), "USD", "GBP")
	require.NoError(t, err)
	assert.True(t, ts.Year() == 2024 && ts.Month() == 12 && ts.Day() == 15, "expected date 2024-12-15, got %s", ts.Format("2006-01-02"))
}

func TestGetRate_InvalidTargetCurrency_Returns422(t *testing.T) {
	t.Parallel()
	// API returns empty rates map for an unknown target.
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		resp := frankfurterResponse{Base: "USD", Date: "2024-12-15", Rates: map[string]float64{}}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	})
	svc, cleanup := newTestService(handler)
	defer cleanup()

	_, _, err := svc.GetRate(t.Context(), "USD", "XYZ")
	require.Error(t, err)

	var se *svcerr.Error
	require.ErrorAs(t, err, &se)
	assert.Equal(t, "invalid_currency", se.Code)
	assert.Equal(t, svcerr.KindUnknown, se.Kind)
}

func TestGetRate_APIReturns404_Returns422(t *testing.T) {
	t.Parallel()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r)
	})
	svc, cleanup := newTestService(handler)
	defer cleanup()

	_, _, err := svc.GetRate(t.Context(), "XYZ", "EUR")
	require.Error(t, err)

	var se *svcerr.Error
	require.ErrorAs(t, err, &se)
	assert.Equal(t, "invalid_currency", se.Code)
	assert.Equal(t, svcerr.KindUnknown, se.Kind)
}

func TestGetRate_APIReturns500_Returns503(t *testing.T) {
	t.Parallel()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "internal error", http.StatusInternalServerError)
	})
	svc, cleanup := newTestService(handler)
	defer cleanup()

	_, _, err := svc.GetRate(t.Context(), "USD", "EUR")
	require.Error(t, err)

	var se *svcerr.Error
	require.ErrorAs(t, err, &se)
	assert.Equal(t, "upstream_error", se.Code)
	assert.Equal(t, svcerr.KindUpstream, se.Kind)
}

func TestParseCache_RoundTrip(t *testing.T) {
	t.Parallel()
	rate := 0.9205
	ts := time.Date(2024, 12, 15, 0, 0, 0, 0, time.UTC)
	val := formatCacheValue(rate, ts)

	gotRate, gotTS, err := parseCache(val)
	require.NoError(t, err)
	assert.Equal(t, rate, gotRate)
	assert.True(t, gotTS.Equal(ts), "ts: want %v, got %v", ts, gotTS)
}

func TestParseCache_InvalidFormat(t *testing.T) {
	t.Parallel()
	_, _, err := parseCache("notvalid")
	assert.Error(t, err, "expected error for invalid cache value")
}

func TestCacheKey_AlwaysUppercase(t *testing.T) {
	t.Parallel()
	assert.Equal(t, "exchange:USD:EUR", cacheKey("usd", "eur"), "cache key must be uppercase")
	assert.Equal(t, "exchange:USD:EUR", cacheKey("USD", "EUR"), "cache key must be uppercase")
}
