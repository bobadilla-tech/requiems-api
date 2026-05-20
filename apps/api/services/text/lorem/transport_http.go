package lorem

import (
	"context"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
)

type Generator interface {
	Generate(paragraphs, sentences int) Lorem
}

// Request defines the parameters for generating lorem ipsum text.
// It is used both for the GET query parameters and the POST batch JSON items.
type Request struct {
	Paragraphs int `query:"paragraphs" json:"paragraphs" validate:"omitempty,min=1,max=20"`
	Sentences  int `query:"sentences" json:"sentences" validate:"omitempty,min=1,max=20"`
}

// BatchRequest represents the incoming JSON body.
type BatchRequest struct {
	Items []Request `json:"items" validate:"required,min=1,max=50,dive"`
}

// BatchItemResponse represents the response for a single item.
type BatchItemResponse struct {
	Data  *Lorem  `json:"data,omitempty"`
	Error *string `json:"error,omitempty"`
}

func RegisterRoutes(r chi.Router, svc Generator) {
	r.Get("/lorem", httpx.HandleGet(func(ctx context.Context, req Request) (Lorem, error) {
		return svc.Generate(req.Paragraphs, req.Sentences), nil
	}, Request{Paragraphs: 1, Sentences: 5}))

	r.Post("/lorem/batch", httpx.HandleBatch(func(ctx context.Context, req BatchRequest) (httpx.BatchResponse[BatchItemResponse], error) {
		var results []BatchItemResponse

		for _, item := range req.Items {
			p := item.Paragraphs
			if p == 0 {
				p = 1
			}
			s := item.Sentences
			if s == 0 {
				s = 5
			}

			lorem := svc.Generate(p, s)
			
			results = append(results, BatchItemResponse{
				Data: &lorem,
			})
		}

		return httpx.BatchResponse[BatchItemResponse]{
			Results: results,
		}, nil
	}))
}
