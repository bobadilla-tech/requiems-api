# Go Auth Foundation — Phase 5: Correctness Cleanup & Playground Fix

Continuation of `docs/plans/2026-08-21-go-auth-foundation-phase-0-1.md`,
`-phase-2.md`, and `-phase-3-4.md`. This plan is scoped to what a single focused
session can safely land end-to-end: independent Rails/Go correctness bugs (the
never-executed Phase 4 backlog), plus the follow-up bugs the Phase 3 session
surfaced mid-implementation but explicitly did not fix. Nothing here touches
Cloudflare, DNS, or the Worker deployment.

**Note on what this actually picks up:** `-phase-3-4.md` has a "Final
notes/Shipped" section only for its own Phase 3 items 1–5 and 8 — Phase 3 items
6–7 were explicitly not started, and **Phase 4 (MRR fix, `trusted_proxies`, doc
sweep, dead-code sweep) was written up in that doc but never executed at all**,
not partially. So this plan is simultaneously (a) finally executing that
never-run Phase 4 backlog and (b) fixing the three bugs Phase 3's session
surfaced mid-implementation. Items 1, 3 (Rails half), 5, and 6 below are that
leftover Phase 4 scope, re-verified against the current tree rather than assumed
still accurate.

**If this session runs short, the minimum viable subset is items 2 and 3** (the
playground fix and the trust-boundary fix) — both are pure code+test, no
external-service dependency, no production-data mutation. Item 1 (MRR) has its
own external dependency and risk profile (see its section below) and can safely
slip to a follow-up session on its own; items 4–6 are independent, low-risk
cleanup that can also slip individually without blocking anything else here.
Ship each item as its own PR, per this series' existing convention.

## Context

**What's already shipped (verified against the current tree, not re-derived from
the plan docs):**

- `apps/api/app/app.go` — `APIKeyAuth` is the sole gate on `/v1`;
  `BackendSecretAuth` and `cfg.BackendSecret` are gone from Go entirely.
- `apps/dashboard/app/models/api_key.rb` — key generation is Rails-local in
  every environment, `requiem_`-prefixed, collision-checked. Cloudflare-sync
  callbacks on `ApiKey`/`Subscription` are deleted; `GoAuthCache.invalidate`
  still fires on revoke/destroy.
- Rate limiting and usage/quota middleware (Phase 2) are mounted and enforcing
  on direct-to-Go traffic.

**What's still true today, confirmed by re-reading the actual files (not just
the plan docs' notes) immediately before writing this plan:**

- `apps/workers/{auth-gateway,api-management,shared}` still exist in the repo
  and are still deployed.
- `infra/caddy/Caddyfile:37` still gates `internal.requiems.xyz` on
  `X-Backend-Secret` — so real Worker-proxied traffic still can't reach Go's
  `/v1` routes at all (the Worker strips `requiems-api-key` before proxying, and
  Go's `APIKeyAuth` requires it). This is the Phase 3 item 6/7 gap (Cloudflare
  origin lockdown + Worker retirement), unstarted, and it needs live Cloudflare
  dashboard/VPS access with human sign-off on each irreversible step —
  **explicitly out of scope for this plan**, not because it's unimportant, but
  because it isn't safely executable inside a single autonomous coding session
  without that access in the room. Everything below is scoped to be genuinely
  completable without it.
- `apps/dashboard/app/services/api_proxy_service.rb:32` still sends only
  `X-Backend-Secret`, never `requiems-api-key` — confirmed by direct read. Since
  Go no longer even checks `X-Backend-Secret` (Phase 3 item 5 removed
  `BackendSecretAuth`), this call 401s at `APIKeyAuth` on every request. The
  public playground (`ApiProxyController`) and `ToolDemosController`'s
  server-side demo forms — both funnel through this one service — have been
  fully non-functional against Go since Phase 3 landed.
- `apps/dashboard/app/lib/app_config.rb:41,165` — `playground_api_key` is read
  from `PLAYGROUND_API_KEY`, defaults to the stale, invalid
  `rq_test_playground_demo_key` literal, and — confirmed by grep — is never
  referenced anywhere outside its own definition and `attr_reader`. It reads as
  an abandoned half-implementation of exactly the fix the previous bug needs,
  not two unrelated findings.
- `apps/dashboard/db/schema.rb` confirms `subscriptions.plan` (the billing-cycle
  column) is still present and, confirmed by grep across `apps/dashboard/app/`,
  still never written by any webhook handler.
  `AnalyticsRevenueService#calculate_mrr`/`#revenue_trend`
  (`analytics_revenue_service.rb:48,76`) still key off
  `sub.plan&.to_sym || :monthly` — every yearly subscriber's MRR contribution is
  still silently computed as monthly pricing.
- No `trusted_proxies` configuration exists anywhere in `apps/dashboard/config/`
  (confirmed, zero matches). `ApiProxyController#create`
  (`api_proxy_controller.rb:23`) still trusts
  `request.headers["CF-Connecting-IP"]` outright.
- The same trust gap exists on the Go side, and it's live product surface, not
  hypothetical: `services/networking/ip/info/transport_http.go`,
  `services/networking/ip/asn/transport_http.go`, and
  `services/systems/global_data/timezone_ip/transport_http.go` all resolve the
  caller's IP via `callerIP()`, which reads `X-Forwarded-For` (first hop,
  unvalidated) with no check that the request actually came through
  Cloudflare/Caddy rather than being crafted directly.
- `apps/dashboard/db/schema.rb` still has `subscriptions.stripe_customer_id`,
  `stripe_subscription_id`, `credit_limit` (zero usages in `app/`, confirmed)
  and the `solid_cache_entries` table (present, unused — cache store is always
  overridden to Redis/null). `credit_adjustments.admin_user_id`,
  `audit_logs.admin_user_id`, `abuse_reports.resolved_by_id` have no FK, unlike
  the structurally identical `subscriptions.promoted_by_id`, which does.
  `app/policies/` contains only `application_policy.rb` — zero concrete
  policies, zero `authorize` calls anywhere in `app/controllers/` (confirmed by
  grep) — dead Pundit scaffolding, exactly as the audit found it.
- `apps/api/app/app_test.go`'s `TestApp_Handler` connects with whatever
  `DATABASE_URL` is set to and has `seedAPIKeyFixture` create
  `api_keys`/`subscriptions`/`plans`/`usage_logs` tables directly against it —
  no separate test database or schema. Confirmed this collides with Rails' real
  migrations if run against the shared dev DB out of order (hit directly during
  the Phase 3 session; documented as a follow-up there, not fixed).

## Approach

### 1. Fix the MRR/revenue billing-cycle bug

**Risk note, stated up front rather than left implicit:** this item mutates
production financial data and depends on external access (LemonSqueezy's
dashboard/API) the same way Phase 3 items 6–7 depend on Cloudflare dashboard
access — the difference is this plan doesn't know whether that access is
available in the room. Confirm LemonSqueezy dashboard/API access is actually
available before starting this item; if it isn't, treat this item the same way
items 6–7 are treated (defer to a session where it is), rather than guessing at
payload fields or backfill values without a way to verify them.

`subscriptions.plan` needs to be populated from LemonSqueezy's webhook payload
in `webhooks/lemonsqueezy_controller.rb`'s
`handle_subscription_created`/`handle_subscription_updated`. **Confirm the exact
payload field at implementation time** — the audit never pinned one down, and
this plan doesn't either; inspect a real webhook payload (replay one from
LemonSqueezy's dashboard, or check `variant_id`/`product_id` naming conventions
already used elsewhere in this controller for the monthly/yearly split) before
writing the assignment.

Then:

- Write a one-off Rails runner script to backfill existing `subscriptions` rows.
  If LemonSqueezy's API can still return each subscription's original billing
  interval, use it; if webhook payload history isn't retrievable for some rows,
  flag those explicitly for manual reconciliation rather than guessing a
  default.
- Verify the backfilled MRR total against LemonSqueezy's own dashboard figure
  before trusting it, and **record both numbers (before/after backfill vs.
  LemonSqueezy's own total) in the PR description** — "spot checked" isn't a
  verifiable exit criterion on its own; the actual comparison needs to be a
  durable artifact someone can review later.
- Add test coverage: `AnalyticsRevenueService#calculate_mrr` and
  `#revenue_trend` currently have no test asserting yearly-vs-monthly pricing
  divergence (verify this gap exists, then close it) — this is the bug that
  should have caught the original defect.

### 2. Fix the playground/demo-form proxy (`api_proxy_service.rb`)

This closes bug #1 and bug #2 from the Phase 3/4 session's follow-up notes as
one coherent fix, not two:

- Provision a real, dedicated system API key for the playground. **Do not mirror
  Phase 0's dev-key seeding pattern literally** — that seed (`db/seeds.rb`) is
  gated `if Rails.env.development?` and never runs in production, since
  production's boot command is `bin/rails db:prepare`, which runs pending
  migrations but never `db:seed` (this is the exact trap Phase 2's own Final
  Notes flagged and worked around by seeding its `plans` table inside a
  migration instead — see `-phase-2.md`'s "Shipped" section). The playground is
  public, production-facing surface, so a `Rails.env.development?`-gated seed
  would leave this bug looking fixed locally while remaining exactly as broken
  in production. Either (a) seed the key inside a migration's `up` block
  (mirroring Phase 2's own fix for the identical problem), or (b) write it as a
  rake task with no environment gate, run once manually against production as
  part of shipping this item — pick one and say so in the PR, don't leave it
  implicit. Whichever mechanism is used, it should create an `ApiKey` row via
  the normal `ApiKeyGenerator` path (so it's `requiem_`-prefixed,
  collision-checked, hashed correctly) tied to a dedicated internal user/plan.
  Decide the plan tier deliberately — it needs enough headroom that playground
  traffic doesn't get rate-limited against a real customer's budget, but should
  still be a real, bounded plan, not `enterprise`-with-null-limits, so a
  demo-abuse spike is still caught by Phase 2's quota/rate-limit middleware
  rather than silently unbounded.
- Store its raw value in `PLAYGROUND_API_KEY`, replacing the current stale
  `rq_test_playground_demo_key` default (which is invalid against the
  `requiem_[0-9a-zA-Z]{24}` format Go's validator now enforces, so it was never
  going to work even if it had been wired up).
- Wire `AppConfig.playground_api_key` into `ApiProxyService#call`: add a
  `"requiems-api-key" => ::AppConfig.playground_api_key` header alongside the
  existing `X-Backend-Secret` one — **don't remove `X-Backend-Secret` from this
  call**, since Caddy still gates on it until Phase 3 item 6 lands (see
  Context); this is additive, not a replacement.
- Test coverage: a request through `ApiProxyController`/`ToolDemosController`
  against a real (or stubbed-at-the-HTTP-boundary) Go backend succeeds with the
  new header present; a regression test asserting `X-Backend-Secret` is still
  sent (catches an accidental future removal before Phase 3 item 6 actually
  lands).

### 3. `trusted_proxies` — both sides of the CF-Connecting-IP trust boundary

Two independent fixes, same underlying issue, land together since they're small
and mutually reinforcing:

- **Rails:** configure `config.action_dispatch.trusted_proxies` (or the
  equivalent Rack middleware boundary) so `CF-Connecting-IP` is only trusted
  when the request actually arrived via Cloudflare's published IP ranges. Update
  `ApiProxyController#create` and any other `CF-Connecting-IP` reader in
  `apps/dashboard/app` to go through the trusted path rather than reading the
  header directly.
- **Go:** `callerIP()` (duplicated across `ip/info`, `ip/asn`,
  `global_data/timezone_ip`) currently trusts the first `X-Forwarded-For` hop
  unconditionally. Add the same trust boundary — a shared helper in
  `platform/httpx` (or wherever the existing cross-cutting HTTP helpers live)
  that only trusts `X-Forwarded-For`/`CF-Connecting-IP` when `RemoteAddr` (or
  the immediate hop) is Cloudflare's or the local Caddy's — and point all three
  call sites at it instead of leaving three copies of the same unvalidated
  logic. Deduplicating three copies into one helper is in scope here since the
  fix itself requires touching all three call sites anyway; don't leave three
  independently-drifting implementations of a security boundary.
- Test coverage for both: a spoofed header from an untrusted source falls back
  to the actual connection IP rather than being trusted verbatim.

### 4. Go integration-test database isolation

Decide and implement one fix for `app_test.go`'s shared-database collision (see
Context): either (a) a dedicated `TEST_DATABASE_URL` env var the test suite
prefers over `DATABASE_URL` when set, with `docker-compose.dev.yml` provisioning
a separate test database/user, or (b) have `seedAPIKeyFixture`/`TestApp_Handler`
create their fixture tables in a dedicated Postgres schema
(`CREATE SCHEMA IF NOT EXISTS go_test` + `search_path`) instead of the default
`public` schema, so they can't collide with Rails' migrated tables regardless of
run order. Pick whichever is less invasive once you're looking at the actual
test setup — this doc doesn't mandate one, both close the gap. Document the
chosen approach in `agents.md`'s existing dev-workflow section so the next
person running both suites back-to-back doesn't hit `PG::DuplicateTable` again.

### 5. Dead-code / schema cleanup (independent of Worker retirement)

All five of these were confirmed safe-to-remove-now by the original audit (not
gated on migration timing) and reconfirmed against the current tree while
writing this plan:

- Drop `subscriptions.stripe_customer_id`, `stripe_subscription_id`,
  `credit_limit` (zero usages confirmed) and the unused `solid_cache_entries`
  table, in one migration. Re-grep immediately before writing the migration in
  case anything changed since this plan was written.
- Add FKs on `credit_adjustments.admin_user_id`, `audit_logs.admin_user_id`,
  `abuse_reports.resolved_by_id`, matching the existing FK pattern on
  `subscriptions.promoted_by_id`. Check for orphaned values first (a data
  cleanup pass, if any exist) before adding the constraint, not after it fails
  to apply.
- Resolve the dead Pundit scaffolding: delete `app/policies/` entirely. Nothing
  in the current codebase indicates an in-progress adoption plan (zero
  `authorize` calls, zero concrete policy classes despite Pundit being in the
  Gemfile) — real authorization already lives in router-level and
  controller-level `before_action` gating, which works today. If a future
  session wants a real policy layer, that's a deliberate new feature, not a
  continuation of this scaffolding.

### 6. Two narrow, non-speculative doc fixes

- `docs/core/rails-app.md`'s "(separate schema)" description of
  `rails_schema_migrations` is factually wrong regardless of Worker-retirement
  timing — it's a renamed table in the same `public` schema, not a separate
  Postgres schema/namespace (confirmed in the architecture audit's PostgreSQL
  section). Fix this one line.
- Grep `docs/` and `agents.md` for any remaining `rq_live_`/`rq_test_` key-
  format literals beyond what Phase 3's session already fixed in the FAQ/ locale
  files — this was explicitly named as a leftover in `-phase-3-4.md`'s own
  follow-ups ("Phase 4's own item 3 ... should also grep ... beyond what this
  session already fixed"), is the same class of stale-literal bug already fixed
  once this series, and — unlike the rest of item 6 below — isn't gated on
  Worker retirement at all, so there's no reason to leave it for later.

**Nothing else in the doc sweep** — nothing that describes the Workers as
currently running is inaccurate yet, because they still are (see Context);
rewriting `architecture.md`/deleting `auth-gateway.md`/`api-management.md`
before Phase 3 items 6–7 actually land would make the docs describe a system
that doesn't exist yet, which is worse than what's there now.

## Test coverage summary

Each item above lists its own tests inline; overall, this plan follows the
repo's existing convention (Phase 0–3's own test suites) of hitting real
Postgres/Redis rather than mocking, for anything that touches the auth/proxy
path.

## Exit criteria

- `AnalyticsRevenueService`'s MRR/revenue-trend figures reflect actual billing
  cycle; backfill applied and spot-checked against LemonSqueezy's dashboard; a
  regression test exists for the yearly-vs-monthly divergence.
- A request through the public playground (`ApiProxyController`) and through a
  `ToolDemosController` demo form both succeed end-to-end against Go — manually
  verified, not just unit-tested, since this is the exact bug class ("looks fine
  in isolation, 401s in practice") that shipped silently last phase.
- `trusted_proxies` configured in Rails; Go's `callerIP()` (all three call
  sites) validates the immediate hop before trusting forwarded-IP headers.
- `go test ./...` and `bin/rails test` can both run back-to-back against a
  freshly-migrated dev stack with no table collision.
- Dead Stripe/credit_limit columns and `solid_cache_entries` dropped; new FKs
  added; `app/policies/` removed.
- `docs/core/rails-app.md`'s schema-separation line corrected; no remaining
  `rq_live_`/`rq_test_` literals anywhere in `docs/` or `agents.md`.

## Explicitly out of scope for this plan

- Phase 3 items 6–7: Cloudflare Authenticated Origin Pulls / origin firewall,
  removing Caddy's `X-Backend-Secret` gate, and deleting
  `apps/workers/{auth-gateway,api-management,shared}` plus the KV namespace/D1
  database. All human-gated, live-infra work requiring Cloudflare dashboard/VPS
  access in the room — pick this up as its own session when that access is
  available, following Phase 3's own item 6 ordering (origin lockdown verified
  live _before_ removing the existing secret-header gate).
- The rest of the Phase 3/4 doc's item 3 (full doc sweep deleting Worker/KV/D1
  references) — contingent on the above landing first.
- **Considered and deliberately left out, not overlooked:** two pieces of
  zero-infra-access prep work for the eventual items 6–7 session — writing (and
  dry-running against a non-prod namespace, if one exists) the KV/D1
  backup-export script (`wrangler kv key list/get`, `wrangler d1 export`) Phase
  3 item 7 will need, and staging the Caddy Authenticated-Origin-Pulls config
  change as a reviewable diff ahead of time. Both are genuinely doable without
  live Cloudflare credentials and would shrink the eventual human-gated session.
  Left out of this plan to keep this session's scope to the correctness bugs
  above rather than growing it into infra-adjacent prep — pick either up
  explicitly, as its own small item, if there's spare time after items 1–6 land
  clean.
- Backfilling `httpx.UsageCounter` across the ~220 endpoints that don't
  implement it (unrelated, separately-scoped, large).
- A dedicated/second Redis instance or logical-DB split; real load-based
  retuning of the placeholder pool sizes / cache TTLs / `maxmemory` value — all
  still waiting on real traffic to exist (audit's Open Question 5).
- **The rate limiter's structural inability to bound the auth-cache
  prefix-guessing exposure** (Phase 2's own carried-forward gap) — unrelated to
  the bugs in this plan, but worth flagging explicitly: this is now the _third_
  consecutive phase doc to re-defer this exact item (`-phase-0-1.md`'s Final
  Notes → `-phase-2.md`'s Context → here) with no session ever picking it up. It
  has no external dependency and a cheap candidate mitigation already named in
  `-phase-0-1.md` (shorten the `apikey:{prefix}` cache TTL further). If it's
  deferred again, the next plan in this series should either schedule it
  explicitly or make a deliberate, written decision to accept the risk — not
  defer it a fourth time by default. **Resolved — see
  standing-issues-hardening.md: candidate-only cache +
  bcrypt-reverify-every-hit,
  `apps/api/platform/middleware/apikeyauth.go:120-126`.**
- Confirming whether `api-management`'s `/analytics/*` endpoints have a hidden
  caller — irrelevant until Phase 3 item 7 actually deletes them.

## Open questions worth resolving before or during this session

- What field in LemonSqueezy's webhook payload actually carries the
  monthly-vs-yearly signal for item 1 — confirm against a real payload sample
  before writing the backfill script.
- What plan tier the playground's system API key should sit on (item 2) — needs
  a product judgment call, not just a technical one; default to something
  bounded and observable rather than unlimited.
