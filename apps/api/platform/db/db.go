package db

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Pool sizing is fixed rather than measured, since there is no production
// traffic yet to size against (a single Go replica against Postgres'
// max_connections, leaving headroom for Rails' own pool on the same
// instance). Revisit once real traffic exists.
const (
	maxPoolConns = 20
	minPoolConns = 5
)

func Connect(ctx context.Context, dsn string) (*pgxpool.Pool, error) {
	cfg, err := pgxpool.ParseConfig(dsn)

	if err != nil {
		return nil, fmt.Errorf("parse DATABASE_URL: %w", err)
	}

	cfg.MaxConns = maxPoolConns
	cfg.MinConns = minPoolConns

	pool, err := pgxpool.NewWithConfig(ctx, cfg)

	if err != nil {
		return nil, fmt.Errorf("connect database: %w", err)
	}

	pingCtx, cancel := context.WithTimeout(ctx, 5*time.Second)

	defer cancel()

	if err := pool.Ping(pingCtx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("ping database: %w", err)
	}

	return pool, nil
}
