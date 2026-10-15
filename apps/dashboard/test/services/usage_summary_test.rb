# frozen_string_literal: true

require "test_helper"
require "minitest/mock"

class UsageSummaryTest < ActiveSupport::TestCase
  LEDGER_FIXTURE = Rails.root.join("../api/platform/middleware/testdata/usage_ledger_fixture.json")
  NOW = Time.utc(2026, 10, 15, 12)

  def setup
    @user = create_user(email: "summary@example.com")
    @key = @user.api_keys.create!(name: "Main", environment: "test")
    @key.update_columns(created_at: Time.utc(2026, 4, 16, 8))
  end

  def summary(now: NOW)
    UsageSummary.new(@user, now: now)
  end

  def log(used_at, credits)
    create_usage_log(user: @user, key: @key, used_at: used_at, credits_used: credits)
  end

  def subscribe(plan, status: "active", period_start: Time.utc(2026, 4, 16), **attrs)
    Subscription.create!(user: @user, plan_name: plan, status: status, current_period_start: period_start, **attrs)
  end

  test "used equals the ledger sum the Go sumUsage query computes on the shared fixture" do
    unless LEDGER_FIXTURE.exist?
      flunk "#{LEDGER_FIXTURE} is missing" if ENV["CI"]

      skip "shared fixture is not in this checkout"
    end

    fixture = JSON.parse(LEDGER_FIXTURE.read)
    other = create_user(email: "other-ledger@example.com")
    other_key = other.api_keys.create!(name: "Other", environment: "test")
    owners = { "subject" => [ @user, @key ], "other" => [ other, other_key ] }

    fixture["rows"].each do |row|
      user, key = owners.fetch(row["owner"])
      create_usage_log(user: user, key: key, used_at: Time.iso8601(row["used_at"]), credits_used: row["credits_used"])
    end

    cycle_start = Time.iso8601(fixture["cycle_start"])
    # An anchor on the 16th puts the cycle start on the fixture's boundary.
    @key.update_columns(created_at: Time.utc(2026, 1, 16))
    result = UsageSummary.new(@user, now: Time.utc(2026, 10, 15))

    assert_equal cycle_start, result.cycle_start
    assert_equal fixture["expected_used"], result.used
  end

  test "screenshot user: promo anchor in April and July logs show 0 used in October" do
    admin = create_user(email: "admin-summary@example.com", admin: true)
    subscribe("professional", period_start: Time.utc(2026, 4, 16), promoted_by: admin, promotion_reason: "x")
    3.times { |i| log(Time.utc(2026, 7, 10 + i), 1) }

    result = summary

    assert_equal 0, result.used
    assert_equal Time.utc(2026, 7, 12), result.last_request_at
    assert_equal Time.utc(2026, 9, 16), result.cycle_start
    assert_equal Time.utc(2026, 10, 16), result.resets_at
    assert_equal "0", format("%g", result.percent)
  end

  test "2,018 of 10M is 0.02 percent" do
    subscribe("professional", period_start: Time.utc(2026, 10, 1))
    log(Time.utc(2026, 10, 5), 2018)

    result = summary

    assert_equal 2018, result.used
    assert_equal 10_000_000, result.limit
    assert_in_delta 0.02, result.percent, 0.0001
    assert_not result.percent_below_floor?
    assert_equal 9_997_982, result.remaining
    assert_equal :ok, result.threshold_state
  end

  test "usage that rounds to nothing is flagged as below the floor" do
    subscribe("professional", period_start: Time.utc(2026, 10, 1))
    log(Time.utc(2026, 10, 5), 1)

    result = summary

    assert_equal 0.0, result.percent
    assert result.percent_below_floor?
  end

  test "NULL credits count as zero and 2 credit calls count twice" do
    log(Time.utc(2026, 10, 1, 1), nil)
    log(Time.utc(2026, 10, 1, 2), 0)
    log(Time.utc(2026, 10, 1, 3), 2)
    log(Time.utc(2026, 10, 1, 4), 1)

    assert_equal 3, summary.used
  end

  test "the exact cycle boundary is included and the second before is not" do
    log(Time.utc(2026, 9, 16, 0, 0, 0), 3)
    log(Time.utc(2026, 9, 15, 23, 59, 59), 5)

    assert_equal 3, summary.used
  end

  test "rows after now still count because Go has no upper bound" do
    log(Time.utc(2026, 12, 1), 4)

    assert_equal 4, summary.used
  end

  test "another user's usage is not counted" do
    other = create_user(email: "other-summary@example.com")
    other_key = other.api_keys.create!(name: "Other", environment: "test")
    create_usage_log(user: other, key: other_key, used_at: Time.utc(2026, 10, 1), credits_used: 99)

    assert_equal 0, summary.used
  end

  test "free users with no subscription use the free limit" do
    result = summary

    assert_equal "free", result.plan_name
    assert_equal 500, result.limit
    assert_equal 500, result.remaining
  end

  test "yearly and monthly subscriptions use the same monthly cycle" do
    subscribe("developer", period_start: Time.utc(2026, 1, 20), current_period_end: Time.utc(2027, 1, 20))

    assert_equal Time.utc(2026, 10, 20) - 1.month, summary.cycle_start
    assert_equal Time.utc(2026, 10, 20), summary.resets_at
  end

  test "a cancelled paid subscription falls back to the free plan and key anchor" do
    subscribe("professional", status: "cancelled")

    result = summary

    assert_equal "free", result.plan_name
    assert_equal 500, result.limit
    assert_equal Time.utc(2026, 9, 16), result.cycle_start
  end

  test "past_due paid subscription is not enforced as paid" do
    subscribe("business", status: "past_due")

    assert_equal 500, summary.limit
  end

  test "an unknown plan has no limit: unlimited state without percent or divide by zero" do
    sub = subscribe("developer")
    sub.update_columns(plan_name: "mystery")

    result = summary

    assert_nil result.limit
    assert result.unlimited?
    assert_nil result.remaining
    assert_nil result.percent
    assert_equal :ok, result.threshold_state
    assert_not result.percent_below_floor?
  end

  test "a zero limit is treated as unlimited, not a divide by zero" do
    log(Time.utc(2026, 10, 1), 5)

    result = summary
    result.stub(:limit, 0) do
      assert result.unlimited?
      assert_nil result.percent
      assert_nil result.remaining
    end
  end

  test "threshold states follow the 75, 90 and 100 percent lines of a 500 credit plan" do
    {
      374 => :ok,
      375 => :warn,
      449 => :warn,
      450 => :high,
      499 => :high,
      500 => :exhausted,
      900 => :exhausted
    }.each do |used, state|
      UsageLog.delete_all
      log(Time.utc(2026, 10, 1), used)

      assert_equal state, summary.threshold_state, "used=#{used}"
    end
  end

  test "percent never rounds up to 100 before the limit is reached" do
    log(Time.utc(2026, 10, 1), 499)

    assert_equal :high, summary.threshold_state
    assert_operator summary.percent, :<, 100
    assert_equal 1, summary.remaining

    sub = subscribe("professional")
    UsageLog.delete_all
    log(Time.utc(2026, 10, 1), 9_999_999)
    assert_operator summary.percent, :<, 100
    sub.destroy
  end

  test "over the limit never reports negative remaining" do
    log(Time.utc(2026, 10, 1), 800)

    assert_equal 0, summary.remaining
    assert_in_delta 160.0, summary.percent, 0.01
  end

  test "no cycle means nothing is used and no last request" do
    @key.update_columns(active: false, revoked_at: Time.current)
    result = summary

    assert_not result.anchored?
    assert_equal 0, result.used
    assert_nil result.cycle_start
    assert_nil result.resets_at
    assert_nil result.last_request_at
  end

  test "used is not cached between calls on separate instances" do
    assert_equal 0, summary.used
    log(Time.utc(2026, 10, 1), 7)

    assert_equal 7, summary.used
  end
end
