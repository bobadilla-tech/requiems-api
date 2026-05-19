package disposable

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
)

// CheckEmailRequest is the body for a single-email disposable check.
type CheckEmailRequest struct {
	Email string `json:"email" validate:"required,email"`
}

// BatchCheckRequest is the body for checking multiple emails at once.
type BatchCheckRequest struct {
	Emails []string `json:"emails" validate:"required,min=1,max=100,dive,email"`
}

// DomainsListQuery holds optional pagination parameters for domain listing.
type DomainsListQuery struct {
	Page    int `query:"page"     validate:"min=1"`
	PerPage int `query:"per_page" validate:"min=1"`
}

func RegisterRoutes(router chi.Router, svc *Service) {
	router.Post("/disposable/check", httpx.Handle(
		func(_ context.Context, req CheckEmailRequest) (CheckEmailResponse, error) {
			return svc.CheckEmail(req.Email), nil
		},
	))

	router.Post("/disposable/batch", httpx.HandleBatch(
		func(_ context.Context, req BatchCheckRequest) (httpx.BatchResponse[CheckEmailResponse], error) {
			return httpx.BatchResponse[CheckEmailResponse]{Results: svc.CheckBatch(req.Emails)}, nil
		},
	))

	router.Get("/disposable/domain/{domain}", func(w http.ResponseWriter, r *http.Request) {
		domain := chi.URLParam(r, "domain")

		if domain == "" {
			httpx.Error(w, http.StatusBadRequest, "bad_request", "domain is required")
			return
		}

		httpx.JSON(w, http.StatusOK, svc.CheckDomain(domain))
	})

	router.Get("/disposable/domains", httpx.HandleGet(func(ctx context.Context, q DomainsListQuery) (DomainsListResponse, error) {
		return svc.GetDomains(q.Page, q.PerPage)
	}, DomainsListQuery{Page: 1, PerPage: 100}))

	router.Get("/disposable/stats", func(w http.ResponseWriter, r *http.Request) {
		httpx.JSON(w, http.StatusOK, svc.GetStats())
	})
}
