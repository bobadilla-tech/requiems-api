package profanity

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
)

// Request is the input for the profanity check endpoint.
type Request struct {
	Text string `json:"text" validate:"required"`
}

// BatchRequest is the input for the profanity batch endpoint.
type BatchRequest struct {
	Texts []string `json:"texts" validate:"required,min=1,max=50,dive,required"`
}

// RegisterRoutes mounts the profanity check handler on the given router.
func RegisterRoutes(r chi.Router, svc *Service) {
	r.Post("/profanity", handleProfanityCheck(svc))
	r.Post("/profanity/batch", handleProfanityCheckBatch(svc))
}

// handleProfanityCheck godoc
//
//	@Summary		Check Profanity
//	@Description	Checks text for profanity; returns censored version and flagged words.
//	@Tags			profanity
//	@Accept			json
//	@Produce		json
//	@Param			request	body		Request	true	"Text to check"
//	@Success		200		{object}	httpx.Response[Result]
//	@Failure		400		{object}	httpx.ErrorResponse
//	@Failure		422		{object}	httpx.ErrorResponse
//	@Failure		500		{object}	httpx.ErrorResponse
//	@Router			/v1/validation/profanity [post]
func handleProfanityCheck(svc *Service) http.HandlerFunc {
	return httpx.Handle(
		func(ctx context.Context, req Request) (Result, error) {
			return svc.Check(ctx, req.Text), nil
		},
	)
}

// handleProfanityCheckBatch godoc
//
//	@Summary		Batch Check Profanity
//	@Description	Check up to 50 texts; results in input order.
//	@Tags			profanity
//	@Accept			json
//	@Produce		json
//	@Param			request	body		BatchRequest	true	"List of texts to check"
//	@Success		200		{object}	httpx.Response[httpx.BatchResponse[BatchResult]]
//	@Failure		400		{object}	httpx.ErrorResponse
//	@Failure		422		{object}	httpx.ErrorResponse
//	@Router			/v1/validation/profanity/batch [post]
func handleProfanityCheckBatch(svc *Service) http.HandlerFunc {
	return httpx.HandleBatch(
		func(ctx context.Context, req BatchRequest) (httpx.BatchResponse[BatchResult], error) {
			return httpx.BatchResponse[BatchResult]{Results: svc.CheckBatch(ctx, req.Texts)}, nil
		},
	)
}
