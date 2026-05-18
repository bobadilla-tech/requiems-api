import { jsonError } from "@requiem/workers-shared";

import type { Context, Next } from "hono";
import type { WorkerBindings } from "../env";

async function timingSafeEqual(a: string, b: string): Promise<boolean> {
  const enc = new TextEncoder();
  const [aHash, bHash] = await Promise.all([
    crypto.subtle.digest("SHA-256", enc.encode(a)),
    crypto.subtle.digest("SHA-256", enc.encode(b)),
  ]);
  return crypto.subtle.timingSafeEqual(aHash, bHash);
}

/**
 * API Management key authentication middleware
 * Validates X-API-Management-Key header for all private endpoints
 * Only the Rails dashboard should have this key
 */
export async function apiKeyAuthMiddleware(c: Context<{ Bindings: WorkerBindings }>, next: Next) {
  const apiKey = c.req.header("X-API-Management-Key");
  const expectedKey = c.env.API_MANAGEMENT_API_KEY;

  if (!apiKey || !(await timingSafeEqual(apiKey, expectedKey))) {
    return jsonError(401, "Unauthorized - Invalid or missing API management key");
  }

  await next();
}
