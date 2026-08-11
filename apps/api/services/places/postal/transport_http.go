package postal

import (
	"context"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
)

// BatchPostalRequest is the input for the batch postal code endpoint.
type BatchPostalRequest struct {
	Items []Query `json:"items" validate:"required,min=1,max=50,dive"`
}

func RegisterRoutes(r chi.Router, svc *Service) {
	r.Get("/postal/{code}", handlePostalLookup(svc))
	r.Post("/postal/batch", handlePostalBatch(svc))
}

// handlePostalLookup godoc
//
//	@Summary		Lookup Postal Code
//	@Description	Returns city, state, country, and coordinates for the given postal code.
//	@Tags			postal-code
//	@Produce		json
//	@Param			code	path		string	true	"Postal code"
//	@Param			country	query		string	false	"ISO 3166-1 alpha-2 country code (default US)"
//	@Success		200		{object}	httpx.Response[PostalCode]
//	@Failure		400		{object}	httpx.ErrorResponse
//	@Failure		404		{object}	httpx.ErrorResponse
//	@Failure		500		{object}	httpx.ErrorResponse
//	@Router			/v1/places/postal/{code} [get]
func handlePostalLookup(svc *Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		code := chi.URLParam(r, "code")
		if code == "" {
			httpx.Error(w, http.StatusBadRequest, "bad_request", "postal code is required")
			return
		}

		country := strings.ToUpper(r.URL.Query().Get("country"))
		if country == "" {
			country = "US"
		}

		result, ok := svc.Lookup(code, country)
		if !ok {
			httpx.Error(w, http.StatusNotFound, "not_found", "postal code not found")
			return
		}

		httpx.JSON(w, http.StatusOK, result)
	}
}

// handlePostalBatch godoc
//
//	@Summary		Batch Lookup Postal Codes
//	@Description	Looks up up to 50 postal codes.
//	@Tags			postal-code
//	@Accept			json
//	@Produce		json
//	@Param			request	body		BatchPostalRequest	true	"List of postal code queries"
//	@Success		200		{object}	httpx.Response[httpx.BatchResponse[BatchResult]]
//	@Failure		400		{object}	httpx.ErrorResponse
//	@Failure		422		{object}	httpx.ErrorResponse
//	@Router			/v1/places/postal/batch [post]
func handlePostalBatch(svc *Service) http.HandlerFunc {
	return httpx.HandleBatch(
		func(_ context.Context, req BatchPostalRequest) (httpx.BatchResponse[BatchResult], error) {
			return httpx.BatchResponse[BatchResult]{Results: svc.LookupBatch(req.Items)}, nil
		},
	)
}
