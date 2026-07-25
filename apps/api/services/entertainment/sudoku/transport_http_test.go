package sudoku

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

func TestSudoku_DefaultDifficulty(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	req := httptest.NewRequest(http.MethodGet, "/sudoku", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[Puzzle]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	assert.Equal(t, "medium", resp.Data.Difficulty)
}

func TestSudoku_AllDifficulties(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	difficulties := []string{"easy", "medium", "hard"}

	for _, d := range difficulties {
		t.Run(d, func(t *testing.T) {
			t.Parallel()
			req := httptest.NewRequest(http.MethodGet, "/sudoku?difficulty="+d, http.NoBody)
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			assert.Equal(t, http.StatusOK, w.Code, "expected status 200 for difficulty %q, got %d", d, w.Code)

			var resp httpx.Response[Puzzle]
			err := json.NewDecoder(w.Body).Decode(&resp)
			require.NoError(t, err)

			assert.Equal(t, d, resp.Data.Difficulty)
		})
	}
}

func TestSudoku_InvalidDifficulty(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	req := httptest.NewRequest(http.MethodGet, "/sudoku?difficulty=impossible", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestSudoku_PuzzleHasEmptyCells(t *testing.T) {
	t.Parallel()
	svc := NewService()

	for _, d := range []string{"easy", "medium", "hard"} {
		t.Run(d, func(t *testing.T) {
			t.Parallel()
			p, err := svc.Generate(d)
			require.NoError(t, err)

			empty := 0
			for r := range 9 {
				for c := range 9 {
					if p.Puzzle[r][c] == 0 {
						empty++
					}
				}
			}

			expectedEmpty := map[string]int{"easy": 36, "medium": 46, "hard": 52}
			assert.Equal(t, expectedEmpty[d], empty, "difficulty %q: expected %d empty cells, got %d", d, expectedEmpty[d], empty)
		})
	}
}

func TestSudoku_SolutionIsComplete(t *testing.T) {
	t.Parallel()
	svc := NewService()
	p, err := svc.Generate("hard")
	require.NoError(t, err)

	for r := range 9 {
		for c := range 9 {
			assert.True(t, p.Solution[r][c] >= 1 && p.Solution[r][c] <= 9, "solution[%d][%d] = %d, want 1-9", r, c, p.Solution[r][c])
		}
	}
}

func TestSudoku_SolutionIsValid(t *testing.T) {
	t.Parallel()
	svc := NewService()
	p, err := svc.Generate("medium")
	require.NoError(t, err)

	// Check each row contains 1-9.
	for r := range 9 {
		assert.True(t, hasAllDigits(p.Solution[r][:]), "row %d does not contain all digits 1-9", r)
	}

	// Check each column contains 1-9.
	for c := range 9 {
		col := make([]int, 9)
		for r := range 9 {
			col[r] = p.Solution[r][c]
		}
		assert.True(t, hasAllDigits(col), "column %d does not contain all digits 1-9", c)
	}

	// Check each 3×3 box contains 1-9.
	for br := range 3 {
		for bc := range 3 {
			box := make([]int, 0, 9)
			for r := range 3 {
				for c := range 3 {
					box = append(box, p.Solution[br*3+r][bc*3+c])
				}
			}
			assert.True(t, hasAllDigits(box), "box [%d,%d] does not contain all digits 1-9", br, bc)
		}
	}
}

func TestSudoku_PuzzleMatchesSolution(t *testing.T) {
	t.Parallel()
	svc := NewService()
	p, err := svc.Generate("easy")
	require.NoError(t, err)

	for r := range 9 {
		for c := range 9 {
			if p.Puzzle[r][c] != 0 {
				assert.Equal(t, p.Solution[r][c], p.Puzzle[r][c], "puzzle[%d][%d]=%d differs from solution[%d][%d]=%d",
					r, c, p.Puzzle[r][c], r, c, p.Solution[r][c])
			}
		}
	}
}

func TestSudokuBatch_HappyPath(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	body := `{"puzzles":["easy","medium","hard"]}`
	req := httptest.NewRequest(http.MethodPost, "/sudoku/batch", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code, "expected status 200, got %d: %s", w.Code, w.Body.String())

	var resp httpx.Response[httpx.BatchResponse[Puzzle]]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	assert.Equal(t, 3, resp.Data.Total)
	require.Len(t, resp.Data.Results, 3)

	// Results must preserve input order.
	expected := []string{"easy", "medium", "hard"}
	for i, p := range resp.Data.Results {
		assert.Equal(t, expected[i], p.Difficulty, "result[%d]: expected difficulty %q, got %q", i, expected[i], p.Difficulty)
	}
}

func TestSudokuBatch_SinglePuzzle(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	body := `{"puzzles":["hard"]}`
	req := httptest.NewRequest(http.MethodPost, "/sudoku/batch", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code, "expected status 200, got %d", w.Code)

	assert.Equal(t, "1", w.Header().Get("X-Usage-Count"))
}

func TestSudokuBatch_UsageCountMatchesBatchSize(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	body := `{"puzzles":["easy","easy","medium","hard","medium"]}`
	req := httptest.NewRequest(http.MethodPost, "/sudoku/batch", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code, "expected status 200, got %d", w.Code)

	assert.Equal(t, "5", w.Header().Get("X-Usage-Count"))
}

func TestSudokuBatch_EmptyArray(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	body := `{"puzzles":[]}`
	req := httptest.NewRequest(http.MethodPost, "/sudoku/batch", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
}

func TestSudokuBatch_MissingField(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	req := httptest.NewRequest(http.MethodPost, "/sudoku/batch", strings.NewReader(`{}`))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
}

func TestSudokuBatch_ExceedsMaxSize(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	// Build a 21-element array, one more than the allowed max of 20.
	items := make([]string, 21)
	for i := range items {
		items[i] = `"easy"`
	}
	body := `{"puzzles":[` + strings.Join(items, ",") + `]}`
	req := httptest.NewRequest(http.MethodPost, "/sudoku/batch", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
}

func TestSudokuBatch_InvalidDifficulty(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	body := `{"puzzles":["easy","impossible"]}`
	req := httptest.NewRequest(http.MethodPost, "/sudoku/batch", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code, "expected status 422 for invalid difficulty, got %d", w.Code)
}

func TestSudokuBatch_MaxAllowed(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	// Exactly 20 items — should succeed.
	items := make([]string, 20)
	for i := range items {
		items[i] = `"medium"`
	}
	body := `{"puzzles":[` + strings.Join(items, ",") + `]}`
	req := httptest.NewRequest(http.MethodPost, "/sudoku/batch", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "expected status 200 for max batch size, got %d: %s", w.Code, w.Body.String())
	assert.Equal(t, "20", w.Header().Get("X-Usage-Count"))
}

// hasAllDigits returns true when values contains each of 1-9 exactly once.
func hasAllDigits(values []int) bool {
	seen := make(map[int]bool, 9)
	for _, v := range values {
		if v < 1 || v > 9 || seen[v] {
			return false
		}
		seen[v] = true
	}
	return len(seen) == 9
}
