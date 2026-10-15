# frozen_string_literal: true

require "test_helper"

class BillingCycleTest < ActiveSupport::TestCase
  VECTORS_PATH = Rails.root.join("../api/platform/middleware/testdata/billing_cycle_vectors.json")

  def setup
    @user = create_user(email: "cycle@example.com")
  end

  def vectors
    skip "shared vectors are not in this checkout (outside the monorepo)" unless VECTORS_PATH.exist?

    JSON.parse(VECTORS_PATH.read)
  end

  test "start and next_start match the vectors shared with the Go API" do
    vectors.each do |v|
      anchor = Time.iso8601(v["anchor"])
      now = Time.iso8601(v["now"])

      assert_equal Time.iso8601(v["start"]), BillingCycle.start(anchor, now), "#{v['name']} (start)"
      assert_equal Time.iso8601(v["next_start"]), BillingCycle.next_start(anchor, now), "#{v['name']} (next_start)"
    end
  end

  test "shared vectors cover the edge cases that matter" do
    names = vectors.map { |v| v["name"] }.join(" ")

    %w[leap future boundary non\ midnight year\ wrap day\ 30 day\ 31 jan\ 31].each do |needle|
      assert_includes names, needle
    end
  end

  test "start accepts zoned times and is always midnight UTC" do
    anchor = Time.find_zone("Pacific/Auckland").parse("2026-04-17 08:00") # 2026-04-16 20:00 UTC
    start = BillingCycle.start(anchor, Time.utc(2026, 10, 15, 12))

    assert_equal Time.utc(2026, 9, 16), start
    assert start.utc?
  end

  test "next_start is never earlier than now and start never later" do
    anchor = Time.utc(2026, 1, 31, 9)
    t = Time.utc(2026, 1, 1)
    400.times do
      assert_operator BillingCycle.start(anchor, t), :<=, t
      assert_operator BillingCycle.next_start(anchor, t), :>, t
      t += 1.day
    end
  end

  test "anchors on the active subscription period start" do
    create_subscription("developer", "active", period_start: Time.utc(2026, 4, 16, 10))

    cycle = BillingCycle.for(@user, now: Time.utc(2026, 10, 15, 12))

    assert_equal Time.utc(2026, 9, 16), cycle.start
    assert_equal Time.utc(2026, 10, 16), cycle.next_start
    assert_equal "developer", cycle.plan_name
  end

  test "trialing subscriptions anchor like active ones" do
    create_subscription("business", "trialing", period_start: Time.utc(2026, 9, 3))

    cycle = BillingCycle.for(@user, now: Time.utc(2026, 10, 15))

    assert_equal Time.utc(2026, 10, 3), cycle.start
    assert_equal "business", cycle.plan_name
  end

  test "cancelled, past_due and expired subscriptions are ignored like Go does" do
    %w[cancelled past_due expired].each do |status|
      create_subscription("professional", status, period_start: Time.utc(2026, 4, 16))
    end
    key = @user.api_keys.create!(name: "k", environment: "test")
    key.update_columns(created_at: Time.utc(2026, 1, 5))

    cycle = BillingCycle.for(@user, now: Time.utc(2026, 10, 15))

    assert_equal "free", cycle.plan_name
    assert_equal Time.utc(2026, 10, 5), cycle.start
  end

  test "an active subscription wins over a trialing one, then the newest wins" do
    create_subscription("developer", "trialing", period_start: Time.utc(2026, 1, 1))
    older = create_subscription("business", "active", period_start: Time.utc(2026, 2, 2))
    older.update_columns(updated_at: 2.days.ago)
    newer = create_subscription("professional", "active", period_start: Time.utc(2026, 3, 3))
    newer.update_columns(updated_at: 1.day.ago)

    cycle = BillingCycle.for(@user, now: Time.utc(2026, 10, 15))

    assert_equal "professional", cycle.plan_name
    assert_equal Time.utc(2026, 10, 3), cycle.start
  end

  test "a subscription without a period start keeps its plan but anchors on the key" do
    create_subscription("developer", "active", period_start: nil)
    key = @user.api_keys.create!(name: "k", environment: "test")
    key.update_columns(created_at: Time.utc(2026, 6, 20))

    cycle = BillingCycle.for(@user, now: Time.utc(2026, 10, 15))

    assert_equal "developer", cycle.plan_name
    assert_equal Time.utc(2026, 9, 20), cycle.start
  end

  test "free users anchor on their oldest active, non-revoked key" do
    oldest = @user.api_keys.create!(name: "oldest", environment: "test")
    oldest.update_columns(created_at: Time.utc(2026, 1, 3))
    newer = @user.api_keys.create!(name: "newer", environment: "test")
    newer.update_columns(created_at: Time.utc(2026, 5, 20))

    assert_equal Time.utc(2026, 10, 3), BillingCycle.for(@user, now: Time.utc(2026, 10, 15)).start

    oldest.update_columns(revoked_at: Time.current, active: false)

    assert_equal Time.utc(2026, 9, 20), BillingCycle.for(@user, now: Time.utc(2026, 10, 15)).start
  end

  test "has no cycle without a subscription period or an active key" do
    cycle = BillingCycle.for(@user)

    assert_not cycle.anchored?
    assert_nil cycle.start
    assert_nil cycle.next_start
    assert_equal "free", cycle.plan_name
  end

  test "an anchor in the future yields the previous cycle" do
    create_subscription("developer", "active", period_start: Time.utc(2027, 4, 16))

    cycle = BillingCycle.for(@user, now: Time.utc(2026, 10, 15, 12))

    assert_equal Time.utc(2026, 9, 16), cycle.start
    assert_equal Time.utc(2026, 10, 16), cycle.next_start
  end

  private

  def create_subscription(plan, status, period_start:)
    Subscription.create!(user: @user, plan_name: plan, status: status, current_period_start: period_start)
  end
end
