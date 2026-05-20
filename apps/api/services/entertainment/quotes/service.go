package quotes

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Quote struct {
	ID     int    `json:"id"`
	Text   string `json:"text"`
	Author string `json:"author,omitempty"`
}

type quotesDB interface {
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
}

type Service struct {
	db quotesDB
}

func NewService(pool *pgxpool.Pool) *Service {
	return &Service{db: pool}
}

func (s *Service) Random(ctx context.Context) (Quote, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	row := s.db.QueryRow(ctx, `
SELECT id, text, author
FROM quotes
ORDER BY random()
LIMIT 1;
`)

	var q Quote

	if err := row.Scan(&q.ID, &q.Text, &q.Author); err != nil {
		return Quote{}, fmt.Errorf("scan quote: %w", err)
	}

	return q, nil
}

// RandomBatch returns n random quotes in a single query.
func (s *Service) RandomBatch(ctx context.Context, n int) ([]Quote, error) {
	if n < 1 {
		return nil, fmt.Errorf("count must be at least 1")
	}

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	rows, err := s.db.Query(ctx, `
SELECT id, text, author
FROM quotes
ORDER BY random()
LIMIT $1`, n)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	results := make([]Quote, 0, n)

	for rows.Next() {
		var q Quote
		if err := rows.Scan(&q.ID, &q.Text, &q.Author); err != nil {
			continue
		}

		results = append(results, q)
	}

	return results, rows.Err()
}
