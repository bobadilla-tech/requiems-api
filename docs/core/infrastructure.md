# Infrastructure

## Production

`api.requiems.xyz` is an orange-cloud DNS record pointing to the VPS.
Cloudflare terminates the public edge and requires the Cloudflare Origin Pull
client certificate at Caddy. Caddy forwards only to the Kamal Go proxy. The
Go process listens on port 8080 inside the application network.

Cloudflare remains responsible for DNS, proxying, WAF, DDoS protection, TLS,
and AOP. It does not run application code and does not store API keys or
usage.

The other production applications are:

- Rails dashboard at `requiems.xyz`
- MCP server at `mcp.requiems.xyz`
- Postgres and Redis as Kamal accessories on the VPS
- LanguageTool as an API accessory

The retired `api-management.requiems.xyz` hostname has no application
service and must not be recreated.

## Local Compose services

From `infra/docker`:

| Service | Port | Role |
| --- | ---: | --- |
| api | 8080 | Direct Go API |
| dashboard | 3000 | Rails dashboard |
| mcp | 3100 | MCP HTTP server |
| db | 5433 | PostgreSQL |
| redis | 6379 | Redis |
| languagetool | 8010 | Optional Go dependency |

Dashboard, Sidekiq, and MCP depend on the Go `api` service, not a gateway or
management service. Set `LOCAL_DEV_API_KEY` in `.env.local`; it must match
`requiem_<24 alphanumeric characters>`.

## Trust boundaries

Public API authentication and authorization are Go middleware concerns.
Caddy's mTLS check prevents direct-origin bypass. Rails' private Playground
connection is network-private and still sends a Go API key. Private
deployment customers have a separate `tenant_secret` contract; that is not
the retired normal Rails-to-Go `BACKEND_SECRET` setting.

## Observability

Use the Cloudflare Ray ID, Caddy access log, Go structured request ID, and
Postgres/Redis logs together when tracing a request. Do not treat a Cloudflare
edge response alone as proof that Go handled the request.

