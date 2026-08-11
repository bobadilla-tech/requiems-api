package password

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
)

// Request holds the optional query parameters for the password endpoint.
type Request struct {
	Length    int  `query:"length"    validate:"min=8,max=128"`
	Uppercase bool `query:"uppercase"`
	Numbers   bool `query:"numbers"`
	Symbols   bool `query:"symbols"`
}

// BatchPasswordRequest is the input for the batch password endpoint.
type BatchPasswordRequest struct {
	Items []Query `json:"items" validate:"required,min=1,max=50,dive"`
}

func RegisterRoutes(r chi.Router, svc *Service) {
	r.Get("/password", handleGeneratePassword(svc))
	r.Post("/password/batch", handlePasswordBatch(svc))
}

// handleGeneratePassword godoc
//
//	@Summary		Generate Password
//	@Description	Generates a cryptographically secure random password.
//	@Tags			password-generator
//	@Produce		json
//	@Param			length		query		integer	false	"Password length (8–128, default 16)"
//	@Param			uppercase	query		boolean	false	"Include uppercase letters (default false)"
//	@Param			numbers		query		boolean	false	"Include numbers (default false)"
//	@Param			symbols		query		boolean	false	"Include symbols (default false)"
//	@Success		200			{object}	httpx.Response[Password]
//	@Failure		400			{object}	httpx.ErrorResponse
//	@Failure		500			{object}	httpx.ErrorResponse
//	@Router			/v1/technology/password [get]
func handleGeneratePassword(svc *Service) http.HandlerFunc {
	return httpx.HandleGet(func(ctx context.Context, req Request) (Password, error) {
		return svc.Generate(req.Length, req.Uppercase, req.Numbers, req.Symbols)
	}, Request{Length: 16})
}

// handlePasswordBatch godoc
//
//	@Summary		Generate Passwords (Batch)
//	@Description	Generates up to 50 passwords; each item can have its own options.
//	@Tags			password-generator
//	@Accept			json
//	@Produce		json
//	@Param			request	body		BatchPasswordRequest	true	"List of password options"
//	@Success		200		{object}	httpx.Response[httpx.BatchResponse[BatchResult]]
//	@Failure		400		{object}	httpx.ErrorResponse
//	@Failure		422		{object}	httpx.ErrorResponse
//	@Router			/v1/technology/password/batch [post]
func handlePasswordBatch(svc *Service) http.HandlerFunc {
	return httpx.HandleBatch(
		func(_ context.Context, req BatchPasswordRequest) (httpx.BatchResponse[BatchResult], error) {
			return httpx.BatchResponse[BatchResult]{Results: svc.GenerateBatch(req.Items)}, nil
		},
	)
}
