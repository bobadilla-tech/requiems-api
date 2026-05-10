package geocode

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
	"requiems-api/platform/svcerr"
)

func RegisterRoutes(r chi.Router, svc *Service) {
	r.Get("/geocode", func(w http.ResponseWriter, r *http.Request) {
		type req struct {
			Address string `query:"address" validate:"required"`
		}

		var q req
		if err := httpx.BindQuery(r, &q); err != nil {
			httpx.Error(w, http.StatusBadRequest, "bad_request", err.Error())
			return
		}

		result, err := svc.Geocode(r.Context(), q.Address)
		if err != nil {
			var se *svcerr.Error
			if errors.As(err, &se) {
				httpx.Error(w, svcerr.HTTPStatus(se), se.Code, se.Message)
				return
			}
			httpx.Error(w, http.StatusInternalServerError, "internal_error", "geocoding failed")
			return
		}

		httpx.JSON(w, http.StatusOK, result)
	})

	r.Get("/reverse-geocode", func(w http.ResponseWriter, r *http.Request) {
		type req struct {
			Lat float64 `query:"lat" validate:"required,min=-90,max=90"`
			Lon float64 `query:"lon" validate:"required,min=-180,max=180"`
		}

		var q req
		if err := httpx.BindQuery(r, &q); err != nil {
			httpx.Error(w, http.StatusBadRequest, "bad_request", err.Error())
			return
		}

		result, err := svc.ReverseGeocode(r.Context(), q.Lat, q.Lon)
		if err != nil {
			var se *svcerr.Error
			if errors.As(err, &se) {
				httpx.Error(w, svcerr.HTTPStatus(se), se.Code, se.Message)
				return
			}
			httpx.Error(w, http.StatusInternalServerError, "internal_error", "reverse geocoding failed")
			return
		}

		httpx.JSON(w, http.StatusOK, result)
	})
}
