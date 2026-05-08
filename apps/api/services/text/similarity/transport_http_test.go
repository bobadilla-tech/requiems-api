package similarity

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

func TestSimilarity_IdenticalTexts(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	body := `{"text1":"The cat sat on the mat","text2":"The cat sat on the mat"}`
	req := httptest.NewRequest(http.MethodPost, "/similarity", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[Result]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	assert.Equal(t, 1.0, resp.Data.Similarity)
	assert.Equal(t, "cosine", resp.Data.Method)
}

func TestSimilarity_UnrelatedTexts(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	body := `{"text1":"The cat sat on the mat","text2":"quantum physics nuclear reactor"}`
	req := httptest.NewRequest(http.MethodPost, "/similarity", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[Result]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	assert.Equal(t, 0.0, resp.Data.Similarity)
}

func TestSimilarity_SimilarTexts(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	body := `{"text1":"The cat sat on the mat","text2":"A cat was sitting on a mat"}`
	req := httptest.NewRequest(http.MethodPost, "/similarity", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[Result]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	// Texts share words (cat, on, mat), expect non-zero similarity.
	assert.True(t, resp.Data.Similarity > 0 && resp.Data.Similarity < 1, "expected similarity between 0 and 1 (exclusive), got %f", resp.Data.Similarity)
}

func TestSimilarity_MissingText1(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	req := httptest.NewRequest(http.MethodPost, "/similarity", strings.NewReader(`{"text2":"hello"}`))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
}

func TestSimilarity_MissingText2(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	req := httptest.NewRequest(http.MethodPost, "/similarity", strings.NewReader(`{"text1":"hello"}`))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
}

func TestSimilarity_MissingBody(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	req := httptest.NewRequest(http.MethodPost, "/similarity", http.NoBody)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}
