# Infrastructure

The production path is:

`Client -> Cloudflare proxy/WAF/DDoS/TLS/AOP -> Caddy -> Kamal -> Go`.

Cloudflare is DNS and security infrastructure only. The VPS runs the
Kamal-managed Go API, Rails dashboard/Sidekiq, MCP server, Postgres, Redis,
Caddy, and LanguageTool.

For local development use:

```bash
cd infra/docker
docker compose -f docker-compose.dev.yml up
```

The direct API is `http://localhost:8080`; the local Rails dashboard is
`http://localhost:3000`. Set `LOCAL_DEV_API_KEY` in `.env.local`.

