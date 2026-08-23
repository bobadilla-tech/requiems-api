# frozen_string_literal: true

namespace :usage do
  desc "Aggregate daily usage summaries for a date range"
  task aggregate_daily: :environment do
    # Get date range from environment variables
    start_date = Date.parse(ENV.fetch("START_DATE", 30.days.ago.to_date.to_s))
    end_date = Date.parse(ENV.fetch("END_DATE", Date.yesterday.to_s))

    puts "Aggregating daily usage summaries"
    puts "Date range: #{start_date} to #{end_date}"
    puts ""

    total_summaries = 0

    (start_date..end_date).each do |date|
      print "Processing #{date}... "

      count = AggregateDailyUsageJob.new.perform(date: date)
      total_summaries += count

      puts "#{count} summaries created"
    end

    puts ""
    puts "✓ Daily aggregation completed!"
    puts "  - Days processed: #{(end_date - start_date).to_i + 1}"
    puts "  - Total summaries created: #{total_summaries}"
  end

  desc "Show PostgreSQL usage ledger status"
  task status: :environment do
    puts "PostgreSQL Usage Status"
    puts "=" * 60
    puts ""

    # Check usage_logs table
    usage_logs_count = UsageLog.count
    puts "Usage logs in PostgreSQL: #{usage_logs_count.to_s.reverse.gsub(/(\d{3})(?=\d)/, '\\1,').reverse}"

    if usage_logs_count > 0
      oldest = UsageLog.order(:used_at).first&.used_at
      newest = UsageLog.order(:used_at).last&.used_at
      puts "  Oldest record: #{oldest}"
      puts "  Newest record: #{newest}"
    end
    puts ""

    # Check daily summaries
    summaries_count = DailyUsageSummary.count
    puts "Daily summaries: #{summaries_count}"

    if summaries_count > 0
      oldest_date = DailyUsageSummary.order(:date).first&.date
      newest_date = DailyUsageSummary.order(:date).last&.date
      puts "  Oldest date: #{oldest_date}"
      puts "  Newest date: #{newest_date}"
    end
    puts ""

    puts "Recurring jobs configured:"
    puts "  - Daily aggregation: Every day at 00:05 UTC"
    puts "  - Promotional expiry: Every hour at minute 30"
  end
end
