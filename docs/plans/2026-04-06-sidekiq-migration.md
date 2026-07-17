# Sidekiq Migration Design Plan

Rails 8 ships with Solid Queue as the default background job backend. Solid
Queue stores jobs in PostgreSQL using a dedicated schema loaded via
`db:schema:load:queue`. This schema was never loaded in the production database,
which meant:

- `deliver_later` calls silently failed (no worker to process the queue)
- Recurring jobs (`SyncD1UsageJob`, `AggregateDailyUsageJob`,
  `ExpirePromotionalSubscriptionsJob`) never ran
- The `sidekiq` Docker service was already defined in `docker-compose.yml` but
  was running the Solid Queue CLI (`bin/jobs start`) — a name/implementation
  mismatch

Similarly, Solid Cable (ActionCable backed by PostgreSQL) was configured in
`cable.yml` but the cable schema tables were also absent in production.

## Current State (before migration)

| Component       | Configured backend | Reality in production         |
| --------------- | ------------------ | ----------------------------- |
| Background jobs | Solid Queue        | Schema missing — jobs lost    |
| ActionCable     | solid_cable        | Schema missing — cable broken |
| Recurring jobs  | `recurring.yml`    | Never executed                |
| Docker worker   | `bin/jobs start`   | Running Solid Queue CLI       |
| Redis           | Cache only         | Already deployed and healthy  |

## Decision

Migrate from Solid Queue + Solid Cable to **Sidekiq** (jobs) + **Redis**
(cable).

### Why Sidekiq over fixing Solid Queue

1. **Redis is already running** — used for Rails cache. No new infrastructure
   needed.
2. **Zero schema maintenance** — Sidekiq uses Redis, not PostgreSQL. The missing
   schema problem goes away entirely.
3. **Proven reliability** — Sidekiq is the industry standard, with years of
   production hardening.
4. **Better observability** — Sidekiq ships with a web UI (`/sidekiq`) that
   shows queue depth, failed jobs, retries, and cron schedule — Solid Queue has
   no UI.
5. **sidekiq-cron** replaces `recurring.yml` cleanly with the same schedule
   semantics.

### Tradeoffs

| Concern                          | Notes                                                              |
| -------------------------------- | ------------------------------------------------------------------ |
| Extra gem dependency             | `sidekiq` + `sidekiq-cron` — well-maintained, low risk             |
| Redis as single point of failure | Redis already carries cache; adding jobs doesn't increase exposure |
| Job loss on Redis restart        | Acceptable: jobs are idempotent; cron re-fires on next tick        |

## Changes Made

### Gems (`Gemfile`)

- Removed `gem "solid_queue"`
- Added `gem "sidekiq", "~> 7.0"`
- Added `gem "sidekiq-cron", "~> 1.12"`

### Queue adapter (`config/environments/production.rb`)

```ruby
config.active_job.queue_adapter = :sidekiq  # was :solid_queue
# removed: config.solid_queue.connects_to = ...
```

### Sidekiq initializer (`config/initializers/sidekiq.rb`) — new file

Wires Sidekiq to Redis and loads the cron schedule on server boot:

```ruby
redis_config = { url: ENV.fetch("REDIS_URL", "redis://localhost:6379") }

Sidekiq.configure_server do |config|
  config.redis = redis_config
  schedule_file = Rails.root.join("config/sidekiq_schedule.yml")
  Sidekiq::Cron::Job.load_from_hash(YAML.load_file(schedule_file)) if schedule_file.exist?
end

Sidekiq.configure_client do |config|
  config.redis = redis_config
end
```

### Cron schedule (`config/sidekiq_schedule.yml`) — new file

Replaces `config/recurring.yml` (Solid Queue format):

```yaml
sync_d1_usage:
  cron: "*/5 * * * *"
  class: SyncD1UsageJob
  queue: default
  description: "Sync usage data from Cloudflare D1 to PostgreSQL every 5 minutes"

aggregate_daily_usage:
  cron: "5 0 * * *"
  class: AggregateDailyUsageJob
  queue: default
  description: "Aggregate raw usage_logs into daily_usage_summaries at 00:05 UTC"

expire_promotional_subscriptions:
  cron: "30 * * * *"
  class: ExpirePromotionalSubscriptionsJob
  queue: default
  description: "Revert expired admin-granted promotional plan upgrades to free"
```

### ActionCable (`config/cable.yml`)

Production switched from `solid_cable` to the Redis adapter:

```yaml
production:
  adapter: redis
  url: <%= ENV.fetch("REDIS_URL", "redis://localhost:6379") %>
  channel_prefix: requiem_production
```

### Database config (`config/database.yml`)

Removed the `queue`, `cable`, and `cache` database entries from the production
section. Only `primary` remains. These entries existed solely for Solid Queue,
Solid Cable, and Solid Cache.

### Puma (`config/puma.rb`)

Removed: `plugin :solid_queue if ENV["SOLID_QUEUE_IN_PUMA"]`

This line would have launched Solid Queue inline with the web process — no
longer needed.

### Sidekiq Web UI (`config/routes.rb`)

Mounted outside the locale scope so the Rack app works correctly:

```ruby
require "sidekiq/web"
require "sidekiq/cron/web"

authenticate :user, ->(u) { u.admin? } do
  mount Sidekiq::Web => "/sidekiq"
end
```

Accessible at `/sidekiq` for admin users. Shows live queue stats, failed jobs,
retries, and the cron job schedule.

### Docker worker service (`infra/docker/docker-compose.yml`)

- Renamed service from `sidekiq` to `worker` (was running Solid Queue CLI)
- Command changed from `bundle exec ruby bin/jobs start` to
  `bundle exec sidekiq`
- Removed `db:schema:load:queue` from dashboard startup sequence

### Deleted files

| File                 | Reason                                                                  |
| -------------------- | ----------------------------------------------------------------------- |
| `config/queue.yml`   | Solid Queue worker/dispatcher config — no longer used                   |
| `bin/jobs`           | Solid Queue CLI entrypoint — replaced by `bundle exec sidekiq`          |
| `db/queue_schema.rb` | Solid Queue DB schema — Sidekiq uses Redis, not PostgreSQL              |
| `db/cable_schema.rb` | Solid Cable DB schema — cable now uses Redis adapter                    |
| `db/cache_schema.rb` | Solid Cache DB schema — cache still uses Redis via `config.cache_store` |

## Verification

1. `bundle install` — resolves `sidekiq` and `sidekiq-cron`
2. `bin/rails test` — all tests pass (no Solid Queue references in test code)
3. In Rails console: `Sidekiq::Cron::Job.all` → returns 3 jobs
4. `/sidekiq` in browser → Sidekiq Web UI accessible to admin users
5. Worker container starts with `bundle exec sidekiq` and connects to Redis

## Environment Variables

No new environment variables required. Sidekiq reads `REDIS_URL`, which was
already set for the cache store.
