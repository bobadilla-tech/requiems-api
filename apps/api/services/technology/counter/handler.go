package counter

import (
	"context"
	"net/http"
	"regexp"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
	"requiems-api/platform/middleware"
)

var namespaceRe = regexp.MustCompile(`^[a-zA-Z0-9_-]{1,64}$`)

const namespaceValidationErrorMessage = "invalid namespace: must be 1-64 chars, alphanumeric, hyphen or underscore only"

// BatchCounterRequest is the input for the counter batch endpoint.
type BatchCounterRequest struct {
	Namespaces []string `json:"namespaces" validate:"required,min=1,max=50,dive,required"`
}

func RegisterRoutes(r chi.Router, svc Service) {
	r.Group(func(validated chi.Router) {
		validated.Use(middleware.ValidateURLParam("namespace", namespaceRe, namespaceValidationErrorMessage))

		validated.Post("/counter/{namespace}", handleCounterIncrement(svc))
		validated.Get("/counter/{namespace}", handleCounterGet(svc))
	})

	r.Post("/counter/batch", handleCounterBatch(svc))
}

// handleCounterIncrement godoc
//
//	@Summary		Increment Counter
//	@Description	Atomically increments a counter and returns the new value.
//	@Tags			counter
//	@Produce		json
//	@Param			namespace	path		string	true	"Counter namespace (1–64 chars, alphanumeric, hyphen or underscore)"
//	@Success		200			{object}	httpx.Response[Counter]
//	@Failure		400			{object}	httpx.ErrorResponse
//	@Failure		500			{object}	httpx.ErrorResponse
//	@Router			/v1/technology/counter/{namespace} [post]
func handleCounterIncrement(svc Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ns := chi.URLParam(r, "namespace")

		val, err := svc.Increment(r.Context(), ns)
		if err != nil {
			httpx.Error(w, http.StatusInternalServerError, "internal_error", "Failed to increment counter")
			return
		}

		httpx.JSON(w, http.StatusOK, Counter{Namespace: ns, Value: val})
	}
}

// handleCounterGet godoc
//
//	@Summary		Get Counter Value
//	@Description	Gets the current value of a counter without incrementing (returns 0 if it does not exist).
//	@Tags			counter
//	@Produce		json
//	@Param			namespace	path		string	true	"Counter namespace (1–64 chars, alphanumeric, hyphen or underscore)"
//	@Success		200			{object}	httpx.Response[Counter]
//	@Failure		400			{object}	httpx.ErrorResponse
//	@Failure		500			{object}	httpx.ErrorResponse
//	@Router			/v1/technology/counter/{namespace} [get]
func handleCounterGet(svc Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ns := chi.URLParam(r, "namespace")

		val, err := svc.Get(r.Context(), ns)
		if err != nil {
			httpx.Error(w, http.StatusInternalServerError, "internal_error", "Failed to get counter")
			return
		}

		httpx.JSON(w, http.StatusOK, Counter{Namespace: ns, Value: val})
	}
}

// handleCounterBatch godoc
//
//	@Summary		Batch Counter Operations
//	@Description	Increments up to 50 counters in a single request.
//	@Tags			counter
//	@Accept			json
//	@Produce		json
//	@Param			request	body		BatchCounterRequest	true	"List of counter namespaces"
//	@Success		200		{object}	httpx.Response[httpx.BatchResponse[BatchCounterItem]]
//	@Failure		400		{object}	httpx.ErrorResponse
//	@Failure		422		{object}	httpx.ErrorResponse
//	@Router			/v1/technology/counter/batch [post]
func handleCounterBatch(svc Service) http.HandlerFunc {
	return httpx.HandleBatch(
		func(ctx context.Context, req BatchCounterRequest) (httpx.BatchResponse[BatchCounterItem], error) {
			return httpx.BatchResponse[BatchCounterItem]{Results: svc.IncrementBatch(ctx, req.Namespaces)}, nil
		},
	)
}
