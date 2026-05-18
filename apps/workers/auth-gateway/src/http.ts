/**
 * Auth gateway specific HTTP utilities
 * Base utilities (jsonResponse, jsonError, corsResponse) are in @requiem/workers-shared
 */

/**
 * Filter headers before forwarding to backend
 * Removes Cloudflare headers and sensitive data
 */
export function filterHeaders(headers: Headers, backendSecret: string): Headers {
  const filtered = new Headers();

  // Capture real client IP before stripping Cloudflare headers
  const connectingIp = headers.get("CF-Connecting-IP");

  for (const [key, value] of headers.entries()) {
    const lowerKey = key.toLowerCase();

    if (lowerKey.startsWith("cf-")) continue;
    if (lowerKey === "requiems-api-key") continue;
    if (lowerKey === "connection") continue;
    if (lowerKey === "keep-alive") continue;

    filtered.set(key, value);
  }

  if (connectingIp) {
    filtered.set("X-Forwarded-For", connectingIp);
  }

  filtered.set("X-Backend-Secret", backendSecret);

  return filtered;
}

/**
 * Add request usage and rate limit headers to response
 */
export function addUsageHeaders(
  response: Response,
  headers: {
    requestsUsed: number;
    requestsRemaining: number;
    requestsReset: string;
    plan: string;
    rateLimitLimit: number;
    rateLimitRemaining: number;
  },
): Response {
  const newResponse = new Response(response.body, {
    status: response.status,
    statusText: response.statusText,
    headers: response.headers,
  });

  newResponse.headers.delete("X-Usage-Count");
  newResponse.headers.set("Access-Control-Allow-Origin", "*");
  newResponse.headers.set("X-Requests-Used", headers.requestsUsed.toString());
  newResponse.headers.set("X-Requests-Remaining", headers.requestsRemaining.toString());
  newResponse.headers.set("X-Requests-Reset", headers.requestsReset);
  newResponse.headers.set("X-Plan", headers.plan);
  newResponse.headers.set("X-RateLimit-Limit", headers.rateLimitLimit.toString());
  newResponse.headers.set("X-RateLimit-Remaining", headers.rateLimitRemaining.toString());

  return newResponse;
}

export type BackendResult =
  | { ok: true; response: Response }
  | { ok: false; error: string; status: 502 | 504 };

const BACKEND_TIMEOUT_MS = 10_000;

/**
 * Fetch from backend with error handling and a hard timeout.
 * Returns status 504 on timeout, 502 on other network errors.
 */
export async function fetchBackend(
  url: string | URL,
  init: RequestInit,
  timeoutMs = BACKEND_TIMEOUT_MS,
): Promise<BackendResult> {
  const controller = new AbortController();
  const timer = setTimeout(() => controller.abort(), timeoutMs);

  try {
    const response = await fetch(url, { ...init, signal: controller.signal });
    return { ok: true, response };
  } catch (error) {
    if (error instanceof Error && error.name === "AbortError") {
      console.error("Backend timeout:", error);
      return { ok: false, error: "Backend timeout", status: 504 };
    }
    console.error("Backend error:", error);
    return { ok: false, error: "Backend unavailable", status: 502 };
  } finally {
    clearTimeout(timer);
  }
}
