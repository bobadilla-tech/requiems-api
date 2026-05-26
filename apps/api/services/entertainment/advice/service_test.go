package advice

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
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

	var (
		call int64
		mu   sync.Mutex
	)

	svc := newTestService(&mockRow{
		scanFn: func(dest ...any) error {
			n := int(atomic.AddInt64(&call, 1))

			mu.Lock()
			defer mu.Unlock()

			*dest[0].(*int) = n
			*dest[1].(*string) = fmt.Sprintf("Advice %d", n)

			return nil
		},
	})

	got, err := svc.RandomBatch(context.Background(), 3)

	require.NoError(t, err)
	require.Len(t, got, 3)

	ids := []int{
		got[0].ID,
		got[1].ID,
		got[2].ID,
	}

	texts := []string{
		got[0].Text,
		got[1].Text,
		got[2].Text,
	}

	assert.ElementsMatch(t, []int{1, 2, 3}, ids)
	assert.ElementsMatch(t,
		[]string{"Advice 1", "Advice 2", "Advice 3"},
		texts,
	)
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

	var (
		mu   sync.Mutex
		call int
	)
	scanErr := errors.New("scan failed")

	svc := newTestService(&mockRow{
		scanFn: func(dest ...any) error {
			mu.Lock()
			c := call
			call++
			mu.Unlock()

			if c == 3 {
				return scanErr
			}

			*dest[0].(*int) = c
			*dest[1].(*string) = fmt.Sprintf("Advice %d", c)

			return nil
		},
	})

	got, err := svc.RandomBatch(context.Background(), 5)

	assert.Nil(t, got)
	assert.ErrorIs(t, err, scanErr)
}
