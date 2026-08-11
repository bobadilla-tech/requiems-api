package base64 //nolint:revive // package name matches the service domain it implements

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
)

// EncodeRequest is the request body for the Base64 encode endpoint.
type EncodeRequest struct {
	Value   string `json:"value"   validate:"required"`
	Variant string `json:"variant" validate:"omitempty,oneof=standard url"`
}

// DecodeRequest is the request body for the Base64 decode endpoint.
type DecodeRequest struct {
	Value   string `json:"value"   validate:"required"`
	Variant string `json:"variant" validate:"omitempty,oneof=standard url"`
}

// RegisterRoutes mounts Base64 encode and decode handlers on the given router.
// Paths are relative to the parent mount point.
func RegisterRoutes(r chi.Router, svc *Service) {
	r.Post("/base64/encode", handleBase64Encode(svc))
	r.Post("/base64/decode", handleBase64Decode(svc))
	r.Post("/base64/encode/batch", handleBase64EncodeBatch(svc))
	r.Post("/base64/decode/batch", handleBase64DecodeBatch(svc))
}

// handleBase64Encode godoc
//
//	@Summary		Encode
//	@Description	Encodes a plain-text string to Base64.
//	@Tags			base64
//	@Accept			json
//	@Produce		json
//	@Param			request	body		EncodeRequest	true	"String to encode"
//	@Success		200		{object}	httpx.Response[Result]
//	@Failure		400		{object}	httpx.ErrorResponse
//	@Failure		422		{object}	httpx.ErrorResponse
//	@Router			/v1/technology/base64/encode [post]
func handleBase64Encode(svc *Service) http.HandlerFunc {
	return httpx.Handle(
		func(_ context.Context, req EncodeRequest) (Result, error) {
			return svc.Encode(req.Value, req.Variant), nil
		},
	)
}

// handleBase64Decode godoc
//
//	@Summary		Decode
//	@Description	Decodes a Base64-encoded string back to plain text.
//	@Tags			base64
//	@Accept			json
//	@Produce		json
//	@Param			request	body		DecodeRequest	true	"Base64 string to decode"
//	@Success		200		{object}	httpx.Response[Result]
//	@Failure		400		{object}	httpx.ErrorResponse
//	@Failure		422		{object}	httpx.ErrorResponse
//	@Router			/v1/technology/base64/decode [post]
func handleBase64Decode(svc *Service) http.HandlerFunc {
	return httpx.Handle(
		func(_ context.Context, req DecodeRequest) (Result, error) {
			return svc.Decode(req.Value, req.Variant)
		},
	)
}

// handleBase64EncodeBatch godoc
//
//	@Summary		Encode (Batch)
//	@Description	Encodes multiple strings to Base64 in a single request.
//	@Tags			base64
//	@Accept			json
//	@Produce		json
//	@Param			request	body		BatchRequest	true	"Values to encode"
//	@Success		200		{object}	httpx.Response[httpx.BatchResponse[Result]]
//	@Failure		422		{object}	httpx.ErrorResponse
//	@Router			/v1/technology/base64/encode/batch [post]
func handleBase64EncodeBatch(svc *Service) http.HandlerFunc {
	return httpx.HandleBatch(
		func(_ context.Context, req BatchRequest) (httpx.BatchResponse[Result], error) {
			return httpx.BatchResponse[Result]{Results: svc.EncodeBatch(req.Values, req.Variant)}, nil
		},
	)
}

// handleBase64DecodeBatch godoc
//
//	@Summary		Decode (Batch)
//	@Description	Decodes multiple Base64 strings. Invalid entries do not fail the request (return empty string, partial success).
//	@Tags			base64
//	@Accept			json
//	@Produce		json
//	@Param			request	body		BatchRequest	true	"Values to decode"
//	@Success		200		{object}	httpx.Response[httpx.BatchResponse[Result]]
//	@Failure		422		{object}	httpx.ErrorResponse
//	@Router			/v1/technology/base64/decode/batch [post]
func handleBase64DecodeBatch(svc *Service) http.HandlerFunc {
	return httpx.HandleBatch(
		func(_ context.Context, req BatchRequest) (httpx.BatchResponse[Result], error) {
			return httpx.BatchResponse[Result]{Results: svc.DecodeBatch(req.Values, req.Variant)}, nil
		},
	)
}
