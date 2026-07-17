# Deployment Guide

Maintainer trick:

```sh
cd ~/server/requiems-api/ && git pull && cd infra/docker && docker compose up -d --build
```

## Overview

This guide covers deploying Requiem API to production on a Hetzner VPS with
Cloudflare Workers for edge authentication.

## Architecture

```
User → Cloudflare Worker (Edge) → Hetzner VPS (Backend)
       (auth, rate limiting)      (API, database, dashboard)
```

**Components to Deploy:**

1. Cloudflare Worker — Auth Gateway (apps/workers/auth-gateway)
2. Cloudflare Worker — API Management (apps/workers/api-management)
3. VPS Backend (Go API + Rails Dashboard + PostgreSQL + Redis)
4. DNS and domain configuration

---

## Pre-Deployment Checklist

- [ ] Hetzner VPS created and accessible via SSH
- [ ] Domain purchased and added to Cloudflare
- [ ] GitHub repository accessible
- [ ] Environment variables documented
- [ ] Database backup strategy planned
- [ ] SSL certificates configured (handled by Caddy)

---

## Part 1: VPS Setup (Hetzner)

### 1.1 Create VPS

**Recommended Configuration:**

- **Server Type:** CPX21 (3 vCPU, 4GB RAM, 80GB SSD) - €7.95/month
- **Location:** Choose based on audience
  - Nuremberg (nbg1) - Europe
  - Falkenstein (fsn1) - Europe (Germany)
  - Helsinki (hel1) - Northern Europe
  - Ashburn (ash) - US East
- **Image:** Ubuntu 24.04 LTS
- **Networking:** IPv4 (required), IPv6 (optional)
- **Backups:** Enable for $1.19/month (recommended)

### 1.2 Initial Server Setup

```bash
# SSH into server
ssh root@YOUR_VPS_IP

# Update system
apt update && apt upgrade -y

# Install Docker
curl -fsSL https://get.docker.com -o get-docker.sh
sh get-docker.sh
rm get-docker.sh

# Install Docker Compose plugin
apt install docker-compose-plugin -y

# Verify installation
docker --version
docker compose version
```

### 1.3 Configure Firewall

```bash
# Install UFW
apt install ufw

# Allow SSH, HTTP, HTTPS
ufw allow 22/tcp   # SSH
ufw allow 80/tcp   # HTTP
ufw allow 443/tcp  # HTTPS

# Enable firewall
ufw enable

# Verify
ufw status verbose
```

### 1.4 Create Deploy User (Optional but Recommended)

```bash
# Create user
adduser deploy

# Add to docker group
usermod -aG docker deploy

# Switch to deploy user
su - deploy
```

### 1.5 Clone Repository

```bash
# Using HTTPS
git clone https://github.com/bobadilla-tech/requiems-api.git
cd requiems-api

# Or using SSH (if you've added deploy key to GitHub)
git clone git@github.com:bobadilla-tech/requiems-api.git
cd requiems-api
```

---

## Part 2: Environment Configuration

### 2.1 Create Environment File

```bash
# Navigate to docker directory
cd infra/docker

# Create .env file
nano .env
```

**Production .env file:**

```bash
# PostgreSQL Database
POSTGRES_USER=requiem
POSTGRES_PASSWORD=CHANGE_THIS_TO_STRONG_PASSWORD
POSTGRES_DB=requiem

DATABASE_URL=postgres://requiem:CHANGE_THIS_TO_STRONG_PASSWORD@db:5432/requiem?sslmode=disable

# Redis
REDIS_URL=redis://redis:6379

# Rails Security
SECRET_KEY_BASE=your_secret_key_base_here
RAILS_MASTER_KEY=your_rails_master_key_here

# Backend Secret (32+ characters)
# Must match Cloudflare Worker BACKEND_SECRET
BACKEND_SECRET=your_32_char_minimum_secret_here

# API Management Worker (Rails -> api-management, never writes to KV directly)
API_MANAGEMENT_URL=https://api-management.requiems.xyz
API_MANAGEMENT_API_KEY=must_match_api_management_worker_secret

# LemonSqueezy billing (store ID and variant IDs are public - see .env.example
# for the full list of per-plan VARIANT_ID / CHECKOUT_UUID vars)
LEMONSQUEEZY_STORE_ID=your_store_id
LEMONSQUEEZY_SIGNING_SECRET=your_webhook_signing_secret
LEMONSQUEEZY_API_KEY=your_lemonsqueezy_api_key

# SMTP (Devise + ApplicationMailer)
SMTP_FROM_EMAIL=noreply@mail.requiems.xyz
```

**Note:** `AppConfig#load_config` hard-requires the LemonSqueezy store ID,
signing secret, and all six variant ID / checkout UUID pairs at boot — Rails
will not start without them. See `infra/docker/.env.example` for the complete
list.

**Generate secrets:**

```bash
# SECRET_KEY_BASE (Rails)
docker run --rm ruby:3.4-alpine sh -c "gem install rails && rails secret"

# BACKEND_SECRET (Worker authentication)
openssl rand -base64 48

# RAILS_MASTER_KEY
# Found in apps/dashboard/config/master.key (create if missing)
```

### 2.2 Update Caddyfile

```bash
nano ../caddy/Caddyfile
```

**Update with your domain:**

```caddyfile
{
  email admin@requiems.xyz
}

# Main application (landing page, dashboard, admin, docs)
requiems.xyz {
  encode gzip
  reverse_proxy dashboard:80
}

# Internal backend API (only accessible via Cloudflare Worker)
internal.requiems.xyz {
  encode gzip

  # Add X-Backend-Secret header verification
  @authorized {
    header X-Backend-Secret {env.BACKEND_SECRET}
  }

  handle @authorized {
    reverse_proxy api:8080
  }

  handle {
    respond "Unauthorized" 403
  }
}
```

---

## Part 3: Deploy Backend Services

### 3.1 Build and Start Services

```bash
cd infra/docker

# Start all services (builds and starts automatically)
docker compose up -d --build

# Check status
docker compose ps

# View logs
docker compose logs -f
```

**Services should be running:**

- ✅ api (Go backend)
- ✅ db (PostgreSQL)
- ✅ redis (Redis)
- ✅ dashboard (Rails)
- ✅ worker (Sidekiq background jobs)
- ✅ caddy (reverse proxy with auto-HTTPS)

**Note:** Migrations run automatically:

- Rails migrations execute on dashboard startup (via `db:migrate` in startup
  command)
- Go migrations execute on API startup (embedded in the application)

### 3.2 Verify Services

```bash
# Check API health
docker compose exec api wget -qO- http://localhost:8080/healthz

# Check Rails
docker compose exec dashboard curl http://localhost:80

# Check logs
docker compose logs api
docker compose logs dashboard
docker compose logs caddy
```

---

## Part 4: DNS Configuration (Cloudflare)

### 4.1 Add Domain to Cloudflare

1. Go to [Cloudflare Dashboard](https://dash.cloudflare.com)
2. Click "Add a Site"
3. Enter your domain (e.g., `requiems.xyz`)
4. Select Free plan
5. Copy the Cloudflare nameservers

### 4.2 Update Nameservers at Registrar

Go to your domain registrar and update nameservers to Cloudflare's:

- `ns1.cloudflare.com`
- `ns2.cloudflare.com`

Wait for DNS propagation (5-30 minutes).

### 4.3 Create DNS Records

In Cloudflare DNS settings, create these A records:

| Type | Name     | Content     | Proxy Status | TTL  |
| ---- | -------- | ----------- | ------------ | ---- |
| A    | @        | YOUR_VPS_IP | ☁️ Proxied   | Auto |
| A    | internal | YOUR_VPS_IP | ☁️ Proxied   | Auto |

The Workers custom domains (`api.requiems.xyz`, `api-management.requiems.xyz`)
are configured in `wrangler.toml` and created automatically when you deploy each
Worker.

**Important:**

- Both domains should be **Proxied (orange cloud)** for:
  - DDoS protection
  - Cloudflare CDN caching
  - SSL/TLS termination
- Caddy will obtain Let's Encrypt certificates automatically using HTTP-01
  challenge (works with Cloudflare proxy)

### 4.4 Verify DNS

```bash
# Check DNS propagation
dig requiems.xyz
dig internal.requiems.xyz

# Or use nslookup
nslookup requiems.xyz
nslookup internal.requiems.xyz
```

---

## Part 5: Deploy Cloudflare Workers

Both workers share the same KV namespace and D1 database, so you only create
those once.

### 5.1 Setup Wrangler

```bash
# Install wrangler globally
npm install -g wrangler

# Login to Cloudflare
wrangler login
```

### 5.2 Create Shared Edge Resources (first deploy only)

```bash
cd apps/workers/auth-gateway

# Create KV namespace for API keys and rate limiting
wrangler kv namespace create KV --env production
# Output: { binding = "KV", id = "abc123..." }
# → Update id in both wrangler.toml files

# Create D1 database for usage tracking
wrangler d1 create requiem-usage --env production
# Output: { database_id = "xyz456..." }
# → Update database_id in both wrangler.toml files

# Apply D1 schema
pnpm run db:migrate:prod
```

### 5.3 Deploy Auth Gateway

The Auth Gateway is the public-facing edge service at `api.requiems.xyz`.

```bash
cd apps/workers/auth-gateway

# Set secrets
wrangler secret put BACKEND_URL --env production
# Enter: https://internal.requiems.xyz

wrangler secret put BACKEND_SECRET --env production
# Enter: (same value as VPS .env BACKEND_SECRET)

# Deploy
pnpm run deploy:prod
```

### 5.4 Deploy API Management Worker

The API Management Worker is the internal service used only by the Rails
dashboard. It is reachable at `api-management.requiems.xyz`.

```bash
cd apps/workers/api-management

# Set secrets
wrangler secret put API_MANAGEMENT_API_KEY --env production
# Enter: a strong random key (32+ chars) — Rails uses this to authenticate
# Generate: openssl rand -base64 48

# Deploy
pnpm run deploy:prod
```

Set the matching key in the Rails dashboard `.env` on the VPS:

```bash
API_MANAGEMENT_URL=https://api-management.requiems.xyz
API_MANAGEMENT_API_KEY=same_key_as_above
```

Then restart the dashboard and sidekiq:

```bash
cd infra/docker
docker compose restart dashboard worker
```

### 5.5 Verify Worker Routes

Both workers use `custom_domain = true` in `wrangler.toml` — Cloudflare
registers the hostnames automatically on first deploy. Confirm in the Cloudflare
Dashboard:

1. Workers & Pages → `requiem-auth-gateway-production` → Settings → Domains &
   Routes
   - Should show: `api.requiems.xyz`
2. Workers & Pages → `requiem-api-management-production` → Settings → Domains &
   Routes
   - Should show: `api-management.requiems.xyz`

---

## Part 6: Verification

### 6.1 Test Full Request Flow

```bash
# 1. Test backend directly (should fail without secret header)
curl https://internal.requiems.xyz/healthz

# 2. Get a test API key from Rails dashboard
# Sign up at https://requiems.xyz

# 3. Test through Worker (should work)
curl -H "requiems-api-key: YOUR_API_KEY" https://api.requiems.xyz/v1/text/advice

# Should return:
# {
#   "advice": "Don't be afraid to make mistakes.",
#   "id": 42
# }
```

### 6.2 Check Response Headers

```bash
curl -I -H "requiems-api-key: YOUR_KEY" https://api.requiems.xyz/v1/text/advice
```

Should include:

```
X-Requests-Used: 1
X-Requests-Remaining: 49
X-Requests-Reset: 2026-02-10T00:00:00.000Z
X-Plan: free
X-RateLimit-Limit: 30
X-RateLimit-Remaining: 29
```

### 6.3 Test Rate Limiting

```bash
# Send multiple requests rapidly
for i in {1..35}; do
  curl -H "requiems-api-key: YOUR_KEY" https://api.requiems.xyz/v1/text/advice
done

# Should receive 429 Too Many Requests after 30 requests (free tier)
```

---

## Part 7: Post-Deployment

### 7.1 Seed Static Data

Some endpoints require data to be seeded after the first deploy. Migrations
create the tables automatically on startup, but the data must be populated
manually once.

Seed binaries are compiled into the production image at `/app/seed-*`. The
production image does not include the Go toolchain, so `go run` will not work.

**Run all seeds in order:**

```bash
cd infra/docker

# BIN/IIN card data (~700k records — takes ~5 min)
docker compose exec api sh -c '/app/seed-bins --db-url "$DATABASE_URL"'

# Inflation data (World Bank CPI, ~241 countries × 30 years)
docker compose exec api sh -c '/app/seed-inflation --db-url "$DATABASE_URL"'

# IBAN country data
docker compose exec api sh -c '/app/seed-iban --db-url "$DATABASE_URL"'

# Commodity prices (FRED + Yahoo Finance, 16 commodities × ~30 years)
docker compose exec api sh -c '/app/seed-commodities --db-url "$DATABASE_URL"'

# SWIFT/BIC codes (~100k+ bank records)
docker compose exec api sh -c '/app/seed-swift --db-url "$DATABASE_URL"'
```

`DATABASE_URL` is set inside the container via `env_file: .env`. The `sh -c`
wrapper ensures the variable is expanded inside the container, not on the host.

**Seeding exercises** — requires uploading the JSON data file first:

```bash
docker cp ~/exercises.json requiem-backend-api-1:/tmp/exercises.json
docker exec requiem-backend-api-1 sh -c '/app/seed-exercises --data-dir /tmp --db-url "$DATABASE_URL"'
```

**Re-seeding is safe** — all seed commands use
`INSERT ... ON CONFLICT DO UPDATE`, so running them again only updates changed
records.

**Keeping commodity prices current** — commodity prices are annual averages from
FRED and Yahoo Finance. Re-run the commodity seed once per year (or whenever you
want fresh data):

```bash
docker compose exec api sh -c '/app/seed-commodities --db-url "$DATABASE_URL"'
```

---

### 7.3 Create Admin User

```bash
docker compose exec dashboard rails runner "
User.create!(
  name: 'Admin',
  email: 'admin@requiems.xyz',
  password: 'SECURE_PASSWORD_HERE',
  password_confirmation: 'SECURE_PASSWORD_HERE',
  admin: true,
  confirmed_at: Time.current
)
"
```

### 7.4 Setup Monitoring

**Cloudflare Analytics:**

- Worker metrics available in Cloudflare Dashboard
- Track requests, errors, latency

**VPS Monitoring:**

```bash
# Resource usage
docker compose stats

# Logs
docker compose logs -f

# Disk usage
df -h
du -sh /var/lib/docker/volumes
```

### 7.5 Setup Backups

**Database Backups (cron job):**

```bash
# Create backup script
nano /home/deploy/backup-db.sh
```

```bash
#!/bin/bash
BACKUP_DIR="/home/deploy/backups"
TIMESTAMP=$(date +%Y%m%d_%H%M%S)

mkdir -p $BACKUP_DIR

docker compose -f /home/deploy/requiems-api/infra/docker/docker-compose.yml \
  exec -T db pg_dump -U requiem requiem > $BACKUP_DIR/backup_$TIMESTAMP.sql

# Keep only last 7 days
find $BACKUP_DIR -name "backup_*.sql" -mtime +7 -delete
```

```bash
# Make executable
chmod +x /home/deploy/backup-db.sh

# Add to crontab
crontab -e

# Add line (daily at 3 AM):
0 3 * * * /home/deploy/backup-db.sh
```

---

## Environment Variables Summary

### VPS (.env file)

- `SECRET_KEY_BASE` - Rails secret
- `RAILS_MASTER_KEY` - Rails encrypted credentials
- `DATABASE_URL` - PostgreSQL connection
- `REDIS_URL` - Redis connection
- `BACKEND_SECRET` - Worker authentication (32+ chars)
- `LEMONSQUEEZY_STORE_ID`, `LEMONSQUEEZY_SIGNING_SECRET`,
  `LEMONSQUEEZY_API_KEY` - LemonSqueezy billing
- `LEMONSQUEEZY_{DEVELOPER,BUSINESS,PROFESSIONAL}_{MONTHLY,YEARLY}_VARIANT_ID` /
  `_CHECKOUT_UUID` - per-plan billing config (required at boot, see
  `.env.example`)
- `API_MANAGEMENT_URL` - `https://api-management.requiems.xyz`
- `API_MANAGEMENT_API_KEY` - Must match the api-management worker secret
- `SMTP_FROM_EMAIL` - Sender address for all outgoing mail (must use the
  verified Resend domain, e.g. `noreply@mail.requiems.xyz`). Used by both Devise
  and ApplicationMailer.

### Auth Gateway Worker (Wrangler secrets)

- `BACKEND_URL` - Internal backend URL (`https://internal.requiems.xyz`)
- `BACKEND_SECRET` - Must match VPS `BACKEND_SECRET`

### API Management Worker (Wrangler secrets)

- `API_MANAGEMENT_API_KEY` - Must match VPS `API_MANAGEMENT_API_KEY`

---

## Troubleshooting

### Worker can't reach backend

**Check:**

1. DNS for `internal.yourdomain.com` resolves to VPS IP
2. Caddy is running: `docker compose ps caddy`
3. `BACKEND_SECRET` matches between Worker and VPS
4. Firewall allows ports 80/443

### Database connection errors

```bash
# Check PostgreSQL is running
docker compose ps db

# Check logs
docker compose logs db

# Restart database
docker compose restart db
```

### Rails won't start

```bash
# Check logs
docker compose logs dashboard

# Common issues:
# - Missing SECRET_KEY_BASE
# - Missing RAILS_MASTER_KEY
# - Database migration needed

# After fixing, rebuild without cache:
docker builder prune -a -f
docker compose build --no-cache dashboard sidekiq
docker compose up -d
```

### Caddy can't get SSL certificates

**Check:**

1. DNS records pointing to VPS (both domains proxied via Cloudflare is OK)
2. Ports 80/443 open in firewall
3. Domain in Caddyfile matches DNS
4. Wait 2-3 minutes for Let's Encrypt to issue certificates

---

## Updating Production

### Standard Update (Normal Deployment)

Pull changes and rebuild automatically:

```bash
cd ~/server/requiems-api
git pull

cd infra/docker
docker compose up -d --build
```

The `--build` flag rebuilds any changed images automatically. Migrations run
automatically on container startup.

**Check deployment:**

```bash
docker compose ps
docker compose logs -f
```

**If the update adds a new seeded table** (e.g. a new commodity, inflation
country set, or BIN dataset), re-run the relevant seed command after deploy. See
[Section 7.1](#71-seed-static-data).

### Clean Deployment (Wipe Database)

Use this when changing database credentials or starting fresh:

```bash
cd ~/server/requiems-api
git pull

cd infra/docker
docker compose down -v    # -v removes volumes
docker compose up -d --build
```

⚠️ **Warning:** The `-v` flag deletes all database data!

### Docker Cache Issues

If you encounter build errors like "parent snapshot does not exist":

```bash
# Clear build cache
docker builder prune -a -f

# Rebuild everything
docker compose up -d --build
```

Or force rebuild without cache:

```bash
docker compose build --no-cache
docker compose up -d
```

### Update Specific Service

Rebuild only one service:

```bash
docker compose up -d --build api
# or
docker compose up -d --build dashboard
```

### Update Workers

```bash
# Auth Gateway
cd apps/workers/auth-gateway
pnpm run deploy:prod

# API Management
cd apps/workers/api-management
pnpm run deploy:prod
```

### Quick Restart (No Rebuild)

Restart services without rebuilding:

```bash
docker compose restart
# or specific service
docker compose restart api
```

---

## Related Documentation

- [Infrastructure Guide](./infrastructure.md)
- [Local Development](../../infra/readme.md)
- [Docker Setup](../../infra/docker/README.md)
