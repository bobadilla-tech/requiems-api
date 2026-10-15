# frozen_string_literal: true

class Dashboard::OverviewController < Dashboard::BaseController
  def index
    @usage = current_user.usage_summary
    @current_plan = @usage.plan_name
    @subscription = @usage.cycle.subscription
    @recent_activity = current_user.recent_activity
    @api_keys_count = current_user.api_keys.active_keys.count
    @next_billing_date = @subscription&.current_period_end
  end
end
