# frozen_string_literal: true

class Dashboard::OverviewController < Dashboard::BaseController
  def index
    # A filter click is a Turbo frame request for the same URL; keep caches from
    # serving one body for the other (Back button, shared caches).
    response.headers["Vary"] = [ response.headers["Vary"], "Turbo-Frame" ].compact_blank.join(", ")
    @usage = current_user.usage_summary
    @current_plan = @usage.plan_name
    @subscription = @usage.cycle.subscription
    @activity_filter = UsageLog::ACTIVITY_FILTERS.include?(params[:activity]) ? params[:activity] : "all"
    @recent_activity = current_user.usage_logs.for_activity_filter(@activity_filter).recent(10)
    @api_keys_count = current_user.api_keys.active_keys.count
    @next_billing_date = @subscription&.current_period_end
  end
end
