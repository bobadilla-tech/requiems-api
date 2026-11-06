# API Performance Metrics Pipeline

The dashboard had no latency figures on API documentation pages. Developers
integrating an API had no way to know what p50/p95/p99 response times looked
like in production before committing to it. The ask was to surface real,
measured latency numbers — not synthetic benchmarks or made-up estimates —
directly in the dashboard docs and in the "Copy as Markdown" export used to
paste context into AI assistants.

The industry approach (Stripe, Cloudflare) is a committed JSON snapshot updated
by a benchmark script, not a live metrics endpoint. That's the pattern we
followed.

---

## What Changed

### 1. Integration Test Performance Instrumentation

`tests/integration/src/stats.ts` — a singleton that accumulates per-path timing
samples across every HTTP request made during a test run. Each call through
`client.ts` records a `(path, durationMs)` pair.

Key design decisions:

- **Merge on persist, not overwrite.** With vitest's `singleFork: true`, each
  test file gets its own module instance. `persist()` reads the existing temp
  file, merges its own samples in, and writes back — so data accumulates across
  files rather than each file clobbering the previous one.
- **`globalSetup.ts` clears the temp file once** before the run starts so
  successive benchmark runs don't compound.

### 2. Benchmark Script and `perf-baseline.json`

`tests/integration/package.json` — the `benchmark` script:

```
INTEGRATION_RUNS=50 UPDATE_PERF_BASELINE=true vitest run; pnpm run docs:perf:yaml
```

When `UPDATE_PERF_BASELINE=true` the custom reporter writes
`tests/integration/perf-baseline.json` — a committed file containing p50, p95,
p99, avg, min, max, and sample count per endpoint. The `;` (not `&&`) ensures
the YAML injection always runs even when some tests fail due to production
instability.

### 3. YAML Injection into Dashboard Docs

`tests/integration/scripts/inject-perf-yaml.ts` upserts a `performance:` block
into every `apps/dashboard/config/api_docs/{api_id}.yml` file that has matching
benchmark data.

Block format:

```yaml
performance:
  p50_ms: 871
  p95_ms: 1111
  p99_ms: 1597
  avg_ms: 949
  samples: 50
  measured_at: "2026-04-16"
```

`tests/integration/scripts/endpoint-yaml-map.ts` maps the 38 measured endpoint
paths to their 33 dashboard `api_id` slugs. When multiple paths share one doc
(e.g. `/v1/technology/base64/encode` and `/decode` both map to `base64`), the
entry with the most samples wins.

The script is idempotent — re-running it updates values without corrupting
multiline strings or comments in the YAML files. 33 out of 33 dashboard YAML
files now carry a `performance:` block.

### 4. Dashboard Performance Section

`apps/dashboard/app/views/apis/show.html.erb` — a Performance section was added
between Overview and Endpoints, conditional on
`@documentation["performance"].present?`.

Three stat cards (Tailwind grid, responsive 1→3 columns):

| Card | Subtitle               | Background     |
| ---- | ---------------------- | -------------- |
| p50  | Median                 | `bg-blue-50`   |
| p95  | 95th percentile        | `bg-indigo-50` |
| p99  | 99th percentile (tail) | `bg-violet-50` |

Each card: large bold ms value, label row, muted description. The sidebar nav
gains a "Performance" link (also conditional) so the section is reachable by
keyboard navigation and page anchor.

### 5. Markdown and Text Export Helpers

`apps/dashboard/app/helpers/apis_helper.rb` — both
`api_documentation_as_markdown()` and `api_documentation_as_text()` were updated
to include performance data. This means the "Copy as Markdown" button and "Open
in Claude / ChatGPT" prompt exports include the p50/p95/p99 table — useful
context for AI assistants helping developers integrate.

### 6. Integration Test Reliability Fixes

Several rounds of fixes made the benchmark and test suite more robust:

**`repeat()` — mode-aware resilience**

In test mode (`pnpm test`), any error propagates immediately — fast failure
signal when the backend is down. In benchmark mode
(`UPDATE_PERF_BASELINE=true`), errors are caught so partial data is collected
from a flaky backend, but the loop bails out after 3 consecutive failures (3 × 8
s = 24 s max per dead endpoint) to prevent runaway async work.

The original naive resilience (catch everything, retry all 50 times) caused test
files to take 941 s because `singleFork: true` meant subsequent tests waited for
all the background async work to drain.

**Per-request `AbortSignal.timeout`**

Every `fetch()` in `client.ts` and `gateway.test.ts` now carries
`AbortSignal.timeout(cfg.requestTimeoutMs)` (default 8 s). Previously the
gateway health check test used raw `fetch()` with no timeout, so a down backend
made it hang for the full vitest 30 s test timeout.

**`testTimeout` by mode**

`vitest.config.ts` sets `testTimeout: 120_000` in benchmark mode and `30_000` in
normal test mode. A well-behaved benchmark test legitimately runs 50 iterations;
the extended timeout prevents vitest from killing it before all samples are
collected.

**Snapshot reorganization**

`/v1/entertainment/advice` and `/v1/entertainment/quotes/random` were being
tested inside `text.test.ts` (a historical accident). They were moved to
`entertainment.test.ts` and their snapshot entries were relocated from
`tests/snapshots/text.snap.json` to `tests/snapshots/entertainment.snap.json`.

---

## Files Changed

| File                                                 | Change                                                  |
| ---------------------------------------------------- | ------------------------------------------------------- |
| `tests/integration/src/stats.ts`                     | Merge-on-persist logic                                  |
| `tests/integration/src/globalSetup.ts`               | Clear temp file before run                              |
| `tests/integration/src/reporter.ts`                  | Write `perf-baseline.json` in benchmark mode            |
| `tests/integration/src/config.ts`                    | `INTEGRATION_RUNS`, `REQUEST_TIMEOUT_MS` env vars       |
| `tests/integration/src/client.ts`                    | `AbortSignal.timeout` on every fetch                    |
| `tests/integration/src/helpers.ts`                   | Mode-aware `repeat()` with consecutive-failure bail-out |
| `tests/integration/vitest.config.ts`                 | Dynamic `testTimeout` based on benchmark mode           |
| `tests/integration/package.json`                     | `benchmark` and `docs:perf:yaml` scripts                |
| `tests/integration/scripts/endpoint-yaml-map.ts`     | Path → api_id map (38 endpoints → 33 docs)              |
| `tests/integration/scripts/inject-perf-yaml.ts`      | Upserts `performance:` into YAML docs                   |
| `tests/integration/perf-baseline.json`               | Committed snapshot (38 endpoints, 50 samples each)      |
| `tests/integration/src/suites/entertainment.test.ts` | Advice + quotes tests moved here                        |
| `tests/integration/src/suites/text.test.ts`          | Advice + quotes tests removed                           |
| `tests/integration/src/suites/gateway.test.ts`       | `AbortSignal.timeout` on all raw fetch calls            |
| `tests/snapshots/entertainment.snap.json`            | `advice` + `quotes_random` entries added                |
| `tests/snapshots/text.snap.json`                     | `advice` + `quotes_random` entries removed              |
| `apps/dashboard/app/views/apis/show.html.erb`        | Performance section + sidebar link                      |
| `apps/dashboard/app/helpers/apis_helper.rb`          | Markdown + text export include perf table               |
| `apps/dashboard/config/api_docs/*.yml` (33 files)    | `performance:` block injected                           |
