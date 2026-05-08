package thesaurus

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
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
