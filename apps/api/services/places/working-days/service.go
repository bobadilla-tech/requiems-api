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

func (WorkingDays) IsData() {}

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
	return businessdayscalculator.CountBusinessDaysWithHolidays(from, to, opts)
}
