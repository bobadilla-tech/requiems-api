package domain

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
)

type stubInfoBatcher struct {
	fn func([]string) []BatchDomainItem
}

func (s *stubInfoBatcher) GetInfoBatch(_ context.Context, domains []string) []BatchDomainItem {
	return s.fn(domains)
}

func setupBatchRouter(b InfoBatcher) chi.Router {
	r := chi.NewRouter()
	registerDomainBatchRoutes(r, b)
	return r
}

func setupRouter() chi.Router {
	r := chi.NewRouter()
	RegisterRoutes(r, NewService())
	return r
}

func TestDomain_InvalidFormat(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		domain string
	}{
		{"bare label", "localhost"},
		{"leading hyphen", "-bad.com"},
		{"trailing hyphen", "bad-.com"},
		{"numeric TLD only", "123"},
		{"just a dot", "."},
	}

	r := setupRouter()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			req := httptest.NewRequest(http.MethodGet, "/domain/"+tt.domain, http.NoBody)
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
		})
	}
}

func TestDomain_KnownDomain(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	req := httptest.NewRequest(http.MethodGet, "/domain/example.com", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[InfoResponse]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	assert.Equal(t, "example.com", resp.Data.Domain)
	assert.NotNil(t, resp.Data.DNS.A)
	assert.NotNil(t, resp.Data.DNS.AAAA)
	assert.NotNil(t, resp.Data.DNS.MX)
	assert.NotNil(t, resp.Data.DNS.NS)
	assert.NotNil(t, resp.Data.DNS.TXT)

	// DNS record content is only asserted when network resolution is available.
	if len(resp.Data.DNS.NS) > 0 && resp.Data.Available {
		assert.False(t, resp.Data.Available, "expected available=false when NS records are present")
	}
}

func TestDomain_ResponseShape(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	req := httptest.NewRequest(http.MethodGet, "/domain/example.com", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	// Verify the raw JSON shape has the expected keys.
	var raw map[string]any
	err := json.NewDecoder(w.Body).Decode(&raw)
	require.NoError(t, err)

	data, ok := raw["data"].(map[string]any)
	require.True(t, ok, "expected 'data' key in response")
	for _, key := range []string{"domain", "available", "dns"} {
		_, exists := data[key]
		assert.True(t, exists, "expected key %q in data", key)
	}

	dns, ok := data["dns"].(map[string]any)
	require.True(t, ok, "expected 'dns' key in data")
	for _, key := range []string{"a", "aaaa", "mx", "ns", "txt"} {
		_, exists := dns[key]
		assert.True(t, exists, "expected key %q in dns", key)
	}
}

func TestDomainBatch_EmptyDomains422(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	req := httptest.NewRequest(http.MethodPost, "/domain/batch", strings.NewReader(`{"domains":[]}`))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
}

func TestDomainBatch_InvalidDomain_InBandError(t *testing.T) {
	t.Parallel()
	stub := &stubInfoBatcher{fn: func(domains []string) []BatchDomainItem {
		results := make([]BatchDomainItem, len(domains))
		for i, d := range domains {
			results[i] = BatchDomainItem{Domain: d, Error: "invalid domain format"}
		}
		return results
	}}
	r := setupBatchRouter(stub)

	req := httptest.NewRequest(http.MethodPost, "/domain/batch", strings.NewReader(`{"domains":["not-a-domain","also-invalid"]}`))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)
	var resp httpx.Response[httpx.BatchResponse[BatchDomainItem]]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)
	assert.Len(t, resp.Data.Results, 2)
	assert.Equal(t, 2, resp.Data.Total)
	for _, item := range resp.Data.Results {
		assert.NotEmpty(t, item.Error)
	}
}

func TestDomainBatch_ResponseShape(t *testing.T) {
	t.Parallel()
	info := InfoResponse{Domain: "example.com", Available: false}
	stub := &stubInfoBatcher{fn: func(domains []string) []BatchDomainItem {
		return []BatchDomainItem{{Domain: "example.com", Result: &info}}
	}}
	r := setupBatchRouter(stub)

	req := httptest.NewRequest(http.MethodPost, "/domain/batch", strings.NewReader(`{"domains":["example.com"]}`))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)
	var resp httpx.Response[httpx.BatchResponse[BatchDomainItem]]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)
	assert.Len(t, resp.Data.Results, 1)
	assert.Equal(t, "example.com", resp.Data.Results[0].Domain)
	assert.NotNil(t, resp.Data.Results[0].Result)
	assert.Empty(t, resp.Data.Results[0].Error)
}

func TestService_IsNXDomain(t *testing.T) {
	t.Parallel()
	svc := NewService()

	// A clearly invented domain should either be unavailable (registered) or
	// available (NXDOMAIN). Either way the service should return 200 with the
	// domain name echoed back, without panicking.
	resp := svc.GetInfo(context.Background(), "example.com")
	assert.Equal(t, "example.com", resp.Domain)
}
