# Docker development

```bash
docker compose -f docker-compose.dev.yml config --quiet
docker compose -f docker-compose.dev.yml up -d api db redis dashboard sidekiq mcp
```

The Compose stack contains Go, Rails, Sidekiq, MCP, Postgres, Redis, and the
optional LanguageTool dependency. Dashboard, Sidekiq, and MCP call the direct
Go service at `api:8080`. There are no Worker containers, Wrangler state
volumes, D1 migrations, or KV mounts.

Set `LOCAL_DEV_API_KEY` in `.env.local` before running the dashboard seed.
Use `TEST_DATABASE_URL` for Go integration tests and never run validation
against the development database.

