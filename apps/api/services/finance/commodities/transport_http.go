package commodities

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
	"requiems-api/platform/svcerr"
)

// RegisterRoutes mounts commodity price handlers on the given router.
// Paths are relative to the parent mount point (/v1/finance).
func RegisterRoutes(r chi.Router, svc *Service) {
	registerCommodityRoutes(r, svc)
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
			var se *svcerr.Error
			if errors.As(err, &se) {
				httpx.Error(w, svcerr.HTTPStatus(se), se.Code, se.Message)
				return
			}
			httpx.Error(w, http.StatusInternalServerError, "internal_error", "internal server error")
			return
		}

		httpx.JSON(w, http.StatusOK, result)
	})
}
