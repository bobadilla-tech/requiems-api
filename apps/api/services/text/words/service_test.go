package words

import (
	"context"
	"testing"

	dictionarydata "github.com/bobadilla-tech/go-dictionary"
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

func TestCuratedToDictionaryEntry_MapsFieldsAsIs(t *testing.T) {
	e := dictionarydata.CuratedEntry{
		Phonetic: "/ɪˈfɛm(ə)rəl/",
		Definitions: []dictionarydata.CuratedDefinition{
			{PartOfSpeech: "adjective", Definition: "lasting a very short time", Example: "ephemeral pleasures"},
			{PartOfSpeech: "noun", Definition: "a plant that lives for a short time", Example: ""},
		},
		Synonyms: []string{"transient", "fleeting"},
	}

	got := curatedToDictionaryEntry("ephemeral", e)

	assert.Equal(t, "ephemeral", got.Word)
	assert.Equal(t, "/ɪˈfɛm(ə)rəl/", got.Phonetic)
	require.Len(t, got.Definitions, 2)
	assert.Equal(t, Definition{
		PartOfSpeech: "adjective",
		Definition:   "lasting a very short time",
		Example:      "ephemeral pleasures",
	}, got.Definitions[0])
	assert.Equal(t, []string{"transient", "fleeting"}, got.Synonyms)
}

func TestCuratedToDictionaryEntry_NilSynonymsBecomeEmptySlice(t *testing.T) {
	e := dictionarydata.CuratedEntry{
		Phonetic:    "/test/",
		Definitions: []dictionarydata.CuratedDefinition{{PartOfSpeech: "noun", Definition: "a test"}},
		Synonyms:    nil,
	}

	got := curatedToDictionaryEntry("test", e)

	assert.NotNil(t, got.Synonyms, "Synonyms must be [] not nil, for stable JSON encoding")
	assert.Empty(t, got.Synonyms)
}

// ---- wiktionaryToDictionaryEntry: phonetic priority ----

func TestWiktionaryToDictionaryEntry_PhoneticPrefersUK(t *testing.T) {
	e := dictionarydata.Entry{
		Word:          "name",
		PhoneticUK:    "/neɪm-uk/",
		PhoneticUS:    "/neɪm-us/",
		PhoneticOther: "/neɪm-other/",
		Variants:      []dictionarydata.Variant{{Definitions: []dictionarydata.Definition{{PartOfSpeech: "noun", Definition: "an identifier"}}}},
	}

	got := wiktionaryToDictionaryEntry("name", e)

	assert.Equal(t, "/neɪm-uk/", got.Phonetic)
}

func TestWiktionaryToDictionaryEntry_PhoneticFallsBackToUS(t *testing.T) {
	e := dictionarydata.Entry{
		Word:          "name",
		PhoneticUK:    "",
		PhoneticUS:    "/neɪm-us/",
		PhoneticOther: "/neɪm-other/",
		Variants:      []dictionarydata.Variant{{Definitions: []dictionarydata.Definition{{PartOfSpeech: "noun", Definition: "an identifier"}}}},
	}

	got := wiktionaryToDictionaryEntry("name", e)

	assert.Equal(t, "/neɪm-us/", got.Phonetic)
}

func TestWiktionaryToDictionaryEntry_PhoneticFallsBackToOther(t *testing.T) {
	e := dictionarydata.Entry{
		Word:          "name",
		PhoneticUK:    "",
		PhoneticUS:    "",
		PhoneticOther: "/neɪm-other/",
		Variants:      []dictionarydata.Variant{{Definitions: []dictionarydata.Definition{{PartOfSpeech: "noun", Definition: "an identifier"}}}},
	}

	got := wiktionaryToDictionaryEntry("name", e)

	assert.Equal(t, "/neɪm-other/", got.Phonetic)
}

func TestWiktionaryToDictionaryEntry_PhoneticEmptyWhenNoneSet(t *testing.T) {
	e := dictionarydata.Entry{
		Word:     "name",
		Variants: []dictionarydata.Variant{{Definitions: []dictionarydata.Definition{{PartOfSpeech: "noun", Definition: "an identifier"}}}},
	}

	got := wiktionaryToDictionaryEntry("name", e)

	assert.Empty(t, got.Phonetic)
}

// ---- wiktionaryToDictionaryEntry: Variant selection ----

func TestWiktionaryToDictionaryEntry_UsesOnlyFirstVariant(t *testing.T) {
	// "name" the identifier vs. "name" the Caribbean yam — two unrelated
	// etymologies. Only Variants[0]'s definitions should surface.
	e := dictionarydata.Entry{
		Word: "name",
		Variants: []dictionarydata.Variant{
			{
				Etymology:   "From Middle English name, from Old English nama.",
				Definitions: []dictionarydata.Definition{{PartOfSpeech: "noun", Definition: "an identifier"}},
			},
			{
				Etymology:   "Borrowed from Spanish ñame.",
				Definitions: []dictionarydata.Definition{{PartOfSpeech: "noun", Definition: "the Caribbean yam"}},
			},
		},
	}

	got := wiktionaryToDictionaryEntry("name", e)

	require.Len(t, got.Definitions, 1)
	assert.Equal(t, "an identifier", got.Definitions[0].Definition,
		"only the first Variant's definitions should be surfaced")
}

func TestWiktionaryToDictionaryEntry_NoVariantsReturnsEmptyNotNilDefinitions(t *testing.T) {
	e := dictionarydata.Entry{Word: "orphan", Variants: nil}

	got := wiktionaryToDictionaryEntry("orphan", e)

	assert.NotNil(t, got.Definitions, "Definitions must be [] not nil, for stable JSON encoding")
	assert.Empty(t, got.Definitions)
}

// ---- wiktionaryToDictionaryEntry: Examples[] -> Example ----

func TestWiktionaryToDictionaryEntry_TakesFirstExampleOnly(t *testing.T) {
	e := dictionarydata.Entry{
		Word: "run",
		Variants: []dictionarydata.Variant{{
			Definitions: []dictionarydata.Definition{{
				PartOfSpeech: "verb",
				Definition:   "to move fast on foot",
				Examples:     []string{"She ran to the store.", "He runs every morning.", "They ran a marathon."},
			}},
		}},
	}

	got := wiktionaryToDictionaryEntry("run", e)

	require.Len(t, got.Definitions, 1)
	assert.Equal(t, "She ran to the store.", got.Definitions[0].Example)
}

func TestWiktionaryToDictionaryEntry_EmptyExampleWhenNoneProvided(t *testing.T) {
	e := dictionarydata.Entry{
		Word: "run",
		Variants: []dictionarydata.Variant{{
			Definitions: []dictionarydata.Definition{{
				PartOfSpeech: "verb",
				Definition:   "to move fast on foot",
				Examples:     nil,
			}},
		}},
	}

	got := wiktionaryToDictionaryEntry("run", e)

	require.Len(t, got.Definitions, 1)
	assert.Empty(t, got.Definitions[0].Example)
}

func TestWiktionaryToDictionaryEntry_MultipleDefinitionsAllMapped(t *testing.T) {
	e := dictionarydata.Entry{
		Word: "head",
		Variants: []dictionarydata.Variant{{
			Definitions: []dictionarydata.Definition{
				{PartOfSpeech: "noun", Definition: "the topmost part of the body", Examples: []string{"He nodded his head."}},
				{PartOfSpeech: "noun", Definition: "a leader", Examples: []string{"the head of the department"}},
				{PartOfSpeech: "verb", Definition: "to lead", Examples: nil},
			},
		}},
	}

	got := wiktionaryToDictionaryEntry("head", e)

	require.Len(t, got.Definitions, 3)
	assert.Equal(t, "noun", got.Definitions[0].PartOfSpeech)
	assert.Equal(t, "He nodded his head.", got.Definitions[0].Example)
	assert.Equal(t, "verb", got.Definitions[2].PartOfSpeech)
	assert.Empty(t, got.Definitions[2].Example)
}
