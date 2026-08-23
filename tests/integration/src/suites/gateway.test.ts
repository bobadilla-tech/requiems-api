/**
 * Integration tests — Gateway / Auth layer
 *
 * These tests validate the direct public Go API and its auth/usage response
 * contract, including missing and invalid key paths.
 */

import { describe, expect, it } from "vitest";
import { getConfig } from "../config.js";

/** Issue a request with no API key */
async function unauthenticated(path: string): Promise<Response> {
  const cfg = getConfig();
  return fetch(new URL(path, cfg.baseUrl).toString(), {
    signal: AbortSignal.timeout(cfg.requestTimeoutMs),
  });
}

/** Issue a request with an obviously invalid key */
async function withBadKey(path: string): Promise<Response> {
  const cfg = getConfig();
  return fetch(new URL(path, cfg.baseUrl).toString(), {
    headers: { "requiems-api-key": "requiem_obviously_invalid_key" },
    signal: AbortSignal.timeout(cfg.requestTimeoutMs),
  });
}

/** Issue an authenticated request and return the raw Response */
async function authenticated(path: string): Promise<Response> {
  const cfg = getConfig();
  return fetch(new URL(path, cfg.baseUrl).toString(), {
    headers: { "requiems-api-key": cfg.apiKey },
    signal: AbortSignal.timeout(cfg.requestTimeoutMs),
  });
}

describe("Direct public API", () => {
  describe("Health check", () => {
    it("GET /healthz returns 200 without an API key", async () => {
      const cfg = getConfig();
      const res = await fetch(new URL("/healthz", cfg.baseUrl).toString(), {
        signal: AbortSignal.timeout(cfg.requestTimeoutMs),
      });
      expect(res.status).toBe(200);
      const body = (await res.json()) as Record<string, unknown>;
      expect(body["status"]).toBe("ok");
    });
  });

  describe("Authentication", () => {
    it("returns 401 when no API key is provided", async () => {
      const res = await unauthenticated("/v1/entertainment/advice");
      expect(res.status).toBe(401);
    });

    it("returns 401 for an invalid API key", async () => {
      const res = await withBadKey("/v1/entertainment/advice");
      expect(res.status).toBe(401);
    });

    it("returns 200 for a valid API key", async () => {
      const res = await authenticated("/v1/entertainment/advice");
      expect(res.status).toBe(200);
    });
  });

  describe("Usage headers on successful requests", () => {
    it("response includes X-Requests-Used header", async () => {
      const res = await authenticated("/v1/entertainment/advice");
      expect(
        res.headers.has("x-requests-used") ||
          res.headers.has("X-Requests-Used"),
      ).toBe(true);
    });

    it("response includes X-Requests-Remaining header", async () => {
      const res = await authenticated("/v1/entertainment/advice");
      expect(
        res.headers.has("x-requests-remaining") ||
          res.headers.has("X-Requests-Remaining"),
      ).toBe(true);
    });

    it("response includes X-RateLimit-Remaining header", async () => {
      const res = await authenticated("/v1/entertainment/advice");
      expect(
        res.headers.has("x-ratelimit-remaining") ||
          res.headers.has("X-RateLimit-Remaining"),
      ).toBe(true);
    });

    it("response includes X-Plan header", async () => {
      const res = await authenticated("/v1/entertainment/advice");
      expect(res.headers.has("x-plan") || res.headers.has("X-Plan")).toBe(true);
    });

    it("CORS header is present", async () => {
      const res = await authenticated("/v1/entertainment/advice");
      expect(res.headers.get("access-control-allow-origin")).toBe("*");
    });
  });

  describe("Not found", () => {
    it("returns 404 for an unknown route", async () => {
      const res = await authenticated("/v1/this-route-does-not-exist-xyz");
      expect(res.status).toBe(404);
    });
  });
});
