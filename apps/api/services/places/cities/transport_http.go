package cities

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
)

// BatchCitiesRequest is the input for the cities batch endpoint.
type BatchCitiesRequest struct {
	Names []string `json:"names" validate:"required,min=1,max=50,dive,required"`
}

func RegisterRoutes(r chi.Router, svc *Service) {
	r.Get("/cities/{city}", func(w http.ResponseWriter, r *http.Request) {
		name := chi.URLParam(r, "city")

		if name == "" {
			httpx.Error(w, http.StatusBadRequest, "bad_request", "city name is required")
			return
		}

		city, ok := svc.Find(name)

		if !ok {
			httpx.Error(w, http.StatusNotFound, "not_found", "city not found")
			return
		}

		httpx.JSON(w, http.StatusOK, city)
	})

	r.Post("/cities/batch", httpx.HandleBatch(
		func(_ context.Context, req BatchCitiesRequest) (httpx.BatchResponse[BatchCityItem], error) {
			return httpx.BatchResponse[BatchCityItem]{Results: svc.FindBatch(req.Names)}, nil
		},
	))
}
