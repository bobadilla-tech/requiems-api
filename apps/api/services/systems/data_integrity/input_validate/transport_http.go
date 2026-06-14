package inputvalidate

import (
	"context"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
)

// RegisterRoutes mounts the input validate handler on the given router.
func RegisterRoutes(r chi.Router, svc *Service) {
	r.Post("/input/validate", httpx.Handle(
		func(ctx context.Context, req Request) (Response, error) {
			return svc.Validate(ctx, req.Email, req.Phone, req.Text), nil
		},
	))
}
