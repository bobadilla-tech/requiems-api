package facts

import (
	"context"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
)

// Request holds optional query parameters for the facts endpoint.
type Request struct {
	Category string `query:"category"`
}

// BatchRequest holds the request payload for the batch facts endpoint.
type BatchRequest struct {
	Categories []string `json:"categories" validate:"required,min=1,max=50,dive"`
}

// RegisterRoutes mounts facts handlers on the given router.
// Paths are relative to the parent mount point (e.g. /v1/entertainment).
func RegisterRoutes(r chi.Router, svc *Service) {
	r.Get("/facts", handleRandomFact(svc))
	r.Post("/facts/batch", handleFactsBatch(svc))
}

// handleRandomFact godoc
//
//	@Summary		Get Random Fact
//	@Description	Returns a randomly selected fact, optionally filtered by category.
//	@Tags			facts
//	@Produce		json
//	@Param			category	query		string	false	"Category filter: science, history, technology, nature, space, food"
//	@Success		200			{object}	httpx.Response[Fact]
//	@Failure		400			{object}	httpx.ErrorResponse
//	@Router			/v1/entertainment/facts [get]
func handleRandomFact(svc *Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
	}
}

// handleFactsBatch godoc
//
//	@Summary		Get Random Facts (Batch)
//	@Description	Returns up to 50 random facts, one per requested category.
//	@Tags			facts
//	@Accept			json
//	@Produce		json
//	@Param			request	body		BatchRequest	true	"List of categories to fetch facts for"
//	@Success		200		{object}	httpx.Response[httpx.BatchResponse[BatchItem]]
//	@Failure		400		{object}	httpx.ErrorResponse
//	@Failure		422		{object}	httpx.ErrorResponse
//	@Router			/v1/entertainment/facts/batch [post]
func handleFactsBatch(svc *Service) http.HandlerFunc {
	return httpx.HandleBatch(
		func(ctx context.Context, req BatchRequest) (httpx.BatchResponse[BatchItem], error) {
			// Normalise to lowercase — consistent with GET /facts?category=
			normalised := make([]string, len(req.Categories))
			for i, c := range req.Categories {
				normalised[i] = strings.ToLower(c)
			}
			return httpx.BatchResponse[BatchItem]{Results: svc.RandomBatch(ctx, normalised)}, nil
		},
	)
}
