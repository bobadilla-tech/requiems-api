# frozen_string_literal: true

class Admin::DashboardController < Admin::BaseController
  STATS_CACHE_TTL = 5.minutes

  def index
    stats = Rails.cache.fetch("admin/dashboard/stats", expires_in: STATS_CACHE_TTL) do
      # System Statistics
      total_users = User.count
      active_users = User.where("last_sign_in_at >= ?", 30.days.ago).count
      total_api_keys = ApiKey.count
      active_api_keys = ApiKey.active_keys.count

      # Usage Statistics (last 30 days)
      total_requests_30d = UsageLog.where("used_at >= ?", 30.days.ago).count
      total_requests_today = UsageLog.where("used_at >= ?", Time.current.beginning_of_day).count

      # Revenue Statistics (exclude admin-granted promotional upgrades)
      active_subscriptions = Subscription.paying.where(promoted_by_id: nil).count
      mrr = calculate_mrr

      # System Health
      avg_response_time = calculate_avg_response_time
      error_rate = calculate_error_rate

      {
        total_users: total_users,
        active_users: active_users,
        total_api_keys: total_api_keys,
        active_api_keys: active_api_keys,
        total_requests_30d: total_requests_30d,
        total_requests_today: total_requests_today,
        active_subscriptions: active_subscriptions,
        mrr: mrr,
        avg_response_time: avg_response_time,
        error_rate: error_rate
      }
    end

    @total_users = stats[:total_users]
    @active_users = stats[:active_users]
    @total_api_keys = stats[:total_api_keys]
    @active_api_keys = stats[:active_api_keys]
    @total_requests_30d = stats[:total_requests_30d]
    @total_requests_today = stats[:total_requests_today]
    @active_subscriptions = stats[:active_subscriptions]
    @mrr = stats[:mrr]
    @avg_response_time = stats[:avg_response_time]
    @error_rate = stats[:error_rate]

    # Recent Activity
    @recent_users = User.order(created_at: :desc).limit(5)
    @recent_api_keys = ApiKey.order(created_at: :desc).limit(5)
    @recent_subscriptions = Subscription.paid.where(promoted_by_id: nil).order(created_at: :desc).limit(5)

    # Chart data (for AJAX loading)
    # Will be loaded via separate endpoints
  end

  private

  def calculate_mrr
    # Calculate Monthly Recurring Revenue
    plan_prices = {
      "developer" => 30,
      "business" => 75,
      "professional" => 150
    }

    Subscription.paying
      .where(promoted_by_id: nil)
      .group(:plan_name)
      .count
      .sum { |plan, count| (plan_prices[plan] || 0) * count }
  end

  def calculate_avg_response_time
    avg = UsageLog.where("used_at >= ?", 7.days.ago)
      .where.not(response_time_ms: nil)
      .average(:response_time_ms)

    avg&.round || 0
  end

  def calculate_error_rate
    total = UsageLog.where("used_at >= ?", 7.days.ago).count
    return 0 if total.zero?

    errors = UsageLog.where("used_at >= ?", 7.days.ago)
      .where("status_code >= ?", 400)
      .count

    ((errors.to_f / total) * 100).round(2)
  end
end
