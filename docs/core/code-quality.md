# Code quality

There is no single long guide: lint, tests, and security checks are documented
per stack. Use this page as an index.

## Canonical command reference

The repo root **[`agents.md`](../../agents.md)** (Cursor/Claude project rules)
has the full **Development Commands** tables for Go, Rails, Auth Gateway, API
Management, and what to run before push.

## Where each stack is covered

| Stack               | Tests & lint in docs                                                                                                                                                                                            |
| ------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **Go API**          | [Getting started](./getting-started.md) (Docker lint snippet), [Adding Go endpoints](./adding-go-endpoints.md) (pre-merge checklist includes `golangci-lint`), [Backend](./backend.md) for architecture context |
| **Rails dashboard** | [Rails app](./rails-app.md) (running tests and the app in Docker)                                                                                                                                               |
| **Auth Gateway**    | [Auth Gateway](./auth-gateway.md) (`pnpm` typecheck, Vitest, lint)                                                                                                                                              |
| **API Management**  | [API Management](./api-management.md) (same worker-style commands)                                                                                                                                              |

## Go API quick reminder

With the dev API container up (`requiem-dev-api-1`):

```bash
docker exec requiem-dev-api-1 sh -lc 'go fmt ./... && golangci-lint run'
docker exec requiem-dev-api-1 go test ./...
```

For more options (race, coverage, single package), see
[`agents.md`](../../agents.md).
