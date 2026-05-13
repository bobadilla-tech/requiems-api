package quotes

import (
	"context"
	"errors"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// mockRow implements pgx.Row.
type mockRow struct {
	scanFn func(dest ...any) error
}

func (m *mockRow) Scan(dest ...any) error { return m.scanFn(dest...) }

// mockQuerier implements querier, always returning the same row.
type mockQuerier struct {
	row pgx.Row
}

func (m *mockQuerier) QueryRow(_ context.Context, _ string, _ ...any) pgx.Row {
	return m.row
}

// multiMockQuerier implements querier returning a different row on each call.
// This allows testing that results are returned in the same order as requests.
type multiMockQuerier struct {
	rows  []pgx.Row
	index int
}

func (m *multiMockQuerier) QueryRow(_ context.Context, _ string, _ ...any) pgx.Row {
	if m.index >= len(m.rows) {
		return m.rows[len(m.rows)-1]
	}
	row := m.rows[m.index]
	m.index++
	return row
}

func newTestService(row pgx.Row) *Service {
	return &Service{db: &mockQuerier{row: row}}
}

func TestRandom_EmptyTable(t *testing.T) {
	t.Parallel()
	svc := newTestService(&mockRow{
		scanFn: func(_ ...any) error { return pgx.ErrNoRows },
	})

	_, err := svc.Random(context.Background())
	assert.ErrorIs(t, err, pgx.ErrNoRows)
}

func TestRandom_SingleRow(t *testing.T) {
	t.Parallel()
	svc := newTestService(&mockRow{
		scanFn: func(dest ...any) error {
			*dest[0].(*int) = 7
			*dest[1].(*string) = "Be yourself; everyone else is already taken."
			*dest[2].(*string) = "Oscar Wilde"
			return nil
		},
	})

	got, err := svc.Random(context.Background())
	require.NoError(t, err)
	assert.Equal(t, 7, got.ID)
	assert.Equal(t, "Be yourself; everyone else is already taken.", got.Text)
	assert.Equal(t, "Oscar Wilde", got.Author)
}

func TestRandom_ScanError(t *testing.T) {
	t.Parallel()
	scanErr := errors.New("scan failed")
	svc := newTestService(&mockRow{
		scanFn: func(_ ...any) error { return scanErr },
	})

	_, err := svc.Random(context.Background())
	assert.ErrorIs(t, err, scanErr)
}

func TestRandom_ReturnsZeroValueOnError(t *testing.T) {
	t.Parallel()
	svc := newTestService(&mockRow{
		scanFn: func(_ ...any) error { return pgx.ErrNoRows },
	})

	got, _ := svc.Random(context.Background())
	assert.Equal(t, 0, got.ID)
	assert.Equal(t, "", got.Text)
	assert.Equal(t, "", got.Author)
}

func TestRandom_EmptyAuthorAllowed(t *testing.T) {
	t.Parallel()
	svc := newTestService(&mockRow{
		scanFn: func(dest ...any) error {
			*dest[0].(*int) = 3
			*dest[1].(*string) = "Anonymous wisdom."
			*dest[2].(*string) = ""
			return nil
		},
	})

	got, err := svc.Random(context.Background())
	require.NoError(t, err)
	assert.Equal(t, "", got.Author)
}

func TestRandomBatch_ReturnsNQuotes(t *testing.T) {
	t.Parallel()
	svc := newTestService(&mockRow{
		scanFn: func(dest ...any) error {
			*dest[0].(*int) = 1
			*dest[1].(*string) = "Test quote."
			*dest[2].(*string) = "Test Author"
			return nil
		},
	})

	got, err := svc.RandomBatch(context.Background(), 3)
	require.NoError(t, err)
	assert.Len(t, got, 3)
	assert.Equal(t, "Test quote.", got[0].Text)
	assert.Equal(t, "Test quote.", got[2].Text)
}

func TestRandomBatch_ContinuesOnItemError(t *testing.T) {
	t.Parallel()
	// When a single quote fails to scan, the batch must not abort:
	// it returns n zero-value Quotes with no error.
	svc := newTestService(&mockRow{
		scanFn: func(_ ...any) error { return pgx.ErrNoRows },
	})

	results, err := svc.RandomBatch(context.Background(), 2)
	assert.NoError(t, err)
	assert.Len(t, results, 2)
	assert.Equal(t, 0, results[0].ID)
	assert.Equal(t, "", results[0].Text)
}

func TestRandomBatch_SingleItem(t *testing.T) {
	t.Parallel()
	svc := newTestService(&mockRow{
		scanFn: func(dest ...any) error {
			*dest[0].(*int) = 5
			*dest[1].(*string) = "One is enough."
			*dest[2].(*string) = "Someone"
			return nil
		},
	})

	got, err := svc.RandomBatch(context.Background(), 1)
	require.NoError(t, err)
	assert.Len(t, got, 1)
}

func TestBatchRandomResponse_IsData(t *testing.T) {
	t.Parallel()
	// IsData() must be callable — verifies the interface is satisfied.
	BatchRandomResponse{}.IsData()
}

func TestRandomBatch_ReturnsResultsInOrder(t *testing.T) {
	t.Parallel()
	// Each call to QueryRow should return a different quote.
	// The batch must preserve the order: result[0] comes from the first call, etc.
	makeRow := func(id int, text, author string) pgx.Row {
		return &mockRow{
			scanFn: func(dest ...any) error {
				*dest[0].(*int) = id
				*dest[1].(*string) = text
				*dest[2].(*string) = author
				return nil
			},
		}
	}

	svc := &Service{db: &multiMockQuerier{
		rows: []pgx.Row{
			makeRow(1, "First quote.", "Author A"),
			makeRow(2, "Second quote.", "Author B"),
			makeRow(3, "Third quote.", "Author C"),
		},
	}}

	got, err := svc.RandomBatch(context.Background(), 3)
	require.NoError(t, err)
	assert.Equal(t, "First quote.", got[0].Text)
	assert.Equal(t, "Second quote.", got[1].Text)
	assert.Equal(t, "Third quote.", got[2].Text)
}
