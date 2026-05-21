package profanity

import (
	"context"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
)

// Request is the input for the profanity check endpoint.
type Request struct {
	Text string `json:"text" validate:"required"`
}

// BatchRequest is the input for the profanity batch endpoint.
type BatchRequest struct {
	Texts []string `json:"texts" validate:"required,min=1,max=50,dive,required"`
}

// RegisterRoutes mounts the profanity check handler on the given router.
func RegisterRoutes(r chi.Router, svc *Service) {
	r.Post("/profanity", httpx.Handle(
		func(ctx context.Context, req Request) (Result, error) {
			return svc.Check(ctx, req.Text), nil
		},
	))

	r.Post("/profanity/batch", httpx.HandleBatch(
		func(ctx context.Context, req BatchRequest) (httpx.BatchResponse[BatchResult], error) {
			return httpx.BatchResponse[BatchResult]{Results: svc.CheckBatch(ctx, req.Texts)}, nil
		},
	))
}
