package color //nolint:revive // package name matches the service domain it implements

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
)

// Request holds the validated query parameters for GET /convert/color.
type Request struct {
	From  string `query:"from"  validate:"required,oneof=hex rgb hsl cmyk"`
	To    string `query:"to"    validate:"required,oneof=hex rgb hsl cmyk"`
	Value string `query:"value" validate:"required"`
}

// ConvertQuery is a single conversion item in a batch request.
type ConvertQuery struct {
	From  string `json:"from"  validate:"required,oneof=hex rgb hsl cmyk"`
	To    string `json:"to"    validate:"required,oneof=hex rgb hsl cmyk"`
	Value string `json:"value" validate:"required"`
}

// BatchColorRequest is the input for the color batch endpoint.
type BatchColorRequest struct {
	Items []ConvertQuery `json:"items" validate:"required,min=1,max=50,dive"`
}

// RegisterRoutes mounts the color conversion handler on the given router.
// Path is relative to the parent mount point.
func RegisterRoutes(r chi.Router, svc *Service) {
	r.Get("/color", handleColorConvert(svc))
	r.Post("/color/batch", handleColorBatch(svc))
}

// handleColorConvert godoc
//
//	@Summary		Convert Color
//	@Description	Converts a color value from one format to another; response always includes all four formats.
//	@Tags			color-conversion
//	@Produce		json
//	@Param			from	query		string	true	"Source format: hex, rgb, hsl, or cmyk"
//	@Param			to		query		string	true	"Target format: hex, rgb, hsl, or cmyk"
//	@Param			value	query		string	true	"Color value in the source format"
//	@Success		200		{object}	httpx.Response[Response]
//	@Failure		400		{object}	httpx.ErrorResponse
//	@Failure		422		{object}	httpx.ErrorResponse
//	@Failure		500		{object}	httpx.ErrorResponse
//	@Router			/v1/technology/color [get]
func handleColorConvert(svc *Service) http.HandlerFunc {
	return httpx.HandleGet(func(ctx context.Context, req Request) (Response, error) {
		return svc.Convert(req.From, req.To, req.Value)
	})
}

// handleColorBatch godoc
//
//	@Summary		Convert Colors (Batch)
//	@Description	Converts up to 50 colors; each response includes all four representations.
//	@Tags			color-conversion
//	@Accept			json
//	@Produce		json
//	@Param			request	body		BatchColorRequest	true	"List of color conversions"
//	@Success		200		{object}	httpx.Response[httpx.BatchResponse[BatchColorItem]]
//	@Failure		400		{object}	httpx.ErrorResponse
//	@Failure		422		{object}	httpx.ErrorResponse
//	@Router			/v1/technology/color/batch [post]
func handleColorBatch(svc *Service) http.HandlerFunc {
	return httpx.HandleBatch(
		func(_ context.Context, req BatchColorRequest) (httpx.BatchResponse[BatchColorItem], error) {
			return httpx.BatchResponse[BatchColorItem]{Results: svc.ConvertBatch(req.Items)}, nil
		},
	)
}
