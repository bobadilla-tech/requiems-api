package mortgage

import "math"

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

// Service computes mortgage payments and amortization schedules.
type Service struct{}

// NewService creates a new Service.
func NewService() *Service { return &Service{} }

// Calculate returns the monthly payment, totals, and full amortization
// schedule for a fixed-rate mortgage.
func (s *Service) Calculate(principal, annualRate float64, years int) Response {
	n := years * 12
	monthlyRate := annualRate / 100.0 / 12.0

	var monthlyPayment float64
	if monthlyRate == 0 {
		monthlyPayment = principal / float64(n)
	} else {
		factor := math.Pow(1+monthlyRate, float64(n))
		monthlyPayment = principal * (monthlyRate * factor) / (factor - 1)
	}

	schedule := make([]ScheduleEntry, n)
	balance := principal

	for i := range n {
		interest := balance * monthlyRate
		principalPaid := monthlyPayment - interest
		balance -= principalPaid
		if balance < 0 {
			balance = 0
		}

		schedule[i] = ScheduleEntry{
			Month:     i + 1,
			Payment:   round2(monthlyPayment),
			Principal: round2(principalPaid),
			Interest:  round2(interest),
			Balance:   round2(balance),
		}
	}

	totalPayment := monthlyPayment * float64(n)

	return Response{
		Principal:      principal,
		Rate:           annualRate,
		Years:          years,
		MonthlyPayment: round2(monthlyPayment),
		TotalPayment:   round2(totalPayment),
		TotalInterest:  round2(totalPayment - principal),
		Schedule:       schedule,
	}
}

func round2(v float64) float64 {
	return math.Round(v*100) / 100
}

// CalculateBatch calculates multiple mortgages and returns the results in the same order as the input.
func (s *Service) CalculateBatch(mortgages []Request) []Response {
	results := make([]Response, len(mortgages))
	for i, m := range mortgages {
		results[i] = s.Calculate(m.Principal, m.Rate, m.Years)
	}
	return results
}
