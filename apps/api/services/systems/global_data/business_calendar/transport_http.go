package businesscalendar

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
	"requiems-api/platform/svcerr"
)

type Request struct {
	Country string
	Year    int
	Month   int // 0 = full year scope
}

func RegisterRoutes(r chi.Router, svc *Service) {
	r.Get("/business-calendar/{country}", func(w http.ResponseWriter, r *http.Request) {
		country := strings.ToUpper(chi.URLParam(r, "country"))
		if len(country) != 2 {
			httpx.Error(w, http.StatusUnprocessableEntity, "validation_failed", "country must be a 2-letter ISO 3166-1 alpha-2 code")
			return
		}

		now := time.Now()
		req := Request{Country: country, Year: now.Year()}

		if y := r.URL.Query().Get("year"); y != "" {
			n, err := strconv.Atoi(y)
			if err != nil || n < 1900 || n > 2100 {
				httpx.Error(w, http.StatusUnprocessableEntity, "validation_failed", "year must be between 1900 and 2100")
				return
			}
			req.Year = n
		}
		if m := r.URL.Query().Get("month"); m != "" {
			n, err := strconv.Atoi(m)
			if err != nil || n < 1 || n > 12 {
				httpx.Error(w, http.StatusUnprocessableEntity, "validation_failed", "month must be between 1 and 12")
				return
			}
			req.Month = n
		}

		res, err := svc.Get(r.Context(), req)
		if err != nil {
			if se, ok := err.(*svcerr.Error); ok {
				httpx.Error(w, svcerr.HTTPStatus(se), se.Code, se.Message)
				return
			}
			httpx.Error(w, http.StatusInternalServerError, "internal_error", "internal server error")
			return
		}
		httpx.JSON(w, http.StatusOK, res)
	})
}
