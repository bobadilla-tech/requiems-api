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
	r.Get("/holidays", func(w http.ResponseWriter, r *http.Request) {
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
	})

	// POST /holidays/batch — return holidays for up to 50 (country, year) pairs at once.
	// Uses HandleBatch so the gateway charges one credit per query (X-Usage-Count).
	r.Post("/holidays/batch", httpx.HandleBatch(
		func(_ context.Context, req BatchRequest) (httpx.BatchResponse[BatchItem], error) {
			return httpx.BatchResponse[BatchItem]{Results: svc.GetHolidaysBatch(req.Queries)}, nil
		},
	))
}
