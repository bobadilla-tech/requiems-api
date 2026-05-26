# Docker Setup

## Development Mode (Hot Reloading)

Start all services with **hot reloading enabled**:

```bash
cd infra/docker
docker compose -f docker-compose.dev.yml up --build
```

### What You Get:

✅ **Go API** with Air hot reloading on port **8080**

- Changes to `.go` files automatically rebuild and restart
- Access: http://localhost:8080/healthz

✅ **Rails Dashboard** with native hot reloading on port **3000**

- Changes to Ruby files automatically reload
- Changes to views refresh on browser reload
- Access: http://localhost:3000

✅ **PostgreSQL** on port **5432**

- Database persists between restarts
- Shared between Go and Rails

✅ **Redis** on port **6379**

- For Sidekiq background jobs

✅ **Auth Gateway** (Cloudflare Worker via Wrangler) on port **4455**

- Public-facing API entry point
- Validates API keys, enforces rate limits, proxies to Go backend
- Seeded with dev API keys automatically on start
- Applies local D1 schema/migrations on start for dev
- Access: http://localhost:4455/healthz

✅ **API Management** (Cloudflare Worker via Wrangler) on port **5544**

- Internal service for API key CRUD, usage exports, analytics
- Uses the same local Wrangler state as Auth Gateway (shared D1 + KV in dev)
- Swagger docs at http://localhost:5544/docs (user: `local`, pass: `password`)
- Access: http://localhost:5544/healthz

✅ **Sidekiq** background worker

- Processes Rails jobs automatically

### Optional: LanguageTool (spell-check)

LanguageTool is **opt-in** — it's a large JVM service and not needed for most
development work. Enable it with the `languagetool` profile:

```bash
docker compose -f docker-compose.dev.yml --profile languagetool up
```

Without this profile, the spell-check API endpoint returns `503 upstream_error`
instead of results. All other endpoints are unaffected.

### First Time Setup:

The dev images will automatically:

1. Build the local Go, Rails, and Worker development images
2. Install dependencies (Go modules, Ruby gems) during image build
3. Start all services

On later restarts, containers reuse those built images, so startup is faster and
less sensitive to runtime package-install failures.

### When to rebuild:

Rebuild after dependency changes or the first time you start the stack:

1. `apps/api/go.mod` or `apps/api/go.sum`
2. `apps/dashboard/Gemfile` or `apps/dashboard/Gemfile.lock`
3. `infra/docker/api.dev.Dockerfile` or `infra/docker/dashboard.dev.Dockerfile`
4. `apps/workers/auth-gateway/wrangler.toml` or
   `apps/workers/api-management/wrangler.toml`
5. `infra/docker/auth-gateway.dev.Dockerfile` or
   `infra/docker/api-management.dev.Dockerfile`

### Worker local state and migrations

- Auth Gateway and API Management share the same local Wrangler persistence path
  in Docker, so both services use the same local D1 database and KV namespace.
- Auth Gateway startup seeds dev keys and applies D1 schema/migrations.
- API Management startup also applies D1 schema before serving requests.

If local D1 schema drifts, reset Docker volumes and rebuild:

```bash
docker compose -f docker-compose.dev.yml down -v
docker compose -f docker-compose.dev.yml up --build
```

### Development Workflow:

```bash
# Start everything
docker compose -f docker-compose.dev.yml up

# Rebuild if you change dependencies or dev Dockerfiles
docker compose -f docker-compose.dev.yml up --build

# View logs for specific service
docker compose -f docker-compose.dev.yml logs -f api
docker compose -f docker-compose.dev.yml logs -f dashboard

# Stop everything
docker compose -f docker-compose.dev.yml down

# Stop and remove volumes (reset database)
docker compose -f docker-compose.dev.yml down -v
```

### Accessing Services:

| Service         | URL                   | Notes                                         |
| --------------- | --------------------- | --------------------------------------------- |
| Auth Gateway    | http://localhost:4455 | Public API entry point                        |
| API Management  | http://localhost:5544 | Internal management service                   |
| Rails Dashboard | http://localhost:3000 | Sign up, sign in, dashboard                   |
| Go API          | http://localhost:8080 | Internal API (gateway → backend)              |
| PostgreSQL      | localhost:5432        | User: `requiem`, Password: `requiem`          |
| Redis           | localhost:6379        | For Sidekiq                                   |
| LanguageTool    | http://localhost:8010 | Spell-check backend (profile: `languagetool`) |

### Hot Reloading:

**Go (Air):**

- Edit any `.go` file
- Air detects changes automatically
- Rebuilds and restarts the server (~2-3 seconds)

**Rails:**

- Edit controllers, models, views
- Rails reloads code automatically (no restart needed)
- Refresh browser to see view changes

### Running Commands Inside Containers:

```bash
# Rails console
docker compose -f docker-compose.dev.yml exec dashboard rails console

# Go build/test
docker compose -f docker-compose.dev.yml exec api go test ./...

# Database migrations
docker compose -f docker-compose.dev.yml exec dashboard rails db:migrate

# Create admin user
docker compose -f docker-compose.dev.yml exec dashboard rails runner "
User.create!(
  name: 'Admin',
  email: 'admin@requiems.xyz',
  password: 'password123',
  password_confirmation: 'password123',
  admin: true,
  confirmed_at: Time.current
)
"
```

### Troubleshooting:

**Port already in use:**

```bash
# Find and kill process using port 3000
lsof -ti:3000 | xargs kill -9

# Or change port in docker-compose.dev.yml
ports:
  - "3001:3000"  # Access on localhost:3001 instead
```

**Database connection errors:**

```bash
# Wait for PostgreSQL to fully start (check logs)
docker compose -f docker-compose.dev.yml logs db

# If still failing, restart database
docker compose -f docker-compose.dev.yml restart db
```

**Gem or bundle issues (Rails):**

```bash
# Rebuild the dashboard and sidekiq dev image
docker compose -f docker-compose.dev.yml build dashboard sidekiq
docker compose -f docker-compose.dev.yml up
```

**Go module errors:**

```bash
# Run go mod tidy
docker compose -f docker-compose.dev.yml exec api go mod tidy

# Or rebuild
docker compose -f docker-compose.dev.yml up --build api
```

---

## Production Mode

For production deployment:

```bash
cd infra/docker
docker compose up --build
```

This uses optimized Dockerfiles with:

- Multi-stage builds
- Compiled binaries
- No development dependencies
- Caddy as reverse proxy (handles TLS automatically)

### Environment Variables

All services load environment variables from a `.env` file in this directory.
Copy the example and fill in real values before starting:

```bash
cp .env.example .env
# edit .env with production values
```

All variables must be set — there are no runtime defaults. See `.env.example`
for the full list with descriptions.

### Production Differences:

- Go: Compiled to static binary (no Air)
- Rails: Uses Thruster + Puma in production mode
- All code baked into images (no volume mounts)
- `RAILS_ENV=production` and `RAILS_LOG_LEVEL=warn` are hardcoded in compose

### Cloudflare Workers

The Auth Gateway and API Management workers run on Cloudflare, not in Docker.
Deploy them separately via Wrangler from their respective directories:

```bash
# Auth Gateway
cd apps/workers/auth-gateway
pnpm run deploy

# API Management
cd apps/workers/api-management
pnpm run deploy
```

Worker secrets (`BACKEND_URL`, `BACKEND_SECRET`, `API_MANAGEMENT_API_KEY`) must
be set via `wrangler secret put` — they are not read from `.env`.

---

## Tips:

1. **Keep dev running:** Leave `docker-compose.dev.yml` running while you code.
   All changes are picked up automatically.

2. **Database GUI:** Connect with Datagrip, DBeaver, or pgAdmin using:
   - Host: `localhost`
   - Port: `5432`
   - Database: `requiem`
   - User: `requiem`
   - Password: `requiem`

3. **Reset database:**

   ```bash
   docker compose -f docker-compose.dev.yml down -v
   docker compose -f docker-compose.dev.yml up
   ```

4. **Dependency changes:** Rebuild the affected dev image after changing
   `Gemfile.lock`, `go.sum`, or the dev Dockerfiles.
