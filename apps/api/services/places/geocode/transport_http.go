package geocode

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
)

// Request holds the query parameters for the geocode endpoint.
type Request struct {
	Address string `query:"address" validate:"required"`
}

// ReverseGeocodeRequest holds the query parameters for the reverse geocode endpoint.
type ReverseGeocodeRequest struct {
	Lat float64 `query:"lat" validate:"required,min=-90,max=90"`
	Lon float64 `query:"lon" validate:"required,min=-180,max=180"`
}

// BatchRequest is the input for the forward geocode batch endpoint.
type BatchRequest struct {
	Addresses []string `json:"addresses" validate:"required,min=1,max=20,dive,required"`
}

// ReverseBatchRequest is the input for the reverse geocode batch endpoint.
type ReverseBatchRequest struct {
	Items []ReverseQuery `json:"items" validate:"required,min=1,max=20,dive"`
}

func RegisterRoutes(r chi.Router, svc *Service) {
	r.Get("/geocode", handleGeocode(svc))
	r.Get("/reverse-geocode", handleReverseGeocode(svc))
	r.Post("/geocode/batch", handleGeocodeBatch(svc))
	r.Post("/reverse-geocode/batch", handleReverseGeocodeBatch(svc))
}

// handleGeocode godoc
//
//	@Summary		Geocode Address
//	@Description	Converts a free-text address into latitude/longitude.
//	@Tags			geocode
//	@Produce		json
//	@Param			address	query		string	true	"Free-text address"
//	@Success		200		{object}	httpx.Response[GeocodeResponse]
//	@Failure		400		{object}	httpx.ErrorResponse
//	@Failure		404		{object}	httpx.ErrorResponse
//	@Failure		503		{object}	httpx.ErrorResponse
//	@Router			/v1/places/geocode [get]
func handleGeocode(svc *Service) http.HandlerFunc {
	return httpx.HandleGet(func(ctx context.Context, req Request) (GeocodeResponse, error) {
		return svc.Geocode(ctx, req.Address)
	})
}

// handleReverseGeocode godoc
//
//	@Summary		Reverse Geocode
//	@Description	Converts geographic coordinates into a human-readable address.
//	@Tags			geocode
//	@Produce		json
//	@Param			lat	query		number	true	"Latitude (-90..90)"
//	@Param			lon	query		number	true	"Longitude (-180..180)"
//	@Success		200	{object}	httpx.Response[ReverseGeocodeResponse]
//	@Failure		400	{object}	httpx.ErrorResponse
//	@Failure		404	{object}	httpx.ErrorResponse
//	@Failure		503	{object}	httpx.ErrorResponse
//	@Router			/v1/places/reverse-geocode [get]
func handleReverseGeocode(svc *Service) http.HandlerFunc {
	return httpx.HandleGet(func(ctx context.Context, req ReverseGeocodeRequest) (ReverseGeocodeResponse, error) {
		return svc.ReverseGeocode(ctx, req.Lat, req.Lon)
	})
}

// handleGeocodeBatch godoc
//
//	@Summary		Batch Geocode Addresses
//	@Description	Converts up to 20 addresses to coordinates. Processed concurrently; per-item errors inline.
//	@Tags			geocode
//	@Accept			json
//	@Produce		json
//	@Param			request	body		BatchRequest	true	"List of addresses to geocode"
//	@Success		200		{object}	httpx.Response[httpx.BatchResponse[BatchGeocodeItem]]
//	@Failure		400		{object}	httpx.ErrorResponse
//	@Failure		422		{object}	httpx.ErrorResponse
//	@Router			/v1/places/geocode/batch [post]
func handleGeocodeBatch(svc *Service) http.HandlerFunc {
	return httpx.HandleBatch(
		func(ctx context.Context, req BatchRequest) (httpx.BatchResponse[BatchGeocodeItem], error) {
			return httpx.BatchResponse[BatchGeocodeItem]{Results: svc.GeocodeBatch(ctx, req.Addresses)}, nil
		},
	)
}

// handleReverseGeocodeBatch godoc
//
//	@Summary		Batch Reverse Geocode
//	@Description	Converts up to 20 coordinate pairs to addresses.
//	@Tags			geocode
//	@Accept			json
//	@Produce		json
//	@Param			request	body		ReverseBatchRequest	true	"List of coordinate pairs"
//	@Success		200		{object}	httpx.Response[httpx.BatchResponse[BatchReverseGeocodeItem]]
//	@Failure		400		{object}	httpx.ErrorResponse
//	@Failure		422		{object}	httpx.ErrorResponse
//	@Router			/v1/places/reverse-geocode/batch [post]
func handleReverseGeocodeBatch(svc *Service) http.HandlerFunc {
	return httpx.HandleBatch(
		func(ctx context.Context, req ReverseBatchRequest) (httpx.BatchResponse[BatchReverseGeocodeItem], error) {
			return httpx.BatchResponse[BatchReverseGeocodeItem]{Results: svc.ReverseGeocodeBatch(ctx, req.Items)}, nil
		},
	)
}
