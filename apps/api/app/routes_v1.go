package app

import (
	"context"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"

	"requiems-api/platform/config"
	"requiems-api/platform/middleware"
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

// newServiceRouter builds a service router whose 404 and 405 handlers tell
// the usage middleware that no real route ran. They must be set before the
// service registers routes and before it is mounted: chi's Mount only copies
// the parent's handlers onto a sub-router whose own are still nil, so nested
// Route groups inside a service inherit them too.
func newServiceRouter() *chi.Mux {
	r := chi.NewRouter()
	r.NotFound(middleware.NotFoundHandler())
	r.MethodNotAllowed(middleware.MethodNotAllowedHandler())

	return r
}

func registerV1Routes(ctx context.Context, r chi.Router, pool *pgxpool.Pool, rdb *redis.Client, cfg config.Config) {
	// Covers an unknown /v1/x and services that are disabled.
	r.NotFound(middleware.NotFoundHandler())
	r.MethodNotAllowed(middleware.MethodNotAllowedHandler())

	if serviceEnabled(cfg, "entertainment") {
		entertainmentRouter := newServiceRouter()
		entertainment.RegisterRoutes(entertainmentRouter, pool)
		r.Mount("/entertainment", entertainmentRouter)
	}

	if serviceEnabled(cfg, "finance") {
		financeRouter := newServiceRouter()
		finance.RegisterRoutes(financeRouter, pool, rdb)
		r.Mount("/finance", financeRouter)
	}

	if serviceEnabled(cfg, "health") {
		healthRouter := newServiceRouter()
		health.RegisterRoutes(healthRouter, pool)
		r.Mount("/health", healthRouter)
	}

	if serviceEnabled(cfg, "networking") {
		networkingRouter := newServiceRouter()
		networking.RegisterRoutes(networkingRouter, cfg)
		r.Mount("/networking", networkingRouter)
	}

	if serviceEnabled(cfg, "places") {
		placesRouter := newServiceRouter()
		places.RegisterRoutes(placesRouter, cfg, rdb)
		r.Mount("/places", placesRouter)
	}

	if serviceEnabled(cfg, "technology") {
		technologyRouter := newServiceRouter()
		technology.RegisterRoutes(ctx, technologyRouter, pool, rdb)
		r.Mount("/technology", technologyRouter)
	}

	if serviceEnabled(cfg, "text") {
		textRouter := newServiceRouter()
		text.RegisterRoutes(textRouter, pool, cfg)
		r.Mount("/text", textRouter)
	}

	if serviceEnabled(cfg, "validation") {
		validationRouter := newServiceRouter()
		validation.RegisterRoutes(validationRouter)
		r.Mount("/validation", validationRouter)
	}

	if serviceEnabled(cfg, "systems") {
		systemsRouter := newServiceRouter()
		systems.RegisterRoutes(systemsRouter, pool, rdb, cfg)
		r.Mount("/systems", systemsRouter)
	}
}
