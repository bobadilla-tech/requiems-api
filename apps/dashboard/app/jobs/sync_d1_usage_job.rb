# frozen_string_literal: true

# Background job to sync usage data from Cloudflare D1 to PostgreSQL
#
# This job is scheduled via Sidekiq Cron (see config/sidekiq_schedule.yml).
# It fetches new usage records from D1 and inserts them into the usage_logs table.
#
# Usage:
#   SyncD1UsageJob.perform_later
#
class SyncD1UsageJob < ApplicationJob
  queue_as :default

  retry_on D1SyncService::Error, attempts: 5, wait: :polynomially_longer do |job, error|
    Rails.logger.error("SyncD1UsageJob permanently failed: #{error.message} (job_id=#{job.job_id})")
  end

  LAST_SYNC_KEY = "d1_sync:last_sync_at"
  SYNC_INTERVAL = 5.minutes

  def perform
    start_time = Time.current
    last_sync_at = get_last_sync_time

    Rails.logger.info("SyncD1UsageJob: starting since=#{last_sync_at.iso8601} (#{(start_time - last_sync_at).round}s window)")

    service = D1SyncService.new
    result = service.fetch_usage(since: last_sync_at)

    if result[:usage].empty?
      window_seconds = (start_time - last_sync_at).round
      if window_seconds > SYNC_INTERVAL * 2
        Rails.logger.warn(
          "D1 sync returned 0 records over a #{window_seconds}s window " \
          "(since: #{last_sync_at.iso8601}). Possible API outage or misconfiguration."
        )
      else
        Rails.logger.info("No new usage records to sync (since: #{last_sync_at.iso8601})")
      end
      return
    end

    inserted = service.bulk_insert(result[:usage])

    # Update last sync timestamp
    set_last_sync_time(start_time)

    Rails.logger.info("D1 sync completed: #{inserted} records inserted")

    # Track metrics
    track_sync_metrics(inserted, start_time)
  rescue D1SyncService::Error => e
    Rails.logger.error("D1 sync failed: #{e.message}")
    # Don't update last_sync_time on failure - will retry from same point
    raise
  end

  private

  def get_last_sync_time
    timestamp = Rails.cache.read(LAST_SYNC_KEY)
    return Time.parse(timestamp) if timestamp

    # Fallback: get timestamp of most recent usage log
    last_log = UsageLog.order(used_at: :desc).first
    last_log ? last_log.used_at : 1.hour.ago
  end

  def set_last_sync_time(time)
    Rails.cache.write(LAST_SYNC_KEY, time.iso8601, expires_in: 7.days)
  end

  def track_sync_metrics(inserted, start_time)
    duration = (Time.current - start_time).round(2)
    rps = duration > 0 ? (inserted.to_f / duration).round(2) : 0
    Rails.logger.info("D1 sync metrics: records_inserted=#{inserted} duration_seconds=#{duration} records_per_second=#{rps}")
  end
end
