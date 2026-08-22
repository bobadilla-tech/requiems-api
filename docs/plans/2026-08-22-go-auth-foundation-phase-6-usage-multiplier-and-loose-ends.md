# Go Auth Foundation — Phase 6: Usage-Multiplier Correctness & Migration Loose Ends

Continuation of `docs/plans/2026-08-21-go-auth-foundation-phase-0-1.md`,
`-phase-2.md`, `-phase-3-4.md`, `-phase-5.md` (undocumented but fully shipped —
see Context), and `2026-08-22-go-auth-foundation-standing-issues-hardening.md`
(PR #966 security hardening, also shipped). This plan was written after
verifying every prior plan doc's claims against the current tree directly —
grep, `git log`, and direct file reads, not the docs' own "Shipped" sections —
specifically because Phase 5's own doc was found to be fully executed in code
with no "Final notes" ever written up for it. Nothing in this plan assumes a
prior doc's stated status is still accurate; everything cited below was
re-verified today, 2026-08-22.

## Context — what is actually still open

**Fully shipped, reverified today, not revisited by this plan:** Go-native
API-key auth (candidate-by-`key_prefix` + bcrypt, Redis-cached, fail-closed on
key-existence), Redis rate limiting + usage/quota middleware, `plans` table,
Rails-local `requiem_`-prefixed key generation with collision retry, all
Cloudflare-sync callbacks deleted from `ApiKey`/ `Subscription`,
`BackendSecretAuth` removed from Go, the auth-cache prefix-guessing exposure
(candidate-only cache, bcrypt re-verified on every hit — confirmed at
`apps/api/platform/middleware/apikeyauth.go:120-126`, this was fixed by the PR
#966 hardening pass, not left open as three earlier plan docs' "carried-forward
gap" notes claimed), Go's `callerIP()` trust boundary
(`apps/api/platform/httpx/trustedproxy.go`, shared by all three IP-reading
endpoints), Rails' `trusted_proxies`
(`apps/dashboard/app/lib/trusted_proxy.rb`), the MRR write-path fix
(`subscriptions.plan` now populated from `determine_billing_cycle` in the
LemonSqueezy webhook controller, with test coverage in
`analytics_revenue_service_test.rb` asserting yearly-vs-monthly divergence), the
dead-code sweep (Stripe columns, `solid_cache_entries`, Pundit's
`app/policies/`, all three missing FKs — one migration,
`20260822000000_dead_code_and_fk_cleanup.rb`), and Go's test-database isolation
(`TEST_DATABASE_URL`, both `app_test.go` and `apikeyauth_test.go`).

**Confirmed still open, verified against the live tree just now:**

1. **The Worker's per-endpoint usage-multiplier config is silently a no-op for
   the two endpoints it was written for, on both the Worker and Go paths — a
   real, previously-unflagged billing-accuracy bug.**
   `apps/workers/shared/src/config.ts`'s `ENDPOINT_MULTIPLIERS` map only lists
   two entries: `GET /v1/text/words/define` and `GET /v1/text/words/synonyms`,
   both at 2x. Those paths do not exist anymore — the services-reorganization
   work (`docs/plans/2026-04-15-services-reorganization.md`) renamed them to
   `GET /v1/text/dictionary/{word}`
   (`apps/api/services/text/words/
   transport_http.go:35`) and
   `GET /v1/text/thesaurus/{word}`
   (`apps/api/services/text/thesaurus/transport_http.go:20`). `config.ts`'s
   `getRequestMultiplier()` does an exact match first, then a prefix match built
   from the _same_ map — since the map's keys were never updated, both the
   Worker's own live billing and any future Go port would silently charge 1
   credit instead of the intended 2 for a dictionary or thesaurus lookup. This
   has nothing to do with the Worker-retirement timeline; it's a config-drift
   bug in currently-live billing logic, independent of Phase 3 items 6–7 below.
2. **`httpx.UsageCounter` (`apps/api/platform/httpx/handler.go:33-36`) is still
   implemented by exactly one endpoint** —
   `services/systems/data_integrity/input_validate_batch/service.go` —
   reconfirmed by grep just now, matching Phase 2's own count. Every other
   endpoint's `usage_logs.credits_used` defaults to 1, including the two
   dictionary/thesaurus endpoints above once their multiplier is fixed.
3. **Phase 3 items 6–7 (Cloudflare Authenticated Origin Pulls / origin firewall,
   and deleting `apps/workers/{auth-gateway,api-management,shared}` plus the KV
   namespace and D1 database) are still not started.**
   `infra/caddy/Caddyfile:36-40` still gates `internal.requiems.xyz` on
   `X-Backend-Secret`; `apps/workers/{auth-gateway,api-management,shared}` still
   exist and are still deployed. **Correction to Phase 5's own Context section,
   which claimed the Worker "strips `requiems-api-key` before proxying"** —
   rereading `apps/workers/auth-gateway/src/http.ts:16-28` directly shows
   `filterHeaders()` forwards every header except `cf-*`, `connection`, and
   `keep-alive`, explicitly preserving `requiems-api-key` (the in-code comment
   even says so). The Worker-proxied path is not currently broken; Phase 5's
   note to the contrary was stale by the time it was written, not a
   currently-accurate description. This doesn't change the scope of items 6–7
   (still human-gated, still needs live Cloudflare/VPS access), just corrects
   why they're being deferred — it's readiness for full cutover, not an active
   outage.
4. **Stale `rq_live_`/`rq_test_` key-format literals remain in eight live
   (non-historical-plan-doc) files** — an independent full-repo grep found four
   more than an earlier draft of this plan listed, so treat this list, not the
   earlier partial one, as authoritative:
   `apps/dashboard/docs/app-config.md:81`,
   `apps/mcp/src/server.test.ts:178,207`, `tests/integration/README.md:21,33`,
   `tests/integration/src/suites/gateway.test.ts:25`,
   `tests/integration/.env.example:11,12`,
   `apps/dashboard/config/locales/en/home.en.yml:527`,
   `apps/dashboard/config/locales/es/home.es.yml:551`,
   `apps/dashboard/config/locales/fr/home.fr.yml:532`. This was named as a
   leftover in `-phase-3-4.md`'s own follow-ups, named again in `-phase-5.md`
   item 6, and still hasn't been touched — third consecutive plan doc to note it
   without fixing it, and the first two attempts at listing the affected files
   were themselves incomplete.
5. **No backfill was ever run for pre-existing `subscriptions.plan` rows.** The
   write-path fix (`determine_billing_cycle`) only populates `plan` for
   subscriptions created or updated after it shipped; any row from before that
   still has `plan = nil` and silently prices as monthly. `-phase-5.md`'s own
   item 1 flagged this as needing LemonSqueezy dashboard/API access to verify,
   and no backfill script exists anywhere in the repo (confirmed by find). Given
   there are no real customers yet, this is very likely moot in practice, but
   nobody has confirmed that explicitly — see Approach below.
6. **The doc corpus itself is now inconsistent about the auth-cache
   prefix-guessing issue's status — noted here, not fixed by this plan.**
   `phase-2.md`, `phase-3-4.md`, and `phase-5.md` all still contain "known,
   accepted gap"/"third consecutive phase doc to re-defer this" language
   describing that exposure as live and unmitigated. It isn't (see this
   Context's opening paragraph) — the standing-issues-hardening doc claims to
   have updated "the Phase 0–1, Phase 2, and Phase 3–4 planning documents," but
   only `phase-0-1.md`'s text actually reflects the fix; the other two, plus
   `phase-5.md` (written after the fix shipped), still read as if it's open.
   Left uncorrected here to keep this plan's scope to code and the items above —
   if a future session touches those docs for another reason, fix this in
   passing.

## Approach

### Phase 6a — Fix the endpoint-multiplier config drift, then port it to Go

1. **Fix `apps/workers/shared/src/config.ts`'s `ENDPOINT_MULTIPLIERS` map.**
   Change the two stale entries to match the current routes. Because
   `getRequestMultiplier()` does an exact `${method} ${pathname}` match first
   and only falls through to prefix matching built from the same map's entries,
   the replacement keys need to be prefixes that work through that fallback for
   a dynamic `{word}` segment — e.g. `GET /v1/text/dictionary` and
   `GET /v1/text/thesaurus` (no trailing word), relying on
   `pathname.startsWith(routePath)`. Confirm this doesn't accidentally also
   match `/v1/text/dictionary` itself if such a bare route exists (it doesn't,
   per `transport_http.go` above — only `/dictionary/{word}`) and doesn't
   collide with `/words/batch` or `/thesaurus/batch` (different path segment,
   `/words/batch` vs `/dictionary/{word}` — confirm the batch routes are
   intentionally excluded from this multiplier, matching the map's historical
   intent of "define/synonyms cost more, batch does not", not an oversight; the
   method check in `getRequestMultiplier` also prevents any collision on its
   own, since this multiplier only ever matches `GET` and the batch routes are
   `POST`). Update `auth-gateway`'s existing multiplier tests — the real file is
   `apps/workers/auth-gateway/src/__tests__/config.test.ts`, which currently
   hardcodes assertions against the stale `/v1/text/words/define`/`/synonyms`
   paths (lines 67-68, 98-149) — to assert against the corrected paths, and add
   a regression test asserting the _old_ stale paths no longer match anything
   (catches a future rename from silently reintroducing this exact bug).

2. **Port the fix to Go — via a manual header set, not `httpx.UsageCounter`.**
   **Verified correction to an earlier draft of this plan, which proposed a
   mechanism that doesn't apply here:**
   `apps/api/services/text/words/
   transport_http.go`'s `/dictionary/{word}`
   handler and `apps/api/services/
   text/thesaurus/transport_http.go`'s
   `/thesaurus/{word}` handler are both raw `http.HandlerFunc`s that call
   `httpx.JSON()` (`httpx/httpx.go:38`) directly — `httpx.JSON` never checks
   `UsageCounter`. Only `httpx.Handle()` (`httpx/handler.go:73-74`) does that
   type-assertion, and these two handlers don't go through `Handle` (its
   request-binding is JSON-body-only and has no path for `chi.URLParam`-sourced
   input, so converting them onto `Handle` is a real refactor, not required
   here). The correct, verified fix: inside each handler, call
   `w.Header().Set("X-Usage-Count", "2")` before calling `httpx.JSON` (headers
   must be set before the response is written). `usage.go`'s `responseCredits()`
   (`middleware/usage.go:246-251`) reads the `X-Usage-Count` header directly off
   whatever's present on the `http.ResponseWriter`, regardless of how it got set
   — confirmed this does not require going through `UsageCounter`/`Handle` at
   all. Do not try to route these two handlers through `httpx.Handle` or
   implement `UsageCounter` on their response types; that mechanism doesn't fit
   `chi.URLParam`-based handlers as currently structured.

3. **Test coverage.** A Go integration test hitting `/v1/text/dictionary/{word}`
   and `/v1/text/thesaurus/{word}` through the full middleware chain
   (`APIKeyAuth` → rate limit → usage quota → handler) and asserting the
   response's `X-Usage-Count` header is `2` and the resulting `usage_logs` row
   has `credits_used = 2` — following this series' existing
   real-Postgres/real-Redis convention. **Do not assert `/words/batch`/
   `/thesaurus/batch` bill flat 1 per request — verified they don't, and never
   have.** Both already use `httpx.HandleBatch()` (`words/transport_http.go:51`,
   `thesaurus/transport_http.go:36`), which unconditionally sets `X-Usage-Count`
   to the batch's item count (`handler.go:122`) — universal, pre-existing,
   tested behavior across every batch endpoint in this codebase, including
   `input_validate_batch`, this series' own cited reference pattern (its
   `UsageCount()` also returns the item total, not 1). The correct assertion is
   that this phase leaves that per-item batch billing **unchanged**: add a test
   confirming a multi-item batch call still reports `credits_used` equal to the
   item count, not a different number, so a future edit to these same two files
   doesn't accidentally flatten batch billing to 1 while touching them for the
   GET-route fix above.

**Explicitly out of scope for 6a:** any broader multiplier scheme for the other
~218 endpoints. The Worker's own config only ever defined multipliers for these
two routes — this phase restores parity with what the Worker's config actually
says (once its own drift is fixed), it does not invent new multiplier policy for
endpoints nobody has ever flagged as expensive. If a future session wants to
price other endpoints above 1x, that's a product decision made explicitly, not
something to infer here. Also out of scope:
`docs/core/api-management.md:338,443` and
`apps/workers/api-management/src/__tests__/analytics.test.ts:169` still cite the
old `/v1/text/words/define` path as example/test data — neither file is part of
`ENDPOINT_MULTIPLIERS`'s call graph (confirmed: only `config.ts`,
`config.test.ts`, and `proxy.ts` reference it), so these two references don't
affect billing correctness; leave them for the eventual doc sweep gated on Phase
3 items 6–7, don't fold them into this item.

### Phase 6b — Migration loose ends (small, independent, no ordering constraint between items)

1. **Fix the eight stale `rq_live_`/`rq_test_` literals** (see Context item 4
   for the authoritative, corrected file list — an earlier draft of this plan
   named only four and missed `tests/integration/.env.example` and all three
   locale files). Update `apps/dashboard/docs/app-config.md:81`'s documented
   default to match `PLAYGROUND_API_KEY`'s real current default
   (`requiem_notprovisioned0000000000` per `app_config.rb:185`, or whatever
   value item 3 below leaves it at — don't just relabel the old literal, confirm
   the doc describes the real current behavior once item 3 lands).
   `apps/mcp/src/server.test.ts:178,207`,
   `tests/integration/{README.md,.env.example,src/suites/gateway.test.ts}` are
   test fixtures/examples — update them to `requiem_`-shaped placeholder values
   (keeping `gateway.test.ts:25`'s "obviously invalid" intent — the point of
   that literal is that it's rejected, so it just needs to be `requiem_`-shaped
   _and_ invalid, e.g. wrong checksum/length, not a real key).
   `config/locales/{en,es,fr}/home.*.yml`'s
   `example_code:
   rq_live_1234567890abcdef` strings are user-facing
   FAQ/marketing copy — update to a `requiem_`-shaped example, consistent with
   the FAQ fix Phase 3 already made elsewhere in the same files (per
   `-phase-3-4.md`'s Final Notes bug #4).

2. **Resolve the MRR-backfill open question explicitly, with an explicit stop
   condition — don't let this item silently expand into a full backfill
   project.** Query production (or confirm via whoever has DB/LemonSqueezy
   access) whether any `subscriptions` row with a non-null `plan_name` and
   `status IN ('active','trialing')` predates the `determine_billing_cycle` fix
   and still has `plan IS NULL`. **If the answer is zero rows** (plausible
   pre-launch, per this plan's opening premise that there are no real customers
   yet), write that finding down — a one-line note in this doc's Final Notes
   once shipped, or a short comment in `analytics_revenue_service.rb` — so the
   next person doesn't reopen this as a mystery, and stop there. **If there are
   affected rows, do not execute the backfill inside this session.** That work
   (fetching each row's billing interval from LemonSqueezy's API, flagging
   anything unretrievable for manual reconciliation, verifying the total against
   LemonSqueezy's own dashboard MRR figure before trusting it) is real,
   separately-scoped engineering with its own risk profile — Phase 5's item 1
   already sized it as such. Treat a nonzero result the same way this plan
   treats Phase 3 items 6–7: stop, write up what was found (row count, date
   range, whether LemonSqueezy history looks retrievable), and hand it to a
   follow-up session rather than executing an unscoped financial-data backfill
   under this item's momentum.

3. **Confirm (or provision) the production `PLAYGROUND_API_KEY`.**
   `apps/dashboard/app/lib/app_config.rb:185`'s default,
   `requiem_notprovisioned0000000000`, is not a placeholder that was
   deliberately left invalid for tests — its name and the code path it feeds
   (`ApiProxyService`, the live public playground and demo forms) strongly
   suggest Phase 5's own item 2 (provision a real, dedicated system API key for
   the playground) was written up but the actual provisioning step — creating
   the `ApiKey` row and setting the real `PLAYGROUND_API_KEY` environment value
   in production — was never run. Confirm this directly (check whether
   production's `PLAYGROUND_API_KEY` env var is still the `notprovisioned`
   default) before assuming the playground fix from Phase 5 is actually live. If
   it's still unprovisioned, run Phase 5 item 2's own provisioning step now:
   create the dedicated `ApiKey` via the normal `ApiKeyGenerator` path on a
   real, bounded (not `enterprise`) plan, and set the production env var. This
   is the same class of "code shipped, ops step never completed" gap this plan
   already treats as first-class (see item 2 above and Context item 5) — it just
   wasn't caught until this rewrite. **Setting a production environment variable
   and provisioning a live API key is a real-infra action, unlike the rest of
   this item's own code changes** — confirm with whoever owns production access
   before setting it, the same way this plan asks for confirmation before Phase
   6b item 2's database query.

4. **Delete the now-fully-dead Cloudflare API-key-management service, since
   nothing calls it anymore — independent of Phase 3 items 6–7's live-infra
   gate.** `apps/dashboard/app/services/cloudflare/api_management_service.rb`
   has zero callers left in the codebase (confirmed by grep) — its only callers
   were the `ApiKey`/`Subscription` Cloudflare-sync callbacks Phase 3 already
   deleted. Deleting this one file touches no live infrastructure (the Worker it
   used to call keeps running regardless; nothing in Go or Rails invokes this
   file today) — it's the same category of change as the already-shipped
   Stripe-columns/`solid_cache_entries`/Pundit sweep, just missed by that sweep
   because this file wasn't literally named in Phase 4's original item list.
   **Do not delete `d1_sync_service.rb` or `sync_d1_usage_job.rb` as part of
   this item** — those are still actively scheduled
   (`config/recurring.yml`/`config/sidekiq_schedule.yml`) and pull real data
   from the still-live D1 database; they stay untouched until Phase 3 item 7's
   actual Worker retirement. Also remove the already-dead
   `CLOUDFLARE_ACCOUNT_ID`/`CLOUDFLARE_KV_NAMESPACE_ID`/`CLOUDFLARE_API_TOKEN`
   template entries from `infra/docker/.env.example` in the same PR — the audit
   already flagged these as safe to remove now, independent of migration timing,
   and no phase doc has picked them up yet.

5. **Worker-retirement prep artifacts — zero live-infra-access work, explicitly
   deferred in `-phase-5.md` as "left out, not overlooked."** Two small,
   self-contained deliverables that shrink the eventual human-gated Phase 3
   items 6–7 session without touching anything live:
   - Write the KV/D1 backup-export script (`wrangler kv key list`/`get` for a
     bulk export, `wrangler d1 export`) as a runnable, reviewable script
     committed under `apps/workers/auth-gateway/scripts/` or similar — ready to
     run the moment someone with Cloudflare access is in the room, not written
     from scratch during that session.
   - Stage the Caddy Authenticated-Origin-Pulls config change
     (`infra/caddy/Caddyfile`) as a reviewable diff/branch — the actual cert
     installation and DNS-side toggle still requires live Cloudflare dashboard
     access and is not applied here, but the Caddy-side config shape can be
     written and reviewed now. Do not merge or deploy this diff in this session
     — it's prep, and Phase 3 item 6's own ordering requirement (verify
     Authenticated Origin Pulls works _before_ removing the existing
     `X-Backend-Secret` gate) still applies whenever it's actually executed.
   - Both artifacts need a real correctness bar, not just "exists": the export
     script should be dry-run against a non-prod KV namespace/D1 database if one
     exists (or at minimum lint/type-check cleanly and be reviewed for the right
     `wrangler` invocation), and the Caddy diff should be validated structurally
     (`caddy validate` or equivalent against the staged config) before being
     called done — "committed" alone isn't a verifiable exit bar for either.

## Exit criteria

- `apps/workers/shared/src/config.ts`'s `ENDPOINT_MULTIPLIERS` map matches the
  current, post-reorg route paths; its own test suite covers the corrected paths
  and regression-tests the old stale ones.
- `GET /v1/text/dictionary/{word}` and `GET /v1/text/thesaurus/{word}` write
  `credits_used = 2` to `usage_logs` and set `X-Usage-Count: 2`, verified by a
  real end-to-end Go test, via a manually-set header (not `UsageCounter`) per
  the corrected mechanism in Phase 6a item 2; `/words/batch` and
  `/thesaurus/batch` remain unaffected — still per-item billing, exactly as they
  already were before this phase.
- Zero `rq_live_`/`rq_test_` literals remain outside historical plan/audit docs
  (grep across the repo, excluding `docs/plans/` and `docs/audits/`, returns
  nothing) — this now requires all eight files in Context item 4 to be fixed,
  not the four an earlier draft named.
- The MRR-backfill question has an explicit, written answer — either "zero
  affected rows, confirmed" or, if rows were found, a written handoff to a
  follow-up session (not a backfill executed inside this one — see item 2's stop
  condition).
- Production's `PLAYGROUND_API_KEY` is confirmed already provisioned, or freshly
  provisioned this session — not left at the `requiem_notprovisioned...`
  default.
- `apps/dashboard/app/services/cloudflare/api_management_service.rb` and the
  dead `CLOUDFLARE_*` `.env.example` template entries are deleted.
- KV/D1 backup-export script exists and is committed; Caddy AOP config diff
  exists as a reviewable, unmerged artifact.

## Explicitly out of scope for this plan

- **Phase 3 items 6–7** (Cloudflare Authenticated Origin Pulls / origin firewall
  live configuration, DNS/proxy-mode changes, and deleting
  `apps/workers/{auth-gateway,api-management,shared}` plus the live KV namespace
  and D1 database). Still human-gated, still requires live Cloudflare
  dashboard/VPS access in the room for each irreversible step, per Phase 3's own
  ordering requirements. This plan's Phase 6b item 5 prepares artifacts for that
  session; it does not start it. (Phase 6b item 4 deletes one already-dead file
  that used to be part of this cleanup — that deletion needs no live access and
  isn't gated on items 6–7, see that item's own reasoning.)
- The full doc sweep deleting `docs/core/auth-gateway.md`/`api-management.md`
  and rewriting `architecture.md`/`infrastructure.md`/`deployment.md` — still
  correctly contingent on items 6–7 actually landing (the Workers are still
  running; deleting their docs now would describe a system that doesn't exist
  yet, per `-phase-5.md`'s own reasoning, which still holds).
- Any multiplier/pricing policy beyond the two routes the Worker's config
  already names — see Phase 6a's own "explicitly out of scope" note.
- A dedicated/second Redis instance or logical-DB split; real load-based
  retuning of placeholder pool sizes, cache TTLs, or the Redis `maxmemory` value
  — all still waiting on real traffic to exist (audit's Open Question 5,
  restated unchanged in every prior phase doc).
- Confirming `api-management`'s `/analytics/*` endpoints' caller — irrelevant
  until Phase 3 item 7 actually deletes them.

## Open questions worth resolving before or during this session

- Who can query production (or confirm there is no meaningful production data
  yet) to answer Phase 6b item 2's MRR-backfill question? If nobody has DB
  access in this session, state that explicitly in the PR rather than guessing
  "probably zero rows."
- Confirm at implementation time whether any other endpoint besides
  `/dictionary/{word}`/`/thesaurus/{word}` was ever intended to carry a
  multiplier > 1 — the Worker's `config.ts` comment block describes hypothetical
  future categories (AI/ML inference, external API calls) with no current
  entries; this plan takes the config's actual current content as the source of
  truth, not the aspirational comment.

## Final Notes

Executed 2026-08-22. All of Phase 6a and Phase 6b landed except the two items
with explicit human/infra gates, which were deliberately not pushed past their
stop conditions.

**Phase 6a (all three items, all on `feat/go-auth-foundations`):**

1. `apps/workers/shared/src/config.ts`'s `ENDPOINT_MULTIPLIERS` map fixed to the
   current `/v1/text/dictionary`/`/v1/text/thesaurus` prefixes; `config.test.ts`
   updated plus a regression test asserting the old
   `/v1/text/words/define`/`/synonyms` keys no longer match anything.
2. Go port: `X-Usage-Count: 2` set directly on the two GET handlers
   (`words/transport_http.go`, `thesaurus/transport_http.go`), per the plan's
   own corrected mechanism — not `httpx.UsageCounter`/`Handle`.
3. Real-Postgres/real-Redis integration test added to `app_test.go`, exercising
   the full `APIKeyAuth -> rate limit -> usage quota -> handler` chain. Building
   it surfaced a real test-isolation bug, fixed in the same commit: the
   `APIKeyAuth` Redis candidate cache is keyed by the 12-char key prefix and
   survives across separate `go test` process invocations, so a fixture reusing
   a hardcoded literal key could authenticate against a stale, already-deleted
   `api_key_id` left over from an earlier run. `seedAPIKeyFixture` is now
   parameterized (`seedAPIKeyFixtureWithKey`) and the new test generates a
   random key per run. Full `go test ./...` in `apps/api` passes (zero failures
   across every service package).

**Phase 6b:**

1. All eight stale `rq_live_`/`rq_test_` literals fixed (the corrected,
   authoritative eight-file list from this doc's Context item 4, not the four an
   earlier draft named). `apps/dashboard/docs/app-config.md`'s documented
   `PLAYGROUND_API_KEY` default corrected to what `app_config.rb` actually ships
   (`requiem_notprovisioned0000000000`) — item 3 below is still open, so this
   reflects current code, not a post-provisioning value. Repo-wide grep
   (excluding `docs/plans/` and `docs/audits/`) now returns zero matches.
2. **MRR backfill: resolved without running the production query.** The session
   had no production DB/LemonSqueezy credentials available, but the project
   owner confirmed directly, out-of-band, that there are no real customers yet —
   the scenario this item's own text flagged as "plausible pre-launch." Treated
   as the authoritative zero-rows answer per this item's exit criteria. No
   backfill code was written or run. If that changes before this is revisited,
   rerun the query this item originally specified (non-null `plan_name`,
   `status IN ('active','trialing')`, `plan IS NULL`, predating the
   `determine_billing_cycle` fix).
3. **PLAYGROUND_API_KEY: deliberately not touched — real-infra action, stop
   condition honored.** Setting a production env var and provisioning a live
   `ApiKey` needs production access confirmation this session didn't have.
   Tracked instead in new `docs/core/v2-deployment-playbook.md`, so it isn't
   forgotten at actual v2 deploy time.
4. `apps/dashboard/app/services/cloudflare/api_management_service.rb` deleted
   (zero callers, confirmed by grep; `zeitwerk:check` clean after).
   `CLOUDFLARE_ACCOUNT_ID`/`CLOUDFLARE_KV_NAMESPACE_ID`/`CLOUDFLARE_API_TOKEN`
   removed from `infra/docker/.env.example` only — confirmed
   `d1_sync_service.rb`/`sync_d1_usage_job.rb` (still live, still scheduled)
   don't read these, calling the api-management Worker's HTTP API instead. The
   still-running `apps/workers/auth-gateway/.env.example` entries for the same
   names were left alone (that Worker is still deployed).
5. **Both prep artifacts shipped, with a correctness bar beyond "exists":**
   - `apps/workers/auth-gateway/scripts/export-backup.ts` (on
     `feat/go-auth-foundations`, safe to merge — it's inert until someone runs
     it with `--remote`) — typechecks and lints clean, and was dry-run with
     `--local` against the shared dev KV/D1 (docker-compose's `wrangler_state`
     volume), correctly exporting all 8 seeded KV keys and the full D1
     schema+data. The dry-run caught two real invocation bugs before they'd have
     hit a live export: `d1 export` uses `-y`/`--skip-confirmation`, not
     `--yes`, and unlike `d1 execute`/ `kv key`, `d1 export` has no
     `--persist-to` flag at all.
   - The Caddy AOP diff (`infra/caddy/Caddyfile`,
     `infra/caddy/certs/cloudflare-origin-pull-ca.pem`,
     `infra/docker/docker-compose.yml`'s cert mount) is **on a separate,
     unmerged branch — `prep/caddy-authenticated-origin-pulls`, pushed to
     origin, not opened as a PR** — not on `feat/go-auth-foundations`. Reason:
     `feat/go-auth-foundations` is PR #966, which does get merged and deployed;
     landing a `require_and_verify` client-cert requirement there would 403 all
     Worker-to-origin traffic the moment it ships, since Cloudflare's dashboard
     toggle for Authenticated Origin Pulls isn't on yet. Structurally validated
     with `caddy validate` (adapts successfully, reports "Valid configuration")
     against a local copy with the vendored cert present at the referenced path.
     The vendored cert is Cloudflare's fixed, publicly-published global-AOP CA
     (`developers.cloudflare.com/ssl/static/authenticated_origin_pull_ca.pem`),
     not a secret. Whoever runs Phase 3 item 6 should: flip the Cloudflare
     dashboard toggle, confirm AOP works, merge that branch, then remove the
     `X-Backend-Secret` gate as its own follow-up — per this plan's own ordering
     requirement.

**Not touched, as instructed:** Phase 3 items 6-7 themselves (Cloudflare
dashboard/DNS changes, deleting the Workers or their KV/D1 stores).
