package contentmoderate

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
)

// RegisterRoutes mounts the content moderate handler on the given router.
func RegisterRoutes(r chi.Router, svc *Service) {
	r.Post("/content/moderate", handleContentModerate(svc))
}

// handleContentModerate godoc
//
//	@Summary		Moderate Content
//	@Description	Checks text for profanity, toxicity, sentiment, and language. Returns per-category flags and `is_safe`.
//	@Tags			data-integrity
//	@Accept			json
//	@Produce		json
//	@Param			request	body		Request	true	"Text to moderate"
//	@Success		200		{object}	httpx.Response[Response]
//	@Failure		401		{object}	httpx.ErrorResponse
//	@Failure		422		{object}	httpx.ErrorResponse
//	@Router			/v1/systems/content/moderate [post]
func handleContentModerate(svc *Service) http.HandlerFunc {
	return httpx.Handle(
		func(ctx context.Context, req Request) (Response, error) {
			return svc.Moderate(ctx, req.Text, req.Language), nil
		},
	)
}
