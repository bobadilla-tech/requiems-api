# Dashboard Overview & Usage Accuracy

Status: IMPLEMENTED in PR-0, PR-A, PR-B and PR-C (see Final notes). It was
drafted as v9 (planning only; rounds 1–8 of review folded in). Deviation from
`docs/plans/readme.md` ("save after work lands"): written up-front at the user's
request; **Final notes** is filled in as each PR ships.

## Context

User-visible symptoms on `/dashboard` (Overview):

1. "Requests used this month" shows 2.02K though the last activity is 3 months
   old; bar says "0% used".
2. 404 requests (`/v1/text/advice`, `/api/v1/text/advice`) appear in Recent API
   Activity.
3. The "Requests" column is only ever 0 or 1; the panel is dense and hard to
   read.

### Verified root causes

- **R1 — wrong "month".** `User#usage_this_month` (`app/models/user.rb:87`)
  counts rows since `subscription.current_period_start` (else calendar month).
  That column is not a monthly anchor: admin promotions set it to the grant time
  (`promotions_controller.rb:33`) and the expiry job leaves it;
  `subscription_created` sets `renews_at - 1.month`
  (`lemonsqueezy_controller.rb:93`, wrong for yearly); `payment_success` sets
  `Time.current` (`:217`). The screenshot (next billing 2027-04-16) fits an
  April-2026 anchor ⇒ ~6 months shown as "this month". _Unconfirmed against real
  data — a human must check with a read-only query; agents never touch prod._
- **R2 — Rails and Go disagree.** Go (`usage.go`) rolls the cycle monthly on the
  anchor's UTC day-of-month and `SUM(credits_used) WHERE used_at >= cycle` (no
  upper bound). Rails counts rows over a different window. Go's anchor
  (`apikeyauth.go:211-217`) =
  `COALESCE(<active|trialing subscription>.current_period_start, <authenticating key>.created_at)`.
  Rails ignores status. `subscriptions.user_id` index is **not unique**, so
  multiple rows are possible.
- **R3 — percent.** integer `.round` of 0.02%. `compact_number` renders 2,018 as
  "2.02K" and ignores locale.
- **R4 — 404s.** `UsageQuota` wraps all `/v1/*` incl. sub-router misses and logs
  every response that cleared quota (401 never reaches it; 429/503 return before
  logging). Go `headerCredits` is clamped to ≥1, so today any 404/4xx/5xx _costs
  1 credit_. Existing 0-credit 404 and `/api/v1/*` rows are legacy Worker/D1-era
  rows (Worker billed 0 for non-2xx). The "Requests" column is `credits_used`,
  i.e. billed/not billed.
- **R5 — public copy** says "each API call costs 1 request … resets at the start
  of each billing period" (`home.en.yml:385-396,534`, + es/fr). No
  docs/OpenAPI/MCP text describes error charging.

### Adjacent defects found by review (tracked, mostly deferred)

Security: Usage-page JS injects endpoint into `innerHTML` unescaped (XSS); CSV
export has formula injection and `.each` over full range + N+1. Go: no
`Recoverer` anywhere ⇒ a handler panic skips refund _and_ log (Redis
over-counts); multi-key free users get per-key cycles ⇒ two Redis counters;
`AddDate(0,1,0)` overflow for 29–31 anchors; cycle taken at request start but
`used_at` after handler (boundary skew); `plancache.go:73` unknown plan ⇒ nil
limits ⇒ **unlimited** (fail-open); `PlanConfig.fetch` silently maps unknown
plans to free (display); `HandleBatch` empty batch bills 1; dedup index drops
same-µs rows after Redis incremented; no auth-cache invalidation on subscription
changes (30 s TTL lag). Rails: Billing page has its own copy of the usage math
and warning banners; admin users/abuse controllers use a _third_ "month"
(calendar); admin analytics read `usage_logs` for scanner/abuse signal
(`common_errors`, failed-auth); `usage_alerts` setting sends nothing; summaries
job drops error/latency aggregates; usage-page JS stacks `turbo:load` listeners,
charts not dark-mode/a11y aware, labels hardcoded English; `requests_col` still
used by admin views; `_stat_card` has no progress slot and is used on the home
page.

## Approach

Ship in three focused PRs for the asked outcomes, a standalone security PR, then
deferred follow-ups. Each PR is safe in either Rails/Go deploy order
(independent Kamal jobs).

### PR-0 (standalone, ship first) — security

- Usage page: build endpoint labels with `textContent` (or escape). CSV: prefix
  cells starting `= + - @ \t \r` with `'`; `includes(:api_key)`. Tests: endpoint
  `"><script>` escaped; CSV injection cell.
- Prefer a server-side fix for testability: render the endpoint-usage list
  server-side (escaped by ERB) or add a JS/system test that injects
  `<img onerror>` as an endpoint and asserts no element is created. CSV test is
  a plain controller test.
- Acceptance: both tests fail before / pass after.

### PR-A — correct numbers (Rails only; fixes R1, R2-display, R3)

- `BillingCycle` (`app/lib/billing_cycle.rb`): **mirror Go exactly as it stands
  today**: anchor = active/trialing subscription start (deterministic pick if
  several: active first, newest; a subscription row with NULL
  `current_period_start` contributes its plan but falls through to the key
  anchor, as Go's COALESCE does) else the key anchor — for the dashboard (no
  authenticating key) use the user's **oldest active, non-revoked key**
  (`api_keys.active_keys.minimum(:created_at)`; Go can only anchor on a key that
  can authenticate), and document that multi-key free users may legitimately
  differ from Go until the Go anchor fix (deferred D1). `start` = most recent
  clamped anchor-day ≤ now (UTC); `next_start` = next clamped anchor-day (not
  `+1.month`).
- Parity vectors at
  `apps/api/platform/middleware/testdata/billing_cycle_vectors.json` (covered by
  CI path filters; Rails test reads it from the repo in CI only; not needed in
  the Kamal image). Cases: Jan 31→Feb, leap, anchor in future, nil, status
  filter, day-30 months, `next_start`.
- `UsageSummary` (`app/services` or `app/lib`): computed once per request, **no
  cache on Overview** (avoid counter lagging the feed). Fields: `used` =
  `SUM(credits_used) WHERE used_at >= start` **exactly as Go** (NULL ⇒ 0; legacy
  0-rows count 0; do **not** COALESCE to 1), `limit` **from the same
  active|trialing subscription Go uses (else free)** — `User#current_plan`
  ignores status today, so decide: fix `current_plan` to use the active scope
  (preferred, affects all callers; list them) or only UsageSummary; add a
  cancelled/past_due paid-row test; `limit` (nil/≤0/unknown plan ⇒
  "unlimited/unknown" state: no bar, no percent, no divide by zero),
  `remaining`, `percent` (Float 1 decimal, `<0.01` floor), `cycle_start`,
  `resets_at`, `threshold_state` (ok/warn ≥75/high ≥90/exhausted, single
  definition replacing the hardcoded `< 100`).
- Overview: new dashboard card partial (not `_stat_card`): exact used count with
  delimiter (compact only the limit; document K/M/B vs es/fr),
  `role=progressbar` + aria values, float width with `.` decimal, min visible
  width only when used>0, text `0.02% used`, sub-line e.g. "Cycle Sep 16 – Oct
  15 · resets Oct 16 (in 1 day)" for anchor-day 16 on Oct 15
  (day-before-boundary; add as a test vector; UTC dates, labelled UTC).
  Remaining card uses `threshold_state` (icon+text, not colour only). Replace
  redundant "Monthly Quota" card by last-request/cycle info. Banner subtitle
  shows the cycle. Keep **"Next billing"** (`current_period_end`) distinct from
  **"Usage resets"** (`next_start`): for promo subscriptions show "Plan ends
  <date>" instead of billing, and never present both as the same date. Admin
  `users/show.html.erb:112` relabels `current_period_start` ("Anchor") or hides
  it for promos. No-subscription/free/unlimited/new-user (no keys → link to API
  Keys first) states.
- Billing page: same `UsageSummary`; move its warning banners to
  `threshold_state`. `User#usage_this_month`/`requests_remaining` delegate to
  it; delete dead `User::PLAN_LIMITS` (`user.rb:19`, confirmed unreferenced).
- Admin: relabel/route admin "requests this month" to `UsageSummary` (or label
  "calendar month"); leave analytics queries alone.
- i18n: new keys en/es/fr; remove only `time_col`, `monthly_quota`,
  `resets_monthly` if unused elsewhere (**keep `requests_col`**, admin uses it);
  `time.formats.*` go in root `en.yml/es.yml/fr.yml`; i18n-tasks missing/unused
  as explicit CI step.
- Copy that promises things that don't exist (found in review):
  `home.*.yml:390-392` FAQ billing a2 ("we'll send you an alert at 75%/90%") and
  `dashboard.*.yml` `usage_alerts_hint` / `weekly_reports_hint` (no alert job,
  no weekly mailer exists). Reword or mark "coming soon" in en/es/fr until D5
  ships. PlanConfig labels ("500 requests/month", `plan_config.rb:12-55`) become
  "credits" wording or get a footnote (cosmetic). Check `home.en.yml:375` ("both
  key types count against…") against the new wording.
- i18n scope: delete/rename overview keys in `dashboard.{en,es,fr}.yml`; also
  rename `dashboard.usage.cards.total_requests`/`failed_requests` labels to
  "calls". **Do not run `i18n-tasks normalize` blindly**
  (`config/i18n-tasks.yml data.write` targets `config/locales/%{locale}.yml`,
  would rewrite the nested files); use `missing`/`unused` only.
- Public copy: update FAQ/glossary wording (en/es/fr) to "credits… resets on
  your billing day"; changelog entry (`home/changelog` + locales) worded as a
  correction.
- Cross-language check: a ledger-equality test (SQL/fixture of rows incl. NULL,
  0, 2× credits, boundary) asserting Rails `used` equals the same sum Go's
  `sumUsage` query computes (run the identical SQL in a Go test against the same
  fixture rows).
- Tests: BillingCycle vectors; UsageSummary (promo anchor 6 months ago, yearly,
  free, no subscription, cancelled, 2× credits, NULL credits, exact boundary,
  unlimited, limit 0); overview + billing controller tests; tests live in
  `test/lib/` (BillingCycle), `test/services/` (UsageSummary if placed there),
  `test/models/`, `test/controllers/dashboard/`; `test/helpers/` does not exist
  — put sparkline/percent helper tests under `test/controllers` or create
  `test/helpers/dashboard/` deliberately.
- **Acceptance:** seeded screenshot user (anchor 2026-04-16, July logs, now
  Oct 15) shows **0 used** with "last request 3 months ago"; 2,018 of 10M shows
  `0.02%`; Rails `used` equals Go `sumUsage` on the same ledger (the dashboard
  shows the **ledger**; Redis, which enforces, can legitimately be higher after
  dedup drops/panics/failed refunds — state this in help text/PR);
  revoked-oldest-key test passes; no `usage_this_month` caller left on the old
  query (checklist: `user.rb:87` and `requests_remaining` `user.rb:98` delegate;
  a multiline-aware guard test bans `used_at >=` outside
  `UsageSummary`/`BillingCycle`/`UsageLog`; callers to repoint:
  `billing_controller.rb:12-19`, `UsageLog.this_month` (usage_log.rb:10; delete
  if unused), `admin/users_controller.rb:95`,
  `admin/abuse_reports_controller.rb:48`). `total_requests` stays a row count
  and is labelled "calls", never "credits"; `BillingCycle.start` is **midnight
  UTC** of the anchor day (Go `cycleStart`), with a non-midnight-anchor vector.

### PR-B — stop recording route misses and stop charging errors (Go; fixes R4)

- Add a helper in `app/routes_v1.go` that builds every service router with
  explicit `NotFound` and `MethodNotAllowed` handlers that flag the request
  (context pointer) and write the normal 404/405. Also set them on `v1` itself
  (covers disabled services via `serviceEnabled`). Do **not** use
  `RoutePattern()`. Test walks all mounted prefixes plus an unknown top-level
  `/v1/x`.
- `UsageQuota`: if flagged ⇒ refund reserved credits, **no ledger row**. Scanner
  traffic still consumes a rate-limit token. To preserve abuse signal for admin
  analytics, increment a best-effort per-user bucket
  (`HINCRBY route_miss:<date> <user_id>`, with TTL; a Redis error never fails
  the request). Rails does not read this Redis for app data today, so **admin
  surfacing is a follow-on**; at minimum emit a structured log line.
- Panic safety (independent bug fix): recover inside `UsageQuota` around `next`
  (re-panic `http.ErrAbortHandler`); if `ww.Status()==0` (response not yet
  committed; `BytesWritten()` stays 0 after a bare `WriteHeader(200)`) write a
  JSON 500, log the row with status **500** (not 0, even when the committed
  client status was 200), report to Sentry. A recovered panic then continues
  through the normal post-handler flow (status 500 sets credits to 0 and the
  reconcile refunds the reservation), so there is exactly one refund path: no
  separate deferred refund, or the reservation is refunded twice. Tests: an
  uncommitted panic returns 500 and records 500; a panic after
  `WriteHeader(200)` keeps the 200 response, records 500; both leave the counter
  at its pre-request value. Panics in earlier middleware (auth, rate limiter)
  remain uncovered — stated, not fixed. Test: response 500 and row
  `status_code` 500.
- Compute `now` **once** and use it for both the cycle key and `used_at`
  (request start). Side effect: ledger order becomes start-time order (slow
  requests shift in Recent Activity); mention in PR.
- Order of operations after `next`: derive final status (0 ⇒ 200; recovered
  panic ⇒ 500), then set `creditsUsed = 0` if status ≥ 400, **before** the
  `creditsUsed != requestedCredits` reconcile and before `recordUsage`. Test:
  dictionary (requested 2) returning 422 ⇒ counter drops by 2, row has 0.
- **Refunds must survive client disconnects:** run every PR-B `adjust`
  (reconcile, refund, deferred panic refund) under
  `context.WithoutCancel(r.Context())` + a short timeout, like `recordUsage`
  already does (today `adjust` uses the request context, so a cancelled request
  leaves Redis over-counted while the ledger says 0). Test: handler cancels the
  request context ⇒ counter restored. Acceptance wording is "counter equals the
  pre-request value **when `adjust` succeeds**"; add a failure-injection test
  showing a failed refund is logged and the next cold bootstrap (from the
  ledger) self-heals.
- Persist the **normalised** status (0 ⇒ 200, recovered panic ⇒ 500) in
  `insertUsageRow`, not raw `ww.Status()`; PR-C chip treats legacy `status_code`
  0 as success.
- Treat `ww.Status()==0` as 200 (handler never called WriteHeader) unless a
  panic was recovered; test it.
- Empty batch: `headerCredits` must distinguish header **absent** (⇒ default)
  from `"0"` (⇒ 0 credits on 2xx), reconciled with `adjust(-1)`; test.
- Write-amplification: logged free errors still cost a Redis reserve+refund and
  a Postgres insert. Accepted: the per-key rate limiter runs earlier and bounds
  it; revisit with retention (D5) if table growth shows up. Pre-matching (below)
  removes the Redis cost for route misses.
- Prefer **pre-matching the route**
  (`router.Match(chi.NewRouteContext(), method, path)` on the **root** router)
  before reserving credits: no reserve/refund inflation (a scanner can't burn
  headroom or turn a typo into a 429 for an over-quota user) and no refund path.
  Wiring: `usageQuota` is built before the router exists (`app.go:52`) so the
  matcher is a **late-bound settable field** that fails open (treated as a
  match) when nil; match on the root router with the full path, using `RawPath`
  if set else `Path`, with a fresh `chi.NewRouteContext()` (chi routes on
  RawPath); add a percent-encoded-path test. If pre-match proves impractical,
  keep flag+refund and document the transient. Either way NotFound/405 handlers
  must be set at the **top** of `registerV1Routes` and on each sub-router
  _before_ `RegisterRoutes`/`Mount` (chi's Mount copies them only when the
  sub-router's own is nil); unit-test one nested `Route` inside a service.
- Scope note: only pure route misses are unrecorded; handler-level 404s on
  parameterised routes (e.g. `/v1/places/{code}`) are logged but free (policy
  above). Confirm `/v1/text/advice` is a true miss (no such route found in
  `apps/api/services/text`).
- PR-B does not clean the screenshot's existing 404 rows; those stay as stored
  and count as stored in `UsageSummary`.
- **Policy (decided by the owner): 4xx and 5xx responses are never charged.**
  Handler-level 4xx/5xx (422, legit 404 on parameterised routes, 500, recovered
  panics) are still **logged** (useful for debugging) with `credits_used = 0`;
  the reserved credit is refunded via the existing `adjust(-requested)` after
  `next`. The status override applies even if `X-Usage-Count` is set. 3xx/2xx
  charge as today (batch header cost only on 2xx). No env flag: the Rails reader
  sums the stored ledger, so Rails/Go deploy order does not matter. Empty-batch
  quirk (`X-Usage-Count: 0` billed 1) is fixed in the same change (2xx with 0
  items ⇒ 0 credits) or explicitly left with a test documenting it. Dedup
  `ON CONFLICT DO NOTHING` can still drop same-µs 0-credit rows; accepted, noted
  in tests.
- Public wording (en/es/fr FAQ/glossary + the API's served spec (`apps/api/app`
  `OpenAPISpec`) and `apps/mcp/openapi.json` descriptions, locate the
  generator/source, + changelog): "Client (4xx) and server (5xx) errors are not
  charged; unknown routes are not recorded." (No claim about 3xx.) PR-A copy
  must not mention errors (PR-A/PR-B use separate keys and separate changelog
  entries to avoid en/es/fr conflicts).
- Pricing/ToS check: confirm no page promises otherwise.
- Rewrite (not just extend) `usage_test.go` and `app_test.go` 404 cases; tests:
  422/500/404-handler ⇒ row with `credits_used=0` and Redis counter back to
  pre-request value; 2xx dictionary still 2, batch still N; route-miss no row +
  Redis unchanged; 405; panic refunds; refund when key expired mid-request (-1
  path); handler 404 still logged.
- **Acceptance:** `GET /v1/nonexistent` and `/v1/text/nonexistent` ⇒ no
  `usage_logs` row and Redis counter unchanged; `GET /v1/text/advice` valid ⇒
  one row; a handler 422/500 ⇒ one row with 0 credits and unchanged Redis
  counter; panic test leaves counter correct.

### PR-C — readable Recent API Activity (Rails view; fixes R3 readability)

- Replace table by a two-line feed row inside
  `turbo_frame_tag "overview-activity"`: status chip (text + icon + sr-only
  label), method chip, endpoint (ERB-escaped; no `html_safe`; `min-w-0`,
  wrap/truncate), latency, **credits as `1 credit` / `not charged`** (rename
  column semantic), key name (keys are only soft-revoked — `api_key_id` is NOT
  NULL + FK — so show "Revoked key" when `revoked_at` is set; no nil fallback
  needed; note `ApiKey has_many :usage_logs, dependent: :destroy` would delete
  ledger rows if a key were ever destroyed), `<time datetime>` relative +
  localized absolute title (UTC labelled). `flex-wrap` at 360px, no horizontal
  scroll.
- Filter pills (All / Successful / Errors) as real links with `aria-current`,
  `data-turbo-action="advance"`, allow-listed `params[:activity]`, scoped to
  `current_user`. Empty/stale state ("No requests this cycle · last request N
  ago"; no-keys state).
- Grouping (collapse consecutive identical endpoint+status, fetch ≈50 → cap 10
  groups) and the 7-day sparkline are **in scope only if time allows; otherwise
  D5**. Dark-mode contrast pass on icons/links/warnings (verify the "white
  panel" in a browser first).
- No flag mechanism exists in the app: **no ship-dark flag**; rely on staged PRs
  (A before C) and revert.
- Legacy rows: pre-PR-B error rows were stored with 1 credit (Worker-era ones
  with 0). `UsageSummary` mirrors Go's stored-credit sum (no error filter), so
  the feed uses neutral wording that is true regardless of PR-B ordering: error
  row with `credits_used > 0` ⇒ "1 credit (charged)" (count shown as stored),
  `credits_used == 0` ⇒ "not charged", **NULL ⇒ "not recorded"** (never "1
  credit"; the counter treats NULL as 0). The Usage page/CSV keep `|| 1` until
  D4 (known mismatch, noted). Seeded acceptance data must include both.
- **Depends on PR-A** (cycle/`last_request_at`); `not charged` label keys off
  `credits_used == 0` as stored (also true for legacy rows). PR-0 and PR-B are
  independent of everything.
- Tests: feed rendering (zero/stale/errors/revoked key), filter allow-list,
  cross-user isolation, escaped endpoint, a11y labels.
- **Acceptance:** no horizontal scroll at 360px; axe clean on `/dashboard`;
  seeded rows render correctly in all three variants — Go-era 404 with 1 credit
  ("1 credit (charged)"), Worker-era 404 with 0 ("not charged"), NULL ("not
  recorded") — each with an error chip; user B never sees user A rows via any
  param.

### Deferred (separate plans)

- **D1 Go anchor/cycle correctness:** user-level anchor (`users.created_at`
  preferred over MIN(key.created_at) — stable under key deletion/revocation),
  `cycle_end` fix for 29–31, rolling-deploy cache-key semantics (cache JSON
  field rename/versioning), subscription `after_commit` →
  `GoAuthCache.invalidate`, deterministic subscription pick / unique index,
  fail-closed unknown plan in `plancache.go`. Ships with a Ruby `BillingCycle`
  update + same vectors **in the same release**; quota-impacting; sign-off.
- (D2 removed: error charging is now part of PR-B.)
- **D3 Webhook anchor hygiene** (don't set `Time.current` on payment; keep
  monthly renewal day for yearly; promo expiry resets anchor). Changes real
  quotas — explicit sign-off; no backfill unless approved (dry-run diff,
  `CONFIRM=1`, human-run).
- **D4 Usage page rework:** default "this billing cycle", i18n labels, span cap,
  UTC day grouping, Stimulus chart controller, theme-aware colours/table
  fallback, CSV streaming + row cap + `Credits`/`Billed` columns + rate-limit.
- **D5 Sparkline/grouping/alerts job**, summaries schema (error_count, avg
  latency) + backfill, then retention. No deletion of legacy rows (kept for
  disputes/analytics; filter in reads). PII note: raw request paths are shown
  back to the user; no privacy-text change since no new data is collected.

## Open decisions (defaults)

1. ~~Charge 4xx/5xx?~~ **Decided: never charged** (now part of PR-B).
2. Keep route-miss abuse signal as a per-user Redis bucket + log line? Default
   yes.
3. Go anchor change accepted (D1)? Default yes, later, separately.

## Risks

- Rolling back PR-A's `User#usage_this_month` delegation changes
  `requests_remaining` for all callers (user.rb, billing, overview, admin) —
  list them in the PR description.
- Displayed usage drops for affected users (a correction; changelog).
  Enforcement is unchanged by PR-0/A/C.
- PR-B changes billing (errors/route misses become free) — customer-visible, in
  copy + changelog; rollback = revert (no data migration). Monitor quota-reject
  rate after deploy. Existing historical error rows keep their stored credits.
- Rails/Go may disagree briefly for multi-key free users until D1 (documented).

## Final notes

### PR-0 (#1012, merged)

- Shipped: the endpoint usage list on the Usage page is rendered server-side
  (ERB-escaped, ordered by count, capped at 100, nil/blank endpoints skipped)
  and shares one scope with the `by_endpoint` JSON action; the CSV export
  prefixes cells that start with `= + - @ \t \r \n`, full-width `＝＋－＠`, or
  those after leading whitespace, with `'`.
- Deviations: `find_each` was tried for the export and reverted because it
  ignores the newest-first order (streaming and row caps stay in D4). The
  `by_endpoint` action and route stay (now only used by tests and any external
  caller). Review feedback to use a tab prefix and cursor streaming was not
  taken (plan specifies `'`; streaming is D4). Seven review rounds.
- The panic-recovery guidance for PR-B was tightened after review: use
  `ww.Status()==0` (not `BytesWritten()==0`) to decide if a JSON 500 can still
  be written, and keep a single refund path (the normal reconcile) so a
  recovered panic cannot refund the reservation twice.

### PR-A (Rails numbers)

- Shipped: `BillingCycle` and `UsageSummary`, `User#current_plan` now follows
  the active/trialing subscription (callers: Overview, `rack_attack.rb`
  free-plan throttles), `User#usage_this_month`/`requests_remaining` delegate to
  `UsageSummary`, dead `User::PLAN_LIMITS` and `UsageLog.this_month` deleted,
  Overview cards and Billing page use the summary, admin "this month" is now
  "this cycle" (credits), admin "Started" is "Cycle anchor" and hidden for
  promotions, copy that promised alerts and weekly reports is marked "coming
  soon", FAQ/glossary say credits and "billing day", changelog entry added.
  Shared Go/Rails parity vectors and a ledger fixture live in
  `apps/api/platform/middleware/testdata/` (the CI dashboard path filter now
  includes that folder).
- Percent is shown with two decimals (`0.02%`), because the plan's "1 decimal"
  could not render the required `0.02%`; below `0.01` shows `<0.01%`.
- Deviations: no explicit blocking `i18n-tasks unused` CI step (main already
  reports about 80 unused keys); the unused check was done as a before/after
  diff instead, and `dashboard.overview.usage.state.*` is ignored (dynamic key).
  `PlanConfig` feature labels ("500 requests/month") were left as is (cosmetic,
  plan allowed a footnote instead). The Usage page and CSV still say requests
  (D4).
- Out-of-scope bugs found (not fixed): `admin/users/show.html.erb:256` calls
  `UsageLog#http_method` (the column is `request_method`), so the admin user
  page raises for any user with recent usage logs;
  `admin/abuse_reports/show.html.erb:107` calls the nonexistent `ApiKey#prefix`,
  so the abuse report page always raises; the Billing page raises in
  `PlanConfig.for` for a subscription whose `plan_name` is not a known plan
  (reachable only by bypassing validations); `UsageLog.credits_used` is
  nullable, so Usage page/CSV `|| 1` and the ledger sum disagree (D4).

PR-A merged as #1013 after four review rounds. Late additions: the Billing page
now describes the effective (active/trialing) subscription like the Overview and
keeps a customer-portal button for a paid row that is not active (for example
`past_due`), `UsageSummary#percent` never shows 100% before the limit is
reached, a `subscription_pick_vectors.json` fixture proves the Go and Rails
subscription pick (status filter, active before trialing, newest by `updated_at`
then id, NULL period start) from the same rows, and the parity tests fail
instead of skipping when `CI` is set. A CodeQL false positive on a test helper
parameter named `api_key` was avoided by renaming it.

### PR-B (Go: route misses and free errors)

- Shipped: `UsageQuota` asks a late-bound matcher (`SetRouteMatcher`, wired in
  `app.go` to `router.Match` on the root router, using `RawPath` when set)
  before doing anything; a request with no route (unknown path, wrong method)
  skips reservation and ledger, still reaches the router for its normal answer,
  and bumps a best-effort `route_miss:<date>` Redis hash per user (8 day TTL;
  nothing reads it yet, admin surfacing is a follow-on). The matcher fails open
  (nil or panicking). Every `/v1` router (and each service router, so nested
  `Route`s inherit) also carries `NotFoundHandler`/`MethodNotAllowedHandler`,
  which flag the request so the same refund-and-no-row outcome holds if
  pre-matching is unavailable.
- `now` is taken once (cycle key and `used_at` are the request start). After the
  handler: status 0 becomes 200, a recovered panic becomes 500 (JSON 500 when
  nothing was committed, `ww.Status()==0`; Sentry; `http.ErrAbortHandler` is
  accounted then re-panicked), status >= 400 sets credits to 0 before the
  reconcile, and the single reconcile path refunds. `adjust` runs under
  `context.WithoutCancel` so a client disconnect cannot leave Redis
  over-counted. `X-Usage-Count: 0` means 0 credits; an absent header still
  means 1. Every batch endpoint validates `min=1`, so an empty batch is not
  reachable over HTTP today: this is defensive, and the changelog does not claim
  it.
- Public copy: new FAQ entry and changelog block with their own keys (en/es/fr),
  and the `info.description` of the served spec (`apps/api/app/openapi.json`,
  the only source; there is no generator in the repo) and its byte-identical
  copy `apps/mcp/openapi.json`.
- Deviations: the 404/405 handlers write chi's default responses, and a 405
  keeps its `Allow` header by probing the root matcher for each method (chi
  hands the allowed methods only to its own default handler); nothing reads the
  route-miss bucket yet (Rails surfacing is a follow-on); the Go middleware
  tests that need a real Redis (`CONFIG SET`, `DEBUG SLEEP`) were run locally
  against miniredis and two of them cannot pass there (rate limiter timeout and
  OOM), CI runs them against real Redis.
- Pricing/ToS check: no page promises that errors are billed.

PR-B merged as #1022 after five review rounds. Review also moved the 405 `Allow`
header back in, split the post-handler accounting into `settle`, kept
`response_time_ms` as the handler's own time while `used_at` is the request
start, pinned the OpenAPI wording and the byte-identical MCP copy with a test,
and removed the changelog claim about empty batches: every batch endpoint
validates `min=1`, so that path is unreachable over HTTP.

### PR-C (activity feed)

- Shipped: the Recent API Activity table is now a feed inside
  `turbo_frame_tag "overview-activity"`: a status chip (code, icon and a
  screen-reader label for success, redirect, client error, server error; legacy
  status 0 or NULL reads as success), method chip, escaped endpoint
  (`break-all`, wraps), latency, credits label (`N credit(s)`,
  `N credit(s) (charged)` on an error row that still carries credits,
  `not charged`, `not recorded` for NULL), key name or `Revoked key`, and a
  `<time>` with relative text and a localized absolute UTC title (formats in the
  root `en/es/fr.yml`). All / Successful / Errors are real links
  (`aria-current`, `data-turbo-action="advance"`) with an allow-listed
  `params[:activity]`, always scoped to `current_user`. States: empty, no keys
  (create-key CTA), no matches for a filter, and a stale line ("No requests this
  cycle · last request 3 months ago"). Ten rows, newest first.
- Verified in a real browser (Chrome driven by playwright-core against the Rails
  test environment on a throwaway database, seeded with every legacy variant, a
  revoked key, a hostile endpoint and a very long endpoint): no horizontal
  scroll at 360px or 1280px, axe (WCAG 2 A/AA, 2.1 AA) clean in light and dark,
  the Turbo filter click updates only the frame and advances the URL, Back
  restores it. Axe found three real problems that were fixed: the scrollable
  list was not keyboard focusable (`tabindex="0"` and a focus ring), and the two
  sidebar "Manage" links were `text-blue-600` on the dark card (2.79:1; now
  `dark:text-blue-400`). Stale prebuilt Tailwind output makes new utility
  classes missing locally (run `bin/rails tailwindcss:build`); CI and production
  build it.
- The "white panel in dark mode" claim did not reproduce: a scan of every large
  element in dark mode found no white background on the Overview, so nothing was
  changed for it.
- Review found, and a browser check confirmed, that the dashboard layout's
  `turbo-visit-control: reload` meta turned the advancing filter links into a
  full page reload (a `window` marker did not survive the click). The Overview
  now opts out of that meta (`content_for :allow_turbo_visit`; other dashboard
  pages keep it), so the click swaps the page without a reload, advances the URL
  and Back restores it, and the active chip is autofocused only on a Turbo frame
  request so keyboard focus is not lost. The list is a labelled, focusable
  `role="region"` scroller. Statuses above 599 count as server errors.
- Follow-up (not done): the Errors/Successful filters query `usage_logs` by
  `user_id`, `status_code` and `used_at DESC`; only `(user_id, used_at)` is
  indexed, so a user with very few matching rows makes a filter click walk their
  history. A partial index on `(user_id, used_at) WHERE status_code >= 400` (or
  summaries, D5) would fix it; it needs a schema migration, which this plan does
  not authorise, so it is left for the retention/summaries work.
- Not done (left for D5 as the plan allows): grouping consecutive identical rows
  and the 7-day sparkline.
- Removed keys: `overview.time_col` and `overview.status_col` (the admin views
  still use `endpoint_col`, `response_col` and `requests_col`).

PR-C merged as #1023 after five review rounds (the first attempt at round five
failed on a rate limit and was re-run, not counted as clean). All four PRs are
merged: #1012 (PR-0, seven rounds), #1013 (PR-A, four), #1022 (PR-B, five) and
#1023 (PR-C, five).

Process notes: an editor-side sync committed and pushed unrelated work to
`origin/fix/usage-pr0-security` early on, so PR-0 shipped from
`fix/usage-pr0-security-escaping` and the stray remote branch was left alone for
its owner to delete.

Follow-ups not done: the filter index for `usage_logs` (above), grouping and the
sparkline (D5), surfacing the Redis `route_miss:<date>` buckets to admin abuse
tooling, D1, D3 and D4. Out-of-scope bugs found and left alone: the admin user
page (`UsageLog#http_method`), the admin abuse report page (`ApiKey#prefix`),
the Billing page for a subscription with an unknown plan name, and the Usage
page and CSV still presenting `credits_used || 1` (D4).
