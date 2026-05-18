import { captureException, wrapRequestHandler } from "@sentry/cloudflare";
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

// Capture Hono-handled errors (these are swallowed by onError and never rethrow)
app.onError((err, c) => {
  captureException(err);
  return errorHandler(err, c);
});

const baseHandler = createWorkerFetch(app, validateEnv);

export default {
  fetch(request: Request, env: WorkerBindings, ctx: ExecutionContext): Promise<Response> {
    return wrapRequestHandler(
      {
        options: {
          dsn: env.SENTRY_DSN ?? "",
          tracesSampleRate: 0.01,
        },
        request,
        context: ctx,
      },
      () => baseHandler.fetch(request, env, ctx),
    );
  },
} satisfies ExportedHandler<WorkerBindings>;
