## Backend (Go) Overview

### Goals

- Single monolithic service supporting 500+ endpoints.
- Feature-oriented structure.

### Project layout

```
apps/api/
├── main.go               # Loads config, builds app, starts HTTP server
├── migrations/           # SQL migration files (auto-run on startup)
├── app/
│   ├── app.go            # Router setup, runs migrations with retry
│   ├── routes_v1.go      # Mounts all domain routers under /v1
│   └── healthz.go        # GET /healthz
├── platform/
│   ├── config/           # Config struct, reads env vars
│   ├── db/               # PostgreSQL pool + migrations runner
│   ├── reqredis/         # Redis client
│   └── httpx/            # JSON() and Error() response helpers
└── services/
    ├── {domain}/
    │   ├── router.go     # Instantiates features and registers routes
    │   └── {feature}/
    │       ├── type.go           # Request/response types
    │       ├── service.go        # Business logic
    │       └── transport_http.go # HTTP handlers, RegisterRoutes()
    └── ...
```

Domains: `email`, `text`, `tech`, `places`, `entertainment`, `misc`, etc...

### Adding a new feature

1. **Migrations** (if needed) — add a pair in `migrations/`:
   - `000X_feature_name.up.sql`
   - `000X_feature_name.down.sql`

2. **Feature package** — create `services/<domain>/<feature>/`:
   - `type.go` — request/response structs
   - `service.go` — business logic with `NewService(pool)` constructor
   - `transport_http.go` — `func RegisterRoutes(r chi.Router, svc *Service)`

3. **Wire up** in `services/<domain>/router.go`:

   ```go
   svc := feature.NewService(pool)
   feature.RegisterRoutes(r, svc)
   ```

4. **New top-level domain** — create the domain router and mount it in
   `app/routes_v1.go`:
   ```go
   r.Mount("/v1/newdomain", newdomain.NewRouter(pool))
   ```

Features that need Redis pass `*redis.Client` into the service. Background
workers (e.g. sync jobs) are started via `go worker.Start(ctx, ...)` from the
domain router.
