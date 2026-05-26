package inflation

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
	"requiems-api/platform/svcerr"
)

// Request holds the validated query parameters for the inflation endpoint.
type Request struct {
	Country string `query:"country" validate:"required,iso3166_1_alpha2"`
}

// BatchRequest is the body for fetching inflation data for multiple countries at once.
type BatchRequest struct {
	Countries []string `json:"countries" validate:"required,min=1,max=50,dive,iso3166_1_alpha2"`
}

// Getter is the interface used by the HTTP transport layer.
type Getter interface {
	GetInflation(ctx context.Context, countryCode string) (Response, error)
	GetInflationBatch(ctx context.Context, countries []string) []BatchItem
}

// RegisterRoutes mounts inflation handlers on the given router.
// Paths are relative to the parent mount point (/v1/finance).
func RegisterRoutes(r chi.Router, svc *Service) {
	registerInflationRoutes(r, svc)
}

// registerInflationRoutes wires the Getter interface to the router. Kept
// unexported so tests can inject a stub without going through the concrete
// *Service type.
func registerInflationRoutes(r chi.Router, g Getter) {
	// GET /inflation?country=US — return latest and historical CPI inflation rate
	r.Get("/inflation", func(w http.ResponseWriter, r *http.Request) {
		// Uppercase the country param before binding — iso3166_1_alpha2 is case-sensitive.
		if country := r.URL.Query().Get("country"); country != "" {
			q := r.URL.Query()
			q.Set("country", strings.ToUpper(country))
			r.URL.RawQuery = q.Encode()
		}

		var req Request
		if err := httpx.BindQuery(r, &req); err != nil {
			httpx.Error(w, http.StatusBadRequest, "bad_request", err.Error())
			return
		}

		resp, err := g.GetInflation(r.Context(), req.Country)
		if err != nil {
			if se, ok := errors.AsType[*svcerr.Error](err); ok {
				httpx.Error(w, svcerr.HTTPStatus(se), se.Code, se.Message)
				return
			}
			httpx.Error(w, http.StatusInternalServerError, "internal_error", "internal server error")
			return
		}

		httpx.JSON(w, http.StatusOK, resp)
	})

	// POST /inflation/batch — return inflation data for up to 50 countries at once.
	// Uses HandleBatch so the gateway charges one credit per country (X-Usage-Count).
	r.Post("/inflation/batch", httpx.HandleBatch(
		func(ctx context.Context, req BatchRequest) (httpx.BatchResponse[BatchItem], error) {
			return httpx.BatchResponse[BatchItem]{Results: g.GetInflationBatch(ctx, req.Countries)}, nil
		},
	))
}
