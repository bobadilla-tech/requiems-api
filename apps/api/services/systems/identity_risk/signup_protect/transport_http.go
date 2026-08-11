package signupprotect

import (
	"context"
	"net/http"

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
	r.Post("/signup/protect", handleSignupProtect(svc))
}

// handleSignupProtect godoc
//
//	@Summary		Protect Signup
//	@Description	Evaluates a new user at signup; full risk decision with per-signal breakdown across email, phone, and IP.
//	@Tags			identity-risk
//	@Accept			json
//	@Produce		json
//	@Param			request	body		Request	true	"At least one of email, phone, or ip_address"
//	@Success		200		{object}	httpx.Response[Result]
//	@Failure		401		{object}	httpx.ErrorResponse
//	@Failure		422		{object}	httpx.ErrorResponse
//	@Router			/v1/systems/signup/protect [post]
func handleSignupProtect(svc *Service) http.HandlerFunc {
	return httpx.Handle(
		func(ctx context.Context, req Request) (Result, error) {
			if req.Email == "" && req.Phone == "" && req.IPAddress == "" {
				return Result{}, svcerr.Unknown("validation_failed", "at least one of email, phone, or ip_address is required")
			}
			return svc.Protect(ctx, req)
		},
	)
}
