package postal

import (
	"context"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
)

// BatchPostalRequest is the input for the batch postal code endpoint.
type BatchPostalRequest struct {
	Items []Query `json:"items" validate:"required,min=1,max=50,dive"`
}

func RegisterRoutes(r chi.Router, svc *Service) {
	r.Get("/postal/{code}", func(w http.ResponseWriter, r *http.Request) {
		code := chi.URLParam(r, "code")
		if code == "" {
			httpx.Error(w, http.StatusBadRequest, "bad_request", "postal code is required")
			return
		}

		country := strings.ToUpper(r.URL.Query().Get("country"))
		if country == "" {
			country = "US"
		}

		result, ok := svc.Lookup(code, country)
		if !ok {
			httpx.Error(w, http.StatusNotFound, "not_found", "postal code not found")
			return
		}

		httpx.JSON(w, http.StatusOK, result)
	})

	r.Post("/postal/batch", httpx.HandleBatch(
		func(_ context.Context, req BatchPostalRequest) (httpx.BatchResponse[BatchResult], error) {
			return httpx.BatchResponse[BatchResult]{Results: svc.LookupBatch(req.Items)}, nil
		},
	))
}
