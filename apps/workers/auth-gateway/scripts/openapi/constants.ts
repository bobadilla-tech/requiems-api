import { join, resolve } from "node:path";

export const TYPE_SCHEMAS: Record<string, Record<string, unknown>> = {
  array: { type: "array", items: {} },
  object: { type: "object" },
  integer: { type: "integer" },
  number: { type: "number" },
  boolean: { type: "boolean" },
  string: { type: "string" },
};

export const METHODS_WITH_BODY = ["POST", "PUT", "PATCH"];

const monorepoRoot = resolve(import.meta.dirname, "../../../../../");

export const apiDocsDir = join(monorepoRoot, "apps/dashboard/config/api_docs");
export const catalogPath = join(monorepoRoot, "apps/dashboard/config/api_catalog.yml");

export const baseSpec = {
  openapi: "3.0.3",
  info: {
    title: "Requiems API",
    version: "1.0.0",
    description:
      "Unified access to enterprise-grade APIs — email validation, text utilities, and more. Authenticate with the `requiems-api-key` header.",
  },
  servers: [{ url: "https://api.requiems.xyz", description: "Production" }],
  components: {
    securitySchemes: {
      "requiems-api-key": {
        type: "apiKey",
        in: "header",
        name: "requiems-api-key",
        description: "Your Requiems API key",
      },
    },
  },
  security: [{ "requiems-api-key": [] }],
};
