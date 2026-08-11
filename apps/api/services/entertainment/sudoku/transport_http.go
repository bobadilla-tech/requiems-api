package sudoku

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
)

// Request holds the query parameters for the sudoku endpoint.
type Request struct {
	Difficulty string `query:"difficulty" validate:"omitempty,oneof=easy medium hard"`
}

// BatchRequest is the body for generating multiple Sudoku puzzles in a single request.
// Each element must be one of: easy, medium, hard.
type BatchRequest struct {
	Puzzles []string `json:"puzzles" validate:"required,min=1,max=20,dive,oneof=easy medium hard"`
}

// RegisterRoutes mounts sudoku handlers on the given router.
// Paths are relative to the parent mount point (e.g. /v1/entertainment).
func RegisterRoutes(r chi.Router, svc *Service) {
	r.Get("/sudoku", handleSudokuPuzzle(svc))
	r.Post("/sudoku/batch", handleSudokuBatch(svc))
}

// handleSudokuPuzzle godoc
//
//	@Summary		Get Sudoku Puzzle
//	@Description	Returns a randomly generated Sudoku puzzle and its solution. Difficulty defaults to medium.
//	@Tags			sudoku
//	@Produce		json
//	@Param			difficulty	query		string	false	"Puzzle difficulty: easy, medium, hard (default medium)"
//	@Success		200			{object}	httpx.Response[Puzzle]
//	@Failure		400			{object}	httpx.ErrorResponse
//	@Router			/v1/entertainment/sudoku [get]
func handleSudokuPuzzle(svc *Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := Request{Difficulty: "medium"}
		if err := httpx.BindQuery(r, &req); err != nil {
			httpx.Error(w, http.StatusBadRequest, "bad_request", err.Error())
			return
		}

		puzzle, err := svc.Generate(req.Difficulty)
		if err != nil {
			httpx.Error(w, http.StatusBadRequest, "bad_request", err.Error())
			return
		}
		httpx.JSON(w, http.StatusOK, puzzle)
	}
}

// handleSudokuBatch godoc
//
//	@Summary		Batch Generate Sudoku Puzzles
//	@Description	Generate up to 20 Sudoku puzzles. Each puzzle = 1 unit of usage.
//	@Tags			sudoku
//	@Accept			json
//	@Produce		json
//	@Param			request	body		BatchRequest	true	"List of puzzle difficulties"
//	@Success		200		{object}	httpx.Response[httpx.BatchResponse[Puzzle]]
//	@Failure		400		{object}	httpx.ErrorResponse
//	@Failure		422		{object}	httpx.ErrorResponse
//	@Router			/v1/entertainment/sudoku/batch [post]
func handleSudokuBatch(svc *Service) http.HandlerFunc {
	return httpx.HandleBatch(
		func(_ context.Context, req BatchRequest) (httpx.BatchResponse[Puzzle], error) {
			results, err := svc.GenerateBatch(req.Puzzles)
			if err != nil {
				return httpx.BatchResponse[Puzzle]{}, err
			}
			return httpx.BatchResponse[Puzzle]{Results: results}, nil
		},
	)
}
