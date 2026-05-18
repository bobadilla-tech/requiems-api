import { afterEach, beforeEach, describe, expect, it, vi } from "vitest";
import worker from "../index";
import { PLANS } from "@requiem/workers-shared";
import type { ApiKeyData } from "@requiem/workers-shared";
import type { WorkerBindings } from "../env";

// Valid key format: requiem_ + 24 alphanumeric chars
const VALID_API_KEY = "requiem_abcdefghijklmnopqrstuvwx";
const TEST_USER_ID = "user-quota-test-123";
const BACKEND_SECRET = "test-backend-secret-32-chars-pad";

/**
 * Build a mock ExecutionContext where waitUntil is a no-op spy.
 */
function makeCtx(): ExecutionContext {
  return {
    waitUntil: vi.fn(),
    passThroughOnException: vi.fn(),
  } as unknown as ExecutionContext;
}

/**
 * Compute the quota cache key that getRequestUsage() would write/read.
 * Mirrors the logic in requests.ts:getMonthStart().
 */
function monthStartKey(userId: string): string {
  const now = new Date();
  now.setUTCDate(1);
  now.setUTCHours(0, 0, 0, 0);
  return `quota:${userId}:${now.toISOString()}`;
}

describe("Quota exceeded — integration", () => {
  let kvStore: Map<string, string>;
  let bindings: WorkerBindings;

  afterEach(() => {
    vi.useRealTimers();
  });

  beforeEach(() => {
    vi.useFakeTimers();
    vi.setSystemTime(new Date("2024-06-15T12:00:00Z"));

    kvStore = new Map();

    // Seed KV with a valid free-plan API key
    const keyData: ApiKeyData = {
      userId: TEST_USER_ID,
      plan: "free",
      createdAt: "2024-01-01T00:00:00Z",
    };
    kvStore.set(`key:${VALID_API_KEY}`, JSON.stringify(keyData));
    // Quota cache key is intentionally absent — forces a D1 look-up

    bindings = {
      KV: {
        get: async (key: string, type?: string) => {
          const value = kvStore.get(key);
          if (value === undefined) return null;
          return type === "json" ? JSON.parse(value) : value;
        },
        put: async (key: string, value: string, _opts?: unknown) => {
          kvStore.set(key, value);
        },
      } as unknown as KVNamespace,

      // D1 always returns usage above the free plan limit (500)
      DB: {
        prepare: (_sql: string) => ({
          bind: (..._args: unknown[]) => ({
            first: async <T>() => ({ total: PLANS.free.requestLimit + 1 }) as T,
            run: async () => ({ success: true, meta: {} }),
          }),
        }),
      } as unknown as D1Database,

      BACKEND_URL: "http://test-backend",
      BACKEND_SECRET,
      ENVIRONMENT: "development" as const,
    };
  });

  it("returns 429 when D1 reports usage above the monthly limit", async () => {
    const req = new Request("http://localhost/v1/text/advice", {
      headers: { "requiems-api-key": VALID_API_KEY },
    });

    const res = await worker.fetch(req, bindings, makeCtx());

    expect(res.status).toBe(429);
  });

  it("sets X-Requests-Remaining and X-Requests-Used to 0", async () => {
    const req = new Request("http://localhost/v1/text/advice", {
      headers: { "requiems-api-key": VALID_API_KEY },
    });

    const res = await worker.fetch(req, bindings, makeCtx());

    expect(res.headers.get("X-Requests-Remaining")).toBe("0");
    expect(res.headers.get("X-Requests-Used")).toBe("0");
  });

  it("sets X-Plan header to the user plan", async () => {
    const req = new Request("http://localhost/v1/text/advice", {
      headers: { "requiems-api-key": VALID_API_KEY },
    });

    const res = await worker.fetch(req, bindings, makeCtx());

    expect(res.headers.get("X-Plan")).toBe("free");
  });

  it("sets X-Requests-Reset to a non-empty value", async () => {
    const req = new Request("http://localhost/v1/text/advice", {
      headers: { "requiems-api-key": VALID_API_KEY },
    });

    const res = await worker.fetch(req, bindings, makeCtx());

    expect(res.headers.get("X-Requests-Reset")).toBeTruthy();
  });

  it("includes a 'request limit exceeded' message in the body", async () => {
    const req = new Request("http://localhost/v1/text/advice", {
      headers: { "requiems-api-key": VALID_API_KEY },
    });

    const res = await worker.fetch(req, bindings, makeCtx());
    const body = (await res.json()) as { error: string };

    expect(body.error).toContain("request limit exceeded");
  });

  it("does not query D1 when quota is already cached in KV", async () => {
    // Warm the cache with an over-limit value
    kvStore.set(monthStartKey(TEST_USER_ID), String(PLANS.free.requestLimit + 1));

    const dbPrepare = vi.fn().mockReturnValue({
      bind: vi.fn().mockReturnValue({
        first: vi.fn().mockResolvedValue({
          total: PLANS.free.requestLimit + 1,
        }),
        run: vi.fn().mockResolvedValue({ success: true, meta: {} }),
      }),
    });
    bindings.DB = { prepare: dbPrepare } as unknown as D1Database;

    const req = new Request("http://localhost/v1/text/advice", {
      headers: { "requiems-api-key": VALID_API_KEY },
    });

    const res = await worker.fetch(req, bindings, makeCtx());

    expect(res.status).toBe(429);
    // D1 must not be called because usage was served from KV cache
    expect(dbPrepare).not.toHaveBeenCalled();
  });
});
