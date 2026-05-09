package bin

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"requiems-api/platform/httpx"
)

// stubService implements Looker for transport tests. It returns a fixed result
// or a fixed error on every call, keeping tests DB-free and fast.
type stubService struct {
	result LookupResponse
	err    error
}

func (s *stubService) Lookup(_ context.Context, bin string) (LookupResponse, error) {
	if s.err != nil {
		return LookupResponse{}, s.err
	}
	r := s.result
	r.BIN = bin
	return r, nil
}

// setupRouter wires up a stub service into a chi router for handler testing.
func setupRouter(svc Looker) chi.Router {
	r := chi.NewRouter()
	registerBINRoutes(r, svc)
	return r
}

// ---- helper ----

func decodeResponse(t *testing.T, body *httptest.ResponseRecorder) httpx.Response[LookupResponse] {
	t.Helper()
	var resp httpx.Response[LookupResponse]
	err := json.NewDecoder(body.Body).Decode(&resp)
	require.NoError(t, err)
	return resp
}

// ---- tests ----

func TestBINLookup_KnownBIN_Returns200(t *testing.T) {
	t.Parallel()
	svc := &stubService{result: LookupResponse{
		Scheme:      "visa",
		CardType:    "credit",
		CardLevel:   "classic",
		IssuerName:  "Chase",
		CountryCode: "US",
		CountryName: "United States",
		Prepaid:     false,
		Luhn:        true,
		Confidence:  0.92,
	}}

	r := setupRouter(svc)
	req := httptest.NewRequest(http.MethodGet, "/bin/424242", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code, "expected 200, got %d: %s", w.Code, w.Body.String())

	resp := decodeResponse(t, w)
	assert.Equal(t, "424242", resp.Data.BIN)
	assert.Equal(t, "visa", resp.Data.Scheme)
	assert.NotEmpty(t, resp.Metadata.Timestamp, "expected metadata.timestamp to be set")
}

func TestBINLookup_ResponseEnvelope(t *testing.T) {
	t.Parallel()
	svc := &stubService{result: LookupResponse{Scheme: "mastercard"}}
	r := setupRouter(svc)

	req := httptest.NewRequest(http.MethodGet, "/bin/510000", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code, "expected 200, got %d", w.Code)

	var raw map[string]json.RawMessage
	err := json.NewDecoder(w.Body).Decode(&raw)
	require.NoError(t, err)
	assert.Contains(t, raw, "data", "response must have a 'data' key")
	assert.Contains(t, raw, "metadata", "response must have a 'metadata' key")
}

func TestBINLookup_UnknownBIN_Returns404(t *testing.T) {
	t.Parallel()
	svc := &stubService{err: &httpx.AppError{
		Status:  http.StatusNotFound,
		Code:    "not_found",
		Message: "BIN not found",
	}}

	r := setupRouter(svc)
	req := httptest.NewRequest(http.MethodGet, "/bin/999999", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusNotFound, w.Code, "expected 404, got %d: %s", w.Code, w.Body.String())
}

func TestBINLookup_TooShort_Returns400(t *testing.T) {
	t.Parallel()
	svc := &stubService{err: &httpx.AppError{
		Status:  http.StatusBadRequest,
		Code:    "bad_request",
		Message: "BIN must be between 6 and 8 digits",
	}}

	r := setupRouter(svc)
	req := httptest.NewRequest(http.MethodGet, "/bin/4242", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusBadRequest, w.Code, "expected 400, got %d: %s", w.Code, w.Body.String())
}

func TestBINLookup_TooLong_Returns400(t *testing.T) {
	t.Parallel()
	svc := &stubService{err: &httpx.AppError{
		Status:  http.StatusBadRequest,
		Code:    "bad_request",
		Message: "BIN must be between 6 and 8 digits",
	}}

	r := setupRouter(svc)
	req := httptest.NewRequest(http.MethodGet, "/bin/424242424", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusBadRequest, w.Code, "expected 400, got %d: %s", w.Code, w.Body.String())
}

func TestBINLookup_NonDigits_Returns400(t *testing.T) {
	t.Parallel()
	svc := &stubService{err: &httpx.AppError{
		Status:  http.StatusBadRequest,
		Code:    "bad_request",
		Message: "BIN must contain digits only",
	}}

	r := setupRouter(svc)
	req := httptest.NewRequest(http.MethodGet, "/bin/abcdef", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusBadRequest, w.Code, "expected 400, got %d: %s", w.Code, w.Body.String())
}

func TestBINLookup_LuhnTrue(t *testing.T) {
	t.Parallel()
	svc := &stubService{result: LookupResponse{Luhn: true}}
	r := setupRouter(svc)

	req := httptest.NewRequest(http.MethodGet, "/bin/424242", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	resp := decodeResponse(t, w)
	assert.True(t, resp.Data.Luhn, "expected luhn = true")
}

func TestBINLookup_LuhnFalse(t *testing.T) {
	t.Parallel()
	svc := &stubService{result: LookupResponse{Luhn: false}}
	r := setupRouter(svc)

	req := httptest.NewRequest(http.MethodGet, "/bin/123456", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	resp := decodeResponse(t, w)
	assert.False(t, resp.Data.Luhn, "expected luhn = false")
}

func TestBINLookup_PrepaidTrue(t *testing.T) {
	t.Parallel()
	svc := &stubService{result: LookupResponse{Prepaid: true, CardType: "prepaid"}}
	r := setupRouter(svc)

	req := httptest.NewRequest(http.MethodGet, "/bin/630400", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	resp := decodeResponse(t, w)
	assert.True(t, resp.Data.Prepaid, "expected prepaid = true")
}

func TestBINLookup_ConfidenceField(t *testing.T) {
	t.Parallel()
	svc := &stubService{result: LookupResponse{Confidence: 0.87}}
	r := setupRouter(svc)

	req := httptest.NewRequest(http.MethodGet, "/bin/424242", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	resp := decodeResponse(t, w)
	assert.Equal(t, 0.87, resp.Data.Confidence)
}

func TestBINLookup_8DigitBIN(t *testing.T) {
	t.Parallel()
	svc := &stubService{result: LookupResponse{Scheme: "visa", CardType: "credit"}}
	r := setupRouter(svc)

	req := httptest.NewRequest(http.MethodGet, "/bin/42424242", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code, "expected 200 for 8-digit BIN, got %d: %s", w.Code, w.Body.String())

	resp := decodeResponse(t, w)
	assert.Equal(t, "42424242", resp.Data.BIN)
}

func TestBINLookup_AllResponseFieldsPresent(t *testing.T) {
	t.Parallel()
	svc := &stubService{result: LookupResponse{
		Scheme:      "mastercard",
		CardType:    "debit",
		CardLevel:   "gold",
		IssuerName:  "Bank of America",
		IssuerURL:   "www.bankofamerica.com",
		IssuerPhone: "+18004321000",
		CountryCode: "US",
		CountryName: "United States",
		Prepaid:     false,
		Luhn:        true,
		Confidence:  0.95,
	}}

	r := setupRouter(svc)
	req := httptest.NewRequest(http.MethodGet, "/bin/510000", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code, "expected 200, got %d", w.Code)

	resp := decodeResponse(t, w)
	d := resp.Data

	checks := map[string]bool{
		"scheme non-empty":       d.Scheme != "",
		"card_type non-empty":    d.CardType != "",
		"card_level non-empty":   d.CardLevel != "",
		"issuer_name non-empty":  d.IssuerName != "",
		"issuer_url non-empty":   d.IssuerURL != "",
		"issuer_phone non-empty": d.IssuerPhone != "",
		"country_code non-empty": d.CountryCode != "",
		"country_name non-empty": d.CountryName != "",
		"confidence > 0":         d.Confidence > 0,
	}

	for name, ok := range checks {
		assert.True(t, ok, "field check failed: %s", name)
	}
}
