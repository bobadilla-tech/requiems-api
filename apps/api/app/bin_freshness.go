package app

import (
	"context"
	"log/slog"
	"time"

	"github.com/jackc/pgx/v5"
)

const binDataStaleAfter = 6 * 30 * 24 * time.Hour // ~6 months

// Querier is the subset of pgxpool.Pool used by checkBINDataFreshness,
// narrowed so it can be faked in tests without a real Postgres connection.
type Querier interface {
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
}

func checkBINDataFreshness(ctx context.Context, q Querier, logger *slog.Logger) {
	var lastUpdated *time.Time
	err := q.QueryRow(ctx, `SELECT MAX(last_updated) FROM bin_data`).Scan(&lastUpdated)
	if err != nil {
		logger.Warn("bin_data freshness check failed", "error", err)
		return
	}
	if lastUpdated == nil {
		logger.Warn("bin_data table is empty — no BIN data loaded")
		return
	}
	if age := time.Since(*lastUpdated); age > binDataStaleAfter {
		logger.Warn("bin_data is stale",
			"last_updated", lastUpdated.Format(time.RFC3339),
			"age_days", int(age.Hours()/24))
	}
}
