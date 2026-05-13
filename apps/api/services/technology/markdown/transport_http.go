package markdown

import (
	"context"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
)

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
