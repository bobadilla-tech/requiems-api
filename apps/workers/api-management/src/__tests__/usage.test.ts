import { describe, expect, it } from "vitest";
import worker from "../index";
import type { UsageRecord } from "../routes/usage/types";
import { authedRequest, makeBindings, makeCtx } from "./helpers";

const SINCE = "2024-01-01T00:00:00.000Z";

describe("GET /usage/export", () => {
  it("returns 400 when `since` query param is missing", async () => {
    const req = authedRequest("/usage/export");
    const res = await worker.fetch(req, makeBindings(), makeCtx());

    expect(res.status).toBe(400);
  });

  it("returns 400 when limit exceeds 5000", async () => {
    const req = authedRequest(`/usage/export?since=${SINCE}&limit=9999`);
    const res = await worker.fetch(req, makeBindings(), makeCtx());

    expect(res.status).toBe(400);
  });

  it("returns usage records from D1", async () => {
    const record: UsageRecord = {
      id: 1,
      api_key: "requiem_abc123",
      user_id: "u1",
      endpoint: "/v1/text/validate",
      credits_used: 1,
      request_method: "POST",
      status_code: 200,
      response_time_ms: 82,
      used_at: "2024-06-01T10:00:00Z",
    };
    const bindings = makeBindings({ DB: makeDbWithResults([record]) });
    const req = authedRequest(`/usage/export?since=${SINCE}`);
    const res = await worker.fetch(req, bindings, makeCtx());

    expect(res.status).toBe(200);
    const body = (await res.json()) as {
      usage: Omit<UsageRecord, "id">[];
      hasMore: boolean;
      nextCursor?: string;
    };
    expect(body.usage).toHaveLength(1);
    expect(body.usage[0].user_id).toBe("u1");
    expect(body.usage[0].endpoint).toBe("/v1/text/validate");
    expect(body.usage[0].request_method).toBe("POST");
    expect(body.usage[0].status_code).toBe(200);
    expect(body.usage[0].response_time_ms).toBe(82);
    // Internal id must not be exposed
    expect((body.usage[0] as Record<string, unknown>).id).toBeUndefined();
  });

  it("sets hasMore: false and no nextCursor when result count is below limit", async () => {
    const bindings = makeBindings({ DB: makeDbWithResults([]) });
    const req = authedRequest(`/usage/export?since=${SINCE}&limit=100`);
    const res = await worker.fetch(req, bindings, makeCtx());

    const body = (await res.json()) as {
      hasMore: boolean;
      nextCursor?: string;
    };
    expect(body.hasMore).toBe(false);
    expect(body.nextCursor).toBeUndefined();
  });

  it("sets hasMore: true and provides nextCursor when result count equals limit", async () => {
    // Return exactly `limit` records — the route infers there are more pages
    const records: UsageRecord[] = Array.from({ length: 2 }, (_, i) => ({
      id: i + 1,
      api_key: "requiem_key",
      user_id: "u1",
      endpoint: "/v1/text/validate",
      credits_used: 1,
      request_method: "GET",
      status_code: 200,
      response_time_ms: 11,
      used_at: "2024-06-01T10:00:00Z",
    }));
    const bindings = makeBindings({ DB: makeDbWithResults(records) });
    const req = authedRequest(`/usage/export?since=${SINCE}&limit=2`);
    const res = await worker.fetch(req, bindings, makeCtx());

    const body = (await res.json()) as {
      hasMore: boolean;
      nextCursor?: string;
    };
    expect(body.hasMore).toBe(true);
    expect(body.nextCursor).toBe("2");
  });

  it("accepts a cursor parameter for pagination", async () => {
    const req = authedRequest(`/usage/export?since=${SINCE}&cursor=100`);
    const res = await worker.fetch(req, makeBindings(), makeCtx());

    // Any non-error response means the cursor was accepted
    expect(res.status).toBe(200);
  });
});

/** Helper: D1 stub that returns the given rows from `.all()`. */
function makeDbWithResults(results: unknown[]) {
  return {
    prepare: (_sql: string) => ({
      bind: (..._args: unknown[]) => ({
        all: async <T>() => ({
          success: true,
          results: results as T[],
          meta: {},
        }),
        first: async <T>() => null as T,
        run: async () => ({ success: true, meta: {} }),
      }),
    }),
  } as unknown as D1Database;
}
