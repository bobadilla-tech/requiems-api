package bin

import (
	"context"
	"errors"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"requiems-api/platform/svcerr"
)

// LookupResponse is the response payload for GET /v1/finance/bin/:bin.
type LookupResponse struct {
	BIN         string `json:"bin"`
	Scheme      string `json:"scheme"`
	CardType    string `json:"card_type"`
	CardLevel   string `json:"card_level"`
	IssuerName  string `json:"issuer_name"`
	IssuerURL   string `json:"issuer_url"`
	IssuerPhone string `json:"issuer_phone"`
	CountryCode string `json:"country_code"`
	CountryName string `json:"country_name"`
	Prepaid     bool   `json:"prepaid"`
	// LuhnPrefixValid reports whether the BIN prefix (6-8 digits) itself
	// passes the Luhn checksum — it does not validate a full card PAN.
	LuhnPrefixValid bool    `json:"luhn_prefix_valid"`
	Confidence      float64 `json:"confidence"`
	DataFreshness   string  `json:"data_freshness,omitempty"`
}

// BatchBINItem is the result for a single item in a batch BIN lookup request.
type BatchBINItem struct {
	BIN    string          `json:"bin"`
	Found  bool            `json:"found"`
	Result *LookupResponse `json:"result,omitempty"`
	Error  string          `json:"error,omitempty"`
}

type batchBINValid struct {
	idx  int
	bin  string
	luhn bool
}

// Service provides BIN lookup against the bin_data PostgreSQL table.
type Service struct {
	db *pgxpool.Pool
}

// NewService creates a new Service backed by the given connection pool.
func NewService(db *pgxpool.Pool) *Service {
	return &Service{db: db}
}

// Lookup validates the raw BIN string and queries the database for card
// metadata. It tries an exact match first; if the input is 8 digits and no
// row is found, it falls back to the 6-digit prefix.
func (s *Service) Lookup(ctx context.Context, raw string) (LookupResponse, error) {
	bin, err := sanitizeBIN(raw)
	if err != nil {
		return LookupResponse{}, err
	}

	luhn := luhnValid(bin)

	result, err := s.queryBIN(ctx, bin)
	if errors.Is(err, pgx.ErrNoRows) && len(bin) == 8 {
		// Fall back to the 6-digit prefix using a LEFT match so it finds both
		// 6-digit rows stored as-is and 8-digit rows whose prefix matches.
		result, err = s.queryBINByPrefix6(ctx, bin[:6])
	}

	if errors.Is(err, pgx.ErrNoRows) {
		return LookupResponse{}, svcerr.NotFound("not_found", "BIN not found")
	}
	if err != nil {
		return LookupResponse{}, err
	}

	// If the database row has no scheme, derive it from the prefix.
	if result.Scheme == "" {
		result.Scheme = detectScheme(bin)
	}

	result.BIN = bin
	result.LuhnPrefixValid = luhn
	return result, nil
}

// LookupBatch validates and looks up each BIN in one or two SQL queries.
// Invalid BINs and not-found BINs return in-band errors/found:false.
func (s *Service) LookupBatch(ctx context.Context, rawBINs []string) []BatchBINItem {
	results := make([]BatchBINItem, len(rawBINs))

	var valid []batchBINValid
	for i, raw := range rawBINs {
		bin, err := sanitizeBIN(raw)
		if err != nil {
			results[i] = BatchBINItem{BIN: raw, Error: err.Error()}
			continue
		}
		valid = append(valid, batchBINValid{idx: i, bin: bin, luhn: luhnValid(bin)})
	}

	if len(valid) == 0 {
		return results
	}

	// Exact-match query for all valid BINs in a single round-trip.
	prefixes := make([]string, len(valid))
	for i, v := range valid {
		prefixes[i] = v.bin
	}
	exactHits, err := s.queryBINBatch(ctx, prefixes)
	if err != nil {
		log.Printf("bin: LookupBatch exact query failed: %v", err)
		for _, v := range valid {
			results[v.idx] = BatchBINItem{BIN: v.bin, Error: "lookup failed"}
		}
		return results
	}

	// Collect 8-digit BINs that had no exact match — need prefix-6 fallback.
	type fallbackEntry struct {
		p6 string
	}
	var fallbacks []fallbackEntry
	for _, v := range valid {
		if _, found := exactHits[v.bin]; !found && len(v.bin) == 8 {
			fallbacks = append(fallbacks, fallbackEntry{p6: v.bin[:6]})
		}
	}

	var fallbackHits map[string]LookupResponse
	if len(fallbacks) > 0 {
		p6s := make([]string, len(fallbacks))
		for i, f := range fallbacks {
			p6s[i] = f.p6
		}
		fallbackHits, err = s.queryBINByPrefix6Batch(ctx, p6s)
		if err != nil {
			log.Printf("bin: LookupBatch prefix-6 fallback query failed: %v", err)
			s.resolveFallbackError(results, valid, exactHits)
			return results
		}
	}

	for _, v := range valid {
		results[v.idx] = s.lookupBatchItem(v, exactHits, fallbackHits)
	}

	return results
}

// resolveFallbackError handles a DB failure during the prefix-6 fallback query.
// Items that had an exact hit are resolved normally; 8-digit misses receive a
// generic in-band error; 6-digit misses are returned as not found.
func (s *Service) resolveFallbackError(results []BatchBINItem, valid []batchBINValid, exactHits map[string]LookupResponse) {
	for _, v := range valid {
		if _, found := exactHits[v.bin]; found {
			results[v.idx] = s.lookupBatchItem(v, exactHits, nil)
			continue
		}
		if len(v.bin) == 8 {
			results[v.idx] = BatchBINItem{BIN: v.bin, Error: "lookup failed"}
		} else {
			results[v.idx] = BatchBINItem{BIN: v.bin, Found: false}
		}
	}
}

func (s *Service) lookupBatchItem(v batchBINValid, exactHits, fallbackHits map[string]LookupResponse) BatchBINItem {
	if r, ok := exactHits[v.bin]; ok {
		return s.batchBINItem(v, r)
	}

	if len(v.bin) != 8 {
		return BatchBINItem{BIN: v.bin, Found: false}
	}

	if r, ok := fallbackHits[v.bin[:6]]; ok {
		return s.batchBINItem(v, r)
	}

	return BatchBINItem{BIN: v.bin, Found: false}
}

func (s *Service) batchBINItem(v batchBINValid, r LookupResponse) BatchBINItem {
	if r.Scheme == "" {
		r.Scheme = detectScheme(v.bin)
	}
	r.BIN = v.bin
	r.LuhnPrefixValid = v.luhn
	return BatchBINItem{BIN: v.bin, Found: true, Result: &r}
}
