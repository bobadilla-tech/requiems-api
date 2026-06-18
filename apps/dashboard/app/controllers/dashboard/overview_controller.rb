# frozen_string_literal: true

class Dashboard::OverviewController < Dashboard::BaseController

  def index
    @current_plan = current_user.current_plan
    @usage_this_month = current_user.usage_this_month
    @requests_remaining = current_user.requests_remaining
    @plan_limit = PlanConfig.requests_per_month(@current_plan)
    @recent_activity = current_user.recent_activity
    @api_keys_count = current_user.api_keys.active_keys.count
    @next_billing_date = current_user.subscription&.current_period_end
    @usage_percentage = calculate_usage_percentage
    @bar_color = calculate_bar_color
  end

  private

  def calculate_usage_this_month
    # Get usage from current month
    start_of_month = Time.current.beginning_of_month

    current_user.usage_logs
      .where("used_at >= ?", start_of_month)
      .count
  end

  def calculate_total_requests
    current_user.usage_logs.count
  end

  def calculate_requests_remaining
    limit = PlanConfig.requests_per_month(@current_plan)
    [ limit - @usage_this_month, 0 ].max
  end

  def calculate_avg_response_time
    # Calculate average response time from recent requests
    recent_logs = current_user.usage_logs
      .where("used_at >= ?", 7.days.ago)
      .where.not(response_time_ms: nil)

    return 0 if recent_logs.empty?

    (recent_logs.average(:response_time_ms) || 0).round
  end

  def fetch_recent_activity
    current_user.usage_logs
      .order(used_at: :desc)
      .limit(10)
      .includes(:api_key)
  end

  def calculate_usage_percentage
    total_limit = @usage_this_month + @requests_remaining
    return 0 if total_limit <= 0

    ((@usage_this_month.to_f / total_limit) * 100).round
  end

  def calculate_bar_color
    if @usage_percentage >= 90
      "bg-red-500"
    elsif @usage_percentage >= 75
      "bg-yellow-500"
    else
      "bg-blue-500"
    end
  end
end
