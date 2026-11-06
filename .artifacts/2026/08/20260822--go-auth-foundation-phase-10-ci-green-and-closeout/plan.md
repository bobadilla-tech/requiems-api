# Go Auth Foundation — Phase 10: CI Green & Pre-Merge Closeout

**Status:** proposed **Depends on:** Phases 0-1, 2, 3-4,
standing-issues-hardening, 5, 6, 7, 8-9 (all shipped on
`feat/go-auth-foundations`, PR #966, currently `MERGEABLE`). **Goal of this
session:** this is very likely the last working session on PR #966. Everything
architectural is done — Cloudflare Workers/KV/D1 are deleted, the Go API is the
sole auth/rate-limit/usage enforcer, live traffic is smoke-tested. What's left
is not new feature work, it's making the PR itself clean and mergeable: real CI
failures, doc corpus rot, and one stray branch. Nothing here should touch
runtime behavior.

**Before doing anything else:** run `git log --oneline -5` and
`gh pr checks
966` fresh. This plan was drafted against commit `04ef902e`, but a
follow-up commit (`7420312e`, "docs: formatting") landed on the branch afterward
and already fixed 10.1's three `gocritic` findings, confirmed green on
`Go Lint (Advisory)`. The state below is what was true when this plan was
written — treat it as a starting hypothesis to re-verify, not as ground truth.
Don't redo work that already landed.

## Context: what's actually failing right now

Checked live via `gh pr checks 966` / `gh api .../check-runs`, last confirmed
fresh as of commit `7420312e`:

| Check                       | State                                                         | Real issue?                                                                                                                                                 |
| --------------------------- | ------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Go Tests                    | pass                                                          | —                                                                                                                                                           |
| Rails Tests                 | pass                                                          | —                                                                                                                                                           |
| MCP Tests                   | pass                                                          | —                                                                                                                                                           |
| Go Lint (Advisory)          | pass (already fixed by `7420312e`)                            | no — 10.1 below is effectively done, keep as a re-verify step only                                                                                          |
| codecov/patch               | **fail**                                                      | yes — but the specific per-file numbers below may themselves be stale, see 10.2's caution                                                                   |
| Codacy Static Code Analysis | **fail** (was `action_required` when last checked, re-verify) | likely false positive — 6 "critical SQL Injection" findings, all in `_test.go` files, all on fully parameterized `pool.Exec(ctx, "...", $1, $2, ...)` calls |
| CodeQL                      | neutral/skip                                                  | not a blocker                                                                                                                                               |

**Important scoping fact:** `main` has no branch protection
(`gh api repos/.../branches/main/protection` → `404 Branch not protected`).
Nothing here is mechanically blocking the merge button — `gh pr view 966`
already reports `mergeable: MERGEABLE`. This phase exists because a production
auth-migration PR going in red is bad practice, not because GitHub is stopping
it. That changes prioritization: real bugs first, coverage next, Codacy triage
last (it may not be fixable from the repo side at all — see 10.3).

## 10.1 — Go Lint: re-verify the 3 `gocritic` fixes are actually in

**Likely already done.** Commit `7420312e` (already on the branch) changed
`apps/api/platform/middleware/usage_test.go` at lines 156, 167, 319, swapping
`httptest.NewRequest(method, path, nil)` for
`httptest.NewRequest(method, path,
http.NoBody)` — exactly gocritic's
`httpNoBody` fix. `gh pr checks 966` shows `Go Lint (Advisory)` passing as of
that commit.

- Confirm: `cd apps/api && golangci-lint run ./...` comes back clean (CI pins
  `golangci-lint-action` at `v2.10`, matching what's installed).
- If it's already clean, mark this done and move on — don't re-touch the file.

## 10.2 — Raise `codecov/patch` coverage (53.57% → passing)

The PR comment (`gh pr view 966 --json comments`, codecov bot) lists the worst
offenders; the comment itself truncates at "and 7 more" so the first task here
is getting the _complete_ list, not just the top 10.

**Caution — the numbers in Step 2 below may already be stale.** They were
transcribed from a codecov comment posted against an earlier commit. A local
`go test -race -coverprofile=coverage.out ./...` run against the current branch
already shows `ratelimit.go`'s `Middleware` at 89.3% and `NewRateLimiter` at
100%, and `app.go`'s `New`/`Handler` at 83-100% — nowhere near the "0.00%"
figures quoted below. Treat Step 2's table as directional history, not current
fact. **Always trust a fresh Step 1 run over this table.** If the fresh numbers
already clear codecov's bar for a file, don't invent test cases to "fix" a gap
that no longer exists.

**Step 1 — get the full picture.** Run locally with real Postgres/Redis (same as
CI's `apps/api` job:
`DATABASE_URL=postgres://requiem:requiem@localhost:5432/requiem_test?sslmode=disable`,
`REDIS_URL=redis://127.0.0.1:6379/0`):

```
cd apps/api && go test -race -coverprofile=coverage.out ./...
go tool cover -func=coverage.out | sort -k3 -n
```

Cross-reference against `git diff main...feat/go-auth-foundations` to see
exactly which _new/changed_ lines are uncovered (codecov patch coverage is
diff-relative, not file-relative — a file can have decent overall coverage and
still fail patch coverage if the new lines specifically aren't hit). Also pull
the Rails side (`apps/dashboard/lib/tasks/playground_api_key.rake`,
`apps/dashboard/app/controllers/webhooks/lemonsqueezy_controller.rb`) via
`bin/rails test` with coverage.

**Step 2 — known gaps to close, ranked by size (from the partial codecov
comment):**

- `apps/api/platform/middleware/usage.go` — 38.88%, 74 missing + 3 partials
  (largest gap)
- `apps/api/platform/middleware/apikeyauth.go` — 45.63%, 53 missing + 3 partials
- `apps/api/platform/middleware/ratelimit.go` — 0.00%, 33 missing
- `apps/api/app/app.go` — 0.00%, 15 missing
- `apps/api/platform/middleware/plancache.go` — 59.25%, 8 missing + 3 partials
- `apps/api/platform/httpx/trustedproxy.go` — 77.77%, 5 missing + 3 partials
- `apps/dashboard/lib/tasks/playground_api_key.rake` — 76.19%, 5 missing
- `apps/dashboard/app/controllers/webhooks/lemonsqueezy_controller.rb` — 60.00%,
  4 missing
- `apps/api/platform/db/db.go` — 0.00%, 2 missing
- `apps/api/platform/reqredis/redis.go` — 0.00%, 2 missing
- plus 7 more files not enumerated in the truncated comment — surface these in
  Step 1.

**Step 3 — investigate whichever files the fresh Step 1 run actually shows as
gapped.** `ratelimit_test.go` and `app_test.go` already exist and already gate
on live Postgres/Redis (`t.Skip("TEST_DATABASE_URL/DATABASE_URL not set...")` at
`ratelimit_test.go:25` and `app_test.go:176,271`), and a local run already shows
`ratelimit.go`/`app.go` well-covered (see the caution above) — so those two are
probably not real gaps. `db.go`/`redis.go` are a different story: they have no
`_test.go` of their own, and `go test ./...` (no `-coverpkg`) only credits
coverage to a package's own test binary, so exercising them indirectly via
`app_test.go` doesn't count — that 0% is real and expected until a small test is
added directly. For whatever files the fresh run confirms as genuinely gapped
(start with `usage.go` and `apikeyauth.go`, the two large ones), read that
file's diff against `main`, identify which branches are new and untested, and
add targeted test cases — don't chase the coverage number, make sure the added
tests actually exercise real failure/edge paths (bad connection strings, Lua
script errors, plan lookup misses, etc.), consistent with this project's
existing pattern of real-infra integration tests over mocks.

**Step 4 — don't over-invest in db.go/redis.go.** These are thin infra-wiring
files (pool/client construction). If a line is pure boilerplate with no
meaningful branch to test (e.g. `return pgxpool.New(ctx, dsn)`), a one-line
"constructs successfully against a live pool" test is enough — don't invent
elaborate mocking for wrapper code.

**Exit bar for 10.2:** patch coverage comfortably above whatever `codecov/patch`
computes as its `auto` target (see `codecov.yml` —
`patch: target: auto, threshold: N/A`, i.e. it wants the diff at least as
well-covered as project average). Confirm by re-running the local coverage diff,
not by guessing.

## 10.3 — Triage the 6 Codacy "critical SQL Injection" findings

Pulled directly via `gh api repos/.../check-runs/<codacy-run-id>/annotations`.
All 6 are on parameterized `pgx` calls in test helpers:

- `apps/api/app/app_test.go:131`
- `apps/api/platform/middleware/apikeyauth_test.go:182`, `:189`, `:372`
- `apps/api/platform/middleware/ratelimit_test.go:56`, `:63`

Every one of these is `pool.Exec(ctx, "...$1, $2...", arg1, arg2, ...)` —
textbook-safe pgx parameter binding, no string concatenation or `fmt.Sprintf`
building the query. These read as false positives from Codacy's pattern matcher
(it likely flags any `.Exec()` call inside a multi-line raw string literal
regardless of placeholder usage).

- Confirm there is genuinely no unsafe concatenation anywhere nearby in these 6
  files (quick manual re-read, already spot-checked during planning — looked
  clean).
- Codacy's check conclusion is `action_required`, not `failure` — this is very
  likely a "needs human triage in the Codacy dashboard" state (dismiss as false
  positive), not something fixable by changing code, since the code is already
  correct. Confirm this reading by checking if the repo has a `.codacy.yml` or
  similar config file that can add an ignore rule for test files matching this
  specific pattern — if one exists and a scoped exclusion is reasonable (e.g.
  exclude `_test.go` files from this specific SQL-injection rule, not from all
  security rules), add it there. **Do not** broadly disable Codacy security
  scanning to make this go away.
- If no in-repo fix is available (i.e. it truly requires clicking "false
  positive" in the Codacy web UI), stop and say so plainly — flag it to the user
  as a 2-minute manual action they need to do themselves, don't spend session
  time trying to route around a dashboard-only gate.

## 10.4 — Correct stale doc corpus (auth-cache brute-force mischaracterization)

Per the digest of all prior phase docs:
`docs/plans/2026-08-21-go-auth-foundation-phase-2.md`,
`docs/plans/2026-08-21-go-auth-foundation-phase-3-4.md`, and
`docs/plans/2026-08-22-go-auth-foundation-phase-5.md` still describe the
auth-cache brute-force exposure as an **open** gap. It was actually fixed during
the PR #966 security-hardening pass
(`docs/plans/2026-08-22-go-auth-foundation-standing-issues-hardening.md`,
candidate-only cache + bcrypt-reverify-every-hit,
`apps/api/platform/middleware/apikeyauth.go:120-126`).

- Grep each of the three docs for the relevant section, add a short "Resolved —
  see standing-issues-hardening.md" note (don't rewrite history, just annotate).
- Phase 2's Final Notes paragraph (around line 477) lists the brute-force item
  alongside two others in the same breath — don't annotate only the brute-force
  one and leave the other two looking equally open by omission:
  - The "three manually-synced copies of plan-tier values" gap is now moot — the
    Worker's copy was deleted wholesale in Phase 8-9 (`apps/workers` has zero
    tracked files). Note it as moot, not "fixed."
  - The "usage-logs row-level dedup collision under rapid same-second traffic"
    item was explicitly logged as an _accepted design tradeoff_, not a bug to
    fix — note it as "accepted, not a gap," don't imply it needs action.
- This is pure documentation hygiene — no code changes. Keep it brief; these are
  historical planning docs, not living reference docs.

## 10.5 — Retire the superseded `prep/caddy-authenticated-origin-pulls` branch

Diffed directly: this branch has exactly one commit not on
`feat/go-auth-foundations` (`5c6d6d6a`, "prep: stage Caddy Authenticated Origin
Pulls config, NOT for merge yet"). Comparing its `Caddyfile` against the current
one on `feat/go-auth-foundations` shows the live branch's Phase 7a cutover
already implemented the same intent more completely — `api.requiems.xyz` (not
the prep branch's `internal.requiems.xyz`) with live `tls client_auth` AOP
enforcement, and the `X-Backend-Secret` gate the prep branch kept "until AOP is
confirmed" has already been removed on the live branch per Phase 7a. The prep
branch is stale/superseded, not a source of pending work.

- Confirm with the user before deleting anything (this is a branch-delete, treat
  as the standard destructive-op confirmation, not a rubber stamp).
- If confirmed: `git branch -d prep/caddy-authenticated-origin-pulls` (or `-D`
  since it has unmerged content w.r.t. `main` — but that content is superseded,
  not lost work) and
  `git push origin --delete prep/caddy-authenticated-origin-pulls`.

## 10.6 — Final pre-merge pass

- Re-run full local test suites one more time after 10.1–10.3's changes:
  `cd apps/api && go test -race ./...`, Rails `bin/rails test` (dashboard), MCP
  tests — confirm nothing regressed.
- Re-run `gh pr checks 966` and confirm Go Lint and codecov/patch are now green
  (Codacy per 10.3's outcome).
- Confirm the `PLAYGROUND_API_KEY` production-provisioning gap (tracked since
  Phase 6, still open per Phase 8-9) is captured somewhere durable as a
  post-merge owner action — it requires live production secrets access
  (`rails playground:provision_key` against a real key), which is not something
  this session can do. This is already tracked in
  `docs/core/v2-deployment-playbook.md` as an explicit checklist item, so this
  step is likely just a quick confirm, not new documentation work — only add a
  line if it's genuinely missing.
- Optional: PR title currently reads generically ("Requiems API v2" per
  CodeRabbit's pre-merge check, which passed but noted the title "does not
  specify the Go authentication, quota, or infrastructure work"). Consider
  tightening the PR title/description before merge — ask the user, don't do this
  unilaterally since it's their PR framing.

## Out of scope for this session

Anything gated on real production traffic (pool sizes, cache TTLs, Redis
`maxmemory`, dedicated Redis instance/DB split, `httpx.UsageCounter` backfill
across ~220 endpoints) — every prior phase doc agrees these are deliberately
deferred until there's real traffic to tune against, and that reasoning still
holds. Do not touch them.

## Exit criteria

- [ ] Go Lint (Advisory) check passes
- [ ] codecov/patch check passes (or, if genuinely infeasible in-session, a
      clear written explanation of exactly which lines remain uncovered and why)
- [ ] Codacy's 6 findings are either resolved via a scoped in-repo exclusion, or
      handed to the user as an explicit "click false-positive in the dashboard"
      action
- [ ] `phase-2.md` / `phase-3-4.md` / `phase-5.md` no longer misstate the
      auth-cache brute-force item as open
- [ ] `prep/caddy-authenticated-origin-pulls` deleted (both local and remote),
      only after explicit user confirmation
- [ ] Full local test suites (Go, Rails, MCP) green after all changes
- [ ] `PLAYGROUND_API_KEY` production provisioning confirmed captured as a
      durable post-merge action item
