## Getting Started with Requiem API

## Local Development Setup

> Prerequisites: Docker & Docker Compose

> Editor note: If you're developing the Rails dashboard (`apps/dashboard/`),
> install a Ruby LSP editor/IDE extension and configure it to use the project's
> `ruby-lsp` gem (added to `apps/dashboard/Gemfile`). This enables code
> navigation, in-editor diagnostics, and other language-server features.

### 1. Start the stack

```bash
cd infra/docker
docker compose -f docker-compose.dev.yml up --build
```

Services available:

| Service         | URL                           | Notes                                                |
| --------------- | ----------------------------- | ---------------------------------------------------- |
| Auth Gateway    | http://localhost:4455/healthz | Dev keys seeded automatically, use `rq_free_000001`  |
| API Management  | http://localhost:5544/healthz | Swagger at `/docs` (user: `local`, pass: `password`) |
| Rails Dashboard | http://localhost:3000         | Hot reload                                           |
| Go API          | http://localhost:8080/healthz | Air hot reload, no auth (bypass gateway)             |
| PostgreSQL      | localhost:5432                | user/pass: `requiem`                                 |
| Redis           | localhost:6379                |                                                      |

### 2. Environment variables

All services load `infra/docker/.env.example` by default: **no extra setup
required** for most development work.

For features that require real credentials (payments, email sending), create
`infra/docker/.env.local` (gitignored) and add only the secrets you need:

```bash
# infra/docker/.env.local — DO NOT COMMIT
LEMONSQUEEZY_API_KEY=        # required for testing payment flows
LEMONSQUEEZY_SIGNING_SECRET= # required for webhook verification
SMTP_PASSWORD=               # required for sending real emails
```

Ask a teammate for the actual values. The app starts fine without `.env.local` —
payments and email delivery just won't work end-to-end locally.

> **Note:** In development, emails open in Letter Opener Web (not sent). Visit
> `http://localhost:3000/letter_opener` to preview them.

### 1. Sign up and get an API key

- Visit the Requiem API dashboard,
- Create an account and generate an API key.
- All requests must include your key in the `requiems-api-key` header.

### 2. Make your first request

Example using `curl` against the edge (replace `YOUR_API_KEY`):

```bash
curl https://api.requiems.xyz/v1/text/advice \
  -H "requiems-api-key: YOUR_API_KEY"
```

Development — through the Auth Gateway (port 4455, same flow as production):

```bash
# Health check (no auth required)
curl http://localhost:4455/healthz

# API endpoints — use a seeded dev key
curl http://localhost:4455/v1/text/advice \
  -H "requiems-api-key: rq_free_000001"
```

Seeded dev API keys (available automatically after stack starts):

| Key              | Plan         |
| ---------------- | ------------ |
| `rq_free_000001` | free         |
| `rq_devl_000001` | developer    |
| `rq_bizz_000001` | business     |
| `rq_prof_000001` | professional |

Auth Gateway and API Management share the same local Wrangler persistence state
in Docker, so they operate on the same local D1 and KV resources (matching
production architecture expectations).

> **Tip:** You can also hit the Go backend directly on port 8080.

### When to rebuild

Rebuild the dev images after changing dependency manifests:

- Go API: `apps/api/go.mod` or `apps/api/go.sum`
- Rails Dashboard: `apps/dashboard/Gemfile` or `apps/dashboard/Gemfile.lock`
- Workers: `apps/workers/*/wrangler.toml` or
  `infra/docker/*worker*.dev.Dockerfile`

```bash
cd infra/docker
docker compose -f docker-compose.dev.yml up --build
```

If local usage export fails with D1 schema errors, reset volumes and rebuild:

```bash
cd infra/docker
docker compose -f docker-compose.dev.yml down -v
docker compose -f docker-compose.dev.yml up --build
```

### Go API lint in Docker

```bash
docker exec requiem-dev-api-1 golangci-lint run
```
