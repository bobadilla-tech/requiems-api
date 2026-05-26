package userverify

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

	"requiems-api/platform/httpx"
	"requiems-api/services/networking/domain"
	ipvpn "requiems-api/services/networking/ip/vpn"
	"requiems-api/services/networking/mx"
	"requiems-api/services/networking/whois"
	"requiems-api/services/validation/email"
)

// -- stubs -------------------------------------------------------------------

type stubEmail struct{ r email.Validation }

func (s *stubEmail) ValidateEmail(_ context.Context, _ string) email.Validation { return s.r }

type stubWHOIS struct {
	r   whois.LookupResponse
	err error
}

func (s *stubWHOIS) Lookup(_ context.Context, _ string) (whois.LookupResponse, error) {
	return s.r, s.err
}

type stubDomain struct{ r domain.InfoResponse }

func (s *stubDomain) GetInfo(_ context.Context, _ string) domain.InfoResponse { return s.r }

type stubMX struct {
	r   mx.LookupResponse
	err error
}

func (s *stubMX) Lookup(_ context.Context, _ string) (mx.LookupResponse, error) {
	return s.r, s.err
}

type stubVPN struct {
	r   ipvpn.IPCheckResponse
	err error
}

func (s *stubVPN) CheckIP(_ context.Context, _ net.IP) (ipvpn.IPCheckResponse, error) {
	return s.r, s.err
}

// -- helpers -----------------------------------------------------------------

func validEmail() email.Validation {
	norm := "alice@example.com"
	dom := "example.com"
	return email.Validation{Valid: true, SyntaxValid: true, MxValid: true, Disposable: false, Normalized: &norm, Domain: &dom}
}

func goodDomain() domain.InfoResponse {
	return domain.InfoResponse{
		Domain:    "example.com",
		Available: false,
		DNS: domain.DNSRecords{
			A:   []string{"93.184.216.34"},
			MX:  []domain.MXRecord{{Host: "mail.example.com", Priority: 10}},
			NS:  []string{"ns1.example.com"},
			TXT: []string{},
		},
	}
}

func goodMX() mx.LookupResponse {
	return mx.LookupResponse{
		Domain:  "example.com",
		Records: []mx.Record{{Host: "mail.example.com", Priority: 10}},
	}
}

func oldWHOIS() whois.LookupResponse {
	return whois.LookupResponse{
		Domain:      "example.com",
		CreatedDate: "1995-08-14T04:00:00Z",
	}
}

func setupRouter(e *stubEmail, w *stubWHOIS, d *stubDomain, m *stubMX, v *stubVPN) chi.Router {
	r := chi.NewRouter()
	RegisterRoutes(r, NewService(e, w, d, m, v))
	return r
}

func post(t *testing.T, r chi.Router, body string) *httptest.ResponseRecorder {
	t.Helper()
	req := httptest.NewRequest(http.MethodPost, "/user/verify", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

// -- tests -------------------------------------------------------------------

func TestUserVerify_ValidEmailEstablishedDomain(t *testing.T) {
	t.Parallel()
	r := setupRouter(
		&stubEmail{r: validEmail()},
		&stubWHOIS{r: oldWHOIS()},
		&stubDomain{r: goodDomain()},
		&stubMX{r: goodMX()},
		&stubVPN{},
	)
	w := post(t, r, `{"email":"alice@example.com"}`)
	require.Equal(t, http.StatusOK, w.Code)
	var resp httpx.Response[Result]
	require.NoError(t, json.NewDecoder(w.Body).Decode(&resp))
	assert.True(t, resp.Data.Verified)
	assert.Less(t, resp.Data.RiskScore, 0.3)
}

func TestUserVerify_DisposableEmail(t *testing.T) {
	t.Parallel()
	norm := "user@tempmail.io"
	dom := "tempmail.io"
	r := setupRouter(
		&stubEmail{r: email.Validation{Valid: true, SyntaxValid: true, MxValid: true, Disposable: true, Normalized: &norm, Domain: &dom}},
		&stubWHOIS{r: oldWHOIS()},
		&stubDomain{r: goodDomain()},
		&stubMX{r: goodMX()},
		&stubVPN{},
	)
	w := post(t, r, `{"email":"user@tempmail.io"}`)
	require.Equal(t, http.StatusOK, w.Code)
	var resp httpx.Response[Result]
	require.NoError(t, json.NewDecoder(w.Body).Decode(&resp))
	assert.False(t, resp.Data.Verified)
	assert.Contains(t, resp.Data.Flags, "disposable_email")
}

func TestUserVerify_DomainNotRegistered(t *testing.T) {
	t.Parallel()
	r := setupRouter(
		&stubEmail{r: validEmail()},
		&stubWHOIS{r: whois.LookupResponse{}},
		&stubDomain{r: domain.InfoResponse{Available: true, DNS: domain.DNSRecords{MX: []domain.MXRecord{}}}},
		&stubMX{err: assert.AnError},
		&stubVPN{},
	)
	w := post(t, r, `{"email":"alice@example.com"}`)
	require.Equal(t, http.StatusOK, w.Code)
	var resp httpx.Response[Result]
	require.NoError(t, json.NewDecoder(w.Body).Decode(&resp))
	assert.False(t, resp.Data.Verified)
	assert.Contains(t, resp.Data.Flags, "domain_not_registered")
}

func TestUserVerify_WHOISFailure(t *testing.T) {
	t.Parallel()
	r := setupRouter(
		&stubEmail{r: validEmail()},
		&stubWHOIS{err: assert.AnError},
		&stubDomain{r: goodDomain()},
		&stubMX{r: goodMX()},
		&stubVPN{},
	)
	w := post(t, r, `{"email":"alice@example.com"}`)
	require.Equal(t, http.StatusOK, w.Code)
	var resp httpx.Response[Result]
	require.NoError(t, json.NewDecoder(w.Body).Decode(&resp))
	assert.Contains(t, resp.Data.Flags, "whois_unavailable")
}

func TestUserVerify_MissingEmail(t *testing.T) {
	t.Parallel()
	r := setupRouter(&stubEmail{}, &stubWHOIS{}, &stubDomain{}, &stubMX{}, &stubVPN{})
	w := post(t, r, `{}`)
	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
}
