package sudoku

import (
	"fmt"
	"math/rand/v2"
)

// Grid is a 9×9 Sudoku board; 0 represents an empty cell.
type Grid [9][9]int

// Puzzle is the response returned by the sudoku endpoint.
type Puzzle struct {
	Difficulty string `json:"difficulty"`
	Puzzle     Grid   `json:"puzzle"`
	Solution   Grid   `json:"solution"`
}

// cellsToRemove maps a difficulty to the number of cells removed from the
// solved board when constructing the puzzle. The remaining given cells are:
//   - easy:   45 givens  (36 removed)
//   - medium: 35 givens  (46 removed)
//   - hard:   29 givens  (52 removed)
var cellsToRemove = map[string]int{
	"easy":   36,
	"medium": 46,
	"hard":   52,
}

// base is a pre-verified valid Sudoku solution used as the starting template.
// All row/column/box constraints are satisfied.
var base = Grid{
	{1, 2, 3, 4, 5, 6, 7, 8, 9},
	{4, 5, 6, 7, 8, 9, 1, 2, 3},
	{7, 8, 9, 1, 2, 3, 4, 5, 6},
	{2, 3, 4, 5, 6, 7, 8, 9, 1},
	{5, 6, 7, 8, 9, 1, 2, 3, 4},
	{8, 9, 1, 2, 3, 4, 5, 6, 7},
	{3, 4, 5, 6, 7, 8, 9, 1, 2},
	{6, 7, 8, 9, 1, 2, 3, 4, 5},
	{9, 1, 2, 3, 4, 5, 6, 7, 8},
}

// Service generates Sudoku puzzles.
type Service struct{}

// NewService returns a new Service.
func NewService() *Service {
	return &Service{}
}

// Generate returns a new Sudoku puzzle at the requested difficulty level.
// Returns an error if difficulty is not one of: easy, medium, hard.
func (s *Service) Generate(difficulty string) (Puzzle, error) {
	if _, ok := cellsToRemove[difficulty]; !ok {
		return Puzzle{}, fmt.Errorf("invalid difficulty %q: must be one of easy, medium, hard", difficulty)
	}

	rng := rand.New(rand.NewPCG(rand.Uint64(), rand.Uint64())) //nolint:gosec // Good enough for puzzle generation.
	solution := shuffle(base, rng)
	puzzle := removeCells(solution, cellsToRemove[difficulty], rng)

	return Puzzle{
		Difficulty: difficulty,
		Puzzle:     puzzle,
		Solution:   solution,
	}, nil
}

// GenerateBatch generates multiple Sudoku puzzles from the given difficulty list.
// Returns an error if any difficulty value is invalid.
func (s *Service) GenerateBatch(difficulties []string) ([]Puzzle, error) {
	results := make([]Puzzle, len(difficulties))
	for i, d := range difficulties {
		p, err := s.Generate(d)
		if err != nil {
			return nil, err
		}
		results[i] = p
	}
	return results, nil
}

// shuffle produces a new valid solution by permuting the base grid.
//
// The transformations applied are:
//  1. Remap digits — replace each digit 1-9 with a random permutation so the
//     numbers themselves look different.
//  2. Shuffle rows within each band — each band contains rows {0,1,2},
//     {3,4,5}, or {6,7,8}. Permuting rows inside a band preserves row and
//     column uniqueness while breaking pattern regularity.
//  3. Shuffle columns within each stack — analogous to row shuffling but for
//     columns in stacks {0,1,2}, {3,4,5}, {6,7,8}.
//  4. Shuffle bands — permuting the three row-bands themselves also preserves
//     validity.
//  5. Shuffle stacks — permuting the three column-stacks preserves validity.
func shuffle(g Grid, rng *rand.Rand) Grid {
	// Step 1: digit remapping.
	perm := rng.Perm(9)
	for r := range 9 {
		for c := range 9 {
			g[r][c] = perm[g[r][c]-1] + 1
		}
	}

	// Step 2: shuffle rows within each band.
	for band := range 3 {
		rows := rng.Perm(3)
		base := band * 3
		var tmp [3][9]int
		for i, row := range rows {
			tmp[i] = g[base+row]
		}
		for i := range 3 {
			g[base+i] = tmp[i]
		}
	}

	// Step 3: shuffle columns within each stack.
	for stack := range 3 {
		cols := rng.Perm(3)
		base := stack * 3
		for r := range 9 {
			var tmp [3]int
			for i, col := range cols {
				tmp[i] = g[r][base+col]
			}
			for i := range 3 {
				g[r][base+i] = tmp[i]
			}
		}
	}

	// Step 4: shuffle bands.
	bandOrder := rng.Perm(3)
	var tmpGrid Grid
	for i, b := range bandOrder {
		for j := range 3 {
			tmpGrid[i*3+j] = g[b*3+j]
		}
	}
	g = tmpGrid

	// Step 5: shuffle stacks.
	stackOrder := rng.Perm(3)
	for r := range 9 {
		var tmpRow [9]int
		for i, s := range stackOrder {
			for j := range 3 {
				tmpRow[i*3+j] = g[r][s*3+j]
			}
		}
		g[r] = tmpRow
	}

	return g
}

// removeCells blanks out the requested number of cells from the solution to
// create the puzzle. Cells are removed in a random order.
func removeCells(solution Grid, count int, rng *rand.Rand) Grid {
	puzzle := solution

	positions := rng.Perm(81)
	for _, pos := range positions[:count] {
		puzzle[pos/9][pos%9] = 0
	}

	return puzzle
}
