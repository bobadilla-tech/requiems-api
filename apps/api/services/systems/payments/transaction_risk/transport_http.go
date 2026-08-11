package transactionrisk

import (
	"context"
	"net"
	"net/http"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
	"requiems-api/platform/svcerr"
)

type Request struct {
	CardBIN        string   `json:"card_bin"    validate:"required"`
	IPAddress      string   `json:"ip_address"  validate:"required"`
	BillingCountry string   `json:"billing_country"`
	AmountUSD      *float64 `json:"amount_usd"`
}

func RegisterRoutes(r chi.Router, svc *Service) {
	r.Post("/transaction/risk", handleTransactionRisk(svc))
}

// handleTransactionRisk godoc
//
//	@Summary		Score Transaction Risk
//	@Description	Scores a transaction for fraud risk by cross-checking card BIN country against IP geolocation and billing country. Detects VPN, proxy, TOR.
//	@Tags			payments-intelligence
//	@Accept			json
//	@Produce		json
//	@Param			request	body		Request	true	"Transaction details to score"
//	@Success		200		{object}	httpx.Response[Result]
//	@Failure		401		{object}	httpx.ErrorResponse
//	@Failure		422		{object}	httpx.ErrorResponse
//	@Router			/v1/systems/transaction/risk [post]
func handleTransactionRisk(svc *Service) http.HandlerFunc {
	return httpx.Handle(
		func(ctx context.Context, req Request) (Result, error) {
			if net.ParseIP(req.IPAddress) == nil {
				return Result{}, svcerr.Unknown("validation_failed", "invalid ip_address format")
			}
			return svc.Score(ctx, req)
		},
	)
}
