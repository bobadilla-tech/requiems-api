# frozen_string_literal: true

require "test_helper"

class Dashboard::BillingControllerTest < ActionDispatch::IntegrationTest
  NOW = Time.utc(2026, 10, 15, 12)

  def setup
    @user = create_user(email: "billing-usage@example.com")
    @key = @user.api_keys.create!(name: "Main", environment: "test")
    @key.update_columns(created_at: Time.utc(2026, 4, 16, 8))
    sign_in @user
  end

  def log(credits, at: Time.utc(2026, 10, 1))
    create_usage_log(user: @user, key: @key, used_at: at, credits_used: credits)
  end

  test "show requires authentication" do
    sign_out @user
    get dashboard_billing_path

    assert_redirected_to new_user_session_path
  end

  %w[en es fr].each do |locale|
    test "show renders in #{locale}" do
      travel_to(NOW) { get dashboard_billing_path(locale: locale) }

      assert_response :success
    end
  end

  test "usage uses the billing cycle and ignores older logs" do
    log(7_000, at: Time.utc(2026, 7, 1))
    log(12)

    travel_to(NOW) { get dashboard_billing_path }

    assert_response :success
    assert_select "span.tabular-nums", text: "12"
    assert_select "[role=progressbar][aria-valuenow='2.4']"
  end

  test "no banner below 75 percent" do
    log(100)

    travel_to(NOW) { get dashboard_billing_path }

    assert_select "[role=alert]", 0
    assert_select ".bg-yellow-50", 0
  end

  test "warn state shows the moderate banner with the real percentage" do
    log(380)

    travel_to(NOW) { get dashboard_billing_path }

    assert_select ".bg-yellow-50", 1
    assert_select ".bg-yellow-50", text: /76% of your credits/
    assert_select "[role=alert]", 0
  end

  test "high and exhausted states show the critical banner" do
    log(460)
    travel_to(NOW) { get dashboard_billing_path }
    assert_select "[role=alert].bg-red-50", 1

    assert_select "[role=alert]", text: /Running low on credits/

    log(100, at: Time.utc(2026, 10, 2))
    travel_to(NOW) { get dashboard_billing_path }
    assert_select "[role=alert].bg-red-50", 1
    assert_select "[role=alert]", text: /Credit limit reached/
    assert_select "[role=alert]", text: /429/
    assert_select "[role=progressbar] > div[style*='width: 100%']"
  end

  test "a cancelled paid subscription is shown and measured as the free plan" do
    Subscription.create!(user: @user, plan_name: "professional", status: "cancelled", current_period_start: Time.utc(2026, 4, 16))
    log(460)

    travel_to(NOW) { get dashboard_billing_path }

    assert_select "[role=alert].bg-red-50", 1
    assert_select "span", text: "of 500 credits"
    assert_select "h3", text: /Free Plan/
    assert_select "h3", text: /\$150/, count: 0
  end

  test "the plan shown and the usage limit come from the same active subscription" do
    Subscription.create!(user: @user, plan_name: "business", status: "cancelled", current_period_start: Time.utc(2026, 1, 1))
    Subscription.create!(user: @user, plan_name: "developer", status: "active", current_period_start: Time.utc(2026, 4, 16))

    travel_to(NOW) { get dashboard_billing_path }

    assert_select "span", text: "of 100,000 credits"
  end

  test "a past_due paid subscriber keeps a way to the customer portal" do
    Subscription.create!(
      user: @user, plan_name: "developer", status: "past_due",
      current_period_start: Time.utc(2026, 4, 16), lemonsqueezy_subscription_id: "ls_past_due"
    )

    travel_to(NOW) { get dashboard_billing_path }

    assert_select "form[action=?]", dashboard_portal_billing_path
    assert_select "a", text: /Upgrade/i
  end

  test "free users without a payment record see no portal button" do
    travel_to(NOW) { get dashboard_billing_path }

    assert_select "form[action=?]", dashboard_portal_billing_path, count: 0
  end
end
