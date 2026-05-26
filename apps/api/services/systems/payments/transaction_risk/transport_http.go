package transactionrisk

import (
	"context"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
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
			return svc.Score(ctx, req)
		},
	))
}
