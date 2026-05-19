package geocode

import (
	"context"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
)

// GeocodeRequest holds the query parameters for the geocode endpoint.
type GeocodeRequest struct {
	Address string `query:"address" validate:"required"`
}

// ReverseGeocodeRequest holds the query parameters for the reverse geocode endpoint.
type ReverseGeocodeRequest struct {
	Lat float64 `query:"lat" validate:"required,min=-90,max=90"`
	Lon float64 `query:"lon" validate:"required,min=-180,max=180"`
}

func RegisterRoutes(r chi.Router, svc *Service) {
	r.Get("/geocode", httpx.HandleGet(func(ctx context.Context, req GeocodeRequest) (GeocodeResponse, error) {
		return svc.Geocode(ctx, req.Address)
	}))

	r.Get("/reverse-geocode", httpx.HandleGet(func(ctx context.Context, req ReverseGeocodeRequest) (ReverseGeocodeResponse, error) {
		return svc.ReverseGeocode(ctx, req.Lat, req.Lon)
	}))
}
