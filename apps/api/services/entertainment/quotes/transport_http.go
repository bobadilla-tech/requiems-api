package quotes

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
)

type BatchRandomRequest struct {
	Count int `json:"count" validate:"required,min=1,max=50"`
}

type BatchRandomResponse = httpx.BatchResponse[Quote]

func RegisterRoutes(r chi.Router, svc *Service) {
	r.Get("/quotes/random", handleRandomQuote(svc))
	r.Post("/quotes/random/batch", handleQuotesBatch(svc))
}

// handleRandomQuote godoc
//
//	@Summary		Get Random Quote
//	@Description	Returns a random inspirational quote with author attribution.
//	@Tags			quotes
//	@Produce		json
//	@Success		200	{object}	httpx.Response[Quote]
//	@Failure		503	{object}	httpx.ErrorResponse
//	@Router			/v1/entertainment/quotes/random [get]
func handleRandomQuote(svc *Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		q, err := svc.Random(r.Context())

		if err != nil {
			httpx.Error(w, http.StatusServiceUnavailable, "service_unavailable", "no quotes available")
			return
		}

		httpx.JSON(w, http.StatusOK, q)
	}
}

// handleQuotesBatch godoc
//
//	@Summary		Get Random Quotes (Batch)
//	@Description	Returns up to 50 random quotes; each quote = 1 unit of usage.
//	@Tags			quotes
//	@Accept			json
//	@Produce		json
//	@Param			request	body		BatchRandomRequest	true	"Number of quotes to return"
//	@Success		200		{object}	httpx.Response[BatchRandomResponse]
//	@Failure		400		{object}	httpx.ErrorResponse
//	@Failure		422		{object}	httpx.ErrorResponse
//	@Router			/v1/entertainment/quotes/random/batch [post]
func handleQuotesBatch(svc *Service) http.HandlerFunc {
	return httpx.HandleBatch(
		func(ctx context.Context, req BatchRandomRequest) (BatchRandomResponse, error) {
			results, err := svc.RandomBatch(ctx, req.Count)
			if err != nil {
				return BatchRandomResponse{}, err
			}
			return BatchRandomResponse{Results: results}, nil
		},
	)
}
