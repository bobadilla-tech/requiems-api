package mortgage

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
)

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
	r.Post("/mortgage/batch", httpx.HandleBatch(func(ctx context.Context, req BatchRequest) (BatchResponse, error) {
		return c.CalculateBatch(req.Mortgages), nil
	}))
}
