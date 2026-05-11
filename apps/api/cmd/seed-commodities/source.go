package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// SourceType selects the data provider for a commodity.
type SourceType string

const (
	SourceFRED  SourceType = "fred"
	SourceYahoo SourceType = "yahoo"
)

// CommodityConfig defines one commodity and its data source.
type CommodityConfig struct {
	Slug       string     // primary key slug used in the API (e.g. "gold")
	Name       string     // human-readable name
	Source     SourceType // fred or yahoo
	SeriesID   string     // FRED series ID (if Source == fred)
	Symbol     string     // Yahoo Finance symbol (if Source == yahoo)
	Unit       string     // display unit for the API response
	Currency   string     // always "USD"
	ConvFactor float64    // multiply raw source value by this to get target unit price
}

// CommodityRecord is a single annual average price ready for upsert.
type CommodityRecord struct {
	Slug     string
	Name     string
	Unit     string
	Currency string
	Year     int
	Price    float64
}

// commodities is the authoritative list of supported commodity slugs.
//
// FRED (fredgraph.csv, no API key required):
//   - World Bank monthly series work without licensing restrictions.
//   - ICE/LBMA precious metals series are no longer available via free CSV download.
//
// Yahoo Finance (v8/finance/chart, no API key required):
//   - Used for precious metals (COMEX continuous front-month contracts) and
//     commodities whose World Bank FRED series return 404.
//   - ZC=F (corn) quotes in USX (US cents per bushel) → ConvFactor = 0.01.
var commodities = []CommodityConfig{
	// Precious metals — Yahoo Finance (COMEX/NYMEX continuous front-month)
	{Slug: "gold", Name: "Gold", Source: SourceYahoo, Symbol: "GC=F", Unit: "oz", Currency: "USD", ConvFactor: 1.0},
	{Slug: "silver", Name: "Silver", Source: SourceYahoo, Symbol: "SI=F", Unit: "oz", Currency: "USD", ConvFactor: 1.0},
	{Slug: "platinum", Name: "Platinum", Source: SourceYahoo, Symbol: "PL=F", Unit: "oz", Currency: "USD", ConvFactor: 1.0},
	{Slug: "palladium", Name: "Palladium", Source: SourceYahoo, Symbol: "PA=F", Unit: "oz", Currency: "USD", ConvFactor: 1.0},

	// Energy — FRED (EIA daily series)
	{Slug: "oil", Name: "Crude Oil (WTI)", Source: SourceFRED, SeriesID: "DCOILWTICO", Unit: "barrel", Currency: "USD", ConvFactor: 1.0},
	{Slug: "brent", Name: "Brent Crude", Source: SourceFRED, SeriesID: "DCOILBRENTEU", Unit: "barrel", Currency: "USD", ConvFactor: 1.0},
	{Slug: "natural-gas", Name: "Natural Gas", Source: SourceFRED, SeriesID: "MHHNGSP", Unit: "mmbtu", Currency: "USD", ConvFactor: 1.0},

	// Base metals — FRED (World Bank monthly, USD/metric ton)
	{Slug: "copper", Name: "Copper", Source: SourceFRED, SeriesID: "PCOPPUSDM", Unit: "lb", Currency: "USD", ConvFactor: 1.0 / 2204.623},
	{Slug: "aluminum", Name: "Aluminum", Source: SourceFRED, SeriesID: "PALUMUSDM", Unit: "metric_ton", Currency: "USD", ConvFactor: 1.0},

	// Agricultural — FRED (World Bank) + Yahoo Finance for missing series
	{Slug: "wheat", Name: "Wheat", Source: SourceFRED, SeriesID: "PWHEAMTUSDM", Unit: "metric_ton", Currency: "USD", ConvFactor: 1.0},
	{Slug: "corn", Name: "Corn", Source: SourceYahoo, Symbol: "ZC=F", Unit: "bushel", Currency: "USD", ConvFactor: 0.01}, // ZC=F is USX/bushel
	{Slug: "soybeans", Name: "Soybeans", Source: SourceFRED, SeriesID: "PSOYBUSDM", Unit: "metric_ton", Currency: "USD", ConvFactor: 1.0},
	{Slug: "coffee", Name: "Coffee", Source: SourceFRED, SeriesID: "PCOFFOTMUSDM", Unit: "lb", Currency: "USD", ConvFactor: 1.0 / 2.20462},
	{Slug: "sugar", Name: "Sugar", Source: SourceFRED, SeriesID: "PSUGAISAUSDM", Unit: "lb", Currency: "USD", ConvFactor: 0.01},
	{Slug: "cotton", Name: "Cotton", Source: SourceFRED, SeriesID: "PCOTTINDUSDM", Unit: "lb", Currency: "USD", ConvFactor: 0.01},
	{Slug: "cocoa", Name: "Cocoa", Source: SourceYahoo, Symbol: "CC=F", Unit: "metric_ton", Currency: "USD", ConvFactor: 1.0},
}

// fetchAndAggregate dispatches to the right fetcher based on the commodity source.
func fetchAndAggregate(cfg CommodityConfig) ([]CommodityRecord, error) {
	switch cfg.Source {
	case SourceFRED:
		return fetchFRED(cfg)
	case SourceYahoo:
		return fetchYahoo(cfg)
	default:
		return nil, fmt.Errorf("unknown source: %s", cfg.Source)
	}
}

// ---- FRED ----

const fredBaseURL = "https://fred.stlouisfed.org/graph/fredgraph.csv"

func fetchFRED(cfg CommodityConfig) ([]CommodityRecord, error) {
	url := fredBaseURL + "?id=" + cfg.SeriesID

	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("download: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(io.LimitReader(resp.Body, 512))
		return nil, fmt.Errorf("download: HTTP %d — %s", resp.StatusCode, strings.TrimSpace(string(body)))
	}

	byYear := make(map[int]*yearAcc)

	r := csv.NewReader(resp.Body)
	if _, err = r.Read(); err != nil {
		return nil, fmt.Errorf("read header: %w", err)
	}

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("read csv: %w", err)
		}
		val, year, ok := parseFREDRow(record)
		if !ok {
			continue
		}
		accumulate(byYear, year, val)
	}

	return buildRecords(cfg, byYear), nil
}

// ---- Yahoo Finance ----

const yahooBaseURL = "https://query1.finance.yahoo.com/v8/finance/chart"

type yahooChart struct {
	Chart struct {
		Result []struct {
			Timestamps []int64 `json:"timestamp"`
			Indicators struct {
				Quote []struct {
					Close []any `json:"close"`
				} `json:"quote"`
			} `json:"indicators"`
		} `json:"result"`
	} `json:"chart"`
}

// fetchYahoo downloads 30 years of monthly data from Yahoo Finance, groups by
// year, and returns annual averages.
func fetchYahoo(cfg CommodityConfig) ([]CommodityRecord, error) {
	url := fmt.Sprintf("%s/%s?interval=1mo&range=30y", yahooBaseURL, cfg.Symbol)

	client := &http.Client{Timeout: 30 * time.Second}
	req, err := http.NewRequest(http.MethodGet, url, http.NoBody)
	if err != nil {
		return nil, fmt.Errorf("build request: %w", err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36")

	resp, err := client.Do(req) //nolint:gosec // URL is constructed from a hard-coded constant and internal config, not user input
	if err != nil {
		return nil, fmt.Errorf("download: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(io.LimitReader(resp.Body, 512))
		return nil, fmt.Errorf("download: HTTP %d — %s", resp.StatusCode, strings.TrimSpace(string(body)))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read body: %w", err)
	}

	var chart yahooChart
	if err = json.Unmarshal(body, &chart); err != nil {
		return nil, fmt.Errorf("parse JSON: %w", err)
	}
	if len(chart.Chart.Result) == 0 {
		return nil, fmt.Errorf("no data returned for symbol %s", cfg.Symbol)
	}

	result := chart.Chart.Result[0]
	if len(result.Indicators.Quote) == 0 {
		return nil, fmt.Errorf("no quote data for symbol %s", cfg.Symbol)
	}

	closes := result.Indicators.Quote[0].Close
	timestamps := result.Timestamps

	byYear := make(map[int]*yearAcc)

	for i, ts := range timestamps {
		val, ok := parseYahooClose(closes, i)
		if !ok {
			continue
		}
		year := time.Unix(ts, 0).UTC().Year()
		if year < 1960 || year > time.Now().Year() {
			continue
		}
		accumulate(byYear, year, val)
	}

	return buildRecords(cfg, byYear), nil
}

// ---- shared helpers ----

// parseFREDRow validates and parses a CSV record into (val, year, ok).
func parseFREDRow(record []string) (val float64, year int, ok bool) {
	if len(record) < 2 {
		return 0, 0, false
	}
	valStr := strings.TrimSpace(record[1])
	if valStr == "." || valStr == "" {
		return 0, 0, false
	}
	var err error
	val, err = strconv.ParseFloat(valStr, 64)
	if err != nil || math.IsNaN(val) || math.IsInf(val, 0) || val <= 0 {
		return 0, 0, false
	}
	year, err = parseYear(strings.TrimSpace(record[0]))
	if err != nil || year < 1960 || year > time.Now().Year() {
		return 0, 0, false
	}
	return val, year, true
}

// parseYahooClose extracts a positive finite float64 from closes at index i.
func parseYahooClose(closes []any, i int) (float64, bool) {
	if i >= len(closes) || closes[i] == nil {
		return 0, false
	}
	var val float64
	switch v := closes[i].(type) {
	case float64:
		val = v
	case json.Number:
		val, _ = v.Float64()
	default:
		return 0, false
	}
	if math.IsNaN(val) || math.IsInf(val, 0) || val <= 0 {
		return 0, false
	}
	return val, true
}

type yearAcc struct {
	sum   float64
	count int
}

func accumulate(m map[int]*yearAcc, year int, val float64) {
	if m[year] == nil {
		m[year] = &yearAcc{}
	}
	m[year].sum += val
	m[year].count++
}

func buildRecords(cfg CommodityConfig, byYear map[int]*yearAcc) []CommodityRecord {
	records := make([]CommodityRecord, 0, len(byYear))
	for year, a := range byYear {
		if a.count == 0 {
			continue
		}
		avg := a.sum / float64(a.count)
		price := math.Round(avg*cfg.ConvFactor*10000) / 10000

		records = append(records, CommodityRecord{
			Slug:     cfg.Slug,
			Name:     cfg.Name,
			Unit:     cfg.Unit,
			Currency: cfg.Currency,
			Year:     year,
			Price:    price,
		})
	}
	return records
}

func parseYear(s string) (int, error) {
	if len(s) < 4 {
		return 0, fmt.Errorf("date too short: %q", s)
	}
	return strconv.Atoi(s[:4])
}

func printStats(records []CommodityRecord) {
	bySlug := make(map[string]int)
	for _, r := range records {
		bySlug[r.Slug]++
	}
	fmt.Printf("\n=== Commodity Seed Stats ===\n")
	fmt.Printf("Total records: %d\n", len(records))
	fmt.Printf("Commodities:   %d\n", len(bySlug))
	fmt.Printf("\nYears per commodity:\n")
	for slug, count := range bySlug {
		fmt.Printf("  %-20s %d years\n", slug, count)
	}
}
