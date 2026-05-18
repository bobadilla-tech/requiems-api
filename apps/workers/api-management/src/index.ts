import { captureException, wrapRequestHandler } from "@sentry/cloudflare";
import { createWorkerFetch, errorHandler, notFoundHandler } from "@requiem/workers-shared";
import { Hono } from "hono";

import { validateEnv, type WorkerBindings } from "./env";
import { apiKeyAuthMiddleware, docsMiddleware } from "./middleware/";
import { analyticsRoute, apiKeysRoute, healthzRoute, swaggerRoute, usageRoute } from "./routes";

const app = new Hono<{ Bindings: WorkerBindings }>();

app.route("/", healthzRoute);

app.use("/docs", docsMiddleware);

app.route("/", swaggerRoute);

app.use("/api-keys/*", apiKeyAuthMiddleware);
app.use("/usage/*", apiKeyAuthMiddleware);
app.use("/analytics/*", apiKeyAuthMiddleware);

app.route("/api-keys", apiKeysRoute);
app.route("/usage", usageRoute);
app.route("/analytics", analyticsRoute);

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
