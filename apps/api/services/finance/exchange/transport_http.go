package exchange

import (
	"context"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
	"requiems-api/platform/svcerr"
)

// RateRequest holds the validated query parameters for GET /exchange-rate.
type RateRequest struct {
	From string `query:"from" validate:"required,len=3,alpha"`
	To   string `query:"to"   validate:"required,len=3,alpha"`
}

// ConvertRequest holds the validated query parameters for GET /convert.
type ConvertRequest struct {
	From   string  `query:"from"   validate:"required,len=3,alpha"`
	To     string  `query:"to"     validate:"required,len=3,alpha"`
	Amount float64 `query:"amount" validate:"required,gt=0"`
}

// Fetcher is the interface used by the HTTP transport layer.
type Fetcher interface {
	GetRate(ctx context.Context, from, to string) (rate float64, fetchedAt time.Time, err error)
}

// RegisterRoutes mounts exchange rate handlers on the given router.
// Paths are relative to the parent mount point (/v1/finance).
func RegisterRoutes(r chi.Router, svc *Service) {
	registerExchangeRoutes(r, svc)
}

// registerExchangeRoutes wires the Fetcher interface to the router. Kept
// unexported so tests can inject a stub without going through the concrete
// *Service type.
func registerExchangeRoutes(r chi.Router, f Fetcher) {
	// GET /exchange-rate?from=USD&to=EUR
	r.Get("/exchange-rate", func(w http.ResponseWriter, r *http.Request) {
		var req RateRequest
		if err := httpx.BindQuery(r, &req); err != nil {
			httpx.Error(w, http.StatusBadRequest, "bad_request", err.Error())
			return
		}

		from := strings.ToUpper(req.From)
		to := strings.ToUpper(req.To)

		rate, ts, err := f.GetRate(r.Context(), from, to)
		if err != nil {
			if se, ok := errors.AsType[*svcerr.Error](err); ok {
				httpx.Error(w, svcerr.HTTPStatus(se), se.Code, se.Message)
				return
			}
			httpx.Error(w, http.StatusServiceUnavailable, "upstream_error", "exchange rate service unavailable")
			return
		}

		httpx.JSON(w, http.StatusOK, RateResponse{
			From:      from,
			To:        to,
			Rate:      rate,
			Timestamp: ts.UTC().Format("2006-01-02T15:04:05Z"),
		})
	})

	// GET /convert?from=USD&to=EUR&amount=100
	r.Get("/convert", func(w http.ResponseWriter, r *http.Request) {
		var req ConvertRequest
		if err := httpx.BindQuery(r, &req); err != nil {
			httpx.Error(w, http.StatusBadRequest, "bad_request", err.Error())
			return
		}

		from := strings.ToUpper(req.From)
		to := strings.ToUpper(req.To)

		rate, ts, err := f.GetRate(r.Context(), from, to)
		if err != nil {
			if se, ok := errors.AsType[*svcerr.Error](err); ok {
				httpx.Error(w, svcerr.HTTPStatus(se), se.Code, se.Message)
				return
			}
			httpx.Error(w, http.StatusServiceUnavailable, "upstream_error", "exchange rate service unavailable")
			return
		}

		httpx.JSON(w, http.StatusOK, ConvertResponse{
			From:      from,
			To:        to,
			Rate:      rate,
			Amount:    req.Amount,
			Converted: round2(rate * req.Amount),
			Timestamp: ts.UTC().Format("2006-01-02T15:04:05Z"),
		})
	})
}
