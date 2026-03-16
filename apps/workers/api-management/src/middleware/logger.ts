import type { MiddlewareHandler } from "hono";
import { createLogger } from "@requiem/workers-shared";
import type { WorkerEnv } from "../env";

/**
 * Logging middleware — creates a structured logger from the incoming request
 * and attaches it to c.var.log so all route handlers can use it without boilerplate.
 */
export const loggerMiddleware: MiddlewareHandler<WorkerEnv> = async (c, next) => {
  c.set("log", createLogger(c.req.raw));
  await next();
};
