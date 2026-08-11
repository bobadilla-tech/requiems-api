package textnormalize

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
)

// RegisterRoutes mounts the text normalize handler on the given router.
func RegisterRoutes(r chi.Router, svc *Service) {
	r.Post("/text/normalize", handleTextNormalize(svc))
}

// handleTextNormalize godoc
//
//	@Summary		Normalize Text
//	@Description	Applies a composable pipeline of normalization operations to a string.
//	@Tags			data-integrity
//	@Accept			json
//	@Produce		json
//	@Param			request	body		Request	true	"Text and normalization operations"
//	@Success		200		{object}	httpx.Response[Result]
//	@Failure		401		{object}	httpx.ErrorResponse
//	@Failure		422		{object}	httpx.ErrorResponse
//	@Router			/v1/systems/text/normalize [post]
func handleTextNormalize(svc *Service) http.HandlerFunc {
	return httpx.Handle(
		func(ctx context.Context, req Request) (Result, error) {
			return svc.Normalize(req.Text, req.Operations), nil
		},
	)
}
