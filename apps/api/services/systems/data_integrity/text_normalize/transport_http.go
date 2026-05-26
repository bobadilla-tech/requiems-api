package textnormalize

import (
	"context"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
)

// RegisterRoutes mounts the text normalize handler on the given router.
func RegisterRoutes(r chi.Router, svc *Service) {
	r.Post("/text/normalize", httpx.Handle(
		func(ctx context.Context, req Request) (Result, error) {
			return svc.Normalize(req.Text, req.Operations), nil
		},
	))
}
