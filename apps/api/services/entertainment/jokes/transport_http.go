package jokes

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
)

// BatchRequest is the input for the dad jokes batch endpoint.
type BatchRequest struct {
	Count int `json:"count" validate:"required,min=1,max=50"`
}

// RegisterRoutes mounts jokes handlers on the given router.
// Paths are relative to the parent mount point (/v1/entertainment).
func RegisterRoutes(r chi.Router, svc *Service) {
	r.Get("/jokes/dad", func(w http.ResponseWriter, r *http.Request) {
		httpx.JSON(w, http.StatusOK, svc.Random())
	})

	r.Post("/jokes/dad/batch", httpx.HandleBatch(
		func(_ context.Context, req BatchRequest) (httpx.BatchResponse[DadJoke], error) {
			return httpx.BatchResponse[DadJoke]{Results: svc.RandomBatch(req.Count)}, nil
		},
	))
}
