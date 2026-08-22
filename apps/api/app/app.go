package app

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"

	"requiems-api/platform/config"
	"requiems-api/platform/db"
	"requiems-api/platform/middleware"
	"requiems-api/platform/reqredis"
)

// apiKeyCacheTTL bounds how long a verified API key stays cached in Redis
// before re-checking Postgres; kept short since it's also the accepted
// window a revoked key can keep authenticating if Rails' cache-invalidating
// DEL fails (see platform/middleware/apikeyauth.go).
const apiKeyCacheTTL = 30 * time.Second

type App struct {
	cfg     config.Config
	handler http.Handler
	pool    *pgxpool.Pool
	rdb     *redis.Client
}

func New(ctx context.Context, cfg config.Config) (*App, error) {
	if err := db.MigrateWithRetry(cfg.DatabaseURL, "migrations"); err != nil {
		return nil, err
	}

	pool, err := db.Connect(ctx, cfg.DatabaseURL)
	if err != nil {
		return nil, err
	}

	rdb, err := reqredis.Connect(ctx, cfg.RedisURL)
	if err != nil {
		return nil, err
	}

	apiKeyAuth := middleware.NewAPIKeyAuth(pool, rdb, apiKeyCacheTTL)

	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.RequestLogger(slog.Default()))

	router.Get("/healthz", Healthz(pool))

	router.Group(func(protected chi.Router) {
		protected.Use(middleware.BackendSecretAuth(cfg.BackendSecret))
		// Enforcing, not shadow — but note what this actually gates today:
		// the Worker (apps/workers/auth-gateway/src/http.ts) strips any
		// incoming requiems-api-key header before proxying to Go, so real
		// Worker-proxied traffic never carries one and will 401 here. This
		// middleware is the live enforcing check only for requests that
		// reach Go directly (local dev, integration tests, future
		// direct-to-Go paths) — see docs/plans/2026-08-21-go-auth-foundation-phase-0-1.md
		// Phase 1 item 6 for why that's an accepted gap, not a bug.
		protected.Use(apiKeyAuth.Middleware())

		protected.Route("/v1", func(v1 chi.Router) {
			registerV1Routes(ctx, v1, pool, rdb, cfg)
		})
	})

	return &App{
		cfg:     cfg,
		handler: router,
		pool:    pool,
		rdb:     rdb,
	}, nil
}

func (a *App) Handler() http.Handler {
	return a.handler
}

// Close releases the database pool and Redis client. Safe to call once,
// after the HTTP server has stopped accepting new requests.
func (a *App) Close() {
	if a.pool != nil {
		a.pool.Close()
	}

	if a.rdb != nil {
		_ = a.rdb.Close()
	}
}
