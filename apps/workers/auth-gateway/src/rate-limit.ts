import type { PlanConfig, PlanName, RateLimitResult } from "@requiem/workers-shared";
import type { WorkerBindings } from "./env";

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
  const resetAt = (currentMinute + 1) * 60_000;

  // Enterprise has no rate limit — skip the DO call entirely
  if (!Number.isFinite(plan.ratePerMinute)) {
    return { allowed: true, remaining: Number.POSITIVE_INFINITY, resetAt };
  }

  // Each (apiKey, minute) maps to its own DO instance, so state is isolated
  // per window. The DO serializes concurrent requests, making the increment atomic.
  const minuteKey = `rl:m:${apiKey}:${currentMinute}`;
  const id = bindings.RATE_LIMITER.idFromName(minuteKey);
  const stub = bindings.RATE_LIMITER.get(id);

  const response = await stub.fetch("https://rate-limiter/check", {
    method: "POST",
    body: JSON.stringify({ limit: plan.ratePerMinute, resetAt }),
    headers: { "Content-Type": "application/json" },
  });

  return (await response.json()) as RateLimitResult;
}
