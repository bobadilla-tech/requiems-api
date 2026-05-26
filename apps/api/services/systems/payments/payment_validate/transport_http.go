package paymentvalidate

import (
	"context"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
	"requiems-api/platform/svcerr"
)

// RegisterRoutes mounts POST /payment/validate on the given router.
func RegisterRoutes(r chi.Router, svc *Service) {
	r.Post("/payment/validate", httpx.Handle(
		func(ctx context.Context, req Request) (Result, error) {
			if req.BIN == "" && req.IBAN == "" && req.SWIFT == "" {
				return Result{}, svcerr.Unknown("validation_failed", "at least one of bin, iban, or swift is required")
			}
			return svc.Validate(ctx, req)
		},
	))
}
