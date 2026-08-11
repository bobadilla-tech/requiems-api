package chucknorris

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
)

// BatchRequest is the input for the chuck-norris batch endpoint.
type BatchRequest struct {
	Count int `json:"count" validate:"required,min=1,max=50"`
}

// RegisterRoutes mounts Chuck Norris fact handlers on the given router.
// Paths are relative to the parent mount point (/v1/entertainment).
func RegisterRoutes(r chi.Router, svc *Service) {
	r.Get("/chuck-norris", handleChuckNorris(svc))
	r.Post("/chuck-norris/batch", handleChuckNorrisBatch(svc))
}

// handleChuckNorris godoc
//
//	@Summary		Get Random Chuck Norris Fact
//	@Description	Returns a randomly selected Chuck Norris fact.
//	@Tags			chuck-norris
//	@Produce		json
//	@Success		200	{object}	httpx.Response[Fact]
//	@Router			/v1/entertainment/chuck-norris [get]
func handleChuckNorris(svc *Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		httpx.JSON(w, http.StatusOK, svc.Random())
	}
}

// handleChuckNorrisBatch godoc
//
//	@Summary		Get Random Chuck Norris Facts (Batch)
//	@Description	Returns up to 50 random Chuck Norris facts.
//	@Tags			chuck-norris
//	@Accept			json
//	@Produce		json
//	@Param			request	body		BatchRequest	true	"Number of facts to return"
//	@Success		200		{object}	httpx.Response[httpx.BatchResponse[Fact]]
//	@Failure		400		{object}	httpx.ErrorResponse
//	@Failure		422		{object}	httpx.ErrorResponse
//	@Router			/v1/entertainment/chuck-norris/batch [post]
func handleChuckNorrisBatch(svc *Service) http.HandlerFunc {
	return httpx.HandleBatch(
		func(_ context.Context, req BatchRequest) (httpx.BatchResponse[Fact], error) {
			return httpx.BatchResponse[Fact]{Results: svc.RandomBatch(req.Count)}, nil
		},
	)
}
