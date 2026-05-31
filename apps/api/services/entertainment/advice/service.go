package advice

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Advice struct {
	ID   int    `json:"id"`
	Text string `json:"advice"`
}

type dbRows interface {
	Close()
	Err() error
	Next() bool
	Scan(dest ...any) error
}

type dbPool interface {
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	Query(ctx context.Context, sql string, args ...any) (dbRows, error)
}

type poolWrapper struct {
	p *pgxpool.Pool
}

func (pw *poolWrapper) QueryRow(ctx context.Context, sql string, args ...any) pgx.Row {
	return pw.p.QueryRow(ctx, sql, args...)
}

func (pw *poolWrapper) Query(ctx context.Context, sql string, args ...any) (dbRows, error) {
	return pw.p.Query(ctx, sql, args...)
}

type Service struct {
	db dbPool
}

func NewService(pool *pgxpool.Pool) *Service {
	return &Service{db: &poolWrapper{p: pool}}
}

func (s *Service) Random(ctx context.Context) (Advice, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	row := s.db.QueryRow(ctx, `
	SELECT id, text
	FROM advice
	ORDER BY random()
	LIMIT 1;
	`)

	var a Advice
	if err := row.Scan(&a.ID, &a.Text); err != nil {
		return Advice{}, fmt.Errorf("scan advice: %w", err)
	}
	return a, nil
}

func (s *Service) RandomBatch(ctx context.Context, count int) ([]Advice, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	rows, err := s.db.Query(ctx, `
		SELECT id, text
		FROM advice
		ORDER BY random()
		LIMIT $1;
	`, count)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	results := make([]Advice, 0, count)
	for rows.Next() {
		var a Advice
		if err := rows.Scan(&a.ID, &a.Text); err != nil {
			return nil, err
		}
		results = append(results, a)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return results, nil
}
