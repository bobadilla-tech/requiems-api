package randomuser

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
)

// RegisterRoutes mounts the random-user handler on the given router.
// Paths are relative to the parent mount point (e.g. /v1/technology).
func RegisterRoutes(r chi.Router, svc *Service) {
	r.Get("/random-user", func(w http.ResponseWriter, r *http.Request) {
		httpx.JSON(w, http.StatusOK, svc.Generate())
	})

	r.Post("/random-user/batch", httpx.HandleBatch(
		func(_ context.Context, req BatchGenerateRequest) (BatchGenerateResponse, int, error) {
			return svc.GenerateBatch(req.Count), req.Count, nil
		},
	))
}
