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
	plans := middleware.NewPlanCache(pool)
	rateLimiter := middleware.NewRateLimiter(rdb, plans, slog.Default())
	usageQuota := middleware.NewUsageQuota(pool, rdb, plans, slog.Default())

	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.RequestLogger(slog.Default()))

	router.Get("/healthz", Healthz(pool))

	router.Group(func(protected chi.Router) {
		// APIKeyAuth is the sole enforcing auth check for all traffic to this
		// group, direct or Cloudflare-proxied. API-key authentication is the
		// sole enforcing application-layer auth boundary.
		protected.Use(apiKeyAuth.Middleware())
		// Rate limiting and usage/quota tracking both read the principal
		// APIKeyAuth just attached, so they're mounted right after it, in
		// this order: cheap per-minute rate check before the (Postgres
		// fallback-capable, potentially synchronous-write) quota check.
		protected.Use(rateLimiter.Middleware())
		protected.Use(usageQuota.Middleware())

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
