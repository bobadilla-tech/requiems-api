package similarity

import (
	"context"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
)

// RegisterRoutes mounts the text similarity handler on the given router.
// Request is the input for the text similarity endpoint.
type Request struct {
	Text1 string `json:"text1" validate:"required"`
	Text2 string `json:"text2" validate:"required"`
}

func RegisterRoutes(r chi.Router, svc *Service) {
	r.Post("/similarity", httpx.Handle(
		func(_ context.Context, req Request) (Result, error) {
			return svc.Cosine(req.Text1, req.Text2), nil
		},
	))
}
