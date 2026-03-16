import { Hono } from "hono";
import { sValidator } from "@hono/standard-validator";
import * as z from "zod";
import {
  type ApiKeyData,
  extractKeyPrefix,
  generateApiKey,
  internalError,
  jsonError,
  jsonResponse,
} from "@requiem/workers-shared";
import type { WorkerEnv } from "../../env";
import { planSchema } from "./schemas";

const app = new Hono<WorkerEnv>();

const createApiKeySchema = z.object({
  userId: z.string().min(1),
  plan: planSchema,
  name: z.string().min(1),
  billingCycleStart: z.string().optional(),
});

interface CreateApiKeyResponse {
  apiKey: string; // Full key (only returned once)
  keyPrefix: string;
  userId: string;
  plan: string;
  createdAt: string;
}

/**
 * POST /api-keys
 * Create a new API key
 * Auth: X-API-Management-Key header (only Rails dashboard has this)
 */
app.post(
  "/",
  sValidator("json", createApiKeySchema, (result, _c) => {
    if (!result.success) {
      return jsonError(400, result.error[0]?.message ?? "Validation error");
    }
  }),
  async (c) => {
    const log = c.var.log;
    const body = c.req.valid("json");

    try {
      // Generate new API key
      const fullKey = generateApiKey();
      const keyPrefix = extractKeyPrefix(fullKey);
      const keyName = `key:${fullKey}`;

      // Check if key already exists (extremely unlikely with nanoid, but good practice)
      const existing = await c.env.KV.get(keyName);
      if (existing) {
        log.warn("Generated key already exists (collision)", { keyPrefix });
        return jsonError(409, "Key collision detected, please retry");
      }

      // Prepare key data for KV
      const now = new Date().toISOString();
      const keyData: ApiKeyData = {
        userId: body.userId,
        plan: body.plan,
        createdAt: now,
        billingCycleStart: body.billingCycleStart || now,
      };

      // Write metadata to D1 first — if this fails we haven't touched KV yet, clean failure
      await c.env.DB.prepare(
        `INSERT INTO api_keys (key_prefix, user_id, plan, created_at, billing_cycle_start, active)
         VALUES (?, ?, ?, ?, ?, 1)`,
      )
        .bind(keyPrefix, body.userId, body.plan, now, keyData.billingCycleStart)
        .run();

      // Write auth key to KV
      await c.env.KV.put(keyName, JSON.stringify(keyData));

      // Write reverse-lookup index: prefix:{keyPrefix} → fullKey
      // Used by delete/patch for O(1) lookup instead of KV.list() scan
      await c.env.KV.put(`prefix:${keyPrefix}`, fullKey);

      log.info("API key created successfully", {
        userId: body.userId,
        plan: body.plan,
        keyPrefix,
      });

      const response: CreateApiKeyResponse = {
        apiKey: fullKey, // Return full key (Rails will store hash)
        keyPrefix,
        userId: body.userId,
        plan: body.plan,
        createdAt: now,
      };

      return jsonResponse(response, 201);
    } catch (error) {
      log.error("Failed to create API key", {
        error,
        params: { userId: body.userId, plan: body.plan },
      });

      return internalError(error, "Failed to create API key", c.env.ENVIRONMENT);
    }
  },
);

export default app;
