package chucknorris

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

func TestChuckNorris_Random(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	req := httptest.NewRequest(http.MethodGet, "/chuck-norris", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[Fact]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	f := resp.Data

	assert.NotEmpty(t, f.Fact)
	assert.True(t, strings.HasPrefix(f.ID, "cn_"), "expected ID to start with 'cn_', got %q", f.ID)
}

func TestChuckNorris_Randomness(t *testing.T) {
	t.Parallel()
	svc := NewService()

	seen := make(map[string]bool)
	for range 50 {
		f := svc.Random()
		seen[f.ID] = true
	}

	// With 30 facts and 50 draws, expect at least 5 distinct facts.
	assert.GreaterOrEqual(t, len(seen), 5, "expected variety in random facts, got only %d distinct IDs in 50 draws", len(seen))
}

func TestChuckNorris_FactsNonEmpty(t *testing.T) {
	t.Parallel()
	svc := NewService()
	for i := range 10 {
		f := svc.Random()
		assert.NotEmpty(t, f.Fact, "call %d: expected non-empty fact", i)
		assert.NotEmpty(t, f.ID, "call %d: expected non-empty ID", i)
	}
}
