## Infrastructure Guide

This document explains how to run Requiem API locally for development and how to
deploy it to production using either Docker Compose (simple VPS) or Kubernetes 
(scalable cluster).

---

## Deployment Options

### Option 1: Docker Compose (Recommended for Simple Deployments)
- Best for: Single server, development, small-scale production
- Location: `infra/docker/`
- See sections 1 and 2 below

### Option 2: Kubernetes (Recommended for Production at Scale)
- Best for: Multi-server clusters, scalable production deployments
- Location: `infra/kubernetes/`
- See `infra/kubernetes/README.md` for detailed Kubernetes guide

---

## 1. Local Development Setup

### 1.1 Requirements

- Docker and Docker Compose
- Go 1.22+ (optional, only needed if you want to run the API directly)
- Node/PNPM/Yarn (optional, for Cloudflare Worker tooling later)

### 1.2 Run everything with Docker (recommended)

From the project root:

```bash
cd infra/docker
docker compose up --build
```

This starts:

- `api` – Go backend on internal port `8080` (exposed as `localhost:6969`)
- `db` – PostgreSQL (`requiem` / `requiem` / `requiem`, exposed as
  `localhost:5432`)
- `redis` – Redis for future queues/cache

Once the stack is up:

- Health check: `http://localhost:6969/healthz`
- Advice endpoint: `http://localhost:6969/v1/advice`

> Note: Caddy is mainly for the VPS setup. For local development you can hit the
> API directly on `localhost:6969`.

### 1.3 Run the API directly (without Docker)

If you prefer to run the Go server directly:

1. Make sure PostgreSQL is running locally (matching the default DSN or set
   `DATABASE_URL`).
2. From the project root:

```bash
export DATABASE_URL="postgres://requiem:requiem@localhost:5432/requiem?sslmode=disable"
go run ./apps/api
```

The API listens on `:8080` by default:

- `http://localhost:8080/healthz`
- `http://localhost:8080/v1/advice`

### 1.4 Hybrid dev workflow (Docker infra + local Go with hot reload)

For the best developer experience, run **Postgres and Redis in Docker**, and the
Go API locally with a watcher:

- Start infra only:

```bash
cd infra/docker
docker compose up db redis
```

- In another terminal, from the repo root:

```bash
export DATABASE_URL="postgres://requiem:requiem@localhost:5432/requiem?sslmode=disable"
go run ./apps/api
```

Or, with a hot-reload tool like `air`:

```bash
export DATABASE_URL="postgres://requiem:requiem@localhost:5432/requiem?sslmode=disable"
air
```

The API is still available on:

- `http://localhost:6969/healthz`
- `http://localhost:6969/v1/advice`

### 1.5 Cloudflare Worker (edge auth) – dev notes

The Worker lives in `apps/edge-auth/index.ts`. A typical dev setup will:

- Use `wrangler dev` to run the worker locally.
- Set environment variables:
  - `BACKEND_ORIGIN` – e.g. `http://localhost:8080`
  - `API_KEY_SECRET` – your shared API key used by clients in `x-api-key`.

> The exact `wrangler.toml` configuration can be added once you hook up your
> Cloudflare account.

---

## 2. VPS Deployment Guide

Target: single VPS (Docker + Docker Compose) running API, Postgres, Redis, and
Caddy. Cloudflare Worker sits in front and forwards authorized traffic.

### 2.1 Prepare the VPS

1. Create a new VPS (any provider).
2. SSH into the server.
3. Install Docker and Docker Compose (or Docker Compose plugin).
4. Clone the repository:

```bash
git clone <your-repo-url> requiems-api
cd requiems-api/infra/docker
```

### 2.2 Start the stack

From `infra/docker`:

```bash
docker compose up -d --build
```

Services:

- `api` – Go API (port `8080` in the Docker network)
- `db` – PostgreSQL
- `redis` – Redis
- `caddy` – HTTPS reverse proxy on ports `80` and `443`

You can check logs with:

```bash
docker compose logs -f api
docker compose logs -f caddy
```

### 2.3 Caddy configuration (HTTPS + reverse proxy)

Caddy is configured via `infra/caddy/Caddyfile`. Example:

```bash
api.yourdomain.com {
  encode gzip

  reverse_proxy api:8080
}
```

What this does:

- Terminates HTTPS and manages TLS certificates automatically.
- Proxies `https://api.yourdomain.com` → `api:8080` (the API container).

### 2.4 DNS setup

- Create an `A` record for `api.yourdomain.com` pointing to your VPS public IP.
- Wait for DNS to propagate.
- With DNS + Caddy running, `https://api.yourdomain.com/healthz` should respond
  successfully.

### 2.5 Cloudflare Worker configuration (edge auth)

1. Deploy `apps/edge-auth/index.ts` as a Worker in your Cloudflare account.
2. Configure environment variables/secrets:
   - `BACKEND_ORIGIN` – `https://api.yourdomain.com`
   - `BACKEND_SECRET` – a strong secret, used by clients in `x-api-key`.
3. Route your public API endpoint to the Worker (e.g.
   `https://v1.yourdomain.com/*` → Worker).

Request flow in production:

1. Client → Worker (with `x-api-key`).
2. Worker validates `x-api-key`.
3. Worker forwards to `https://api.yourdomain.com/...`.
4. Cloudflare → Caddy on VPS → Go API.

---

## 3. Environment Variables Summary

- **API container**

  - `PORT` – API listen port (default `8080`).
  - `DATABASE_URL` – Postgres DSN (set by Docker Compose for container, or
    manually in local dev).
  - `REDIS_URL` – Redis URL for future queues/cache.

- **Cloudflare Worker**
  - `BACKEND_ORIGIN` – Base URL of the API behind Caddy.
  - `BACKEND_SECRET` – Shared secret checked against `x-api-key`.

---

## 4. Kubernetes Deployment (Alternative to Docker Compose)

For production deployments requiring scalability, high availability, and advanced 
orchestration features, you can deploy to Kubernetes instead of using Docker Compose.

### 4.1 When to Use Kubernetes

Use Kubernetes when you need:
- **Horizontal scaling**: Automatically scale services based on load
- **High availability**: Run multiple replicas with automatic failover
- **Rolling updates**: Zero-downtime deployments
- **Advanced networking**: Service mesh, network policies, ingress controllers
- **Multi-node clusters**: Distribute workload across multiple servers
- **Production-grade orchestration**: Self-healing, automated restarts, health checks

### 4.2 Kubernetes Setup

All Kubernetes manifests are located in `infra/kubernetes/`:

```
infra/kubernetes/
├── namespace.yaml      # Kubernetes namespace
├── configmap.yaml      # Configuration data
├── secrets.yaml        # Sensitive data (passwords, API keys)
├── database.yaml       # PostgreSQL service and deployment
├── redis.yaml          # Redis service and deployment
├── api.yaml            # Go API service and deployment
├── dashboard.yaml      # Rails dashboard, sidekiq deployments
├── ingress.yaml        # Ingress for external access
├── deploy.sh           # Automated deployment script
├── cleanup.sh          # Cleanup script
└── README.md           # Detailed Kubernetes guide
```

### 4.3 Service Architecture in Kubernetes

All services use **ClusterIP** for internal cluster communication:

- `db:5432` - PostgreSQL (internal only)
- `redis:6379` - Redis (internal only)  
- `api:8080` - Go API (internal, exposed via Ingress to internal.requiems.xyz)
- `dashboard:80` - Rails app (exposed via Ingress to requiems.xyz)
- `sidekiq` - Background jobs (no service, deployment only)

External access is managed through an Ingress controller.

### 4.4 Quick Deploy

```bash
cd infra/kubernetes

# Review and update secrets
vi secrets.yaml

# Deploy everything
./deploy.sh

# Check status
kubectl get pods -n requiem
kubectl get svc -n requiem
kubectl get ingress -n requiem
```

### 4.5 Single-Server Kubernetes with k3s

For a single-server deployment with Kubernetes benefits:

```bash
# On your server, install k3s (lightweight Kubernetes)
curl -sfL https://get.k3s.io | sh -

# Deploy Requiem API
cd infra/kubernetes
./deploy.sh
```

### 4.6 Scaling Services

```bash
# Scale API instances
kubectl scale deployment api -n requiem --replicas=3

# Scale dashboard instances  
kubectl scale deployment dashboard -n requiem --replicas=3
```

For complete Kubernetes deployment instructions, troubleshooting, and production 
considerations, see **[infra/kubernetes/README.md](kubernetes/README.md)**.

---

## 5. Choosing Between Docker Compose and Kubernetes

| Feature | Docker Compose | Kubernetes |
|---------|---------------|------------|
| **Complexity** | Simple | More complex |
| **Best For** | Single server, dev | Production at scale |
| **Scaling** | Manual | Automatic |
| **High Availability** | Limited | Built-in |
| **Updates** | Manual restart | Rolling updates |
| **Learning Curve** | Easy | Moderate |
| **Resource Overhead** | Low | Higher |
| **Use Case** | Simple deployments | Enterprise production |

**Recommendation:**
- Use **Docker Compose** for simple VPS deployments and development
- Use **Kubernetes** for production deployments requiring scale and resilience
