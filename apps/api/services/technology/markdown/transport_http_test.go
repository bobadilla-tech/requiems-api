package markdown

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
)

func setupRouter() chi.Router {
	r := chi.NewRouter()
	RegisterRoutes(r, NewService())
	return r
}

func TestMarkdown_HappyPath(t *testing.T) {
	r := setupRouter()

	body := `{"markdown":"# Hello\n\nThis is **bold** text."}`
	req := httptest.NewRequest(http.MethodPost, "/markdown", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d: %s", w.Code, w.Body.String())
	}

	var resp httpx.Response[Response]
	if err := json.NewDecoder(w.Body).Decode(&resp); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	want := "<h1>Hello</h1>\n<p>This is <strong>bold</strong> text.</p>"
	if resp.Data.HTML != want {
		t.Errorf("html mismatch\ngot:  %q\nwant: %q", resp.Data.HTML, want)
	}
}

func TestMarkdown_Sanitize(t *testing.T) {
	r := setupRouter()

	body := `{"markdown":"Hello <script>alert('xss')</script>","sanitize":true}`
	req := httptest.NewRequest(http.MethodPost, "/markdown", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d: %s", w.Code, w.Body.String())
	}

	var resp httpx.Response[Response]
	if err := json.NewDecoder(w.Body).Decode(&resp); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if strings.Contains(resp.Data.HTML, "<script>") {
		t.Errorf("expected script tag to be stripped, got: %q", resp.Data.HTML)
	}
}

func TestMarkdown_MissingBody(t *testing.T) {
	r := setupRouter()

	req := httptest.NewRequest(http.MethodPost, "/markdown", http.NoBody)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected 400, got %d", w.Code)
	}

}

func TestMarkdown_EmptyMarkdown(t *testing.T) {
	r := setupRouter()

	body := `{"markdown":""}`
	req := httptest.NewRequest(http.MethodPost, "/markdown", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	// Empty markdown triggers validation failure (required field)
	if w.Code != http.StatusUnprocessableEntity {
		t.Errorf("expected 422, got %d: %s", w.Code, w.Body.String())
	}
}

func TestMarkdownBatch_HappyPath(t *testing.T) {
	r := setupRouter()

	body := `{
		"markdowns": [
			"# Hello",
			"**bold**",
			"- item one"
		]
	}`

	req := httptest.NewRequest(
		http.MethodPost,
		"/markdown/batch",
		strings.NewReader(body),
	)

	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d: %s", w.Code, w.Body.String())
	}

	var resp httpx.Response[BatchResponse]

	if err := json.NewDecoder(w.Body).Decode(&resp); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if len(resp.Data.Results) != 3 {
		t.Fatalf(
			"expected 3 results, got %d",
			len(resp.Data.Results),
		)
	}

	for i, result := range resp.Data.Results {
		if result.HTML == "" {
			t.Errorf("expected non-empty html in result[%d]", i)
		}
	}
}

func TestMarkdownBatch_Sanitize(t *testing.T) {
	r := setupRouter()

	body := `{
		"markdowns": [
			"Hello <script>alert('xss')</script>",
			"<strong>safe?</strong>"
		],
		"sanitize": true
	}`

	req := httptest.NewRequest(
		http.MethodPost,
		"/markdown/batch",
		strings.NewReader(body),
	)

	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d: %s", w.Code, w.Body.String())
	}

	var resp httpx.Response[BatchResponse]

	if err := json.NewDecoder(w.Body).Decode(&resp); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	for i, result := range resp.Data.Results {
		if strings.Contains(result.HTML, "<script>") {
			t.Errorf(
				"expected script tag to be stripped in result[%d], got: %q",
				i,
				result.HTML,
			)
		}
	}
}

func TestMarkdownBatch_MissingBody(t *testing.T) { // check
	r := setupRouter()

	req := httptest.NewRequest(
		http.MethodPost,
		"/markdown/batch",
		http.NoBody,
	)

	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected 400, got %d", w.Code)
	}
}

func TestMarkdownBatch_EmptyMarkdowns(t *testing.T) {
	r := setupRouter()

	body := `{
		"markdowns": []
	}`

	req := httptest.NewRequest(
		http.MethodPost,
		"/markdown/batch",
		strings.NewReader(body),
	)

	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusUnprocessableEntity {
		t.Errorf("expected 422, got %d: %s", w.Code, w.Body.String())
	}
}

func TestMarkdownBatch_WithEmptyMarkdownItem(t *testing.T) {
	r := setupRouter()

	body := `{
		"markdowns": [
			"",
			"# Title"
		]
	}`

	req := httptest.NewRequest(
		http.MethodPost,
		"/markdown/batch",
		strings.NewReader(body),
	)

	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusUnprocessableEntity {
		t.Errorf("expected 422, got %d: %s", w.Code, w.Body.String())
	}
}

func TestMarkdownBatch_TooManyMarkdowns(t *testing.T) {
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
	if err != nil {
		t.Fatalf("failed to marshal payload: %v", err)
	}

	req := httptest.NewRequest(
		http.MethodPost,
		"/markdown/batch",
		strings.NewReader(string(bodyBytes)),
	)

	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusUnprocessableEntity {
		t.Errorf("expected 422, got %d: %s", w.Code, w.Body.String())
	}
}
