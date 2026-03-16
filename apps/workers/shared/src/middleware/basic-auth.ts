import type { Context, Next } from "hono";
import { jsonResponse } from "../http";
import { createLogger } from "../logger";

// Generic bindings interface for basic auth
interface BasicAuthBindings {
  SWAGGER_USERNAME?: string;
  SWAGGER_PASSWORD?: string;
}

/**
 * Custom Basic Authentication Middleware
 *
 * NOTE: We implemented this custom basic auth instead of using hono/basic-auth
 * after extensive troubleshooting with the Hono library version.
 *
 * The Hono basicAuth middleware was throwing cryptic errors with no message
 * (Error at basicAuth2 in bundled code) that caused 500 responses instead of
 * proper 401 authentication challenges. After multiple debugging attempts
 * including:
 * - Fixing route patterns (/docs vs /docs/*)
 * - Making verifyUser async
 * - Adding explicit type casting
 * - Adding realm parameter
 * - Improving error handling
 *
 * None of these resolved the issue. This custom implementation gives us full
 * control and works reliably in production.
 */
export const basicAuthMiddleware = async (c: Context, next: Next) => {
  const authHeader = c.req.header("Authorization");

  if (!authHeader || !authHeader.startsWith("Basic ")) {
    return new Response("Unauthorized", {
      status: 401,
      headers: {
        "WWW-Authenticate": 'Basic realm="API Management Documentation"',
      },
    });
  }

  try {
    const base64Credentials = authHeader.substring(6);
    const credentials = atob(base64Credentials);
    const [username, password] = credentials.split(":");

    const env = c.env as BasicAuthBindings;
    const validUsername = env.SWAGGER_USERNAME;
    const validPassword = env.SWAGGER_PASSWORD;

    if (!validUsername || !validPassword) {
      const log = createLogger(c.req.raw);
      log.error("SWAGGER credentials not configured");
      return jsonResponse({ error: "Service unavailable" }, 503);
    }

    if (username === validUsername && password === validPassword) {
      await next();
      return;
    }

    return new Response("Unauthorized", {
      status: 401,
      headers: {
        "WWW-Authenticate": 'Basic realm="API Management Documentation"',
      },
    });
  } catch (error) {
    const log = createLogger(c.req.raw);
    log.error("Basic auth error", { error });
    return new Response("Unauthorized", {
      status: 401,
      headers: {
        "WWW-Authenticate": 'Basic realm="API Management Documentation"',
      },
    });
  }
};
