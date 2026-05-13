package disposable

import (
	"context"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
	"requiems-api/platform/svcerr"
)

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

	router.Get("/disposable/domains", func(w http.ResponseWriter, r *http.Request) {
		q := DomainsListQuery{Page: 1, PerPage: 100}

		if err := httpx.BindQuery(r, &q); err != nil {
			if vf, ok := errors.AsType[*httpx.ValidationFailure](err); ok {
				httpx.ValidationError(w, vf)
				return
			}

			httpx.Error(w, http.StatusBadRequest, "bad_request", err.Error())
			return
		}

		result, err := svc.GetDomains(q.Page, q.PerPage)

		if err != nil {
			if se, ok := errors.AsType[*svcerr.Error](err); ok {
				httpx.Error(w, svcerr.HTTPStatus(se), se.Code, se.Message)
				return
			}
			httpx.Error(w, http.StatusInternalServerError, "internal_error", "unexpected error")
			return
		}

		httpx.JSON(w, http.StatusOK, result)
	})

	router.Get("/disposable/stats", func(w http.ResponseWriter, r *http.Request) {
		httpx.JSON(w, http.StatusOK, svc.GetStats())
	})
}
