package useragent

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
)

func RegisterRoutes(r chi.Router, svc *Service) {
	r.Get("/useragent", func(w http.ResponseWriter, r *http.Request) {
		var req ParseRequest

		if err := httpx.BindQuery(r, &req); err != nil {
			httpx.Error(w, http.StatusBadRequest, "bad_request", err.Error())
			return
		}

		httpx.JSON(w, http.StatusOK, svc.Parse(req.UA))
	})
	r.Post("/useragent/batch", httpx.HandleBatch(
		func(ctx context.Context, req BatchParseRequest) (httpx.BatchResponse[BatchParseItem], error) {
			return httpx.BatchResponse[BatchParseItem]{Results: svc.ParseBatch(ctx, req.UserAgents), Total: len(req.UserAgents)}, nil
		},
	))
}
