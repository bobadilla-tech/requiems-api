package spellcheck

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

func setupRouter(ltURL string) chi.Router {
	r := chi.NewRouter()
	RegisterRoutes(r, NewService(ltURL))
	return r
}

// noopLT returns a mock LanguageTool server that always responds with no matches.
// Use it for tests that only check validation errors (422, 400) where LanguageTool
// is never actually reached.
func noopLT(t *testing.T) *httptest.Server {
	t.Helper()
	return newMockLT(t, []ltMatch{})
}

// ── /spellcheck ──────────────────────────────────────────────────────────────

func TestSpellcheck_CleanText(t *testing.T) {
	t.Parallel()
	lt := newMockLT(t, []ltMatch{})
	defer lt.Close()
	r := setupRouter(lt.URL)

	body := `{"text":"Hello world"}`
	req := httptest.NewRequest(http.MethodPost, "/spellcheck", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[Result]
	require.NoError(t, json.NewDecoder(w.Body).Decode(&resp))
	assert.Empty(t, resp.Data.Corrections)
}

func TestSpellcheck_MisspelledText(t *testing.T) {
	t.Parallel()
	lt := newMockLT(t, []ltMatch{
		{Offset: 0, Length: 3, Replacements: []ltReplacement{{Value: "This"}}},
		{Offset: 9, Length: 4, Replacements: []ltReplacement{{Value: "test"}}},
	})
	defer lt.Close()
	r := setupRouter(lt.URL)

	body := `{"text":"Ths is a tset"}`
	req := httptest.NewRequest(http.MethodPost, "/spellcheck", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[Result]
	require.NoError(t, json.NewDecoder(w.Body).Decode(&resp))
	assert.NotEmpty(t, resp.Data.Corrections)
	assert.NotEqual(t, "Ths is a tset", resp.Data.Corrected)
}

func TestSpellcheck_MissingTextField(t *testing.T) {
	t.Parallel()
	lt := noopLT(t)
	defer lt.Close()
	r := setupRouter(lt.URL)

	req := httptest.NewRequest(http.MethodPost, "/spellcheck", strings.NewReader(`{}`))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
}

func TestSpellcheck_MissingBody(t *testing.T) {
	t.Parallel()
	lt := noopLT(t)
	defer lt.Close()
	r := setupRouter(lt.URL)

	req := httptest.NewRequest(http.MethodPost, "/spellcheck", http.NoBody)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

// ── /spellcheck/batch ────────────────────────────────────────────────────────

func TestSpellcheckBatch_MixedTexts(t *testing.T) {
	t.Parallel()
	lt := newMockLT(t, []ltMatch{
		{Offset: 0, Length: 3, Replacements: []ltReplacement{{Value: "This"}}},
	})
	defer lt.Close()
	r := setupRouter(lt.URL)

	body := `{"texts":["Hello world","Ths is a tset","Simple test"]}`
	req := httptest.NewRequest(http.MethodPost, "/spellcheck/batch", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[BatchCheckResponse]
	require.NoError(t, json.NewDecoder(w.Body).Decode(&resp))
	assert.Equal(t, 3, resp.Data.Total)
	require.Len(t, resp.Data.Results, 3)
}

func TestSpellcheckBatch_OrderPreserved(t *testing.T) {
	t.Parallel()
	lt := newMockLT(t, []ltMatch{})
	defer lt.Close()
	r := setupRouter(lt.URL)

	// RFC requires results to stay in the same order as the input array.
	body := `{"texts":["Hello world","Clean text","More clean text"]}`
	req := httptest.NewRequest(http.MethodPost, "/spellcheck/batch", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[BatchCheckResponse]
	require.NoError(t, json.NewDecoder(w.Body).Decode(&resp))
	require.Len(t, resp.Data.Results, 3)
}

func TestSpellcheckBatch_UsageCountHeader(t *testing.T) {
	t.Parallel()
	lt := newMockLT(t, []ltMatch{})
	defer lt.Close()
	r := setupRouter(lt.URL)

	body := `{"texts":["Hello world","Clean text","More text"]}`
	req := httptest.NewRequest(http.MethodPost, "/spellcheck/batch", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)
	// The gateway reads this header to charge per item instead of per request.
	assert.Equal(t, "3", w.Header().Get("X-Usage-Count"))
}

func TestSpellcheckBatch_EmptyArray(t *testing.T) {
	t.Parallel()
	lt := noopLT(t)
	defer lt.Close()
	r := setupRouter(lt.URL)

	req := httptest.NewRequest(http.MethodPost, "/spellcheck/batch", strings.NewReader(`{"texts":[]}`))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
}

func TestSpellcheckBatch_ExceedsLimit(t *testing.T) {
	t.Parallel()
	lt := noopLT(t)
	defer lt.Close()
	r := setupRouter(lt.URL)

	// 51 texts — one over the maximum of 50.
	texts := make([]string, 51)
	for i := range texts {
		texts[i] = `"hello world"`
	}
	body := `{"texts":[` + strings.Join(texts, ",") + `]}`
	req := httptest.NewRequest(http.MethodPost, "/spellcheck/batch", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
}

func TestSpellcheckBatch_MissingTextsField(t *testing.T) {
	t.Parallel()
	lt := noopLT(t)
	defer lt.Close()
	r := setupRouter(lt.URL)

	req := httptest.NewRequest(http.MethodPost, "/spellcheck/batch", strings.NewReader(`{}`))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
}

func TestSpellcheckBatch_MissingBody(t *testing.T) {
	t.Parallel()
	lt := noopLT(t)
	defer lt.Close()
	r := setupRouter(lt.URL)

	req := httptest.NewRequest(http.MethodPost, "/spellcheck/batch", http.NoBody)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}
