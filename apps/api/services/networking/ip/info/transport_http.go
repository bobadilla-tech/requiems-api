package info

import (
	"context"
	"net"
	"net/http"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
)

// BatchInfoRequest is the input for the batch IP geolocation endpoint.
type BatchInfoRequest struct {
	IPs []string `json:"ips" validate:"required,min=1,max=50,dive,ip"`
}

func RegisterRoutes(r chi.Router, svc *Service) {
	r.Get("/ip/{ip}", httpx.Guard(svc, handleIPInfo(svc)))
	r.Get("/ip", httpx.Guard(svc, handleOwnIPInfo(svc)))

	r.Post("/ip/info/batch", httpx.Guard(svc, handleIPInfoBatch(svc)))
}

// ipInfoHandler returns the shared geolocation handler used by both the
// explicit-IP and caller-IP routes.
func ipInfoHandler(svc *Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ipStr := chi.URLParam(r, "ip")
		if ipStr == "" {
			ipStr = httpx.ClientIP(r)
		}

		if net.ParseIP(ipStr) == nil {
			httpx.Error(w, http.StatusBadRequest, "bad_request", "invalid IP address")
			return
		}

		result, err := svc.CheckInfo(r.Context(), ipStr)
		if err != nil {
			httpx.Error(w, http.StatusInternalServerError, "internal_error", "internal error")
			return
		}

		httpx.JSON(w, http.StatusOK, result)
	}
}

// handleIPInfo godoc
//
//	@Summary		Get IP Info for IP
//	@Description	Geolocation for a specific IP (IPv4 and IPv6).
//	@Tags			ip-info
//	@Produce		json
//	@Param			ip	path		string	true	"IP address (IPv4 or IPv6)"
//	@Success		200	{object}	httpx.Response[LookupResponse]
//	@Failure		400	{object}	httpx.ErrorResponse
//	@Failure		500	{object}	httpx.ErrorResponse
//	@Router			/v1/networking/ip/{ip} [get]
func handleIPInfo(svc *Service) http.HandlerFunc {
	return ipInfoHandler(svc)
}

// handleOwnIPInfo godoc
//
//	@Summary		Get IP Info (Caller IP)
//	@Description	Geolocation for the requesting client's IP.
//	@Tags			ip-info
//	@Produce		json
//	@Success		200	{object}	httpx.Response[LookupResponse]
//	@Failure		500	{object}	httpx.ErrorResponse
//	@Router			/v1/networking/ip [get]
func handleOwnIPInfo(svc *Service) http.HandlerFunc {
	return ipInfoHandler(svc)
}

// handleIPInfoBatch godoc
//
//	@Summary		Batch IP Geolocation
//	@Description	Looks up geolocation for up to 50 IPs; per-item errors inline.
//	@Tags			ip-info
//	@Accept			json
//	@Produce		json
//	@Param			request	body		BatchInfoRequest	true	"List of IP addresses"
//	@Success		200		{object}	httpx.Response[httpx.BatchResponse[BatchIPInfoItem]]
//	@Failure		400		{object}	httpx.ErrorResponse
//	@Failure		422		{object}	httpx.ErrorResponse
//	@Router			/v1/networking/ip/info/batch [post]
func handleIPInfoBatch(svc *Service) http.HandlerFunc {
	return httpx.HandleBatch(
		func(ctx context.Context, req BatchInfoRequest) (httpx.BatchResponse[BatchIPInfoItem], error) {
			return httpx.BatchResponse[BatchIPInfoItem]{Results: svc.CheckInfoBatch(ctx, req.IPs)}, nil
		},
	)
}
