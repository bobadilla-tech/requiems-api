# Requiems Domain Migration Audit

Status: audited implementation plan; audit-only, no implementation performed.

Scope amendment: the project owner confirms there are currently no users,
customers, or meaningful public positioning. Backwards compatibility, SEO
preservation, old-domain redirects, and customer migration are therefore
explicitly out of scope unless a later decision re-enables them.

Audit date: 2026-08-21

Repository: `requiems-api`, branch `v2/requiems-api-v2`

## 1. Executive Summary

The repository does not currently implement the requested target architecture.
The current repository-described production topology is split across five
named hostnames and two
documented VPS deployment paths:

- `requiems.xyz` serves the Rails public website, authentication, dashboard,
  admin pages, webhooks, docs, SEO assets, and server-side API playground.
- `api.requiems.xyz` is a Cloudflare Worker custom domain for the authenticated
  public API gateway. It validates `requiems-api-key`, enforces rate and quota
  limits, records usage, and forwards `/v1/*` to the private Go API.
- `internal.requiems.xyz` is the secret-guarded origin for the Go API.
- `api-management.requiems.xyz` is an internal API-management Worker used by
  Rails for API-key CRUD and usage operations.
- `mcp.requiems.xyz` serves the MCP HTTP server, whose generated tools call the
  current API base URL.

The target requires one apex host, `requiems.xyz`, to have path-dependent
behavior: `/v1/*` must execute the API gateway while all other requests must
redirect to `https://requiemsapi.com`. The repository currently has no such
host/path split. The current apex is routed to Rails by Kamal/Caddy, while the
API Worker owns all paths on the separate `api.requiems.xyz` custom domain.

Overall risk is MEDIUM for a clean break, rising to HIGH only if the active
ingress, DNS/TLS ownership, payment-provider callback, or API security boundary
is changed without verification. The recommended clean-break posture is:

1. move Rails to `requiemsapi.com` as the only public product/dashboard host;
2. expose the current API implementation only at `requiems.xyz/v1/*` (plus
   explicitly selected health/spec endpoints);
3. retire `api.requiems.xyz` without redirecting it after internal references
   and provider settings are updated;
4. do not configure generic old-page redirects or SEO migration machinery;
5. return an explicit 404/410 for non-API paths on `requiems.xyz`, unless a
   later product decision restores the root redirect for convenience.

This is a recommendation, not a completed migration. The exact Cloudflare
route and DNS records must be confirmed against the production account before
any configuration is changed.

The main unresolved implementation gates are the active Cloudflare/VPS ingress,
the exact new API route owner, the payment webhook cutover, the Worker deploy/
rollback path, and whether the separate Go-authentication architecture work is
included or sequenced later.

## 2. Current Architecture

### 2.1 Domain and service map

| Host | Repository evidence | Current responsibility | Confidence |
| --- | --- | --- | --- |
| `requiems.xyz` | `infra/kamal/deploy.dashboard.yml:15-18`; `infra/caddy/Caddyfile:5-23`; `apps/dashboard/config/routes.rb:64-214` | Rails website, locale routes, Devise auth, dashboard, admin, API docs, webhooks, playground | High in repository; active ingress needs production verification |
| `api.requiems.xyz` | `apps/workers/auth-gateway/wrangler.toml:38-42`; `apps/workers/auth-gateway/src/index.ts:20-30` | Public Cloudflare auth/rate/quota/usage gateway, including `/v1/*`, `/healthz`, and `/openapi.json` | High |
| `internal.requiems.xyz` | `infra/kamal/deploy.api.yml:10-18`; `infra/caddy/Caddyfile:31-47` | Go API origin, protected by `X-Backend-Secret` | High in repository; active proxy needs verification |
| `api-management.requiems.xyz` | `apps/workers/api-management/wrangler.toml:31-35`; `apps/dashboard/app/services/cloudflare/api_management_service.rb:12-14` | Internal API-key management, usage export, analytics, Swagger | High |
| `mcp.requiems.xyz` | `infra/kamal/deploy.mcp.yml:10-15`; `apps/mcp/src/server.ts:108-150` | MCP HTTP server; generated handlers call the configured Requiems API base | High |
| `requiemsapi.com` | No repository references found | Target product host; not configured in repository | High |

### 2.2 Request flow

The public API request flow is:

```text
Client
  -> api.requiems.xyz (Auth Gateway Worker)
  -> KV API-key lookup and rate limit
  -> D1 quota/usage operations
  -> internal.requiems.xyz with X-Backend-Secret
  -> Go API /v1/*
  -> PostgreSQL and/or Redis
```

The Worker route is host-wide on the current custom domain. `src/index.ts`
exempts `/healthz` and `/openapi.json`, applies CORS, then applies API-key auth
to the remaining paths and mounts a wildcard proxy route. The proxy preserves
the incoming method, path, query, body, and non-sensitive headers while adding
`X-Backend-Secret` (`apps/workers/auth-gateway/src/index.ts:20-30`,
`apps/workers/auth-gateway/src/routes/proxy.ts:26-61`).

The Go application has no public-host routing. It exposes `/healthz` publicly
and mounts all business routes under `/v1` behind `BackendSecretAuth`
(`apps/api/app/app.go:35-51`). It therefore cannot implement the requested
apex redirect/API split by itself.

The Rails public request flow is:

```text
Client
  -> current Rails host requiems.xyz
  -> optional /en, /es, /fr locale scope
  -> public pages, Devise, /dashboard, /admin, /api/proxy, webhooks
```

The Rails playground and tool demos call the Go origin directly through
`INTERNAL_API_URL` and `X-Backend-Secret`; they do not call the public API
gateway (`apps/dashboard/app/services/api_proxy_service.rb:22-46`). This is an
important distinction: changing the public API hostname alone will not change
the server-side playground path.

### 2.3 Routing and deployment sources of truth

There are two production-looking ingress descriptions in the repository:

- Kamal deploys the API to `internal.requiems.xyz`, the dashboard to
  `requiems.xyz`, and MCP to `mcp.requiems.xyz`
  (`infra/kamal/deploy.*.yml`).
- Docker Compose plus Caddy describes the dashboard at `requiems.xyz`, the
  private API at `internal.requiems.xyz`, and MCP at `mcp.requiems.xyz`
  (`infra/docker/docker-compose.yml`, `infra/caddy/Caddyfile`).

The current CD workflow invokes `kamal setup` for API, dashboard, and MCP; it
does not deploy Caddy (`.github/workflows/cd.yml:75-149`,
`.github/actions/kamal-deploy/action.yml:44-46`). The repository therefore
supports a strong inference that Kamal is the active deployment path, but this
must be checked on the VPS and in Cloudflare before migration. The Caddy files
may be legacy, separately operated, or used by a different production mode.

DNS records, Cloudflare Worker routes/custom domains, Cloudflare SSL settings,
Kamal proxy state, VPS firewall rules, and external monitors are not represented
as authoritative infrastructure-as-code in this repository. They are external
dependencies and must be inventoried before implementation.

Cloudflare's current documentation distinguishes a Custom Domain, which owns
all paths for a hostname, from a Route, which runs in front of an existing
proxied origin. A route requires a proxied DNS record; a Custom Domain creates
the Worker-facing DNS/certificate state. This distinction is directly relevant
to the apex split and is not encoded in the current repository.

External platform references used for this conclusion:

- [Cloudflare Workers Routes](https://developers.cloudflare.com/workers/configuration/routing/routes/)
- [Cloudflare Workers Custom Domains](https://developers.cloudflare.com/workers/configuration/routing/custom-domains/)
- [Cloudflare Workers best practices](https://developers.cloudflare.com/workers/best-practices/workers-best-practices/)
- [Cloudflare proxied DNS records](https://developers.cloudflare.com/dns/proxy-status/)

### 2.4 Rails application and public URLs

Rails has an optional locale scope. The public root and content pages are
redirected to explicit locale paths by `ApplicationController#set_locale`
(`apps/dashboard/app/controllers/application_controller.rb:20-42`). Dashboard
routes are nested under that scope (`apps/dashboard/config/routes.rb:64-111`),
so the repository's canonical dashboard URLs are currently forms such as
`/en/dashboard`, not an explicit unlocalized `/dashboard` contract.

Non-locale routes include:

- `/up`
- `/llms.txt`, `/llms-full.txt`, `/apis/:id/index.md`
- `POST /api/proxy`
- `POST /locale`
- tool-demo POST endpoints
- `POST /webhooks/lemonsqueezy`

The target's literal `https://requiemsapi.com/dashboard` therefore requires a
deliberate URL contract decision: add an unlocalized dashboard entry point, or
document the localized URL as canonical and make `/dashboard` redirect to the
selected locale. It must not be left to the current optional-locale behavior.

The target examples are also not the current API route contract. Repository
routes include `/v1/validation/email`, `/v1/validation/phone`,
`/v1/networking/ip`, and `/v1/systems/signup/protect`; the shorter examples in
the requested architecture are illustrative unless product separately approves
aliases. This audit does not turn a domain migration into a path migration.

### 2.5 Authentication, cookies, and redirects

Devise manages browser authentication and Rails includes CSRF meta tags in the
HTML layout (`apps/dashboard/app/views/layouts/application.html.erb:43-60`). No
custom session-store initializer or explicit cookie domain was found. The
default Rails session cookie is therefore expected to be host-scoped unless
the deployed environment overrides it externally.

Moving the Rails app from `requiems.xyz` to `requiemsapi.com` changes the
registrable domain. A host-only cookie will not be sent to the new host, and a
cookie scoped to `.requiems.xyz` cannot be shared with `requiemsapi.com`.
Users should be expected to authenticate again unless a separately designed,
short-lived, one-time cross-domain handoff is implemented. A shared cookie is
not a safe or technically valid solution across these two registrable domains.

No OAuth provider or OAuth callback route was found in the repository. Devise
confirmation, password reset, account deletion, and other mail links use Rails
mailer URL options. The deployed values currently include
`MAILER_HOST: requiems.xyz`, `SMTP_DOMAIN: mail.requiems.xyz`, and
`SMTP_FROM_EMAIL: noreply@mail.requiems.xyz`
(`infra/kamal/deploy.dashboard.yml:60-65`; `apps/dashboard/config/environments/production.rb:26-40`).
The web host must change in links, while the mail transport domain must only
change if email DNS and deliverability are intentionally migrated.

The Lemon Squeezy webhook is a POST route at the old apex
(`apps/dashboard/config/routes.rb:52-54`; `docs/core/lemonsqueezy-webhook-setup.md:56`).
An indiscriminate old-host redirect would risk dropping or altering webhook
requests and must not be used as the first cutover behavior.

### 2.6 API authentication, CORS, and compatibility behavior

The public API uses the `requiems-api-key` request header. The Worker strips that
header before forwarding and adds the backend secret. Worker responses include
`Access-Control-Allow-Origin: *`; the preflight response allows
`Content-Type` and `requiems-api-key` (`apps/workers/shared/src/http.ts:5-15`,
`apps/workers/auth-gateway/src/http.ts:10-33`). The gateway does not use
cookies for authentication, so API authentication is header-based rather than
browser-session based; however, `filterHeaders` currently forwards an incoming
`Cookie` header to the Go origin (`apps/workers/auth-gateway/src/http.ts:16-31`).
A same-origin request from `requiems.xyz` can therefore carry a Rails cookie
across the API boundary. This must be an explicit reviewed decision and a
tested header-filtering rule.

The repository contains no host allowlist in Go and no API-host validation in
the Worker. The edge routing layer is the boundary that determines whether a
request reaches the gateway. This makes an incorrect `requiems.xyz/*` route a
security/availability issue, not merely a URL issue.

### 2.7 Documentation, OpenAPI, MCP, and SDK-like consumers

The old API hostname is embedded in:

- `apps/workers/auth-gateway/scripts/openapi/constants.ts:19-38` and the
  generated `src/generated/openapi.ts`;
- `apps/mcp/scripts/fetch-spec.ts:1-5` and `apps/mcp/openapi.json`;
- `apps/mcp` generated runtime/configuration and production deployment
  (`infra/kamal/deploy.mcp.yml:33-35`);
- 61 API YAML files under `apps/dashboard/config/api_docs` with `base_url`
  and/or examples;
- generated/handwritten `docs/apis/**/*.md` examples;
- README, integration test defaults, load-test examples, and performance
  reports;
- Rails dashboard quick-start, API-key display, API-reference, LLM, and tool
  pages.

The shared Worker website constant (`apps/workers/shared/src/constants.ts:3`),
the OG-image generator (`scripts/gen-og-image.mjs:3,125`), and tracked sitemap
XML files (`apps/dashboard/public/sitemap*.xml`) are additional generated or
branding surfaces that must be classified separately from API examples.

The exact-domain search performed before this audit file was created found 167
matching files in the repository (including generated artifacts, docs, tests,
fixtures, and the pre-existing audit). The number is an inventory signal, not a
proposed blind replacement count; regenerate the inventory after any URL
changes.

### 2.8 SEO and public-web assets

SEO migration is explicitly out of scope under the current project premise.
The repository does contain canonical, sitemap, robots, JSON-LD, OpenGraph,
and LLM surfaces, but they do not need a backward-compatible migration. They
may be updated opportunistically for consistency on the new product host, or
retired/marked noindex if the public site is not being positioned yet. No SEO
redirect, sitemap submission, canonical preservation, or search-console work
is a release gate.

## 3. Target Architecture

The intended logical contract is:

```text
https://requiemsapi.com/*
    -> Rails public product, auth, dashboard, docs, account, webhooks

https://requiems.xyz/v1/*
    -> Auth Gateway Worker -> internal Go API

https://requiems.xyz/healthz
https://requiems.xyz/openapi.json
    -> explicit API operational/documentation endpoints, if retained

https://requiems.xyz/<non-API public path>
    -> explicit 404/410 (recommended); no generic redirect

https://api.requiems.xyz/v1/*
    -> retired; no redirect
```

The exact treatment of `/healthz` and `/openapi.json` must be part of the
contract. The current Worker exposes both on the old API host, and MCP/spec
automation depends on `/openapi.json`. For the clean break, publish them at the
new API host if operationally needed, update internal generators, and retire
the old endpoints. They must not fall through to a product redirect.

### 3.1 Recommended ingress design

Use Cloudflare path Routes on the proxied apex origin for the new API paths:

1. Route the exact `/v1` path and `/v1/*` to the selected API implementation.
   Preserve the original path, method, query, body, API-key header, and response
   behavior.
2. Route `/healthz` and `/openapi.json` only if the new API contract needs them;
   otherwise remove/update internal consumers and return 404 for them.
3. Return explicit 404/410 for every non-API apex path. Do not add a generic
   Redirect Rule. If HTTP is enabled, define HTTPS behavior separately and
   test API POST handling rather than inheriting an undocumented redirect.
4. Retire the old API Custom Domain and old apex Rails surface only after the
   new product/API paths and any active Lemon Squeezy callback are verified.

Cloudflare routes are appropriate here because the Worker is being placed in
front of an existing apex origin. The current Worker Custom Domain model is
appropriate for `api.requiems.xyz`, where the Worker owns the hostname.

### 3.2 Public product host

Configure the Rails/Kamal public service for `requiemsapi.com`, including:

- the canonical host and application host settings;
- auth, dashboard, admin, docs, public pages, and localized paths;
- mailer URL host for browser links;
- absolute links, auth/account links, and analytics where used;
- the literal dashboard URL contract (`/dashboard` versus localized path);
- any provider callbacks that should now land on the product host.

Keep `requiems.xyz` available at the edge for API routing only. Non-API apex
paths should return the selected 404/410 response; no redirect-only origin is
required for this clean-break scope.

## 4. Dependency Map

```text
api.requiems.xyz (legacy API host)
├── Cloudflare Worker custom domain
├── API-key authentication and rate/quota enforcement
├── D1 usage recording and response usage headers
├── OpenAPI endpoint (/openapi.json)
├── apps/mcp/scripts/fetch-spec.ts
├── apps/mcp/openapi.json and generated tools
├── 61 Rails API documentation YAMLs
├── docs/apis/**/*.md examples
├── README and developer docs
├── integration/load test defaults and performance reports
├── dashboard quick-start/API-key examples
├── external API clients and customer integrations (not measurable from repo)
└── monitoring/analytics outside repo (must be inventoried)

requiems.xyz (current product/apex host)
├── Rails/Kamal dashboard and public product
├── locale redirects and canonical URLs
├── Devise sessions, CSRF, confirmation/reset links
├── dashboard/account/admin URLs
├── POST /webhooks/lemonsqueezy
├── POST /api/proxy and server-side tool demos
├── sitemap, robots, llms.txt, llms-full.txt
├── JSON-LD, OpenGraph, canonical, hreflang
├── mcp/api documentation links and marketing links
└── target API path /v1/* must be carved out before redirecting

requiemsapi.com (new product host)
├── new DNS/TLS/Cloudflare/Kamal ingress
├── all public Rails pages
├── auth and account lifecycle
├── dashboard and admin
├── email links and payment callback destinations
└── new SEO canonical/sitemap/robots surface

internal.requiems.xyz
├── Auth Gateway BACKEND_URL
├── Go API deployment and health check
├── Caddy/Kamal secret guard
└── no public hostname change required by this migration

api-management.requiems.xyz
├── Rails API-management URL
├── API-key CRUD, plan sync, usage export, analytics
└── no public product/API domain change required by this migration

mcp.requiems.xyz
├── MCP ingress and health check
├── REQUIEMS_BASE_URL production configuration
├── OpenAPI fetch/generation workflow
└── public AI/LLM links from Rails
```

### 4.1 Classification of dependency classes

| Class | Verified dependencies |
| --- | --- |
| Application | Rails host/URL generation, locale routing, dashboard paths, webhook route, server-side API proxy, Go `/v1` routing, Worker proxy/auth |
| Infrastructure | Cloudflare Worker custom domain/routes, DNS records, TLS, Kamal proxy, optional Caddy, VPS origin, Cloudflare zone configuration |
| Configuration | `API_BASE_URL`, `INTERNAL_API_URL`, `MAILER_HOST`, `REQUIEMS_BASE_URL`, Wrangler routes, Kamal proxy host, Caddy host blocks |
| Documentation | README, docs/core, docs/apis, Rails API docs, LLM feeds, generated OpenAPI, MCP README/spec |
| Security | API-key header, backend secret, CSRF/session cookies, redirect handling, host/path precedence, webhook signature endpoint |
| Compatibility | Legacy API host, method/body/header preservation, generated clients, customer integrations, cached documentation |
| Operational | Worker health/openapi endpoints, Go health, Kamal/Caddy health checks, Sentry, uptime/traffic dashboards outside repo |
| SEO | sitemap default host, robots sitemap, canonical/hreflang/JSON-LD, OpenGraph URLs, old page redirects |
| Developer experience | curl/examples, integration defaults, load tests, MCP generation, API reference and quick-start snippets |

## 5. Required Changes

The following is an execution plan, not implementation code.

### 5.1 Pre-implementation discovery gates — HIGH

Before changing repository files or production DNS, record the following in a
change ticket/runbook:

1. Query authoritative DNS and Cloudflare DNS for `requiems.xyz`,
   `api.requiems.xyz`, `internal.requiems.xyz`, `api-management.requiems.xyz`,
   `mcp.requiems.xyz`, and `requiemsapi.com`. The purpose is retirement and
   cutover validation, not preserving old traffic.
2. Confirm which production ingress is active: Kamal proxy, Caddy, or another
   external proxy. Inspect the live proxy containers and Cloudflare origin
   settings; do not infer from documentation.
3. Export current Cloudflare Worker routes/custom domains, zone routes, redirect
   rules, cache rules, WAF rules, SSL mode, edge certificates, and DNS records.
4. Check for any non-user operational consumers: MCP, CI, scheduled jobs,
   monitoring, billing callbacks, email links, and internal scripts. No customer
   traffic or sunset telemetry is required under the confirmed scope.
5. Update or remove internal consumers before retiring old hosts: API docs,
   OpenAPI/MCP generation, CI defaults, examples, monitoring, and provider
   callbacks.
6. Confirm ownership and certificate availability for `requiemsapi.com`,
   including Cloudflare zone onboarding and registrar access.
7. Decide whether this clean break uses the current Auth Gateway Worker or the
   separate Go-authentication architecture. The recommended low-scope choice is
   to use the current Worker for the cutover and schedule Go-auth migration
   separately; do not combine two large migrations without an owner and
   rollback boundary.
8. Treat any Worker “staging” or canary as unsafe until isolation is proven.
   Current base Wrangler KV and D1 IDs are the production IDs, including
   `preview_database_id`, and only a production environment is declared
   (`apps/workers/auth-gateway/wrangler.toml:7-16,31-42`). The documented
   `pnpm run deploy` staging command has no environment selector
   (`apps/workers/auth-gateway/package.json:8-13`,
   `apps/workers/auth-gateway/readme.md:173-178`). Require separate staging
   Worker, KV, D1, secrets, origin, hostname, and observability before any
   canary; never test an apex route against production state unintentionally.

The decision record must map the owner of `/v1`, API-key validation, quotas,
usage writes, `/openapi.json`, and MCP fetch. Legacy-host support is explicitly
not an owner requirement because the old API host is being retired.

Do not proceed if the active apex ingress or any active billing/provider
callback is unknown.

### 5.2 Application and configuration changes — HIGH/MEDIUM

Rails/dashboard:

- Introduce one authoritative public-site URL setting for
  `https://requiemsapi.com` and use it for absolute public links, mailer URL
  options, and analytics where used. SEO canonical/sitemap work is optional.
- Change the production deployment host from `requiems.xyz` to
  `requiemsapi.com` only after the new host is serving a validated build.
- Decide and implement the dashboard URL contract. To satisfy the stated
  target literally, expose `/dashboard` as a stable entry point and define how
  it selects/redirects to the user's locale; otherwise explicitly document
  `/en/dashboard` (or the user's locale) as canonical.
- Preserve the Rails route namespace and all auth/account/admin paths during
  the host move.
- Update `MAILER_HOST` to the product host for Devise confirmation, reset,
  invitation/account, and deletion links. Keep `SMTP_DOMAIN`, sender address,
  and `mail.requiems.xyz` unchanged unless the email migration is separately
  approved and DNS-authenticated.
- Keep `/webhooks/lemonsqueezy` available on the new product host. Update the
  Lemon Squeezy webhook configuration and verify signature handling. Since no
  compatibility period is required, disable the old callback after the new
  endpoint is verified; do not add an old-host redirect.
- Keep `INTERNAL_API_URL` unchanged for the server-side Rails playground unless
  the origin topology changes. `API_BASE_URL` is loaded and validated by
  `AppConfig` (`apps/dashboard/app/lib/app_config.rb:164,177-179`) but the
  repository search found no runtime consumer of `AppConfig.api_base_url`.
  Change it only as part of an explicit configuration contract and do not
  assume changing it updates generated docs or UI examples. API documentation
  should use `https://requiems.xyz` with the `/v1` path contract, not
  `https://requiemsapi.com`.
- Audit hardcoded `https://requiems.xyz` values in helpers, mailers, views,
  translations, private deployment text, and generated pages. Classify each as
  public product URL, API URL, email address, private deployment wildcard, or
  external link before changing it.
- Update the shared Worker `WEBSITE_URL` constant and human-facing examples;
  sitemap/robots/OG-image updates are optional cleanup, not migration gates.

Auth/session:

- Treat the host move as a browser session boundary. Test that old cookies are
  not accidentally sent to the wrong site and that login, logout, remember-me,
  confirmation, reset-password, account deletion, and CSRF flows work on the
  new host.
- Verify production session cookie `Secure`, `HttpOnly`, `SameSite`, path,
  expiry, and HTTP-to-HTTPS behavior from live responses; the repository does
  not explicitly configure `force_ssl`, session-store, or cookie options
  (`apps/dashboard/config/application.rb:17-26`,
  `apps/dashboard/config/environments/production.rb:5-18`).
- Do not design a cross-domain session handoff or broad cookie-domain setting;
  there are no existing users to migrate. Validate only new-host login/CSRF
  behavior.

Go/API Worker:

- Preserve the Go route prefix `/v1` and backend secret boundary.
- Make the selected API implementation available only on the new apex API
  route. Verify the new route's methods, bodies, query strings, API-key
  headers, usage headers, error statuses, and binary responses.
- Decide whether `/healthz` and `/openapi.json` are supported on
  `requiems.xyz`; if yes, create explicit path routes and update all internal
  consumers; if no, remove/update those consumers and return 404.
- Do not make the Go backend serve public redirects or accept public traffic
  without `X-Backend-Secret`.
- Decide whether the Worker may forward Rails `Cookie` headers to Go. Current
  `filterHeaders` forwards them, so API requests on the apex can cross a
  session-cookie boundary even though the API authenticates with a key. Add a
  test proving the selected policy.

Additional verified application controls from review:

- Generate the complete `bin/rails routes` method/path inventory and explicitly
  retire the old-apex Rails routes after moving the product host. This includes
  `/api/proxy`, `/locale`, `/tools/demos/*`, `/webhooks/lemonsqueezy`, localized
  forms, Devise, dashboard, and admin mutations
  (`apps/dashboard/config/routes.rb:16-54,67-149,168-186`). No 301/302/307/308
  compatibility behavior is required; test that retired old paths fail closed
  and that their new-host equivalents work.
- Make the product host policy explicit. The repository has no production
  `config.hosts` allowlist and SEO/payment helpers use `request.base_url` or
  request-derived URL helpers (`apps/dashboard/app/helpers/application_helper.rb:172-187,403-484`,
  `apps/dashboard/app/controllers/private_deployments_controller.rb:56-61`).
  Define accepted hosts, trusted proxy handling, `Host`/
  `X-Forwarded-Host` behavior, and fixed production absolute URLs before
  enabling the new host.
- Preserve the internal playground contract: Rails uses `INTERNAL_API_URL`
  and `X-Backend-Secret`, not the public API base. Do not route demos through
  the public gateway unless metering and authentication are deliberately
  redesigned.
- Treat payment return URLs and private-deployment tenant URLs separately.
  Checkout return URLs are request-host-derived, while tenant URLs use the
  `*.requiems.xyz` namespace (`apps/dashboard/app/models/private_deployment_request.rb:83-85`).
  Do not blindly migrate either surface.
- Fix the authentication contract as part of the documentation migration:
  the gateway accepts only `requiems-api-key`, while the root README, FAQ, and
  generated security material advertise Bearer authentication
  (`apps/workers/auth-gateway/src/middleware/api-key-auth.ts:39-46`,
  `readme.md:98-105`, `apps/dashboard/app/helpers/application_helper.rb:110-116`).
  Either add and test Bearer support as a separate API change or correct all
  first-party surfaces to the existing header; test preflight for the chosen
  contract.
- Do not treat the requested illustrative paths as existing API aliases. The
  current Go routes are `/v1/validation/email`, `/v1/validation/phone`,
  `/v1/networking/ip`, and `/v1/systems/signup/protect`; adding shorter aliases
  is outside a domain-only migration and needs a separate product decision.
- Treat `POST /locale` as an existing redirect sink in its own security test:
  it derives the destination from `Referer` and preserves query/fragment
  (`apps/dashboard/app/controllers/locale_controller.rb:17-36`). Fuzz external,
  encoded, malformed, and cross-host referers on the new product host; no old
  apex redirect rule is required.

### 5.3 Infrastructure and routing changes — HIGH

Implement only after the discovery gates identify the real ingress:

1. Add `requiemsapi.com` to the active product ingress and certificate
   configuration. It must reach the Rails service directly, not the API Worker.
2. Choose one authoritative redirect/ingress mechanism after discovery. The
   preferred mechanism is a single apex dispatcher Worker; if separate
   Cloudflare Routes and Redirect Rules are chosen, record the exact
   terminating Rules phase, exclusion expression, path/query handling, HTTP
   behavior, and Cloudflare Trace result. The repository has Kamal and Caddy
   descriptions but no configured Cloudflare redirect implementation. Do not
   implement the rule in an inactive layer.
3. Validate the product host through the complete Rails stack before changing
   the old host. If no isolated preview hostname or service exists, call this
   an in-place host addition and use a controlled hosts/edge validation path;
   do not assume the existing Kamal deployment creates a parallel preview.
4. Keep the apex DNS record proxied and pointing to the active origin required
   for non-API traffic. A Cloudflare Worker Route for `requiems.xyz/v1/*`
   requires a proxied DNS record at that hostname.
5. Attach the selected apex dispatcher/API Worker to both the exact `/v1` path
   and `/v1/*` behavior as required by the Cloudflare route matcher. Specify
   the production zone identifier/ownership, preserve the legacy Custom Domain,
   and prove the Worker receives the original path. Include explicit
   operational/spec routes or dispatcher branches according to the selected
   contract.
6. Provide a concrete old-apex Rails grace path for all retained non-API POSTs.
   Moving Kamal to `requiemsapi.com` alone does not make
   `requiems.xyz/webhooks/lemonsqueezy` or the other old-host POSTs reach Rails.
   Define the old-host TLS, Host handling, proxy/origin target, CSRF behavior,
   raw-body preservation, provider switch timing, and retirement check.
7. Configure the old apex behavior after the API route and grace path are active:
   - `/v1/*` must never redirect to the product host;
   - `GET`/`HEAD` public paths should 301 to the equivalent product path,
     preserving path and query string;
   - POST webhook traffic must be handled by an origin-compatible endpoint
     until the external provider is switched and the old endpoint is retired;
   - retain or explicitly retire the other old-apex POST routes (`/api/proxy`,
     `/locale`, and `/tools/demos/*`) rather than sending them through a
     generic redirect; stale/cached product pages may still submit to those
     paths;
   - define behavior for unknown methods, `/up`, `/llms.txt`, `/llms-full.txt`,
     `/sitemap.xml`, and legacy localized pages explicitly.
8. Ensure redirect/cache precedence is tested at Cloudflare edge, at the active
   origin proxy, and at Rails. A broad apex redirect configured before the API
   route can take precedence incorrectly.
9. Disable caching for authenticated API responses and webhook requests; do
   not cache 301s during the initial validation window beyond an intentionally
   chosen short TTL. Existing cached redirects can outlive a rollback.
   The change ticket must name the Cloudflare Cache Rules/Worker-cache owner,
   exact path exclusions, origin `Cache-Control` headers, purge API/operator,
   TTLs, and acceptance headers (`Age`, `CF-Cache-Status`, `Cache-Control`).
   Apply a temporary no-cache policy to `robots.txt`, `sitemap*.xml`, and
   `sitemap.xsl` while their hosts/canonicals change, then restore only an
   intentional long-lived policy after freshness is verified.
10. Add an explicit Worker deployment step. The repository's CD workflow deploys
    API, dashboard, and MCP with Kamal but has no Auth Gateway Wrangler job
    (`.github/workflows/cd.yml:75-149`). Record the deploy owner, Cloudflare
    credentials/approval, `pnpm run generate:openapi`, `pnpm run deploy:prod`,
    route verification, and Worker rollback/version procedure.
11. Lock down the Go origin independently of the shared header: verify firewall
    and Cloudflare-only reachability, reject direct-origin and direct-Worker
    spoof traffic, and rotate the static `BACKEND_SECRET` only with a tested
    dual-secret/maintenance procedure. Kamal exposes `internal.requiems.xyz`
    and Caddy/Go currently rely on the same header (`infra/kamal/deploy.api.yml:14-18`,
    `infra/caddy/Caddyfile:31-46`, `apps/api/platform/middleware/auth.go:10-35`).

### 5.4 DNS and TLS — HIGH

Repository-managed:

- Update the active Kamal/proxy hostname for the product service.
- If Caddy is active, add the product host and the final apex host/path rules
  there; if Kamal proxy is active, update Kamal only. Do not modify both
  unconditionally.
- Update Worker Wrangler route configuration for the new route and preserve the
  legacy custom domain. Confirm environment-specific route declarations are
  complete.
- The Cloudflare change ticket must name the zone ID/name, Worker script and
  environment, exact route patterns (`requiems.xyz/v1` and
  `requiems.xyz/v1/*` if both are needed), proxied DNS record, generic redirect
  rule owner, precedence/overrides, cache/WAF rule owner, certificate owner,
  deploy credential/approval owner, and rollback version. These values are
  external and currently absent from the repository; the plan is not complete
  until they are recorded.
- Keep `internal.requiems.xyz`, `api-management.requiems.xyz`, and
  `mcp.requiems.xyz` unchanged unless separate decisions require moving them.

External DNS/Cloudflare:

- Onboard and verify `requiemsapi.com` in Cloudflare, set the required A/AAAA/
  CNAME record to the active product origin, and proxy web traffic.
- Preserve the existing `api.requiems.xyz` Worker custom-domain DNS/certificate
  state during the compatibility period.
- Ensure `requiems.xyz` has a proxied record suitable for Worker Routes and the
  non-API product redirect/origin.
- Keep mail MX/TXT/CNAME records and `mail.requiems.xyz` sender authentication
  unchanged unless email ownership is intentionally migrated.
- Verify Cloudflare edge certificates for both new product and API hosts, the
  origin certificate/SSL mode, Caddy/Kamal certificate behavior, and renewal.
- Confirm IPv4/IPv6, DNSSEC, CAA, HSTS, and certificate transparency impacts
  with the infrastructure owner.
- Export the pre-change DNS and certificate state and define deletion order:
  never remove the legacy Worker Custom Domain, its DNS, or its certificate
  while compatibility is supported. Verify registrar ownership, CAA, DNSSEC,
  Cloudflare edge certificates, Worker Custom Domain certificates, product
  host certificates, origin certificates, SSL mode, renewal, and any origin
  IP exposure. The repository cannot prove these external facts.
- Verify actual response headers on every promised host. Generated security
  documentation claims HSTS, but no repository ingress configuration proves
  that `Strict-Transport-Security` is emitted (`scripts/generate-docs.mjs:437-444`).
  Decide `includeSubDomains`/preload deliberately and test HTTP-to-HTTPS
  behavior before publishing an HSTS policy.

### 5.5 Backwards compatibility — HIGH

Keep `api.requiems.xyz` as a permanent dual-host alias to the same Auth Gateway
behavior unless product leadership explicitly approves a supported-by date and
customer migration program.

Do not redirect API traffic. A 301/302/307/308 adds a second request and can
break clients that do not follow redirects, clients that refuse to forward
authorization-like headers, POST/PUT bodies, signed requests, streaming or
binary responses, or clients whose redirect behavior changes the method. It
also makes compatibility dependent on client redirect policy and can produce
cache poisoning or stale redirect behavior.

The legacy host should remain operationally observable and tested with the
same API-key, rate-limit, quota, error, usage-header, binary-response, and
method matrix as the new host. If it is ever retired, use a documented
deprecation period, customer inventory, explicit communication, a separately
verified redirect policy, and a measured rollback window.

The Lemon Squeezy controller verifies HMAC over the raw request body, but the
repository has no event-ID ledger or replay window. Existing tests prove one
subscription case is business-idempotent, not that every side effect is
deduplicated; private-deployment mail enqueueing can repeat
(`apps/dashboard/app/controllers/webhooks/lemonsqueezy_controller.rb:38-64,115-133`,
`apps/dashboard/test/controllers/webhooks/lemonsqueezy_controller_test.rb:70-82`).
Before old and new webhook endpoints are active concurrently, choose one active
provider callback, add or verify event-level idempotency/replay controls, and
test duplicate delivery, cross-host delivery, invalid signatures, and raw-body
preservation.

The D1 usage schema does not include the request hostname, so host adoption
cannot be measured from billing usage rows alone. Add hostname labels in edge
logs/Cloudflare analytics or another non-billing metric; do not change quota
aggregation merely to measure the migration.

### 5.6 Documentation and developer experience — MEDIUM/HIGH

Update the API base URL to `https://requiems.xyz` (with endpoint paths under
`/v1`) in:

- all `apps/dashboard/config/api_docs/*.yml` `base_url` values and handwritten
  examples;
- generated OpenAPI source and regenerated Worker spec;
- `apps/mcp/openapi.json`, MCP fetch URL, production `REQUIEMS_BASE_URL`, and
  any generated artifacts;
- `docs/apis/**/*.md`, root README, docs/core, apps/workers READMEs, and
  examples/quick-start content;
- integration/load test defaults and production test fixtures;
- LLM feeds and API-reference output.

Update the public product URL to `https://requiemsapi.com` in human-facing
links, auth links, account links, marketing links, sitemap/LLM links, and
structured data. Keep email addresses, private-deployment wildcard domains,
MCP host, management host, and origin host distinct.

Generated files must be regenerated by their documented generators and then
reviewed for semantic URL changes. Do not do a repository-wide blind string
replacement.

Before generation, select exactly one canonical OpenAPI fetch and advertised
server contract for the chosen architecture branch. For the Worker-retention
branch, the safe transition is to keep the non-redirecting
`https://api.requiems.xyz/openapi.json` fetch until
`https://requiems.xyz/openapi.json` is explicitly served and tested; then
update `servers[0].url` and MCP runtime configuration to the new API host.
For the Go-cutover branch, the Go-served spec and its compatibility host must
be selected before retiring the Worker. In either branch, a fetch that receives
an HTML redirect is a failed build, not a successful regeneration.

Before publishing any regenerated OpenAPI/MCP artifact, produce a normalized
before/after diff of every `(method, path, operationId)` tuple. Fail the
release on endpoint loss, path drift, or operation reassignment unless a
separately approved API change explains it. This is required because the
repository already contains verified drift: the Worker spec advertises
`/v1/entertainment/advice`, while the MCP snapshot/generated tool use
`/v1/text/advice` and the Go router mounts advice under `/v1/entertainment`
(`apps/workers/auth-gateway/src/generated/openapi.ts:332`,
`apps/mcp/generated/tools/text_advice.ts:2,24`,
`apps/api/app/routes_v1.go:42-47`).

Keep the URL layers separate: OpenAPI `servers.url`, the OpenAPI fetch URL,
MCP `REQUIEMS_BASE_URL`, product links, and API example base paths are separate
contracts. The MCP `--base-url` option is parsed but not used by generation
(`apps/mcp/scripts/generate.ts:74-95,492-517`); require an actual runtime
environment/argument override and do not fetch a new spec through an endpoint
that merely redirects. Audit the external SDK repository linked by the README
(`bobadilla-tech/requiems-api-clients`), the external `requiems-api-skills`
repository, SDK release workflows, and any scheduled OpenAPI consumers.

Assign an owner to the external SDK decision. Before announcing the new API
host, either publish versioned SDK releases whose base URL/header contract
passes both-host tests, or explicitly defer SDK publication and label the
existing client repository as unsupported for this migration. Do not imply SDK
readiness from OpenAPI regeneration alone.

Also correct adjacent verified documentation drift during the same review:
`requiems-api.xyz`, `api.requiems.dev`, the root README's Bearer example, the
gateway's product CTA, and the quota-exceeded upgrade URL. Local Compose MCP
defaults must remain environment-specific rather than silently changing local
development to production.

### 5.7 SEO — MEDIUM

- Set canonical and `hreflang` host generation to `requiemsapi.com`.
- Set JSON-LD Organization, WebSite, BreadcrumbList, WebAPI, Article, and
  Service URLs to the product host where they describe human-facing pages.
- Regenerate all sitemap files with `requiemsapi.com` and update robots.txt.
- Keep API documentation examples pointing to `requiems.xyz/v1`.
- Create path-preserving 301 redirects from old public page URLs to their new
  product-host equivalents, including locale paths and trailing-slash policy.
- Avoid redirect chains caused by Rails locale canonicalization: either map
  old unlocalized paths directly to final localized product URLs or make the
  new product host serve stable unlocalized pages. Verify one-hop behavior for
  `/`, `/pricing`, `/docs`, and localized paths.
- Keep `/v1/*`, API health/spec endpoints, and webhook paths out of generic SEO
  redirect logic.
- Rebase every structured-data and social URL on a configured product host,
  not the incoming `Host`/`request.base_url`; normalize trailing slashes and
  remove query strings from canonical, OpenGraph, and JSON-LD URLs. Update
  Organization, logo, image, WebSite, breadcrumb, BlogPosting, WebAPI,
  Service, Article, and publisher URLs.
- Regenerate sitemap index/sections/XSL and robots. Verify the tracked files are
  fresh, both slash forms of private paths are excluded/noindex, auth pages are
  non-indexable, and blog hreflang entries represent real translations rather
  than every available locale.
- Publish and test `requiemsapi.com/llms.txt`, `/llms-full.txt`, and
  `/apis/{id}/index.md`; machine-readable feeds must use product URLs for human
  pages and `https://requiems.xyz/v1` for API calls.
- Submit the new sitemap and domain property in Search Console; monitor
  indexing, redirect chains, canonical selection, and coverage.

### 5.8 Monitoring and operations — HIGH/MEDIUM

Before cutover, create separate dashboards/alerts for:

- `requiemsapi.com` product health and Rails 4xx/5xx/latency;
- `requiems.xyz/v1/*` and `api.requiems.xyz/v1/*` API success/error/latency,
  auth failures, quota/rate-limit responses, and usage headers;
- Worker invocation/error/CPU metrics and Sentry events for both route forms;
- Go `internal.requiems.xyz/healthz` and backend secret failures;
- API Management, MCP, Rails, Sidekiq, database, Redis, and webhook delivery;
- DNS, TLS/certificate renewal, Cloudflare route activation, origin reachability,
  and redirect-loop detection.

Because D1 usage rows and the shared logger do not carry the request hostname,
define the adoption metric before cutover: Cloudflare Logs/Analytics (or an
equivalent edge stream) must retain `ClientRequestHost`, route/script/version,
status, and request count for a named retention period; assign a query owner;
and reconcile total successful requests against D1 usage totals without
altering billing rows. A screenshot or one-time dashboard view is not enough
for legacy-sunset evidence.

Add synthetic probes for GET and POST API endpoints on both API hosts, the
product root, localized pages, `/dashboard`, login, password reset request,
the webhook endpoint in a safe test mode, `/openapi.json`, `/llms.txt`, and
`/sitemap.xml`.

Add alerts/probes for cache leakage (`Cache-Control`, `Age`,
`CF-Cache-Status`), actual HSTS headers, redirect loops, internal exception
disclosure, and forged client-IP/forwarded-host headers. The Worker currently
has no explicit API response cache policy, its global error handler returns
`err.message`, and its forwarding logic only overwrites `X-Forwarded-For` when
Cloudflare supplies `CF-Connecting-IP` (`apps/workers/shared/src/middleware/error-handler.ts:10-23`,
`apps/workers/auth-gateway/src/http.ts:13-31`). These should be treated as
security acceptance checks, not assumed properties of the edge.

### 5.9 Tests required — HIGH/MEDIUM

Add or execute a host/path contract matrix before production:

| Request | Expected result |
| --- | --- |
| `GET https://requiemsapi.com/` | Rails product 200/canonical product host |
| `GET https://requiemsapi.com/en/` | Rails localized product page |
| `GET https://requiemsapi.com/dashboard` | Defined dashboard contract; no ambiguity |
| `GET https://requiems.xyz/` | Redirect to product root |
| `GET https://requiems.xyz/en/pricing` | Path-preserving redirect to product |
| `GET/POST https://requiems.xyz/v1/...` | Auth Gateway API behavior, no redirect |
| `GET/POST https://api.requiems.xyz/v1/...` | Same legacy behavior, no redirect |
| `OPTIONS https://requiems.xyz/v1/...` | Correct CORS preflight |
| `GET /healthz` and `/openapi.json` on each promised host | Contract-specific success |
| `POST /webhooks/lemonsqueezy` on new and grace-period old host | Signature/body accepted; no generic redirect |
| Rails login/logout/confirmation/reset/account deletion | Correct new-host cookies and URLs |
| MCP HTTP tool call | Calls new API base and propagates caller API key |

Extend the matrix with:

- exact `/v1`, `/v1/`, and `/v1/<endpoint>` routes; HTTP-to-HTTPS, query,
  malformed, and unknown-method cases;
- old-apex `POST /api/proxy`, `/locale`, every tool demo, contact, suggestion,
  sales, private-deployment, Devise, dashboard, admin, and webhook route;
- 301/302/307/308 behavior with authenticated POST bodies, CSRF tokens, API
  headers, and clients configured both to follow and not follow redirects;
- fixed-host redirect fuzz cases for encoded `//`, backslashes, CR/LF, and
  unusual paths to prevent open redirects;
- direct origin requests without and with forged `X-Backend-Secret`, direct
  Worker requests with forged `CF-Connecting-IP`/forwarded headers, and host
  allowlist/trusted-proxy behavior;
- cookie forwarding from same-origin product requests, Rails `CF-Connecting-IP`
  handling for `/api/proxy` and tool demos, Go `X-Real-IP` fallback, generic
  production errors from both Worker and Rails proxy/demo surfaces, and the
  existing `/locale` Referer-derived redirect sink;
- cache and HSTS headers, generic production error bodies, CORS credentials and
  preflight, `Access-Control-Expose-Headers` expectations, and legacy/new
  parity for JSON, PNG/binary responses, status, body hash, usage headers,
  timeout, and error behavior;
- duplicated Lemon Squeezy event delivery and one active provider callback
  during the transition.

API compatibility tests must include GET, POST, PUT/PATCH/DELETE if supported,
JSON bodies, query strings, authorization/API-key header, binary PNG responses,
streaming/large bodies if present, 4xx/5xx responses, rate-limit/quota headers,
and timeout behavior. Test redirect handling explicitly with clients that do
and do not follow redirects.

## 6. Backwards Compatibility Strategy

### Recommendation: permanent dual-host API support

Keep `api.requiems.xyz` on the existing Auth Gateway Worker and add
`requiems.xyz/v1/*` as a second invocation path to the same Worker logic. Do
not redirect the legacy API host and do not move API keys, D1, KV, usage
semantics, or backend secrets as part of this domain migration.

Rationale verified from the repository:

- The old hostname is embedded in 167 files and likely in external clients.
- The API accepts POST bodies and binary responses, and generated examples use
  multiple HTTP clients.
- The Worker is already the stable authentication/usage boundary.
- No customer migration inventory or deprecation policy is present in the repo.
- The repository provides no evidence that all external clients follow redirects
  correctly.

Compatibility requirements:

- Same Worker version and bindings for both host paths.
- Same request path beginning `/v1`; no accidental path stripping.
- Same API-key header and backend-secret behavior.
- Same CORS, usage, quota, rate-limit, status, content-type, and cache policy.
- Same D1 usage ledger and observability dimensions, with host added as a
  metric label if possible.
- No redirect response on API requests.
- Old host remains in DNS and certificate inventory for the supported lifetime.

Only consider deprecation after traffic is measured, customers are contacted,
the docs have been stable on the new API URL, and a separately approved sunset
plan exists. If leadership requires eventual retirement, the plan must specify
the date, client notification, support process, final status code policy, and
rollback limitations.

## 7. Deployment Plan

1. Freeze the migration contract: public host, exact dashboard URL, API
   operational/spec endpoints, old-page redirect policy, webhook grace period,
   email-host policy, and legacy API support lifetime.
2. Complete external discovery gates and save DNS/Cloudflare/Kamal/Caddy
   exports, current certificates, health checks, traffic baselines, and a
   database/Worker configuration backup.
3. Provision/verify `requiemsapi.com` DNS, Cloudflare zone, certificates, and
   origin access without changing old traffic. Capture the exact edge/origin
   mechanism; do not assume a preview exists.
4. Add the new Rails host to the verified active ingress. If an isolated
   preview is unavailable, validate through a controlled edge/hosts path and
   explicitly label the operation as an in-place host addition. Validate auth,
   session behavior, all public routes, dashboard, admin, mail links,
   sitemap/robots/JSON-LD, and payment/webhook test flows.
5. Deploy the Auth Gateway/OpenAPI Worker explicitly with Wrangler, because
   the current CD workflow does not deploy it. Confirm the production Worker,
   bindings, exact `/v1` and `/v1/*` routes, and the legacy Custom Domain.
6. Add the new Auth Gateway route for `requiems.xyz/v1/*` in a disabled,
   staged, or low-risk activation state supported by the Cloudflare account.
   Confirm that the apex DNS record is proxied and the route resolves to the
   correct Worker without changing `/` behavior.
7. Establish and test the old-apex Rails grace path for retained POSTs before
   moving the product host or enabling generic redirects. Keep exactly one
   active Lemon Squeezy callback and verify event deduplication/replay policy.
8. Deploy/update API contract and generated artifacts, but keep legacy docs and
   old API host available until the new path passes live canaries.
9. Run the dual-host API contract tests with a real test key and compare
   responses, headers, latency, usage, and D1 records.
10. Switch the Rails product host to `requiemsapi.com` and update external
   provider callbacks. Keep the old Rails host capable of handling webhook and
   compatibility traffic.
11. Enable old-apex path-preserving redirects only after `/v1/*`, all required
   POST exceptions, and webhook handling are verified. Use temporary 302/307
   with no-cache during validation; promote to 301 only after acceptance, then
   monitor loops, cached redirects, and POST failures.
12. Publish docs and new canonical URLs, regenerate OpenAPI/MCP artifacts,
   update integration defaults, and announce the new product/API URLs.
13. Monitor for at least one complete high-traffic period and one billing/
   webhook cycle. Compare old/new API traffic, auth failures, status codes,
   latency, usage accounting, Rails sign-ins, email clicks, and SEO crawl data.
14. Close the migration only when the external dependency checklist is signed
   off. Keep the legacy API host and rollback records intact.

## 8. Rollback Plan

Rollback must be staged and reversible:

1. If new API path failures occur, disable only the `requiems.xyz/v1/*`
   Worker route and leave `api.requiems.xyz` serving production traffic. Restore
   apex non-API origin behavior if changed.
2. Use a dedicated rollback runbook/workflow with an explicit prior Worker
   version, Wrangler route/binding export, image tag, registry credentials,
   Kamal config, proxy-host state, and acceptance probes. The current composite
   action always runs `kamal setup` and the CD workflow has no rollback input
   (`.github/actions/kamal-deploy/action.yml:44-46`, `.github/workflows/cd.yml:17-28`),
   so a normal deploy rerun is not a rollback procedure.
   For Cloudflare, record the deployed Worker version, deployment ID, route
   bindings, and the exact disable/restore operation for `/v1`, `/healthz`,
   `/openapi.json`, and the generic redirect. Roll back the Worker version and
   route state as one change; do not assume version rollback alone removes a
   newly attached route.
3. If the product host fails, stop the old-host redirects and restore
   `requiems.xyz` to the previous Rails ingress while keeping the new host
   available for diagnosis. Do not change the Rails `SECRET_KEY_BASE` or
   invalidate database sessions unnecessarily.
4. If auth or API behavior differs, roll back the Worker deployment/configuration
   while preserving KV/D1 IDs and secrets. Do not recreate the namespace or
   database as part of rollback.
5. If webhook delivery fails, restore the old webhook endpoint in the provider
   and keep the old POST route live. Replay only provider-supported events after
   signature and idempotency checks.
6. If SEO redirects are wrong, disable the generic redirect rule, preserve
   product availability, and correct canonical/sitemap generation before
   re-enabling. Expect previously cached 301s to persist at some clients.
7. Verify all rollback health checks, API smoke tests, customer-key requests,
   dashboard login, mail links, API Management, MCP, and webhook delivery.

Rollback limitations:

- DNS and certificate propagation are not instantaneous.
- Cached permanent redirects may survive the edge rollback.
- Sessions created only on the new registrable domain will not make old-host
  users logged in.
- External provider settings and customer clients may have already changed.
- A rollback must not delete `requiemsapi.com` DNS/certificates or the legacy
  API custom domain until the incident is closed.

## 9. Findings

Only verified repository findings are listed here. External state is listed as
an assumption/dependency below, not presented as a repository fact.

| Severity | Finding | Evidence | Impact | Recommended Action |
| --- | --- | --- | --- | --- |
| HIGH | The current public API is a host-wide Worker custom domain at `api.requiems.xyz`; the apex is Rails, not the API. | `apps/workers/auth-gateway/wrangler.toml:38-42`; `infra/kamal/deploy.dashboard.yml:15-18`; `infra/caddy/Caddyfile:5-23` | A hostname string change cannot create the required apex path split. | Add and test a Cloudflare path route for `/v1/*` before enabling apex redirects. |
| HIGH | The same exact hostname `requiems.xyz` currently serves a POST Lemon Squeezy webhook. | `apps/dashboard/config/routes.rb:52-54`; `docs/core/lemonsqueezy-webhook-setup.md:56` | A broad redirect can break signed POST delivery or lose request body/method semantics. | Move provider callback to new host and retain an old-host POST grace path; never generic-redirect it first. |
| HIGH | The API legacy hostname is referenced by generated specs, MCP, 61 API YAMLs, docs, tests, and UI examples. | `apps/workers/auth-gateway/scripts/openapi/constants.ts:27`; `apps/mcp/scripts/fetch-spec.ts:1`; `apps/dashboard/config/api_docs`; exact-domain inventory | Blind replacement or legacy shutdown can break clients and generators. | Keep dual-host API support and update consumers by semantic category. |
| HIGH | Rails has no repository-defined custom session cookie domain, and the target uses a different registrable domain. | No `session_store` initializer found; Rails layout includes CSRF; `infra/kamal/deploy.dashboard.yml:65`; routes in `apps/dashboard/config/routes.rb:64-111` | Existing browser sessions will not automatically transfer; insecure cookie workarounds could create auth issues. | Plan reauthentication or separately review a one-time signed handoff. |
| HIGH | Production ingress has two competing repository descriptions; CD deploys Kamal while Caddy/Compose remain documented as production. | `.github/workflows/cd.yml:75-149`; `infra/kamal/deploy.*.yml`; `infra/docker/docker-compose.yml`; `infra/caddy/Caddyfile` | Editing the wrong ingress can cause outage or leave the active path unchanged. | Verify live VPS and Cloudflare ingress before selecting files to change. |
| HIGH | Go `/v1` routes are protected only by backend secret middleware and have no public host/path logic. | `apps/api/app/app.go:35-51`; `apps/api/platform/middleware/auth.go` | Incorrect edge routing can bypass the gateway or route API requests to Rails. | Keep Worker/secret boundary and test direct-origin rejection plus path precedence. |
| HIGH | The old apex has many state-changing POST routes beyond the payment webhook. | `apps/dashboard/config/routes.rb:16-54,67-124,168-186` | A generic redirect can lose bodies, CSRF state, cookies, or method semantics for demos, contact/sales, Devise, dashboard, admin, and webhook traffic. | Inventory each route and retain, proxy, migrate, or retire it explicitly; test 301/302/307/308 behavior. |
| HIGH | Rails has no repository-defined production host allowlist, while absolute URLs are partly derived from the request host. | No production `config.hosts`; `apps/dashboard/app/helpers/application_helper.rb:172-187,403-484`; `apps/dashboard/app/controllers/private_deployments_controller.rb:56-61` | Host-header/proxy mistakes can produce wrong links, redirect abuse, or cross-host security behavior during cutover. | Define accepted hosts and trusted proxy policy; use fixed configured production URLs and test Host/X-Forwarded-Host rejection. |
| HIGH | The Go origin is exposed through a public hostname and relies on a static shared `X-Backend-Secret` guard; repository firewall policy does not prove Cloudflare-only reachability. | `infra/kamal/deploy.api.yml:14-18`; `infra/caddy/Caddyfile:31-46`; `apps/api/platform/middleware/auth.go:10-35` | Direct-origin exposure or secret leakage can bypass the intended gateway boundary. | Verify origin firewall/Cloudflare-only access, direct-origin rejection, secret rotation, and spoofed-header tests before changing routes. |
| HIGH | Webhook HMAC verification has no event-ID ledger or replay window; only one subscription test is business-idempotent. | `apps/dashboard/app/controllers/webhooks/lemonsqueezy_controller.rb:38-64,115-133`; `apps/dashboard/test/controllers/webhooks/lemonsqueezy_controller_test.rb:70-82` | Concurrent old/new delivery or replay can repeat email/deployment side effects. | Keep one active provider callback and add/verify event-level idempotency/replay controls before dual ingress. |
| HIGH | The Auth Gateway forwards incoming `Cookie` headers to the Go origin even though API authentication is header-based. | `apps/workers/auth-gateway/src/http.ts:16-31` | Apex API requests made in a browser session can carry Rails cookies across the API boundary and expand the impact of a routing or origin mistake. | Decide and test cookie filtering, and verify session-cookie attributes and same-origin API behavior. |
| HIGH | The existing CD workflow does not deploy the Auth Gateway Worker, and the repository does not identify an isolated product-host preview. | `.github/workflows/cd.yml:75-149`; `apps/workers/auth-gateway/package.json:8-13`; `infra/kamal/deploy.dashboard.yml:15-18` | A route/config change may never reach production, or a supposed preview may alter the only live ingress. | Add an explicit Wrangler deploy/verification owner and define the actual parallel validation mechanism. |
| HIGH | This migration’s permanent Worker strategy conflicts with the repository’s separate architecture audit, which recommends retiring the Worker/KV/D1/API-Management layer. | `docs/audits/2026-08-21-architecture-audit.md:70-84,1357-1360`; this audit’s dual-host strategy | Implementing both plans without sequencing can create a short-lived route, contradictory auth/usage semantics, and an unsupported legacy host. | Obtain an architecture decision: freeze Worker retirement for the support lifetime, or schedule the domain migration after Go-authenticated ingress and rewrite this plan. |
| HIGH | The documented Worker staging path uses production KV/D1 identifiers and has no staging environment selector. | `apps/workers/auth-gateway/wrangler.toml:7-16,31-42`; `apps/workers/auth-gateway/package.json:8-13`; `apps/workers/auth-gateway/readme.md:173-178` | A canary can mutate production keys, quotas, rate limits, or usage data and does not provide safe route validation. | Provision isolated staging resources/hostname/secrets/observability before any canary; prohibit production apex attachment from the staging command. |
| HIGH | No repository source names the authoritative Cloudflare zone, apex Worker Route, redirect-rule owner, or precedence configuration. | `apps/workers/auth-gateway/wrangler.toml:40-42`; no apex route/zone declaration; `infra/kamal/deploy.dashboard.yml:15-18`; `infra/caddy/Caddyfile:5-23` | An engineer could configure an inactive layer or let a generic redirect intercept `/v1` traffic. | Record exact zone, Worker, route patterns, DNS record, rule order, owners, and Cloudflare Trace acceptance before implementation. |
| HIGH | The old-apex POST grace path is required but not executable from current ingress configuration. | Old host only in `infra/kamal/deploy.dashboard.yml:15-18` and `infra/caddy/Caddyfile:5-23`; no new-host alias/proxy/retirement switch exists | Moving Rails to the product host can strand webhook, forms, Devise, dashboard, and admin POSTs before they are migrated. | Define the concrete old-host TLS/proxy/Host target, raw-body behavior, provider switch, and retirement operation. |
| HIGH | Existing deployment automation has no versioned rollback input or Cloudflare route rollback procedure. | `.github/actions/kamal-deploy/action.yml:44-46`; `.github/workflows/cd.yml:17-28`; no Worker route rollback automation | Re-running deployment cannot reliably restore the prior product/API topology; cached redirects and route state may persist. | Create a rollback runbook/workflow with image/Worker version, route bindings, credentials, disable/restore operations, and probes. |
| MEDIUM | Rails dashboard URLs are locale-scoped while target names `/dashboard` explicitly. | `apps/dashboard/config/routes.rb:64-111`; `ApplicationController#set_locale` | A target URL can redirect unexpectedly or fail if not deliberately defined. | Decide and test an unlocalized dashboard entry point/canonical URL. |
| MEDIUM | API CORS is wildcard and header-based; preflight behavior is implemented by Worker middleware. | `apps/workers/shared/src/http.ts:5-15`; `apps/workers/shared/src/middleware/cors.ts:9-15` | New route must preserve preflight and API-key header behavior; browser failures may be host-specific. | Test OPTIONS and actual requests from product and third-party origins. |
| MEDIUM | OpenAPI and `/healthz` are current Worker routes but target only names `/v1/*`. | `apps/workers/auth-gateway/src/index.ts:20-30`; generated spec URL | Removing/redirecting these endpoints breaks operations and MCP generation. | Define explicit new-host contract or preserve old URLs and update consumers. |
| MEDIUM | `API_BASE_URL` is configured and validated but has no runtime consumer in the repository. | `apps/dashboard/app/lib/app_config.rb:164,177-179`; repository-wide `AppConfig.api_base_url` search | Changing the variable alone will not update UI, docs, or generated API examples and can create false confidence in the cutover. | Treat it as a configuration contract to either wire explicitly or update only after identifying its intended consumer. |
| MEDIUM | The target API examples do not match current Go route paths. | `apps/api/app/routes_v1.go:42-95`; validation, networking, and signup-protect transport routes | Domain-only migration docs could accidentally promise nonexistent endpoint aliases. | Freeze the migration as domain-only; require a separate decision for path aliases and test only actual routes. |
| MEDIUM | First-party authentication documentation contradicts the gateway and its CORS preflight. | `apps/workers/auth-gateway/src/middleware/api-key-auth.ts:39-46`; `readme.md:98-105`; `apps/workers/shared/src/http.ts:9-15` | Users may send an unsupported Bearer header and fail both authentication and browser preflight. | Correct all surfaces to `requiems-api-key` or separately implement/test Bearer support. |
| MEDIUM | Worker API responses have no explicit cache policy. | `apps/workers/shared/src/http.ts:21-33`; `apps/workers/auth-gateway/src/http.ts:50-65` | Future or existing Cloudflare cache rules could cache authenticated or operational responses incorrectly. | Export cache-rule state and require `no-store`/private behavior and `Age`/`CF-Cache-Status` probes. |
| MEDIUM | Generated security material claims HSTS, but repository ingress configuration does not prove the header is emitted. | `scripts/generate-docs.mjs:437-444`; no `Strict-Transport-Security` ingress match | Clients may receive a security policy different from published claims, and an incorrect preload policy complicates rollback. | Verify actual headers and deliberately choose HSTS scope/preload policy. |
| MEDIUM | Worker error handling returns internal exception messages in the global Hono handler. | `apps/workers/shared/src/middleware/error-handler.ts:10-23` | New route failures could disclose internal details to public callers. | Return generic production errors while logging details; add error-body assertions. |
| MEDIUM | Rails public proxy/demo surfaces also return backend exception and timeout messages. | `apps/dashboard/app/controllers/api_proxy_controller.rb:33-40`; `apps/dashboard/app/services/api_proxy_service.rb:53-58`; `apps/dashboard/app/controllers/tool_demos_controller.rb:769-778` | Moving these POST surfaces behind an old-host grace path can preserve public internal-error disclosure. | Include Rails proxy/demo error bodies in the security review and return only approved public messages. |
| MEDIUM | Forwarded client-IP handling trusts an incoming header when Cloudflare does not provide `CF-Connecting-IP`. | `apps/workers/auth-gateway/src/http.ts:13-31`; Go IP handlers read `X-Forwarded-For` first | Direct or spoofed requests can alter IP-dependent API behavior and telemetry. | Require Cloudflare-only ingress, overwrite forwarding headers, and test spoofed direct requests. |
| MEDIUM | API usage rows do not record request hostname. | `apps/workers/auth-gateway/schema.sql:3-12`; `apps/workers/auth-gateway/src/requests.ts:75-90` | D1 cannot measure legacy-to-new host adoption without mixing it into billing semantics. | Use edge hostname analytics/log labels; keep quota/billing aggregation unchanged. |
| MEDIUM | Product-host SEO output has request-derived/hardcoded inconsistencies, stale sitemap artifacts, private-page robots gaps, and potentially false localized blog alternates. | `apps/dashboard/app/helpers/application_helper.rb:403-484`; `apps/dashboard/config/sitemap.rb:8,59-124`; `apps/dashboard/public/robots.txt:1-8`; `apps/dashboard/app/models/blog_post.rb:3-21` | Redirect chains, duplicate content, stale indexing, and incorrect canonical/hreflang signals. | Centralize fixed product URLs, regenerate and freshness-check sitemap/robots/XSL, and test actual translations/private paths. |
| MEDIUM | MCP has separate fetch, snapshot, and runtime URL contracts, and its `--base-url` generation flag is unused. | `apps/mcp/scripts/fetch-spec.ts:1-11`; `apps/mcp/openapi.json:1`; `apps/mcp/scripts/generate.ts:74-95,492-517`; `apps/mcp/generated/runtime.ts:36-44,73-96` | Updating one layer can leave MCP fetching redirects/HTML or calling the wrong host. | Configure fetch and runtime separately, regenerate snapshot/tools as appropriate, and test a live MCP call. |
| MEDIUM | OpenAPI, MCP, and Go already disagree on at least one operation path. | `apps/workers/auth-gateway/src/generated/openapi.ts:332`; `apps/mcp/generated/tools/text_advice.ts:2,24`; `apps/api/app/routes_v1.go:42-47` | Host migration regeneration can silently publish or preserve path drift and break generated clients. | Diff `(method,path,operationId)` before/after generation and fail publication on unexplained loss or reassignment. |
| MEDIUM | Mailer and checkout links contain independent old-host contracts outside `MAILER_HOST`. | `apps/dashboard/app/mailers/private_deployment_mailer.rb:7,30`; `apps/dashboard/app/controllers/private_deployments_controller.rb:56-61`; `apps/dashboard/app/mailers/application_mailer.rb:4` | Password/payment/customer links can remain on the old host or depend on an old-host grace path. | Audit browser-link host, sender domain, checkout return URL, and mail transport separately; test pre- and post-cutover checkouts. |
| MEDIUM | Private-deployment tenant URLs use the `*.requiems.xyz` wildcard namespace. | `apps/dashboard/app/models/private_deployment_request.rb:83-85`; `apps/dashboard/app/views/private_deployment_mailer/deployment_ready.html.erb:20-35` | Blindly changing the product domain could break separately hosted tenant deployments. | Preserve the tenant namespace unless it receives a separate migration plan and DNS/certificate design. |
| MEDIUM | Rails marks all production public files one-year immutable, including robots and sitemap artifacts that must change during cutover. | `apps/dashboard/config/environments/production.rb:10-11`; `apps/dashboard/public/robots.txt:8`; `apps/dashboard/public/sitemap.xml:8` | Regenerated SEO files may remain stale at edge/client caches, and rollback may be obscured by cached content. | Apply short-lived/no-cache migration rules, purge old/new sitemap and robots URLs, and verify cache headers before Search Console submission. |
| MEDIUM | New-host health ownership is unresolved: the Worker owns `/healthz` only on its current Custom Domain while Rails exposes `/up`. | `apps/workers/auth-gateway/src/index.ts:20-28`; `apps/dashboard/config/routes.rb:9-16` | A public probe can test a redirect/origin health endpoint while the API route is unavailable, or the new host can have no promised health endpoint. | Select Worker- or Go-owned `/healthz` and `/openapi.json` contracts per architecture branch; assign probes and test bare/trailing-slash variants. |
| MEDIUM | Certificate ownership and renewal for `requiemsapi.com` are absent from repository configuration. | Existing host settings in `infra/kamal/deploy.dashboard.yml:15-18` and `infra/caddy/Caddyfile:5-23`; no new-domain DNS/certificate files | New product traffic can fail TLS or certificate renewal, and removing legacy certificates can break compatibility. | Name Cloudflare/registrar/origin certificate owners, issuance/renewal path, CAA/DNSSEC state, and deletion rollback. |
| LOW | Many repository docs, fixtures, examples, and performance reports intentionally record the old host. | `readme.md`; `docs/apis`; `tests/reports`; `tests/integration`; `apps/dashboard` views | Stale examples and reports reduce developer confidence but do not alone cause outage. | Update live docs/examples; label historical reports rather than rewriting history. |

## 10. Multi-Agent Review History

Lead research passes before independent review:

- Repository cross-check established the Worker Custom Domain versus apex Rails
  split, exact `/v1` versus `/v1/*` route requirement, `API_BASE_URL`'s lack of
  runtime consumption, generated/SEO surfaces, and Kamal-versus-Caddy ingress
  ambiguity.
- Lead adversarial pass rechecked Rails routes, Devise/session behavior,
  mailers, webhook tests, API proxy, MCP/OpenAPI generation, Worker CORS, Go
  secret middleware, CI/CD, and SEO files before accepting reviewer reports.

Review round 1 (independent subagents; no files edited by reviewers):

- Infrastructure reviewer: confirmed the redirect mechanism, old-apex Rails
  grace ingress, Worker deployment gap, absent isolated preview assumption,
  route ownership, health-check, certificate, cache, and rollback requirements.
- Application reviewer: confirmed the complete old-apex POST inventory,
  session boundary, raw-body webhook handling, internal playground contract,
  API_BASE_URL non-consumption, hardcoded mail links, request-derived payment
  return URLs, MCP layers, and private-tenant wildcard.
- API compatibility reviewer: confirmed permanent no-redirect dual-host API
  support; found illustrative API path drift, Bearer/header contradictions,
  binary and proxy parity gaps, usage-header exposure, stale error URLs, and
  D1's lack of hostname attribution.
- Security reviewer: confirmed old-apex POST redirect risk, absent production
  host policy/trusted-proxy evidence, public-origin/secret boundary requiring
  external lockdown verification, webhook replay/idempotency gap, wildcard
  CORS, absent explicit cache policy, unverified HSTS claim, error disclosure,
  and forwarded-IP spoofing tests.
- Developer-experience reviewer: confirmed Worker deployment is outside CD,
  OpenAPI/MCP URL layers, inert MCP base-url flag, external SDK/skills
  dependencies, LLM feeds, incorrect README auth example, local MCP defaults,
  and missing host-split integration tests.
- SEO/public-web reviewer: confirmed locale redirect chains, stale sitemap and
  robots hosts, request-derived structured data, canonical/OG inconsistency,
  dashboard URL ambiguity, private-page indexing gaps, blog hreflang risk, and
  LLM-feed migration requirements.

Lead verification of round 1:

- Accepted findings only where the cited repository files confirmed the claim.
- Merged duplicate SEO, MCP, webhook, session, and POST findings into the
  execution sections and findings table.
- Rejected live DNS resolution and current production response observations as
  repository findings; retained them as human-verification gates.
- Kept permanent legacy API support as the recommendation because no evidence
  establishes that external clients follow redirects or that a sunset program
  exists.

Adversarial review (round 1 extension):

- Found and verified a cross-document architecture conflict: the separate
  architecture audit recommends retiring the Worker/KV/D1/API-Management
  layer, while this plan initially recommended permanent Worker dual-host
  support. Added a blocking architecture decision gate.
- Found and verified that the documented Worker staging command can use the
  production KV/D1 IDs. Added an isolated-staging gate and prohibited using the
  current staging label as evidence of safe canarying.
- Found and verified one-year immutable caching on Rails public files. Added
  migration cache-control, purge, and freshness checks for robots/sitemaps.

Review round 2 (convergence review; three independent re-reviewers):

- Infrastructure convergence: found that route ownership, old-apex grace
  ingress, Worker staging/deployment, rollback, health ownership, cache rules,
  and certificate ownership were still procedural rather than executable. The
  lead added named change-ticket fields, explicit branch contracts, route and
  rollback artifacts, and owner/acceptance requirements.
- Security convergence: found cookie forwarding across the API boundary,
  Rails forwarded-IP and error-disclosure surfaces, broader webhook side
  effects, explicit cookie-attribute verification, and the `/locale` referer
  redirect sink. The lead verified each cited file and added controls/tests.
- Compatibility convergence: found the admin portion of the POST inventory was
  incomplete, architecture sequencing was still too open-ended, D1 adoption
  measurement needed retention/query/reconciliation details, OpenAPI branch
  selection was incomplete, auth-header drift was broader than README/helper
  text, and external SDK release ownership was missing. The lead added the
  complete route inventory requirement, two executable architecture branches,
  durable metric requirements, localized/auth completeness checks, and SDK
  ownership gates.

Lead verification of round 2:

- Verified the admin resource routes at `apps/dashboard/config/routes.rb:126-149`,
  the Worker cookie behavior, Rails proxy/error/IP behavior, architecture audit
  conflict, Wrangler identifiers, and production immutable cache header.
- Accepted the findings as plan gaps; no reviewer finding was accepted solely
  from external state. External Cloudflare owners, certificates, DNS, and live
  traffic remain human-verification dependencies.
- The revised plan now has explicit Worker-retention and Go-cutover branches;
  implementation must choose one before code or DNS work begins.

Final adversarial review:

- Found and verified that the plan still described Cloudflare precedence too
  loosely. Replaced that with a preferred single apex dispatcher Worker, or a
  fully explicit Redirect Rule exclusion expression and Rules-phase/Trace gate;
  added defined HTTP behavior and exact route tests.
- Found and verified existing OpenAPI/MCP/Go advice-path drift. Added a required
  normalized `(method,path,operationId)` before/after diff and publication gate.
- No further repository-backed migration gap was reported after those changes.
  Remaining DNS, Cloudflare account, active ingress, certificates, traffic,
  provider, monitoring, Search Console, and architecture-branch facts are
  explicitly listed as external human gates, not assumed.

## 11. Final Confidence Assessment

Final assessment after repository research, two multi-agent review rounds, lead
verification, and a final adversarial pass.

- Independent review rounds completed: 2, plus 1 final adversarial pass.
- Reviewer roles completed: infrastructure, application, API compatibility,
  security, developer experience, SEO/public web, and adversarial.
- Verified HIGH findings remaining: 17; these are pre-implementation gates,
  not authorization to proceed without resolution/acceptance.
- Verified MEDIUM findings remaining: 20.
- Verified LOW findings remaining: 1.
- Review convergence: no new relevant HIGH/MEDIUM repository findings remained
  after the final adversarial changes; the legacy API no-redirect strategy was
  independently upheld.
- Unresolved decisions: Worker-retention versus Go-cutover architecture branch;
  exact `/dashboard`, `/healthz`, and `/openapi.json` contracts; full old-apex
  POST disposition; event-level webhook idempotency; cookie forwarding policy;
  auth-header contract; SDK publication decision; and API path-alias scope.
- Human verification required: Cloudflare zone/routes/Rules/Cache/WAF state;
  DNS/registrar/DNSSEC/CAA; certificates and renewal; active VPS ingress;
  origin firewall; Worker/KV/D1 staging resources; customer/API traffic;
  Lemon Squeezy callback configuration; email provider/DNS; monitoring and
  retention; Search Console; external SDK/skills repositories; and customer
  integrations.

## 12. Final Implementation Checklist

### Contract and discovery

- [x] Complete independent infrastructure, application, API-compatibility,
      security, developer-experience, SEO, and adversarial review rounds; verify
      findings against repository evidence.
- [ ] Resolve the architecture branch: Worker-retention or Go-cutover. Record
      ownership of `/v1`, auth, quotas, usage, OpenAPI, MCP, and legacy support.
- [ ] Approve exact product, API, dashboard, legacy API, webhook, health, and
      OpenAPI URL contracts.
- [ ] Generate and classify the complete Rails method/path inventory, including
      every POST/PATCH/PUT/DELETE old-apex route.
- [ ] Verify live DNS, Cloudflare routes/custom domains/rules, certificates,
      active VPS ingress, and origin health checks.
- [ ] Export traffic/client/webhook/monitoring baselines.
- [ ] Back up relevant Worker route/binding settings and database/usage data.
- [ ] Name owners and acceptance artifacts for Cloudflare routes, redirects,
      cache/WAF rules, certificates, Worker deploy, DNS, and rollback.
- [ ] Provision isolated staging Worker/KV/D1/secrets/origin/hostname before
      any canary; prohibit the current production-ID staging command.

### Product host

- [ ] Provision and certificate-verify `requiemsapi.com`.
- [ ] Deploy Rails to the new host without changing old traffic.
- [ ] Decide `/dashboard` versus localized canonical and test all auth/account
      paths.
- [ ] Update public URL generation, mail links, analytics, JSON-LD, OpenGraph,
      canonical, hreflang, sitemap, robots, and LLM feeds.
- [ ] Keep mail transport DNS unchanged unless separately approved.

### API host/path

- [ ] Add exact Cloudflare route(s) for `requiems.xyz/v1` and
      `requiems.xyz/v1/*` to the selected API owner, with precedence verified.
- [ ] Implement the selected non-redirecting `/healthz` and `/openapi.json`
      surfaces on the selected API owner; keep legacy spec/health URLs live.
- [ ] Prove method/body/query/header/path/response parity with legacy host.
- [ ] Keep `api.requiems.xyz` as a no-redirect dual-host API surface.
- [ ] Preserve KV/D1, usage, quota, rate-limit, and backend-secret behavior.

### Webhooks and external systems

- [ ] Update Lemon Squeezy callback to the product host.
- [ ] Retain and test old-host POST handling during the grace period.
- [ ] Verify email confirmation/reset/account links and sender deliverability.
- [ ] Update MCP production base URL and regenerate its spec/tools.
- [ ] Update external monitors, Search Console, support docs, SDK/customer
      configuration, CI secrets, and provider callbacks.

### Redirects, documentation, and SEO

- [ ] Configure API route precedence before generic apex redirects.
- [ ] Redirect old public GET/HEAD paths path-preservingly to product host.
- [ ] Exclude API, webhook, operational, and required machine endpoints from
      generic redirect logic.
- [ ] Preserve/proxy/retire every old-apex state-changing route, including
      admin-generated routes and `/locale`; fuzz Referer-derived redirects and
      encoded redirect targets.
- [ ] Update API docs to `https://requiems.xyz/v1` and product docs to
      `https://requiemsapi.com` by semantic category.
- [ ] Regenerate OpenAPI, MCP, LLM, API docs, and sitemap artifacts.
- [ ] Decide and assign ownership for external SDK/agent-skill updates; publish
      tested versioned releases or explicitly defer and label them unsupported.
- [ ] Submit new sitemap and monitor canonical/indexing/redirect behavior.

### Validation and rollback

- [ ] Run the full host/path/API compatibility matrix with real test keys.
- [ ] Test cookies, CSRF, login, logout, confirmation, reset, dashboard, and
      webhook signatures.
- [ ] Verify cookie filtering to Go, CORS/auth-header completeness, forwarded-IP
      provenance, generic production error bodies, HSTS, cache headers, and
      direct-origin rejection.
- [ ] Add durable hostname-labeled edge telemetry with retention, query owner,
      and reconciliation against D1 usage totals.
- [ ] Run API integration/load/synthetic tests against both API hosts.
- [ ] Monitor Worker, Rails, Go, Redis, D1, API Management, MCP, DNS/TLS,
      webhook, and SEO metrics through at least one billing period.
- [ ] Keep a tested rollback for Worker route, product host, redirects, and
      webhook provider configuration.
