package locationresolve

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
	"requiems-api/platform/svcerr"
)

type Request struct {
	Address     string       `json:"address"`
	Coordinates *Coordinates `json:"coordinates"`
	CountryCode string       `json:"country_code"`
}

func RegisterRoutes(r chi.Router, svc *Service) {
	r.Post("/location/resolve", handleLocationResolve(svc))
}

// handleLocationResolve godoc
//
//	@Summary		Resolve Location
//	@Description	Resolves an address or coordinates into a full location profile — city, country, timezone, UTC offset, current time, working days this month, next holiday.
//	@Tags			global-data
//	@Accept			json
//	@Produce		json
//	@Param			request	body		Request	true	"Address or coordinates to resolve"
//	@Success		200		{object}	httpx.Response[Result]
//	@Failure		401		{object}	httpx.ErrorResponse
//	@Failure		422		{object}	httpx.ErrorResponse
//	@Router			/v1/systems/location/resolve [post]
func handleLocationResolve(svc *Service) http.HandlerFunc {
	return httpx.Handle(
		func(ctx context.Context, req Request) (Result, error) {
			if req.Address == "" && req.Coordinates == nil {
				return Result{}, svcerr.Unknown("validation_failed", "address or coordinates is required")
			}

			res, err := svc.Resolve(ctx, req)
			if err != nil {
				if _, ok := err.(*missingInputError); ok {
					return Result{}, svcerr.Unknown("validation_failed", err.Error())
				}
				if se, ok := err.(*svcerr.Error); ok {
					return Result{}, se
				}
				return Result{}, svcerr.Upstream("upstream_error", "geocoding service unavailable")
			}

			return res, nil
		},
	)
}
