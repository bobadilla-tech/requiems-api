package holidays

import (
	"context"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
)

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
