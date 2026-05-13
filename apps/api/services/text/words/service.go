package words

import (
	"context"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"requiems-api/platform/svcerr"
)

type querier interface {
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
}

type Service struct {
	db querier
}

func NewService(db *pgxpool.Pool) *Service {
	return &Service{db: db}
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

	e, ok := dictionaryData[normalized]
	if !ok {
		return DictionaryEntry{}, svcerr.NotFound("not_found", "word not found in dictionary")
	}

	defs := make([]Definition, 0, len(e.definitions))
	for _, d := range e.definitions {
		defs = append(defs, Definition{
			PartOfSpeech: d.partOfSpeech,
			Definition:   d.definition,
			Example:      d.example,
		})
	}

	synonyms := e.synonyms
	if synonyms == nil {
		synonyms = []string{}
	}

	return DictionaryEntry{
		Word:        normalized,
		Phonetic:    e.phonetic,
		Definitions: defs,
		Synonyms:    synonyms,
	}, nil
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
