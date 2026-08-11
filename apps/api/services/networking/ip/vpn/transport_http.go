package vpn

import (
	"context"
	"net"
	"net/http"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
)

type BatchRequest struct {
	IPs []string `json:"ips" validate:"required,min=1,max=50,dive,ip"`
}

func RegisterRoutes(r chi.Router, svc *Service) {
	r.Get("/ip/vpn/{ip}", httpx.Guard(svc, handleVPNCheck(svc)))
	r.Post("/ip/vpn/batch", handleVPNCheckBatch(svc))
}

// handleVPNCheck godoc
//
//	@Summary		Check IP Address
//	@Description	Analyzes an IP to determine if it is a VPN, proxy, Tor exit node, or hosting provider.
//	@Tags			vpn-detection
//	@Produce		json
//	@Param			ip	path		string	true	"IP address (IPv4 or IPv6)"
//	@Success		200	{object}	httpx.Response[IPCheckResponse]
//	@Failure		400	{object}	httpx.ErrorResponse
//	@Failure		500	{object}	httpx.ErrorResponse
//	@Router			/v1/networking/ip/vpn/{ip} [get]
func handleVPNCheck(svc *Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ip := net.ParseIP(getIP(r))
		if ip == nil {
			httpx.Error(w, http.StatusBadRequest, "bad_request", "invalid IP address")
			return
		}

		result, err := svc.CheckIP(r.Context(), ip)
		if err != nil {
			httpx.Error(w, http.StatusInternalServerError, "internal_error", "failed to check IP address")
			return
		}

		httpx.JSON(w, http.StatusOK, result)
	}
}

// handleVPNCheckBatch godoc
//
//	@Summary		Check IP Addresses (Batch)
//	@Description	Analyzes up to 50 IPs for VPN, proxy, Tor, or hosting provider signals.
//	@Tags			vpn-detection
//	@Accept			json
//	@Produce		json
//	@Param			request	body		BatchRequest	true	"List of IP addresses"
//	@Success		200		{object}	httpx.Response[httpx.BatchResponse[IPCheckResponse]]
//	@Failure		400		{object}	httpx.ErrorResponse
//	@Failure		422		{object}	httpx.ErrorResponse
//	@Router			/v1/networking/ip/vpn/batch [post]
func handleVPNCheckBatch(svc *Service) http.HandlerFunc {
	return httpx.HandleBatch(
		func(ctx context.Context, req BatchRequest) (httpx.BatchResponse[IPCheckResponse], error) {
			results, err := svc.CheckBatch(ctx, req.IPs)
			if err != nil {
				return httpx.BatchResponse[IPCheckResponse]{}, err
			}

			return httpx.BatchResponse[IPCheckResponse]{
				Results: results,
			}, nil
		},
	)
}

func getIP(r *http.Request) string {
	if ip := chi.URLParam(r, "ip"); ip != "" {
		return ip
	}

	addr := r.RemoteAddr
	if host, _, err := net.SplitHostPort(addr); err == nil {
		return host
	}

	return addr
}
