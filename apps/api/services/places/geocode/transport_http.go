package geocode

import (
	"context"

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

// GeocodeBatchRequest is the input for the forward geocode batch endpoint.
type GeocodeBatchRequest struct {
	Addresses []string `json:"addresses" validate:"required,min=1,max=20,dive,required"`
}

// ReverseGeocodeBatchRequest is the input for the reverse geocode batch endpoint.
type ReverseGeocodeBatchRequest struct {
	Items []ReverseQuery `json:"items" validate:"required,min=1,max=20,dive"`
}

func RegisterRoutes(r chi.Router, svc *Service) {
	r.Get("/geocode", httpx.HandleGet(func(ctx context.Context, req Request) (GeocodeResponse, error) {
		return svc.Geocode(ctx, req.Address)
	}))

	r.Get("/reverse-geocode", httpx.HandleGet(func(ctx context.Context, req ReverseGeocodeRequest) (ReverseGeocodeResponse, error) {
		return svc.ReverseGeocode(ctx, req.Lat, req.Lon)
	}))

	r.Post("/geocode/batch", httpx.HandleBatch(
		func(ctx context.Context, req GeocodeBatchRequest) (httpx.BatchResponse[BatchGeocodeItem], error) {
			return httpx.BatchResponse[BatchGeocodeItem]{Results: svc.GeocodeBatch(ctx, req.Addresses)}, nil
		},
	))

	r.Post("/reverse-geocode/batch", httpx.HandleBatch(
		func(ctx context.Context, req ReverseGeocodeBatchRequest) (httpx.BatchResponse[BatchReverseGeocodeItem], error) {
			return httpx.BatchResponse[BatchReverseGeocodeItem]{Results: svc.ReverseGeocodeBatch(ctx, req.Items)}, nil
		},
	))
}
