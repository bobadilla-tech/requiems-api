package facts

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"strings"
	"bytes"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"requiems-api/platform/httpx"
)

func setupRouter() chi.Router {
	r := chi.NewRouter()
	svc := NewService()
	RegisterRoutes(r, svc)
	return r
}

func TestFacts_RandomFact(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	req := httptest.NewRequest(http.MethodGet, "/facts", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[Fact]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)
	f := resp.Data

	assert.NotEmpty(t, f.Fact)
	assert.NotEmpty(t, f.Category)
	assert.NotEmpty(t, f.Source)
}

func TestFacts_CategoryFilter(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	categories := []string{"science", "history", "technology", "nature", "space", "food"}
	for _, cat := range categories {
		t.Run(cat, func(t *testing.T) {
			t.Parallel()
			req := httptest.NewRequest(http.MethodGet, "/facts?category="+cat, http.NoBody)
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			assert.Equal(t, http.StatusOK, w.Code, "expected status 200 for category %q, got %d", cat, w.Code)

			var resp httpx.Response[Fact]
			err := json.NewDecoder(w.Body).Decode(&resp)
			require.NoError(t, err)

			assert.Equal(t, cat, resp.Data.Category)
		})
	}
}

func TestFacts_InvalidCategory(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	req := httptest.NewRequest(http.MethodGet, "/facts?category=invalid", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestFacts_CategoryCaseInsensitive(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	req := httptest.NewRequest(http.MethodGet, "/facts?category=SCIENCE", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "expected status 200 for uppercase category, got %d", w.Code)

	var resp httpx.Response[Fact]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	assert.Equal(t, "science", resp.Data.Category)
}

func TestFacts_ServiceRandom(t *testing.T) {
	t.Parallel()
	svc := NewService()

	f, err := svc.Random("")
	require.NoError(t, err)
	assert.NotEmpty(t, f.Fact)
}

func TestFacts_ServiceRandomByCategory(t *testing.T) {
	t.Parallel()
	svc := NewService()

	f, err := svc.Random("science")
	require.NoError(t, err)
	assert.Equal(t, "science", f.Category)
}

func postBatch(t *testing.T, r chi.Router, body any) *httptest.ResponseRecorder {
	t.Helper()
	b, err := json.Marshal(body)
	require.NoError(t, err)
	req := httptest.NewRequest(http.MethodPost, "/facts/batch", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestBatch_AllValidCategories(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	w := postBatch(t, r, BatchRequest{
		Categories: []string{"science", "history", "technology", "nature", "space", "food"},
	})

	assert.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[httpx.BatchResponse[BatchItem]]
	require.NoError(t, json.NewDecoder(w.Body).Decode(&resp))
	require.Len(t, resp.Data.Results, 6)

	expected := []string{"science", "history", "technology", "nature", "space", "food"}
	for i, item := range resp.Data.Results {
		assert.Equal(t, expected[i], item.Category)
		assert.NotEmpty(t, item.Fact)
		assert.Empty(t, item.Error)
	}
}

func TestBatch_UsageCountHeader(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	w := postBatch(t, r, BatchRequest{
		Categories: []string{"science", "space", "food"},
	})

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "3", w.Header().Get("X-Usage-Count"))
}

func TestBatch_InvalidCategoryInBand(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	w := postBatch(t, r, BatchRequest{
		Categories: []string{"science", "dragons", "space"},
	})

	assert.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[httpx.BatchResponse[BatchItem]]
	require.NoError(t, json.NewDecoder(w.Body).Decode(&resp))
	require.Len(t, resp.Data.Results, 3)

	assert.Empty(t, resp.Data.Results[0].Error)
	assert.NotEmpty(t, resp.Data.Results[1].Error) // dragons — in-band
	assert.Empty(t, resp.Data.Results[1].Fact)
	assert.Empty(t, resp.Data.Results[2].Error)
}

func TestBatch_CaseInsensitive(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	w := postBatch(t, r, BatchRequest{Categories: []string{"SCIENCE", "History", "FOOD"}})

	assert.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[httpx.BatchResponse[BatchItem]]
	require.NoError(t, json.NewDecoder(w.Body).Decode(&resp))

	assert.Equal(t, "science", resp.Data.Results[0].Category)
	assert.Equal(t, "history", resp.Data.Results[1].Category)
	assert.Equal(t, "food", resp.Data.Results[2].Category)
}

func TestBatch_OrderPreserved(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	categories := []string{"food", "science", "history", "nature", "technology", "space"}
	w := postBatch(t, r, BatchRequest{Categories: categories})

	assert.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[httpx.BatchResponse[BatchItem]]
	require.NoError(t, json.NewDecoder(w.Body).Decode(&resp))

	for i, item := range resp.Data.Results {
		assert.Equal(t, categories[i], item.Category)
	}
}

func TestBatch_EmptyCategories_400(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	w := postBatch(t, r, BatchRequest{Categories: []string{}})
	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
}

func TestBatch_TooManyCategories_400(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	cats := make([]string, 51)
	for i := range cats {
		cats[i] = "science"
	}
	w := postBatch(t, r, BatchRequest{Categories: cats})
	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
}

func TestBatch_ExactlyMaxCategories_OK(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	cats := make([]string, 50)
	for i := range cats {
		cats[i] = "science"
	}
	w := postBatch(t, r, BatchRequest{Categories: cats})

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "50", w.Header().Get("X-Usage-Count"))
}

func TestBatch_MalformedJSON_400(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	req := httptest.NewRequest(http.MethodPost, "/facts/batch",
		strings.NewReader(`{not valid json`))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}