package advice

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
	"requiems-api/platform/svcerr"
)

type BatchRequest struct {
	Count int `json:"count" validate:"required,min=1,max=50"`
}

func RegisterRoutes(r chi.Router, svc *Service) {
	r.Get("/advice", handleAdvice(svc))
	r.Post("/advice/batch", handleAdviceBatch(svc))
}

// handleAdvice    godoc
//
//	@Summary		Get a random piece of advice
//	@Description	Returns a single random piece of advice from the curated collection of wisdom.
//	@Tags			advice
//	@Produce		json
//	@Success		200	{object}	httpx.Response[Advice]
//	@Failure		503	{object}	httpx.ErrorResponse
//	@Router			/v1/entertainment/advice [get]
func handleAdvice(svc *Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		a, err := svc.Random(r.Context())
		if err != nil {
			httpx.Error(w, http.StatusServiceUnavailable, "service_unavailable", "no advice available")
			return
		}

		httpx.JSON(w, http.StatusOK, a)
	}
}

// handleAdviceBatch godoc
//
//	@Summary		Get random advice in bulk
//	@Description	Returns a batch of random pieces of advice from the curated collection.
//	@Tags			advice
//	@Accept			json
//	@Produce		json
//	@Param			request	body		BatchRequest	true	"Batch request"
//	@Success		200		{object}	httpx.Response[httpx.BatchResponse[Advice]]
//	@Failure		400		{object}	httpx.ErrorResponse
//	@Failure		422		{object}	httpx.ErrorResponse
//	@Failure		503		{object}	httpx.ErrorResponse
//	@Router			/v1/entertainment/advice/batch [post]
func handleAdviceBatch(svc *Service) http.HandlerFunc {
	return httpx.HandleBatch(
		func(ctx context.Context, req BatchRequest) (httpx.BatchResponse[Advice], error) {
			results, err := svc.RandomBatch(ctx, req.Count)
			if err != nil {
				return httpx.BatchResponse[Advice]{}, svcerr.Upstream("service_unavailable", "no advice available")
			}

			return httpx.BatchResponse[Advice]{Results: results}, nil
		},
	)
}
