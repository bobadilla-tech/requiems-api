package info

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
		r.Get("/ip/{ip}", func(w http.ResponseWriter, r *http.Request) {
			httpx.Error(w, http.StatusServiceUnavailable, "service_unavailable", "IP info service not available")
		})
		r.Get("/ip", func(w http.ResponseWriter, r *http.Request) {
			httpx.Error(w, http.StatusServiceUnavailable, "service_unavailable", "IP info service not available")
		})
	} else {
		RegisterRoutes(r, testSvc)
	}
	return r
}

func skipIfNoService(t *testing.T) {
	t.Helper()
	if testSvc == nil {
		t.Skip("IP info service not available (database not configured)")
	}
}

func TestInfo_HappyPath_PathParam(t *testing.T) {
	t.Parallel()
	skipIfNoService(t)
	r := setupRouter()

	req := httptest.NewRequest(http.MethodGet, "/ip/8.8.8.8", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[LookupResponse]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	assert.Equal(t, "8.8.8.8", resp.Data.IP)
}

func TestInfo_HappyPath_NoPathParam(t *testing.T) {
	t.Parallel()
	skipIfNoService(t)
	r := setupRouter()

	req := httptest.NewRequest(http.MethodGet, "/ip", http.NoBody)
	req.RemoteAddr = "1.1.1.1:54321"
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[LookupResponse]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	assert.NotEmpty(t, resp.Data.IP)
}

func TestInfo_InvalidIPFormat(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	req := httptest.NewRequest(http.MethodGet, "/ip/not-an-ip", http.NoBody)
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

func TestInfo_PrivateIP_ReturnsEmpty(t *testing.T) {
	t.Parallel()
	skipIfNoService(t)
	r := setupRouter()

	req := httptest.NewRequest(http.MethodGet, "/ip/10.0.0.1", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[LookupResponse]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	assert.NotEmpty(t, resp.Data.IP)
	// Country/city should be empty for private IPs
	assert.Empty(t, resp.Data.Country)
}

func TestInfo_XRealIP(t *testing.T) {
	t.Parallel()
	skipIfNoService(t)
	r := setupRouter()

	req := httptest.NewRequest(http.MethodGet, "/ip", http.NoBody)
	req.Header.Set("X-Real-IP", "8.8.4.4")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[LookupResponse]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)
	assert.Equal(t, "8.8.4.4", resp.Data.IP)
}

func TestInfo_IPv6Address(t *testing.T) {
	t.Parallel()
	skipIfNoService(t)
	r := setupRouter()

	req := httptest.NewRequest(http.MethodGet, "/ip/2001:4860:4860::8888", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[LookupResponse]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	assert.NotEmpty(t, resp.Data.IP)
}

func TestInfo_ResponseFields(t *testing.T) {
	t.Parallel()
	skipIfNoService(t)
	r := setupRouter()

	req := httptest.NewRequest(http.MethodGet, "/ip/1.1.1.1", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	body := w.Body.Bytes()

	var resp httpx.Response[LookupResponse]
	err := json.Unmarshal(body, &resp)
	require.NoError(t, err)

	assert.Equal(t, "1.1.1.1", resp.Data.IP)

	var raw map[string]any
	err = json.Unmarshal(body, &raw)
	require.NoError(t, err)
	data, ok := raw["data"].(map[string]any)
	require.True(t, ok, "expected object at data")
	_, ok = data["is_vpn"]
	assert.True(t, ok, "expected data.is_vpn to be present in response JSON")
}

func TestInfo_Batch_HappyPath(t *testing.T) {
	t.Parallel()
	skipIfNoService(t)
	r := setupRouter()

	body := `{"ips":["8.8.8.8","1.1.1.1"]}`
	req := httptest.NewRequest(http.MethodPost, "/ip/info/batch", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)
	var resp httpx.Response[httpx.BatchResponse[BatchIPInfoItem]]
	require.NoError(t, json.NewDecoder(w.Body).Decode(&resp))
	assert.Equal(t, 2, resp.Data.Total)
	require.Len(t, resp.Data.Results, 2)
	assert.NotNil(t, resp.Data.Results[0].Result)
	assert.NotNil(t, resp.Data.Results[1].Result)
}

func TestInfo_Batch_PrivateIP(t *testing.T) {
	t.Parallel()
	skipIfNoService(t)
	r := setupRouter()

	body := `{"ips":["192.168.1.1"]}`
	req := httptest.NewRequest(http.MethodPost, "/ip/info/batch", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)
	var resp httpx.Response[httpx.BatchResponse[BatchIPInfoItem]]
	require.NoError(t, json.NewDecoder(w.Body).Decode(&resp))
	require.Len(t, resp.Data.Results, 1)
	assert.NotNil(t, resp.Data.Results[0].Result)
	assert.Empty(t, resp.Data.Results[0].Result.Country)
}

func TestInfo_Batch_InvalidIP_Rejected(t *testing.T) {
	t.Parallel()
	skipIfNoService(t)
	r := setupRouter()

	body := `{"ips":["not-an-ip"]}`
	req := httptest.NewRequest(http.MethodPost, "/ip/info/batch", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
}
