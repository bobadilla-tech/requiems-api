import { Hono } from "hono";
import { createLogger, internalError, jsonError, jsonResponse } from "@requiem/workers-shared";
import type { WorkerBindings } from "../../env";

const app = new Hono<{ Bindings: WorkerBindings }>();

/**
 * DELETE /api-keys/:keyPrefix
 * Revoke an API key by its prefix
 * Auth: X-API-Management-Key header (only Rails dashboard has this)
 */
app.delete("/:keyPrefix", async (c) => {
  const log = createLogger(c.req.raw);
  const keyPrefix = c.req.param("keyPrefix");

  if (!keyPrefix) {
    return jsonError(400, "Missing keyPrefix parameter");
  }

  try {
    const fullKey = await c.env.KV.get(`prefix:${keyPrefix}`);

    if (!fullKey) {
      log.warn("API key not found for revocation", { keyPrefix });
      return jsonError(404, "API key not found");
    }

    // Update D1 first — audit trail is preserved even if KV cleanup fails
    await c.env.DB.prepare(
      `UPDATE api_keys
       SET revoked_at = ?, active = 0
       WHERE key_prefix = ?`,
    )
      .bind(new Date().toISOString(), keyPrefix)
      .run();

    // Delete auth key and reverse-lookup index from KV
    await c.env.KV.delete(`key:${fullKey}`);
    await c.env.KV.delete(`prefix:${keyPrefix}`);

    log.info("API key revoked successfully", { keyPrefix });

    return jsonResponse({
      success: true,
      message: "API key revoked successfully",
      keyPrefix,
    });
  } catch (error) {
    log.error("Failed to revoke API key", {
      error,
      params: { keyPrefix },
    });

    return internalError(error, "Failed to revoke API key", c.env.ENVIRONMENT);
  }
});

export default app;
