# frozen_string_literal: true

require "test_helper"

class UsageSummaryHelperTest < ActionView::TestCase
  include UsageSummaryHelper

  Fake = Struct.new(:used, :limit, :percent, :threshold_state, :percent_below_floor, :anchored, :cycle_start, :resets_at, keyword_init: true) do
    def unlimited? = limit.nil? || limit <= 0
    def percent_below_floor? = percent_below_floor
    def anchored? = anchored
  end

  def summary(used:, limit: 10_000_000, state: :ok, below: false, **extra)
    percent = limit.to_i.positive? ? (used.to_f / limit * 100).round(2) : nil
    Fake.new(used: used, limit: limit, percent: percent, threshold_state: state, percent_below_floor: below, anchored: true, **extra)
  end

  test "percent text keeps two decimals, strips zeros and floors tiny values" do
    assert_equal "0.02", usage_percent_text(summary(used: 2_018))
    assert_equal "12.5", usage_percent_text(summary(used: 1_250_000))
    assert_equal "100", usage_percent_text(summary(used: 10_000_000))
    assert_equal "0", usage_percent_text(summary(used: 0))
    assert_equal "<0.01", usage_percent_text(summary(used: 1, below: true))
  end

  test "percent text is localised but bar width always uses a decimal point" do
    I18n.with_locale(:es) do
      assert_equal "12,5", usage_percent_text(summary(used: 1_250_000))
      assert_equal "12.5", usage_bar_width(summary(used: 1_250_000))
    end
  end

  test "bar width is zero without usage or limit, at least 1 with usage and at most 100" do
    assert_equal "0", usage_bar_width(summary(used: 0))
    assert_equal "0", usage_bar_width(summary(used: 5, limit: nil))
    assert_equal "1", usage_bar_width(summary(used: 2_018))
    assert_equal "100", usage_bar_width(summary(used: 20_000_000))
  end

  test "bar colour follows the threshold state" do
    assert_equal "bg-blue-500", usage_bar_color(summary(used: 1, state: :ok))
    assert_equal "bg-yellow-500", usage_bar_color(summary(used: 1, state: :warn))
    assert_equal "bg-red-500", usage_bar_color(summary(used: 1, state: :high))
    assert_equal "bg-red-600", usage_bar_color(summary(used: 1, state: :exhausted))
  end

  test "cycle range ends the day before the reset" do
    fake = summary(used: 0, cycle_start: Time.utc(2026, 9, 16), resets_at: Time.utc(2026, 10, 16))

    assert_equal "Sep 16 – Oct 15 UTC", usage_cycle_range(fake)
    assert_equal "Oct 16", usage_resets_on(fake)
  end

  test "an unanchored summary has no range" do
    assert_nil usage_cycle_range(Fake.new(anchored: false))
    assert_nil usage_resets_on(Fake.new(anchored: false))
  end
end
