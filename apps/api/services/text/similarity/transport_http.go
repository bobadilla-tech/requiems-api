package similarity

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
)

// Request is the input for the text similarity endpoint.
type Request struct {
	Text1 string `json:"text1" validate:"required"`
	Text2 string `json:"text2" validate:"required"`
}

func RegisterRoutes(r chi.Router, svc *Service) {
	r.Post("/similarity", handleTextSimilarity(svc))
}

// handleTextSimilarity godoc
//
//	@Summary		Compare Text Similarity
//	@Description	Compares two texts and returns a cosine similarity score.
//	@Tags			text-similarity
//	@Accept			json
//	@Produce		json
//	@Param			request	body		Request	true	"Two texts to compare"
//	@Success		200		{object}	httpx.Response[Result]
//	@Failure		400		{object}	httpx.ErrorResponse
//	@Failure		422		{object}	httpx.ErrorResponse
//	@Failure		500		{object}	httpx.ErrorResponse
//	@Router			/v1/text/similarity [post]
func handleTextSimilarity(svc *Service) http.HandlerFunc {
	return httpx.Handle(
		func(_ context.Context, req Request) (Result, error) {
			return svc.Cosine(req.Text1, req.Text2), nil
		},
	)
}
