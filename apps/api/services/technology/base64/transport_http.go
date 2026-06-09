package base64 //nolint:revive // package name matches the service domain it implements

import (
	"context"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
)

// EncodeRequest is the request body for the Base64 encode endpoint.
type EncodeRequest struct {
	Value   string `json:"value"   validate:"required"`
	Variant string `json:"variant" validate:"omitempty,oneof=standard url"`
}

// DecodeRequest is the request body for the Base64 decode endpoint.
type DecodeRequest struct {
	Value   string `json:"value"   validate:"required"`
	Variant string `json:"variant" validate:"omitempty,oneof=standard url"`
}

// RegisterRoutes mounts Base64 encode and decode handlers on the given router.
// Paths are relative to the parent mount point.
func RegisterRoutes(r chi.Router, svc *Service) {
	r.Post("/base64/encode", httpx.Handle(
		func(_ context.Context, req EncodeRequest) (Result, error) {
			return svc.Encode(req.Value, req.Variant), nil
		},
	))

	r.Post("/base64/decode", httpx.Handle(
		func(_ context.Context, req DecodeRequest) (Result, error) {
			return svc.Decode(req.Value, req.Variant)
		},
	))
	r.Post("/base64/encode/batch", httpx.HandleBatch(
		func(_ context.Context, req BatchRequest) (httpx.BatchResponse[Result], error) {
			results := make([]Result, len(req.Values))

			for i, value := range req.Values {
				results[i] = svc.Encode(value, req.Variant)
			}

			return httpx.BatchResponse[Result]{
				Results: results,
			}, nil
		},
	))

	r.Post("/base64/decode/batch", httpx.HandleBatch(
		func(_ context.Context, req BatchRequest) (httpx.BatchResponse[Result], error) {
			results := make([]Result, len(req.Values))

			for i, value := range req.Values {
				res, err := svc.Decode(value, req.Variant)

				if err != nil {
					results[i] = Result{}
					continue
				}

				results[i] = res
			}

			return httpx.BatchResponse[Result]{
				Results: results,
			}, nil
		},
	))
}
