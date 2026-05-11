package swift

import (
	"context"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"requiems-api/platform/svcerr"
)

func setupSwiftTestDB(t *testing.T) *pgxpool.Pool {
	t.Helper()

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		t.Skip("DATABASE_URL not set; skipping DB-backed swift service tests")
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
		t.Skipf("database unavailable for DB-backed swift tests: %v", err)
	}

	_, err = pool.Exec(ctx, `
		CREATE TABLE IF NOT EXISTS swift_codes (
			swift_code TEXT PRIMARY KEY,
			bank_code TEXT NOT NULL,
			country_code TEXT NOT NULL,
			location_code TEXT NOT NULL,
			branch_code TEXT NOT NULL,
			bank_name TEXT NOT NULL DEFAULT '',
			city TEXT NOT NULL DEFAULT '',
			country_name TEXT NOT NULL DEFAULT '',
			last_updated TIMESTAMPTZ NOT NULL DEFAULT NOW()
		)
	`)
	if err != nil {
		t.Skipf("cannot initialize swift_codes table for tests: %v", err)
	}

	_, err = pool.Exec(ctx, `DELETE FROM swift_codes WHERE swift_code LIKE 'TST%'`)
	if err != nil {
		t.Skipf("cannot clean swift test rows: %v", err)
	}

	t.Cleanup(func() {
		_, _ = pool.Exec(context.Background(), `DELETE FROM swift_codes WHERE swift_code LIKE 'TST%'`)
	})

	return pool
}

func insertSwiftFixture(t *testing.T, pool *pgxpool.Pool, code, bankCode, countryCode, locationCode, branchCode, bankName, city, countryName string) {
	t.Helper()
	_, err := pool.Exec(context.Background(), `
		INSERT INTO swift_codes (
			swift_code, bank_code, country_code, location_code, branch_code,
			bank_name, city, country_name
		) VALUES ($1,$2,$3,$4,$5,$6,$7,$8)
	`, code, bankCode, countryCode, locationCode, branchCode, bankName, city, countryName)
	require.NoError(t, err, "insert fixture %s", code)
}

func TestServiceLookupAndList_DBBacked(t *testing.T) {
	pool := setupSwiftTestDB(t)
	svc := NewService(pool)

	insertSwiftFixture(t, pool, "TSTADEFFXXX", "TSTA", "DE", "FF", "XXX", "Test Bank A", "Frankfurt", "Germany")
	insertSwiftFixture(t, pool, "TSTBUSNYXXX", "TSTB", "US", "NY", "XXX", "Test Bank B", "New York", "United States")
	insertSwiftFixture(t, pool, "TSTADEFF001", "TSTA", "DE", "FF", "001", "Test Bank A Branch", "Berlin", "Germany")

	ctx := context.Background()

	lookup, err := svc.Lookup(ctx, "TSTADEFF")
	require.NoError(t, err, "Lookup")
	require.Equal(t, "TSTADEFFXXX", lookup.SwiftCode)

	list, err := svc.List(ctx, ListFilter{CountryCode: "de", Limit: 10, Offset: 0})
	require.NoError(t, err, "List")
	require.True(t, list.Returned >= 2, "expected at least 2 DE records, got %d", list.Returned)

	filtered, err := svc.List(ctx, ListFilter{CountryCode: "US", BankCode: "TSTB", Limit: 10, Offset: 0})
	require.NoError(t, err, "List filtered")
	require.Equal(t, 1, filtered.Returned, "expected 1 US/TSTB fixture record, got %d", filtered.Returned)
}

func TestServiceList_InvalidFilters(t *testing.T) {
	pool := setupSwiftTestDB(t)
	svc := NewService(pool)

	_, err := svc.List(context.Background(), ListFilter{CountryCode: "D1"})
	require.Error(t, err)
	var se *svcerr.Error
	require.ErrorAs(t, err, &se)
	assert.Equal(t, "bad_request", se.Code)

	_, err = svc.List(context.Background(), ListFilter{BankCode: "TS1A"})
	require.Error(t, err)
}
