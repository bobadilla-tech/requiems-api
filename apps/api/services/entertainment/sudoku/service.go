package sudoku

import (
	"fmt"

	sudokulib "github.com/bobadilla-tech/sudoku-go"
)

// Grid is a 9×9 Sudoku board; 0 represents an empty cell.
type Grid = sudokulib.Grid

// Puzzle is the response returned by the sudoku endpoint.
type Puzzle struct {
	Difficulty string `json:"difficulty"`
	Puzzle     Grid   `json:"puzzle"`
	Solution   Grid   `json:"solution"`
}

// difficultyMap maps the API string difficulty to the library's Difficulty type.
var difficultyMap = map[string]sudokulib.Difficulty{
	"easy":   sudokulib.Easy,
	"medium": sudokulib.Medium,
	"hard":   sudokulib.Hard,
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
	d, ok := difficultyMap[difficulty]
	if !ok {
		return Puzzle{}, fmt.Errorf("invalid difficulty %q: must be one of easy, medium, hard", difficulty)
	}

	p, err := sudokulib.Generate(d)
	if err != nil {
		return Puzzle{}, err
	}

	return Puzzle{
		Difficulty: difficulty,
		Puzzle:     p.Grid,
		Solution:   p.Solution,
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
