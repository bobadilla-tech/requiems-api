# frozen_string_literal: true

require "test_helper"

# The billing-cycle "used" number must come from one place (UsageSummary, backed
# by BillingCycle). A second copy of the query is how the dashboard drifted from
# Go in the first place, so ban the old shapes anywhere else in app/.
class UsageQueryGuardTest < ActiveSupport::TestCase
  ALLOWED = %w[
    app/services/usage_summary.rb
    app/lib/billing_cycle.rb
    app/models/usage_log.rb
  ].freeze

  # A used_at comparison tied to a calendar month or a subscription period,
  # across lines (the old code split these over a variable and a where).
  FORBIDDEN = [
    /used_at[^;]{0,200}beginning_of_month/m,
    /beginning_of_month[^;]{0,200}used_at/m,
    /used_at[^;]{0,200}current_period_start/m,
    /current_period_start[^;]{0,200}used_at/m,
    /used_at\s*>=\s*\?"?,\s*period_start/m
  ].freeze

  def app_files
    Dir[Rails.root.join("app/**/*.{rb,erb}")].map { |f| Pathname(f).relative_path_from(Rails.root).to_s }
  end

  test "no usage query outside UsageSummary and BillingCycle uses calendar months or subscription periods" do
    offenders = (app_files - ALLOWED).select do |path|
      source = File.read(Rails.root.join(path))
      FORBIDDEN.any? { |pattern| source.match?(pattern) }
    end

    assert_empty offenders, "Use UsageSummary for billing-cycle usage: #{offenders.join(', ')}"
  end

  test "the guard actually detects the old multi-line shape" do
    old_code = <<~RUBY
      period_start = subscription&.current_period_start || Time.current.beginning_of_month
      usage_logs.where("used_at >= ?", period_start).count
    RUBY

    assert(FORBIDDEN.any? { |pattern| old_code.match?(pattern) })
  end

  test "UsageLog no longer exposes a calendar month scope" do
    assert_not UsageLog.respond_to?(:this_month)
  end
end
