package detectlanguage

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
)

// Request is the input for the detect-language endpoint.
type Request struct {
	Text string `json:"text" validate:"required"`
}

// BatchDetectRequest is the input for the detect-language batch endpoint.
type BatchDetectRequest struct {
	Texts []string `json:"texts" validate:"required,min=1,max=50,dive,required"`
}

// RegisterRoutes mounts the detect-language handler on the given router.
func RegisterRoutes(r chi.Router, svc *Service) {
	r.Post("/detect-language", handleDetectLanguage(svc))
	r.Post("/detect-language/batch", handleDetectLanguageBatch(svc))
}

// handleDetectLanguage godoc
//
//	@Summary		Detect Language
//	@Description	Identifies the language of the provided text.
//	@Tags			detect-language
//	@Accept			json
//	@Produce		json
//	@Param			request	body		Request	true	"Text to detect"
//	@Success		200		{object}	httpx.Response[Result]
//	@Failure		400		{object}	httpx.ErrorResponse
//	@Failure		422		{object}	httpx.ErrorResponse
//	@Failure		500		{object}	httpx.ErrorResponse
//	@Router			/v1/text/detect-language [post]
func handleDetectLanguage(svc *Service) http.HandlerFunc {
	return httpx.Handle(
		func(_ context.Context, req Request) (Result, error) {
			return svc.Detect(req.Text), nil
		},
	)
}

// handleDetectLanguageBatch godoc
//
//	@Summary		Detect Language (Batch)
//	@Description	Detects the language of up to 50 texts; results in input order.
//	@Tags			detect-language
//	@Accept			json
//	@Produce		json
//	@Param			request	body		BatchDetectRequest	true	"List of texts to detect"
//	@Success		200		{object}	httpx.Response[httpx.BatchResponse[BatchDetectItem]]
//	@Failure		400		{object}	httpx.ErrorResponse
//	@Failure		422		{object}	httpx.ErrorResponse
//	@Failure		500		{object}	httpx.ErrorResponse
//	@Router			/v1/text/detect-language/batch [post]
func handleDetectLanguageBatch(svc *Service) http.HandlerFunc {
	return httpx.HandleBatch(
		func(_ context.Context, req BatchDetectRequest) (httpx.BatchResponse[BatchDetectItem], error) {
			return httpx.BatchResponse[BatchDetectItem]{Results: svc.DetectBatch(req.Texts)}, nil
		},
	)
}
