package email

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
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)
	return resp
}

// ---- single endpoint tests (unchanged) ----
func TestValidate_MissingEmail(t *testing.T) {
	t.Parallel()
	w := postValidate(`{}`)

	// httpx.Handle returns 422 for failed struct validation.
	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
}

func TestValidate_InvalidSyntax(t *testing.T) {
	t.Parallel()
	w := postValidate(`{"email":"notanemail"}`)

	require.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[Validation]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	assert.False(t, resp.Data.Valid)
	assert.False(t, resp.Data.SyntaxValid)
}

func TestValidate_ValidEmail(t *testing.T) {
	t.Parallel()
	w := postValidate(`{"email":"user@gmail.com"}`)

	require.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[Validation]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	assert.True(t, resp.Data.SyntaxValid)
	assert.Equal(t, "user@gmail.com", *resp.Data.Email)
	assert.Equal(t, "gmail.com", *resp.Data.Domain)
}

func TestValidate_SuggestionPresentInResponse(t *testing.T) {
	t.Parallel()
	w := postValidate(`{"email":"user@gmial.com"}`)

	require.Equal(t, http.StatusOK, w.Code)

	// Decode into a raw map to verify "suggestion" is always present, even when null.
	var raw map[string]json.RawMessage
	err := json.NewDecoder(w.Body).Decode(&raw)
	require.NoError(t, err)

	var data map[string]json.RawMessage
	err = json.Unmarshal(raw["data"], &data)
	require.NoError(t, err)

	if _, ok := data["suggestion"]; !ok {
		t.Error("expected 'suggestion' key to be present in response (even when null)")
	}

	var suggestion *string
	err = json.Unmarshal(data["suggestion"], &suggestion)
	require.NoError(t, err)
	require.NotNil(t, suggestion)
	assert.Equal(t, "gmail.com", *suggestion)
}

func TestValidate_SuggestionNullForKnownDomain(t *testing.T) {
	t.Parallel()
	w := postValidate(`{"email":"user@gmail.com"}`)

	require.Equal(t, http.StatusOK, w.Code)

	var raw map[string]json.RawMessage
	err := json.NewDecoder(w.Body).Decode(&raw)
	require.NoError(t, err)

	var data map[string]json.RawMessage
	err = json.Unmarshal(raw["data"], &data)
	require.NoError(t, err)

	if _, ok := data["suggestion"]; !ok {
		t.Error("expected 'suggestion' key to always be present")
	}

	if string(data["suggestion"]) != "null" {
		t.Errorf("expected suggestion=null for gmail.com, got %s", data["suggestion"])
	}
}

// ---- batch endpoint tests ----

func TestBatch_HappyPath_Returns200(t *testing.T) {
	t.Parallel()
	w := postBatch(`{"emails":["user@gmail.com","user@yahoo.com"]}`)

	require.Equal(t, http.StatusOK, w.Code)

	resp := decodeBatchResponse(t, w)

	assert.Equal(t, 2, resp.Data.Total)

	require.Len(t, resp.Data.Results, 2)

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
	t.Parallel()
	w := postBatch(`{"emails":["notanemail"]}`)

	require.Equal(t, http.StatusOK, w.Code)

	item := decodeBatchResponse(t, w).Data.Results[0]

	assert.False(t, item.Valid)
	assert.False(t, item.SyntaxValid)
	assert.Nil(t, item.Email)
}

func TestBatch_OrderPreserved(t *testing.T) {
	t.Parallel()
	w := postBatch(`{"emails":["a@gmail.com","b@gmail.com","c@gmail.com"]}`)

	require.Equal(t, http.StatusOK, w.Code)

	results := decodeBatchResponse(t, w).Data.Results

	expected := []string{"a@gmail.com", "b@gmail.com", "c@gmail.com"}

	for i := range results {
		require.NotNil(t, results[i].Email)
		assert.Equal(t, expected[i], *results[i].Email)
	}
}

func TestBatch_EmptyArray(t *testing.T) {
	t.Parallel()
	w := postBatch(`{"emails":[]}`)

	require.Equal(t, http.StatusUnprocessableEntity, w.Code)
}

func TestBatch_MissingEmailsField(t *testing.T) {
	t.Parallel()
	w := postBatch(`{}`)

	require.Equal(t, http.StatusUnprocessableEntity, w.Code)
}

func TestBatch_OverLimit(t *testing.T) {
	t.Parallel()
	emails := make([]string, 51)
	for i := range emails {
		emails[i] = `"user@gmail.com"`
	}

	body := `{"emails":[` + strings.Join(emails, ",") + `]}`
	w := postBatch(body)

	require.Equal(t, http.StatusUnprocessableEntity, w.Code)
}

func TestBatch_Mixed_RealWorldScenario(t *testing.T) {
	t.Parallel()
	w := postBatch(`{
		"emails":[
			"user@gmail.com",
			"user@gmial.com",
			"notanemail",
			"user@yahoo.com"
		]
	}`)

	require.Equal(t, http.StatusOK, w.Code)

	resp := decodeBatchResponse(t, w)

	require.Equal(t, 4, resp.Data.Total)

	results := resp.Data.Results

	assert.True(t, results[0].SyntaxValid)

	assert.NotNil(t, results[1].Suggestion)

	assert.False(t, results[2].SyntaxValid)

	assert.True(t, results[3].SyntaxValid)
}

func TestBatch_DuplicateEmails(t *testing.T) {
	t.Parallel()
	w := postBatch(`{
		"emails":[
			"user@gmail.com",
			"user@gmail.com",
			"user@gmail.com"
		]
	}`)

	resp := decodeBatchResponse(t, w)

	assert.Equal(t, 3, resp.Data.Total)

	for _, r := range resp.Data.Results {
		if r.Email == nil {
			t.Error("expected email not nil for duplicates")
		}
	}
}
