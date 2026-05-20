package workingdays

import (
	"net/http"
	"context"
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
	
	r.Post("/working-days/batch", httpx.HandleBatch(
	func(ctx context.Context, req BatchRequest) (httpx.BatchResponse[WorkingDays], error) {
		items := make([]Request, 0, len(req.Items))

		for _, item := range req.Items {
			items = append(items, Request{
				From:        item.From.Time,
				To:          item.To.Time,
				Country:     item.Country,
				Subdivision: item.Subdivision,
			})
		}

		results := svc.GetWorkingDaysBatch(items)

		return httpx.BatchResponse[WorkingDays]{
			Results: results,
		}, nil
	},
))
}
