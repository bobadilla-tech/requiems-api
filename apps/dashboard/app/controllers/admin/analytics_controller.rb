# frozen_string_literal: true

class Admin::AnalyticsController < Admin::BaseController

  def usage
    @date_range = params[:date_range] || "30"
    @start_date = @date_range.to_i.days.ago.beginning_of_day
    @end_date = Time.current.end_of_day

    # Total requests in period
    @total_requests = UsageLog.where(used_at: @start_date..@end_date).count

    # Requests per day (for chart)
    @requests_by_day = UsageLog
      .where(used_at: @start_date..@end_date)
      .group(Arel.sql("DATE(used_at)"))
      .count
      .transform_keys { |date| date.to_s }

    # Requests by endpoint (top 10)
    @requests_by_endpoint = UsageLog
      .where(used_at: @start_date..@end_date)
      .group(:endpoint)
      .order(Arel.sql("COUNT(*) DESC"))
      .limit(10)
      .count

    # Requests by plan (single query using CASE WHEN to bucket free/no-subscription users)
    @requests_by_plan = UsageLog
      .where(used_at: @start_date..@end_date)
      .joins(:user)
      .joins("LEFT OUTER JOIN subscriptions ON subscriptions.user_id = users.id")
      .group(Arel.sql("CASE WHEN subscriptions.id IS NULL OR subscriptions.plan_name = 'free' THEN 'free' ELSE subscriptions.plan_name END"))
      .count

    # Average response times by endpoint (top 10)
    @avg_response_by_endpoint = UsageLog
      .where(used_at: @start_date..@end_date)
      .where.not(response_time_ms: nil)
      .group(:endpoint)
      .order(Arel.sql("AVG(response_time_ms) DESC"))
      .limit(10)
      .average(:response_time_ms)
      .transform_values { |v| v.round(2) }

    # Top users by usage
    @top_users_by_usage = UsageLog
      .where(used_at: @start_date..@end_date)
      .joins(:user)
      .group("users.id", "users.email")
      .select("users.id, users.email, COUNT(*) as request_count, SUM(credits_used) as total_requests")
      .order(Arel.sql("request_count DESC"))
      .limit(10)
  end

  def revenue
    data = AnalyticsRevenueService.new.call

    @mrr                  = data.mrr
    @arr                  = data.arr
    @revenue_by_plan      = data.revenue_by_plan
    @revenue_trend        = data.revenue_trend
    @active_subscriptions = data.active_subscriptions
    @subscriptions_by_plan = data.subscriptions_by_plan
    @churn_rate           = data.churn_rate
    @new_vs_canceled      = data.new_vs_canceled
  end

  def system_health
    # Time range for health metrics
    @time_range = params[:time_range] || "24h"
    @start_time = case @time_range
    when "1h" then 1.hour.ago
    when "24h" then 24.hours.ago
    when "7d" then 7.days.ago
    when "30d" then 30.days.ago
    else 24.hours.ago
    end

    # API Uptime (percentage of successful requests — single query)
    total_requests, successful_requests = UsageLog
      .where(used_at: @start_time..Time.current)
      .pick(
        Arel.sql("COUNT(*)"),
        Arel.sql("COUNT(*) FILTER (WHERE status_code BETWEEN 200 AND 299)")
      )
    total_requests = total_requests.to_i
    successful_requests = successful_requests.to_i
    @uptime_percentage = total_requests > 0 ? ((successful_requests.to_f / total_requests) * 100).round(2) : 100.0

    # Average response times (P50, P95, P99) computed in the database
    p50, p95, p99 = UsageLog
      .where(used_at: @start_time..Time.current)
      .where.not(response_time_ms: nil)
      .pick(
        Arel.sql("percentile_cont(0.50) WITHIN GROUP (ORDER BY response_time_ms)"),
        Arel.sql("percentile_cont(0.95) WITHIN GROUP (ORDER BY response_time_ms)"),
        Arel.sql("percentile_cont(0.99) WITHIN GROUP (ORDER BY response_time_ms)")
      )

    if p50
      @p50_response_time = p50.round(2)
      @p95_response_time = p95.round(2)
      @p99_response_time = p99.round(2)
    else
      @p50_response_time = @p95_response_time = @p99_response_time = 0
    end

    @error_rate_trend = build_error_rate_trend(@start_time, @time_range)

    # Rate limit hits (last 24h)
    # Note: This would need to be tracked in the database
    # For now, we'll show 0 as a placeholder
    @rate_limit_hits = 0

    # Failed authentication attempts (last 24h)
    @failed_auth_attempts = UsageLog
      .where(used_at: 24.hours.ago..Time.current)
      .where(status_code: 401)
      .count

    # Most common errors (top 10)
    @common_errors = UsageLog
      .where(used_at: @start_time..Time.current)
      .where("status_code >= ?", 400)
      .group(:status_code)
      .count
      .sort_by { |_, count| -count }
      .first(10)
      .to_h

    # Requests per minute (last hour) for real-time monitoring
    if @time_range == "1h"
      @requests_per_minute = UsageLog
        .where(used_at: 1.hour.ago..Time.current)
        .group(Arel.sql("DATE_TRUNC('minute', used_at)"))
        .count
        .transform_keys { |time| time.strftime("%H:%M") }
    end
  end

  private

  def build_error_rate_trend(start_time, time_range)
    trunc_unit = case time_range
    when "1h"  then "minute"
    when "24h" then "hour"
    else             "day"
    end

    rows = UsageLog
      .where(used_at: start_time..Time.current)
      .group(Arel.sql("DATE_TRUNC('#{trunc_unit}', used_at)"))
      .select(
        Arel.sql("DATE_TRUNC('#{trunc_unit}', used_at) AS bucket"),
        Arel.sql("COUNT(*) AS total"),
        Arel.sql("COUNT(*) FILTER (WHERE status_code >= 400) AS error_count")
      )

    rows.each_with_object({}) do |row, hash|
      label = case time_range
      when "1h"  then row.bucket.strftime("%H:%M")
      when "24h" then row.bucket.strftime("%H:00")
      else            row.bucket.strftime("%Y-%m-%d")
      end

      rate = row.total > 0 ? ((row.error_count.to_f / row.total) * 100).round(2) : 0.0
      hash[label] = rate
    end
  end

end
