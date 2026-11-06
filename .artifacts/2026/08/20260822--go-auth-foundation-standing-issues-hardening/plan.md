# Go Auth Foundation — Standing Issues Hardening

**Status:** Completed and verified 2026-08-22

## Historical context

This artifact records the disposition of the standing issues raised during the
Go auth foundation review, including the PR #966 security findings and the
associated inline and outside-diff comments. Each finding was checked against
the current tree before changes were made. Fixes already present in the tree
were verified rather than duplicated; stale or no-longer-valid findings were
left unchanged.

## Outcome

The current implementation now has the following verified behavior:

- The Auth Gateway preserves and forwards the complete `requiems-api-key` header
  to the Go API. The Go route authenticates the forwarded key, while the gateway
  remains the public edge authentication boundary.
- API-key cache entries keyed by prefix are only candidate records. Every
  request still verifies the complete presented key with bcrypt, including
  shared-prefix, prefix-only, and altered-suffix cases.
- Wrapped response status `0` is logged as `200`, matching `net/http` when a
  handler writes neither headers nor a body.
- Rails API-key invalidation runs after commit for destroy and active/revoked
  updates. `User#ban!` collects the rows affected by `update_all` and explicitly
  invalidates their Go auth-cache entries.
- The API-key prefix index migration uses PostgreSQL’s concurrent algorithm
  outside a migration transaction.
- Seed output contains only the key prefix. A raw local-development key is
  accepted only through the protected `LOCAL_DEV_API_KEY` environment variable,
  and existing keys retain their prior behavior.
- Numeric API-key IDs remain available in operational logs without logging API
  key secrets; the usage log contains a narrowly scoped CodeQL explanation.
- Plan lookup failures preserve stale limits when available. A cold lookup
  without stale limits produces HTTP 503 for quota enforcement, while rate
  limiting continues to fail open as designed.
- Usage accounting uses the request’s credit count consistently in Redis and
  PostgreSQL. Redis reservations are atomic; when Redis is unavailable, the
  quota path fails closed instead of relying on a non-atomic PostgreSQL sum.
  Usage-ledger writes use a five-second context that survives client
  cancellation.
- Billing-cycle anchors are clamped to the final day of the current month, with
  the prior-month fallback preserved.
- Authentication selects one eligible active/trialing subscription
  deterministically, and the plans migration backs the accepted subscription
  values with a database foreign key.
- API-key creation retries only the specific unique-prefix persistence conflict,
  including concurrent insert races.
- Production Redis configuration requires an operator-supplied `REDIS_MAXMEMORY`
  value after capacity validation; development keeps a bounded example value.

## Decisions and non-goals

1. The Go API does not use the prefix cache as an authentication decision. A
   cache hit saves the candidate lookup, not the bcrypt verification.
2. Redis-unavailable quota enforcement does not use an unlocked PostgreSQL
   aggregate as a substitute for a reservation. Rejecting with 503 preserves
   quota correctness until an atomic fallback is available.
3. No server-peppered digest was introduced: the required bcrypt verification
   remains the security boundary, and the cache-hit behavior is covered by
   regression tests.
4. Findings that were already correctly implemented in the current tree were
   recorded as verified rather than reworked. The historical Phase 5 plan that
   was already present in the worktree was preserved untouched.

## Targeted corrections recorded with this artifact

- Corrected the cycle-start regression expectations for February: anchor days
  29, 30, and 31 all clamp to February 28 and therefore fall back to January 28
  for the test date used.
- Updated usage middleware comments to describe atomic Redis reservation and
  fail-closed behavior accurately.
- Updated the Phase 0–1, Phase 2, and Phase 3–4 planning documents so their
  historical security and failure-mode descriptions match the current
  implementation.

## Validation

Completed checks:

- `go test ./...`
- Auth Gateway `pnpm exec vitest run` — 80 tests passed
- Ruby syntax checks for the affected models, seeds, migration, and API-key
  generator
- `docker compose --env-file .env.example -f docker-compose.dev.yml config --quiet`
- `git diff --check`

The Rails test suite was not runnable in this environment because Bundler 2.6.9
is not installed locally and the dashboard container was not running. Production
Compose intentionally remains gated until an operator supplies a
capacity-validated `REDIS_MAXMEMORY` value.
