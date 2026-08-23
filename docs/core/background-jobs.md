# Background jobs

The remaining scheduled work is Rails/Postgres work:

- `AggregateDailyUsageJob` aggregates Postgres `usage_logs`.
- `ExpirePromotionalSubscriptionsJob` expires promotional plans.
- `RefreshSitemapJob` refreshes public sitemap files.
- Solid Queue cleanup removes finished framework jobs.

The retired `SyncD1UsageJob` and all D1 export/sync behavior have been removed.
Do not add a D1 reconciliation or billing backfill.

## Inspect production

```bash
docker ps
docker logs <dashboard-job-container>
docker exec <redis-container> redis-cli
```

Before operational retirement, inspect Sidekiq queued, scheduled, retry, dead,
and running/leased state plus Solid Queue jobs. The D1 sync class must be absent
from all sets. Do not change the aggregation job; it reads the Postgres ledger.

## Troubleshooting

For usage problems, inspect Go request logs, Postgres `usage_logs`, and the
aggregation job. A missing `api-management.requiems.xyz` record is expected; the
hostname is retired.
