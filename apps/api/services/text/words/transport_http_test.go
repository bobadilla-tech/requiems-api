package words

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

func TestWordsBatch_ValidRequest(t *testing.T) {
	t.Parallel()

	r := setupRouter()

	body := `{
		"items": ["ephemeral", "serendipity", "melancholy"]
	}`

	req := httptest.NewRequest(http.MethodPost, "/words/batch", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[BatchResponse]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	assert.Equal(t, 3, resp.Data.Total)
	assert.Len(t, resp.Data.Results, 3)

	for _, item := range resp.Data.Results {
		assert.NotEmpty(t, item.Word)
		assert.True(t, item.Found)
		assert.NotNil(t, item.Entry)
		assert.Empty(t, item.Error)
	}
}

func TestWordsBatch_CaseInsensitive(t *testing.T) {
	t.Parallel()

	r := setupRouter()

	body := `{
		"items": ["EPHEMERAL", "SeReNdiPiTy"]
	}`

	req := httptest.NewRequest(http.MethodPost, "/words/batch", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[BatchResponse]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	assert.Equal(t, 2, resp.Data.Total)

	for _, item := range resp.Data.Results {
		assert.True(t, item.Found)
		assert.NotNil(t, item.Entry)
	}
}

func TestWordsBatch_PartialFailure(t *testing.T) {
	t.Parallel()

	r := setupRouter()

	body := `{
		"items": ["ephemeral", "fakeword123", "serendipity"]
	}`

	req := httptest.NewRequest(http.MethodPost, "/words/batch", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[BatchResponse]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	assert.Equal(t, 3, resp.Data.Total)
	assert.Len(t, resp.Data.Results, 3)

	assert.True(t, resp.Data.Results[0].Found)
	assert.False(t, resp.Data.Results[1].Found)
	assert.NotEmpty(t, resp.Data.Results[1].Error)
	assert.True(t, resp.Data.Results[2].Found)
}

func TestWordsBatch_EmptyItems(t *testing.T) {
	t.Parallel()

	r := setupRouter()

	body := `{"items":[]}`

	req := httptest.NewRequest(http.MethodPost, "/words/batch", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
}

func TestWordsBatch_MissingBody(t *testing.T) {
	t.Parallel()

	r := setupRouter()

	req := httptest.NewRequest(http.MethodPost, "/words/batch", strings.NewReader(`{}`))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
}

func TestWordsBatch_UnknownWords(t *testing.T) {
	t.Parallel()

	r := setupRouter()

	body := `{
		"items": ["notaword1", "notaword2"]
	}`

	req := httptest.NewRequest(http.MethodPost, "/words/batch", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[BatchResponse]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	for _, item := range resp.Data.Results {
		assert.False(t, item.Found)
		assert.NotEmpty(t, item.Error)
		assert.Nil(t, item.Entry)
	}
}
func TestWordsBatch_TooManyItems(t *testing.T) {
	t.Parallel()

	r := setupRouter()

	items := make([]string, 51)
	for i := 0; i < 51; i++ {
		items[i] = "ephemeral"
	}

	bodyBytes, _ := json.Marshal(BatchRequest{Items: items})

	req := httptest.NewRequest(http.MethodPost, "/words/batch", strings.NewReader(string(bodyBytes)))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusUnprocessableEntity, w.Code)

	// opcional: solo validar que devuelve error HTTP correcto
	assert.Contains(t, w.Body.String(), "error")
}
