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
	Items []QRQuery `json:"items" validate:"required,min=1,max=50,dive"`
}

func RegisterRoutes(r chi.Router, svc *Service) {
	// GET /qr — returns a raw PNG image (not JSON, must stay inline).
	r.Get("/qr", func(w http.ResponseWriter, r *http.Request) {
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
	})

	// GET /qr/base64 — returns a JSON envelope with a base64-encoded PNG.
	r.Get("/qr/base64", httpx.HandleGet(func(ctx context.Context, req Request) (Base64Response, error) {
		png, err := svc.Generate(req.Data, req.Size, req.Recovery)
		if err != nil {
			return Base64Response{}, err
		}
		return Base64Response{
			Image:  base64.StdEncoding.EncodeToString(png),
			Width:  req.Size,
			Height: req.Size,
		}, nil
	}, Request{Size: defaultSize}))

	// POST /qr/base64/batch — returns multiple base64-encoded QR codes.
	// The raw PNG endpoint (GET /qr) has no batch variant: it returns binary image/png.
	r.Post("/qr/base64/batch", httpx.HandleBatch(
		func(_ context.Context, req BatchQRRequest) (httpx.BatchResponse[BatchBase64Item], error) {
			return httpx.BatchResponse[BatchBase64Item]{Results: svc.GenerateBatch(req.Items)}, nil
		},
	))
}
