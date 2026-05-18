import type { PlanConfig, PlanName, RateLimitResult } from "@requiem/workers-shared";
import type { WorkerBindings } from "./env";

const RATE_LIMIT_WINDOW_SECONDS = 60;

/**
 * Get request limits description for a plan
 */
export function getPlanLimits(plan: PlanName): string {
  const limits: Record<PlanName, string> = {
    free: "500 requests/month",
    developer: "100k requests/month",
    business: "1M requests/month",
    professional: "10M requests/month",
    enterprise: "unlimited requests/month",
  };

  return limits[plan];
}

/**
 * Get request limit exceeded message
 * All plans are monthly, so this always returns the monthly message
 */
export function getRequestLimitMessage(): string {
  return "Monthly request limit exceeded. Upgrade at requiems-api.xyz";
}

export async function checkRateLimit(
  bindings: WorkerBindings,
  apiKey: string,
  plan: PlanConfig,
): Promise<RateLimitResult> {
  const now = Date.now();
  const currentMinute = Math.floor(now / 60_000);

  const minuteKey = `rl:m:${apiKey}:${currentMinute}`;
  const existing = (await bindings.KV.get(minuteKey)) ?? "0";

  const minuteCount = Number.parseInt(existing, 10);
  const resetAt = (currentMinute + 1) * 60000;

  if (minuteCount >= plan.ratePerMinute) {
    return {
      allowed: false,
      remaining: 0,
      resetAt,
    };
  }

  await bindings.KV.put(minuteKey, (minuteCount + 1).toString(), {
    expirationTtl: RATE_LIMIT_WINDOW_SECONDS,
  });

  return {
    allowed: true,
    remaining: plan.ratePerMinute - minuteCount - 1,
    resetAt,
  };
}
