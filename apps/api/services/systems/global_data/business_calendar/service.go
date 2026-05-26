package businesscalendar

import (
	"context"
	"sort"
	"time"

	"requiems-api/services/places/holidays"
)

type HolidaysGetter interface {
	GetHolidays(country string, year int) (holidays.HolidayList, error)
}

type WorkingDaysGetter interface {
	GetWorkingDays(from, to time.Time, country, subdivision string) int
}

type Service struct {
	holidays    HolidaysGetter
	workingDays WorkingDaysGetter
}

func NewService(h HolidaysGetter, w WorkingDaysGetter) *Service {
	return &Service{holidays: h, workingDays: w}
}

type Holiday struct {
	Date string `json:"date"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type Result struct {
	CountryCode  string    `json:"country_code"`
	Year         int       `json:"year"`
	Month        *int      `json:"month,omitempty"`
	WorkingDays  int       `json:"working_days"`
	TotalDays    *int      `json:"total_days,omitempty"`
	WeekendDays  *int      `json:"weekend_days,omitempty"`
	Holidays     []Holiday `json:"holidays"`
	HolidayCount int       `json:"holiday_count"`
	NextHoliday  *Holiday  `json:"next_holiday"`
}

func (s *Service) Get(_ context.Context, req Request) (Result, error) {
	year := req.Year
	if year == 0 {
		year = time.Now().Year()
	}

	list, err := s.holidays.GetHolidays(req.Country, year)
	if err != nil {
		list = holidays.HolidayList{Country: req.Country, Year: year, Holidays: nil}
	}

	now := time.Now()
	allHolidays := toHolidays(list.Holidays)

	if req.Month != 0 {
		monthStart := time.Date(year, time.Month(req.Month), 1, 0, 0, 0, 0, time.UTC)
		monthEnd := monthStart.AddDate(0, 1, 0).Add(-time.Second)
		totalDays := int(monthEnd.Sub(monthStart).Hours()/24) + 1
		weekendDays := countWeekendDays(monthStart, monthEnd)

		inMonth := filterByMonth(allHolidays, year, req.Month)
		workDays := s.workingDays.GetWorkingDays(monthStart, monthEnd, req.Country, "")
		nextH := findNextHoliday(allHolidays, now)

		m := req.Month
		total := totalDays
		wknd := weekendDays
		return Result{
			CountryCode:  req.Country,
			Year:         year,
			Month:        &m,
			WorkingDays:  workDays,
			TotalDays:    &total,
			WeekendDays:  &wknd,
			Holidays:     inMonth,
			HolidayCount: len(inMonth),
			NextHoliday:  nextH,
		}, nil
	}

	yearStart := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)
	yearEnd := time.Date(year, 12, 31, 23, 59, 59, 0, time.UTC)
	workDays := s.workingDays.GetWorkingDays(yearStart, yearEnd, req.Country, "")
	nextH := findNextHoliday(allHolidays, now)

	return Result{
		CountryCode:  req.Country,
		Year:         year,
		WorkingDays:  workDays,
		Holidays:     allHolidays,
		HolidayCount: len(allHolidays),
		NextHoliday:  nextH,
	}, nil
}

func toHolidays(src []holidays.Holiday) []Holiday {
	out := make([]Holiday, 0, len(src))
	for _, h := range src {
		out = append(out, Holiday{Date: h.Date, Name: h.Name, Type: "national"})
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Date < out[j].Date })
	return out
}

func filterByMonth(all []Holiday, year, month int) []Holiday {
	prefix := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC).Format("2006-01")
	var out []Holiday
	for _, h := range all {
		if len(h.Date) >= 7 && h.Date[:7] == prefix {
			out = append(out, h)
		}
	}
	if out == nil {
		out = []Holiday{}
	}
	return out
}

func findNextHoliday(all []Holiday, after time.Time) *Holiday {
	cutoff := after.AddDate(0, 0, 90)
	todayStr := after.Format("2006-01-02")
	cutoffStr := cutoff.Format("2006-01-02")
	for i := range all {
		if all[i].Date >= todayStr && all[i].Date <= cutoffStr {
			h := all[i]
			return &h
		}
	}
	return nil
}

func countWeekendDays(from, to time.Time) int {
	count := 0
	for d := from; !d.After(to); d = d.AddDate(0, 0, 1) {
		if d.Weekday() == time.Saturday || d.Weekday() == time.Sunday {
			count++
		}
	}
	return count
}
