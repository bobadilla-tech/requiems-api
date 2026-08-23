# Copilot instructions

Requiems API is a monorepo with a Go API, Rails dashboard, and Bun MCP server.

## Runtime boundary

`Client -> Cloudflare -> Caddy/AOP -> Kamal -> Go -> Redis/Postgres`.

Go owns API-key auth, Redis rate limits, quotas, and usage. Rails owns users,
API keys, subscriptions, plans, and dashboard operations. MCP uses the same
Go API. Cloudflare provides DNS, WAF, DDoS, TLS, and AOP only.

## Development

```bash
cd infra/docker
docker compose -f docker-compose.dev.yml up
docker exec requiem-dev-api-1 go test -race ./...
docker exec requiem-dev-dashboard-1 bin/rails test
cd apps/mcp && bun test && bunx tsc --noEmit
```

Use `TEST_DATABASE_URL`, `RAILS_ENV=test`, and an isolated Redis target.
Set `LOCAL_DEV_API_KEY` to a valid `requiem_<24 alphanumeric characters>`
value for local Rails seed, MCP stdio, and load tests.

Do not add Worker, KV, D1, API-management, or normal Rails-to-Go
`BACKEND_SECRET` dependencies. Preserve the separate private-deployment
`tenant_secret` contract.

