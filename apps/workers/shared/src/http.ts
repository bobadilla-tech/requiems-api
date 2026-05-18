/**
 * Shared HTTP utilities for Cloudflare Workers
 */

const CORS_HEADERS = {
  "Access-Control-Allow-Origin": "*",
};

export const corsResponse = new Response(null, {
  headers: {
    ...CORS_HEADERS,
    "Access-Control-Allow-Methods": "GET, POST, PUT, DELETE, OPTIONS",
    "Access-Control-Allow-Headers": "Content-Type, requiems-api-key",
    "Access-Control-Max-Age": "86400",
  },
});

/**
 * JSON response helper
 */
export function jsonResponse(
  data: unknown,
  status = 200,
  headers: Record<string, string> = {},
): Response {
  return new Response(JSON.stringify(data), {
    status,
    headers: {
      "Content-Type": "application/json",
      ...CORS_HEADERS,
      ...headers,
    },
  });
}

/**
 * JSON error response helper
 */
export function jsonError(
  status: number,
  message: string,
  headers: Record<string, string> = {},
): Response {
  return jsonResponse({ error: message }, status, headers);
}

/** One field-level validation failure (matches Go httpx validation_failed shape, with message). */
export type ValidationFieldError = {
  field: string;
  message: string;
};

/**
 * 422 Unprocessable Entity with structured field errors (Zod, etc.).
 */
export function jsonValidationFailed(fields: ValidationFieldError[]): Response {
  return jsonResponse({ error: "validation_failed", fields }, 422);
}

/**
 * Returns a 500 error response. In development, includes the error message.
 */
export function internalError(error: unknown, message: string, environment: string): Response {
  if (environment === "development") {
    return jsonError(500, `${message}: ${error instanceof Error ? error.message : String(error)}`);
  }

  return jsonError(500, message);
}

/**
 * Export CORS headers for reuse in workers
 */
export { CORS_HEADERS };
