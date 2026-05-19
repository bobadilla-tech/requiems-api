package bin

import (
	"context"
	"errors"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"requiems-api/platform/svcerr"
)

// LookupResponse is the response payload for GET /v1/finance/bin/:bin.
type LookupResponse struct {
	BIN         string  `json:"bin"`
	Scheme      string  `json:"scheme"`
	CardType    string  `json:"card_type"`
	CardLevel   string  `json:"card_level"`
	IssuerName  string  `json:"issuer_name"`
	IssuerURL   string  `json:"issuer_url"`
	IssuerPhone string  `json:"issuer_phone"`
	CountryCode string  `json:"country_code"`
	CountryName string  `json:"country_name"`
	Prepaid     bool    `json:"prepaid"`
	Luhn        bool    `json:"luhn"`
	Confidence  float64 `json:"confidence"`
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
	result.Luhn = luhn
	return result, nil
}

// queryBIN executes a single point-lookup against bin_data by exact prefix.
func (s *Service) queryBIN(ctx context.Context, prefix string) (LookupResponse, error) {
	row := s.db.QueryRow(ctx, `
		SELECT
			bin_prefix, scheme, card_type, card_level,
			issuer_name, issuer_url, issuer_phone,
			country_code, country_name,
			prepaid, confidence
		FROM bin_data
		WHERE bin_prefix = $1
	`, prefix)

	var r LookupResponse
	err := row.Scan(
		&r.BIN, &r.Scheme, &r.CardType, &r.CardLevel,
		&r.IssuerName, &r.IssuerURL, &r.IssuerPhone,
		&r.CountryCode, &r.CountryName,
		&r.Prepaid, &r.Confidence,
	)
	return r, err
}

// queryBINByPrefix6 finds a row whose first 6 digits match prefix6, using the
// functional index on LEFT(bin_prefix, 6). Used as a fallback when an 8-digit
// exact lookup returns no rows.
func (s *Service) queryBINByPrefix6(ctx context.Context, prefix6 string) (LookupResponse, error) {
	row := s.db.QueryRow(ctx, `
		SELECT
			bin_prefix, scheme, card_type, card_level,
			issuer_name, issuer_url, issuer_phone,
			country_code, country_name,
			prepaid, confidence
		FROM bin_data
		WHERE LEFT(bin_prefix, 6) = $1
		LIMIT 1
	`, prefix6)

	var r LookupResponse
	err := row.Scan(
		&r.BIN, &r.Scheme, &r.CardType, &r.CardLevel,
		&r.IssuerName, &r.IssuerURL, &r.IssuerPhone,
		&r.CountryCode, &r.CountryName,
		&r.Prepaid, &r.Confidence,
	)
	return r, err
}

// sanitizeBIN strips common separators and validates that the result is 6–8
// decimal digits.
func sanitizeBIN(raw string) (string, error) {
	cleaned := strings.Map(func(r rune) rune {
		if r == '-' || r == ' ' {
			return -1
		}
		return r
	}, strings.TrimSpace(raw))

	if len(cleaned) < 6 || len(cleaned) > 8 {
		return "", svcerr.Invalid("bad_request", "BIN must be between 6 and 8 digits")
	}

	for _, ch := range cleaned {
		if ch < '0' || ch > '9' {
			return "", svcerr.Invalid("bad_request", "BIN must contain digits only")
		}
	}

	return cleaned, nil
}

// luhnValid runs the Luhn algorithm on the digit string.
// For a BIN (6–8 digits) this is a partial check on the prefix only.
func luhnValid(s string) bool {
	sum := 0
	nDigits := len(s)
	parity := nDigits % 2

	for i, ch := range s {
		if ch < '0' || ch > '9' {
			return false
		}
		digit := int(ch - '0')
		if i%2 == parity {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		sum += digit
	}
	return sum%10 == 0
}

// detectScheme derives the card scheme from a BIN prefix using the canonical
// ISO/IEC 7812 prefix ranges. Ranges are checked from most specific to least
// specific to avoid false matches on overlapping prefixes.
func detectScheme(bin string) string {
	if len(bin) < 4 {
		return ""
	}

	n2 := atoiN(bin, 2)
	n3 := atoiN(bin, 3)
	n4 := atoiN(bin, 4)
	n6 := atoiN(bin, 6)

	switch {
	// Mir: 2200–2204 — must come before Mastercard 2-series
	case n4 >= 2200 && n4 <= 2204:
		return "mir"

	// Mastercard 2-series: 2221–2720
	case n4 >= 2221 && n4 <= 2720:
		return "mastercard"

	// Amex: 34, 37 — must come before Visa (both start with 3x)
	case n2 == 34 || n2 == 37:
		return "amex"

	// JCB: 3528–3589
	case n4 >= 3528 && n4 <= 3589:
		return "jcb"

	// Diners Club: 300–305, 36, 38
	case (n4 >= 3000 && n4 <= 3059) || n2 == 36 || n2 == 38:
		return "diners"

	// Visa: starts with 4
	case bin[0] == '4':
		return "visa"

	// Mastercard 5-series: 51–55
	case n2 >= 51 && n2 <= 55:
		return "mastercard"

	// Maestro specific prefixes — check before UnionPay (overlapping 6x space)
	case n4 == 6304 || n4 == 6759 || n4 == 6761 || n4 == 6762 || n4 == 6763:
		return "maestro"

	// Discover: 6011
	case n4 == 6011:
		return "discover"

	// Discover: 622126–622925 — must come before UnionPay 62xx
	case n6 >= 622126 && n6 <= 622925:
		return "discover"

	// RuPay: 6521, 6522 — must come before Discover 65xx range
	case n4 == 6521 || n4 == 6522:
		return "rupay"

	// Discover: 644–649 and 65xx
	case (n3 >= 644 && n3 <= 649) || n2 == 65:
		return "discover"

	// RuPay: 60 — check before UnionPay 62
	case n2 == 60:
		return "rupay"

	// UnionPay: 62, 81
	case n2 == 62 || n2 == 81:
		return "unionpay"
	}

	return ""
}

// atoiN converts the first n digits of s to an integer. Returns 0 if s has
// fewer than n characters or contains non-digit bytes.
func atoiN(s string, n int) int {
	if len(s) < n {
		return 0
	}
	v := 0
	for i := 0; i < n; i++ {
		b := s[i]
		if b < '0' || b > '9' {
			return 0
		}
		v = v*10 + int(b-'0')
	}
	return v
}
