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
	// GET /barcode — returns a raw PNG image.
	r.Get("/barcode", func(w http.ResponseWriter, r *http.Request) {
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
	})

	// GET /barcode/base64 — returns a JSON envelope with a base64-encoded PNG.
	r.Get("/barcode/base64", func(w http.ResponseWriter, r *http.Request) {
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
	})

	// POST /barcode/batch — encodes up to maxBatchSize barcodes in one request.
	// httpx.HandleBatch sets X-Usage-Count = len(items) and wraps the slice
	// in the standard httpx.BatchResponse envelope automatically.
	r.Post("/barcode/batch", httpx.HandleBatch(
		func(ctx context.Context, req BatchRequest) (httpx.BatchResponse[BatchResultItem], error) {
			results := svc.GenerateBatch(ctx, req.Items)
			return httpx.BatchResponse[BatchResultItem]{Results: results}, nil
		},
	))

}
