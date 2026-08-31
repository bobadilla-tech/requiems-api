package signupprotect

import (
	"context"
	"encoding/json"
	"net"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/alicebob/miniredis/v2"
	"github.com/go-chi/chi/v5"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"requiems-api/platform/httpx"
	ipinfo "requiems-api/services/networking/ip/info"
	ipvpn "requiems-api/services/networking/ip/vpn"
	"requiems-api/services/validation/email"
	"requiems-api/services/validation/phone"
)

type stubEmail struct{ r email.Validation }

func (s *stubEmail) ValidateEmail(_ context.Context, _ string) email.Validation { return s.r }

type stubPhone struct{ r phone.ValidateResponse }

func (s *stubPhone) Validate(_ string) phone.ValidateResponse { return s.r }

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

func cleanEmail() email.Validation {
	v, mx := true, true
	normalized := "user@example.com"
	domain := "example.com"
	return email.Validation{Valid: v, SyntaxValid: true, MxValid: mx, Disposable: false, Normalized: &normalized, Domain: &domain}
}

func setupRouter(t *testing.T, e *stubEmail, p *stubPhone, v *stubVPN, i *stubIPInfo) chi.Router {
	t.Helper()
	mr := miniredis.RunT(t)
	rdb := redis.NewClient(&redis.Options{Addr: mr.Addr()})
	r := chi.NewRouter()
	RegisterRoutes(r, NewService(e, p, v, i, rdb))
	return r
}

func post(t *testing.T, r chi.Router, body string) *httptest.ResponseRecorder {
	t.Helper()
	req := httptest.NewRequest(http.MethodPost, "/signup/protect", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestSignupProtect_AllClean(t *testing.T) {
	t.Parallel()
	r := setupRouter(
		t,
		&stubEmail{r: cleanEmail()},
		&stubPhone{r: phone.ValidateResponse{Valid: true, Country: "US", Risk: &phone.Risk{}}},
		&stubVPN{r: ipvpn.IPCheckResponse{}},
		&stubIPInfo{r: ipinfo.LookupResponse{CountryCode: "US"}},
	)
	w := post(t, r, `{"email":"user@example.com","phone":"+14155550100","ip_address":"1.2.3.4"}`)
	require.Equal(t, http.StatusOK, w.Code)
	var resp httpx.Response[Result]
	require.NoError(t, json.NewDecoder(w.Body).Decode(&resp))
	assert.True(t, resp.Data.IsSafe)
	assert.Less(t, resp.Data.RiskScore, 0.5)
	assert.Empty(t, resp.Data.Flags)
	assert.NotNil(t, resp.Data.Signals.Email)
	assert.NotNil(t, resp.Data.Signals.Phone)
	assert.NotNil(t, resp.Data.Signals.IP)
}

func TestSignupProtect_DisposableEmail(t *testing.T) {
	t.Parallel()
	normalized := "user@tempmail.io"
	domain := "tempmail.io"
	r := setupRouter(
		t,
		&stubEmail{r: email.Validation{Valid: true, SyntaxValid: true, MxValid: true, Disposable: true, Normalized: &normalized, Domain: &domain}},
		&stubPhone{r: phone.ValidateResponse{Valid: true, Risk: &phone.Risk{}}},
		&stubVPN{r: ipvpn.IPCheckResponse{}},
		&stubIPInfo{r: ipinfo.LookupResponse{}},
	)
	w := post(t, r, `{"email":"user@tempmail.io","phone":"+14155550100","ip_address":"1.2.3.4"}`)
	require.Equal(t, http.StatusOK, w.Code)
	var resp httpx.Response[Result]
	require.NoError(t, json.NewDecoder(w.Body).Decode(&resp))
	assert.Contains(t, resp.Data.Flags, "disposable_email")
	assert.GreaterOrEqual(t, resp.Data.RiskScore, 0.30)
	require.NotNil(t, resp.Data.Signals.Email)
	assert.True(t, resp.Data.Signals.Email.Disposable)
}

func TestSignupProtect_TORDetected(t *testing.T) {
	t.Parallel()
	r := setupRouter(
		t,
		&stubEmail{r: cleanEmail()},
		&stubPhone{r: phone.ValidateResponse{Valid: true, Risk: &phone.Risk{}}},
		&stubVPN{r: ipvpn.IPCheckResponse{IsTor: true}},
		&stubIPInfo{r: ipinfo.LookupResponse{}},
	)
	w := post(t, r, `{"email":"user@example.com","phone":"+14155550100","ip_address":"1.2.3.4"}`)
	require.Equal(t, http.StatusOK, w.Code)
	var resp httpx.Response[Result]
	require.NoError(t, json.NewDecoder(w.Body).Decode(&resp))
	assert.False(t, resp.Data.IsSafe)
	assert.Contains(t, resp.Data.Flags, "tor_detected")
	assert.GreaterOrEqual(t, resp.Data.RiskScore, 0.40)
	require.NotNil(t, resp.Data.Signals.IP)
	assert.True(t, resp.Data.Signals.IP.IsTOR)
}

func TestSignupProtect_VoIPPhone(t *testing.T) {
	t.Parallel()
	r := setupRouter(
		t,
		&stubEmail{r: cleanEmail()},
		&stubPhone{r: phone.ValidateResponse{Valid: true, Risk: &phone.Risk{IsVoIP: true}}},
		&stubVPN{r: ipvpn.IPCheckResponse{}},
		&stubIPInfo{r: ipinfo.LookupResponse{}},
	)
	w := post(t, r, `{"email":"user@example.com","phone":"+14155550100","ip_address":"1.2.3.4"}`)
	require.Equal(t, http.StatusOK, w.Code)
	var resp httpx.Response[Result]
	require.NoError(t, json.NewDecoder(w.Body).Decode(&resp))
	assert.Contains(t, resp.Data.Flags, "phone_voip")
	require.NotNil(t, resp.Data.Signals.Phone)
	assert.True(t, resp.Data.Signals.Phone.IsVoIP)
}

func TestSignupProtect_GeoMismatchPhoneIP(t *testing.T) {
	t.Parallel()
	r := setupRouter(
		t,
		&stubEmail{r: cleanEmail()},
		&stubPhone{r: phone.ValidateResponse{Valid: true, Country: "US", Risk: &phone.Risk{}}},
		&stubVPN{r: ipvpn.IPCheckResponse{}},
		&stubIPInfo{r: ipinfo.LookupResponse{CountryCode: "NL"}},
	)
	w := post(t, r, `{"email":"user@example.com","phone":"+14155550100","ip_address":"1.2.3.4"}`)
	require.Equal(t, http.StatusOK, w.Code)
	var resp httpx.Response[Result]
	require.NoError(t, json.NewDecoder(w.Body).Decode(&resp))
	assert.Contains(t, resp.Data.Flags, "geo_mismatch_phone_ip")
}

func TestSignupProtect_OnlyEmailSignalsNullPhoneIP(t *testing.T) {
	t.Parallel()
	r := setupRouter(t, &stubEmail{r: cleanEmail()}, &stubPhone{}, &stubVPN{}, &stubIPInfo{})
	w := post(t, r, `{"email":"user@example.com"}`)
	require.Equal(t, http.StatusOK, w.Code)
	var resp httpx.Response[Result]
	require.NoError(t, json.NewDecoder(w.Body).Decode(&resp))
	assert.Nil(t, resp.Data.Signals.Phone)
	assert.Nil(t, resp.Data.Signals.IP)
	assert.Greater(t, resp.Data.Confidence, 0.70)
}

func TestSignupProtect_NoFieldsProvided(t *testing.T) {
	t.Parallel()
	r := setupRouter(t, &stubEmail{}, &stubPhone{}, &stubVPN{}, &stubIPInfo{})
	w := post(t, r, `{}`)
	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
}
