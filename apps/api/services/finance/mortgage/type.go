package mortgage

import "requiems-api/platform/httpx"

// Request holds the validated query parameters for the mortgage endpoint.
type Request struct {
	Principal float64 `query:"principal" validate:"required,gt=0"`
	Rate      float64 `query:"rate"      validate:"required,gt=0"`
	Years     int     `query:"years"     validate:"required,min=1,max=50"`
}

// ScheduleEntry represents a single month in the amortization schedule.
type ScheduleEntry struct {
	Month     int     `json:"month"`
	Payment   float64 `json:"payment"`
	Principal float64 `json:"principal"`
	Interest  float64 `json:"interest"`
	Balance   float64 `json:"balance"`
}

// Response is the response payload for GET /v1/finance/mortgage.
type Response struct {
	Principal      float64         `json:"principal"`
	Rate           float64         `json:"rate"`
	Years          int             `json:"years"`
	MonthlyPayment float64         `json:"monthly_payment"`
	TotalPayment   float64         `json:"total_payment"`
	TotalInterest  float64         `json:"total_interest"`
	Schedule       []ScheduleEntry `json:"schedule"`
}

// BatchRequest is the body for validating miltiple mortgages at once.
type BatchRequest struct {
	Mortgages []Request `json:"mortgages" validate:"required,min=1,max=50,dive"`
}

// BatchResponse is the response payload for POST /v1/finance/mortgage/batch.
type BatchResponse = httpx.BatchResponse[Response]

func (Response) IsData() {}
