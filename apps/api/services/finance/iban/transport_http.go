package iban

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
)

// RegisterRoutes mounts IBAN handlers on the given router.
// Paths are relative to the parent mount point (/v1/finance).
func RegisterRoutes(r chi.Router, svc *Service) {
	registerIBANRoutes(r, svc)
}

// registerIBANRoutes wires the Validator interface to the router. Kept
// unexported so tests can inject a stub without going through the concrete
// *Service type.
func registerIBANRoutes(r chi.Router, v Validator) {
	// GET /iban/{iban} — validate and parse an IBAN
	r.Get("/iban/{iban}", func(w http.ResponseWriter, r *http.Request) {
		raw := chi.URLParam(r, "iban")

		result, err := v.Parse(r.Context(), raw)
		if err != nil {
			httpx.Error(w, http.StatusInternalServerError, "internal_error", "internal server error")
			return
		}

		httpx.JSON(w, http.StatusOK, result)
	})

	// POST /iban/batch — validate and parse IBANs
	r.Post("/iban/batch", httpx.HandleBatch(
		func(ctx context.Context, req BatchParseRequest) (BatchParseResponse, int, error) {
			result, err := v.ParseBatch(ctx, req.Numbers)
			if err != nil {
				return BatchParseResponse{}, 0, err
			}
			return result, result.Total, nil
		},
	))
}
