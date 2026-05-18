package advice

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
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

	results := make([]Advice, 0, count)

	for i := 0; i < count; i++ {
		row := s.db.QueryRow(ctx, `
			SELECT id, text
			FROM advice
			ORDER BY random()
			LIMIT 1;
		`)

		var a Advice

		if err := row.Scan(&a.ID, &a.Text); err != nil {
			return nil, fmt.Errorf("scan advice: %w", err)
		}

		results = append(results, a)
	}

	return results, nil
}