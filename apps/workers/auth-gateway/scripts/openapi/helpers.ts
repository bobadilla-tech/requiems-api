import { readdir, readFile } from "node:fs/promises";
import { join } from "node:path";

import * as yaml from "js-yaml";

import type { CatalogEntry, YamlApiDoc, YamlEndpoint, YamlError } from "./types";
import { apiDocsDir, baseSpec, catalogPath, METHODS_WITH_BODY, TYPE_SCHEMAS } from "./constants";

export function yamlTypeToSchema(type: string): Record<string, unknown> {
  return TYPE_SCHEMAS[type] ?? { type: "string" };
}

const KNOWN_ERROR_CODES: Record<string, number> = {
  validation_failed: 422,
  not_found: 404,
  unauthorized: 401,
  forbidden: 403,
  internal_error: 500,
  bad_request: 400,
  service_unavailable: 503,
};

export function getErrorStatus(error: YamlError): number {
  if (typeof error.code === "number") return error.code;
  if (typeof error.status === "number") return error.status;

  if (typeof error.status === "string") {
    const parsed = Number.parseInt(error.status, 10);
    if (!Number.isNaN(parsed)) return parsed;
  }

  return KNOWN_ERROR_CODES[String(error.code)] ?? 400;
}

export function buildOperation(endpoint: YamlEndpoint, apiId: string): Record<string, unknown> {
  const params = endpoint.parameters ?? [];

  const pathAndQueryParams = params
    .filter((p) => p.location === "path" || p.location === "query")
    .map((p) => {
      const schema: Record<string, unknown> = {
        ...yamlTypeToSchema(p.type),
      };

      if (p.example !== undefined) schema.example = p.example;

      const param: Record<string, unknown> = {
        name: p.name,
        in: p.location,
        required: p.location === "path" ? true : (p.required ?? false),
        schema,
      };

      if (p.description) param.description = p.description;
      return param;
    });

  const bodyParams = params.filter((p) => !p.location || p.location === "body");

  const operation: Record<string, unknown> = {
    summary: endpoint.name,
    tags: [apiId],
    security: [{ "requiems-api-key": [] }],
  };

  if (endpoint.description) operation.description = endpoint.description;
  if (pathAndQueryParams.length > 0) operation.parameters = pathAndQueryParams;

  // Build requestBody for POST/PUT/PATCH with body params
  const method = endpoint.method.toUpperCase();

  if (bodyParams.length > 0 && METHODS_WITH_BODY.includes(method)) {
    const properties: Record<string, unknown> = {};
    const required: string[] = [];

    for (const p of bodyParams) {
      const schema: Record<string, unknown> = { ...yamlTypeToSchema(p.type) };
      if (p.description) schema.description = p.description;
      if (p.example !== undefined) schema.example = p.example;
      properties[p.name] = schema;
      if (p.required) required.push(p.name);
    }

    const bodySchema: Record<string, unknown> = {
      type: "object",
      properties,
    };

    if (required.length > 0) bodySchema.required = required;

    // Include request example if available
    if (endpoint.request_example) {
      try {
        bodySchema.example = JSON.parse(endpoint.request_example);
      } catch {
        // skip malformed examples
      }
    }

    operation.requestBody = {
      required: true,
      content: {
        "application/json": { schema: bodySchema },
      },
    };
  }

  const responses: Record<string, unknown> = {};

  const successResponse: Record<string, unknown> = {
    description: "Successful response",
  };

  if (endpoint.response_example) {
    try {
      const example = JSON.parse(endpoint.response_example);
      const responseFields = endpoint.response_fields ?? [];
      const dataProperties: Record<string, unknown> = {};

      for (const field of responseFields) {
        const fieldSchema: Record<string, unknown> = {
          ...yamlTypeToSchema(field.type),
        };

        if (field.description) fieldSchema.description = field.description;
        dataProperties[field.name] = fieldSchema;
      }

      successResponse.content = {
        "application/json": {
          schema: {
            type: "object",
            properties: {
              data: {
                type: "object",
                properties: dataProperties,
              },
              metadata: {
                type: "object",
                properties: {
                  timestamp: { type: "string", format: "date-time" },
                },
              },
            },
          },
          example,
        },
      };
    } catch {
      // skip malformed response examples
    }
  }

  responses["200"] = successResponse;

  // Error responses
  const errorsByStatus: Record<number, string[]> = {};

  for (const err of endpoint.errors ?? []) {
    const status = getErrorStatus(err);

    if (!errorsByStatus[status]) errorsByStatus[status] = [];

    const msg = err.description ?? err.message ?? String(err.code ?? "Error");
    errorsByStatus[status].push(msg);
  }

  for (const [status, messages] of Object.entries(errorsByStatus)) {
    responses[status] = { description: messages.join("; ") };
  }

  operation.responses = responses;

  return operation;
}

export async function loadCatalog(): Promise<{ apis: CatalogEntry[] }> {
  try {
    const catalogContent = await readFile(catalogPath, "utf8");

    return yaml.load(catalogContent) as {
      apis: CatalogEntry[];
    };
  } catch (err) {
    console.error(`❌ Failed to read catalog at ${catalogPath}:`, err);
    process.exit(1);
    throw err;
  }
}

export async function loadAPIDocs(): Promise<string[]> {
  try {
    const files = await readdir(apiDocsDir);

    return files.filter((f) => f.endsWith(".yml") || f.endsWith(".yaml")).sort();
  } catch (err) {
    console.error(`❌ Failed to read api_docs directory at ${apiDocsDir}:`, err);
    process.exit(1);
    throw err;
  }
}

export async function loadAPIDoc(file: string) {
  try {
    const apiDocContent = await readFile(join(apiDocsDir, file), "utf8");
    return yaml.load(apiDocContent) as YamlApiDoc;
  } catch (err) {
    console.warn(`⚠️  Skipping ${file}: failed to parse YAML —`, err);
    return null;
  }
}

export function buildSpec(
  tags: { name: string; description: string }[],
  paths: Record<string, Record<string, unknown>>,
) {
  const spec = {
    ...baseSpec,
    tags,
    paths,
  };

  return spec;
}
