package mx

import (
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

func setupRouter() chi.Router {
	r := chi.NewRouter()
	RegisterRoutes(r, NewService())
	return r
}

func TestMXLookup_InvalidDomain(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	tests := []struct {
		name   string
		domain string
	}{
		{"empty-like path", "not_a_domain"},
		{"plain label", "localhost"},
		{"starts with dash", "-bad.com"},
		{"IP address", "1.2.3.4"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			req := httptest.NewRequest(http.MethodGet, "/mx/"+tc.domain, http.NoBody)
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			assert.Equal(t, http.StatusBadRequest, w.Code, "expected 400 for %q, got %d", tc.domain, w.Code)

			var resp httpx.ErrorResponse
			err := json.NewDecoder(w.Body).Decode(&resp)
			require.NoError(t, err)
			assert.Equal(t, "bad_request", resp.Error)
		})
	}
}

func TestMXLookup_NonExistentDomain(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	req := httptest.NewRequest(http.MethodGet, "/mx/nonexistent-domain-that-does-not-exist.invalid", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Should return 404 (no MX records / NXDOMAIN) or 500 (network unavailable in CI)
	assert.True(t, w.Code == http.StatusNotFound || w.Code == http.StatusInternalServerError, "expected 404 or 500 for non-existent domain, got %d", w.Code)
}

func TestMXLookup_HappyPath(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	req := httptest.NewRequest(http.MethodGet, "/mx/gmail.com", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Network may not be available in all CI environments; accept 200 or 500/404
	if w.Code == http.StatusInternalServerError || w.Code == http.StatusNotFound {
		t.Skip("DNS not available in this environment")
	}

	require.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[LookupResponse]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	assert.Equal(t, "gmail.com", resp.Data.Domain)
	assert.NotEmpty(t, resp.Data.Records)

	// Verify priority ordering (ascending)
	for i := 1; i < len(resp.Data.Records); i++ {
		assert.True(t, resp.Data.Records[i].Priority >= resp.Data.Records[i-1].Priority,
			"records not sorted by priority: record[%d].Priority=%d < record[%d].Priority=%d",
			i, resp.Data.Records[i].Priority, i-1, resp.Data.Records[i-1].Priority)
	}

	// Each record should have a non-empty host
	for i, rec := range resp.Data.Records {
		assert.NotEmpty(t, rec.Host, "record[%d] has empty host", i)
	}
}

func TestMXLookup_BatchLookup_HappyPath(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	body := `{"domains":["gmail.com","outlook.com","yahoo.com"]}`
	req := httptest.NewRequest(http.MethodPost, "/mx/batch", strings.NewReader(body))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[httpx.BatchResponse[BatchLookupItem]]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	require.Equal(t, 3, resp.Data.Total)
	require.Len(t, resp.Data.Results, 3)

	for _, item := range resp.Data.Results {
		assert.True(t, item.Found)
		assert.NotEmpty(t, item.Data.Domain)
		assert.Empty(t, item.Error)
	}
}

func TestMxLookup_BatchLookup_EmptyBody(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	body := `{"domains":[]}`
	req := httptest.NewRequest(http.MethodPost, "/mx/batch", strings.NewReader(body))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var resp httpx.Response[httpx.BatchResponse[BatchLookupItem]]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	require.Equal(t, http.StatusUnprocessableEntity, w.Code)
}

func TestMxLookup_BatchLookup_ExceedsLimit(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	domains := make([]string, 51)

	for i := range domains {
		domains[i] = "gmail.com"
	}

	body, _ := json.Marshal(BatchRequest{
		Domains: domains,
	})

	req := httptest.NewRequest(http.MethodPost, "/mx/batch", strings.NewReader(string(body)))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusUnprocessableEntity, w.Code)
}

func TestMxLookup_BatchLookup_setUsageCountHeader(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	body := `{"domains":["gmail.com","outlook.com","yahoo.com"]}`
	req := httptest.NewRequest(http.MethodPost, "/mx/batch", strings.NewReader(body))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	require.Equal(t, "3", w.Header().Get("X-Usage-Count"))
}
