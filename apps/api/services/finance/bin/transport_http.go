package bin

import (
	"context"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
	"requiems-api/platform/svcerr"
)

// Looker is the interface used by the HTTP transport layer.
type Looker interface {
	Lookup(ctx context.Context, bin string) (LookupResponse, error)
}

// RegisterRoutes mounts BIN lookup handlers on the given router.
// Paths are relative to the parent mount point (/v1/finance).
func RegisterRoutes(r chi.Router, svc *Service) {
	registerBINRoutes(r, svc)
}

// registerBINRoutes wires the Looker interface to the router. Kept unexported
// so tests can inject a stub without going through the concrete *Service type.
func registerBINRoutes(r chi.Router, l Looker) {
	// GET /bin/{bin} — look up card metadata for a 6–8 digit BIN prefix
	r.Get("/bin/{bin}", func(w http.ResponseWriter, r *http.Request) {
		rawBIN := chi.URLParam(r, "bin")

		result, err := l.Lookup(r.Context(), rawBIN)
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
