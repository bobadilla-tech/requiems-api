package commodities

import (
	"context"
	"math"
	"strconv"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"requiems-api/platform/svcerr"
)

// HistoricalPrice is a single year's average commodity price.
type HistoricalPrice struct {
	Period string  `json:"period"`
	Price  float64 `json:"price"`
}

// CommodityPrice is the response payload for GET /v1/finance/commodities/:commodity.
type CommodityPrice struct {
	Commodity  string            `json:"commodity"`
	Name       string            `json:"name"`
	Price      float64           `json:"price"`
	Unit       string            `json:"unit"`
	Currency   string            `json:"currency"`
	Change24h  float64           `json:"change_24h"`
	Historical []HistoricalPrice `json:"historical"`
}

const historyDepth = 11 // 1 current + 10 historical years

type yearRow struct {
	year     int16
	price    float64
	name     string
	unit     string
	currency string
}

type slugData struct {
	rows []yearRow
}

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

// BatchCommodityItem is the result for a single item in a batch commodities request.
type BatchCommodityItem struct {
	Slug   string          `json:"slug"`
	Found  bool            `json:"found"`
	Result *CommodityPrice `json:"result,omitempty"`
	Error  string          `json:"error,omitempty"`
}

// GetBatch returns price data for each slug in a single SQL query.
// Slugs not found in the database return Found: false.
func (s *Service) GetBatch(ctx context.Context, slugs []string) []BatchCommodityItem {
	results := make([]BatchCommodityItem, len(slugs))
	// Pre-index slug → result position (last position wins for duplicates).
	slugIndex := make(map[string]int, len(slugs))
	for i, slug := range slugs {
		slugIndex[slug] = i
		results[i] = BatchCommodityItem{Slug: slug, Found: false}
	}

	rows, err := s.db.Query(ctx, `
		SELECT slug, year, price, name, unit, currency
		FROM commodity_price_history
		WHERE slug = ANY($1::text[])
		ORDER BY slug, year DESC
	`, slugs)
	if err != nil {
		return commodityBatchErrorResults(slugs)
	}
	defer rows.Close()

	bySlug := make(map[string]*slugData)

	collectCommodityRows(rows, bySlug)
	if err := rows.Err(); err != nil {
		return commodityBatchErrorResults(slugs)
	}

	for slug, d := range bySlug {
		if item, ok := buildCommodityBatchItem(slug, slugIndex[slug], d); ok {
			results[item.Index] = item.Item
		}
	}

	return results
}

type commodityBatchResult struct {
	Index int
	Item  BatchCommodityItem
}

func commodityBatchErrorResults(slugs []string) []BatchCommodityItem {
	results := make([]BatchCommodityItem, len(slugs))
	for i, slug := range slugs {
		results[i] = BatchCommodityItem{Slug: slug, Error: "failed to fetch commodity data"}
	}
	return results
}

func collectCommodityRows(rows pgx.Rows, bySlug map[string]*slugData) {
	for rows.Next() {
		var slug string
		var r yearRow
		if err := rows.Scan(&slug, &r.year, &r.price, &r.name, &r.unit, &r.currency); err != nil {
			continue
		}
		if d, ok := bySlug[slug]; ok {
			if len(d.rows) < historyDepth {
				d.rows = append(d.rows, r)
			}
			continue
		}
		bySlug[slug] = &slugData{rows: []yearRow{r}}
	}
}

func buildCommodityBatchItem(slug string, idx int, d *slugData) (commodityBatchResult, bool) {
	if idx < 0 || len(d.rows) == 0 {
		return commodityBatchResult{}, false
	}

	latest := d.rows[0]
	var change float64
	if len(d.rows) > 1 && d.rows[1].price != 0 {
		raw := (latest.price - d.rows[1].price) / d.rows[1].price * 100
		change = math.Round(raw*100) / 100
	}
	historical := make([]HistoricalPrice, 0, len(d.rows)-1)
	for _, r := range d.rows[1:] {
		historical = append(historical, HistoricalPrice{
			Period: strconv.Itoa(int(r.year)),
			Price:  r.price,
		})
	}
	cp := CommodityPrice{
		Commodity:  slug,
		Name:       latest.name,
		Price:      latest.price,
		Unit:       latest.unit,
		Currency:   latest.currency,
		Change24h:  change,
		Historical: historical,
	}
	return commodityBatchResult{
		Index: idx,
		Item:  BatchCommodityItem{Slug: slug, Found: true, Result: &cp},
	}, true
}
