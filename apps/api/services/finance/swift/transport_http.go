package swift

import (
	"context"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
	"requiems-api/platform/svcerr"
)

// ListFilter describes optional filters and pagination for SWIFT listings.
type ListFilter struct {
	CountryCode string `query:"country_code"`
	BankCode    string `query:"bank_code"`
	Query       string `query:"q"`
	Limit       int    `query:"limit"  validate:"min=1,max=100"`
	Offset      int    `query:"offset" validate:"min=0"`
}

// Looker is the interface used by the HTTP transport layer.
type Looker interface {
	Lookup(ctx context.Context, code string) (LookupResponse, error)
	List(ctx context.Context, filter ListFilter) (ListResponse, error)
}

// RegisterRoutes mounts SWIFT/BIC lookup handlers on the given router.
// Paths are relative to the parent mount point (/v1/finance).
func RegisterRoutes(r chi.Router, svc *Service) {
	registerSWIFTRoutes(r, svc)
}

// registerSWIFTRoutes wires the Looker interface to the router. Kept unexported
// so tests can inject a stub without going through the concrete *Service type.
func registerSWIFTRoutes(r chi.Router, l Looker) {
	r.Get("/swift", handleSwiftList(l))
	r.Get("/swift/{code}", handleSwiftLookup(l))
}

// handleSwiftList godoc
//
//	@Summary		List SWIFT Codes
//	@Description	Lists SWIFT records with optional filters and pagination.
//	@Tags			swift-code
//	@Produce		json
//	@Param			country_code	query		string	false	"2-letter country filter (e.g. DE)"
//	@Param			bank_code		query		string	false	"4-letter bank code filter (e.g. DEUT)"
//	@Param			q				query		string	false	"Text search across swift_code, bank_name, city"
//	@Param			limit			query		integer	false	"Maximum results (default 50, max 200)"
//	@Param			offset			query		integer	false	"Result offset (default 0)"
//	@Success		200				{object}	httpx.Response[ListResponse]
//	@Failure		400				{object}	httpx.ErrorResponse
//	@Failure		500				{object}	httpx.ErrorResponse
//	@Router			/v1/finance/swift [get]
func handleSwiftList(l Looker) http.HandlerFunc {
	return httpx.HandleGet(func(ctx context.Context, req ListFilter) (ListResponse, error) {
		return l.List(ctx, req)
	}, ListFilter{Limit: 50})
}

// handleSwiftLookup godoc
//
//	@Summary		Get SWIFT Code
//	@Description	Looks up bank metadata for a SWIFT/BIC code.
//	@Tags			swift-code
//	@Produce		json
//	@Param			code	path		string	true	"8 or 11 alphanumeric characters"
//	@Success		200		{object}	httpx.Response[LookupResponse]
//	@Failure		400		{object}	httpx.ErrorResponse
//	@Failure		404		{object}	httpx.ErrorResponse
//	@Failure		500		{object}	httpx.ErrorResponse
//	@Router			/v1/finance/swift/{code} [get]
func handleSwiftLookup(l Looker) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rawCode := chi.URLParam(r, "code")

		result, err := l.Lookup(r.Context(), rawCode)
		if err != nil {
			if se, ok := errors.AsType[*svcerr.Error](err); ok {
				httpx.Error(w, svcerr.HTTPStatus(se), se.Code, se.Message)
				return
			}
			httpx.Error(w, http.StatusInternalServerError, "internal_error", "internal server error")
			return
		}

		httpx.JSON(w, http.StatusOK, result)
	}
}
