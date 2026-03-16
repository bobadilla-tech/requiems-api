package advice

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Service struct {
	db *pgxpool.Pool
}

func NewService(db *pgxpool.Pool) *Service {
	return &Service{db: db}
}

func (s *Service) Random(ctx context.Context) (Advice, error) {
	var count int
	if err := s.db.QueryRow(ctx, "SELECT COUNT(*) FROM advice").Scan(&count); err != nil {
		return Advice{}, err
	}
	if count == 0 {
		return Advice{}, pgx.ErrNoRows
	}

	var a Advice
	err := s.db.QueryRow(ctx, `
SELECT id, text
FROM advice
ORDER BY id
OFFSET floor(random() * $1)::int
LIMIT 1
`, count).Scan(&a.ID, &a.Text)
	if err != nil {
		return Advice{}, err
	}
	return a, nil
}
