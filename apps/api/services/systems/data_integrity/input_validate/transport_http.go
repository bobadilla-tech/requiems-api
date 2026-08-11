package inputvalidate

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
)

// RegisterRoutes mounts the input validate handler on the given router.
func RegisterRoutes(r chi.Router, svc *Service) {
	r.Post("/input/validate", handleInputValidate(svc))
}

// handleInputValidate godoc
//
//	@Summary		Input Validate
//	@Description	Validates a single email, phone, or text — or any combination. Returns per-field quality scores, normalized values, risk flags, and overall quality score.
//	@Tags			data-integrity
//	@Accept			json
//	@Produce		json
//	@Param			request	body		Request	true	"At least one of email, phone, or text"
//	@Success		200		{object}	httpx.Response[Response]
//	@Failure		401		{object}	httpx.ErrorResponse
//	@Failure		422		{object}	httpx.ErrorResponse
//	@Router			/v1/systems/input/validate [post]
func handleInputValidate(svc *Service) http.HandlerFunc {
	return httpx.Handle(
		func(ctx context.Context, req Request) (Response, error) {
			return svc.Validate(ctx, req.Email, req.Phone, req.Text), nil
		},
	)
}
