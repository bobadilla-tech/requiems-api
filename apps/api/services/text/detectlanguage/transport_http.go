package detectlanguage

import (
	"context"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
)

// Request is the input for the detect-language endpoint.
type Request struct {
	Text string `json:"text" validate:"required"`
}

// RegisterRoutes mounts the detect-language handler on the given router.
func RegisterRoutes(r chi.Router, svc *Service) {
	r.Post("/detect-language", httpx.Handle(
		func(_ context.Context, req Request) (Result, error) {
			return svc.Detect(req.Text), nil
		},
	))
}
