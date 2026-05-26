package quotes

import (
	"context"
	"errors"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// mockRow implements pgx.Row.
type mockRow struct {
	scanFn func(dest ...any) error
}

func (m *mockRow) Scan(dest ...any) error { return m.scanFn(dest...) }

// mockQuerier implements quotesDB for single-row tests.
type mockQuerier struct {
	row pgx.Row
}

func (m *mockQuerier) QueryRow(_ context.Context, _ string, _ ...any) pgx.Row { return m.row }
func (m *mockQuerier) Query(_ context.Context, _ string, _ ...any) (pgx.Rows, error) {
	return &mockRows{}, nil
}

// mockRows implements pgx.Rows for batch tests.
// Set quotes for happy-path rows; set total+scanErr to simulate scan failures.
type mockRows struct {
	quotes  []Quote
	total   int // used when simulating scan errors (quotes is nil)
	index   int
	err     error
	scanErr error
}

func (m *mockRows) Close()                                       {}
func (m *mockRows) Err() error                                   { return m.err }
func (m *mockRows) CommandTag() pgconn.CommandTag                { return pgconn.CommandTag{} }
func (m *mockRows) FieldDescriptions() []pgconn.FieldDescription { return nil }
func (m *mockRows) Values() ([]any, error)                       { return nil, nil }
func (m *mockRows) RawValues() [][]byte                          { return nil }
func (m *mockRows) Conn() *pgx.Conn                              { return nil }
func (m *mockRows) Next() bool {
	if m.quotes != nil {
		return m.index < len(m.quotes)
	}
	return m.index < m.total
}
func (m *mockRows) Scan(dest ...any) error {
	m.index++
	if m.scanErr != nil {
		return m.scanErr
	}
	q := m.quotes[m.index-1]
	*dest[0].(*int) = q.ID
	*dest[1].(*string) = q.Text
	*dest[2].(*string) = q.Author
	return nil
}

// batchQuerier returns fixed rows from Query and a fixed row from QueryRow.
type batchQuerier struct {
	row  pgx.Row
	rows *mockRows
}

func (b *batchQuerier) QueryRow(_ context.Context, _ string, _ ...any) pgx.Row { return b.row }
func (b *batchQuerier) Query(_ context.Context, _ string, _ ...any) (pgx.Rows, error) {
	return b.rows, nil
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
	svc := &Service{db: &batchQuerier{
		rows: &mockRows{
			quotes: []Quote{
				{ID: 1, Text: "Test quote.", Author: "Test Author"},
				{ID: 2, Text: "Test quote.", Author: "Test Author"},
				{ID: 3, Text: "Test quote.", Author: "Test Author"},
			},
		},
	}}

	got, err := svc.RandomBatch(context.Background(), 3)
	require.NoError(t, err)
	assert.Len(t, got, 3)
	assert.Equal(t, "Test quote.", got[0].Text)
	assert.Equal(t, "Test quote.", got[2].Text)
}

func TestRandomBatch_ContinuesOnItemError(t *testing.T) {
	t.Parallel()
	// Scan errors are skipped; the batch returns only successfully scanned rows.
	svc := &Service{db: &batchQuerier{
		rows: &mockRows{total: 2, scanErr: pgx.ErrNoRows},
	}}

	results, err := svc.RandomBatch(context.Background(), 2)
	assert.NoError(t, err)
	assert.Empty(t, results)
}

func TestRandomBatch_SingleItem(t *testing.T) {
	t.Parallel()
	svc := &Service{db: &batchQuerier{
		rows: &mockRows{
			quotes: []Quote{{ID: 5, Text: "One is enough.", Author: "Someone"}},
		},
	}}

	got, err := svc.RandomBatch(context.Background(), 1)
	require.NoError(t, err)
	assert.Len(t, got, 1)
}

func TestRandomBatch_InvalidCount(t *testing.T) {
	t.Parallel()
	svc := newTestService(&mockRow{
		scanFn: func(_ ...any) error { return nil },
	})

	_, err := svc.RandomBatch(context.Background(), 0)
	assert.Error(t, err)

	_, err = svc.RandomBatch(context.Background(), -1)
	assert.Error(t, err)
}

func TestRandomBatch_ReturnsResultsInOrder(t *testing.T) {
	t.Parallel()
	svc := &Service{db: &batchQuerier{
		rows: &mockRows{
			quotes: []Quote{
				{ID: 1, Text: "First quote.", Author: "Author A"},
				{ID: 2, Text: "Second quote.", Author: "Author B"},
				{ID: 3, Text: "Third quote.", Author: "Author C"},
			},
		},
	}}

	got, err := svc.RandomBatch(context.Background(), 3)
	require.NoError(t, err)
	assert.Equal(t, "First quote.", got[0].Text)
	assert.Equal(t, "Second quote.", got[1].Text)
	assert.Equal(t, "Third quote.", got[2].Text)
}
