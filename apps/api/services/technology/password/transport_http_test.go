package password

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
	svc := NewService()
	RegisterRoutes(r, svc)
	return r
}

func TestPassword_DefaultLength(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	req := httptest.NewRequest(http.MethodGet, "/password", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[Password]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	assert.Equal(t, 16, resp.Data.Length)
	assert.Equal(t, 16, len(resp.Data.Password))
}

func TestPassword_CustomLength(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	req := httptest.NewRequest(http.MethodGet, "/password?length=32", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[Password]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	assert.Equal(t, 32, resp.Data.Length)
	assert.Equal(t, 32, len(resp.Data.Password))
}

func TestPassword_AllCharsets(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	req := httptest.NewRequest(http.MethodGet, "/password?length=64&uppercase=true&numbers=true&symbols=true", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[Password]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	pwd := resp.Data.Password

	hasLower := strings.ContainsAny(pwd, charsetLower)
	hasUpper := strings.ContainsAny(pwd, charsetUpper)
	hasDigit := strings.ContainsAny(pwd, charsetNumbers)
	hasSymbol := strings.ContainsAny(pwd, charsetSymbols)

	assert.True(t, hasLower, "expected at least one lowercase letter")
	assert.True(t, hasUpper, "expected at least one uppercase letter")
	assert.True(t, hasDigit, "expected at least one digit")
	assert.True(t, hasSymbol, "expected at least one symbol")
	assert.Equal(t, "strong", resp.Data.Strength)
}

func TestPassword_LengthTooShort(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	req := httptest.NewRequest(http.MethodGet, "/password?length=4", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
}

func TestPassword_LengthTooLong(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	req := httptest.NewRequest(http.MethodGet, "/password?length=200", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
}

func TestPassword_StrengthWeak(t *testing.T) {
	t.Parallel()
	svc := NewService()

	result, err := svc.Generate(8, false, false, false)
	require.NoError(t, err)

	assert.Equal(t, "weak", result.Strength)
}

func TestPassword_StrengthMedium(t *testing.T) {
	t.Parallel()
	svc := NewService()

	result, err := svc.Generate(8, true, false, false)
	require.NoError(t, err)

	assert.Equal(t, "medium", result.Strength)
}

func TestPassword_StrengthStrong(t *testing.T) {
	t.Parallel()
	svc := NewService()

	result, err := svc.Generate(16, true, true, true)
	require.NoError(t, err)

	assert.Equal(t, "strong", result.Strength)
}

func TestPassword_OnlyLowercase(t *testing.T) {
	t.Parallel()
	svc := NewService()

	result, err := svc.Generate(12, false, false, false)
	require.NoError(t, err)

	for _, c := range result.Password {
		assert.True(t, strings.ContainsRune(charsetLower, c), "unexpected character %q in lowercase-only password", c)
	}
}

func TestPassword_NoSymbolsWhenNotRequested(t *testing.T) {
	t.Parallel()
	svc := NewService()

	result, err := svc.Generate(32, true, true, false)
	require.NoError(t, err)

	assert.False(t, strings.ContainsAny(result.Password, charsetSymbols), "expected no symbols in password when symbols not requested")
}

func TestPassword_Batch_HappyPath(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	body := `{"items":[{"length":12},{"length":24,"uppercase":true,"numbers":true}]}`
	req := httptest.NewRequest(http.MethodPost, "/password/batch", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)
	var resp httpx.Response[httpx.BatchResponse[BatchResult]]
	require.NoError(t, json.NewDecoder(w.Body).Decode(&resp))
	assert.Equal(t, 2, resp.Data.Total)
	require.NotNil(t, resp.Data.Results[0].Result)
	require.NotNil(t, resp.Data.Results[1].Result)
}

func TestPassword_Batch_DefaultLength(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	body := `{"items":[{}]}`
	req := httptest.NewRequest(http.MethodPost, "/password/batch", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)
	var resp httpx.Response[httpx.BatchResponse[BatchResult]]
	require.NoError(t, json.NewDecoder(w.Body).Decode(&resp))
	require.NotNil(t, resp.Data.Results[0].Result)
	assert.Equal(t, 16, resp.Data.Results[0].Result.Length)
}

func TestPassword_Batch_EmptyArray(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	req := httptest.NewRequest(http.MethodPost, "/password/batch", strings.NewReader(`{"items":[]}`))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
}

func TestPassword_Batch_ExceedsLimit(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	items := make([]string, 51)
	for i := range items {
		items[i] = `{}`
	}
	body := `{"items":[` + strings.Join(items, ",") + `]}`
	req := httptest.NewRequest(http.MethodPost, "/password/batch", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
}
