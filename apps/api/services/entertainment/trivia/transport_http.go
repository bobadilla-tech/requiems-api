package trivia

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
)

// Request holds the optional query parameters for the trivia endpoint.
type Request struct {
	Category   string `query:"category"   validate:"omitempty,oneof=science history geography sports music movies literature math technology nature"`
	Difficulty string `query:"difficulty" validate:"omitempty,oneof=easy medium hard"`
}

// BatchRequest is the request payload for fetching multiple sets
// of trivia questions in a single call.
type BatchRequest struct {
	Filters []Request `json:"filters" validate:"required,min=1,max=50,dive"`
}

// RegisterRoutes mounts trivia handlers on the given router.
// Paths are relative to the parent mount point (e.g. /v1/entertainment).
func RegisterRoutes(r chi.Router, svc *Service) {
	r.Get("/trivia", handleTriviaQuestion(svc))
	r.Post("/trivia/batch", handleTriviaBatch(svc))
}

// handleTriviaQuestion godoc
//
//	@Summary		Get Trivia Question
//	@Description	Returns a random trivia question with multiple-choice answers. Optional category and difficulty filters.
//	@Tags			trivia
//	@Produce		json
//	@Param			category	query		string	false	"Category filter: science, history, geography, sports, music, movies, literature, math, technology, nature"
//	@Param			difficulty	query		string	false	"Difficulty filter: easy, medium, hard"
//	@Success		200			{object}	httpx.Response[Question]
//	@Failure		400			{object}	httpx.ErrorResponse
//	@Failure		404			{object}	httpx.ErrorResponse
//	@Router			/v1/entertainment/trivia [get]
func handleTriviaQuestion(svc *Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req Request
		if err := httpx.BindQuery(r, &req); err != nil {
			httpx.Error(w, http.StatusBadRequest, "bad_request", err.Error())
			return
		}

		q, err := svc.Random(req.Category, req.Difficulty)
		if err != nil {
			httpx.Error(w, http.StatusNotFound, "not_found", err.Error())
			return
		}

		httpx.JSON(w, http.StatusOK, q)
	}
}

// handleTriviaBatch godoc
//
//	@Summary		Get Batch Trivia Questions
//	@Description	Returns up to 50 trivia questions.
//	@Tags			trivia
//	@Accept			json
//	@Produce		json
//	@Param			request	body		BatchRequest	true	"List of category/difficulty filters"
//	@Success		200		{object}	httpx.Response[httpx.BatchResponse[BatchResponse]]
//	@Failure		400		{object}	httpx.ErrorResponse
//	@Failure		422		{object}	httpx.ErrorResponse
//	@Router			/v1/entertainment/trivia/batch [post]
func handleTriviaBatch(svc *Service) http.HandlerFunc {
	return httpx.HandleBatch(
		func(_ context.Context, req BatchRequest) (httpx.BatchResponse[BatchResponse], error) {
			return httpx.BatchResponse[BatchResponse]{Results: svc.RandomBatch(req.Filters)}, nil
		},
	)
}
