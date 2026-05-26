package app

import (
	"context"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"

	"requiems-api/platform/config"
	"requiems-api/services/entertainment"
	"requiems-api/services/finance"
	"requiems-api/services/health"
	"requiems-api/services/networking"
	"requiems-api/services/places"
	"requiems-api/services/systems"
	"requiems-api/services/technology"
	"requiems-api/services/text"
	"requiems-api/services/validation"
)

// serviceEnabled reports whether a named service should be mounted.
// When EnabledServices is empty (the default), all services are mounted —
// preserving the existing shared-deployment behaviour unchanged.
// For private deployments, set ENABLED_SERVICES="email,text,tech" to mount
// only the services the tenant purchased.
func serviceEnabled(cfg config.Config, key string) bool {
	if strings.TrimSpace(cfg.EnabledServices) == "" {
		return true
	}

	for s := range strings.SplitSeq(cfg.EnabledServices, ",") {
		if strings.TrimSpace(s) == key {
			return true
		}
	}

	return false
}

func registerV1Routes(ctx context.Context, r chi.Router, pool *pgxpool.Pool, rdb *redis.Client, cfg config.Config) {
	if serviceEnabled(cfg, "entertainment") {
		entertainmentRouter := chi.NewRouter()
		entertainment.RegisterRoutes(entertainmentRouter, pool)
		r.Mount("/entertainment", entertainmentRouter)
	}

	if serviceEnabled(cfg, "finance") {
		financeRouter := chi.NewRouter()
		finance.RegisterRoutes(financeRouter, pool, rdb)
		r.Mount("/finance", financeRouter)
	}

	if serviceEnabled(cfg, "health") {
		healthRouter := chi.NewRouter()
		health.RegisterRoutes(healthRouter, pool)
		r.Mount("/health", healthRouter)
	}

	if serviceEnabled(cfg, "networking") {
		networkingRouter := chi.NewRouter()
		networking.RegisterRoutes(networkingRouter, cfg)
		r.Mount("/networking", networkingRouter)
	}

	if serviceEnabled(cfg, "places") {
		placesRouter := chi.NewRouter()
		places.RegisterRoutes(placesRouter, cfg, rdb)
		r.Mount("/places", placesRouter)
	}

	if serviceEnabled(cfg, "technology") {
		technologyRouter := chi.NewRouter()
		technology.RegisterRoutes(ctx, technologyRouter, pool, rdb)
		r.Mount("/technology", technologyRouter)
	}

	if serviceEnabled(cfg, "text") {
		textRouter := chi.NewRouter()
		text.RegisterRoutes(textRouter, pool, cfg)
		r.Mount("/text", textRouter)
	}

	if serviceEnabled(cfg, "validation") {
		validationRouter := chi.NewRouter()
		validation.RegisterRoutes(validationRouter)
		r.Mount("/validation", validationRouter)
	}

	if serviceEnabled(cfg, "systems") {
		systemsRouter := chi.NewRouter()
		systems.RegisterRoutes(systemsRouter, pool, rdb, cfg)
		r.Mount("/systems", systemsRouter)
	}
}
