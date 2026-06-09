package cryptocoin

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

func setupRouter(svc *Service) chi.Router {
	r := chi.NewRouter()
	RegisterRoutes(r, svc)
	return r
}

func TestCrypto_GetPrice_ValidSymbol(t *testing.T) {
	t.Parallel()
	upstream := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
	defer upstream.Close()

	svc := newServiceWithClient(upstream.Client(), upstream.URL)
	r := setupRouter(svc)

	req := httptest.NewRequest(http.MethodGet, "/crypto/BTC", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code, "expected 200, got %d: %s", w.Code, w.Body.String())

	var resp httpx.Response[Price]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	assert.Equal(t, "BTC", resp.Data.Symbol)
	assert.Equal(t, "Bitcoin", resp.Data.Name)
	assert.Equal(t, 42000.50, resp.Data.PriceUSD)
}

func TestCrypto_GetPrice_UppercaseNormalization(t *testing.T) {
	t.Parallel()
	upstream := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body := coinGeckoResponse{
			"bitcoin": {USD: 42000.50},
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(body)
	}))
	defer upstream.Close()

	svc := newServiceWithClient(upstream.Client(), upstream.URL)
	r := setupRouter(svc)

	// lowercase symbol should be normalized
	req := httptest.NewRequest(http.MethodGet, "/crypto/btc", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code, "expected 200, got %d: %s", w.Code, w.Body.String())
}

func TestCrypto_GetPrice_UnknownSymbol(t *testing.T) {
	t.Parallel()
	svc := newServiceWithClient(http.DefaultClient, "http://unused")
	r := setupRouter(svc)

	req := httptest.NewRequest(http.MethodGet, "/crypto/FAKE", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
}

func TestCrypto_GetPriceBatch_ValidSymbol(t *testing.T) {
	t.Parallel()
	upstream := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body := coinGeckoResponse{
			"bitcoin": {USD: 42000.50, USD24hChange: 2.5, USDMarketCap: 820000000000, USD24hVol: 25000000000},
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(body)
	}))
	defer upstream.Close()

	svc := newServiceWithClient(upstream.Client(), upstream.URL)
	r := setupRouter(svc)

	req := httptest.NewRequest(http.MethodPost, "/crypto/batch", strings.NewReader(`{"symbols":["BTC"]}`))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)
	var resp httpx.Response[httpx.BatchResponse[BatchPriceItem]]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)
	assert.Len(t, resp.Data.Results, 1)
	assert.Equal(t, 1, resp.Data.Total)
	assert.Equal(t, "BTC", resp.Data.Results[0].Symbol)
	assert.NotNil(t, resp.Data.Results[0].Result)
	assert.Empty(t, resp.Data.Results[0].Error)
}

func TestCrypto_GetPriceBatch_UnsupportedSymbol_InBandError(t *testing.T) {
	t.Parallel()
	svc := newServiceWithClient(http.DefaultClient, "http://127.0.0.1:0")
	r := setupRouter(svc)

	req := httptest.NewRequest(http.MethodPost, "/crypto/batch", strings.NewReader(`{"symbols":["FAKE","NOTREAL"]}`))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)
	var resp httpx.Response[httpx.BatchResponse[BatchPriceItem]]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)
	assert.Len(t, resp.Data.Results, 2)
	for _, item := range resp.Data.Results {
		assert.NotEmpty(t, item.Error)
		assert.Nil(t, item.Result)
	}
}

func TestCrypto_GetPriceBatch_EmptySymbols422(t *testing.T) {
	t.Parallel()
	svc := newServiceWithClient(http.DefaultClient, "http://127.0.0.1:0")
	r := setupRouter(svc)

	req := httptest.NewRequest(http.MethodPost, "/crypto/batch", strings.NewReader(`{"symbols":[]}`))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
}

func TestCrypto_GetPrice_UpstreamDown(t *testing.T) {
	t.Parallel()
	upstream := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusServiceUnavailable)
	}))
	defer upstream.Close()

	svc := newServiceWithClient(upstream.Client(), upstream.URL)
	r := setupRouter(svc)

	req := httptest.NewRequest(http.MethodGet, "/crypto/ETH", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusServiceUnavailable, w.Code)
}
