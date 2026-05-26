package userverify

import (
	"context"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
)

type Request struct {
	Email     string `json:"email" validate:"required"`
	IPAddress string `json:"ip_address"`
}

func RegisterRoutes(r chi.Router, svc *Service) {
	r.Post("/user/verify", httpx.Handle(
		func(ctx context.Context, req Request) (Result, error) {
			return svc.Verify(ctx, req)
		},
	))
}
