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
		func(_ context.Context, req BatchCheckRequest) (BatchCheckResponse, int, error) {
			res, err := svc.CheckBatch(req.Texts)
			return res, len(req.Texts), err
		},
	))
}
