package numbase

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
)

// ConvertRequest holds the validated query parameters for the base conversion endpoint.
type ConvertRequest struct {
	From  int    `query:"from"  validate:"required,oneof=2 8 10 16"`
	To    int    `query:"to"    validate:"required,oneof=2 8 10 16"`
	Value string `query:"value" validate:"required"`
}

// BatchConvertRequest is the input for the batch base conversion endpoint.
type BatchConvertRequest struct {
	Items []ConvertQuery `json:"items" validate:"required,min=1,max=50,dive"`
}

// RegisterRoutes mounts the base conversion handler on the given router.
// Paths are relative to the parent mount point.
func RegisterRoutes(r chi.Router, svc *Service) {
	r.Get("/base", handleBaseConvert(svc))
	r.Post("/base/batch", handleBaseBatch(svc))
}

// handleBaseConvert godoc
//
//	@Summary		Convert Base
//	@Description	Converts an integer from one number base to another.
//	@Tags			number-base-conversion
//	@Produce		json
//	@Param			from	query		integer	true	"Source base: 2, 8, 10, or 16"
//	@Param			to		query		integer	true	"Target base: 2, 8, 10, or 16"
//	@Param			value	query		string	true	"Number as string (optional 0x, 0b, or 0o prefixes)"
//	@Success		200		{object}	httpx.Response[Result]
//	@Failure		400		{object}	httpx.ErrorResponse
//	@Router			/v1/technology/base [get]
func handleBaseConvert(svc *Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req ConvertRequest
		if err := httpx.BindQuery(r, &req); err != nil {
			httpx.Error(w, http.StatusBadRequest, "bad_request", err.Error())
			return
		}

		result, err := svc.Convert(req.Value, req.From, req.To)
		if err != nil {
			httpx.Error(w, http.StatusBadRequest, "bad_request", err.Error())
			return
		}

		httpx.JSON(w, http.StatusOK, result)
	}
}

// handleBaseBatch godoc
//
//	@Summary		Convert Base (Batch)
//	@Description	Converts up to 50 numbers between bases; per-item errors inline.
//	@Tags			number-base-conversion
//	@Accept			json
//	@Produce		json
//	@Param			request	body		BatchConvertRequest	true	"List of base conversions"
//	@Success		200		{object}	httpx.Response[httpx.BatchResponse[BatchResult]]
//	@Failure		400		{object}	httpx.ErrorResponse
//	@Failure		422		{object}	httpx.ErrorResponse
//	@Router			/v1/technology/base/batch [post]
func handleBaseBatch(svc *Service) http.HandlerFunc {
	return httpx.HandleBatch(
		func(_ context.Context, req BatchConvertRequest) (httpx.BatchResponse[BatchResult], error) {
			return httpx.BatchResponse[BatchResult]{Results: svc.ConvertBatch(req.Items)}, nil
		},
	)
}
