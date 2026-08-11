package email

import (
	"context"
	"net/http"

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
	r.Post("/email", handleEmailValidate(svc))
	r.Post("/email/batch", handleEmailValidateBatch(svc))
}

// handleEmailValidate godoc
//
//	@Summary		Validate Email
//	@Description	Validates a single email and returns syntax, MX, disposable, normalized form, and typo suggestion.
//	@Tags			email-validate
//	@Accept			json
//	@Produce		json
//	@Param			request	body		Request	true	"Email address to validate"
//	@Success		200		{object}	httpx.Response[Validation]
//	@Failure		400		{object}	httpx.ErrorResponse
//	@Failure		422		{object}	httpx.ErrorResponse
//	@Failure		500		{object}	httpx.ErrorResponse
//	@Router			/v1/validation/email [post]
func handleEmailValidate(svc *Service) http.HandlerFunc {
	return httpx.Handle(func(ctx context.Context, req Request) (Validation, error) {
		return svc.ValidateEmail(ctx, req.Email), nil
	})
}

// handleEmailValidateBatch godoc
//
//	@Summary		Validate Emails (Batch)
//	@Description	Validates up to 50 emails; each processed independently. Billing: 1 credit per email.
//	@Tags			email-validate
//	@Accept			json
//	@Produce		json
//	@Param			request	body		BatchRequest	true	"List of emails to validate"
//	@Success		200		{object}	httpx.Response[httpx.BatchResponse[BatchItem]]
//	@Failure		400		{object}	httpx.ErrorResponse
//	@Failure		422		{object}	httpx.ErrorResponse
//	@Failure		500		{object}	httpx.ErrorResponse
//	@Router			/v1/validation/email/batch [post]
func handleEmailValidateBatch(svc *Service) http.HandlerFunc {
	return httpx.HandleBatch(func(ctx context.Context, req BatchRequest) (httpx.BatchResponse[BatchItem], error) {
		return httpx.BatchResponse[BatchItem]{Results: svc.ValidateEmailBatch(ctx, req.Emails)}, nil
	})
}
