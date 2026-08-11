package paymentvalidate

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
	"requiems-api/platform/svcerr"
)

type Request struct {
	BIN   string `json:"bin"`
	IBAN  string `json:"iban"`
	SWIFT string `json:"swift"`
}

func RegisterRoutes(r chi.Router, svc *Service) {
	r.Post("/payment/validate", handlePaymentValidate(svc))
}

// handlePaymentValidate godoc
//
//	@Summary		Validate Payment
//	@Description	Validates and cross-checks BIN, IBAN, and SWIFT identifiers in one call. Returns per-field results and a consistency check.
//	@Tags			payments-intelligence
//	@Accept			json
//	@Produce		json
//	@Param			request	body		Request	true	"At least one of bin, iban, or swift"
//	@Success		200		{object}	httpx.Response[Result]
//	@Failure		401		{object}	httpx.ErrorResponse
//	@Failure		422		{object}	httpx.ErrorResponse
//	@Router			/v1/systems/payment/validate [post]
func handlePaymentValidate(svc *Service) http.HandlerFunc {
	return httpx.Handle(
		func(ctx context.Context, req Request) (Result, error) {
			if req.BIN == "" && req.IBAN == "" && req.SWIFT == "" {
				return Result{}, svcerr.Unknown("validation_failed", "at least one of bin, iban, or swift is required")
			}
			return svc.Validate(ctx, req)
		},
	)
}
