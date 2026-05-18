package advice

import (
	"context"
	"errors"
	"testing"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// mockRow implements pgx.Row.
type mockRow struct {
	scanFn func(dest ...any) error
}

func (m *mockRow) Scan(dest ...any) error { return m.scanFn(dest...) }

// mockQuerier implements querier.
type mockQuerier struct {
	row pgx.Row
}

func (m *mockQuerier) QueryRow(_ context.Context, _ string, _ ...any) pgx.Row {
	return m.row
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

	svc := newTestService(&mockRow{
		scanFn: func(_ ...any) error { return pgx.ErrNoRows },
	})

	_, err := svc.RandomBatch(context.Background(), 3)
	assert.ErrorIs(t, err, pgx.ErrNoRows)
}

func TestRandomBatch_MultipleRows(t *testing.T) {
	t.Parallel()

	call := 0

	svc := newTestService(&mockRow{
		scanFn: func(dest ...any) error {
			call++

			*dest[0].(*int) = call
			*dest[1].(*string) = fmt.Sprintf("Advice %d", call)

			return nil
		},
	})

	got, err := svc.RandomBatch(context.Background(), 3)

	require.NoError(t, err)
	require.Len(t, got, 3)

	assert.Equal(t, 1, got[0].ID)
	assert.Equal(t, "Advice 1", got[0].Text)

	assert.Equal(t, 2, got[1].ID)
	assert.Equal(t, "Advice 2", got[1].Text)

	assert.Equal(t, 3, got[2].ID)
	assert.Equal(t, "Advice 3", got[2].Text)
}

func TestRandomBatch_ScanError(t *testing.T) {
	t.Parallel()

	scanErr := errors.New("scan failed")

	svc := newTestService(&mockRow{
		scanFn: func(_ ...any) error {
			return scanErr
		},
	})

	_, err := svc.RandomBatch(context.Background(), 3)

	assert.ErrorIs(t, err, scanErr)
}

func TestRandomBatch_ZeroCount(t *testing.T) {
	t.Parallel()

	svc := newTestService(&mockRow{
		scanFn: func(_ ...any) error { return nil },
	})

	got, err := svc.RandomBatch(context.Background(), 0)

	require.NoError(t, err)
	assert.Empty(t, got)
}

func TestRandomBatch_ReturnsPartialResultsOnFailure(t *testing.T) {
	t.Parallel()

	call := 0
	scanErr := errors.New("scan failed")

	svc := newTestService(&mockRow{
		scanFn: func(dest ...any) error {
			call++

			if call == 3 {
				return scanErr
			}

			*dest[0].(*int) = call
			*dest[1].(*string) = fmt.Sprintf("Advice %d", call)

			return nil
		},
	})

	got, err := svc.RandomBatch(context.Background(), 5)

	assert.Nil(t, got)
	assert.ErrorIs(t, err, scanErr)
}
