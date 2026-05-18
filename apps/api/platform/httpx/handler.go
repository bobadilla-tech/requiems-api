package httpx

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	sentry "github.com/getsentry/sentry-go"

	"requiems-api/platform/svcerr"
)

// Handle wraps an endpoint function with automatic JSON binding, validation,
// and structured error responses. The request body is capped at 1 MB.
//
// Req is decoded from the JSON request body and validated via struct tags.
// Res must implement Data (add an IsData() method to your response type).
//
// Error mapping:
//   - *ValidationFailure  → 422 with {"error":"validation_failed","fields":[...]}
//   - *svcerr.Error       → mapped HTTP status with {"error":Code,"message":Message}
//   - any other error     → 500 with {"error":"internal_error"}
//
// Usage:
//
//	router.Post("/check", httpx.Handle(
//	    func(ctx context.Context, req CheckEmailRequest) (CheckEmailResponse, error) {
//	        return svc.CheckEmail(req.Email), nil
//	    },
//	))
func Handle[Req any, Res Data](
	fn func(ctx context.Context, req Req) (Res, error),
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.Body = http.MaxBytesReader(w, r.Body, 1<<20) // 1 MB limit

		var req Req

		if err := BindAndValidate(r, &req); err != nil {
			if vf, ok := errors.AsType[*ValidationFailure](err); ok {
				ValidationError(w, vf)
				return
			}

			Error(w, http.StatusBadRequest, "bad_request", cleanDecodeError(err))
			return
		}

		res, err := fn(r.Context(), req)

		if err != nil {
			if se, ok := errors.AsType[*svcerr.Error](err); ok {
				if se.Kind == svcerr.KindUpstream {
					sentry.CaptureException(err)
				}

				Error(w, svcerr.HTTPStatus(se), se.Code, se.Message)
				return
			}

			sentry.CaptureException(err)
			Error(w, http.StatusInternalServerError, "internal_error", "internal server error")
			return
		}

		JSON(w, http.StatusOK, res)
	}
}

// HandleBatch is like Handle but for batch endpoints. It reads len(res.Results)
// to set the X-Usage-Count header (used by the auth gateway for per-item billing)
// and auto-populates res.Total before writing the response.
func HandleBatch[Req any, Item any](
	fn func(ctx context.Context, req Req) (BatchResponse[Item], error),
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.Body = http.MaxBytesReader(w, r.Body, 1<<20)

		var req Req
		if err := BindAndValidate(r, &req); err != nil {
			if vf, ok := errors.AsType[*ValidationFailure](err); ok {
				ValidationError(w, vf)
				return
			}
			Error(w, http.StatusBadRequest, "bad_request", cleanDecodeError(err))
			return
		}

		res, err := fn(r.Context(), req)
		if err != nil {
			if se, ok := errors.AsType[*svcerr.Error](err); ok {
				if se.Kind == svcerr.KindUpstream {
					sentry.CaptureException(err)
				}
				Error(w, svcerr.HTTPStatus(se), se.Code, se.Message)
				return
			}
			sentry.CaptureException(err)
			Error(w, http.StatusInternalServerError, "internal_error", "internal server error")
			return
		}

		res.Total = len(res.Results)
		w.Header().Set("X-Usage-Count", strconv.Itoa(len(res.Results)))
		JSON(w, http.StatusOK, res)
	}
}

// Guard returns h unchanged when svc is non-nil. When svc is nil (because the
// service failed to initialize at startup), every request to the wrapped
// handler receives a 500 Internal Server Error response instead of panicking
// or crashing the process.
//
// Usage in a transport_http.go RegisterRoutes function:
//
//	r.Get("/endpoint", httpx.Guard(svc, myHandler))
func Guard[S any](svc *S, h http.HandlerFunc) http.HandlerFunc {
	if svc == nil {
		return func(w http.ResponseWriter, r *http.Request) {
			Error(w, http.StatusInternalServerError, "internal_error", "service unavailable")
		}
	}
	return h
}

// cleanDecodeError returns a safe, human-readable message for JSON decode
// errors, hiding internal implementation details from the client.
func cleanDecodeError(err error) string {
	if err == nil {
		return ""
	}

	if _, ok := errors.AsType[*http.MaxBytesError](err); ok {
		return "request body too large (max 1MB)"
	}

	return "invalid request body"
}
