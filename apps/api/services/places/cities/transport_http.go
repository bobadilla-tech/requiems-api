package cities

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
)

func RegisterRoutes(r chi.Router, svc *Service) {
	r.Get("/cities/{city}", func(w http.ResponseWriter, r *http.Request) {
		name := chi.URLParam(r, "city")

		if name == "" {
			httpx.Error(w, http.StatusBadRequest, "bad_request", "city name is required")
			return
		}

		city, ok := svc.Find(name)

		if !ok {
			httpx.Error(w, http.StatusNotFound, "not_found", "city not found")
			return
		}

		httpx.JSON(w, http.StatusOK, city)
	})
}
