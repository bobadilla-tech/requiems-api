# frozen_string_literal: true

# Presentation of a UsageSummary on the Overview and Billing pages.
#
# Counts are shown exactly (delimited); only the plan limit is compacted.
# compact_number uses the English K/M/B suffixes in every locale (10M reads
# fine in es/fr as well).
module UsageSummaryHelper
  MIN_VISIBLE_BAR_PERCENT = 1.0

  BAR_COLORS = {
    ok: "bg-blue-500",
    warn: "bg-yellow-500",
    high: "bg-red-500",
    exhausted: "bg-red-600"
  }.freeze

  # "0.02", "12.5", "100" or "<0.01" (used > 0 but rounds to nothing).
  def usage_percent_text(summary)
    return "<0.01" if summary.percent_below_floor?

    number_with_precision(summary.percent, precision: 2, strip_insignificant_zeros: true)
  end

  # Width in percent with a "." decimal in every locale (it goes into CSS).
  def usage_bar_width(summary)
    return "0" if summary.unlimited? || summary.used.zero?

    width = summary.percent.clamp(MIN_VISIBLE_BAR_PERCENT, 100)
    number_with_precision(width, precision: 2, strip_insignificant_zeros: true, locale: :en)
  end

  def usage_bar_color(summary)
    BAR_COLORS.fetch(summary.threshold_state)
  end

  # "Sep 16 – Oct 15" (the last day of the cycle is the day before it resets).
  def usage_cycle_range(summary)
    return nil unless summary.anchored?

    from = l(summary.cycle_start.to_date, format: :cycle_day)
    to = l((summary.resets_at - 1.day).to_date, format: :cycle_day)
    t("dashboard.overview.usage.cycle_range", from: from, to: to)
  end

  def usage_resets_on(summary)
    l(summary.resets_at.to_date, format: :cycle_day) if summary.anchored?
  end

  def usage_resets_in_days(summary)
    (summary.resets_at.to_date - summary.cycle.now.utc.to_date).to_i
  end

  def usage_last_request_text(summary)
    last = summary.last_request_at
    return t("dashboard.overview.usage.no_requests_yet") unless last

    t("dashboard.overview.usage.last_request", time: time_ago_in_words(last))
  end
end
