# Go Auth Foundation — Phase 7: Cloudflare Workers Retirement (Full Cutover)

Continuation of `docs/plans/2026-08-21-go-auth-foundation-phase-0-1.md`,
`-phase-2.md`, `-phase-3-4.md`, `-phase-5.md`, and
`2026-08-22-go-auth-foundation-phase-6-usage-multiplier-and-loose-ends.md` (and
the standing-issues-hardening doc for PR #966). This plan was written by
re-reading `2026-08-21-architecture-audit.md` end to end and re-verifying the
current tree directly (`git branch`, `grep`, file reads — not trusting any prior
doc's own "Final Notes" section as current truth) on 2026-08-22.

## Context — what is actually still open

**Fully shipped and enforcing, reverified today, not revisited by this plan:**

- Go-native API-key auth (`apps/api/platform/middleware/apikeyauth.go`,
  candidate-by-`key_prefix` + bcrypt, Redis-cached at `apikey:{prefix}`,
  fail-closed) is the **sole** auth gate on Go's `/v1` group —
  `BackendSecretAuth` was deleted from Go entirely in Phase 3
  (`apps/api/app/app.go:60`, `apps/api/app/app_test.go:161`).
- Redis rate limiting (`platform/middleware/ratelimit.go`, Lua
  `INCR`+`EXPIRE 60`, key `ratelimit:{api_key_id}:{unix_minute}`, fail-open on
  Redis error).
- Usage/quota enforcement (`platform/middleware/usage.go`, key
  `usage:{user_id}:{cycle_start_unix}`, fail-closed 503 on Redis failure) —
  writes `usage_logs` rows **synchronously per request**, not via a batched
  Redis→Postgres flush (a deliberate deviation from the audit's proposed design,
  accepted because there's no real traffic yet to make that a performance
  problem — see Notes below on whether to revisit this).
- Rails owns key generation exclusively (`ApiKeyGenerator`, `requiem_`-prefixed,
  collision-checked, `apps/dashboard/app/services/api_key_generator.rb`) — zero
  Cloudflare HTTP calls on key creation.
- All Cloudflare-sync callbacks deleted from `ApiKey`/`Subscription`; the
  duplicate-sync and MRR/`plan` bugs from the audit are both fixed and tested.
- `trusted_proxies` configured on both sides: Rails (`app/lib/trusted_proxy.rb`)
  and Go (`platform/httpx/trustedproxy.go` — already includes Cloudflare's
  published edge CIDRs as "defense in depth once Caddy proxies directly," i.e.
  written in anticipation of this phase).
- The Worker's per-endpoint usage-multiplier config drift is fixed and ported to
  Go (dictionary/thesaurus 2x, Phase 6a).
- Dead-code sweep done: Stripe columns, `solid_cache_entries`, Pundit
  `app/policies/`, missing FKs, all eight stale `rq_live_`/`rq_test_` literals,
  `apps/dashboard/app/services/cloudflare/api_management_service.rb` deleted,
  `CLOUDFLARE_*` vars removed from `infra/docker/.env.example`.
- **Worker-retirement prep artifacts already exist, unmerged:**
  `apps/workers/auth-gateway/scripts/export-backup.ts` (dry-run validated, on
  `feat/go-auth-foundations`) and a full Caddy Authenticated Origin Pulls (AOP)
  diff on branch `prep/caddy-authenticated-origin-pulls`
  (`caddy validate`-clean, cert vendored at
  `infra/caddy/certs/cloudflare-origin-pull-ca.pem`) — confirmed present in the
  local repo today via `git branch -a`.

**Confirmed still open — this is what Phase 7 covers:**

1. **Traffic still routes through the Worker.** `infra/caddy/Caddyfile:32-47`
   still gates `internal.requiems.xyz` on `X-Backend-Secret`;
   `apps/workers/{auth-gateway,api-management,shared}` still exist and are still
   deployed to Cloudflare. Go's own API-key auth is the enforcing gate for
   whatever reaches it, but the Worker is still the only thing Cloudflare's edge
   sends traffic to — full direct-to-Go cutover has not happened.
2. **Cloudflare KV namespace and D1 database are still live** (`credit_usage`
   ledger, `api_keys` audit mirror) and still being written by the Worker and
   read by Rails' `D1SyncService`/`SyncD1UsageJob` on a schedule — confirmed
   both `apps/dashboard/config/recurring.yml` (every 5 minutes) and
   `apps/dashboard/config/sidekiq_schedule.yml` (every 3 minutes,
   duplicate/stale schedule config, see Phase 7b item 1) still register
   `sync_d1_usage`.
3. **CI/CD, dependabot, and docker-compose still reference the Workers**:
   `.github/dependabot.yml` (3 npm ecosystem entries),
   `.github/workflows/ci.yml` (`worker-test`/`api-management-test` jobs — the
   second job's key is `api-management-test`, not `worker-management-test`;
   verify the exact key with a grep before editing — calling the reusable
   `_worker-ci.yml`, lines ~300-380), `.github/workflows/_worker-ci.yml` itself,
   and `infra/docker/docker-compose.dev.yml` (wrangler dev containers,
   `wrangler_state` volume, `auth_gateway_modules`/`api_management_modules`/
   `shared_modules` volumes).
4. **Docs still describe the Worker-based architecture**:
   `docs/core/auth-gateway.md`, `docs/core/api-management.md` (whole docs), plus
   sections in `docs/core/architecture.md`, `docs/core/infrastructure.md`,
   `docs/core/deployment.md` ("Part 5: Deploy Cloudflare Workers"),
   `docs/core/rails-app.md` ("Cloudflare Integration" section), and `agents.md`
   (repo root). (The audit's separately-flagged "(separate schema)"
   mischaracterization at `rails-app.md:207-213` has already been fixed in the
   current tree — reverified today, the doc now correctly says "same `public`
   schema, not a separate Postgres schema/namespace" — nothing to do there.)
5. **`tests/integration/src/suites/gateway.test.ts` and
   `tests/load/scenarios/rate-limit.ts`** carry Worker-specific comments/
   assertions (confirmed: `gateway.test.ts`'s docblock says "validate the full
   Worker → Backend flow"; `rate-limit.ts:6,16` references "the Auth Gateway
   (Cloudflare Worker)" and `apps/workers/shared/src/config.ts` as the source of
   rate-limit values) — both already hit the public gateway URL rather than
   Worker internals, so per the audit's own Phase 8 note, most of this is
   comment/doc-string cleanup plus confirming `API_BASE_URL` points at the right
   place post-cutover, not a rewrite.
6. **`PLAYGROUND_API_KEY` is still unprovisioned in production**
   (`app_config.rb:185`'s `requiem_notprovisioned0000000000` default), tracked
   in `docs/core/v2-deployment-playbook.md` — a real-infra action requiring an
   explicit go-ahead, distinct from (but worth bundling with) this phase's other
   production changes.

## What "no users, no traffic" changes about this plan

The audit's Migration Phases 4-6 (percentage canary, consistent-hash routing,
at-least-one-full-billing-cycle shadow comparison) exist to de-risk a cutover
against real customer traffic and real billing accuracy. **None of that applies
here.** There are no real customers, no real subscriptions with billing history
that matters, and no SLA to protect during the cutover window. This phase is a
**direct, full cutover**, not a canary:

- No percentage-based traffic splitting. Cloudflare stops routing to the Worker
  for 100% of traffic in one change, not a gradual ramp.
- No shadow-comparison period. Go's auth/rate-limit/usage has been the enforcing
  path for everything that reaches it since Phase 0/1-2; the only new thing this
  phase does is make sure _all_ traffic reaches it directly instead of via the
  Worker.
- **DB rows are disposable, with one exception.** Per explicit instruction: it
  is fine to drop/truncate any operational-state row (`credit_usage` in D1,
  `usage_logs`, `daily_usage_summaries`, `credit_adjustments`, `audit_logs` in
  Postgres) without migration or reconciliation effort. The one exception is
  Go's own **reference/data tables** (`advice`, `quotes`, `words`, `bin_data`,
  `inflation_data`, `iban_countries`, `commodity_price_history`, `exercises`,
  `swift_codes`, `counters`) — these are product content, not auth/billing
  state, and must not be touched by anything in this phase. This means: **do not
  build a D1→Postgres reconciliation or backfill step for
  `credit_usage`/`api_keys` audit rows before deleting D1** — run the
  already-built `export-backup.ts` once for the record (cheap insurance, not a
  correctness requirement) and then delete.
- The MRR-backfill question (Phase 6b item 2) was already answered "zero
  affected rows, confirmed by the project owner" — nothing to redo here.
- **This carve-out is scoped narrowly. It covers `credit_usage` (D1),
  `usage_logs`, `daily_usage_summaries`, `credit_adjustments`, and `audit_logs`
  only** — the operational/billing-ledger tables that exist purely to record
  traffic that never happened. It does **not** cover `api_keys`, `users`,
  `subscriptions`, or `plans` — these are identity/ billing configuration, not
  usage history, and Phase 7c item 7's cleanup must not be read as license to
  touch them. Nor does it cover Go's reference/data tables, named explicitly
  above. Any step in this plan that says "truncate" or "drop rows" means the
  five ledger tables listed at the start of this bullet, nothing else.

This does **not** relax the origin-lockdown requirement (Migration Phase 6 in
the audit): the moment `X-Backend-Secret` is removed, Go becomes directly
internet-reachable behind Cloudflare's proxy, and skipping AOP or an equivalent
origin firewall would be a real, currently-exploitable gap regardless of whether
there's billing data at stake. That verification step stays mandatory.

## Approach

**Expect this to span 2-3 separate sessions/PRs, not one sitting.** 7a needs a
human at a live Cloudflare dashboard plus VPS access, ends with an observed bake
period, and its exit criteria must actually land (not just be staged) before 7b
starts; 7b needs that same live access again for the KV/D1 deletion; 7c is a
multi-file doc/test rewrite that depends on 7b's deletions. Don't try to
compress all three into one session — if 7a stalls waiting on dashboard access,
stop there rather than staging 7b/7c work against an unfinished cutover.

### Phase 7a — Cloudflare/infra cutover (human-gated throughout, needs live Cloudflare dashboard + VPS access in the room)

**Every step below that touches live Cloudflare configuration, production
DNS/routing, or a deployed secret requires a human with production access
present and confirming — this is not a phase an agent should run unattended end
to end.** Ordering matters — each step depends on the previous one landing and
being verified, not just applied.

1. **Run the backup-export script once, for the record.**
   `apps/workers/auth-gateway/scripts/export-backup.ts --remote` against the
   live KV namespace and D1 database. Store the output artifact wherever the
   team keeps this kind of one-off ops output (not committed to the repo). Per
   the "DB rows are disposable" note above, this is a courtesy snapshot, not a
   blocking correctness gate — do not let it expand into a reconciliation
   project.

2. **Fix the Rails→Go internal-URL collision before touching Caddy.**
   `apps/dashboard/app/services/api_proxy_service.rb` calls
   `AppConfig.internal_api_url`, which Kamal sets to
   `https://internal.requiems.xyz` in production
   (`infra/kamal/deploy.dashboard.yml`) — the same public hostname AOP is about
   to become the _sole_ gate on. Rails has no Cloudflare origin-pull client
   certificate, so once step 5 below lands, every Playground/demo-form request
   from Rails would fail the TLS handshake outright (not a 403 — the connection
   itself is refused), silently breaking the public Playground. **Before
   removing `X-Backend-Secret` in step 5**, repoint production's
   `INTERNAL_API_URL` to a path that bypasses the AOP-gated public vhost
   entirely — e.g. a private-network address straight to the Go container/Kamal
   accessory, mirroring dev's `INTERNAL_API_URL=http://api:8080`
   (`infra/docker/.env.example`) — and verify a real Playground request succeeds
   against the new URL before proceeding. Do not treat this as covered by step
   8's public smoke test; that test exercises the public API, not Rails'
   internal call path, and would not catch this regression.

3. **Add Cloudflare-side routing for `api.requiems.xyz` before removing the
   Worker.** Today, `api.requiems.xyz` traffic is entirely intercepted by the
   `auth-gateway` Worker's route binding — there is no Caddy site block for
   `api.requiems.xyz` in `infra/caddy/Caddyfile` (only `requiems.xyz`,
   `mcp.requiems.xyz`, and `internal.requiems.xyz` exist today). Once the Worker
   route is removed (step 7), Cloudflare needs somewhere else to send that
   traffic. Before step 7, do one of: (a) add a Cloudflare Origin Rule/Transform
   Rule rewriting the Host/SNI for `api.requiems.xyz` requests to
   `internal.requiems.xyz` so they land on the existing site block, or (b) add a
   new `api.requiems.xyz` site block to `infra/caddy/Caddyfile` with the same
   AOP `tls { client_auth }` config as `internal.requiems.xyz`. Decide which
   before executing this phase (an Origin Rule is likely less code to maintain,
   but confirm it supports SNI/client-cert-required backends the way this setup
   needs). Verify this routing works **before** step 7 removes the Worker, not
   after — test it while the Worker route is still active as a fallback.

4. **Merge and deploy `prep/caddy-authenticated-origin-pulls`, timed tightly
   with the Cloudflare-side toggle (step 5) — do not let these drift apart.**
   This adds the `tls { client_auth { mode
   require_and_verify ... } }` block
   to `internal.requiems.xyz` in `infra/caddy/Caddyfile`. **Important correction
   to the branch's own comment, which is wrong about Caddy's behavior:** the
   comment claims this "only takes effect once the toggle is switched on" — that
   is not how Caddy's `client_auth` works. The moment this config is redeployed,
   Caddy enforces the client-certificate requirement on _every_ TLS connection
   to `internal.requiems.xyz`, regardless of whether Cloudflare's dashboard
   toggle is on yet. Since Cloudflare doesn't attach the origin-pull client cert
   until that toggle is flipped, this config redeploy **alone** breaks every
   current caller of `internal.requiems.xyz` the moment it lands — including the
   Worker's own proxied requests, which are still the production path at this
   point. **Do not redeploy this Caddy change and then go find someone with
   dashboard access** — have the person with dashboard access ready to flip the
   toggle immediately (within the same short window, ideally scripted/rehearsed
   beforehand) so the outage gap is seconds, not an open-ended wait. Have the
   pre-change `Caddyfile` ready to redeploy instantly as a rollback if the
   toggle can't be flipped right away.

5. **Flip the Authenticated Origin Pulls toggle in the Cloudflare dashboard for
   the zone**, immediately following step 4's redeploy per the timing note
   above.

6. **Verify AOP is actually enforcing, not just configured**, per the audit's
   explicit "test it, don't just configure it" requirement — and verify the
   rejection mechanism itself, not just its outcome:
   - A request routed through Cloudflare (via the Worker, still in front at this
     point) still succeeds end-to-end.
   - A direct request to the origin's IP/hostname, bypassing Cloudflare's proxy
     entirely (`curl --resolve` against the VPS IP, or from a host not behind
     Cloudflare, ideally a network path not otherwise allowlisted at the VPS
     firewall level — a firewall-level block would produce a false pass that
     looks like AOP working but isn't), is rejected. **Capture and inspect the
     actual TLS alert** (e.g. "certificate required"/handshake failure) rather
     than just checking that the connection failed, to confirm the rejection is
     coming from Caddy's `client_auth` block specifically, not an unrelated
     network failure. Confirm the rejection happens even with a correctly-formed
     `X-Backend-Secret` header attached (still relevant until step 8 removes
     it), to prove AOP — not the header — is the operative gate.

7. **Remove the `X-Backend-Secret`/`@authorized` gate as its own follow-up
   commit**, once step 6 passes. `infra/caddy/Caddyfile`'s
   `internal.requiems.xyz` block keeps only the AOP `tls { client_auth }` block.
   Redeploy Caddy. Remove `BACKEND_SECRET` from
   `infra/docker/.env.example`/Kamal secrets.

8. **Cut Cloudflare routing from "Worker in front" to "proxy-only" for
   `api.requiems.xyz`.** **Disable** the Worker route (`auth-gateway`'s
   route/trigger binding) — do not delete the Worker itself yet, and prefer
   disabling the route over deleting the binding outright, so there's a fast
   path back to "Worker in front" if the direct-cutover traffic misbehaves. This
   requires step 3's routing to already be verified working. This is a single,
   full cutover for 100% of traffic — no percentage split, per the "no users, no
   traffic" note above. Keep WAF/DDoS/TLS proxying (orange-cloud DNS) — do not
   go grey-cloud/ DNS-only, that would remove Cloudflare's edge protection
   entirely, which the audit explicitly does not recommend.

9. **Review Cloudflare WAF/rate-limiting rules for the zone** (audit Risk #3 —
   the Worker's KV-based per-key rejection at the edge goes away; Cloudflare's
   generic WAF is what's left in front of Go now). Tighten/add rules as needed
   given current traffic (none yet) — this is a lightweight pass given there's
   no real traffic pattern to tune against yet, not a full WAF policy design
   project.

10. **Smoke-test the full public flow post-cutover**: a real `requiems-api-key`
    request through `api.requiems.xyz` (now Cloudflare → Caddy → Go directly, no
    Worker) returns 200 with the expected
    `X-Requests-Used`/`X-RateLimit-Remaining`/`X-Plan` headers; an invalid key
    returns 401; a request over the per-minute limit returns 429; a real
    Playground request through Rails (step 2's repointed URL) succeeds. Confirm
    `CF-Connecting-IP` is correctly read as the client IP by Go's IP-lookup
    endpoint post-cutover (this is the first time
    `trusted_proxies`/`trustedproxy.go` is exercised against real
    Cloudflare-fronted traffic, not local/test traffic).

11. **Let the direct-cutover traffic sit observed-healthy for a short bake
    period (a few hours to a day; a human/on-call judgment call, not a fixed
    SLA) before starting Phase 7b.** The Worker route being merely disabled, not
    deleted, in step 8 is what makes this bake period meaningful — rolling back
    is re-enabling the route, not restoring deleted code. Phase 7b deletes the
    Worker code and the KV/D1 stores it depends on; once that happens, this
    rollback path no longer exists.

**Exit criteria for 7a:** Rails' internal API path repointed and verified
independent of `internal.requiems.xyz`'s AOP gate; `api.requiems.xyz` routing
verified working without the Worker; AOP verified rejecting direct-origin
traffic by TLS alert, not just connection failure; `X-Backend-Secret` gate
removed from Caddy; Worker route disabled in Cloudflare (not yet deleted), 100%
of `api.requiems.xyz` traffic goes Cloudflare → Caddy → Go with no Worker in the
path; WAF rules reviewed; smoke test passes; a bake period has elapsed with no
observed regression.

### Phase 7b — Delete the Workers, KV, D1, and their Rails-side sync machinery (depends on 7a's exit criteria actually landing, including the bake period — not just being staged)

**Items 2 and 3 below require the same live Cloudflare access as Phase 7a — this
is not unattended repo cleanup.** Items 4-7 are repo-only and safe to automate
once items 1-3 have landed.

1. **Disable the still-scheduled D1-reading job before deleting D1**, to avoid a
   window of noisy job failures against a database that no longer exists. Remove
   the `sync_d1_usage` entry from **both** `apps/dashboard/config/recurring.yml`
   and `apps/dashboard/config/sidekiq_schedule.yml` — do this before item 2's
   deletion, not after. Production runs Sidekiq as its queue adapter
   (`config/environments/production.rb`'s `queue_adapter = :sidekiq`, Kamal's
   dashboard accessory runs `bundle exec sidekiq` — no Solid Queue supervisor
   process is started), so `sidekiq_schedule.yml`'s 3-minute cron is very likely
   the one actually firing in production and `recurring.yml`'s 5-minute entry
   likely never runs there — confirm against the live Sidekiq process before
   relying on that, but plan to edit both files regardless (dead config in the
   unused one is still worth cleaning up, and this resolves the pre-existing
   cadence-drift smell either way).
2. **Delete Cloudflare KV namespace and D1 database** (`requiem-usage`) via
   `wrangler`/dashboard, now that the export-backup snapshot exists, the sync
   job is disabled (item 1), and no traffic depends on them. **This requires a
   human with live Cloudflare access present and confirming before running** —
   it's as irreversible as anything in Phase 7a, even though the data itself is
   disposable per this plan's premise; the act of deleting a live Cloudflare
   resource still warrants an explicit go-ahead, not just "the data doesn't
   matter so just run it."
3. **Delete `apps/workers/auth-gateway`, `apps/workers/api-management`,
   `apps/workers/shared`** entirely, including their `wrangler.toml` bindings,
   and confirm any Wrangler-side secrets for these two Workers (e.g.
   `API_MANAGEMENT_API_KEY`, deploy tokens set via `wrangler secret
   put`) are
   actually gone from Cloudflare once the Workers themselves are deleted — don't
   assume deleting the Worker automatically clears its secrets, verify it.
4. **Delete the Rails-side D1 sync machinery**, now genuinely dead (not just
   uncalled, since D1 no longer exists):
   `apps/dashboard/app/services/d1_sync_service.rb`,
   `apps/dashboard/app/jobs/sync_d1_usage_job.rb` (the scheduler entries were
   already removed in item 1). Leave `AggregateDailyUsageJob` and
   `ExpirePromotionalSubscriptionsJob` untouched — they read `usage_logs`/
   `subscriptions` in Postgres directly, unaffected by D1's removal.
5. **Remove CI/CD references**: delete the `worker-test`/ `api-management-test`
   jobs (confirm the exact job key via grep — it is `api-management-test`, not
   `worker-management-test`) and their `paths-filter` blocks from
   `.github/workflows/ci.yml` (~lines 50-55, 300-380, including the
   `needs: [..., worker-test]` dependency and the result-check step around line
   380); delete `.github/workflows/_worker-ci.yml`; remove the three
   `apps/workers/*` entries from `.github/dependabot.yml`.
6. **Remove Worker containers from `infra/docker/docker-compose.dev.yml`**: the
   `auth-gateway`/`api-management` dev services, `wrangler_state`/
   `auth_gateway_modules`/`api_management_modules`/`shared_modules` volumes, and
   the `apps/workers/shared` bind mount.
7. **Remove now-fully-dead env vars** referenced only by the deleted Workers:
   the `apps/workers/auth-gateway/.env.example`-scoped `CLOUDFLARE_*` entries
   (left alone in Phase 6b specifically because the Worker was still deployed
   then — that condition no longer holds), plus `BACKEND_SECRET` if any
   reference survived 7a.

**Exit criteria for 7b:** D1-reading Sidekiq job disabled before KV/D1 deletion
(order verified); KV namespace and D1 database deleted in Cloudflare, with a
human confirming that step; `apps/workers/` directory no longer exists in the
repo; associated Wrangler secrets confirmed gone;
`d1_sync_service.rb`/`sync_d1_usage_job.rb` deleted with no scheduler still
referencing them; CI green with the Worker jobs removed (not skipped — actually
absent); `docker compose -f
infra/docker/docker-compose.dev.yml config --quiet`
succeeds with no Worker services.

### Phase 7c — Docs, tests, and remaining loose ends (depends on 7b; can start once 7b's deletions are staged even if not yet merged)

1. **Delete `docs/core/auth-gateway.md` and `docs/core/api-management.md`**
   entirely.
2. **Rewrite the Worker-era sections**: `docs/core/architecture.md` (full
   diagram + prose rewrite to Go-native auth/rate-limit/usage, no KV/D1),
   `docs/core/infrastructure.md` ("Architecture Components" diagram, Redis
   section — expand to document the `apikey:*`/`ratelimit:*`/`usage:*` Redis
   namespaces now that they're real), `docs/core/deployment.md` (delete "Part 5:
   Deploy Cloudflare Workers", add the AOP setup steps from 7a in its place),
   `docs/core/rails-app.md` ("Cloudflare Integration" section — describe direct
   Postgres key CRUD and Go-native usage writes; the "(separate schema)" line
   the audit flagged independently is already fixed in the current tree,
   confirmed by this plan's review pass — don't reintroduce it while rewriting
   this section). Update `agents.md` at the repo root to drop the Auth
   Gateway/API Management component listings.
3. **Full-text sweep** for `auth-gateway`, `api-management`,
   `Cloudflare
   Worker`, `KV`, `D1` across `docs/` (excluding `docs/plans/`
   and `docs/audits/`, which are historical record and should not be rewritten)
   to catch anything not listed above — the audit's own Documentation Audit
   table says this list isn't exhaustive.
4. **Also fix the two stale references Phase 6a explicitly deferred**:
   `docs/core/api-management.md` is being deleted in item 1, so its stale
   `/v1/text/words/define` references go with it;
   `apps/workers/api-management/src/__tests__/analytics.test.ts` is being
   deleted in 7b anyway.
5. **Update `tests/integration/src/suites/gateway.test.ts` and
   `tests/load/scenarios/rate-limit.ts`**: both already hit the public
   `API_BASE_URL`, so confirm that env var points at the right place
   post-cutover (Cloudflare-fronted `api.requiems.xyz`, same hostname as before
   — cutover doesn't change the public URL). Update the Worker-referencing
   comments/docblocks (`gateway.test.ts`'s "full Worker → Backend flow"
   docblock, `rate-limit.ts`'s "Auth Gateway (Cloudflare Worker)" comment and
   its `apps/workers/shared/src/config.ts` citation — repoint that citation to
   wherever plan-tier limits now live, e.g. the `plans` table / Go's rate-limit
   middleware) so they describe the current system, not a deleted one. Confirm
   no assertion in either file actually depends on Worker-only behavior (early
   read suggests they don't — both test observable HTTP behavior against the
   gateway URL) and fix any that do.
6. **Provision production `PLAYGROUND_API_KEY`** (checklist item already tracked
   in `docs/core/v2-deployment-playbook.md`): confirm with whoever owns
   production access before running this — it's a real-infra action (creating a
   live `ApiKey` row and setting a production env var), distinct from the rest
   of this phase's code/doc changes even though the "direct cutover, no notice
   period" framing applies to traffic routing, not to skipping a check before
   touching prod credentials. If already provisioned by the time this phase
   runs, just check the box.
7. **Confirm the DB row cleanup is actually applied where relevant.** Per the
   "no users, no traffic" note: any lingering test/dev `credit_usage`-derived
   rows in Postgres `usage_logs`/`daily_usage_summaries`/
   `credit_adjustments`/`audit_logs` that only exist because of pre-cutover
   testing/seeding can be truncated freely — do **not** touch Go's
   reference/data tables (`advice`, `quotes`, `words`, `bin_data`,
   `inflation_data`, `iban_countries`, `commodity_price_history`, `exercises`,
   `swift_codes`, `counters`). State explicitly in this phase's Final Notes
   whether any truncation was actually performed, and what was truncated, so
   it's not ambiguous later whether "the DB was cleaned" means test noise or
   actual reference data.

**Exit criteria for 7c:** `docs/core/{auth-gateway,api-management}.md` deleted;
`architecture.md`/`infrastructure.md`/`deployment.md`/
`rails-app.md`/`agents.md` updated; repo-wide grep for `auth-gateway`,
`api-management`, `Cloudflare Worker`, `KV`, `D1` outside `docs/plans/` and
`docs/audits/` returns nothing describing them as live; `tests/integration`/
`tests/load` pass against the post-cutover `API_BASE_URL` with Worker-specific
language removed; `PLAYGROUND_API_KEY` confirmed provisioned; DB cleanup status
stated explicitly in Final Notes.

## Explicitly out of scope for this plan

- Any new multiplier/pricing policy beyond what Phase 6a already restored parity
  for.
- A dedicated/second Redis instance, logical-DB split, or real load-based
  retuning of pool sizes/cache TTLs/Redis `maxmemory` — still gated on real
  traffic existing (audit Open Question 5), unchanged by this phase.
- Revisiting the Phase 2 decision to write `usage_logs` synchronously
  per-request instead of via a batched Redis flush. This was a deliberate,
  documented deviation from the audit's "zero synchronous Postgres ops"
  performance target, accepted only because there's no real traffic yet. It is
  worth flagging as a **future** performance item once real RPS numbers exist
  (audit Open Question 5) — this phase does not change it, since doing so is
  unrelated to retiring the Workers and would conflate two independent pieces of
  work.
- Rotating `BACKEND_SECRET`/any other credential as a security hygiene exercise
  beyond simply removing it once unused.
- Building a general-purpose Postgres-backed `plans`-table config path for the
  Worker's config (`apps/workers/shared/src/config.ts`'s plan-tier limits) —
  moot once the Worker is deleted in 7b; Go already reads limits from the
  `plans` table via `PlanCache`, so there's nothing left to migrate here.

## Open questions worth resolving before or during this session

- Confirm against the live Sidekiq process which of `recurring.yml` vs
  `sidekiq_schedule.yml` is actually driving `SyncD1UsageJob` in production
  before Phase 7b item 1 — production's queue-adapter config strongly suggests
  it's `sidekiq_schedule.yml` (see that item's own reasoning), but verify rather
  than assume before relying on it.
- Confirm whether the Cloudflare zone has any WAF/rate-limit rules configured
  outside the Worker today (audit Open Question 3, still unanswered) — needed to
  actually execute Phase 7a item 9, not just acknowledge it.
- Confirm live Cloudflare dashboard + VPS SSH access is available in the session
  that runs Phase 7a before starting — items 5 and 8 (AOP toggle, Worker route
  disable) cannot be done from the repo alone, and item 4's
  Caddy-redeploy-to-toggle-flip timing needs that access lined up in advance,
  not found mid-step.
- Decide the `api.requiems.xyz` routing mechanism (Cloudflare Origin Rule vs. a
  new Caddy site block, Phase 7a item 3) before the session starts, if possible,
  so it isn't designed live under time pressure right before the Worker route is
  disabled.
- Confirm the actual private-network address to repoint `INTERNAL_API_URL` to
  for Phase 7a item 2 (Rails→Go internal path) — this plan identifies the
  problem and the shape of the fix but not the exact address, which depends on
  the current Kamal/VPS network layout.

## Final Notes

### Session status — 2026-08-22

Phase 7a was not started beyond read-only preparation. The required live
Cloudflare dashboard access was confirmed by the project owner. A direct VPS IP
was subsequently supplied, but SSH authentication rejected the supplied
credential twice; retries were stopped. The Rails internal URL, Caddy
deployment, AOP toggle timing, direct-origin TLS-alert verification,
Worker-route disable, WAF review, and post-cutover smoke test were therefore
deferred rather than guessed or simulated.

Executed:

- Read this plan in full and inspected the prepared AOP branch and relevant
  Kamal/Caddy configuration.
- Installed Cloudflare's published agent skills and registered the five
  Cloudflare MCP servers for Codex; OAuth setup completed for the registered
  servers. No repository commit or PR was created for this environment setup.
- Per the project owner's explicit direction, skipped 7a step 1's optional KV/D1
  backup export. No backup, reconciliation, backfill, or database-row cleanup
  was performed.

Deferred:

- All of Phase 7a, pending working VPS/SSH authentication or equivalent live
  origin-management access and a decision on the exact private Rails→Go address.
  The intended public shape is `api.requiems.xyz` → Cloudflare → the Go origin
  currently represented by `internal.requiems.xyz`, with Rails using a private
  path that bypasses AOP.
- Phases 7b and 7c. KV/D1 deletion remains explicitly human-gated and was not
  attempted. No Worker, KV, D1, Rails sync, CI, compose, documentation, or test
  files were deleted or rewritten.

Open-question answers:

- Live access: Cloudflare dashboard access was confirmed and the VPS IP was
  supplied, but usable direct VPS SSH access was not available because the
  supplied credential was rejected.
- Sidekiq scheduler: not checked against production because 7a did not reach the
  live VPS and 7b was not started.
- Cloudflare WAF/rate-limit rules: not reviewed; still unanswered.
- `api.requiems.xyz` routing mechanism: not finalized. The direct Caddy
  `api.requiems.xyz` site-block approach remains the likely implementation, but
  it must be confirmed against the live DNS/Cloudflare setup before the Worker
  route is disabled.
- Rails `INTERNAL_API_URL`: not changed; the exact private-network address
  remains dependent on the live Kamal/Docker layout.
- Real commits/PRs: none were created in this session.
