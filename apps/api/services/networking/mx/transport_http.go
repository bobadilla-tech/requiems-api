package mx

import (
	"context"
	"net"
	"net/http"
	"regexp"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
	"requiems-api/platform/middleware"
)

type BatchRequest struct {
	Domains []string `json:"domains" validate:"required,min=1,max=50,dive,required,fqdn"`
}

// domainRe matches valid fully-qualified domain names.
var domainRe = regexp.MustCompile(`^[a-zA-Z0-9]([a-zA-Z0-9\-]{0,61}[a-zA-Z0-9])?(\.[a-zA-Z0-9]([a-zA-Z0-9\-]{0,61}[a-zA-Z0-9])?)*\.[a-zA-Z]{2,}$`)

func RegisterRoutes(r chi.Router, svc *Service) {
	r.With(middleware.ValidateURLParam("domain", domainRe, "invalid domain name")).
		Get("/mx/{domain}", handleMXLookup(svc))

	r.Post("/mx/batch", handleMXBatch(svc))
}

// handleMXLookup godoc
//
//	@Summary		MX Lookup
//	@Description	Retrieves all MX records for a domain, sorted by priority ascending.
//	@Tags			mx-lookup
//	@Produce		json
//	@Param			domain	path		string	true	"Domain to look up (e.g. gmail.com)"
//	@Success		200		{object}	httpx.Response[LookupResponse]
//	@Failure		400		{object}	httpx.ErrorResponse
//	@Failure		404		{object}	httpx.ErrorResponse
//	@Failure		500		{object}	httpx.ErrorResponse
//	@Router			/v1/networking/mx/{domain} [get]
func handleMXLookup(svc *Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		domain := chi.URLParam(r, "domain")

		result, err := svc.Lookup(r.Context(), domain)

		if err != nil {
			if isDNSNotFound(err) {
				httpx.Error(w, http.StatusNotFound, "not_found", "no MX records found for domain")
				return
			}

			httpx.Error(w, http.StatusInternalServerError, "internal_error", "internal error")
			return
		}

		httpx.JSON(w, http.StatusOK, result)
	}
}

// handleMXBatch godoc
//
//	@Summary		Batch MX Lookup
//	@Description	Returns MX records for up to 50 domains.
//	@Tags			mx-lookup
//	@Accept			json
//	@Produce		json
//	@Param			request	body		BatchRequest	true	"List of domains to look up"
//	@Success		200		{object}	httpx.Response[httpx.BatchResponse[BatchLookupItem]]
//	@Failure		400		{object}	httpx.ErrorResponse
//	@Failure		422		{object}	httpx.ErrorResponse
//	@Failure		500		{object}	httpx.ErrorResponse
//	@Router			/v1/networking/mx/batch [post]
func handleMXBatch(svc *Service) http.HandlerFunc {
	return httpx.HandleBatch(
		func(ctx context.Context, req BatchRequest) (httpx.BatchResponse[BatchLookupItem], error) {
			return httpx.BatchResponse[BatchLookupItem]{Results: svc.LookupBatch(ctx, req.Domains)}, nil
		},
	)
}

// Reports whether the error is a DNS "no such host" / NXDOMAIN error.
func isDNSNotFound(err error) bool {
	if dnsErr, ok := err.(*net.DNSError); ok {
		return dnsErr.IsNotFound
	}

	return false
}
