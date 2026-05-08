import { afterEach, describe, expect, it, vi } from "vitest";
import { corsResponse, jsonError, jsonResponse } from "@requiem/workers-shared";
import { addUsageHeaders, fetchBackend, filterHeaders } from "../http";

describe("HTTP Utilities", () => {
  describe("jsonResponse", () => {
    it("creates JSON response with correct content type", () => {
      const response = jsonResponse({ status: "ok" });

      expect(response.status).toBe(200);
      expect(response.headers.get("Content-Type")).toBe("application/json");
    });

    it("creates JSON response with CORS headers", () => {
      const response = jsonResponse({ data: "test" });

      expect(response.headers.get("Access-Control-Allow-Origin")).toBe("*");
    });

    it("serializes data correctly", async () => {
      const data = { message: "test", count: 42 };
      const response = jsonResponse(data);

      const parsed = await response.json();
      expect(parsed).toEqual(data);
    });

    it("accepts custom status code", () => {
      const response = jsonResponse({ status: "created" }, 201);

      expect(response.status).toBe(201);
    });

    it("accepts custom headers", () => {
      const response = jsonResponse({ status: "ok" }, 200, {
        "X-Custom-Header": "test",
      });

      expect(response.headers.get("X-Custom-Header")).toBe("test");
    });

    it("merges custom headers with default headers", () => {
      const response = jsonResponse({ status: "ok" }, 200, {
        "X-Custom": "value",
      });

      expect(response.headers.get("Content-Type")).toBe("application/json");
      expect(response.headers.get("Access-Control-Allow-Origin")).toBe("*");
      expect(response.headers.get("X-Custom")).toBe("value");
    });
  });

  describe("jsonError", () => {
    it("creates error response with status code", () => {
      const response = jsonError(400, "Bad request");

      expect(response.status).toBe(400);
    });

    it("includes error message in body", async () => {
      const response = jsonError(404, "Not found");
      const data = (await response.json()) as { error: string };

      expect(data).toHaveProperty("error", "Not found");
    });

    it("creates 401 unauthorized error", async () => {
      const response = jsonError(401, "Unauthorized");

      expect(response.status).toBe(401);
      const data = (await response.json()) as { error: string };
      expect(data.error).toBe("Unauthorized");
    });

    it("creates 403 forbidden error", async () => {
      const response = jsonError(403, "Forbidden");

      expect(response.status).toBe(403);
      const data = (await response.json()) as { error: string };
      expect(data.error).toBe("Forbidden");
    });

    it("creates 429 rate limit error with custom headers", () => {
      const response = jsonError(429, "Rate limited", {
        "X-RateLimit-Remaining": "0",
        "X-RateLimit-Reset": "2024-01-01T00:00:00Z",
      });

      expect(response.status).toBe(429);
      expect(response.headers.get("X-RateLimit-Remaining")).toBe("0");
      expect(response.headers.get("X-RateLimit-Reset")).toBe("2024-01-01T00:00:00Z");
    });
  });

  describe("filterHeaders", () => {
    it("removes Cloudflare headers", () => {
      const headers = new Headers({
        "Content-Type": "application/json",
        "cf-ray": "test-ray",
        "cf-connecting-ip": "1.2.3.4",
        "cf-ipcountry": "US",
      });

      const filtered = filterHeaders(headers, "test-secret");

      expect(filtered.get("Content-Type")).toBe("application/json");
      expect(filtered.get("cf-ray")).toBeNull();
      expect(filtered.get("cf-connecting-ip")).toBeNull();
      expect(filtered.get("cf-ipcountry")).toBeNull();
    });

    it("removes API key header", () => {
      const headers = new Headers({
        "requiems-api-key": "secret-key",
        "Content-Type": "application/json",
      });

      const filtered = filterHeaders(headers, "test-secret");

      expect(filtered.get("requiems-api-key")).toBeNull();
      expect(filtered.get("Content-Type")).toBe("application/json");
    });

    it("removes connection headers", () => {
      const headers = new Headers({
        connection: "keep-alive",
        "keep-alive": "timeout=5",
        "Content-Type": "application/json",
      });

      const filtered = filterHeaders(headers, "test-secret");

      expect(filtered.get("connection")).toBeNull();
      expect(filtered.get("keep-alive")).toBeNull();
      expect(filtered.get("Content-Type")).toBe("application/json");
    });

    it("adds backend secret header", () => {
      const headers = new Headers({
        "Content-Type": "application/json",
      });

      const filtered = filterHeaders(headers, "my-backend-secret");

      expect(filtered.get("X-Backend-Secret")).toBe("my-backend-secret");
    });

    it("preserves custom headers", () => {
      const headers = new Headers({
        "Content-Type": "application/json",
        "X-Custom-Header": "custom-value",
        "User-Agent": "test-agent",
      });

      const filtered = filterHeaders(headers, "secret");

      expect(filtered.get("Content-Type")).toBe("application/json");
      expect(filtered.get("X-Custom-Header")).toBe("custom-value");
      expect(filtered.get("User-Agent")).toBe("test-agent");
    });
  });

  describe("addUsageHeaders", () => {
    it("clones response with usage headers", () => {
      const original = new Response("test body", {
        status: 200,
        headers: { "Content-Type": "text/plain" },
      });

      const modified = addUsageHeaders(original, {
        requestsUsed: 5,
        requestsRemaining: 95,
        requestsReset: "2024-01-01T00:00:00Z",
        plan: "developer",
        rateLimitLimit: 60,
        rateLimitRemaining: 55,
      });

      expect(modified.headers.get("X-Requests-Used")).toBe("5");
      expect(modified.headers.get("X-Requests-Remaining")).toBe("95");
      expect(modified.headers.get("X-Requests-Reset")).toBe("2024-01-01T00:00:00Z");
      expect(modified.headers.get("X-Plan")).toBe("developer");
      expect(modified.headers.get("X-RateLimit-Limit")).toBe("60");
      expect(modified.headers.get("X-RateLimit-Remaining")).toBe("55");
    });

    it("preserves original response status", () => {
      const original = new Response("test", { status: 201 });

      const modified = addUsageHeaders(original, {
        requestsUsed: 1,
        requestsRemaining: 99,
        requestsReset: "2024-01-01",
        plan: "free",
        rateLimitLimit: 10,
        rateLimitRemaining: 9,
      });

      expect(modified.status).toBe(201);
    });

    it("preserves original content-type header", () => {
      const original = new Response("test", {
        headers: { "Content-Type": "application/json" },
      });

      const modified = addUsageHeaders(original, {
        requestsUsed: 1,
        requestsRemaining: 99,
        requestsReset: "2024-01-01",
        plan: "free",
        rateLimitLimit: 10,
        rateLimitRemaining: 9,
      });

      expect(modified.headers.get("Content-Type")).toBe("application/json");
    });

    it("adds CORS header", () => {
      const original = new Response("test");

      const modified = addUsageHeaders(original, {
        requestsUsed: 0,
        requestsRemaining: 100,
        requestsReset: "2024-01-01",
        plan: "free",
        rateLimitLimit: 10,
        rateLimitRemaining: 10,
      });

      expect(modified.headers.get("Access-Control-Allow-Origin")).toBe("*");
    });

    it("strips X-Usage-Count internal header from backend response", () => {
      const original = new Response("test", {
        headers: { "X-Usage-Count": "42" },
      });

      const modified = addUsageHeaders(original, {
        requestsUsed: 42,
        requestsRemaining: 58,
        requestsReset: "2024-01-01",
        plan: "developer",
        rateLimitLimit: 5000,
        rateLimitRemaining: 4999,
      });

      expect(modified.headers.get("X-Usage-Count")).toBeNull();
      expect(modified.headers.get("X-Requests-Used")).toBe("42");
    });
  });

  describe("fetchBackend", () => {
    afterEach(() => {
      vi.restoreAllMocks();
    });

    it("returns successful response", async () => {
      const mockResponse = new Response(JSON.stringify({ data: "test" }), {
        status: 200,
      });
      vi.spyOn(globalThis, "fetch").mockResolvedValueOnce(mockResponse);

      const result = await fetchBackend("https://api.example.com", {
        method: "GET",
      });

      expect(result.ok).toBe(true);
      if (result.ok) {
        expect(result.response.status).toBe(200);
      }
    });

    it("returns error on fetch failure", async () => {
      vi.spyOn(globalThis, "fetch").mockRejectedValueOnce(new Error("Network error"));

      const result = await fetchBackend("https://api.example.com", {
        method: "GET",
      });

      expect(result.ok).toBe(false);
      if (!result.ok) {
        expect(result.error).toBe("Backend unavailable");
        expect(result.status).toBe(502);
      }
    });

    it("returns 504 on timeout", async () => {
      vi.spyOn(globalThis, "fetch").mockImplementationOnce((_url, init) =>
        new Promise((_resolve, reject) => {
          (init?.signal as AbortSignal | undefined)?.addEventListener("abort", () => {
            reject(new DOMException("The operation was aborted.", "AbortError"));
          });
        }),
      );

      const result = await fetchBackend("https://api.example.com", { method: "GET" }, 1);

      expect(result.ok).toBe(false);
      if (!result.ok) {
        expect(result.error).toBe("Backend timeout");
        expect(result.status).toBe(504);
      }
    });
  });

  describe("corsResponse", () => {
    it("has correct CORS headers", () => {
      expect(corsResponse.headers.get("Access-Control-Allow-Origin")).toBe("*");
      expect(corsResponse.headers.get("Access-Control-Allow-Methods")).toBe(
        "GET, POST, PUT, DELETE, OPTIONS",
      );
      expect(corsResponse.headers.get("Access-Control-Allow-Headers")).toBe(
        "Content-Type, requiems-api-key",
      );
      expect(corsResponse.headers.get("Access-Control-Max-Age")).toBe("86400");
    });

    it("has null body", async () => {
      const text = await corsResponse.text();
      expect(text).toBe("");
    });
  });
});
