package chucknorris

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
)

// BatchRequest is the input for the chuck-norris batch endpoint.
type BatchRequest struct {
	Count int `json:"count" validate:"required,min=1,max=50"`
}

// RegisterRoutes mounts Chuck Norris fact handlers on the given router.
// Paths are relative to the parent mount point (/v1/entertainment).
func RegisterRoutes(r chi.Router, svc *Service) {
	r.Get("/chuck-norris", func(w http.ResponseWriter, r *http.Request) {
		httpx.JSON(w, http.StatusOK, svc.Random())
	})

	r.Post("/chuck-norris/batch", httpx.HandleBatch(
		func(_ context.Context, req BatchRequest) (httpx.BatchResponse[Fact], error) {
			return httpx.BatchResponse[Fact]{Results: svc.RandomBatch(req.Count)}, nil
		},
	))
}
