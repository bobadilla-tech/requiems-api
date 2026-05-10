package exchange

import (
	"context"
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"

	"requiems-api/platform/svcerr"
)

const (
	cacheTTL       = time.Hour
	frankfurterURL = "https://api.frankfurter.app"
	httpTimeout    = 5 * time.Second
)

// Fetcher is the interface used by the HTTP transport layer.
type Fetcher interface {
	GetRate(ctx context.Context, from, to string) (rate float64, fetchedAt time.Time, err error)
}

// Service fetches exchange rates from the Frankfurter API and caches them in Redis.
type Service struct {
	rdb        *redis.Client
	httpClient *http.Client
	baseURL    string
}

// NewService creates a Service backed by the Frankfurter API.
func NewService(rdb *redis.Client) *Service {
	return &Service{
		rdb:        rdb,
		httpClient: &http.Client{Timeout: httpTimeout},
		baseURL:    frankfurterURL,
	}
}

// newServiceWithClient is used in tests to inject a custom HTTP client and base URL.
func newServiceWithClient(rdb *redis.Client, client *http.Client, baseURL string) *Service {
	return &Service{rdb: rdb, httpClient: client, baseURL: baseURL}
}

// GetRate returns the exchange rate from `from` to `to`, using a Redis cache
// with a 1-hour TTL to avoid unnecessary upstream calls.
func (s *Service) GetRate(ctx context.Context, from, to string) (float64, time.Time, error) {
	if s.rdb != nil {
		if rate, ts, ok := s.fromCache(ctx, from, to); ok {
			return rate, ts, nil
		}
	}

	rate, ts, err := s.fetchRate(ctx, from, to)
	if err != nil {
		return 0, time.Time{}, err
	}

	if s.rdb != nil {
		s.toCache(ctx, from, to, rate, ts)
	}

	return rate, ts, nil
}

// fromCache reads a cached rate from Redis. Returns ok=false on any error or miss.
func (s *Service) fromCache(ctx context.Context, from, to string) (float64, time.Time, bool) {
	val, err := s.rdb.Get(ctx, cacheKey(from, to)).Result()
	if err != nil {
		return 0, time.Time{}, false
	}

	rate, ts, err := parseCache(val)
	if err != nil {
		return 0, time.Time{}, false
	}

	return rate, ts, true
}

// toCache stores a rate in Redis. Errors are silently ignored.
func (s *Service) toCache(ctx context.Context, from, to string, rate float64, ts time.Time) {
	s.rdb.Set(ctx, cacheKey(from, to), formatCacheValue(rate, ts), cacheTTL)
}

func formatCacheValue(rate float64, ts time.Time) string {
	return fmt.Sprintf("%s|%s", strconv.FormatFloat(rate, 'f', -1, 64), ts.UTC().Format(time.RFC3339))
}

// frankfurterResponse is the JSON shape returned by api.frankfurter.app/latest.
type frankfurterResponse struct {
	Base  string             `json:"base"`
	Date  string             `json:"date"`
	Rates map[string]float64 `json:"rates"`
}

// fetchRate calls the Frankfurter API and returns the rate and the rate date.
func (s *Service) fetchRate(ctx context.Context, from, to string) (float64, time.Time, error) {
	url := fmt.Sprintf("%s/latest?from=%s&to=%s", s.baseURL, from, to)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, http.NoBody)
	if err != nil {
		return 0, time.Time{}, fmt.Errorf("exchange: build request: %w", err)
	}

	resp, err := s.httpClient.Do(req) //nolint:gosec // URL is built from a fixed base URL and validated 3-char alpha currency codes
	if err != nil {
		return 0, time.Time{}, svcerr.Upstream("upstream_error", "exchange rate service unavailable")
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return 0, time.Time{}, svcerr.Unknown("invalid_currency", fmt.Sprintf("unknown currency code: %s or %s", from, to))
	}

	if resp.StatusCode != http.StatusOK {
		return 0, time.Time{}, svcerr.Upstream("upstream_error", "exchange rate service unavailable")
	}

	var body frankfurterResponse
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		return 0, time.Time{}, svcerr.Upstream("upstream_error", "exchange rate service unavailable")
	}

	rate, ok := body.Rates[to]
	if !ok {
		return 0, time.Time{}, svcerr.Unknown("invalid_currency", fmt.Sprintf("unknown currency code: %s", to))
	}

	ts, err := time.Parse("2006-01-02", body.Date)
	if err != nil {
		ts = time.Now().UTC()
	}

	return rate, ts, nil
}

func cacheKey(from, to string) string {
	return fmt.Sprintf("exchange:%s:%s", strings.ToUpper(from), strings.ToUpper(to))
}

func parseCache(val string) (float64, time.Time, error) {
	parts := strings.SplitN(val, "|", 2)
	if len(parts) != 2 {
		return 0, time.Time{}, fmt.Errorf("invalid cache value")
	}

	rate, err := strconv.ParseFloat(parts[0], 64)
	if err != nil {
		return 0, time.Time{}, err
	}

	ts, err := time.Parse(time.RFC3339, parts[1])
	if err != nil {
		return 0, time.Time{}, err
	}

	return rate, ts, nil
}

// Round rounds f to 2 decimal places.
func round2(f float64) float64 {
	return math.Round(f*100) / 100
}
