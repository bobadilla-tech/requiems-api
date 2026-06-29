// generated/runtime.ts
//
// Shared helpers used by every generated tool file. This file is itself
// generated (deterministically) so it has no hand-maintained drift, but it
// contains no per-operation logic — just the plumbing every tool needs.

export interface RequestOptions {
  method: "GET" | "POST" | "PUT" | "PATCH" | "DELETE";
  path: string;
  /** Values to interpolate into `{param}` placeholders in `path`. */
  pathParams?: Record<string, unknown>;
  /** Values to serialize as a query string. */
  query?: Record<string, unknown>;
  /** JSON body for POST/PUT/PATCH requests. */
  body?: unknown;
  /** Per-request timeout override, in ms. Defaults to 5000. */
  timeoutMs?: number;
}

export class RequiemsApiError extends Error {
  constructor(
    public readonly status: number,
    public readonly statusText: string,
    public readonly body: unknown,
  ) {
    super(`Requiems API error: ${status} ${statusText}`);
    this.name = "RequiemsApiError";
  }
}

function baseUrl(): string {
  const url = process.env.REQUIEMS_BASE_URL;
  if (!url) {
    throw new Error(
      "REQUIEMS_BASE_URL is not set. Copy .env.example to .env and fill it in.",
    );
  }
  return url.replace(/\/+$/, "");
}

function buildPath(path: string, pathParams?: Record<string, unknown>): string {
  if (!pathParams) return path;
  let resolved = path;
  for (const [key, value] of Object.entries(pathParams)) {
    resolved = resolved.replace(
      `{${key}}`,
      encodeURIComponent(String(value)),
    );
  }
  return resolved;
}

function buildQuery(query?: Record<string, unknown>): string {
  if (!query) return "";
  const params = new URLSearchParams();
  for (const [key, value] of Object.entries(query)) {
    if (value === undefined || value === null) continue;
    params.set(key, String(value));
  }
  const qs = params.toString();
  return qs ? `?${qs}` : "";
}

/**
 * Perform a single Requiems API request. Shared by every generated tool's
 * handler. Applies auth, timeout, and consistent error handling.
 */
export async function requiemsRequest<T = unknown>(
  options: RequestOptions,
): Promise<T> {
  const url =
    baseUrl() +
    buildPath(options.path, options.pathParams) +
    buildQuery(options.query);

  const controller = new AbortController();
  const timeout = setTimeout(
    () => controller.abort(),
    options.timeoutMs ?? 5000,
  );

  try {
    const res = await fetch(url, {
      method: options.method,
      headers: {
        "requiems-api-key": process.env.REQUIEMS_API_KEY ?? "",
        ...(options.body !== undefined
          ? { "Content-Type": "application/json" }
          : {}),
      },
      body: options.body !== undefined ? JSON.stringify(options.body) : undefined,
      signal: controller.signal,
    });

    if (!res.ok) {
      let parsedBody: unknown;
      try {
        parsedBody = await res.json();
      } catch {
        parsedBody = await res.text().catch(() => undefined);
      }
      throw new RequiemsApiError(res.status, res.statusText, parsedBody);
    }

    const contentType = res.headers.get("content-type") ?? "";
    if (contentType.includes("application/json")) {
      return (await res.json()) as T;
    }

    // Non-JSON responses (e.g. raw PNG from barcode/QR endpoints) are
    // returned as a base64 string so they can flow through MCP as text.
    const buffer = await res.arrayBuffer();
    return Buffer.from(buffer).toString("base64") as unknown as T;
  } finally {
    clearTimeout(timeout);
  }
}
