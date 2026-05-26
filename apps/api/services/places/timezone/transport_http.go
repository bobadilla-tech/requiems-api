package timezone

import (
	"context"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
)

// Request holds the query parameters for the timezone endpoint.
// Either lat+lon together, or city must be provided.
type Request struct {
	Lat  float64 `query:"lat" validate:"min=-90,max=90"`
	Lon  float64 `query:"lon" validate:"min=-180,max=180"`
	City string  `query:"city"`
}

// BatchTimezoneRequest is the input for the batch timezone endpoint.
type BatchTimezoneRequest struct {
	Items []Query `json:"items" validate:"required,min=1,max=50,dive"`
}

func RegisterRoutes(r chi.Router, svc *Service) {
	r.Get("/time/*", httpx.Guard(svc, func(w http.ResponseWriter, r *http.Request) {
		tzName := chi.URLParam(r, "*")
		if tzName == "" {
			httpx.Error(w, http.StatusBadRequest, "bad_request", "timezone is required")
			return
		}

		info, err := svc.GetCurrentTime(tzName)
		if err != nil {
			httpx.Error(w, http.StatusNotFound, "not_found", err.Error())
			return
		}

		httpx.JSON(w, http.StatusOK, *info)
	}))

	r.Get("/timezone", httpx.Guard(svc, handleGetTimezone(svc)))

	r.Post("/timezone/batch", httpx.Guard(svc, httpx.HandleBatch(
		func(_ context.Context, req BatchTimezoneRequest) (httpx.BatchResponse[BatchResult], error) {
			return httpx.BatchResponse[BatchResult]{Results: svc.BatchLookup(req.Items)}, nil
		},
	)))
}

func handleGetTimezone(svc *Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req, hasCity, hasCoords, err := parseTimezoneQuery(r)
		if err != nil {
			httpx.Error(w, http.StatusBadRequest, "bad_request", err.Error())
			return
		}

		var info *Info
		if hasCity {
			info, err = svc.GetTimezoneByCity(req.City)
		} else if hasCoords {
			info, err = svc.GetTimezoneByCoords(req.Lat, req.Lon)
		}

		if err != nil {
			httpx.Error(w, http.StatusNotFound, "not_found", err.Error())
			return
		}

		httpx.JSON(w, http.StatusOK, *info)
	}
}

func parseTimezoneQuery(r *http.Request) (req Request, hasCity, hasCoords bool, err error) {
	if err = httpx.BindQuery(r, &req); err != nil {
		return
	}

	q := r.URL.Query()
	hasCoords = q.Has("lat") && q.Has("lon")
	hasCity = q.Has("city")

	if !hasCoords && !hasCity {
		err = errors.New("provide either 'city' or both 'lat' and 'lon' query parameters")
	}

	return
}
