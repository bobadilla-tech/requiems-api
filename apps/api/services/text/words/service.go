package words

import (
	"context"
	"strings"
	"time"

	dictionary "github.com/bobadilla-tech/go-dictionary"
	"github.com/bobadilla-tech/thesaurus-go/pkg/thesaurus"
	"github.com/jackc/pgx/v5/pgxpool"

	"requiems-api/platform/db"
	"requiems-api/platform/svcerr"
)

type Word struct {
	ID           int    `json:"id"`
	Word         string `json:"word"`
	Definition   string `json:"definition"`
	PartOfSpeech string `json:"part_of_speech,omitempty"`
}

// Definition represents a single definition entry for a word.
type Definition struct {
	PartOfSpeech string `json:"part_of_speech"`
	Definition   string `json:"definition"`
	Example      string `json:"example,omitempty"`
}

// DictionaryEntry is the response payload for the dictionary endpoint.
type DictionaryEntry struct {
	Word        string       `json:"word"`
	Phonetic    string       `json:"phonetic,omitempty"`
	Definitions []Definition `json:"definitions"`
	Synonyms    []string     `json:"synonyms"`
}

type BatchItem struct {
	Word  string           `json:"word"`
	Found bool             `json:"found"`
	Entry *DictionaryEntry `json:"entry,omitempty"`
	Error string           `json:"error,omitempty"`
}

type Service struct {
	db db.Querier
}

func NewService(pool *pgxpool.Pool) *Service {
	return &Service{db: pool}
}

func (s *Service) Random(ctx context.Context) (Word, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	row := s.db.QueryRow(ctx, `
SELECT id, word, definition, part_of_speech
FROM words
ORDER BY random()
LIMIT 1;
`)

	var w Word

	if err := row.Scan(&w.ID, &w.Word, &w.Definition, &w.PartOfSpeech); err != nil {
		return Word{}, svcerr.Upstream("service_unavailable", "no words available")
	}

	return w, nil
}

// Define returns the dictionary entry for the given word.
// Returns an error if the word is not found in the dataset.
func (s *Service) Define(word string) (DictionaryEntry, error) {
	normalized := strings.ToLower(strings.TrimSpace(word))

	entry, ok := dictionary.GetCurated(normalized)
	if ok {
		return curatedToDictionaryEntry(normalized, entry), nil
	}

	if entry, ok := dictionary.Get(normalized); ok {
		return wiktionaryToDictionaryEntry(normalized, entry), nil
	}

	return DictionaryEntry{}, svcerr.NotFound("not_found", "word not found in dictionary")
}

// curatedToDictionaryEntry — Converts a CuratedEntry (the ~30-word curated
// dataset) directly to DictionaryEntry. A straightforward 1:1 mapping with
// no decisions to make: copies Phonetic as-is, transforms each CuratedDefinition into
// a Definition, and uses Synonyms from the curated entry as-is (or []string{} if nil).
func curatedToDictionaryEntry(word string, e dictionary.CuratedEntry) DictionaryEntry {

	defs := make([]Definition, 0, len(e.Definitions))
	for _, d := range e.Definitions {
		defs = append(defs, Definition{
			PartOfSpeech: d.PartOfSpeech,
			Definition:   d.Definition,
			Example:      d.Example,
		})
	}

	synonyms := e.Synonyms
	if synonyms == nil {
		synonyms = []string{}
	}

	return DictionaryEntry{
		Word:        word,
		Phonetic:    e.Phonetic,
		Definitions: defs,
		Synonyms:    synonyms,
	}

}

// wiktionaryToDictionaryEntry — Converts a Wiktionary-derived Entry (richer, with multiple etymologies/dialects)
// into DictionaryEntry, flattening it: picks Phonetic with UK→US→Other priority, takes only
// the first Variant (first etymology) and only the first Example
// from each of its definitions, and resolves Synonyms by calling thesaurus.Lookup(word)
// since Wiktionary-derived entries carry no synonyms of their own.
func wiktionaryToDictionaryEntry(word string, e dictionary.Entry) DictionaryEntry {
	phonetic := e.PhoneticUK

	if phonetic == "" {
		phonetic = e.PhoneticUS
	}

	if phonetic == "" {
		phonetic = e.PhoneticOther
	}

	var defs []Definition
	if len(e.Variants) > 0 {
		v := e.Variants[0]
		defs = make([]Definition, 0, len(v.Definitions))
		for _, d := range v.Definitions {
			example := ""
			if len(d.Examples) > 0 {
				example = d.Examples[0]
			}
			defs = append(defs, Definition{
				PartOfSpeech: d.PartOfSpeech,
				Definition:   d.Definition,
				Example:      example,
			})
		}
	}

	if defs == nil {
		defs = []Definition{}
	}

	synonyms := []string{}
	if syn, ok := thesaurus.Lookup(word); ok && syn.Synonyms != nil {
		synonyms = syn.Synonyms
	}

	return DictionaryEntry{
		Word:        word,
		Phonetic:    phonetic,
		Definitions: defs,
		Synonyms:    synonyms,
	}
}

func (s *Service) BatchDefine(ctx context.Context, req BatchRequest) ([]BatchItem, error) {
	results := make([]BatchItem, len(req.Items))

	for i, raw := range req.Items {
		word := strings.ToLower(strings.TrimSpace(raw))

		entry, err := s.Define(word)
		if err != nil {
			results[i] = BatchItem{
				Word:  word,
				Found: false,
				Error: err.Error(),
			}
			continue
		}

		results[i] = BatchItem{
			Word:  word,
			Found: true,
			Entry: &entry,
		}
	}

	return results, nil
}
