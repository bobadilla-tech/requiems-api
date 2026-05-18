import { type Logger, type RequestCheckResult, withRetry } from "@requiem/workers-shared";
import type { WorkerBindings } from "./env";

/**
 * Get current request usage.
 *
 * Checks a 60-second KV cache first to avoid hitting D1 on every request.
 * On cache miss the result is written back to KV so the next request is served
 * from cache.
 *
 * IMPORTANT: Queries by user_id because all API keys for a user share the same quota.
 *
 * Note: Database tables still use "credit_" naming for historical reasons,
 * but we treat these as request counts in the code.
 */
export async function getRequestUsage(
  bindings: WorkerBindings,
  userId: string,
  period: "daily" | "monthly",
  billingCycleStart?: string,
  logger?: Logger,
): Promise<number> {
  const startDate = period === "daily" ? getTodayStart() : billingCycleStart || getMonthStart();
  const cacheKey = `quota:${userId}:${startDate}`;

  // KV cache hit — avoids a D1 aggregate query on every request
  const cached = await bindings.KV.get(cacheKey);

  if (cached !== null) {
    return Number(cached);
  }

  // Cache miss — query D1 and populate cache
  const result = await bindings.DB.prepare(`
    SELECT COALESCE(SUM(credits_used), 0) as total
    FROM credit_usage
    WHERE user_id = ? AND used_at >= ?
  `)
    .bind(userId, startDate)
    .first<{ total: number }>();

  const usage = result?.total || 0;

  // Write to KV with 60-second TTL (best-effort, don't block on failure)
  bindings.KV.put(cacheKey, usage.toString(), { expirationTtl: 60 }).catch((err) => {
    logger?.warn("KV cache write failed in getRequestUsage", {
      error: err,
      cacheKey,
    });
  });

  return usage;
}

/**
 * Record request usage in D1 and keep the KV quota cache warm.
 *
 * @param billingCycleStart - used to derive the quota cache key; optional, falls
 *   back to the current month start so the key matches what getRequestUsage uses.
 */
export async function recordRequestUsage(
  bindings: WorkerBindings,
  apiKey: string,
  userId: string,
  endpoint: string,
  requests: number,
  statusCode: number,
  responseTimeMs: number,
  requestMethod: string,
  billingCycleStart?: string,
  logger?: Logger,
): Promise<void> {
  // Retry the D1 write up to 3 times with exponential backoff.
  // The KV cache update below remains best-effort — D1 is the source of truth.
  await withRetry(() =>
    bindings.DB.prepare(`
      INSERT INTO credit_usage (api_key, user_id, endpoint, credits_used, request_method, status_code, response_time_ms, used_at)
      VALUES (?, ?, ?, ?, ?, ?, ?, ?)
    `)
      .bind(
        apiKey,
        userId,
        endpoint,
        requests,
        requestMethod,
        statusCode,
        responseTimeMs,
        new Date().toISOString(),
      )
      .run(),
  );

  // Optimistically increment the KV cache so the next quota check stays warm.
  // Race conditions here are acceptable: the 60-second TTL bounds any skew, and
  // D1 is always the authoritative source on a cache miss.
  const startDate = billingCycleStart || getMonthStart();
  const cacheKey = `quota:${userId}:${startDate}`;
  const cached = await bindings.KV.get(cacheKey);
  if (cached !== null) {
    bindings.KV.put(cacheKey, (Number(cached) + requests).toString(), {
      expirationTtl: 60,
    }).catch((err) => {
      logger?.warn("KV cache write failed in recordRequestUsage", {
        error: err,
        cacheKey,
      });
    });
  }
}

/**
 * Check request usage and get current status
 */
export async function checkRequestUsage(
  bindings: WorkerBindings,
  userId: string,
  period: "daily" | "monthly",
  limit: number,
  billingCycleStart?: string,
  logger?: Logger,
): Promise<RequestCheckResult> {
  const usage = await getRequestUsage(bindings, userId, period, billingCycleStart, logger);
  const remaining = Math.max(0, limit - usage);
  const resetAt = getResetTime(period, billingCycleStart);

  return {
    usage,
    remaining,
    limit,
    resetAt,
  };
}

/**
 * Get start of today (midnight UTC)
 * Used for daily reset periods
 */
export function getTodayStart(): string {
  const now = new Date();
  now.setUTCHours(0, 0, 0, 0);
  return now.toISOString();
}

/**
 * Get start of current month (1st at midnight UTC)
 * Default billing cycle start for paid users
 */
export function getMonthStart(): string {
  const now = new Date();
  now.setUTCDate(1);
  now.setUTCHours(0, 0, 0, 0);
  return now.toISOString();
}

/**
 * Get when request quota will reset
 */
export function getResetTime(period: "daily" | "monthly", billingCycleStart?: string): string {
  const now = new Date();

  if (period === "daily") {
    // Tomorrow at midnight UTC
    now.setUTCDate(now.getUTCDate() + 1);
    now.setUTCHours(0, 0, 0, 0);
    return now.toISOString();
  }

  // Monthly: next billing cycle
  if (billingCycleStart) {
    // Calculate next billing date based on cycle start
    const cycleStart = new Date(billingCycleStart);
    const dayOfMonth = cycleStart.getUTCDate();

    const nextReset = new Date(now);
    nextReset.setUTCDate(dayOfMonth);
    nextReset.setUTCHours(0, 0, 0, 0);

    // If we're past this month's reset date, go to next month
    if (nextReset <= now) {
      nextReset.setUTCMonth(nextReset.getUTCMonth() + 1);
    }

    return nextReset.toISOString();
  }

  // Default: first of next month
  now.setUTCMonth(now.getUTCMonth() + 1);
  now.setUTCDate(1);
  now.setUTCHours(0, 0, 0, 0);
  return now.toISOString();
}
