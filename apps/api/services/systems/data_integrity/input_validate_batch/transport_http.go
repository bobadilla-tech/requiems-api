package inputvalidatebatch

import (
	"context"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
)

// RegisterRoutes mounts the input validate batch handler on the given router.
func RegisterRoutes(r chi.Router, svc *Service) {
	r.Post("/input/validate/batch", httpx.Handle(

		func(ctx context.Context, req Request) (BatchResponse, error) {
			return svc.ValidateBatch(ctx, req.Items), nil
		},
	))
}
