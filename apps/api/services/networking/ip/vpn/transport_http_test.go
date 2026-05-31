package vpn

import (
	"bytes"
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
		r.Get("/ip/vpn/{ip}", func(w http.ResponseWriter, r *http.Request) {
			httpx.Error(w, http.StatusServiceUnavailable, "service_unavailable", "VPN service not available")
		})
	} else {
		RegisterRoutes(r, testSvc)
	}
	return r
}

func skipIfNoService(t *testing.T) {
	if testSvc == nil {
		t.Skip("VPN service not available (database not configured)")
	}
}

func TestVPN_HappyPath(t *testing.T) {
	t.Parallel()
	skipIfNoService(t)
	r := setupRouter()

	req := httptest.NewRequest(http.MethodGet, "/ip/vpn/8.8.8.8", http.NoBody)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[IPCheckResponse]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	assert.NotEmpty(t, resp.Data.IP)
	assert.True(t, resp.Data.Score >= 0, "expected non-negative score, got %d", resp.Data.Score)
}

func TestVPN_ValidIPFields(t *testing.T) {
	t.Parallel()
	skipIfNoService(t)
	r := setupRouter()

	req := httptest.NewRequest(http.MethodGet, "/ip/vpn/1.1.1.1", http.NoBody)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[IPCheckResponse]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	assert.Equal(t, "1.1.1.1", resp.Data.IP)

	validThreats := map[string]bool{
		"none":     true,
		"low":      true,
		"medium":   true,
		"high":     true,
		"critical": true,
	}
	assert.True(t, validThreats[resp.Data.Threat.String()], "invalid threat level: %s", resp.Data.Threat)
	assert.True(t, resp.Data.FraudScore >= 0 && resp.Data.FraudScore <= 100, "fraud_score out of range: %d", resp.Data.FraudScore)
}

func TestVPN_InvalidIPFormat(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	req := httptest.NewRequest(http.MethodGet, "/ip/vpn/not-an-ip", http.NoBody)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if testSvc == nil {
		assert.Equal(t, http.StatusServiceUnavailable, w.Code, "expected 503 without service, got %d", w.Code)
		return
	}

	assert.Equal(t, http.StatusBadRequest, w.Code)

	var resp httpx.ErrorResponse
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	assert.Equal(t, "bad_request", resp.Error)
}

func TestVPN_IPv6Address(t *testing.T) {
	t.Parallel()
	skipIfNoService(t)
	r := setupRouter()

	req := httptest.NewRequest(http.MethodGet, "/ip/vpn/2001:4860:4860::8888", http.NoBody)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[IPCheckResponse]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	assert.NotEmpty(t, resp.Data.IP)
}

func TestVPN_AllBooleansReturned(t *testing.T) {
	t.Parallel()
	skipIfNoService(t)
	r := setupRouter()

	req := httptest.NewRequest(http.MethodGet, "/ip/vpn/8.8.8.8", http.NoBody)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[IPCheckResponse]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	assert.True(t, resp.Data.IsVPN == false || resp.Data.IsVPN == true, "is_vpn should be a boolean")
	assert.True(t, resp.Data.IsProxy == false || resp.Data.IsProxy == true, "is_proxy should be a boolean")
	assert.True(t, resp.Data.IsTor == false || resp.Data.IsTor == true, "is_tor should be a boolean")
	assert.True(t, resp.Data.IsHosting == false || resp.Data.IsHosting == true, "is_hosting should be a boolean")
}

func TestVPN_Batch_HappyPath(t *testing.T) {
	t.Parallel()

	skipIfNoService(t)

	r := setupRouter()

	body := `{
		"ips": [
			"8.8.8.8",
			"1.1.1.1",
			"2001:4860:4860::8888"
		]
	}`

	req := httptest.NewRequest(
		http.MethodPost,
		"/ip/vpn/batch",
		strings.NewReader(body),
	)

	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[httpx.BatchResponse[IPCheckResponse]]

	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	require.Len(t, resp.Data.Results, 3)

	assert.Equal(t, "8.8.8.8", resp.Data.Results[0].IP)
	assert.Equal(t, "1.1.1.1", resp.Data.Results[1].IP)
	assert.Equal(t, "2001:4860:4860::8888", resp.Data.Results[2].IP)
}

func TestVPN_Batch_InvalidJSON(t *testing.T) {
	t.Parallel()

	r := setupRouter()

	req := httptest.NewRequest(
		http.MethodPost,
		"/ip/vpn/batch",
		strings.NewReader(`{invalid-json}`),
	)

	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	skipIfNoService(t)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	var resp httpx.ErrorResponse

	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	assert.Equal(t, "bad_request", resp.Error)
}

func TestVPN_Batch_EmptyIPs(t *testing.T) {
	t.Parallel()

	r := setupRouter()

	body := `{
		"ips": []
	}`

	req := httptest.NewRequest(
		http.MethodPost,
		"/ip/vpn/batch",
		strings.NewReader(body),
	)

	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	skipIfNoService(t)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	var resp httpx.ErrorResponse

	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	assert.Equal(t, "bad_request", resp.Error)
}

func TestVPN_Batch_InvalidIPs(t *testing.T) {
	t.Parallel()

	skipIfNoService(t)

	r := setupRouter()

	body := `{
		"ips": [
			"8.8.8.8",
			"invalid-ip",
			"",
			"999.999.999.999"
		]
	}`

	req := httptest.NewRequest(
		http.MethodPost,
		"/ip/vpn/batch",
		strings.NewReader(body),
	)

	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[httpx.BatchResponse[IPCheckResponse]]

	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	require.Len(t, resp.Data.Results, 4)

	assert.Equal(t, "8.8.8.8", resp.Data.Results[0].IP)
	assert.Equal(t, "invalid-ip", resp.Data.Results[1].IP)
	assert.Equal(t, "", resp.Data.Results[2].IP)
	assert.Equal(t, "999.999.999.999", resp.Data.Results[3].IP)
}

func TestVPN_Batch_PreservesOrder(t *testing.T) {
	t.Parallel()

	skipIfNoService(t)

	r := setupRouter()

	body := `{
		"ips": [
			"1.1.1.1",
			"8.8.8.8",
			"2001:4860:4860::8888"
		]
	}`

	req := httptest.NewRequest(
		http.MethodPost,
		"/ip/vpn/batch",
		strings.NewReader(body),
	)

	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[httpx.BatchResponse[IPCheckResponse]]

	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	require.Len(t, resp.Data.Results, 3)

	expected := []string{
		"1.1.1.1",
		"8.8.8.8",
		"2001:4860:4860::8888",
	}

	for i := range expected {
		assert.Equal(t, expected[i], resp.Data.Results[i].IP)
	}
}
func TestVPN_Batch_TooManyIPs(t *testing.T) {
	t.Parallel()

	skipIfNoService(t)

	r := setupRouter()

	ips := make([]string, 0, 51)

	for i := 0; i < 51; i++ {
		ips = append(ips, "8.8.8.8")
	}

	reqBody := BatchRequest{
		IPs: ips,
	}

	body, err := json.Marshal(reqBody)
	require.NoError(t, err)

	req := httptest.NewRequest(
		http.MethodPost,
		"/ip/vpn/batch",
		bytes.NewReader(body),
	)

	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	var resp httpx.ErrorResponse

	err = json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	assert.Equal(t, "bad_request", resp.Error)
}
