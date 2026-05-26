package numbase

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
)

// ConvertRequest holds the validated query parameters for the base conversion endpoint.
type ConvertRequest struct {
	From  int    `query:"from"  validate:"required,oneof=2 8 10 16"`
	To    int    `query:"to"    validate:"required,oneof=2 8 10 16"`
	Value string `query:"value" validate:"required"`
}

// BatchConvertRequest is the input for the batch base conversion endpoint.
type BatchConvertRequest struct {
	Items []ConvertQuery `json:"items" validate:"required,min=1,max=50,dive"`
}

// RegisterRoutes mounts the base conversion handler on the given router.
// Paths are relative to the parent mount point.
func RegisterRoutes(r chi.Router, svc *Service) {
	r.Get("/base", func(w http.ResponseWriter, r *http.Request) {
		var req ConvertRequest
		if err := httpx.BindQuery(r, &req); err != nil {
			httpx.Error(w, http.StatusBadRequest, "bad_request", err.Error())
			return
		}

		result, err := svc.Convert(req.Value, req.From, req.To)
		if err != nil {
			httpx.Error(w, http.StatusBadRequest, "bad_request", err.Error())
			return
		}

		httpx.JSON(w, http.StatusOK, result)
	})

	r.Post("/base/batch", httpx.HandleBatch(
		func(_ context.Context, req BatchConvertRequest) (httpx.BatchResponse[BatchResult], error) {
			return httpx.BatchResponse[BatchResult]{Results: svc.ConvertBatch(req.Items)}, nil
		},
	))
}
