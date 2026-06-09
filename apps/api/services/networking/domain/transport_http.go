package domain

import (
	"context"
	"net/http"
	"regexp"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
	"requiems-api/platform/middleware"
)

// domainRe accepts standard hostnames such as "example.com" or "sub.example.co.uk".
// Each label is 1–63 chars (alphanumeric or hyphens, not starting/ending with a
// hyphen), and there must be at least one dot separating the labels.
var domainRe = regexp.MustCompile(`^([a-zA-Z0-9]([a-zA-Z0-9\-]{0,61}[a-zA-Z0-9])?\.)+[a-zA-Z]{2,}$`)

// BatchDomainRequest is the input for the domain batch endpoint.
type BatchDomainRequest struct {
	Domains []string `json:"domains" validate:"required,min=1,max=10,dive,required"`
}

// InfoBatcher is the interface used by the domain batch HTTP transport layer.
type InfoBatcher interface {
	GetInfoBatch(ctx context.Context, domains []string) []BatchDomainItem
}

func RegisterRoutes(r chi.Router, svc *Service) {
	r.Group(func(validated chi.Router) {
		validated.Use(middleware.ValidateURLParam("domain", domainRe, "invalid domain: must be a valid hostname such as example.com"))

		validated.Get("/domain/{domain}", func(w http.ResponseWriter, r *http.Request) {
			d := chi.URLParam(r, "domain")
			httpx.JSON(w, http.StatusOK, svc.GetInfo(r.Context(), d))
		})
	})

	registerDomainBatchRoutes(r, svc)
}

// registerDomainBatchRoutes wires the InfoBatcher interface to the router. Kept
// unexported so tests can inject a stub without live DNS lookups.
func registerDomainBatchRoutes(r chi.Router, b InfoBatcher) {
	r.Post("/domain/batch", httpx.HandleBatch(
		func(ctx context.Context, req BatchDomainRequest) (httpx.BatchResponse[BatchDomainItem], error) {
			return httpx.BatchResponse[BatchDomainItem]{Results: b.GetInfoBatch(ctx, req.Domains)}, nil
		},
	))
}
