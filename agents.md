# Repository instructions

Requiems API is a monorepo containing:

- `apps/api`: Go API and the enforcing auth/rate-limit/quota/usage path.
- `apps/dashboard`: Rails users, API keys, subscriptions, plans, billing,
  admin, and jobs.
- `apps/mcp`: Bun MCP server calling the Go API.
- `infra`: Cloudflare/DNS/WAF/TLS/AOP and Kamal/Compose deployment.

Current production flow:

```
Client -> Cloudflare proxy/WAF/DDoS/TLS/AOP -> Caddy -> Kamal -> Go
```

Cloudflare has no application Worker, KV, or D1 dependency. The normal
Rails-to-Go path uses `requiems-api-key`; it does not use
`X-Backend-Secret`. The separate encrypted private-deployment
`tenant_secret` contract remains supported.

Local stack:

```bash
docker compose -f infra/docker/docker-compose.dev.yml up
curl http://localhost:8080/healthz
```

Set `LOCAL_DEV_API_KEY` to `requiem_<24 alphanumeric characters>` in
`infra/docker/.env.local`. Go tests use `TEST_DATABASE_URL`; Rails tests
use `RAILS_ENV=test` and isolated Redis. Never run validation against
production or the development database.

Preserve unrelated worktree changes. Use `apply_patch` for edits. Keep
historical audit and plan records unchanged except the implementation notes in
the active completion plan.

