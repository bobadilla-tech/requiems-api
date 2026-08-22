# Go Auth Foundation — Phases 8–9: Workers Retirement Completion

Continuation of:

- `docs/audits/2026-08-21-architecture-audit.md`
- `docs/plans/2026-08-21-go-auth-foundation-phase-0-1.md`
- `docs/plans/2026-08-21-go-auth-foundation-phase-2.md`
- `docs/plans/2026-08-21-go-auth-foundation-phase-3-4.md`
- `docs/plans/2026-08-22-go-auth-foundation-phase-5.md`
- `docs/plans/2026-08-22-go-auth-foundation-standing-issues-hardening.md`
- `docs/plans/2026-08-22-go-auth-foundation-phase-6-usage-multiplier-and-loose-ends.md`
- `docs/plans/2026-08-22-go-auth-foundation-phase-7-cloudflare-workers-retirement.md`

This is the next focused work package after Phase 7a's direct-ingress work. It
finishes the retirement of the Cloudflare Worker application layer and then
removes the repository's remaining Worker-era operational and documentation
surface.

## Current state and decision

The Go API is already the enforcing auth/rate-limit/usage path. Phase 7a's
recorded live work put Cloudflare proxying, Caddy, AOP, Kamal, and Go on the
direct public path; it also removed the Caddy `X-Backend-Secret` gate and
repointed Rails' internal API URL to the private `requiems-api:8080` network
alias.

Phase 7a is not considered closed until the following are observed and
recorded:

1. The owner confirms that `api.requiems.xyz` has no active Worker route or
   custom-domain attachment.
2. A valid production API key succeeds through
   `Cloudflare → Caddy/AOP → Kamal → Go`.
3. The same public path returns 401 for an invalid key and 429 after the
   configured per-minute limit is exceeded.
4. A Rails Playground request succeeds over the private internal URL.
5. The direct-origin request is rejected by AOP (TLS client-certificate
   failure), and the short bake period has no observed regression.

These checks are a prerequisite for deletion, not a reason to build a canary,
shadow path, notice period, or dual billing system. There are no users or
production traffic, so this migration may use a direct 100% cutoff. The Worker
route can be disabled or already absent while these checks are completed, but
the Worker code and Cloudflare stores must remain available until the checks
are recorded.

## Data-retention boundary

No D1-to-Postgres reconciliation, backfill, or customer-notice process is
needed. The project owner has authorized discarding pre-launch operational
rows. Before any destructive database command, verify the exact host, database
name, schema, migration version, and environment; print and review the target
table counts.

Rows may be truncated or discarded only from these operational ledgers when
they are present and are confirmed to be test/pre-launch traffic:

- Cloudflare D1 `credit_usage` (deleted with D1).
- PostgreSQL `usage_logs`.
- PostgreSQL `daily_usage_summaries`.
- PostgreSQL `credit_adjustments`.
- PostgreSQL `audit_logs`.

Do not delete or truncate `users`, `api_keys`, `subscriptions`, `plans`, or
any Go reference/product-data table, including `advice`, `quotes`, `words`,
`bin_data`, `inflation_data`, `iban_countries`, `commodity_price_history`,
`exercises`, `swift_codes`, `counters`, or other tables containing API data.
Do not use `CASCADE`, wildcards, schema-wide truncation, or a migration that
drops a table merely because its rows are disposable. Use an explicit five-table
allowlist, capture before/after counts for those tables, and hard-fail if any
protected table changes. The implementation notes must state the exact SQL
commands (or that no cleanup was run), target database, environment, and row
counts.

## Phase 8 — Irreversible Worker and edge-state retirement

This phase is one focused deletion session, but it has two explicit live-state
gates. Repository edits may be prepared in a branch before the gates, but the
Cloudflare deletions must not be run speculatively.

### 8.1 Re-verify Phase 7a and establish a rollback record

Before deleting anything:

1. Read the Phase 7a final notes and compare them to current live state.
2. Confirm the public hostname, DNS proxy status, Caddy AOP configuration,
   Kamal routing, Rails private URL, and Go health/auth behavior.
3. Run a successful Kamal-managed deploy/restart using the committed routing
   configuration, then re-verify AOP, public routing, and Rails' private API
   URL. Do not treat the one-off live container state from Phase 7a as the
   final deployment proof.
4. Capture the Worker route/custom-domain status and the valid-key/401/429/
   Playground smoke results in the implementation notes. If no suitable live
   key exists, create a disposable pre-launch key through the normal Rails
   generator, use it only for the read-only smoke test, then revoke it; never
   use a load-test credential for production smoke traffic.
5. Preserve the last known-good pre-deletion commit and the direct-cutover
   commit (`6ca862f4`) as the code rollback record. No application rollback
   depends on recreating a deleted KV/D1 resource.

Immediately before the first Worker/KV/D1 deletion, record a separate
point-of-no-return approval: all Phase 7a evidence is complete, the owner
understands that re-enabling the Worker route and restoring KV/D1 will no
longer be rollback options, and the remaining rollback is limited to reverting
Go/Caddy/Kamal code and redeploying it. This approval is distinct from approval
to prepare repository diffs.

If any public-flow check fails, stop Phase 8. Do not delete the Worker or the
Cloudflare stores while the direct path is uncertain.

### 8.2 Stop all D1 consumers before deleting D1

Remove `sync_d1_usage` from both:

- `apps/dashboard/config/recurring.yml`
- `apps/dashboard/config/sidekiq_schedule.yml`

Deploy/restart both scheduler systems so both schedule sources are reloaded.
Quiet or drain the relevant workers, then inspect Solid Queue and Sidekiq
scheduled, queued, retry, and dead sets for `SyncD1UsageJob`; also confirm
there is no currently running or leased instance. Observe at least one former
schedule interval with no D1 job or D1 API log entry. Only after active,
queued, retry, and scheduled work are clear may D1 be deleted. Do not change
`AggregateDailyUsageJob`; it reads Postgres `usage_logs` and remains valid.

### 8.3 Delete Cloudflare edge resources, with explicit confirmation

With the owner present and confirming each destructive action immediately
before it runs:

1. Identify the exact production Worker names, route/custom-domain bindings,
   KV namespace ID, and D1 database name/ID from the live account and current
   Wrangler configuration. Do not rely on a guessed resource name.
2. Confirm the Worker route is disabled/absent and that no Rails or DNS path
   still depends on `api-management.requiems.xyz`.
3. Delete the production `auth-gateway` and `api-management` Workers and
   remove their live custom-domain/route bindings if any remain. Verify their
   deployed secrets are no longer active; in particular, check
   `BACKEND_SECRET` and `API_MANAGEMENT_API_KEY` rather than assuming Worker
   deletion cleaned every secret.
4. Delete the Cloudflare KV namespace and D1 database used by the Workers.
   The data is disposable under this plan, but the exact resource identity and
   owner confirmation remain mandatory because the deletion is irreversible.
5. Verify that requests to the public API still terminate at the direct
   Cloudflare/Caddy/Go path and that no Worker invocation or D1/KV access is
   present in the live path. Capture the before/after Cloudflare Worker list,
   routes, triggers, custom domains, KV/D1 identifiers and deletion output;
   pair a cache-busting public request's Cloudflare Ray ID with matching Caddy
   and Go access-log entries. Check the retired
   `api-management.requiems.xyz` hostname separately.

Do not grey-cloud the DNS record or remove Cloudflare proxying. Cloudflare
remains the DNS/WAF/DDoS/TLS and AOP-facing layer.

### 8.4 Remove the Worker code and Rails D1 integration

Delete the now-dead application code:

- `apps/workers/auth-gateway/`
- `apps/workers/api-management/`
- `apps/workers/shared/`
- `apps/workers/readme.md`
- `apps/dashboard/app/services/d1_sync_service.rb`
- `apps/dashboard/app/jobs/sync_d1_usage_job.rb`
- D1-specific locale entries and tests that have no remaining caller.

Before deleting any Rails file, run a repo-wide reference search and remove
callers/configuration first. Do not remove unrelated Postgres usage jobs.

### 8.5 Remove build, deployment, and local-stack wiring

Remove Worker-only CI/CD and development wiring, including:

- Worker path filters, `worker-test`, `api-management-test`, and their
  reusable `.github/workflows/_worker-ci.yml` workflow.
- Worker entries from `.github/dependabot.yml`.
- Worker install/setup steps and Worker-only test secrets from
  `.github/workflows/copilot-setup-steps.yml` and `.github/workflows/ci.yml`.
- Worker deploy secrets and required-secret lists from
  `.github/workflows/cd.yml`, `infra/kamal/deploy.api.yml`,
  `infra/kamal/deploy.dashboard.yml`, and `infra/kamal/secrets`, after
  confirming the variables are not used by another feature.
- `infra/docker/auth-gateway.dev.Dockerfile` and
  `infra/docker/api-management.dev.Dockerfile`.
- The `auth-gateway` and `api-management` services, shared Worker mounts,
  Wrangler state, module volumes, and stale `depends_on` edges from
  `infra/docker/docker-compose.dev.yml`.
- The `dashboard` and `sidekiq` dependency edges so they depend on the Go
  `api` service (the private `requiems-api:8080` destination), not the deleted
  `api-management` service.
- The MCP dev service's `REQUIEMS_BASE_URL` and dependency edge so MCP uses
  the direct Go API service (`api:8080`) or the explicitly chosen public API
  URL, not the deleted `auth-gateway` service.

Do not remove a generic `BACKEND_SECRET` or tenant/private-deployment secret
solely because its name resembles the retired Worker secret. First classify
each remaining reference. Remove the old Go/Caddy/Rails proxy secret only when
the reference search and tests show it no longer serves a supported feature;
preserve any genuinely separate private-deployment credential contract.

### Phase 8 exit criteria

- Phase 7a's public-flow, AOP, Playground, and bake checks are recorded.
- Both D1 schedules are disabled and no D1 sync job remains queued.
- The exact production Workers, KV namespace, D1 database, routes, and Worker
  secrets are deleted/verified absent with owner confirmation.
- `api.requiems.xyz`, `api-management.requiems.xyz`, and every discovered
  Worker custom domain/route have an intentional post-deletion DNS/HTTP state;
  only `api.requiems.xyz` remains an active public API path.
- No `apps/workers` code or Worker-only Docker/CI/dependabot/Kamal wiring
  remains.
- Rails boots without D1 sync classes, API Management settings, or stale
  scheduler entries.
- Local Compose config validates and starts the API, dashboard, MCP, Postgres,
  Redis, Caddy-related development dependencies, and Sidekiq without a Worker.
- No protected product-data table rows were changed; any operational-ledger
  cleanup is explicitly recorded.

## Phase 9 — Direct-Go repository consolidation and audit closure

This phase is the reviewable cleanup pass after Phase 8's deletion. It should
be a separate commit/PR from the irreversible Cloudflare deletion where
practical, so a documentation or test mistake cannot obscure the production
cutover.

### 9.1 Rewrite canonical architecture and operations docs

Update the current-state documentation to describe:

`Client → Cloudflare proxy/WAF/DDoS/TLS/AOP → Caddy → Kamal → Go → Redis/Postgres`

The Go service owns API-key authentication, Redis rate limiting, quota checks,
and Postgres usage rows. Rails owns users, API keys, subscriptions, plans, and
the dashboard. Cloudflare does not run application Worker logic or hold the
usage ledger.

Delete the obsolete docs:

- `docs/core/auth-gateway.md`
- `docs/core/api-management.md`

Rewrite the affected sections and navigation in:

- `docs/core/architecture.md`
- `docs/core/infrastructure.md`
- `docs/core/deployment.md`
- `docs/core/rails-app.md`
- `docs/core/background-jobs.md`
- `docs/core/code-quality.md`
- `docs/core/readme.md`
- `docs/core/local-payment-testing.md`
- `docs/core/lemonsqueezy-webhook-setup.md`
- `docs/core/maintenance-tasks.md`
- `docs/core/getting-started.md`
- `docs/core/adding-go-endpoints.md`
- `apps/dashboard/docs/app-config.md`
- `apps/readme.md`
- `.github/copilot-instructions.md`
- `scripts/generate-docs.mjs` and any generated API documentation it updates
- `tests/integration/README.md`
- `tests/load/README.md`
- `tests/load/.env.example`
- `tests/load/run.sh`
- `infra/readme.md`
- `infra/docker/readme.md`
- `infra/caddy/readme.md`
- root `agents.md`

Remove stale commands, ports, URLs, secret names, D1/KV schemas, Worker
deployment instructions, and references to the deleted Rails sync path. Keep
historical audit and plan documents unchanged; they are migration records, not
current architecture documentation. Review product/blog copy separately and
only change it if it falsely describes the current runtime rather than an
historical technology choice.

### 9.2 Update integration, load, and local test contracts

Preserve black-box behavior while removing Worker-specific assumptions:

- `tests/integration/src/suites/gateway.test.ts`: retain public HTTP checks,
  rename the suite around the public API/direct ingress, and verify
  `API_BASE_URL` is still the Cloudflare-fronted public hostname.
- `tests/load/config.ts` and `tests/load/scenarios/rate-limit.ts`: use the
  current Go-generated key contract and current rate-limit source; remove
  `localhost:4455`, `rq_*` literals, Worker seed instructions, and Worker-only
  comments. Keep limits sourced from the same plan configuration used by Go.
- `tests/integration/README.md`, `tests/integration/.env.example`,
  `tests/load/README.md`, `tests/load/.env.example`, and `tests/load/run.sh`:
  update setup, ports, credentials, and wrappers to use the direct API. Do not
  leave a Worker-only README or shell script as the way to run a passing test.
- `apps/dashboard/db/seeds.rb` and `infra/docker/.env.example`: define one
  explicit `LOCAL_DEV_API_KEY` (`requiem_<24 alphanumeric characters>`) for
  local Rails seeding, and make MCP/load tests consume that same key through
  environment variables. Remove `rq_*` defaults and Worker seed scripts.
- Rails `ApiProxyService` tests: assert the `requiems-api-key` path and remove
  the obsolete `X-Backend-Secret` assertion only after the application code
  removes that header.
- Replace/delete D1 sync tests and update background-job tests to cover the
  remaining Postgres jobs.
- Ensure `go test ./...`, Rails tests/security checks, MCP tests, and the
  remaining integration/load checks use the direct path.
- Update `.github/workflows/codeql.yml` to retain JavaScript/TypeScript
  analysis for `apps/mcp` while removing the deleted Worker target, and remove
  obsolete Worker flags from `codecov.yml` and Worker-only ignore entries from
  `.gitignore`.

### 9.3 Remove retired configuration and verify secret boundaries

Run a targeted reference audit for:

- `API_MANAGEMENT_URL`
- `API_MANAGEMENT_API_KEY`
- `BACKEND_SECRET`
- `CLOUDFLARE_ACCOUNT_ID`
- `CLOUDFLARE_KV_NAMESPACE_ID`
- `CLOUDFLARE_API_TOKEN`
- `CLOUDFLARE_WORKER_URL`
- `auth-gateway`
- `api-management`
- `SyncD1UsageJob`
- `credit_usage`

Also search tracked and hidden files, excluding only `docs/plans/` and
`docs/audits/`, for these concrete forms:

```text
auth-gateway  api-management  apps/workers  worker-test  _worker-ci
auth_gateway  api_management  api-management.requiems.xyz
wrangler  wrangler_state  localhost:4455  localhost:5544
Cloudflare Worker  CLOUDFLARE_  d1_sync_usage  credit_usage
X-Backend-Secret  API_MANAGEMENT_URL  API_MANAGEMENT_API_KEY
```

Produce a reviewed residual-reference manifest with file, line, classification,
reason for retention, and confirmation that it is not a live runtime or
configuration dependency. A remaining match cannot be described merely as a
"false positive"; it must be removed or explicitly allowlisted as historical
or product content. The allowlist should identify any retained blog,
translation, or historical explanatory copy separately from operational docs.

Delete variables and `AppConfig` readers that have no remaining consumer.
Create a usage matrix for `BACKEND_SECRET`: remove it from the normal
Rails-to-Go Playground request once Go API-key auth is the only supported
credential, while preserving the separate private-deployment tenant-secret
contract and its tests/documentation if it is still real. Remove stale tests,
locales, rake-task output, and app-config documentation along with each deleted
setting.

### 9.4 Full audit sweep and validation

Run searches outside `docs/plans/` and `docs/audits/` for the retired
components. Every remaining match must be one of:

- an intentional historical/product-content reference documented as such;
- a neutral security/infrastructure term that is not referring to the deleted
  Worker/D1/KV runtime; or
- a false positive removed by the cleanup.

Run validation against an isolated test target, never the development or
production database. Set `TEST_DATABASE_URL` to the dedicated disposable
database, use `RAILS_ENV=test` for Rails, and use an isolated Redis database or
disposable Redis instance. Add a guard that aborts if the host/database does
not match the expected test target. Capture before/after counts (and, where
practical, data-only checksums) for all protected product-data tables around
the suite.

Validate at minimum with executable commands appropriate to the containers:

```text
docker compose -f infra/docker/docker-compose.dev.yml config --quiet
docker compose -f infra/docker/docker-compose.dev.yml up -d api db redis dashboard sidekiq mcp
curl --fail http://localhost:8080/healthz
go test ./...                                  # in apps/api/container, TEST_DATABASE_URL set
RAILS_ENV=test bin/rails db:test:prepare       # in apps/dashboard/container
RAILS_ENV=test bin/rails test                  # in apps/dashboard/container
bundle exec brakeman --no-pager
bundle exec bundler-audit
bin/importmap audit
cd apps/mcp && bun test
cd apps/mcp && bunx tsc --noEmit
cd tests/integration && pnpm run typecheck && pnpm test
yaml/workflow/Kamal configuration validation for changed CI/CD files
```

For production, use a dedicated pre-launch key and a safe read-only endpoint.
Record a smoke-test table containing the key/plan (never the raw secret), exact
endpoint, expected status/body/Go-owned headers, request timestamp, Cloudflare
Ray ID, matching Caddy/Go log lines, expected usage-log impact, and cleanup
action. The rate-limit check must use a disposable low-quota-safe key or an
explicitly reset pre-launch ledger; never use integration/load credentials.
Repeat valid-key, invalid-key, 429, Playground, and cache-busting checks after
deletion. Verify direct-origin AOP with the actual TLS `certificate_required`
or equivalent client-certificate failure, not just a generic connection
failure. Check every prior scheduler (Solid Queue and Sidekiq) for no D1 job or
D1 log activity.

The post-deletion evidence matrix must include: Worker list/routes/triggers and
custom domains before and after; exact KV/D1 identifiers and deletion output;
no Worker invocation evidence; the retired API Management hostname's DNS and
HTTP behavior; Cloudflare Ray ID plus origin logs for a cache-busting request;
and timestamps for the valid-key, invalid-key, 429, Playground, and AOP tests.

### Phase 9 exit criteria

- Canonical docs and repo instructions describe direct Go auth and no longer
  present Workers/KV/D1 as live components.
- Public integration/load tests target observable API behavior without
  Worker-only setup or literals.
- CI, CD, Compose, Kamal, Copilot setup, Dependabot, Rails config, and test
  commands contain no dead Worker integration.
- The full targeted test/security suite passes.
- The final reference sweep is reviewed, with intentional historical/content
  matches listed rather than hidden.
- Implementation notes state the exact production resources deleted, the exact
  database rows (if any) removed, the protected data tables left untouched,
  validation results, and any follow-up work deferred until real traffic.

## Explicitly out of scope

- Cloudflare DNS/WAF/DDoS/TLS retirement; Cloudflare remains in front.
- A percentage canary, shadow comparison, notice period, or billing-cycle
  reconciliation; there are no users or traffic.
- New usage-multiplier policy, Redis performance tuning, or a second Redis
  instance; these require observed traffic.
- Dropping identity/configuration tables or Go product/reference data.
- Rewriting historical plans/audits merely to make grep return zero matches.
- Deleting or rotating unrelated private-deployment tenant secrets without a
  separate confirmed owner and contract.

## Review findings

Three independent reviewers audited this plan after it was drafted. Findings
are classified below and incorporated into the plan unless explicitly marked
otherwise.

### High

- **Test isolation — fixed in 9.4.** Require dedicated `TEST_DATABASE_URL`,
  `RAILS_ENV=test`, isolated Redis, target guards, and protected-table
  before/after checks. This prevents validation/load traffic from mutating
  development, production, or product-data rows.
- **Point of no return and rollback semantics — fixed in 8.1/8.3.** Require
  owner sign-off immediately before irreversible deletion and state that the
  post-deletion rollback is limited to Go/Caddy/Kamal redeployment.
- **Production completion evidence — fixed in 8.3/9.4.** Require before/after
  Cloudflare resource evidence, Ray IDs, matching origin logs, retired-hostname
  checks, and timestamped post-deletion smoke results.
- **Product-data protection — fixed in the data-retention boundary and 9.4.**
  Require exact database identity, an explicit five-table allowlist, no
  `CASCADE`/wildcards, protected-table counts/checksums, and exact cleanup
  notes.

### Medium

- **Running/leased D1 jobs — fixed in 8.2.** Cover both Solid Queue and
  Sidekiq, drain/restart workers, inspect scheduled/queued/retry/dead sets, and
  observe a former schedule interval before deleting D1.
- **One-off Phase 7a deployment state — fixed in 8.1.** Require a
  Kamal-managed deploy/restart followed by AOP, routing, and Rails private-URL
  verification.
- **Repository reference completeness — fixed in 8.5/9.1/9.3.** Add MCP,
  CodeQL, Copilot, generated docs, integration/load setup files, Compose seed
  wiring, ports, Wrangler forms, and a residual-reference manifest.
- **Local key contract — fixed in 9.2.** Make `LOCAL_DEV_API_KEY` the explicit
  source for Rails seeds and share it with MCP/load tests; remove old `rq_*`
  Worker defaults.
- **Secret-boundary ambiguity — fixed in 8.5/9.3.** Require a usage matrix
  before removing `BACKEND_SECRET`, preserving any genuinely separate
  private-deployment tenant-secret contract while removing the obsolete normal
  Rails-to-Go header.
- **Production smoke and hostname checks — fixed in 8.1/8.3/9.4.** Specify a
  dedicated key, safe endpoint, expected outputs, usage impact, Ray/origin-log
  evidence, AOP TLS failure, and checks for both public hostnames.
- **Validation commands — fixed in 9.4.** Replace vague JS/Compose language
  with direct stack, health, Go, Rails, MCP, integration, security, and config
  validation commands.

### Low

- **Post-deletion evidence source — fixed in 8.3/9.4.** Evidence now names
  Cloudflare outputs, Ray IDs, Caddy/Go logs, and timestamps.
- **Historical-reference handling — fixed in 9.1/9.3.** Historical plans and
  audits remain excluded, while every other retained match requires an
  explicit manifest entry.
- **Coverage/ignore cleanup — fixed in 9.2.** Codecov and `.gitignore` Worker
  remnants are included with CodeQL cleanup.

## Implementation notes

To be completed by the implementation session. Record:

- Phase 7a verification and bake evidence.
- Live Cloudflare resources and exact deletion confirmations.
- Files/resources deleted or retained and why.
- Database cleanup commands, environment, tables, and row counts, or an
  explicit statement that no rows were removed.
- Validation commands and results.
- Any unresolved findings and the next plan that owns them.
