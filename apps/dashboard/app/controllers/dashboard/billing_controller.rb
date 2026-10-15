# frozen_string_literal: true

class Dashboard::BillingController < Dashboard::BaseController
  helper_method :valid_plan?

  def show
    # The subscription the Go API enforces (active/trialing), so the plan,
    # dates and usage numbers on this page all describe the same row.
    @usage = current_user.usage_summary
    @subscription = @usage.cycle.subscription
    @current_plan = @usage.plan_name
    # A paid row that is not active (e.g. past_due) is enforced as free, but its
    # owner still needs the customer portal to fix the payment.
    legacy = current_user.subscription
    @portal_for_inactive_subscription = @current_plan == "free" && legacy&.lemonsqueezy_subscription_id.present? && !legacy.promoted?

    # Get plan limits and pricing
    @plan_info = get_plan_info(@current_plan)
    @available_plans = get_all_plans

    # Handle success/cancel callbacks from LemonSqueezy
    if params[:success] == "true"
      flash.now[:notice] = t("dashboard.billing.flash.activated")
    elsif params[:canceled] == "true"
      flash.now[:alert] = t("dashboard.billing.flash.canceled")
    end
  end

  def checkout
    # Create LemonSqueezy checkout session for plan upgrade
    plan = params[:plan]

    unless valid_plan?(plan)
      redirect_to dashboard_billing_path, alert: t("dashboard.billing.flash.invalid_plan")
      return
    end

    billing_cycle = params[:billing_cycle] || "monthly"
    checkout_url = create_lemonsqueezy_checkout_url(plan, billing_cycle)

    redirect_to checkout_url, allow_other_host: true
  end

  def portal
    # Redirect to LemonSqueezy customer portal
    if current_user.subscription&.promoted?
      redirect_to dashboard_billing_path, alert: t("dashboard.billing.flash.promotional_no_portal")
      return
    end

    unless current_user.subscription&.lemonsqueezy_subscription_id
      redirect_to dashboard_billing_path, alert: t("dashboard.billing.flash.no_subscription")
      return
    end

    # LemonSqueezy customer portal URL
    # Format: https://[store-slug].lemonsqueezy.com/billing
    store_slug = AppConfig.lemonsqueezy_store_slug
    portal_url = "https://#{store_slug}.lemonsqueezy.com/billing"

    redirect_to portal_url, allow_other_host: true
  end

  def cancel_subscription
    unless current_user.subscription
      redirect_to dashboard_billing_path, alert: t("dashboard.billing.flash.no_subscription_to_cancel")
      return
    end

    # For LemonSqueezy, we'll direct users to the customer portal
    # where they can manage their subscription
    redirect_to portal_dashboard_billing_path
  end

  private

  def get_plan_info(plan_name)
    PlanConfig.for(plan_name)
  end

  def get_all_plans
    PlanConfig.all
  end

  def valid_plan?(plan)
    PlanConfig.paid?(plan)
  end

  def create_lemonsqueezy_checkout_url(plan, billing_cycle)
    checkout_uuid = AppConfig.checkout_uuid_for(plan: plan, billing_cycle: billing_cycle)

    store_slug = AppConfig.lemonsqueezy_store_slug
    checkout_url = "https://#{store_slug}.lemonsqueezy.com/checkout/buy/#{checkout_uuid}"

    # Add query parameters
    params = {
      "checkout[email]" => current_user.email,
      "checkout[custom][user_id]" => current_user.id,
      "checkout[custom][plan]" => plan
    }

    referral = current_user.referral_received
    params["checkout[custom][referred_by_code]"] = referral.referrer.referral_code if referral

    "#{checkout_url}?#{params.to_query}"
  end
end
