# frozen_string_literal: true

# The usage cycle the Go API enforces quota against.
#
# Mirrors apps/api/platform/middleware (cycleStart + the apikeyauth.go
# subscription join) as it stands today: the anchor is the start of the user's
# active/trialing subscription, else the creation time of an API key; the cycle
# rolls over monthly on the anchor's UTC day-of-month at midnight UTC. Shared
# vectors in apps/api/platform/middleware/testdata/billing_cycle_vectors.json
# keep both implementations in step.
#
# Known difference (deferred D1): Go anchors a free user on the key that
# authenticates the request, the dashboard has no such key and uses the oldest
# active, non-revoked one. They only differ for free users with several keys.
class BillingCycle
  attr_reader :subscription, :now

  def self.for(user, now: Time.current)
    new(user, now: now)
  end

  # Most recent anchor day-of-month (clamped to the month length) at midnight
  # UTC, at or before now.
  def self.start(anchor, now)
    anchor = anchor.utc
    now = now.utc
    month = Date.new(now.year, now.month, 1)
    month = month.prev_month if midnight(clamped_day(month, anchor.day)) > now

    midnight(clamped_day(month, anchor.day))
  end

  # The next clamped anchor day after the current cycle start.
  def self.next_start(anchor, now)
    current = start(anchor, now)
    month = Date.new(current.year, current.month, 1).next_month

    midnight(clamped_day(month, anchor.utc.day))
  end

  def self.clamped_day(month_start, day)
    Date.new(month_start.year, month_start.month, [ day, month_start.end_of_month.day ].min)
  end
  private_class_method :clamped_day

  def self.midnight(date)
    Time.utc(date.year, date.month, date.day)
  end
  private_class_method :midnight

  def initialize(user, now: Time.current)
    @now = now
    @subscription = Subscription.effective_for(user)
    @anchor = subscription&.current_period_start || user.api_keys.active_keys.minimum(:created_at)
  end

  attr_reader :anchor

  def plan_name
    subscription&.plan_name || "free"
  end

  def anchored?
    anchor.present?
  end

  def start
    self.class.start(anchor, now) if anchored?
  end

  def next_start
    self.class.next_start(anchor, now) if anchored?
  end
end
