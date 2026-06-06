package domaintrust

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"requiems-api/platform/httpx"
	"requiems-api/services/networking/domain"
	"requiems-api/services/networking/whois"
)

func setupRouter(whoIsSvc WhoIsService, domainSvc DomainService) chi.Router {
	r := chi.NewRouter()
	svc := NewService(whoIsSvc, domainSvc)
	RegisterRoutes(r, svc)
	return r
}

type mockDomainService struct {
	available bool
	noMx      bool
}

func (m *mockDomainService) GetInfo(ctx context.Context, domainName string) domain.InfoResponse {
	mx := []domain.MXRecord{{Host: "mail.example.com", Priority: 10}}
	if m.noMx {
		mx = []domain.MXRecord{}
	}

	return domain.InfoResponse{
		Available: m.available,
		DNS: domain.DNSRecords{
			A:  []string{"1.2.3.4"},
			MX: mx,
			NS: []string{"ns1.example.com"},
		},
	}
}

type mockWhoIsService struct {
	createdDate string
	expiryDate  string
}

func (m *mockWhoIsService) Lookup(_ context.Context, domainName string) (whois.LookupResponse, error) {
	return whois.LookupResponse{
		Domain:      domainName,
		Registrar:   "Mock Registrar",
		NameServers: []string{"ns1.example.com", "ns2.example.com"},
		Status:      []string{"clientTransferProhibited"},
		CreatedDate: m.createdDate,
		UpdatedDate: m.createdDate,
		ExpiryDate:  m.expiryDate,
		DNSSec:      false,
	}, nil
}

func Get(t *testing.T, r chi.Router, domainName string) *httptest.ResponseRecorder {
	t.Helper()

	req := httptest.NewRequest(http.MethodGet, "/domain/trust/"+domainName, http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestDomainTrust_HappyPath(t *testing.T) {
	t.Parallel()

	whoIsSvc := whois.NewService()
	domainSvc := domain.NewService()

	r := setupRouter(whoIsSvc, domainSvc)

	w := Get(t, r, "google.com")
	require.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[Response]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	assert.Equal(t, "high", resp.Data.TrustLevel)
	assert.Empty(t, resp.Data.Flags)

}

func TestDomainTrust_AgeDomain(t *testing.T) {
	t.Parallel()

	mockDomain := &mockDomainService{available: false}

	// new domain
	mockWhoIs := &mockWhoIsService{
		createdDate: time.Now().AddDate(0, 0, -10).Format(time.RFC3339), // 10 días
		expiryDate:  time.Now().AddDate(1, 0, 0).Format(time.RFC3339),
	}

	r := setupRouter(mockWhoIs, mockDomain)

	w := Get(t, r, "newdomain.com")
	require.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[Response]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	assert.Contains(t, resp.Data.Flags, "new_domain")
	assert.LessOrEqual(t, resp.Data.TrustScore, 0.5)
	assert.Equal(t, "medium", resp.Data.TrustLevel)

}

func TestDomainTrust_NoMX(t *testing.T) {
	t.Parallel()

	mockDomain := &mockDomainService{
		available: false,
		noMx:      true,
	}
	mockWhoIs := &mockWhoIsService{
		createdDate: time.Now().AddDate(-2, 0, 0).Format(time.RFC3339),
		expiryDate:  time.Now().AddDate(1, 0, 0).Format(time.RFC3339),
	}

	r := setupRouter(mockWhoIs, mockDomain)
	w := Get(t, r, "nomx.com")

	require.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[Response]
	require.NoError(t, json.NewDecoder(w.Body).Decode(&resp))

	assert.Contains(t, resp.Data.Flags, "no_mx")
	assert.LessOrEqual(t, resp.Data.TrustScore, 0.65)
	assert.Empty(t, resp.Data.MxRecords)
}

func TestDomainTrust_NotRegistered(t *testing.T) {
	t.Parallel()

	mockDomain := &mockDomainService{available: true}
	mockWhoIs := &mockWhoIsService{}

	r := setupRouter(mockWhoIs, mockDomain)
	w := Get(t, r, "notregistered.com")

	require.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[Response]
	require.NoError(t, json.NewDecoder(w.Body).Decode(&resp))

	assert.Contains(t, resp.Data.Flags, "domain_not_registered")
	assert.Equal(t, 0.0, resp.Data.TrustScore)
	assert.Equal(t, "low", resp.Data.TrustLevel)
}

func TestDomainTrust_InvalidFormat(t *testing.T) {
	t.Parallel()

	r := setupRouter(&mockWhoIsService{}, &mockDomainService{})
	w := Get(t, r, "not_a_valid_domain")

	require.Equal(t, http.StatusBadRequest, w.Code)
}
