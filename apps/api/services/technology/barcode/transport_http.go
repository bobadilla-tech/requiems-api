package barcode

import (
	"context"
	"encoding/base64"
	"net/http"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
)

// Request holds the query parameters for the barcode endpoints.
type Request struct {
	Data string `query:"data" validate:"required"`
	Type string `query:"type" validate:"required,oneof=code128 code93 code39 ean8 ean13"`
}

func RegisterRoutes(r chi.Router, svc *Service) {
	r.Get("/barcode", handleBarcodePNG(svc))
	r.Get("/barcode/base64", handleBarcodeBase64(svc))
	r.Post("/barcode/batch", handleBarcodeBatch(svc))
}

// handleBarcodePNG godoc
//
//	@Summary		Generate Barcode (PNG)
//	@Description	Returns a raw PNG image of the barcode (Content-Type image/png).
//	@Tags			barcode
//	@Produce		image/png
//	@Param			data	query		string	true	"Text or numeric string to encode"
//	@Param			type	query		string	true	"Barcode format: code128, code93, code39, ean8, ean13"
//	@Success		200		{file}		binary	"Barcode PNG image"
//	@Failure		400		{object}	httpx.ErrorResponse
//	@Failure		422		{object}	httpx.ErrorResponse
//	@Router			/v1/technology/barcode [get]
func handleBarcodePNG(svc *Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req Request

		if err := httpx.BindQuery(r, &req); err != nil {
			httpx.Error(w, http.StatusBadRequest, "bad_request", err.Error())
			return
		}

		png, _, _, err := svc.Generate(req.Data, req.Type)
		if err != nil {
			httpx.Error(w, http.StatusUnprocessableEntity, "unprocessable_entity", err.Error())
			return
		}

		w.Header().Set("Content-Type", "image/png")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(png)
	}
}

// handleBarcodeBase64 godoc
//
//	@Summary		Generate Barcode (Base64 JSON)
//	@Description	Returns JSON with base64-encoded PNG, type, and dimensions.
//	@Tags			barcode
//	@Produce		json
//	@Param			data	query		string	true	"Text or numeric string to encode"
//	@Param			type	query		string	true	"Barcode format: code128, code93, code39, ean8, ean13"
//	@Success		200		{object}	httpx.Response[Base64Response]
//	@Failure		400		{object}	httpx.ErrorResponse
//	@Failure		422		{object}	httpx.ErrorResponse
//	@Router			/v1/technology/barcode/base64 [get]
func handleBarcodeBase64(svc *Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req Request

		if err := httpx.BindQuery(r, &req); err != nil {
			httpx.Error(w, http.StatusBadRequest, "bad_request", err.Error())
			return
		}

		png, width, height, err := svc.Generate(req.Data, req.Type)
		if err != nil {
			httpx.Error(w, http.StatusUnprocessableEntity, "unprocessable_entity", err.Error())
			return
		}

		httpx.JSON(w, http.StatusOK, Base64Response{
			Image:  base64.StdEncoding.EncodeToString(png),
			Type:   req.Type,
			Width:  width,
			Height: height,
		})
	}
}

// handleBarcodeBatch godoc
//
//	@Summary		Generate Barcodes (Batch)
//	@Description	Generates up to 20 barcodes. Invalid items do not fail the request.
//	@Tags			barcode
//	@Accept			json
//	@Produce		json
//	@Param			request	body		BatchRequest	true	"List of barcode options"
//	@Success		200		{object}	httpx.Response[httpx.BatchResponse[BatchResultItem]]
//	@Failure		400		{object}	httpx.ErrorResponse
//	@Failure		422		{object}	httpx.ErrorResponse
//	@Failure		500		{object}	httpx.ErrorResponse
//	@Router			/v1/technology/barcode/batch [post]
func handleBarcodeBatch(svc *Service) http.HandlerFunc {
	return httpx.HandleBatch(
		func(ctx context.Context, req BatchRequest) (httpx.BatchResponse[BatchResultItem], error) {
			results := svc.GenerateBatch(ctx, req.Items)
			return httpx.BatchResponse[BatchResultItem]{Results: results}, nil
		},
	)
}
