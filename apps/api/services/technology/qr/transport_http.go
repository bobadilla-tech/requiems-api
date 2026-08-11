package qr

import (
	"context"
	"encoding/base64"
	"net/http"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
)

// Request holds the query parameters for the QR code endpoints.
type Request struct {
	Data     string `query:"data"     validate:"required"`
	Size     int    `query:"size"     validate:"min=50,max=1000"`
	Recovery string `query:"recovery" validate:"omitempty,oneof=low medium high highest"`
}

const defaultSize = 256

// BatchQRRequest is the input for the batch QR base64 endpoint.
type BatchQRRequest struct {
	Items []Query `json:"items" validate:"required,min=1,max=50,dive"`
}

func RegisterRoutes(r chi.Router, svc *Service) {
	r.Get("/qr", handleQRCodePNG(svc))
	r.Get("/qr/base64", handleQRCodeBase64(svc))
	r.Post("/qr/base64/batch", handleQRCodeBatch(svc))
}

// handleQRCodePNG godoc
//
//	@Summary		Generate QR Code (PNG)
//	@Description	Returns a raw PNG image of the QR code (Content-Type image/png).
//	@Tags			qr-code
//	@Produce		image/png
//	@Param			data		query		string	true	"Text or URL to encode"
//	@Param			size		query		integer	false	"Image size in pixels (50–1000, default 256)"
//	@Param			recovery	query		string	false	"Error correction level: low, medium, high, highest (default medium)"
//	@Success		200			{file}		binary	"QR code PNG image"
//	@Failure		400			{object}	httpx.ErrorResponse
//	@Failure		500			{object}	httpx.ErrorResponse
//	@Router			/v1/technology/qr [get]
func handleQRCodePNG(svc *Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := Request{Size: defaultSize}

		if err := httpx.BindQuery(r, &req); err != nil {
			httpx.Error(w, http.StatusBadRequest, "bad_request", err.Error())
			return
		}

		png, err := svc.Generate(req.Data, req.Size, req.Recovery)
		if err != nil {
			httpx.Error(w, http.StatusInternalServerError, "internal_error", "failed to generate QR code")
			return
		}

		w.Header().Set("Content-Type", "image/png")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(png)
	}
}

// handleQRCodeBase64 godoc
//
//	@Summary		Generate QR Code (Base64 JSON)
//	@Description	Returns JSON with base64-encoded PNG and dimensions.
//	@Tags			qr-code
//	@Produce		json
//	@Param			data		query		string	true	"Text or URL to encode"
//	@Param			size		query		integer	false	"Image size in pixels (50–1000, default 256)"
//	@Param			recovery	query		string	false	"Error correction level: low, medium, high, highest (default medium)"
//	@Success		200			{object}	httpx.Response[Base64Response]
//	@Failure		400			{object}	httpx.ErrorResponse
//	@Failure		500			{object}	httpx.ErrorResponse
//	@Router			/v1/technology/qr/base64 [get]
func handleQRCodeBase64(svc *Service) http.HandlerFunc {
	return httpx.HandleGet(func(ctx context.Context, req Request) (Base64Response, error) {
		png, err := svc.Generate(req.Data, req.Size, req.Recovery)
		if err != nil {
			return Base64Response{}, err
		}
		return Base64Response{
			Image:  base64.StdEncoding.EncodeToString(png),
			Width:  req.Size,
			Height: req.Size,
		}, nil
	}, Request{Size: defaultSize})
}

// handleQRCodeBatch godoc
//
//	@Summary		Batch Generate QR Codes (Base64)
//	@Description	Generates up to 50 base64 QR codes. The PNG endpoint has no batch variant.
//	@Tags			qr-code
//	@Accept			json
//	@Produce		json
//	@Param			request	body		BatchQRRequest	true	"List of QR generation options"
//	@Success		200		{object}	httpx.Response[httpx.BatchResponse[BatchBase64Item]]
//	@Failure		400		{object}	httpx.ErrorResponse
//	@Failure		422		{object}	httpx.ErrorResponse
//	@Router			/v1/technology/qr/base64/batch [post]
func handleQRCodeBatch(svc *Service) http.HandlerFunc {
	return httpx.HandleBatch(
		func(_ context.Context, req BatchQRRequest) (httpx.BatchResponse[BatchBase64Item], error) {
			return httpx.BatchResponse[BatchBase64Item]{Results: svc.GenerateBatch(req.Items)}, nil
		},
	)
}
