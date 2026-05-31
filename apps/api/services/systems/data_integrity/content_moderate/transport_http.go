package contentmoderate

import (
	"context"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
)

// RegisterRoutes mounts the content moderate handler on the given router.
func RegisterRoutes(r chi.Router, svc *Service) {
	r.Post("/content/moderate", httpx.Handle(
		func(ctx context.Context, req Request) (Response, error) {
			return svc.Moderate(ctx, req.Text, req.Language), nil
		},
	))
}
