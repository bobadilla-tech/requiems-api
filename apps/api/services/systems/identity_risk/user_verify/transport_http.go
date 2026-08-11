package userverify

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
)

type Request struct {
	Email     string `json:"email" validate:"required"`
	IPAddress string `json:"ip_address"`
}

func RegisterRoutes(r chi.Router, svc *Service) {
	r.Post("/user/verify", handleUserVerify(svc))
}

// handleUserVerify godoc
//
//	@Summary		Verify User
//	@Description	Deep-verifies an email address using domain-level signals (WHOIS age, MX, availability). Optional IP check.
//	@Tags			identity-risk
//	@Accept			json
//	@Produce		json
//	@Param			request	body		Request	true	"Email to verify"
//	@Success		200		{object}	httpx.Response[Result]
//	@Failure		401		{object}	httpx.ErrorResponse
//	@Failure		422		{object}	httpx.ErrorResponse
//	@Router			/v1/systems/user/verify [post]
func handleUserVerify(svc *Service) http.HandlerFunc {
	return httpx.Handle(
		func(ctx context.Context, req Request) (Result, error) {
			return svc.Verify(ctx, req)
		},
	)
}
