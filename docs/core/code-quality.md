# Code quality

## Required checks

Run these checks against isolated test services:

```bash
docker compose -f infra/docker/docker-compose.dev.yml config --quiet
docker exec requiem-dev-api-1 go test -race ./...
docker exec requiem-dev-dashboard-1 bin/rails test
docker exec requiem-dev-dashboard-1 bundle exec brakeman --no-pager
docker exec requiem-dev-dashboard-1 bundle exec bundler-audit
docker exec requiem-dev-dashboard-1 bin/importmap audit
cd apps/mcp && bun test && bunx tsc --noEmit
```

Use `TEST_DATABASE_URL` for Go and `RAILS_ENV=test` for Rails. Capture
protected-table counts before and after validation.

CI analyzes Go, Ruby, and MCP JavaScript/TypeScript. There are no Worker
targets, flags, path filters, or reusable Worker workflows.

