package transactionrisk

import (
	"context"
	"net"

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
	r.Post("/transaction/risk", httpx.Handle(
		func(ctx context.Context, req Request) (Result, error) {
			if net.ParseIP(req.IPAddress) == nil {
				return Result{}, svcerr.Unknown("validation_failed", "invalid ip_address format")
			}
			return svc.Score(ctx, req)
		},
	))
}
