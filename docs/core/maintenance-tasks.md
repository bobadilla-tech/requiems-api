# Maintenance tasks

- Run `rails usage:status` to inspect Postgres usage ledgers.
- Run `rails usage:aggregate_daily START_DATE=... END_DATE=...` for an
  intentional aggregation window.
- Review Sidekiq and Solid Queue health before changing schedules.
- Regenerate MCP tools after an OpenAPI change:
  `cd apps/mcp && bun run fetch-spec && bun run generate`.
- Keep Cloudflare proxying, WAF/DDoS controls, TLS, and AOP enabled.
- Keep test database and Redis guards in validation scripts.

Do not restore Worker deployments, D1 sync, KV key storage, or schema-wide
cleanup commands.

