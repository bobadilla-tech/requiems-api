package commodities

import (
	"context"
	"math"
	"strconv"

	"github.com/jackc/pgx/v5/pgxpool"

	"requiems-api/platform/svcerr"
)

const historyDepth = 11 // 1 current + 10 historical years

// Service provides commodity price lookups against the commodity_price_history PostgreSQL table.
type Service struct {
	db *pgxpool.Pool
}

// NewService creates a new Service backed by the given connection pool.
func NewService(db *pgxpool.Pool) *Service {
	return &Service{db: db}
}

// Get returns the latest annual average price and historical data for the given slug.
func (s *Service) Get(ctx context.Context, slug string) (CommodityPrice, error) {
	rows, err := s.db.Query(ctx, `
		SELECT year, price, name, unit, currency
		FROM commodity_price_history
		WHERE slug = $1
		ORDER BY year DESC
		LIMIT $2
	`, slug, historyDepth)
	if err != nil {
		return CommodityPrice{}, err
	}
	defer rows.Close()

	type yearRow struct {
		year     int16
		price    float64
		name     string
		unit     string
		currency string
	}

	var results []yearRow
	for rows.Next() {
		var r yearRow
		if err := rows.Scan(&r.year, &r.price, &r.name, &r.unit, &r.currency); err != nil {
			return CommodityPrice{}, err
		}
		results = append(results, r)
	}
	if err := rows.Err(); err != nil {
		return CommodityPrice{}, err
	}

	if len(results) == 0 {
		return CommodityPrice{}, svcerr.NotFound("not_found", "commodity not found")
	}

	latest := results[0]

	var change float64
	if len(results) > 1 && results[1].price != 0 {
		raw := (latest.price - results[1].price) / results[1].price * 100
		change = math.Round(raw*100) / 100
	}

	historical := make([]HistoricalPrice, 0, len(results)-1)
	for _, r := range results[1:] {
		historical = append(historical, HistoricalPrice{
			Period: strconv.Itoa(int(r.year)),
			Price:  r.price,
		})
	}

	return CommodityPrice{
		Commodity:  slug,
		Name:       latest.name,
		Price:      latest.price,
		Unit:       latest.unit,
		Currency:   latest.currency,
		Change24h:  change,
		Historical: historical,
	}, nil
}
