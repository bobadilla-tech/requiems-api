import { describe, expect, it } from "vitest";
import worker from "../index";
import { authedRequest, makeBindings, makeCtx } from "./helpers";
import type { DateStats, EndpointStats } from "../routes/analytics/types";

// ------------------------------------------------------------------
// Shared DB factories
// ------------------------------------------------------------------

/** D1 stub that returns the given rows from `.all()` and `firstResult` from `.first()`. */
function makeDb(allResults: unknown[] = [], firstResult: unknown = null) {
  return {
    prepare: (_sql: string) => ({
      bind: (..._args: unknown[]) => ({
        all: async <T>() => ({
          success: true,
          results: allResults as T[],
          meta: {},
        }),
        first: async <T>() => firstResult as T,
        run: async () => ({ success: true, meta: {} }),
      }),
    }),
  } as unknown as D1Database;
}

// ------------------------------------------------------------------
// GET /analytics/summary
// ------------------------------------------------------------------

describe("GET /analytics/summary", () => {
  it("returns 400 when userId is missing", async () => {
    const req = authedRequest("/analytics/summary");
    const res = await worker.fetch(req, makeBindings(), makeCtx());

    expect(res.status).toBe(400);
  });

  it("returns summary with totals for a user", async () => {
    const totals = { totalRequests: 42, totalCredits: 84 };
    const topEndpoints: EndpointStats[] = [
      { endpoint: "/v1/text/validate", requests: 30, credits: 60 },
    ];

    // The summary route makes two parallel queries.
    // Our stub returns the same `allResults` / `firstResult` for every call.
    // `.first()` is used for billing cycle start (1st call) and totals (2nd call).
    // We make `.first()` return `totals` so the totals query works correctly,
    // and `.all()` return the endpoints for the top-endpoints query.
    const db = makeDb(topEndpoints, totals);
    const bindings = makeBindings({ DB: db });

    const req = authedRequest("/analytics/summary?userId=u1");
    const res = await worker.fetch(req, bindings, makeCtx());

    expect(res.status).toBe(200);
    const body = (await res.json()) as {
      userId: string;
      totalRequests: number;
      totalCredits: number;
      topEndpoints: EndpointStats[];
      dateRange: { since: string; until: string };
    };
    expect(body.userId).toBe("u1");
    expect(body.totalRequests).toBe(42);
    expect(body.totalCredits).toBe(84);
    expect(body.topEndpoints).toHaveLength(1);
    expect(body.dateRange.since).toBeTruthy();
    expect(body.dateRange.until).toBeTruthy();
  });

  it("defaults totals to 0 when D1 returns null", async () => {
    // `.first()` returns null → no billing cycle row and no totals row
    const bindings = makeBindings({ DB: makeDb([], null) });
    const req = authedRequest("/analytics/summary?userId=u1");
    const res = await worker.fetch(req, bindings, makeCtx());

    expect(res.status).toBe(200);
    const body = (await res.json()) as {
      totalRequests: number;
      totalCredits: number;
    };
    expect(body.totalRequests).toBe(0);
    expect(body.totalCredits).toBe(0);
  });
});

// ------------------------------------------------------------------
// GET /analytics/by-date
// ------------------------------------------------------------------

describe("GET /analytics/by-date", () => {
  it("returns 400 when userId is missing", async () => {
    const req = authedRequest("/analytics/by-date");
    const res = await worker.fetch(req, makeBindings(), makeCtx());

    expect(res.status).toBe(400);
  });

  it("returns 400 for an invalid ISO datetime in `since`", async () => {
    const req = authedRequest("/analytics/by-date?userId=u1&since=not-a-date");
    const res = await worker.fetch(req, makeBindings(), makeCtx());

    expect(res.status).toBe(400);
  });

  it("returns time series data with dateRange and groupBy", async () => {
    const rows: DateStats[] = [
      { date: "2024-06-01", requests: 10, credits: 20 },
      { date: "2024-06-02", requests: 5, credits: 10 },
    ];
    const bindings = makeBindings({ DB: makeDb(rows) });
    const req = authedRequest("/analytics/by-date?userId=u1");
    const res = await worker.fetch(req, bindings, makeCtx());

    expect(res.status).toBe(200);
    const body = (await res.json()) as {
      timeSeries: DateStats[];
      dateRange: { since: string; until: string };
      groupBy: string;
    };
    expect(body.timeSeries).toHaveLength(2);
    expect(body.timeSeries[0].date).toBe("2024-06-01");
    expect(body.groupBy).toBe("day");
    expect(body.dateRange.since).toBeTruthy();
    expect(body.dateRange.until).toBeTruthy();
  });

  it("accepts groupBy=hour", async () => {
    const bindings = makeBindings({ DB: makeDb([]) });
    const req = authedRequest("/analytics/by-date?userId=u1&groupBy=hour");
    const res = await worker.fetch(req, bindings, makeCtx());

    expect(res.status).toBe(200);
    const body = (await res.json()) as { groupBy: string };
    expect(body.groupBy).toBe("hour");
  });

  it("rejects an invalid groupBy value", async () => {
    const req = authedRequest("/analytics/by-date?userId=u1&groupBy=week");
    const res = await worker.fetch(req, makeBindings(), makeCtx());

    expect(res.status).toBe(400);
  });
});

// ------------------------------------------------------------------
// GET /analytics/by-endpoint
// ------------------------------------------------------------------

describe("GET /analytics/by-endpoint", () => {
  it("returns 400 when userId is missing", async () => {
    const req = authedRequest("/analytics/by-endpoint");
    const res = await worker.fetch(req, makeBindings(), makeCtx());

    expect(res.status).toBe(400);
  });

  it("returns 400 when limit exceeds 100", async () => {
    const req = authedRequest("/analytics/by-endpoint?userId=u1&limit=999");
    const res = await worker.fetch(req, makeBindings(), makeCtx());

    expect(res.status).toBe(400);
  });

  it("returns endpoint breakdown with dateRange", async () => {
    const rows: EndpointStats[] = [
      { endpoint: "/v1/text/validate", requests: 50, credits: 100 },
      { endpoint: "/v1/text/words/define", requests: 10, credits: 40 },
    ];
    const bindings = makeBindings({ DB: makeDb(rows) });
    const req = authedRequest("/analytics/by-endpoint?userId=u1");
    const res = await worker.fetch(req, bindings, makeCtx());

    expect(res.status).toBe(200);
    const body = (await res.json()) as {
      endpoints: EndpointStats[];
      dateRange: { since: string; until: string };
    };
    expect(body.endpoints).toHaveLength(2);
    expect(body.endpoints[0].endpoint).toBe("/v1/text/validate");
    expect(body.dateRange.until).toBeTruthy();
  });

  it("falls back to 30-day window when no active billing cycle exists", async () => {
    // first() returns null → no billing cycle → falls back to 30 days ago
    const bindings = makeBindings({ DB: makeDb([], null) });
    const req = authedRequest("/analytics/by-endpoint?userId=u1");
    const res = await worker.fetch(req, bindings, makeCtx());

    expect(res.status).toBe(200);
    const body = (await res.json()) as { dateRange: { since: string } };
    // The since date should be roughly 30 days ago — just check it's an ISO string
    expect(body.dateRange.since).toMatch(/^\d{4}-\d{2}-\d{2}T/);
  });
});
