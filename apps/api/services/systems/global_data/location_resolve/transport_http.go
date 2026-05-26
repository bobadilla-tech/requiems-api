package locationresolve

import (
	"context"

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
	r.Post("/location/resolve", httpx.Handle(
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
	))
}
