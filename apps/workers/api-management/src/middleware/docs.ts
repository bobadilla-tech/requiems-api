import type { MiddlewareHandler } from "hono";
import { basicAuthMiddleware } from "@requiem/workers-shared";
import type { WorkerBindings } from "../env";

/**
 * Protects /docs with basic auth in production only.
 * In development the page is open so the swagger UI is directly accessible.
 */
export const docsMiddleware: MiddlewareHandler<{ Bindings: WorkerBindings }> = (c, next) =>
  c.env.ENVIRONMENT === "production" ? basicAuthMiddleware(c, next) : next();
