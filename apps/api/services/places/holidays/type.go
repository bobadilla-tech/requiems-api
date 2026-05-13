package holidays

import "requiems-api/platform/httpx"

type Request struct {
	Country string `query:"country" validate:"required,iso3166_1_alpha2"`
	Year    int    `query:"year" validate:"required,min=1"`
}

type Holiday struct {
	Date string `json:"date"`
	Name string `json:"name"`
}

func (Holiday) IsData() {}

type HolidayList struct {
	Country  string    `json:"country"`
	Year     int       `json:"year"`
	Holidays []Holiday `json:"holidays"`
	Total    int       `json:"total"`
}

func (HolidayList) IsData() {}

// BatchQuery holds a single (country, year) pair within a batch request.
type BatchQuery struct {
	Country string `json:"country" validate:"required,iso3166_1_alpha2"`
	Year    int    `json:"year" validate:"required,min=1"`
}

// BatchRequest is the body for POST /holidays/batch.
// Up to 50 (country, year) pairs may be submitted in a single call.
type BatchRequest struct {
	Queries []BatchQuery `json:"queries" validate:"required,min=1,max=50,dive"`
}

// BatchItem holds the result for a single (country, year) pair in a batch response.
// Found is false when no holidays exist for that combination.
type BatchItem struct {
	Country  string    `json:"country"`
	Year     int       `json:"year"`
	Found    bool      `json:"found"`
	Holidays []Holiday `json:"holidays,omitempty"`
	Total    int       `json:"total,omitempty"`
}

// BatchResponse is the response payload for POST /v1/places/holidays/batch.
type BatchResponse = httpx.BatchResponse[BatchItem]
