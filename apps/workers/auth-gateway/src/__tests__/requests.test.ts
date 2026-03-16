import { beforeEach, describe, expect, it, vi } from "vitest";
import {
  _resetKvStateForTesting,
  checkRequestUsage,
  getRequestUsage,
  recordRequestUsage,
} from "../requests";
import type { WorkerBindings } from "../env";

// ---------------------------------------------------------------------------
// Helpers
// ---------------------------------------------------------------------------

function makeBindings(overrides?: {
  kvGet?: (key: string) => Promise<string | null>;
  kvPut?: (key: string, value: string, opts?: { expirationTtl?: number }) => Promise<void>;
  dbFirst?: () => Promise<{ total: number } | null>;
  dbRun?: () => Promise<void>;
}): WorkerBindings {
  return {
    KV: {
      get: overrides?.kvGet ?? (async () => null),
      put: overrides?.kvPut ?? (async () => {}),
    } as any,
    DB: {
      prepare: () => ({
        bind: () => ({
          first: overrides?.dbFirst ?? (async () => ({ total: 42 })),
          run: overrides?.dbRun ?? (async () => {}),
        }),
      }),
    } as any,
    BACKEND_URL: "http://test",
    BACKEND_SECRET: "secret",
    ENVIRONMENT: "development",
  };
}

// Trigger the circuit to open by simulating N failures (threshold = 3)
async function openCircuit(bindings: WorkerBindings): Promise<void> {
  const failingBindings = makeBindings({
    kvGet: async () => {
      throw new Error("KV unavailable");
    },
    dbFirst: async () => ({ total: 0 }),
  });
  // 3 calls will each hit a KV read failure → opens circuit
  for (let i = 0; i < 3; i++) {
    await getRequestUsage(failingBindings, `circuit-open-user-${i}`, "monthly");
  }
}

// ---------------------------------------------------------------------------
// Tests
// ---------------------------------------------------------------------------

describe("requests", () => {
  beforeEach(() => {
    _resetKvStateForTesting();
    vi.useFakeTimers();
    vi.setSystemTime(new Date("2025-01-15T12:00:00Z"));
  });

  // -------------------------------------------------------------------------
  // getRequestUsage — normal path
  // -------------------------------------------------------------------------
  describe("getRequestUsage", () => {
    it("returns cached value from KV on hit", async () => {
      const bindings = makeBindings({ kvGet: async () => "77" });
      const result = await getRequestUsage(bindings, "user1", "monthly");
      expect(result).toBe(77);
    });

    it("queries D1 on KV miss and returns result", async () => {
      const bindings = makeBindings({
        kvGet: async () => null,
        dbFirst: async () => ({ total: 42 }),
      });
      const result = await getRequestUsage(bindings, "user1", "monthly");
      expect(result).toBe(42);
    });

    it("returns 0 when D1 has no rows", async () => {
      const bindings = makeBindings({
        kvGet: async () => null,
        dbFirst: async () => null,
      });
      const result = await getRequestUsage(bindings, "user1", "monthly");
      expect(result).toBe(0);
    });

    it("writes D1 result back to KV", async () => {
      const puts: Array<{ key: string; value: string }> = [];
      const bindings = makeBindings({
        kvGet: async () => null,
        kvPut: async (key, value) => {
          puts.push({ key, value });
        },
        dbFirst: async () => ({ total: 10 }),
      });
      await getRequestUsage(bindings, "user1", "monthly");
      // KV.put is fire-and-forget; give the microtask queue a turn
      await Promise.resolve();
      expect(puts.some((p) => p.value === "10")).toBe(true);
    });
  });

  // -------------------------------------------------------------------------
  // KV failure handling — circuit breaker
  // -------------------------------------------------------------------------
  describe("KV circuit breaker", () => {
    it("falls back to D1 on a single KV read error", async () => {
      const bindings = makeBindings({
        kvGet: async () => {
          throw new Error("KV timeout");
        },
        dbFirst: async () => ({ total: 55 }),
      });
      const result = await getRequestUsage(bindings, "user1", "monthly");
      expect(result).toBe(55);
    });

    it("opens circuit after threshold failures and stops hitting D1", async () => {
      let d1Calls = 0;
      const failingBindings = makeBindings({
        kvGet: async () => {
          throw new Error("KV unavailable");
        },
        dbFirst: async () => {
          d1Calls++;
          return { total: 5 };
        },
      });

      // 3 calls open the circuit (threshold = 3)
      for (let i = 0; i < 3; i++) {
        await getRequestUsage(failingBindings, "user1", "monthly");
      }
      const d1CallsAfterOpen = d1Calls;

      // After circuit opens, in-memory cache should be warm for "user1"
      // so subsequent calls must NOT hit D1
      await getRequestUsage(failingBindings, "user1", "monthly");
      await getRequestUsage(failingBindings, "user1", "monthly");

      expect(d1Calls).toBe(d1CallsAfterOpen);
    });

    it("uses in-memory cache when circuit is open", async () => {
      await openCircuit(makeBindings({ kvGet: async () => { throw new Error(); }, dbFirst: async () => ({ total: 0 }) }));

      // Seed the in-memory cache by querying through a failing KV binding
      const bindings = makeBindings({
        kvGet: async () => {
          throw new Error("KV unavailable");
        },
        dbFirst: async () => ({ total: 99 }),
      });

      // First call populates memCache from D1
      const first = await getRequestUsage(bindings, "userX", "monthly");
      expect(first).toBe(99);

      // Second call: D1 should NOT be called again
      let d1Called = false;
      const bindings2 = makeBindings({
        kvGet: async () => {
          throw new Error("KV unavailable");
        },
        dbFirst: async () => {
          d1Called = true;
          return { total: 200 };
        },
      });
      const second = await getRequestUsage(bindings2, "userX", "monthly");
      expect(second).toBe(99);
      expect(d1Called).toBe(false);
    });

    it("recovers after the recovery window", async () => {
      // Open the circuit
      await openCircuit(makeBindings({ kvGet: async () => { throw new Error(); }, dbFirst: async () => ({ total: 0 }) }));

      // Advance past the 30-second recovery window
      vi.advanceTimersByTime(31_000);

      // Now KV is healthy again
      let kvCalled = false;
      const bindings = makeBindings({
        kvGet: async () => {
          kvCalled = true;
          return "7";
        },
      });

      const result = await getRequestUsage(bindings, "user1", "monthly");
      expect(kvCalled).toBe(true);
      expect(result).toBe(7);
    });
  });

  // -------------------------------------------------------------------------
  // recordRequestUsage
  // -------------------------------------------------------------------------
  describe("recordRequestUsage", () => {
    it("inserts into D1", async () => {
      let ran = false;
      const bindings = makeBindings({
        dbRun: async () => {
          ran = true;
        },
      });
      await recordRequestUsage(bindings, "key1", "user1", "/v1/text/advice", 1);
      expect(ran).toBe(true);
    });

    it("increments KV cache when entry exists", async () => {
      const puts: Array<{ key: string; value: string }> = [];
      const bindings = makeBindings({
        kvGet: async () => "10",
        kvPut: async (key, value) => {
          puts.push({ key, value });
        },
        dbRun: async () => {},
      });
      await recordRequestUsage(bindings, "key1", "user1", "/v1/text/advice", 3);
      await Promise.resolve();
      expect(puts.some((p) => p.value === "13")).toBe(true);
    });

    it("skips KV when circuit is open", async () => {
      await openCircuit(makeBindings({ kvGet: async () => { throw new Error(); }, dbFirst: async () => ({ total: 0 }) }));

      let kvGetCalled = false;
      const bindings = makeBindings({
        kvGet: async () => {
          kvGetCalled = true;
          return "5";
        },
        dbRun: async () => {},
      });

      await recordRequestUsage(bindings, "key1", "user1", "/v1/text/advice", 1);
      expect(kvGetCalled).toBe(false);
    });

    it("updates in-memory cache when entry exists", async () => {
      // First seed the memCache via getRequestUsage
      const bindings = makeBindings({
        kvGet: async () => null,
        dbFirst: async () => ({ total: 50 }),
        dbRun: async () => {},
      });
      await getRequestUsage(bindings, "user1", "monthly");

      // Now record usage — memCache should be incremented
      const recordBindings = makeBindings({
        kvGet: async () => null,
        dbRun: async () => {},
        dbFirst: async () => ({ total: 999 }), // should NOT be called
      });

      let d1GetCalled = false;
      const spyBindings: WorkerBindings = {
        ...recordBindings,
        DB: {
          prepare: () => ({
            bind: () => ({
              first: async () => {
                d1GetCalled = true;
                return { total: 999 };
              },
              run: async () => {},
            }),
          }),
        } as any,
      };

      await recordRequestUsage(spyBindings, "key1", "user1", "/v1/text/advice", 5);

      // After recording, getRequestUsage should return 55 from memCache (no D1 call)
      const getBindings = makeBindings({
        kvGet: async () => null, // KV miss
        dbFirst: async () => {
          d1GetCalled = true;
          return { total: 999 };
        },
      });
      const usage = await getRequestUsage(getBindings, "user1", "monthly");
      expect(usage).toBe(55);
      expect(d1GetCalled).toBe(false);
    });
  });

  // -------------------------------------------------------------------------
  // checkRequestUsage
  // -------------------------------------------------------------------------
  describe("checkRequestUsage", () => {
    it("returns correct remaining and limit", async () => {
      const bindings = makeBindings({ kvGet: async () => "100" });
      const result = await checkRequestUsage(bindings, "user1", "monthly", 500);
      expect(result.usage).toBe(100);
      expect(result.remaining).toBe(400);
      expect(result.limit).toBe(500);
    });

    it("clamps remaining to 0 when over limit", async () => {
      const bindings = makeBindings({ kvGet: async () => "600" });
      const result = await checkRequestUsage(bindings, "user1", "monthly", 500);
      expect(result.remaining).toBe(0);
    });
  });
});
