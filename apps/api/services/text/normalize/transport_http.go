package normalize

import (
	"context"
	"net/http"

	"requiems-api/platform/httpx"

	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(router chi.Router, svc *Service) {
	router.Post("/normalize/batch", httpx.HandleBatch(func(_ context.Context, req BatchEmailNormalizationRequest) (EmailNormalizationBatchResponse, int, error) {
		return svc.NormalizeBatch(req.Emails), len(req.Emails), nil
	}))

	router.Post("/normalize", httpx.Handle(func(_ context.Context, req EmailNormalizationRequest) (EmailNormalization, error) {
		res, err := svc.Normalize(req.Email)
		
		if err != nil {
			return EmailNormalization{}, &httpx.AppError{Status: http.StatusBadRequest, Code: "bad_request", Message: err.Error()}
		}

		return res, nil
	}))
}
