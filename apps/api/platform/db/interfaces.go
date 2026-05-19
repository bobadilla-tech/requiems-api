package db

import (
	"context"

	"github.com/jackc/pgx/v5"
)

// Querier is the minimal DB interface for services that only need single-row queries.
// *pgxpool.Pool satisfies this interface directly.
type Querier interface {
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
}
