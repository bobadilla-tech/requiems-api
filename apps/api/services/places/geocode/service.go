package geocode

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"

	"requiems-api/platform/svcerr"
)

// GeocodeResponse is returned for address-to-coordinates lookups.
type GeocodeResponse struct { //nolint:revive // established public API type name
	Address string  `json:"address"`
	City    string  `json:"city"`
	Country string  `json:"country"`
	Lat     float64 `json:"lat"`
	Lon     float64 `json:"lon"`
}

// ReverseGeocodeResponse is returned for coordinates-to-address lookups.
type ReverseGeocodeResponse struct {
	Lat     float64 `json:"lat"`
	Lon     float64 `json:"lon"`
	Address string  `json:"address"`
	City    string  `json:"city"`
	Country string  `json:"country"`
}

const (
	cacheTTL    = 24 * time.Hour
	httpTimeout = 10 * time.Second
	userAgent   = "requiems-api/1.0 (https://requiems.xyz)"
)

// Service performs geocoding and reverse geocoding via the Nominatim API,
// caching results in Redis.
type Service struct {
	baseURL    string
	httpClient *http.Client
	rdb        *redis.Client
}

// NewService creates a Service backed by the Nominatim API.
func NewService(baseURL string, httpClient *http.Client, rdb *redis.Client) *Service {
	if httpClient == nil {
		httpClient = &http.Client{Timeout: httpTimeout}
	}
	return &Service{
		baseURL:    strings.TrimRight(baseURL, "/"),
		httpClient: httpClient,
		rdb:        rdb,
	}
}

// Geocode converts a free-text address into coordinates.
func (s *Service) Geocode(ctx context.Context, address string) (GeocodeResponse, error) {
	cacheKey := "geocode:" + url.QueryEscape(strings.ToLower(strings.TrimSpace(address)))

	if s.rdb != nil {
		if cached, ok := s.fromCache(ctx, cacheKey); ok {
			var r GeocodeResponse
			if err := json.Unmarshal([]byte(cached), &r); err == nil {
				return r, nil
			}
		}
	}

	apiURL := fmt.Sprintf("%s/search?format=json&q=%s&limit=1&addressdetails=1",
		s.baseURL, url.QueryEscape(address))

	body, err := s.doRequest(ctx, apiURL)
	if err != nil {
		return GeocodeResponse{}, err
	}

	var results []nominatimSearchResult
	if err := json.Unmarshal(body, &results); err != nil {
		return GeocodeResponse{}, svcerr.Upstream("upstream_error", "geocoding service unavailable")
	}
	if len(results) == 0 {
		return GeocodeResponse{}, svcerr.NotFound("not_found", "no results found for the given address")
	}

	first := results[0]
	lat, _ := strconv.ParseFloat(first.Lat, 64)
	lon, _ := strconv.ParseFloat(first.Lon, 64)

	resp := GeocodeResponse{
		Address: first.DisplayName,
		City:    first.Address.resolveCity(),
		Country: strings.ToUpper(first.Address.CountryCode),
		Lat:     lat,
		Lon:     lon,
	}

	s.toCache(ctx, cacheKey, resp)
	return resp, nil
}

// ReverseGeocode converts coordinates into a human-readable address.
func (s *Service) ReverseGeocode(ctx context.Context, lat, lon float64) (ReverseGeocodeResponse, error) {
	cacheKey := fmt.Sprintf("revgeocode:%.4f:%.4f", lat, lon)

	if s.rdb != nil {
		if cached, ok := s.fromCache(ctx, cacheKey); ok {
			var r ReverseGeocodeResponse
			if err := json.Unmarshal([]byte(cached), &r); err == nil {
				return r, nil
			}
		}
	}

	apiURL := fmt.Sprintf("%s/reverse?format=json&lat=%f&lon=%f&addressdetails=1",
		s.baseURL, lat, lon)

	body, err := s.doRequest(ctx, apiURL)
	if err != nil {
		return ReverseGeocodeResponse{}, err
	}

	var result nominatimReverseResult
	if err := json.Unmarshal(body, &result); err != nil {
		return ReverseGeocodeResponse{}, svcerr.Upstream("upstream_error", "geocoding service unavailable")
	}
	if result.DisplayName == "" {
		return ReverseGeocodeResponse{}, svcerr.NotFound("not_found", "no results found for the given coordinates")
	}

	resp := ReverseGeocodeResponse{
		Lat:     lat,
		Lon:     lon,
		Address: result.DisplayName,
		City:    result.Address.resolveCity(),
		Country: strings.ToUpper(result.Address.CountryCode),
	}

	s.toCache(ctx, cacheKey, resp)
	return resp, nil
}

func (s *Service) doRequest(ctx context.Context, apiURL string) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, apiURL, http.NoBody)
	if err != nil {
		return nil, fmt.Errorf("geocode: build request: %w", err)
	}
	req.Header.Set("User-Agent", userAgent)

	resp, err := s.httpClient.Do(req) //nolint:gosec // URL is built from a trusted base URL + encoded user input
	if err != nil {
		return nil, svcerr.Upstream("upstream_error", "geocoding service unavailable")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, svcerr.Upstream("upstream_error", "geocoding service unavailable")
	}

	buf, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, svcerr.Upstream("upstream_error", "geocoding service unavailable")
	}

	return buf, nil
}

// ReverseQuery is the per-item input for the reverse geocode batch endpoint.
type ReverseQuery struct {
	Lat float64 `json:"lat" validate:"required,min=-90,max=90"`
	Lon float64 `json:"lon" validate:"required,min=-180,max=180"`
}

// BatchGeocodeItem is the per-item result returned by GeocodeBatch.
type BatchGeocodeItem struct {
	Address string           `json:"address"`
	Result  *GeocodeResponse `json:"result,omitempty"`
	Error   string           `json:"error,omitempty"`
}

// BatchReverseGeocodeItem is the per-item result returned by ReverseGeocodeBatch.
type BatchReverseGeocodeItem struct {
	Lat    float64                 `json:"lat"`
	Lon    float64                 `json:"lon"`
	Result *ReverseGeocodeResponse `json:"result,omitempty"`
	Error  string                  `json:"error,omitempty"`
}

const maxGeoWorkers = 10

// GeocodeBatch geocodes each address concurrently and returns results in input order.
// Per-item errors are absorbed in-band; processing continues for all items.
func (s *Service) GeocodeBatch(ctx context.Context, addresses []string) []BatchGeocodeItem {
	results := make([]BatchGeocodeItem, len(addresses))
	sem := make(chan struct{}, maxGeoWorkers)
	var wg sync.WaitGroup

	for i, addr := range addresses {
		wg.Add(1)
		sem <- struct{}{}
		go func(i int, addr string) {
			defer wg.Done()
			defer func() { <-sem }()

			itemCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
			defer cancel()

			r, err := s.Geocode(itemCtx, addr)
			if err != nil {
				var msg string
				if se, ok := err.(interface{ Error() string }); ok {
					msg = se.Error()
				} else {
					msg = "geocoding failed"
				}
				results[i] = BatchGeocodeItem{Address: addr, Error: msg}
			} else {
				results[i] = BatchGeocodeItem{Address: addr, Result: &r}
			}
		}(i, addr)
	}

	wg.Wait()
	return results
}

// ReverseGeocodeBatch reverse-geocodes each coordinate pair concurrently and returns results in input order.
// Per-item errors are absorbed in-band; processing continues for all items.
func (s *Service) ReverseGeocodeBatch(ctx context.Context, items []ReverseQuery) []BatchReverseGeocodeItem {
	results := make([]BatchReverseGeocodeItem, len(items))
	sem := make(chan struct{}, maxGeoWorkers)
	var wg sync.WaitGroup

	for i, item := range items {
		wg.Add(1)
		sem <- struct{}{}
		go func(i int, item ReverseQuery) {
			defer wg.Done()
			defer func() { <-sem }()

			itemCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
			defer cancel()

			r, err := s.ReverseGeocode(itemCtx, item.Lat, item.Lon)
			if err != nil {
				results[i] = BatchReverseGeocodeItem{Lat: item.Lat, Lon: item.Lon, Error: err.Error()}
			} else {
				results[i] = BatchReverseGeocodeItem{Lat: item.Lat, Lon: item.Lon, Result: &r}
			}
		}(i, item)
	}

	wg.Wait()
	return results
}

func (s *Service) fromCache(ctx context.Context, key string) (string, bool) {
	val, err := s.rdb.Get(ctx, key).Result()
	if err != nil {
		return "", false
	}
	return val, true
}

func (s *Service) toCache(ctx context.Context, key string, v any) {
	if s.rdb == nil {
		return
	}
	b, err := json.Marshal(v)
	if err != nil {
		return
	}
	s.rdb.Set(ctx, key, string(b), cacheTTL)
}

// nominatimAddress is the address detail block in Nominatim responses.
type nominatimAddress struct {
	City        string `json:"city"`
	Town        string `json:"town"`
	Village     string `json:"village"`
	County      string `json:"county"`
	CountryCode string `json:"country_code"`
}

// resolveCity returns the most specific city-level place name available.
func (a nominatimAddress) resolveCity() string {
	if a.City != "" {
		return a.City
	}
	if a.Town != "" {
		return a.Town
	}
	if a.Village != "" {
		return a.Village
	}
	return a.County
}

type nominatimSearchResult struct {
	Lat         string           `json:"lat"`
	Lon         string           `json:"lon"`
	DisplayName string           `json:"display_name"`
	Address     nominatimAddress `json:"address"`
}

type nominatimReverseResult struct {
	DisplayName string           `json:"display_name"`
	Address     nominatimAddress `json:"address"`
}
