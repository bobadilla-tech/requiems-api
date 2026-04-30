package normalize

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

func TestNormalize_HappyPath(t *testing.T) {
	r := setupRouter()

	body := `{"email":"user@example.com"}`
	req := httptest.NewRequest(http.MethodPost, "/normalize", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d: %s", w.Code, w.Body.String())
	}

	var resp httpx.Response[EmailNormalization]
	if err := json.NewDecoder(w.Body).Decode(&resp); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if resp.Data.Normalized == "" {
		t.Error("expected non-empty Normalized field")
	}
	if resp.Data.Original != "user@example.com" {
		t.Errorf("expected Original %q, got %q", "user@example.com", resp.Data.Original)
	}
	if resp.Data.Local != "user" {
		t.Errorf("expected Local %q, got %q", "user", resp.Data.Local)
	}
	if resp.Data.Domain != "example.com" {
		t.Errorf("expected Domain %q, got %q", "example.com", resp.Data.Domain)
	}
}

func TestNormalize_GmailNormalization(t *testing.T) {
	r := setupRouter()

	body := `{"email":"te.st.user+spam@gmail.com"}`
	req := httptest.NewRequest(http.MethodPost, "/normalize", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d: %s", w.Code, w.Body.String())
	}

	var resp httpx.Response[EmailNormalization]
	if err := json.NewDecoder(w.Body).Decode(&resp); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if resp.Data.Normalized != "testuser@gmail.com" {
		t.Errorf("expected normalized %q, got %q", "testuser@gmail.com", resp.Data.Normalized)
	}
	if len(resp.Data.Changes) == 0 {
		t.Error("expected at least one change for gmail normalization")
	}
}

func TestNormalize_UppercaseDomainLowercased(t *testing.T) {
	r := setupRouter()

	// For unknown providers the local part is preserved (case-sensitive per
	// RFC 5321); only the domain is lowercased.
	body := `{"email":"USER@EXAMPLE.COM"}`
	req := httptest.NewRequest(http.MethodPost, "/normalize", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d: %s", w.Code, w.Body.String())
	}

	var resp httpx.Response[EmailNormalization]
	if err := json.NewDecoder(w.Body).Decode(&resp); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if resp.Data.Normalized != "USER@example.com" {
		t.Errorf("expected normalized %q, got %q", "USER@example.com", resp.Data.Normalized)
	}
}

func TestNormalize_OriginalIsAlwaysUnmodifiedInput(t *testing.T) {
	r := setupRouter()

	input := "Test.User+tag@Gmail.com"
	body := `{"email":"` + input + `"}`
	req := httptest.NewRequest(http.MethodPost, "/normalize", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d: %s", w.Code, w.Body.String())
	}

	var resp httpx.Response[EmailNormalization]
	if err := json.NewDecoder(w.Body).Decode(&resp); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if resp.Data.Original != input {
		t.Errorf("expected Original %q, got %q", input, resp.Data.Original)
	}
}

func TestNormalize_MissingEmailField(t *testing.T) {
	r := setupRouter()

	req := httptest.NewRequest(http.MethodPost, "/normalize", strings.NewReader(`{}`))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusUnprocessableEntity {
		t.Errorf("expected 422, got %d: %s", w.Code, w.Body.String())
	}
}

func TestNormalize_InvalidEmailFormat(t *testing.T) {
	r := setupRouter()

	body := `{"email":"not-an-email"}`
	req := httptest.NewRequest(http.MethodPost, "/normalize", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected 400, got %d: %s", w.Code, w.Body.String())
	}
}

func TestNormalize_MissingBody(t *testing.T) {
	r := setupRouter()

	req := httptest.NewRequest(http.MethodPost, "/normalize", http.NoBody)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected 400, got %d: %s", w.Code, w.Body.String())
	}
}

func TestNormalize_UnknownFieldsRejected(t *testing.T) {
	r := setupRouter()

	body := `{"email":"user@example.com","unexpected_field":"value"}`
	req := httptest.NewRequest(http.MethodPost, "/normalize", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected 400 for unknown fields, got %d: %s", w.Code, w.Body.String())
	}
}

func TestNormalizeBatch_HappyPathSetsUsageCount(t *testing.T) {
	r := setupRouter()

	body := `{"emails":["a@b.co","user@example.com"]}`
	req := httptest.NewRequest(http.MethodPost, "/normalize/batch", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d: %s", w.Code, w.Body.String())
	}
	if got := w.Header().Get("X-Usage-Count"); got != "2" {
		t.Errorf("X-Usage-Count: want 2, got %q", got)
	}

	var resp httpx.Response[EmailNormalizationBatchResponse]
	if err := json.NewDecoder(w.Body).Decode(&resp); err != nil {
		t.Fatalf("decode: %v", err)
	}
	if resp.Data.Total != 2 || len(resp.Data.Results) != 2 {
		t.Fatalf("total/results: %+v", resp.Data)
	}
	if !resp.Data.Results[0].Valid || !resp.Data.Results[1].Valid {
		t.Fatalf("expected both valid, got %+v", resp.Data.Results)
	}
}

func TestNormalizeBatch_InvalidItemInBand(t *testing.T) {
	r := setupRouter()

	body := `{"emails":["user@example.com","not-an-email"]}`
	req := httptest.NewRequest(http.MethodPost, "/normalize/batch", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d: %s", w.Code, w.Body.String())
	}

	var resp httpx.Response[EmailNormalizationBatchResponse]
	if err := json.NewDecoder(w.Body).Decode(&resp); err != nil {
		t.Fatalf("decode: %v", err)
	}
	if !resp.Data.Results[0].Valid || resp.Data.Results[1].Valid {
		t.Fatalf("want first valid second invalid, got %+v", resp.Data.Results)
	}
}

func TestNormalizeBatch_EmptyArrayValidation(t *testing.T) {
	r := setupRouter()

	body := `{"emails":[]}`
	req := httptest.NewRequest(http.MethodPost, "/normalize/batch", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusUnprocessableEntity {
		t.Errorf("expected 422, got %d: %s", w.Code, w.Body.String())
	}
}
