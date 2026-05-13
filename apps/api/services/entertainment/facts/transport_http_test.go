package facts

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
