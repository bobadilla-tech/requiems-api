package iban

import (
	"context"
	"errors"
	"strings"
	"unicode"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type querier interface {
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
}

// Service provides IBAN validation and parsing against the iban_countries table.
type Service struct {
	db querier
}

// NewService creates a new Service backed by the given connection pool.
func NewService(db *pgxpool.Pool) *Service {
	return &Service{db: db}
}

// countryRow holds the format data for a single country from iban_countries.
type countryRow struct {
	name          string
	ibanLength    int16
	bankOffset    int16
	bankLength    int16
	accountOffset int16
	accountLength int16
}

// Parse normalises raw input, validates the IBAN checksum, and extracts the
// bank code and account number for countries present in the iban_countries
// table.
//
// Parse always returns a ParseResponse; the Valid field indicates whether the
// IBAN passed all validation checks. A non-nil error indicates an
// infrastructure failure (e.g. database unreachable).
func (s *Service) Parse(ctx context.Context, raw string) (ParseResponse, error) {
	iban := normalizeIBAN(raw)

	// Reject input that cannot possibly be an IBAN.
	if !basicFormatOK(iban) {
		return ParseResponse{IBAN: iban, Valid: false}, nil
	}

	countryCode := iban[:2]

	// Look up the country format from the database.
	row, err := s.queryCountry(ctx, countryCode)
	if err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return ParseResponse{}, err
	}

	country := row.name
	known := !errors.Is(err, pgx.ErrNoRows)

	// Validate length when the country is known.
	if known && int(row.ibanLength) != len(iban) {
		return ParseResponse{IBAN: iban, Valid: false, Country: country}, nil
	}

	// Validate ISO 13616 checksum (mod-97 of the rearranged string must be 1).
	if !validateChecksum(iban) {
		return ParseResponse{IBAN: iban, Valid: false, Country: country}, nil
	}

	// Extract bank code and account number from the BBAN.
	bankCode, account := "", ""
	if known {
		bban := iban[4:]
		bankCode = extract(bban, int(row.bankOffset), int(row.bankLength))
		account = extract(bban, int(row.accountOffset), int(row.accountLength))
	}

	return ParseResponse{
		IBAN:     iban,
		Valid:    true,
		Country:  country,
		BankCode: bankCode,
		Account:  account,
	}, nil
}

func (s *Service) queryCountry(ctx context.Context, code string) (countryRow, error) {
	var r countryRow
	err := s.db.QueryRow(ctx, `
		SELECT country_name, iban_length, bank_offset, bank_length, account_offset, account_length
		FROM iban_countries
		WHERE country_code = $1
	`, code).Scan(
		&r.name, &r.ibanLength,
		&r.bankOffset, &r.bankLength,
		&r.accountOffset, &r.accountLength,
	)
	return r, err
}

// normalizeIBAN strips internal spaces and uppercases the input.
func normalizeIBAN(raw string) string {
	return strings.ToUpper(strings.Map(func(r rune) rune {
		if r == ' ' {
			return -1
		}
		return r
	}, strings.TrimSpace(raw)))
}

// basicFormatOK returns true when s has at least 5 characters, begins with
// two uppercase ASCII letters (country code) followed by two ASCII digits
// (check digits), and contains only alphanumeric characters thereafter.
func basicFormatOK(s string) bool {
	if len(s) < 5 {
		return false
	}
	for i, ch := range s {
		switch {
		case i < 2:
			if ch < 'A' || ch > 'Z' {
				return false
			}
		case i < 4:
			if ch < '0' || ch > '9' {
				return false
			}
		default:
			if !unicode.IsLetter(ch) && !unicode.IsDigit(ch) {
				return false
			}
		}
	}
	return true
}

// validateChecksum implements the ISO 13616 mod-97 IBAN checksum:
//  1. Rearrange: move the first 4 characters to the end.
//  2. Replace each letter with its numeric equivalent (A=10 … Z=35).
//  3. Compute the large integer modulo 97.
//
// A valid IBAN produces a remainder of 1.
func validateChecksum(iban string) bool {
	return mod97(iban[4:]+iban[:4]) == 1
}

// mod97 computes the integer value of the numeric string s modulo 97.
// Letters are expanded inline to their 2-digit equivalents (A=10 … Z=35).
// The computation is done digit-by-digit to avoid big-integer arithmetic.
func mod97(s string) int {
	rem := 0
	for _, ch := range s {
		switch {
		case ch >= '0' && ch <= '9':
			rem = (rem*10 + int(ch-'0')) % 97
		case ch >= 'A' && ch <= 'Z':
			// A letter expands to two decimal digits, so multiply remainder by 100.
			rem = (rem*100 + int(ch-'A') + 10) % 97
		}
	}
	return rem
}

// extract returns the substring of s at [offset, offset+length). Returns an
// empty string when length is zero, offset is negative, or the slice would
// exceed the string bounds.
func extract(s string, offset, length int) string {
	if length <= 0 || offset < 0 || offset+length > len(s) {
		return ""
	}
	return s[offset : offset+length]
}

// ParseBatch parses a slice of IBANs and returns the results in the same order as the input.
func (s *Service) ParseBatch(ctx context.Context, numbers []string) ([]ParseResponse, error) {
	results := make([]ParseResponse, len(numbers))
	for i, n := range numbers {
		result, err := s.Parse(ctx, n)
		if err != nil {
			return nil, err
		}
		results[i] = result
	}
	return results, nil
}
