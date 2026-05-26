package paymentvalidate

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"requiems-api/platform/httpx"
	"requiems-api/services/finance/bin"
	"requiems-api/services/finance/iban"
	"requiems-api/services/finance/swift"
)

type stubBIN struct {
	r   bin.LookupResponse
	err error
}

func (s *stubBIN) Lookup(_ context.Context, _ string) (bin.LookupResponse, error) {
	return s.r, s.err
}

type stubIBAN struct {
	r   iban.ParseResponse
	err error
}

func (s *stubIBAN) Parse(_ context.Context, _ string) (iban.ParseResponse, error) {
	return s.r, s.err
}

type stubSWIFT struct {
	r   swift.LookupResponse
	err error
}

func (s *stubSWIFT) Lookup(_ context.Context, _ string) (swift.LookupResponse, error) {
	return s.r, s.err
}

func setupRouter(b *stubBIN, i *stubIBAN, s *stubSWIFT) chi.Router {
	r := chi.NewRouter()
	RegisterRoutes(r, NewService(b, i, s))
	return r
}

func post(t *testing.T, r chi.Router, body string) *httptest.ResponseRecorder {
	t.Helper()
	req := httptest.NewRequest(http.MethodPost, "/payment/validate", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func usBIN() bin.LookupResponse {
	return bin.LookupResponse{
		BIN: "424242", Scheme: "visa", CardType: "credit", CardLevel: "classic",
		IssuerName: "Chase", CountryCode: "US", Prepaid: false, Luhn: true,
	}
}

func gbIBAN() iban.ParseResponse {
	return iban.ParseResponse{
		IBAN: "GB29NWBK60161331926819", Valid: true, Country: "United Kingdom",
		BankCode: "NWBK", Account: "31926819",
	}
}

func gbSWIFT() swift.LookupResponse {
	return swift.LookupResponse{
		SwiftCode: "NWBKGB2LXXX", BankCode: "NWBK", CountryCode: "GB",
		BankName: "NatWest", City: "London",
	}
}

func TestPaymentValidate_ValidBINOnly(t *testing.T) {
	t.Parallel()
	r := setupRouter(&stubBIN{r: usBIN()}, &stubIBAN{}, &stubSWIFT{})
	w := post(t, r, `{"bin":"424242"}`)
	require.Equal(t, http.StatusOK, w.Code)
	var resp httpx.Response[Result]
	require.NoError(t, json.NewDecoder(w.Body).Decode(&resp))
	require.NotNil(t, resp.Data.BIN)
	assert.True(t, resp.Data.BIN.Valid)
	assert.Nil(t, resp.Data.IBAN)
	assert.Nil(t, resp.Data.SWIFT)
	assert.True(t, resp.Data.Consistency.OK)
	assert.Empty(t, resp.Data.Consistency.Flags)
}

func TestPaymentValidate_BINandIBANSameCountry(t *testing.T) {
	t.Parallel()
	gbBIN := bin.LookupResponse{
		BIN: "400000", Scheme: "visa", CountryCode: "GB", Luhn: true,
	}
	r := setupRouter(&stubBIN{r: gbBIN}, &stubIBAN{r: gbIBAN()}, &stubSWIFT{})
	w := post(t, r, `{"bin":"400000","iban":"GB29NWBK60161331926819"}`)
	require.Equal(t, http.StatusOK, w.Code)
	var resp httpx.Response[Result]
	require.NoError(t, json.NewDecoder(w.Body).Decode(&resp))
	assert.True(t, resp.Data.Consistency.OK)
	assert.Empty(t, resp.Data.Consistency.Flags)
}

func TestPaymentValidate_BINUSandIBANGB_Mismatch(t *testing.T) {
	t.Parallel()
	r := setupRouter(&stubBIN{r: usBIN()}, &stubIBAN{r: gbIBAN()}, &stubSWIFT{})
	w := post(t, r, `{"bin":"424242","iban":"GB29NWBK60161331926819"}`)
	require.Equal(t, http.StatusOK, w.Code)
	var resp httpx.Response[Result]
	require.NoError(t, json.NewDecoder(w.Body).Decode(&resp))
	assert.False(t, resp.Data.Consistency.OK)
	assert.Contains(t, resp.Data.Consistency.Flags, "country_mismatch_bin_iban")
}

func TestPaymentValidate_AllThreeConsistent(t *testing.T) {
	t.Parallel()
	r := setupRouter(&stubBIN{r: bin.LookupResponse{BIN: "400000", CountryCode: "GB", Luhn: true}},
		&stubIBAN{r: gbIBAN()}, &stubSWIFT{r: gbSWIFT()})
	w := post(t, r, `{"bin":"400000","iban":"GB29NWBK60161331926819","swift":"NWBKGB2L"}`)
	require.Equal(t, http.StatusOK, w.Code)
	var resp httpx.Response[Result]
	require.NoError(t, json.NewDecoder(w.Body).Decode(&resp))
	assert.True(t, resp.Data.Consistency.OK)
}

func TestPaymentValidate_NoInstrumentsProvided(t *testing.T) {
	t.Parallel()
	r := setupRouter(&stubBIN{}, &stubIBAN{}, &stubSWIFT{})
	w := post(t, r, `{}`)
	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
}
