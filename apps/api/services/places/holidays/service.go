package holidays

import (
	"errors"

	h "github.com/bobadilla-tech/holidays-per-country"
)

type Holiday struct {
	Date string `json:"date"`
	Name string `json:"name"`
}


type HolidayList struct {
	Country  string    `json:"country"`
	Year     int       `json:"year"`
	Holidays []Holiday `json:"holidays"`
	Total    int       `json:"total"`
}


// BatchQuery holds a single (country, year) pair within a batch request.
type BatchQuery struct {
	Country string `json:"country" validate:"required,iso3166_1_alpha2"`
	Year    int    `json:"year" validate:"required,min=1"`
}

// BatchItem holds the result for a single (country, year) pair in a batch response.
type BatchItem struct {
	Country  string    `json:"country"`
	Year     int       `json:"year"`
	Found    bool      `json:"found"`
	Holidays []Holiday `json:"holidays,omitempty"`
	Total    int       `json:"total,omitempty"`
}

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) GetHolidays(country string, year int) (HolidayList, error) {
	holidays := h.GetHolidays(country, year)
	if len(holidays) == 0 {
		return HolidayList{}, errors.New("no holidays found for the specified country and year")
	}

	holidayList := make([]Holiday, len(holidays))

	for i, holiday := range holidays {
		holidayList[i] = Holiday{
			Name: holiday.Name,
			Date: holiday.Date.Format("2006-01-02"),
		}
	}

	return HolidayList{
		Country:  country,
		Year:     year,
		Holidays: holidayList,
		Total:    len(holidayList),
	}, nil
}

// GetHolidaysBatch returns holidays for each (country, year) pair in the given
// slice. Pairs with no data are returned with Found: false instead of failing
// the entire request (partial failure pattern).
func (s *Service) GetHolidaysBatch(queries []BatchQuery) []BatchItem {
	results := make([]BatchItem, len(queries))

	for i, q := range queries {
		list, err := s.GetHolidays(q.Country, q.Year)
		if err != nil {
			results[i] = BatchItem{
				Country: q.Country,
				Year:    q.Year,
				Found:   false,
			}
			continue
		}

		results[i] = BatchItem{
			Country:  list.Country,
			Year:     list.Year,
			Found:    true,
			Holidays: list.Holidays,
			Total:    list.Total,
		}
	}

	return results
}
