import { Hono } from "hono";
import * as z from "zod";
import {
  createLogger,
  internalError,
  jsonError,
  jsonResponse,
  THIRTY_DAYS_AGO_MS,
} from "@requiem/workers-shared";
import type { WorkerBindings } from "../../env";
import type { EndpointStats, UsageSummary } from "./types";

const app = new Hono<{ Bindings: WorkerBindings }>();

const summaryQuerySchema = z.object({
  userId: z.string().min(1, "Missing required parameter: userId"),
  since: z.string().optional(),
  until: z
    .string()
    .optional()
    .default(() => new Date().toISOString()),
});

/**
 * GET /analytics/summary
 * Get overall usage summary for a user
 *
 * Query parameters:
 * - userId: string (required)
 * - since: ISO timestamp (optional, defaults to billing cycle start)
 * - until: ISO timestamp (optional, defaults to now)
 */
app.get("/summary", async (c) => {
  const log = createLogger(c.req.raw);

  const parsed = summaryQuerySchema.safeParse(c.req.query());
  if (!parsed.success) {
    return jsonError(400, parsed.error.issues[0]?.message ?? "Validation error");
  }
  const { userId, since, until } = parsed.data;

  try {
    // If no "since" provided, get the earliest billing cycle start
    let sinceDate = since;
    if (!sinceDate) {
      const billingResult = await c.env.DB.prepare(`
          SELECT MIN(billing_cycle_start) as earliest
          FROM api_keys
          WHERE user_id = ? AND active = 1
        `)
        .bind(userId)
        .first<{ earliest: string }>();

      sinceDate =
        billingResult?.earliest || new Date(Date.now() - THIRTY_DAYS_AGO_MS).toISOString();
    }

    // Run independent queries in parallel
    const [totalsResult, topEndpointsResult] = await Promise.all([
      c.env.DB.prepare(`
          SELECT
            COUNT(*) as totalRequests,
            SUM(credits_used) as totalCredits
          FROM credit_usage
          WHERE user_id = ?
            AND used_at >= ?
            AND used_at <= ?
        `)
        .bind(userId, sinceDate, until)
        .first<{ totalRequests: number; totalCredits: number }>(),

      c.env.DB.prepare(`
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
          LIMIT 5
        `)
        .bind(userId, sinceDate, until)
        .all<EndpointStats>(),
    ]);

    const summary: UsageSummary = {
      userId,
      totalRequests: totalsResult?.totalRequests || 0,
      totalCredits: totalsResult?.totalCredits || 0,
      dateRange: {
        since: sinceDate,
        until,
      },
      topEndpoints: topEndpointsResult.results || [],
    };

    log.info("Analytics summary fetched", { userId });

    return jsonResponse(summary);
  } catch (error) {
    log.error("Error fetching analytics summary", {
      error,
      params: { userId, since, until },
    });

    return internalError(error, "Failed to fetch analytics", c.env.ENVIRONMENT);
  }
});

export default app;
