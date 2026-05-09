package cryptocoin

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"requiems-api/platform/httpx"
)

func TestGetPrice_ValidSymbol(t *testing.T) {
	t.Parallel()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body := coinGeckoResponse{
			"bitcoin": {
				USD:          42000.50,
				USD24hChange: 2.5,
				USDMarketCap: 820000000000,
				USD24hVol:    25000000000,
			},
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(body)
	}))
	defer srv.Close()

	svc := newServiceWithClient(srv.Client(), srv.URL)
	p, err := svc.GetPrice(context.Background(), "BTC")
	require.NoError(t, err)

	assert.Equal(t, "BTC", p.Symbol)
	assert.Equal(t, "Bitcoin", p.Name)
	assert.Equal(t, 42000.50, p.PriceUSD)
	assert.Equal(t, 2.5, p.Change24h)
}

func TestGetPrice_UnknownSymbol(t *testing.T) {
	t.Parallel()
	svc := newServiceWithClient(http.DefaultClient, "http://unused")
	_, err := svc.GetPrice(context.Background(), "FAKE")
	require.Error(t, err)

	var ae *httpx.AppError
	require.ErrorAs(t, err, &ae)
	assert.Equal(t, "unknown_symbol", ae.Code)
	assert.Equal(t, http.StatusUnprocessableEntity, ae.Status)
}

func TestGetPrice_UpstreamError(t *testing.T) {
	t.Parallel()
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer srv.Close()

	svc := newServiceWithClient(srv.Client(), srv.URL)
	_, err := svc.GetPrice(context.Background(), "BTC")
	require.Error(t, err)

	var ae *httpx.AppError
	require.ErrorAs(t, err, &ae)
	assert.Equal(t, "upstream_error", ae.Code)
}

func TestGetPrice_NoRedis_CallsUpstream(t *testing.T) {
	t.Parallel()
	callCount := 0
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		callCount++
		body := coinGeckoResponse{
			"bitcoin": {USD: 50000},
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(body)
	}))
	defer srv.Close()

	svc := newServiceWithClient(srv.Client(), srv.URL)

	for i := 0; i < 2; i++ {
		_, err := svc.GetPrice(context.Background(), "BTC")
		require.NoError(t, err, "call %d failed", i+1)
	}

	assert.Equal(t, 2, callCount, "expected 2 upstream calls (no Redis), got %d", callCount)
}
