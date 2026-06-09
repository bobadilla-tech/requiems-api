package detectlanguage

import (
	"context"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
)

// Request is the input for the detect-language endpoint.
type Request struct {
	Text string `json:"text" validate:"required"`
}

// BatchDetectRequest is the input for the detect-language batch endpoint.
type BatchDetectRequest struct {
	Texts []string `json:"texts" validate:"required,min=1,max=50,dive,required"`
}

// RegisterRoutes mounts the detect-language handler on the given router.
func RegisterRoutes(r chi.Router, svc *Service) {
	r.Post("/detect-language", httpx.Handle(
		func(_ context.Context, req Request) (Result, error) {
			return svc.Detect(req.Text), nil
		},
	))

	r.Post("/detect-language/batch", httpx.HandleBatch(
		func(_ context.Context, req BatchDetectRequest) (httpx.BatchResponse[BatchDetectItem], error) {
			return httpx.BatchResponse[BatchDetectItem]{Results: svc.DetectBatch(req.Texts)}, nil
		},
	))
}
