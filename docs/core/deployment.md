# Deployment

## Components

Kamal deploys three application images:

- `infra/kamal/deploy.api.yml`: Go API, Caddy, database, Redis, and LanguageTool
  accessories.
- `infra/kamal/deploy.dashboard.yml`: Rails web and Sidekiq.
- `infra/kamal/deploy.mcp.yml`: MCP server.

Cloudflare is not an application deployment target. Do not deploy Workers or
restore the retired KV/D1 resources.

## Deploy

The normal path is the GitHub Actions CD workflow. It supplies the Kamal
registry and VPS secrets, then runs the same checked-in Kamal configuration used
for a manual deploy. A direct local deploy requires the same environment:

```bash
kamal deploy -c infra/kamal/deploy.api.yml
kamal deploy -c infra/kamal/deploy.dashboard.yml
kamal deploy -c infra/kamal/deploy.mcp.yml
```

The API health check is `GET /healthz`. After every production deploy verify the
public hostname, a valid-key request, invalid-key 401, and origin AOP
certificate rejection.

## Required secrets

- Kamal registry and SSH credentials
- Postgres credentials and `DATABASE_URL`
- `REDIS_URL`
- Rails `SECRET_KEY_BASE`
- Lemon Squeezy signing secret
- SMTP password

The normal public API path has no `BACKEND_SECRET`. The private-deployment
`tenant_secret` is stored and delivered through its own Rails workflow and must
remain supported.

## Cloudflare

Keep `requiems.xyz` (Go API) and `requiemsapi.com` (Rails dashboard) both
proxied — two separate Cloudflare zones. Cloudflare should expose the DNS
record, WAF/DDoS/TLS controls, and origin-pull authentication only.
Origin-pull (AOP) mTLS is enforced on `requiems.xyz` only; `requiemsapi.com`
has never been behind AOP. The retired Worker routes, custom domains, KV
namespace, and D1 database are not rollback dependencies.

## Rollback

Rollback means reverting the Go/Caddy/Kamal commit and redeploying the
application images. It does not mean recreating a deleted Worker, KV namespace,
or D1 database.
