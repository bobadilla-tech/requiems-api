package words

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"requiems-api/platform/svcerr"
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
	var se *svcerr.Error
	assert.ErrorAs(t, err, &se)
	assert.Equal(t, svcerr.KindUpstream, se.Kind)
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
	svc := newTestService(&mockRow{
		scanFn: func(_ ...any) error { return pgx.ErrNoRows },
	})

	_, err := svc.Random(context.Background())
	var se *svcerr.Error
	assert.ErrorAs(t, err, &se)
	assert.Equal(t, svcerr.KindUpstream, se.Kind)
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
func TestBatchDefine_MixedWords(t *testing.T) {
	t.Parallel()

	svc := &Service{}

	req := BatchRequest{
		Items: []string{
			"ephemeral",
			"SERENDIPITY",
			"zzyzx",
		},
	}

	resp, err := svc.BatchDefine(context.Background(), req)
	require.NoError(t, err)

	assert.Equal(t, 3, len(resp))
	require.Len(t, resp, 3)

	assert.True(t, resp[0].Found)
	assert.Equal(t, "ephemeral", resp[0].Word)

	assert.True(t, resp[1].Found)
	assert.Equal(t, "serendipity", resp[1].Word)

	assert.False(t, resp[2].Found)
	assert.Nil(t, resp[2].Entry)
	assert.NotEmpty(t, resp[2].Error)
}

func TestBatchDefine_AllValid(t *testing.T) {
	t.Parallel()

	svc := &Service{}

	req := BatchRequest{
		Items: []string{
			"melancholy",
			"resilience",
			"eloquent",
		},
	}

	resp, err := svc.BatchDefine(context.Background(), req)
	require.NoError(t, err)

	assert.Equal(t, 3, len(resp))
	require.Len(t, resp, 3)

	for i, r := range resp {
		assert.True(t, r.Found, "item %d should be found", i)
		assert.NotNil(t, r.Entry)
		assert.NotEmpty(t, r.Entry.Definitions)
	}
}

func TestBatchDefine_AllInvalid(t *testing.T) {
	t.Parallel()

	svc := &Service{}

	req := BatchRequest{
		Items: []string{"yyy", "zzz"},
	}

	resp, err := svc.BatchDefine(context.Background(), req)
	require.NoError(t, err)

	assert.Equal(t, 2, len(resp))

	for _, r := range resp {
		assert.False(t, r.Found)
		assert.Nil(t, r.Entry)
		assert.NotEmpty(t, r.Error)
	}
}

func TestBatchDefine_TrimsAndNormalizes(t *testing.T) {
	t.Parallel()

	svc := &Service{}

	req := BatchRequest{
		Items: []string{
			"  Ephemeral  ",
			"  SERENDIPITY ",
		},
	}

	resp, err := svc.BatchDefine(context.Background(), req)
	require.NoError(t, err)

	assert.True(t, resp[0].Found)
	assert.Equal(t, "ephemeral", resp[0].Word)

	assert.True(t, resp[1].Found)
	assert.Equal(t, "serendipity", resp[1].Word)
}

func TestBatchDefine_EmptyRequest(t *testing.T) {
	t.Parallel()

	svc := &Service{}

	req := BatchRequest{
		Items: []string{},
	}

	resp, err := svc.BatchDefine(context.Background(), req)
	require.NoError(t, err)

	assert.Equal(t, 0, len(resp))
	assert.Len(t, resp, 0)
}
