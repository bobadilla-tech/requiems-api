package userverify

import (
	"context"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
)

// RegisterRoutes mounts POST /user/verify on the given router.
func RegisterRoutes(r chi.Router, svc *Service) {
	r.Post("/user/verify", httpx.Handle(
		func(ctx context.Context, req Request) (Result, error) {
			return svc.Verify(ctx, req)
		},
	))
}
