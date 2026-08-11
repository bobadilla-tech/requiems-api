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
	r.Get("/cities/{city}", handleCityLookup(svc))
	r.Post("/cities/batch", handleCitiesBatch(svc))
}

// handleCityLookup godoc
//
//	@Summary		Get City Info
//	@Description	Returns metadata for a city by name. Case-insensitive.
//	@Tags			cities
//	@Produce		json
//	@Param			city	path		string	true	"City name"
//	@Success		200		{object}	httpx.Response[City]
//	@Failure		400		{object}	httpx.ErrorResponse
//	@Failure		404		{object}	httpx.ErrorResponse
//	@Failure		500		{object}	httpx.ErrorResponse
//	@Router			/v1/places/cities/{city} [get]
func handleCityLookup(svc *Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
	}
}

// handleCitiesBatch godoc
//
//	@Summary		Get Cities (Batch)
//	@Description	Returns metadata for up to 50 cities by name.
//	@Tags			cities
//	@Accept			json
//	@Produce		json
//	@Param			request	body		BatchCitiesRequest	true	"List of city names"
//	@Success		200		{object}	httpx.Response[httpx.BatchResponse[BatchCityItem]]
//	@Failure		400		{object}	httpx.ErrorResponse
//	@Failure		422		{object}	httpx.ErrorResponse
//	@Router			/v1/places/cities/batch [post]
func handleCitiesBatch(svc *Service) http.HandlerFunc {
	return httpx.HandleBatch(
		func(_ context.Context, req BatchCitiesRequest) (httpx.BatchResponse[BatchCityItem], error) {
			return httpx.BatchResponse[BatchCityItem]{Results: svc.FindBatch(req.Names)}, nil
		},
	)
}
