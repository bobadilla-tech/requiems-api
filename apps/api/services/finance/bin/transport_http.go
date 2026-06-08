package bin

import (
	"context"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
	"requiems-api/platform/svcerr"
)

// Looker is the interface used by the HTTP transport layer.
type Looker interface {
	Lookup(ctx context.Context, bin string) (LookupResponse, error)
}

// BatchBINRequest is the input for the BIN batch endpoint.
type BatchBINRequest struct {
	BINs []string `json:"bins" validate:"required,min=1,max=50,dive,required"`
}

// RegisterRoutes mounts BIN lookup handlers on the given router.
// Paths are relative to the parent mount point (/v1/finance).
func RegisterRoutes(r chi.Router, svc *Service) {
	registerBINRoutes(r, svc)
}

// registerBINRoutes wires the Service to the router.
func registerBINRoutes(r chi.Router, svc *Service) {
	r.Get("/bin/{bin}", func(w http.ResponseWriter, r *http.Request) {
		rawBIN := chi.URLParam(r, "bin")

		result, err := svc.Lookup(r.Context(), rawBIN)
		if err != nil {
			if se, ok := errors.AsType[*svcerr.Error](err); ok {
				httpx.Error(w, svcerr.HTTPStatus(se), se.Code, se.Message)
				return
			}
			httpx.Error(w, http.StatusInternalServerError, "internal_error", "internal server error")
			return
		}

		httpx.JSON(w, http.StatusOK, result)
	})

	r.Post("/bin/batch", httpx.HandleBatch(
		func(ctx context.Context, req BatchBINRequest) (httpx.BatchResponse[BatchBINItem], error) {
			return httpx.BatchResponse[BatchBINItem]{Results: svc.LookupBatch(ctx, req.BINs)}, nil
		},
	))
}
