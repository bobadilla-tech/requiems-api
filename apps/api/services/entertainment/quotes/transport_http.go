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
	r.Get("/quotes/random", func(w http.ResponseWriter, r *http.Request) {
		q, err := svc.Random(r.Context())

		if err != nil {
			httpx.Error(w, http.StatusServiceUnavailable, "service_unavailable", "no quotes available")
			return
		}

		httpx.JSON(w, http.StatusOK, q)
	})

	r.Post("/quotes/random/batch", httpx.HandleBatch(
		func(ctx context.Context, req BatchRandomRequest) (BatchRandomResponse, error) {
			results, err := svc.RandomBatch(ctx, req.Count)
			if err != nil {
				return BatchRandomResponse{}, err
			}
			return BatchRandomResponse{Results: results}, nil
		},
	))
}
