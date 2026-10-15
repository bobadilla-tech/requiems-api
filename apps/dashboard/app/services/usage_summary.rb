# frozen_string_literal: true

# What the dashboard shows for "usage this cycle". Computed per request, never
# cached, so the counter cannot lag the activity feed.
#
# `used` is the ledger sum Go's sumUsage computes (credits_used summed from the
# cycle start with no upper bound; NULL counts as 0, never as 1). Redis, which
# enforces the quota, can legitimately be higher than the ledger after dedup
# drops, panics or failed refunds.
class UsageSummary
  WARN_PERCENT = 75
  HIGH_PERCENT = 90

  attr_reader :user, :cycle

  delegate :plan_name, :anchored?, to: :cycle

  def initialize(user, now: Time.current)
    @user = user
    @cycle = BillingCycle.for(user, now: now)
  end

  def cycle_start
    cycle.start
  end

  def resets_at
    cycle.next_start
  end

  def used
    @used ||=
      if cycle_start
        user.usage_logs.where("used_at >= ?", cycle_start).sum(:credits_used).to_i
      else
        0
      end
  end

  # nil for an unknown plan, as Go treats a plan without limits.
  def limit
    PlanConfig::PLANS.dig(plan_name, :requests_per_month)
  end

  def unlimited?
    limit.nil? || limit <= 0
  end

  def remaining
    return nil if unlimited?

    [ limit - used, 0 ].max
  end

  def percent
    return nil if unlimited?

    value = (used.to_f / limit * 100).round(2)
    # Never show 100% until the limit is actually reached.
    used < limit ? [ value, 99.99 ].min : value
  end

  # True when something was used but it rounds to 0.00%.
  def percent_below_floor?
    !unlimited? && used.positive? && percent < 0.01
  end

  def threshold_state
    return :ok if unlimited?
    return :exhausted if used >= limit

    ratio = used.to_f / limit * 100
    if ratio >= HIGH_PERCENT then :high
    elsif ratio >= WARN_PERCENT then :warn
    else :ok
    end
  end

  def last_request_at
    return @last_request_at if defined?(@last_request_at)

    @last_request_at = user.usage_logs.maximum(:used_at)
  end
end
