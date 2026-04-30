package disposable

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
)

func RegisterRoutes(router chi.Router, svc *Service) {
	// POST /disposable/check — single email check
	router.Post("/disposable/check", httpx.Handle(
		func(_ context.Context, req CheckEmailRequest) (CheckEmailResponse, error) {
			return svc.CheckEmail(req.Email), nil
		},
	))

	// POST /disposable/batch — batch email check (max 100); X-Usage-Count = len(emails)
	router.Post("/disposable/batch", httpx.HandleBatch(
		func(_ context.Context, req BatchCheckRequest) (BatchCheckResponse, int, error) {
			return svc.CheckBatch(req.Emails), len(req.Emails), nil
		},
	))

	// GET /disposable/domain/{domain} — check a specific domain
	router.Get("/disposable/domain/{domain}", func(w http.ResponseWriter, r *http.Request) {
		domain := chi.URLParam(r, "domain")

		if domain == "" {
			httpx.Error(w, http.StatusBadRequest, "bad_request", "domain is required")
			return
		}

		httpx.JSON(w, http.StatusOK, svc.CheckDomain(domain))
	})

	// GET /disposable/domains — paginated list of all disposable domains
	router.Get("/disposable/domains", func(w http.ResponseWriter, r *http.Request) {
		page := 1
		perPage := 100

		if pageStr := r.URL.Query().Get("page"); pageStr != "" {
			p, err := strconv.Atoi(pageStr)
			if err != nil || p <= 0 {
				httpx.Error(w, http.StatusBadRequest, "bad_request", "page must be a positive integer")
				return
			}
			page = p
		}

		if perPageStr := r.URL.Query().Get("per_page"); perPageStr != "" {
			pp, err := strconv.Atoi(perPageStr)
			if err != nil || pp <= 0 {
				httpx.Error(w, http.StatusBadRequest, "bad_request", "per_page must be a positive integer")
				return
			}
			perPage = pp
		}

		result, err := svc.GetDomains(page, perPage)
		if err != nil {
			if appErr, ok := err.(*httpx.AppError); ok {
				httpx.Error(w, appErr.Status, appErr.Code, appErr.Message)
				return
			}
			httpx.Error(w, http.StatusInternalServerError, "internal_error", "unexpected error")
			return
		}

		httpx.JSON(w, http.StatusOK, result)
	})

	// GET /disposable/stats — blocklist statistics
	router.Get("/disposable/stats", func(w http.ResponseWriter, r *http.Request) {
		httpx.JSON(w, http.StatusOK, svc.GetStats())
	})
}
