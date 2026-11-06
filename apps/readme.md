# Applications

- **api** — Go API. Owns API-key authentication, Redis rate limiting, quota,
  usage rows, and product endpoints.
- **dashboard** — Rails UI, users, API keys, subscriptions, plans, billing,
  admin, and jobs.
- **mcp** — Bun MCP server that calls the Go API.
- **clients** — OpenAPI-generated SDKs (TypeScript, C#, Python, Ruby, Go) and
  example apps, regenerated weekly.
- **skills** — AI agent skills (SKILL.md per endpoint) generated from the API
  docs YAML; npm package `@requiems/api-skills` and Claude Code plugin.
- **tutorials** — Self-contained example projects and tutorials built on the
  API. Contribute new ones via `apps/tutorials/contributing.md`.
- **api data/tools** — Go product data and generated MCP/OpenAPI artifacts.

The former Cloudflare Worker applications, KV namespace, and D1 usage ledger
are retired and are not part of local or production runtime.

