import { Hono } from "hono";
import * as z from "zod";
import { createLogger, internalError, jsonError, jsonResponse } from "@requiem/workers-shared";
import type { WorkerBindings } from "../../env";

const app = new Hono<{ Bindings: WorkerBindings }>();

const listQuerySchema = z.object({
  userId: z.string().optional(),
  active: z.string().optional(),
});

interface ApiKeyRecord {
  keyPrefix: string;
  userId: string;
  plan: string;
  active: boolean;
  createdAt: string;
  updatedAt: string | null;
  revokedAt: string | null;
  billingCycleStart: string;
}

/**
 * GET /api-keys?userId=...
 * List API keys. Optionally filter by userId.
 * Never returns full key values — only safe metadata.
 * Auth: X-API-Management-Key header (only Rails dashboard has this)
 */
app.get("/", async (c) => {
  const log = createLogger(c.req.raw);
  const parsed = listQuerySchema.safeParse(c.req.query());
  if (!parsed.success) {
    return jsonError(400, parsed.error.issues[0]?.message ?? "Validation error");
  }
  const { userId, active } = parsed.data;
  const activeOnly = active !== "false";

  try {
    let query: string;
    let bindings: (string | number | boolean | null)[];

    if (userId) {
      query = activeOnly
        ? `SELECT key_prefix, user_id, plan, active, created_at, updated_at, revoked_at, billing_cycle_start
             FROM api_keys WHERE user_id = ? AND active = 1 ORDER BY created_at DESC`
        : `SELECT key_prefix, user_id, plan, active, created_at, updated_at, revoked_at, billing_cycle_start
             FROM api_keys WHERE user_id = ? ORDER BY created_at DESC`;
      bindings = [userId];
    } else {
      query = activeOnly
        ? `SELECT key_prefix, user_id, plan, active, created_at, updated_at, revoked_at, billing_cycle_start
             FROM api_keys WHERE active = 1 ORDER BY created_at DESC`
        : `SELECT key_prefix, user_id, plan, active, created_at, updated_at, revoked_at, billing_cycle_start
             FROM api_keys ORDER BY created_at DESC`;
      bindings = [];
    }

    const result = await c.env.DB.prepare(query)
      .bind(...bindings)
      .all<{
        key_prefix: string;
        user_id: string;
        plan: string;
        active: number;
        created_at: string;
        updated_at: string | null;
        revoked_at: string | null;
        billing_cycle_start: string;
      }>();

    const keys: ApiKeyRecord[] = result.results.map((row) => ({
      keyPrefix: row.key_prefix,
      userId: row.user_id,
      plan: row.plan,
      active: row.active === 1,
      createdAt: row.created_at,
      updatedAt: row.updated_at,
      revokedAt: row.revoked_at,
      billingCycleStart: row.billing_cycle_start,
    }));

    log.info("Listed API keys", { count: keys.length, userId, activeOnly });

    return jsonResponse({ keys, total: keys.length });
  } catch (error) {
    log.error("Failed to list API keys", { error, params: { userId } });

    return internalError(error, "Failed to list API keys", c.env.ENVIRONMENT);
  }
});

export default app;
