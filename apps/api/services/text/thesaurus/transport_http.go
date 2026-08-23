package thesaurus

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
)

// BatchThesaurusRequest is the input for the thesaurus batch endpoint.
type BatchThesaurusRequest struct {
	Words []string `json:"words" validate:"required,min=1,max=50,dive,required"`
}

// RegisterRoutes mounts thesaurus handlers on the given router.
// Paths are relative to the parent mount point.
func RegisterRoutes(r chi.Router, svc *Service) {
	r.Get("/thesaurus/{word}", func(w http.ResponseWriter, r *http.Request) {
		word := chi.URLParam(r, "word")
		if word == "" {
			httpx.Error(w, http.StatusBadRequest, "bad_request", "word is required")
			return
		}

		result, err := svc.Lookup(word)
		if err != nil {
			httpx.Error(w, http.StatusNotFound, "not_found", "word not found in thesaurus")
			return
		}

		w.Header().Set("X-Usage-Count", "2")
		httpx.JSON(w, http.StatusOK, result)
	})

	r.Post("/thesaurus/batch", httpx.HandleBatch(
		func(_ context.Context, req BatchThesaurusRequest) (httpx.BatchResponse[BatchThesaurusItem], error) {
			return httpx.BatchResponse[BatchThesaurusItem]{Results: svc.LookupBatch(req.Words)}, nil
		},
	))
}
