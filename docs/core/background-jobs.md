# Background Jobs

The Rails dashboard runs background jobs via **Sidekiq** with **Sidekiq-Cron**
for scheduled tasks. Redis is the job queue backend (the same Redis instance
used for caching).

## Jobs

Cron entries (schedule, queue, descriptions) live in
**[`config/sidekiq_schedule.yml`](../../apps/dashboard/config/sidekiq_schedule.yml)**.

Implementations (behavior, logging, `retry_on`, and idempotency notes):

- [`app/jobs/sync_d1_usage_job.rb`](../../apps/dashboard/app/jobs/sync_d1_usage_job.rb)
- [`app/jobs/aggregate_daily_usage_job.rb`](../../apps/dashboard/app/jobs/aggregate_daily_usage_job.rb)
- [`app/jobs/expire_promotional_subscriptions_job.rb`](../../apps/dashboard/app/jobs/expire_promotional_subscriptions_job.rb)

Base class:
[`app/jobs/application_job.rb`](../../apps/dashboard/app/jobs/application_job.rb)

## Checking Status in Production

### Web UI (easiest)

The Sidekiq dashboard is mounted at `/sidekiq` — admin users only.

It shows:

- Live queue depth (pending, processing, scheduled)
- Retry queue with stack traces
- Dead letter queue (jobs that exhausted retries)
- Cron jobs and their last/next run times

### Logs

```bash
# Tail the worker container logs
docker compose logs -f worker
```

### Rails console

```bash
docker exec -it requiem-prod-dashboard-1 rails console
```

```ruby
# Overall stats
Sidekiq::Stats.new.to_h
# => { enqueued: 0, processed: 1234, failed: 2, retry_size: 0, dead_size: 0, ... }

# Queue depth
Sidekiq::Queue.new("default").count

# Cron jobs and their next fire time
Sidekiq::Cron::Job.all.each { |j| puts "#{j.name}: last=#{j.last_time} next=#{j.next_time}" }

# Retry queue
Sidekiq::RetrySet.new.map { |j| [j.display_class, j.error_message, j.at] }

# Dead queue (exhausted retries)
Sidekiq::DeadSet.new.map { |j| [j.display_class, j.error_message] }
```

## Triggering Jobs Manually

```bash
docker exec -it requiem-prod-dashboard-1 rails console
```

```ruby
# Re-sync D1 usage immediately
SyncD1UsageJob.perform_later

# Aggregate a specific date (defaults to yesterday if omitted)
AggregateDailyUsageJob.perform_later(date: Date.yesterday)
AggregateDailyUsageJob.perform_later(date: "2026-04-15")

# Expire promotional subscriptions now
ExpirePromotionalSubscriptionsJob.perform_later
```

## Infrastructure

In production the worker runs as a separate Docker service:

```yaml
# infra/docker/docker-compose.yml
worker:
  build: ../../apps/dashboard
  command: bundle exec sidekiq
  depends_on: [db, redis]
```

Sidekiq connects to Redis via the `REDIS_URL` environment variable (configured
in `.env`).

## Retry behavior

`retry_on` and attempt counts are defined in each job file above (not duplicated
here). `ExpirePromotionalSubscriptionsJob` uses Sidekiq defaults unless
overridden in that class.

Jobs are idempotent — re-running them on the same data is safe. If a cron job
misses a tick (e.g., worker restart), it fires again on the next scheduled tick.

## Common Failure Scenarios

**SyncD1UsageJob stops inserting records**\
The job logs a warning if zero records come back for an extended window. Check:

1. Cloudflare D1 / API Management service health
2. `CLOUDFLARE_*` environment variables in the worker container
3. `/sidekiq` retry queue for stack traces

**AggregateDailyUsageJob failed for a date**\
Re-run manually with the specific date — the upsert logic makes it safe to run
multiple times.

**Promotions not expiring**\
Check the dead letter queue at `/sidekiq`. Each subscription is processed
independently so a single failure does not block the rest.
