package spellcheck

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
)

// Request is the input for the spell check endpoint.
type Request struct {
	Text string `json:"text" validate:"required"`
}

func RegisterRoutes(r chi.Router, svc *Service) {
	r.Post("/spellcheck", handleSpellCheck(svc))
	r.Post("/spellcheck/batch", handleSpellCheckBatch(svc))
}

// handleSpellCheck godoc
//
//	@Summary		Check Spelling
//	@Description	Checks input text for spelling mistakes; returns corrected version plus per-word corrections.
//	@Tags			spell-check
//	@Accept			json
//	@Produce		json
//	@Param			request	body		Request	true	"Text to spell-check"
//	@Success		200		{object}	httpx.Response[Result]
//	@Failure		400		{object}	httpx.ErrorResponse
//	@Failure		422		{object}	httpx.ErrorResponse
//	@Failure		500		{object}	httpx.ErrorResponse
//	@Router			/v1/text/spellcheck [post]
func handleSpellCheck(svc *Service) http.HandlerFunc {
	return httpx.Handle(
		func(ctx context.Context, req Request) (Result, error) {
			return svc.Check(req.Text)
		},
	)
}

// handleSpellCheckBatch godoc
//
//	@Summary		Check Spelling (Batch)
//	@Description	Checks multiple texts; results in input order.
//	@Tags			spell-check
//	@Accept			json
//	@Produce		json
//	@Param			request	body		BatchCheckRequest	true	"List of texts to spell-check"
//	@Success		200		{object}	httpx.Response[BatchCheckResponse]
//	@Failure		400		{object}	httpx.ErrorResponse
//	@Failure		422		{object}	httpx.ErrorResponse
//	@Failure		500		{object}	httpx.ErrorResponse
//	@Router			/v1/text/spellcheck/batch [post]
func handleSpellCheckBatch(svc *Service) http.HandlerFunc {
	return httpx.HandleBatch(
		func(_ context.Context, req BatchCheckRequest) (BatchCheckResponse, error) {
			results := svc.CheckBatch(req.Texts)
			return BatchCheckResponse{Results: results}, nil
		},
	)
}
