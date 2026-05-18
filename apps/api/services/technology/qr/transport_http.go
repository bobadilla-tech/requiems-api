package qr

import (
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

func RegisterRoutes(r chi.Router, svc *Service) {
	// GET /qr — returns a raw PNG image.
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
	r.Get("/qr/base64", func(w http.ResponseWriter, r *http.Request) {
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

		httpx.JSON(w, http.StatusOK, Base64Response{
			Image:  base64.StdEncoding.EncodeToString(png),
			Width:  req.Size,
			Height: req.Size,
		})
	})
}
