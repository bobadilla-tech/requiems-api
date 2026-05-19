package cities

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type City struct {
	Name       string  `json:"name"`
	Country    string  `json:"country"`
	Population int64   `json:"population"`
	Timezone   string  `json:"timezone"`
	Lat        float64 `json:"lat"`
	Lon        float64 `json:"lon"`
}

func (City) IsData() {}

// Service looks up cities from the GeoNames cities15000 dataset.
// The dataset is loaded once at startup and held in memory.
type Service struct {
	// index maps lowercase ASCII city name to the best matching City record.
	// When multiple cities share the same name the one with the highest
	// population is kept.
	index map[string]City
}

// NewService loads the GeoNames cities TSV file from dbPath.
// If the file is missing the service starts in degraded mode (all lookups
// return not found) rather than crashing the application.
func NewService(dbPath string) *Service {
	s := &Service{index: make(map[string]City)}

	f, err := os.Open(dbPath) //nolint:gosec // path comes from application config, not user input
	if err != nil {
		log.Printf("cities: failed to open database %q: %v (all lookups will return not_found)", dbPath, err)
		return s
	}
	defer f.Close()

	// GeoNames cities15000 TSV columns (tab-separated, 19 fields):
	//  0  geonameid
	//  1  name
	//  2  asciiname
	//  3  alternatenames
	//  4  latitude
	//  5  longitude
	//  6  feature_class
	//  7  feature_code
	//  8  country_code
	//  9  cc2
	// 10  admin1_code
	// 11  admin2_code
	// 12  admin3_code
	// 13  admin4_code
	// 14  population
	// 15  elevation
	// 16  dem
	// 17  timezone
	// 18  modification_date
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		cols := strings.Split(line, "\t")
		if len(cols) < 18 {
			continue
		}

		lat, err := strconv.ParseFloat(cols[4], 64)
		if err != nil {
			continue
		}

		lon, err := strconv.ParseFloat(cols[5], 64)
		if err != nil {
			continue
		}

		pop, _ := strconv.ParseInt(cols[14], 10, 64)

		city := City{
			Name:       cols[1],
			Country:    strings.ToUpper(cols[8]),
			Population: pop,
			Timezone:   cols[17],
			Lat:        lat,
			Lon:        lon,
		}

		// Index by lowercase ASCII name. When there are duplicates keep the
		// most populous city so that "london" returns London, UK rather than
		// London, Ontario.
		key := strings.ToLower(cols[2])
		if existing, exists := s.index[key]; !exists || city.Population > existing.Population {
			s.index[key] = city
		}
	}

	if err := scanner.Err(); err != nil {
		log.Printf("cities: error reading database: %v", err)
	}

	return s
}

// Find returns the city record for the given name (case-insensitive).
func (s *Service) Find(name string) (City, bool) {
	key := strings.ToLower(strings.TrimSpace(name))
	city, ok := s.index[key]
	return city, ok
}
