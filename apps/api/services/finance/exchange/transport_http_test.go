package exchange

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"requiems-api/platform/httpx"
	"requiems-api/platform/svcerr"
)

type stubFetcher struct {
	fn func(ctx context.Context, from, to string) (float64, time.Time, error)
}

func (s *stubFetcher) GetRate(ctx context.Context, from, to string) (float64, time.Time, error) {
	return s.fn(ctx, from, to)
}

func fixedRate(rate float64) *stubFetcher {
	ts := time.Date(2024, 12, 15, 0, 0, 0, 0, time.UTC)
	return &stubFetcher{fn: func(_ context.Context, _, _ string) (float64, time.Time, error) {
		return rate, ts, nil
	}}
}

func errFetcher(err error) *stubFetcher {
	return &stubFetcher{fn: func(_ context.Context, _, _ string) (float64, time.Time, error) {
		return 0, time.Time{}, err
	}}
}

func setupRouter(f Fetcher) chi.Router {
	r := chi.NewRouter()
	registerExchangeRoutes(r, f)
	return r
}

// — /exchange-rate tests —

func TestExchangeRate_HappyPath(t *testing.T) {
	t.Parallel()
	r := setupRouter(fixedRate(0.92))

	req := httptest.NewRequest(http.MethodGet, "/exchange-rate?from=USD&to=EUR", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code, "expected 200, got %d: %s", w.Code, w.Body.String())

	var resp httpx.Response[RateResponse]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)
	assert.Equal(t, "USD", resp.Data.From)
	assert.Equal(t, "EUR", resp.Data.To)
	assert.Equal(t, 0.92, resp.Data.Rate)
	assert.NotEmpty(t, resp.Data.Timestamp, "timestamp must not be empty")
	assert.NotEmpty(t, resp.Metadata.Timestamp, "metadata.timestamp must not be empty")
}

func TestExchangeRate_LowercaseCodes_Normalized(t *testing.T) {
	t.Parallel()
	var gotFrom, gotTo string
	f := &stubFetcher{fn: func(_ context.Context, from, to string) (float64, time.Time, error) {
		gotFrom, gotTo = from, to
		return 0.92, time.Now(), nil
	}}
	r := setupRouter(f)

	req := httptest.NewRequest(http.MethodGet, "/exchange-rate?from=usd&to=eur", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code, "expected 200, got %d", w.Code)
	assert.Equal(t, "USD", gotFrom)
	assert.Equal(t, "EUR", gotTo)
}

func TestExchangeRate_MissingFrom_Returns400(t *testing.T) {
	t.Parallel()
	r := setupRouter(fixedRate(0.92))

	req := httptest.NewRequest(http.MethodGet, "/exchange-rate?to=EUR", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestExchangeRate_MissingTo_Returns400(t *testing.T) {
	t.Parallel()
	r := setupRouter(fixedRate(0.92))

	req := httptest.NewRequest(http.MethodGet, "/exchange-rate?from=USD", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestExchangeRate_InvalidCurrencyCode_Returns400(t *testing.T) {
	t.Parallel()
	r := setupRouter(fixedRate(0.92))

	req := httptest.NewRequest(http.MethodGet, "/exchange-rate?from=US&to=EUR", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code, "expected 400 for 2-char code, got %d", w.Code)
}

func TestExchangeRate_UnknownCurrency_Returns422(t *testing.T) {
	t.Parallel()
	svcErr := svcerr.Unknown("invalid_currency", "unknown currency code: XYZ")
	r := setupRouter(errFetcher(svcErr))

	req := httptest.NewRequest(http.MethodGet, "/exchange-rate?from=XYZ&to=EUR", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)

	var errResp httpx.ErrorResponse
	err := json.NewDecoder(w.Body).Decode(&errResp)
	require.NoError(t, err)
	assert.Equal(t, "invalid_currency", errResp.Error)
}

func TestExchangeRate_UpstreamError_Returns503(t *testing.T) {
	t.Parallel()
	r := setupRouter(errFetcher(errors.New("connection refused")))

	req := httptest.NewRequest(http.MethodGet, "/exchange-rate?from=USD&to=EUR", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusServiceUnavailable, w.Code)
}

// — /convert tests —

func TestConvert_HappyPath(t *testing.T) {
	t.Parallel()
	r := setupRouter(fixedRate(0.92))

	req := httptest.NewRequest(http.MethodGet, "/convert?from=USD&to=EUR&amount=100", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code, "expected 200, got %d: %s", w.Code, w.Body.String())

	var resp httpx.Response[ConvertResponse]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)
	assert.Equal(t, float64(100), resp.Data.Amount)
	assert.Equal(t, 92.00, resp.Data.Converted)
	assert.Equal(t, 0.92, resp.Data.Rate)
}

func TestConvert_MissingAmount_Returns400(t *testing.T) {
	t.Parallel()
	r := setupRouter(fixedRate(0.92))

	req := httptest.NewRequest(http.MethodGet, "/convert?from=USD&to=EUR", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestConvert_ZeroAmount_Returns400(t *testing.T) {
	t.Parallel()
	r := setupRouter(fixedRate(0.92))

	req := httptest.NewRequest(http.MethodGet, "/convert?from=USD&to=EUR&amount=0", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code, "expected 400 for amount=0, got %d", w.Code)
}

func TestConvert_UnknownCurrency_Returns422(t *testing.T) {
	t.Parallel()
	svcErr := svcerr.Unknown("invalid_currency", "unknown currency code: XYZ")
	r := setupRouter(errFetcher(svcErr))

	req := httptest.NewRequest(http.MethodGet, "/convert?from=XYZ&to=EUR&amount=50", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
}

func TestConvert_UpstreamError_Returns503(t *testing.T) {
	t.Parallel()
	r := setupRouter(errFetcher(errors.New("timeout")))

	req := httptest.NewRequest(http.MethodGet, "/convert?from=USD&to=EUR&amount=100", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusServiceUnavailable, w.Code)
}

func TestConvert_ConversionRounding(t *testing.T) {
	t.Parallel()
	r := setupRouter(fixedRate(0.9205))

	req := httptest.NewRequest(http.MethodGet, "/convert?from=USD&to=EUR&amount=100", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code, "expected 200, got %d", w.Code)

	var resp httpx.Response[ConvertResponse]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)
	assert.Equal(t, 92.05, resp.Data.Converted)
}
