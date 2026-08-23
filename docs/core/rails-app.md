# Rails Dashboard

The Rails application owns user-facing state: users, API keys, subscriptions,
plans, billing webhooks, dashboard pages, and private deployment requests.

## API keys

`Dashboard::ApiKeysController` creates keys through the local
`ApiKeyGenerator`. The full value is shown once. Rails stores the key prefix
and a BCrypt hash; Go verifies the presented value against Postgres. Revocation
invalidates Go's Redis auth cache.

The dashboard never calls a separate API-management service. A failed create
should be investigated in the Rails request log and database validation path,
not by looking for a Worker or D1 endpoint.

## Playground

`ApiProxyService` uses `INTERNAL_API_URL` (production:
`http://requiems-api:8080`) and sends `requiems-api-key` from
`PLAYGROUND_API_KEY`. It does not send the retired `X-Backend-Secret`
header.

## Background work

Sidekiq remains for application jobs such as daily usage aggregation,
promotional subscription expiry, sitemap refresh, and mail. Solid Queue
remains for Rails framework jobs where configured. The former
`SyncD1UsageJob` and `D1SyncService` no longer exist.

## Local setup

Set `LOCAL_DEV_API_KEY` before the development seed creates the local key.
Use an isolated `TEST_DATABASE_URL` and isolated Redis for tests. Never point
test commands at production or the development database.

