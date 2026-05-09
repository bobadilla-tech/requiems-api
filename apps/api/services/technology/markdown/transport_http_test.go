package markdown

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

func TestMarkdown_HappyPath(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	body := `{"markdown":"# Hello\n\nThis is **bold** text."}`
	req := httptest.NewRequest(http.MethodPost, "/markdown", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[Response]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	want := "<h1>Hello</h1>\n<p>This is <strong>bold</strong> text.</p>"
	assert.Equal(t, want, resp.Data.HTML)
}

func TestMarkdown_Sanitize(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	body := `{"markdown":"Hello <script>alert('xss')</script>","sanitize":true}`
	req := httptest.NewRequest(http.MethodPost, "/markdown", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[Response]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	assert.NotContains(t, resp.Data.HTML, "<script>")
}

func TestMarkdown_MissingBody(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	req := httptest.NewRequest(http.MethodPost, "/markdown", http.NoBody)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestMarkdown_EmptyMarkdown(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	body := `{"markdown":""}`
	req := httptest.NewRequest(http.MethodPost, "/markdown", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	// Empty markdown triggers validation failure (required field)
	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
}

func TestMarkdownBatch_HappyPath(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	body := `{
		"markdowns": [
			"# Hello",
			"**bold**",
			"- item one"
		]
	}`
	req := httptest.NewRequest(http.MethodPost, "/markdown/batch", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code, w.Body.String())

	var resp httpx.Response[BatchResponse]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)
	require.Len(t, resp.Data.Results, 3)

	for i, result := range resp.Data.Results {
		assert.NotEmpty(t, result.HTML, "result[%d] should have non-empty HTML", i)
	}
}

func TestMarkdownBatch_Sanitize(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	body := `{
		"markdowns": [
			"Hello <script>alert('xss')</script>",
			"<strong>safe?</strong>"
		],
		"sanitize": true
	}`
	req := httptest.NewRequest(http.MethodPost, "/markdown/batch", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code, w.Body.String())

	var resp httpx.Response[BatchResponse]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	for i, result := range resp.Data.Results {
		assert.NotContains(t, result.HTML, "<script>", "result[%d] should not contain script tag", i)
	}
}

func TestMarkdownBatch_MissingBody(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	req := httptest.NewRequest(http.MethodPost, "/markdown/batch", http.NoBody)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestMarkdownBatch_EmptyMarkdowns(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	body := `{"markdowns": []}`
	req := httptest.NewRequest(http.MethodPost, "/markdown/batch", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code, w.Body.String())
}

func TestMarkdownBatch_WithEmptyMarkdownItem(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	body := `{"markdowns": ["", "# Title"]}`
	req := httptest.NewRequest(http.MethodPost, "/markdown/batch", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code, w.Body.String())
}

func TestMarkdownBatch_TooManyMarkdowns(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	items := make([]string, 51)
	for i := range items {
		items[i] = "# Title"
	}

	payload := struct {
		Markdowns []string `json:"markdowns"`
	}{
		Markdowns: items,
	}

	bodyBytes, err := json.Marshal(payload)
	require.NoError(t, err)

	req := httptest.NewRequest(http.MethodPost, "/markdown/batch", strings.NewReader(string(bodyBytes)))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code, w.Body.String())
}
