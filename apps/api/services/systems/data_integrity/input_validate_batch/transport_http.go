package inputvalidatebatch

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
)

// RegisterRoutes mounts the input validate batch handler on the given router.
func RegisterRoutes(r chi.Router, svc *Service) {
	r.Post("/input/validate/batch", handleInputValidateBatch(svc))
}

// handleInputValidateBatch godoc
//
//	@Summary		Batch Input Validate
//	@Description	Validates and scores a batch of up to 50 contact records; per-item email, phone, and text validation plus batch aggregates.
//	@Tags			data-integrity
//	@Accept			json
//	@Produce		json
//	@Param			request	body		Request	true	"Batch of contact records to validate"
//	@Success		200		{object}	httpx.Response[BatchResponse]
//	@Failure		401		{object}	httpx.ErrorResponse
//	@Failure		422		{object}	httpx.ErrorResponse
//	@Router			/v1/systems/input/validate/batch [post]
func handleInputValidateBatch(svc *Service) http.HandlerFunc {
	return httpx.Handle(
		func(ctx context.Context, req Request) (BatchResponse, error) {
			return svc.ValidateBatch(ctx, req.Items), nil
		},
	)
}
