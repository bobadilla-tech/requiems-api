package transactionrisk

import (
	"context"
	"encoding/json"
	"net"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	ipinfo "requiems-api/services/networking/ip/info"
	ipvpn "requiems-api/services/networking/ip/vpn"
	"requiems-api/platform/httpx"
	"requiems-api/services/finance/bin"
)

// -- stubs -------------------------------------------------------------------

type stubBIN struct {
	r   bin.LookupResponse
	err error
}

func (s *stubBIN) Lookup(_ context.Context, _ string) (bin.LookupResponse, error) {
	return s.r, s.err
}

type stubVPN struct {
	r   ipvpn.IPCheckResponse
	err error
}

func (s *stubVPN) CheckIP(_ context.Context, _ net.IP) (ipvpn.IPCheckResponse, error) {
	return s.r, s.err
}

type stubIPInfo struct {
	r   ipinfo.LookupResponse
	err error
}

func (s *stubIPInfo) CheckInfo(_ context.Context, _ string) (ipinfo.LookupResponse, error) {
	return s.r, s.err
}

// -- helpers -----------------------------------------------------------------

func setupRouter(b *stubBIN, v *stubVPN, i *stubIPInfo) chi.Router {
	r := chi.NewRouter()
	RegisterRoutes(r, NewService(b, v, i))
	return r
}

func post(t *testing.T, r chi.Router, body string) *httptest.ResponseRecorder {
	t.Helper()
	req := httptest.NewRequest(http.MethodPost, "/transaction/risk", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

// -- tests -------------------------------------------------------------------

func TestTransactionRisk_CleanTransaction(t *testing.T) {
	t.Parallel()
	r := setupRouter(
		&stubBIN{r: bin.LookupResponse{CountryCode: "US"}},
		&stubVPN{r: ipvpn.IPCheckResponse{}},
		&stubIPInfo{r: ipinfo.LookupResponse{CountryCode: "US"}},
	)
	w := post(t, r, `{"card_bin":"424242","ip_address":"1.2.3.4","billing_country":"US"}`)
	require.Equal(t, http.StatusOK, w.Code)
	var resp httpx.Response[Result]
	require.NoError(t, json.NewDecoder(w.Body).Decode(&resp))
	assert.True(t, resp.Data.IsSafe)
	assert.Less(t, resp.Data.RiskScore, 0.3)
}

func TestTransactionRisk_IPCountryMismatch(t *testing.T) {
	t.Parallel()
	r := setupRouter(
		&stubBIN{r: bin.LookupResponse{CountryCode: "US"}},
		&stubVPN{r: ipvpn.IPCheckResponse{}},
		&stubIPInfo{r: ipinfo.LookupResponse{CountryCode: "RU"}},
	)
	w := post(t, r, `{"card_bin":"424242","ip_address":"1.2.3.4","billing_country":"US"}`)
	require.Equal(t, http.StatusOK, w.Code)
	var resp httpx.Response[Result]
	require.NoError(t, json.NewDecoder(w.Body).Decode(&resp))
	assert.Contains(t, resp.Data.Flags, "country_mismatch")
	assert.GreaterOrEqual(t, resp.Data.RiskScore, 0.35)
}

func TestTransactionRisk_VPNDetected(t *testing.T) {
	t.Parallel()
	r := setupRouter(
		&stubBIN{r: bin.LookupResponse{CountryCode: "US"}},
		&stubVPN{r: ipvpn.IPCheckResponse{IsVPN: true}},
		&stubIPInfo{r: ipinfo.LookupResponse{CountryCode: "US"}},
	)
	w := post(t, r, `{"card_bin":"424242","ip_address":"1.2.3.4","billing_country":"US"}`)
	require.Equal(t, http.StatusOK, w.Code)
	var resp httpx.Response[Result]
	require.NoError(t, json.NewDecoder(w.Body).Decode(&resp))
	assert.Contains(t, resp.Data.Flags, "vpn_detected")
	assert.GreaterOrEqual(t, resp.Data.RiskScore, 0.20)
}

func TestTransactionRisk_TORDetected(t *testing.T) {
	t.Parallel()
	r := setupRouter(
		&stubBIN{r: bin.LookupResponse{CountryCode: "US"}},
		&stubVPN{r: ipvpn.IPCheckResponse{IsTor: true}},
		&stubIPInfo{r: ipinfo.LookupResponse{CountryCode: "US"}},
	)
	w := post(t, r, `{"card_bin":"424242","ip_address":"1.2.3.4","billing_country":"US"}`)
	require.Equal(t, http.StatusOK, w.Code)
	var resp httpx.Response[Result]
	require.NoError(t, json.NewDecoder(w.Body).Decode(&resp))
	assert.False(t, resp.Data.IsSafe)
	assert.Contains(t, resp.Data.Flags, "tor_detected")
	assert.GreaterOrEqual(t, resp.Data.RiskScore, 0.40)
}

func TestTransactionRisk_HighValueVPN(t *testing.T) {
	t.Parallel()
	amount := 600.0
	r := setupRouter(
		&stubBIN{r: bin.LookupResponse{CountryCode: "US"}},
		&stubVPN{r: ipvpn.IPCheckResponse{IsVPN: true}},
		&stubIPInfo{r: ipinfo.LookupResponse{CountryCode: "US"}},
	)
	body := `{"card_bin":"424242","ip_address":"1.2.3.4","billing_country":"US","amount_usd":600.0}`
	_ = amount
	w := post(t, r, body)
	require.Equal(t, http.StatusOK, w.Code)
	var resp httpx.Response[Result]
	require.NoError(t, json.NewDecoder(w.Body).Decode(&resp))
	assert.Contains(t, resp.Data.Flags, "high_value_vpn")
}

func TestTransactionRisk_BINCountryMismatch(t *testing.T) {
	t.Parallel()
	r := setupRouter(
		&stubBIN{r: bin.LookupResponse{CountryCode: "GB"}},
		&stubVPN{r: ipvpn.IPCheckResponse{}},
		&stubIPInfo{r: ipinfo.LookupResponse{CountryCode: "US"}},
	)
	w := post(t, r, `{"card_bin":"424242","ip_address":"1.2.3.4","billing_country":"US"}`)
	require.Equal(t, http.StatusOK, w.Code)
	var resp httpx.Response[Result]
	require.NoError(t, json.NewDecoder(w.Body).Decode(&resp))
	assert.Contains(t, resp.Data.Flags, "bin_country_mismatch")
}

func TestTransactionRisk_MissingCardBIN(t *testing.T) {
	t.Parallel()
	r := setupRouter(&stubBIN{}, &stubVPN{}, &stubIPInfo{})
	w := post(t, r, `{"ip_address":"1.2.3.4"}`)
	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
}

func TestTransactionRisk_MissingIPAddress(t *testing.T) {
	t.Parallel()
	r := setupRouter(&stubBIN{}, &stubVPN{}, &stubIPInfo{})
	w := post(t, r, `{"card_bin":"424242"}`)
	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
}
