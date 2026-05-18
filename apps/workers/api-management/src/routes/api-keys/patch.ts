import { Hono } from "hono";
import * as z from "zod";
import {
  type ApiKeyData,
  createLogger,
  internalError,
  jsonError,
  jsonResponse,
} from "@requiem/workers-shared";
import type { WorkerBindings } from "../../env";
import { planSchema } from "./schemas";

const app = new Hono<{ Bindings: WorkerBindings }>();

const patchApiKeySchema = z
  .object({
    plan: planSchema.optional(),
    billingCycleStart: z.string().optional(),
  })
  .refine((d) => d.plan !== undefined || d.billingCycleStart !== undefined, {
    message: "Must provide at least one field to update: plan, billingCycleStart",
  });

/**
 * PATCH /api-keys/:keyPrefix
 * Update an API key (plan change, billing cycle)
 * Auth: X-API-Management-Key header (only Rails dashboard has this)
 */
app.patch("/:keyPrefix", async (c) => {
  const log = createLogger(c.req.raw);
  const keyPrefix = c.req.param("keyPrefix");

  const rawBody = await c.req.json().catch(() => null);
  const parsed = patchApiKeySchema.safeParse(rawBody);
  if (!parsed.success) {
    return jsonError(400, parsed.error.issues[0]?.message ?? "Validation error");
  }
  const body = parsed.data;

  try {
    // O(1) reverse-lookup
    const fullKey = await c.env.KV.get(`prefix:${keyPrefix}`);

    if (!fullKey) {
      log.warn("API key not found for update", { keyPrefix });
      return jsonError(404, "API key not found");
    }

    const keyName = `key:${fullKey}`;

    // Get existing key data
    const existingData = await c.env.KV.get<ApiKeyData>(keyName, "json");
    if (!existingData) {
      log.warn("API key data not found in KV", { keyPrefix });
      return jsonError(404, "API key not found");
    }

    // Update key data
    const updatedData: ApiKeyData = {
      ...existingData,
      ...(body.plan && { plan: body.plan }),
      ...(body.billingCycleStart && { billingCycleStart: body.billingCycleStart }),
    };

    // Write updated data to KV
    await c.env.KV.put(keyName, JSON.stringify(updatedData));

    // Update in D1
    const updates: string[] = [];
    const bindings: (string | number)[] = [];

    if (body.plan) {
      updates.push("plan = ?");
      bindings.push(body.plan);
    }
    if (body.billingCycleStart) {
      updates.push("billing_cycle_start = ?");
      bindings.push(body.billingCycleStart);
    }
    updates.push("updated_at = ?");
    bindings.push(new Date().toISOString());
    bindings.push(keyPrefix);

    await c.env.DB.prepare(
      `UPDATE api_keys
         SET ${updates.join(", ")}
         WHERE key_prefix = ?`,
    )
      .bind(...bindings)
      .run();

    log.info("API key updated successfully", {
      keyPrefix,
      updates: body,
    });

    return jsonResponse({
      success: true,
      message: "API key updated successfully",
      keyPrefix,
      plan: updatedData.plan,
      billingCycleStart: updatedData.billingCycleStart,
    });
  } catch (error) {
    log.error("Failed to update API key", {
      error,
      params: { keyPrefix, updates: body },
    });

    return internalError(error, "Failed to update API key", c.env.ENVIRONMENT);
  }
});

export default app;
