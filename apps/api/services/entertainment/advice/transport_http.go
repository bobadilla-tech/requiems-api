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
	r.Get("/advice", func(w http.ResponseWriter, r *http.Request) {
		a, err := svc.Random(r.Context())
		if err != nil {
			httpx.Error(w, http.StatusServiceUnavailable, "service_unavailable", "no advice available")
			return
		}

		httpx.JSON(w, http.StatusOK, a)
	})
	r.Post("/advice/batch", httpx.HandleBatch(
		func(ctx context.Context, req BatchRequest) (httpx.BatchResponse[Advice], error) {
			results, err := svc.RandomBatch(ctx, req.Count)
			if err != nil {
				return httpx.BatchResponse[Advice]{},
					svcerr.Upstream("service_unavailable", "no advice available")
			}

			return httpx.BatchResponse[Advice]{Results: results}, nil
		},
	))
}
