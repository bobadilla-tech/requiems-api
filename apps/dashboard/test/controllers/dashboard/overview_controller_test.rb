# frozen_string_literal: true

require "test_helper"

class Dashboard::OverviewControllerTest < ActionDispatch::IntegrationTest
  def setup
    @user = create_user(
      email: "test@example.com",
      password: "password123",
      password_confirmation: "password123"
    )

    sign_in @user
  end

  test "index requires authentication" do
    sign_out @user
    get dashboard_root_path

    assert_redirected_to new_user_session_path
  end

  test "index renders successfully" do
    get dashboard_root_path

    assert_response :success
  end

  %w[en es fr].each do |locale|
    test "index renders successfully in #{locale} locale" do
      get dashboard_root_path(locale: locale)

      assert_response :success
    end
  end

  NOW = Time.utc(2026, 10, 15, 12)

  def seed_key(created_at: Time.utc(2026, 4, 16, 8))
    key = @user.api_keys.create!(name: "Main", environment: "test")
    key.update_columns(created_at: created_at)
    key
  end

  def subscribe(plan, **attrs)
    Subscription.create!(user: @user, plan_name: plan, status: "active", current_period_start: Time.utc(2026, 4, 16), **attrs)
  end

  test "screenshot user: April anchor, July logs, October shows 0 used and last request 3 months ago" do
    admin = create_user(email: "overview-admin@example.com", admin: true)
    subscribe("professional", promoted_by: admin, promotion_reason: "x", current_period_end: Time.utc(2027, 4, 16))
    key = seed_key
    3.times { |i| create_usage_log(user: @user, key: key, used_at: Time.utc(2026, 7, 10 + i)) }

    travel_to(NOW) { get dashboard_root_path }

    assert_response :success
    assert_select "[data-testid=credits-used]", text: /\A\s*0\s*\z/
    assert_select "[data-testid=percent-used]", text: /0% used/
    assert_select "[data-testid=last-request]", text: /Last request .*3 months ago/
    assert_select "[data-testid=cycle-line]", text: /Sep 16 – Oct 15 UTC · resets Oct 16 \(in 1 day\)/
    assert_select "[data-testid=resets-on]", text: "Oct 16"
  end

  test "2,018 of 10M shows 0.02 percent with the exact count" do
    subscribe("professional", current_period_start: Time.utc(2026, 10, 1))
    key = seed_key
    create_usage_log(user: @user, key: key, used_at: Time.utc(2026, 10, 5), credits_used: 2018)

    travel_to(NOW) { get dashboard_root_path }

    assert_select "[data-testid=credits-used]", text: /2,018/
    assert_select "[data-testid=percent-used]", text: /0\.02% used/
    assert_select "[data-testid=credits-remaining]", text: /9,997,982/
    assert_select "[role=progressbar][aria-valuenow='0.02']"
    assert_select "[role=progressbar] > div[style*='width: 1%']"
  end

  test "the bar width uses a decimal point in every locale and is hidden at zero usage" do
    subscribe("developer", current_period_start: Time.utc(2026, 10, 1))
    key = seed_key
    create_usage_log(user: @user, key: key, used_at: Time.utc(2026, 10, 5), credits_used: 12_500)

    %w[en es fr].each do |locale|
      travel_to(NOW) { get dashboard_root_path(locale: locale) }

      assert_select "[role=progressbar] > div[style*='width: 12.5%']", 1, locale
      assert_select "[role=progressbar][aria-valuenow='12.5']", 1, locale
    end
  end

  test "zero usage renders a zero width bar" do
    seed_key

    travel_to(NOW) { get dashboard_root_path }

    assert_select "[role=progressbar] > div[style*='width: 0%']"
  end

  test "a user over the free limit sees the exhausted state with text and a red bar" do
    key = seed_key
    create_usage_log(user: @user, key: key, used_at: Time.utc(2026, 10, 1), credits_used: 600)

    travel_to(NOW) { get dashboard_root_path }

    assert_select "[data-testid=threshold-state][data-state=exhausted]", text: /Limit reached/
    assert_select "[data-testid=credits-remaining]", text: /\A\s*0\s*\z/
    assert_select "[role=progressbar] > div.bg-red-600"
    assert_select "[role=progressbar] > div[style*='width: 100%']"
  end

  test "threshold states are conveyed with text, not colour alone" do
    key = seed_key
    create_usage_log(user: @user, key: key, used_at: Time.utc(2026, 10, 1), credits_used: 380)

    travel_to(NOW) { get dashboard_root_path }

    assert_select "[data-testid=threshold-state][data-state=warn]", text: /Running low/
    assert_select "[data-testid=threshold-state] svg[aria-hidden=true]"
  end

  test "an unknown plan renders the no-limit state without a bar or percentage" do
    subscribe("developer").update_columns(plan_name: "mystery")
    seed_key

    travel_to(NOW) { get dashboard_root_path }

    assert_response :success
    assert_select "[role=progressbar]", 0
    assert_select "[data-testid=percent-used]", 0
    assert_select "#usage-card", text: /No credit limit/
  end

  test "a new user without keys is pointed at API keys first" do
    get dashboard_root_path

    assert_response :success
    assert_select "#cycle-card a[href=?]", new_dashboard_api_key_path
    assert_select "[data-testid=cycle-line]", 0
    assert_select "[data-testid=last-request]", text: /No requests yet/
  end

  test "promotional subscriptions show Plan Ends instead of Next Billing" do
    admin = create_user(email: "overview-admin2@example.com", admin: true)
    subscribe("business", promoted_by: admin, promotion_reason: "x", current_period_end: Time.utc(2027, 4, 16))
    seed_key

    travel_to(NOW) { get dashboard_root_path }

    assert_select "p", text: "Plan Ends"
    assert_select "p", text: "Next Billing Date", count: 0
    assert_select "[data-testid=resets-on]", text: "Oct 16"
  end

  test "paid subscriptions keep Next Billing distinct from usage resets" do
    subscribe("developer", current_period_start: Time.utc(2026, 10, 3), current_period_end: Time.utc(2026, 11, 20))
    seed_key

    travel_to(NOW) { get dashboard_root_path }

    assert_select "p", text: "Next Billing Date"
    assert_select "[data-testid=resets-on]", text: "Nov 3"
  end

  test "the banner subtitle shows the cycle" do
    seed_key

    travel_to(NOW) { get dashboard_root_path }

    assert_select "p", text: /Sep 16 – Oct 15 UTC/
  end

  test "the page never promises alerts or weekly reports" do
    seed_key
    get dashboard_root_path

    assert_not_includes response.body.downcase, "weekly report"
  end

  test "the overview does not leak other users' usage" do
    other = create_user(email: "overview-other@example.com")
    other_key = other.api_keys.create!(name: "Other", environment: "test")
    other_key.update_columns(created_at: Time.utc(2026, 4, 16))
    create_usage_log(user: other, key: other_key, used_at: Time.utc(2026, 10, 1), credits_used: 400)
    seed_key

    travel_to(NOW) { get dashboard_root_path }

    assert_select "[data-testid=credits-used]", text: /\A\s*0\s*\z/
  end
end
