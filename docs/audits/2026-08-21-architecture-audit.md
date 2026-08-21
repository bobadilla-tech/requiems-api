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
Redis (via Go) holds an unrelated set of response caches and one internal
counter primitive. None of these four stores is aware of the others in real time
— the two Workers coordinate over HTTP, and Rails pulls D1 into Postgres on a
3-minute cron. The Go backend, which does all the actual business logic, sits
behind a single shared secret and has **no concept of an API key, a user, a
plan, or a rate limit** — every one of those concerns lives in a
220-line-of-glue-logic edge Worker in front of it.

This split was a reasonable choice when the goal was "sub-10ms global auth
checks," but the audit found the edge tier is not actually buying that guarantee
cleanly: the KV-based rate limiter is a **non-atomic get-then-put**
(`apps/workers/auth-gateway/src/rate-limit.ts:38,51`) that races under
concurrent requests from the same key, the quota-check path has no timeout or
fallback if KV/D1 are slow (`middleware/api-key-auth.ts:50,64`), and the proxy's
hot path is essentially untested end-to-end (no test exercises a full successful
proxied request). Meanwhile the Go backend already imports `pgx` and `go-redis`
and already runs a small, well-designed Redis→Postgres batched-counter-sync
worker (`apps/api/services/technology/counter/sync_worker.go`) that is
architecturally identical to what a Redis-based rate-limiter/quota-tracker would
need — the target stack isn't hypothetical, a working prototype of its core
mechanism already ships in production for an unrelated feature.

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
hot-path state, Postgres for durable state) using a pattern the Go codebase has
already proven. This removes an entire deployment target (two Workers,
unautomated in CI), an entire edge database (D1, whose SQLite-per-database write
model does not scale to the 10k–100k RPS target the product hypothesis
mentions), and ~2,000 lines of Rails/TS glue code whose only job is keeping
those stores in sync.

**Recommended target architecture:** Cloudflare stays as DNS/WAF/DDoS/TLS only
(no Worker logic). The Go backend gains its own API-key authentication, atomic
Redis-based rate limiting, and Redis-based usage counting flushed to Postgres on
the same batched-sync pattern already in production. Rails keeps owning
`users`/`api_keys`/`subscriptions` as the durable source of truth and generates
keys directly (the code path already exists in test mode —
`apps/dashboard/app/models/api_key.rb:60-68` — and just needs promoting).
`auth-gateway`, `api-management`, `apps/workers/shared`, Cloudflare KV, and
Cloudflare D1 are retired in a staged migration (Section: Migration Plan). This
is not a hypothetical rewrite — every piece of the target stack (`pgx`,
`go-redis`, the dirty-set-swap sync pattern, Rails' local key generation)
already exists in the codebase today; the work is extension and cutover, not
invention.

---

## Current Architecture

```
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
   for one internal        from D1 via api-management], daily_usage_
   billable metric)        summaries, credit_adjustments, audit_logs …)

  Cloudflare Worker: api-management.requiems.xyz (internal-only)
    — API key CRUD (dual-writes KV + D1)
    — usage export endpoint (Rails' cron reads this)
    — analytics endpoints (no confirmed caller in Rails)
    Auth: static X-API-Management-Key shared secret
```

**Four operational-state stores for one concern (auth/rate-limit/usage):**

| Store         | Holds                                                                                                                                                 | Written by                                                        | Read by                          | Propagation lag to Postgres                          |
| ------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------- | -------------------------------- | ---------------------------------------------------- |
| Cloudflare KV | `key:{apiKey}` → user/plan, `rl:m:{apiKey}:{minute}`, `quota:{userId}:{cycleStart}` (60s cache)                                                       | api-management (key data), auth-gateway (rate/quota cache)        | auth-gateway                     | N/A — KV state itself is never the sync source       |
| Cloudflare D1 | `credit_usage` (every request), `api_keys` (audit mirror)                                                                                             | auth-gateway (usage rows), api-management (key CRUD)              | api-management (`/usage/export`) | up to 3 minutes (Sidekiq cron)                       |
| PostgreSQL    | `usage_logs`, `daily_usage_summaries`, `api_keys`, `subscriptions`, `users`, …                                                                        | Rails (`SyncD1UsageJob`, `AggregateDailyUsageJob`, direct writes) | Rails dashboard/admin            | — (destination)                                      |
| Redis         | Go response caches (`geocode:*`, `crypto:*`, `exchange:*`), `counter:*` (one internal metric), Rails `Rack::Attack` throttle counters, Sidekiq queues | Go, Rails                                                         | Go, Rails                        | `counter:*` flushed to Postgres `counters` every 60s |

No store above is authoritative for more than its own slice, and nothing
reconciles them faster than the 3-minute D1→Postgres cron. A revoked key
propagates from Rails → api-management → KV delete typically within one HTTP
round trip (fast), but a plan/quota change synced only into KV can silently
diverge from Postgres if that HTTP call fails — both
`ApiKey#sync_revocation_to_cloudflare`
(`apps/dashboard/app/models/api_key.rb:99-103`) and
`Subscription#sync_to_cloudflare`
(`apps/dashboard/app/models/subscription.rb:30-35`) swallow errors and only log;
there is no reconciliation job that would catch a Postgres/KV divergence.

---

## Proposed Architecture

```
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
│   atomic INCR+EXPIRE)   │
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
zero edge-to-origin HTTP hops for auth (currently: client → Worker → KV → KV →
D1 → Go → Postgres; proposed: client → Go → Redis), and Rails talks to Postgres
directly for key CRUD instead of round-tripping to a Worker. Cloudflare is
retained purely as network infrastructure — its removal is explicitly **not**
recommended (see Question 11 below).

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

Best case: **4 KV reads + 2 KV writes + 1 async D1 write** per request. Worst
case (quota cache miss, ~once/user/minute): **+1 D1 read +1 KV write**.

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

CRUD-correct: create writes D1 before KV (so a D1 failure never leaves an
orphaned KV entry, `routes/api-keys/create.ts:76-88`), revoke/patch are D1-first
for audit-trail integrity, list never returns full key values. Auth is a
constant-time SHA-256 compare against a static shared secret
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
No N+1 patterns found; batch endpoints consistently use bounded goroutine
fan-out.

**But it is not yet fit to be the sole internet-facing tier**, independent of
the Workers question:

- **No graceful shutdown.** `main.go` calls `server.ListenAndServe()` with no
  `signal.NotifyContext`/`server.Shutdown()` anywhere — SIGTERM during a rolling
  deploy kills in-flight requests, the pgx pool, and the Redis client outright.
  Acceptable today because the Worker in front absorbs client-facing disruption
  during a Go redeploy; not acceptable if Go becomes the thing clients connect
  to directly.
- **No structured logging, no metrics, no tracing.** `log.Printf` in 25 files,
  Sentry only on uncaught errors (`TracesSampleRate: 0.01`). The observability
  audit's baseline questions ("which endpoint is slow," "is Postgres slow")
  cannot be answered from what exists today.
- **No connection-pool tuning.** `pgxpool` and `go-redis` both run on library
  defaults (`platform/db/db.go`, `platform/reqredis/redis.go`) — fine at current
  load, needs explicit sizing before any RPS target discussion is meaningful.
- **A ready-made usage-tracking hook exists but is used in 1 of ~220
  endpoints.** `platform/httpx/handler.go:33-36` already defines a
  `UsageCounter` interface that auto-sets `X-Usage-Count` for any response
  implementing it — the exact mechanism the auth-gateway reads to apply
  per-endpoint billing multipliers. A migration would need this backfilled
  broadly, but the extension point already exists and is proven.
- **The counter-sync pattern is the template to reuse, not reinvent.**
  `services/technology/counter/{redis_mutations.go,sync_worker.go,repository.go}`
  implements exactly the shape a Redis-based rate-limiter/usage-tracker needs:
  Lua-scripted atomic increment-and-mark-dirty, a `RENAMENX`-based dirty-set
  swap that's safe against crash-mid-cycle (idempotent because it upserts
  _absolute_ values, not deltas), and a single batched Postgres upsert every
  60s. This is real, running, tested code — not a design sketch.

### Rails (`apps/dashboard`)

Solid separation from Go's tables (confirmed: no raw SQL crosses the ownership
boundary in either direction; Rails explicitly excludes Go-owned tables from its
schema dumper, `config/application.rb:37-52`). Devise + Rack::Attack cover authn
and throttling for Rails' own surfaces reasonably well; the LemonSqueezy
webhook's signature verification is correctly constant-time and well-isolated
from the DB transaction.

**Concrete bugs found, independent of the Workers decision:**

- **Duplicate Cloudflare sync on every plan change.**
  `Subscription#sync_to_cloudflare` fires via an `after_update` callback
  (`subscription.rb:26`) _and_ the LemonSqueezy webhook controller explicitly
  calls the same sync again right after
  (`webhooks/lemonsqueezy_controller.rb:108,160,186,210`) — every real plan
  change today does two full passes over the user's active keys, each issuing
  one HTTP `PATCH` per key to api-management.
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

**The key-generation code path Rails would need for a migration already
exists**, gated behind `Rails.env.test?`: `ApiKey#generate_key_locally`
(`api_key.rb:60-68`) generates and hashes a key entirely in Rails without
calling Cloudflare. Today it's only exercised in tests; production always routes
through `request_key_from_server` → api-management (`api_key.rb:29-58`).
Promoting the local path to production and dropping the Cloudflare round trip is
a small, well-precedented change, not new architecture.

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
colliding with `golang-migrate`'s default `schema_migrations` table (which Go's
migration runner uses with no `x-migrations-table` override, confirmed against
`apps/api/platform/config/config.go` and `infra/docker/.env.example`).
**`docs/core/rails-app.md:207-213`'s claim of "separate schema" is imprecise** —
it's a renamed table in the same `public` schema, not a Postgres
schema/namespace — but the practical collision is handled. No code change needed
here; only a documentation correction.

**Real, minor schema issues found regardless of the Workers question:**
`api_keys.key_hash` — the actual authentication lookup column — has no unique
index (only the non-secret `key_prefix` is indexed); `words` has no index on the
`word` lookup column; the FK gaps noted above in the Rails section.

### Redis

**Verdict on Question 4 (can Redis replace KV for high-write operational state):
Yes, and the pattern to do it already exists in this codebase.**

Today Redis is single-instance (`redis:7-alpine`, no clustering/replication,
default RDB persistence, **no `maxmemory` configured** — unbounded growth is
possible, not just unlikely), shared across Go's response caches, Go's one
counter feature, Rails' `Rails.cache`, Rails' `Rack::Attack` throttles, and
unnamespaced Sidekiq queues. No collisions observed (key prefixes differ), but
there's no structural isolation (e.g., no separate Redis DB index per consumer).

The one place Redis already does exactly the job a rate-limiter/quota-tracker
needs — `services/technology/counter`'s Lua-scripted atomic increment +
dirty-set-swap + batched Postgres upsert — is a genuinely good design:
idempotent (writes absolute values, safe to retry), crash-safe between the
increment and the next flush (worst case, up to 60s of increments are lost only
if Redis itself crashes without a persisted snapshot — a real, bounded window,
not an unbounded one), and already load-tested implicitly by being in
production. This is the concrete evidence behind the recommendation to
generalize this pattern rather than design a new one.

**What Redis is not being asked to do today, and would need to do in the
proposed design:** synchronous atomic rate-limit checks in the request hot path
(`INCR` + `EXPIRE`, O(1), fixes the current KV get-then-put race by
construction) and a cache of `api_keys.key_hash → {user_id, plan}` invalidated
on revoke (mirrors the existing `geocode`/`crypto`/`exchange` TTL-cache pattern
already in `apps/api/services`).

---

## Data Ownership

| Data                                                     | Source of truth today                                                                                | Proposed source of truth                                                                                                             |
| -------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------ |
| API key existence/plan/revocation                        | Cloudflare KV (fast path) + D1 `api_keys` (audit) + Postgres `api_keys` (Rails' view) — three copies | PostgreSQL `api_keys` (Rails-owned), Redis as a read-through cache only                                                              |
| Rate-limit counters                                      | Cloudflare KV, non-atomic, no persistent record                                                      | Redis, atomic `INCR`+`EXPIRE`, ephemeral by design (no Postgres mirror needed — these are enforcement-only, not billing)             |
| Usage/billing ledger                                     | Cloudflare D1 `credit_usage` → (3 min lag) → Postgres `usage_logs`                                   | Redis counters (per-user/per-key) → (batched, same pattern as `counters` today) → Postgres `usage_logs` directly, no edge SQLite hop |
| Subscriptions/plans                                      | Postgres `subscriptions` (Rails), mirrored into KV                                                   | Postgres `subscriptions` only; Go reads it (directly or via a Redis-cached view) instead of a separate KV copy                       |
| Business/reference data (advice, quotes, BIN data, etc.) | Postgres, Go-owned                                                                                   | unchanged                                                                                                                            |
| Response caches (geocode, crypto, FX)                    | Redis, Go-owned                                                                                      | unchanged                                                                                                                            |

---

## Request Lifecycle

**Current (authenticated API call, quota-cache-warm case):** Client → Cloudflare
edge → auth-gateway Worker → KV read (key) → KV read+write (rate limit, racy) →
KV read (quota cache) → fetch to Go over the internet (10s timeout) → Go (no
auth of its own, trusts `X-Backend-Secret`) → Postgres/Redis for business logic
→ response → Worker records usage: async D1 write + KV read/write. **7–8 storage
operations, 2 network hops between edge and origin**, before business logic even
runs.

**Proposed:** Client → Cloudflare (proxy only, no Worker) → Go: Redis `GET` (key
cache) → Redis `INCR`+`EXPIRE` (rate limit, atomic) → business logic → Redis
`INCR` (usage, same Lua pattern as `counters` today) → response. **~3 atomic
Redis ops, one process, one network hop.** Usage flush to Postgres happens
out-of-band on the existing 60s batched-sync cadence, not per-request.

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

**Proposed:** Go validates the customer's key directly — hash it, look up
`api_keys.key_hash` (needs the missing unique index added first), cache the
`{user_id, plan}` result in Redis with a short TTL (mirrors the existing
`geocode`/`crypto` cache pattern) so steady-state traffic doesn't hit Postgres
per request. Revocation invalidates the Redis cache entry explicitly (an active
`DEL` on revoke, not just a TTL wait) — this is _faster_ to propagate than
today's system, where a revoked key still works against a stale 60s KV cache in
some paths and depends on a successful `sync_revocation_to_cloudflare` HTTP call
that can silently fail.

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
implementation. The fix is purely mechanical: `Redis.INCR(rl:{key}:{minute})` +
`EXPIRE` on first increment is atomic by construction (single Redis command
round trip, no read-modify-write race), which the current KV approach cannot
offer without a Durable Object or similar (out of scope). This directly resolves
the concurrency-correctness gap found in `rate-limit.ts:38-53`.

---

## Usage/Billing Architecture

**Current:** auth-gateway writes one row to D1 `credit_usage` per request
(async, retried, non-blocking) → api-management exports it paginated → Rails'
`SyncD1UsageJob` polls every 3 minutes and bulk-inserts into Postgres
`usage_logs`, deduplicated by a unique index on
`(api_key_id, used_at, endpoint)`. This is **at-least-once with idempotent
dedup** — a sound design, just with an unnecessary edge-SQLite hop and up to 3
minutes of lag before usage is visible in the dashboard.

**Can this move to Go + Redis + Postgres without losing billing accuracy? Yes**,
using the same idempotent-upsert principle already proven in
`services/technology/counter`: increment a Redis counter per `(api_key_id, day)`
or similar granularity atomically at request time, periodically flush _absolute_
values (not deltas) to Postgres via upsert, exactly as `counter/sync_worker.go`
does today. Because the flush writes absolute values, a crash mid-cycle is
self-healing on the next tick — no double-counting, no silent loss beyond the
same bounded ~60s window the `counters` feature already accepts in production.
This directly answers the audit's core billing question: **exactly-once is not
required and is not what the current system provides either** (D1's design is
at-least-once + dedup-by-unique-index); the proposed design preserves that same
guarantee with fewer moving parts and less lag.

---

## Security Findings (ranked)

1. **Non-atomic rate-limit counter**
   (`apps/workers/auth-gateway/src/rate-limit.ts:38-53`) — get-then-put race
   allows a customer to exceed their per-minute limit under concurrent requests.
   Low severity today (soft abuse-prevention, not a hard security boundary), but
   worth fixing regardless of the migration decision since it's cheap to fix in
   place (Cloudflare KV has no atomic increment; this would require either a
   Durable Object or moving the check to Redis).
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
4. **`api_keys.key_hash` has no unique index** in Postgres — not currently
   exploitable (nothing queries by it in production yet, since validation
   happens in KV), but must be added before any migration that makes this column
   the live authentication lookup.
5. No hardcoded secrets, no committed `.env` files, no SQL injection patterns,
   no CORS misconfiguration, and no plaintext API keys at rest were found
   anywhere in the repo during this audit — the baseline hygiene is good.

---

## Performance Findings

- **KV/D1 storage-operation count scales linearly and un-batchably with RPS** —
  6–8 edge storage operations per request today, most of them synchronous to the
  response. At the product hypothesis's 10k–100k RPS targets, this is 60k–800k
  KV operations/sec plus 10k–100k D1 writes/sec; D1 (SQLite-per-database) is not
  designed for that write concurrency, which matches the "critical problem"
  already observed with KV write pressure described in the audit brief.
- **Proposed design cuts this to ~3 atomic Redis ops/request, zero synchronous
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
  design should explicitly decide fail-open-vs-fail-closed for a Redis outage
  (recommendation: fail-open on the rate-limit/usage-counter check, fail-closed
  on the api-key-existence check — a Redis outage should not let unauthenticated
  traffic through, but it also shouldn't hard-block all authenticated traffic
  over a soft rate-limit check).
- **Single, unclustered Redis instance with no `maxmemory` configured** —
  currently low-risk (only response caches + one counter + Rails
  cache/throttles), becomes a single point of failure for the entire
  request-serving path once auth/rate-limit/usage move onto it. Needs a
  `maxmemory`+eviction policy and a documented "what happens if Redis is down"
  story before cutover — today, an unbounded Redis a real risk to plan for at
  the RPS targets in scope.
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

| Item                                                                                   | Evidence                                       | Why unused                               | Safe to remove?                              | Dependencies                                                                                    |
| -------------------------------------------------------------------------------------- | ---------------------------------------------- | ---------------------------------------- | -------------------------------------------- | ----------------------------------------------------------------------------------------------- |
| `apps/workers/auth-gateway`, `apps/workers/api-management`, `apps/workers/shared`      | Full audit above                               | Logic moves to Go/Rails/Postgres/Redis   | Yes, after staged migration (see below)      | Cloudflare KV namespace, D1 database, Wrangler secrets, `cd.yml`-adjacent manual deploy scripts |
| Cloudflare KV namespace `7cc847da...`                                                  | `wrangler.toml` in both Workers                | No consumer once Workers are removed     | Yes, after migration                         | —                                                                                               |
| Cloudflare D1 database `requiem-usage`                                                 | `wrangler.toml` in both Workers, `schema.sql`  | No consumer once Workers are removed     | Yes, after migration                         | `D1SyncService`, `SyncD1UsageJob`                                                               |
| `apps/dashboard/app/services/d1_sync_service.rb`                                       | Full file read                                 | Only pulls from D1                       | Yes, after migration                         | `SyncD1UsageJob`                                                                                |
| `apps/dashboard/app/services/cloudflare/api_management_service.rb`                     | Full file read                                 | Only talks to api-management             | Yes, after migration                         | `ApiKey`, `Subscription` callbacks                                                              |
| `apps/dashboard/app/jobs/sync_d1_usage_job.rb`                                         | Full file read                                 | Sole purpose is D1→Postgres sync         | Yes, after migration                         | Sidekiq schedule entry                                                                          |
| `CLOUDFLARE_ACCOUNT_ID`, `CLOUDFLARE_KV_NAMESPACE_ID`, `CLOUDFLARE_API_TOKEN` env vars | `.env.example:107-110`                         | Only used for KV management from Rails   | Yes, after migration                         | —                                                                                               |
| `apps/workers/shared/src/types.ts` `ApiKeyManagementRequest`/`Response`                | grep, zero usages                              | Never consumed                           | Yes, now (independent of migration)          | none                                                                                            |
| `apps/workers/shared/src/config.ts` `PLAN_NAMES`                                       | grep, zero usages                              | Never consumed                           | Yes, now                                     | none                                                                                            |
| `auth-gateway/src/rate-limit.ts` `getPlanLimits()`                                     | grep, only used in its own test                | Superseded by `getRequestLimitMessage()` | Yes, now                                     | none                                                                                            |
| Rails `subscriptions.stripe_customer_id/stripe_subscription_id/credit_limit`           | grep, zero usages in `app/`                    | Pre-LemonSqueezy remnants                | Yes, now (independent of migration)          | none confirmed — verify with a migration before dropping                                        |
| Rails `solid_cache_entries` table                                                      | `cache_store` always overridden away from it   | Rails-8 default never used               | Yes, now                                     | none                                                                                            |
| Rails `app/policies/` (Pundit scaffolding)                                             | Zero concrete policies, zero `authorize` calls | Dead scaffolding                         | Yes, now, or: finish wiring it up — pick one | none                                                                                            |
| `api-management`'s `/analytics/*` endpoints                                            | No caller found in Rails                       | Built, tested, unconsumed                | Needs confirmation before deleting           | none found                                                                                      |

---

## Architectural Smells (ranked)

1. **Four uncoordinated stores for one concern (auth/rate-limit/usage).**
   Evidence: Section "Current Architecture" table above. Impact: propagation
   lag, silent divergence risk, and the KV write-pressure problem already
   observed in production. Root cause: auth/rate-limit logic was pushed to the
   edge for latency reasons without a plan for keeping edge state and origin
   state reconciled. Recommendation: consolidate per the Proposed Architecture.
   Priority: P0.
2. **Non-atomic rate limiting under the exact conditions it exists to prevent
   (bursty, concurrent, abusive traffic).** Evidence: `rate-limit.ts:38-53`.
   Impact: correctness gap in abuse prevention scales with concurrency, i.e.,
   gets worse exactly as traffic grows. Root cause: KV has no atomic increment
   primitive. Recommendation: Redis `INCR`+`EXPIRE`. Priority: P0 (folds into
   the broader migration; a standalone fix isn't worth building twice).
3. **Duplicate side-effecting HTTP call on every subscription plan change.**
   Evidence: `subscription.rb:26` +
   `lemonsqueezy_controller.rb:108,160,186,210`. Impact: 2x unnecessary load on
   api-management, 2x unnecessary KV writes, harder-to-reason-about failure
   modes (which of the two calls failed?). Root cause: the callback was likely
   added after the explicit call already existed, or vice versa, without
   noticing the overlap. Recommendation: remove one (keep the callback, since it
   also covers non-webhook plan changes like admin promotions). Priority: P1,
   independent of the migration.
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

## Migration Plan

This assumes the architecture decision below is accepted. Each phase preserves
current behavior until the next phase explicitly changes it — no "delete Worker,
deploy Go, hope" step exists in this plan.

**Phase 0 — Fix the blocking Go gaps (prerequisite, not migration-specific).**
Add graceful shutdown (`signal.NotifyContext` + `server.Shutdown`), structured
logging (`log/slog` at minimum), and explicit `pgxpool`/`go-redis` pool sizing
to `apps/api`. Add a unique index on `api_keys.key_hash` in Rails. These are
needed regardless of the migration and should land first so later phases build
on a healthier base.

**Phase 1 — Introduce Go-side API-key authentication, dual-running alongside the
Worker.** Add a new Go middleware that validates `requiems-api-key` against
Postgres `api_keys.key_hash` (Redis-cached), applied _only_ to a feature-flagged
subset of routes or behind a header the Worker sets to say "I already
authenticated this." Worker authentication remains the production gate; Go's new
check runs in shadow/log-only mode to compare decisions against the Worker's.

**Phase 2 — Introduce Redis-based rate limiting and usage counting in shadow
mode.** Implement atomic `INCR`+`EXPIRE` rate limiting and generalize the
existing `counter` sync-worker pattern to a per-`api_key`/per-`user` usage
counter, flushing to a new (or the existing) `usage_logs`-compatible table. Run
in parallel with the Worker's KV/D1 path, log-only, no enforcement yet.

**Phase 3 — Dual-operation comparison.** Compare Go's shadow
auth/rate-limit/usage decisions against the Worker's live decisions for a
sustained period (recommend at least one full billing cycle, so quota-boundary
behavior is exercised). Reconcile any discrepancies found — this is where subtle
bugs (e.g., timezone handling in monthly quota resets, edge cases in the
multiplier logic for batch endpoints) will surface before they can affect
customers.

**Phase 4 — Move key-generation ownership to Rails.** Promote
`ApiKey#generate_key_locally` to the production path, removing the
`request_key_from_server` call to api-management for _new_ key creation only
(existing keys are unaffected — they already validate the same way in Postgres).
Revocation/update still round-trip to Cloudflare KV at this stage so the
Worker's auth path stays correct for existing traffic.

**Phase 5 — Move a small percentage of production traffic directly to Go (bypass
the Worker).** Requires Cloudflare routing changes (weighted routing or a canary
DNS/route split) so a fraction of `api.requiems.xyz` traffic hits Go directly
instead of through auth-gateway. Monitor error rates, latency, and
rate-limit/quota accuracy against the Worker-routed control group.

**Phase 6 — Full cutover.** Route all traffic directly to Go. Cloudflare config
changes from "Worker in front" to "proxy-only" (orange-cloud DNS, WAF/DDoS rules
retained, Worker route removed). `X-Backend-Secret` and the Caddy `@authorized`
gate (`infra/caddy/Caddyfile:32-47`) are removed since Go is now directly the
internet-facing origin behind Cloudflare's proxy — replace with
Cloudflare-specific origin-verification (e.g., Cloudflare's own origin pull
certificates or an Authenticated Origin Pulls setup) if origin-lockdown is still
desired at the Caddy layer.

**Phase 7 — Remove D1/KV synchronization.** Delete `SyncD1UsageJob`,
`D1SyncService`, `Cloudflare::ApiManagementService`, the
`sync_to_cloudflare`/`sync_revocation_to_cloudflare` callbacks on
`Subscription`/`ApiKey`. Remove `CLOUDFLARE_*` env vars.

**Phase 8 — Remove Workers.** Delete `apps/workers/auth-gateway`,
`apps/workers/api-management`, `apps/workers/shared`. Delete the Cloudflare KV
namespace and D1 database. Remove their entries from `.dependabot.yml`,
`docker-compose.dev.yml`, `_worker-ci.yml` and its callers in `ci.yml`. Update
`docs/core/{architecture,auth-gateway,api-management}.md` (see Documentation
Audit below) and remove `tests/integration/src/suites/gateway.test.ts`'s
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
- Add structured logging (`log/slog`) to `apps/api`. _Independently doable,
  touches many files, medium complexity, low risk._
- Explicit `pgxpool`/`go-redis` pool config in
  `apps/api/platform/{db,reqredis}`. _Independently doable, small, needs
  load-test validation._
- Add unique index on `apps/dashboard` `api_keys.key_hash`. _Independently
  doable, one migration, needs a live-DB check that no duplicate hashes already
  exist first._
- Fix the non-atomic rate-limit race — either patch in place (accept added KV
  complexity) or fold directly into Phase 2 of the migration (recommended: fold
  in, don't build the KV fix twice).

**P1 — High priority**

- Fix duplicate Cloudflare sync on subscription plan change
  (`subscription.rb:26` vs `lemonsqueezy_controller.rb`). _Independent of
  migration, small, needs a test added first since neither path is currently
  tested._
- Fix MRR calculation (populate `subscriptions.plan` from LemonSqueezy webhook
  payload, backfill). _Independent of migration, needs care around backfill
  correctness — recommend writing a one-off Rails runner script, verifying
  against LemonSqueezy's own dashboard totals before trusting the backfill._
- Migration Phases 0–3 (Go-side auth/rate-limit/usage in shadow mode). _Depends
  on P0 items above. Medium-high complexity — this is the architectural core of
  the audit's recommendation. Needs new Go tests for the auth/rate-limit/usage
  paths and a comparison/reconciliation script for Phase 3._
- Configure `trusted_proxies` in Rails (and equivalent in Go post-migration) for
  `CF-Connecting-IP` handling. _Small, but must land before/at Phase 6 (full
  cutover), since Cloudflare's routing to origin changes then._

**P2 — Medium priority**

- Migration Phases 4–7 (key-generation ownership, traffic cutover, D1/KV sync
  removal). _Depends on P1 migration phases. Requires coordinated Cloudflare
  routing changes and careful monitoring during Phase 5's canary period._
- Add missing FKs on `credit_adjustments.admin_user_id`,
  `audit_logs.admin_user_id`, `abuse_reports.resolved_by_id`. _Independent,
  small, needs a data-cleanup pass first if any orphaned values exist._
- Add index on Go's `words.word` column if lookup-by-word is actually a query
  pattern (verify against `apps/api/services/text/words/service.go` query shape
  first). _Independent, trivial once confirmed necessary._
- Add test coverage for `Cloudflare::ApiManagementService` and the three
  untested Sidekiq jobs — _only worth doing if Phase 4+ of the migration is
  delayed; otherwise this code is being deleted, don't invest in testing code on
  its way out._

**P3 — Nice to have**

- Migration Phase 8 (delete Workers, KV, D1, related CI/docs/deps).
- Drop dead Rails columns (`stripe_*`, `credit_limit`) and the unused
  `solid_cache_entries` table.
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
- **Cloudflare's role changes from "does auth" to "does nothing but proxy,"
  which changes the DDoS/abuse posture** — today, a volumetric or
  credential-stuffing attack against a specific API key is rejected at the edge
  (KV lookup, cheap) before ever reaching Go; post-migration, the same rejection
  happens in Go behind Cloudflare's generic WAF/rate-limiting (not the
  plan-aware, per-key logic being migrated). Cloudflare's WAF rules should be
  reviewed/tightened as part of Phase 6, not assumed to be equivalent.
- **Redis becomes a single point of failure for the entire request-serving
  path**, not just for a handful of internal caches. The "what happens if Redis
  is down" question needs an explicit, tested answer (recommended: fail-open on
  rate-limit/usage checks, fail-closed on key-existence checks) before Phase 6,
  not discovered live during an incident.
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

---

## Final Recommendation

Proceed with consolidating auth, rate limiting, and usage tracking into the Go
backend + Redis + PostgreSQL, retiring `auth-gateway`, `api-management`,
`apps/workers/shared`, Cloudflare KV, and Cloudflare D1. The original hypothesis
is correct, and the audit found concrete, current evidence supporting it beyond
the original concern about KV write pressure: a real correctness bug in the KV
rate limiter, thin test coverage on the exact code path every request depends
on, four uncoordinated stores for one concern, and — most importantly — a
working, production-proven template for the Redis+Postgres pattern already
living in the codebase (`services/technology/counter`). This is not a bet on an
unproven architecture; it's an extension of one already running.

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
