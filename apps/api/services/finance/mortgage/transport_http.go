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
	// GET /mortgage?principal=300000&rate=6.5&years=30
	r.Get("/mortgage", func(w http.ResponseWriter, r *http.Request) {
		var req Request
		if err := httpx.BindQuery(r, &req); err != nil {
			httpx.Error(w, http.StatusBadRequest, "bad_request", err.Error())
			return
		}

		httpx.JSON(w, http.StatusOK, c.Calculate(req.Principal, req.Rate, req.Years))
	})
	r.Post("/mortgage/batch", httpx.HandleBatch(func(_ context.Context, req BatchRequest) (httpx.BatchResponse[Response], error) {
		return httpx.BatchResponse[Response]{Results: c.CalculateBatch(req.Mortgages)}, nil
	}))
}
