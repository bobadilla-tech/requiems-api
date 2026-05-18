package sentiment

import (
	"context"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
)

// RegisterRoutes mounts the sentiment analysis handler on the given router.
// Request is the input for the sentiment analysis endpoint.
type Request struct {
	Text string `json:"text" validate:"required"`
}

// BatchAnalyzeRequest is the request body for analyzing multiple texts at once.
type BatchAnalyzeRequest struct {
	Texts []string `json:"texts" validate:"required,min=1,max=50,dive,required"`
}

func RegisterRoutes(r chi.Router, svc *Service) {
	r.Post("/sentiment", httpx.Handle(
		func(_ context.Context, req Request) (Result, error) {
			return svc.Analyze(req.Text), nil
		},
	))

	r.Post("/sentiment/batch", httpx.HandleBatch(
		func(_ context.Context, req BatchAnalyzeRequest) (httpx.BatchResponse[Result], error) {
			return httpx.BatchResponse[Result]{Results: svc.AnalyzeBatch(req.Texts)}, nil
		},
	))
}
