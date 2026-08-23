# Architecture

## Current request path

```
Client
  -> Cloudflare proxy / WAF / DDoS / TLS / AOP
  -> Caddy
  -> Kamal-managed Go API
  -> Redis + PostgreSQL
```

Cloudflare is infrastructure only. There are no application Workers, KV
namespaces, or D1 databases in the request path.

The Go API owns API-key authentication, per-key Redis rate limiting, monthly
quota checks, and Postgres usage recording. API keys are generated and revoked
by Rails and verified by Go against the shared Postgres database. Requests use
the `requiems-api-key` header.

Rails owns users, API keys, subscriptions, plans, the dashboard, billing
webhooks, and private deployment requests. The Playground calls the Go API over
the private Docker/Kamal network at `http://requiems-api:8080` and passes a real
Go API key.

## Data ownership

- PostgreSQL: users, api_keys, subscriptions, plans, usage_logs, and the Go
  product/reference tables.
- Redis: Go auth verification cache, rate-limit counters, quota state, Rails
  jobs, and bounded application caches.
- Cloudflare: DNS, proxying, WAF, DDoS protection, TLS, and authenticated origin
  pulls only.

Operational usage ledgers may be discarded during the pre-launch cleanup, but
protected identity, billing, plan, and product data is never truncated.

## Service boundaries

`apps/api` is the enforcing API. `apps/dashboard` is the Rails web and job
application. `apps/mcp` is a client/tool server that calls the same Go API.
Caddy accepts origin traffic only with Cloudflare's client certificate. Kamal
deploys the API, dashboard/Sidekiq, and MCP images.

See [infrastructure](infrastructure.md), [deployment](deployment.md), and
[getting started](getting-started.md) for operational details.
