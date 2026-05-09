# Caddy

Caddy is the reverse proxy that runs on the VPS in production. It sits between
Cloudflare and the backend services, handling TLS termination and routing.

## Architecture

```
Internet → Cloudflare (CDN/DDoS) → Caddy (VPS) → Services
```

## What Caddy Does

### 1. Automatic HTTPS

Caddy automatically obtains and renews Let's Encrypt certificates for all
configured domains — no certbot, no cron jobs, no manual renewal.

### 2. Reverse Proxy

| Domain                  | Target         | Notes                   |
| ----------------------- | -------------- | ----------------------- |
| `requiems.xyz`          | `dashboard:80` | Rails app               |
| `internal.requiems.xyz` | `api:8080`     | Go API (secret-guarded) |

### 3. Backend Secret Guard

`internal.requiems.xyz` enforces the `X-Backend-Secret` header before forwarding
to the Go API:

```caddyfile
@authorized {
  header X-Backend-Secret {env.BACKEND_SECRET}
}
handle @authorized {
  reverse_proxy api:8080
}
handle {
  respond "Unauthorized" 403
}
```

Any request without the correct secret gets a `403 Unauthorized`. Only the
Cloudflare Worker knows this secret, so the Go API is effectively private even
though `internal.requiems.xyz` is a public domain.

## Local Development

Caddy is **not** used in local development. Services are accessed directly via
localhost ports:

| Service   | Local URL                     |
| --------- | ----------------------------- |
| Dashboard | http://localhost:3000         |
| Go API    | http://localhost:8080/healthz |
| Auth GW   | http://localhost:4455         |

## Production Setup

Caddy runs as a Docker container alongside the other services. It reads
`BACKEND_SECRET` from the environment to configure the guard matcher.

DNS records in Cloudflare must point both `requiems.xyz` and
`internal.requiems.xyz` to the VPS IP. Both should be **proxied (orange cloud)**
for DDoS protection.

See [deployment guide](../../docs/core/deployment.md) for full setup
instructions.
