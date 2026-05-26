package advice

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"requiems-api/platform/db"
)

type Service struct {
	db db.Querier
}

func NewService(pool *pgxpool.Pool) *Service {
	return &Service{db: pool}
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
	if count < 0 {
		return nil, fmt.Errorf("count must be >= 0")
	}
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	results := make([]Advice, count)

	var (
		wg    sync.WaitGroup
		errCh = make(chan error, count)
	)

	for i := 0; i < count; i++ {
		wg.Add(1)

		go func(idx int) {
			defer wg.Done()

			row := s.db.QueryRow(ctx, `
				SELECT id, text
				FROM advice
				ORDER BY random()
				LIMIT 1;
			`)

			var a Advice

			if err := row.Scan(&a.ID, &a.Text); err != nil {
				errCh <- fmt.Errorf("scan advice: %w", err)
				return
			}

			results[idx] = a
		}(i)
	}

	wg.Wait()
	close(errCh)

	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	return results, nil
}
