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
	LookupBatch(ctx context.Context, rawBINs []string) []BatchBINItem
}

// BatchBINRequest is the input for the BIN batch endpoint.
type BatchBINRequest struct {
	BINs []string `json:"bins" validate:"required,min=1,max=50,dive,required"`
}

// RegisterRoutes mounts BIN lookup handlers on the given router.
// Paths are relative to the parent mount point (/v1/finance).
func RegisterRoutes(r chi.Router, svc Looker) {
	registerBINRoutes(r, svc)
}

// registerBINRoutes wires the Service to the router.
func registerBINRoutes(r chi.Router, svc Looker) {
	r.Get("/bin/{bin}", handleBINLookup(svc))
	r.Post("/bin/batch", handleBINBatch(svc))
}

// handleBINLookup godoc
//
//	@Summary		BIN Lookup
//	@Description	Returns card metadata for the given 6–8 digit BIN prefix.
//	@Tags			bin-lookup
//	@Produce		json
//	@Param			bin	path		string	true	"6–8 digit BIN prefix"
//	@Success		200	{object}	httpx.Response[LookupResponse]
//	@Failure		400	{object}	httpx.ErrorResponse
//	@Failure		404	{object}	httpx.ErrorResponse
//	@Failure		500	{object}	httpx.ErrorResponse
//	@Router			/v1/finance/bin/{bin} [get]
func handleBINLookup(svc Looker) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
	}
}

// handleBINBatch godoc
//
//	@Summary		BIN Lookup (Batch)
//	@Description	Looks up card metadata for up to 50 BIN prefixes.
//	@Tags			bin-lookup
//	@Accept			json
//	@Produce		json
//	@Param			request	body		BatchBINRequest	true	"List of BIN prefixes"
//	@Success		200		{object}	httpx.Response[httpx.BatchResponse[BatchBINItem]]
//	@Failure		400		{object}	httpx.ErrorResponse
//	@Failure		422		{object}	httpx.ErrorResponse
//	@Router			/v1/finance/bin/batch [post]
func handleBINBatch(svc Looker) http.HandlerFunc {
	return httpx.HandleBatch(
		func(ctx context.Context, req BatchBINRequest) (httpx.BatchResponse[BatchBINItem], error) {
			return httpx.BatchResponse[BatchBINItem]{Results: svc.LookupBatch(ctx, req.BINs)}, nil
		},
	)
}
