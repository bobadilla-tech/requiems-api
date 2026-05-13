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

type querier interface {
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
}

type Service struct {
	db querier
}

func NewService(db *pgxpool.Pool) *Service {
	return &Service{db: db}
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

// RandomBatch returns n random quotes in a single operation.
// If an individual quote fails to scan (e.g. transient DB error), a zero-value
// Quote is used for that slot so the batch always returns exactly n results.
// The batch is aborted only if the context is cancelled or times out.
func (s *Service) RandomBatch(ctx context.Context, n int) ([]Quote, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	results := make([]Quote, n)
	for i := 0; i < n; i++ {
		row := s.db.QueryRow(ctx, `
SELECT id, text, author
FROM quotes
ORDER BY random()
LIMIT 1;
`)
		var q Quote
		if err := row.Scan(&q.ID, &q.Text, &q.Author); err != nil {
			if ctx.Err() != nil {
				return nil, ctx.Err()
			}
			// Individual scan error: keep zero-value Quote and continue.
			continue
		}
		results[i] = q
	}
	return results, nil
}
