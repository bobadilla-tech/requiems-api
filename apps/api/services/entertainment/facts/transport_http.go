package facts

import (
	"net/http"
	"strings"
	"context"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
)

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
	r.Post("/facts/batch", httpx.HandleBatch(
		func(ctx context.Context, req BatchRequest) (httpx.BatchResponse[BatchItem], error) {
			// Normalise to lowercase — consistent with GET /facts?category=
			normalised := make([]string, len(req.Categories))
			for i, c := range req.Categories {
				normalised[i] = strings.ToLower(c)
			}
			return svc.RandomBatch(ctx, normalised), nil
		},
	))
}
