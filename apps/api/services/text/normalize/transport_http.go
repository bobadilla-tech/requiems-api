package normalize

import (
	"context"

	"requiems-api/platform/httpx"
	"requiems-api/platform/svcerr"

	"github.com/go-chi/chi/v5"
)

type EmailNormalizationRequest struct {
	Email string `json:"email" validate:"required"`
}

type BatchEmailNormalizationRequest struct {
	Emails []string `json:"emails" validate:"required,min=1,max=100,dive,required"`
}

func RegisterRoutes(router chi.Router, svc *Service) {
	router.Post("/normalize/batch", httpx.HandleBatch(func(_ context.Context, req BatchEmailNormalizationRequest) (httpx.BatchResponse[EmailNormalizationBatchItem], error) {
		return httpx.BatchResponse[EmailNormalizationBatchItem]{Results: svc.NormalizeBatch(req.Emails)}, nil
	}))

	router.Post("/normalize", httpx.Handle(func(_ context.Context, req EmailNormalizationRequest) (EmailNormalization, error) {
		res, err := svc.Normalize(req.Email)

		if err != nil {
			return EmailNormalization{}, svcerr.Invalid("bad_request", err.Error())
		}

		return res, nil
	}))
}
