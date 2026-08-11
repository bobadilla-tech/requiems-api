package holidays

import (
	"context"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
)

// Request holds the query parameters for the holidays endpoint.
type Request struct {
	Country string `query:"country" validate:"required,iso3166_1_alpha2"`
	Year    int    `query:"year" validate:"required,min=1"`
}

// BatchRequest is the body for POST /holidays/batch.
type BatchRequest struct {
	Queries []BatchQuery `json:"queries" validate:"required,min=1,max=50,dive"`
}

func RegisterRoutes(r chi.Router, svc *Service) {
	r.Get("/holidays", handleGetHolidays(svc))
	r.Post("/holidays/batch", handleHolidaysBatch(svc))
}

// handleGetHolidays godoc
//
//	@Summary		Get Holidays
//	@Description	Returns a list of public holidays for the specified country and year.
//	@Tags			holidays
//	@Produce		json
//	@Param			country	query		string	true	"ISO 3166-1 alpha-2 country code"
//	@Param			year	query		integer	true	"Year to fetch holidays for"
//	@Success		200		{object}	httpx.Response[HolidayList]
//	@Failure		400		{object}	httpx.ErrorResponse
//	@Failure		404		{object}	httpx.ErrorResponse
//	@Router			/v1/places/holidays [get]
func handleGetHolidays(svc *Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if country := r.URL.Query().Get("country"); country != "" {
			q := r.URL.Query()
			q.Set("country", strings.ToUpper(country))
			r.URL.RawQuery = q.Encode()
		}

		var req Request
		if err := httpx.BindQuery(r, &req); err != nil {
			httpx.Error(w, http.StatusBadRequest, "bad_request", err.Error())
			return
		}

		resp, err := svc.GetHolidays(req.Country, req.Year)
		if err != nil {
			httpx.Error(w, http.StatusNotFound, "not_found", err.Error())
			return
		}

		httpx.JSON(w, http.StatusOK, resp)
	}
}

// handleHolidaysBatch godoc
//
//	@Summary		Batch Get Holidays
//	@Description	Returns holidays for up to 50 (country, year) pairs; missing combos return `found: false`.
//	@Tags			holidays
//	@Accept			json
//	@Produce		json
//	@Param			request	body		BatchRequest	true	"List of country/year queries"
//	@Success		200		{object}	httpx.Response[httpx.BatchResponse[BatchItem]]
//	@Failure		400		{object}	httpx.ErrorResponse
//	@Failure		422		{object}	httpx.ErrorResponse
//	@Router			/v1/places/holidays/batch [post]
func handleHolidaysBatch(svc *Service) http.HandlerFunc {
	return httpx.HandleBatch(
		func(_ context.Context, req BatchRequest) (httpx.BatchResponse[BatchItem], error) {
			return httpx.BatchResponse[BatchItem]{Results: svc.GetHolidaysBatch(req.Queries)}, nil
		},
	)
}
