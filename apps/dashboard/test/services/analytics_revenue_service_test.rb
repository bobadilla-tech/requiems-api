# frozen_string_literal: true

require "test_helper"

class AnalyticsRevenueServiceTest < ActiveSupport::TestCase
  test "calculate_mrr prices a yearly subscription at its monthly-equivalent rate, not the monthly price" do
    monthly_user = create_user(email: "monthly@example.com")
    Subscription.create!(user: monthly_user, plan_name: "developer", plan: "monthly", status: "active")

    yearly_user = create_user(email: "yearly@example.com")
    Subscription.create!(user: yearly_user, plan_name: "developer", plan: "yearly", status: "active")

    mrr = AnalyticsRevenueService.new.call.mrr

    expected = PlanConfig::PLANS["developer"][:price_monthly] + PlanConfig.price_yearly_monthly("developer")
    assert_equal expected, mrr
    assert_not_equal PlanConfig::PLANS["developer"][:price_monthly] * 2, mrr,
      "a yearly subscriber must not be priced as if billed monthly"
  end

  test "calculate_mrr defaults a subscription with no plan (billing cycle) to monthly pricing" do
    user = create_user(email: "nocycle@example.com")
    Subscription.create!(user: user, plan_name: "developer", plan: nil, status: "active")

    mrr = AnalyticsRevenueService.new.call.mrr

    assert_equal PlanConfig::PLANS["developer"][:price_monthly], mrr
  end
end
