package inflation

import (
	"context"
	"strconv"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"

	"requiems-api/platform/svcerr"
)

// HistoricalRate is a single year's inflation rate.
type HistoricalRate struct {
	Period string  `json:"period"`
	Rate   float64 `json:"rate"`
}

// Response is the response payload for GET /v1/finance/inflation.
type Response struct {
	Country    string           `json:"country"`
	Rate       float64          `json:"rate"`
	Period     string           `json:"period"`
	Historical []HistoricalRate `json:"historical"`
}


// BatchItem holds the result for a single country in a batch request.
// Found is false when no data exists for that country code.
type BatchItem struct {
	Country    string           `json:"country"`
	Found      bool             `json:"found"`
	Rate       float64          `json:"rate,omitempty"`
	Period     string           `json:"period,omitempty"`
	Historical []HistoricalRate `json:"historical,omitempty"`
}

const historyDepth = 11 // 1 current + 10 historical years

// Service provides inflation data lookups against the inflation_data PostgreSQL table.
type Service struct {
	db *pgxpool.Pool
}

// NewService creates a new Service backed by the given connection pool.
func NewService(db *pgxpool.Pool) *Service {
	return &Service{db: db}
}

// GetInflation returns the latest inflation rate and historical data for the
// given 2-letter ISO 3166-1 alpha-2 country code.
func (s *Service) GetInflation(ctx context.Context, rawCode string) (Response, error) {
	code := strings.ToUpper(strings.TrimSpace(rawCode))

	rows, err := s.db.Query(ctx, `
		SELECT year, rate
		FROM inflation_data
		WHERE country_code = $1
		ORDER BY year DESC
		LIMIT $2
	`, code, historyDepth)
	if err != nil {
		return Response{}, err
	}
	defer rows.Close()

	type yearRate struct {
		year int16
		rate float64
	}

	var results []yearRate
	for rows.Next() {
		var yr yearRate
		if err := rows.Scan(&yr.year, &yr.rate); err != nil {
			return Response{}, err
		}
		results = append(results, yr)
	}
	if err := rows.Err(); err != nil {
		return Response{}, err
	}

	if len(results) == 0 {
		return Response{}, svcerr.NotFound("not_found", "no inflation data found for country")
	}

	latest := results[0]
	historical := make([]HistoricalRate, 0, len(results)-1)
	for _, r := range results[1:] {
		historical = append(historical, HistoricalRate{
			Period: strconv.Itoa(int(r.year)),
			Rate:   r.rate,
		})
	}

	return Response{
		Country:    code,
		Rate:       latest.rate,
		Period:     strconv.Itoa(int(latest.year)),
		Historical: historical,
	}, nil
}

// GetInflationBatch returns inflation data for multiple countries in the same
// order as the input slice. Countries with no data are returned with Found: false.
func (s *Service) GetInflationBatch(ctx context.Context, countries []string) []BatchItem {
	results := make([]BatchItem, len(countries))

	for i, c := range countries {
		resp, err := s.GetInflation(ctx, c)
		if err != nil {
			// Country not found or unexpected error: return in-band with Found: false.
			results[i] = BatchItem{
				Country: strings.ToUpper(c),
				Found:   false,
			}
			continue
		}

		results[i] = BatchItem{
			Country:    resp.Country,
			Found:      true,
			Rate:       resp.Rate,
			Period:     resp.Period,
			Historical: resp.Historical,
		}
	}

	return results
}
