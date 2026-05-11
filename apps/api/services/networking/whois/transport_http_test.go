package whois

import (
	"bytes"
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

// fakeQuerier returns a fixed raw WHOIS text or an error.
type fakeQuerier struct {
	result string
	err    error
}

func (f *fakeQuerier) Whois(_ string, _ ...string) (string, error) {
	return f.result, f.err
}

// sampleWHOIS is a minimal but parseable WHOIS response for example.com.
const sampleWHOIS = `Domain Name: EXAMPLE.COM
Registry Domain ID: 2336799_DOMAIN_COM-VRSN
Registrar WHOIS Server: whois.iana.org
Registrar URL: http://res-dom.iana.org
Updated Date: 2023-08-14T07:01:38Z
Creation Date: 1995-08-14T04:00:00Z
Registrar Registration Expiration Date: 2024-08-13T04:00:00Z
Registrar: RESERVED-Internet Assigned Numbers Authority
Registrar IANA ID: 376
Domain Status: clientDeleteProhibited
Domain Status: clientTransferProhibited
Domain Status: clientUpdateProhibited
Name Server: A.IANA-SERVERS.NET
Name Server: B.IANA-SERVERS.NET
DNSSEC: signedDelegation
`

const notFoundWHOIS = `No match for "DOESNOTEXIST123456789.COM".
>>> Last update of whois database: 2024-01-01T00:00:00Z <<<`

func setupRouter(q Querier) chi.Router {
	r := chi.NewRouter()
	svc := &Service{q: q}
	RegisterRoutes(r, svc)
	return r
}

func TestWhois_ValidDomain(t *testing.T) {
	t.Parallel()
	r := setupRouter(&fakeQuerier{result: sampleWHOIS})

	req := httptest.NewRequest(http.MethodGet, "/whois/example.com", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[LookupResponse]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	assert.Equal(t, "example.com", resp.Data.Domain)
	assert.NotEmpty(t, resp.Data.Registrar)
	assert.NotEmpty(t, resp.Data.NameServers)
	assert.NotEmpty(t, resp.Data.CreatedDate)
}

func TestWhois_DomainNotFound(t *testing.T) {
	t.Parallel()
	r := setupRouter(&fakeQuerier{result: notFoundWHOIS})

	req := httptest.NewRequest(http.MethodGet, "/whois/doesnotexist123456789.com", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestWhois_InvalidDomainFormat(t *testing.T) {
	t.Parallel()
	r := setupRouter(&fakeQuerier{result: sampleWHOIS})

	tests := []struct {
		name   string
		domain string
	}{
		{"empty-like path segment with dots only", "..."},
		{"starts with hyphen", "-bad.com"},
		{"no TLD", "nodot"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			req := httptest.NewRequest(http.MethodGet, "/whois/"+tt.domain, http.NoBody)
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			assert.Equal(t, http.StatusBadRequest, w.Code, "expected 400 for %q, got %d: %s", tt.domain, w.Code, w.Body.String())
		})
	}
}

func TestWhois_QueryError(t *testing.T) {
	t.Parallel()
	r := setupRouter(&fakeQuerier{err: ErrDomainNotFound})

	req := httptest.NewRequest(http.MethodGet, "/whois/example.com", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestService_Lookup_NotFound(t *testing.T) {
	t.Parallel()
	svc := &Service{q: &fakeQuerier{result: notFoundWHOIS}}

	_, err := svc.Lookup(context.Background(), "doesnotexist.com")
	require.Error(t, err)
}

func TestService_Lookup_ValidDomain(t *testing.T) {
	t.Parallel()
	svc := &Service{q: &fakeQuerier{result: sampleWHOIS}}

	resp, err := svc.Lookup(context.Background(), "example.com")
	require.NoError(t, err)

	assert.Equal(t, "example.com", resp.Domain)
	assert.NotEmpty(t, resp.CreatedDate)
	assert.NotEmpty(t, resp.ExpiryDate)
}

func TestWhois_BatchLookup(t *testing.T) {
	r := setupRouter(&fakeQuerier{result: sampleWHOIS})

	body := `{
		"domains": [
			"example.com",
			"google.com"
		]
	}`

	req := httptest.NewRequest(
		http.MethodPost,
		"/whois/batch",
		strings.NewReader(body),
	)

	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[BatchLookupResponse]

	err := json.NewDecoder(w.Body).Decode(&resp)

	assert.NoError(t, err)
	assert.Equal(t, 2, resp.Data.Total)
	assert.Len(t, resp.Data.Results, 2)

	for _, item := range resp.Data.Results {
		assert.True(t, item.Found)
		assert.NotEmpty(t, item.Data.Domain)
		assert.Empty(t, item.Error)
	}
}

func TestWhois_BatchLookup_InvalidJSON(t *testing.T) {
	r := setupRouter(&fakeQuerier{result: sampleWHOIS})

	req := httptest.NewRequest(
		http.MethodPost,
		"/whois/batch",
		strings.NewReader(`{"domains":`),
	)

	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestWhois_BatchLookup_EmptyDomains(t *testing.T) {
	r := setupRouter(&fakeQuerier{result: sampleWHOIS})

	body := `{
		"domains": []
	}`

	req := httptest.NewRequest(
		http.MethodPost,
		"/whois/batch",
		strings.NewReader(body),
	)

	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
}
func TestWhois_BatchLookup_NotFound(t *testing.T) {
	r := setupRouter(&fakeQuerier{result: notFoundWHOIS})

	body := `{
		"domains": [
			"doesnotexist.com"
		]
	}`

	req := httptest.NewRequest(
		http.MethodPost,
		"/whois/batch",
		strings.NewReader(body),
	)

	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[BatchLookupResponse]

	err := json.NewDecoder(w.Body).Decode(&resp)

	assert.NoError(t, err)
	assert.Len(t, resp.Data.Results, 1)

	item := resp.Data.Results[0]

	assert.False(t, item.Found)
	assert.NotEmpty(t, item.Error)
}
func TestWhois_BatchLookup_TooManyDomains(t *testing.T) {
	r := setupRouter(&fakeQuerier{result: sampleWHOIS})

	domains := make([]string, 51)

	for i := range domains {
		domains[i] = "example.com"
	}

	body, err := json.Marshal(map[string]any{
		"domains": domains,
	})

	assert.NoError(t, err)

	req := httptest.NewRequest(
		http.MethodPost,
		"/whois/batch",
		bytes.NewReader(body),
	)

	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
}
