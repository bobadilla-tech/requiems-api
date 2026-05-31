package workingdays

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/go-chi/chi/v5"

	"requiems-api/platform/httpx"
)

// Request holds the query parameters for the GET endpoint.
// Defaults should be set before calling httpx.BindQuery.

type Request struct {
	From        time.Time `query:"from" validate:"required"`
	To          time.Time `query:"to" validate:"required,gtfield=From"`
	Country     string    `query:"country" validate:"omitempty,iso3166_1_alpha2"`
	Subdivision string    `query:"subdivision" validate:"omitempty,iso3166_2,required_with=Country"`
}
type Date struct {
	time.Time
}

// Date supports JSON dates in:
// - YYYY-MM-DD
// - RFC3339
func (d *Date) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)

	if s == "" || s == "null" {
		return nil
	}

	// Accept YYYY-MM-DD
	t, err := time.Parse(time.DateOnly, s)
	if err != nil {
		// Fallback to RFC3339
		t, err = time.Parse(time.RFC3339, s)
		if err != nil {
			return fmt.Errorf(
				"date must be YYYY-MM-DD or RFC3339, got %q",
				s,
			)
		}
	}

	d.Time = t

	return nil
}

func (d Date) MarshalJSON() ([]byte, error) {
	return []byte(`"` + d.Format(time.DateOnly) + `"`), nil
}

// BatchItem is used only for the batch JSON endpoint.
type BatchItem struct {
	From        Date   `json:"from" validate:"required"`
	To          Date   `json:"to" validate:"required"`
	Country     string `json:"country" validate:"omitempty,iso3166_1_alpha2"`
	Subdivision string `json:"subdivision" validate:"omitempty,iso3166_2,required_with=Country"`
}

type BatchRequest struct {
	Items []BatchItem `json:"items" validate:"required,min=1,max=50,dive"`
}

func RegisterRoutes(r chi.Router, svc *Service) {
	r.Get("/working-days", httpx.HandleGet(func(ctx context.Context, req Request) (WorkingDays, error) {
		return WorkingDays{
			WorkingDays: svc.GetWorkingDays(req.From, req.To, req.Country, req.Subdivision),
			From:        req.From.Format("2006-01-02"),
			To:          req.To.Format("2006-01-02"),
			Country:     req.Country,
			Subdivision: req.Subdivision,
		}, nil
	}))

	r.Post("/working-days/batch", httpx.HandleBatch(
		func(ctx context.Context, req BatchRequest) (httpx.BatchResponse[WorkingDays], error) {
			results := make([]WorkingDays, len(req.Items))
			for i, item := range req.Items {
				results[i] = WorkingDays{
					WorkingDays: svc.GetWorkingDays(item.From.Time, item.To.Time, item.Country, item.Subdivision),
					From:        item.From.Format(time.DateOnly),
					To:          item.To.Format(time.DateOnly),
					Country:     item.Country,
					Subdivision: item.Subdivision,
				}
			}
			return httpx.BatchResponse[WorkingDays]{Results: results}, nil
		},
	))
}
