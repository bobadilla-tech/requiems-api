# Requiems API MCP Server

The project parses the Requiems OpenAPI specification, generates one MCP tool per API operation, builds a tool registry, and runs a lightweight MCP server capable of serving those tools to any MCP-compatible client (Claude Desktop, Cursor, VS Code MCP, etc.).

The OpenAPI specification is the single source of truth. Whenever the API changes, simply regenerate the tools.

---

# Features

* Built entirely with Bun
* Automatic MCP tool generation from OpenAPI
* One generated file per API endpoint
* Automatic Zod schema generation
* Native fetch (no axios or node-fetch)
* Deterministic and idempotent code generation
* Runtime environment validation
* Minimal handwritten MCP runtime
* Compatible with the official MCP SDK

---


# Project Structure

```
apps/mcp/
├── generated/
│   ├── index.ts
│   ├── runtime.ts
│   └── tools/
│       ├── technology_convert.ts
│       ├── validation_email.ts
│       ├── networking_ip.ts
│       └── ...
│
├── scripts/
│   └── generate.ts
│
├── src/
│   └── server.ts
│
├── openapi.json
├── package.json
├── tsconfig.json
└── bunfig.toml
```

---

# How It Works

The project consists of two independent parts.

## 1. Code Generator

The generator reads the OpenAPI specification and produces:

* one TypeScript file per API operation
* a central tool registry
* strongly typed Zod schemas
* request handlers

Generated files should never be edited manually.

Whenever the API specification changes:

```
OpenAPI
    ↓
generate.ts
    ↓
generated/*
```

---

## 2. Runtime

The runtime is intentionally handwritten and remains stable.

Its responsibilities are:

* validate environment variables
* register generated tools
* initialize the MCP server
* connect the stdio transport
* report startup/runtime errors

The runtime contains **no API-specific business logic**.

All endpoint-specific logic lives inside the generated tools.

---

# Generated Tool Structure

Every API operation generates an independent module.

Example:

```ts
export const technology_convert = {
    name: "technology_convert",
    description: "...",
    inputSchema,
    handler: async (...) => { ... }
}
```

Each generated tool contains:

* metadata
* description
* Zod input validation
* request handler
* TypeScript types


# Generating Tools

Generate tools from the OpenAPI specification.

```
bun run generate
```

Equivalent command:

```
bun run scripts/generate.ts \
    --input openapi.json \
    --output generated
```

The generator will:

* load the OpenAPI specification
* validate the document
* extract supported operations
* skip batch endpoints
* generate one tool per endpoint
* generate the central tool registry

---

# Running the MCP Server

Start the runtime:

```
bun run start
```

The server uses the standard MCP stdio transport.

If startup succeeds you'll see something similar to:

```
[server] Starting MCP server...
[server] Registered 42 tools
[server] MCP server running
```

---

# Development Workflow

Whenever the API changes:

Generate new tools

```
bun run generate
```

Restart the MCP server

```
bun run start
```

No manual edits inside the `generated/` directory are required.

---

# Available Scripts

Fetch OpenAPI specification

```
bun run fetch-openapi

```

Generate MCP tools

```
bun run generate
```

Start the MCP runtime

```
bun run start
```

---

# Documentation

This repository documents only the MCP server.

For the complete Requiems API documentation, endpoint reference, and examples, refer to the main Requiems documentation.
