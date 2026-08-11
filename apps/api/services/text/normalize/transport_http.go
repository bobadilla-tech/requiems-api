package normalize

import (
	"context"
	"net/http"

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
	router.Post("/normalize", handleEmailNormalize(svc))
	router.Post("/normalize/batch", handleEmailNormalizeBatch(svc))
}

// handleEmailNormalize godoc
//
//	@Summary		Normalize Email
//	@Description	Normalizes a single email address and returns the canonical form plus list of transformations.
//	@Tags			email-normalize
//	@Accept			json
//	@Produce		json
//	@Param			request	body		EmailNormalizationRequest	true	"Email address to normalize"
//	@Success		200		{object}	httpx.Response[EmailNormalization]
//	@Failure		400		{object}	httpx.ErrorResponse
//	@Failure		422		{object}	httpx.ErrorResponse
//	@Failure		500		{object}	httpx.ErrorResponse
//	@Router			/v1/text/normalize [post]
func handleEmailNormalize(svc *Service) http.HandlerFunc {
	return httpx.Handle(func(_ context.Context, req EmailNormalizationRequest) (EmailNormalization, error) {
		res, err := svc.Normalize(req.Email)

		if err != nil {
			return EmailNormalization{}, svcerr.Invalid("bad_request", err.Error())
		}

		return res, nil
	})
}

// handleEmailNormalizeBatch godoc
//
//	@Summary		Normalize Email (Batch)
//	@Description	Normalizes up to 100 emails; each item includes `valid`. Billed per email.
//	@Tags			email-normalize
//	@Accept			json
//	@Produce		json
//	@Param			request	body		BatchEmailNormalizationRequest	true	"List of emails to normalize"
//	@Success		200		{object}	httpx.Response[httpx.BatchResponse[EmailNormalizationBatchItem]]
//	@Failure		400		{object}	httpx.ErrorResponse
//	@Failure		422		{object}	httpx.ErrorResponse
//	@Router			/v1/text/normalize/batch [post]
func handleEmailNormalizeBatch(svc *Service) http.HandlerFunc {
	return httpx.HandleBatch(func(_ context.Context, req BatchEmailNormalizationRequest) (httpx.BatchResponse[EmailNormalizationBatchItem], error) {
		return httpx.BatchResponse[EmailNormalizationBatchItem]{Results: svc.NormalizeBatch(req.Emails)}, nil
	})
}
