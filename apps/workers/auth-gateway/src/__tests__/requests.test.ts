import { beforeEach, describe, expect, it, vi } from "vitest";
import { getRequestUsage } from "../requests";
import type { WorkerBindings } from "../env";

describe("getRequestUsage", () => {
  let bindings: WorkerBindings;
  let mockKV: { get: ReturnType<typeof vi.fn>; put: ReturnType<typeof vi.fn> };
  let mockDB: { prepare: ReturnType<typeof vi.fn> };

  beforeEach(() => {
    vi.useFakeTimers();
    vi.setSystemTime(new Date("2024-06-15T12:00:00Z"));

    mockKV = {
      get: vi.fn(),
      put: vi.fn().mockResolvedValue(undefined),
    };

    mockDB = {
      prepare: vi.fn().mockReturnValue({
        bind: vi.fn().mockReturnValue({
          first: vi.fn().mockResolvedValue({ total: 42 }),
        }),
      }),
    };

    bindings = {
      KV: mockKV as any,
      DB: mockDB as any,
      BACKEND_URL: "http://test",
      BACKEND_SECRET: "secret",
      ENVIRONMENT: "development",
    };
  });

  it("returns cached value on KV hit", async () => {
    mockKV.get.mockResolvedValue("10");

    const result = await getRequestUsage(bindings, "user-1", "monthly");

    expect(result).toBe(10);
    expect(mockDB.prepare).not.toHaveBeenCalled();
  });

  it("queries D1 and caches result on KV miss", async () => {
    mockKV.get.mockResolvedValue(null);

    const result = await getRequestUsage(bindings, "user-1", "monthly");

    expect(result).toBe(42);
    expect(mockDB.prepare).toHaveBeenCalledOnce();
    expect(mockKV.put).toHaveBeenCalledOnce();
  });

  it("retries KV once on transient error before returning cached value", async () => {
    mockKV.get
      .mockRejectedValueOnce(new Error("transient KV error"))
      .mockResolvedValue("7");

    const result = await getRequestUsage(bindings, "user-1", "monthly");

    expect(result).toBe(7);
    expect(mockKV.get).toHaveBeenCalledTimes(2);
    expect(mockDB.prepare).not.toHaveBeenCalled();
  });

  it("falls back to D1 when both KV attempts fail", async () => {
    const logger = { warn: vi.fn() } as any;
    mockKV.get.mockRejectedValue(new Error("KV down"));

    const result = await getRequestUsage(bindings, "user-1", "monthly", undefined, logger);

    expect(result).toBe(42);
    expect(mockKV.get).toHaveBeenCalledTimes(2);
    expect(mockDB.prepare).toHaveBeenCalledOnce();
    expect(logger.warn).toHaveBeenCalledWith(
      expect.stringContaining("KV cache read failed"),
      expect.objectContaining({ cacheKey: expect.any(String) }),
    );
  });
});
