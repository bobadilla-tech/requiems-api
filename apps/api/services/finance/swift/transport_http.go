package swift

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
	"requiems-api/platform/svcerr"
)

// RegisterRoutes mounts SWIFT/BIC lookup handlers on the given router.
// Paths are relative to the parent mount point (/v1/finance).
func RegisterRoutes(r chi.Router, svc *Service) {
	registerSWIFTRoutes(r, svc)
}

// registerSWIFTRoutes wires the Looker interface to the router. Kept unexported
// so tests can inject a stub without going through the concrete *Service type.
func registerSWIFTRoutes(r chi.Router, l Looker) {
	// GET /swift — list SWIFT records with optional filters.
	r.Get("/swift", func(w http.ResponseWriter, r *http.Request) {
		filter := ListFilter{Limit: 50}
		if err := httpx.BindQuery(r, &filter); err != nil {
			httpx.Error(w, http.StatusBadRequest, "bad_request", err.Error())
			return
		}

		result, err := l.List(r.Context(), filter)
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

	// GET /swift/{code} — look up bank metadata for a SWIFT/BIC code
	r.Get("/swift/{code}", func(w http.ResponseWriter, r *http.Request) {
		rawCode := chi.URLParam(r, "code")

		result, err := l.Lookup(r.Context(), rawCode)
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
