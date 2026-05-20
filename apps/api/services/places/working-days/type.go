package workingdays

import (
	"fmt"
	"strings"
	"time"
)

// Request holds the query parameters for the GET endpoint.
// Defaults should be set before calling httpx.BindQuery.
type Request struct {
	From        time.Time `query:"from" validate:"required"`
	To          time.Time `query:"to" validate:"required,gtfield=From"`
	Country     string    `query:"country" validate:"omitempty,iso3166_1_alpha2"`
	Subdivision string    `query:"subdivision" validate:"omitempty,iso3166_2,required_with=Country"`
}

// Date supports JSON dates in:
// - YYYY-MM-DD
// - RFC3339
type Date struct {
	time.Time
}

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

// WorkingDays represents the response for working days calculation.
type WorkingDays struct {
	WorkingDays int    `json:"working_days"`
	From        string `json:"from"`
	To          string `json:"to"`
	Country     string `json:"country,omitempty"`
	Subdivision string `json:"subdivision,omitempty"`
}

func (WorkingDays) IsData() {}

type BatchRequest struct {
	Items []BatchItem `json:"items" validate:"required,min=1,max=50,dive"`
}

type BatchResponse struct {
	Results []WorkingDays `json:"results"`
}

func (BatchResponse) IsData() {}