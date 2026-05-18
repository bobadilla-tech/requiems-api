/**
 * Simple structured logger using cf-ray as trace ID
 *
 * Cloudflare Workers Logs will capture these as JSON for easy filtering.
 * Use `wrangler tail` to see logs in real-time during development.
 */

export interface Logger {
  info: (msg: string, data?: object) => void;
  warn: (msg: string, data?: object) => void;
  error: (msg: string, data?: object) => void;
  debug: (msg: string, data?: object) => void;
}

interface LogEntry {
  level: string;
  rayId: string;
  msg: string;
  [key: string]: unknown;
}

/**
 * Serialize an error object for logging
 */
function serializeError(error: unknown): object {
  if (error instanceof Error) {
    return {
      name: error.name,
      message: error.message,
      stack: error.stack,
      ...(error.cause ? { cause: serializeError(error.cause) } : {}),
    };
  }

  if (typeof error === "object" && error !== null) {
    return error;
  }

  return { value: String(error) };
}

function formatLog(level: string, rayId: string, msg: string, data?: object): string {
  const processedData = data
    ? Object.entries(data).reduce(
        (acc, [key, value]) => {
          // Serialize error objects properly
          if (key === "error" || value instanceof Error) {
            acc[key] = serializeError(value);
          } else {
            acc[key] = value;
          }
          return acc;
        },
        {} as Record<string, unknown>,
      )
    : {};

  const entry: LogEntry = { level, rayId, msg, ...processedData };
  return JSON.stringify(entry);
}

/**
 * Create a logger instance with the request's trace ID
 *
 * @param request - The incoming request (uses cf-ray header as trace ID)
 */
export function createLogger(request: Request): Logger {
  const rayId = request.headers.get("cf-ray") ?? crypto.randomUUID().slice(0, 16);

  return {
    info: (msg, data) => console.log(formatLog("info", rayId, msg, data)),
    warn: (msg, data) => console.warn(formatLog("warn", rayId, msg, data)),
    error: (msg, data) => console.error(formatLog("error", rayId, msg, data)),
    debug: (msg, data) => console.debug(formatLog("debug", rayId, msg, data)),
  };
}

/**
 * Mask an API key for logging (show first 8 chars only)
 */
export function maskApiKey(key: string): string {
  return key.length > 8 ? `${key.slice(0, 8)}...` : key;
}
