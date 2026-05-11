package trivia

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

func TestTrivia_NoFilters(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	req := httptest.NewRequest(http.MethodGet, "/trivia", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[Question]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	assert.NotEmpty(t, resp.Data.Question)
	assert.NotEmpty(t, resp.Data.Options)
	assert.NotEmpty(t, resp.Data.Answer)
	assert.NotEmpty(t, resp.Data.Category)
	assert.NotEmpty(t, resp.Data.Difficulty)
}

func TestTrivia_FilterByCategory(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	categories := []string{"science", "history", "geography", "sports", "music", "movies", "literature", "math", "technology", "nature"}
	for _, cat := range categories {
		t.Run(cat, func(t *testing.T) {
			t.Parallel()
			req := httptest.NewRequest(http.MethodGet, "/trivia?category="+cat, http.NoBody)
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			assert.Equal(t, http.StatusOK, w.Code, "expected status 200 for category %q, got %d", cat, w.Code)

			var resp httpx.Response[Question]
			err := json.NewDecoder(w.Body).Decode(&resp)
			require.NoError(t, err)

			assert.Equal(t, cat, resp.Data.Category)
		})
	}
}

func TestTrivia_FilterByDifficulty(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	difficulties := []string{"easy", "medium", "hard"}
	for _, d := range difficulties {
		t.Run(d, func(t *testing.T) {
			t.Parallel()
			req := httptest.NewRequest(http.MethodGet, "/trivia?difficulty="+d, http.NoBody)
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			assert.Equal(t, http.StatusOK, w.Code, "expected status 200 for difficulty %q, got %d", d, w.Code)

			var resp httpx.Response[Question]
			err := json.NewDecoder(w.Body).Decode(&resp)
			require.NoError(t, err)

			assert.Equal(t, d, resp.Data.Difficulty)
		})
	}
}

func TestTrivia_FilterByCategoryAndDifficulty(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	req := httptest.NewRequest(http.MethodGet, "/trivia?category=science&difficulty=easy", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[Question]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	assert.Equal(t, "science", resp.Data.Category)
	assert.Equal(t, "easy", resp.Data.Difficulty)
}

func TestTrivia_InvalidCategory(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	req := httptest.NewRequest(http.MethodGet, "/trivia?category=invalid", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestTrivia_InvalidDifficulty(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	req := httptest.NewRequest(http.MethodGet, "/trivia?difficulty=impossible", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestTrivia_AnswerIsInOptions(t *testing.T) {
	t.Parallel()
	for _, q := range questions {
		found := false
		for _, opt := range q.Options {
			if opt == q.Answer {
				found = true
				break
			}
		}
		assert.True(t, found, "question %q: answer %q is not in options %v", q.Question, q.Answer, q.Options)
	}
}

func TestTrivia_AllQuestionsHaveFourOptions(t *testing.T) {
	t.Parallel()
	for _, q := range questions {
		assert.Len(t, q.Options, 4, "question %q: expected 4 options, got %d", q.Question, len(q.Options))
	}
}

func TestService_Random_NoMatch(t *testing.T) {
	t.Parallel()
	svc := NewService()
	_, err := svc.Random("science", "impossible")
	assert.Error(t, err, "expected error for filters with no matching questions, got nil")
}
