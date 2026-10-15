# frozen_string_literal: true

module Dashboard::OverviewHelper
  STATUS_CHIPS = {
    success: "bg-green-100 text-green-800 dark:bg-green-900/40 dark:text-green-300",
    redirect: "bg-blue-100 text-blue-800 dark:bg-blue-900/40 dark:text-blue-300",
    client_error: "bg-yellow-100 text-yellow-900 dark:bg-yellow-900/40 dark:text-yellow-200",
    server_error: "bg-red-100 text-red-800 dark:bg-red-900/40 dark:text-red-300",
    unknown: "bg-gray-100 text-gray-800 dark:bg-gray-700 dark:text-gray-200"
  }.freeze

  # Legacy rows stored status_code 0 (or NULL); they are treated as successful.
  def activity_status_kind(log)
    code = log.status_code.to_i
    case code
    when 0, 200..299 then :success
    when 300..399 then :redirect
    when 400..499 then :client_error
    when 500..599 then :server_error
    else :unknown
    end
  end

  def activity_error?(log)
    %i[client_error server_error].include?(activity_status_kind(log))
  end

  def activity_status_chip_classes(log)
    STATUS_CHIPS.fetch(activity_status_kind(log))
  end

  # What was charged, as stored. NULL is "not recorded" (the quota counter treats
  # it as 0); an error row that still carries credits came from before errors
  # became free and says so.
  def activity_credits_label(log)
    credits = log.credits_used
    return t("dashboard.overview.activity.credits_unrecorded") if credits.nil?
    return t("dashboard.overview.activity.credits_none") if credits.zero?

    key = activity_error?(log) ? "credits_charged" : "credits"
    t("dashboard.overview.activity.#{key}", count: credits)
  end

  def activity_filter_classes(active)
    base = "px-3 py-1.5 rounded-full text-xs font-medium transition-colors focus:outline-none focus-visible:ring-2 focus-visible:ring-blue-500"
    state = if active
      "bg-blue-600 text-white"
    else
      "bg-gray-100 text-gray-700 hover:bg-gray-200 dark:bg-gray-700 dark:text-gray-200 dark:hover:bg-gray-600"
    end
    "#{base} #{state}"
  end
end
