# Project Overview

Requiems API is a production-ready API platform providing unified access to
multiple enterprise-grade APIs. Built as a multi-language monorepo with:

- **Go API** (apps/api) - Internal business logic backend
- **Rails Dashboard** (apps/dashboard) - Public web UI, user management, admin
  panel
- **Auth Gateway** (apps/workers/auth-gateway) - Public edge gateway for auth,
  rate limiting, and request proxying
- **API Management** (apps/workers/api-management) - Internal service for API
  key management, usage exports, and analytics

## Before Implementing — Read the Relevant Doc First

Before writing any code, read the doc for the area you are working in:

| Task                                            | Read first                                                |
| ----------------------------------------------- | --------------------------------------------------------- |
| Add a Go endpoint (new feature, new domain)     | `docs/core/adding-go-endpoints.md`                        |
| Work on the Go backend (architecture, patterns) | `docs/core/backend.md`                                    |
| Work on the Auth Gateway worker                 | `docs/core/auth-gateway.md`                               |
| Work on the API Management worker               | `docs/core/api-management.md`                             |
| Work on the Rails dashboard                     | `docs/core/rails-app.md`                                  |
| Infrastructure / deployment                     | `docs/core/infrastructure.md` / `docs/core/deployment.md` |
| Dev environment setup                           | `docs/core/getting-started.md`                            |

All docs are in `docs/`, see `docs/core/readme.md` for the full index.

## Development Commands

### Full Stack Development

See [docs/core/getting-started.md](docs/core/getting-started.md) for environment
setup and local development configuration.

Start all services with hot reload:

```bash
docker compose -f docker-compose.dev.yml up
```

Services:

- Go API: <http://localhost:8080/healthz> (Air hot reload)
- Rails Dashboard: <http://localhost:3000> (Rails hot reload)
- Auth Gateway: <http://localhost:4455/healthz> (public API entry point)
- API Management: <http://localhost:5544/healthz> (internal management service)
- PostgreSQL: localhost:5432 (user: `requiem`, password: `requiem`)
- Redis: localhost:6379

### Go API (apps/api)

**Note:** Commands assume containers are running. Container: `requiem-dev-api-1`

Run tests:

```bash
docker exec requiem-dev-api-1 go test ./...
```

Run tests with coverage:

```bash
docker exec requiem-dev-api-1 go test -race -coverprofile=coverage.out ./...
```

Run specific test:

```bash
docker exec requiem-dev-api-1 go test ./services/text/advice -v -run TestGetAdvice
```

Run lint:

```bash
docker exec requiem-dev-api-1 golangci-lint run
```

Regenerate the OpenAPI spec (`docs/swagger.json`), from `apps/api`:

```bash
swag init -g main.go --outputTypes json --parseDependency   # regenerate docs/swagger.json
swag fmt -g main.go                                          # normalize annotation formatting
```

The spec is generated from swagger annotations on the handler functions in
`services/**/transport_http.go` (see the `// handleX godoc` blocks). `swag` is a
global binary; run these commands locally, not in the container. After changing
annotations, regenerate and commit the new `docs/swagger.json`. The drift-guard
test `app/spec_coverage_test.go` fails if a registered route is undocumented.

### Rails Dashboard (apps/dashboard)

**Note:** Commands assume containers are running. Container:
`requiem-dev-dashboard-1`

Rails console:

```bash
docker exec -it requiem-dev-dashboard-1 rails console
```

Run tests:

```bash
docker exec requiem-dev-dashboard-1 bin/rails test
```

Run specific test:

```bash
docker exec requiem-dev-dashboard-1 bin/rails test test/models/user_test.rb
```

Run security scans:

```bash
docker exec requiem-dev-dashboard-1 bundle exec brakeman --no-pager
docker exec requiem-dev-dashboard-1 bundle exec bundler-audit
docker exec requiem-dev-dashboard-1 bin/importmap audit
```

Run linting:

```bash
docker exec requiem-dev-dashboard-1 bundle exec rubocop
```

Migrations:

```bash
docker exec requiem-dev-dashboard-1 bin/rails db:migrate
docker exec requiem-dev-dashboard-1 bin/rails db:rollback
```

Generate migration:

```bash
docker exec requiem-dev-dashboard-1 bin/rails generate migration AddFieldToTable
```

### Auth Gateway (apps/workers/auth-gateway)

Public-facing edge gateway for request authentication and proxying.

**Note:** Commands assume containers are running. Container:
`requiem-dev-auth-gateway-1`

Type check:

```bash
docker exec -e CI=true requiem-dev-auth-gateway-1 pnpm run typecheck
```

Run tests:

```bash
docker exec requiem-dev-auth-gateway-1 pnpm exec vitest run              # Run all tests
docker exec requiem-dev-auth-gateway-1 pnpm exec vitest run --coverage   # With coverage report
```

Run lint/format locally (not in Docker):

```bash
cd apps/workers/auth-gateway
pnpm run lint          # Lint code
pnpm run lint:fix      # Auto-fix lint issues
pnpm run format:check  # Check formatting
pnpm run format        # Auto-format code
```

### API Management (apps/workers/api-management)

Internal service for API key management, usage exports, and analytics.

**Note:** Commands assume containers are running. Container:
`requiem-dev-api-management-1`

Type check:

```bash
docker exec -e CI=true requiem-dev-api-management-1 pnpm run typecheck
```

Run tests:

```bash
docker exec requiem-dev-api-management-1 pnpm exec vitest run              # Run all tests
docker exec requiem-dev-api-management-1 pnpm exec vitest run --coverage   # With coverage report
```

### Local Testing Before Push

Run these commands locally to catch issues before CI.

**Note:** Containers must be running. Use
`docker compose -f docker-compose.dev.yml up` in `infra/docker` first.

```bash
# Go Backend
docker exec requiem-dev-api-1 go test ./...                    # Tests (must pass)
docker exec requiem-dev-api-1 golangci-lint run  # Linting (advisory)

# Rails Dashboard
docker exec requiem-dev-dashboard-1 bin/rails test             # Tests (must pass)
docker exec requiem-dev-dashboard-1 bundle exec brakeman --no-pager  # Security (must pass)
docker exec requiem-dev-dashboard-1 bundle exec bundler-audit  # Dependency audit (must pass)
docker exec requiem-dev-dashboard-1 bin/importmap audit        # JS audit (must pass)
docker exec requiem-dev-dashboard-1 bundle exec rubocop        # Linting (advisory)

# Auth Gateway
docker exec requiem-dev-auth-gateway-1 pnpm exec vitest run       # Tests (must pass - 71 tests)
docker exec -e CI=true requiem-dev-auth-gateway-1 pnpm run typecheck  # TypeScript (must pass)
```

Run lint/format locally for workers:

```bash
cd apps/workers/auth-gateway && pnpm run lint && pnpm run format:check  # Auth Gateway (advisory)
cd apps/workers/api-management && pnpm run lint && pnpm run format:check # API Management (advisory)

# API Management tests (must pass)
docker exec requiem-dev-api-management-1 pnpm exec vitest run
docker exec -e CI=true requiem-dev-api-management-1 pnpm run typecheck
```

## Architecture Overview

### Request Flow

```
Public API Requests:
User → Auth Gateway (api.requiems.xyz) → Go Backend → PostgreSQL
       ↓                                   ↓
       KV (auth, rate limits)             Business logic
       D1 (usage recording)

Internal Management:
Rails Dashboard → API Management (api-management.requiems.xyz) → KV + D1
                 ↓
                 API key CRUD, usage exports, analytics
```

1. **Auth Gateway** (`apps/workers/auth-gateway/` - Port 4455):
   - **Public-facing** service at api.requiems.xyz
   - Validates API keys from Cloudflare KV
   - Checks per-minute rate limits (KV counters)
   - Checks monthly quota limits (D1 queries)
   - Records usage to D1 (asynchronous)
   - Proxies requests to Go backend with `X-Backend-Secret` header
   - Adds usage headers to responses
   - **Authentication:** `requiems-api-key` header from end users

2. **API Management** (`apps/workers/api-management/` - Port 5544):
   - **Internal-only** service at api-management.requiems.xyz
   - API key management (create, revoke, update in KV + D1)
   - Usage data export for Rails sync (D1 → PostgreSQL)
   - Analytics queries (by endpoint, by date, summary)
   - Swagger documentation at `/docs` (basic auth protected in production)
   - **Authentication:** `X-API-Management-Key` header (only Rails has this)

3. **Go Backend** (`apps/api/`):
   - Receives requests from auth-gateway only
   - Executes business logic
   - Queries PostgreSQL for data
   - Returns JSON responses
   - **No authentication** - trusts the gateway

4. **Rails Dashboard** (`apps/dashboard/`):
   - User registration/login
   - Subscription management
   - API key management (via API Management service)
   - Usage statistics and analytics
   - Admin panel
   - Background jobs (D1 sync, aggregation)

### Code Organization

#### Go API (apps/api/)

Domain-driven design with feature modules:

- `app/` - Application initialization, routing
- `platform/` - Shared infrastructure
  - `config/` - Environment configuration
  - `db/` - PostgreSQL connection, migrations
  - `httpx/` - HTTP utilities
  - `middleware/` - Auth middleware
  - `reqredis/` - Redis connection
- `services/` - Self-contained business domain modules, one directory per API
  category
  - `entertainment/` - Jokes, trivia, facts, horoscope, sudoku, emoji, quotes,
    advice (`/v1/entertainment/*`)
  - `finance/` - Exchange rates, crypto, BIN, IBAN, SWIFT, mortgage,
    commodities, inflation (`/v1/finance/*`)
  - `health/` - Fitness exercises (`/v1/health/*`)
  - `networking/` - IP geolocation/ASN/VPN, WHOIS, domain info, MX lookup,
    disposable email (`/v1/networking/*`)
  - `places/` - Timezone, working days, holidays, geocoding, postal codes,
    cities (`/v1/places/*`)
  - `technology/` - QR/barcode, password, user agent, counter, random user,
    base64, number base, color, unit/format conversion, markdown
    (`/v1/technology/*`)
  - `text/` - Words/dictionary, lorem ipsum, spell check, thesaurus, sentiment,
    language detection, text similarity, email normalize (`/v1/text/*`)
  - `validation/` - Email validation, phone validation, profanity filter
    (`/v1/validation/*`)

  Each subdomain follows the same pattern: `service.go` (business logic),
  `transport_http.go` (HTTP handlers), `type.go` (data types), and parent
  `router.go` (routes).

**Pattern**: Each feature has `service.go` (business logic), `transport_http.go`
(HTTP handlers), `type.go` (data types), and parent `router.go` (routes).

#### Cloudflare Workers (apps/workers/)

**Shared Package** (`apps/workers/shared/`):

Common utilities and types used by all Cloudflare Workers:

- `types.ts` - Core types (PlanName, ApiKeyData, RateLimitResult, etc.)
- `config.ts` - Plan definitions and endpoint multipliers
- `logger.ts` - Structured logging with cf-ray tracing
- `http.ts` - HTTP utilities (jsonResponse, jsonError, CORS)

Workers import from shared via `@requiem/workers-shared` path alias.

**Auth Gateway** (`apps/workers/auth-gateway/src/`):

Lean public gateway focused on request flow:

- `index.ts` - Main request handler (auth validation, rate limiting, proxying)
- `env.ts` - Worker-specific environment (BACKEND_URL, BACKEND_SECRET)
- `requests.ts` - Usage tracking (D1 queries and recording)
- `rate-limit.ts` - Rate limiting logic (KV counters)
- `http.ts` - Backend-specific HTTP utilities (filterHeaders, addUsageHeaders)

**API Management** (`apps/workers/api-management/src/`):

Internal management service with Hono framework:

- `index.ts` - Hono app with Swagger documentation
- `shared/env.ts` - Worker-specific environment (API_MANAGEMENT_API_KEY)
- `shared/api-key-generator.ts` - API key generation utilities
- `routes/api-keys/` - API key CRUD operations
- `routes/usage/` - Usage data export with pagination
- `routes/analytics/` - Analytics queries (endpoint, date, summary)
- `middleware/` - API key auth, basic auth, error handling

#### Rails Dashboard (apps/dashboard/)

Standard Rails 8 structure with:

- Turbo + Stimulus for interactivity
- Tailwind CSS for styling
- Sidekiq for background jobs
- Solid Queue/Cache for Rails 8 features

**View Organization Pattern**:

Rails views are organized to keep controller directories clean with
page-specific partials in a dedicated `partials/` directory:

```
app/views/
├── {controller}/
│   ├── {action}.html.erb      # Main views only (one file per page)
│   └── {another_action}.html.erb
└── partials/
    ├── {page_name}/
    │   ├── _section_name.html.erb
    │   └── _another_section.html.erb
    └── shared/                 # Truly shared across multiple pages
        └── _component.html.erb
```

Example structure:

```
app/views/
├── home/
│   ├── contact.html.erb       # Clean! Main views only
│   ├── about.html.erb
│   └── pricing.html.erb
└── partials/
    ├── contact/
    │   ├── _info_sections.html.erb
    │   ├── _additional_links.html.erb
    │   └── _contact_form.html.erb
    └── shared/
        └── _footer.html.erb
```

## Database Architecture

### PostgreSQL (Shared)

Single database with two migration systems:

- **Go migrations**: `apps/api/migrations/*.sql` (business data tables: advice,
  quotes, words)
- **Rails migrations**: `apps/dashboard/db/migrate/*.rb` (user tables: users,
  api_keys, subscriptions, usage_logs)

Separate migration tracking tables prevent conflicts.

### Cloudflare KV (Edge)

Key-value store for:

- API key lookup: `key:{api_key}` → `{userId, plan, billingCycleStart, ...}`
- Rate limiting: `rl:m:{key}:{minute}` → request count (auto-expires)

### Cloudflare D1 (Edge SQLite)

Usage tracking:

- `credit_usage` table: Records every API call (requests used per billing
  period)
- SQL queries for billing period aggregations

## Key Dependencies

### Go

- `chi` - HTTP router
- `pgx` - PostgreSQL driver
- `golang-migrate` - Database migrations
- `bobadilla-tech/is-email-disposable` - Email validation
- `bobadilla-tech/lorelai` - Lorem ipsum generation

### Rails

- Rails 8.1
- Tailwind CSS
- Turbo/Stimulus (Hotwire)
- Solid Cache/Queue
- Sidekiq

### Cloudflare Worker

- Wrangler - Deployment tool
- `@cloudflare/workers-types` - TypeScript types
- `zod` - Schema validation
- `@t3-oss/env-core` - Environment variables

## Important Notes

### Adding New Go Endpoints

See `docs/core/adding-go-endpoints.md` for the full step-by-step guide including
code patterns, testing, documentation YAML, and the pre-merge checklist.

### Go Backend Security

The Go backend trusts the Cloudflare Worker gateway completely:

- No API key validation in Go
- No rate limiting in Go
- Expects `X-Backend-Secret` header from gateway
- Only processes business logic and database queries

### Database Migrations

**Go migrations** (business data):

```bash
# Add new migration in apps/api/migrations/
# Named: NNNN_description.up.sql
# Runs automatically on app startup via app/app.go:migrateWithRetry()
```

**Rails migrations** (user data):

```bash
docker exec requiem-dev-dashboard-1 bin/rails generate migration AddFieldToTable
docker exec requiem-dev-dashboard-1 bin/rails db:migrate
```

### Cloudflare Worker Development

1. Local development uses miniflare (simulates KV, D1)
2. Secrets must be set via `wrangler secret put`:
   - `BACKEND_URL`
   - `BACKEND_SECRET`
3. KV/D1 bindings configured in `wrangler.toml`
4. Seed KV with: `pnpm run kv:seed`

## Common Development Tasks

**Note:** These commands manage the Docker containers. Run from `infra/docker`
directory.

View service logs:

```bash
cd infra/docker
docker compose -f docker-compose.dev.yml logs -f api
docker compose -f docker-compose.dev.yml logs -f dashboard
```

Reset database:

```bash
cd infra/docker
docker compose -f docker-compose.dev.yml down -v
docker compose -f docker-compose.dev.yml up
```

Restart a specific service:

```bash
cd infra/docker
docker compose -f docker-compose.dev.yml restart api
docker compose -f docker-compose.dev.yml restart dashboard
```

Connect to database:

- Host: localhost
- Port: 5432
- Database: requiem
- User: requiem
- Password: requiem
