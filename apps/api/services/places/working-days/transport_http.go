package workingdays

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
)

func RegisterRoutes(r chi.Router, svc *Service) {
	r.Get("/working-days", func(w http.ResponseWriter, r *http.Request) {
		req := Request{}

		if err := httpx.BindQuery(r, &req); err != nil {
			httpx.Error(w, http.StatusBadRequest, "bad_request", err.Error())
			return
		}

		// Calculate working days
		workingDays := svc.GetWorkingDays(req.From, req.To, req.Country, req.Subdivision)

		// Build response
		response := WorkingDays{
			WorkingDays: workingDays,
			From:        req.From.Format("2006-01-02"),
			To:          req.To.Format("2006-01-02"),
			Country:     req.Country,
			Subdivision: req.Subdivision,
		}

		httpx.JSON(w, http.StatusOK, response)
	})
	r.Post("/working-days/batch", func(w http.ResponseWriter, r *http.Request) {
		var req BatchRequest

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			httpx.Error(w, http.StatusBadRequest, "invalid_request", "invalid json body")
			return
		}

		if len(req.Items) == 0 {
			httpx.Error(w, http.StatusBadRequest, "invalid_request", "items must not be empty")
			return
		}

		if len(req.Items) > 50 {
			httpx.Error(w, http.StatusBadRequest, "invalid_request", "items must not exceed 50")
			return
		}

		results := svc.GetWorkingDaysBatch(req.Items)

		httpx.JSON(w, http.StatusOK, BatchResponse{Results: results})
	})
}
