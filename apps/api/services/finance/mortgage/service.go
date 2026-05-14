package mortgage

import "math"

// Calculator is the interface used by the HTTP transport layer, allowing
// transport tests to inject a stub without requiring the concrete *Service.
type Calculator interface {
	Calculate(principal, annualRate float64, years int) Response
	CalculateBatch(mortgages []Request) BatchResponse
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
func (s *Service) CalculateBatch(mortgages []Request) BatchResponse {
	results := make([]Response, len(mortgages))
	for i, m := range mortgages {
		results[i] = s.Calculate(m.Principal, m.Rate, m.Years)
	}
	return BatchResponse{Results: results, Total: len(results)}
}
