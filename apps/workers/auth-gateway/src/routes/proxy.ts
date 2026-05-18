import {
  type ApiKeyData,
  createLogger,
  getRequestMultiplier,
  jsonError,
  type PlanConfig,
  type RateLimitResult,
  type RequestCheckResult,
} from "@requiem/workers-shared";
import { Hono } from "hono";

import type { WorkerBindings } from "../env";
import { addUsageHeaders, fetchBackend, filterHeaders } from "../http";
import { recordRequestUsage } from "../requests";

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
  const startedAt = Date.now();

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

  const responseTimeMs = Date.now() - startedAt;

  if (!result.ok) {
    log.error("Backend fetch failed", { error: result.error });
    return jsonError(result.status, result.error);
  }

  const backendResponse = result.response;

  // If backend returned an error, record a zero-credit telemetry entry (for observability)
  if (!backendResponse.ok) {
    log.warn("Backend error response", {
      status: backendResponse.status,
      path: url.pathname,
    });

    log.debug("Backend Response", result);

    const response = addUsageHeaders(backendResponse, {
      requestsUsed: 0,
      requestsRemaining: requestUsage.remaining,
      requestsReset: requestUsage.resetAt,
      plan: keyData.plan,
      rateLimitLimit: plan.ratePerMinute,
      rateLimitRemaining: rateLimit.remaining,
    });

    c.executionCtx.waitUntil(
      recordRequestUsage(
        c.env,
        apiKey,
        keyData.userId,
        url.pathname,
        0,
        backendResponse.status,
        responseTimeMs,
        c.req.method,
        keyData.billingCycleStart,
        log,
      ).catch((err) => {
        log.error("Failed to record backend error usage after retries", {
          error: err,
          path: url.pathname,
          userId: keyData.userId,
        });
      }),
    );

    return response;
  }

  // Use dynamic multiplier from backend if present (batch endpoints set X-Usage-Count),
  // otherwise fall back to the static per-endpoint multiplier.
  const usageCountHeader = backendResponse.headers.get("X-Usage-Count");
  const parsedCount = usageCountHeader ? parseInt(usageCountHeader, 10) : NaN;
  const effectiveMultiplier =
    !Number.isNaN(parsedCount) && parsedCount > 0 ? parsedCount : requestMultiplier;

  // Record usage after response is sent — waitUntil keeps the worker alive for the write.
  // recordRequestUsage retries up to 3 times internally; log if all attempts fail.
  c.executionCtx.waitUntil(
    recordRequestUsage(
      c.env,
      apiKey,
      keyData.userId,
      url.pathname,
      effectiveMultiplier,
      backendResponse.status,
      responseTimeMs,
      c.req.method,
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

  // Add usage headers to successful response
  const response = addUsageHeaders(backendResponse, {
    requestsUsed: effectiveMultiplier,
    requestsRemaining: Math.max(0, requestUsage.remaining - effectiveMultiplier),
    requestsReset: requestUsage.resetAt,
    plan: keyData.plan,
    rateLimitLimit: plan.ratePerMinute,
    rateLimitRemaining: rateLimit.remaining,
  });

  return response;
});

export default app;
