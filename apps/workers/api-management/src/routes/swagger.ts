import { Hono } from "hono";
import { SwaggerUI } from "@hono/swagger-ui";
import type { WorkerBindings } from "../env";

const swaggerRoute = new Hono<{ Bindings: WorkerBindings }>();

function getOpenApiMeta(isDev: boolean) {
  const servers = isDev
    ? [
      { url: "http://localhost:5544", description: "Local development" },
      { url: "https://api-management.requiems.xyz", description: "Production" },
    ]
    : [
      { url: "https://api-management.requiems.xyz", description: "Production" },
      { url: "http://localhost:5544", description: "Local development" },
    ];

  const description = isDev
    ? "Internal API for managing API keys, usage data, and analytics. Only accessible by Rails dashboard.\n\n**Local dev key:** `dev_api_mgmt_key_for_local_dev_only`"
    : "Internal API for managing API keys, usage data, and analytics. Only accessible by Rails dashboard.";

  return { servers, description };
}

/**
 * OpenAPI specification endpoint
 * Provides machine-readable API documentation
 */
swaggerRoute.get("/openapi.json", (c) => {
  const isDev = c.env.ENVIRONMENT !== "production";
  const { servers, description } = getOpenApiMeta(isDev);

  return c.json({
    openapi: "3.0.0",
    info: {
      title: "Requiems API Management",
      version: "1.0.0",
      description,
    },
    servers,
    components: {
      securitySchemes: {
        ApiManagementKey: {
          type: "apiKey",
          in: "header",
          name: "X-API-Management-Key",
          description: "API Management key (only Rails dashboard has this)",
        },
      },
    },
    security: [{ ApiManagementKey: [] }],
    paths: {
      "/healthz": {
        get: {
          summary: "Health check",
          security: [],
          responses: {
            "200": {
              description: "Service is healthy",
            },
          },
        },
      },
      "/api-keys": {
        get: {
          summary: "List API keys",
          description:
            "Returns metadata for all API keys. Never includes full key values. Filter by userId or active status.",
          parameters: [
            {
              name: "userId",
              in: "query",
              required: false,
              schema: { type: "string" },
              description: "Filter by user ID",
            },
            {
              name: "active",
              in: "query",
              required: false,
              schema: { type: "boolean", default: true },
              description: "Set to false to include revoked keys",
            },
          ],
          responses: {
            "200": {
              description: "List of API key records (no full keys)",
              content: {
                "application/json": {
                  schema: {
                    type: "object",
                    properties: {
                      total: { type: "integer" },
                      keys: {
                        type: "array",
                        items: {
                          type: "object",
                          properties: {
                            keyPrefix: {
                              type: "string",
                              description: "First 12 chars of the key",
                            },
                            userId: { type: "string" },
                            plan: { type: "string" },
                            active: { type: "boolean" },
                            createdAt: { type: "string", format: "date-time" },
                            updatedAt: {
                              type: "string",
                              format: "date-time",
                              nullable: true,
                            },
                            revokedAt: {
                              type: "string",
                              format: "date-time",
                              nullable: true,
                            },
                            billingCycleStart: {
                              type: "string",
                              format: "date-time",
                            },
                          },
                        },
                      },
                    },
                  },
                },
              },
            },
            "401": { description: "Unauthorized" },
          },
        },
        post: {
          summary: "Create a new API key",
          description:
            "Generates a new API key on the server and returns it. The full key is only returned once.",
          requestBody: {
            required: true,
            content: {
              "application/json": {
                schema: {
                  type: "object",
                  required: ["userId", "plan", "name"],
                  properties: {
                    userId: { type: "string" },
                    plan: {
                      type: "string",
                      enum: [
                        "free",
                        "developer",
                        "business",
                        "professional",
                        "enterprise",
                      ],
                    },
                    name: {
                      type: "string",
                      description: "Human-readable name for the key",
                    },
                    billingCycleStart: { type: "string", format: "date-time" },
                  },
                },
              },
            },
          },
          responses: {
            "201": {
              description: "API key created successfully",
              content: {
                "application/json": {
                  schema: {
                    type: "object",
                    properties: {
                      apiKey: {
                        type: "string",
                        description: "Full API key (store securely)",
                      },
                      keyPrefix: {
                        type: "string",
                        description: "First 12 chars for display",
                      },
                      userId: { type: "string" },
                      plan: { type: "string" },
                      createdAt: { type: "string", format: "date-time" },
                    },
                  },
                },
              },
            },
            "422": {
              description: "Request body failed validation",
              content: {
                "application/json": {
                  schema: {
                    type: "object",
                    properties: {
                      error: { type: "string", enum: ["validation_failed"] },
                      fields: {
                        type: "array",
                        items: {
                          type: "object",
                          properties: {
                            field: { type: "string" },
                            message: { type: "string" },
                          },
                        },
                      },
                    },
                  },
                },
              },
            },
            "401": { description: "Unauthorized" },
          },
        },
      },
      "/api-keys/{keyPrefix}": {
        delete: {
          summary: "Revoke an API key",
          description: "Deletes the key from KV and marks as revoked in D1",
          parameters: [
            {
              name: "keyPrefix",
              in: "path",
              required: true,
              schema: { type: "string" },
              description: "First 12 characters of the API key",
            },
          ],
          responses: {
            "200": { description: "API key revoked successfully" },
            "404": { description: "API key not found" },
            "401": { description: "Unauthorized" },
          },
        },
        patch: {
          summary: "Update an API key",
          description: "Update plan or billing cycle for an existing key",
          parameters: [
            {
              name: "keyPrefix",
              in: "path",
              required: true,
              schema: { type: "string" },
              description: "First 12 characters of the API key",
            },
          ],
          requestBody: {
            required: true,
            content: {
              "application/json": {
                schema: {
                  type: "object",
                  properties: {
                    plan: {
                      type: "string",
                      enum: [
                        "free",
                        "developer",
                        "business",
                        "professional",
                        "enterprise",
                      ],
                    },
                    billingCycleStart: { type: "string", format: "date-time" },
                  },
                },
              },
            },
          },
          responses: {
            "200": { description: "API key updated successfully" },
            "404": { description: "API key not found" },
            "401": { description: "Unauthorized" },
          },
        },
      },
      "/usage/export": {
        get: {
          summary: "Export usage data from D1",
          parameters: [
            {
              name: "since",
              in: "query",
              required: true,
              schema: { type: "string" },
            },
            { name: "limit", in: "query", schema: { type: "integer" } },
            { name: "cursor", in: "query", schema: { type: "string" } },
          ],
          responses: {
            "200": { description: "Usage data with pagination" },
          },
        },
      },
      "/analytics/by-endpoint": {
        get: {
          summary: "Usage breakdown by endpoint",
          parameters: [
            {
              name: "userId",
              in: "query",
              required: true,
              schema: { type: "string" },
            },
          ],
          responses: {
            "200": { description: "Endpoint usage statistics" },
          },
        },
      },
      "/analytics/by-date": {
        get: {
          summary: "Usage trends over time",
          parameters: [
            {
              name: "userId",
              in: "query",
              required: true,
              schema: { type: "string" },
            },
          ],
          responses: {
            "200": { description: "Time series usage data" },
          },
        },
      },
      "/analytics/summary": {
        get: {
          summary: "Overall usage summary",
          parameters: [
            {
              name: "userId",
              in: "query",
              required: true,
              schema: { type: "string" },
            },
          ],
          responses: {
            "200": { description: "Usage summary" },
          },
        },
      },
    },
  });
});

swaggerRoute.get("/docs", (c) => {
  const isDev = c.env.ENVIRONMENT !== "production";
  const onComplete = isDev
    ? `() => { window.ui.preauthorizeApiKey('ApiManagementKey', '${c.env.API_MANAGEMENT_API_KEY}') }`
    : undefined;
  return c.html(`<!doctype html>
<html lang="en">
  <head>
    <meta charset="utf-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1" />
    <title>API Management</title>
  </head>
  <body>
    ${
    SwaggerUI({ url: "/openapi.json", persistAuthorization: true, onComplete })
  }
  </body>
</html>`);
});

export { swaggerRoute };
