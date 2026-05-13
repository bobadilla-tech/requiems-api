package phone

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
)

func RegisterRoutes(r chi.Router, svc *Service) {
	r.Get("/phone", func(w http.ResponseWriter, r *http.Request) {
		req := ValidateRequest{}

		if err := httpx.BindQuery(r, &req); err != nil {
			httpx.Error(w, http.StatusBadRequest, "bad_request", err.Error())
			return
		}

		httpx.JSON(w, http.StatusOK, svc.Validate(req.Number))
	})

	r.Post("/phone/batch", httpx.HandleBatch(
		func(_ context.Context, req BatchValidateRequest) (httpx.BatchResponse[ValidateResponse], error) {
			return httpx.BatchResponse[ValidateResponse]{Results: svc.ValidateBatch(req.Numbers)}, nil
		},
	))
}
