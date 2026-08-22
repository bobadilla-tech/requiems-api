# Go Auth Foundation — Phase 3 & 4: Key Ownership Cutover + Worker Retirement + Correctness Cleanup

Continuation of `docs/plans/2026-08-21-go-auth-foundation-phase-0-1.md`
(Go-native API-key auth, enforcing) and
`docs/plans/2026-08-21-go-auth-foundation-phase-2.md` (Redis rate limiting +
usage/quota, enforcing). Both shipped against
`docs/audits/2026-08-21-architecture-audit.md`'s Migration Phases 1–2, using the
same "no shadow mode, build it as the enforcing mechanism directly" rationale
throughout: there are no users and no production traffic, so the audit's
multi-week shadow-comparison/canary machinery (its Phases 3 and 5) protects
nothing that exists yet and is dropped, not deferred.

This plan finishes the job under the same rationale: it takes the audit's
Migration Phases 4, 6, and 7 (key-generation ownership, full traffic cutover,
D1/KV sync removal) and collapses them into one phase, because the
canary/percentage-rollout machinery in the audit's Phase 5 — which Phases 4/6/7
are staged around — has nothing to canary against. Phase 8 (delete the Worker
code itself, plus docs/CI cleanup) and the two standalone P1 correctness bugs
the audit flagged as unrelated to the migration timeline become Phase 4 of this
plan.

**Note on how to execute this document:** Phase 3 and Phase 4 are written
together because they're both in-scope for the same focused block of work, but
they carry very different risk profiles — Phase 3 touches live Cloudflare
DNS/proxy config and deletes infrastructure with no undo (see its item 6/7
gating below); Phase 4 is ordinary application-code/doc cleanup. Recommend
running them as separate PRs and, ideally, separate sessions — land and verify
Phase 3 (including its human-gated infra steps) before starting Phase 4, rather
than batching both into one continuous run. This also means Phase 3 alone is
enough to satisfy "1–2 phases for a focused session" if the infra-gated steps
end up taking longer than expected; Phase 4 can slip to a follow-up session
without blocking anything.

## Context

**The Worker→Go path is already non-functional, which changes the risk profile
of this work.** Phase 1's own "Final notes" flagged this explicitly: mounting
`APIKeyAuth` in the same `/v1` route group as `BackendSecretAuth`
(`apps/api/app/app.go:58-68`, both `protected.Use(...)`, AND-composed) means
real Worker-proxied traffic — which only ever carries `X-Backend-Secret`, never
`requiems-api-key`, because the Worker strips that header before proxying
(`apps/workers/auth-gateway/src/http.ts:20`) — 401s at `apiKeyAuth.Middleware()`
today. So the "before" state for this plan is not "a working Worker-fronted
system that needs a careful cutover"; it's "a Worker that cannot successfully
proxy a single authenticated request to Go right now." Cutting traffic direct to
Go is not a risky migration off a working path — it's the only way to get a
working path back.

**Verified against the current tree** (not re-derived from the audit, which is
same-day but was read-only): `apps/dashboard/app/models/api_key.rb` still gates
`generate_key_locally` behind `Rails.env.test?` and calls
`Cloudflare::ApiManagementService.new.create_key(...)` in every other
environment (`request_key_from_server`, `api_key.rb:38-67`);
`ApiKeyGenerator.generate`
(`apps/dashboard/app/services/api_key_generator.rb:9`) still produces
`rq_live_`/`rq_test_`-prefixed keys, not `requiem_`-prefixed; `ApiKey`'s
`after_destroy :remove_from_cloudflare` and
`after_update :sync_revocation_to_cloudflare, if: :saved_change_to_active?`
callbacks (`api_key.rb:24,25`) and `Subscription`'s
`after_create`/`after_update :sync_to_cloudflare` callbacks
(`apps/dashboard/app/models/subscription.rb:25-26`) are all still live, and the
LemonSqueezy webhook controller still makes its own explicit
`Cloudflare::ApiManagementService.new.sync_user_plan(...)` calls at
`apps/dashboard/app/controllers/webhooks/lemonsqueezy_controller.rb:108,160,186,210`
— the duplicate-sync bug the audit flagged is unfixed. `subscriptions.plan` (the
billing-cycle column) is confirmed present in `db/schema.rb` and confirmed never
written by any webhook handler above (only `plan_name`, `status`,
`current_period_*`, etc. are set) — `AnalyticsRevenueService#calculate_mrr` and
`#revenue_by_plan`/`#revenue_trend`
(`apps/dashboard/app/services/analytics_revenue_service.rb:36-38,51,66-67`) all
key off `sub.plan&.to_sym || :monthly`, so every subscription's MRR contribution
is still computed as monthly pricing regardless of actual billing cycle. None of
this was touched by Phase 0/1/2, which only worked in `apps/api` and in the
narrow Go-auth-cache slice of `apps/dashboard` (`GoAuthCache`, `apikeyauth.go`)
— so the audit's citations for all of the above are still accurate today.
`AnalyticsRevenueService`'s `sub.plan&.to_sym` reads are at
`apps/dashboard/app/services/analytics_revenue_service.rb:44,48`
(`calculate_mrr`) and `:64,76` (`revenue_trend`).

**Go's prefix-extraction contract is fixed and must not change.**
`apps/api/platform/middleware/apikeyauth.go:26` hardcodes
`keyPrefixLength = 12`, already shipped and tested in Phase 1. Whatever key
format Phase 3 below produces, the first 12 characters must remain the
`key_prefix` Postgres stores and Redis caches by — this plan fixes the literal
prefix string (`rq_live_`/`rq_test_` → `requiem_`) and adds collision handling,
it does not touch the 12-character slicing convention itself.

## Approach

### Phase 3 — Rails-owned key generation, full Worker/KV/D1 retirement, origin lockdown

**1. Fix `ApiKeyGenerator`'s output format.**
`apps/dashboard/app/services/api_key_generator.rb`: change `generate` to produce
`requiem_<24-char-alnum>` (matching the validator Go already runs against,
`^requiem_[0-9a-zA-Z]{24}$`, and the shape Phase 0's dev-seed data already uses)
instead of `rq_live_`/`rq_test_`. Confirm `extract_prefix`'s `full_key[0..11]`
still yields a 12-character prefix under the new format (it does — `requiem_` is
8 chars + 4 chars of the random part, same slice logic, no code change needed
there beyond the literal prefix string used in `generate`). Update any existing
Rails tests/fixtures asserting on the old `rq_live_`/`rq_test_` shape.

**2. Add collision-check-and-retry.** `ApiKeyGenerator.generate` (or a wrapping
method) should check for an existing `key_prefix` before returning — Phase 0's
btree index (`index_api_keys_on_key_prefix_btree`) makes this an efficient
`EXISTS` query, not a table scan — and retry with a freshly generated key on
collision, capped at a small fixed number of attempts (e.g. 5) before raising.
Mirrors the "extremely unlikely with nanoid, but good practice" comment the
audit found in the Worker's own `create.ts:53-64`.

**3. Make `generate_key_locally` (or equivalent) the only code path, in every
environment.** `ApiKey#request_key_from_server`
(`apps/dashboard/app/models/api_key.rb:32-51`): delete the `Rails.env.test?`
branch and the `Cloudflare::ApiManagementService` branch entirely; always
generate locally via `ApiKeyGenerator` (now producing the correct format with
collision handling per items 1–2). No Worker/KV round trip on key creation, in
any environment, from this point on — there is no transitional dual-write to
Cloudflare KV to build (unlike the audit's own Migration Phase 4, which needed
one only because it assumed live traffic still flowing through the Worker; that
premise doesn't hold here, see Context).

**4. Delete the Cloudflare-sync callbacks, not convert them.** Remove `ApiKey`'s
`after_destroy :remove_from_cloudflare` and
`after_update :sync_revocation_to_cloudflare, if: :saved_change_to_active?`
(`api_key.rb:24,25`) and their private methods — `remove_from_cloudflare`
(`api_key.rb:100-106`) **and** `sync_revocation_to_cloudflare`
(`api_key.rb:108-112`, easy to miss since it's not adjacent to the first method
— confirm both are gone, not just the one nearer the callback declarations);
remove `Subscription`'s `after_create`/`after_update :sync_to_cloudflare`
callbacks (`subscription.rb:25-26,30-35`); remove the four explicit
`Cloudflare::ApiManagementService.new.sync_user_plan(...)` calls in
`lemonsqueezy_controller.rb:108,160,186,210`. This is a straight deletion, not
the audit's suggested `after_commit`-conversion fix for the duplicate-sync bug —
that fix mattered only if the sync itself was being kept; since Cloudflare sync
is retired wholesale in this same phase, converting it first and deleting it a
phase later would be pure churn. `invalidate_go_auth_cache` (the
`after_destroy`/`after_update` callbacks calling `GoAuthCache.invalidate`,
`api_key.rb:25,27`) stays — that's the Redis cache invalidation Phase 1 built,
unrelated to Cloudflare.

**5. Remove `BackendSecretAuth` from Go's protected route group.**
`apps/api/app/app.go:58-59`: delete
`protected.Use(middleware.BackendSecretAuth(cfg.BackendSecret))` and the
now-stale comment block above `apiKeyAuth.Middleware()` at `:60-67` describing
the Worker's header-stripping behavior (replace with a short note that this is
now the sole enforcing auth check for all traffic, direct or proxied). This is
the one-line change that actually fixes the 401-on-all-traffic state described
in Context — without it, deleting the Worker just changes _why_ every request
fails, not _whether_ it does. Remove `BackendSecretAuth` itself
(`apps/api/platform/middleware/auth.go`) and its test file once nothing
references it; remove `cfg.BackendSecret` from `platform/config/config.go` and
`BACKEND_SECRET` from env files once the Go side no longer reads it (Caddy still
needs the literal env var for item 6 below until that item also lands — sequence
these in the same PR, not two separate merges that leave a window where Caddy
demands a header Go no longer checks).

**6. Cut Cloudflare over to proxy-only and lock down the origin — a separate,
explicitly human-gated sub-phase, not a code diff an agent should run straight
through.** This step touches real production infrastructure (the VPS, the live
Cloudflare zone) outside this repo, and two of its actions are effectively
irreversible (KV/D1 deletion in item 7 has no undo; a botched DNS/proxy-mode
flip or firewall rule can take the origin down or expose it, independent of
whether any customers exist yet — "no users" justifies skipping the audit's
_shadow-mode_ machinery for application code, it does not make an infra change
reversible or lower-risk). Do not treat asking "should I proceed?" and getting a
"yes" in chat as sufficient gating for the deletions in item 7 — get an explicit
go-ahead **per irreversible action**, immediately before taking it, not once at
the start of the sub-phase. Concretely, in this order — **the order matters, do
not reverse it**:

1. Configure the replacement origin-lockdown control _first, while the existing
   `X-Backend-Secret` gate is still live_: Cloudflare Authenticated Origin Pulls
   (mTLS between Cloudflare's edge and the origin — requires installing
   Cloudflare's origin CA cert into Caddy's TLS config) or an origin firewall
   rule restricting inbound to Cloudflare's published IP ranges. This requires
   Cloudflare dashboard/API access outside this repo — get explicit confirmation
   from whoever holds that access before applying it, since it's the first touch
   on live DNS/firewall config.
2. **Verify the new control actually works before removing the old one**:
   confirm a direct request to the origin's IP/hostname (bypassing Cloudflare's
   proxy) is rejected. Configuring AOP/the firewall rule is not evidence it
   works — test it live.
3. Only after step 2 passes, remove `infra/caddy/Caddyfile`'s
   `@authorized`/`X-Backend-Secret` gate on the `internal.requiems.xyz` block
   (`:32-47`) and proxy directly to `api:8080`. **Doing this before step 2 is
   verified re-creates the exact gap the audit's Migration Phase 6 calls a hard
   requirement to avoid** — a window where the origin can be reached directly,
   bypassing Cloudflare's WAF/DDoS/rate-limiting (not a fully-open API —
   `APIKeyAuth` in Go still gates every request regardless of this ordering —
   but still a real bypass of the protection this step exists to keep in place).
4. Add `trusted_proxies`-equivalent client-IP handling in Go (or Caddy) for
   `CF-Connecting-IP`, since it now applies directly to Go instead of being
   pre-filtered by a Worker that no longer exists. Do this alongside step 3,
   before real traffic relies on it. Overlaps with, and can be done together
   with, the independent Rails `trusted_proxies` fix in Phase 4 item 2 below —
   same underlying issue, two different services.
5. **Rollback plan, in case step 2's verification fails or step 3 breaks
   something once live**: keep the pre-change Caddyfile block committed but
   unreleased (don't squash it away) so reverting is `git revert` + Caddy
   redeploy, not a from-memory rewrite. If AOP/the firewall rule breaks
   legitimate Cloudflare-routed traffic after step 3, the fastest recovery is
   reverting the Caddyfile change (re-adding the secret-header gate) while the
   AOP/firewall config is debugged separately — don't try to fix both at once
   under live traffic.

**7. Delete the Workers and their data stores — only after item 6 is fully
verified live, and only with explicit human sign-off immediately before each
deletion.** Delete `apps/workers/auth-gateway`, `apps/workers/api-management`,
`apps/workers/shared` (reversible — it's a `git revert` away). Before touching
Cloudflare's KV namespace or D1 database: **export/back up their contents
first** (`wrangler kv key list`/`get` or a bulk export, `wrangler d1 export`) —
these are the two genuinely irreversible actions in this whole plan, so get an
explicit go-ahead from whoever controls the Cloudflare account for the deletion
specifically, not folded into an earlier "does this plan look OK" approval.
Delete `apps/dashboard/app/services/d1_sync_service.rb`,
`apps/dashboard/app/services/cloudflare/api_management_service.rb`,
`apps/dashboard/app/jobs/sync_d1_usage_job.rb` and its Sidekiq schedule entry.
Remove
`CLOUDFLARE_ACCOUNT_ID`/`CLOUDFLARE_KV_NAMESPACE_ID`/`CLOUDFLARE_API_TOKEN` from
`infra/docker/.env.example:108-110` (already confirmed dead by the audit
independent of this migration, so this part carries no risk). Remove Worker
entries from `.dependabot.yml`, `infra/docker/docker-compose.dev.yml`,
`_worker-ci.yml` and its callers in `ci.yml`.

**8. Test coverage for this phase**, following the repo's
real-Postgres/real-Redis convention:

- `ApiKeyGenerator`: format matches `^requiem_[0-9a-zA-Z]{24}$`; collision
  triggers a retry with a new key, not a validation error, up to the retry cap;
  exhausting the retry cap raises a clear error.
- `ApiKey` creation in `development`/`production`-like environments (not just
  `test`) produces a valid local key with no HTTP call attempted — assert no
  `Cloudflare::ApiManagementService` reference remains reachable at all
  (compile-time via deletion, not just a runtime stub).
- Revoking/destroying an `ApiKey` no longer attempts any Cloudflare call
  (trivially true once the callbacks are deleted, but keep a regression test
  asserting `GoAuthCache.invalidate` still fires — Phase 1's existing test for
  this should keep passing unmodified, confirm it does).
- Go integration test: a Rails-created dev key (via the fixed generator)
  authenticates through `apps/api` with **no** `X-Backend-Secret` header sent at
  all — this is the test that would have caught today's 401-everything state,
  and it's the one that proves item 5 actually fixed it.

**Phase 3 exit criteria:**

- A key created through the normal Rails signup flow (any environment) is
  `requiem_`-prefixed, collision-checked, and immediately authenticates against
  `apps/api` directly with no Cloudflare involvement anywhere in the create
  path.
- `apps/api`'s protected route group has exactly one auth gate (`APIKeyAuth`) —
  `BackendSecretAuth` and `BACKEND_SECRET` are gone from Go's config and route
  setup.
- `internal.requiems.xyz` proxies directly to Go with no header-secret gate;
  Cloudflare origin-lockdown (AOP or IP firewall) is live and independently
  verified by a direct-to-origin request test, not just configured.
- `apps/workers/{auth-gateway,api-management,shared}`, the KV namespace, the D1
  database, and their Rails-side sync services/jobs no longer exist in the repo
  or in Cloudflare.
- Full test suite (item 8) green.

### Phase 4 — Correctness cleanup (independent bugs + doc/dead-code sweep)

Everything here is either already-independent of the migration per the audit
(items 1–2 could in principle ship before Phase 3, but are grouped here since
Phase 3 already touches every file item 1 touches) or is only unblocked once
Phase 3's deletions land (items 3–5).

**1. Fix the MRR/revenue billing-cycle bug.** `subscriptions.plan` is never
written by any webhook handler (confirmed in Context above). Populate it from
the LemonSqueezy variant/webhook payload in `lemonsqueezy_controller.rb`'s
`handle_subscription_created`/ `handle_subscription_updated` (wherever the
yearly-vs-monthly signal is actually present in the payload — confirm the exact
field at implementation time, since the audit didn't pin one down;
LemonSqueezy's variant metadata or the webhook's `data.attributes` interval
fields are the likely source). Backfill existing subscription rows from
LemonSqueezy's own records if they're still retrievable via their API (per the
audit's Open Question 6 — if webhook payload history isn't retained, flag rows
needing manual reconciliation rather than guessing). Write this as a one-off
Rails runner script, verify the backfilled totals against LemonSqueezy's own
dashboard MRR figure before trusting it.

**2. Configure `trusted_proxies` in Rails.** `ApiProxyController` (or
equivalent) currently trusts `CF-Connecting-IP` with no configured trust
boundary — independent of the migration per the audit (Rails/`requiems.xyz` has
always been fronted by Cloudflare directly via its own Caddyfile block, never by
the Worker). Can be done alongside Phase 3 item 6's Go-side equivalent, same
underlying issue.

**3. Delete the retired Worker/KV/D1 documentation and rewrite what depends on
it**, per the audit's Documentation Audit table:

- Delete `docs/core/auth-gateway.md`, `docs/core/api-management.md`.
- Rewrite `docs/core/architecture.md` and `docs/core/infrastructure.md` to
  describe Go-native auth/rate-limit/usage backed by Redis/Postgres, with no
  KV/D1/Worker references.
- Remove `docs/core/deployment.md`'s "Part 5: Deploy Cloudflare Workers"
  section; add the origin-lockdown (AOP/IP-firewall) setup steps from Phase 3
  item 6 in its place.
- Update `docs/core/rails-app.md`'s "Cloudflare Integration" section to describe
  direct Postgres key CRUD and Go-native usage writes; also fix the unrelated
  "(separate schema)" mischaracterization the audit flagged
  (`rails_schema_migrations` is a renamed table in the same `public` schema, not
  a separate Postgres schema — a doc correction, not a code change).
- Update `agents.md` (repo root) — **this is a bigger edit than "drop two
  component descriptions" implies**: the file has ~30+ live references to the
  Workers, including a full dev-workflow section (container names like
  `requiem-dev-auth-gateway-1`/`requiem-dev-api-management-1`, `docker exec`
  test/typecheck commands scoped to those containers, ports, env vars, and a
  "must pass — 71 tests" CI-gate line tied to Worker test counts). Treat this as
  a real rewrite of a load-bearing contributor-instructions file, not a
  find-replace — budget time for it accordingly, don't scope it as trivial.
- Run a full-text search for `auth-gateway`, `api-management`, `KV`, and `D1`
  across `docs/` **and `agents.md`** as a final sweep — the audit's own list is
  explicitly non-exhaustive.
- Update `tests/integration/src/suites/gateway.test.ts` and
  `tests/load/scenarios/rate-limit.ts` to drop Worker/KV-specific assertions in
  favor of Go-equivalent checks (both already hit the public gateway URL, not
  Worker internals, per the audit — expect this to mostly be `API_BASE_URL`
  re-pointing plus deleting a handful of KV-specific assertions, not a rewrite).

**4. Dead-code sweep, now safe since Phase 3 deleted their only remaining
consumers:**

- Drop Rails' dead Stripe-era columns (`subscriptions.stripe_customer_id`,
  `stripe_subscription_id`, `credit_limit`) and the unused `solid_cache_entries`
  table — confirm zero references first with a fresh grep (the audit already
  did, but do it again post-Phase-3 in case anything changed), one migration.
- Add missing FKs on `credit_adjustments.admin_user_id`,
  `audit_logs.admin_user_id`, `abuse_reports.resolved_by_id`, matching the
  existing FK on `subscriptions.promoted_by_id`. Run a data-cleanup pass first
  if any orphaned values exist (check before adding the constraint, not after it
  fails to apply).
- Resolve the dead Pundit scaffolding one way or the other (`app/policies/`) —
  either delete it or wire up at least one real `authorize` call; the audit
  takes no side on which, just flags it as misleading as-is.
- Confirm whether `api-management`'s `/analytics/*` endpoints had any caller
  outside Rails before their deletion in Phase 3 item 7 — if this wasn't
  independently confirmed before that deletion, note it in the PR description as
  a known unknown rather than silently assuming "no caller found" meant "no
  caller."

**Phase 4 exit criteria:**

- `AnalyticsRevenueService`'s MRR/revenue-trend figures reflect actual billing
  cycle (yearly subscribers no longer priced as monthly); backfill applied and
  spot-checked against LemonSqueezy's dashboard.
- `trusted_proxies` configured in Rails.
- `docs/core/*` and `agents.md` contain no stale Worker/KV/D1 references
  (verified by the full-text sweep in item 3's last bullet).
- Dead columns/table dropped; new FKs added; Pundit scaffolding resolved.

## Explicitly out of scope for this plan

- Backfilling `httpx.UsageCounter` across the ~220 endpoints that don't
  implement it yet (audit flagged this as separate, larger work in Phase 2's
  plan already; still true here).
- A dedicated/second Redis instance or logical-DB split (audit's fuller
  Reliability recommendation beyond the `noeviction`+`maxmemory` config Phase 2
  already shipped).
- Real load/traffic-based re-tuning of pool sizes, cache TTLs, or the 256mb
  Redis `maxmemory` value — all still placeholders per Phase 0–2's own notes,
  revisit once real traffic exists (audit's Open Question 5).
- The rate limiter's structural inability to bound the auth-cache
  prefix-guessing exposure (Phase 2's own flagged, carried-forward gap) —
  unrelated to key _generation_ (this plan's Phase 3), only to key
  _verification_ cache behavior (Phase 1, already shipped).
- Rewriting `AnalyticsRevenueService` beyond fixing the `plan` column input — no
  other correctness issue was found in it.

## Open questions worth resolving before or during Phase 3

- Who has Cloudflare dashboard / VPS access to execute item 6's DNS-proxy-mode
  and origin-firewall changes, and item 7's KV-namespace/D1-database deletion?
  These need a human in the loop regardless of how the code-side PRs are
  structured.
- Does LemonSqueezy's webhook payload actually carry a monthly-vs-yearly signal
  in a field this codebase hasn't parsed yet, or does Phase 4 item 1 need a
  separate LemonSqueezy API call to fetch it per-subscription? Confirm against a
  real webhook payload sample (LemonSqueezy's dashboard can usually replay one)
  before writing the backfill script.
