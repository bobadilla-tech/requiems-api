package iban

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"requiems-api/platform/httpx"
)

// stubValidator implements Validator for transport tests. It returns a fixed
// result or a fixed error on every call, keeping tests DB-free and fast.
type stubValidator struct {
	result  ParseResponse
	results BatchParseResponse
	err     error
}

func (s *stubValidator) Parse(_ context.Context, raw string) (ParseResponse, error) {
	if s.err != nil {
		return ParseResponse{}, s.err
	}
	r := s.result
	r.IBAN = raw
	return r, nil
}

func (s *stubValidator) ParseBatch(_ context.Context, numbers []string) (BatchParseResponse, error) {
	if s.err != nil {
		return BatchParseResponse{}, s.err
	}
	r := s.results
	return r, nil
}

func setupRouter(v Validator) chi.Router {
	r := chi.NewRouter()
	registerIBANRoutes(r, v)
	return r
}

func decodeResponse(t *testing.T, w *httptest.ResponseRecorder) httpx.Response[ParseResponse] {
	t.Helper()
	var resp httpx.Response[ParseResponse]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)
	return resp
}

// ---- tests ----

func TestIBAN_ValidDE_Returns200(t *testing.T) {
	t.Parallel()
	svc := &stubValidator{result: ParseResponse{
		Valid:    true,
		Country:  "Germany",
		BankCode: "37040044",
		Account:  "0532013000",
	}}
	r := setupRouter(svc)

	req := httptest.NewRequest(http.MethodGet, "/iban/DE89370400440532013000", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	resp := decodeResponse(t, w)
	assert.True(t, resp.Data.Valid)
	assert.Equal(t, "DE89370400440532013000", resp.Data.IBAN)
	assert.Equal(t, "Germany", resp.Data.Country)
	assert.Equal(t, "37040044", resp.Data.BankCode)
	assert.Equal(t, "0532013000", resp.Data.Account)
	assert.NotEmpty(t, resp.Metadata.Timestamp)
}

func TestIBAN_ResponseEnvelope(t *testing.T) {
	t.Parallel()
	svc := &stubValidator{result: ParseResponse{Valid: true, Country: "Netherlands"}}
	r := setupRouter(svc)

	req := httptest.NewRequest(http.MethodGet, "/iban/NL91ABNA0417164300", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	var raw map[string]json.RawMessage
	err := json.NewDecoder(w.Body).Decode(&raw)
	require.NoError(t, err)
	_, ok := raw["data"]
	assert.True(t, ok, "response must have a 'data' key")
	_, ok = raw["metadata"]
	assert.True(t, ok, "response must have a 'metadata' key")
}

func TestIBAN_InvalidChecksum_Returns200WithValidFalse(t *testing.T) {
	t.Parallel()
	svc := &stubValidator{result: ParseResponse{Valid: false, Country: "Germany"}}
	r := setupRouter(svc)

	req := httptest.NewRequest(http.MethodGet, "/iban/DE00370400440532013000", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	resp := decodeResponse(t, w)
	assert.False(t, resp.Data.Valid)
}

func TestIBAN_DBError_Returns500(t *testing.T) {
	t.Parallel()
	svc := &stubValidator{err: errors.New("db unavailable")}
	r := setupRouter(svc)

	req := httptest.NewRequest(http.MethodGet, "/iban/DE89370400440532013000", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestIBAN_UnknownCountry_Returns200(t *testing.T) {
	t.Parallel()
	// IBAN from country not in DB — valid checksum, empty bank_code/account.
	svc := &stubValidator{result: ParseResponse{Valid: true}}
	r := setupRouter(svc)

	req := httptest.NewRequest(http.MethodGet, "/iban/XX00TEST12345678", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestIBAN_AllResponseFieldsPresent(t *testing.T) {
	t.Parallel()
	svc := &stubValidator{result: ParseResponse{
		Valid:    true,
		Country:  "Netherlands",
		BankCode: "ABNA",
		Account:  "0417164300",
	}}
	r := setupRouter(svc)

	req := httptest.NewRequest(http.MethodGet, "/iban/NL91ABNA0417164300", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	resp := decodeResponse(t, w)
	d := resp.Data

	checks := map[string]bool{
		"iban non-empty":      d.IBAN != "",
		"country non-empty":   d.Country != "",
		"bank_code non-empty": d.BankCode != "",
		"account non-empty":   d.Account != "",
		"valid is true":       d.Valid,
	}
	for name, ok := range checks {
		assert.True(t, ok, "field check failed: %s", name)
	}
}

func TestIBAN_GBParsing_Returns200(t *testing.T) {
	t.Parallel()
	svc := &stubValidator{result: ParseResponse{
		Valid:    true,
		Country:  "United Kingdom",
		BankCode: "WEST",
		Account:  "98765432",
	}}
	r := setupRouter(svc)

	req := httptest.NewRequest(http.MethodGet, "/iban/GB82WEST12345698765432", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	resp := decodeResponse(t, w)
	assert.Equal(t, "WEST", resp.Data.BankCode)
	assert.Equal(t, "98765432", resp.Data.Account)
}

func TestIBAN_BatchParse(t *testing.T) {
	t.Parallel()
	svc := &stubValidator{results: BatchParseResponse{
		Results: []ParseResponse{
			{IBAN: "GB29NWBK60161331926819", Valid: true, Country: "United Kingdom", BankCode: "NWBK", Account: "31926819"},
			{IBAN: "DE89370400440532013000", Valid: true, Country: "Germany", BankCode: "37040044", Account: "0532013000"},
			{IBAN: "XX89370400440532013000", Valid: false},
		},
		Total: 3,
	}}

	r := setupRouter(svc)

	body := `{"numbers": ["GB29NWBK60161331926819","DE89370400440532013000","XX89370400440532013000"]}`
	req := httptest.NewRequest(http.MethodPost, "/iban/batch", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[BatchParseResponse]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	assert.Equal(t, 3, resp.Data.Total)

	require.Len(t, resp.Data.Results, 3)

	// valid UK IBAN
	assert.True(t, resp.Data.Results[0].Valid)
	assert.Equal(t, "United Kingdom", resp.Data.Results[0].Country)
	assert.Equal(t, "NWBK", resp.Data.Results[0].BankCode)

	// valid DE IBAN
	assert.True(t, resp.Data.Results[1].Valid)
	assert.Equal(t, "Germany", resp.Data.Results[1].Country)
	assert.Equal(t, "37040044", resp.Data.Results[1].BankCode)

	// invalid IBAN
	assert.False(t, resp.Data.Results[2].Valid)
}

func TestIBAN_BatchParse_EmptyBody(t *testing.T) {
	t.Parallel()
	svc := &stubValidator{}
	r := setupRouter(svc)

	req := httptest.NewRequest(http.MethodPost, "/iban/batch", strings.NewReader(`{"numbers": []}`))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
}

func TestIBAN_BatchParse_ExceedsLimmit(t *testing.T) {
	t.Parallel()
	svc := &stubValidator{}
	r := setupRouter(svc)

	numbers := make([]string, 51)

	for i := range numbers {
		numbers[i] = `"GB29NWBK60161331926819"`
	}

	body := `{"numbers":[` + strings.Join(numbers, ",") + `]}`
	req := httptest.NewRequest(http.MethodPost, "/iban/batch", strings.NewReader(body))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
}

func TestIBAN_BatchParse_MissingBody(t *testing.T) {
	t.Parallel()
	svc := &stubValidator{}
	r := setupRouter(svc)

	req := httptest.NewRequest(http.MethodPost, "/iban/batch", strings.NewReader(`{}`))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
}

func TestIBAN_BatchParse_SetsUsageCountHeader(t *testing.T) {
	t.Parallel()
	svc := &stubValidator{}
	r := setupRouter(svc)

	body := `{"numbers": ["GB29NWBK60161331926819","DE89370400440532013000","XX89370400440532013000"]}`
	req := httptest.NewRequest(http.MethodPost, "/iban/batch", strings.NewReader(body))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "3", w.Header().Get("X-Usage-Count"))
}
