# Requiems API — Architecture & Codebase Audit

**Date:** 2026-08-21 **Scope:** Full monorepo — `apps/api` (Go),
`apps/dashboard` (Rails), `apps/workers/{auth-gateway,api-management,shared}`
(Cloudflare Workers), PostgreSQL, Redis, Cloudflare KV/D1, CI/CD, infra.
**Method:** Code-level tracing (not doc-trusting) of every component listed
above, cross-referenced against `docs/core/*.md`. Findings below cite
`file:line`. No code was changed as part of this audit.

---

## Executive Summary

Requiems API runs four separate stateful stores to serve one thing: an
authenticated, rate-limited, metered API request. Cloudflare KV holds the API
key and a per-minute rate-limit counter. Cloudflare D1 holds the usage ledger.
PostgreSQL (via Rails) holds the durable API-key/subscription/usage record.
Redis (via Go) holds an unrelated set of response caches plus the state for one
public product feature (a hit-counter API). None of these four stores is aware
of the others in real time — the two Workers coordinate over HTTP, and Rails
pulls D1 into Postgres on a 3-minute cron. The Go backend, which does all the
actual business logic, sits behind a single shared secret and has **no concept
of an API key, a user, a plan, or a rate limit** — every one of those concerns
lives in the edge Worker in front of it (~700 lines of glue logic in
auth-gateway alone; more across both Workers combined).

This split was a reasonable choice when the goal was "sub-10ms global auth
checks," but the audit found the edge tier is not actually buying that guarantee
cleanly: the KV-based rate limiter is a **non-atomic get-then-put**
(`apps/workers/auth-gateway/src/rate-limit.ts:38,51`) that races under
concurrent requests from the same key, the quota-check path
(`middleware/api-key-auth.ts:81` → `requests.ts`) has no timeout or fallback if
KV/D1 are slow, and the proxy's hot path is essentially untested end-to-end (no
test exercises a full successful proxied request). Meanwhile the Go backend
already imports `pgx` and `go-redis` and already runs a small Redis→Postgres
batched-counter-sync worker
(`apps/api/services/technology/counter/sync_worker.go`) whose crash-safe,
idempotent flush mechanism is a genuinely solid foundation to build a
Redis-based rate-limiter/usage-tracker on. **This is a foundation, not a
finished prototype** — the existing pattern proves the sync mechanism works, not
that rate limiting or multi-dimensional billing already work; see the "Go
Backend" component audit below for exactly what it does and does not
demonstrate.

**Biggest problems found, independent of the Workers question:** a
duplicate-HTTP-call bug that double-syncs every plan change to Cloudflare
(`apps/dashboard/app/models/subscription.rb:26` +
`apps/dashboard/app/controllers/webhooks/lemonsqueezy_controller.rb:108` et
al.), a revenue-reporting bug where MRR is always computed as monthly pricing
because the billing-cycle column is never written
(`apps/dashboard/app/models/subscription.rb` `plan` field, consumed by
`analytics_revenue_service.rb:44-82`), a possible Postgres
migration-tracking-table collision between Rails and Go that needs live-database
verification, a Go service with no graceful shutdown and no structured logging,
and an auth-gateway proxy path with materially weaker test coverage than the
rest of the system it protects.

**Biggest opportunity:** collapsing four operational-state stores (KV, D1,
Postgres via 3-min cron, Redis-for-unrelated-things) into one (Redis for
hot-path state, Postgres for durable state), building on a sync mechanism the Go
codebase has already proven. This removes an entire deployment target (two
Workers, unautomated in CI), an entire edge database (D1, whose
SQLite-per-database write model does not scale to the 10k–100k RPS target the
product hypothesis mentions), and roughly 3,000 lines of Rails/TS glue code
whose only job is keeping those stores in sync (auth-gateway ~700 lines,
api-management ~1,400, `apps/workers/shared` ~600, plus the Rails-side
`D1SyncService`/`Cloudflare::ApiManagementService`/`SyncD1UsageJob` and the
model sync callbacks, ~360 lines).

**Recommended target architecture:** Cloudflare stays as DNS/WAF/DDoS/TLS only
(no Worker logic). The Go backend gains its own API-key authentication, atomic
Redis-based rate limiting, and Redis-based usage counting flushed to Postgres,
building on the batched-sync pattern already in production (with real extensions
— see the Go Backend and Usage/Billing sections below). Rails keeps owning
`users`/`api_keys`/`subscriptions` as the durable source of truth and generates
keys directly — a test-only code path for this already exists
(`apps/dashboard/app/models/api_key.rb:60-68`) but is **not** ready to promote
as-is: it currently produces keys in the wrong format (`rq_live_`/`rq_test_`
prefix instead of the `requiem_` prefix the live validator requires) and lacks
the collision-check-and-retry logic and DB-level uniqueness guard the
Cloudflare-backed path currently provides (see Rails component audit and
Migration Phase 4 below for what actually needs to change). `auth-gateway`,
`api-management`, `apps/workers/shared`, Cloudflare KV, and Cloudflare D1 are
retired in a staged migration (Section: Migration Plan). This is not a
hypothetical rewrite — the target stack's core building blocks (`pgx`,
`go-redis`, the dirty-set-swap sync pattern) already exist in the codebase today
— but it is extension plus some genuinely new work (rate limiting,
multi-dimensional usage tracking, key-format/uniqueness fixes), not a pure
cutover of already-finished code.

---

## Current Architecture

```text
                              INTERNET
                                 │
                 ┌───────────────┴───────────────┐
                 ▼                                ▼
        api.requiems.xyz                    requiems.xyz
        (Cloudflare Worker:                 (Rails Dashboard,
         auth-gateway)                       Caddy → Puma)
                 │                                │
        ┌────────┼────────┐                       │
        ▼        ▼        ▼                       │
       KV    (rate limit  D1                       │
    (API key   in same    (credit_usage             │
     lookup)     KV)      ledger)                   │
        │                  │  ▲                     │
        │                  │  │ GET /usage/export    │
        │                  │  │ (via api-management) │
        │                  │  └─────────────┐        │
        │                  │                │        │
        └────────┬─────────┘                │        │
                  │ X-Backend-Secret         │        │
                  ▼                          │        │
        internal.requiems.xyz (Go, Chi)      │        │
        NO per-key auth, NO rate limit,      │        │
        NO usage tracking of its own          │        │
                  │                           │        │
        ┌─────────┴─────────┐                 │        │
        ▼                   ▼                 │        │
     Redis              PostgreSQL ◄───────────┴────────┘
  (response caches:     (Go tables: reference data + `counters`
   geocode/crypto/fx,    Rails tables: users, api_keys, subscriptions,
   + `counters` sink      usage_logs [fed by 3-min Sidekiq cron pulling
   backing the public      from D1 via api-management], daily_usage_
   `counter` product API)  summaries, credit_adjustments, audit_logs …)

  Cloudflare Worker: api-management.requiems.xyz (internal-only)
    — API key CRUD (dual-writes KV + D1)
    — usage export endpoint (Rails' cron reads this)
    — analytics endpoints (no confirmed caller in Rails)
    Auth: static X-API-Management-Key shared secret
```

**Four operational-state stores for one concern (auth/rate-limit/usage):**

| Store         | Holds                                                                                                                                                                                                               | Written by                                                        | Read by                          | Propagation lag to Postgres                          |
| ------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------- | -------------------------------- | ---------------------------------------------------- |
| Cloudflare KV | `key:{apiKey}` → user/plan, `rl:m:{apiKey}:{minute}`, `quota:{userId}:{cycleStart}` (60s cache)                                                                                                                     | api-management (key data), auth-gateway (rate/quota cache)        | auth-gateway                     | N/A — KV state itself is never the sync source       |
| Cloudflare D1 | `credit_usage` (every request), `api_keys` (audit mirror)                                                                                                                                                           | auth-gateway (usage rows), api-management (key CRUD)              | api-management (`/usage/export`) | up to 3 minutes (Sidekiq cron)                       |
| PostgreSQL    | `usage_logs`, `daily_usage_summaries`, `api_keys`, `subscriptions`, `users`, …                                                                                                                                      | Rails (`SyncD1UsageJob`, `AggregateDailyUsageJob`, direct writes) | Rails dashboard/admin            | — (destination)                                      |
| Redis         | Go response caches (`geocode:*`, `crypto:*`, `exchange:*`), `counter:*` (backs the public `counter` hit-counter API, see `docs/apis/technology/counter.md`), Rails `Rack::Attack` throttle counters, Sidekiq queues | Go, Rails                                                         | Go, Rails                        | `counter:*` flushed to Postgres `counters` every 60s |

No store above is authoritative for more than its own slice, and nothing
reconciles them faster than the 3-minute D1→Postgres cron. A revoked key
propagates from Rails → api-management → KV delete typically within one HTTP
round trip (fast), but a plan/quota change synced only into KV can silently
diverge from Postgres if that HTTP call fails — both
`ApiKey#remove_from_cloudflare` (`apps/dashboard/app/models/api_key.rb:91-97`,
called from `sync_revocation_to_cloudflare` at `:99-103`) and
`Subscription#sync_to_cloudflare`
(`apps/dashboard/app/models/subscription.rb:30-35`) swallow errors and only log;
there is no reconciliation job that would catch a Postgres/KV divergence.

---

## Proposed Architecture

```text
         INTERNET
            │
┌───────────┴───────────┐
│  Cloudflare (DNS/WAF/  │
│  DDoS/TLS proxy only — │
│  no Worker logic)      │
└───────────┬───────────┘
            ▼
┌─────────────────────────┐
│        Go API           │
│  API-key auth (Redis-   │
│   cached, Postgres-     │
│   backed)               │
│  Rate limiting (Redis   │
│   atomic incr+expire,   │
│   Lua-scripted — see    │
│   Rate Limiting below)  │
│  Usage counting (Redis  │
│   dirty-set-swap →      │
│   Postgres, existing    │
│   pattern, generalized) │
│  Business logic (220+   │
│   endpoints, unchanged) │
└──────┬──────────┬───────┘
       ▼          ▼
    Redis     PostgreSQL
       ▲          ▲
       │          │
   (shared)   ┌───┴───┐
       │      │ Rails │
       └──────┤ (owns  │
              │ users/ │
              │ keys/  │
              │ subs)  │
              └────────┘
```

Key differences from today: one hot-path store (Redis) instead of two (KV + D1),
zero edge-to-origin HTTP hops for auth. Today's synchronous path is client →
Worker → KV (key lookup) → KV (rate limit) → KV (quota cache, with a D1 read
only on a quota-cache miss) → Go → Postgres/Redis, followed by an asynchronous
D1 usage write dispatched via `waitUntil` _after_ the response has already gone
back to the client (`routes/proxy.ts`) — D1 is not on the synchronous critical
path the way an earlier draft of this diagram implied. Proposed: client → Go →
Redis, one hop, no separate edge tier at all. Rails talks to Postgres directly
for key CRUD instead of round-tripping to a Worker. Cloudflare is retained
purely as network infrastructure — its removal is explicitly **not** recommended
(see the DDoS/abuse-posture discussion under Risks below).

---

## Component-by-Component Audit

### Cloudflare `auth-gateway`

**Verdict on Question 1 (does it justify its complexity): No, not as currently
built.**

Full request chain per authenticated call
(`apps/workers/auth-gateway/src/middleware/api-key-auth.ts:34-113`,
`src/routes/proxy.ts:30-157`):

1. Format-validate the key (regex, no I/O).
2. **KV read** `key:{apiKey}` — 401 if missing.
3. **KV read** `rl:m:{apiKey}:{minute}`, then (usually) **KV write** of the
   incremented count — `rate-limit.ts:38,51-53`. This is a plain get-then-put
   with no compare-and-swap; two concurrent requests from the same key in the
   same minute can both read the same count and both write `count+1`,
   undercounting the true rate. At low concurrency this is invisible; at the
   10k–100k RPS the product hypothesis targets, it becomes a real correctness
   gap in exactly the mechanism meant to prevent abuse.
4. **KV read** `quota:{userId}:{cycleStart}` for the monthly quota cache; on a
   cache miss (roughly once every 60s per active user), a **D1 read**
   (`SELECT SUM(credits_used) …`, `requests.ts:34-40`) followed by a best-effort
   **KV write** to warm the cache.
5. Proxy to Go with `X-Backend-Secret` (`http.ts:31`), 10s timeout, strips all
   `cf-*` headers and re-adds `CF-Connecting-IP` as `X-Forwarded-For`.
6. On success, a **D1 write** to `credit_usage` (`requests.ts:76-90`) fires
   inside `waitUntil` (non-blocking) with 3x retry — and, separately, the quota
   KV cache is **read again** and **written again** to keep it warm
   (`requests.ts:98-101`). This is a second read of the same key already read in
   step 4, purely because the two call sites don't share state — a minor
   redundant round trip, not a correctness bug.

Best case — quota cache warm, the steady-state condition since the quota cache
TTL is 60s: **6 KV ops (4 reads + 2 writes) + 1 async D1 write = 7 total ops**
per request. Worst case — quota cache miss, which recurs roughly once per active
user per 60-second window (so a small, bounded fraction of overall traffic, not
the typical case): **7 KV ops (4 reads + 3 writes) + 1 D1 read + 1 D1 write = 9
total ops**. (These two figures are used consistently in the Request Lifecycle
and Performance Findings sections below — treat the best case as the realistic
per-request cost under sustained load, and the worst case as a periodic per-user
spike, not a per-request multiplier.)

None of the KV/D1 reads in the hot path are wrapped in error handling — a slow
or erroring KV/D1 call is not caught, falls through to Hono's generic `onError`,
and returns a bare 500 (`index.ts:33-36`). There is no fail-open or fail-soft
mode for edge-store unavailability; the doc's implicit promise of "auth is
ultra-fast and resilient at the edge" is not backed by the code's error
handling.

**Test coverage gap:** no test file exercises a full successful proxied request
(backend 2xx → usage recorded → response headers assembled). `quota.test.ts`
only covers the 429 path; `requests.test.ts` covers one D1 bind-argument
assertion. The header-filtering, dynamic `X-Usage-Count` multiplier logic, and
`waitUntil` usage-recording in `proxy.ts` are effectively untested beyond
low-level unit tests of their helper functions.

**Local dev is likely broken today:** the documented dev curl example and
`scripts/seed-dev.ts` both use keys shaped like `rq_free_000001`
(`docs/core/auth-gateway.md:160`, `scripts/seed-dev.ts:22-30`), but the live
validator requires `^requiem_[0-9a-zA-Z]{24}$`
(`apps/workers/shared/src/api-key-generator.ts:15-19`). A seeded dev key would
be rejected with 401 before ever reaching KV — either the docs/seed script are
stale relative to a key-format change, or nobody has run the documented dev curl
command recently. This independently explains the thin test coverage: the happy
path is hard to exercise manually right now.

**Nothing in the repo automates deploying this Worker** —
`.github/workflows/cd.yml` deploys `api`, `dashboard`, `mcp` only; Worker
deploys are `pnpm run deploy:prod` run by hand, with no CI gate on the deploy
itself (only on typecheck/vitest pre-merge).

### Cloudflare `api-management`

**Verdict on Question 2 (does it need to exist as a Worker): No, but it is
better-built than auth-gateway and its logic (not its transport) should be
preserved.**

Mostly CRUD-correct, with one inconsistency: create and revoke (`delete.ts`)
write D1 before KV (`routes/api-keys/create.ts:76-88`, `delete.ts:29-39`) — so a
D1 failure never leaves an orphaned KV entry, and audit-trail integrity is
preserved for those two paths. **`patch.ts` does the opposite** — it writes KV
first and D1 second (`patch.ts:66,84-90`), so a KV write that succeeds followed
by a failed D1 write leaves KV serving the updated plan/billing-cycle data with
no corresponding D1 audit record, exactly the failure mode the D1-first ordering
on the other two routes is meant to prevent. List never returns full key values.
Auth is a constant-time SHA-256 compare against a static shared secret
(`middleware/api-key-auth.ts:20-29`) — same shared-secret-everywhere pattern as
`X-Backend-Secret`, just for a different boundary. Test coverage here is
substantially better than auth-gateway (full CRUD lifecycle, collision handling,
validation edge cases all covered).

Its **only confirmed caller** is Rails, through two independently-implemented
Faraday clients that duplicate connection setup, retry policy, and auth-header
logic against the same base URL: `Cloudflare::ApiManagementService` (key CRUD +
plan sync) and the separately-named `D1SyncService` (usage export) —
`apps/dashboard/app/services/cloudflare/api_management_service.rb`,
`apps/dashboard/app/services/d1_sync_service.rb`. Its `/analytics/*` endpoints
have no discovered caller anywhere in Rails — built, tested, and (as far as this
audit can tell) unused surface area.

There is no reason this needs to run at the edge: nothing about API-key CRUD or
usage export is latency-critical in the way the auth hot path is — these are
admin/background operations. Moving key CRUD directly into Rails/Postgres (which
already holds the authoritative `key_hash`/`key_prefix`) and moving usage
recording directly into Go/Postgres eliminates the need for this service
entirely, along with both of the Rails-side HTTP clients that talk to it.

### `apps/workers/shared`

Shared TS package (`config.ts` plan limits/endpoint costs, types, HTTP helpers,
retry/backoff, custom Basic Auth). Its only consumers are the two Workers above;
nothing else in the monorepo imports it (Rails/Go read the same _values_ via a
hand-copied config, per the comment at `config.ts:1-9`, not this package).
Confirmed dead exports within it:
`ApiKeyManagementRequest`/`ApiKeyManagementResponse` types, `PLAN_NAMES`,
`getPlanLimits()`. Entirely retireable alongside the two Workers.

### Go Backend (`apps/api`)

Clean, consistently-layered (`router.go` → `service.go` → `transport_http.go`)
across ~220 routes in 9 domains; test LOC (23k) exceeds production LOC (20k) and
every leaf feature package has a matching test file. No circular dependencies
(`systems` is a one-way composition root over the other domains, not a cycle).
No N+1 patterns found; most batch endpoints use bounded goroutine fan-out — one
confirmed exception is `counter.IncrementBatch`
(`services/technology/counter/service.go:85-102`), which loops sequentially with
no concurrency. Low practical impact given batches are capped, but worth noting
since it's inside the exact package this audit leans on most heavily as evidence
(see below).

**But it is not yet fit to be the sole internet-facing tier**, independent of
the Workers question:

- **No graceful shutdown.** `main.go` calls `server.ListenAndServe()` with no
  `signal.NotifyContext`/`server.Shutdown()` anywhere — SIGTERM during a rolling
  deploy kills in-flight requests, the pgx pool, and the Redis client outright.
  Acceptable today because the Worker in front absorbs client-facing disruption
  during a Go redeploy; not acceptable if Go becomes the thing clients connect
  to directly.
- **No structured logging, no metrics, no tracing.** `log.Printf`/`log.Println`
  calls scattered across ~23 files, Sentry only on uncaught errors
  (`TracesSampleRate: 0.01`). The observability audit's baseline questions
  ("which endpoint is slow," "is Postgres slow") cannot be answered from what
  exists today.
- **No connection-pool tuning.** `pgxpool` and `go-redis` both run on library
  defaults (`platform/db/db.go`, `platform/reqredis/redis.go`) — fine at current
  load, needs explicit sizing before any RPS target discussion is meaningful.
- **A ready-made usage-tracking hook exists but is used in 1 of ~220
  endpoints.** `platform/httpx/handler.go:33-36` already defines a
  `UsageCounter` interface that auto-sets `X-Usage-Count` for any response
  implementing it — the exact mechanism the auth-gateway reads to apply
  per-endpoint billing multipliers. A migration would need this backfilled
  broadly, but the extension point already exists and is proven.
- **The counter-sync pattern is the template to reuse, not reinvent — but it
  proves less than earlier framing in this document implies.**
  `services/technology/counter/{redis_mutations.go,sync_worker.go,repository.go}`
  implements a genuinely production-proven mechanism: Lua-scripted atomic
  increment-and-mark-dirty, a `RENAMENX`-based dirty-set swap that's safe
  against crash-mid-cycle (idempotent because it upserts _absolute_ values, not
  deltas), and a single batched Postgres upsert every 60s. This part is real,
  running, tested code, not a design sketch — the sync/flush mechanism
  generalizes well. **What it does not prove:** the Lua scripts do a bare `INCR`
  with no `EXPIRE`/TTL at all — there is no windowing or limit-checking anywhere
  in this pattern, so on its own it does not demonstrate a working rate limiter
  (see the Rate Limiting Architecture section below, which recommends
  `INCR`+`EXPIRE`, a related but distinct and not-yet-proven mechanism). It also
  only supports a fixed `+1` per call (no `INCRBY`), whereas per-endpoint
  billing multipliers need variable increments, and its `namespace → total` data
  model is a flat aggregate — extending it to per-API-key × per-user ×
  per-endpoint × per-day usage tracking with monthly resets is real, unbuilt
  design work (see the Usage/Billing Architecture caveat below), not a
  copy-paste. `sync_worker.go`'s own code comment flags a scaling concern
  directly relevant to that higher-cardinality use case ("For very large sets
  (10K+), consider SSCAN for incremental retrieval"), and `repository.go`'s
  `UpsertBatch` builds a single unchunked multi-row `INSERT` with no batching
  fallback for larger row counts. None of this invalidates building on this
  pattern — its crash-safety and idempotency properties do generalize, and it
  remains the right foundation — but a rate limiter and a multi-dimensional
  usage-counter should be scoped as new work built _on_ this proven sync
  mechanism, not as work already done. It's also worth correcting the framing
  found elsewhere in earlier audit drafts: `counter` is not an internal-only
  mechanism — it backs a public, documented, customer-facing product API
  (`docs/apis/technology/counter.md`), which is a point in favor of its
  production-readiness as a _sync mechanism_, but means claims like "already
  load-tested implicitly by being in production" should be read as "has served
  real traffic," not as verified evidence at 10k–100k RPS — no data on its
  actual traffic volume was found in this repo.

### Rails (`apps/dashboard`)

Solid separation from Go's tables (confirmed: no raw SQL crosses the ownership
boundary in either direction; Rails explicitly excludes Go-owned tables from its
schema dumper, `config/application.rb:37-52`). Devise + Rack::Attack cover authn
and throttling for Rails' own surfaces reasonably well; the LemonSqueezy
webhook's signature verification is correctly constant-time and well-isolated
from the DB transaction.

**Concrete bugs found, independent of the Workers decision:**

- **Duplicate Cloudflare sync on every plan change.**
  `Subscription#sync_to_cloudflare` fires via `after_create`/`after_update`
  callbacks (`subscription.rb:25-26`) _and_ the LemonSqueezy webhook controller
  explicitly calls the same sync again right after
  (`webhooks/lemonsqueezy_controller.rb:108,160,186,210`) — every real plan
  change today does two full passes over the user's active keys, each issuing
  one HTTP `PATCH` per key to api-management. (Note:
  `handle_subscription_created` at line 108 pairs with the `after_create`
  callback at line 25, not `after_update` at line 26; the other three call sites
  pair with `after_update` at line 26 — same underlying bug either way.)
- **MRR/revenue figures are silently wrong for yearly subscribers.**
  `subscriptions.plan` (the billing-cycle column, distinct from `plan_name`) is
  never written by any code path in `app/` — confirmed by grep — yet
  `Admin::DashboardController#calculate_mrr` and `AnalyticsRevenueService` both
  branch on it, defaulting to `:monthly` pricing for every subscription
  regardless of actual billing cycle.
- **`api_keys.last_used_at`/`last_used_ip`** are displayed in the dashboard UI
  but never written by any Rails code — always shows "never used."
- **Dead Stripe-era columns** (`subscriptions.stripe_customer_id`,
  `stripe_subscription_id`, `credit_limit`) and a **dead `solid_cache_entries`
  table** (cache store is always overridden to Redis or null_store, so this
  Rails-8-default table is provisioned but never used).
- **Pundit is installed and scaffolded (`app/policies/application_policy.rb`)
  but zero concrete policies exist and no controller calls `authorize`** — all
  real authorization is ad-hoc `before_action` checks (functionally fine today,
  given router-level `admin?` gating is also present, but it's dead scaffolding
  creating a false impression of a policy layer).
- **`admin_user_id`/`resolved_by_id` columns on
  `credit_adjustments`/`audit_logs`/`abuse_reports` have no foreign key**,
  inconsistent with the structurally identical `promoted_by_id` on
  `subscriptions`, which does.
- **`Cloudflare::ApiManagementService` has zero test coverage** — the service
  responsible for every API-key create/revoke/update HTTP call has no test file
  at all, nor do any of the three "core" scheduled jobs (`SyncD1UsageJob`,
  `AggregateDailyUsageJob`, `ExpirePromotionalSubscriptionsJob`).
- **`CF-Connecting-IP` is trusted with no `trusted_proxies` configured**
  (`ApiProxyController`, matches the existing caveat already documented in
  `docs/core/rails-app.md:277-308`) — if Rails is ever reachable other than
  through Cloudflare, this header is spoofable and feeds both the IP-lookup
  endpoint and Rack::Attack's IP-based throttles.
- **`RefreshSitemapJob` is scheduled but undocumented** in both
  `docs/core/rails-app.md` and `docs/core/background-jobs.md` — a
  doc-completeness gap, not a code issue.

**The starting point for the key-generation code path Rails would need for a
migration already exists, but it is not ready to promote as-is.**
`ApiKey#generate_key_locally` (`api_key.rb:60-68`), gated behind
`Rails.env.test?`, generates and hashes a key entirely in Rails without calling
Cloudflare — via `apps/dashboard/app/services/api_key_generator.rb`. Two real
gaps versus the production Cloudflare-backed path
(`apps/workers/api-management/src/lib/generate-api-key.ts`):

1. **Wrong key format.** Rails' local generator produces keys prefixed
   `rq_live_`/`rq_test_`; the live auth-gateway validator requires
   `^requiem_[0-9a-zA-Z]{24}$`
   (`apps/workers/shared/src/api-key-generator.ts:15-19`). A key generated by
   `generate_key_locally` today would be **rejected by the Worker's own auth
   check** — this is the same root-cause format drift already flagged elsewhere
   in this audit for the broken dev-seed keys. Promoting this path to production
   requires fixing the prefix first.
2. **No collision guard.** The production path explicitly checks KV for the
   generated key and returns a `409` on collision before writing
   (`create.ts:53-64`, comment: "extremely unlikely with nanoid, but good
   practice"). `generate_key_locally` relies solely on ActiveRecord's
   `validates :key_prefix, uniqueness: true`, which is a race-prone
   check-then-insert with no database-level backing — `key_prefix`'s only index
   is a trigram index for fuzzy search (`db/schema.rb`), not an efficient
   exact-match/uniqueness-enforcing one (also flagged separately above as a P0
   item). A validation failure on collision also just errors out to the user
   today rather than retrying with a new key.

Promoting the local path to production is a good direction, but requires (a)
fixing the key-prefix format, (b) adding a proper btree index on `key_prefix`
(already a P0 item for other reasons — see Authentication Architecture above for
why `key_prefix`, not `key_hash`, is the column that needs it), and (c) adding
retry-on-collision logic — not just flipping the `Rails.env.test?` gate. See
Migration Phase 4 below.

### PostgreSQL

Go and Rails cleanly partition table ownership — verified via full grep of both
codebases' SQL, zero cross-boundary queries found in either direction. Go's
tables (`advice`, `quotes`, `words`, `bin_data`, `inflation_data`,
`iban_countries`, `commodity_price_history`, `exercises`, `swift_codes`, plus
`counters`) are all read-mostly reference data or the one counter sink — none
reference a user or API key. Rails' tables (`users`, `api_keys`,
`subscriptions`, `usage_logs`, `daily_usage_summaries`, `credit_adjustments`,
`audit_logs`, `abuse_reports`, `referrals`, `private_deployment_requests`) carry
the actual relational business model.

**Migration-tracking table: initially flagged as a collision risk, then resolved
on closer inspection.** The Rails app renames its own migration-bookkeeping
tables via `apps/dashboard/config/initializers/schema_migrations.rb` to
`rails_schema_migrations`/`rails_internal_metadata`, specifically to avoid
colliding with `golang-migrate`'s default `schema_migrations` table (Go's
migration runner invokes `migrate.New(sourceURL, dsn)` with the bare DSN and no
`x-migrations-table` override — `apps/api/platform/db/migrate.go`; the DSN
itself, confirmed via `apps/api/platform/config/config.go` and
`infra/docker/.env.example`, carries no such override either).
**`docs/core/rails-app.md:207-213`'s claim of "separate schema" is imprecise** —
it's a renamed table in the same `public` schema, not a Postgres
schema/namespace — but the practical collision is handled. No code change needed
here; only a documentation correction.

**Real, minor schema issues found regardless of the Workers question:**
`api_keys.key_prefix` — the column any Postgres-backed auth lookup would
actually query by, since bcrypt's `key_hash` can't be looked up by equality (see
Authentication Architecture above) — has only a trigram index built for fuzzy
admin search, not an efficient exact-match btree index; the FK gaps noted above
in the Rails section. (An earlier version of this audit also flagged `words` as
missing an index on its `word` column — on closer check, no code path anywhere
in `apps/api` currently queries `words` by word value at all [`Random()` does
`ORDER BY random()`, and `Define()`/`BatchDefine()` resolve from an in-memory
dataset, not this table], so there's currently no query pattern that index would
even serve. Dropped as a non-issue; verify again if a word-lookup query is ever
added.)

### Redis

**Verdict on Question 4 (can Redis replace KV for high-write operational state):
Yes, and the pattern to do it already exists in this codebase.**

Today Redis is single-instance (`redis:7-alpine`, no clustering/replication,
default RDB persistence, **no `maxmemory` configured** — unbounded growth is
possible, not just unlikely), shared across Go's response caches, Go's one
counter feature, Rails' `Rails.cache`, Rails' `Rack::Attack` throttles, and
unnamespaced Sidekiq queues. No collisions observed (key prefixes differ), but
there's no structural isolation (e.g., no separate Redis DB index per consumer).

`services/technology/counter`'s Lua-scripted atomic increment + dirty-set-swap

- batched Postgres upsert is a genuinely good design for the specific thing it
  does — sync an accumulating Redis counter to Postgres losslessly and
  idempotently: crash-safe between the increment and the next flush (worst case,
  up to 60s of increments are lost only if Redis itself crashes without a
  persisted snapshot — a real, bounded window, not an unbounded one). It is
  **not yet** a rate limiter or a multi-dimensional usage tracker — see the
  caveats in the Go Backend component audit above — but its sync/flush mechanism
  is the concrete evidence behind the recommendation to build the rate limiter
  and usage tracker on the same foundation rather than inventing a new one from
  scratch.

**What Redis is not being asked to do today, and would need to do in the
proposed design:** synchronous atomic rate-limit checks in the request hot path
(a Lua-scripted `INCR`+`EXPIRE`, O(1) — a _plain_ two-command `INCR` then
`EXPIRE` is not atomic as a pair, see Rate Limiting Architecture below — but
either way this fixes the current KV get-then-put race by construction) and a
cache of `key_prefix → {user_id, plan}` invalidated on revoke (keyed by
`key_prefix`, not a hash of the raw key — see the corrected Authentication
Architecture section above for why; mirrors the existing
`geocode`/`crypto`/`exchange` TTL-cache pattern already in `apps/api/services`).

---

## Data Ownership

| Data                                                      | Source of truth today                                                                                | Proposed source of truth                                                                                                                                                                                        |
| --------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| API key existence/plan/revocation                         | Cloudflare KV (fast path) + D1 `api_keys` (audit) + Postgres `api_keys` (Rails' view) — three copies | PostgreSQL `api_keys` (Rails-owned), Redis as a read-through cache only                                                                                                                                         |
| Rate-limit counters                                       | Cloudflare KV, non-atomic, no persistent record                                                      | Redis, Lua-scripted atomic `INCR`+`EXPIRE` (see Rate Limiting Architecture), ephemeral by design (no Postgres mirror needed — these are enforcement-only, not billing)                                          |
| Usage quota enforcement (aggregate)                       | Cloudflare D1 `credit_usage`, summed on read for the quota check                                     | Redis counters (per-user/per-key, aggregate only) → batched flush to a new aggregate table (same pattern as `counters`, not `usage_logs` itself — see the row-level caveat in Usage/Billing Architecture above) |
| Usage row-level ledger (billing/analytics detail)         | Cloudflare D1 `credit_usage` → (3 min lag) → Postgres `usage_logs`                                   | Written directly from Go to Postgres `usage_logs` per request (not via the Redis aggregate counters, which cannot reconstruct individual-request rows) — see Open Question 7                                    |
| Subscriptions/plans                                       | Postgres `subscriptions` (Rails), mirrored into KV                                                   | Postgres `subscriptions` only; Go reads it (directly or via a Redis-cached view) instead of a separate KV copy                                                                                                  |
| Business/reference data (advice, quotes, BIN data, etc.)  | Postgres, Go-owned                                                                                   | unchanged                                                                                                                                                                                                       |
| Response caches (geocode, crypto, FX)                     | Redis, Go-owned                                                                                      | unchanged                                                                                                                                                                                                       |
| Rails throttle counters (`Rack::Attack`) / Sidekiq queues | Redis, Rails-owned                                                                                   | unchanged — unaffected by this migration                                                                                                                                                                        |

---

## Request Lifecycle

**Current (authenticated API call, quota-cache-warm case):** Client → Cloudflare
edge → auth-gateway Worker → KV read (key) → KV read+write (rate limit, racy) →
KV read (quota cache) → fetch to Go over the internet (10s timeout) → Go (no
auth of its own, trusts `X-Backend-Secret`) → Postgres/Redis for business logic
→ response → Worker records usage: async D1 write + KV read/write. **7 storage
operations in the steady-state case (9 during a quota-cache-miss), 2 network
hops between edge and origin**, before business logic even runs.

**Proposed:** Client → Cloudflare (proxy only, no Worker) → Go: Redis `GET`
(key-prefix cache) → Redis Lua-scripted `INCR`+`EXPIRE` (per-minute rate limit,
atomic as a unit — see Rate Limiting Architecture below) → Redis `GET` (monthly
usage counter, compared against `plan.requestLimit`; reject with 429 if
`usage >= limit`, honoring `billingCycleStart` — this quota-enforcement step
exists in the current system's `checkRequestUsage` and must not be dropped in
the redesign, it was omitted from an earlier draft of this lifecycle) → business
logic → Redis `INCR` (usage, extending the Lua pattern already used by
`counters`) → response. **~4 atomic Redis ops, one process, one network hop.**
Usage flush to Postgres happens out-of-band on a batched-sync cadence, not
per-request.

---

## Authentication Architecture

**Current:** Two independent shared-secret boundaries (`X-Backend-Secret`
between Worker and Go; `X-API-Management-Key` between Rails and api-management),
plus a KV-backed per-customer key lookup that the Go backend never sees or
validates. Keys are stored as bcrypt hash + prefix in Postgres
(`apps/dashboard/app/models/api_key.rb`, correct — no plaintext key at rest
anywhere in Postgres or D1), but the actual _validation_ of a customer's key
never touches Postgres or bcrypt at request time — it's a flat KV JSON blob
lookup with no hash comparison at all (the KV value is looked up by the full key
as the literal cache key, `key:{apiKey}`, not by a hash).

**Proposed — corrected from an earlier draft, which described a lookup mechanism
that doesn't actually work with bcrypt:** you cannot look up a row by computing
`bcrypt(presented_key)` and matching it against a stored `key_hash` column —
bcrypt embeds a random salt per hash, so the same input produces a different
output every time it's hashed, unlike a deterministic digest (SHA-256/HMAC).
Equality-lookup-by-hash simply isn't possible with bcrypt; the real flow has to
be candidate-then-verify: extract `key_prefix` from the presented key
(deterministic, cheap), `SELECT ... WHERE key_prefix =
?` (needs a proper btree
index — see the corrected item below, not a unique index on `key_hash`, which
was this audit's original, incorrect target), then
`bcrypt.compare(presented_key, candidate.key_hash)` against each candidate row
(in practice almost always exactly one, given the prefix's entropy, but the code
should not assume DB-enforced uniqueness).

To avoid paying Postgres-round-trip **and** bcrypt's deliberately-slow compare
on every request, cache the verified `{user_id, plan}` result in Redis keyed by
`key_prefix` (not by a hash of the raw key — Rails discards the raw key after
creation, so it would have no way to compute an invalidation key for a
raw-key-keyed cache entry later), mirroring the existing `geocode`/`crypto`
TTL-cache pattern. **Revocation must invalidate this cache durably, not just
optimistically:** an active `DEL key_prefix` on revoke is the fast path (and,
unlike a raw-key-derived cache key, is actually computable at revoke time since
Postgres has `key_prefix`), but if that `DEL` fails (Redis unreachable,
timeout), a stale cache entry would otherwise keep serving a revoked key as
valid until its TTL expires. Recommend either retrying the `DEL` via a small
durable queue, or keeping the TTL short enough that this window is an accepted,
explicitly-documented risk — either way, this is faster to propagate than
today's system, where a revoked key can persist against a stale 60s KV cache and
depends on `sync_revocation_to_cloudflare`, which already silently fails with no
retry today (see Reliability Findings). The
Postgres-revocation-succeeds/Redis-invalidation-fails path needs explicit test
coverage before cutover, not just before/after happy-path testing.

---

## Rate Limiting Architecture

**Current:** Fixed-window per-minute counter in KV, get-then-put (non-atomic),
keyed `rl:m:{apiKey}:{minute}`, 60s TTL. Plan-tiered limits (30–50,000 req/min)
hardcoded in `apps/workers/shared/src/config.ts`.

**Recommendation: keep the algorithm, fix the atomicity.** The product's actual
requirement (per-plan requests-per-minute, documented and load-tested in
`tests/load/scenarios/rate-limit.ts`) is served fine by a fixed window — nothing
about the current product surfaces a burst-smoothing or fairness requirement
that would justify the added complexity of a token-bucket or sliding-window-log
implementation. `INCR` itself is atomic and directly resolves the KV
get-then-put race found in `rate-limit.ts:38-53`. **One correction to an earlier
framing:** a bare `INCR` followed by a separate `EXPIRE` call is _not_ atomic as
a pair — if the process crashes between the two calls, the key never gets a TTL
and persists indefinitely, which is a real contributor to the "unbounded Redis
growth" risk flagged in the Reliability Findings section below. The fix is to
make the increment-and-expire a single atomic unit, using the same Lua-script
approach already proven (for a different purpose) in
`apps/api/services/technology/counter/redis_mutations.go` — e.g.
`if
redis.call("INCR", KEYS[1]) == 1 then redis.call("EXPIRE", KEYS[1], 60) end`
in one script call. This is a small, well-precedented addition on top of the
existing pattern, not new technique. Needs test coverage beyond the happy path
before shipping: concurrent increments against the same key (confirm no
undercounting under load, unlike today's KV race), and window-rollover behavior
at the minute boundary (confirm a key from the previous window expires and a
fresh window starts cleanly rather than any off-by-one letting a burst span two
windows).

The plan-tiered limit _values_ themselves (30–50,000 req/min, currently
hardcoded in `apps/workers/shared/src/config.ts` and hand-copied into Rails per
that file's own comment) need a single new source of truth once the Workers are
retired — recommend a Postgres `plans` table Go reads (directly or via a
short-TTL Redis-cached view), replacing both the Worker's config file and any
Rails-side hand copy, so the exact config-drift risk this audit flags as an
existing problem isn't recreated in the new architecture.

---

## Usage/Billing Architecture

**Current:** auth-gateway writes one row to D1 `credit_usage` per request
(async, retried, non-blocking) → api-management exports it paginated → Rails'
`SyncD1UsageJob` polls every 3 minutes and bulk-inserts into Postgres
`usage_logs`, deduplicated by a unique index on
`(api_key_id, used_at, endpoint)`. This is **at-least-once with idempotent
dedup** — a sound design, just with an unnecessary edge-SQLite hop and up to 3
minutes of lag before usage is visible in the dashboard.

**Can this move to Go + Redis + Postgres without losing billing accuracy? For
the aggregate/enforcement side, yes** — using the same idempotent-upsert
principle already proven in `services/technology/counter`: increment a Redis
counter per `(api_key_id, day)` or similar granularity atomically at request
time, periodically flush _absolute_ values (not deltas) to Postgres via upsert,
generalizing `counter/sync_worker.go`'s mechanism. Because the flush writes
absolute values, a crash mid-cycle is self-healing on the next tick — no
double-counting, no silent loss beyond the same bounded ~60s window the
`counters` feature already accepts in production. This directly answers the
audit's core billing question for quota _enforcement_: **exactly-once is not
required and is not what the current system provides either** (D1's design is
at-least-once + dedup-by-unique-index); an aggregate-counter design preserves
that same guarantee with fewer moving parts and less lag.

**Important caveat this audit initially glossed over:** `usage_logs` today is
not just an aggregate total — it's a **row-level ledger**
(`api_key_id, endpoint, credits_used, request_method, status_code,
response_time_ms, used_at`
per request), and Rails' usage dashboard
(`apps/dashboard/app/controllers/dashboard/usage_controller.rb`) reads it as
such: a paginated list of individual recent requests, plus by-endpoint/by-date
breakdowns and an error-rate calculation that depend on those per-request
fields. A flat `namespace → total` counter (the pattern actually proven by
`counter/sync_worker.go`) can produce an aggregate total per `(api_key_id,
day)`
but **cannot reconstruct a list of individual requests with per-request status
code and response time** — that's a different data shape. The migration needs an
explicit decision here, not an assumption that "the counter pattern covers it":
either (a) keep writing per-request rows directly from Go to Postgres
`usage_logs` for the row-level/analytics use case, using Redis only for the
real-time aggregate counters that drive quota enforcement and the dashboard's
summary numbers, or (b) accept that the "recent requests" list and per-request
analytics degrade to aggregates-only and redesign that part of the dashboard.
Option (a) is very likely the right call — direct Postgres writes for the audit
trail, Redis only for the hot-path counter — and does not conflict with anything
else in this document, but it should be stated explicitly in the implementation
plan rather than assumed. See Open Questions.

---

## Security Findings (ranked)

1. **Non-atomic rate-limit counter**
   (`apps/workers/auth-gateway/src/rate-limit.ts:38-53`) — get-then-put race
   allows a customer to exceed their per-minute limit under concurrent requests.
   Low severity today (soft abuse-prevention, not a hard security boundary).
   **This audit's actual recommendation is not to patch it in place** — fixing
   it standalone in KV would mean building either a Durable Object or an interim
   Redis check, then rebuilding the real thing again during Migration Phase 2;
   given the low severity, the fix is scoped as part of Phase 2 only (see
   Architectural Smell #2 and the Implementation Plan below for why this is
   deliberately not tracked as independent work).
2. **`CF-Connecting-IP` trusted with no `trusted_proxies` configured in Rails**
   (`apps/dashboard`, no config found in `config/`) — spoofable if Rails is ever
   reached other than through Cloudflare; feeds both an IP-geolocation endpoint
   and Rack::Attack's IP throttles. Already flagged in the repo's own docs as a
   known limitation; fix is a `trusted_proxies` config, independent of the
   Workers decision.
3. **Two Rails-owned Cloudflare secrets (`BACKEND_SECRET`,
   `API_MANAGEMENT_API_KEY`) are static, unrotatable-without-coordinated-deploy
   shared secrets** spanning three independently-deployed services (Worker, VPS,
   Worker) — not a vulnerability per se, but a real operational fragility; a
   single leaked value grants broad trust with no scoping or expiry.
4. **`api_keys.key_prefix` has no efficient exact-match index** in Postgres —
   only a trigram/GIN index exists (built for fuzzy admin search), not a plain
   btree. Not currently exploitable (nothing queries by it in production yet,
   since validation happens in KV), but this is what actually needs an index
   before any migration makes Postgres the live authentication lookup — **not**
   `key_hash`, as an earlier draft of this finding claimed: bcrypt hashes are
   salted and non-deterministic, so a hash-equality lookup against `key_hash`
   isn't possible in the first place; the real lookup path is
   candidate-selection by `key_prefix` followed by a bcrypt verify (see
   Authentication Architecture above).
5. **Customer API keys are looked up in Cloudflare KV by their literal plaintext
   value**, not a hash — `key:{apiKey}` is the actual KV cache key
   (`middleware/api-key-auth.ts:50`) — unlike the bcrypt-hashed copy Postgres
   holds for the same key. This is a real inconsistency in security posture
   between the two stores (KV is Cloudflare's own encrypted-at-rest storage, so
   this is not an active exploit path today, but it means the "hash, don't store
   plaintext" property this audit credits Postgres for does not hold for the
   store actually used to authenticate every request). This is resolved by
   construction once the proposed Redis-cache-in-front-of-Postgres-hash design
   (Authentication Architecture, above) replaces the KV lookup.
6. No hardcoded secrets, no committed `.env` files, and no SQL injection or CORS
   misconfiguration were found anywhere in the repo during this audit — the
   baseline hygiene in source-controlled code is good (this does not extend to
   the KV plaintext-key finding above, which is a production-data-store fact,
   not a source-code one).

---

## Performance Findings

- **KV/D1 storage-operation count scales linearly and un-batchably with RPS** —
  using the steady-state best case established above (6 KV ops + 1 async D1
  write per request), 10k–100k RPS means **60k–600k KV operations/sec**
  sustained, plus **10k–100k D1 writes/sec** (one per request, async but still
  real write load), plus a smaller trickle of D1 reads from the periodic
  quota-cache-miss case. D1 (SQLite-per-database) is not designed for write
  concurrency anywhere near the 10k–100k/sec end of that range, which matches
  the "critical problem" already observed with KV write pressure described in
  the audit brief — the KV figure alone is large but Cloudflare KV is built for
  high read/write fan-out; the D1 figure is the harder limit.
- **Proposed design cuts this to ~4 atomic Redis ops/request (key-cache lookup,
  rate-limit increment, quota check, usage increment), zero synchronous
  Postgres/D1 ops** — usage flush is batched at a fixed 60s cadence regardless
  of request volume (the existing `counters` sync worker already demonstrates
  this scales independent of RPS, since it batches via `MGET`+one upsert per
  cycle, not per request).
- **Go's connection pools are unconfigured (library defaults)** for both
  `pgxpool` and `go-redis` — this needs explicit sizing (not "fixed," just
  _decided_) before Go becomes the sole request-serving tier; not a bottleneck
  today because the Worker/D1 tier is the current bottleneck, but it will become
  one immediately after migration if left untouched.
- **No N+1 queries, no unbounded queries, and consistent bounded-concurrency
  fan-out for batch endpoints** were found anywhere in the Go codebase — the
  business-logic tier itself is not a performance risk.

---

## Reliability Findings

- **No graceful shutdown in Go** — acceptable today (Worker absorbs client
  impact during Go redeploys), becomes a direct customer-facing outage risk once
  Go is the request-terminating tier. Must be fixed before/during migration, not
  after.
- **auth-gateway hard-fails (bare 500) on any KV/D1 error in the hot path** — no
  fail-open, no fallback, no timeout on D1 reads. The proposed Redis-based
  design should explicitly decide fail-open-vs-fail-closed for a Redis outage,
  and **rate limiting and usage/quota accounting need different answers here,
  not one shared policy**: fail-closed on api-key existence (a Redis outage
  should not let unauthenticated traffic through); fail-open on the per-minute
  rate-limit check is fine (it's a soft abuse-prevention feature — the cost of
  under-enforcing it briefly during an outage is low); but **usage/quota
  accounting must not silently fail-open**, since that would mean requests get
  served and billed-for-nothing during the outage window, permanently losing
  those usage records with no reconciliation path (this is a real
  billing-accuracy risk, not just an availability trade-off — see the Risks
  section below). Recommend either: (a) a durable fallback for usage writes
  during a Redis outage — e.g., write the row directly to Postgres synchronously
  as a degraded-but-not-lossy path — or (b) documenting and accepting a bounded
  reconciliation gap (Go still serves the request, quota enforcement is
  temporarily best-effort, but a follow-up job cross-checks Postgres
  `usage_logs` against expected volume once Redis recovers). Do not ship this
  decision unmade.
- **Single, unclustered Redis instance with no `maxmemory` configured** —
  currently low-risk (only response caches + one counter + Rails
  cache/throttles), becomes a single point of failure for the entire
  request-serving path once auth/rate-limit/usage move onto it. Needs a
  `maxmemory`+eviction policy and a documented "what happens if Redis is down"
  story before cutover — today, an unbounded Redis is a real risk to plan for at
  the RPS targets in scope. **Eviction and crash-loss are different risks that
  need different mitigations**: a crash loses at most the last ~60s of
  increments (bounded, self-healing on restart per the counter-sync design
  above); an eviction policy like `allkeys-lru` under memory pressure could
  evict a not-yet-flushed rate-limit or usage-counter key at any time, well
  before its 60s flush — an unbounded-in-frequency risk the crash scenario
  doesn't capture. The same risk applies to **Sidekiq's own queue/job state**,
  which already shares this Redis instance today — an eviction policy that
  reclaims a Sidekiq job payload under memory pressure would silently drop a
  background job, not just degrade a cache, and this instance currently has no
  such protection either. Recommend isolating all non-cache state (rate-limit
  counters, usage counters, api-key cache, and Sidekiq's queues) into a separate
  Redis logical DB or a dedicated instance with `noeviction`, keeping
  `allkeys-lru` (or similar) only for the genuinely disposable response caches
  (`geocode:*`, `crypto:*`, `exchange:*`).
- **Silent divergence between Postgres and Cloudflare KV on sync failure** —
  `ApiKey#sync_revocation_to_cloudflare` and `Subscription#sync_to_cloudflare`
  both swallow errors with only a log line, no retry, no reconciliation job.
  This class of bug (durable-store says X, cache says stale-X) disappears by
  construction once Go reads `api_keys`/`subscriptions` from Postgres directly
  instead of through a separately-synced KV mirror.

---

## Code Quality Findings

- Go: clean, well-tested, well-layered — no significant findings beyond the
  observability/pool-tuning/shutdown gaps already covered under
  Reliability/Performance.
- Rails: the concrete bugs listed under Component Audit (duplicate Cloudflare
  sync, dead MRR-affecting column, dead Pundit scaffolding, missing FKs,
  untested Cloudflare service) are the actionable findings; overall structure
  and security posture are otherwise solid.
- Workers: `auth-gateway` has materially thinner test coverage than the rest of
  the system relative to its criticality (it's the one thing every single API
  request passes through); `api-management` is comparatively well-tested. Both
  have small amounts of dead code (unused exported types/functions in
  `apps/workers/shared`, an unreachable `"daily"` quota-period branch in
  `auth-gateway/src/requests.ts`).
- CI/CD: permissions are well-scoped everywhere; CD pipeline is SHA-pinned
  (good) while CI/CodeQL use mutable tags (inconsistent but lower-risk given
  CI's read-only permissions); Brakeman/golangci-lint/RuboCop are all
  non-blocking (advisory), which is a reasonable trade-off already in place, not
  a new finding requiring action.

---

## Dead Code / Infrastructure

| Item                                                                                   | Evidence                                       | Why unused                                                                                                                                                             | Safe to remove?                                | Dependencies                                                                                    |
| -------------------------------------------------------------------------------------- | ---------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------- | ----------------------------------------------------------------------------------------------- |
| `apps/workers/auth-gateway`, `apps/workers/api-management`, `apps/workers/shared`      | Full audit above                               | Logic moves to Go/Rails/Postgres/Redis                                                                                                                                 | Yes, after staged migration (see below)        | Cloudflare KV namespace, D1 database, Wrangler secrets, `cd.yml`-adjacent manual deploy scripts |
| Cloudflare KV namespace `7cc847da...`                                                  | `wrangler.toml` in both Workers                | No consumer once Workers are removed                                                                                                                                   | Yes, after migration                           | —                                                                                               |
| Cloudflare D1 database `requiem-usage`                                                 | `wrangler.toml` in both Workers, `schema.sql`  | No consumer once Workers are removed                                                                                                                                   | Yes, after migration                           | `D1SyncService`, `SyncD1UsageJob`                                                               |
| `apps/dashboard/app/services/d1_sync_service.rb`                                       | Full file read                                 | Only pulls from D1                                                                                                                                                     | Yes, after migration                           | `SyncD1UsageJob`                                                                                |
| `apps/dashboard/app/services/cloudflare/api_management_service.rb`                     | Full file read                                 | Only talks to api-management                                                                                                                                           | Yes, after migration                           | `ApiKey`, `Subscription` callbacks                                                              |
| `apps/dashboard/app/jobs/sync_d1_usage_job.rb`                                         | Full file read                                 | Sole purpose is D1→Postgres sync                                                                                                                                       | Yes, after migration                           | Sidekiq schedule entry                                                                          |
| `CLOUDFLARE_ACCOUNT_ID`, `CLOUDFLARE_KV_NAMESPACE_ID`, `CLOUDFLARE_API_TOKEN` env vars | `.env.example:108-110`                         | Grepped zero usages anywhere in `apps/dashboard/app` (not even in `app_config.rb`, which loads every other Rails secret) — already dead today, not just post-migration | **Yes, now** — independent of migration timing | —                                                                                               |
| `apps/workers/shared/src/types.ts` `ApiKeyManagementRequest`/`Response`                | grep, zero usages                              | Never consumed                                                                                                                                                         | Yes, now (independent of migration)            | none                                                                                            |
| `apps/workers/shared/src/config.ts` `PLAN_NAMES`                                       | grep, zero usages                              | Never consumed                                                                                                                                                         | Yes, now                                       | none                                                                                            |
| `auth-gateway/src/rate-limit.ts` `getPlanLimits()`                                     | grep, only used in its own test                | Superseded by `getRequestLimitMessage()`                                                                                                                               | Yes, now                                       | none                                                                                            |
| Rails `subscriptions.stripe_customer_id/stripe_subscription_id/credit_limit`           | grep, zero usages in `app/`                    | Pre-LemonSqueezy remnants                                                                                                                                              | Yes, now (independent of migration)            | none confirmed — verify with a migration before dropping                                        |
| Rails `solid_cache_entries` table                                                      | `cache_store` always overridden away from it   | Rails-8 default never used                                                                                                                                             | Yes, now                                       | none                                                                                            |
| Rails `app/policies/` (Pundit scaffolding)                                             | Zero concrete policies, zero `authorize` calls | Dead scaffolding                                                                                                                                                       | Yes, now, or: finish wiring it up — pick one   | none                                                                                            |
| `api-management`'s `/analytics/*` endpoints                                            | No caller found in Rails                       | Built, tested, unconsumed                                                                                                                                              | Needs confirmation before deleting             | none found                                                                                      |

---

## Architectural Smells (ranked)

1. **Four uncoordinated stores for one concern (auth/rate-limit/usage).**
   Evidence: Section "Current Architecture" table above. Impact: propagation
   lag, silent divergence risk, and the KV write-pressure problem already
   observed in production. Root cause: auth/rate-limit logic was pushed to the
   edge for latency reasons without a plan for keeping edge state and origin
   state reconciled. Recommendation: consolidate per the Proposed Architecture.
   Priority: P0 at the problem level (this is the audit's top overall finding) —
   but note the actual fix is not one ticket; it's executed across the full
   P0–P3 breakdown below (P0 prerequisites, P1 = Migration Phases 1–3, P2 =
   Phases 4–7, P3 = Phase 8), not a single standalone P0 line item.
2. **Non-atomic rate limiting under the exact conditions it exists to prevent
   (bursty, concurrent, abusive traffic).** Evidence: `rate-limit.ts:38-53`.
   Impact: correctness gap in abuse prevention scales with concurrency, i.e.,
   gets worse exactly as traffic grows. Root cause: KV has no atomic increment
   primitive. Recommendation: a Lua-scripted atomic `INCR`+`EXPIRE` in Redis (a
   plain two-command `INCR` then `EXPIRE` is not atomic as a pair — see Rate
   Limiting Architecture above). Priority: **not a standalone P0** — this
   audit's recommended fix builds it as part of Migration Phase 2, so it's
   tracked under the P1 migration-phases bucket in the Implementation Plan
   below, not as independent standalone work.
3. **Duplicate side-effecting HTTP call on every subscription plan change.**
   Evidence: `subscription.rb:26` +
   `lemonsqueezy_controller.rb:108,160,186,210`. Impact: 2x unnecessary load on
   api-management, 2x unnecessary KV writes, harder-to-reason-about failure
   modes (which of the two calls failed?). Root cause: the callback was likely
   added after the explicit call already existed, or vice versa, without
   noticing the overlap. Recommendation: remove one (keep the callback, since it
   also covers non-webhook plan changes like admin promotions). Priority: P1,
   independent of the migration. 3b. **Correction to the recommendation above:**
   `Subscription#sync_to_cloudflare` fires via a plain
   `after_update`/`after_create` callback (`subscription.rb:25-26`), which runs
   _inside_ the enclosing DB transaction — the LemonSqueezy webhook controller's
   explicit call was deliberately placed _after_
   `ActiveRecord::Base.transaction do ... end` specifically to keep HTTP I/O out
   of the transaction (see the comment at `lemonsqueezy_controller.rb:105-107`).
   Naively "keeping the callback, dropping the explicit call" would leave the
   transaction-blocking version as the only remaining path — the opposite of
   what the existing code already tried to avoid. The correct fix is to convert
   the callback to
   `after_commit`/`after_create_commit ... if: :saved_change_to_plan_name?` _and
   then_ drop the explicit calls — this gets both properties (single call site,
   no HTTP I/O inside the transaction) instead of trading one problem for the
   other.
4. **Silent MRR miscalculation.** Evidence: `subscriptions.plan` never written,
   consumed by two live analytics surfaces. Impact: revenue reporting is wrong
   for every yearly subscriber, silently. Root cause: incomplete LemonSqueezy
   integration — billing-cycle metadata was never wired from the webhook payload
   into this column. Recommendation: populate `plan` from the LemonSqueezy
   variant/webhook payload, backfill existing rows. Priority: P1, independent of
   the migration, arguably more urgent than the architecture work since it
   affects real business reporting today.
5. **Go has no graceful shutdown or structured observability.** Evidence:
   Section "Go Backend" above. Impact: currently masked by the Worker sitting in
   front; becomes a direct reliability and debuggability gap the moment Go is
   internet-facing. Root cause: built for an internal-trust-boundary role, not
   scrutinized for internet-facing-service hygiene. Recommendation: add
   before/during Phase 1 of the migration, not after. Priority: P0 (blocking for
   the migration, not optional).
6. **Two independently-implemented Faraday clients in Rails against the same
   Worker.** Evidence: `Cloudflare::ApiManagementService` + `D1SyncService`.
   Impact: duplicated retry/timeout/auth logic, inconsistent naming/namespacing.
   Root cause: organic growth, no consolidation pass. Recommendation: both are
   deleted by the migration anyway; not worth consolidating in place. Priority:
   P3 (superseded by migration).
7. **Dead Pundit scaffolding creating a false impression of a policy layer.**
   Evidence: Section "Rails" above. Impact: low — actual authorization is
   enforced elsewhere (router + controller `before_action`), but a future
   contributor could reasonably assume Pundit is the enforcement mechanism and
   be wrong. Recommendation: either delete the scaffolding or actually wire it
   up; low priority either way. Priority: P3.

---

## Documentation Audit

Docs that describe the current Worker-based design and would become incorrect
once `auth-gateway`/`api-management`/KV/D1 are retired (Migration Phase 8):

| File                           | Section                                                                 | Current claim                                                                                  | Required change                                                                                                                                                                                                                    |
| ------------------------------ | ----------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `docs/core/architecture.md`    | Whole doc                                                               | Diagrams and prose describe Cloudflare KV/D1 as the auth/rate-limit/usage layer in front of Go | Rewrite to describe Go-native auth/rate-limit/usage backed by Redis/Postgres; remove KV/D1 references entirely                                                                                                                     |
| `docs/core/auth-gateway.md`    | Whole doc                                                               | Documents `auth-gateway`'s routes, KV/D1 schema, deployment                                    | Delete, or replace with a short pointer to the Go auth/rate-limit middleware's own doc                                                                                                                                             |
| `docs/core/api-management.md`  | Whole doc                                                               | Documents `api-management`'s API-key CRUD/usage-export endpoints                               | Delete; API-key CRUD becomes plain Rails model/controller documentation                                                                                                                                                            |
| `docs/core/infrastructure.md`  | "Architecture Components" diagram, Redis section (`:143-153`)           | Shows Cloudflare Worker + KV/D1 layer; describes Redis as "Sidekiq + Go counter storage" only  | Update diagram to remove the Worker/KV/D1 tier; expand the Redis section to include auth-cache/rate-limit/usage-counter namespaces                                                                                                 |
| `docs/core/deployment.md`      | "Part 5: Deploy Cloudflare Workers" (`:333-428`)                        | Full deploy walkthrough for both Workers, KV/D1 provisioning                                   | Delete this Part entirely; add Cloudflare origin-lockdown (Authenticated Origin Pulls / IP firewall) setup steps in its place, per Migration Phase 6                                                                               |
| `docs/core/rails-app.md`       | "Cloudflare Integration" (`:275-329`), API Management Worker references | Describes `CloudflareApiManagementService`/D1 usage-data flow as current                       | Update to describe direct Postgres key CRUD and Go-native usage writes; also independently correct the "(separate schema)" mischaracterization at `:207-213` regardless of migration timing (see PostgreSQL component audit above) |
| `agents.md` (repo root)        | Component descriptions                                                  | Lists "Auth Gateway" and "API Management" as active components                                 | Update once Phase 8 completes                                                                                                                                                                                                      |
| `docs/apis/**/*.md` (64 files) | curl examples                                                           | Reference `api.requiems.xyz` as the public endpoint                                            | No change needed — this is the public-facing URL regardless of what serves it behind the scenes; Cloudflare's DNS/proxy layer keeps the same hostname throughout the migration                                                     |

This list is not exhaustive — a full-text search for `auth-gateway`,
`api-management`, `KV`, and `D1` across `docs/` immediately before Phase 8 is
the reliable way to catch anything missed here.

---

## Migration Plan

This assumes the architecture decision below is accepted. Each phase preserves
current behavior until the next phase explicitly changes it — no "delete Worker,
deploy Go, hope" step exists in this plan. **Two corrections from the first
draft of this plan are folded in below: a sequencing bug in what is now Phase 4
that would have broken authentication for new signups, and an unworkable
mechanism in what is now Phase 5 — Cloudflare Workers intercept 100% of traffic
to a matched route by default, so "weighted DNS routing" cannot split traffic
between "through the Worker" and "direct to Go" the way the original phase
implied.**

**Phase 0 — Fix the blocking Go gaps (prerequisite, not migration-specific).**
Add graceful shutdown (`signal.NotifyContext` + `server.Shutdown`), structured
logging, explicit `pgxpool`/`go-redis` pool sizing to `apps/api`, and a btree
index on `api_keys.key_prefix` in Rails (not `key_hash` — see the corrected
Authentication Architecture section above for why). Concrete acceptance criteria
(these are underspecified if left as prose): structured logging = `log/slog`,
wired into Go's HTTP middleware so every request logs request-id, method, route,
status, and latency as JSON to stdout (picked up by Docker/Kamal's log
collection), with `Sentry` continuing to own uncaught-exception capture; pool
sizing = explicit `pgxpool.Config.MaxConns`/`MinConns` and `go-redis`'s
`PoolSize`, set relative to Postgres's own `max_connections` and the number of
Go replicas (get the real current-traffic numbers per Open Question 5 before
picking a value, don't guess). Also fix the dev-seed key-format bug in this
phase (cheap, independent, and needed before Phase 3's shadow-comparison work
can be manually exercised): `scripts/seed-dev.ts` and
`docs/core/auth-gateway.md`'s example both use `rq_free_000001`-shaped keys that
the live `requiem_[0-9a-zA-Z]{24}` validator rejects.

**Phase 1 — Introduce Go-side API-key authentication, dual-running alongside the
Worker, shadow-only.** Add a new Go middleware that validates `requiems-api-key`
by candidate-selecting on `key_prefix` and bcrypt-verifying against
`api_keys.key_hash` (Redis-cached by `key_prefix` — see Authentication
Architecture above). Pick one gating mechanism (a feature-flagged route subset,
not both a flag and a Worker-set header) and commit to it for this phase.
**Explicit exit criterion:** the middleware logs its auth decision for every
request without affecting the response, for a defined minimum period (recommend
two weeks) with zero unexplained false-positive/false-negative divergence from
the Worker's live decision on a sampled comparison.

**Phase 2 — Introduce Redis-based rate limiting and usage counting,
shadow-only.** Implement atomic rate limiting via a Lua-scripted `INCR`+`EXPIRE`
(see Rate Limiting Architecture above — not a bare two-command `INCR`+`EXPIRE`,
which is not atomic as a pair), and extend the `counter` package's sync
mechanism — not copy its data model as-is — to a usage tracker with
per-`api_key`/per-`user` granularity and monthly-cycle resets, flushing to
Postgres. Decide up front whether row-level `usage_logs` detail (individual
requests with status code/response time, needed by the Rails usage dashboard) is
written directly from Go per-request or reconstructed some other way — see the
Usage/Billing Architecture caveat above; do not assume the aggregate counter
alone covers this. Assign a single source of truth for the plan-tier rate-limit
_values_ (currently hardcoded in `apps/workers/shared/src/config.ts` and
hand-copied into Rails) — recommend a Postgres `plans` table. Run everything in
this phase in parallel with the Worker's KV/D1 path, logging only, no
enforcement. **Exit criterion:** rate-limit and usage counts logged by Go match
the Worker's KV/D1-derived numbers within an agreed tolerance for a defined
period.

**Phase 3 — Dual-operation comparison and explicit enforcement cutover for
directly-routed traffic.** Compare Go's shadow auth/rate-limit/usage decisions
against the Worker's live decisions for a sustained period — **at least one full
billing cycle**, so monthly quota-reset behavior is exercised. **Acceptance
threshold:** zero unexplained usage/quota discrepancies across the full cycle;
every discrepancy found must be root-caused (e.g., a timezone bug in the monthly
reset, an edge case in the batch-endpoint multiplier logic), not merely logged
and waved through. Do not proceed to Phase 5 on a partial or unexplained pass.
**State the transition explicitly, since it's easy to leave implicit:** at the
end of this phase, Go's auth/rate-limit/usage logic is declared _ready to
enforce_, even though it isn't yet serving any real traffic — that readiness
gate is what Phase 5 depends on.

**Phase 4 — Move key-generation ownership to Rails, with the Cloudflare mirror
kept in sync.** First, fix `ApiKey#generate_key_locally`'s output format
(`requiem_` prefix, not `rq_live_`/`rq_test_`) and add real collision handling
(the `key_prefix` index from Phase 0 makes a pre-write collision check feasible,
plus a retry-on-collision loop mirroring api-management's `create.ts` behavior)
— see the Rails component audit above for why this isn't optional polish. **Do
not simply stop calling api-management for new keys at this point** — an earlier
version of this plan did, which would 401 every new signup between this phase
and full cutover, since the Worker (still the production auth gate through
Phase 5) only knows about keys present in Cloudflare KV. Instead, Rails should
generate the key locally and push it into KV for the transition period — **but
note this requires new work, not just a dual call to the existing endpoint**:
api-management's `POST /api-keys` (`create.ts`) always generates its own key
server-side and has no request field to accept an externally-supplied key, so
naively "pushing" Rails' locally-generated key to the current endpoint would
make api-management mint and store a _different_ key than the one Rails shows
the user, silently reintroducing the same 401-on-signup failure via a different
path. Either (a) extend api-management's create endpoint (or add a new one) to
accept and persist an externally-supplied key, or (b) have Rails write directly
to the KV namespace via the Cloudflare API for this transitional dual-write,
bypassing api-management's key-generation logic entirely for this one call while
still going through it for revoke/update. Also decide and document the failure
mode if this dual-write's Cloudflare-side call fails mid-signup (block signup
and surface an error, or proceed and rely on a retry/reconciliation job) — don't
let it silently degrade to the same swallow-and-log pattern already flagged as a
problem elsewhere in this document. The Cloudflare round trip for new-key
creation is only fully removed once Phase 6 completes and the Worker is no
longer the auth gate for any traffic. Revocation/update continue round-tripping
to Cloudflare KV through this phase for the same reason.

**Phase 5 — Move a small percentage of production traffic directly to Go,
enforcing.** The mechanism must live _inside_ auth-gateway, not at the DNS/route
level — a Cloudflare Worker route match is all-or-nothing per hostname/pattern,
so "weighted DNS routing" cannot split traffic between "goes through the Worker"
and "bypasses it" the way a naive canary plan implies. Concretely: add a
config-driven percentage (or consistent-hash-by-api-key, for a stable
per-customer experience) check near the top of the Worker's request handling
that, for the selected slice, skips the Worker's own KV/D1 auth/rate-limit/
usage logic entirely and passthrough-proxies straight to Go — which, per Phase
3's exit criterion, is now the _enforcing_ authority for that slice, not a
shadow logger. Define a rollback trigger before starting (e.g., error-rate delta
or rate-limit/quota mismatch rate crossing a threshold vs. the Worker-routed
control group) and who/what monitors it during the canary window.

**Phase 6 — Full cutover.** Route all traffic directly to Go (the Worker
passthrough percentage from Phase 5 goes to 100%, or the Worker route is removed
entirely — either achieves the same end state). Cloudflare config changes from
"Worker in front" to "proxy-only" (orange-cloud DNS, WAF/DDoS rules retained).
`X-Backend-Secret` and the Caddy `@authorized` gate
(`infra/caddy/Caddyfile:32-47`) are removed since Go is now directly the
internet-facing origin behind Cloudflare's proxy. **This is a hard requirement,
not an optional step:** an equivalent origin-lockdown control (e.g.,
Cloudflare's Authenticated Origin Pulls, or origin firewall rules restricting
inbound to Cloudflare's published IP ranges) must be in place _before_ this
phase completes — without it, Go becomes directly internet-reachable with zero
origin authentication the moment the old secret-header gate is removed, which
directly contradicts the Proposed Architecture's premise that "Cloudflare stays
as DNS/WAF/DDoS/TLS." **Verification step, not just configuration:** before
declaring this phase done, confirm a direct request to the origin's IP/hostname
(bypassing Cloudflare's proxy) is actually rejected — configuring Authenticated
Origin Pulls or a firewall rule is not sufficient on its own; test it. Configure
`trusted_proxies`/equivalent client-IP handling in Go at this point too, since
`CF-Connecting-IP` semantics now apply directly to Go instead of being
pre-filtered by the Worker.

**Phase 7 — Remove D1/KV synchronization.** Delete `SyncD1UsageJob`,
`D1SyncService`, `Cloudflare::ApiManagementService`,
`Subscription#sync_to_cloudflare` (converted to `after_commit` per the P1 bug
fix above, then deleted entirely here),
`ApiKey#sync_revocation_to_cloudflare`/`remove_from_cloudflare` (these were
never part of the P1 `after_commit` fix — that fix only covered `Subscription`;
`ApiKey`'s callbacks are simply deleted whole in this phase since Cloudflare
sync is being removed, not converted), and the temporary dual-write mechanism
added in Phase 4. Remove `CLOUDFLARE_*` env vars.

**Phase 8 — Remove Workers.** Delete `apps/workers/auth-gateway`,
`apps/workers/api-management`, `apps/workers/shared`. Delete the Cloudflare KV
namespace and D1 database. Remove their entries from `.dependabot.yml`,
`docker-compose.dev.yml`, `_worker-ci.yml` and its callers in `ci.yml`. Update
`docs/core/{architecture,auth-gateway,api-management}.md` (see Documentation
Audit above) and remove `tests/integration/src/suites/gateway.test.ts`'s
Worker-specific assertions / `tests/load/scenarios/rate-limit.ts`'s KV-specific
assertions in favor of Go-equivalent checks (both suites already hit the public
gateway URL, not Worker internals directly, so most of
`tests/integration`/`tests/load` needs no changes beyond re-pointing
`API_BASE_URL`).

---

## Implementation Plan

**P0 — Critical (blocking, do first)**

- Add graceful shutdown to `apps/api/main.go`. _Independently doable, no
  dependencies, small._
- Add structured logging to `apps/api`: `log/slog` in HTTP middleware, JSON to
  stdout, fields = request-id/method/route/status/latency at minimum. _touches
  many files, medium complexity, low risk._
- Explicit `pgxpool`/`go-redis` pool config in
  `apps/api/platform/{db,reqredis}`, sized against Postgres's `max_connections`
  and the real number of Go replicas/current traffic (see Open Question 5 —
  don't size blind). _small, needs load-test validation._
- Add a plain btree index on `apps/dashboard` `api_keys.key_prefix` for
  efficient exact-match candidate lookup during auth (the existing index there
  is a trigram index for fuzzy admin search, not this) — **not** a unique index
  on `key_hash`, which was this item's original target in an earlier draft:
  bcrypt hashes can't be looked up by equality in the first place, see
  Authentication Architecture above. _One migration, small._
- Fix the dev-seed API-key format bug (`scripts/seed-dev.ts` /
  `docs/core/auth-gateway.md` use `rq_free_*`-shaped keys the live
  `requiem_[0-9a-zA-Z]{24}` validator rejects). _Independently doable, trivial,
  and needed before Phase 3's shadow-comparison work can be exercised manually._
- **Not actually an independent P0 item, despite the framing below:** the
  non-atomic KV rate-limit race is real (see Security Findings #1), but this
  audit's own recommended fix is to build the atomic replacement as part of
  Migration Phase 2, not to patch KV in place — so it's tracked under the P1
  migration-phases item below, not standalone here. Listed for visibility only.

**P1 — High priority**

- Fix duplicate Cloudflare sync on subscription plan change: convert
  `Subscription#sync_to_cloudflare`'s callback to `after_commit`/
  `after_create_commit` (not a plain `after_update`/`after_create`, which would
  otherwise block the callback's HTTP call inside the DB transaction), then drop
  the now-redundant explicit calls in
  `webhooks/lemonsqueezy_controller.rb:108,160,186,210`. _Independent of
  migration, small, needs a test added first since neither path is currently
  tested._
- Fix MRR calculation (populate `subscriptions.plan` from LemonSqueezy webhook
  payload, backfill). _Independent of migration, needs care around backfill
  correctness — recommend writing a one-off Rails runner script, verifying
  against LemonSqueezy's own dashboard totals before trusting the backfill._
- Migration Phases 1–3 (Go-side auth/rate-limit/usage in shadow mode; Phase 0's
  prerequisites are the standalone P0 bullets above). _Depends on P0 items
  above. Medium-high complexity — this is the architectural core of the audit's
  recommendation. Needs new Go tests for the auth/rate-limit/usage paths and a
  comparison/reconciliation script for Phase 3._
- Configure `trusted_proxies` in Rails for `CF-Connecting-IP` handling
  (`ApiProxyController`). _Small, and — unlike originally framed — this is
  **not** tied to the Workers migration timeline: Rails (`requiems.xyz`) has
  always been fronted by Cloudflare directly, never by auth-gateway (confirmed
  in `infra/caddy/Caddyfile`, separate site blocks), so this can and should ship
  independently, any time._

**P2 — Medium priority**

- Migration Phases 4–7 (key-generation ownership, traffic cutover, D1/KV sync
  removal). _Depends on P1 migration phases. Requires coordinated Cloudflare
  routing changes and careful monitoring during Phase 5's canary period._
- Add equivalent client-IP/origin-trust handling in Go once it becomes the
  direct target of Cloudflare's proxy. _Genuinely tied to Migration Phase 6
  (full cutover, itself in the P2 bucket above) — see that phase for the
  origin-lockdown requirement it's bundled with. Listed here rather than under
  P1 specifically because it cannot happen before Phase 6 regardless of how
  early the code for it is written._
- Add missing FKs on `credit_adjustments.admin_user_id`,
  `audit_logs.admin_user_id`, `abuse_reports.resolved_by_id`. _Independent,
  small, needs a data-cleanup pass first if any orphaned values exist._
- Add test coverage for `Cloudflare::ApiManagementService` and the three
  untested Sidekiq jobs — _only worth doing if Phase 4+ of the migration is
  delayed; otherwise this code is being deleted, don't invest in testing code on
  its way out._

**P3 — Nice to have**

- Migration Phase 8 (delete Workers, KV, D1, related CI/docs/deps).
- Drop dead Rails columns (`stripe_*`, `credit_limit`) and the unused
  `solid_cache_entries` table.
- Remove the already-unused
  `CLOUDFLARE_ACCOUNT_ID`/`CLOUDFLARE_KV_NAMESPACE_ID`/ `CLOUDFLARE_API_TOKEN`
  template entries from `.env.example` — zero code references found in
  `apps/dashboard/app`, so this doesn't need to wait for the migration.
- Resolve the Pundit scaffolding one way or the other.
- Confirm and remove api-management's unused `/analytics/*` endpoints (or find
  their intended Rails caller and wire it up, if this was simply an incomplete
  feature rather than dead code).
- Add dependabot coverage for `apps/mcp`.
- Fix the dev-compose Redis port binding (`0.0.0.0:6379` → `127.0.0.1:6379`) for
  consistency with `db`'s binding.
- Verify live-VPS Kamal `db`/`redis` accessories aren't accidentally duplicating
  the docker-compose-managed instances.

---

## Risks

- **The migration touches the billing-critical path.** Any bug in the
  Redis-based usage-counting design that under- or over-counts would directly
  affect customer invoices. Mitigation: Phase 3's shadow-comparison period is
  not optional — do not skip or shorten it, and specifically test month-boundary
  quota reset behavior, which is the kind of edge case most likely to diverge
  between the old and new systems.
- **Removing `X-Backend-Secret`/Caddy's `@authorized` gate without a real
  replacement makes Go directly internet-reachable with zero origin
  authentication.** This is now called out as a hard requirement in Migration
  Phase 6, not an optional step, precisely because it's easy to treat as a
  cleanup nicety and skip under deadline pressure — it isn't one. Verify the
  replacement (Authenticated Origin Pulls or origin-IP firewalling) is actually
  in place and tested before Phase 6 is considered complete, not just planned.
- **Cloudflare's role changes from "does auth" to "does nothing but proxy,"
  which changes the DDoS/abuse posture** — today, a volumetric or
  credential-stuffing attack against a specific API key is rejected at the edge
  (KV lookup, cheap) before ever reaching Go; post-migration, the same rejection
  happens in Go behind Cloudflare's generic WAF/rate-limiting (not the
  plan-aware, per-key logic being migrated). Cloudflare's WAF rules should be
  reviewed/tightened as part of Phase 6, not assumed to be equivalent.
- **Redis becomes a single point of failure for the entire request-serving
  path**, not just for a handful of internal caches. The "what happens if Redis
  is down" question needs an explicit, tested answer before Phase 6, not
  discovered live during an incident — and rate-limit and usage/quota checks
  need different answers, not one shared "fail open" policy: fail-open on the
  rate-limit check is acceptable (soft abuse prevention), fail-closed on
  key-existence is required (no unauthenticated traffic), and usage/quota
  accounting needs a durable fallback or an explicitly-accepted reconciliation
  gap rather than silently fail-open — see Reliability Findings above for the
  full reasoning; failing open there means billed usage is permanently lost,
  which is a different class of risk than briefly under-enforcing a rate limit.
- **The 3-minute D1→Postgres lag is currently a soft, mostly cosmetic delay**
  (dashboard usage numbers a few minutes stale); a bug in the new Redis-based
  flush could turn "slightly stale" into "silently wrong," which is a worse
  failure mode even if rarer. Mitigation: keep the flush's
  idempotent-absolute-value design (matching `counter/sync_worker.go`), not a
  delta-based design, specifically because it self-heals after a crash.

---

## Open Questions

1. Live-database verification of the migration-tracking-table situation was not
   performed (out of scope for a read-only audit) — confirm
   `rails_schema_migrations` vs `schema_migrations` actually behave as expected
   on the production database, not just in the initializer's intent.
2. Whether the Kamal-deployed `db`/`redis` accessories
   (`infra/kamal/deploy.api.yml`, `deploy.dashboard.yml`) are the same physical
   containers as the docker-compose-managed ones, or accidental duplicates —
   needs a live-VPS check.
3. What Cloudflare WAF/rate-limiting rules exist today outside the Worker (not
   visible from the repo) — needed to design Phase 6's replacement
   abuse-prevention posture.
4. Whether `api-management`'s `/analytics/*` endpoints are truly unused by Rails
   or called from somewhere not found by static analysis (e.g., an ops script, a
   manual curl workflow) — worth a direct question to whoever built that feature
   before deleting it.
5. What the actual current production RPS is — this audit worked from the stated
   10k–100k RPS _target_, not measured current traffic; the urgency of Phase 0's
   pool-tuning/observability work depends on how close current traffic already
   is to today's un-tuned defaults' limits.
6. Whether the `plan` billing-cycle backfill for the MRR fix can be reliably
   reconstructed from LemonSqueezy's own records (webhook payloads may not have
   been retained), or whether some subscriptions will need manual
   reconciliation.
7. Whether Rails' usage dashboard's row-level views (recent-requests list,
   per-request status code/response time) need to keep working exactly as they
   do today. A flat Redis aggregate counter (the pattern generalized from
   `services/technology/counter`) cannot reconstruct that row-level detail — the
   likely answer is that Go writes per-request rows directly to Postgres
   `usage_logs` for this purpose, independent of the Redis-based aggregate
   counters used for quota enforcement, but this should be a decision made
   explicitly during Migration Phase 2, not discovered midway through it. See
   the Usage/Billing Architecture caveat above.

---

## Final Recommendation

Proceed with consolidating auth, rate limiting, and usage tracking into the Go
backend + Redis + PostgreSQL, retiring `auth-gateway`, `api-management`,
`apps/workers/shared`, Cloudflare KV, and Cloudflare D1. The original hypothesis
is correct, and the audit found concrete, current evidence supporting it beyond
the original concern about KV write pressure: a real correctness bug in the KV
rate limiter, thin test coverage on the exact code path every request depends
on, four uncoordinated stores for one concern, and a production-proven
Redis→Postgres sync mechanism already living in the codebase
(`services/technology/counter`) that the new rate limiter and usage tracker can
be built on. **Be precise about what that existing code proves, since an earlier
draft of this audit overstated it:** it proves the crash-safe, idempotent
flush-to-Postgres mechanism works — it does not itself implement rate-limit
windowing (no TTL), variable-increment billing multipliers, or multi-dimensional
(per-key/per-user/per-endpoint) usage tracking, all of which are real, scoped,
buildable work on top of that foundation, not code that already exists. This is
not a bet on an unproven architecture, but it is more than a pure cutover —
treat the sync mechanism as validated, and the rate-limiter/usage-tracker as new
work informed by a good template.

This is **not** a "delete the Workers this sprint" recommendation. Go needs
graceful shutdown and basic observability before it can be trusted as the sole
internet-facing tier (Phase 0), and the migration must run in shadow/comparison
mode through at least one full billing cycle (Phase 3) before any production
traffic is cut over, given that usage/billing accuracy is the single
highest-consequence failure mode in scope. Keep Cloudflare in front throughout
and after the migration — its DNS/WAF/DDoS/TLS role is sound and was never in
question; only its Worker-based business logic is being retired.

Independent of the migration timeline, fix the duplicate-Cloudflare-sync bug and
the MRR calculation bug now — both are live correctness issues affecting real
business operations today, cost little to fix, and have nothing to do with the
Workers/KV/D1 decision.
