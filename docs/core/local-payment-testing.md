# Local payment testing

Use the Rails test environment and an isolated database. Payment webhook tests
exercise subscriptions and plans locally; they do not synchronize API keys
to Cloudflare.

```bash
cd apps/dashboard
RAILS_ENV=test bin/rails db:test:prepare
RAILS_ENV=test bin/rails test test/controllers/webhooks
```

For a local browser flow, set the Lemon Squeezy test variables in
`infra/docker/.env.local`, run the dashboard, and use the test checkout
configuration. The Go API reads the resulting Rails/Postgres subscription and
API-key state directly.

