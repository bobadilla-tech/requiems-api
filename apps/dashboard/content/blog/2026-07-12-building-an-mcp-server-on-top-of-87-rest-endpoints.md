---
title: "Building an MCP Server on Top of 87 REST Endpoints"
slug: "building-an-mcp-server-on-top-of-87-rest-endpoints"
date: 2026-07-12
author: "Eliaz Bobadilla"
description: "How we turned Requiems API's OpenAPI spec into an MCP server with 87 tools — the codegen approach, the stdio/HTTP transport split, and the bugs we hit shipping it."
---

Requiems API is a REST API platform: dozens of small, focused endpoints for
things like email validation, IBAN checks, geocoding, and fraud scoring. Every
one of those endpoints is already described by an OpenAPI spec we generate from
the API itself. When [MCP](https://modelcontextprotocol.io) started showing up
in Claude Desktop, Cursor, and every other agentic tool, the obvious question
was: can an LLM just call our API directly, as tools, without us hand-writing 87
tool definitions?

This is how we built [`mcp.requiems.xyz`](https://mcp.requiems.xyz) — a real MCP
server sitting in front of our existing API, generated from our OpenAPI spec,
and what broke along the way.

## Codegen over hand-written tools

The core decision was to treat our OpenAPI spec as the source of truth and
generate MCP tool definitions from it, rather than writing and maintaining 87
tool schemas by hand. `apps/mcp/scripts/generate.ts` does the work:

1. Pull the live `openapi.json` from the API (`fetch-spec.ts`).
2. For each operation, derive a tool name — `operationId` if the spec has one,
   otherwise derived from the path and method (`GET /v1/technology/convert` →
   `technology_convert`; non-GET methods get a `_post`/`_delete` suffix so a
   `GET`/`POST` pair on the same path doesn't collide).
3. Convert the OpenAPI parameter and request-body schemas into Zod schemas.
4. Write one generated file per tool, plus an `index.ts` that imports and
   registers all of them.

```typescript
// apps/mcp/generated/tools/technology_convert.ts (shape, simplified)
export const technology_convert = {
  name: "technology_convert",
  description: "Convert a value between unit systems",
  inputSchema: z.object({
    value: z.number(),
    from: z.string(),
    to: z.string(),
  }),
  handler: async (input, { apiKey }) => {
    return requiemsRequest("/v1/technology/convert", { apiKey, query: input });
  },
};
```

Generation is deterministic and idempotent: tools are sorted alphabetically,
`generated/tools/` is wiped and rewritten on every run, and the whole directory
is disposable — the spec is the source of truth, not the generated code. The one
thing codegen doesn't touch is `src/server.ts` and `generated/runtime.ts`, which
are hand-maintained.

That determinism came with an explicit scope cut: `zodForParamSchema()` only
handles flat objects and primitive types cleanly. Anything using `oneOf`,
`anyOf`, or deeply nested objects falls back to `z.any()` / `z.record(z.any())`
instead of failing the whole generation run. Batch endpoints are excluded
entirely for the same reason. We'd rather ship 87 well-typed tools today than
block on perfectly typing the long tail of the spec.

## One codebase, two very different trust models

The generated tool code needs to run in two contexts that don't share much
except the handler logic:

- **stdio**, for a local MCP client (Claude Desktop, Cursor, VS Code) — one
  process, one user, one static API key set in the environment.
- **Streamable HTTP**, for `mcp.requiems.xyz` — a single public process serving
  many different callers, each with their own API key.

The stdio case is the easy one: read `REQUIEMS_API_KEY` from the environment
once at startup and thread it through. The HTTP case is the interesting one,
because the _same_ generated handler code needs the caller's key, and there's no
per-process global to put it in — the process is shared across every request
from every user.

We solved this with `AsyncLocalStorage`:

```typescript
// apps/mcp/generated/runtime.ts (simplified)
export const apiKeyContext = new AsyncLocalStorage<{ apiKey: string }>();

export function requiemsRequest(path: string, init: RequestInit = {}) {
  const { apiKey } = apiKeyContext.getStore()!;
  return fetch(`${API_BASE}${path}`, {
    ...init,
    headers: { ...init.headers, "requiems-api-key": apiKey },
  });
}
```

Every generated handler calls `requiemsRequest` without knowing or caring
whether it's running under stdio or HTTP. In HTTP mode, each request reads the
`requiems-api-key` header, runs the whole request inside
`apiKeyContext.run({ apiKey }, ...)`, and the correct key falls out the other
end automatically — no key ever needs to be passed explicitly through a call
chain that codegen doesn't fully control.

## Stateless by design, not by accident

The MCP SDK supports session-based Streamable HTTP transports, where a client
opens a session and reuses it across requests. We didn't use that. Instead,
`mcp.requiems.xyz` creates a **fresh `McpServer` and transport instance per HTTP
request** (`sessionIdGenerator: undefined`), and tears it down in a `finally`
block once the request completes.

This was a deliberate trade against session reuse: no session-affinity
requirements, nothing to garbage-collect server-side if a client disconnects
mid-session, and the whole endpoint stays horizontally scalable behind a plain
reverse proxy. The cost is a small amount of per-request setup overhead, which
is cheap compared to the request itself hitting our API.

## Two bugs worth admitting to

**Silently dropped results.** Early on, tool calls against the HTTP endpoint
would succeed with a 200 but return nothing useful to the model. The generated
handlers were returning raw API JSON straight from `requiemsRequest`, not the
`{ content: [...] }` shape MCP's `CallToolResultSchema` expects — and the SDK
doesn't throw when the shape is wrong, it just defaults `content` to `[]`. Every
tool call "worked" and returned nothing. The fix was wrapping every handler's
return value:

```typescript
return { content: [{ type: "text", text: JSON.stringify(result) }] };
```

**Missing keys failing open.** The first version of the HTTP transport didn't
reject a request that was missing the `requiems-api-key` header — it fell
through with an empty string, which then failed downstream in a way that was
harder to debug than a clean 401 at the door. We added an explicit check to
reject the request immediately if the header is absent, plus a top-level error
handler on the HTTP server so a thrown exception returns a proper JSON-RPC error
object instead of a bare 500.

Both are the kind of bug that only shows up once real traffic — not your own
manual testing — starts hitting the endpoint, which is exactly why we're writing
this down.

## Where it actually runs

Most of Requiems API is built on Cloudflare Workers. The MCP server isn't — it
runs as a plain **Bun** process in a **Docker** container, reverse-proxied by
**Caddy** at `mcp.requiems.xyz`. That's a deliberate choice, not an oversight:
the stateless-per-request model above doesn't need Workers' edge distribution to
be fast, Bun gives us native `fetch`/`AbortController` without extra
dependencies, and a single always-on container is simpler to reason about for a
process that's mostly proxying to our own API anyway. Authentication for the
tools themselves happens per-request via the `requiems-api-key` header, so the
container itself doesn't need to gate anything at the network layer.

## What's next

87 tools across categories like validation, finance, networking, and identity
risk are live today at `mcp.requiems.xyz` — point Claude Desktop, Cursor, or any
MCP-compatible client at it and the whole Requiems API surface shows up as
callable tools. Batch endpoints and deeply nested schemas are next on the list
now that the core generation pipeline has proven itself in production.

If you build something with it, we'd like to hear about it.
