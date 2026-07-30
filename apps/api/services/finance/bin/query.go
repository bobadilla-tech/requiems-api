package bin

import (
	"context"
	"log"
)

// This file contains all database query helpers for the BIN service.
// Methods are defined on *Service; the DB pool lives in service.go.

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
		ORDER BY LENGTH(bin_prefix) DESC, bin_prefix ASC
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

// queryBINBatch fetches all rows whose bin_prefix exactly matches one of the
// given prefixes. Returns a map keyed by bin_prefix for O(1) lookup.
func (s *Service) queryBINBatch(ctx context.Context, prefixes []string) (map[string]LookupResponse, error) {
	rows, err := s.db.Query(ctx, `
		SELECT
			bin_prefix, scheme, card_type, card_level,
			issuer_name, issuer_url, issuer_phone,
			country_code, country_name,
			prepaid, confidence
		FROM bin_data
		WHERE bin_prefix = ANY($1::text[])
	`, prefixes)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	hits := make(map[string]LookupResponse)
	for rows.Next() {
		var r LookupResponse
		if err := rows.Scan(
			&r.BIN, &r.Scheme, &r.CardType, &r.CardLevel,
			&r.IssuerName, &r.IssuerURL, &r.IssuerPhone,
			&r.CountryCode, &r.CountryName,
			&r.Prepaid, &r.Confidence,
		); err != nil {
			log.Printf("bin: queryBINBatch scan error: %v", err)
			return nil, err
		}
		hits[r.BIN] = r
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return hits, nil
}

// queryBINByPrefix6Batch fetches one best-match row per 6-digit prefix for all
// entries in prefixes6. Returns a map keyed by the 6-digit prefix.
func (s *Service) queryBINByPrefix6Batch(ctx context.Context, prefixes6 []string) (map[string]LookupResponse, error) {
	rows, err := s.db.Query(ctx, `
		SELECT DISTINCT ON (LEFT(bin_prefix, 6))
			bin_prefix, scheme, card_type, card_level,
			issuer_name, issuer_url, issuer_phone,
			country_code, country_name,
			prepaid, confidence
		FROM bin_data
		WHERE LEFT(bin_prefix, 6) = ANY($1::text[])
		ORDER BY LEFT(bin_prefix, 6), LENGTH(bin_prefix) DESC, bin_prefix ASC
	`, prefixes6)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	hits := make(map[string]LookupResponse)
	for rows.Next() {
		var r LookupResponse
		if err := rows.Scan(
			&r.BIN, &r.Scheme, &r.CardType, &r.CardLevel,
			&r.IssuerName, &r.IssuerURL, &r.IssuerPhone,
			&r.CountryCode, &r.CountryName,
			&r.Prepaid, &r.Confidence,
		); err != nil {
			log.Printf("bin: queryBINByPrefix6Batch scan error: %v", err)
			return nil, err
		}
		prefix6 := r.BIN
		if len(r.BIN) > 6 {
			prefix6 = r.BIN[:6]
		}
		hits[prefix6] = r
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return hits, nil
}
