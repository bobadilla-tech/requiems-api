import { Hono } from "hono";

import { validateEnv, type WorkerBindings } from "./env";
import {
  corsMiddleware,
  createWorkerFetch,
  errorHandler,
  jsonResponse,
  notFoundHandler,
} from "@requiem/workers-shared";

import { apiKeyAuthMiddleware } from "./middleware/api-key-auth";
import { openApiSpec } from "./generated/openapi";

import proxyRoute from "./routes/proxy";

const app = new Hono<{ Bindings: WorkerBindings }>();

app.get("/healthz", (_c) => jsonResponse({ status: "ok", service: "auth-gateway" }));

app.get("/openapi.json", (c) => c.json(openApiSpec));

app.use("*", corsMiddleware);

app.use("/*", apiKeyAuthMiddleware);

app.route("/", proxyRoute);

app.notFound(notFoundHandler);
app.onError(errorHandler);

export default createWorkerFetch(app, validateEnv);

export { RateLimiterDO } from "./rate-limiter-do";
