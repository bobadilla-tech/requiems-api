# frozen_string_literal: true

class Admin::AnalyticsController < ApplicationController
  before_action :authenticate_user!
  before_action :require_admin!
  layout "admin"

  def usage
    @date_range = params[:date_range] || "30"

    data = AnalyticsUsageService.new(date_range: @date_range).call

    @total_requests           = data.total_requests
    @requests_by_day          = data.requests_by_day
    @requests_by_endpoint     = data.requests_by_endpoint
    @requests_by_plan         = data.requests_by_plan
    @avg_response_by_endpoint = data.avg_response_by_endpoint
    @top_users_by_usage       = data.top_users_by_usage
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

    # API Uptime (percentage of successful requests)
    total_requests = UsageLog.where(used_at: @start_time..Time.current).count
    successful_requests = UsageLog.where(used_at: @start_time..Time.current).where(status_code: 200..299).count
    @uptime_percentage = total_requests > 0 ? ((successful_requests.to_f / total_requests) * 100).round(2) : 100.0

    # Average response times (P50, P95, P99)
    response_times = UsageLog
      .where(used_at: @start_time..Time.current)
      .where.not(response_time_ms: nil)
      .order(:response_time_ms)
      .pluck(:response_time_ms)

    if response_times.any?
      @p50_response_time = percentile(response_times, 50).round(2)
      @p95_response_time = percentile(response_times, 95).round(2)
      @p99_response_time = percentile(response_times, 99).round(2)
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
        .group("DATE_TRUNC('minute', used_at)")
        .count
        .transform_keys { |time| time.strftime("%H:%M") }
    end
  end

  private

  def require_admin!
    unless current_user.admin?
      redirect_to root_path, alert: "Access denied. Admin privileges required."
    end
  end

  def percentile(sorted_array, percentile)
    return 0 if sorted_array.empty?

    index = (percentile / 100.0) * (sorted_array.length - 1)
    lower = sorted_array[index.floor]
    upper = sorted_array[index.ceil]

    lower + (upper - lower) * (index - index.floor)
  end
end
