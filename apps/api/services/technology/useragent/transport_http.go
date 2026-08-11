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
	r.Get("/useragent", handleParseUserAgent(svc))
	r.Post("/useragent/batch", handleUserAgentBatch(svc))
}

// handleParseUserAgent godoc
//
//	@Summary		Parse User Agent
//	@Description	Parses a user agent string and returns browser, OS, device, and bot status.
//	@Tags			useragent
//	@Produce		json
//	@Param			ua	query		string	true	"Full user agent string"
//	@Success		200	{object}	httpx.Response[Result]
//	@Failure		400	{object}	httpx.ErrorResponse
//	@Router			/v1/technology/useragent [get]
func handleParseUserAgent(svc *Service) http.HandlerFunc {
	return httpx.HandleGet(func(ctx context.Context, req ParseRequest) (Result, error) {
		return svc.Parse(req.UA), nil
	})
}

// handleUserAgentBatch godoc
//
//	@Summary		Batch Parse User Agents
//	@Description	Parses up to 50 user agents.
//	@Tags			useragent
//	@Accept			json
//	@Produce		json
//	@Param			request	body		BatchParseRequest	true	"List of user agent strings"
//	@Success		200		{object}	httpx.Response[httpx.BatchResponse[BatchParseItem]]
//	@Failure		400		{object}	httpx.ErrorResponse
//	@Failure		422		{object}	httpx.ErrorResponse
//	@Router			/v1/technology/useragent/batch [post]
func handleUserAgentBatch(svc *Service) http.HandlerFunc {
	return httpx.HandleBatch(
		func(ctx context.Context, req BatchParseRequest) (httpx.BatchResponse[BatchParseItem], error) {
			return httpx.BatchResponse[BatchParseItem]{Results: svc.ParseBatch(ctx, req.UserAgents)}, nil
		},
	)
}
