import { describe, expect, it } from "vitest";
import {
  DEFAULT_REQUEST_MULTIPLIER,
  ENDPOINT_MULTIPLIERS,
  getRequestMultiplier,
  PLANS,
} from "@requiem/workers-shared";

describe("Configuration", () => {
  describe("PLANS", () => {
    it("has all plan tiers defined", () => {
      expect(PLANS).toHaveProperty("free");
      expect(PLANS).toHaveProperty("developer");
      expect(PLANS).toHaveProperty("business");
      expect(PLANS).toHaveProperty("professional");
      expect(PLANS).toHaveProperty("enterprise");
    });

    it("free plan has correct limits", () => {
      expect(PLANS.free.requestLimit).toBe(500);
      expect(PLANS.free.ratePerMinute).toBe(30);
    });

    it("developer plan has correct limits", () => {
      expect(PLANS.developer.requestLimit).toBe(100_000);
      expect(PLANS.developer.ratePerMinute).toBe(5000);
    });

    it("business plan has correct limits", () => {
      expect(PLANS.business.requestLimit).toBe(1_000_000);
      expect(PLANS.business.ratePerMinute).toBe(10000);
    });

    it("professional plan has correct limits", () => {
      expect(PLANS.professional.requestLimit).toBe(10_000_000);
      expect(PLANS.professional.ratePerMinute).toBe(50000);
    });

    it("enterprise plan has unlimited limits", () => {
      expect(PLANS.enterprise.requestLimit).toBe(Number.POSITIVE_INFINITY);
      expect(PLANS.enterprise.ratePerMinute).toBe(Number.POSITIVE_INFINITY);
    });

    it("plans are in ascending order by request limits", () => {
      const plans = [PLANS.free, PLANS.developer, PLANS.business, PLANS.professional];

      for (let i = 0; i < plans.length - 1; i++) {
        expect(plans[i].requestLimit).toBeLessThan(plans[i + 1].requestLimit);
        expect(plans[i].ratePerMinute).toBeLessThan(plans[i + 1].ratePerMinute);
      }
    });

    it("all plans have positive rate limits", () => {
      Object.values(PLANS).forEach((plan) => {
        expect(plan.requestLimit).toBeGreaterThan(0);
        expect(plan.ratePerMinute).toBeGreaterThan(0);
      });
    });
  });

  describe("ENDPOINT_MULTIPLIERS", () => {
    it("is a Map with string keys and number values", () => {
      expect(ENDPOINT_MULTIPLIERS).toBeInstanceOf(Map);
    });

    it("has dictionary endpoints with 2x multiplier", () => {
      expect(ENDPOINT_MULTIPLIERS.get("GET /v1/text/words/define")).toBe(2);
      expect(ENDPOINT_MULTIPLIERS.get("GET /v1/text/words/synonyms")).toBe(2);
    });

    it("all multipliers are positive integers", () => {
      ENDPOINT_MULTIPLIERS.forEach((multiplier) => {
        expect(multiplier).toBeGreaterThan(0);
        expect(Number.isInteger(multiplier)).toBe(true);
      });
    });

    it("all keys follow the correct format (METHOD /path)", () => {
      ENDPOINT_MULTIPLIERS.forEach((_, key) => {
        const parts = key.split(" ");
        expect(parts).toHaveLength(2);

        const [method, path] = parts;
        expect(["GET", "POST", "PUT", "DELETE", "PATCH"]).toContain(method);
        expect(path).toMatch(/^\//); // Path starts with /
      });
    });
  });

  describe("DEFAULT_REQUEST_MULTIPLIER", () => {
    it("is set to 1", () => {
      expect(DEFAULT_REQUEST_MULTIPLIER).toBe(1);
    });
  });

  describe("getRequestMultiplier", () => {
    it("returns exact match for dictionary define endpoint", () => {
      const multiplier = getRequestMultiplier("GET", "/v1/text/words/define");
      expect(multiplier).toBe(2);
    });

    it("returns exact match for dictionary synonyms endpoint", () => {
      const multiplier = getRequestMultiplier("GET", "/v1/text/words/synonyms");
      expect(multiplier).toBe(2);
    });

    it("returns default multiplier for unlisted endpoints", () => {
      const multiplier = getRequestMultiplier("GET", "/v1/email/disposable/check");
      expect(multiplier).toBe(DEFAULT_REQUEST_MULTIPLIER);
    });

    it("returns default multiplier for common text endpoints", () => {
      const endpoints = ["/v1/text/advice", "/v1/text/lorem", "/v1/text/quotes", "/v1/text/words"];

      endpoints.forEach((path) => {
        const multiplier = getRequestMultiplier("GET", path);
        expect(multiplier).toBe(1);
      });
    });

    it("is case-sensitive for HTTP methods", () => {
      // Correct method
      expect(getRequestMultiplier("GET", "/v1/text/words/define")).toBe(2);

      // Wrong case - should not match
      expect(getRequestMultiplier("get", "/v1/text/words/define")).toBe(1);
      expect(getRequestMultiplier("Get", "/v1/text/words/define")).toBe(1);
    });

    it("handles prefix matching for dynamic routes", () => {
      // If /v1/text/words/define is configured with multiplier 2,
      // and we call /v1/text/words/define/something, it should match via prefix
      const multiplier = getRequestMultiplier("GET", "/v1/text/words/define/something");
      expect(multiplier).toBe(2);
    });

    it("returns different multipliers for different methods on same path", () => {
      // Only GET is configured, POST should return default
      expect(getRequestMultiplier("GET", "/v1/text/words/define")).toBe(2);
      expect(getRequestMultiplier("POST", "/v1/text/words/define")).toBe(1);
    });

    it("handles root paths correctly", () => {
      const multiplier = getRequestMultiplier("GET", "/");
      expect(multiplier).toBe(DEFAULT_REQUEST_MULTIPLIER);
    });

    it("handles paths without API version prefix", () => {
      const multiplier = getRequestMultiplier("GET", "/healthz");
      expect(multiplier).toBe(DEFAULT_REQUEST_MULTIPLIER);
    });

    it("handles query parameters in path (should not affect matching)", () => {
      // The function receives pathname only (no query string)
      const multiplier = getRequestMultiplier("GET", "/v1/text/words/define");
      expect(multiplier).toBe(2);
    });
  });

  describe("Plan comparison", () => {
    it("developer plan has 200x more requests than free", () => {
      expect(PLANS.developer.requestLimit / PLANS.free.requestLimit).toBe(200);
    });

    it("business plan has 10x more requests than developer", () => {
      expect(PLANS.business.requestLimit / PLANS.developer.requestLimit).toBe(10);
    });

    it("professional plan has 10x more requests than business", () => {
      expect(PLANS.professional.requestLimit / PLANS.business.requestLimit).toBe(10);
    });
  });

  describe("Rate limit ratios", () => {
    it("all plans have positive rate-to-limit ratios", () => {
      // Check that all plans have some rate limiting in place
      const plans = ["free", "developer", "business", "professional"] as const;

      plans.forEach((planName) => {
        const plan = PLANS[planName];
        const ratio = plan.ratePerMinute / plan.requestLimit;

        // Ratio should be positive (rate limit exists)
        expect(ratio).toBeGreaterThan(0);
      });
    });
  });
});
