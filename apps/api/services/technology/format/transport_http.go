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

func RegisterRoutes(r chi.Router, svc *Service) {
	r.Post("/format", httpx.Handle(
		func(_ context.Context, req Request) (Response, error) {
			return svc.Convert(req)
		},
	))
}
