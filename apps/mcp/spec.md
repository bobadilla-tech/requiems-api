# Requiems API — MCP Server (POC Spec)

## 1. Goals

Build an MCP server that exposes Requiems API endpoints as MCP tools so AI
assistants can call them directly.

1. Parse the Requiems `openapi.json` spec
2. Generate one MCP tool file per API operation
3. Generate an index that registers all tools
4. Run a TypeScript MCP server backed by the generated tools

### Non-goals (POC phase)

- Advanced auth flows (API key via header is enough)

---

## 2. Architecture Overview

```
requiems openapi.json
     ↓
codegen script
     ↓
generated/
  tools/
    technology_convert.ts
    validation_email.ts
    entertainment_quotes_random.ts
  index.ts
     ↓
src/server.ts (manual, stable)
     ↓
MCP server runtime
```

### Separation of concerns

- **Generated code**: disposable, reproducible from the spec
- **Server runtime**: stable, minimal, hand-written
- **OpenAPI spec**: source of truth

---

## 3. Project Structure

```
apps/mcp/
  src/
    server.ts                # MCP server entrypoint (manual)
  generated/
    tools/                   # generated tool files
    index.ts                 # generated registry
  scripts/
    generate.ts              # codegen CLI
  openapi.json               # Requiems API spec
  package.json               # bun-managed
  bunfig.toml
  tsconfig.json
```

---

## 4. Runtime & Tooling

Use **Bun** throughout — no Node.js, no `ts-node`, no `tsx`.

- `bun run` executes TypeScript directly, no compilation step needed
- `bun install` for dependency management
- Native `fetch` (global) — **do not use `node-fetch`** or any polyfill
- Native `AbortController`, `URL`, `ReadableStream` — no shims needed
- Prefer Bun-native APIs over npm packages wherever possible

### Common commands

```
bun install
bun run scripts/generate.ts --input openapi.json --output generated
bun run src/server.ts
```

---

## 5. Codegen CLI (`scripts/generate.ts`)

### Responsibilities

- Parse the Requiems OpenAPI spec
- Extract all operations
- Generate one tool file per operation
- Generate a central index

### CLI interface

```
bun run scripts/generate.ts \
  --input openapi.json \
  --output generated \
  --base-url https://api.requiems.io
```

### Internal pipeline

```
load spec
  ↓
validate minimal structure
  ↓
extract operations
  ↓
normalize names
  ↓
generate tool files
  ↓
generate index
```

---

## 6. Operation → Tool Mapping

### Naming

Use `operationId` if present.

Fallback — combine path segments:

```
GET /v1/technology/convert
→ technology_convert

GET /v1/entertainment/quotes/random
→ entertainment_quotes_random

POST /v1/validation/email
→ validation_email_post
```

Normalize:

- strip `/v1/` prefix
- replace `/` with `_`
- remove `{param}` braces (param name stays as suffix)
- lowercase

---

## 7. Generated Tool File

### Example: `generated/tools/technology_convert.ts`

Maps `GET /v1/technology/convert` — converts units (miles → km, kg → lb, °C →
°F, etc.).

```ts
export const technology_convert = {
  name: "technology_convert",
  description:
    "Convert a value between units. Supports length, weight, volume, temperature, area, and speed.",

  inputSchema: {
    type: "object",
    properties: {
      from: {
        type: "string",
        description: "Source unit (e.g. miles, kg, celsius)",
      },
      to: {
        type: "string",
        description: "Target unit (e.g. km, lb, fahrenheit)",
      },
      value: { type: "number", description: "Numeric value to convert" },
    },
    required: ["from", "to", "value"],
  },

  handler: async (args: Record<string, unknown>) => {
    const params = new URLSearchParams({
      from: String(args.from),
      to: String(args.to),
      value: String(args.value),
    });

    const controller = new AbortController();
    const timeout = setTimeout(() => controller.abort(), 5000);

    try {
      const res = await fetch(
        `${process.env.REQUIEMS_BASE_URL}/v1/technology/convert?${params}`,
        {
          headers: { "requiems-api-key": process.env.REQUIEMS_API_KEY ?? "" },
          signal: controller.signal,
        },
      );

      if (!res.ok) {
        throw new Error(`Request failed: ${res.status}`);
      }

      return await res.json();
    } finally {
      clearTimeout(timeout);
    }
  },
};
```

### Example: `generated/tools/validation_email.ts`

Maps `POST /v1/validation/email` — validates an email address (syntax, MX,
disposable check).

```ts
export const validation_email = {
  name: "validation_email",
  description:
    "Validate an email address. Returns syntax check, MX record lookup, and disposable domain detection.",

  inputSchema: {
    type: "object",
    properties: {
      email: { type: "string", description: "Email address to validate" },
    },
    required: ["email"],
  },

  handler: async (args: Record<string, unknown>) => {
    const controller = new AbortController();
    const timeout = setTimeout(() => controller.abort(), 5000);

    try {
      const res = await fetch(
        `${process.env.REQUIEMS_BASE_URL}/v1/validation/email`,
        {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
            "requiems-api-key": process.env.REQUIEMS_API_KEY ?? "",
          },
          body: JSON.stringify({ email: args.email }),
          signal: controller.signal,
        },
      );

      if (!res.ok) {
        throw new Error(`Request failed: ${res.status}`);
      }

      return await res.json();
    } finally {
      clearTimeout(timeout);
    }
  },
};
```

### Design decisions

- Export a plain object — no SDK coupling inside generated files
- `REQUIEMS_BASE_URL` and `REQUIEMS_API_KEY` injected via env
- Native `fetch` + `AbortController` for timeouts — no extra deps
- Keep handler isolated and unit-testable

---

## 8. Generated Index

### `generated/index.ts`

```ts
import { technology_convert } from "./tools/technology_convert";
import { validation_email } from "./tools/validation_email";
import { entertainment_quotes_random } from "./tools/entertainment_quotes_random";

export const tools = [
  technology_convert,
  validation_email,
  entertainment_quotes_random,
];
```

---

## 9. MCP Server (Manual Code)

### `src/server.ts`

```ts
import { McpServer } from "@modelcontextprotocol/sdk/server/mcp.js";
import { StdioServerTransport } from "@modelcontextprotocol/sdk/server/stdio.js";
import { tools } from "../generated";

const server = new McpServer({
  name: "requiems-api",
  version: "0.1.0",
});

for (const tool of tools) {
  server.tool(
    tool.name,
    tool.description,
    tool.inputSchema.properties,
    tool.handler,
  );
}

const transport = new StdioServerTransport();
await server.connect(transport);
```

### Why this is manual

- Keeps control over runtime
- Allows middleware, logging, auth injection later
- Prevents regeneration from breaking server behavior

---

## 10. Requiems API — Target Operations (POC)

Focus on simple GET + single-body-param POST endpoints for the POC:

| Tool name                     | Method | Path                              | Description              |
| ----------------------------- | ------ | --------------------------------- | ------------------------ |
| `technology_convert`          | GET    | `/v1/technology/convert`          | Unit conversion          |
| `technology_useragent`        | GET    | `/v1/technology/useragent`        | Parse user-agent string  |
| `technology_password`         | GET    | `/v1/technology/password`         | Generate secure password |
| `technology_color`            | GET    | `/v1/technology/color`            | Color format conversion  |
| `validation_email`            | POST   | `/v1/validation/email`            | Email validation         |
| `validation_phone`            | GET    | `/v1/validation/phone`            | Phone number validation  |
| `entertainment_quotes_random` | GET    | `/v1/entertainment/quotes/random` | Random quote             |
| `entertainment_trivia`        | GET    | `/v1/entertainment/trivia`        | Random trivia            |
| `networking_ip`               | GET    | `/v1/networking/ip/{ip}`          | IP geolocation / info    |

Batch endpoints (e.g. `/email/batch`) deferred to phase 2.

---

## 11. Schema Handling (POC Scope)

### Supported initially

- `path` parameters
- `query` parameters
- `requestBody` with flat object schema (single depth)
- basic types: `string`, `number`, `boolean`

### Deferred

- deeply nested schemas
- `oneOf`, `anyOf`
- response streaming

---

## 12. Code Generation Strategy

### Deterministic output

- Same spec → same files
- No timestamps or randomness
- Stable ordering (alphabetical by operationId)

### Idempotency

Running codegen multiple times:

- overwrites `generated/` files
- does not touch `src/server.ts` or any manual files

### File boundaries

- One operation = one file
- Keeps diffs small and reviewable

---

## 13. Reliability

### API failures

```ts
if (!res.ok) {
  throw new Error(`Requiems API error: ${res.status}`);
}
```

### Timeouts

5 s default via native `AbortController` — included in every generated handler
(see examples above).

### Auth

`requiems-api-key` header injected from `process.env.REQUIEMS_API_KEY`. Never
hardcoded.

---

## 14. Extensibility Hooks

Plan for:

- custom request transformers per tool
- response shaping
- per-tool overrides (drop a file in `src/overrides/technology_convert.ts` to
  replace generated handler)

---

## 15. Risks & Tradeoffs

### Static vs dynamic generation

Static generation chosen for POC:

- type safety, debuggable, versionable diffs
- requires regeneration when spec changes (acceptable — spec is stable)

---

## 16. Success Criteria

- `bun run scripts/generate.ts` produces tool files from `openapi.json`
- `bun run src/server.ts` starts without error
- Claude (or any MCP client) can call `technology_convert`, `validation_email`,
  and `entertainment_quotes_random` and receive real Requiems API responses

---

## 17. Next Steps After POC

- Add remaining Requiems operations
- Batch endpoint support (phase 2)
- CI: regenerate on spec change, diff in PR
