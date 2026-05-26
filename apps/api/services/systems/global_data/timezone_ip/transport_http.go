package timezoneip

import (
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
)

// RegisterRoutes mounts GET /timezone/from-ip/{ip} on the given router.
func RegisterRoutes(r chi.Router, svc *Service) {
	r.Get("/timezone/from-ip/{ip}", func(w http.ResponseWriter, r *http.Request) {
		ip := chi.URLParam(r, "ip")
		if ip == "" || ip == "me" {
			// Fall back to caller IP.
			ip = r.Header.Get("X-Forwarded-For")
			if ip == "" {
				ip = strings.Split(r.RemoteAddr, ":")[0]
			}
		}

		res, err := svc.Resolve(r.Context(), ip)
		if err != nil {
			httpx.Error(w, http.StatusNotFound, "ip_not_found", "IP address not found in geolocation database")
			return
		}
		httpx.JSON(w, http.StatusOK, res)
	})
}
