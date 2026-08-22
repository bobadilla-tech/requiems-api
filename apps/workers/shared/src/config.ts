/**
 * Plan request limits (monthly quotas) shared with the Rails dashboard.
 *
 * IMPORTANT: This JSON file is the single source of truth for per-plan
 * request limits. It is also read by Rails (apps/dashboard) to populate
 * User::PLAN_LIMITS. Update plan-limits.json instead of editing the
 * numbers here directly.
 */
import PLAN_REQUEST_LIMITS from "../plan-limits.json";

/**
 * Plan configurations
 *
 * All plans use monthly request quotas
 * Quotas reset at the start of each billing cycle (monthly)
 *
 * Note: Some endpoints count as multiple requests (see ENDPOINT_MULTIPLIERS)
 *
 * Pricing:
 * - Free: $0/month - 500 requests
 * - Developer: $30/month - 100,000 requests
 * - Business: $75/month - 1,000,000 requests
 * - Professional: $150/month - 10,000,000 requests
 * - Enterprise: Custom - Unlimited requests
 */
export const PLANS = {
  free: {
    requestLimit: PLAN_REQUEST_LIMITS.free,
    ratePerMinute: 30,
  },
  developer: {
    requestLimit: PLAN_REQUEST_LIMITS.developer,
    ratePerMinute: 5000,
  },
  business: {
    requestLimit: PLAN_REQUEST_LIMITS.business,
    ratePerMinute: 10000,
  },
  professional: {
    requestLimit: PLAN_REQUEST_LIMITS.professional,
    ratePerMinute: 50000,
  },
  enterprise: {
    requestLimit: Number.POSITIVE_INFINITY,
    ratePerMinute: Number.POSITIVE_INFINITY,
  },
} as const;

/**
 * Plan names
 * https://github.com/bobadilla-tech/requiems-api/docs/core/business.md
 */
export type PlanName = keyof typeof PLANS;

export const PLAN_NAMES = Object.keys(PLANS) as PlanName[];

/**
 * Request multipliers - how many requests an endpoint counts as
 *
 * DEFAULT: 1 request per API call
 * This map ONLY lists endpoints that count as MORE than 1 request.
 *
 * Most endpoints = 1 request (default)
 * Some endpoints require expensive operations = 2x+ requests (listed below)
 *
 * Examples of expensive operations:
 * - API calls to external services (2-3x)
 * - Complex AI/ML inference (3-5x)
 * - Large data processing (2-3x)
 *
 * IMPORTANT: Keep this in sync with your Go backend routes!
 * When adding expensive endpoints, add them here.
 */
export const ENDPOINT_MULTIPLIERS = new Map<string, number>([
  // Dictionary operations count as 2 requests
  ["GET /v1/text/dictionary", 2],
  ["GET /v1/text/thesaurus", 2],
  // Future expensive endpoints:
  // ["GET /v1/ai/image-recognition", 5],
  // ["POST /v1/text/translate", 3],
]);

/**
 * Pre-computed prefix patterns for efficient matching
 * Format: [method, pathPrefix, multiplier]
 */
const ENDPOINT_PREFIXES: Array<[string, string, number]> = Array.from(
  ENDPOINT_MULTIPLIERS.entries(),
).map(([route, multiplier]) => {
  const [method, path] = route.split(" ", 2);
  return [method, path, multiplier];
});

/**
 * Default multiplier for endpoints not listed above
 * This is the multiplier for 90%+ of endpoints
 */
export const DEFAULT_REQUEST_MULTIPLIER = 1;

/**
 * Maximum number of records returned by the usage export endpoint.
 * Caps the `limit` query parameter to prevent oversized D1 result sets.
 */
export const USAGE_EXPORT_MAX_LIMIT = 5000;

/**
 * Maximum number of endpoints returned by the analytics by-endpoint endpoint.
 * Caps the `limit` query parameter to bound the result set size.
 */
export const ANALYTICS_ENDPOINT_MAX_LIMIT = 100;

/**
 * Get the request multiplier for an endpoint
 * Matches exact path first, then tries prefix matching for dynamic routes
 */
export function getRequestMultiplier(method: string, pathname: string): number {
  const exactKey = `${method} ${pathname}`;

  const exactMatch = ENDPOINT_MULTIPLIERS.get(exactKey);
  if (exactMatch !== undefined) {
    return exactMatch;
  }

  // Prefix matching for dynamic routes (e.g., /v1/finance/stocks/:symbol)
  for (const [routeMethod, routePath, multiplier] of ENDPOINT_PREFIXES) {
    if (method === routeMethod && pathname.startsWith(routePath)) {
      return multiplier;
    }
  }

  return DEFAULT_REQUEST_MULTIPLIER;
}
