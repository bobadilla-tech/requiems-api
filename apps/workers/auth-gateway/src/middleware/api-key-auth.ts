import type { MiddlewareHandler } from "hono";
import {
  type ApiKeyData,
  createLogger,
  isValidKeyFormat,
  jsonError,
  maskApiKey,
  type PlanConfig,
  PLANS,
  type RateLimitResult,
  type RequestCheckResult,
  WEBSITE_URL,
} from "@requiem/workers-shared";

import { checkRateLimit, getRequestLimitMessage } from "../rate-limit";
import { checkRequestUsage } from "../requests";
import type { WorkerBindings } from "../env";

// Define variables that will be attached to context
type Variables = {
  apiKey: string;
  keyData: ApiKeyData;
  plan: PlanConfig;
  rateLimit: RateLimitResult;
  requestUsage: RequestCheckResult;
};

/**
 * API Key Authentication Middleware
 *
 * Validates API key, checks rate limits, and verifies request quotas.
 * Attaches validated key data to context for downstream handlers.
 */
export const apiKeyAuthMiddleware: MiddlewareHandler<{
  Bindings: WorkerBindings;
  Variables: Variables;
}> = async (c, next) => {
  const log = createLogger(c.req.raw);
  const apiKey = c.req.header("requiems-api-key");

  if (!apiKey) {
    return jsonError(401, `Get your key at ${WEBSITE_URL}`);
  }

  if (!isValidKeyFormat(apiKey)) {
    return jsonError(401, "Invalid API key");
  }

  // Fetch API key data from KV
  const keyData = await c.env.KV.get<ApiKeyData>(`key:${apiKey}`, "json");

  if (!keyData) {
    return jsonError(401, "Invalid API key");
  }

  // Validate plan configuration
  const plan = PLANS[keyData.plan as keyof typeof PLANS];

  if (!plan) {
    return jsonError(500, "Invalid plan configuration");
  }

  // Check per-minute rate limit
  const rateLimit = await checkRateLimit(c.env, apiKey, plan);

  if (!rateLimit.allowed) {
    log.info("Rate limit exceeded", {
      key: maskApiKey(apiKey),
      plan: keyData.plan,
    });

    return jsonError(429, "Rate limit exceeded", {
      "X-RateLimit-Limit": plan.ratePerMinute.toString(),
      "X-RateLimit-Remaining": "0",
      "X-RateLimit-Reset": Math.ceil(rateLimit.resetAt / 1000).toString(),
      "Retry-After": Math.ceil((rateLimit.resetAt - Date.now()) / 1000).toString(),
    });
  }

  // Check monthly request quota
  const requestUsage = await checkRequestUsage(
    c.env,
    keyData.userId,
    "monthly",
    plan.requestLimit,
    keyData.billingCycleStart,
    log,
  );

  if (requestUsage.usage >= plan.requestLimit) {
    log.info("Request limit exceeded", {
      key: maskApiKey(apiKey),
      plan: keyData.plan,
      usage: requestUsage.usage,
    });

    return jsonError(429, getRequestLimitMessage(), {
      "X-Requests-Used": "0",
      "X-Requests-Remaining": "0",
      "X-Requests-Reset": requestUsage.resetAt,
      "X-Plan": keyData.plan,
    });
  }

  // Attach auth context for downstream handlers
  c.set("apiKey", apiKey);
  c.set("keyData", keyData);
  c.set("plan", plan);
  c.set("rateLimit", rateLimit);
  c.set("requestUsage", requestUsage);

  await next();
};
