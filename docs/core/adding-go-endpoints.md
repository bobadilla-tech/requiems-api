# Adding Go endpoints

1. Read [backend patterns](backend.md) and the service's existing `service.go`,
   `transport_http.go`, and route registration.
2. Add the handler and service under the appropriate `apps/api/services` domain.
3. Register the route under the Go router and keep the `requiems-api-key`
   middleware boundary intact.
4. Add unit/integration tests using the dedicated test database. Never use
   development or production data.
5. Update the generated OpenAPI source and regenerate MCP tools with
   `cd apps/mcp && bun run generate`.
6. Update the dashboard API documentation snippets if the public contract
   changed.
7. Run `go test ./...`, MCP tests/typecheck, Rails checks when affected, and the
   integration smoke checks.

The public flow is Cloudflare -> Caddy/AOP -> Kamal -> Go. No Worker, KV, D1,
API-management sync, or proxy secret is part of a new endpoint.
