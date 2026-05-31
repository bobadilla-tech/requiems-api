package advice

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

// mockRows implements dbRows.
type mockRows struct {
	items [][]any
	pos   int
	err   error
}

func (m *mockRows) Close() {}

func (m *mockRows) Err() error { return m.err }

func (m *mockRows) Next() bool {
	if m.pos < len(m.items) {
		return true
	}
	return false
}

func (m *mockRows) Scan(dest ...any) error {
	row := m.items[m.pos]
	m.pos++
	for i, d := range dest {
		switch p := d.(type) {
		case *int:
			*p = row[i].(int)
		case *string:
			*p = row[i].(string)
		}
	}
	return nil
}

// mockDB implements dbPool.
type mockDB struct {
	row  pgx.Row
	rows *mockRows
}

func (m *mockDB) QueryRow(_ context.Context, _ string, _ ...any) pgx.Row {
	return m.row
}

func (m *mockDB) Query(_ context.Context, _ string, _ ...any) (dbRows, error) {
	if m.rows.err != nil {
		return nil, m.rows.err
	}
	return m.rows, nil
}

func newTestService(row pgx.Row) *Service {
	return &Service{db: &mockDB{
		row:  row,
		rows: &mockRows{},
	}}
}

func newTestServiceWithRows(rows *mockRows) *Service {
	return &Service{db: &mockDB{
		row:  &mockRow{scanFn: func(_ ...any) error { return pgx.ErrNoRows }},
		rows: rows,
	}}
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
			*dest[0].(*int) = 1
			*dest[1].(*string) = "Do one thing every day that scares you."
			return nil
		},
	})

	got, err := svc.Random(context.Background())
	require.NoError(t, err)
	assert.Equal(t, 1, got.ID)
	assert.Equal(t, "Do one thing every day that scares you.", got.Text)
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
}

func TestRandomBatch_EmptyTable(t *testing.T) {
	t.Parallel()
	svc := newTestServiceWithRows(&mockRows{items: [][]any{}})

	got, err := svc.RandomBatch(context.Background(), 3)
	require.NoError(t, err)
	assert.Empty(t, got)
}

func TestRandomBatch_MultipleRows(t *testing.T) {
	t.Parallel()
	svc := newTestServiceWithRows(&mockRows{
		items: [][]any{
			{1, "Advice 1"},
			{2, "Advice 2"},
			{3, "Advice 3"},
		},
	})

	got, err := svc.RandomBatch(context.Background(), 3)
	require.NoError(t, err)
	require.Len(t, got, 3)

	assert.ElementsMatch(t, []int{1, 2, 3}, []int{got[0].ID, got[1].ID, got[2].ID})
	assert.ElementsMatch(t,
		[]string{"Advice 1", "Advice 2", "Advice 3"},
		[]string{got[0].Text, got[1].Text, got[2].Text},
	)
}

func TestRandomBatch_QueryError(t *testing.T) {
	t.Parallel()
	queryErr := errors.New("db unavailable")
	svc := newTestServiceWithRows(&mockRows{err: queryErr})

	_, err := svc.RandomBatch(context.Background(), 3)
	assert.ErrorIs(t, err, queryErr)
}

func TestRandomBatch_ZeroCount(t *testing.T) {
	t.Parallel()
	svc := newTestServiceWithRows(&mockRows{items: [][]any{}})

	got, err := svc.RandomBatch(context.Background(), 0)
	require.NoError(t, err)
	assert.Empty(t, got)
}
