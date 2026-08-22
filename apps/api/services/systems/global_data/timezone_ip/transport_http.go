package timezoneip

import (
	"net"
	"net/http"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
)

// RegisterRoutes mounts GET /timezone/from-ip/{ip} on the given router.
func RegisterRoutes(r chi.Router, svc *Service) {
	r.Get("/timezone/from-ip/{ip}", func(w http.ResponseWriter, r *http.Request) {
		ip := chi.URLParam(r, "ip")
		if ip == "" || ip == "me" {
			ip = httpx.ClientIP(r)
		}

		if net.ParseIP(ip) == nil {
			httpx.Error(w, http.StatusBadRequest, "bad_request", "invalid IP address")
			return
		}

		res, err := svc.Resolve(r.Context(), ip)
		if err != nil {
			httpx.Error(w, http.StatusNotFound, "ip_not_found", "IP address not found in geolocation database")
			return
		}
		httpx.JSON(w, http.StatusOK, res)
	})
}
