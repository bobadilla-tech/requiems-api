# frozen_string_literal: true

class Dashboard::BillingController < ApplicationController
  before_action :authenticate_user!
  layout "dashboard"

  helper_method :valid_plan?

  def show
    @subscription = current_user.subscription
    @current_plan = @subscription&.plan_name || "free"

    # Calculate usage for current billing period
    if @subscription
      period_start = @subscription.current_period_start || Time.current.beginning_of_month
      @usage_this_period = current_user.usage_logs
        .where("used_at >= ?", period_start)
        .count
    else
      @usage_this_period = current_user.usage_logs
        .where("used_at >= ?", Time.current.beginning_of_month)
        .count
    end

    # Get plan limits and pricing
    @plan_info = get_plan_info(@current_plan)
    @available_plans = get_all_plans

    # Calculate quota usage percentage
    quota_limit = @plan_info[:requests_per_month]
    @quota_used_percent = quota_limit > 0 ? ((@usage_this_period.to_f / quota_limit) * 100).round(1) : 0

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
