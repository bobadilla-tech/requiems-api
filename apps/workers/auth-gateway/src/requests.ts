import { type Logger, type RequestCheckResult, withRetry } from "@requiem/workers-shared";
import type { WorkerBindings } from "./env";

// ---------------------------------------------------------------------------
// KV circuit breaker
//
// Tracks consecutive KV failures. After FAILURE_THRESHOLD failures the circuit
// opens and all KV access is skipped for RECOVERY_MS. After that window the
// circuit enters half-open: one probe is allowed through; on success it closes,
// on failure it stays open for another RECOVERY_MS window.
// ---------------------------------------------------------------------------
const KV_FAILURE_THRESHOLD = 3;
const KV_RECOVERY_MS = 30_000; // 30 seconds

interface KvCircuitState {
  failures: number;
  openedAt: number;
}

const kvCircuit: KvCircuitState = { failures: 0, openedAt: 0 };

/** Returns true when KV is healthy enough to try. Resets state on recovery. */
function shouldUseKv(): boolean {
  if (kvCircuit.failures < KV_FAILURE_THRESHOLD) return true;
  if (Date.now() - kvCircuit.openedAt > KV_RECOVERY_MS) {
    // Half-open: allow one probe
    kvCircuit.failures = 0;
    return true;
  }
  return false;
}

function recordKvFailure(logger?: Logger, ctx?: string): void {
  kvCircuit.failures++;
  if (kvCircuit.failures >= KV_FAILURE_THRESHOLD) {
    kvCircuit.openedAt = Date.now();
    logger?.warn("KV circuit opened — falling back to in-memory cache", { ctx });
  }
}

function recordKvSuccess(): void {
  kvCircuit.failures = 0;
}

// ---------------------------------------------------------------------------
// In-memory fallback cache
//
// Used when the KV circuit is open so D1 is not hit on every request within
// the same Worker isolate. Bounded to MEM_CACHE_MAX_ENTRIES to avoid unbounded
// memory growth; oldest entry is evicted when the limit is reached.
// ---------------------------------------------------------------------------
const MEM_CACHE_TTL_MS = 60_000;
const MEM_CACHE_MAX_ENTRIES = 1_000;

interface MemCacheEntry {
  value: number;
  expiresAt: number;
}

const memCache = new Map<string, MemCacheEntry>();

function memCacheGet(key: string): number | null {
  const entry = memCache.get(key);
  if (!entry) return null;
  if (Date.now() > entry.expiresAt) {
    memCache.delete(key);
    return null;
  }
  return entry.value;
}

function memCacheSet(key: string, value: number): void {
  if (memCache.size >= MEM_CACHE_MAX_ENTRIES && !memCache.has(key)) {
    const firstKey = memCache.keys().next().value;
    if (firstKey !== undefined) memCache.delete(firstKey);
  }
  memCache.set(key, { value, expiresAt: Date.now() + MEM_CACHE_TTL_MS });
}

/** Reset all circuit-breaker and cache state. Only intended for tests. */
export function _resetKvStateForTesting(): void {
  kvCircuit.failures = 0;
  kvCircuit.openedAt = 0;
  memCache.clear();
}

// ---------------------------------------------------------------------------
// Public API
// ---------------------------------------------------------------------------

/**
 * Get current request usage.
 *
 * Checks a 60-second KV cache first to avoid hitting D1 on every request.
 * On cache miss the result is written back to KV so the next request is served
 * from cache.
 *
 * When KV is degraded (circuit open) an in-memory fallback cache is used so
 * D1 is not hammered on every request for the duration of the outage.
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
  if (shouldUseKv()) {
    try {
      const cached = await bindings.KV.get(cacheKey);
      if (cached !== null) {
        recordKvSuccess();
        return Number(cached);
      }
      recordKvSuccess();
    } catch (err) {
      recordKvFailure(logger, "getRequestUsage:read");
      logger?.warn("KV read failed in getRequestUsage, falling back", { error: err, cacheKey });
    }
  }

  // In-memory fallback — avoids D1 when KV circuit is open
  const memCached = memCacheGet(cacheKey);
  if (memCached !== null) {
    return memCached;
  }

  // Cache miss — query D1 and populate caches
  const result = await bindings.DB.prepare(`
    SELECT COALESCE(SUM(credits_used), 0) as total
    FROM credit_usage
    WHERE user_id = ? AND used_at >= ?
  `)
    .bind(userId, startDate)
    .first<{ total: number }>();

  const usage = result?.total || 0;

  // Always keep the in-memory fallback warm
  memCacheSet(cacheKey, usage);

  // Write to KV with 60-second TTL (best-effort, don't block on failure)
  if (shouldUseKv()) {
    bindings.KV.put(cacheKey, usage.toString(), { expirationTtl: 60 }).catch((err) => {
      recordKvFailure(logger, "getRequestUsage:write");
      logger?.warn("KV cache write failed in getRequestUsage", { error: err, cacheKey });
    });
  }

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
  billingCycleStart?: string,
  logger?: Logger,
): Promise<void> {
  // Retry the D1 write up to 3 times with exponential backoff.
  // The KV cache update below remains best-effort — D1 is the source of truth.
  await withRetry(() =>
    bindings.DB.prepare(`
      INSERT INTO credit_usage (api_key, user_id, endpoint, credits_used, used_at)
      VALUES (?, ?, ?, ?, datetime('now'))
    `)
      .bind(apiKey, userId, endpoint, requests)
      .run(),
  );

  const startDate = billingCycleStart || getMonthStart();
  const cacheKey = `quota:${userId}:${startDate}`;

  // Always increment the in-memory fallback so it stays warm during KV outages
  const memCached = memCacheGet(cacheKey);
  if (memCached !== null) {
    memCacheSet(cacheKey, memCached + requests);
  }

  // Optimistically increment the KV cache so the next quota check stays warm.
  // Race conditions here are acceptable: the 60-second TTL bounds any skew, and
  // D1 is always the authoritative source on a cache miss.
  if (shouldUseKv()) {
    try {
      const cached = await bindings.KV.get(cacheKey);
      recordKvSuccess();
      if (cached !== null) {
        bindings.KV.put(cacheKey, (Number(cached) + requests).toString(), {
          expirationTtl: 60,
        }).catch((err) => {
          recordKvFailure(logger, "recordRequestUsage:write");
          logger?.warn("KV cache write failed in recordRequestUsage", { error: err, cacheKey });
        });
      }
    } catch (err) {
      recordKvFailure(logger, "recordRequestUsage:read");
      logger?.warn("KV read failed in recordRequestUsage", { error: err, cacheKey });
    }
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
