# Private Deployments — Design Plan

## Overview

This document describes the design for the **Private Deployment** product tier:
dedicated, per-tenant Go API instances hosted on Hetzner VPS, accessible at
`{slug}.requiems.xyz`.

Unlike the existing shared plans (Free → Professional) which go through the
Cloudflare auth-gateway with per-request quotas, a private deployment is a fully
isolated environment the tenant owns exclusively. They call the Go backend
directly over HTTPS with no Cloudflare layer in front.

---

## Why This Exists

High-volume businesses (CRM hygiene, registration protection, data-warehouse
processing) need predictable flat-rate costs instead of per-request billing.
Competitors like BulkEmailChecker sell "unlimited verifications" priced by
concurrent threads. We extend that idea: sell dedicated infrastructure where the
tenant controls which endpoints they need and pays for the server tier, not per
call.

---

## Request Flow (per-tenant vs. shared)

**Shared (current):**

```text
Client → Cloudflare Auth Gateway (api.requiems.xyz)
           ├─ validates API key (KV)
           ├─ checks rate limits (KV counters)
           ├─ records usage (D1)
           └─ proxies to → Go backend (internal, protected by X-Backend-Secret)
```

**Private Deployment (new):**

```text
Client → Caddy (HTTPS, {slug}.requiems.xyz)
              └─ reverse-proxies to → Go backend (port 8080)
                                         X-Backend-Secret: {tenant-secret}
                                         ENABLED_SERVICES=email,tech,...
```

No Cloudflare Worker, no usage tracking, no quota enforcement. The tenant
includes `X-Backend-Secret: {their-secret}` in every request. The Go API is
otherwise identical.

---

## Server Tiers

Priced against Hetzner AMD VPS nodes with a healthy margin:

| Tier       | Hetzner Node | vCPU   | RAM   | SSD    | Monthly | Yearly (eff/mo) | LemonSqueezy charge   |
| ---------- | ------------ | ------ | ----- | ------ | ------- | --------------- | --------------------- |
| Starter    | CPX21        | 3 AMD  | 4 GB  | 80 GB  | $200/mo | $170/mo         | $2,040 once/year      |
| Growth     | CPX31        | 4 AMD  | 8 GB  | 160 GB | $300/mo | $255/mo         | $3,060 once/year      |
| Scale      | CPX41        | 8 AMD  | 16 GB | 240 GB | $500/mo | $425/mo         | $2,550 every 6 months |
| Enterprise | CPX51        | 16 AMD | 32 GB | 360 GB | $960/mo | $816/mo         | $4,896 every 6 months |

> Scale and Enterprise use semi-annual billing to stay under LemonSqueezy's
> $5,000 product price cap.

---

## Endpoint Selection

Tenants choose which service domains they want mounted. All 10 are available as
checkboxes:

| Key             | What it includes                                                                     |
| --------------- | ------------------------------------------------------------------------------------ |
| `email`         | Validate, disposable check, normalize                                                |
| `text`          | Quotes, dictionary, lorem ipsum, advice, profanity, spellcheck                       |
| `tech`          | IP info/VPN/ASN, phone, password gen, QR/barcode, domain, WHOIS, MX, user-agent      |
| `places`        | Cities, geocoding, reverse-geocoding, timezone, holidays, postal codes, working days |
| `finance`       | BIN lookup, crypto, exchange rates, inflation, IBAN, SWIFT, commodities, mortgage    |
| `entertainment` | Jokes, facts, horoscope, trivia, emoji, sudoku, chuck-norris                         |
| `ai`            | Similarity, sentiment, language detection                                            |
| `convert`       | Base64, color, markdown, numbase, format                                             |
| `fitness`       | Exercises database                                                                   |
| `misc`          | Counter, random-user, unit conversion                                                |

This is implemented via the `ENABLED_SERVICES` env var on the Go process (empty
= all services, comma-separated list = only those mounted). Unselected services
simply return 404.

---

## Customer Flow

```text
1. Customer logs in → visits /private-deployment
2. Selects: billing cycle (monthly/yearly) + server tier + endpoint checkboxes + contact info
3. Submits form
      └─ DB: PrivateDeploymentRequest created (status: pending_payment)
      └─ Redirect → LemonSqueezy checkout
            custom_data: { private_deployment_request_id, user_id }

4. Customer pays on LemonSqueezy
      └─ Webhook: subscription_created fires
            ├─ Detects private_deployment_request_id in custom_data
            ├─ status → pending, lemonsqueezy_subscription_id saved
            ├─ Email → customer: "Request received, deployment starting soon"
            └─ Email → admin (OBSERVER_EMAILS): "New paid deployment — action required"

5. [Manual] Admin provisions Hetzner VPS:
      - docker compose: Go API + Postgres + Redis + Caddy
      - ENABLED_SERVICES="{selected services}"
      - BACKEND_SECRET="{generated unique secret}"
      - DNS: {slug}.requiems.xyz → VPS IP

6. Admin visits /admin/private_deployments/{id}
      - Enters: subdomain_slug + tenant_secret + optional notes
      - Clicks "Mark as Deployed"
            ├─ status → active, deployed_at → now
            └─ Email → customer: "Your API is live!"
                   • URL: https://{slug}.requiems.xyz
                   • Auth header: X-Backend-Secret: {secret}
                   • Example curl
                   • Link to API docs
```

---

## Implementation Scope

### 1. Go Backend — `apps/api/`

**`platform/config/config.go`**

- Add `EnabledServices string` field
- Load from `ENABLED_SERVICES` env var (default empty = all services)

**`app/routes_v1.go`**

- Add `serviceEnabled(cfg config.Config, key string) bool` helper — returns true
  when `EnabledServices` is empty or contains the key
- Wrap each of the 10 `r.Mount()` calls with
  `if serviceEnabled(cfg, "email") { ... }`

No auth changes. `BackendSecretAuth` middleware is unchanged; tenants just
supply their own unique `BACKEND_SECRET`.

---

### 2. Rails Dashboard — `apps/dashboard/`

#### Database

New table: `private_deployment_requests`

| Column                | Type       | Notes                                              |
| --------------------- | ---------- | -------------------------------------------------- |
| `user_id`             | references | FK to users, not null                              |
| `company`             | string     | not null                                           |
| `contact_name`        | string     | not null                                           |
| `contact_email`       | string     | not null                                           |
| `server_tier`         | string     | starter/growth/scale/enterprise                    |
| `monthly_price_cents` | integer    | set from tier on create                            |
| `selected_services`   | jsonb      | `["email","text"]`                                 |
| `subdomain_slug`      | string     | set by admin, unique                               |
| `tenant_secret`       | string     | set by admin                                       |
| `status`              | string     | pending_payment/pending/deploying/active/cancelled |
| `admin_notes`         | text       | optional                                           |
| `deployed_at`         | datetime   | set on activate                                    |
| `timestamps`          |            |                                                    |

#### Model — `app/models/private_deployment_request.rb`

- `belongs_to :user`
- Validates tier, status, services length ≥ 1, subdomain_slug format
- `TIER_PRICES` hash (cents)
- `#live_url` helper

#### Public Controller — `app/controllers/private_deployments_controller.rb`

- `before_action :authenticate_user!`
- `GET  /private-deployment` → `new`
- `POST /private-deployment` → `create` — saves request, fires two emails,
  redirects to dashboard with notice

#### Admin Controller — `app/controllers/admin/private_deployments_controller.rb`

- `GET  /admin/private_deployments` → `index` (filter by status)
- `GET  /admin/private_deployments/:id` → `show`
- `PATCH /admin/private_deployments/:id/activate` — fills slug/secret, status →
  active, sends `deployment_ready` email
- `PATCH /admin/private_deployments/:id/cancel` — status → cancelled

#### Mailer — `app/mailers/private_deployment_mailer.rb`

| Method                        | To                | Purpose                           |
| ----------------------------- | ----------------- | --------------------------------- |
| `request_received(request)`   | `contact_email`   | Confirm receipt, set expectations |
| `admin_notification(request)` | `OBSERVER_EMAILS` | Alert admin with link to panel    |
| `deployment_ready(request)`   | `contact_email`   | Go-live email with URL + secret   |

#### Routes — `config/routes.rb`

```ruby
# Public (login required)
get  "private-deployment", to: "private_deployments#new",    as: "new_private_deployment"
post "private-deployment", to: "private_deployments#create", as: "private_deployments"

# Admin
namespace :admin do
  resources :private_deployments, only: [:index, :show] do
    member do
      patch :activate
      patch :cancel
    end
  end
end
```

#### Views

- `app/views/private_deployments/new.html.erb` — 3-section form: tier cards →
  endpoint checkboxes → contact info
- `app/views/admin/private_deployments/index.html.erb` — table with status tabs
- `app/views/admin/private_deployments/show.html.erb` — detail + activate/cancel
  form
- Email templates: `request_received`, `admin_notification`, `deployment_ready`
  (HTML + text pairs)

#### Pricing Page Update

Add a "Private Deployment" callout below the existing plan cards in the pricing
page partials, with a CTA linking to `/private-deployment`.

---

## Future Work: Zero-Downtime Updates with Kamal

Once there are multiple tenants, each Hetzner VPS becomes a Kamal deployment
target. When a new Go version ships, `kamal deploy` rolls it out to all tenant
servers with zero downtime. Each tenant has its own env file with its unique
`BACKEND_SECRET` and `ENABLED_SERVICES`.

---

## Non-Goals (out of scope for this iteration)

- Automated Hetzner provisioning — deployment stays manual for now
- Per-tenant usage tracking or billing automation
- Multi-region tenant deployments
- Tenant self-service management portal
