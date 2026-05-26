package markdown

import (
	"context"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
)

// Request is the input for the markdown-to-HTML conversion endpoint.
type Request struct {
	Markdown string `json:"markdown" validate:"required"`
	Sanitize bool   `json:"sanitize"`
}

// BatchRequest is the input for batch markdown-to-HTML conversion.
type BatchRequest struct {
	Markdowns []string `json:"markdowns" validate:"required,min=1,max=50,dive,required"`
	Sanitize  bool     `json:"sanitize"`
}

func RegisterRoutes(r chi.Router, svc *Service) {
	r.Post("/markdown", httpx.Handle(
		func(_ context.Context, req Request) (Response, error) {
			return svc.Convert(req.Markdown, req.Sanitize)
		},
	))

	r.Post("/markdown/batch", httpx.HandleBatch(
		func(_ context.Context, req BatchRequest) (httpx.BatchResponse[Response], error) {
			items, err := svc.ConvertBatch(req.Markdowns, req.Sanitize)
			return httpx.BatchResponse[Response]{Results: items}, err
		},
	))
}
