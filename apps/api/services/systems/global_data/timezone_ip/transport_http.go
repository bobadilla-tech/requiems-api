package timezoneip

import (
	"net"
	"net/http"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
)

// RegisterRoutes mounts GET /timezone/from-ip/{ip} on the given router.
func RegisterRoutes(r chi.Router, svc *Service) {
	r.Get("/timezone/from-ip/{ip}", handleTimezoneFromIP(svc))
}

// handleTimezoneFromIP godoc
//
//	@Summary		Timezone from IP
//	@Description	Resolves timezone, UTC offset, and DST status for an IPv4/IPv6 address.
//	@Tags			global-data
//	@Produce		json
//	@Param			ip	path		string	true	"IP address (IPv4 or IPv6); use `me` for the caller IP"
//	@Success		200	{object}	httpx.Response[Result]
//	@Failure		401	{object}	httpx.ErrorResponse
//	@Failure		422	{object}	httpx.ErrorResponse
//	@Router			/v1/systems/timezone/from-ip/{ip} [get]
func handleTimezoneFromIP(svc *Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
	}
}
