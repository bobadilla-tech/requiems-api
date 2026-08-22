# The Auth Gateway (Cloudflare Worker)

## Overview

The Auth Gateway is a Cloudflare Worker that sits at the edge of the Requiem API
architecture, handling authentication, rate limiting, and request tracking
before requests reach the backend.

**Location:** [apps/workers/auth-gateway/](../apps/workers/auth-gateway/)
**Technology:** TypeScript on Cloudflare Workers **Runtime:** V8 isolates
running globally across Cloudflare's network

## Architecture

The Worker acts as the first line of defense and management for all API
requests:

```
Client Request
    ↓
Worker (validates requiems-api-key header)
    ↓
KV Lookup (API key → user data)
    ↓
Rate Limit Check (KV counter)
    ↓
Usage Check (D1 query)
    ↓
Forward to Backend (with X-Backend-Secret)
    ↓
Record Usage (D1 insert)
    ↓
Add Usage Headers & Return Response
```

## Responsibilities

### 1. API Key Validation

- Extracts `requiems-api-key` header from requests
- Looks up key data in Cloudflare KV (~10ms global latency)
- Returns 401 if key is invalid or missing

### 2. Rate Limiting

- Enforces per-minute rate limits based on plan tier
- Stores counters in KV with 60-second TTL
- Returns 429 when limits exceeded

### 3. Usage Tracking

- Checks request usage against plan limits
- Queries Cloudflare D1 (SQLite at the edge)
- Tracks per-endpoint request costs
- Monthly reset

### 4. Request Forwarding

- Adds `X-Backend-Secret` header for backend authentication
- Strips Cloudflare transport headers while preserving `requiems-api-key` for
  Go's trusted backend to verify the complete credential
- Forwards to internal backend URL

### 5. Usage Recording

- Records request usage in D1 database
- Tracks endpoint, timestamp, and requests used
- Enables usage analytics and billing

## Key Files

### Core Logic

- [src/index.ts](../apps/workers/auth-gateway/src/index.ts) - Hono app wiring
  (thin entry point)
- [src/middleware/api-key-auth.ts](../apps/workers/auth-gateway/src/middleware/api-key-auth.ts) -
  API key validation, rate limiting, and quota check middleware
- [src/routes/proxy.ts](../apps/workers/auth-gateway/src/routes/proxy.ts) -
  Proxies to the Go backend and records usage
- [src/env.ts](../apps/workers/auth-gateway/src/env.ts) - Worker environment
  bindings
- [src/requests.ts](../apps/workers/auth-gateway/src/requests.ts) - Usage
  tracking (KV cache-first, falls back to D1 aggregate query on miss)
- [src/rate-limit.ts](../apps/workers/auth-gateway/src/rate-limit.ts) - Rate
  limiting logic
- [src/http.ts](../apps/workers/auth-gateway/src/http.ts) - HTTP helpers and
  header filtering
- [src/generated/openapi.ts](../apps/workers/auth-gateway/src/generated/openapi.ts) -
  Generated OpenAPI spec, served at `/openapi.json`

Plan limits and endpoint costs live in the shared package:
[apps/workers/shared/src/config.ts](../apps/workers/shared/src/config.ts)

### Configuration

- [wrangler.toml](../apps/workers/auth-gateway/wrangler.toml) - Worker
  configuration
- [schema.sql](../apps/workers/auth-gateway/schema.sql) - D1 database schema

### Development

- [scripts/seed-dev.ts](../apps/workers/auth-gateway/scripts/seed-dev.ts) -
  Seeds KV and D1 with dev API keys

## Development

### Local Development

The worker runs inside Docker via the dev compose stack — no manual setup
needed:

```bash
cd infra/docker
docker compose -f docker-compose.dev.yml up
# Auth Gateway available at http://localhost:4455
```

For standalone type checking or linting outside Docker:

```bash
cd apps/workers/auth-gateway
pnpm run typecheck
pnpm run lint
```

### Environment Variables

Set these as Wrangler secrets (never commit):

- `BACKEND_URL` - Internal backend endpoint (e.g.,
  `https://internal.requiems.xyz`)
- `BACKEND_SECRET` - Shared secret for backend authentication (min 32 chars)

```bash
wrangler secret put BACKEND_URL
wrangler secret put BACKEND_SECRET
```

### Cloudflare Resources

**KV Namespace (API keys & rate limits):**

```bash
wrangler kv:namespace create KV
# Add ID to wrangler.toml
```

**D1 Database (usage tracking):**

```bash
wrangler d1 create requiem-usage
wrangler d1 execute requiem-usage --file=schema.sql
# Add ID to wrangler.toml
```

### Testing Locally

Dev API keys are seeded automatically when the stack starts. Use them directly:

```bash
curl -H "requiems-api-key: requiem_free00000000000000000001" http://localhost:4455/v1/text/advice
```

## Response Headers

Every successful response includes usage information:

```
X-Requests-Used: 1
X-Requests-Remaining: 499
X-Requests-Reset: 2026-03-01T00:00:00.000Z
X-Plan: developer
X-RateLimit-Limit: 5000
X-RateLimit-Remaining: 4999
```

## Plan Tiers

| Plan         | Request Limit | Period  | Rate Limit | Price      |
| ------------ | ------------- | ------- | ---------- | ---------- |
| Free         | 500           | Monthly | 30/min     | $0/month   |
| Developer    | 100,000       | Monthly | 5,000/min  | $30/month  |
| Business     | 1,000,000     | Monthly | 10,000/min | $75/month  |
| Professional | 10,000,000    | Monthly | 50,000/min | $150/month |
| Enterprise   | Unlimited     | Monthly | Unlimited  | Custom     |

## Endpoint Costs

Some endpoints cost more than one request.

See [apps/workers/shared/src/config.ts](../apps/workers/shared/src/config.ts)
for the full list.

## Deployment

```bash
# Deploy to production
pnpm run deploy

# Or deploy to specific environment
pnpm run deploy:prod
```

## Integration with Rails Dashboard

The Rails dashboard manages API keys through the API Management worker, which
keeps KV in sync automatically:

- New API key created/updated/revoked → goes through
  [CloudflareApiManagementService](../apps/dashboard/app/services/cloudflare/api_management_service.rb)
- Rails never writes to KV directly — all changes go through the API Management
  worker

This ensures the Auth Gateway always has up-to-date key data.

## Performance

- **KV Lookups:** ~10ms globally
- **D1 Queries:** ~15-30ms
- **Total Overhead:** ~50-100ms before reaching backend
- **Global Distribution:** Runs in 300+ Cloudflare data centers worldwide

## Related Documentation

- [Architecture Overview](./architecture.md)
- [Backend Documentation](./backend.md)
- [Rails App Documentation](./rails-app.md)
