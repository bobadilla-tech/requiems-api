package signupprotect

import (
	"context"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
	"requiems-api/platform/svcerr"
)

type Request struct {
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	IPAddress string `json:"ip_address"`
}

func RegisterRoutes(r chi.Router, svc *Service) {
	r.Post("/signup/protect", httpx.Handle(
		func(ctx context.Context, req Request) (Result, error) {
			if req.Email == "" && req.Phone == "" && req.IPAddress == "" {
				return Result{}, svcerr.Unknown("validation_failed", "at least one of email, phone, or ip_address is required")
			}
			return svc.Protect(ctx, req)
		},
	))
}
