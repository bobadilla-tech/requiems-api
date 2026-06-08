package color //nolint:revive // package name matches the service domain it implements

import (
	"context"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
)

// Request holds the validated query parameters for GET /convert/color.
type Request struct {
	From  string `query:"from"  validate:"required,oneof=hex rgb hsl cmyk"`
	To    string `query:"to"    validate:"required,oneof=hex rgb hsl cmyk"`
	Value string `query:"value" validate:"required"`
}

// ColorConvertQuery is a single conversion item in a batch request.
type ColorConvertQuery struct {
	From  string `json:"from"  validate:"required,oneof=hex rgb hsl cmyk"`
	To    string `json:"to"    validate:"required,oneof=hex rgb hsl cmyk"`
	Value string `json:"value" validate:"required"`
}

// BatchColorRequest is the input for the color batch endpoint.
type BatchColorRequest struct {
	Items []ColorConvertQuery `json:"items" validate:"required,min=1,max=50,dive"`
}

// RegisterRoutes mounts the color conversion handler on the given router.
// Path is relative to the parent mount point.
func RegisterRoutes(r chi.Router, svc *Service) {
	r.Get("/color", httpx.HandleGet(func(ctx context.Context, req Request) (Response, error) {
		return svc.Convert(req.From, req.To, req.Value)
	}))

	r.Post("/color/batch", httpx.HandleBatch(
		func(_ context.Context, req BatchColorRequest) (httpx.BatchResponse[BatchColorItem], error) {
			return httpx.BatchResponse[BatchColorItem]{Results: svc.ConvertBatch(req.Items)}, nil
		},
	))
}
