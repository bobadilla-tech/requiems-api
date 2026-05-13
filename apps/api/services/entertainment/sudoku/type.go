package sudoku

// Request holds the query parameters for the sudoku endpoint.
// Defaults should be set before calling httpx.BindQuery.
type Request struct {
	Difficulty string `query:"difficulty" validate:"omitempty,oneof=easy medium hard"`
}

// Grid is a 9×9 Sudoku board; 0 represents an empty cell.
type Grid [9][9]int

// Puzzle is the response returned by the sudoku endpoint.
type Puzzle struct {
	Difficulty string `json:"difficulty"`
	Puzzle     Grid   `json:"puzzle"`
	Solution   Grid   `json:"solution"`
}

func (Puzzle) IsData() {}

// BatchRequest is the body for generating multiple Sudoku puzzles in a single request.
// Each element must be one of: easy, medium, hard.
type BatchRequest struct {
	Puzzles []string `json:"puzzles" validate:"required,min=1,max=20,dive,oneof=easy medium hard"`
}
