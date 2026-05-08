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

	assert.False(t, strings.Contains(resp.Data.HTML, "<script>"), "expected script tag to be stripped, got: %q", resp.Data.HTML)
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
