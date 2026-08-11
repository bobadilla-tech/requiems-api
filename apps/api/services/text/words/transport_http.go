package words

import (
	"context"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
	"requiems-api/platform/svcerr"
)

type BatchRequest struct {
	Items []string `json:"items" validate:"required,min=1,max=50,dive,required"`
}

// RegisterRoutes mounts words handlers on the given router.
// Paths are relative to the parent mount point.
func RegisterRoutes(r chi.Router, svc *Service) {
	r.Get("/words/random", handleRandomWord(svc))
	r.Get("/dictionary/{word}", handleDictionaryLookup(svc))
	r.Post("/words/batch", handleWordsBatch(svc))
}

// handleRandomWord godoc
//
//	@Summary		Get Random Word
//	@Description	Returns a random word with its definition and part of speech.
//	@Tags			words
//	@Produce		json
//	@Success		200	{object}	httpx.Response[Word]
//	@Failure		500	{object}	httpx.ErrorResponse
//	@Failure		503	{object}	httpx.ErrorResponse
//	@Router			/v1/text/words/random [get]
func handleRandomWord(svc *Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		wrd, err := svc.Random(r.Context())
		if err != nil {
			if se, ok := errors.AsType[*svcerr.Error](err); ok {
				httpx.Error(w, svcerr.HTTPStatus(se), se.Code, se.Message)
				return
			}
			httpx.Error(w, http.StatusInternalServerError, "internal_error", "internal server error")
			return
		}

		httpx.JSON(w, http.StatusOK, wrd)
	}
}

// handleDictionaryLookup godoc
//
//	@Summary		Dictionary Lookup
//	@Description	Returns definition, phonetics, examples, and synonyms for the given word.
//	@Tags			dictionary
//	@Produce		json
//	@Param			word	path		string	true	"Word to look up"
//	@Success		200		{object}	httpx.Response[DictionaryEntry]
//	@Failure		400		{object}	httpx.ErrorResponse
//	@Failure		404		{object}	httpx.ErrorResponse
//	@Failure		500		{object}	httpx.ErrorResponse
//	@Router			/v1/text/dictionary/{word} [get]
func handleDictionaryLookup(svc *Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		word := chi.URLParam(r, "word")

		entry, err := svc.Define(word)
		if err != nil {
			if se, ok := errors.AsType[*svcerr.Error](err); ok {
				httpx.Error(w, svcerr.HTTPStatus(se), se.Code, se.Message)
				return
			}
			httpx.Error(w, http.StatusInternalServerError, "internal_error", "internal server error")
			return
		}

		httpx.JSON(w, http.StatusOK, entry)
	}
}

// handleWordsBatch godoc
//
//	@Summary		Batch Define Words
//	@Description	Resolve multiple words in a single request; returns entries or per-word errors.
//	@Tags			words
//	@Accept			json
//	@Produce		json
//	@Param			request	body		BatchRequest	true	"List of words to resolve"
//	@Success		200		{object}	httpx.Response[httpx.BatchResponse[BatchItem]]
//	@Failure		422		{object}	httpx.ErrorResponse
//	@Router			/v1/text/words/batch [post]
func handleWordsBatch(svc *Service) http.HandlerFunc {
	return httpx.HandleBatch(
		func(ctx context.Context, req BatchRequest) (httpx.BatchResponse[BatchItem], error) {
			items, err := svc.BatchDefine(ctx, req)

			return httpx.BatchResponse[BatchItem]{Results: items}, err
		},
	)
}
