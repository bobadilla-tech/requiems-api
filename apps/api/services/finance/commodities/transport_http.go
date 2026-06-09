package commodities

import (
	"context"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
	"requiems-api/platform/svcerr"
)

// Getter is the interface used by the single-item HTTP transport layer.
type Getter interface {
	Get(ctx context.Context, slug string) (CommodityPrice, error)
}

// Batcher is the interface used by the batch HTTP transport layer.
type Batcher interface {
	GetBatch(ctx context.Context, slugs []string) []BatchCommodityItem
}

// BatchSlugsRequest is the input for the commodities batch endpoint.
type BatchSlugsRequest struct {
	Slugs []string `json:"slugs" validate:"required,min=1,max=20,dive,required"`
}

// RegisterRoutes mounts commodity price handlers on the given router.
// Paths are relative to the parent mount point (/v1/finance).
func RegisterRoutes(r chi.Router, svc *Service) {
	registerCommodityRoutes(r, svc)
	registerCommodityBatchRoutes(r, svc)
}

// registerCommodityBatchRoutes wires the Batcher interface to the router. Kept
// unexported so tests can inject a stub without going through the concrete
// *Service type.
func registerCommodityBatchRoutes(r chi.Router, b Batcher) {
	r.Post("/commodities/batch", httpx.HandleBatch(
		func(ctx context.Context, req BatchSlugsRequest) (httpx.BatchResponse[BatchCommodityItem], error) {
			return httpx.BatchResponse[BatchCommodityItem]{Results: b.GetBatch(ctx, req.Slugs)}, nil
		},
	))
}

// registerCommodityRoutes wires the Getter interface to the router. Kept
// unexported so tests can inject a stub without going through the concrete
// *Service type.
func registerCommodityRoutes(r chi.Router, g Getter) {
	// GET /commodities/{commodity} — return price data for a commodity slug
	r.Get("/commodities/{commodity}", func(w http.ResponseWriter, r *http.Request) {
		slug := chi.URLParam(r, "commodity")

		result, err := g.Get(r.Context(), slug)
		if err != nil {
			if se, ok := errors.AsType[*svcerr.Error](err); ok {
				httpx.Error(w, svcerr.HTTPStatus(se), se.Code, se.Message)
				return
			}
			httpx.Error(w, http.StatusInternalServerError, "internal_error", "internal server error")
			return
		}

		httpx.JSON(w, http.StatusOK, result)
	})
}
