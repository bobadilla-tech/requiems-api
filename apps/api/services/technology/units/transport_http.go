package units

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
)

// BatchItem represents a single unit conversion operation.
type BatchItem struct {
	From  string   `json:"from" validate:"required"`
	To    string   `json:"to"  validate:"required"`
	Value *float64 `json:"value" validate:"required"`
}

// BatchRequest represents a request containing multiple conversion operations.
type BatchRequest struct {
	Operations []BatchItem `json:"operations" validate:"required,min=1,max=50,dive"`
}

func RegisterRoutes(r chi.Router, svc *Service) {
	r.Get("/convert/units", func(w http.ResponseWriter, r *http.Request) {
		httpx.JSON(w, http.StatusOK, svc.Units())
	})

	r.Get("/convert", func(w http.ResponseWriter, r *http.Request) {
		from := r.URL.Query().Get("from")
		to := r.URL.Query().Get("to")
		valueStr := r.URL.Query().Get("value")

		if from == "" || to == "" || valueStr == "" {
			httpx.Error(w, http.StatusBadRequest, "bad_request", "from, to, and value parameters are required")
			return
		}

		value, err := strconv.ParseFloat(valueStr, 64)
		if err != nil {
			httpx.Error(w, http.StatusBadRequest, "bad_request", "value must be a valid number")
			return
		}

		result, err := svc.Convert(from, to, value)
		if err != nil {
			httpx.Error(w, http.StatusBadRequest, "bad_request", err.Error())
			return
		}

		httpx.JSON(w, http.StatusOK, result)
	})

	r.Post("/convert/batch", httpx.HandleBatch(
		func(ctx context.Context, req BatchRequest) (httpx.BatchResponse[BatchResponse], error) {
			return httpx.BatchResponse[BatchResponse]{Results: svc.ConvertBatch(ctx, req.Operations)}, nil
		}),
	)
}
