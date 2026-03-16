interface CheckRequest {
  limit: number;
  resetAt: number;
}

/**
 * RateLimiterDO - Durable Object for atomic per-minute rate limiting.
 *
 * Each instance manages a single (apiKey, minute) window. Because DO
 * instances are single-threaded, read-increment-write is fully serialized —
 * no race condition is possible.
 *
 * Key format used by the caller: `rl:m:{apiKey}:{minuteEpoch}`
 */
export class RateLimiterDO implements DurableObject {
  constructor(private readonly state: DurableObjectState) {}

  async fetch(request: Request): Promise<Response> {
    const { limit, resetAt } = (await request.json()) as CheckRequest;

    const count = (await this.state.storage.get<number>("count")) ?? 0;

    if (count >= limit) {
      return Response.json({ allowed: false, remaining: 0, resetAt });
    }

    await this.state.storage.put("count", count + 1);

    // Schedule cleanup so storage doesn't outlive the window
    if ((await this.state.storage.getAlarm()) === null) {
      await this.state.storage.setAlarm(resetAt + 1000);
    }

    return Response.json({ allowed: true, remaining: limit - count - 1, resetAt });
  }

  async alarm(): Promise<void> {
    await this.state.storage.deleteAll();
  }
}
