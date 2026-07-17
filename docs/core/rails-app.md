# The Ruby on Rails App

## Overview

The Rails application serves as the public-facing dashboard and admin panel for
Requiem API, handling user authentication, API key management, usage statistics,
and billing.

**Location:** [apps/dashboard/](../apps/dashboard/) **Technology:** Ruby on
Rails 8.1, Tailwind CSS, Hotwire (Turbo + Stimulus) **Database:** PostgreSQL
(shared with Go backend) **Background Jobs:** Sidekiq with Redis

## Responsibilities

### Public Website

- Landing page ([requiems.xyz](https://requiems.xyz))
- Marketing content and feature showcase
- Documentation portal
- Live API playground

### User Dashboard

- User authentication and registration (Devise)
- API key generation and management
- Usage statistics and analytics
- Billing and subscription management
- Account settings

### Admin Panel

- User management
- API key oversight
- Usage monitoring
- Abuse detection and reporting
- System health monitoring

## Architecture

```
┌─────────────────────────────────────┐
│         Rails App (Port 3000)       │
│                                     │
│  ┌──────────────────────────────┐   │
│  │     Controllers              │   │
│  │  - HomeController            │   │
│  │  - DashboardController       │   │
│  │  - ApiKeysController         │   │
│  └──────────────────────────────┘   │
│                                     │
│  ┌──────────────────────────────┐   │
│  │     Models                   │   │
│  │  - User (Devise)             │   │
│  │  - ApiKey                    │   │
│  │  - Subscription              │   │
│  │  - UsageLog                  │   │
│  └──────────────────────────────┘   │
│                                     │
│  ┌──────────────────────────────┐   │
│  │     Services                 │   │
│  │  - D1SyncService              │   │
│  │  - Cloudflare::ApiManagementService │
│  └──────────────────────────────┘   │
│                                     │
│  ┌──────────────────────────────┐   │
│  │     Background Jobs          │   │
│  │  - SyncD1UsageJob             │   │
│  │  - AggregateDailyUsageJob     │   │
│  │  - ExpirePromotionalSubscriptionsJob │
│  └──────────────────────────────┘   │
└─────────────────────────────────────┘
         ↓                    ↓
    PostgreSQL             Redis
```

## Database Tables

### User Management

- `users` - User accounts (email, password, plan, etc.)
- `api_keys` - API keys associated with users
- `subscriptions` - Billing and subscription data

### Usage Tracking

- `usage_logs` - Detailed API usage records
- `daily_usage_summaries` - Aggregated daily usage
- `credit_adjustments` - Manual request quota adjustments

### Administration

- `audit_logs` - System activity audit trail
- `abuse_reports` - Flagged suspicious activity

## Key Features

### API Key Management

Users can:

- Generate new API keys
- View key statistics (requests used, requests remaining)
- Regenerate keys
- Delete keys
- Set key permissions (future feature)

When a key is created, updated, or deleted, it's managed through the API
Management worker via
[CloudflareApiManagementService](../apps/dashboard/app/services/cloudflare/api_management_service.rb).

### Usage Analytics

- Real-time usage tracking
- Historical usage charts
- Request consumption breakdown
- Top endpoints by usage
- Monthly/daily aggregations

### Background Jobs (Sidekiq-Cron)

**[SyncD1UsageJob](../apps/dashboard/app/jobs/sync_d1_usage_job.rb)**

- Pulls usage data from Cloudflare D1 into PostgreSQL
- Runs every 3 minutes via Sidekiq-Cron (`config/sidekiq_schedule.yml`)

**[AggregateDailyUsageJob](../apps/dashboard/app/jobs/aggregate_daily_usage_job.rb)**

- Builds daily usage summaries for analytics/reporting
- Runs once per day at 00:05 UTC via Sidekiq-Cron

**[ExpirePromotionalSubscriptionsJob](../apps/dashboard/app/jobs/expire_promotional_subscriptions_job.rb)**

- Reverts expired admin-granted promotional plan upgrades back to free
- Runs hourly via Sidekiq-Cron

See [Background Jobs](./background-jobs.md) for full details.

## Development

### Local Setup with Docker

```bash
cd infra/docker
docker compose -f docker-compose.dev.yml up
```

The Rails app runs on http://localhost:3000 with:

- Hot reloading (changes reflected immediately)
- Tailwind CSS compilation
- Sidekiq worker running
- PostgreSQL and Redis available

### Editor support

For the best development experience when working on the Rails dashboard, install
a Ruby LSP editor/IDE extension and point it at the repository. The project
includes the `ruby-lsp` gem in `apps/dashboard/Gemfile`; configuring your editor
to use this language server enables code navigation, in-editor diagnostics, and
formatting features.

### Running Commands

```bash
# Rails console
docker compose -f docker-compose.dev.yml exec dashboard rails console

# Database migrations
docker compose -f docker-compose.dev.yml exec dashboard rails db:migrate

# Run tests
docker compose -f docker-compose.dev.yml exec dashboard rails test

# Create admin user
docker compose -f docker-compose.dev.yml exec dashboard rails runner "
User.create!(
  name: 'Admin',
  email: 'admin@requiems.xyz',
  password: 'password123',
  password_confirmation: 'password123',
  admin: true,
  confirmed_at: Time.current
)
"
```

### Environment Variables

Required in production:

- `RAILS_ENV` - `production`
- `DATABASE_URL` - PostgreSQL connection string
- `REDIS_URL` - Redis connection string
- `SECRET_KEY_BASE` - Rails secret (generate with `rails secret`)
- `RAILS_MASTER_KEY` - Encrypted credentials key
- `API_MANAGEMENT_API_KEY` - Authenticates Rails to the API Management worker
- `LEMONSQUEEZY_STORE_ID`, `LEMONSQUEEZY_SIGNING_SECRET` (+ `_TEST` variant) -
  LemonSqueezy billing integration
- `LEMONSQUEEZY_{DEVELOPER,BUSINESS,PROFESSIONAL}_{MONTHLY,YEARLY}_VARIANT_ID` /
  `_CHECKOUT_UUID` (+ `_TEST` variants) - per-plan billing config

See [AppConfig](../apps/dashboard/app/lib/app_config.rb) for the full list this
app requires at boot.

## Database Migrations

Rails uses separate migration tracking from the Go backend to avoid conflicts:

**Go migrations table:** `schema_migrations` **Rails migrations table:**
`schema_migrations` (separate schema)

Both apps share the same PostgreSQL database but maintain independent migration
histories.

## Styling

- **Tailwind CSS** - Utility-first CSS framework
- **Hotwire** - Modern SPA-like experience without JavaScript frameworks
  - **Turbo Drive** - Fast navigation
  - **Turbo Frames** - Partial page updates
  - **Stimulus** - JavaScript sprinkles

### Syntax Highlighting

Always use the `highlight` Stimulus controller (`highlight_controller.js`) —
never add a separate hljs call or import.

The hljs CSS is loaded via `<link>` tag in both layouts (`application.html.erb`
and `dashboard.html.erb`). Do **not** use `@import` for it in `application.css`
— Tailwind v4 strips external URL imports during compilation.

**Standalone code block:**

```erb
<div data-controller="highlight">
  <pre class="bg-gray-900 text-gray-100 rounded-lg p-4 text-sm overflow-x-auto"><code class="language-bash">your code here</code></pre>
</div>
```

**With code tabs** — add `highlight` alongside `code-tabs` on the same element:

```erb
<div data-controller="code-tabs highlight">
  ...tabs and content blocks...
</div>
```

The controller runs `hljs.highlightElement()` on every `pre code` inside the
element on connect. Always set `class="language-{lang}"` on the `<code>` tag
(`language-bash` for curl/shell, `language-javascript`, `language-python`,
`language-go`, etc.).

## Testing

```bash
# Run all tests
docker compose -f docker-compose.dev.yml exec dashboard rails test

# Run specific test
docker compose -f docker-compose.dev.yml exec dashboard rails test test/models/api_key_test.rb

# Run system tests (browser tests)
docker compose -f docker-compose.dev.yml exec dashboard rails test:system
```

## Deployment

### With Docker Compose (VPS)

```bash
cd infra/docker
docker compose up -d --build
```

## Cloudflare Integration

### Client IP Resolution (Cloudflare Dependency)

The Rails dashboard sits behind Cloudflare as a reverse proxy. This means
`request.remote_ip` in Rails returns Cloudflare's **edge node IP**, not the
user's real browser IP. Rails has no `config.action_dispatch.trusted_proxies`
configured, so it does not strip Cloudflare from the forwarded chain.

Cloudflare injects the real client IP via the `CF-Connecting-IP` header on every
request it forwards. Code that needs the actual client IP must read this header
explicitly:

```ruby
real_ip = request.headers["CF-Connecting-IP"] || request.remote_ip
```

The fallback to `request.remote_ip` handles local development and any traffic
that reaches Rails directly (e.g. internal health checks, Docker networking).

**Affected code:**

- `ApiProxyController` — forwards `X-Forwarded-For` to the Go backend so the IP
  endpoint (`GET /v1/networking/ip`) returns the user's real IP instead of a
  Cloudflare datacenter IP.

**Note:** The auth gateway (`apps/workers/auth-gateway/`) has the same
constraint — it reads `CF-Connecting-IP` from Cloudflare and forwards it as
`X-Forwarded-For` when proxying to the Go backend. See
[auth-gateway.md](./auth-gateway.md).

**If Cloudflare is removed:** replace `CF-Connecting-IP` reads with proper
`config.action_dispatch.trusted_proxies` pointing at the upstream load balancer
/ reverse proxy, and update `filterHeaders` in the auth gateway similarly.

### API Management Service

API key operations (create, update, revoke) are sent to the API Management
worker via
[CloudflareApiManagementService](../apps/dashboard/app/services/cloudflare/api_management_service.rb),
which keeps KV in sync. Rails never writes to KV directly.

### Usage Data Flow

```
User makes API request
    ↓
Worker records usage in D1
    ↓
Sidekiq job pulls from D1 (every 3 minutes)
    ↓
Usage stored in PostgreSQL
    ↓
Displayed in Rails dashboard
```

## Performance Considerations

- **Database Indexes:** All foreign keys and frequently queried columns indexed
- **Caching:** Fragment caching for dashboard views
- **Background Jobs:** Long-running operations moved to Sidekiq
- **Asset Pipeline:** Tailwind CSS compiled and minified
- **CDN:** Static assets served via CDN in production

## Related Documentation

- [Architecture Overview](./architecture.md)
- [Auth Gateway Documentation](./auth-gateway.md)
- [Infrastructure Guide](./infrastructure.md)
