package words

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
)

// RegisterRoutes mounts words handlers on the given router.
// Paths are relative to the parent mount point.
func RegisterRoutes(r chi.Router, svc *Service) {
	r.Get("/words/random", func(w http.ResponseWriter, r *http.Request) {
		wrd, err := svc.Random(r.Context())
		if err != nil {
			httpx.Error(w, http.StatusServiceUnavailable, "service_unavailable", "no words available")
			return
		}

		httpx.JSON(w, http.StatusOK, wrd)
	})

	r.Get("/dictionary/{word}", func(w http.ResponseWriter, r *http.Request) {
		word := chi.URLParam(r, "word")
		if word == "" {
			httpx.Error(w, http.StatusBadRequest, "bad_request", "word is required")
			return
		}

		entry, err := svc.Define(word)
		if err != nil {
			httpx.Error(w, http.StatusNotFound, "not_found", "word not found in dictionary")
			return
		}

		httpx.JSON(w, http.StatusOK, entry)
	})
	
	r.Post("/words/batch", httpx.HandleBatch(
		func(ctx context.Context, req BatchRequest) (BatchResponse, int, error) {

			resp, err := svc.BatchDefine(ctx, req)

			if err != nil {
				return BatchResponse{}, 0, err
			}

			return resp, len(req.Items), nil
		},
	))
}
