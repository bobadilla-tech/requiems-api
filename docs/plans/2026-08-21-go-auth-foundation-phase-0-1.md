# Go Auth Foundation — Phase 0 & 1

Prerequisite hardening of `apps/api` plus a Go-native, Postgres/Redis-backed
API-key auth path, built directly as the enforcing mechanism rather than run in
shadow mode alongside the Cloudflare Worker — because the product has no users
or production traffic yet, the multi-week comparison/canary machinery that full
architecture audit's migration plan calls for for a live cutover buys nothing
here and is deliberately dropped.

## Context

`docs/audits/2026-08-21-architecture-audit.md` is a full-repo audit that found
four uncoordinated stateful stores (Cloudflare KV, Cloudflare D1, Postgres,
Redis) serving one concern — authenticated, rate-limited, metered API requests —
and recommends consolidating onto Go + Redis + Postgres, retiring
`apps/workers/auth-gateway`, `apps/workers/api-management`, and
`apps/workers/shared`. That audit's "Migration Plan" section lays out 8 phases
built for a live-traffic cutover: Phases 1–3 run the new Go auth path in
shadow-only mode for weeks, comparing its decisions against the Worker's live
decisions, specifically so a bug in the new path never affects a real customer's
request or bill mid-transition.

That caution is the right call when customers exist. They don't yet — this
project has no users and no production traffic. There is nothing to protect
mid-transition, no notice period owed, no backward-compatibility surface to
preserve. This plan takes the audit's Phase 0 and Phase 1 scope and executes it
directly: build the Go auth path correctly, test it for correctness, then make
it the actual enforcing path. No shadow-logging period, no dual-gating flag, no
sampled-divergence tracking against the Worker. The Worker (`auth-gateway`)
keeps running during this plan (it isn't being deleted here — that's a later
phase, out of scope), but nothing in this plan depends on keeping its decisions
in sync with Go's.

Today's auth reality, per the audit's Authentication Architecture section: a
customer's API key is validated by the Worker as a flat Cloudflare KV lookup
keyed by the literal plaintext key (`key:{apiKey}`,
`apps/workers/auth-gateway/src/middleware/api-key-auth.ts:50`) — Postgres holds
a bcrypt hash (`api_keys.key_hash`) and a `key_prefix` column, but nothing in
the live request path ever reads or verifies against them. Because bcrypt hashes
are salted, there is no way to look up a row by `bcrypt(presented_key)` — the
only viable Postgres-backed flow is candidate-then-verify: extract `key_prefix`
(deterministic, cheap) from the presented key,
`SELECT ... WHERE key_prefix = ?`, then `bcrypt.compare` against each
candidate's `key_hash`. `api_keys.key_prefix` currently has only a trigram/GIN
index (built for fuzzy admin search) — no plain btree for efficient exact-match
lookup, which this auth path needs.

The Go backend itself is not yet fit to be an internet-facing tier independent
of the auth question: no graceful shutdown (`apps/api/main.go` calls
`server.ListenAndServe()` with no `signal.NotifyContext`/`server.Shutdown()`),
no structured logging (scattered `log.Printf`/`log.Println` across ~23 files,
Sentry only on uncaught errors), and unconfigured connection pools (`pgxpool`
and `go-redis` both on library defaults in `apps/api/platform/{db,reqredis}`).
These are masked today because the Worker absorbs client-facing disruption
during a Go redeploy — but they're real gaps worth closing regardless of the
auth work, and the auth work is a reasonable forcing function to close them now.

Separately, the documented dev curl example and `scripts/seed-dev.ts` both seed
keys shaped like `rq_free_000001` (`docs/core/auth-gateway.md:160`,
`scripts/seed-dev.ts:22-30`), but the live validator requires
`^requiem_[0-9a-zA-Z]{24}$`
(`apps/workers/shared/src/api-key-generator.ts:15-19`). A seeded dev key is
rejected with 401 today — this needs fixing before the new Go auth path can be
exercised manually in dev at all.

## Approach

### Phase 0 — Go hardening + prerequisite fixes

Five independent items, no ordering dependency between them, each shippable as
its own small PR.

1. **Graceful shutdown** — `apps/api/main.go`. Wrap the server start in
   `signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)`,
   run `ListenAndServe` in a goroutine, and on context cancellation call
   `server.Shutdown(ctx)` with a bounded timeout so in-flight requests drain
   before the process exits. Also close the `pgxpool` and Redis client after
   shutdown returns.

2. **Structured logging** — `log/slog`, wired into Go's HTTP middleware chain
   (wherever the existing request-logging middleware lives today, alongside the
   router setup), emitting JSON to stdout with at minimum: request ID, method,
   route, status code, latency. Sentry's role is unchanged — it stays the
   uncaught-exception capture mechanism, not the request-logging mechanism. No
   log aggregation service is being added here; stdout JSON is picked up by the
   existing Docker/Kamal log collection as-is.

3. **Connection pool sizing** — `apps/api/platform/db/db.go` and
   `apps/api/platform/reqredis/redis.go`. Set explicit
   `pgxpool.Config.MaxConns`/`MinConns` and `go-redis`'s `PoolSize` instead of
   library defaults. There's no production traffic to size against yet (the
   audit's Open Question 5 doesn't apply pre-launch) — pick conservative fixed
   values consistent with Postgres's configured `max_connections` and a single
   Go replica (e.g. `MaxConns` in the 15–25 range, leaving headroom for Rails'
   own pool against the same Postgres instance), and leave a short code comment
   noting these are placeholder values to revisit once real traffic exists.
   Don't block this item on measuring traffic that doesn't exist yet.

4. **btree index on `api_keys.key_prefix`** — new Rails migration in
   `apps/dashboard/db/migrate/`. Plain `add_index :api_keys, :key_prefix`
   (btree, not the existing trigram/GIN index used for fuzzy admin search). This
   is what Phase 1's candidate-lookup query needs for an efficient exact-match
   `WHERE key_prefix = ?`.

5. **Fix dev-seed key format, and make the seed script produce a key Postgres
   actually knows about.** The seed script is
   `apps/workers/auth-gateway/scripts/seed-dev.ts` (not a repo-root
   `scripts/seed-dev.ts`), and today it only seeds Cloudflare KV/D1 via
   `wrangler kv key put` / `wrangler d1 execute` — D1's `api_keys` table
   (`apps/workers/auth-gateway/schema.sql`) has no `key_hash` column at all,
   it's just `key_prefix, user_id, plan, active, timestamps`. A key seeded this
   way has no corresponding Rails/Postgres `api_keys` row (Rails requires
   `key_hash` to be present), so it's invisible to Phase 1's Postgres-backed
   lookup no matter what prefix format it uses. This item is two changes, not
   one: (a) fix the key-shape mismatch — the seed script and
   `docs/core/auth-gateway.md:160`'s curl example currently use
   `rq_free_000001`-shaped keys against a live validator that requires
   `^requiem_[0-9a-zA-Z]{24}$`; (b) add a Postgres-side seed path — a Rails seed
   task (e.g. `rails db:seed` addition or a `rake` task) that produces a real
   `key_hash`+`key_prefix` row for local dev, independent of the Worker/KV/D1
   seed script. **Don't just call `ApiKey.create!` naively** — `ApiKey`'s
   `before_validation :request_key_from_server`
   (`apps/dashboard/app/models/api_key.rb`) only takes the local-generation
   branch (`generate_key_locally`) under `Rails.env.test?`; in any other env,
   including plain `development`, it calls out to
   `Cloudflare::ApiManagementService` over HTTP, which defeats the "no Worker
   dependency" point of this seed task and requires live Worker credentials to
   even run. Instead, either (i) call `generate_key_locally`/ `ApiKeyGenerator`
   directly and pass the resulting `key_prefix`/`key_hash` into `ApiKey.create!`
   — `request_key_from_server` early-returns once `key_prefix` is already
   present, so this skips the network branch — or (ii) run the seed task with
   `RAILS_ENV=test` semantics for this call specifically. Pick one and say so in
   the actual PR; don't leave the seed task making a live Cloudflare call by
   accident. Both the KV-seeded key (for exercising the Worker path) and the
   Postgres-seeded key (for exercising Phase 1's Go path) should exist in dev
   after this item, and they don't need to be the same key.

**Phase 0 exit criteria:** all five items merged; `apps/api` restarts cleanly on
SIGTERM with no dropped in-flight requests (verify with a manual slow-endpoint +
`docker stop`/`kamal` redeploy check); a Postgres-seeded dev key (from item 5b)
exists and is queryable by `key_prefix`.

### Phase 1 — Go-native API-key authentication (built as enforcing, not shadow)

New Go middleware, added to the existing `apps/api/platform/middleware` package
— which already holds exactly this kind of cross-cutting auth/request middleware
(`auth.go`'s `BackendSecretAuth`, `requestid.go`, `urlparam.go`, each with a
matching `_test.go`). Extend this package rather than inventing a new one.

1. **Candidate-then-verify lookup.** On each request carrying a
   `requiems-api-key` header (confirmed as the live header name in both
   `apps/workers/auth-gateway/src/middleware/api-key-auth.ts:39` and the OpenAPI
   spec): extract `key_prefix` from the presented key (same deterministic
   slicing logic the key-generator uses),
   `SELECT * FROM api_keys WHERE key_prefix = ?` (uses Phase 0's index), then
   `bcrypt.compare(presented_key, candidate.key_hash)` against each returned row
   (expect exactly one in practice, but don't assume DB-enforced uniqueness —
   iterate the candidate set). No plaintext key is ever persisted or logged.

2. **Redis-backed verification cache**, keyed by `key_prefix` — not by a hash of
   the raw key. Rails discards the raw key after creation, so a raw-key-derived
   cache key would have no way to be invalidated later; a `key_prefix`-keyed
   cache can be invalidated because Postgres always has `key_prefix` on hand at
   revoke time. Cache value: `{user_id, plan,
   revoked: bool}` (or
   equivalent), short TTL (mirror the existing `geocode`/`crypto`/`exchange`
   TTL-cache pattern already in `apps/api/services` for the TTL convention used
   elsewhere in this codebase). This avoids paying a Postgres round-trip and
   bcrypt's deliberately-slow compare on every request once a key's been seen
   once.

3. **Revocation invalidation.** On revoke, Rails must issue
   `DEL apikey:{key_prefix}` against the **same, unnamespaced Redis keyspace Go
   writes to** — this is not as simple as "Rails already talks to Redis for
   Rack::Attack/Sidekiq, so reuse that." `Rails.cache` (what `Rack::Attack`
   uses) is a `:redis_cache_store` configured with `namespace: "rails_cache"`
   (`apps/dashboard/config/environments/{development,production}.rb`), so
   `Rails.cache.delete(key_prefix)` would actually target
   `rails_cache:{key_prefix}` in Redis, silently missing the unnamespaced key
   Go's raw `go-redis` client wrote — a permanent no-op, not a best-effort
   fallback. Sidekiq's Redis connection, by contrast, is raw/unnamespaced and
   would work. Use a raw Redis connection (`Redis.new(url: ENV["REDIS_URL"])` or
   equivalent, matching Go's key format exactly, not `Rails.cache`) for this
   specific `DEL` call. If the `DEL` fails, a stale cache entry could keep
   serving a revoked key until TTL expiry — at zero scale this is an accepted,
   explicitly-documented gap for now (short TTL is the mitigation), not
   something to build a durable-retry-queue for in this phase. Revisit if/when
   this actually matters.

4. **Fail-closed on key-existence.** If Redis is unreachable and Postgres must
   be hit directly and also fails, or if the key genuinely doesn't resolve,
   reject the request (401/403) — never let an auth-store outage fall through to
   authenticated-as-nobody or authenticated-as-anybody. This is a correctness
   property, not migration-timeline-dependent caution, so it's not part of the
   "skip because no users" trim.

5. **Test coverage**, following this repo's existing pattern of hitting a real
   Postgres/Redis rather than mocking them:
   - Valid key → resolves correct `{user_id, plan}`.
   - Revoked key → rejected, both immediately (cache miss path) and while a
     stale cache entry exists pre-TTL-expiry (documents the accepted gap from
     item 3, doesn't silently pass).
   - **Revoke a key that has a warm cache entry, then re-request it before TTL
     expiry → must be rejected.** This is the one that actually catches a
     namespace/keyspace mismatch in the `DEL` call from item 3 (e.g. Rails
     deleting `rails_cache:{key_prefix}` while Go reads plain `{key_prefix}`) —
     without this exact test, that bug ships silently and only surfaces as
     "revoked keys keep working for up to the cache TTL" in production.
   - Two rows sharing a `key_prefix` (collision candidate set) → correct row
     selected via bcrypt compare, no false-accept on the wrong candidate.
   - Redis down, Postgres reachable → falls through to direct DB lookup and
     still enforces correctly (or: fails closed, if that's the chosen
     degraded-mode behavior — decide explicitly in code, not by omission).
   - Malformed/garbage key → rejected without panicking or leaking whether the
     prefix matched anything.

6. **Wire it up as real, request-blocking middleware — with an explicit note on
   what traffic it actually gates today.** Once the middleware's own test suite
   is green, mount it on `apps/api`'s router as enforcing (not shadow/log-only)
   middleware. At the time Phase 1 shipped, the Worker preserved
   `requiems-api-key` on its trusted hop to Go, and Go verified it against
   Postgres. Direct requests used the same header contract, so this middleware
   enforced both Worker-proxied and local traffic. Making this middleware the
   enforcing path for real customer traffic requires Phase 5/6 of the audit's
   migration plan (direct traffic cutover) — don't claim this phase alone flips
   production auth over, because it doesn't.

**Phase 1 exit criteria:** middleware test suite green (all cases in item 5); a
Postgres-seeded dev key (from Phase 0 item 5b) authenticates through the new Go
path end-to-end when the request is sent directly to Go (bypassing the Worker);
a revoked key is rejected within its cache TTL window; killing the local Redis
container mid-session doesn't silently let unauthenticated traffic through.

### Explicitly out of scope for this plan

- Rate limiting and usage counting (audit's Phase 2) — separate plan, next.
- Any Worker code changes, retirement, or traffic cutover (audit's Phases 4–8).
- The audit's multi-week shadow-comparison, canary-percentage routing, and
  sampled-divergence tracking — dropped per the no-users rationale above, not
  deferred.
- Rails-side key-generation ownership changes (audit's Phase 4) — Rails
  continues generating keys via the existing Cloudflare-backed path for now;
  this plan only adds a second, Go-native way to _validate_ a key that already
  exists in Postgres.

## Final notes

**Shipped, Phase 0:**

- Graceful shutdown in `apps/api/main.go`: `signal.NotifyContext` +
  `server.Shutdown` with a 15s drain timeout, pool/Redis client closed after.
  Manually verified with a running binary + `SIGTERM` mid-request — in-flight
  requests completed, new connections were refused once shutdown began, clean
  exit logged.
- Structured logging via `log/slog`: `platform/middleware/logging.go`
  (`RequestLogger`) emits one JSON line per request (request ID, method, route
  pattern, status, latency), mounted in `app.go` alongside `RequestID`.
  `main.go`'s own startup/shutdown logs were also moved to `slog` since it was a
  small, directly-related change. **Scope call:** the ~23 files with scattered
  `log.Printf`/`log.Println` elsewhere in `apps/api/services` were _not_ touched
  — the plan's acceptance criterion was "wired into Go's HTTP middleware," not a
  repo-wide logging migration, and rewriting 23 unrelated files would have been
  scope creep beyond what this plan asked for.
- Pool sizing: `pgxpool.Config.MaxConns=20/MinConns=5` in `platform/db/db.go`,
  `go-redis`'s `PoolSize=20` in `platform/reqredis/redis.go`, both with a code
  comment flagging them as placeholders to revisit under real traffic.
- Btree index:
  `db/migrate/20260821000000_add_btree_index_to_api_keys_key_prefix.rb`
  (`index_api_keys_on_key_prefix_btree`), migrated and `schema.rb` regenerated.
- Dev-seed fix: `apps/workers/auth-gateway/scripts/seed-dev.ts` and
  `docs/core/auth-gateway.md`'s curl example now use
  `requiem_[plan]00000000000000000001`-shaped keys (valid against
  `^requiem_[0-9a-zA-Z]{24}$`, unique 12-char prefixes for D1). Postgres-side
  seeding was added to `db/seeds.rb`: calls `ApiKeyGenerator` directly (option
  (i) from the plan) and passes `key_prefix`/`key_hash` straight into
  `ApiKey.create!`, so `request_key_from_server`'s before_validation
  early-returns and no Cloudflare HTTP call happens in development. Verified
  idempotent (`bin/rails db:seed` twice — second run detects the existing key
  and skips). The KV-seeded and Postgres-seeded keys are intentionally different
  keys, per the plan.

**Shipped, Phase 1:**

- `apps/api/platform/middleware/apikeyauth.go`: `APIKeyAuth`,
  candidate-then-verify by `key_prefix` (btree-indexed lookup +
  `bcrypt.CompareHashAndPassword` against each candidate), Redis-cached by
  `key_prefix` under the `apikey:{prefix}` key (documented in-code as the exact
  contract Rails' revocation `DEL` must match). Fails closed whenever Postgres
  can't affirmatively resolve the key (including when Redis is _also_ down);
  falls through to a direct Postgres lookup when only Redis is unavailable. Full
  test suite in `apikeyauth_test.go` covers every case item 5 lists, including
  the namespace-mismatch regression test (warm cache + revoke + raw `DEL` +
  re-request before TTL) — all run against real Postgres/Redis, not mocks, per
  this repo's convention.
- Rails side: `app/services/go_auth_cache.rb` (`GoAuthCache.invalidate`) uses a
  raw `Redis.new` connection, never `Rails.cache`, wired into `ApiKey` via new
  `after_update`/`after_destroy` callbacks alongside the existing
  Cloudflare-sync ones. Tested against real Redis
  (`test/services/go_auth_cache_test.rb`, plus two new `ApiKey` model tests) —
  including a regression test asserting invalidation does _not_ touch a
  `rails_cache:`-namespaced key.
- Mounted in `app.go`: `apiKeyAuth.Middleware()` is chained in the _same_ `/v1`
  protected group as `BackendSecretAuth`, both required (AND-composed), with a
  30s cache TTL constant.
- Manually verified end-to-end against a real Postgres+Redis with the
  Rails-generated schema: seeded key → 200 through Go directly; wrong key/no key
  → 401; Redis killed mid-session + correct key → still 200 (falls through to
  Postgres); Redis killed + wrong key → still 401 (no fail-open). All four match
  the Phase 1 exit criteria.

**Worker-proxy status at Phase 1 implementation time:** the Worker preserved
`requiems-api-key` on its trusted hop to Go, and Go re-verified the complete
credential against Postgres. The Worker's edge validation and Go's backend
validation therefore both remained in the request path; local direct-to-Go
requests used the same header contract. Later retirement/cutover work is tracked
in the Phase 7 plan.

**Security status of the former cache-hit concern:** Redis entries under
`apikey:{key_prefix}` are candidates only. Every cache hit re-runs bcrypt
against the complete presented key; a mismatch falls through to the full
Postgres candidate query, which also handles shared-prefix keys. Prefix-only and
altered-suffix presentations therefore fail closed. The shipped per-key rate
limiter and monthly quota run after authentication; Redis failures fail open for
the soft rate limiter and fail closed with `503` for quota reservation.

**Follow-ups:**

- Pool sizes (`MaxConns`/`PoolSize`) and the 30s API-key cache TTL are all
  placeholders picked without real traffic to measure against — revisit once
  there's load to look at.
