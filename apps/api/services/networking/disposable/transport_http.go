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
	router.Post("/disposable/check", handleDisposableCheck(svc))
	router.Post("/disposable/batch", handleDisposableBatch(svc))
	router.Get("/disposable/domain/{domain}", handleDisposableDomain(svc))
	router.Get("/disposable/domains", handleDisposableDomains(svc))
	router.Get("/disposable/stats", handleDisposableStats(svc))
}

// handleDisposableCheck godoc
//
//	@Summary		Check Single Email
//	@Description	Validates whether an email address uses a disposable domain.
//	@Tags			disposable-email
//	@Accept			json
//	@Produce		json
//	@Param			request	body		CheckEmailRequest	true	"Email address to check"
//	@Success		200		{object}	httpx.Response[CheckEmailResponse]
//	@Failure		400		{object}	httpx.ErrorResponse
//	@Router			/v1/networking/disposable/check [post]
func handleDisposableCheck(svc *Service) http.HandlerFunc {
	return httpx.Handle(
		func(_ context.Context, req CheckEmailRequest) (CheckEmailResponse, error) {
			return svc.CheckEmail(req.Email), nil
		},
	)
}

// handleDisposableBatch godoc
//
//	@Summary		Check Batch Emails
//	@Description	Validates up to 100 email addresses for disposable domains.
//	@Tags			disposable-email
//	@Accept			json
//	@Produce		json
//	@Param			request	body		BatchCheckRequest	true	"List of email addresses to check"
//	@Success		200		{object}	httpx.Response[httpx.BatchResponse[CheckEmailResponse]]
//	@Failure		400		{object}	httpx.ErrorResponse
//	@Router			/v1/networking/disposable/batch [post]
func handleDisposableBatch(svc *Service) http.HandlerFunc {
	return httpx.HandleBatch(
		func(_ context.Context, req BatchCheckRequest) (httpx.BatchResponse[CheckEmailResponse], error) {
			return httpx.BatchResponse[CheckEmailResponse]{Results: svc.CheckBatch(req.Emails)}, nil
		},
	)
}

// handleDisposableDomain godoc
//
//	@Summary		Check Domain
//	@Description	Checks if a specific domain is in the disposable blocklist.
//	@Tags			disposable-email
//	@Produce		json
//	@Param			domain	path		string	true	"Domain to check (e.g. tempmail.com)"
//	@Success		200		{object}	httpx.Response[DomainCheckResponse]
//	@Failure		400		{object}	httpx.ErrorResponse
//	@Router			/v1/networking/disposable/domain/{domain} [get]
func handleDisposableDomain(svc *Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		domain := chi.URLParam(r, "domain")

		if domain == "" {
			httpx.Error(w, http.StatusBadRequest, "bad_request", "domain is required")
			return
		}

		httpx.JSON(w, http.StatusOK, svc.CheckDomain(domain))
	}
}

// handleDisposableDomains godoc
//
//	@Summary		List Domains (Paginated)
//	@Description	Gets a paginated list of all disposable domains in the blocklist.
//	@Tags			disposable-email
//	@Produce		json
//	@Param			page		query		integer	false	"Page number (default 1)"
//	@Param			per_page	query		integer	false	"Results per page (default 100, max 1000)"
//	@Success		200			{object}	httpx.Response[DomainsListResponse]
//	@Router			/v1/networking/disposable/domains [get]
func handleDisposableDomains(svc *Service) http.HandlerFunc {
	return httpx.HandleGet(func(ctx context.Context, q DomainsListQuery) (DomainsListResponse, error) {
		return svc.GetDomains(q.Page, q.PerPage)
	}, DomainsListQuery{Page: 1, PerPage: 100})
}

// handleDisposableStats godoc
//
//	@Summary		Get Statistics
//	@Description	Gets statistics about the disposable email blocklist.
//	@Tags			disposable-email
//	@Produce		json
//	@Success		200	{object}	httpx.Response[StatsResponse]
//	@Router			/v1/networking/disposable/stats [get]
func handleDisposableStats(svc *Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		httpx.JSON(w, http.StatusOK, svc.GetStats())
	}
}
