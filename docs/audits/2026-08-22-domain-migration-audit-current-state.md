# Requiems Domain Migration — Current-State Audit

Status: current-state re-audit. The 2026-08-21 domain migration audit's
recommended target (`requiemsapi.com` as the public product host, apex
`requiems.xyz/v1/*` path-split for the API) was **not implemented**. A different
initiative — the "Go Auth Foundation" work (phases 0–10, 2026-08-21 through
2026-08-23) — was executed instead, and is now complete. It solved the same
underlying problem (Cloudflare Worker as a fragile, production-ID-only,
undeployed-by-CD single point of failure) by retiring the Worker application
layer rather than by moving domains.

This document does not edit or supersede
`docs/audits/2026-08-21-domain-migration-audit.md` in place — per this
repository's own convention
(`docs/plans/2026-08-22-go-auth-foundation-phase-8-9-workers-retirement-completion.md`
§9.1: "Keep historical audit and plan documents unchanged; they are migration
records, not current architecture documentation"), that file is left as the
historical record of the original brief. This file is the as-built reassessment.

Audit date: 2026-08-22

Repository: `requiems-api`, branch `main` (post-merge of `v2/requiems-api-v2`,
commit `ac6b611a`)

## 1. Executive Summary

**The domain migration to `requiemsapi.com` did not happen and is not in
progress.** No repository file, DNS record, plan, or doc references
`requiemsapi.com` outside the historical 2026-08-21 audit itself. There is no
evidence this is still a live goal — it appears to have been quietly dropped in
favor of a narrower fix.

What actually happened instead, fully completed and verified in production
(`docs/plans/2026-08-22-go-auth-foundation-phase-8-9-workers-retirement-completion.md`,
"Implementation notes", completed 2026-08-23):

1. The Cloudflare Worker application layer (`auth-gateway`, `api-management`,
   `shared`) was deleted from the repository and from the live Cloudflare
   account (KV namespace `requiems_api_cf` and D1 database `requiem-usage` both
   deleted; confirmed empty before deletion).
2. `api.requiems.xyz` **kept its hostname** but changed backend: it is now a
   plain orange-cloud DNS A record to the VPS, terminated by Caddy with
   Cloudflare Origin Pull mTLS enforced (`infra/caddy/Caddyfile:31-47`),
   forwarding directly to the Kamal-managed Go API. No Worker, KV, or D1 sits in
   the request path anymore.
3. `internal.requiems.xyz` was retired entirely — it no longer exists in Caddy,
   Kamal, or Rails config. `api.requiems.xyz` is now Go's own Kamal `proxy.host`
   (`infra/kamal/deploy.api.yml:15`).
4. `api-management.requiems.xyz` was retired entirely (Worker deleted, DNS
   absent, confirmed by post-deletion `curl` exit 6).
5. `requiems.xyz` (Rails) and `mcp.requiems.xyz` (MCP) are unchanged.
6. The API's auth boundary moved from a Cloudflare Worker doing header-based key
   lookup against KV, to Go's own `apikeyauth.go` middleware validating against
   Postgres with a Redis cache — `X-Backend-Secret` and the Worker-to-origin
   secret gate are gone.

None of the original audit's five recommended actions (move Rails to
`requiemsapi.com`, split `requiems.xyz` by path, retire `api.requiems.xyz`, apex
404/410 for non-API paths, no redirects) were done. The apex (`requiems.xyz`)
still serves Rails for everything, exactly as before the audit. The clean-break
policy in the old audit's Section 3 is **moot** — there is no apex path-split to
test, because the API never moved to the apex.

Overall risk of the _current, as-shipped_ topology is LOW-MEDIUM: it is a
narrower, better-tested change than the original plan (single Worker retirement
with explicit rollback gates, owner-confirmed destructive-action sign-off, and a
documented production smoke-test matrix), but it leaves a few of the original
audit's non-Worker-specific findings genuinely unresolved (see §4), and
introduces one new one (no CORS headers on the direct Go API path).

## 2. Current Architecture (as-built, verified against source)

### 2.1 Domain and service map

| Host                          | Evidence                                                                                               | Current responsibility                                                                                                                                                    | Confidence |
| ----------------------------- | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------- |
| `requiems.xyz`                | `infra/kamal/deploy.dashboard.yml:15`; `infra/caddy/Caddyfile:9-24`                                    | Rails: public site, locale routes, Devise auth, dashboard, admin, `/api/proxy`, tool demos, webhooks, docs                                                                | High       |
| `api.requiems.xyz`            | `infra/kamal/deploy.api.yml:15`; `infra/caddy/Caddyfile:31-47`; `apps/api/app/app.go`                  | Public API — Caddy (Cloudflare AOP mTLS) → Kamal → Go directly. Go owns API-key auth, rate limit, quota, usage. No Worker, no KV, no D1.                                  | High       |
| `mcp.requiems.xyz`            | `infra/kamal/deploy.mcp.yml:15`; `infra/caddy/Caddyfile:27-30`                                         | MCP HTTP server, `REQUIEMS_BASE_URL: https://api.requiems.xyz`                                                                                                            | High       |
| `internal.requiems.xyz`       | Absent from `infra/caddy/Caddyfile`, `infra/kamal/*.yml`, and `apps/dashboard` config                  | **Retired.** Rails' `INTERNAL_API_URL` is now the private Docker network alias `http://requiems-api:8080` (`infra/kamal/deploy.dashboard.yml:31`), not a public hostname. | High       |
| `api-management.requiems.xyz` | `docs/core/infrastructure.md:20`; deletion evidence in the phase 8/9 plan                              | **Retired.** Worker deleted, DNS confirmed absent.                                                                                                                        | High       |
| `requiemsapi.com`             | No repository, DNS, or plan reference found anywhere (checked `docs/`, `infra/`, `apps/`, `agents.md`) | **Never provisioned. Not a current goal per any tracked document.**                                                                                                       | High       |

### 2.2 Current request flow

```text
Public API:
Client -> Cloudflare (DNS/WAF/DDoS/TLS, AOP mTLS enforced)
       -> Caddy (api.requiems.xyz, client_auth require_and_verify against
          Cloudflare Origin Pull CA)
       -> Kamal-managed Go API (:8080)
       -> apikeyauth (Postgres + Redis cache) -> ratelimit (Redis) -> usage/quota
          (Postgres + Redis) -> /v1/* handler
       -> PostgreSQL / Redis

Rails public/product:
Client -> Cloudflare -> Caddy (requiems.xyz) -> Kamal-managed Rails
       -> locale scope, Devise, dashboard, admin, /api/proxy, webhooks

Rails Playground (server-side, not public API path):
Rails -> private Docker network requiems-api:8080 (no hostname, no secret
       header) -> Go, using a real requiems-api-key
```

Source: `apps/api/app/app.go:52-73` (chi router: `RequestID` → `RequestLogger` →
protected group with `apiKeyAuth.Middleware()` → `rateLimiter.Middleware()` →
`usageQuota.Middleware()` → `/v1` routes);
`apps/api/platform/middleware/apikeyauth.go:18`
(`clientAuthHeader = "requiems-api-key"`); `docs/core/architecture.md:1-24`
(current-state doc, matches source).

### 2.3 What replaced the Worker's responsibilities

| Old Worker responsibility                          | Current owner                                                                                                                                                                                                                                            |
| -------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `requiems-api-key` validation against KV           | Go `apikeyauth.go` against Postgres, Redis-cached 30s (`apps/api/app/app.go:23`)                                                                                                                                                                         |
| Rate limiting                                      | Go `ratelimit.go` against Redis, plan-aware via `plancache.go`                                                                                                                                                                                           |
| Quota/usage recording                              | Go `usage.go` against Postgres + Redis, plan-aware                                                                                                                                                                                                       |
| `X-Backend-Secret` origin gate                     | Removed. Caddy's Cloudflare Origin Pull mTLS (`tls { client_auth }`) is now the only edge-bypass defense; Go itself has no secondary secret check on `/v1`.                                                                                              |
| CORS (`Access-Control-Allow-Origin: *`, preflight) | **Nothing.** `grep -rn "Access-Control\|cors(" apps/api --include='*.go'` returns no matches outside an unrelated string literal. See §4 New Finding.                                                                                                    |
| `/healthz`, `/openapi.json` on the API host        | `/healthz` only — Go registers `router.Get("/healthz", ...)` (`apps/api/app/app.go:56`); no `/openapi.json` route exists in `apps/api`. Spec generation now lives entirely in `apps/dashboard`/`apps/mcp` tooling, not served by the API process itself. |
| Forwarded-IP handling                              | Go `platform/httpx/trustedproxy.go` and Rails `app/lib/trusted_proxy.rb` (new — mirrors each other by design, static Cloudflare CIDR snapshot).                                                                                                          |

## 3. Disposition of the original 2026-08-21 findings

Legend: **RESOLVED** = moot because the Worker/KV/D1 layer that caused it is
gone; **STILL OPEN** = applies unchanged to the current topology; **NEW** =
introduced by the current topology; **N/A** = only applied to the
`requiemsapi.com` migration itself, which never started.

| Original finding (severity)                                                           | Disposition                                                                                                                                                                                                                                                                                                                                                                                                              |
| ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| Host-wide Worker custom domain vs. apex split (HIGH)                                  | N/A — apex split abandoned                                                                                                                                                                                                                                                                                                                                                                                               |
| Lemon Squeezy webhook bound to old apex, must switch before retiring host (HIGH)      | N/A — apex/host never changed, webhook still on `requiems.xyz`, nothing to switch                                                                                                                                                                                                                                                                                                                                        |
| Two competing ingress descriptions, Kamal vs. Caddy (HIGH)                            | RESOLVED — both are now the _same_ verified live path: Caddy in front of Kamal, on one VPS (`infra/kamal/deploy.api.yml` accessories now run Caddy itself)                                                                                                                                                                                                                                                               |
| Go `/v1` protected only by backend secret, no host/path logic (HIGH)                  | RESOLVED differently — Go now does its own API-key auth; there is no backend-secret layer left to bypass. Direct-origin access is blocked by Caddy AOP mTLS, verified live (`certificate_required` TLS failure recorded in phase 8/9 notes).                                                                                                                                                                             |
| No production `config.hosts` allowlist, absolute URLs partly request-derived (HIGH)   | **STILL OPEN** — `grep -n "config.hosts" apps/dashboard/config/**/*.rb` returns nothing. No allowlist was added.                                                                                                                                                                                                                                                                                                         |
| Static shared secret + unproven origin firewall (HIGH)                                | RESOLVED — static shared secret removed from the normal path; origin is now behind Caddy mTLS, verified live with an actual TLS `certificate_required` rejection test.                                                                                                                                                                                                                                                   |
| Webhook HMAC has no event-ID replay ledger (HIGH)                                     | **STILL OPEN** — `apps/dashboard/app/controllers/webhooks/lemonsqueezy_controller.rb` still has no idempotency/event-ID check; only signature verification. Untouched by the v2 changes (its 27-line diff was plan-name/billing-cycle logic, not idempotency).                                                                                                                                                           |
| Worker forwards `Cookie` header to Go origin (HIGH)                                   | RESOLVED — no Worker, no forwarding layer exists at all                                                                                                                                                                                                                                                                                                                                                                  |
| CD workflow never deployed the Worker; no isolated preview (HIGH)                     | RESOLVED — Worker deleted, so there's nothing to fail to deploy; Go/Rails/MCP are all in the existing Kamal CD path                                                                                                                                                                                                                                                                                                      |
| Worker-retention vs. Go-cutover architecture branch open (MEDIUM, cross-doc conflict) | RESOLVED — Go-cutover was chosen and completed                                                                                                                                                                                                                                                                                                                                                                           |
| Documented Worker "staging" used production KV/D1 IDs (HIGH)                          | RESOLVED — no Worker, no KV/D1, nothing to canary unsafely                                                                                                                                                                                                                                                                                                                                                               |
| No named authoritative Cloudflare zone/route owner (HIGH)                             | RESOLVED for the current topology — phase 8/9 notes name the exact zone (`9dceb9681679d346c9afff8d5e92cf2d`) and confirm zero Worker routes/custom domains exist                                                                                                                                                                                                                                                         |
| No versioned rollback procedure (HIGH)                                                | Partially resolved — the Worker retirement had an explicit rollback record (commit `6ca862f4`); Kamal itself still has no documented route/version rollback runbook beyond "redeploy a prior image"                                                                                                                                                                                                                      |
| Dashboard URL locale-scoped, target wants `/dashboard` (MEDIUM)                       | N/A — no `requiemsapi.com`/`/dashboard` contract was ever adopted; `apps/dashboard/config/routes.rb` locale scoping is unchanged                                                                                                                                                                                                                                                                                         |
| Wildcard CORS behavior needs preservation on new route (MEDIUM)                       | Superseded by a worse outcome — see §4 New Finding: CORS was **dropped**, not preserved                                                                                                                                                                                                                                                                                                                                  |
| `/healthz`/`/openapi.json` ownership unresolved (MEDIUM)                              | Partially resolved differently — `/healthz` is now on `api.requiems.xyz` directly from Go; `/openapi.json` is **not** served by any host at all anymore (see §4)                                                                                                                                                                                                                                                         |
| `API_BASE_URL` configured, no runtime consumer (MEDIUM)                               | Still true but now used correctly at the infra layer — `infra/kamal/deploy.dashboard.yml:33` sets `API_BASE_URL: https://api.requiems.xyz` (unchanged value, host never moved)                                                                                                                                                                                                                                           |
| Target API examples don't match real Go routes (MEDIUM)                               | N/A — path-alias question never arose since no path migration happened                                                                                                                                                                                                                                                                                                                                                   |
| Bearer-auth documentation contradicts header-only auth (MEDIUM)                       | **STILL OPEN, and now worse** — `readme.md:103` and `apps/dashboard/app/helpers/application_helper.rb:114` still show `Authorization: Bearer YOUR_API_KEY`; `apps/dashboard/config/locales/{en,es,fr}/comparisons.*.yml` also independently show a Bearer example. Actual contract is `requiems-api-key` only (`apps/api/platform/middleware/apikeyauth.go:18`), and Go now has no CORS/preflight support for it at all. |
| No explicit Worker cache policy for authenticated responses (MEDIUM)                  | RESOLVED — no Worker/cache layer between Cloudflare and Go for `/v1`                                                                                                                                                                                                                                                                                                                                                     |
| HSTS claimed but not proven (MEDIUM)                                                  | **STILL OPEN, unverified** — no repository ingress config (`infra/caddy/Caddyfile`, Kamal proxy) emits `Strict-Transport-Security` explicitly; still needs a live header check                                                                                                                                                                                                                                           |
| Worker error handler leaks `err.message` (MEDIUM)                                     | RESOLVED — Worker deleted; Go's own error handling was not audited here and should be checked separately                                                                                                                                                                                                                                                                                                                 |
| Forwarded-IP spoofing via untrusted header trust (MEDIUM)                             | RESOLVED — replaced with `platform/httpx/trustedproxy.go` (Go) and `app/lib/trusted_proxy.rb` (Rails), both validating the immediate peer against a Cloudflare CIDR allowlist before trusting `X-Forwarded-For`                                                                                                                                                                                                          |
| MCP fetch/runtime/snapshot URL layers inconsistent (MEDIUM)                           | Not reverified in this pass — `apps/mcp` still points at `https://api.requiems.xyz` (`infra/kamal/deploy.mcp.yml`), consistent with the unchanged hostname; internal layer consistency wasn't re-audited here                                                                                                                                                                                                            |
| OpenAPI/MCP/Go path drift (`/v1/entertainment/advice` vs `/v1/text/advice`) (MEDIUM)  | Not reverified in this pass                                                                                                                                                                                                                                                                                                                                                                                              |
| Mailer/checkout links have independent old-host contracts (MEDIUM)                    | N/A — mail/checkout host is still `requiems.xyz`/`mail.requiems.xyz`, unchanged (`infra/kamal/deploy.dashboard.yml:46-52`)                                                                                                                                                                                                                                                                                               |
| Private-deployment tenant URLs use `*.requiems.xyz` wildcard (MEDIUM)                 | N/A — apex never moved, wildcard namespace untouched                                                                                                                                                                                                                                                                                                                                                                     |
| Certificate ownership for `requiemsapi.com` absent (MEDIUM)                           | N/A — domain never provisioned                                                                                                                                                                                                                                                                                                                                                                                           |
| Stale domains (`requiems-api.xyz`, `api.requiems.dev`) in docs (LOW)                  | **PARTIALLY RESOLVED** — `requiems-api.xyz` no longer appears anywhere in the repo. `api.requiems.dev` is **still present**, in 30 lines across `apps/dashboard/config/locales/{en,es,fr}/comparisons.*.yml` (competitor-comparison marketing copy, including a Bearer-auth example).                                                                                                                                    |

## 4. New findings introduced by the current architecture

| Severity | Finding                                                                                                                                                              | Evidence                                                                                        | Impact                                                                                                                                                                                                                                                                | Recommended action                                                                                                                                                                                                                  |
| -------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| MEDIUM   | The direct Go API path emits no CORS headers at all (no `Access-Control-Allow-Origin`, no preflight handling). The Worker used to add wildcard CORS unconditionally. | `grep -rn "Access-Control\|cors(" apps/api --include='*.go'` — no matches in non-test code      | Any third-party developer calling `api.requiems.xyz` directly from browser JS (as the README/docs imply is supported) will fail on the CORS preflight, with no clear error surfaced to them. Server-to-server callers (Rails `/api/proxy`, MCP, curl) are unaffected. | Decide deliberately: either add CORS middleware to Go for the intentionally-public `/v1/*` surface, or update all API docs/examples to state server-side-only browser usage is unsupported. Do not leave this as an accidental gap. |
| LOW      | `/openapi.json` is no longer served by any host. The Worker used to serve it on the old API custom domain; Go's router only registers `/healthz`.                    | `apps/api/app/app.go:56` (only `/healthz` registered); no `/openapi.json` handler in `apps/api` | MCP spec-fetch and any external OpenAPI consumer pointed at `api.requiems.xyz/openapi.json` will get a 404 from Go rather than a spec. Not yet verified whether MCP now fetches the spec from a different source (e.g., a static file) or is currently broken.        | Confirm `apps/mcp/scripts/fetch-spec.ts`'s actual current source and either restore a served spec endpoint or document the new fetch path.                                                                                          |

## 5. Open question for the owner

`requiemsapi.com` is absent from every tracked document except the original
audit. Two readings are possible:

1. The domain migration was deliberately shelved in favor of the narrower
   Worker-retirement fix, and should be formally dropped from any backlog
   referencing it (the 2026-08-21 audit file itself should stay as historical
   record, per the repo's own convention, but nothing should still point to it
   as a live plan).
2. It's still wanted, just not yet scheduled — in which case the 2026-08-21
   audit's Section 3 target and Section 5 execution plan are still the right
   starting point, but should be re-scoped now that the Worker/KV/D1 layer no
   longer exists (several of its steps, e.g. "retire the old Worker Custom
   Domain," are already done for unrelated reasons).

This audit does not assume either answer. Recommend the owner explicitly confirm
intent before further planning work references `requiemsapi.com`.

## 6. Recommended next steps (independent of the `requiemsapi.com` decision)

1. Fix or intentionally accept the CORS gap (§4) — this is a live behavior
   regression from the pre-Worker-retirement state, not a pre-existing audit
   finding.
2. Resolve the Bearer-vs-header-key documentation contradiction across
   `readme.md`, `apps/dashboard/app/helpers/application_helper.rb`, and the
   three `comparisons.*.yml` locale files — one coherent pass, not per-file.
3. Replace the remaining `api.requiems.dev` references in
   `apps/dashboard/config/locales/{en,es,fr}/comparisons.*.yml` with the real
   host.
4. Add a Rails `config.hosts` allowlist for production — still absent, and was
   flagged HIGH in the original audit before any of this work started.
5. Add webhook event-ID idempotency to `Webhooks::LemonsqueezyController` —
   still absent, still HIGH-equivalent risk (duplicate delivery can repeat
   subscription side effects), and untouched by v2.
6. Verify `Strict-Transport-Security` is actually emitted end-to-end through
   Cloudflare → Caddy, with a live header check.
7. Confirm the current MCP OpenAPI fetch/runtime source now that Go no longer
   serves `/openapi.json`, and re-run the `(method, path,
   operationId)` drift
   check the original audit specified for `advice`.
