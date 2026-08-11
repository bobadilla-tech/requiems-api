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
	r.Post("/commodities/batch", handleCommoditiesBatch(b))
}

// registerCommodityRoutes wires the Getter interface to the router. Kept
// unexported so tests can inject a stub without going through the concrete
// *Service type.
func registerCommodityRoutes(r chi.Router, g Getter) {
	r.Get("/commodities/{commodity}", handleCommodityPrice(g))
}

// handleCommodityPrice godoc
//
//	@Summary		Get Commodity Price
//	@Description	Returns latest annual average price and up to 10 years of history for the commodity slug.
//	@Tags			commodities
//	@Produce		json
//	@Param			commodity	path		string	true	"Commodity slug (e.g. gold, silver, oil)"
//	@Success		200			{object}	httpx.Response[CommodityPrice]
//	@Failure		404			{object}	httpx.ErrorResponse
//	@Failure		500			{object}	httpx.ErrorResponse
//	@Router			/v1/finance/commodities/{commodity} [get]
func handleCommodityPrice(g Getter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
	}
}

// handleCommoditiesBatch godoc
//
//	@Summary		Get Commodity Prices (Batch)
//	@Description	Returns latest prices for up to 20 commodity slugs.
//	@Tags			commodities
//	@Accept			json
//	@Produce		json
//	@Param			request	body		BatchSlugsRequest	true	"List of commodity slugs"
//	@Success		200		{object}	httpx.Response[httpx.BatchResponse[BatchCommodityItem]]
//	@Failure		422		{object}	httpx.ErrorResponse
//	@Router			/v1/finance/commodities/batch [post]
func handleCommoditiesBatch(b Batcher) http.HandlerFunc {
	return httpx.HandleBatch(
		func(ctx context.Context, req BatchSlugsRequest) (httpx.BatchResponse[BatchCommodityItem], error) {
			return httpx.BatchResponse[BatchCommodityItem]{Results: b.GetBatch(ctx, req.Slugs)}, nil
		},
	)
}
