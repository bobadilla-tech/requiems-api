import { jsonResponse } from "../http";

import type { ExecutionContext } from "@cloudflare/workers-types";
import type { Hono, NotFoundHandler } from "hono";

/**
 * Standard 404 handler for all workers.
 */
export const notFoundHandler: NotFoundHandler = (_c) => jsonResponse({ error: "Not found" }, 404);

/**
 * Wraps a Hono app in the Cloudflare Worker `fetch` export, running env
 * validation before dispatching so misconfigured deployments fail fast.
 */
export function createWorkerFetch<TEnv extends object>(
  app: Hono<{ Bindings: TEnv }>,
  validateEnv: (env: TEnv) => void,
) {
  return {
    async fetch(request: Request, env: TEnv, ctx: ExecutionContext): Promise<Response> {
      try {
        validateEnv(env);
      } catch (error) {
        console.error("Environment validation failed:", error);
        return jsonResponse({ error: "Configuration error" }, 500);
      }

      return app.fetch(request, env, ctx);
    },
  };
}
