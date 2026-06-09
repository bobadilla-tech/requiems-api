package cryptocoin

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"

	"requiems-api/platform/svcerr"
)

const (
	cacheTTL     = 5 * time.Minute
	coinGeckoURL = "https://api.coingecko.com/api/v3"
	httpTimeout  = 10 * time.Second
)

// Price is the response payload for GET /v1/finance/crypto/{symbol}.
type Price struct {
	Symbol    string  `json:"symbol"`
	Name      string  `json:"name"`
	PriceUSD  float64 `json:"price_usd"`
	Change24h float64 `json:"change_24h"`
	MarketCap float64 `json:"market_cap"`
	Volume24h float64 `json:"volume_24h"`
}

type coinInfo struct {
	id   string
	name string
}

// coinMap maps uppercase ticker symbols to CoinGecko IDs and display names.
var coinMap = map[string]coinInfo{
	"BTC":   {id: "bitcoin", name: "Bitcoin"},
	"ETH":   {id: "ethereum", name: "Ethereum"},
	"BNB":   {id: "binancecoin", name: "BNB"},
	"XRP":   {id: "ripple", name: "XRP"},
	"ADA":   {id: "cardano", name: "Cardano"},
	"SOL":   {id: "solana", name: "Solana"},
	"DOGE":  {id: "dogecoin", name: "Dogecoin"},
	"DOT":   {id: "polkadot", name: "Polkadot"},
	"MATIC": {id: "matic-network", name: "Polygon"},
	"AVAX":  {id: "avalanche-2", name: "Avalanche"},
	"LINK":  {id: "chainlink", name: "Chainlink"},
	"LTC":   {id: "litecoin", name: "Litecoin"},
	"UNI":   {id: "uniswap", name: "Uniswap"},
	"ATOM":  {id: "cosmos", name: "Cosmos"},
	"TRX":   {id: "tron", name: "TRON"},
	"XLM":   {id: "stellar", name: "Stellar"},
	"ALGO":  {id: "algorand", name: "Algorand"},
	"NEAR":  {id: "near", name: "NEAR Protocol"},
	"FTM":   {id: "fantom", name: "Fantom"},
	"SHIB":  {id: "shiba-inu", name: "Shiba Inu"},
}

// Service fetches cryptocurrency prices from CoinGecko and caches them in Redis.
type Service struct {
	rdb        *redis.Client
	httpClient *http.Client
	baseURL    string
}

// NewService creates a Service backed by the CoinGecko API.
func NewService(rdb *redis.Client) *Service {
	return &Service{
		rdb:        rdb,
		httpClient: &http.Client{Timeout: httpTimeout},
		baseURL:    coinGeckoURL,
	}
}

// newServiceWithClient is used in tests to inject a custom HTTP client and base URL.
func newServiceWithClient(client *http.Client, baseURL string) *Service {
	return &Service{httpClient: client, baseURL: baseURL}
}

// GetPrice returns current price data for the given ticker symbol.
func (s *Service) GetPrice(ctx context.Context, symbol string) (Price, error) {
	coin, ok := coinMap[symbol]
	if !ok {
		return Price{}, svcerr.Unknown("unknown_symbol", fmt.Sprintf("unsupported symbol: %s", symbol))
	}

	if s.rdb != nil {
		if p, ok := s.fromCache(ctx, symbol); ok {
			return p, nil
		}
	}

	price, err := s.fetchPrice(ctx, coin.id, symbol, coin.name)
	if err != nil {
		return Price{}, err
	}

	if s.rdb != nil {
		s.toCache(ctx, symbol, price)
	}

	return price, nil
}

// BatchPriceItem is the result for a single item in a batch crypto price request.
type BatchPriceItem struct {
	Symbol string `json:"symbol"`
	Result *Price `json:"result,omitempty"`
	Error  string `json:"error,omitempty"`
}

// GetPriceBatch returns prices for each symbol. Redis cache is checked per symbol;
// all cache misses are resolved with a single CoinGecko call.
func (s *Service) GetPriceBatch(ctx context.Context, symbols []string) []BatchPriceItem {
	results := make([]BatchPriceItem, len(symbols))

	type validSym struct {
		idx    int
		upper  string
		coinID string
		name   string
	}
	var valid []validSym
	for i, sym := range symbols {
		upper := strings.ToUpper(sym)
		coin, ok := coinMap[upper]
		if !ok {
			results[i] = BatchPriceItem{Symbol: sym, Error: fmt.Sprintf("unsupported symbol: %s", sym)}
			continue
		}
		valid = append(valid, validSym{idx: i, upper: upper, coinID: coin.id, name: coin.name})
	}

	if len(valid) == 0 {
		return results
	}

	// Check Redis cache for each valid symbol; collect misses.
	var cacheMisses []validSym
	for _, v := range valid {
		if s.rdb != nil {
			if p, ok := s.fromCache(ctx, v.upper); ok {
				results[v.idx] = BatchPriceItem{Symbol: v.upper, Result: &p}
				continue
			}
		}
		cacheMisses = append(cacheMisses, v)
	}

	if len(cacheMisses) == 0 {
		return results
	}

	// Single CoinGecko call for all cache misses.
	ids := make([]string, len(cacheMisses))
	for i, v := range cacheMisses {
		ids[i] = v.coinID
	}
	prices, err := s.fetchPriceBatch(ctx, ids)
	if err != nil {
		for _, v := range cacheMisses {
			results[v.idx] = BatchPriceItem{Symbol: v.upper, Error: "crypto price service unavailable"}
		}
		return results
	}

	for _, v := range cacheMisses {
		data, ok := prices[v.coinID]
		if !ok {
			results[v.idx] = BatchPriceItem{Symbol: v.upper, Error: "price not available"}
			continue
		}
		p := Price{
			Symbol:    v.upper,
			Name:      v.name,
			PriceUSD:  data.USD,
			Change24h: data.USD24hChange,
			MarketCap: data.USDMarketCap,
			Volume24h: data.USD24hVol,
		}
		if s.rdb != nil {
			s.toCache(ctx, v.upper, p)
		}
		results[v.idx] = BatchPriceItem{Symbol: v.upper, Result: &p}
	}

	return results
}

// fetchPriceBatch fetches prices for multiple coin IDs in a single CoinGecko call.
func (s *Service) fetchPriceBatch(ctx context.Context, coinIDs []string) (coinGeckoResponse, error) {
	url := fmt.Sprintf( //nolint:gosec // URL built from hardcoded base and coin IDs from coinMap
		"%s/simple/price?ids=%s&vs_currencies=usd&include_market_cap=true&include_24hr_vol=true&include_24hr_change=true",
		s.baseURL, strings.Join(coinIDs, ","),
	)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, http.NoBody)
	if err != nil {
		return nil, svcerr.Upstream("upstream_error", "crypto price service unavailable")
	}

	resp, err := s.httpClient.Do(req) //nolint:gosec // same URL, already validated above
	if err != nil {
		return nil, svcerr.Upstream("upstream_error", "crypto price service unavailable")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, svcerr.Upstream("upstream_error", "crypto price service unavailable")
	}

	var body coinGeckoResponse
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		return nil, svcerr.Upstream("upstream_error", "crypto price service unavailable")
	}

	return body, nil
}

func (s *Service) fromCache(ctx context.Context, symbol string) (Price, bool) {
	val, err := s.rdb.Get(ctx, cacheKey(symbol)).Result()
	if err != nil {
		return Price{}, false
	}

	var p Price
	if err := json.Unmarshal([]byte(val), &p); err != nil {
		return Price{}, false
	}

	return p, true
}

func (s *Service) toCache(ctx context.Context, symbol string, p Price) {
	b, err := json.Marshal(p)
	if err != nil {
		return
	}
	s.rdb.Set(ctx, cacheKey(symbol), string(b), cacheTTL)
}

// coinGeckoResponse is the JSON shape returned by CoinGecko /simple/price.
type coinGeckoResponse map[string]struct {
	USD          float64 `json:"usd"`
	USD24hChange float64 `json:"usd_24h_change"`
	USDMarketCap float64 `json:"usd_market_cap"`
	USD24hVol    float64 `json:"usd_24h_vol"`
}

func (s *Service) fetchPrice(ctx context.Context, coinID, symbol, name string) (Price, error) {
	url := fmt.Sprintf(
		"%s/simple/price?ids=%s&vs_currencies=usd&include_market_cap=true&include_24hr_vol=true&include_24hr_change=true",
		s.baseURL, coinID,
	)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, http.NoBody) //nolint:gosec // URL is built from a hardcoded base, not user input
	if err != nil {
		return Price{}, fmt.Errorf("crypto: build request: %w", err)
	}

	resp, err := s.httpClient.Do(req) //nolint:gosec // same URL, already validated above

	if err != nil {
		return Price{}, svcerr.Upstream("upstream_error", "crypto price service unavailable")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return Price{}, svcerr.Upstream("upstream_error", "crypto price service unavailable")
	}

	var body coinGeckoResponse
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		return Price{}, svcerr.Upstream("upstream_error", "crypto price service unavailable")
	}

	data, ok := body[coinID]
	if !ok {
		return Price{}, svcerr.Upstream("upstream_error", "crypto price service unavailable")
	}

	return Price{
		Symbol:    strings.ToUpper(symbol),
		Name:      name,
		PriceUSD:  data.USD,
		Change24h: data.USD24hChange,
		MarketCap: data.USDMarketCap,
		Volume24h: data.USD24hVol,
	}, nil
}

func cacheKey(symbol string) string {
	return fmt.Sprintf("crypto:%s", strings.ToUpper(symbol))
}
