# Getting started

## Services

Start the isolated development stack:

```bash
cd infra/docker
docker compose -f docker-compose.dev.yml up
```

| Service         | URL                           |
| --------------- | ----------------------------- |
| Go API          | http://localhost:8080/healthz |
| Rails dashboard | http://localhost:3000         |
| MCP HTTP        | http://localhost:3100         |
| PostgreSQL      | localhost:5433                |
| Redis           | localhost:6379                |

The Compose default development key is:

```text
LOCAL_DEV_API_KEY=requiem_AAAAAAAAAAAAAAAAAAAAAAAA
```

You can override it with any disposable `requiem_<24 alphanumeric characters>`
value in `infra/docker/.env.local`. The development seed uses the configured
value and reconciles the test user's key on startup, so an old Docker database
volume does not retain an unusable key. Do not use a production key for
development.

## Direct API smoke test

```bash
curl --fail http://localhost:8080/healthz
curl -H "requiems-api-key: ${LOCAL_DEV_API_KEY:-requiem_AAAAAAAAAAAAAAAAAAAAAAAA}" \
  http://localhost:8080/v1/entertainment/advice
```

Go integration tests must use the dedicated `requiem_test` database through
`TEST_DATABASE_URL`. Rails tests use `RAILS_ENV=test` and an isolated Redis
database.
