package mortgage

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
)

// Request holds the validated query parameters for the mortgage endpoint.
type Request struct {
	Principal float64 `query:"principal" validate:"required,gt=0"`
	Rate      float64 `query:"rate"      validate:"required,gt=0"`
	Years     int     `query:"years"     validate:"required,min=1,max=50"`
}

// BatchRequest is the body for calculating multiple mortgages at once.
type BatchRequest struct {
	Mortgages []Request `json:"mortgages" validate:"required,min=1,max=50,dive"`
}

// Calculator is the interface used by the HTTP transport layer.
type Calculator interface {
	Calculate(principal, annualRate float64, years int) Response
	CalculateBatch(mortgages []Request) []Response
}

// RegisterRoutes mounts mortgage handlers on the given router.
// Paths are relative to the parent mount point (/v1/finance).
func RegisterRoutes(r chi.Router, svc *Service) {
	registerMortgageRoutes(r, svc)
}

// registerMortgageRoutes wires the Calculator interface to the router. Kept
// unexported so tests can inject a stub without going through the concrete
// *Service type.
func registerMortgageRoutes(r chi.Router, c Calculator) {
	r.Get("/mortgage", handleMortgageCalculate(c))
	r.Post("/mortgage/batch", handleMortgageBatch(c))
}

// handleMortgageCalculate godoc
//
//	@Summary		Calculate Mortgage
//	@Description	Returns monthly payment, total cost, and full amortization schedule.
//	@Tags			mortgage
//	@Produce		json
//	@Param			principal	query		number	true	"Loan amount"
//	@Param			rate		query		number	true	"Annual interest rate in percent, greater than 0"
//	@Param			years		query		integer	true	"Loan term in years (1–50)"
//	@Success		200			{object}	httpx.Response[Response]
//	@Failure		400			{object}	httpx.ErrorResponse
//	@Failure		500			{object}	httpx.ErrorResponse
//	@Router			/v1/finance/mortgage [get]
func handleMortgageCalculate(c Calculator) http.HandlerFunc {
	return httpx.HandleGet(func(ctx context.Context, req Request) (Response, error) {
		return c.Calculate(req.Principal, req.Rate, req.Years), nil
	})
}

// handleMortgageBatch godoc
//
//	@Summary		Batch Calculate Mortgages
//	@Description	Calculates up to 50 mortgages in a single request.
//	@Tags			mortgage
//	@Accept			json
//	@Produce		json
//	@Param			request	body		BatchRequest	true	"List of mortgage parameters"
//	@Success		200		{object}	httpx.Response[httpx.BatchResponse[Response]]
//	@Failure		422		{object}	httpx.ErrorResponse
//	@Router			/v1/finance/mortgage/batch [post]
func handleMortgageBatch(c Calculator) http.HandlerFunc {
	return httpx.HandleBatch(func(_ context.Context, req BatchRequest) (httpx.BatchResponse[Response], error) {
		return httpx.BatchResponse[Response]{Results: c.CalculateBatch(req.Mortgages)}, nil
	})
}
