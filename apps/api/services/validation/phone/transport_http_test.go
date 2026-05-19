package phone

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

func TestPhone_ValidUSNumber(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	req := httptest.NewRequest(http.MethodGet, "/phone?number=%2B12015551234", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[ValidateResponse]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	assert.True(t, resp.Data.Valid)
	assert.Equal(t, "US", resp.Data.Country)
	assert.NotEmpty(t, resp.Data.Formatted)
	assert.NotEmpty(t, resp.Data.Type)
}

func TestPhone_InvalidNumber(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	req := httptest.NewRequest(http.MethodGet, "/phone?number=12345", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[ValidateResponse]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	assert.False(t, resp.Data.Valid)
	assert.Equal(t, "", resp.Data.Country)
}

func TestPhone_MissingNumber(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	req := httptest.NewRequest(http.MethodGet, "/phone", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
}

func TestPhone_UKMobile(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	req := httptest.NewRequest(http.MethodGet, "/phone?number=%2B447400123456", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[ValidateResponse]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	assert.True(t, resp.Data.Valid)
	assert.Equal(t, "GB", resp.Data.Country)
	assert.Equal(t, "mobile", resp.Data.Type)
}

func TestPhone_CarrierPresent(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	req := httptest.NewRequest(http.MethodGet, "/phone?number=%2B447400123456", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[ValidateResponse]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	require.True(t, resp.Data.Valid)
	require.NotNil(t, resp.Data.Carrier)
	assert.NotEmpty(t, resp.Data.Carrier.Name)
	assert.Equal(t, "metadata", resp.Data.Carrier.Source)
}

func TestPhone_RiskVoIP(t *testing.T) {
	t.Parallel()
	svc := NewService()
	// Google Voice numbers are VOIP type in the US (area code 202 VOIP range)
	// Use a number whose type we know via the service
	tests := []struct {
		name   string
		number string
	}{
		// +1-500 numbers are personal/VOIP in the US
		{"US personal number", "+15005550006"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := svc.Validate(tt.number)
			require.NotNil(t, result.Risk)
		})
	}
}

func TestPhone_RiskMobile(t *testing.T) {
	t.Parallel()
	svc := NewService()
	result := svc.Validate("+447400123456")

	require.True(t, result.Valid)
	require.NotNil(t, result.Risk)
	assert.False(t, result.Risk.IsVoIP)
	assert.False(t, result.Risk.IsVirtual)
}

func TestPhone_InvalidHasNoCarrierOrRisk(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	req := httptest.NewRequest(http.MethodGet, "/phone?number=12345", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var resp httpx.Response[ValidateResponse]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	require.False(t, resp.Data.Valid)
	assert.Nil(t, resp.Data.Carrier)
	assert.Nil(t, resp.Data.Risk)
}

func TestPhone_BatchValidate(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	body := `{"numbers":["+447400123456","+12015551234","12345"]}`
	req := httptest.NewRequest(http.MethodPost, "/phone/batch", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[httpx.BatchResponse[ValidateResponse]]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	assert.Equal(t, 3, resp.Data.Total)
	require.Len(t, resp.Data.Results, 3)
	assert.True(t, resp.Data.Results[0].Valid)
	assert.True(t, resp.Data.Results[1].Valid)
	assert.False(t, resp.Data.Results[2].Valid)
}

func TestPhone_BatchValidate_EmptyBody(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	req := httptest.NewRequest(http.MethodPost, "/phone/batch", strings.NewReader(`{"numbers":[]}`))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
}

func TestPhone_BatchValidate_ExceedsLimit(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	numbers := make([]string, 51)
	for i := range numbers {
		numbers[i] = `"+447400123456"`
	}
	body := `{"numbers":[` + strings.Join(numbers, ",") + `]}`
	req := httptest.NewRequest(http.MethodPost, "/phone/batch", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
}

func TestService_NumberType(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		number   string
		wantType string
	}{
		{"UK landline", "+441613281234", "landline"},
		{"UK mobile", "+447400123456", "mobile"},
		{"US toll-free", "+18005551234", "toll_free"},
	}

	svc := NewService()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := svc.Validate(tt.number)
			require.True(t, result.Valid)
			assert.Equal(t, tt.wantType, result.Type)
		})
	}
}
