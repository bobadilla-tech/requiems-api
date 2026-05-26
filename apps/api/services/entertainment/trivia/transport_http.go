package trivia

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
)

// Request holds the optional query parameters for the trivia endpoint.
type Request struct {
	Category   string `query:"category"   validate:"omitempty,oneof=science history geography sports music movies literature math technology nature"`
	Difficulty string `query:"difficulty" validate:"omitempty,oneof=easy medium hard"`
}

// BatchRequest is the request payload for fetching multiple sets
// of trivia questions in a single call.
type BatchRequest struct {
	Filters []Request `json:"filters" validate:"required,min=1,max=50,dive"`
}

// RegisterRoutes mounts trivia handlers on the given router.
// Paths are relative to the parent mount point (e.g. /v1/entertainment).
func RegisterRoutes(r chi.Router, svc *Service) {
	r.Get("/trivia", func(w http.ResponseWriter, r *http.Request) {
		var req Request
		if err := httpx.BindQuery(r, &req); err != nil {
			httpx.Error(w, http.StatusBadRequest, "bad_request", err.Error())
			return
		}

		q, err := svc.Random(req.Category, req.Difficulty)
		if err != nil {
			httpx.Error(w, http.StatusNotFound, "not_found", err.Error())
			return
		}

		httpx.JSON(w, http.StatusOK, q)
	})

	r.Post("/trivia/batch", httpx.HandleBatch(
		func(_ context.Context, req BatchRequest) (httpx.BatchResponse[BatchResponse], error) {
			return httpx.BatchResponse[BatchResponse]{Results: svc.RandomBatch(req.Filters)}, nil
		},
	))
}
