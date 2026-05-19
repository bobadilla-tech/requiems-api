package facts

import (
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
)

// Request holds optional query parameters for the facts endpoint.
type Request struct {
	Category string `query:"category"`
}

// RegisterRoutes mounts facts handlers on the given router.
// Paths are relative to the parent mount point (e.g. /v1/entertainment).
func RegisterRoutes(r chi.Router, svc *Service) {
	r.Get("/facts", func(w http.ResponseWriter, r *http.Request) {
		req := Request{}
		if err := httpx.BindQuery(r, &req); err != nil {
			httpx.Error(w, http.StatusBadRequest, "bad_request", err.Error())
			return
		}

		category := strings.ToLower(req.Category)
		if category != "" && !IsValidCategory(category) {
			httpx.Error(w, http.StatusBadRequest, "bad_request", "invalid category")
			return
		}

		fact, err := svc.Random(category)
		if err != nil {
			httpx.Error(w, http.StatusInternalServerError, "internal_error", "failed to retrieve fact")
			return
		}

		httpx.JSON(w, http.StatusOK, fact)
	})
}
