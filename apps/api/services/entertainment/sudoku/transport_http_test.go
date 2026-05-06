package sudoku

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
)

func setupRouter() chi.Router {
	r := chi.NewRouter()
	RegisterRoutes(r, NewService())
	return r
}

func TestSudoku_DefaultDifficulty(t *testing.T) {
	r := setupRouter()

	req := httptest.NewRequest(http.MethodGet, "/sudoku", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", w.Code)
	}

	var resp httpx.Response[Puzzle]
	if err := json.NewDecoder(w.Body).Decode(&resp); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if resp.Data.Difficulty != "medium" {
		t.Errorf("expected default difficulty 'medium', got %q", resp.Data.Difficulty)
	}
}

func TestSudoku_AllDifficulties(t *testing.T) {
	r := setupRouter()

	difficulties := []string{"easy", "medium", "hard"}

	for _, d := range difficulties {
		t.Run(d, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/sudoku?difficulty="+d, http.NoBody)
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			if w.Code != http.StatusOK {
				t.Errorf("expected status 200 for difficulty %q, got %d", d, w.Code)
			}

			var resp httpx.Response[Puzzle]
			if err := json.NewDecoder(w.Body).Decode(&resp); err != nil {
				t.Fatalf("failed to decode response: %v", err)
			}

			if resp.Data.Difficulty != d {
				t.Errorf("expected difficulty %q, got %q", d, resp.Data.Difficulty)
			}
		})
	}
}

func TestSudoku_InvalidDifficulty(t *testing.T) {
	r := setupRouter()

	req := httptest.NewRequest(http.MethodGet, "/sudoku?difficulty=impossible", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected status 400, got %d", w.Code)
	}
}

func TestSudoku_PuzzleHasEmptyCells(t *testing.T) {
	svc := NewService()

	for _, d := range []string{"easy", "medium", "hard"} {
		t.Run(d, func(t *testing.T) {
			p := svc.Generate(d)

			empty := 0
			for r := range 9 {
				for c := range 9 {
					if p.Puzzle[r][c] == 0 {
						empty++
					}
				}
			}

			expected := cellsToRemove[d]
			if empty != expected {
				t.Errorf("difficulty %q: expected %d empty cells, got %d", d, expected, empty)
			}
		})
	}
}

func TestSudoku_SolutionIsComplete(t *testing.T) {
	svc := NewService()
	p := svc.Generate("hard")

	for r := range 9 {
		for c := range 9 {
			if p.Solution[r][c] < 1 || p.Solution[r][c] > 9 {
				t.Errorf("solution[%d][%d] = %d, want 1-9", r, c, p.Solution[r][c])
			}
		}
	}
}

func TestSudoku_SolutionIsValid(t *testing.T) {
	svc := NewService()
	p := svc.Generate("medium")

	// Check each row contains 1-9.
	for r := range 9 {
		if !hasAllDigits(p.Solution[r][:]) {
			t.Errorf("row %d does not contain all digits 1-9", r)
		}
	}

	// Check each column contains 1-9.
	for c := range 9 {
		col := make([]int, 9)
		for r := range 9 {
			col[r] = p.Solution[r][c]
		}
		if !hasAllDigits(col) {
			t.Errorf("column %d does not contain all digits 1-9", c)
		}
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
			if !hasAllDigits(box) {
				t.Errorf("box [%d,%d] does not contain all digits 1-9", br, bc)
			}
		}
	}
}

func TestSudoku_PuzzleMatchesSolution(t *testing.T) {
	svc := NewService()
	p := svc.Generate("easy")

	for r := range 9 {
		for c := range 9 {
			if p.Puzzle[r][c] != 0 && p.Puzzle[r][c] != p.Solution[r][c] {
				t.Errorf("puzzle[%d][%d]=%d differs from solution[%d][%d]=%d",
					r, c, p.Puzzle[r][c], r, c, p.Solution[r][c])
			}
		}
	}
}

func TestSudokuBatch_HappyPath(t *testing.T) {
	r := setupRouter()

	body := `{"puzzles":["easy","medium","hard"]}`
	req := httptest.NewRequest(http.MethodPost, "/sudoku/batch", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d: %s", w.Code, w.Body.String())
	}

	var resp httpx.Response[BatchResponse]
	if err := json.NewDecoder(w.Body).Decode(&resp); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if resp.Data.Total != 3 {
		t.Errorf("expected total 3, got %d", resp.Data.Total)
	}

	if len(resp.Data.Results) != 3 {
		t.Fatalf("expected 3 results, got %d", len(resp.Data.Results))
	}

	// Results must preserve input order.
	expected := []string{"easy", "medium", "hard"}
	for i, p := range resp.Data.Results {
		if p.Difficulty != expected[i] {
			t.Errorf("result[%d]: expected difficulty %q, got %q", i, expected[i], p.Difficulty)
		}
	}
}

func TestSudokuBatch_SinglePuzzle(t *testing.T) {
	r := setupRouter()

	body := `{"puzzles":["hard"]}`
	req := httptest.NewRequest(http.MethodPost, "/sudoku/batch", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", w.Code)
	}

	if w.Header().Get("X-Usage-Count") != "1" {
		t.Errorf("expected X-Usage-Count header to be '1', got %q", w.Header().Get("X-Usage-Count"))
	}
}

func TestSudokuBatch_UsageCountMatchesBatchSize(t *testing.T) {
	r := setupRouter()

	body := `{"puzzles":["easy","easy","medium","hard","medium"]}`
	req := httptest.NewRequest(http.MethodPost, "/sudoku/batch", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", w.Code)
	}

	if w.Header().Get("X-Usage-Count") != "5" {
		t.Errorf("expected X-Usage-Count header to be '5', got %q", w.Header().Get("X-Usage-Count"))
	}
}

func TestSudokuBatch_EmptyArray(t *testing.T) {
	r := setupRouter()

	body := `{"puzzles":[]}`
	req := httptest.NewRequest(http.MethodPost, "/sudoku/batch", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusUnprocessableEntity {
		t.Errorf("expected status 422, got %d", w.Code)
	}
}

func TestSudokuBatch_MissingField(t *testing.T) {
	r := setupRouter()

	req := httptest.NewRequest(http.MethodPost, "/sudoku/batch", strings.NewReader(`{}`))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusUnprocessableEntity {
		t.Errorf("expected status 422, got %d", w.Code)
	}
}

func TestSudokuBatch_ExceedsMaxSize(t *testing.T) {
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

	if w.Code != http.StatusUnprocessableEntity {
		t.Errorf("expected status 422, got %d", w.Code)
	}
}

func TestSudokuBatch_InvalidDifficulty(t *testing.T) {
	r := setupRouter()

	body := `{"puzzles":["easy","impossible"]}`
	req := httptest.NewRequest(http.MethodPost, "/sudoku/batch", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusUnprocessableEntity {
		t.Errorf("expected status 422 for invalid difficulty, got %d", w.Code)
	}
}

func TestSudokuBatch_MaxAllowed(t *testing.T) {
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

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200 for max batch size, got %d: %s", w.Code, w.Body.String())
	}

	if w.Header().Get("X-Usage-Count") != "20" {
		t.Errorf("expected X-Usage-Count '20', got %q", w.Header().Get("X-Usage-Count"))
	}
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
