package normalize

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

func TestNormalize_HappyPath(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	body := `{"email":"user@example.com"}`
	req := httptest.NewRequest(http.MethodPost, "/normalize", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[EmailNormalization]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	assert.NotEmpty(t, resp.Data.Normalized)
	assert.Equal(t, "user@example.com", resp.Data.Original)
	assert.Equal(t, "user", resp.Data.Local)
	assert.Equal(t, "example.com", resp.Data.Domain)
}

func TestNormalize_GmailNormalization(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	body := `{"email":"te.st.user+spam@gmail.com"}`
	req := httptest.NewRequest(http.MethodPost, "/normalize", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[EmailNormalization]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	assert.Equal(t, "testuser@gmail.com", resp.Data.Normalized)
	assert.NotEmpty(t, resp.Data.Changes)
}

func TestNormalize_UppercaseDomainLowercased(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	// For unknown providers the local part is preserved (case-sensitive per
	// RFC 5321); only the domain is lowercased.
	body := `{"email":"USER@EXAMPLE.COM"}`
	req := httptest.NewRequest(http.MethodPost, "/normalize", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[EmailNormalization]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	assert.Equal(t, "USER@example.com", resp.Data.Normalized)
}

func TestNormalize_OriginalIsAlwaysUnmodifiedInput(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	input := "Test.User+tag@Gmail.com"
	body := `{"email":"` + input + `"}`
	req := httptest.NewRequest(http.MethodPost, "/normalize", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[EmailNormalization]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	assert.Equal(t, input, resp.Data.Original)
}

func TestNormalize_MissingEmailField(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	req := httptest.NewRequest(http.MethodPost, "/normalize", strings.NewReader(`{}`))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
}

func TestNormalize_InvalidEmailFormat(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	body := `{"email":"not-an-email"}`
	req := httptest.NewRequest(http.MethodPost, "/normalize", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestNormalize_MissingBody(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	req := httptest.NewRequest(http.MethodPost, "/normalize", http.NoBody)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestNormalize_UnknownFieldsRejected(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	body := `{"email":"user@example.com","unexpected_field":"value"}`
	req := httptest.NewRequest(http.MethodPost, "/normalize", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestNormalizeBatch_HappyPathSetsUsageCount(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	body := `{"emails":["a@b.co","user@example.com"]}`
	req := httptest.NewRequest(http.MethodPost, "/normalize/batch", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "2", w.Header().Get("X-Usage-Count"))

	var resp httpx.Response[httpx.BatchResponse[EmailNormalizationBatchItem]]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)
	require.True(t, resp.Data.Total == 2 && len(resp.Data.Results) == 2, "total/results: %+v", resp.Data)
	require.True(t, resp.Data.Results[0].Valid && resp.Data.Results[1].Valid, "expected both valid, got %+v", resp.Data.Results)
}

func TestNormalizeBatch_InvalidItemInBand(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	body := `{"emails":["user@example.com","not-an-email"]}`
	req := httptest.NewRequest(http.MethodPost, "/normalize/batch", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[httpx.BatchResponse[EmailNormalizationBatchItem]]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)
	assert.True(t, resp.Data.Results[0].Valid && !resp.Data.Results[1].Valid, "want first valid second invalid, got %+v", resp.Data.Results)
}

func TestNormalizeBatch_EmptyArrayValidation(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	body := `{"emails":[]}`
	req := httptest.NewRequest(http.MethodPost, "/normalize/batch", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
}
