import { Hono } from "hono";
import * as z from "zod";
import {
  ANALYTICS_ENDPOINT_MAX_LIMIT,
  createLogger,
  internalError,
  jsonError,
  jsonResponse,
  THIRTY_DAYS_AGO_MS,
} from "@requiem/workers-shared";
import type { WorkerBindings } from "../../env";
import type { EndpointStats } from "./types";

const app = new Hono<{ Bindings: WorkerBindings }>();

const byEndpointQuerySchema = z.object({
  userId: z.string().min(1, "Missing required parameter: userId"),
  since: z.string().optional(),
  until: z
    .string()
    .optional()
    .default(() => new Date().toISOString()),
  limit: z.coerce.number().min(1).max(ANALYTICS_ENDPOINT_MAX_LIMIT).default(10),
});

/**
 * GET /analytics/by-endpoint
 * Get usage breakdown by endpoint for a user
 *
 * Query parameters:
 * - userId: string (required)
 * - since: ISO timestamp (optional, defaults to billing cycle start)
 * - until: ISO timestamp (optional, defaults to now)
 * - limit: number (optional, max top endpoints to return, default: 10)
 */
app.get("/by-endpoint", async (c) => {
  const log = createLogger(c.req.raw);

  const parsed = byEndpointQuerySchema.safeParse(c.req.query());
  if (!parsed.success) {
    return jsonError(400, parsed.error.issues[0]?.message ?? "Validation error");
  }
  const { userId, limit, since, until } = parsed.data;

  try {
    // If no "since" provided, get the earliest billing cycle start for this user
    let sinceDate = since;
    if (!sinceDate) {
      const billingResult = await c.env.DB.prepare(`
          SELECT MIN(billing_cycle_start) as earliest
          FROM api_keys
          WHERE user_id = ? AND active = 1
        `)
        .bind(userId)
        .first<{ earliest: string }>();

      if (billingResult?.earliest) {
        sinceDate = billingResult.earliest;
      } else {
        sinceDate = new Date(Date.now() - THIRTY_DAYS_AGO_MS).toISOString();
        log.warn("No active billing cycle found for user; falling back to 30-day window", {
          userId,
        });
      }
    }

    // Query usage by endpoint
    const result = await c.env.DB.prepare(`
        SELECT
          endpoint,
          COUNT(*) as requests,
          SUM(credits_used) as credits
        FROM credit_usage
        WHERE user_id = ?
          AND used_at >= ?
          AND used_at <= ?
        GROUP BY endpoint
        ORDER BY credits DESC
        LIMIT ?
      `)
      .bind(userId, sinceDate, until, limit)
      .all<EndpointStats>();

    log.info("Analytics by endpoint fetched", {
      userId,
      endpoints: result.results?.length || 0,
    });

    return jsonResponse({
      endpoints: result.results,
      dateRange: { since: sinceDate, until },
    });
  } catch (error) {
    log.error("Error fetching endpoint analytics", {
      error,
      params: { userId, since, until, limit },
    });

    return internalError(error, "Failed to fetch analytics", c.env.ENVIRONMENT);
  }
});

export default app;
