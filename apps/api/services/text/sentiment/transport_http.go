package sentiment

import (
	"context"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
)

// RegisterRoutes mounts the sentiment analysis handler on the given router.
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
