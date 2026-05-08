package words

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
	RegisterRoutes(r, &Service{db: &mockQuerier{}})
	return r
}

func TestDictionary_KnownWord(t *testing.T) {
	t.Parallel()
	req := httptest.NewRequest(http.MethodGet, "/dictionary/ephemeral", http.NoBody)
	w := httptest.NewRecorder()

	setupRouter().ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[DictionaryEntry]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	assert.Equal(t, "ephemeral", resp.Data.Word)
	assert.NotEmpty(t, resp.Data.Phonetic)
	assert.NotEmpty(t, resp.Data.Definitions)
	assert.NotEmpty(t, resp.Data.Definitions[0].PartOfSpeech)
	assert.NotEmpty(t, resp.Data.Definitions[0].Definition)
	assert.NotEmpty(t, resp.Data.Synonyms)
}

func TestDictionary_CaseInsensitive(t *testing.T) {
	t.Parallel()
	req := httptest.NewRequest(http.MethodGet, "/dictionary/EPHEMERAL", http.NoBody)
	w := httptest.NewRecorder()

	setupRouter().ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[DictionaryEntry]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	assert.Equal(t, "ephemeral", resp.Data.Word)
}

func TestDictionary_UnknownWord(t *testing.T) {
	t.Parallel()
	req := httptest.NewRequest(http.MethodGet, "/dictionary/zzyzx", http.NoBody)
	w := httptest.NewRecorder()

	setupRouter().ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestDictionary_ExampleField(t *testing.T) {
	t.Parallel()
	req := httptest.NewRequest(http.MethodGet, "/dictionary/ephemeral", http.NoBody)
	w := httptest.NewRecorder()

	setupRouter().ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[DictionaryEntry]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	assert.NotEmpty(t, resp.Data.Definitions[0].Example)
}

func TestDictionary_MultipleDefinitions(t *testing.T) {
	t.Parallel()
	// melancholy has two definitions (noun and adjective)
	req := httptest.NewRequest(http.MethodGet, "/dictionary/melancholy", http.NoBody)
	w := httptest.NewRecorder()

	setupRouter().ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[DictionaryEntry]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	assert.True(t, len(resp.Data.Definitions) >= 2, "expected at least 2 definitions, got %d", len(resp.Data.Definitions))
}
