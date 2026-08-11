package iban

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
)

// Validator is the interface used by the HTTP transport layer.
type Validator interface {
	Parse(ctx context.Context, raw string) (ParseResponse, error)
	ParseBatch(ctx context.Context, numbers []string) []ParseResponse
}

// BatchParseRequest is the body for validating multiple IBANs at once.
type BatchParseRequest struct {
	Numbers []string `json:"numbers" validate:"required,min=1,max=50,dive,required"`
}

// RegisterRoutes mounts IBAN handlers on the given router.
// Paths are relative to the parent mount point (/v1/finance).
func RegisterRoutes(r chi.Router, svc *Service) {
	registerIBANRoutes(r, svc)
}

// registerIBANRoutes wires the Validator interface to the router. Kept
// unexported so tests can inject a stub without going through the concrete
// *Service type.
func registerIBANRoutes(r chi.Router, v Validator) {
	r.Get("/iban/{iban}", handleIBANParse(v))
	r.Post("/iban/batch", handleIBANBatch(v))
}

// handleIBANParse godoc
//
//	@Summary		Validate IBAN
//	@Description	Validates an IBAN and returns country, bank code, and account number. Always returns HTTP 200 — check the `valid` field.
//	@Tags			iban
//	@Produce		json
//	@Param			iban	path		string	true	"IBAN to validate (case-insensitive, spaces stripped)"
//	@Success		200		{object}	httpx.Response[ParseResponse]
//	@Failure		500		{object}	httpx.ErrorResponse
//	@Router			/v1/finance/iban/{iban} [get]
func handleIBANParse(v Validator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		raw := chi.URLParam(r, "iban")

		result, err := v.Parse(r.Context(), raw)
		if err != nil {
			httpx.Error(w, http.StatusInternalServerError, "internal_error", "internal server error")
			return
		}

		httpx.JSON(w, http.StatusOK, result)
	}
}

// handleIBANBatch godoc
//
//	@Summary		Batch Validate IBANs
//	@Description	Validates up to 50 IBANs in a single request; results in input order.
//	@Tags			iban
//	@Accept			json
//	@Produce		json
//	@Param			request	body		BatchParseRequest	true	"List of IBANs to validate"
//	@Success		200		{object}	httpx.Response[httpx.BatchResponse[ParseResponse]]
//	@Failure		422		{object}	httpx.ErrorResponse
//	@Router			/v1/finance/iban/batch [post]
func handleIBANBatch(v Validator) http.HandlerFunc {
	return httpx.HandleBatch(
		func(ctx context.Context, req BatchParseRequest) (httpx.BatchResponse[ParseResponse], error) {
			return httpx.BatchResponse[ParseResponse]{Results: v.ParseBatch(ctx, req.Numbers)}, nil
		},
	)
}
