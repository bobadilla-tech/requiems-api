/**
 * Shared k6 configuration for the Requiem API load testing suite.
 *
 * Targets the direct Go API, which owns authentication, rate limiting, quota,
 * and usage recording.
 */

import { Params } from "k6/http";

/** Plan tier names */
export type PlanKey = "free" | "developer" | "business" | "professional";

/** A single API endpoint descriptor used across scenarios */
export interface SampleEndpoint {
  method: "GET" | "POST";
  path: string;
  body?: string;
}

/**
 * Shape of a single metric's values in the handleSummary data object.
 * k6 exposes named statistics (rate, count, avg, p(95), etc.) as plain numbers.
 */
export interface MetricValues {
  [key: string]: number;
}

/** A single metric entry in the handleSummary data object. */
export interface MetricData {
  values: MetricValues;
}

/**
 * Data object passed to handleSummary() at the end of a k6 run.
 * Only the `metrics` field is typed here — k6 also provides `rootGroup`, etc.
 */
export interface SummaryData {
  metrics: Record<string, MetricData | undefined>;
}

/** Base URL of the Auth Gateway (edge proxy). Override with BASE_URL env var. */
export const BASE_URL: string = __ENV.BASE_URL || "http://localhost:8080";

/**
 * Dev API keys — one per plan tier.
 * Matches the keys written by seed-dev.ts.
 */
const localKey = __ENV.LOCAL_DEV_API_KEY || __ENV.REQUIEMS_API_KEY || "";
if (!/^requiem_[0-9A-Za-z]{24}$/.test(localKey)) {
  throw new Error("Set LOCAL_DEV_API_KEY or REQUIEMS_API_KEY to a valid requiem_<24 alphanumeric characters> key");
}
export const API_KEYS: Record<PlanKey, string> = {
  free: __ENV.API_KEY_FREE || localKey,
  developer: __ENV.API_KEY_DEVELOPER || localKey,
  business: __ENV.API_KEY_BUSINESS || localKey,
  professional: __ENV.API_KEY_PROFESSIONAL || localKey,
};

/**
 * Per-plan rate limits (requests per minute).
 * Values match the rate_limit_per_minute column in PostgreSQL plans.
 */
export const RATE_LIMITS: Record<PlanKey, number> = {
  free: 30,
  developer: 5000,
  business: 10000,
  professional: 50000,
};

/**
 * Returns k6 request params that authenticate as the given plan tier.
 */
export function authParams(plan: PlanKey = "developer"): Params {
  return {
    headers: {
      "requiems-api-key": API_KEYS[plan],
      "Content-Type": "application/json",
    },
  };
}

/**
 * Common performance thresholds applied to baseline and concurrent scenarios.
 *
 * - http_req_failed: fewer than 1 % of requests may fail (non-2xx/3xx).
 * - http_req_duration p(95): 95 % of requests must complete within 500 ms.
 * - http_req_duration p(99): 99 % of requests must complete within 1 000 ms.
 */
export const DEFAULT_THRESHOLDS: Record<string, string[]> = {
  http_req_failed: ["rate<0.01"],
  http_req_duration: ["p(95)<500", "p(99)<1000"],
};

/**
 * Lightweight set of API endpoints used across multiple scenarios.
 * Each entry is { method, path, body? } relative to BASE_URL.
 *
 * Service mount points (from apps/api/app/routes_v1.go):
 *   /v1/entertainment → entertainment services
 *   /v1/finance       → finance services
 *   /v1/health        → health / fitness services
 *   /v1/networking    → IP, domain, MX, WHOIS, disposable email
 *   /v1/places        → places services
 *   /v1/technology    → conversion, QR, barcode, counter, random user, etc.
 *   /v1/text          → text, AI, and email normalization
 *   /v1/validation    → email validate, phone, profanity
 */
export const SAMPLE_ENDPOINTS: SampleEndpoint[] = [
  // health check (no auth required)
  { method: "GET", path: "/healthz" },

  // text
  { method: "GET", path: "/v1/text/lorem?paragraphs=1" },
  { method: "GET", path: "/v1/text/words/random" },
  {
    method: "POST",
    path: "/v1/text/spellcheck",
    body: JSON.stringify({ text: "Ths sentence has erors." }),
  },
  {
    method: "POST",
    path: "/v1/text/sentiment",
    body: JSON.stringify({ text: "I love this product!" }),
  },

  // entertainment
  { method: "GET", path: "/v1/entertainment/advice" },
  { method: "GET", path: "/v1/entertainment/quotes/random" },
  { method: "GET", path: "/v1/entertainment/facts" },
  { method: "GET", path: "/v1/entertainment/trivia" },
  { method: "GET", path: "/v1/entertainment/jokes/dad" },

  // networking
  { method: "GET", path: "/v1/networking/disposable/domain/mailinator.com" },

  // validation
  {
    method: "POST",
    path: "/v1/validation/email",
    body: JSON.stringify({ email: "test@example.com" }),
  },

  // technology
  { method: "GET", path: "/v1/technology/random-user" },
  {
    method: "GET",
    path: "/v1/technology/convert?from=km&to=mi&value=10",
  },
  { method: "GET", path: "/v1/technology/useragent" },
  { method: "GET", path: "/v1/technology/password" },
  { method: "GET", path: "/v1/technology/units" },

  // places
  { method: "GET", path: "/v1/places/cities/London" },
  { method: "GET", path: "/v1/places/postal/10001?country=US" },

  // networking — IP
  { method: "GET", path: "/v1/networking/ip" },

  // finance
  {
    method: "GET",
    path: "/v1/finance/mortgage?principal=300000&rate=6.5&years=30",
  },
  { method: "GET", path: "/v1/finance/exchange-rate?from=USD&to=EUR" },

  // health
  { method: "GET", path: "/v1/health/exercises/random" },
];
