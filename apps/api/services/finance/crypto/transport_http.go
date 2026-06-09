package cryptocoin

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
	"requiems-api/platform/svcerr"
)

// Getter is the interface used by the HTTP transport layer.
type Getter interface {
	GetPrice(ctx context.Context, symbol string) (Price, error)
}

// BatchSymbolsRequest is the input for the crypto batch endpoint.
type BatchSymbolsRequest struct {
	Symbols []string `json:"symbols" validate:"required,min=1,max=20,dive,required"`
}

// RegisterRoutes mounts crypto price handlers on the given router.
// Paths are relative to the parent mount point (/v1/finance).
func RegisterRoutes(r chi.Router, svc *Service) {
	registerCryptoRoutes(r, svc)

	r.Post("/crypto/batch", httpx.HandleBatch(
		func(ctx context.Context, req BatchSymbolsRequest) (httpx.BatchResponse[BatchPriceItem], error) {
			return httpx.BatchResponse[BatchPriceItem]{Results: svc.GetPriceBatch(ctx, req.Symbols)}, nil
		},
	))
}

// registerCryptoRoutes wires the Getter interface to the router. Kept
// unexported so tests can inject a stub without going through the concrete
// *Service type.
func registerCryptoRoutes(r chi.Router, g Getter) {
	// GET /crypto/{symbol}
	r.Get("/crypto/{symbol}", func(w http.ResponseWriter, r *http.Request) {
		symbol := strings.ToUpper(chi.URLParam(r, "symbol"))

		price, err := g.GetPrice(r.Context(), symbol)
		if err != nil {
			if se, ok := errors.AsType[*svcerr.Error](err); ok {
				httpx.Error(w, svcerr.HTTPStatus(se), se.Code, se.Message)
				return
			}
			httpx.Error(w, http.StatusServiceUnavailable, "upstream_error", "crypto price service unavailable")
			return
		}

		httpx.JSON(w, http.StatusOK, price)
	})
}
