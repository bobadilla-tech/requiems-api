package jokes

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
)

// BatchRequest is the input for the dad jokes batch endpoint.
type BatchRequest struct {
	Count int `json:"count" validate:"required,min=1,max=50"`
}

// RegisterRoutes mounts jokes handlers on the given router.
// Paths are relative to the parent mount point (/v1/entertainment).
func RegisterRoutes(r chi.Router, svc *Service) {
	r.Get("/jokes/dad", handleDadJoke(svc))
	r.Post("/jokes/dad/batch", handleDadJokeBatch(svc))
}

// handleDadJoke godoc
//
//	@Summary		Get Random Dad Joke
//	@Description	Returns a randomly selected dad joke from the collection.
//	@Tags			dad-jokes
//	@Produce		json
//	@Success		200	{object}	httpx.Response[DadJoke]
//	@Router			/v1/entertainment/jokes/dad [get]
func handleDadJoke(svc *Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		httpx.JSON(w, http.StatusOK, svc.Random())
	}
}

// handleDadJokeBatch godoc
//
//	@Summary		Get Random Dad Jokes (Batch)
//	@Description	Returns up to 50 random dad jokes.
//	@Tags			dad-jokes
//	@Accept			json
//	@Produce		json
//	@Param			request	body		BatchRequest	true	"Number of jokes to return"
//	@Success		200		{object}	httpx.Response[httpx.BatchResponse[DadJoke]]
//	@Failure		400		{object}	httpx.ErrorResponse
//	@Failure		422		{object}	httpx.ErrorResponse
//	@Router			/v1/entertainment/jokes/dad/batch [post]
func handleDadJokeBatch(svc *Service) http.HandlerFunc {
	return httpx.HandleBatch(
		func(_ context.Context, req BatchRequest) (httpx.BatchResponse[DadJoke], error) {
			return httpx.BatchResponse[DadJoke]{Results: svc.RandomBatch(req.Count)}, nil
		},
	)
}
