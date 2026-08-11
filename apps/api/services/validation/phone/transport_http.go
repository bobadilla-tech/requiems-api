package phone

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
)

// ValidateRequest holds query parameters for the phone validation endpoint.
type ValidateRequest struct {
	Number string `query:"number" validate:"required"`
}

// BatchValidateRequest is the body for validating multiple phone numbers at once.
type BatchValidateRequest struct {
	Numbers []string `json:"numbers" validate:"required,min=1,max=50,dive,required"`
}

func RegisterRoutes(r chi.Router, svc *Service) {
	r.Get("/phone", handlePhoneValidate(svc))
	r.Post("/phone/batch", handlePhoneValidateBatch(svc))
}

// handlePhoneValidate godoc
//
//	@Summary		Validate Phone Number
//	@Description	Validates a single phone number; returns country, type, formatted, carrier, and risk flags.
//	@Tags			phone-validation
//	@Produce		json
//	@Param			number	query		string	true	"Phone number to validate (must include country calling code)"
//	@Success		200		{object}	httpx.Response[ValidateResponse]
//	@Failure		400		{object}	httpx.ErrorResponse
//	@Router			/v1/validation/phone [get]
func handlePhoneValidate(svc *Service) http.HandlerFunc {
	return httpx.HandleGet(func(ctx context.Context, req ValidateRequest) (ValidateResponse, error) {
		return svc.Validate(req.Number), nil
	})
}

// handlePhoneValidateBatch godoc
//
//	@Summary		Batch Validate Phone Numbers
//	@Description	Validates up to 50 phone numbers; results in input order.
//	@Tags			phone-validation
//	@Accept			json
//	@Produce		json
//	@Param			request	body		BatchValidateRequest	true	"List of phone numbers to validate"
//	@Success		200		{object}	httpx.Response[httpx.BatchResponse[ValidateResponse]]
//	@Failure		422		{object}	httpx.ErrorResponse
//	@Router			/v1/validation/phone/batch [post]
func handlePhoneValidateBatch(svc *Service) http.HandlerFunc {
	return httpx.HandleBatch(
		func(_ context.Context, req BatchValidateRequest) (httpx.BatchResponse[ValidateResponse], error) {
			return httpx.BatchResponse[ValidateResponse]{Results: svc.ValidateBatch(req.Numbers)}, nil
		},
	)
}
