// Package base64 tests the HTTP transport for the base64 service.
package base64 //nolint:revive // name matches feature dir; stdlib stays encoding/base64 with import aliases in app code

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

const (
	encodedHello = "SGVsbG8sIHdvcmxkIQ=="
	decodedHello = "Hello, world!"
)

// setupRouter builds a chi router with the base64 service wired up.
func setupRouter() chi.Router {
	r := chi.NewRouter()
	RegisterRoutes(r, NewService())
	return r
}

// assertJSON verifies that the response has a JSON Content-Type and a valid JSON body.
func assertJSON(t *testing.T, w *httptest.ResponseRecorder) {
	t.Helper()
	ct := w.Header().Get("Content-Type")
	if !strings.HasPrefix(ct, "application/json") {
		t.Errorf("Content-Type: got %q, want application/json", ct)
	}
	if !json.Valid(w.Body.Bytes()) {
		t.Errorf("body is not valid JSON: %s", w.Body.String())
	}
}

// ── /base64/encode ────────────────────────────────────────────────────────────

func TestEncode_HappyPath(t *testing.T) {
	t.Parallel()
	req := httptest.NewRequest(
		http.MethodPost,
		"/base64/encode",
		strings.NewReader(`{"value":"`+decodedHello+`"}`),
	)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	setupRouter().ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)
	assertJSON(t, w)

	var resp httpx.Response[Result]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)
	assert.Equal(t, encodedHello, resp.Data.Result)
}

// TestEncode_MissingValue verifies that the endpoint rejects a request
// when the required "value" field is missing.
func TestEncode_MissingValue(t *testing.T) {
	t.Parallel()
	req := httptest.NewRequest(http.MethodPost, "/base64/encode",
		strings.NewReader(`{}`))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	setupRouter().ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
	assertJSON(t, w)
}

// TestEncode_InvalidVariant verifies that the endpoint rejects a request
// when "variant" contains a value other than "standard" or "url".
func TestEncode_InvalidVariant(t *testing.T) {
	t.Parallel()
	req := httptest.NewRequest(http.MethodPost, "/base64/encode",
		strings.NewReader(`{"value":"`+decodedHello+`","variant":"invalid"}`))

	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	setupRouter().ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
	assertJSON(t, w)
}

// ── /base64/decode ────────────────────────────────────────────────────────────

func TestDecode_HappyPath(t *testing.T) {
	t.Parallel()
	req := httptest.NewRequest(http.MethodPost, "/base64/decode",
		strings.NewReader(`{"value":"`+encodedHello+`"}`))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	setupRouter().ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)
	assertJSON(t, w)

	var resp httpx.Response[Result]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)
	assert.Equal(t, decodedHello, resp.Data.Result)
}

// TestDecode_MissingValue verifies that the endpoint rejects a request
// when the required "value" field is missing.
func TestDecode_MissingValue(t *testing.T) {
	t.Parallel()
	req := httptest.NewRequest(http.MethodPost, "/base64/decode",
		strings.NewReader(`{}`))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	setupRouter().ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
	assertJSON(t, w)
}

// TestDecode_InvalidVariant verifies that the endpoint rejects a request
// when "variant" contains a value other than "standard" or "url".
func TestDecode_InvalidVariant(t *testing.T) {
	t.Parallel()
	req := httptest.NewRequest(http.MethodPost, "/base64/decode",
		strings.NewReader(`{"value":"`+encodedHello+`","variant":"invalid"}`))

	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	setupRouter().ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
	assertJSON(t, w)
}

// TestDecode_ServiceError verifies that the endpoint returns 422 when the
// value passes validation but is not valid base64.
func TestDecode_ServiceError(t *testing.T) {
	t.Parallel()
	req := httptest.NewRequest(http.MethodPost, "/base64/decode",
		strings.NewReader(`{"value":"not-valid-base64!!!"}`))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	setupRouter().ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
	assertJSON(t, w)
}
