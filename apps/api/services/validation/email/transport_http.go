package email

import (
	"context"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
)

// RegisterRoutes registers HTTP routes for the validate package on r.
// It registers a POST "/validate" endpoint that accepts a Request containing an email
// and responds with a Validation produced by svc for the provided email.
func RegisterRoutes(r chi.Router, svc *Service) {
	r.Post("/email", httpx.Handle(func(ctx context.Context, req Request) (Validation, error) {
		return svc.ValidateEmail(ctx, req.Email), nil
	}))

	r.Post("/email/batch", httpx.HandleBatch(func(ctx context.Context, req BatchRequest) (httpx.BatchResponse[BatchItem], error) {
		return httpx.BatchResponse[BatchItem]{Results: svc.ValidateEmailBatch(ctx, req.Emails)}, nil
	}))
}
