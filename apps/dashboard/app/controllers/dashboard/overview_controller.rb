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
