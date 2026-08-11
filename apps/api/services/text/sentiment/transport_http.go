package sentiment

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
)

// Request is the input for the sentiment analysis endpoint.
type Request struct {
	Text string `json:"text" validate:"required"`
}

// BatchAnalyzeRequest is the request body for analyzing multiple texts at once.
type BatchAnalyzeRequest struct {
	Texts []string `json:"texts" validate:"required,min=1,max=50,dive,required"`
}

func RegisterRoutes(r chi.Router, svc *Service) {
	r.Post("/sentiment", handleSentimentAnalyze(svc))
	r.Post("/sentiment/batch", handleSentimentAnalyzeBatch(svc))
}

// handleSentimentAnalyze godoc
//
//	@Summary		Analyze Sentiment
//	@Description	Analyzes text; returns classification, confidence score, and full breakdown.
//	@Tags			sentiment
//	@Accept			json
//	@Produce		json
//	@Param			request	body		Request	true	"Text to analyze"
//	@Success		200		{object}	httpx.Response[Result]
//	@Failure		422		{object}	httpx.ErrorResponse
//	@Router			/v1/text/sentiment [post]
func handleSentimentAnalyze(svc *Service) http.HandlerFunc {
	return httpx.Handle(
		func(_ context.Context, req Request) (Result, error) {
			return svc.Analyze(req.Text), nil
		},
	)
}

// handleSentimentAnalyzeBatch godoc
//
//	@Summary		Analyze Sentiment (Batch)
//	@Description	Analyzes up to 50 texts; each text = 1 unit of usage.
//	@Tags			sentiment
//	@Accept			json
//	@Produce		json
//	@Param			request	body		BatchAnalyzeRequest	true	"List of texts to analyze"
//	@Success		200		{object}	httpx.Response[httpx.BatchResponse[Result]]
//	@Failure		400		{object}	httpx.ErrorResponse
//	@Failure		422		{object}	httpx.ErrorResponse
//	@Router			/v1/text/sentiment/batch [post]
func handleSentimentAnalyzeBatch(svc *Service) http.HandlerFunc {
	return httpx.HandleBatch(
		func(_ context.Context, req BatchAnalyzeRequest) (httpx.BatchResponse[Result], error) {
			return httpx.BatchResponse[Result]{Results: svc.AnalyzeBatch(req.Texts)}, nil
		},
	)
}
