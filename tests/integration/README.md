# Integration Tests — Full Request Flow

Local end-to-end tests for the complete `Worker → Backend → Database` path.
These tests hit the **real production API** (or a local dev stack) using a
genuine API key and are designed to:

- Verify the full request flow works end-to-end
- Detect **flaky behaviour** by running each scenario multiple times
- Track **response-time regressions** (min / avg / p95 / p99 / max per endpoint)
- Snapshot **response body shapes** so structural API changes are caught early

> **These tests are not part of CI.** Run them locally when you want to validate
> a change against production.

## Prerequisites

| Requirement           | Notes                                            |
| --------------------- | ------------------------------------------------ |
| Node.js ≥ 20          | Any LTS release                                  |
| pnpm ≥ 10             | `npm i -g pnpm`                                  |
| A `requiem_*` API key | Obtain from [requiems.xyz](https://requiems.xyz) |

## Setup

```bash
cd tests/integration

# Install dependencies
pnpm install

# Copy the example env file and fill in your API key
cp .env.example .env
# Edit .env and set REQUIEMS_API_KEY=requiem_...
```

## Running the Tests

```bash
# Run all integration suites once
pnpm test

# Watch mode (re-runs on file changes — useful during local development)
pnpm test:watch

# Open the Vitest UI for interactive exploration
pnpm test:ui
```

## Configuration

All settings live in `.env` (see `.env.example` for defaults):

| Variable           | Default                    | Description                                              |
| ------------------ | -------------------------- | -------------------------------------------------------- |
| `REQUIEMS_API_KEY` | _(required)_               | Your production API key                                  |
| `API_BASE_URL`     | `https://api.requiems.xyz` | Gateway URL to test against                              |
| `INTEGRATION_RUNS` | `3`                        | How many times each scenario is repeated for timing data |

Set `API_BASE_URL=http://localhost:4455` to test against a local `wrangler dev`
instance instead of production.

## Output

### Console table

A timing summary is printed after every run:

```
📊  Integration Test Performance Summary
──────────────────────────────────────────────────────────────────────────────────
Path                              Samples  Min(ms)  Avg(ms)  P50(ms)  P95(ms)  P99(ms)  Max(ms)
──────────────────────────────────────────────────────────────────────────────────
/v1/convert/base64/decode               3      183      195      198      206      206      206
/v1/convert/base64/encode               3      165      178      181      188      188      188
/v1/email/disposable/check              6      201      223      218      267      267      267
/v1/text/advice                         9      145      162      158      189      189      189
...
```

### JSON performance report

A machine-readable report is saved to `tests/integration/reports/` after each
run so you can track trends over time:

```
📁  Performance report saved to: reports/perf-2025-04-13T03-24-51-725Z.json
```

Each report contains `min`, `avg`, `p50`, `p95`, `p99`, `max` for every
endpoint, along with all raw sample values so you can build charts or compare
runs programmatically.

### Response shape snapshots

The first time each endpoint is called its response shape (key names + JSON
types) is written to `tests/integration/snapshots/`. Subsequent runs compare the
live shape against the snapshot and fail if they differ:

```
Response shape mismatch for "advice":
  Expected : {data:{advice:string,id:number},metadata:{...}}
  Received : {data:{advice:string,id:number,new_field:string},metadata:{...}}

If this change is intentional, delete the snapshot and re-run.
```

To accept a shape change, delete the relevant `.snap.json` file and re-run.

## Directory Structure

```
tests/integration/
├── .env.example         # Environment variable template
├── package.json
├── tsconfig.json
├── vitest.config.ts
├── reports/             # Auto-generated JSON perf reports (git-ignored)
├── snapshots/           # Auto-generated response shape snapshots
└── src/
    ├── client.ts        # HTTP client with timing instrumentation
    ├── config.ts        # Env-based configuration
    ├── helpers.ts       # Shared test helpers (assertEnvelope, repeat)
    ├── reporter.ts      # Custom Vitest reporter (perf table + JSON output)
    ├── setup.ts         # Global setup: loads .env
    ├── snapshot.ts      # Snapshot capture and comparison utilities
    ├── stats.ts         # Response-time statistics (p50/p95/p99/…)
    └── suites/
        ├── gateway.test.ts      # Auth, headers, error paths
        ├── text.test.ts         # /v1/text/*
        ├── email.test.ts        # /v1/email/*
        ├── finance.test.ts      # /v1/finance/*
        ├── entertainment.test.ts # /v1/entertainment/*
        ├── tech.test.ts         # /v1/tech/*
        ├── convert.test.ts      # /v1/convert/*
        └── misc.test.ts         # /v1/misc/*
```

## Detecting Regressions

Run the suite before and after a change, then diff the JSON report files:

```bash
# Before your change
pnpm test
cp reports/perf-*.json /tmp/before.json

# After your change
pnpm test
cp reports/perf-*.json /tmp/after.json

# Quick comparison
node -e "
const before = require('/tmp/before.json');
const after  = require('/tmp/after.json');
const map = Object.fromEntries(before.endpoints.map(e => [e.path, e]));
for (const e of after.endpoints) {
  const b = map[e.path];
  if (!b) { console.log('NEW:', e.path); continue; }
  const delta = e.avg - b.avg;
  if (Math.abs(delta) > 50) {
    console.log(\`\${delta > 0 ? '⬆️ SLOWER' : '⬇️ FASTER'}  \${e.path}  \${delta > 0 ? '+' : ''}\${delta}ms (avg)\`);
  }
}
"
```
