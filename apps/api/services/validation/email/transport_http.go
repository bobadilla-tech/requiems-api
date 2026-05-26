package email

import (
	"context"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
)

// Request holds the JSON body for email validation.
type Request struct {
	Email string `json:"email" validate:"required"`
}

// BatchRequest is the body for validating multiple emails at once.
type BatchRequest struct {
	Emails []string `json:"emails" validate:"required,min=1,max=50,dive,required"`
}

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
