import {
  createLogger,
  internalError,
  jsonError,
  jsonResponse,
  USAGE_EXPORT_MAX_LIMIT,
} from "@requiem/workers-shared";
import { Hono } from "hono";
import * as z from "zod";
import type { WorkerBindings } from "../../env";
import type { UsageExportResponse, UsageRecord } from "./types";

const app = new Hono<{ Bindings: WorkerBindings }>();

const exportQuerySchema = z.object({
  since: z.string().min(1, "Missing required parameter: since"),
  limit: z.coerce.number().min(1).max(USAGE_EXPORT_MAX_LIMIT).default(1000),
  cursor: z.coerce.number().min(0).default(0),
});

/**
 * GET /usage/export
 * Export usage data from D1 for sync to PostgreSQL
 *
 * Query parameters:
 * - since: ISO timestamp - get records after this time (required)
 * - limit: number - max records to return (default: 1000, max: 5000)
 * - cursor: string - pagination cursor (last seen record id)
 *
 * Auth: X-API-Management-Key header (only Rails dashboard has this)
 */
app.get("/export", async (c) => {
  const log = createLogger(c.req.raw);

  const parsed = exportQuerySchema.safeParse(c.req.query());
  if (!parsed.success) {
    return jsonError(400, parsed.error.issues[0]?.message ?? "Validation error");
  }
  const { since, limit, cursor: afterId } = parsed.data;

  try {
    // Fetch usage records using keyset pagination on the autoincrement id.
    // This is stable under concurrent inserts (unlike OFFSET which can skip rows).
    const result = await c.env.DB.prepare(`
        SELECT
          id,
          api_key,
          user_id,
          endpoint,
          credits_used,
          request_method,
          status_code,
          response_time_ms,
          used_at
        FROM credit_usage
        WHERE used_at >= ? AND id > ?
        ORDER BY id ASC
        LIMIT ?
      `)
      .bind(since, afterId, limit)
      .all<UsageRecord>();

    if (!result.success) {
      throw new Error("D1 query failed");
    }
    if (!Array.isArray(result.results)) {
      throw new Error(`Unexpected D1 response shape: results is ${typeof result.results}`);
    }
    const records = result.results;
    const hasMore = records.length === limit;
    const nextCursor = hasMore ? records[records.length - 1].id.toString() : undefined;

    log.info("Usage export successful", {
      returned: records.length,
      hasMore,
    });

    // Strip internal id field before returning to callers
    const usage = records.map(({ id: _id, ...rest }) => rest);

    const response: UsageExportResponse = {
      usage,
      hasMore,
      nextCursor,
    };

    return jsonResponse(response);
  } catch (error) {
    log.error("Error exporting usage data", {
      error,
      params: { since, limit, afterId },
    });

    return internalError(error, "Failed to export usage data", c.env.ENVIRONMENT);
  }
});

export default app;
