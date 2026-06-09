package convformat

import (
	"context"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
)

// Request is the input for the format conversion endpoint.
type Request struct {
	From    string `json:"from"    validate:"required,oneof=json yaml csv xml toml"`
	To      string `json:"to"      validate:"required,oneof=json yaml csv xml toml"`
	Content string `json:"content" validate:"required"`
}

// BatchFormatRequest is the input for the format conversion batch endpoint.
// Max 20 items — each item content may be up to 512 KB; total body stays ≤ 1 MiB.
type BatchFormatRequest struct {
	Items []Request `json:"items" validate:"required,min=1,max=20,dive"`
}

func RegisterRoutes(r chi.Router, svc *Service) {
	r.Post("/format", httpx.Handle(
		func(_ context.Context, req Request) (Response, error) {
			return svc.Convert(req)
		},
	))
	r.Post("/format/batch", httpx.HandleBatch(
		func(_ context.Context, req BatchFormatRequest) (httpx.BatchResponse[BatchFormatItem], error) {
			return httpx.BatchResponse[BatchFormatItem]{Results: svc.ConvertBatch(req.Items)}, nil
		},
	))
}
