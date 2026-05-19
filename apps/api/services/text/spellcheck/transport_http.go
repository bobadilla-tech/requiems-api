package spellcheck

import (
	"context"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
)

// RegisterRoutes mounts the spell check handler on the given router.
// Request is the input for the spell check endpoint.
type Request struct {
	Text string `json:"text" validate:"required"`
}

func RegisterRoutes(r chi.Router, svc *Service) {
	r.Post("/spellcheck", httpx.Handle(
		func(ctx context.Context, req Request) (Result, error) {
			return svc.Check(req.Text)
		},
	))

	r.Post("/spellcheck/batch", httpx.HandleBatch(
		func(_ context.Context, req BatchCheckRequest) (BatchCheckResponse, error) {
			results, err := svc.CheckBatch(req.Texts)
			if err != nil {
				return BatchCheckResponse{}, err
			}
			return BatchCheckResponse{Results: results}, nil
		},
	))
}
