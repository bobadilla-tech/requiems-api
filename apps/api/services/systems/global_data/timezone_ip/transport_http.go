package timezoneip

import (
	"net"
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
			ip = callerIP(r)
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

// callerIP extracts the real client IP from the request, checking
// X-Forwarded-For, X-Real-IP, and RemoteAddr in that order.
func callerIP(r *http.Request) string {
	if xff := r.Header.Get("X-Forwarded-For"); xff != "" {
		if before, _, ok := strings.Cut(xff, ","); ok {
			return strings.TrimSpace(before)
		}
		return strings.TrimSpace(xff)
	}
	if xri := r.Header.Get("X-Real-IP"); xri != "" {
		return strings.TrimSpace(xri)
	}
	addr := r.RemoteAddr
	if host, _, err := net.SplitHostPort(addr); err == nil {
		return host
	}
	return addr
}
