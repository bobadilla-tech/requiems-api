package units

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
)

// BatchItem represents a single unit conversion operation.
type BatchItem struct {
	From  string   `json:"from" validate:"required"`
	To    string   `json:"to"  validate:"required"`
	Value *float64 `json:"value" validate:"required"`
}

// BatchRequest represents a request containing multiple conversion operations.
type BatchRequest struct {
	Operations []BatchItem `json:"operations" validate:"required,min=1,max=50,dive"`
}

func RegisterRoutes(r chi.Router, svc *Service) {
	r.Get("/convert/units", handleListUnits(svc))
	r.Get("/convert", handleUnitConvert(svc))
	r.Post("/convert/batch", handleUnitsBatch(svc))
}

// handleListUnits godoc
//
//	@Summary		List Available Units
//	@Description	Returns all available unit conversion types grouped by measurement category.
//	@Tags			unit-conversion
//	@Produce		json
//	@Success		200	{object}	httpx.Response[Results]
//	@Router			/v1/technology/convert/units [get]
func handleListUnits(svc *Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		httpx.JSON(w, http.StatusOK, svc.Units())
	}
}

// handleUnitConvert godoc
//
//	@Summary		Convert Units
//	@Description	Converts a value from one unit to another.
//	@Tags			unit-conversion
//	@Produce		json
//	@Param			from	query		string	true	"Source unit key (e.g. miles)"
//	@Param			to		query		string	true	"Target unit key (e.g. km)"
//	@Param			value	query		number	true	"Numeric value to convert"
//	@Success		200		{object}	httpx.Response[Result]
//	@Failure		400		{object}	httpx.ErrorResponse
//	@Router			/v1/technology/convert [get]
func handleUnitConvert(svc *Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		from := r.URL.Query().Get("from")
		to := r.URL.Query().Get("to")
		valueStr := r.URL.Query().Get("value")

		if from == "" || to == "" || valueStr == "" {
			httpx.Error(w, http.StatusBadRequest, "bad_request", "from, to, and value parameters are required")
			return
		}

		value, err := strconv.ParseFloat(valueStr, 64)
		if err != nil {
			httpx.Error(w, http.StatusBadRequest, "bad_request", "value must be a valid number")
			return
		}

		result, err := svc.Convert(from, to, value)
		if err != nil {
			httpx.Error(w, http.StatusBadRequest, "bad_request", err.Error())
			return
		}

		httpx.JSON(w, http.StatusOK, result)
	}
}

// handleUnitsBatch godoc
//
//	@Summary		Convert Units (Batch)
//	@Description	Converts up to 50 unit conversion operations.
//	@Tags			unit-conversion
//	@Accept			json
//	@Produce		json
//	@Param			request	body		BatchRequest	true	"List of unit conversion operations"
//	@Success		200		{object}	httpx.Response[httpx.BatchResponse[BatchResponse]]
//	@Failure		400		{object}	httpx.ErrorResponse
//	@Failure		422		{object}	httpx.ErrorResponse
//	@Router			/v1/technology/convert/batch [post]
func handleUnitsBatch(svc *Service) http.HandlerFunc {
	return httpx.HandleBatch(
		func(ctx context.Context, req BatchRequest) (httpx.BatchResponse[BatchResponse], error) {
			return httpx.BatchResponse[BatchResponse]{Results: svc.ConvertBatch(ctx, req.Operations)}, nil
		},
	)
}
