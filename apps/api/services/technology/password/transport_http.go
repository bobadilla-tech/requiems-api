package password

import (
	"context"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
)

// Request holds the optional query parameters for the password endpoint.
type Request struct {
	Length    int  `query:"length"    validate:"min=8,max=128"`
	Uppercase bool `query:"uppercase"`
	Numbers   bool `query:"numbers"`
	Symbols   bool `query:"symbols"`
}

// BatchPasswordRequest is the input for the batch password endpoint.
type BatchPasswordRequest struct {
	Items []Query `json:"items" validate:"required,min=1,max=50,dive"`
}

func RegisterRoutes(r chi.Router, svc *Service) {
	r.Get("/password", httpx.HandleGet(func(ctx context.Context, req Request) (Password, error) {
		return svc.Generate(req.Length, req.Uppercase, req.Numbers, req.Symbols)
	}, Request{Length: 16}))

	r.Post("/password/batch", httpx.HandleBatch(
		func(_ context.Context, req BatchPasswordRequest) (httpx.BatchResponse[BatchResult], error) {
			return httpx.BatchResponse[BatchResult]{Results: svc.GenerateBatch(req.Items)}, nil
		},
	))
}
