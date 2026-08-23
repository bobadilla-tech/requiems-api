/**
 * Rate Limit Validation — tests/load/scenarios/rate-limit.ts
 *
 * Purpose
 * -------
 * Confirm that the Go API directly enforces
 * per-plan rate limits.  Three behaviours are validated:
 *
 *   1. Burst over limit  — rapid-fire >N req/min on a free-plan key triggers
 *                          429 responses once the limit (30 req/min) is hit.
 *   2. Under-limit pass  — a developer-plan key (5 000 req/min) never gets
 *                          a 429 even when many requests are sent quickly.
 *   3. Recovery          — after the 60-second window resets, the same
 *                          free-plan key is accepted again.
 *
 * Rate limits are read from the PostgreSQL plans table by Go:
 *   free         →  30 req / min
 *   developer    →  5 000 req / min
 *   business     →  10 000 req / min
 *   professional → 50 000 req / min
 *
 * Usage
 * -----
 *   k6 run tests/load/scenarios/rate-limit.ts
 *
 * The scenario is deliberately short (≈ 3 minutes total) so it fits in a
 * typical CI pipeline step.
 */

import http from "k6/http";
import { check, group, sleep } from "k6";
import { Counter, Rate } from "k6/metrics";
import { Options } from "k6/options";

import { API_KEYS, BASE_URL, RATE_LIMITS, SummaryData } from "../config.ts";

// ---------------------------------------------------------------------------
// Custom metrics
// ---------------------------------------------------------------------------

/** Total 429 responses received */
const rateLimitHits = new Counter("rate_limit_hits");

/** Requests that returned a non-429 error */
const unexpectedErrors = new Counter("unexpected_errors");

/** Rate of requests that were correctly allowed (2xx) */
const allowedRate = new Rate("requests_allowed");

/** Rate of requests that were correctly rate-limited (429) */
const blockedRate = new Rate("requests_blocked");

// ---------------------------------------------------------------------------
// k6 options
// ---------------------------------------------------------------------------

export const options: Options = {
  /**
   * Three sequential scenarios:
   *  1. free_burst     — 1 VU, 35 iterations in ≤10 s  (expect 429 after ~30)
   *  2. developer_safe — 1 VU, 60 iterations quickly    (expect 0 x 429)
   *  3. recovery       — after 65 s window, free key ok again
   */
  scenarios: {
    free_burst: {
      executor: "shared-iterations",
      vus: 1,
      iterations: 35,
      maxDuration: "20s",
      startTime: "0s",
    },
    developer_safe: {
      executor: "shared-iterations",
      vus: 1,
      iterations: 60,
      maxDuration: "20s",
      // start after free_burst finishes (+5 s buffer)
      startTime: "25s",
    },
    recovery: {
      executor: "shared-iterations",
      vus: 1,
      iterations: 5,
      maxDuration: "20s",
      // start 65+ seconds after free_burst so the 60-s window resets
      startTime: "95s",
    },
  },

  thresholds: {
    // At least some requests must have been rate-limited in free_burst
    rate_limit_hits: ["count>0"],
    // No unexpected (non-429 / non-2xx) errors
    unexpected_errors: ["count==0"],
  },
};

// Shared probe endpoint — a lightweight GET that every plan can access
const PROBE_PATH = "/v1/text/advice";

// ---------------------------------------------------------------------------
// Default function
// ---------------------------------------------------------------------------

export default function (): void {
  const scenario = __ENV["K6_SCENARIO_NAME"];

  if (scenario === "free_burst") {
    runFreeBurst();
  } else if (scenario === "developer_safe") {
    runDeveloperSafe();
  } else if (scenario === "recovery") {
    runRecovery();
  } else {
    // Fallback: run free_burst logic so the script is usable standalone
    runFreeBurst();
  }
}

// ---------------------------------------------------------------------------
// Scenario implementations
// ---------------------------------------------------------------------------

/**
 * Send 35 rapid requests using the free-plan key.
 * The first ~30 should return 2xx; the rest should return 429.
 */
function runFreeBurst(): void {
  group("free_plan_burst", () => {
    const url = `${BASE_URL}${PROBE_PATH}`;
    const params = {
      headers: { "requiems-api-key": API_KEYS.free },
      tags: { scenario: "free_burst" },
    };

    const res = http.get(url, params);

    if (res.status === 429) {
      rateLimitHits.add(1);
      blockedRate.add(1);

      const ok = check(res, {
        "429 has rate-limit error body": (r) => {
          try {
            const body = JSON.parse(r.body as string) as { error?: string };
            return (
              body.error === "rate_limit_exceeded" ||
              body.error === "too_many_requests"
            );
          } catch {
            // Some gateways return a plain string — still a valid 429
            return (r.body as string).length > 0;
          }
        },
        "Retry-After or X-RateLimit headers present": (r) =>
          r.headers["Retry-After"] !== undefined ||
          r.headers["X-Ratelimit-Limit"] !== undefined ||
          r.headers["X-Ratelimit-Remaining"] !== undefined,
      });

      if (!ok) {
        unexpectedErrors.add(1);
      }
    } else if (res.status >= 200 && res.status < 300) {
      allowedRate.add(1);

      check(res, {
        "allowed response has usage headers": (r) =>
          r.headers["X-Requests-Remaining"] !== undefined ||
          r.headers["X-Ratelimit-Remaining"] !== undefined,
      });
    } else {
      unexpectedErrors.add(1);
      console.error(
        `[rate-limit/free_burst] Unexpected status ${res.status}: ${
          (res.body as string).substring(0, 200)
        }`,
      );
    }

    // No sleep — we intentionally burst to saturate the 30 req/min limit
  });
}

/**
 * Send 60 rapid requests using the developer-plan key (5 000 req/min).
 * Every response should be 2xx — no rate limiting expected at this volume.
 */
function runDeveloperSafe(): void {
  group("developer_plan_safe", () => {
    const url = `${BASE_URL}${PROBE_PATH}`;
    const params = {
      headers: { "requiems-api-key": API_KEYS.developer },
      tags: { scenario: "developer_safe" },
    };

    const res = http.get(url, params);

    const ok = check(res, {
      "developer plan not rate-limited": (r) => r.status !== 429,
      "developer plan returns 2xx": (r) => r.status >= 200 && r.status < 300,
    });

    if (!ok) {
      if (res.status === 429) {
        unexpectedErrors.add(1);
        console.error(
          "[rate-limit/developer_safe] Unexpected 429 — developer plan should not be rate-limited at 60 req/min",
        );
      } else {
        unexpectedErrors.add(1);
        console.error(
          `[rate-limit/developer_safe] Unexpected status ${res.status}: ${
            (res.body as string).substring(0, 200)
          }`,
        );
      }
    } else {
      allowedRate.add(1);
    }
  });
}

/**
 * After ~65 seconds the 60-second rate-limit window for the free plan resets.
 * These requests should succeed again (2xx).
 */
function runRecovery(): void {
  group("free_plan_recovery", () => {
    const url = `${BASE_URL}${PROBE_PATH}`;
    const params = {
      headers: { "requiems-api-key": API_KEYS.free },
      tags: { scenario: "recovery" },
    };

    const res = http.get(url, params);

    const ok = check(res, {
      "recovered after rate-limit window": (r) =>
        r.status >= 200 && r.status < 300,
    });

    if (ok) {
      allowedRate.add(1);
    } else {
      unexpectedErrors.add(1);
      console.error(
        `[rate-limit/recovery] Expected 2xx after window reset but got ${res.status}: ${
          (res.body as string).substring(0, 200)
        }`,
      );
    }

    sleep(0.5);
  });
}

// ---------------------------------------------------------------------------
// Summary
// ---------------------------------------------------------------------------

export function handleSummary(data: SummaryData): Record<string, string> {
  const m = data.metrics;

  const hits = m["rate_limit_hits"]?.values["count"] ?? 0;
  const unexpected = m["unexpected_errors"]?.values["count"] ?? 0;
  const allowed = (m["requests_allowed"]?.values["rate"] ?? 0) * 100;
  const freeLimit = RATE_LIMITS.free;

  const summary = [
    "=".repeat(60),
    "  RATE LIMIT VALIDATION RESULTS",
    "=".repeat(60),
    `  Target URL         : ${BASE_URL}`,
    `  Free plan limit    : ${freeLimit} req/min`,
    `  429 responses seen : ${hits}  (expected > 0 for free burst)`,
    `  Unexpected errors  : ${unexpected}`,
    `  Allowed rate       : ${allowed.toFixed(1)} %`,
    "=".repeat(60),
    hits > 0
      ? "  ✓ Rate limiting is ACTIVE"
      : "  ✗ Rate limiting NOT triggered — check gateway config",
    unexpected === 0
      ? "  ✓ No unexpected errors"
      : `  ✗ ${unexpected} unexpected error(s) — investigate logs`,
    "=".repeat(60),
  ].join("\n");

  console.log(summary);

  return { stdout: summary + "\n" };
}
