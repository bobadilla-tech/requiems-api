package sudoku

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
)

// RegisterRoutes mounts sudoku handlers on the given router.
// Paths are relative to the parent mount point (e.g. /v1/entertainment).
func RegisterRoutes(r chi.Router, svc *Service) {
	r.Get("/sudoku", func(w http.ResponseWriter, r *http.Request) {
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
	})

	r.Post("/sudoku/batch", httpx.HandleBatch(
		func(_ context.Context, req BatchRequest) (httpx.BatchResponse[Puzzle], error) {
			results, err := svc.GenerateBatch(req.Puzzles)
			if err != nil {
				return httpx.BatchResponse[Puzzle]{}, err
			}
			return httpx.BatchResponse[Puzzle]{Results: results}, nil
		},
	))
}
