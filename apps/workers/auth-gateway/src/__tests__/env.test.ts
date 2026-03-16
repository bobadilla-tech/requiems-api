import { describe, expect, it } from "vitest";
import { validateEnv } from "../env";

const validEnv = {
  BACKEND_URL: "https://api.internal.example.com",
  BACKEND_SECRET: "a".repeat(32),
  CORS_ALLOWED_ORIGINS: "https://example.com",
  ENVIRONMENT: "production" as const,
  KV: {} as KVNamespace,
  DB: {} as D1Database,
};

describe("validateEnv", () => {
  it("accepts valid environment", () => {
    expect(() => validateEnv(validEnv)).not.toThrow();
  });

  it("throws when CORS_ALLOWED_ORIGINS is missing", () => {
    const env = { ...validEnv, CORS_ALLOWED_ORIGINS: undefined as unknown as string };
    expect(() => validateEnv(env)).toThrow();
  });

  it("throws when CORS_ALLOWED_ORIGINS is empty", () => {
    const env = { ...validEnv, CORS_ALLOWED_ORIGINS: "" };
    expect(() => validateEnv(env)).toThrow(/CORS_ALLOWED_ORIGINS must be a non-empty string/);
  });

  it("accepts wildcard CORS_ALLOWED_ORIGINS", () => {
    const env = { ...validEnv, CORS_ALLOWED_ORIGINS: "*" };
    expect(() => validateEnv(env)).not.toThrow();
  });

  it("accepts comma-separated origins in CORS_ALLOWED_ORIGINS", () => {
    const env = {
      ...validEnv,
      CORS_ALLOWED_ORIGINS: "https://app.example.com,https://admin.example.com",
    };
    expect(() => validateEnv(env)).not.toThrow();
  });

  it("throws when BACKEND_URL is missing", () => {
    const env = { ...validEnv, BACKEND_URL: undefined as unknown as string };
    expect(() => validateEnv(env)).toThrow();
  });

  it("throws when BACKEND_SECRET is too short", () => {
    const env = { ...validEnv, BACKEND_SECRET: "short" };
    expect(() => validateEnv(env)).toThrow();
  });
});
