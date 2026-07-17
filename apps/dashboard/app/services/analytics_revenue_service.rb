# frozen_string_literal: true

# Computes all revenue-related analytics for the admin dashboard.
#
# Usage:
#   data = AnalyticsRevenueService.new.call
#   data.mrr           # => 2450.0
#   data.arr           # => 29400.0
#   data.churn_rate    # => 2.5
#
class AnalyticsRevenueService
  PLAN_PRICES = PlanConfig::PAID_PLAN_NAMES.index_with do |plan|
    { monthly: PlanConfig::PLANS[plan][:price_monthly], yearly: PlanConfig.price_yearly_monthly(plan) }
  end.freeze

  Result = Data.define(
    :mrr,
    :arr,
    :revenue_by_plan,
    :revenue_trend,
    :active_subscriptions,
    :subscriptions_by_plan,
    :churn_rate,
    :new_vs_canceled
  )

  def call
    mrr = calculate_mrr

    Result.new(
      mrr:                  mrr,
      arr:                  mrr * 12,
      revenue_by_plan:      revenue_by_plan,
      revenue_trend:        revenue_trend,
      active_subscriptions: active_subscriptions,
      subscriptions_by_plan: subscriptions_by_plan,
      churn_rate:           churn_rate,
      new_vs_canceled:      new_vs_canceled
    )
  end

  private

  def calculate_mrr
    Subscription
      .paying
      .sum do |sub|
        price = PLAN_PRICES[sub.plan_name]&.fetch(sub.plan&.to_sym || :monthly, 0) || 0
        sub.plan == "yearly" ? (price * 12 / 12.0) : price
      end
  end

  def revenue_by_plan
    Subscription
      .paying
      .group(:plan_name, :plan)
      .count
      .each_with_object({}) do |((plan, cycle), count), hash|
        key = "#{plan.titleize} (#{cycle&.titleize || 'Monthly'})"
        price = PLAN_PRICES[plan]&.fetch(cycle&.to_sym || :monthly, 0) || 0
        hash[key] = price * count
      end
  end

  def revenue_trend
    trend = {}
    (0..11).each do |i|
      month_start = i.months.ago.beginning_of_month
      month_end   = i.months.ago.end_of_month
      month_label = month_start.strftime("%b %Y")

      month_revenue = Subscription
        .paid
        .where("current_period_start <= ?", month_end)
        .where("current_period_end IS NULL OR current_period_end >= ?", month_start)
        .sum do |sub|
          PLAN_PRICES[sub.plan_name]&.fetch(sub.plan&.to_sym || :monthly, 0) || 0
        end

      trend[month_label] = month_revenue
    end
    trend.reverse_each.to_h
  end

  def active_subscriptions
    Subscription.paying.count
  end

  def subscriptions_by_plan
    Subscription
      .paid
      .group(:plan_name)
      .count
      .transform_keys(&:titleize)
  end

  def churn_rate
    start_of_period = 30.days.ago

    subscriptions_at_start = Subscription
      .paid
      .where("created_at < ?", start_of_period)
      .count

    canceled_in_period = Subscription
      .paid
      .where(cancel_at_period_end: true)
      .where("canceled_at >= ?", start_of_period)
      .count

    return 0 if subscriptions_at_start == 0

    ((canceled_in_period.to_f / subscriptions_at_start) * 100).round(2)
  end

  def new_vs_canceled
    result = {}
    (0..11).each do |i|
      month_start = i.months.ago.beginning_of_month
      month_end   = i.months.ago.end_of_month
      month_label = month_start.strftime("%b %Y")

      new_subs = Subscription
        .paid
        .where(created_at: month_start..month_end)
        .count

      canceled_subs = Subscription
        .paid
        .where(cancel_at_period_end: true)
        .where(canceled_at: month_start..month_end)
        .count

      result[month_label] = { new: new_subs, canceled: canceled_subs }
    end
    result.reverse_each.to_h
  end
end
