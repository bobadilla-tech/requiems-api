## Infrastructure Guide

This document explains how to run Requiem API locally for development and how to
deploy it to a single VPS.

---

## 1. Local Development Setup

### 1.1 Requirements

- Docker and Docker Compose

That's it. No other tools needed locally.

### 1.2 Start the full stack

```bash
cd infra/docker
docker compose -f docker-compose.dev.yml up
```

`infra/docker/.env.example` is committed with safe dev defaults, no
configuration needed on a fresh clone.

This starts all services:

| Service          | URL                             | Description                            |
| ---------------- | ------------------------------- | -------------------------------------- |
| `api`            | `http://localhost:8080/healthz` | Go backend (Air hot reload)            |
| `dashboard`      | `http://localhost:3000`         | Rails UI (hot reload)                  |
| `auth-gateway`   | `http://localhost:4455`         | Edge gateway (Cloudflare Worker)       |
| `api-management` | `http://localhost:5544/docs`    | API key management + Swagger UI        |
| `db`             | `localhost:5432`                | PostgreSQL (`requiem/requiem/requiem`) |
| `redis`          | `localhost:6379`                | Redis (queues/cache)                   |
| `sidekiq`        | —                               | Background jobs                        |

#### Swagger UI credentials (local)

- **Username:** `local`
- **Password:** `password`

#### Dev API keys (seeded automatically)

| Plan           | Key              | Header             |
| -------------- | ---------------- | ------------------ |
| `free`         | `rq_free_000001` | `requiems-api-key` |
| `developer`    | `rq_devl_000001` | `requiems-api-key` |
| `business`     | `rq_bizz_000001` | `requiems-api-key` |
| `professional` | `rq_prof_000001` | `requiems-api-key` |

```bash
curl -H 'requiems-api-key: rq_free_000001' http://localhost:4455/v1/text/advice
```

#### API Management key (local)

```
dev_api_mgmt_key_for_local_dev_only
```

Pass as `X-API-Management-Key` header to `http://localhost:5544`.

## 2. VPS Deployment Guide

Target: single VPS (Docker + Docker Compose) running API, Postgres, Redis, and
Caddy. Cloudflare Worker sits in front and forwards authorized traffic.

> **Continuous delivery:** production deploys are automated with Kamal via
> `.github/workflows/cd.yml`. See
> [docs/core/deployment.md](../docs/core/deployment.md#continuous-delivery-github-actions)
> for the workflow, the `production` environment secrets it requires, and manual
> redeploy/rollback instructions.

### 2.1 Prepare the VPS

#### Recommended Hetzner Configuration

- **Server Type:** CPX21 or better (3 vCPU, 4GB RAM, 80GB SSD)
- **Location:** Choose based on target audience
  - Nuremberg (nbg1) - Europe/Germany
  - Helsinki (hel1) - Northern Europe
  - Ashburn (ash) - US East
- **Image:** Ubuntu 24.04 LTS
- **Networking:** Enable IPv4 (required), IPv6 (optional)

#### Initial Setup

1. **Create VPS** in Hetzner Cloud Console

2. **SSH into the server:**

```bash
ssh root@YOUR_VPS_IP
```

3. **Update system:**

```bash
apt update && apt upgrade -y
```

4. **Install Docker:**

```bash
curl -fsSL https://get.docker.com -o get-docker.sh
sh get-docker.sh
rm get-docker.sh

# Install Docker Compose plugin
apt install docker-compose-plugin -y

# Verify installation
docker --version
docker compose version
```

5. **Configure firewall:**

```bash
apt install ufw
ufw allow 22/tcp   # SSH
ufw allow 80/tcp   # HTTP
ufw allow 443/tcp  # HTTPS
ufw enable
ufw status
```

> **Warning:** Docker inserts iptables rules directly into its own `DOCKER`
> chain, ahead of UFW's chain — a `ports:` mapping in docker-compose.yml (e.g.
> exposing `db` on `5432:5432`) publishes to `0.0.0.0` regardless of UFW rules.
> Never publish datastore ports to the host; see
> [infrastructure.md](../docs/core/infrastructure.md#firewall-configuration-ufw).

6. **Clone repository:**

```bash
git clone https://github.com/bobadilla-tech/requiems-api.git
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

### 2.4 DNS Setup (Cloudflare)

#### Add Domain to Cloudflare

1. **Add your domain** to Cloudflare (if not already):
   - Sign up/login at [cloudflare.com](https://cloudflare.com)
   - Add your domain
   - Update nameservers at your registrar to Cloudflare's nameservers

2. **Create DNS records:**

| Type | Name     | Content     | Proxy Status           | TTL  |
| ---- | -------- | ----------- | ---------------------- | ---- |
| A    | api      | YOUR_VPS_IP | DNS only (grey cloud)  | Auto |
| A    | internal | YOUR_VPS_IP | DNS only (grey cloud)  | Auto |
| A    | @        | YOUR_VPS_IP | Proxied (orange cloud) | Auto |

**Important:**

- Use "DNS only" (grey cloud) for `api` and `internal` subdomains
- This allows Caddy to obtain Let's Encrypt certificates directly
- The root domain (@) can be proxied for DDoS protection

3. **Verify DNS propagation:**

```bash
dig api.requiems.xyz
dig internal.requiems.xyz
```

4. **Wait for DNS** to propagate (5-30 minutes)

5. **Cloudflare Worker routing** (for public API):
   - In Cloudflare dashboard, go to Workers & Pages
   - Deploy your edge-auth worker
   - Add a route: `api.requiems.xyz/*` → your worker
   - Set `BACKEND_URL` environment variable to `https://internal.requiems.xyz`

#### Domain Architecture

```
User Request
    ↓
api.yourdomain.com (Cloudflare Worker - handles auth)
    ↓ (validates requiems-api-key, forwards with X-Backend-Secret)
internal.yourdomain.com (Direct to VPS)
    ↓
Caddy on VPS (HTTPS termination)
    ↓
Go API container (port 8080)
```

With DNS + Caddy running:

- `https://api.yourdomain.com/healthz` → should work (via Worker)
- `https://internal.yourdomain.com/healthz` → should work (direct to VPS)

### 2.5 Cloudflare Worker configuration (edge auth)

1. Deploy `apps/edge-auth/index.ts` as a Worker in your Cloudflare account.
2. Configure environment variables/secrets:
   - `BACKEND_ORIGIN` – `https://api.requiems.xyz.com`
   - `BACKEND_SECRET` – a strong secret, used by clients in `X-Backend-Secret`.
3. Route your public API endpoint to the Worker (e.g.
   `https://api.requiems.xyz/*` → Worker).

Request flow in production:

1. Client → Worker (with `X-Backend-Secret`).
2. Worker validates `X-Backend-Secret`.
3. Worker forwards to `https://internal.requiems.xyz/...`.
4. Cloudflare → Caddy on VPS → Go API.
