# Copilot Instructions

## Project Overview

Multi-app monorepo with four services:

| App             | Path                           | Language                       |
| --------------- | ------------------------------ | ------------------------------ |
| Go API          | `apps/api/`                    | Go 1.26                        |
| Rails Dashboard | `apps/dashboard/`              | Ruby 3.4.8 / Rails 8           |
| Auth Gateway    | `apps/workers/auth-gateway/`   | TypeScript (Cloudflare Worker) |
| API Management  | `apps/workers/api-management/` | TypeScript (Cloudflare Worker) |

Shared worker utilities live in `apps/workers/shared/`.

## Architecture

```
User → Auth Gateway (port 4455) → Go API (port 8080) → PostgreSQL
                                         ↓
Rails Dashboard → API Management (port 5544) → Cloudflare KV/D1
```

## Go API — Code Patterns

Domain-driven design under `services/`. Each feature follows:

```
services/{domain}/{feature}/
├── type.go            # Request/response types
├── service.go         # Business logic, NewService() constructor
└── transport_http.go  # HTTP handlers, RegisterRoutes(r chi.Router, svc *Service)
```

Routes are wired in `services/{domain}/router.go` and mounted in
`app/routes_v1.go`.

Use `httpx.JSON()` for success responses and `httpx.Error()` for errors.

## Validating Your Work

Your environment has all tools pre-installed. After making changes, run the
checks for each app you modified.

### Go API (`apps/api/`)

```bash
cd apps/api

# Must pass
go test -race ./...

# Advisory (fix if you can, but not a blocker)
golangci-lint run

# Docker dev container equivalent
docker exec requiem-dev-api-1 golangci-lint run
```

Environment required for tests:

```
DATABASE_URL=postgres://requiem:requiem@localhost:5432/requiem_test?sslmode=disable
BACKEND_SECRET=test_secret_min_32_chars_long_for_testing_only
```

### Rails Dashboard (`apps/dashboard/`)

```bash
cd apps/dashboard

# Must pass
RAILS_ENV=test DATABASE_URL=postgres://requiem:requiem@localhost:5432/requiem_test?sslmode=disable REDIS_URL=redis://localhost:6379 BACKEND_SECRET=test_secret_min_32_chars_long_for_testing_only bin/rails test

# Must pass
bundle exec bundler-audit
bin/importmap audit
bundle exec brakeman --no-pager

# Advisory
bundle exec rubocop
```

### Auth Gateway (`apps/workers/auth-gateway/`)

```bash
cd apps/workers/auth-gateway

# Must pass
pnpm exec vitest run
pnpm run typecheck

# Advisory
pnpm run lint
pnpm run format:check
```

### API Management (`apps/workers/api-management/`)

```bash
cd apps/workers/api-management

# Must pass
pnpm exec vitest run
pnpm run typecheck
```

## Rails Dashboard — i18n (rori18n)

`apps/dashboard` uses [rori18n](https://github.com/bobadilla-tech/rori18n) — a Go CLI that manages i18n YAML files. Install once:

```bash
go install github.com/bobadilla-tech/rori18n@latest
```

Common tasks (run from repo root):

```bash
# Lint: fail if any t() call has no YAML key
rori18n lint --root apps/dashboard

# Add a key across all locales
rori18n add-key --root apps/dashboard --key shared.buttons.save --value "Save changes"

# Rename a key everywhere (YAML + source callers)
rori18n refactor-key --root apps/dashboard --old shared.old.key --new shared.new.key

# Find orphaned / missing keys
rori18n audit --root apps/dashboard --all
```

Brand names protected from translation live in `apps/dashboard/.translate-dictionary.txt`.
Do not use `i18n-tasks` write commands — use rori18n instead. `bundle exec i18n-tasks health` is still useful as a passive check.

## Before Marking a Task Done

1. Run the "must pass" checks for every app you touched.
2. Fix any failures before finishing — do not leave broken tests.
3. Advisory checks (lint, rubocop) are nice-to-fix but not required.
