# Usage Pipeline Fixes — Design Plan

This document describes two bugs that caused the Usage & Analytics dashboard to
show no data despite API calls being recorded in Cloudflare D1, and the fixes
applied to each.

---

## Why This Exists

The usage pipeline moves data from Cloudflare D1 (written by the Auth Gateway on
every request) into PostgreSQL (read by the Rails dashboard). Two independent
bugs, one in each half of that pipeline, caused every sync attempt to fail
silently and every analytics query to return empty results.

---

## Pipeline Overview

```
Auth Gateway → D1 (credit_usage)
                  ↓
         SyncD1UsageJob (every 5 min)
                  ↓
         D1SyncService → PostgreSQL (usage_logs)
                  ↓
         Dashboard::UsageController → Rails views
```

Analytics queries (by-endpoint, by-date, summary) read directly from D1 via the
API Management worker and are also affected by the storage-format bug.

---

## Bug 1 — D1 Datetime Format Mismatch

### Root cause

The Auth Gateway INSERT used SQLite's `datetime('now')` function, which produces
`'2026-04-16 13:01:00'` (space separator, no timezone). Every consumer of D1
data — the API Management analytics routes and the Auth Gateway's own quota
check — passes ISO 8601 timestamps as query parameters: `'2026-04-16T01:06:04Z'`
(T separator + Z).

SQLite has no native datetime type; it compares datetime columns as plain
strings. The space character (ASCII 32) sorts before `'T'` (ASCII 84), so
`'2026-04-16 13:01:00'` is lexicographically less than any `'2026-04-16T...'`
value. Every `WHERE used_at >= ?` returned zero rows regardless of the actual
time values.

This also silently broke quota enforcement: `getRequestUsage` in the Auth
Gateway uses the same `WHERE used_at >= ?` pattern, so all users appeared to
have 0 usage for their billing period.

### Fix

`apps/workers/auth-gateway/src/requests.ts` — replace `datetime('now')` in the
INSERT with a bound JavaScript parameter:

```typescript
// Before
.bind(apiKey, userId, endpoint, requests)
// VALUES (?, ?, ?, ?, datetime('now'))

// After
.bind(apiKey, userId, endpoint, requests, new Date().toISOString())
// VALUES (?, ?, ?, ?, ?)
```

ISO timestamps compare correctly as strings because they are zero-padded and use
a consistent separator. No query-side changes are needed — the analytics and
export queries already emit ISO timestamps as parameters.

Records written before this fix (space format) remain invisible to range
queries. No D1 data migration is performed; the existing records represent a
small window of data and the platform is early-stage.

---

## Bug 2 — String/Symbol Key Mismatch in D1SyncService

### Root cause

`D1SyncService#bulk_insert` (`apps/dashboard/app/services/d1_sync_service.rb`)
accessed usage records with Ruby symbol keys (`:api_key`, `:endpoint`,
`:credits_used`, `:used_at`). Faraday 2.x parses JSON responses with string keys
by default; `parse_response` forwarded the raw `body["usage"]` array without
conversion. Symbol access on a string-keyed Hash returns `nil`, so
`r[:api_key][0...12]` immediately raised `NoMethodError`.

`SyncD1UsageJob` only rescues `D1SyncService::Error`, not `NoMethodError`, so
the job failed on every run without updating the sync checkpoint. The next run
retried from the same timestamp, hit the same crash, and cycled indefinitely
into Sidekiq's dead queue.

### Fix

`apps/dashboard/app/services/d1_sync_service.rb` — symbolize keys at the
deserialization boundary in `parse_response`:

```ruby
# Before
usage: body["usage"] || []

# After
usage: (body["usage"] || []).map { |r| r.transform_keys(&:to_sym) }
```

---

## Bug 3 — Missing Unique Index for `insert_all`

### Root cause

`bulk_insert` calls:

```ruby
UsageLog.insert_all(values, unique_by: [:api_key_id, :used_at, :endpoint])
```

Rails' `insert_all` with `unique_by` generates
`INSERT ... ON CONFLICT (api_key_id, used_at, endpoint) DO NOTHING`. PostgreSQL
requires a unique index on exactly those columns to resolve the conflict target.
No such index existed — migration
`20260316170000_add_composite_index_to_usage_logs.rb` was created but left with
an empty `change` body.

Without the index, PostgreSQL raises `PG::InvalidColumnReference`, which (after
Bug 2 is fixed) would have caused the same silent failure cycle.

### Fix

`apps/dashboard/db/migrate/20260316170000_add_composite_index_to_usage_logs.rb`:

```ruby
def change
  add_index :usage_logs, [:api_key_id, :used_at, :endpoint],
            unique: true,
            name: "index_usage_logs_dedup"
end
```

The index also serves as a deduplication guard if the sync job ever processes an
overlapping window (e.g. after a Redis cache miss rewinds the checkpoint).
