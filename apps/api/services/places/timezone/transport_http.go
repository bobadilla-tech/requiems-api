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
	r.Get("/time/*", httpx.Guard(svc, handleWorldTime(svc)))
	r.Get("/timezone", httpx.Guard(svc, handleTimezoneLookup(svc)))
	r.Post("/timezone/batch", httpx.Guard(svc, handleTimezoneBatch(svc)))
}

// handleWorldTime godoc
//
//	@Summary		Get Current Time by Timezone
//	@Description	Returns the current time for the given IANA timezone identifier.
//	@Tags			world-time
//	@Produce		json
//	@Param			timezone	path		string	true	"IANA timezone identifier (e.g. America/New_York, Europe/London, UTC)"
//	@Success		200			{object}	httpx.Response[Info]
//	@Failure		400			{object}	httpx.ErrorResponse
//	@Failure		404			{object}	httpx.ErrorResponse
//	@Router			/v1/places/time/{timezone} [get]
func handleWorldTime(svc *Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
	}
}

// handleTimezoneLookup godoc
//
//	@Summary		Get Timezone
//	@Description	Returns timezone info for given coordinates or city name. Provide either `city` or both `lat` and `lon`.
//	@Tags			timezone
//	@Produce		json
//	@Param			lat		query		number	false	"Latitude (-90..90)"
//	@Param			lon		query		number	false	"Longitude (-180..180)"
//	@Param			city	query		string	false	"City name (e.g. Tokyo, London)"
//	@Success		200		{object}	httpx.Response[Info]
//	@Failure		400		{object}	httpx.ErrorResponse
//	@Failure		404		{object}	httpx.ErrorResponse
//	@Router			/v1/places/timezone [get]
func handleTimezoneLookup(svc *Service) http.HandlerFunc {
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

// handleTimezoneBatch godoc
//
//	@Summary		Batch Timezone Lookup
//	@Description	Looks up timezone for up to 50 locations. Priority: timezone name > city > coordinates.
//	@Tags			timezone
//	@Accept			json
//	@Produce		json
//	@Param			request	body		BatchTimezoneRequest	true	"List of location lookups"
//	@Success		200		{object}	httpx.Response[httpx.BatchResponse[BatchResult]]
//	@Failure		400		{object}	httpx.ErrorResponse
//	@Failure		422		{object}	httpx.ErrorResponse
//	@Router			/v1/places/timezone/batch [post]
func handleTimezoneBatch(svc *Service) http.HandlerFunc {
	return httpx.HandleBatch(
		func(_ context.Context, req BatchTimezoneRequest) (httpx.BatchResponse[BatchResult], error) {
			return httpx.BatchResponse[BatchResult]{Results: svc.BatchLookup(req.Items)}, nil
		},
	)
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
