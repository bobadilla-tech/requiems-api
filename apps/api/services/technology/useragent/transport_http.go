package useragent

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
)

// ParseRequest holds the query parameters for the user agent parse endpoint.
type ParseRequest struct {
	UA string `query:"ua" validate:"required"`
}

// BatchParseRequest holds the body for validating multiple user agents at once.
type BatchParseRequest struct {
	UserAgents []string `json:"user_agents" validate:"required,min=1,max=50,dive,required"`
}

// BatchParseItem represents the parsed result for a single user agent in a batch request.
type BatchParseItem struct {
	UserAgent string `json:"user_agent"`
	Data      Result `json:"data"`
}

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
