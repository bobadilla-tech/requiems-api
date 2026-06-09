package thesaurus

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

func TestThesaurus_KnownWord(t *testing.T) {
	t.Parallel()
	req := httptest.NewRequest(http.MethodGet, "/thesaurus/happy", http.NoBody)
	w := httptest.NewRecorder()

	setupRouter().ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[Result]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	assert.Equal(t, "happy", resp.Data.Word)
	assert.NotEmpty(t, resp.Data.Synonyms)
	assert.NotEmpty(t, resp.Data.Antonyms)
}

func TestThesaurus_CaseInsensitive(t *testing.T) {
	t.Parallel()
	req := httptest.NewRequest(http.MethodGet, "/thesaurus/HAPPY", http.NoBody)
	w := httptest.NewRecorder()

	setupRouter().ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[Result]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	assert.Equal(t, "happy", resp.Data.Word)
}

func TestThesaurus_UnknownWord(t *testing.T) {
	t.Parallel()
	req := httptest.NewRequest(http.MethodGet, "/thesaurus/zzyzx", http.NoBody)
	w := httptest.NewRecorder()

	setupRouter().ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestThesaurus_Batch_KnownWords(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	req := httptest.NewRequest(http.MethodPost, "/thesaurus/batch", strings.NewReader(`{"words":["happy","fast"]}`))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[httpx.BatchResponse[BatchThesaurusItem]]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	assert.Len(t, resp.Data.Results, 2)
	assert.Equal(t, 2, resp.Data.Total)
	assert.Equal(t, "happy", resp.Data.Results[0].Word)
	assert.NotNil(t, resp.Data.Results[0].Result)
	assert.Empty(t, resp.Data.Results[0].Error)
}

func TestThesaurus_Batch_UnknownWord_InBandError(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	req := httptest.NewRequest(http.MethodPost, "/thesaurus/batch", strings.NewReader(`{"words":["happy","zzyzxnotaword"]}`))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[httpx.BatchResponse[BatchThesaurusItem]]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	assert.Len(t, resp.Data.Results, 2)
	assert.NotNil(t, resp.Data.Results[0].Result, "known word should have result")
	assert.NotEmpty(t, resp.Data.Results[1].Error, "unknown word should have in-band error")
	assert.Nil(t, resp.Data.Results[1].Result)
}

func TestThesaurus_Batch_EmptyWords422(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	req := httptest.NewRequest(http.MethodPost, "/thesaurus/batch", strings.NewReader(`{"words":[]}`))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
}
