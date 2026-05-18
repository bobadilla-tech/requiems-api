package timezone

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/ringsaturn/tzf"
)

var locationCache sync.Map

// Service handles timezone lookups.
type Service struct {
	finder tzf.F
}

// NewService creates a new Service with the default timezone finder.
func NewService() (*Service, error) {
	finder, err := tzf.NewDefaultFinder()
	if err != nil {
		return nil, fmt.Errorf("failed to init timezone finder: %w", err)
	}
	return &Service{finder: finder}, nil
}

// GetTimezoneByCoords looks up timezone info for the given latitude/longitude.
// Note: tzf.GetTimezoneName expects (longitude, latitude) — reversed from the
// conventional (lat, lon) order used by this function's signature.
func (s *Service) GetTimezoneByCoords(lat, lon float64) (*Info, error) {
	name := s.finder.GetTimezoneName(lon, lat)
	if name == "" {
		return nil, fmt.Errorf("timezone not found for coordinates %.4f, %.4f", lat, lon)
	}
	return buildTimezoneInfo(name)
}

// GetCurrentTime returns the current time information for the given IANA timezone name.
func (s *Service) GetCurrentTime(tzName string) (*Info, error) {
	return buildTimezoneInfo(tzName)
}

// GetTimezoneByCity looks up timezone info for the given city name.
func (s *Service) GetTimezoneByCity(city string) (*Info, error) {
	key := strings.ToLower(strings.TrimSpace(city))
	name, ok := cityTimezones[key]
	if !ok {
		return nil, fmt.Errorf("city %q not found", city)
	}
	return buildTimezoneInfo(name)
}

// loadLocation returns a cached *time.Location for tzName, loading it on first use.
func loadLocation(tzName string) (*time.Location, error) {
	if v, ok := locationCache.Load(tzName); ok {
		if loc, ok := v.(*time.Location); ok {
			return loc, nil
		}
	}
	loc, err := time.LoadLocation(tzName)
	if err != nil {
		return nil, err
	}
	locationCache.Store(tzName, loc)
	return loc, nil
}

// buildTimezoneInfo loads the IANA timezone and computes the current info.
func buildTimezoneInfo(tzName string) (*Info, error) {
	loc, err := loadLocation(tzName)
	if err != nil {
		return nil, fmt.Errorf("invalid timezone %q: %w", tzName, err)
	}

	now := time.Now().In(loc)
	_, offsetSecs := now.Zone()

	return &Info{
		Timezone:    tzName,
		Offset:      formatOffset(offsetSecs),
		CurrentTime: now.UTC().Format(time.RFC3339),
		IsDST:       isDST(now),
	}, nil
}

// formatOffset formats a UTC offset in seconds as "+HH:MM" or "-HH:MM".
func formatOffset(offsetSecs int) string {
	sign := "+"
	if offsetSecs < 0 {
		sign = "-"
		offsetSecs = -offsetSecs
	}
	hours := offsetSecs / 3600
	mins := (offsetSecs % 3600) / 60
	return fmt.Sprintf("%s%02d:%02d", sign, hours, mins)
}

// isDST reports whether the given time is in daylight saving time.
// It compares the current UTC offset against the minimum offset used by the
// location across a full year (which represents standard time).
func isDST(t time.Time) bool {
	_, offsetNow := t.Zone()

	jan := time.Date(t.Year(), time.January, 1, 12, 0, 0, 0, t.Location())
	jul := time.Date(t.Year(), time.July, 1, 12, 0, 0, 0, t.Location())

	_, offsetJan := jan.Zone()
	_, offsetJul := jul.Zone()

	if offsetJan == offsetJul {
		return false // no DST observed
	}

	// Standard time is the offset that is less positive (more negative) of the two.
	standardOffset := offsetJan
	if offsetJul < offsetJan {
		standardOffset = offsetJul
	}

	return offsetNow != standardOffset
}

func (s *Service) GetTimezoneBatch(
	ctx context.Context,
	cities []string,
) (BatchResponse, error) {
	results := make([]BatchResult, len(cities))

	for i, city := range cities {
		select {
		case <-ctx.Done():
			return BatchResponse{}, ctx.Err()
		default:
		}

		info, err := s.GetTimezoneByCity(city)
		if err != nil {
			results[i] = BatchResult{
				City: city,
			}
			continue
		}

		results[i] = BatchResult{
			City: city,
			Info: info,
		}
	}

	return BatchResponse{
		Results: results,
	}, nil
}
