package locationresolve

import (
	"context"
	"sync"
	"time"

	"requiems-api/services/places/geocode"
	"requiems-api/services/places/holidays"
	"requiems-api/services/places/timezone"
)

// -- Dependency interfaces ---------------------------------------------------

type Geocoder interface {
	Geocode(ctx context.Context, address string) (geocode.GeocodeResponse, error)
	ReverseGeocode(ctx context.Context, lat, lon float64) (geocode.ReverseGeocodeResponse, error)
}

type TimezoneGetter interface {
	GetTimezoneByCoords(lat, lon float64) (*timezone.Info, error)
}

type HolidaysGetter interface {
	GetHolidays(country string, year int) (holidays.HolidayList, error)
}

type WorkingDaysGetter interface {
	GetWorkingDays(from, to time.Time, country, subdivision string) int
}

// -- Service -----------------------------------------------------------------

// Service resolves full location context from an address or coordinates.
type Service struct {
	geocoder    Geocoder
	timezone    TimezoneGetter
	holidays    HolidaysGetter
	workingDays WorkingDaysGetter
}

// NewService returns a new location resolve Service.
func NewService(g Geocoder, t TimezoneGetter, h HolidaysGetter, w WorkingDaysGetter) *Service {
	return &Service{geocoder: g, timezone: t, holidays: h, workingDays: w}
}

// Coordinates holds lat/lon.
type Coordinates struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

// Request is the input for POST /location/resolve.
type Request struct {
	Address     string       `json:"address"`
	Coordinates *Coordinates `json:"coordinates"`
	CountryCode string       `json:"country_code"`
}

// NextHoliday is the next upcoming holiday within 90 days.
type NextHoliday struct {
	Date string `json:"date"`
	Name string `json:"name"`
	Type string `json:"type"`
}

// Result is the full location context response.
type Result struct {
	Address              string       `json:"address"`
	City                 string       `json:"city"`
	Country              string       `json:"country"`
	CountryCode          string       `json:"country_code"`
	Coordinates          Coordinates  `json:"coordinates"`
	Timezone             *string      `json:"timezone"`
	UTCOffset            *string      `json:"utc_offset"`
	CurrentTime          *string      `json:"current_time"`
	IsHolidayToday       bool         `json:"is_holiday_today"`
	WorkingDaysThisMonth int          `json:"working_days_this_month"`
	NextHoliday          *NextHoliday `json:"next_holiday"`
	Flags                []string     `json:"flags"`
}

// Resolve resolves the location context from address or coordinates.
func (s *Service) Resolve(ctx context.Context, req Request) (Result, error) {
	var (
		lat, lon    float64
		address     string
		city        string
		country     string
		countryCode string
	)

	// Phase 1: resolve coordinates (sequential).
	if req.Coordinates != nil {
		lat = req.Coordinates.Lat
		lon = req.Coordinates.Lng
		rev, err := s.geocoder.ReverseGeocode(ctx, lat, lon)
		if err == nil {
			address = rev.Address
			city = rev.City
			country = rev.Country
			countryCode = rev.Country // nominatim returns country_code in Country field (uppercase)
			if req.CountryCode != "" {
				countryCode = req.CountryCode
			}
		}
	} else if req.Address != "" {
		res, err := s.geocoder.Geocode(ctx, req.Address)
		if err != nil {
			return Result{}, err
		}
		lat = res.Lat
		lon = res.Lon
		address = res.Address
		city = res.City
		country = res.Country
		countryCode = res.Country
		if req.CountryCode != "" {
			countryCode = req.CountryCode
		}
	} else {
		return Result{}, &missingInputError{}
	}

	// Phase 2: parallel fan-out — timezone + holidays + working-days.
	now := time.Now()
	year := now.Year()
	month := now.Month()
	monthStart := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)
	monthEnd := monthStart.AddDate(0, 1, 0).Add(-time.Second)

	type tzOut struct {
		r   *timezone.Info
		err error
	}
	type holOut struct {
		r   holidays.HolidayList
		err error
	}
	type wdOut struct{ n int }

	tzCh := make(chan tzOut, 1)
	holCh := make(chan holOut, 1)
	wdCh := make(chan wdOut, 1)

	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		defer wg.Done()
		r, err := s.timezone.GetTimezoneByCoords(lat, lon)
		tzCh <- tzOut{r, err}
	}()
	go func() {
		defer wg.Done()
		r, err := s.holidays.GetHolidays(countryCode, year)
		holCh <- holOut{r, err}
	}()
	go func() {
		defer wg.Done()
		n := s.workingDays.GetWorkingDays(monthStart, monthEnd, countryCode, "")
		wdCh <- wdOut{n}
	}()

	wg.Wait()

	tzResult := <-tzCh
	holResult := <-holCh
	wdResult := <-wdCh

	flags := make([]string, 0, 2)
	out := Result{
		Address:     address,
		City:        city,
		Country:     country,
		CountryCode: countryCode,
		Coordinates: Coordinates{Lat: lat, Lng: lon},
		Flags:       flags,
	}

	if tzResult.err != nil {
		flags = append(flags, "timezone_unavailable")
	} else {
		out.Timezone = &tzResult.r.Timezone
		out.UTCOffset = &tzResult.r.Offset
		out.CurrentTime = &tzResult.r.CurrentTime
	}

	if holResult.err != nil {
		flags = append(flags, "calendar_unavailable")
	} else {
		todayStr := now.Format("2006-01-02")
		for _, h := range holResult.r.Holidays {
			if h.Date == todayStr {
				out.IsHolidayToday = true
				break
			}
		}
		out.NextHoliday = findNext(holResult.r.Holidays, now)
	}

	out.WorkingDaysThisMonth = wdResult.n
	out.Flags = flags
	return out, nil
}

func findNext(all []holidays.Holiday, after time.Time) *NextHoliday {
	cutoff := after.AddDate(0, 0, 90)
	todayStr := after.Format("2006-01-02")
	cutoffStr := cutoff.Format("2006-01-02")
	for _, h := range all {
		if h.Date >= todayStr && h.Date <= cutoffStr {
			return &NextHoliday{Date: h.Date, Name: h.Name, Type: "national"}
		}
	}
	return nil
}

type missingInputError struct{}

func (e *missingInputError) Error() string { return "address or coordinates required" }
