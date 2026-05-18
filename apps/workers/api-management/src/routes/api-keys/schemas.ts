import * as z from "zod";

export const planSchema = z.enum([
  "free",
  "developer",
  "business",
  "professional",
  "enterprise",
]);

/**
 * Trimmed string, 1–255 characters. `.trim()` runs before `.min`/`.max` so length checks apply to trimmed text.
 */
export const apiKeyLabelStringSchema = z
  .string()
  .trim()
  .min(1, { error: "must not be empty or whitespace only" })
  .max(255, { error: "must be at most 255 characters" });
