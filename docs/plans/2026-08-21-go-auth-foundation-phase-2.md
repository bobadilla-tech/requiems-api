# Go Auth Foundation — Phase 2: Rate Limiting & Usage Tracking

Continuation of `docs/plans/2026-08-21-go-auth-foundation-phase-0-1.md`, which
shipped Go-native, Postgres/Redis-backed API-key authentication as the enforcing
path for any request that reaches Go directly. This plan is the audit's
Migration Phase 2 (`docs/audits/2026-08-21-architecture-audit.md`, "Migration
Plan" section): atomic Redis rate limiting and per-user usage/quota tracking.
Same rationale as Phase 0/1 for skipping the audit's shadow-mode machinery — no
users, no production traffic, nothing to protect mid-transition — so this is
built directly as the enforcing mechanism, tested for correctness, and shipped.
The Worker (`auth-gateway`) keeps running unmodified; retiring it is still out
of scope (audit Phases 4–8).

## Context

Phase 1 authenticates a request and attaches an `APIKeyPrincipal{UserID, Plan}`
to the request context (`apps/api/platform/middleware/apikeyauth.go`). Nothing
downstream of that yet enforces a rate limit or tracks usage — every request
that clears `APIKeyAuth` today runs unbounded. `APIKeyPrincipal` also doesn't
carry `APIKeyID` or the user's billing-cycle start, both of which this phase's
middleware needs (rate limiting is per-key, quota is per-user-per-cycle, per the
legacy Worker's own key scheme in
`apps/workers/auth-gateway/src/middleware/api-key-auth.ts`) — extending
`lookupDB`'s query and the cached JSON shape to carry both is part of this
phase, not a prerequisite fixed elsewhere.

**Three duplicated copies of plan-tier values exist today**, not two as an
earlier audit draft implied: `apps/workers/shared/src/config.ts` +
`apps/workers/shared/plan-limits.json` (Worker), and
`apps/dashboard/app/lib/plan_config.rb`'s hardcoded `PLANS` hash (Rails). The
audit recommends a Postgres `plans` table as the single new source of truth.
This phase adds that table and points Go at it. **Explicitly out of scope:**
migrating `PlanConfig` or the Worker's config off their hardcoded copies — the
Worker keeps running unchanged, and Rails' existing hardcoded values are already
correct, just duplicated. Leaving three copies in sync manually for one more
phase is an accepted, temporary cost; collapsing to one true copy everywhere is
Phase 8 territory. The new `plans` table just gives Go somewhere real to read
from instead of adding a fourth hardcoded copy.

**Design deviation from the audit's literal suggestion, reasoned explicitly here
so it isn't second-guessed as an oversight:** the audit's Usage/Billing
Architecture section suggests generalizing `counter/sync_worker.go`'s
dirty-set-swap flush mechanism for the aggregate/quota side. That mechanism
exists to solve a specific problem — the `counter` feature has no independent
durable per-event record, so Redis is the only copy of the truth until it's
flushed. **That problem doesn't exist here.** This phase writes each request's
usage row directly and synchronously to Postgres `usage_logs` (the row-level
ledger Rails' dashboard already reads, using the existing
`(api_key_id, used_at, endpoint)` unique dedup index from
`20260316170000_add_composite_index_to_usage_logs.rb` for idempotency) — this
directly answers the audit's Open Question 7 by choosing option (a) from the
Usage/Billing Architecture section explicitly. Once that write is durable and
synchronous, a separate flush-to-Postgres path for the aggregate counter would
be redundant: the aggregate is only needed as a **fast, atomic check**
in-request, not as a second source of truth. So the quota counter uses the same
tiny Lua bootstrap-and-increment shape already proven in
`counter/redis_mutations.go`'s `incrementWithBootstrapScript`, but with no
flush-out step: on a cache miss it bootstraps its baseline from
`SELECT COALESCE(SUM(credits_used),0) FROM usage_logs WHERE user_id = ? AND
used_at >= ?`
— **scoped by `user_id`, not `api_key_id`**, matching the
`usage:{user_id}:{cycle_start_unix}` key and the quota-is-per-user (not per-key)
design stated above; a user with more than one active API key must have all of
them summed into one baseline, or the bootstrap silently under-counts (mirrors
the Worker's old quota-cache-miss D1 read, which also summed at the user level,
not per-key) — and sets a TTL equal to time-remaining-in-the-billing-cycle, so a
new cycle just misses and rebootstraps from zero automatically — no explicit
monthly-reset job needed. This is simpler than generalizing the sync worker, not
a shortcut: `daily_usage_summaries` keeps being populated exactly as it is today
by Rails' existing `AggregateDailyUsageJob` reading `usage_logs`, completely
unaffected by anything in this phase.

**Accepted, intentional behavior worth stating explicitly:** a Redis quota
reservation is required for each request. Redis loss or an atomic reservation
failure returns `503`; the implementation does not rebuild a counter through a
non-atomic Postgres `SUM` fallback and cannot silently forgive or double-count
checked requests during an outage.

Rate limiting has no equivalent row-level ledger and doesn't need one — the
audit's Data Ownership table already says these counters are "ephemeral by
design, no Postgres mirror needed." Plain Lua `INCR`+`EXPIRE`, no bootstrap, no
flush.

## Approach

### 1. `plans` table (Rails migration + seed)

New table: `id` (string PK, e.g. `"free"`/`"developer"`/...), `request_limit`
(integer, nullable = unlimited, for `enterprise`), `rate_limit_per_minute`
(integer, nullable = unlimited). Seed from `PlanConfig::PLANS`'
`requests_per_month`/`rate_limit_per_minute` values (`free`/`developer`/
`business`/`professional`) plus an `enterprise` row with both columns null.
Rails' `PlanConfig` and the Worker's `config.ts` are **not** migrated to read
from this table in this phase (see Context) — it exists so Go has one real
source, seeded once, and manually kept in sync with the other two copies for now
(flag this manual-sync obligation in the PR description).

### 2. Extend `APIKeyPrincipal` and the Redis-cached blob

`apps/api/platform/middleware/apikeyauth.go`: add `APIKeyID int64` and
`CurrentPeriodStart time.Time` (from `subscriptions.current_period_start`,
falling back to the key's `created_at` if the user has no subscription row —
matches free-tier users, who may never have a `subscriptions` row) to both
`APIKeyPrincipal` and `cachedAPIKey`. Update `lookupDB`'s query (join adds
`api_keys.id`, `subscriptions.current_period_start`) and its two call sites.
Existing Phase 1 tests that assert on the cached JSON shape need updating
alongside this, not left to bit-rot.

### 3. Rate limiter middleware

New file, same `platform/middleware` package (extends it, per the Phase 0/1
convention): `ratelimit.go`. Lua script, single round trip:

```lua
local count = redis.call("INCR", KEYS[1])
if count == 1 then redis.call("EXPIRE", KEYS[1], 60) end
return count
```

Key: `ratelimit:{api_key_id}:{unix_minute}`. Limit resolved from the `plans`
table (Go-side, in-process cache with an explicit **60s TTL** — same convention,
and same concrete value, as the `geocode`/`crypto`/`exchange` TTL-cache pattern
already in `apps/api/services` per the audit's Authentication Architecture note;
don't leave this as a vague "short TTL" decision for the next session to make
from scratch) keyed by `principal.Plan`. `enterprise` (or any plan with a null
`rate_limit_per_minute`) skips the check entirely. Exceeding the limit returns
`429` with a `Retry-After` header.

**Known, unmitigated gap in this phase, stated explicitly rather than left
implicit:** this rate limiter is keyed by `api_key_id`, which only exists
_after_ `APIKeyAuth` has already resolved an identity. It does nothing to
throttle the brute-force exposure Phase 0/1's own notes flagged in the auth
cache (a warm `key_prefix` cache entry trusts any presented suffix, with no
re-verification against the full key) — that exposure is a single-shot success
on the first correct-prefix guess, not a multi-attempt brute force against a
fixed identity, so per-key rate limiting structurally can't bound it. Phase
0/1's note that "rate limiting... would bound the brute-force risk" does not
hold for this design as built; the real fix (shortening the cache TTL further,
re-verifying a suffix fragment on cache hit, or lengthening the prefix) is out
of scope here and should be picked up as its own follow-up, not assumed solved
by this phase.

**Fail-open on Redis unavailability** — a Lua-eval error against Redis lets the
request through uncounted, logged at `warn`. This matches the audit's explicit
guidance (Reliability Findings): rate limiting is soft abuse-prevention, not a
security boundary, and fail-closed here would turn a Redis blip into a full
outage for something lower-stakes than auth. Mounted after `APIKeyAuth` in the
same `/v1` group, reading `APIKeyPrincipalFromContext`.

### 4. Usage/quota middleware

New file: `usage.go`, same package. Two independent pieces, both running per
request, in this order:

**a. Quota check (before the handler runs).** Lua bootstrap-and-increment script
(same shape as `counter/redis_mutations.go`'s `incrementWithBootstrapScript`,
reimplemented here rather than imported — `counter` is a leaf feature package
backing a public product API and shouldn't gain an auth-domain dependency, per
the audit's own note that `counter`'s data model is namespace→total and not
meant to be reused as-is). Key: `usage:{user_id}:{cycle_start_unix}`. On
bootstrap, baseline query as described in Context. The Lua script increments
unconditionally, then the post-increment value is compared against the plan's
`request_limit` (null = unlimited, skip the check entirely). If over the limit,
the request is rejected with `429` and the reservation is rolled back — a
rejected request consumes no quota and produces no `usage_logs` row. The
reservation uses the authoritative static route cost before dispatch; dynamic
batch cost is reconciled from the backend response after an admitted request.

**b. Row-level write (after the handler runs, on response).** Synchronous
`INSERT INTO usage_logs (user_id, api_key_id, endpoint, credits_used,
request_method, status_code, response_time_ms, used_at, usage_date, created_at,
updated_at) VALUES (...) ON CONFLICT (api_key_id, used_at, endpoint) DO
NOTHING`
— **column-list conflict target, not `ON CONFLICT ON CONSTRAINT`**: the existing
unique index (`index_usage_logs_on_api_key_id_and_used_at_and_endpoint` in
current `schema.rb`) is a plain unique index, not a named Postgres constraint in
`pg_constraint`, so `ON CONFLICT ON CONSTRAINT <name>` would raise
`constraint "..." does not exist` at runtime regardless of which name is used;
Postgres's column-list form matches against a unique index directly and is the
correct syntax here. If this write itself fails (transient Postgres error, not a
conflict — a real error), log at `error` and drop the row rather than retrying
synchronously post-response; this is the one place in this phase where a failure
is silently swallowed rather than fail-closed, and it's called out here
explicitly as an accepted gap (the response has already been sent; there's
nothing left to reject) rather than left implicit. `credits_used` reads the
endpoint's multiplier — backfill the existing `httpx.UsageCounter` interface
(`platform/httpx/handler.go:33-36`, already defined) is **not** in scope for
this phase (220 endpoints is real, separate work); confirm at implementation
time exactly how many endpoints implement it today (the number in this doc,
"one," was not independently re-verified and may be stale — a wrong count here
is a live quota-accuracy bug, not a neutral scope note), then default every
other endpoint to `credits_used = 1` and read `UsageCounter.UsageCount()` only
where a response already implements it, mirroring today's actual coverage.
Widening `UsageCounter` adoption is a natural, explicitly-flagged follow-up, not
something to rush into this phase's scope.

**Explicit, reasoned trade-off — this deviates from the audit's Performance
Findings, not just its Usage/Billing Architecture section.** The audit's
Performance Findings section separately states the target design should reach
"~4 atomic Redis ops/request... zero synchronous Postgres/D1 ops" at the
10k–100k RPS the product hypothesis targets. Step (b) above is a synchronous
Postgres write on every request, which conflicts with that target directly —
this conflict is real, not reconciled by the Usage/Billing Architecture
section's option (a) reasoning alone (that section only argues row-level detail
can't be reconstructed from an aggregate; it doesn't address write concurrency).
Accepted here because there is no real traffic yet — a synchronous per-request
insert is fine at dev/staging volume — but this is exactly the kind of decision
Open Question 5 says shouldn't be made blind: **flag explicitly, don't silently
ship past it.** If this remains synchronous once real RPS exists, revisit as a
batched/async write (a bounded, buffered writer flushing periodically, distinct
from the Redis quota counter) before that becomes the bottleneck the audit
warned about.

**Fail-closed on quota-check failure:** Redis quota reservation is the atomic
counter. If Redis errors _or times out_ on the quota Lua script, reject the
request with `503`; do not fall through to a bare Postgres `SUM`, because a
sum-then-write sequence is not an atomic reservation and concurrent requests
could all pass against the same total. A transactional reservation table or a
locked Postgres counter can replace this policy later, but until one exists
Redis-unavailable quota requests must fail closed. Both the rate-limiter's
`INCR` and this quota script must be called with a short, explicit timeout (e.g.
a bounded `context.Context`, on the order of tens of milliseconds) — **a Redis
instance that's alive but slow under memory pressure is a distinct failure mode
from a clean connection error**, and without an explicit timeout it isn't caught
by the same error-handling path at all; it just makes every request hang on a
slow/degraded Redis instead of falling through. Any Redis error — connection
refused, timeout, or an `OOM command not allowed` write rejection once
`maxmemory` is actually hit under the `noeviction` policy from section 5 below —
must be handled identically and must reject the request; don't special-case on
error string/type. A bare Postgres `SUM` fallback is intentionally not used
because it cannot reserve quota atomically under concurrency. If Postgres is
also unreachable, the request is likewise rejected. Log at `error`, not `warn`,
since this path means a real reconciliation gap if it happens repeatedly.
**Known, accepted risk left unaddressed by this design:** if Redis degrades (not
fails) under sustained load, requests receive `503` until Redis recovers. A
transactional Postgres reservation can replace this policy later; it is not
silently assumed to be safe here.

### 5. Redis key-space isolation

Audit's Reliability Findings flag unbounded Redis with no `maxmemory`/eviction
policy as a growing risk as more state (auth cache, rate limits, quota counters)
piles onto the same instance as disposable response caches and Sidekiq's queues.
This phase adds three new categories of Redis state (`apikey:*` from Phase 1,
`ratelimit:*`, `usage:*` from this phase) that must never be evicted under
memory pressure, sharing an instance with `geocode:*`/`crypto:*`/`exchange:*`,
which are fine to evict. Local docker-compose and Kamal Redis config: set
`maxmemory-policy noeviction` on the whole instance (simplest fix available
without standing up a second Redis deployment target, which is out of scope for
this phase) and set an explicit `maxmemory` sized with headroom over current
usage — **placeholder value, same caveat Phase 0/1 gave its pool sizes:** picked
without real traffic to measure against, and needs to be revisited once there's
load to look at; say so in the config comment, don't present it as a tuned
number. Document in-repo (a short comment in the compose/Kamal Redis config, not
a new doc file) that `noeviction` means the disposable response caches also
never get LRU-evicted under this config — acceptable since they're small and
TTL'd, but worth stating so a future contributor doesn't assume `allkeys-lru`
protection that isn't there. **Also document, explicitly, what `noeviction`
actually does at capacity:** Redis doesn't silently drop data — it starts
rejecting _writes_ with an `OOM command not allowed` error while reads keep
succeeding. This is a distinct failure mode from "Redis unreachable" and must be
handled by the same fallback paths in sections 3 and 4 (see the
timeout/error-handling note there) — it is not automatically covered just
because `noeviction` is configured; the application code has to actually treat
it the same way it treats a connection error. Splitting into a separate logical
Redis DB or a dedicated instance is real, larger infra work — explicitly
deferred, noted as a follow-up.

### 6. Test coverage

Following the repo's real-Postgres/real-Redis convention (Phase 1's
`apikeyauth_test.go`):

- Rate limiter: concurrent requests against the same key don't undercount (the
  exact race the audit flags in the legacy KV limiter, `rate-limit.ts:38-53`); a
  key from a previous minute window doesn't leak into the next; Redis down →
  request proceeds (fail-open), logged; Redis reachable but exceeding the call
  timeout also falls open, not hangs.
- Quota: bootstrap-from-Postgres on cold cache produces the correct baseline
  **summed across all of a user's API keys**, not just the triggering one (seed
  `usage_logs` rows under two different keys for the same user, confirm the
  counter starts from their combined sum, not just one key's); crossing the
  limit rejects with 429 and no further Postgres write for that request; a new
  billing cycle (via TTL expiry or an explicit `current_period_start` change)
  resets to zero, not carrying over the previous cycle's count; Redis down →
  fails closed with 503, as does both Redis and Postgres being down; a simulated
  `OOM command not allowed` write rejection (e.g. a test Redis instance with
  `maxmemory` set near-zero) is handled identically to a connection error, not
  left as an unhandled 500.
- Row-level write: two requests that would produce the same
  `(api_key_id, used_at, endpoint)` — simulate via a fixed clock or explicit
  `used_at` — dedup correctly via `ON CONFLICT DO NOTHING`, no error surfaced to
  the client either way. **Known, accepted gap, not something to fix in this
  phase:** this dedup key was designed for the old at-least-once D1→Postgres
  sync, where retries could produce genuine duplicates of the _same_ request.
  Here, two genuinely distinct rapid requests to the same endpoint within the
  same stored time granularity would collide and the second is silently dropped
  from `usage_logs` — while the Redis quota counter (which increments
  unconditionally per checked request) still counts both. This means the
  row-level ledger can under-report relative to the quota counter under rapid
  same-second traffic to one endpoint. Not addressed here; flag as a follow-up
  if it turns out to matter (e.g. widening `used_at`'s precision or dropping the
  dedup constraint now that the write path is
  synchronous-and-idempotent-by-construction rather than retried).
- `plans` table: a plan with a null `request_limit`/`rate_limit_per_minute`
  (enterprise) is never rate-limited or quota-rejected.

## Exit criteria

- `plans` table migrated and seeded; values match `PlanConfig::PLANS`.
- Rate limiter and quota middleware mounted, enforcing, in the same `/v1` group
  as `APIKeyAuth` (same "what traffic this actually gates" caveat from Phase 1
  applies unchanged: only direct-to-Go traffic, not Worker-proxied traffic, per
  that plan's own flagged deviation).
- Full test suite (item 6) green.
- Manually verified end-to-end: a seeded dev key rate-limited past its
  per-minute cap returns 429 with `Retry-After`; a key pushed past its monthly
  quota (seed `usage_logs` rows near the limit) returns 429 on the next request;
  `usage_logs` gets a real row per successful request, visible via
  `bin/rails console` against the dev Postgres.
- Redis `maxmemory-policy` set to `noeviction` with an explicit `maxmemory` in
  both docker-compose and Kamal configs.

## Explicitly out of scope for this plan

- Backfilling `httpx.UsageCounter` across the ~220 endpoints that don't
  implement it yet (separate, larger plan).
- Migrating `PlanConfig`/the Worker's `config.ts` off their hardcoded plan
  values onto the new `plans` table (Phase 8 territory — see Context).
- Any Worker code changes, retirement, or traffic cutover (audit Phases 4–8).
- A dedicated/second Redis instance or logical-DB split (audit Reliability
  Findings' fuller recommendation) — this phase only sets `noeviction` on the
  existing shared instance.
- The Phase 1 deviation already shipped (Worker→Go proxy path currently 401s
  since `APIKeyAuth` is mounted in the same group as `BackendSecretAuth`) is
  unchanged by this plan and not revisited here.

## Final notes

**Shipped:**

- `plans` table: `apps/dashboard/db/migrate/20260821010000_create_plans.rb`,
  seeded inside the migration's `up` (not `db/seeds.rb`) — production's boot
  command runs `db:prepare`, which runs pending migrations but never `db:seed`,
  so seeding in `db/seeds.rb` would have left the table empty in every
  environment that matters most. `id` is a string PK; `request_limit` and
  `rate_limit_per_minute` are nullable integers, seeded with the five rows the
  plan doc specifies (four from `PlanConfig::PLANS`, `enterprise` with both
  columns null). `schema.rb` regenerated.
- `apps/api/platform/middleware/apikeyauth.go`: `APIKeyPrincipal` and
  `cachedAPIKey` both extended with `APIKeyID int64` and
  `CurrentPeriodStart time.Time`; `lookupDB`'s query now also selects
  `api_keys.id` and
  `COALESCE(subscriptions.current_period_start,
  api_keys.created_at)`.
  Existing Phase 1 tests updated in place (new
  `subscriptions.current_period_start` column with a `DEFAULT NOW()` in the test
  table, plus new assertions on the two added fields) rather than left to
  bit-rot, per the plan's own instruction.
- `apps/api/platform/middleware/plancache.go`: `PlanCache`, a shared in-process,
  60s-TTL cache over the `plans` table, constructed once in `app.go` and passed
  to both `RateLimiter` and `UsageQuota` — one cache, one Postgres read path,
  not two independent implementations. Unknown plan names remain unlimited, but
  a Postgres error is returned separately: the rate limiter fails open as soft
  abuse prevention, while quota uses stale limits when available and returns
  `503` when no stale limits exist.
- `apps/api/platform/middleware/ratelimit.go`: `RateLimiter`, the exact Lua
  script from the plan doc (`INCR` + first-hit `EXPIRE 60`), keyed
  `ratelimit:{api_key_id}:{unix_minute}`. Fails open on any Redis error
  (including a timeout), logged at `warn`; `429` with `Retry-After` set to
  seconds-remaining-in-the-current-minute on exceeding the limit.
- `apps/api/platform/middleware/usage.go`: `UsageQuota`. Quota check reuses
  `counter/redis_mutations.go`'s EXISTS-then-INCR / SET-then-INCR shape,
  reimplemented (not imported) per the plan's own reasoning, with an added
  `EXPIRE` on bootstrap so a key just misses at the next cycle boundary — no
  reset job. Baseline sums across all of a user's API keys. Cycle boundary is
  computed by `cycleStart()`, mirroring the legacy Worker's `getResetTime`
  (`apps/workers/auth-gateway/src/requests.ts`): the most recent occurrence of
  the billing anchor's day-of-month, at midnight UTC, at or before now — this is
  what lets a static anchor (a free-tier key's `created_at`, which never
  changes) still roll into a new monthly cycle automatically, which a naive "use
  `CurrentPeriodStart` as the literal cycle key" approach would not have done.
  Redis error/timeout fails closed with `503`, logged at `error`; there is no
  non-atomic Postgres `SUM` fallback. Row-level write is a synchronous
  `INSERT ... ON
  CONFLICT (api_key_id, used_at, endpoint) DO NOTHING`,
  column-list form as specified; `credits_used` reads `X-Usage-Count`
  (defaulting to 1) — the same header the Worker's `proxy.ts` already reads for
  the identical purpose, confirming the doc's "one endpoint implements
  `UsageCounter`" guess (`services/systems/data_integrity/input_validate_batch`)
  was correct, not stale.
- Mounted in `app.go`: `rateLimiter.Middleware()` then
  `usageQuota.Middleware()`, both after `apiKeyAuth.Middleware()` in the same
  protected `/v1` group — same "what traffic this actually gates" caveat from
  Phase 1 applies unchanged.
- Redis config: shared Redis uses `maxmemory-policy noeviction`, but production
  now requires an operator-supplied `REDIS_MAXMEMORY` chosen after measuring
  protected auth/quota state, Sidekiq, and bounded TTL caches with headroom;
  development retains the bounded example value.
- Tests: `ratelimit_test.go` and `usage_test.go` cover every case section 6
  lists (concurrent-increment correctness, minute-window isolation, fail-open on
  Redis-down and Redis-timeout, cross-key-sum bootstrap, cycle rollover
  excluding prior-cycle usage, 429-rolls-back-and-skips-the-row-write, repeated
  rejected batch reservations, Redis-down fail-closed, both-down fail-closed,
  simulated `OOM command not allowed` rejected identically to a connection error
  via `CONFIG SET maxmemory 1`, row-level dedup via `ON CONFLICT`, and
  enterprise/null-limit plans never triggering either check) — all against real
  Postgres/Redis, following the repo's existing convention. `usage.go`'s
  row-insert was split into a small `insertUsageRow` helper specifically so the
  dedup test could force an exact `used_at` collision instead of racing
  `time.Now()`'s microsecond precision; this is the one production-code shape
  driven by testability rather than the plan doc's own text.
- `apps/api/app/app_test.go`'s `seedAPIKeyFixture` updated to also create
  self-contained `subscriptions`/`plans`/`usage_logs` tables and seed a `free`
  plan row with null limits. This wasn't optional cleanup: the protected route
  group now runs `APIKeyAuth`'s `LEFT JOIN subscriptions` query plus the new
  rate-limit/quota middleware on every request, so the existing end-to-end
  fixture needed the same tables to keep passing — the same "existing tests need
  updating alongside this" instruction the plan doc gives for the
  `apikeyauth_test.go` cached-JSON-shape tests applies here too, just for a file
  the doc didn't happen to name.

**Bug found and fixed while implementing, not called out in the plan doc:**
`go-redis` v9 ignores a per-call `context.WithTimeout` entirely unless the
client's `ContextTimeoutEnabled` option is set — without it, a command blocks on
the client's own (much longer) default `ReadTimeout` regardless of the context
passed to `Script.Run`. This would have silently defeated the plan's explicit
"tens of milliseconds" timeout requirement for exactly the "alive but slow"
Redis failure mode section 4 calls out as the reason a timeout is needed at all
— discovered via `TestRateLimiter_RedisTimeoutFailsOpenWithoutHanging` actually
hanging for ~300ms instead of failing open at 50ms. Fixed by setting
`ContextTimeoutEnabled = true` in `platform/reqredis/redis.go`'s `Connect`
(applies to every caller of the shared client, not just this phase's middleware
— a strictly-more-correct default, not a behavior change for any caller that
doesn't set a context deadline) and the equivalent test-client constructor in
`apikeyauth_test.go`.

**Manually verified end-to-end**, against the dev docker-compose stack
(`db`+`redis`) with a Rails-migrated schema and a `db:seed`-generated dev key,
hitting Go directly on its own port (bypassing the Worker, matching Phase 1's
own verification method):

- 29 requests within a minute → all `200`; the 30th → `429` with
  `Retry-After: 24`.
- Redis's `usage:{user_id}:{cycle}` counter set to 499 → next request → `200`
  (the 500th, at the limit); the request after that → `429
  quota_exceeded`.
- `usage_logs` held exactly one row per served (non-429) request — 30 rows for
  30 served requests, confirmed via direct SQL (equivalent to the
  `bin/rails console` check the exit criteria names); none of the rate-limited
  or quota-rejected requests produced a row.

**Deviations from the doc's literal text, both reasoned, neither silent:**

1. The doc's rate-limiter section describes the in-process plan-limit cache's
   60s TTL as "the same convention, and same concrete value, as the
   geocode/crypto/exchange TTL-cache pattern" — in fact those three use
   Redis-backed caches with TTLs of 24h, 1h, and 5min respectively, not 60s and
   not in-process. The doc's own text for _this_ cache is unambiguous ("Go-side,
   in-process cache with an explicit 60s TTL"), so that's what was built; the
   "same convention" framing just doesn't hold up against the actual code in
   `apps/api/services`, worth noting so nobody goes looking for a 60s Redis
   cache over there.
2. `cycleStart()`'s day-of-month rollover logic isn't spelled out in the plan
   doc at the level of "recompute the boundary each request from the anchor plus
   now" vs. "use `CurrentPeriodStart` as a static cycle key" — the doc only says
   the key is `usage:{user_id}:{cycle_start_unix}` and that a new cycle "just
   misses and rebootstraps." A static reading of `CurrentPeriodStart` as the
   literal cycle key would never roll over for free-tier users (whose anchor is
   `api_keys.created_at`, which never changes, since they have no subscription
   webhook updating it) — the SUM's `WHERE used_at >= cycle_start` would keep
   including every month since signup forever. `cycleStart()` closes that gap by
   mirroring the legacy Worker's own `getResetTime` day-of-month logic, so this
   isn't new behavior for existing accounts, just the same monthly-rollover
   semantics reimplemented in Go.

**Known, accepted gaps carried forward as-is (already reasoned about in the plan
doc's text, not revisited here):** the rate limiter's inability to bound the
auth-cache prefix-guessing exposure; the usage-logs row-level dedup collision
under rapid same-second traffic to one endpoint; the three manually-synced
copies of plan-tier values.

**Follow-ups (new, surfaced by this phase):**

- The 256mb `maxmemory` value and the plan-cache's 60s TTL are placeholders
  picked without real traffic to measure against, same caveat as every other
  placeholder constant from Phase 0/1 — revisit once there's load to look at.
- `PlanCache`'s fail-open-on-Postgres-error behavior (see above) was a judgment
  call filling a gap the plan doc left implicit. Worth confirming explicitly
  once there's a real incident to reason from, rather than leaving it as an
  assumption baked into the code.
