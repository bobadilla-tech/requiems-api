# frozen_string_literal: true

require "test_helper"

class AnalyticsRevenueServiceTest < ActiveSupport::TestCase
  # Asserts on the delta calculate_mrr attributes to the subscriptions this
  # test creates, not the absolute total — the dev/test database can already
  # hold other paying subscriptions (seeded data, the playground system key's
  # "developer" subscription, etc.), so an absolute-total assertion is
  # inherently fragile to shared database state unrelated to this test.
  def mrr_delta
    before = AnalyticsRevenueService.new.call.mrr
    yield
    AnalyticsRevenueService.new.call.mrr - before
  end

  test "calculate_mrr prices a yearly subscription at its monthly-equivalent rate, not the monthly price" do
    delta = mrr_delta do
      monthly_user = create_user(email: "monthly@example.com")
      Subscription.create!(user: monthly_user, plan_name: "developer", plan: "monthly", status: "active")

      yearly_user = create_user(email: "yearly@example.com")
      Subscription.create!(user: yearly_user, plan_name: "developer", plan: "yearly", status: "active")
    end

    expected = PlanConfig::PLANS["developer"][:price_monthly] + PlanConfig.price_yearly_monthly("developer")
    assert_equal expected, delta
    assert_not_equal PlanConfig::PLANS["developer"][:price_monthly] * 2, delta,
      "a yearly subscriber must not be priced as if billed monthly"
  end

  test "calculate_mrr defaults a subscription with no plan (billing cycle) to monthly pricing" do
    delta = mrr_delta do
      user = create_user(email: "nocycle@example.com")
      Subscription.create!(user: user, plan_name: "developer", plan: nil, status: "active")
    end

    assert_equal PlanConfig::PLANS["developer"][:price_monthly], delta
  end
end
