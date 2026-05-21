package postal

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

// PostalCode is the response returned for a postal code lookup.
type PostalCode struct { //nolint:revive // established public API type name
	PostalCode string  `json:"postal_code"`
	City       string  `json:"city"`
	State      string  `json:"state"`
	Country    string  `json:"country"`
	Lat        float64 `json:"lat"`
	Lon        float64 `json:"lon"`
}

// PostalQuery is the per-item input for the batch postal code endpoint.
type PostalQuery struct {
	Code    string `json:"code"    validate:"required"`
	Country string `json:"country" validate:"omitempty,len=2"`
}

// BatchResult is the per-item result returned by LookupBatch.
type BatchResult struct {
	Code    string      `json:"code"`
	Country string      `json:"country"`
	Found   bool        `json:"found"`
	Result  *PostalCode `json:"result,omitempty"`
}

// Service looks up postal codes from the GeoNames postal code dataset.
// The dataset is loaded once at startup and held in memory.
type Service struct {
	index map[string]PostalCode // key: "COUNTRY:postalcode"
}

// NewService loads the GeoNames postal code TSV file from dbPath.
// If the file is missing the service starts in degraded mode (all lookups
// return not found) rather than crashing the application.
func NewService(dbPath string) *Service {
	s := &Service{index: make(map[string]PostalCode)}

	f, err := os.Open(dbPath) //nolint:gosec // path comes from application config, not user input
	if err != nil {
		log.Printf("postal: failed to open database %q: %v (all lookups will return not_found)", dbPath, err)
		return s
	}
	defer f.Close()

	// GeoNames postal code TSV columns (tab-separated):
	//  0  country_code
	//  1  postal_code
	//  2  place_name  (city)
	//  3  admin_name1 (state / province)
	//  4  admin_code1
	//  5  admin_name2
	//  6  admin_code2
	//  7  admin_name3
	//  8  admin_code3
	//  9  latitude
	// 10  longitude
	// 11  accuracy
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		cols := strings.Split(line, "\t")
		if len(cols) < 11 {
			continue
		}

		lat, err := strconv.ParseFloat(cols[9], 64)
		if err != nil {
			continue
		}

		lon, err := strconv.ParseFloat(cols[10], 64)
		if err != nil {
			continue
		}

		country := strings.ToUpper(strings.TrimSpace(cols[0]))
		code := strings.TrimSpace(cols[1])
		key := country + ":" + strings.ToUpper(code)

		// Keep the first record for each key (GeoNames lists the primary
		// administrative division first).
		if _, exists := s.index[key]; !exists {
			s.index[key] = PostalCode{
				PostalCode: code,
				City:       cols[2],
				State:      cols[3],
				Country:    country,
				Lat:        lat,
				Lon:        lon,
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Printf("postal: error reading database: %v", err)
	}

	return s
}

// Lookup returns the location for a given postal code and country.
// country should be an ISO 3166-1 alpha-2 code (e.g. "US").
func (s *Service) Lookup(code, country string) (PostalCode, bool) {
	key := strings.ToUpper(strings.TrimSpace(country)) + ":" + strings.ToUpper(strings.TrimSpace(code))
	p, ok := s.index[key]
	return p, ok
}

// LookupBatch looks up each item and returns results in input order.
// Country defaults to "US" when empty. Not-found items have Found: false and no error.
func (s *Service) LookupBatch(items []PostalQuery) []BatchResult {
	results := make([]BatchResult, len(items))
	for i, item := range items {
		country := item.Country
		if country == "" {
			country = "US"
		}
		p, ok := s.Lookup(item.Code, country)
		if ok {
			results[i] = BatchResult{Code: item.Code, Country: strings.ToUpper(country), Found: true, Result: &p}
		} else {
			results[i] = BatchResult{Code: item.Code, Country: strings.ToUpper(country), Found: false}
		}
	}
	return results
}
