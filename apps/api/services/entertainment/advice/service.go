package advice

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"requiems-api/platform/db"
)

type Advice struct {
	ID   int    `json:"id"`
	Text string `json:"advice"`
}

func (Advice) IsData() {}

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
