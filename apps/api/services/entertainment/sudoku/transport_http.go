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

		httpx.JSON(w, http.StatusOK, svc.Generate(req.Difficulty))
	})

	r.Post("/sudoku/batch", httpx.HandleBatch(
		func(_ context.Context, req BatchRequest) (BatchResponse, int, error) {
			results := svc.GenerateBatch(req.Puzzles)
			return BatchResponse{Results: results, Total: len(results)}, len(results), nil
		},
	))
}
