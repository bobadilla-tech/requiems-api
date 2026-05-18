package workingdays

import (
	"time"

	businessdayscalculator "github.com/bobadilla-tech/business-days-calculator"
)

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

// GetWorkingDaysBatch calculates working days for multiple date ranges.
func (s *Service) GetWorkingDaysBatch(items []Request) []WorkingDays {
	results := make([]WorkingDays, 0, len(items))

	for _, item := range items {
		days := s.GetWorkingDays(item.From, item.To, item.Country, item.Subdivision)

		results = append(results, WorkingDays{
			WorkingDays: days,
			From:        item.From.Format("2006-01-02"),
			To:          item.To.Format("2006-01-02"),
			Country:     item.Country,
			Subdivision: item.Subdivision,
		})
	}

	return results
}
