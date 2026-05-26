package textnormalize

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

func post(t *testing.T, r chi.Router, body string) *httptest.ResponseRecorder {
	t.Helper()
	req := httptest.NewRequest(http.MethodPost, "/text/normalize", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestNormalize_Trim(t *testing.T) {
	t.Parallel()
	r := setupRouter()
	w := post(t, r, `{"text":"  hello world  ","operations":["trim"]}`)
	require.Equal(t, http.StatusOK, w.Code)
	var resp httpx.Response[Result]
	require.NoError(t, json.NewDecoder(w.Body).Decode(&resp))
	assert.Equal(t, "hello world", resp.Data.Normalized)
	assert.Equal(t, "  hello world  ", resp.Data.Original)
}

func TestNormalize_Lowercase(t *testing.T) {
	t.Parallel()
	r := setupRouter()
	w := post(t, r, `{"text":"Hello WORLD","operations":["lowercase"]}`)
	require.Equal(t, http.StatusOK, w.Code)
	var resp httpx.Response[Result]
	require.NoError(t, json.NewDecoder(w.Body).Decode(&resp))
	assert.Equal(t, "hello world", resp.Data.Normalized)
}

func TestNormalize_Uppercase(t *testing.T) {
	t.Parallel()
	r := setupRouter()
	w := post(t, r, `{"text":"hello world","operations":["uppercase"]}`)
	require.Equal(t, http.StatusOK, w.Code)
	var resp httpx.Response[Result]
	require.NoError(t, json.NewDecoder(w.Body).Decode(&resp))
	assert.Equal(t, "HELLO WORLD", resp.Data.Normalized)
}

func TestNormalize_NormalizeUnicode(t *testing.T) {
	t.Parallel()
	r := setupRouter()
	// café composed vs decomposed — both should produce NFC
	w := post(t, r, `{"text":"café","operations":["normalize_unicode"]}`)
	require.Equal(t, http.StatusOK, w.Code)
	var resp httpx.Response[Result]
	require.NoError(t, json.NewDecoder(w.Body).Decode(&resp))
	assert.Equal(t, "café", resp.Data.Normalized)
}

func TestNormalize_CollapseWhitespace(t *testing.T) {
	t.Parallel()
	r := setupRouter()
	w := post(t, r, `{"text":"hello   \t world\n\nfoo","operations":["collapse_whitespace"]}`)
	require.Equal(t, http.StatusOK, w.Code)
	var resp httpx.Response[Result]
	require.NoError(t, json.NewDecoder(w.Body).Decode(&resp))
	assert.Equal(t, "hello world foo", resp.Data.Normalized)
}

func TestNormalize_StripHTML(t *testing.T) {
	t.Parallel()
	r := setupRouter()
	w := post(t, r, `{"text":"<p>Hello <b>world</b>&amp;!</p>","operations":["strip_html"]}`)
	require.Equal(t, http.StatusOK, w.Code)
	var resp httpx.Response[Result]
	require.NoError(t, json.NewDecoder(w.Body).Decode(&resp))
	assert.Contains(t, resp.Data.Normalized, "Hello")
	assert.Contains(t, resp.Data.Normalized, "world")
	assert.Contains(t, resp.Data.Normalized, "&!")
	assert.NotContains(t, resp.Data.Normalized, "<p>")
	assert.NotContains(t, resp.Data.Normalized, "<b>")
}

func TestNormalize_RemovePunctuation(t *testing.T) {
	t.Parallel()
	r := setupRouter()
	w := post(t, r, `{"text":"Hello, world! How are you?","operations":["remove_punctuation"]}`)
	require.Equal(t, http.StatusOK, w.Code)
	var resp httpx.Response[Result]
	require.NoError(t, json.NewDecoder(w.Body).Decode(&resp))
	assert.Equal(t, "Hello world How are you", resp.Data.Normalized)
}

func TestNormalize_MultipleOperationsOrdered(t *testing.T) {
	t.Parallel()
	r := setupRouter()
	// trim → collapse_whitespace → lowercase applied in order
	w := post(t, r, `{"text":"  Hello   WORLD  ","operations":["trim","collapse_whitespace","lowercase"]}`)
	require.Equal(t, http.StatusOK, w.Code)
	var resp httpx.Response[Result]
	require.NoError(t, json.NewDecoder(w.Body).Decode(&resp))
	assert.Equal(t, "hello world", resp.Data.Normalized)
}

func TestNormalize_OperationsEchoedInResponse(t *testing.T) {
	t.Parallel()
	r := setupRouter()
	w := post(t, r, `{"text":"hello","operations":["trim","lowercase"]}`)
	require.Equal(t, http.StatusOK, w.Code)
	var resp httpx.Response[Result]
	require.NoError(t, json.NewDecoder(w.Body).Decode(&resp))
	assert.Equal(t, []string{"trim", "lowercase"}, resp.Data.Operations)
}

func TestNormalize_UnknownOperationIsNoOp(t *testing.T) {
	t.Parallel()
	r := setupRouter()
	w := post(t, r, `{"text":"hello","operations":["unknown_op"]}`)
	require.Equal(t, http.StatusOK, w.Code)
	var resp httpx.Response[Result]
	require.NoError(t, json.NewDecoder(w.Body).Decode(&resp))
	assert.Equal(t, "hello", resp.Data.Normalized)
}

func TestNormalize_MissingText(t *testing.T) {
	t.Parallel()
	r := setupRouter()
	w := post(t, r, `{"operations":["trim"]}`)
	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
}

func TestNormalize_MissingOperations(t *testing.T) {
	t.Parallel()
	r := setupRouter()
	w := post(t, r, `{"text":"hello"}`)
	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
}

func TestNormalize_EmptyOperationsArray(t *testing.T) {
	t.Parallel()
	r := setupRouter()
	w := post(t, r, `{"text":"hello","operations":[]}`)
	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
}

func TestNormalize_EmptyBody(t *testing.T) {
	t.Parallel()
	r := setupRouter()
	req := httptest.NewRequest(http.MethodPost, "/text/normalize", http.NoBody)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}
