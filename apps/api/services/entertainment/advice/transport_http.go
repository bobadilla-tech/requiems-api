package advice

import (
	"encoding/json"
	"net/http"
	
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
	r.Post("/advice/batch", func(w http.ResponseWriter, r *http.Request) {
		var req BatchRequest

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			httpx.Error(w, http.StatusBadRequest, "invalid_request", "invalid json body")
			return
		}

		if req.Count <= 0 {
			httpx.Error(w, http.StatusBadRequest, "invalid_count", "count must be greater than zero")
			return
		}

		results, err := svc.RandomBatch(r.Context(), req.Count)
		if err != nil {
			httpx.Error(w, http.StatusServiceUnavailable, "service_unavailable", "no advice available")
			return
		}

		httpx.JSON(w, http.StatusOK, BatchResponse[Advice]{
			Results: results,
		})
	})
}
