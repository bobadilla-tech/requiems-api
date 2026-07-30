package bin

import (
	"context"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupBINTestDB(t *testing.T) *pgxpool.Pool {
	t.Helper()

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		t.Skip("DATABASE_URL not set; skipping DB-backed BIN service tests")
	}

	ctx := context.Background()
	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		t.Skipf("unable to create pgx pool: %v", err)
	}

	t.Cleanup(func() {
		pool.Close()
	})

	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		t.Skipf("database unavailable for DB-backed BIN tests: %v", err)
	}

	_, err = pool.Exec(ctx, `
		CREATE TABLE IF NOT EXISTS bin_data (
			bin_prefix    VARCHAR(8)   PRIMARY KEY,
			prefix_length SMALLINT     NOT NULL DEFAULT 6,
			scheme        TEXT         NOT NULL DEFAULT '',
			card_type     TEXT         NOT NULL DEFAULT '',
			card_level    TEXT         NOT NULL DEFAULT '',
			issuer_name   TEXT         NOT NULL DEFAULT '',
			issuer_url    TEXT         NOT NULL DEFAULT '',
			issuer_phone  TEXT         NOT NULL DEFAULT '',
			country_code  CHAR(2)      NOT NULL DEFAULT '',
			country_name  TEXT         NOT NULL DEFAULT '',
			prepaid       BOOLEAN      NOT NULL DEFAULT FALSE,
			source        TEXT         NOT NULL DEFAULT '',
			confidence    NUMERIC(3,2) NOT NULL DEFAULT 0.50,
			first_seen    TIMESTAMPTZ  NOT NULL DEFAULT NOW(),
			last_updated  TIMESTAMPTZ  NOT NULL DEFAULT NOW()
		)
	`)
	if err != nil {
		t.Skipf("cannot initialize bin_data table for tests: %v", err)
	}

	_, err = pool.Exec(ctx, `DELETE FROM bin_data WHERE bin_prefix LIKE 'TST%' OR bin_prefix IN ('424242','42424242','510000','999001')`)
	if err != nil {
		t.Skipf("cannot clean BIN test rows: %v", err)
	}

	t.Cleanup(func() {
		_, _ = pool.Exec(context.Background(), `DELETE FROM bin_data WHERE bin_prefix LIKE 'TST%' OR bin_prefix IN ('424242','42424242','510000','999001')`)
	})

	return pool
}

func insertBINFixture(t *testing.T, pool *pgxpool.Pool, prefix, scheme, cardType, cardLevel, issuerName, countryCode, countryName string, prepaid bool, confidence float64) {
	t.Helper()
	_, err := pool.Exec(context.Background(), `
		INSERT INTO bin_data (bin_prefix, prefix_length, scheme, card_type, card_level, issuer_name, country_code, country_name, prepaid, confidence)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
		ON CONFLICT (bin_prefix) DO NOTHING
	`, prefix, len(prefix), scheme, cardType, cardLevel, issuerName, countryCode, countryName, prepaid, confidence)
	require.NoError(t, err, "insert BIN fixture %s", prefix)
}

// ---- Lookup (single) ----

func TestServiceLookup_ExactMatch(t *testing.T) {
	pool := setupBINTestDB(t)
	svc := NewService(pool)

	insertBINFixture(t, pool, "424242", "visa", "credit", "classic", "Chase", "US", "United States", false, 0.95)

	result, err := svc.Lookup(context.Background(), "424242")
	require.NoError(t, err)
	assert.Equal(t, "424242", result.BIN)
	assert.Equal(t, "visa", result.Scheme)
	assert.Equal(t, "credit", result.CardType)
	assert.Equal(t, "US", result.CountryCode)
}

func TestServiceLookup_8DigitFallsBackTo6(t *testing.T) {
	pool := setupBINTestDB(t)
	svc := NewService(pool)

	// Only the 6-digit prefix exists in DB; 8-digit lookup should fall back.
	insertBINFixture(t, pool, "424242", "visa", "credit", "classic", "Chase", "US", "United States", false, 0.90)

	result, err := svc.Lookup(context.Background(), "42424299")
	require.NoError(t, err)
	assert.Equal(t, "42424299", result.BIN)
	assert.Equal(t, "visa", result.Scheme)
}

func TestServiceLookup_NotFound(t *testing.T) {
	pool := setupBINTestDB(t)
	svc := NewService(pool)

	_, err := svc.Lookup(context.Background(), "000000")
	require.Error(t, err)
}

func TestServiceLookup_SchemeDetectedFromPrefix(t *testing.T) {
	pool := setupBINTestDB(t)
	svc := NewService(pool)

	// Row has empty scheme — detectScheme should fill it in.
	insertBINFixture(t, pool, "510000", "", "credit", "classic", "Test Bank", "US", "United States", false, 0.80)

	result, err := svc.Lookup(context.Background(), "510000")
	require.NoError(t, err)
	assert.Equal(t, "mastercard", result.Scheme)
}

// ---- LookupBatch ----

func TestServiceLookupBatch_HappyPath(t *testing.T) {
	pool := setupBINTestDB(t)
	svc := NewService(pool)

	insertBINFixture(t, pool, "424242", "visa", "credit", "classic", "Chase", "US", "United States", false, 0.95)
	insertBINFixture(t, pool, "510000", "mastercard", "credit", "gold", "Citi", "US", "United States", false, 0.92)

	results := svc.LookupBatch(context.Background(), []string{"424242", "510000", "000000"})

	require.Len(t, results, 3)
	assert.True(t, results[0].Found)
	assert.Equal(t, "visa", results[0].Result.Scheme)
	assert.True(t, results[1].Found)
	assert.Equal(t, "mastercard", results[1].Result.Scheme)
	assert.False(t, results[2].Found)
	assert.Empty(t, results[2].Error)
}

func TestServiceLookupBatch_8DigitFallback(t *testing.T) {
	pool := setupBINTestDB(t)
	svc := NewService(pool)

	insertBINFixture(t, pool, "999001", "visa", "debit", "classic", "Test Bank", "AR", "Argentina", true, 0.75)

	// 99900199 has no exact 8-digit match — should fall back to 999001.
	results := svc.LookupBatch(context.Background(), []string{"99900199"})

	require.Len(t, results, 1)
	assert.True(t, results[0].Found)
	assert.Equal(t, "99900199", results[0].BIN)
}

func TestServiceLookupBatch_MixedValidAndInvalid(t *testing.T) {
	pool := setupBINTestDB(t)
	svc := NewService(pool)

	insertBINFixture(t, pool, "424242", "visa", "credit", "classic", "Chase", "US", "United States", false, 0.95)

	results := svc.LookupBatch(context.Background(), []string{"424242", "abc", "000000"})

	require.Len(t, results, 3)
	assert.True(t, results[0].Found)
	assert.NotEmpty(t, results[1].Error) // "abc" → validation error
	assert.False(t, results[2].Found)    // "000000" → not in DB
}

func TestServiceLookupBatch_OrderPreserved(t *testing.T) {
	pool := setupBINTestDB(t)
	svc := NewService(pool)

	insertBINFixture(t, pool, "424242", "visa", "credit", "classic", "Chase", "US", "United States", false, 0.95)
	insertBINFixture(t, pool, "510000", "mastercard", "credit", "gold", "Citi", "US", "United States", false, 0.92)

	results := svc.LookupBatch(context.Background(), []string{"510000", "424242"})

	require.Len(t, results, 2)
	assert.Equal(t, "510000", results[0].BIN)
	assert.Equal(t, "424242", results[1].BIN)
}
