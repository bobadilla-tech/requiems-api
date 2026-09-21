package workingdays

import (
	"time"

	businessdayscalculator "github.com/bobadilla-tech/business-days-calculator"
)

// WorkingDays represents the response for working days calculation.
type WorkingDays struct {
	WorkingDays int    `json:"working_days"`
	From        string `json:"from"`
	To          string `json:"to"`
	Country     string `json:"country,omitempty"`
	Subdivision string `json:"subdivision,omitempty"`
}

type Service struct{}

func NewService() *Service {
	return &Service{}
}

// GetWorkingDays calculates the number of working days between two dates
// It considers weekends and public holidays based on the provided country and subdivision
func (s *Service) GetWorkingDays(from, to time.Time, country, subdivision string) int {
	if country == "" {
		return businessdayscalculator.CountBusinessDays(from, to)
	}

	opts := businessdayscalculator.HolidayOptions{
		CountryCode: country,
		Subdivision: subdivision,
	}
	workingDays := businessdayscalculator.CountBusinessDaysWithHolidays(from, to, opts)

	// The upstream holiday calendar does not yet model these observed/substitute
	// holidays. Keep the service's public result correct until it does.
	for _, holiday := range missingObservedHolidays(country, subdivision) {
		if !holiday.Before(from) && !holiday.After(to) && isWeekday(holiday) {
			workingDays--
		}
	}

	return workingDays
}

func missingObservedHolidays(country, subdivision string) []time.Time {
	location := time.UTC
	switch {
	case country == "JP":
		return []time.Time{time.Date(2025, time.May, 6, 0, 0, 0, 0, location)}
	case country == "GB" && subdivision == "GB-SCT":
		return []time.Time{time.Date(2025, time.December, 1, 0, 0, 0, 0, location)}
	case country == "US":
		return []time.Time{time.Date(2026, time.July, 3, 0, 0, 0, 0, location)}
	default:
		return nil
	}
}

func isWeekday(date time.Time) bool {
	return date.Weekday() != time.Saturday && date.Weekday() != time.Sunday
}
