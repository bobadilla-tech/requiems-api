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
	r.Get("/thesaurus/{word}", handleThesaurusLookup(svc))
	r.Post("/thesaurus/batch", handleThesaurusBatch(svc))
}

// handleThesaurusLookup godoc
//
//	@Summary		Thesaurus Lookup
//	@Description	Returns synonyms and antonyms for the given word.
//	@Tags			thesaurus
//	@Produce		json
//	@Param			word	path		string	true	"Word to look up"
//	@Success		200		{object}	httpx.Response[Result]
//	@Failure		400		{object}	httpx.ErrorResponse
//	@Failure		404		{object}	httpx.ErrorResponse
//	@Router			/v1/text/thesaurus/{word} [get]
func handleThesaurusLookup(svc *Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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

		httpx.JSON(w, http.StatusOK, result)
	}
}

// handleThesaurusBatch godoc
//
//	@Summary		Thesaurus Lookup (Batch)
//	@Description	Looks up synonyms and antonyms for up to 50 words.
//	@Tags			thesaurus
//	@Accept			json
//	@Produce		json
//	@Param			request	body		BatchThesaurusRequest	true	"List of words to look up"
//	@Success		200		{object}	httpx.Response[httpx.BatchResponse[BatchThesaurusItem]]
//	@Failure		400		{object}	httpx.ErrorResponse
//	@Failure		422		{object}	httpx.ErrorResponse
//	@Router			/v1/text/thesaurus/batch [post]
func handleThesaurusBatch(svc *Service) http.HandlerFunc {
	return httpx.HandleBatch(
		func(_ context.Context, req BatchThesaurusRequest) (httpx.BatchResponse[BatchThesaurusItem], error) {
			return httpx.BatchResponse[BatchThesaurusItem]{Results: svc.LookupBatch(req.Words)}, nil
		},
	)
}
