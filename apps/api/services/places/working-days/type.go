package workingdays

import "time"

// Request holds the query parameters for the working days endpoint.
// Defaults should be set before calling httpx.BindQuery.
type Request struct {
	From        time.Time `query:"from" validate:"required"`
	To          time.Time `query:"to" validate:"required,gtfield=From"`
	Country     string    `query:"country" validate:"omitempty,iso3166_1_alpha2"`
	Subdivision string    `query:"subdivision" validate:"omitempty,iso3166_2,required_with=Country"`
}

// WorkingDays represents the response for working days calculation
type WorkingDays struct {
	WorkingDays int    `json:"working_days"`
	From        string `json:"from"`
	To          string `json:"to"`
	Country     string `json:"country,omitempty"`
	Subdivision string `json:"subdivision,omitempty"`
}

func (WorkingDays) IsData() {}

type BatchRequest struct {
	Items []Request `json:"items" validate:"required,min=1,max=50,dive"`
}

type BatchResponse struct {
	Results []WorkingDays `json:"results"`
}

func (BatchResponse) IsData() {}
