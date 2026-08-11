package markdown

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
)

// Request is the input for the markdown-to-HTML conversion endpoint.
type Request struct {
	Markdown string `json:"markdown" validate:"required"`
	Sanitize bool   `json:"sanitize"`
}

// BatchRequest is the input for batch markdown-to-HTML conversion.
type BatchRequest struct {
	Markdowns []string `json:"markdowns" validate:"required,min=1,max=50,dive,required"`
	Sanitize  bool     `json:"sanitize"`
}

func RegisterRoutes(r chi.Router, svc *Service) {
	r.Post("/markdown", handleMarkdownConvert(svc))
	r.Post("/markdown/batch", handleMarkdownBatch(svc))
}

// handleMarkdownConvert godoc
//
//	@Summary		Convert Markdown to HTML
//	@Description	Converts a Markdown string to HTML. `sanitize: true` strips unsafe tags like script/iframe.
//	@Tags			markdown
//	@Accept			json
//	@Produce		json
//	@Param			request	body		Request	true	"Markdown text to convert"
//	@Success		200		{object}	httpx.Response[Response]
//	@Failure		422		{object}	httpx.ErrorResponse
//	@Router			/v1/technology/markdown [post]
func handleMarkdownConvert(svc *Service) http.HandlerFunc {
	return httpx.Handle(
		func(_ context.Context, req Request) (Response, error) {
			return svc.Convert(req.Markdown, req.Sanitize)
		},
	)
}

// handleMarkdownBatch godoc
//
//	@Summary		Convert Markdown (Batch)
//	@Description	Converts multiple Markdown strings to HTML; results in input order.
//	@Tags			markdown
//	@Accept			json
//	@Produce		json
//	@Param			request	body		BatchRequest	true	"Markdown texts to convert"
//	@Success		200		{object}	httpx.Response[httpx.BatchResponse[Response]]
//	@Failure		400		{object}	httpx.ErrorResponse
//	@Failure		422		{object}	httpx.ErrorResponse
//	@Failure		500		{object}	httpx.ErrorResponse
//	@Router			/v1/technology/markdown/batch [post]
func handleMarkdownBatch(svc *Service) http.HandlerFunc {
	return httpx.HandleBatch(
		func(_ context.Context, req BatchRequest) (httpx.BatchResponse[Response], error) {
			items, err := svc.ConvertBatch(req.Markdowns, req.Sanitize)
			return httpx.BatchResponse[Response]{Results: items}, err
		},
	)
}
