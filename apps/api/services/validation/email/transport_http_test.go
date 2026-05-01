package email

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

// ---- helpers ----

func postValidate(body string) *httptest.ResponseRecorder {
	req := httptest.NewRequest(http.MethodPost, "/email", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	setupRouter().ServeHTTP(w, req)
	return w
}

func postBatch(body string) *httptest.ResponseRecorder {
	req := httptest.NewRequest(http.MethodPost, "/email/batch", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	setupRouter().ServeHTTP(w, req)
	return w
}

func decodeBatchResponse(t *testing.T, w *httptest.ResponseRecorder) httpx.Response[BatchResponse] {
	t.Helper()
	var resp httpx.Response[BatchResponse]
	if err := json.NewDecoder(w.Body).Decode(&resp); err != nil {
		t.Fatalf("failed to decode batch response: %v", err)
	}
	return resp
}


// ---- single endpoint tests (unchanged) ----
func TestValidate_MissingEmail(t *testing.T) {
	w := postValidate(`{}`)

	// httpx.Handle returns 422 for failed struct validation.
	if w.Code != http.StatusUnprocessableEntity {
		t.Errorf("expected 422, got %d: %s", w.Code, w.Body.String())
	}
}

func TestValidate_InvalidSyntax(t *testing.T) {
	w := postValidate(`{"email":"notanemail"}`)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d: %s", w.Code, w.Body.String())
	}

	var resp httpx.Response[Validation]
	if err := json.NewDecoder(w.Body).Decode(&resp); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if resp.Data.Valid {
		t.Error("expected Valid=false for invalid syntax")
	}
	if resp.Data.SyntaxValid {
		t.Error("expected SyntaxValid=false for invalid syntax")
	}
}

func TestValidate_ValidEmail(t *testing.T) {
	w := postValidate(`{"email":"user@gmail.com"}`)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d: %s", w.Code, w.Body.String())
	}

	var resp httpx.Response[Validation]
	if err := json.NewDecoder(w.Body).Decode(&resp); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if !resp.Data.SyntaxValid {
		t.Error("expected SyntaxValid=true")
	}
	if *resp.Data.Email != "user@gmail.com" {
		t.Errorf("expected Email=user@gmail.com, got %q", *resp.Data.Email)
	}
	if *resp.Data.Domain != "gmail.com" {
		t.Errorf("expected Domain=gmail.com, got %q", *resp.Data.Domain)
	}
}

func TestValidate_SuggestionPresentInResponse(t *testing.T) {
	w := postValidate(`{"email":"user@gmial.com"}`)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d: %s", w.Code, w.Body.String())
	}

	// Decode into a raw map to verify "suggestion" is always present, even when null.
	var raw map[string]json.RawMessage
	if err := json.NewDecoder(w.Body).Decode(&raw); err != nil {
		t.Fatalf("failed to decode: %v", err)
	}

	var data map[string]json.RawMessage
	if err := json.Unmarshal(raw["data"], &data); err != nil {
		t.Fatalf("failed to decode data: %v", err)
	}

	if _, ok := data["suggestion"]; !ok {
		t.Error("expected 'suggestion' key to be present in response (even when null)")
	}

	var suggestion *string
	if err := json.Unmarshal(data["suggestion"], &suggestion); err != nil {
		t.Fatalf("failed to decode suggestion: %v", err)
	}
	if suggestion == nil {
		t.Fatal("expected non-nil suggestion for gmial.com")
	}
	if *suggestion != "gmail.com" {
		t.Errorf("expected suggestion=gmail.com, got %q", *suggestion)
	}
}

func TestValidate_SuggestionNullForKnownDomain(t *testing.T) {
	w := postValidate(`{"email":"user@gmail.com"}`)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}

	var raw map[string]json.RawMessage
	if err := json.NewDecoder(w.Body).Decode(&raw); err != nil {
		t.Fatalf("failed to decode: %v", err)
	}

	var data map[string]json.RawMessage
	if err := json.Unmarshal(raw["data"], &data); err != nil {
		t.Fatalf("failed to decode data: %v", err)
	}

	if _, ok := data["suggestion"]; !ok {
		t.Error("expected 'suggestion' key to always be present")
	}

	if string(data["suggestion"]) != "null" {
		t.Errorf("expected suggestion=null for gmail.com, got %s", data["suggestion"])
	}
}

// ---- batch endpoint tests ----

func TestBatch_HappyPath_Returns200(t *testing.T) {
	w := postBatch(`{"emails":["user@gmail.com","user@yahoo.com"]}`)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d: %s", w.Code, w.Body.String())
	}

	resp := decodeBatchResponse(t, w)

	if resp.Data.Total != 2 {
		t.Errorf("total = %d, want 2", resp.Data.Total)
	}

	if len(resp.Data.Results) != 2 {
		t.Fatalf("results len = %d, want 2", len(resp.Data.Results))
	}

	for i, item := range resp.Data.Results {
		if !item.SyntaxValid {
			t.Errorf("results[%d].SyntaxValid = false, want true", i)
		}
		if item.Email == nil {
			t.Errorf("results[%d].Email = nil, want non-nil", i)
		}
	}
}

func TestBatch_InvalidSyntax_InBand(t *testing.T) {
	w := postBatch(`{"emails":["notanemail"]}`)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}

	item := decodeBatchResponse(t, w).Data.Results[0]

	if item.Valid {
		t.Error("Valid = true, want false")
	}
	if item.SyntaxValid {
		t.Error("SyntaxValid = true, want false")
	}
	if item.Email != nil {
		t.Error("Email != nil, want nil")
	}
}

func TestBatch_OrderPreserved(t *testing.T) {
	w := postBatch(`{"emails":["a@gmail.com","b@gmail.com","c@gmail.com"]}`)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}

	results := decodeBatchResponse(t, w).Data.Results

	expected := []string{"a@gmail.com", "b@gmail.com", "c@gmail.com"}

	for i := range results {
		if results[i].Email == nil {
			t.Fatalf("results[%d].Email is nil", i)
		}
		if *results[i].Email != expected[i] {
			t.Errorf("results[%d] = %s, want %s", i, *results[i].Email, expected[i])
		}
	}
}

func TestBatch_EmptyArray(t *testing.T) {
	w := postBatch(`{"emails":[]}`)

	if w.Code != http.StatusUnprocessableEntity {
		t.Fatalf("expected 422, got %d: %s", w.Code, w.Body.String())
	}
}

func TestBatch_MissingEmailsField(t *testing.T) {
	w := postBatch(`{}`)

	if w.Code != http.StatusUnprocessableEntity {
		t.Fatalf("expected 422, got %d: %s", w.Code, w.Body.String())
	}
}

func TestBatch_OverLimit(t *testing.T) {
	emails := make([]string, 51)
	for i := range emails {
		emails[i] = `"user@gmail.com"`
	}

	body := `{"emails":[` + strings.Join(emails, ",") + `]}`
	w := postBatch(body)

	if w.Code != http.StatusUnprocessableEntity {
		t.Fatalf("expected 422, got %d: %s", w.Code, w.Body.String())
	}
}
func TestBatch_Mixed_RealWorldScenario(t *testing.T) {
	w := postBatch(`{
		"emails":[
			"user@gmail.com",
			"user@gmial.com",
			"notanemail",
			"user@yahoo.com"
		]
	}`)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}

	resp := decodeBatchResponse(t, w)

	if resp.Data.Total != 4 {
		t.Fatalf("expected 4 results, got %d", resp.Data.Total)
	}

	results := resp.Data.Results

	if !results[0].SyntaxValid {
		t.Error("gmail should be valid syntax")
	}

	if results[1].Suggestion == nil {
		t.Error("expected suggestion for gmial.com")
	}

	if results[2].SyntaxValid {
		t.Error("expected invalid syntax for notanemail")
	}

	if !results[3].SyntaxValid {
		t.Error("yahoo should be valid syntax")
	}
}
func TestBatch_DuplicateEmails(t *testing.T) {
	w := postBatch(`{
		"emails":[
			"user@gmail.com",
			"user@gmail.com",
			"user@gmail.com"
		]
	}`)

	resp := decodeBatchResponse(t, w)

	if resp.Data.Total != 3 {
		t.Errorf("expected 3 results, got %d", resp.Data.Total)
	}

	for _, r := range resp.Data.Results {
		if r.Email == nil {
			t.Error("expected email not nil for duplicates")
		}
	}
}