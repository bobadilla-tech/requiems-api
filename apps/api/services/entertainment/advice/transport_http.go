package advice

import (
	"net/http"
	"fmt"
	"context"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
)

// RegisterRoutes mounts advice handlers on the given router.
// Paths are relative to the parent mount point.
func RegisterRoutes(r chi.Router, svc *Service) {
	r.Get("/advice", func(w http.ResponseWriter, r *http.Request) {
		a, err := svc.Random(r.Context())
		if err != nil {
			httpx.Error(w, http.StatusServiceUnavailable, "service_unavailable", "no advice available")
			return
		}

		httpx.JSON(w, http.StatusOK, a)
	})
	r.Post("/advice/batch", httpx.HandleBatch(
	func(ctx context.Context, req BatchRequest) (httpx.BatchResponse[Advice], error) {

		if req.Count <= 0 {
			return httpx.BatchResponse[Advice]{},
				fmt.Errorf("count must be greater than zero")
		}

		if req.Count > 50 {
			return httpx.BatchResponse[Advice]{},
				fmt.Errorf("max 50 items allowed")
		}

		results, err := svc.RandomBatch(ctx, req.Count)
		if err != nil {
			return httpx.BatchResponse[Advice]{},
				fmt.Errorf("no advice available: %w", err)
		}

		return httpx.BatchResponse[Advice]{
			Results: results,
		}, nil
	},
))
}
