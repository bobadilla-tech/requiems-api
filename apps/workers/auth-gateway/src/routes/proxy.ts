import { Hono } from "hono";
import {
  type ApiKeyData,
  createLogger,
  getRequestMultiplier,
  jsonError,
  type PlanConfig,
  type RateLimitResult,
  type RequestCheckResult,
} from "@requiem/workers-shared";

import { addUsageHeaders, fetchBackend, filterHeaders } from "../http";
import { recordRequestUsage } from "../requests";
import type { WorkerBindings } from "../env";

type Variables = {
  apiKey: string;
  keyData: ApiKeyData;
  plan: PlanConfig;
  rateLimit: RateLimitResult;
  requestUsage: RequestCheckResult;
};

const app = new Hono<{ Bindings: WorkerBindings; Variables: Variables }>();

/**
 * Wildcard route handler - proxies all requests to backend
 * Auth middleware has already validated API key, rate limits, and quotas
 */
app.all("/*", async (c) => {
  const log = createLogger(c.req.raw);
  const url = new URL(c.req.url);

  // Get auth context from middleware
  const apiKey = c.get("apiKey");
  const keyData = c.get("keyData");
  const plan = c.get("plan");
  const rateLimit = c.get("rateLimit");
  const requestUsage = c.get("requestUsage");

  // Calculate request multiplier for this endpoint
  const requestMultiplier = getRequestMultiplier(c.req.method, url.pathname);

  // Construct backend URL
  const backendUrl = new URL(url.pathname + url.search, c.env.BACKEND_URL);

  // Filter headers and add backend secret
  const backendHeaders = filterHeaders(c.req.raw.headers, c.env.BACKEND_SECRET);

  // Buffer the body so it can be retransmitted if the backend redirects.
  // ReadableStream bodies cannot be re-sent after a redirect.
  const hasBody = c.req.method !== "GET" && c.req.method !== "HEAD";
  const body = hasBody ? await c.req.arrayBuffer() : null;

  // Fetch from backend
  const result = await fetchBackend(backendUrl, {
    method: c.req.method,
    headers: backendHeaders,
    body,
  });

  if (!result.ok) {
    log.error("Backend fetch failed", { error: result.error });
    return jsonError(result.status, result.error);
  }

  const backendResponse = result.response;

  if (!backendResponse.ok) {
    log.warn("Backend error response", {
      status: backendResponse.status,
      path: url.pathname,
    });
  }

  // Record usage for every request that reaches the backend, matching rate-limit
  // behavior (which ticks eagerly in auth middleware regardless of backend outcome).
  c.executionCtx.waitUntil(
    recordRequestUsage(
      c.env,
      apiKey,
      keyData.userId,
      url.pathname,
      requestMultiplier,
      keyData.billingCycleStart,
      log,
    ).catch((err) => {
      log.error("Failed to record usage after retries", {
        error: err,
        path: url.pathname,
        userId: keyData.userId,
      });
    }),
  );

  const response = addUsageHeaders(backendResponse, {
    requestsUsed: requestMultiplier,
    requestsRemaining: Math.max(0, requestUsage.remaining - requestMultiplier),
    requestsReset: requestUsage.resetAt,
    plan: keyData.plan,
    rateLimitLimit: plan.ratePerMinute,
    rateLimitRemaining: rateLimit.remaining,
  });

  return response;
});

export default app;
