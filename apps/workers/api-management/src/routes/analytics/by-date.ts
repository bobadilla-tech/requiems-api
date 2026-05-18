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
import type { DateStats } from "./types";

const app = new Hono<{ Bindings: WorkerBindings }>();

const byDateQuerySchema = z.object({
  userId: z.string().min(1, "Missing required parameter: userId"),
  since: z
    .string()
    .datetime({ message: "since must be a valid ISO 8601 datetime" })
    .optional()
    .default(() => new Date(Date.now() - THIRTY_DAYS_AGO_MS).toISOString()),
  until: z
    .string()
    .datetime({ message: "until must be a valid ISO 8601 datetime" })
    .optional()
    .default(() => new Date().toISOString()),
  groupBy: z.enum(["day", "hour"]).default("day"),
});

/**
 * GET /analytics/by-date
 * Get usage trends over time for a user
 *
 * Query parameters:
 * - userId: string (required)
 * - since: ISO timestamp (optional, defaults to 30 days ago)
 * - until: ISO timestamp (optional, defaults to now)
 * - groupBy: "day" | "hour" (optional, defaults to "day")
 */
app.get("/by-date", async (c) => {
  const log = createLogger(c.req.raw);

  const parsed = byDateQuerySchema.safeParse(c.req.query());
  if (!parsed.success) {
    return jsonError(400, parsed.error.issues[0]?.message ?? "Validation error");
  }
  const { userId, groupBy, since, until } = parsed.data;

  try {
    const dateFormat = groupBy === "day" ? "%Y-%m-%d" : "%Y-%m-%d %H:00:00";

    const result = await c.env.DB.prepare(`
        SELECT
          strftime('${dateFormat}', used_at) as date,
          COUNT(*) as requests,
          SUM(credits_used) as credits
        FROM credit_usage
        WHERE user_id = ?
          AND used_at >= ?
          AND used_at <= ?
        GROUP BY date
        ORDER BY date ASC
      `)
      .bind(userId, since, until)
      .all<DateStats>();

    log.info("Analytics by date fetched", {
      userId,
      dataPoints: result.results?.length || 0,
      groupBy,
    });

    return jsonResponse({
      timeSeries: result.results,
      dateRange: { since, until },
      groupBy,
    });
  } catch (error) {
    log.error("Error fetching date analytics", {
      error,
      params: { userId, since, until, groupBy },
    });

    return internalError(error, "Failed to fetch analytics", c.env.ENVIRONMENT);
  }
});

export default app;
