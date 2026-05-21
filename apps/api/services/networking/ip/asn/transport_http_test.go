package asn

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/bobadilla-tech/go-ip-intelligence/v2/ipi"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"requiems-api/platform/httpx"
)

var testSvc *Service

func init() {
	client, err := ipi.New(
		ipi.WithDatabasePath(""),
		ipi.WithASNDatabasePath(""),
		ipi.WithCityDatabasePath(""),
	)
	if err == nil {
		testSvc = NewService(client)
	}
}

func setupRouter() chi.Router {
	r := chi.NewRouter()
	if testSvc == nil {
		r.Get("/ip/asn/{ip}", func(w http.ResponseWriter, r *http.Request) {
			httpx.Error(w, http.StatusServiceUnavailable, "service_unavailable", "ASN service not available")
		})
		r.Get("/ip/asn", func(w http.ResponseWriter, r *http.Request) {
			httpx.Error(w, http.StatusServiceUnavailable, "service_unavailable", "ASN service not available")
		})
	} else {
		RegisterRoutes(r, testSvc)
	}
	return r
}

func skipIfNoService(t *testing.T) {
	t.Helper()
	if testSvc == nil {
		t.Skip("ASN service not available (database not configured)")
	}
}

func TestASN_HappyPath(t *testing.T) {
	t.Parallel()
	skipIfNoService(t)
	r := setupRouter()

	req := httptest.NewRequest(http.MethodGet, "/ip/asn/8.8.8.8", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[IPAddressASNResponse]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	assert.NotEmpty(t, resp.Data.IP)
	assert.NotEmpty(t, resp.Data.ASN)
}

func TestASN_InvalidIPFormat(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	req := httptest.NewRequest(http.MethodGet, "/ip/asn/not-an-ip", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if testSvc == nil {
		assert.Equal(t, http.StatusServiceUnavailable, w.Code)
		return
	}

	assert.Equal(t, http.StatusBadRequest, w.Code)

	var resp httpx.ErrorResponse
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)
	assert.Equal(t, "bad_request", resp.Error)
}

func TestASN_PrivateIP(t *testing.T) {
	t.Parallel()
	skipIfNoService(t)
	r := setupRouter()

	req := httptest.NewRequest(http.MethodGet, "/ip/asn/192.168.1.1", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[IPAddressASNResponse]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	assert.Equal(t, "192.168.1.1", resp.Data.IP)
	assert.Empty(t, resp.Data.ASN)
}

func TestASN_NoIPParam_UsesRemoteAddr(t *testing.T) {
	t.Parallel()
	skipIfNoService(t)
	r := setupRouter()

	req := httptest.NewRequest(http.MethodGet, "/ip/asn", http.NoBody)
	req.RemoteAddr = "8.8.8.8:12345"
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)
}

func TestASN_XForwardedFor(t *testing.T) {
	t.Parallel()
	skipIfNoService(t)
	r := setupRouter()

	req := httptest.NewRequest(http.MethodGet, "/ip/asn", http.NoBody)
	req.Header.Set("X-Forwarded-For", "1.1.1.1, 10.0.0.1")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[IPAddressASNResponse]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)
	assert.Equal(t, "1.1.1.1", resp.Data.IP)
}

func TestASN_Batch_HappyPath(t *testing.T) {
	t.Parallel()
	skipIfNoService(t)
	r := setupRouter()

	body := `{"ips":["8.8.8.8"]}`
	req := httptest.NewRequest(http.MethodPost, "/ip/asn/batch", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)
	var resp httpx.Response[httpx.BatchResponse[BatchASNItem]]
	require.NoError(t, json.NewDecoder(w.Body).Decode(&resp))
	require.Len(t, resp.Data.Results, 1)
	assert.NotNil(t, resp.Data.Results[0].Result)
	assert.NotEmpty(t, resp.Data.Results[0].Result.ASN)
}

func TestASN_Batch_PrivateIP(t *testing.T) {
	t.Parallel()
	skipIfNoService(t)
	r := setupRouter()

	body := `{"ips":["10.0.0.1"]}`
	req := httptest.NewRequest(http.MethodPost, "/ip/asn/batch", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)
	var resp httpx.Response[httpx.BatchResponse[BatchASNItem]]
	require.NoError(t, json.NewDecoder(w.Body).Decode(&resp))
	require.Len(t, resp.Data.Results, 1)
	assert.NotNil(t, resp.Data.Results[0].Result)
	assert.Empty(t, resp.Data.Results[0].Result.ASN)
}

func TestASN_Batch_InvalidIP_Rejected(t *testing.T) {
	t.Parallel()
	skipIfNoService(t)
	r := setupRouter()

	body := `{"ips":["bad"]}`
	req := httptest.NewRequest(http.MethodPost, "/ip/asn/batch", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
}
