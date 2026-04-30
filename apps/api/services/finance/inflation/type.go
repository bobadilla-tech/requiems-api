package inflation

import "context"

// HistoricalRate is a single year's inflation rate.
type HistoricalRate struct {
	Period string  `json:"period"`
	Rate   float64 `json:"rate"`
}

// Response is the response payload for GET /v1/finance/inflation.
type Response struct {
	Country    string           `json:"country"`
	Rate       float64          `json:"rate"`
	Period     string           `json:"period"`
	Historical []HistoricalRate `json:"historical"`
}

func (Response) IsData() {}

// Request holds the validated query parameters for the inflation endpoint.
type Request struct {
	Country string `query:"country" validate:"required,iso3166_1_alpha2"`
}

// Getter is the interface used by the HTTP transport layer, allowing transport
// tests to inject a stub without requiring a database connection.
type Getter interface {
	GetInflation(ctx context.Context, countryCode string) (Response, error)
	GetInflationBatch(ctx context.Context, countries []string) BatchResponse
}

// BatchRequest is the body for fetching inflation data for multiple countries at once.
type BatchRequest struct {
	Countries []string `json:"countries" validate:"required,min=1,max=50,dive,iso3166_1_alpha2"`
}

// BatchItem holds the result for a single country in a batch request.
// Found is false when no data exists for that country code.
type BatchItem struct {
	Country    string           `json:"country"`
	Found      bool             `json:"found"`
	Rate       float64          `json:"rate,omitempty"`
	Period     string           `json:"period,omitempty"`
	Historical []HistoricalRate `json:"historical,omitempty"`
}

// BatchResponse is the response payload for POST /v1/finance/inflation/batch.
type BatchResponse struct {
	Results []BatchItem `json:"results"`
	Total   int         `json:"total"`
}

func (BatchResponse) IsData() {}
