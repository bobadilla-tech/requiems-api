import { describe, expect, it } from "vitest";
import { generateApiKey } from "../lib/generate-api-key";
import { extractKeyPrefix, isValidKeyFormat } from "@requiem/workers-shared";

describe("generateApiKey", () => {
  it("produces a key with the requiem_ prefix", () => {
    expect(generateApiKey()).toMatch(/^requiem_/);
  });

  it("produces a key that passes isValidKeyFormat", () => {
    expect(isValidKeyFormat(generateApiKey())).toBe(true);
  });

  it("random part is 24 alphanumeric characters", () => {
    const key = generateApiKey();
    const randomPart = key.slice("requiem_".length);
    expect(randomPart).toHaveLength(24);
    expect(randomPart).toMatch(/^[0-9a-zA-Z]{24}$/);
  });

  it("successive calls return different keys", () => {
    const keys = new Set(Array.from({ length: 20 }, () => generateApiKey()));
    expect(keys.size).toBe(20);
  });

  it("extractKeyPrefix returns the first 12 characters", () => {
    const key = generateApiKey();
    const prefix = extractKeyPrefix(key);
    expect(prefix).toHaveLength(12);
    expect(key.startsWith(prefix)).toBe(true);
  });
});
