# Requiems Domain Migration Audit

Status: audited implementation plan; audit-only, no implementation performed.

Scope amendment: the project owner confirms there are currently no users,
customers, or meaningful public positioning. Backwards compatibility, SEO
preservation, old-domain redirects, and customer migration are therefore
explicitly out of scope unless a later decision re-enables them.

Authoritative migration policy: this owner-approved clean-break policy
supersedes the initially described root redirect and any earlier draft language
about dual-host support or an old-apex grace period. The route behavior matrix
in Section 3 is the source of truth for implementation, tests, deployment, and
rollback.

Audit date: 2026-08-21

Repository: `requiems-api`, branch `v2/requiems-api-v2`

## 1. Executive Summary

The repository does not currently implement the requested target architecture.
The current repository-described production topology is split across five named
hostnames and two documented VPS deployment paths:

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

The initially requested target requires one apex host, `requiems.xyz`, to have
path-dependent behavior. Under the owner-approved clean break, `/v1*` executes
the API gateway, explicitly retained operational endpoints execute on the API
gateway, and every other apex request is retired with 404/410 rather than
redirected. The repository currently has no such host/path split. The current
apex is routed to Rails by Kamal/Caddy, while the API Worker owns all paths on
the separate `api.requiems.xyz` custom domain.

Overall risk is MEDIUM for a clean break, rising to HIGH only if the active
ingress, DNS/TLS ownership, payment-provider callback, or API security boundary
is changed without verification. The recommended clean-break posture is:

1. move Rails to `requiemsapi.com` as the only public product/dashboard host;
2. expose the current API implementation only at `requiems.xyz/v1/*` (plus
   explicitly selected health/spec endpoints);
3. retire `api.requiems.xyz` without redirecting it after internal references
   and provider settings are updated;
4. do not configure generic old-page redirects or SEO migration machinery;
5. return an explicit 404/410 for every non-API path and method on
   `requiems.xyz`; no root or old-page redirect is part of this migration.

This is a recommendation, not a completed migration. The exact Cloudflare route
and DNS records must be confirmed against the production account before any
configuration is changed.

The main unresolved implementation gates are the active Cloudflare/VPS ingress,
the exact new API route owner, the payment webhook cutover, the Worker deploy/
rollback path, and whether the separate Go-authentication architecture work is
included or sequenced later.

## 2. Current Architecture

### 2.1 Domain and service map

| Host                          | Repository evidence                                                                                                         | Current responsibility                                                                              | Confidence                                                       |
| ----------------------------- | --------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------- |
| `requiems.xyz`                | `infra/kamal/deploy.dashboard.yml:15-18`; `infra/caddy/Caddyfile:5-23`; `apps/dashboard/config/routes.rb:64-214`            | Rails website, locale routes, Devise auth, dashboard, admin, API docs, webhooks, playground         | High in repository; active ingress needs production verification |
| `api.requiems.xyz`            | `apps/workers/auth-gateway/wrangler.toml:38-42`; `apps/workers/auth-gateway/src/index.ts:20-30`                             | Public Cloudflare auth/rate/quota/usage gateway, including `/v1/*`, `/healthz`, and `/openapi.json` | High                                                             |
| `internal.requiems.xyz`       | `infra/kamal/deploy.api.yml:10-18`; `infra/caddy/Caddyfile:31-47`                                                           | Go API origin, protected by `X-Backend-Secret`                                                      | High in repository; active proxy needs verification              |
| `api-management.requiems.xyz` | `apps/workers/api-management/wrangler.toml:31-35`; `apps/dashboard/app/services/cloudflare/api_management_service.rb:12-14` | Internal API-key management, usage export, analytics, Swagger                                       | High                                                             |
| `mcp.requiems.xyz`            | `infra/kamal/deploy.mcp.yml:10-15`; `apps/mcp/src/server.ts:108-150`                                                        | MCP HTTP server; generated handlers call the configured Requiems API base                           | High                                                             |
| `requiemsapi.com`             | No repository references found                                                                                              | Target product host; not configured in repository                                                   | High                                                             |

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
(`apps/api/app/app.go:35-51`). It therefore cannot implement the requested apex
redirect/API split by itself.

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
  `requiems.xyz`, and MCP to `mcp.requiems.xyz` (`infra/kamal/deploy.*.yml`).
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

Cloudflare's current documentation distinguishes a Custom Domain, which owns all
paths for a hostname, from a Route, which runs in front of an existing proxied
origin. A route requires a proxied DNS record; a Custom Domain creates the
Worker-facing DNS/certificate state. This distinction is directly relevant to
the apex split and is not encoded in the current repository.

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
default Rails session cookie is therefore expected to be host-scoped unless the
deployed environment overrides it externally.

Moving the Rails app from `requiems.xyz` to `requiemsapi.com` changes the
registrable domain. A host-only cookie will not be sent to the new host, and a
cookie scoped to `.requiems.xyz` cannot be shared with `requiemsapi.com`. Users
should be expected to authenticate again unless a separately designed,
short-lived, one-time cross-domain handoff is implemented. A shared cookie is
not a safe or technically valid solution across these two registrable domains.

No OAuth provider or OAuth callback route was found in the repository. Devise
confirmation, password reset, account deletion, and other mail links use Rails
mailer URL options. The deployed values currently include
`MAILER_HOST: requiems.xyz`, `SMTP_DOMAIN: mail.requiems.xyz`, and
`SMTP_FROM_EMAIL: noreply@mail.requiems.xyz`
(`infra/kamal/deploy.dashboard.yml:60-65`;
`apps/dashboard/config/environments/production.rb:26-40`). The web host must
change in links, while the mail transport domain must only change if email DNS
and deliverability are intentionally migrated.

The Lemon Squeezy webhook is a POST route at the old apex
(`apps/dashboard/config/routes.rb:52-54`;
`docs/core/lemonsqueezy-webhook-setup.md:56`). An indiscriminate old-host
redirect would risk dropping or altering webhook requests and must not be used
as the first cutover behavior.

### 2.6 API authentication, CORS, and compatibility behavior

The public API uses the `requiems-api-key` request header. The Worker strips
that header before forwarding and adds the backend secret. Worker responses
include `Access-Control-Allow-Origin: *`; the preflight response allows
`Content-Type` and `requiems-api-key` (`apps/workers/shared/src/http.ts:5-15`,
`apps/workers/auth-gateway/src/http.ts:10-33`). The gateway does not use cookies
for authentication, so API authentication is header-based rather than
browser-session based; however, `filterHeaders` currently forwards an incoming
`Cookie` header to the Go origin
(`apps/workers/auth-gateway/src/http.ts:16-31`). A same-origin request from
`requiems.xyz` can therefore carry a Rails cookie across the API boundary. This
must be an explicit reviewed decision and a tested header-filtering rule.

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

SEO migration is explicitly out of scope under the current project premise. The
repository does contain canonical, sitemap, robots, JSON-LD, OpenGraph, and LLM
surfaces, but they do not need a backward-compatible migration. They may be
updated opportunistically for consistency on the new product host, or
retired/marked noindex if the public site is not being positioned yet. No SEO
redirect, sitemap submission, canonical preservation, or search-console work is
a release gate.

## 3. Target Architecture

The intended logical contract is:

```text
https://requiemsapi.com/*
    -> Rails public product, auth, dashboard, docs, account, webhooks

https://requiems.xyz/v1/*
    -> Auth Gateway Worker -> internal Go API

https://requiems.xyz/healthz
https://requiems.xyz/openapi.json
    -> Auth Gateway operational/documentation endpoints

https://requiems.xyz/<non-API public path>
    -> explicit 404/410; no generic redirect

https://api.requiems.xyz/v1/*
    -> retired; no redirect
```

### 3.0 Authoritative `requiems.xyz` method/path matrix

The following matrix is normative for this migration. Every result is for
`https://requiems.xyz`; no row redirects to `requiemsapi.com`.

| Path class              | GET/HEAD                                                         | POST/PUT/PATCH/DELETE                | OPTIONS or other/unknown methods                                                                             |
| ----------------------- | ---------------------------------------------------------------- | ------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `/`                     | 404/410, no body for `HEAD`                                      | 404/410                              | 404/410                                                                                                      |
| Any public non-API path | 404/410, no body for `HEAD`                                      | 404/410                              | 404/410                                                                                                      |
| `/v1` and `/v1/*`       | Forward to Auth Gateway/API contract                             | Forward to Auth Gateway/API contract | `OPTIONS` is the API CORS preflight; unsupported methods return the API's 405/404 response, never a redirect |
| `/healthz`              | Auth Gateway health response; `HEAD` follows the health contract | 405/404                              | 405/404                                                                                                      |
| `/openapi.json`         | Auth Gateway OpenAPI response; `HEAD` follows the spec contract  | 405/404                              | 405/404                                                                                                      |
| Unknown path            | 404/410                                                          | 404/410                              | 404/410                                                                                                      |

The API implementation must preserve its existing method/body/query/header
semantics for `/v1` and `/v1/*`; the edge must not turn unsupported API methods
into redirects. The selected API owner is the Auth Gateway Worker, with the Go
service remaining private behind `X-Backend-Secret`.

### 3.1 Endpoint lifecycle

| Host/path                     | Owner                                 | Current state                                           | Transition                                                                          | Retirement condition                                                                                                             |
| ----------------------------- | ------------------------------------- | ------------------------------------------------------- | ----------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- |
| `requiemsapi.com/*`           | Rails/Kamal product service           | Not configured in repository                            | Add DNS/TLS and product ingress; validate auth, dashboard, webhooks, and mail links | Separate product shutdown decision                                                                                               |
| `requiems.xyz/v1` and `/v1/*` | Auth Gateway Worker -> private Go API | Not routed on apex; Worker owns old custom domain       | Add explicit proxied apex routes and update active internal consumers               | Separate API shutdown decision                                                                                                   |
| `requiems.xyz/healthz`        | Auth Gateway Worker                   | Not exposed on apex; exists on old Worker custom domain | Publish as the canonical API health endpoint with the new API route                 | Retire only when the API route and its probes are retired                                                                        |
| `requiems.xyz/openapi.json`   | Auth Gateway Worker                   | Not exposed on apex; MCP/spec automation uses old host  | Publish as the canonical spec endpoint; update MCP fetch/generation                 | Retire only after MCP and all spec consumers are removed or repointed                                                            |
| `requiems.xyz/<non-API>`      | Apex edge retirement handler          | Rails currently owns these paths                        | Replace Rails exposure with explicit 404/410; do not redirect                       | Remains the terminal response for this migration                                                                                 |
| `api.requiems.xyz/*`          | Existing Auth Gateway Custom Domain   | Current public API, health, and spec host               | Update consumers to new host, verify no redirect, then remove Custom Domain/DNS     | DNS is absent and no active certificate/route remains; if the provider requires a temporary resolving state, return 404/410 only |

This table is also the contract for MCP generation, health/spec probes,
deployment acceptance, and rollback. Rollback restores the selected new-host
route/version and its endpoint ownership; it does not restore customer-facing
legacy compatibility.

### 3.2 Recommended ingress design

Use Cloudflare path Routes on the proxied apex origin for the new API paths:

1. Route the exact `/v1` path and `/v1/*` to the selected API implementation.
   Preserve the original path, method, query, body, API-key header, and response
   behavior.
2. Route `/healthz` and `/openapi.json` explicitly to the Auth Gateway Worker
   and update internal consumers before retiring the old host.
3. Return explicit 404/410 for every non-API apex path and method. Do not add a
   generic Redirect Rule. If HTTP is enabled, define HTTPS behavior separately
   and test API POST handling rather than inheriting an undocumented redirect.
4. Retire the old API Custom Domain and old apex Rails surface only after the
   new product/API paths and any active Lemon Squeezy callback are verified.

Cloudflare routes are appropriate here because the Worker is being placed in
front of an existing apex origin. The current Worker Custom Domain model is
appropriate for `api.requiems.xyz`, where the Worker owns the hostname.

### 3.3 Public product host

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
├── external consumers (not measurable from repo; no customer migration required)
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
└── target API path /v1/* must be carved out from apex non-API retirement

requiemsapi.com (new product host)
├── new DNS/TLS/Cloudflare/Kamal ingress
├── all public Rails pages
├── auth and account lifecycle
├── dashboard and admin
├── email links and payment callback destinations
└── optional public metadata surface; SEO migration is out of scope

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

| Class                | Verified dependencies                                                                                                                 |
| -------------------- | ------------------------------------------------------------------------------------------------------------------------------------- |
| Application          | Rails host/URL generation, locale routing, dashboard paths, webhook route, server-side API proxy, Go `/v1` routing, Worker proxy/auth |
| Infrastructure       | Cloudflare Worker custom domain/routes, DNS records, TLS, Kamal proxy, optional Caddy, VPS origin, Cloudflare zone configuration      |
| Configuration        | `API_BASE_URL`, `INTERNAL_API_URL`, `MAILER_HOST`, `REQUIEMS_BASE_URL`, Wrangler routes, Kamal proxy host, Caddy host blocks          |
| Documentation        | README, docs/core, docs/apis, Rails API docs, LLM feeds, generated OpenAPI, MCP README/spec                                           |
| Security             | API-key header, backend secret, CSRF/session cookies, redirect handling, host/path precedence, webhook signature endpoint             |
| Compatibility        | Clean-break retirement of legacy API host, internal generated clients, method/body/header correctness on the new host                 |
| Operational          | Worker health/openapi endpoints, Go health, Kamal/Caddy health checks, Sentry, uptime/traffic dashboards outside repo                 |
| SEO                  | Explicitly out of scope under the current project premise                                                                             |
| Developer experience | curl/examples, integration defaults, load tests, MCP generation, API reference and quick-start snippets                               |

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
   separately; do not combine two large migrations without an owner and rollback
   boundary.
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
- Change the production deployment host from `requiems.xyz` to `requiemsapi.com`
  only after the new host is serving a validated build.
- Decide and implement the dashboard URL contract. To satisfy the stated target
  literally, expose `/dashboard` as a stable entry point and define how it
  selects/redirects to the user's locale; otherwise explicitly document
  `/en/dashboard` (or the user's locale) as canonical.
- Preserve the Rails route namespace and all auth/account/admin paths during the
  host move.
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
  Change it only as part of an explicit configuration contract and do not assume
  changing it updates generated docs or UI examples. API documentation should
  use `https://requiems.xyz` with the `/v1` path contract, not
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
  confirmation, reset-password, account deletion, and CSRF flows work on the new
  host.
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
- Make the selected API implementation available only on the new apex API route.
  Verify the new route's methods, bodies, query strings, API-key headers, usage
  headers, error statuses, and binary responses.
- Publish `/healthz` and `/openapi.json` on `requiems.xyz` as explicit Auth
  Gateway routes and update all internal consumers before retiring the old host.
  These are retained operational contracts, not fall-through paths.
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
  request-derived URL helpers
  (`apps/dashboard/app/helpers/application_helper.rb:172-187,403-484`,
  `apps/dashboard/app/controllers/private_deployments_controller.rb:56-61`).
  Define accepted hosts, trusted proxy handling, `Host`/ `X-Forwarded-Host`
  behavior, and fixed production absolute URLs before enabling the new host.
- Preserve the internal playground contract: Rails uses `INTERNAL_API_URL` and
  `X-Backend-Secret`, not the public API base. Do not route demos through the
  public gateway unless metering and authentication are deliberately redesigned.
- Treat payment return URLs and private-deployment tenant URLs separately.
  Checkout return URLs are request-host-derived, while tenant URLs use the
  `*.requiems.xyz` namespace
  (`apps/dashboard/app/models/private_deployment_request.rb:83-85`). Do not
  blindly migrate either surface.
- Fix the authentication contract as part of the documentation migration: the
  gateway accepts only `requiems-api-key`, while the root README, FAQ, and
  generated security material advertise Bearer authentication
  (`apps/workers/auth-gateway/src/middleware/api-key-auth.ts:39-46`,
  `readme.md:98-105`,
  `apps/dashboard/app/helpers/application_helper.rb:110-116`). Either add and
  test Bearer support as a separate API change or correct all first-party
  surfaces to the existing header; test preflight for the chosen contract.
- Do not treat the requested illustrative paths as existing API aliases. The
  current Go routes are `/v1/validation/email`, `/v1/validation/phone`,
  `/v1/networking/ip`, and `/v1/systems/signup/protect`; adding shorter aliases
  is outside a domain-only migration and needs a separate product decision.
- Treat `POST /locale` as an existing redirect sink in its own security test: it
  derives the destination from `Referer` and preserves query/fragment
  (`apps/dashboard/app/controllers/locale_controller.rb:17-36`). Fuzz external,
  encoded, malformed, and cross-host referers on the new product host; no old
  apex redirect rule is required.

### 5.3 Infrastructure and routing changes — HIGH

Implement only after the discovery gates identify the real ingress:

1. Add `requiemsapi.com` to the active product ingress and certificate
   configuration. It must reach the Rails service directly, not the API Worker.
2. Choose one authoritative API ingress mechanism after discovery. No generic
   apex redirect is required. The repository has Kamal and Caddy descriptions
   but no authoritative Cloudflare configuration; do not implement the API route
   in an inactive layer.
3. Validate the product host through the complete Rails stack before changing
   the old host. If no isolated preview hostname or service exists, call this an
   in-place host addition and use a controlled hosts/edge validation path; do
   not assume the existing Kamal deployment creates a parallel preview.
4. Keep the apex DNS record proxied and pointing to the selected API/origin
   path. A Cloudflare Worker Route for `requiems.xyz/v1/*` requires a proxied
   DNS record at that hostname.
5. Attach the selected apex dispatcher/API Worker to both the exact `/v1` path
   and `/v1/*` behavior as required by the Cloudflare route matcher. Specify the
   production zone identifier/ownership, remove the legacy Custom Domain after
   cutover, and prove the Worker receives the original path. Include explicit
   operational/spec routes or dispatcher branches according to the selected
   contract.
6. Switch the Lemon Squeezy callback to the new product host, verify raw-body
   HMAC handling, then disable the old callback and old Rails ingress. No grace
   proxy is required unless the provider configuration cannot switch atomically.
7. Return explicit 404/410 for non-API apex paths and test that no redirect,
   HTML product page, or Rails POST handler is accidentally exposed there.
8. Disable caching for authenticated API responses and webhook requests. The
   change ticket must name the cache owner, path exclusions, origin headers,
   purge operator, TTLs, and acceptance headers (`Age`, `CF-Cache-Status`,
   `Cache-Control`). SEO asset cache changes are out of scope.
9. Add an explicit Worker deployment step. The repository's CD workflow deploys
   API, dashboard, and MCP with Kamal but has no Auth Gateway Wrangler job
   (`.github/workflows/cd.yml:75-149`). Record the deploy owner, Cloudflare
   credentials/approval, `pnpm run generate:openapi`, `pnpm run deploy:prod`,
   route verification, and Worker rollback/version procedure.
10. Lock down the Go origin independently of the shared header: verify firewall
    and Cloudflare-only reachability, reject direct-origin and direct-Worker
    spoof traffic, and rotate the static `BACKEND_SECRET` only with a tested
    dual-secret/maintenance procedure. Kamal exposes `internal.requiems.xyz` and
    Caddy/Go currently rely on the same header
    (`infra/kamal/deploy.api.yml:14-18`, `infra/caddy/Caddyfile:31-46`,
    `apps/api/platform/middleware/auth.go:10-35`).

### 5.4 DNS and TLS — HIGH

Repository-managed:

- Update the active Kamal/proxy hostname for the product service.
- If Caddy is active, add the product host and the final apex host/path rules
  there; if Kamal proxy is active, update Kamal only. Do not modify both
  unconditionally.
- Update Worker Wrangler route configuration for the new route. Retire the
  legacy custom domain after the new endpoint passes validation; confirm
  environment-specific route declarations are complete.
- The Cloudflare change ticket must name the zone ID/name, Worker script and
  environment, exact route patterns (`requiems.xyz/v1` and `requiems.xyz/v1/*`
  if both are needed), proxied DNS record, API cache/WAF rule owner, certificate
  owner, deploy credential/approval owner, and rollback version. These values
  are external and currently absent from the repository; the plan is not
  complete until they are recorded.
- Keep `internal.requiems.xyz`, `api-management.requiems.xyz`, and
  `mcp.requiems.xyz` unchanged unless separate decisions require moving them.

External DNS/Cloudflare:

- Onboard and verify `requiemsapi.com` in Cloudflare, set the required A/AAAA/
  CNAME record to the active product origin, and proxy web traffic.
- Remove the existing `api.requiems.xyz` Worker custom-domain DNS/certificate
  state after the new API route and internal consumers are validated.
- Ensure `requiems.xyz` has a proxied record suitable for the API Worker Route.
- Keep mail MX/TXT/CNAME records and `mail.requiems.xyz` sender authentication
  unchanged unless email ownership is intentionally migrated.
- Verify Cloudflare edge certificates for both new product and API hosts, the
  origin certificate/SSL mode, Caddy/Kamal certificate behavior, and renewal.
- Confirm IPv4/IPv6, DNSSEC, CAA, HSTS, and certificate transparency impacts
  with the infrastructure owner.
- Export the pre-change DNS and certificate state and define deletion order for
  the retired old API host. Verify registrar ownership, CAA, DNSSEC, Cloudflare
  edge certificates, Worker Custom Domain certificates, product host
  certificates, origin certificates, SSL mode, renewal, and any origin IP
  exposure. The repository cannot prove these external facts.
- Verify actual response headers on every promised host. Generated security
  documentation claims HSTS, but no repository ingress configuration proves that
  `Strict-Transport-Security` is emitted (`scripts/generate-docs.mjs:437-444`).
  Decide `includeSubDomains`/preload deliberately and test HTTP-to-HTTPS
  behavior before publishing an HSTS policy.

### 5.5 Backwards compatibility — OUT OF SCOPE

The project owner has explicitly accepted a clean break. Do not preserve
`api.requiems.xyz`, old apex pages, old API paths, old cookies, or old examples
for customer compatibility. After internal consumers and provider settings are
updated, retire the old API Custom Domain. Any old surface that still resolves
during propagation must return 404/410, never redirect; the final state is DNS
absence.

This is not permission to skip correctness on the new system: the new API still
needs method/body/query/header parity within its own contract, and the new Lemon
Squeezy endpoint still needs raw-body HMAC verification. The existing webhook
has no event-ID replay ledger and can repeat side effects
(`apps/dashboard/app/controllers/webhooks/lemonsqueezy_controller.rb:38-64,115-133`),
so either add idempotency or explicitly accept the provider's retry semantics as
a separate payment-risk decision.

### 5.6 Documentation and developer experience — MEDIUM/HIGH

Update the API base URL to `https://requiems.xyz` (with endpoint paths under
`/v1`) in:

- all `apps/dashboard/config/api_docs/*.yml` `base_url` values and handwritten
  examples;
- generated OpenAPI source and regenerated Worker spec;
- `apps/mcp/openapi.json`, MCP fetch URL, production `REQUIEMS_BASE_URL`, and
  any generated artifacts used internally;
- `docs/apis/**/*.md`, root README, docs/core, apps/workers READMEs, and
  examples/quick-start content;
- integration/load test defaults and test fixtures.

Update the public product URL to `https://requiemsapi.com` in active human-
facing links, auth links, account links, and internal docs. Keep email
addresses, private-deployment wildcard domains, MCP host, management host, and
origin host distinct. Historical SEO artifacts do not need migration.

Generated files must be regenerated by their documented generators and then
reviewed for semantic URL changes. Do not do a repository-wide blind string
replacement.

Before generation, select one canonical OpenAPI fetch and advertised server
contract for the selected implementation. Update internal fetchers directly to
the new API host; a fetch that receives HTML or a retired-host response is a
failed build, not a successful regeneration.

Before publishing any regenerated OpenAPI/MCP artifact, produce a normalized
before/after diff of every `(method, path, operationId)` tuple. Fail the release
on endpoint loss, path drift, or operation reassignment unless a separately
approved API change explains it. This is required because the repository already
contains verified drift: the Worker spec advertises `/v1/entertainment/advice`,
while the MCP snapshot/generated tool use `/v1/text/advice` and the Go router
mounts advice under `/v1/entertainment`
(`apps/workers/auth-gateway/src/generated/openapi.ts:332`,
`apps/mcp/generated/tools/text_advice.ts:2,24`,
`apps/api/app/routes_v1.go:42-47`).

Keep the URL layers separate: OpenAPI `servers.url`, the OpenAPI fetch URL, MCP
`REQUIEMS_BASE_URL`, product links, and API example base paths are separate
contracts. The MCP `--base-url` option is parsed but not used by generation
(`apps/mcp/scripts/generate.ts:74-95,492-517`); require an actual runtime
environment/argument override and do not fetch a new spec through an endpoint
that merely redirects. Audit the external SDK repository linked by the README
(`bobadilla-tech/requiems-api-clients`), the external `requiems-api-skills`
repository, SDK release workflows, and any scheduled OpenAPI consumers.

External SDK publication is not a migration gate because there are no users;
update or defer it according to the later product launch plan. Do not imply SDK
readiness from OpenAPI regeneration alone.

Also correct adjacent verified documentation drift during the same review:
`requiems-api.xyz`, `api.requiems.dev`, the root README's Bearer example, the
gateway's product CTA, and the quota-exceeded upgrade URL. Local Compose MCP
defaults must remain environment-specific rather than silently changing local
development to production.

### 5.7 SEO — OUT OF SCOPE

No SEO migration is required under the confirmed project premise. Do not spend
cutover effort on canonical tags, sitemap regeneration, robots changes, Search
Console, hreflang, indexing, LLM-feed migration, or old-page redirects. Update
public metadata or LLM feeds only if needed for the new product to render or for
an active operational consumer; otherwise defer them to later positioning/launch
work. None is a release gate.

### 5.8 Monitoring and operations — HIGH/MEDIUM

Before cutover, create separate dashboards/alerts for:

- `requiemsapi.com` product health and Rails 4xx/5xx/latency;
- `requiems.xyz/v1/*` API success/error/latency, auth failures, quota/rate-limit
  responses, and usage headers;
- Worker invocation/error/CPU metrics and Sentry events for the new route;
- Go `internal.requiems.xyz/healthz` and backend secret failures;
- API Management, MCP, Rails, Sidekiq, database, Redis, and webhook delivery;
- DNS, TLS/certificate renewal, Cloudflare route activation, and origin
  reachability.

No legacy-host adoption metric is required. Keep only the operational metrics
needed to confirm the new API route and clean retirement.

Add synthetic probes for GET and POST API endpoints on the new API host, the
product root, localized pages, `/dashboard`, login, password reset request, the
new webhook endpoint in a safe test mode, `/healthz`, and `/openapi.json`.

Add alerts/probes for cache leakage (`Cache-Control`, `Age`, `CF-Cache-Status`),
actual HSTS headers, redirect loops, internal exception disclosure, and forged
client-IP/forwarded-host headers. The Worker currently has no explicit API
response cache policy, its global error handler returns `err.message`, and its
forwarding logic only overwrites `X-Forwarded-For` when Cloudflare supplies
`CF-Connecting-IP` (`apps/workers/shared/src/middleware/error-handler.ts:10-23`,
`apps/workers/auth-gateway/src/http.ts:13-31`). These should be treated as
security acceptance checks, not assumed properties of the edge.

### 5.9 Tests required — HIGH/MEDIUM

Add or execute a host/path contract matrix before production:

| Request                                                   | Expected result                                                       |
| --------------------------------------------------------- | --------------------------------------------------------------------- |
| `GET https://requiemsapi.com/`                            | Rails product 200/canonical product host                              |
| `GET https://requiemsapi.com/en/`                         | Rails localized product page                                          |
| `GET https://requiemsapi.com/dashboard`                   | Defined dashboard contract; no ambiguity                              |
| `GET https://requiems.xyz/`                               | Explicit 404/410; no generic redirect                                 |
| `HEAD https://requiems.xyz/`                              | Same 404/410 status as `GET`, with no body                            |
| `GET https://requiems.xyz/en/pricing`                     | Explicit 404/410; no product HTML                                     |
| `HEAD https://requiems.xyz/en/pricing`                    | Same 404/410 status as `GET`, with no body                            |
| `POST https://requiems.xyz/en/pricing`                    | Explicit 404/410; no redirect or Rails handler                        |
| `OPTIONS https://requiems.xyz/unknown`                    | Explicit 404/410; no redirect                                         |
| `GET/POST https://requiems.xyz/v1/...`                    | Auth Gateway API behavior, no redirect                                |
| `GET/POST https://api.requiems.xyz/v1/...`                | Before DNS deletion: no redirect; after deletion: no active DNS/route |
| `OPTIONS https://requiems.xyz/v1/...`                     | Correct CORS preflight                                                |
| `GET/HEAD https://requiems.xyz/healthz`                   | Auth Gateway health contract                                          |
| `GET/HEAD https://requiems.xyz/openapi.json`              | Auth Gateway OpenAPI contract; MCP fetch succeeds                     |
| `GET https://api.requiems.xyz/healthz` or `/openapi.json` | Before DNS deletion: no redirect; after deletion: no active DNS/route |
| `POST /webhooks/lemonsqueezy` on new host                 | Signature/body accepted                                               |
| Rails login/logout/confirmation/reset/account deletion    | Correct new-host cookies and URLs                                     |
| MCP HTTP tool call                                        | Calls new API base and propagates caller API key                      |

Extend the matrix with:

- exact `/v1`, `/v1/`, and `/v1/<endpoint>` routes; HTTP-to-HTTPS, query,
  malformed, and unknown-method cases;
- retired old-apex paths return 404/410 without redirecting, while all new-host
  auth/account/admin/webhook routes work;
- direct origin requests without and with forged `X-Backend-Secret`, direct
  Worker requests with forged `CF-Connecting-IP`/forwarded headers, and host
  allowlist/trusted-proxy behavior;
- cookie forwarding from same-origin product requests, Rails `CF-Connecting-IP`
  handling for `/api/proxy` and tool demos, Go `X-Real-IP` fallback, generic
  production errors from both Worker and Rails proxy/demo surfaces, and the
  existing `/locale` Referer-derived redirect sink;
- cache and HSTS headers, generic production error bodies, CORS credentials and
  preflight, `Access-Control-Expose-Headers` expectations, and new-host JSON/
  PNG/binary response behavior, status, usage headers, timeout, and errors;
- duplicated Lemon Squeezy event delivery and the selected idempotency policy.

API contract tests must include GET, POST, PUT/PATCH/DELETE if supported, JSON
bodies, query strings, authorization/API-key header, binary PNG responses,
streaming/large bodies if present, 4xx/5xx responses, rate-limit/quota headers,
and timeout behavior. No cross-host redirect-following tests are required.

## 6. Backwards Compatibility Strategy

### Recommendation: clean break

Do not preserve `api.requiems.xyz`, old apex pages, old API examples, old
cookies, or old paths. Update all known internal consumers to the new canonical
URLs, verify the new product/API surfaces, then remove the old DNS/custom-domain
configuration. During DNS propagation, any still-resolving retired URL must
return 404/410 rather than redirect; the final old-host state is DNS absence.

The only exceptions are active operational dependencies: Lemon Squeezy must be
configured to call the new product endpoint before its old callback is disabled,
and MCP/OpenAPI/CI/monitoring must be updated before their old endpoints are
removed. These are cutover prerequisites, not compatibility support.

## 7. Deployment Plan

1. Freeze the clean-break contract: public host, exact dashboard URL, API
   operational/spec endpoints, no-redirect 404/410 policy, webhook callback,
   email-host policy, and selected API implementation.
2. Complete external discovery gates and save DNS/Cloudflare/Kamal/Caddy
   exports, current certificates, health checks, and a database/Worker
   configuration backup.
3. Provision/verify `requiemsapi.com` DNS, Cloudflare zone, certificates, and
   origin access. Capture the exact edge/origin mechanism; do not assume a
   preview exists.
4. Add the new Rails host to the verified active ingress. If an isolated preview
   is unavailable, validate through a controlled edge/hosts path and explicitly
   label the operation as an in-place host addition. Validate auth, session
   behavior, all public routes, dashboard, admin, mail links, and
   payment/webhook test flows.
5. Deploy the Auth Gateway/OpenAPI Worker explicitly with Wrangler, because the
   current CD workflow does not deploy it. Confirm the production Worker,
   bindings, exact `/v1` and `/v1/*` routes, and the new health/spec contract.
6. Add the new API route for `requiems.xyz/v1/*` in an isolated staging or
   controlled activation state supported by the Cloudflare account. Confirm that
   the apex DNS record is proxied and the route resolves to the correct Worker
   without exposing product pages at `/`.
7. Update API docs, OpenAPI/MCP artifacts, CI defaults, monitoring, and all
   internal provider configuration to the new hosts. Run the endpoint tuple
   diff.
8. Switch Rails to `requiemsapi.com`, switch Lemon Squeezy to the new callback,
   and validate new-host authentication, webhook signatures, and API requests.
9. Before deletion, verify `api.requiems.xyz` and non-API `requiems.xyz` paths
   never redirect. Remove the old Worker Custom Domain, DNS records, and old
   Rails host; afterward verify the old API hostname has no active DNS,
   certificate, or route.
10. Monitor the new product/API, webhook, DNS/TLS, Worker, Rails, Go, MCP, and
    internal management paths through at least one controlled webhook/test cycle
    (or the first real billing cycle if billing becomes active).
11. Close the migration only when the clean-break checklist and external
    ownership checks are signed off.

## 8. Rollback Plan

Rollback must be staged and reversible:

1. If new API path failures occur, disable the new `requiems.xyz/v1/*` route and
   restore the prior known-good API implementation from the saved configuration.
   There is no requirement to restore the retired public API hostname for
   customers.
2. Use a dedicated rollback runbook/workflow with an explicit prior Worker
   version, Wrangler route/binding export, image tag, registry credentials,
   Kamal config, proxy-host state, and acceptance probes. The current composite
   action always runs `kamal setup` and the CD workflow has no rollback input
   (`.github/actions/kamal-deploy/action.yml:44-46`,
   `.github/workflows/cd.yml:17-28`), so a normal deploy rerun is not a rollback
   procedure. For Cloudflare, record the deployed Worker version, deployment ID,
   route bindings, and exact disable/restore operation for the new API route and
   health/spec endpoints. Roll back version and route state together.
3. If the product host fails, restore the last known-good Rails deployment and
   product-host ingress. Do not change `SECRET_KEY_BASE` or invalidate new-host
   sessions unnecessarily.
4. If auth or API behavior differs, roll back the Worker
   deployment/configuration while preserving KV/D1 IDs and secrets. Do not
   recreate the namespace or database as part of rollback.
5. If webhook delivery fails, point Lemon Squeezy back only if the old callback
   is still intentionally retained in the rollback window; otherwise pause
   billing events and replay after signature/idempotency checks.
6. Verify all rollback health checks, API smoke tests, test-key requests,
   dashboard login, mail links, API Management, MCP, and webhook delivery.

Rollback limitations:

- DNS and certificate propagation are not instantaneous.
- External provider settings may require manual restoration.
- There are no customer sessions or API clients to preserve under the confirmed
  scope.
- Keep the new product certificate and a restorable API configuration until the
  incident is closed; deletion of retired DNS is otherwise reversible only by
  re-provisioning.

## 9. Findings

Only verified repository findings are listed here. External state is listed as
an assumption/dependency below, not presented as a repository fact.

| Severity | Finding                                                                                                                                                                                            | Evidence                                                                                                                                                                                       | Impact                                                                                                                                            | Recommended Action                                                                                                                                      |
| -------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------- |
| HIGH     | The current public API is a host-wide Worker custom domain at `api.requiems.xyz`; the apex is Rails, not the API.                                                                                  | `apps/workers/auth-gateway/wrangler.toml:38-42`; `infra/kamal/deploy.dashboard.yml:15-18`; `infra/caddy/Caddyfile:5-23`                                                                        | A hostname change alone cannot create the required apex API path.                                                                                 | Add and test the new `requiems.xyz/v1/*` route, then retire the old Custom Domain.                                                                      |
| HIGH     | The current Lemon Squeezy webhook is bound to the old apex.                                                                                                                                        | `apps/dashboard/config/routes.rb:52-54`; `docs/core/lemonsqueezy-webhook-setup.md:56`                                                                                                          | Retiring the old Rails host before changing the provider callback can stop payment processing.                                                    | Switch and verify the callback on `requiemsapi.com`, then disable the old route.                                                                        |
| HIGH     | Production ingress has two competing repository descriptions; CD deploys Kamal while Caddy/Compose remain documented as production.                                                                | `.github/workflows/cd.yml:75-149`; `infra/kamal/deploy.*.yml`; `infra/docker/docker-compose.yml`; `infra/caddy/Caddyfile`                                                                      | Editing the wrong ingress can cause outage or leave the active path unchanged.                                                                    | Verify live VPS and Cloudflare ingress before selecting files to change.                                                                                |
| HIGH     | Go `/v1` routes are protected only by backend secret middleware and have no public host/path logic.                                                                                                | `apps/api/app/app.go:35-51`; `apps/api/platform/middleware/auth.go`                                                                                                                            | Incorrect edge routing can bypass the gateway or route API requests to Rails.                                                                     | Keep Worker/secret boundary and test direct-origin rejection plus path precedence.                                                                      |
| HIGH     | Rails has no repository-defined production host allowlist, while absolute URLs are partly derived from the request host.                                                                           | No production `config.hosts`; `apps/dashboard/app/helpers/application_helper.rb:172-187,403-484`; `apps/dashboard/app/controllers/private_deployments_controller.rb:56-61`                     | Host-header/proxy mistakes can produce wrong links, redirect abuse, or cross-host security behavior during cutover.                               | Define accepted hosts and trusted proxy policy; use fixed configured production URLs and test Host/X-Forwarded-Host rejection.                          |
| HIGH     | The Go origin is exposed through a public hostname and relies on a static shared `X-Backend-Secret` guard; repository firewall policy does not prove Cloudflare-only reachability.                 | `infra/kamal/deploy.api.yml:14-18`; `infra/caddy/Caddyfile:31-46`; `apps/api/platform/middleware/auth.go:10-35`                                                                                | Direct-origin exposure or secret leakage can bypass the intended gateway boundary.                                                                | Verify origin firewall/Cloudflare-only access, direct-origin rejection, secret rotation, and spoofed-header tests before changing routes.               |
| HIGH     | Webhook HMAC verification has no event-ID ledger or replay window; only one subscription test is business-idempotent.                                                                              | `apps/dashboard/app/controllers/webhooks/lemonsqueezy_controller.rb:38-64,115-133`; `apps/dashboard/test/controllers/webhooks/lemonsqueezy_controller_test.rb:70-82`                           | Concurrent old/new delivery or replay can repeat email/deployment side effects.                                                                   | Keep one active provider callback and add/verify event-level idempotency/replay controls before dual ingress.                                           |
| HIGH     | The Auth Gateway forwards incoming `Cookie` headers to the Go origin even though API authentication is header-based.                                                                               | `apps/workers/auth-gateway/src/http.ts:16-31`                                                                                                                                                  | Apex API requests made in a browser session can carry Rails cookies across the API boundary and expand the impact of a routing or origin mistake. | Decide and test cookie filtering, and verify session-cookie attributes and same-origin API behavior.                                                    |
| HIGH     | The existing CD workflow does not deploy the Auth Gateway Worker, and the repository does not identify an isolated product-host preview.                                                           | `.github/workflows/cd.yml:75-149`; `apps/workers/auth-gateway/package.json:8-13`; `infra/kamal/deploy.dashboard.yml:15-18`                                                                     | A route/config change may never reach production, or a supposed preview may alter the only live ingress.                                          | Add an explicit Wrangler deploy/verification owner and define the actual parallel validation mechanism.                                                 |
| MEDIUM   | The repository’s separate architecture audit proposes retiring the Worker/KV/D1/API-Management layer, while this clean break can use the current Worker for speed.                                 | `docs/audits/2026-08-21-architecture-audit.md:70-84,1357-1360`                                                                                                                                 | Combining both migrations can expand scope and complicate rollback even without customers.                                                        | Use the current Worker for this domain cutover unless Go auth is explicitly in scope; schedule the architecture rewrite separately.                     |
| HIGH     | The documented Worker staging path uses production KV/D1 identifiers and has no staging environment selector.                                                                                      | `apps/workers/auth-gateway/wrangler.toml:7-16,31-42`; `apps/workers/auth-gateway/package.json:8-13`; `apps/workers/auth-gateway/readme.md:173-178`                                             | A canary can mutate production keys, quotas, rate limits, or usage data and does not provide safe route validation.                               | Provision isolated staging resources/hostname/secrets/observability before any canary; prohibit production apex attachment from the staging command.    |
| HIGH     | No repository source names the authoritative Cloudflare zone, apex Worker Route, or route owner.                                                                                                   | `apps/workers/auth-gateway/wrangler.toml:40-42`; no apex route/zone declaration; `infra/kamal/deploy.dashboard.yml:15-18`; `infra/caddy/Caddyfile:5-23`                                        | An engineer could configure an inactive layer or leave the new API path unreachable.                                                              | Record exact zone, Worker, route patterns, DNS record, owners, and route-trace acceptance before implementation.                                        |
| HIGH     | Existing deployment automation has no versioned rollback input or Cloudflare route rollback procedure.                                                                                             | `.github/actions/kamal-deploy/action.yml:44-46`; `.github/workflows/cd.yml:17-28`; no Worker route rollback automation                                                                         | Re-running deployment cannot reliably restore the prior product/API topology; cached redirects and route state may persist.                       | Create a rollback runbook/workflow with image/Worker version, route bindings, credentials, disable/restore operations, and probes.                      |
| MEDIUM   | Rails dashboard URLs are locale-scoped while target names `/dashboard` explicitly.                                                                                                                 | `apps/dashboard/config/routes.rb:64-111`; `ApplicationController#set_locale`                                                                                                                   | A target URL can redirect unexpectedly or fail if not deliberately defined.                                                                       | Decide and test an unlocalized dashboard entry point/canonical URL.                                                                                     |
| MEDIUM   | API CORS is wildcard and header-based; preflight behavior is implemented by Worker middleware.                                                                                                     | `apps/workers/shared/src/http.ts:5-15`; `apps/workers/shared/src/middleware/cors.ts:9-15`                                                                                                      | New route must preserve preflight and API-key header behavior; browser failures may be host-specific.                                             | Test OPTIONS and actual requests from product and third-party origins.                                                                                  |
| MEDIUM   | OpenAPI and `/healthz` are current Worker routes on the old API host but are not yet assigned to the new contract.                                                                                 | `apps/workers/auth-gateway/src/index.ts:20-30`; generated spec URL                                                                                                                             | Removing them without updating internal consumers breaks operations and MCP generation.                                                           | Define new-host endpoints or remove/update internal consumers before retiring the old host.                                                             |
| MEDIUM   | `API_BASE_URL` is configured and validated but has no runtime consumer in the repository.                                                                                                          | `apps/dashboard/app/lib/app_config.rb:164,177-179`; repository-wide `AppConfig.api_base_url` search                                                                                            | Changing the variable alone will not update UI, docs, or generated API examples and can create false confidence in the cutover.                   | Treat it as a configuration contract to either wire explicitly or update only after identifying its intended consumer.                                  |
| MEDIUM   | The target API examples do not match current Go route paths.                                                                                                                                       | `apps/api/app/routes_v1.go:42-95`; validation, networking, and signup-protect transport routes                                                                                                 | Domain-only migration docs could accidentally promise nonexistent endpoint aliases.                                                               | Freeze the migration as domain-only; require a separate decision for path aliases and test only actual routes.                                          |
| MEDIUM   | First-party authentication documentation contradicts the gateway and its CORS preflight.                                                                                                           | `apps/workers/auth-gateway/src/middleware/api-key-auth.ts:39-46`; `readme.md:98-105`; `apps/workers/shared/src/http.ts:9-15`                                                                   | Users may send an unsupported Bearer header and fail both authentication and browser preflight.                                                   | Correct all surfaces to `requiems-api-key` or separately implement/test Bearer support.                                                                 |
| MEDIUM   | Worker API responses have no explicit cache policy.                                                                                                                                                | `apps/workers/shared/src/http.ts:21-33`; `apps/workers/auth-gateway/src/http.ts:50-65`                                                                                                         | Future or existing Cloudflare cache rules could cache authenticated or operational responses incorrectly.                                         | Export cache-rule state and require `no-store`/private behavior and `Age`/`CF-Cache-Status` probes.                                                     |
| MEDIUM   | Generated security material claims HSTS, but repository ingress configuration does not prove the header is emitted.                                                                                | `scripts/generate-docs.mjs:437-444`; no `Strict-Transport-Security` ingress match                                                                                                              | Clients may receive a security policy different from published claims, and an incorrect preload policy complicates rollback.                      | Verify actual headers and deliberately choose HSTS scope/preload policy.                                                                                |
| MEDIUM   | Worker error handling returns internal exception messages in the global Hono handler.                                                                                                              | `apps/workers/shared/src/middleware/error-handler.ts:10-23`                                                                                                                                    | New route failures could disclose internal details to public callers.                                                                             | Return generic production errors while logging details; add error-body assertions.                                                                      |
| MEDIUM   | Rails public proxy/demo surfaces also return backend exception and timeout messages.                                                                                                               | `apps/dashboard/app/controllers/api_proxy_controller.rb:33-40`; `apps/dashboard/app/services/api_proxy_service.rb:53-58`; `apps/dashboard/app/controllers/tool_demos_controller.rb:769-778`    | Retiring the old apex does not remove the disclosure risk if these surfaces are re-exposed on the product host.                                   | Include Rails proxy/demo error bodies in the security review and return only approved public messages.                                                  |
| MEDIUM   | Forwarded client-IP handling trusts an incoming header when Cloudflare does not provide `CF-Connecting-IP`.                                                                                        | `apps/workers/auth-gateway/src/http.ts:13-31`; Go IP handlers read `X-Forwarded-For` first                                                                                                     | Direct or spoofed requests can alter IP-dependent API behavior and telemetry.                                                                     | Require Cloudflare-only ingress, overwrite forwarding headers, and test spoofed direct requests.                                                        |
| MEDIUM   | MCP has separate fetch, snapshot, and runtime URL contracts, and its `--base-url` generation flag is unused.                                                                                       | `apps/mcp/scripts/fetch-spec.ts:1-11`; `apps/mcp/openapi.json:1`; `apps/mcp/scripts/generate.ts:74-95,492-517`; `apps/mcp/generated/runtime.ts:36-44,73-96`                                    | Updating one layer can leave MCP fetching redirects/HTML or calling the wrong host.                                                               | Configure fetch and runtime separately, regenerate snapshot/tools as appropriate, and test a live MCP call.                                             |
| MEDIUM   | OpenAPI, MCP, and Go already disagree on at least one operation path.                                                                                                                              | `apps/workers/auth-gateway/src/generated/openapi.ts:332`; `apps/mcp/generated/tools/text_advice.ts:2,24`; `apps/api/app/routes_v1.go:42-47`                                                    | Host migration regeneration can silently publish or preserve path drift and break generated clients.                                              | Diff `(method,path,operationId)` before/after generation and fail publication on unexplained loss or reassignment.                                      |
| MEDIUM   | Mailer and checkout links contain independent old-host contracts outside `MAILER_HOST`.                                                                                                            | `apps/dashboard/app/mailers/private_deployment_mailer.rb:7,30`; `apps/dashboard/app/controllers/private_deployments_controller.rb:56-61`; `apps/dashboard/app/mailers/application_mailer.rb:4` | New product/payment links can remain on the retired host.                                                                                         | Audit browser-link host, sender domain, checkout return URL, and mail transport separately; test new-host checkouts.                                    |
| MEDIUM   | Private-deployment tenant URLs use the `*.requiems.xyz` wildcard namespace.                                                                                                                        | `apps/dashboard/app/models/private_deployment_request.rb:83-85`; `apps/dashboard/app/views/private_deployment_mailer/deployment_ready.html.erb:20-35`                                          | Blindly changing the product domain could break separately hosted tenant deployments.                                                             | Preserve the tenant namespace unless it receives a separate migration plan and DNS/certificate design.                                                  |
| MEDIUM   | The canonical new-host health/spec routes are not configured; the Worker currently exposes `/healthz` and `/openapi.json` only through its old Custom Domain while Rails separately exposes `/up`. | `apps/workers/auth-gateway/src/index.ts:20-28`; `apps/dashboard/config/routes.rb:9-16`                                                                                                         | Monitoring and MCP generation can fail or probe the wrong service during cutover.                                                                 | Route both endpoints to the Auth Gateway on `requiems.xyz`, update consumers, and test exact/bare/trailing-slash behavior before retiring the old host. |
| MEDIUM   | Certificate ownership and renewal for `requiemsapi.com` are absent from repository configuration.                                                                                                  | Existing host settings in `infra/kamal/deploy.dashboard.yml:15-18` and `infra/caddy/Caddyfile:5-23`; no new-domain DNS/certificate files                                                       | New product traffic can fail TLS or certificate renewal.                                                                                          | Name Cloudflare/registrar/origin certificate owners, issuance/renewal path, CAA/DNSSEC state, and deletion rollback.                                    |
| LOW      | Many repository docs, fixtures, examples, and performance reports record the old host.                                                                                                             | `readme.md`; `docs/apis`; `tests/reports`; `tests/integration`; `apps/dashboard` views                                                                                                         | Internal docs and test defaults can point at retired endpoints, but no customer compatibility impact exists.                                      | Update active docs/tests; leave clearly historical reports unchanged.                                                                                   |

## 10. Multi-Agent Review History

Lead research passes before independent review:

- Repository cross-check established the Worker Custom Domain versus apex Rails
  split, exact `/v1` versus `/v1/*` route requirement, `API_BASE_URL`'s lack of
  runtime consumption, generated/SEO surfaces, and Kamal-versus-Caddy ingress
  ambiguity.
- Lead adversarial pass rechecked Rails routes, Devise/session behavior,
  mailers, webhook tests, API proxy, MCP/OpenAPI generation, Worker CORS, Go
  secret middleware, CI/CD, and SEO files before accepting reviewer reports.

The first review rounds were conducted against the original brief, before the
owner confirmed the clean-break scope. Any round-1 wording about preserving
legacy hosts, old-page redirects, customer telemetry, or SEO is historical
review context only; the scope amendment below supersedes it and is not a
current implementation requirement.

Review round 1 (independent subagents; no files edited by reviewers):

- Infrastructure reviewer: confirmed the redirect mechanism, old-apex Rails
  grace ingress, Worker deployment gap, absent isolated preview assumption,
  route ownership, health-check, certificate, cache, and rollback requirements.
- Application reviewer: confirmed the complete old-apex POST inventory, session
  boundary, raw-body webhook handling, internal playground contract,
  API_BASE_URL non-consumption, hardcoded mail links, request-derived payment
  return URLs, MCP layers, and private-tenant wildcard.
- API compatibility reviewer: confirmed permanent no-redirect dual-host API
  support; found illustrative API path drift, Bearer/header contradictions,
  binary and proxy parity gaps, usage-header exposure, stale error URLs, and
  D1's lack of hostname attribution.
- Security reviewer: confirmed old-apex POST redirect risk, absent production
  host policy/trusted-proxy evidence, public-origin/secret boundary requiring
  external lockdown verification, webhook replay/idempotency gap, wildcard CORS,
  absent explicit cache policy, unverified HSTS claim, error disclosure, and
  forwarded-IP spoofing tests.
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
- Under the original audit premise, retained legacy API support as the safest
  default because external-client behavior was unknown. The later owner scope
  decision explicitly supersedes that recommendation and authorizes retirement
  without compatibility support.

Adversarial review (round 1 extension):

- Found and verified a cross-document architecture conflict: the separate
  architecture audit recommends retiring the Worker/KV/D1/API-Management layer,
  while this plan initially recommended permanent Worker dual-host support.
  Added a blocking architecture decision gate.
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
- Security convergence: found cookie forwarding across the API boundary, Rails
  forwarded-IP and error-disclosure surfaces, broader webhook side effects,
  explicit cookie-attribute verification, and the `/locale` referer redirect
  sink. The lead verified each cited file and added controls/tests.
- Compatibility convergence: found the admin portion of the POST inventory was
  incomplete, architecture sequencing was still too open-ended, D1 adoption
  measurement needed retention/query/reconciliation details, OpenAPI branch
  selection was incomplete, auth-header drift was broader than README/helper
  text, and external SDK release ownership was missing. The lead added the
  complete route inventory requirement, two executable architecture branches,
  durable metric requirements, localized/auth completeness checks, and SDK
  ownership gates.

Lead verification of round 2:

- Verified the admin resource routes at
  `apps/dashboard/config/routes.rb:126-149`, the Worker cookie behavior, Rails
  proxy/error/IP behavior, architecture audit conflict, Wrangler identifiers,
  and production immutable cache header.
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
  provider, monitoring, and architecture-branch facts are explicitly listed as
  external human gates, not assumed.

Scope amendment after review:

- The project owner confirmed that there are no users, customers, or meaningful
  public positioning to preserve. Accordingly, legacy API compatibility,
  old-page redirects, customer sunset telemetry, and SEO migration work were
  removed from the required deliverable.
- The lead reclassified the corresponding compatibility/SEO concerns as
  out-of-scope, retained only active operational dependencies (payment
  callbacks, MCP, CI, monitoring, email, and internal services), and rewrote the
  target as a clean break with explicit 404/410 retirement behavior.
- This scope amendment does not waive security, authentication, DNS/TLS,
  ingress, deployment, rollback, or new-system correctness gates.

## 11. Final Confidence Assessment

Final assessment after repository research, two multi-agent review rounds, lead
verification, and a final adversarial pass.

- Independent review rounds completed: 2, plus 1 final adversarial pass.
- Reviewer roles completed: infrastructure, application, API compatibility,
  security, developer experience, SEO/public web, and adversarial.
- Verified HIGH findings remaining: 12; these are pre-implementation gates, not
  authorization to proceed without resolution/acceptance.
- Verified MEDIUM findings remaining: 18.
- Verified LOW findings remaining: 1.
- Review convergence: no new relevant HIGH/MEDIUM repository findings remained
  after the final adversarial changes; the legacy API no-redirect strategy was
  replaced by the owner-approved clean-break strategy.
- Unresolved decisions: Worker-retention versus Go-cutover architecture branch;
  exact `/dashboard` contract; full old-apex POST disposition; event-level
  webhook idempotency; cookie forwarding policy; auth-header contract; SDK
  publication decision; and API path-alias scope. The `/healthz` and
  `/openapi.json` owner/transition contract is fixed by Section 3.
- Human verification required: Cloudflare zone/routes/Rules/Cache/WAF state;
  DNS/registrar/DNSSEC/CAA; certificates and renewal; active VPS ingress; origin
  firewall; Worker/KV/D1 staging resources; Lemon Squeezy callback
  configuration; email provider/DNS; monitoring and retention; external
  SDK/skills repositories; and any active internal consumers.

## 12. Final Implementation Checklist

### Contract and discovery

- [x] Complete independent infrastructure, application, API-compatibility,
      security, developer-experience, SEO, and adversarial review rounds; verify
      findings against repository evidence.
- [ ] Resolve the architecture branch: Worker-retention or Go-cutover. Record
      ownership of `/v1`, auth, quotas, usage, OpenAPI, MCP, and retirement.
- [ ] Approve exact product, API, dashboard, webhook, health, OpenAPI, and
      non-API 404/410 contracts; explicitly record that old hosts are not
      redirected or compatibility-supported.
- [ ] Generate and classify the complete Rails method/path inventory, including
      every POST/PATCH/PUT/DELETE old-apex route.
- [ ] Verify live DNS, Cloudflare routes/custom domains/rules, certificates,
      active VPS ingress, and origin health checks.
- [ ] Export webhook/monitoring baselines and any active internal-consumer
      inventory; customer traffic/sunset measurement is not required.
- [ ] Back up relevant Worker route/binding settings and database/usage data.
- [ ] Name owners and acceptance artifacts for Cloudflare routes, retirement,
      cache/WAF rules, certificates, Worker deploy, DNS, and rollback.
- [ ] Provision isolated staging Worker/KV/D1/secrets/origin/hostname before any
      canary; prohibit the current production-ID staging command.

### Product host

- [ ] Provision and certificate-verify `requiemsapi.com`.
- [ ] Deploy Rails to the new host and validate it before retiring the old
      product ingress.
- [ ] Decide `/dashboard` versus localized canonical and test all auth/account
      paths.
- [ ] Update public URL generation, mail links, analytics, and active human-
      facing links. Defer SEO-only assets unless needed for rendering.
- [ ] Keep mail transport DNS unchanged unless separately approved.

### API host/path

- [ ] Add exact Cloudflare route(s) for `requiems.xyz/v1` and
      `requiems.xyz/v1/*` to the selected API owner, with precedence verified.
- [ ] Implement the canonical non-redirecting `/healthz` and `/openapi.json`
      surfaces on the Auth Gateway at the new API host and update all internal
      consumers before retiring the old host.
- [ ] Prove method/body/query/header/path/response correctness on the new host.
- [ ] Before deleting `api.requiems.xyz`, verify old API paths never redirect;
      then remove its Custom Domain/DNS and verify no active DNS, certificate,
      or route remains. Any still-resolving propagation path must return
      404/410.
- [ ] Preserve KV/D1, usage, quota, rate-limit, and backend-secret behavior.

### Webhooks and external systems

- [ ] Update Lemon Squeezy callback to the product host.
- [ ] Disable the old Lemon Squeezy callback after the new callback is verified;
      no old-host POST grace period is required.
- [ ] Verify email confirmation/reset/account links and sender deliverability.
- [ ] Update MCP production base URL and regenerate its spec/tools.
- [ ] Update external monitors, active support/internal docs, SDK/agent
      configuration where operationally used, CI secrets, and provider
      callbacks.

### Retirement, documentation, and optional SEO cleanup

- [ ] Configure API route precedence before the non-API apex 404/410 handler.
- [ ] Verify old public/API paths are retired without 301/302/307/308
      compatibility redirects.
- [ ] Retire every old-apex state-changing route, including admin-generated
      routes and `/locale`; fuzz any retained Referer-derived redirect on the
      new product host and encoded redirect targets.
- [ ] Update API docs to `https://requiems.xyz/v1` and product docs to
      `https://requiemsapi.com` by semantic category.
- [ ] Regenerate OpenAPI, MCP, and active API docs; update LLM/SDK artifacts
      only when they are active operational consumers.
- [ ] Defer Search Console, sitemap submission, canonical/indexing, and other
      SEO migration work until product positioning is in scope.

### Validation and rollback

- [ ] Run the new-host route/API correctness matrix with real test keys and
      explicit old-surface 404/410 assertions.
- [ ] Test cookies, CSRF, login, logout, confirmation, reset, dashboard, and
      webhook signatures.
- [ ] Verify cookie filtering to Go, CORS/auth-header completeness, forwarded-IP
      provenance, generic production error bodies, HSTS, cache headers, and
      direct-origin rejection.
- [ ] Add durable hostname-labeled edge telemetry with retention, query owner,
      and reconciliation against D1 usage totals.
- [ ] Run API integration/load/synthetic tests against the new API host only.
- [ ] Monitor Worker, Rails, Go, Redis, D1, API Management, MCP, DNS/TLS, and
      webhook metrics through at least one controlled webhook/test cycle (or the
      first real billing cycle if billing becomes active).
- [ ] Keep a tested rollback for Worker route, product host, retirement, and
      webhook provider configuration.
