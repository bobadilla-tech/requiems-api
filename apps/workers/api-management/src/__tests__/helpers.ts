import { vi } from "vitest";
import type { WorkerBindings } from "../env";

export const API_MANAGEMENT_KEY = "test-api-management-key-32-chars!!";

/**
 * Build a no-op ExecutionContext spy (required by the Worker fetch signature).
 */
export function makeCtx(): ExecutionContext {
  return {
    waitUntil: vi.fn(),
    passThroughOnException: vi.fn(),
  } as unknown as ExecutionContext;
}

/**
 * Build a minimal in-memory KV namespace backed by a Map.
 * The returned `store` reference lets tests inspect and seed KV state.
 */
export function makeKV(store = new Map<string, string>()): KVNamespace {
  return {
    get: async (key: string, type?: string) => {
      const value = store.get(key);
      if (value === undefined) return null;
      return type === "json" ? JSON.parse(value) : value;
    },
    put: async (key: string, value: string) => {
      store.set(key, value);
    },
    delete: async (key: string) => {
      store.delete(key);
    },
    list: async () => ({ keys: [], list_complete: true, cursor: "" }),
    getWithMetadata: async () => ({ value: null, metadata: null }),
  } as unknown as KVNamespace;
}

/**
 * Build a minimal D1 stub. `results` is returned from `.all()`, and
 * `firstResult` from `.first()`.  Override either per-test as needed.
 */
export function makeDB(results: unknown[] = [], firstResult: unknown = null): D1Database {
  return {
    prepare: (_sql: string) => ({
      bind: (..._args: unknown[]) => ({
        all: async <T>() => ({
          success: true,
          results: results as T[],
          meta: {},
        }),
        first: async <T>() => firstResult as T,
        run: async () => ({ success: true, meta: {} }),
      }),
    }),
  } as unknown as D1Database;
}

/**
 * Assemble default WorkerBindings with overrides.
 * Tests that need specific KV / DB behaviour should pass in their own mocks.
 */
export function makeBindings(overrides: Partial<WorkerBindings> = {}): WorkerBindings {
  return {
    API_MANAGEMENT_API_KEY: API_MANAGEMENT_KEY,
    ENVIRONMENT: "development",
    KV: makeKV(),
    DB: makeDB(),
    ...overrides,
  };
}

/**
 * Build a request to the api-management worker, including the auth header.
 */
export function authedRequest(
  path: string,
  init: RequestInit = {},
  key = API_MANAGEMENT_KEY,
): Request {
  const headers = new Headers(init.headers);
  headers.set("X-API-Management-Key", key);
  return new Request(`http://localhost${path}`, { ...init, headers });
}
