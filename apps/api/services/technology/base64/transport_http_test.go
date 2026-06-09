// Package base64 tests the HTTP transport for the base64 service.
package base64 //nolint:revive // name matches feature dir; stdlib stays encoding/base64 with import aliases in app code

import (
	"bytes"
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

// ── /base64/encode/batch ──────────────────────────────────────────────────────

func TestEncodeBatch_HappyPath(t *testing.T) {
	t.Parallel()

	req := httptest.NewRequest(
		http.MethodPost,
		"/base64/encode/batch",
		strings.NewReader(`{
			"values":["Hello","World"]
		}`),
	)

	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	setupRouter().ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)
	assertJSON(t, w)

	var resp httpx.Response[httpx.BatchResponse[Result]]

	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	require.Len(t, resp.Data.Results, 2)

	assert.Equal(t, "SGVsbG8=", resp.Data.Results[0].Result)
	assert.Equal(t, "V29ybGQ=", resp.Data.Results[1].Result)

	assert.Equal(t, 2, resp.Data.Total)
	assert.Equal(t, "2", w.Header().Get("X-Usage-Count"))
}

// TestEncodeBatch_InvalidVariant verifies that the endpoint rejects a batch
// request when "variant" contains a value other than "standard" or "url".
func TestEncodeBatch_InvalidVariant(t *testing.T) {
	t.Parallel()

	req := httptest.NewRequest(
		http.MethodPost,
		"/base64/encode/batch",
		strings.NewReader(`{
			"values":["Hello"],
			"variant":"invalid"
		}`),
	)

	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	setupRouter().ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
	assertJSON(t, w)
}

// TestEncodeBatch_EmptyValues verifies that the endpoint rejects a batch
// request when the values list is empty.
func TestEncodeBatch_EmptyValues(t *testing.T) {
	t.Parallel()

	req := httptest.NewRequest(
		http.MethodPost,
		"/base64/encode/batch",
		strings.NewReader(`{
			"values":[]
		}`),
	)

	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	setupRouter().ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
	assertJSON(t, w)
}

// TestEncodeBatch_TooManyValues verifies that the endpoint rejects a batch
// request containing more than the allowed 50 values.
func TestEncodeBatch_TooManyValues(t *testing.T) {
	t.Parallel()

	values := make([]string, 51)

	for i := range values {
		values[i] = "hello"
	}

	body, err := json.Marshal(map[string]any{
		"values": values,
	})
	require.NoError(t, err)

	req := httptest.NewRequest(
		http.MethodPost,
		"/base64/encode/batch",
		bytes.NewReader(body),
	)

	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	setupRouter().ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
	assertJSON(t, w)
}

// ── /base64/decode/batch ──────────────────────────────────────────────────────

func TestDecodeBatch_HappyPath(t *testing.T) {
	t.Parallel()

	req := httptest.NewRequest(
		http.MethodPost,
		"/base64/decode/batch",
		strings.NewReader(`{
			"values":["SGVsbG8=","V29ybGQ="]
		}`),
	)

	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	setupRouter().ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)
	assertJSON(t, w)

	var resp httpx.Response[httpx.BatchResponse[Result]]

	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	require.Len(t, resp.Data.Results, 2)

	assert.Equal(t, "Hello", resp.Data.Results[0].Result)
	assert.Equal(t, "World", resp.Data.Results[1].Result)

	assert.Equal(t, 2, resp.Data.Total)
	assert.Equal(t, "2", w.Header().Get("X-Usage-Count"))
}

// TestDecodeBatch_InvalidVariant verifies that the endpoint rejects a batch
// request when "variant" contains a value other than "standard" or "url".
func TestDecodeBatch_InvalidVariant(t *testing.T) {
	t.Parallel()

	req := httptest.NewRequest(
		http.MethodPost,
		"/base64/decode/batch",
		strings.NewReader(`{
			"values":["SGVsbG8="],
			"variant":"invalid"
		}`),
	)

	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	setupRouter().ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
	assertJSON(t, w)
}

// TestDecodeBatch_EmptyValues verifies that the endpoint rejects a batch
// request when the values list is empty.
func TestDecodeBatch_EmptyValues(t *testing.T) {
	t.Parallel()

	req := httptest.NewRequest(
		http.MethodPost,
		"/base64/decode/batch",
		strings.NewReader(`{
			"values":[]
		}`),
	)

	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	setupRouter().ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
	assertJSON(t, w)
}

// TestDecodeBatch_PartialFailure verifies that invalid Base64 entries are
// returned as empty results while valid entries are decoded successfully.
func TestDecodeBatch_PartialFailure(t *testing.T) {
	t.Parallel()

	req := httptest.NewRequest(
		http.MethodPost,
		"/base64/decode/batch",
		strings.NewReader(`{
			"values":["SGVsbG8=","not-valid-base64!!!"]
		}`),
	)

	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	setupRouter().ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)
	assertJSON(t, w)

	var resp httpx.Response[httpx.BatchResponse[Result]]

	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	require.Len(t, resp.Data.Results, 2)

	assert.Equal(t, "Hello", resp.Data.Results[0].Result)

	// Partial success: invalid entries are returned as empty strings.
	assert.Equal(t, "", resp.Data.Results[1].Result)

	assert.Equal(t, 2, resp.Data.Total)
	assert.Equal(t, "2", w.Header().Get("X-Usage-Count"))
}
