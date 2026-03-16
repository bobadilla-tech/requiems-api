package quotes

import (
	"context"
	"math/rand/v2"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Quote struct {
	ID     int    `json:"id"`
	Text   string `json:"text"`
	Author string `json:"author,omitempty"`
}

func (Quote) IsData() {}

type Service struct {
	db *pgxpool.Pool
}

func NewService(db *pgxpool.Pool) *Service {
	return &Service{db: db}
}

func (s *Service) Random(ctx context.Context) (Quote, error) {
	var count int
	if err := s.db.QueryRow(ctx, `SELECT count(*) FROM quotes`).Scan(&count); err != nil {
		return Quote{}, err
	}

	offset := rand.IntN(count)

	var q Quote
	row := s.db.QueryRow(ctx, `
SELECT id, text, author
FROM quotes
LIMIT 1
OFFSET $1;
`, offset)

	if err := row.Scan(&q.ID, &q.Text, &q.Author); err != nil {
		return Quote{}, err
	}

	return q, nil
}
