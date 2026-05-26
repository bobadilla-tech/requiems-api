package whois

import (
	"context"
	"errors"
	"net/http"
	"regexp"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
	"requiems-api/platform/middleware"
)

// BatchLookupRequest is the body for looking up multiple domains at once.
type BatchLookupRequest struct {
	Domains []string `json:"domains" validate:"required,min=1,max=50,dive,required,hostname_rfc1123"`
}

// domainPattern matches valid domain names (e.g., "example.com", "sub.example.co.uk").
var domainPattern = regexp.MustCompile(`^[a-zA-Z0-9]([a-zA-Z0-9\-]{0,61}[a-zA-Z0-9])?(\.[a-zA-Z0-9]([a-zA-Z0-9\-]{0,61}[a-zA-Z0-9])?)+$`)

func RegisterRoutes(r chi.Router, svc *Service) {
	r.With(middleware.ValidateURLParam("domain", domainPattern, "invalid domain name")).
		Get("/whois/{domain}", func(w http.ResponseWriter, r *http.Request) {
			domain := chi.URLParam(r, "domain")

			result, err := svc.Lookup(r.Context(), domain)

			if err != nil {
				if errors.Is(err, ErrDomainNotFound) {
					httpx.Error(w, http.StatusNotFound, "not_found", "domain not found")
					return
				}

				httpx.Error(w, http.StatusInternalServerError, "internal_error", "whois lookup failed")
				return
			}

			httpx.JSON(w, http.StatusOK, result)
		})

	r.Post("/whois/batch", httpx.HandleBatch(
		func(ctx context.Context, req BatchLookupRequest) (httpx.BatchResponse[BatchLookupItem], error) {
			items, err := svc.LookupBatch(ctx, req.Domains)
			return httpx.BatchResponse[BatchLookupItem]{Results: items}, err
		},
	))

}
