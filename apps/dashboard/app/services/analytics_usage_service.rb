# frozen_string_literal: true

# Computes all usage-related analytics for the admin dashboard.
#
# Usage:
#   data = AnalyticsUsageService.new(date_range: "30").call
#   data.total_requests          # => 12_450
#   data.requests_by_day         # => { "2026-02-14" => 312, ... }
#   data.top_users_by_usage      # => ActiveRecord::Relation
#
class AnalyticsUsageService
  Result = Data.define(
    :total_requests,
    :requests_by_day,
    :requests_by_endpoint,
    :requests_by_plan,
    :avg_response_by_endpoint,
    :top_users_by_usage
  )

  def initialize(date_range: "30")
    @start_date = date_range.to_i.days.ago.beginning_of_day
    @end_date   = Time.current.end_of_day
  end

  def call
    Result.new(
      total_requests:           total_requests,
      requests_by_day:          requests_by_day,
      requests_by_endpoint:     requests_by_endpoint,
      requests_by_plan:         requests_by_plan,
      avg_response_by_endpoint: avg_response_by_endpoint,
      top_users_by_usage:       top_users_by_usage
    )
  end

  private

  def in_period
    UsageLog.where(used_at: @start_date..@end_date)
  end

  def total_requests
    in_period.count
  end

  def requests_by_day
    in_period
      .group("DATE(used_at)")
      .count
      .transform_keys { |date| date.to_s }
  end

  def requests_by_endpoint
    in_period
      .group(:endpoint)
      .count
      .sort_by { |_, count| -count }
      .first(10)
      .to_h
  end

  def requests_by_plan
    paid = in_period
      .joins(user: :subscription)
      .group("subscriptions.plan_name")
      .count

    free_count = in_period
      .joins(:user)
      .joins("LEFT OUTER JOIN subscriptions ON subscriptions.user_id = users.id")
      .where("subscriptions.id IS NULL OR subscriptions.plan_name = 'free'")
      .count

    paid["free"] = free_count if free_count > 0
    paid
  end

  def avg_response_by_endpoint
    in_period
      .where.not(response_time_ms: nil)
      .group(:endpoint)
      .average(:response_time_ms)
      .sort_by { |_, avg| -avg }
      .first(10)
      .to_h
      .transform_values { |v| v.round(2) }
  end

  def top_users_by_usage
    in_period
      .joins(:user)
      .group("users.id", "users.email")
      .select("users.id, users.email, COUNT(*) as request_count, SUM(credits_used) as total_requests")
      .order("request_count DESC")
      .limit(10)
  end
end
