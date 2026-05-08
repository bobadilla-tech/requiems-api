package words

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
			*dest[0].(*int) = 42
			*dest[1].(*string) = "ephemeral"
			*dest[2].(*string) = "Lasting for a very short time."
			*dest[3].(*string) = "adjective"
			return nil
		},
	})

	got, err := svc.Random(context.Background())
	require.NoError(t, err)
	assert.Equal(t, 42, got.ID)
	assert.Equal(t, "ephemeral", got.Word)
	assert.Equal(t, "Lasting for a very short time.", got.Definition)
	assert.Equal(t, "adjective", got.PartOfSpeech)
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
	assert.True(t, got.ID == 0 && got.Word == "" && got.Definition == "" && got.PartOfSpeech == "", "expected zero Word on error, got %+v", got)
}

func TestRandom_EmptyPartOfSpeechAllowed(t *testing.T) {
	t.Parallel()
	svc := newTestService(&mockRow{
		scanFn: func(dest ...any) error {
			*dest[0].(*int) = 1
			*dest[1].(*string) = "run"
			*dest[2].(*string) = "Move at a speed faster than walking."
			*dest[3].(*string) = ""
			return nil
		},
	})

	got, err := svc.Random(context.Background())
	require.NoError(t, err)
	assert.Equal(t, "", got.PartOfSpeech)
}
