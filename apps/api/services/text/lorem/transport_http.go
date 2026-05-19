package lorem

import (
	"context"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
)

type Generator interface {
	Generate(paragraphs, sentences int) Lorem
}

// Request holds the optional query parameters for the lorem endpoint.
type Request struct {
	Paragraphs int `query:"paragraphs" validate:"min=1,max=20"`
	Sentences  int `query:"sentences"  validate:"min=1,max=20"`
}

func RegisterRoutes(r chi.Router, svc Generator) {
	r.Get("/lorem", httpx.HandleGet(func(ctx context.Context, req Request) (Lorem, error) {
		return svc.Generate(req.Paragraphs, req.Sentences), nil
	}, Request{Paragraphs: 1, Sentences: 5}))
}
