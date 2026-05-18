package workingdays

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
)

// Request holds the query parameters for the working days endpoint.
type Request struct {
	From        time.Time `query:"from" validate:"required"`
	To          time.Time `query:"to" validate:"required,gtfield=From"`
	Country     string    `query:"country" validate:"omitempty,iso3166_1_alpha2"`
	Subdivision string    `query:"subdivision" validate:"omitempty,iso3166_2,required_with=Country"`
}

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
}
