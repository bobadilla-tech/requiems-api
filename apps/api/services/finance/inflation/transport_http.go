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
	r.Get("/inflation", handleInflationRate(g))
	r.Post("/inflation/batch", handleInflationBatch(g))
}

// handleInflationRate godoc
//
//	@Summary		Get Inflation Rate
//	@Description	Returns latest annual CPI inflation rate for a country plus previous 10 years.
//	@Tags			inflation
//	@Produce		json
//	@Param			country	query		string	true	"ISO 3166-1 alpha-2 country code (case-insensitive)"
//	@Success		200		{object}	httpx.Response[Response]
//	@Failure		400		{object}	httpx.ErrorResponse
//	@Failure		404		{object}	httpx.ErrorResponse
//	@Failure		500		{object}	httpx.ErrorResponse
//	@Router			/v1/finance/inflation [get]
func handleInflationRate(g Getter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
	}
}

// handleInflationBatch godoc
//
//	@Summary		Batch Inflation Rates
//	@Description	Returns inflation data for up to 50 countries; countries with no data return `found: false`. Billing: 1 credit per country.
//	@Tags			inflation
//	@Accept			json
//	@Produce		json
//	@Param			request	body		BatchRequest	true	"List of ISO country codes"
//	@Success		200		{object}	httpx.Response[httpx.BatchResponse[BatchItem]]
//	@Failure		422		{object}	httpx.ErrorResponse
//	@Failure		500		{object}	httpx.ErrorResponse
//	@Router			/v1/finance/inflation/batch [post]
func handleInflationBatch(g Getter) http.HandlerFunc {
	return httpx.HandleBatch(
		func(ctx context.Context, req BatchRequest) (httpx.BatchResponse[BatchItem], error) {
			return httpx.BatchResponse[BatchItem]{Results: g.GetInflationBatch(ctx, req.Countries)}, nil
		},
	)
}
