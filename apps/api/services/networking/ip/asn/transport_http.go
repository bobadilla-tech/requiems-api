package asn

import (
	"context"
	"net"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
)

// BatchASNRequest is the input for the batch ASN lookup endpoint.
type BatchASNRequest struct {
	IPs []string `json:"ips" validate:"required,min=1,max=50,dive,ip"`
}

func RegisterRoutes(r chi.Router, svc *Service) {
	r.Get("/ip/asn/{ip}", httpx.Guard(svc, handleASNLookup(svc)))
	r.Get("/ip/asn", httpx.Guard(svc, handleOwnASNLookup(svc)))

	r.Post("/ip/asn/batch", httpx.Guard(svc, handleASNBatch(svc)))
}

// asnHandler returns the shared ASN lookup handler used by both the
// explicit-IP and caller-IP routes.
func asnHandler(svc *Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ipStr := chi.URLParam(r, "ip")

		if ipStr == "" {
			ipStr = httpx.ClientIP(r)
		}

		ip := net.ParseIP(ipStr)

		if ip == nil {
			httpx.Error(w, http.StatusBadRequest, "bad_request", "invalid IP address")
			return
		}

		result, err := svc.CheckASN(r.Context(), ip.String())

		if err != nil {
			if strings.Contains(err.Error(), "private/reserved") {
				httpx.JSON(w, http.StatusOK, IPAddressASNResponse{IP: ip.String()})
				return
			}

			httpx.Error(w, http.StatusInternalServerError, "internal_error", "internal error")
			return
		}

		httpx.JSON(w, http.StatusOK, result)
	}
}

// handleASNLookup godoc
//
//	@Summary		Lookup ASN for IP
//	@Description	Returns ASN, organization, ISP, domain, and network route information for a specific IP.
//	@Tags			ip-asn
//	@Produce		json
//	@Param			ip	path		string	true	"IP address (IPv4 or IPv6)"
//	@Success		200	{object}	httpx.Response[IPAddressASNResponse]
//	@Failure		400	{object}	httpx.ErrorResponse
//	@Failure		500	{object}	httpx.ErrorResponse
//	@Router			/v1/networking/ip/asn/{ip} [get]
func handleASNLookup(svc *Service) http.HandlerFunc {
	return asnHandler(svc)
}

// handleOwnASNLookup godoc
//
//	@Summary		Lookup ASN (Caller IP)
//	@Description	Returns ASN, organization, ISP, domain, and network route information for the requesting client's IP.
//	@Tags			ip-asn
//	@Produce		json
//	@Success		200	{object}	httpx.Response[IPAddressASNResponse]
//	@Failure		500	{object}	httpx.ErrorResponse
//	@Router			/v1/networking/ip/asn [get]
func handleOwnASNLookup(svc *Service) http.HandlerFunc {
	return asnHandler(svc)
}

// handleASNBatch godoc
//
//	@Summary		Batch ASN Lookup
//	@Description	Looks up ASN for up to 50 IPs. Private/reserved IPs return an empty result without error.
//	@Tags			ip-asn
//	@Accept			json
//	@Produce		json
//	@Param			request	body		BatchASNRequest	true	"List of IP addresses"
//	@Success		200		{object}	httpx.Response[httpx.BatchResponse[BatchASNItem]]
//	@Failure		400		{object}	httpx.ErrorResponse
//	@Failure		422		{object}	httpx.ErrorResponse
//	@Router			/v1/networking/ip/asn/batch [post]
func handleASNBatch(svc *Service) http.HandlerFunc {
	return httpx.HandleBatch(
		func(ctx context.Context, req BatchASNRequest) (httpx.BatchResponse[BatchASNItem], error) {
			return httpx.BatchResponse[BatchASNItem]{Results: svc.CheckASNBatch(ctx, req.IPs)}, nil
		},
	)
}
