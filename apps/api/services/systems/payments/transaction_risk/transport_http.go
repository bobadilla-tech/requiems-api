package transactionrisk

import (
	"context"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
)

// RegisterRoutes mounts POST /transaction/risk on the given router.
func RegisterRoutes(r chi.Router, svc *Service) {
	r.Post("/transaction/risk", httpx.Handle(
		func(ctx context.Context, req Request) (Result, error) {
			return svc.Score(ctx, req)
		},
	))
}
