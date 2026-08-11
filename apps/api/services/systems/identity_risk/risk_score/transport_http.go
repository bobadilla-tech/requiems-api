package riskscore

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
	r.Post("/risk/score", handleRiskScore(svc))
}

// handleRiskScore godoc
//
//	@Summary		Score Risk
//	@Description	Scores a user for risk without the full signal breakdown. Lower latency than /signup/protect.
//	@Tags			identity-risk
//	@Accept			json
//	@Produce		json
//	@Param			request	body		Request	true	"At least one of email, phone, or ip_address"
//	@Success		200		{object}	httpx.Response[Result]
//	@Failure		401		{object}	httpx.ErrorResponse
//	@Failure		422		{object}	httpx.ErrorResponse
//	@Router			/v1/systems/risk/score [post]
func handleRiskScore(svc *Service) http.HandlerFunc {
	return httpx.Handle(
		func(ctx context.Context, req Request) (Result, error) {
			if req.Email == "" && req.Phone == "" && req.IPAddress == "" {
				return Result{}, svcerr.Unknown("validation_failed", "at least one of email, phone, or ip_address is required")
			}
			return svc.Score(ctx, req)
		},
	)
}
