package exchange

import (
	"context"
	"errors"
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
	r.Get("/exchange-rate", httpx.HandleGet(func(ctx context.Context, req RateRequest) (RateResponse, error) {
		from := strings.ToUpper(req.From)
		to := strings.ToUpper(req.To)
		rate, ts, err := f.GetRate(ctx, from, to)
		if err != nil {
			if _, ok := errors.AsType[*svcerr.Error](err); !ok {
				return RateResponse{}, svcerr.Upstream("upstream_error", "exchange rate service unavailable")
			}
			return RateResponse{}, err
		}
		return RateResponse{
			From:      from,
			To:        to,
			Rate:      rate,
			Timestamp: ts.UTC().Format("2006-01-02T15:04:05Z"),
		}, nil
	}))

	r.Get("/convert", httpx.HandleGet(func(ctx context.Context, req ConvertRequest) (ConvertResponse, error) {
		from := strings.ToUpper(req.From)
		to := strings.ToUpper(req.To)
		rate, ts, err := f.GetRate(ctx, from, to)
		if err != nil {
			if _, ok := errors.AsType[*svcerr.Error](err); !ok {
				return ConvertResponse{}, svcerr.Upstream("upstream_error", "exchange rate service unavailable")
			}
			return ConvertResponse{}, err
		}
		return ConvertResponse{
			From:      from,
			To:        to,
			Rate:      rate,
			Amount:    req.Amount,
			Converted: round2(rate * req.Amount),
			Timestamp: ts.UTC().Format("2006-01-02T15:04:05Z"),
		}, nil
	}))
}
