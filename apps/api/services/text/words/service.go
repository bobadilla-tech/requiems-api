package words

import (
	"context"
	"math/rand/v2"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Word struct {
	ID           int    `json:"id"`
	Word         string `json:"word"`
	Definition   string `json:"definition"`
	PartOfSpeech string `json:"part_of_speech,omitempty"`
}

func (Word) IsData() {}

type Service struct {
	db *pgxpool.Pool
}

func NewService(db *pgxpool.Pool) *Service {
	return &Service{db: db}
}

func (s *Service) Random(ctx context.Context) (Word, error) {
	var count int
	if err := s.db.QueryRow(ctx, `SELECT COUNT(*) FROM words`).Scan(&count); err != nil {
		return Word{}, err
	}

	offset := rand.IntN(count)

	row := s.db.QueryRow(ctx, `
SELECT id, word, definition, part_of_speech
FROM words
OFFSET $1
LIMIT 1;
`, offset)

	var w Word
	if err := row.Scan(&w.ID, &w.Word, &w.Definition, &w.PartOfSpeech); err != nil {
		return Word{}, err
	}

	return w, nil
}
