import { PLANS } from "@requiem/workers-shared";
import { afterEach, beforeEach, describe, expect, it, vi } from "vitest";
import type { WorkerBindings } from "../env";
import { checkRateLimit, getPlanLimits, getRequestLimitMessage } from "../rate-limit";

describe("Rate Limiting", () => {
  // Mock KV store
  let mockKV: Map<string, { value: string; expirationTtl?: number }>;
  let bindings: WorkerBindings;

  afterEach(() => {
    vi.useRealTimers();
  });

  beforeEach(() => {
    // Reset mock KV store before each test
    mockKV = new Map();

    bindings = {
      KV: {
        get: async (key: string) => {
          const entry = mockKV.get(key);
          return entry ? entry.value : null;
        },
        put: async (key: string, value: string, options?: { expirationTtl?: number }) => {
          mockKV.set(key, { value, expirationTtl: options?.expirationTtl });
        },
      } as unknown as KVNamespace,
      DB: {} as unknown as D1Database,
      BACKEND_URL: "http://test",
      BACKEND_SECRET: "test-secret",
      ENVIRONMENT: "development",
    };

    // Mock Date.now to control time in tests
    vi.useFakeTimers();
  });

  describe("checkRateLimit", () => {
    it("allows request when under rate limit", async () => {
      vi.setSystemTime(new Date("2024-01-01T00:00:00Z"));

      const result = await checkRateLimit(bindings, "test-key", PLANS.developer);

      expect(result.allowed).toBe(true);
      expect(result.remaining).toBe(PLANS.developer.ratePerMinute - 1);
      expect(result.resetAt).toBeGreaterThan(Date.now());
    });

    it("denies request when rate limit exceeded", async () => {
      vi.setSystemTime(new Date("2024-01-01T00:00:00Z"));

      // Simulate rate limit already hit
      const currentMinute = Math.floor(Date.now() / 60_000);
      const minuteKey = `rl:m:test-key:${currentMinute}`;
      await bindings.KV.put(minuteKey, "30", { expirationTtl: 60 });

      const result = await checkRateLimit(bindings, "test-key", PLANS.free);

      expect(result.allowed).toBe(false);
      expect(result.remaining).toBe(0);
    });

    it("increments counter for each request", async () => {
      vi.setSystemTime(new Date("2024-01-01T00:00:00Z"));

      const currentMinute = Math.floor(Date.now() / 60_000);
      const minuteKey = `rl:m:test-key:${currentMinute}`;

      // First request
      await checkRateLimit(bindings, "test-key", PLANS.free);
      expect(mockKV.get(minuteKey)?.value).toBe("1");

      // Second request
      await checkRateLimit(bindings, "test-key", PLANS.free);
      expect(mockKV.get(minuteKey)?.value).toBe("2");

      // Third request
      await checkRateLimit(bindings, "test-key", PLANS.free);
      expect(mockKV.get(minuteKey)?.value).toBe("3");
    });

    it("sets correct TTL on rate limit key", async () => {
      vi.setSystemTime(new Date("2024-01-01T00:00:00Z"));

      await checkRateLimit(bindings, "test-key", PLANS.developer);

      const currentMinute = Math.floor(Date.now() / 60_000);
      const minuteKey = `rl:m:test-key:${currentMinute}`;

      const entry = mockKV.get(minuteKey);
      expect(entry?.expirationTtl).toBe(60); // 60 seconds TTL
    });

    it("resets counter after minute boundary", async () => {
      // Start at 00:00:00
      vi.setSystemTime(new Date("2024-01-01T00:00:00Z"));

      const firstMinute = Math.floor(Date.now() / 60_000);
      await checkRateLimit(bindings, "test-key", PLANS.free);

      // Fast forward to next minute (00:01:00)
      vi.setSystemTime(new Date("2024-01-01T00:01:00Z"));

      const secondMinute = Math.floor(Date.now() / 60_000);
      expect(secondMinute).not.toBe(firstMinute);

      // Counter should be reset (new minute = new key)
      const result = await checkRateLimit(bindings, "test-key", PLANS.free);
      expect(result.remaining).toBe(PLANS.free.ratePerMinute - 1);
    });

    it("uses different keys for different API keys", async () => {
      vi.setSystemTime(new Date("2024-01-01T00:00:00Z"));

      await checkRateLimit(bindings, "key-1", PLANS.free);
      await checkRateLimit(bindings, "key-2", PLANS.free);

      const currentMinute = Math.floor(Date.now() / 60_000);
      const key1Counter = mockKV.get(`rl:m:key-1:${currentMinute}`)?.value;
      const key2Counter = mockKV.get(`rl:m:key-2:${currentMinute}`)?.value;

      expect(key1Counter).toBe("1");
      expect(key2Counter).toBe("1");
    });

    it("calculates correct reset time", async () => {
      vi.setSystemTime(new Date("2024-01-01T00:00:30Z")); // 30 seconds into minute

      const result = await checkRateLimit(bindings, "test-key", PLANS.free);

      // Reset should be at end of current minute
      const currentMinute = Math.floor(Date.now() / 60_000);
      const expectedResetAt = (currentMinute + 1) * 60000;

      expect(result.resetAt).toBe(expectedResetAt);
    });

    it("handles enterprise plan with infinite rate limit", async () => {
      vi.setSystemTime(new Date("2024-01-01T00:00:00Z"));

      // Even with high existing count, enterprise should pass
      const currentMinute = Math.floor(Date.now() / 60_000);
      const minuteKey = `rl:m:enterprise-key:${currentMinute}`;
      await bindings.KV.put(minuteKey, "999999", { expirationTtl: 60 });

      const result = await checkRateLimit(bindings, "enterprise-key", PLANS.enterprise);

      expect(result.allowed).toBe(true);
    });

    it("handles free plan rate limit correctly", async () => {
      vi.setSystemTime(new Date("2024-01-01T00:00:00Z"));

      const freePlan = PLANS.free;
      const results = [];

      // Make requests up to the limit
      for (let i = 0; i < freePlan.ratePerMinute; i++) {
        const result = await checkRateLimit(bindings, "free-key", freePlan);
        results.push(result);
      }

      // All requests up to limit should be allowed
      expect(results.every((r) => r.allowed)).toBe(true);

      // Next request should be denied
      const exceededResult = await checkRateLimit(bindings, "free-key", freePlan);
      expect(exceededResult.allowed).toBe(false);
      expect(exceededResult.remaining).toBe(0);
    });

    it("handles missing KV entry (first request)", async () => {
      vi.setSystemTime(new Date("2024-01-01T00:00:00Z"));

      // KV returns null for non-existent key
      const result = await checkRateLimit(bindings, "new-key", PLANS.developer);

      expect(result.allowed).toBe(true);
      expect(result.remaining).toBe(PLANS.developer.ratePerMinute - 1);
    });

    it("correctly decrements remaining count", async () => {
      vi.setSystemTime(new Date("2024-01-01T00:00:00Z"));

      const plan = PLANS.developer;

      // First request
      const result1 = await checkRateLimit(bindings, "test-key", plan);
      expect(result1.remaining).toBe(plan.ratePerMinute - 1);

      // Second request
      const result2 = await checkRateLimit(bindings, "test-key", plan);
      expect(result2.remaining).toBe(plan.ratePerMinute - 2);

      // Third request
      const result3 = await checkRateLimit(bindings, "test-key", plan);
      expect(result3.remaining).toBe(plan.ratePerMinute - 3);
    });
  });

  describe("getPlanLimits", () => {
    it("returns correct description for free plan", () => {
      expect(getPlanLimits("free")).toBe("500 requests/month");
    });

    it("returns correct description for developer plan", () => {
      expect(getPlanLimits("developer")).toBe("100k requests/month");
    });

    it("returns correct description for business plan", () => {
      expect(getPlanLimits("business")).toBe("1M requests/month");
    });

    it("returns correct description for professional plan", () => {
      expect(getPlanLimits("professional")).toBe("10M requests/month");
    });

    it("returns correct description for enterprise plan", () => {
      expect(getPlanLimits("enterprise")).toBe("unlimited requests/month");
    });

    it("descriptions use abbreviated formats for large numbers", () => {
      expect(getPlanLimits("developer")).toContain("k"); // 100k
      expect(getPlanLimits("business")).toContain("M"); // 1M
      expect(getPlanLimits("professional")).toContain("M"); // 10M
    });
  });

  describe("getRequestLimitMessage", () => {
    it("returns monthly limit exceeded message", () => {
      const message = getRequestLimitMessage();

      expect(message).toContain("Monthly");
      expect(message).toContain("request limit exceeded");
    });

    it("includes upgrade URL", () => {
      const message = getRequestLimitMessage();

      expect(message).toContain("requiems-api.xyz");
    });

    it("always returns the same message", () => {
      // Since all plans are monthly, message should be consistent
      const message1 = getRequestLimitMessage();
      const message2 = getRequestLimitMessage();

      expect(message1).toBe(message2);
    });
  });
});
