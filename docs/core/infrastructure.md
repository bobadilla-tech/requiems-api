# Infrastructure

## Overview

Requiem API runs on a distributed architecture combining edge computing
(Cloudflare) with VPS hosting (Hetzner) to deliver fast, reliable API services
globally.

## Architecture Components

```
┌──────────────────────────────────────────────────────────────┐
│                     Cloudflare Global Network                │
│                                                              │
│  ┌────────────────────────────────────────────────────┐      │
│  │  Worker (Auth Gateway)                             │      │
│  │  - API key validation                              │      │
│  │  - Rate limiting                                   │      │
│  │  - Request tracking                                │      │
│  └────────────────────────────────────────────────────┘      │
│                                                              │
│  ┌───────────────┐    ┌───────────────────────────────┐      │
│  │  KV (Keys)    │    │  D1 (Usage Tracking)          │      │
│  │  ~10ms global │    │  SQLite at the edge           │      │
│  └───────────────┘    └───────────────────────────────┘      │
└──────────────────────────────────────────────────────────────┘
                              ↓
                  X-Backend-Secret header
                              ↓
┌──────────────────────────────────────────────────────────────┐
│              Hetzner VPS (Docker Compose)                    │
│                                                              │
│  ┌────────────────────────────────────────────────────┐      │
│  │  Caddy Reverse Proxy                               │      │
│  │  - HTTPS termination (Let's Encrypt)               │      │
│  │  - Ports 80/443                                    │      │
│  └────────────────────────────────────────────────────┘      │
│                              ↓                               │
│  ┌────────────────┐    ┌─────────────────┐                   │
│  │  Go API        │    │  Rails Dashboard│                   │
│  │  Port 8080     │    │  Port 3000      │                   │
│  └────────────────┘    └─────────────────┘                   │
│                              ↓                               │
│  ┌────────────────────────────────────────────────────┐      │
│  │  PostgreSQL 16                                     │      │
│  │  - Shared database for Go & Rails                  │      │
│  │  - Port 5432                                       │      │
│  └────────────────────────────────────────────────────┘      │
│                                                              │
│  ┌────────────────────────────────────────────────────┐      │
│  │  Redis 7                                           │      │
│  │  - Sidekiq background jobs                         │      │
│  │  - Port 6379                                       │      │
│  └────────────────────────────────────────────────────┘      │
└──────────────────────────────────────────────────────────────┘
```

## Docker Compose Setup

All backend services run in Docker containers managed by Docker Compose.

### Development Environment

**File:**
[infra/docker/docker-compose.dev.yml](../../infra/docker/docker-compose.dev.yml)

**Services:**

- `api` - Go backend with Air hot reloading
- `dashboard` - Rails with native hot reloading
- `db` - PostgreSQL 16
- `redis` - Redis 7
- `sidekiq` - Background job processor
- `auth-gateway` - Cloudflare Worker (Wrangler) — public API entry point
- `api-management` - Cloudflare Worker (Wrangler) — internal API key management

**Features:**

- Volume mounts for live code reloading
- Development dependencies included
- Exposed ports for direct access
- Development-optimized configurations

### Production Environment

**File:**
[infra/docker/docker-compose.yml](../../infra/docker/docker-compose.yml)

**Services:**

- `api` - Compiled Go binary
- `dashboard` - Rails with Thruster + Puma
- `db` - PostgreSQL 16
- `redis` - Redis 7
- `sidekiq` - Background job processor
- `caddy` - HTTPS reverse proxy (optional, prod profile)

**Features:**

- Multi-stage Docker builds
- Optimized production images
- No development dependencies
- Secure configurations

## Database (PostgreSQL)

**Version:** 16 (Alpine) **Port:** 5432 **Credentials:** `requiem` / `requiem` /
`requiem` (dev), use env vars in production

### Shared Database Strategy

Both Go and Rails connect to the same PostgreSQL instance but maintain separate
migration tracking:

- **Go tables:** Business data (advice, quotes, words, etc.)
- **Rails tables:** User data (users, api_keys, subscriptions, etc.)
- **Migration tracking:** Separate `schema_migrations` tables

This approach:

- ✅ Simplifies infrastructure (one database)
- ✅ Enables data sharing when needed
- ✅ Reduces operational complexity
- ⚠️ Requires coordination on schema changes

### Database Volumes

Data persists in Docker volumes:

- Development: `requiem-dev_db_data`
- Production: `requiem-backend_db_data`

### Backup Strategy

```bash
# Backup database
docker compose exec db pg_dump -U requiem requiem > backup.sql

# Restore database
cat backup.sql | docker compose exec -T db psql -U requiem requiem
```

## Redis

**Version:** 7 (Alpine) **Port:** 6379 **Purpose:** Background job queue for
Sidekiq and real-time counter storage for the Go API

### Usage

- **Go API** — Atomic counter increments (`INCR counter:{namespace}`); counter
  values are synced to PostgreSQL every 60 seconds by the background sync worker
- Rails background jobs (email sending, usage sync, cleanup)
- Future: Caching layer for API responses

## Caddy Reverse Proxy

**Version:** 2 (Alpine) **Ports:** 80 (HTTP), 443 (HTTPS) **Configuration:**
[infra/caddy/Caddyfile](../../infra/caddy/Caddyfile)

### Features

- Automatic HTTPS with Let's Encrypt
- Zero-configuration TLS certificate management
- HTTP/2 and HTTP/3 support
- Gzip compression
- Reverse proxy to backend services

## Port Mapping

### Development

- `4455` - Auth Gateway (Cloudflare Worker)
- `5544` - API Management (Cloudflare Worker)
- `8080` - Go API
- `3000` - Rails Dashboard
- `5432` - PostgreSQL
- `6379` - Redis

### Production (with Caddy)

- `80` - HTTP (redirects to HTTPS)
- `443` - HTTPS (Caddy → services)
- `8080` - Go API (internal only)
- `3000` - Rails Dashboard (internal only)

## Networking

### Development

All services communicate via Docker's default bridge network with service names
as hostnames:

- `db` - PostgreSQL
- `redis` - Redis
- `api` - Go API
- `dashboard` - Rails

### Production

Same Docker network, plus:

- Caddy handles external HTTPS
- Backend services not directly exposed
- Worker authenticates via `X-Backend-Secret` header

## Resource Requirements

### Minimum (Development)

- **CPU:** 2 cores
- **RAM:** 4GB
- **Disk:** 10GB

### Recommended (Production)

- **CPU:** 4 cores (Hetzner CPX21)
- **RAM:** 8GB
- **Disk:** 80GB SSD
- **Network:** 20TB bandwidth/month

## Monitoring

### Health Checks

**Go API:**

```bash
curl http://localhost:8080/healthz
# Response: {"status":"ok"}
```

**Rails Dashboard:**

```bash
curl http://localhost:3000
# Response: 200 OK
```

**PostgreSQL:**

```bash
docker compose exec db pg_isready -U requiem
# Response: accepting connections
```

### Logs

```bash
# All services
docker compose logs -f

# Specific service
docker compose logs -f api
docker compose logs -f dashboard
docker compose logs -f sidekiq
```

### Resource Usage

```bash
# Container stats
docker compose stats

# Disk usage
docker compose exec db du -sh /var/lib/postgresql/data
```

## Maintenance

### Update Dependencies

**Go API:**

```bash
docker compose exec api go get -u ./...
docker compose exec api go mod tidy
docker compose up --build api
```

**Rails Dashboard:**

```bash
docker compose exec dashboard bundle update
docker compose up --build dashboard
```

### Database Migrations

**Go:**

```bash
docker compose exec api go run ./cmd/migrate
```

**Rails:**

```bash
docker compose exec dashboard rails db:migrate
```

### Clean Up

```bash
# Remove stopped containers
docker compose down

# Remove volumes (WARNING: deletes data)
docker compose down -v

# Remove images
docker compose down --rmi all

# Clean system
docker system prune -a
```

## Security

### Firewall Configuration (UFW)

```bash
# Allow SSH, HTTP, HTTPS
ufw allow 22/tcp
ufw allow 80/tcp
ufw allow 443/tcp
ufw enable

# Verify
ufw status
```

**Warning:** Docker writes iptables DNAT rules directly into its own `DOCKER`
chain, ahead of UFW's chain. Any service with a `ports:` mapping in
docker-compose is published to `0.0.0.0` regardless of UFW rules — UFW does NOT
protect published container ports. Never add a `ports:` entry for `db` (or any
datastore); use `docker compose exec` for one-off host access, or bind
explicitly to `127.0.0.1:<port>:<port>` if local tooling truly needs it.

### Environment Variables

Never commit secrets to git. Use:

- `.env` files (gitignored)
- Docker Compose environment variables
- Wrangler secrets for Cloudflare Worker

### Database Security

- Change default credentials in production
- Use SSL connections in production
- Restrict PostgreSQL to local connections — enforced by NOT publishing a
  `ports:` entry on `db` in docker-compose.yml (see UFW warning above)
- Regular backups

## Related Documentation

- [Local Development Setup](../../infra/readme.md)
- [Docker Compose README](../../infra/docker/README.md)
- [Deployment Guide](./deployment.md)
