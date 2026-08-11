package convformat

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
)

// Request is the input for the format conversion endpoint.
type Request struct {
	From    string `json:"from"    validate:"required,oneof=json yaml csv xml toml"`
	To      string `json:"to"      validate:"required,oneof=json yaml csv xml toml"`
	Content string `json:"content" validate:"required"`
}

// BatchFormatRequest is the input for the format conversion batch endpoint.
// Max 20 items — each item content may be up to 512 KB; total body stays ≤ 1 MiB.
type BatchFormatRequest struct {
	Items []Request `json:"items" validate:"required,min=1,max=20,dive"`
}

func RegisterRoutes(r chi.Router, svc *Service) {
	r.Post("/format", handleFormatConvert(svc))
	r.Post("/format/batch", handleFormatBatch(svc))
}

// handleFormatConvert godoc
//
//	@Summary		Convert Format
//	@Description	Converts content from one structured data format to another. Supported formats: json, yaml, csv, xml, toml.
//	@Tags			data-format-conversion
//	@Accept			json
//	@Produce		json
//	@Param			request	body		Request	true	"Format conversion request"
//	@Success		200		{object}	httpx.Response[Response]
//	@Failure		413		{object}	httpx.ErrorResponse
//	@Failure		422		{object}	httpx.ErrorResponse
//	@Router			/v1/technology/format [post]
func handleFormatConvert(svc *Service) http.HandlerFunc {
	return httpx.Handle(
		func(_ context.Context, req Request) (Response, error) {
			return svc.Convert(req)
		},
	)
}

// handleFormatBatch godoc
//
//	@Summary		Convert Format (Batch)
//	@Description	Converts up to 20 structured-data conversions in a single request.
//	@Tags			data-format-conversion
//	@Accept			json
//	@Produce		json
//	@Param			request	body		BatchFormatRequest	true	"List of format conversions"
//	@Success		200		{object}	httpx.Response[httpx.BatchResponse[BatchFormatItem]]
//	@Failure		400		{object}	httpx.ErrorResponse
//	@Failure		422		{object}	httpx.ErrorResponse
//	@Router			/v1/technology/format/batch [post]
func handleFormatBatch(svc *Service) http.HandlerFunc {
	return httpx.HandleBatch(
		func(_ context.Context, req BatchFormatRequest) (httpx.BatchResponse[BatchFormatItem], error) {
			return httpx.BatchResponse[BatchFormatItem]{Results: svc.ConvertBatch(req.Items)}, nil
		},
	)
}
