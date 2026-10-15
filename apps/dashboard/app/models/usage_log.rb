# frozen_string_literal: true

class UsageLog < ApplicationRecord
  belongs_to :user
  belongs_to :api_key

  scope :in_range, ->(start_date, end_date) { where(used_at: start_date..end_date) }
  scope :successful, -> { where(status_code: 200..299) }
  scope :with_errors, -> { where("status_code >= ?", 400) }
  scope :last_7_days, -> { where("used_at >= ?", 7.days.ago) }
  scope :with_response_time, -> { where.not(response_time_ms: nil) }
  scope :recent, ->(limit = 10) { order(used_at: :desc, id: :desc).limit(limit).includes(:api_key) }
  # Legacy rows stored status_code 0 (or NULL) for "no status"; they count as successful.
  scope :activity_success, -> { where(status_code: [ nil, 0 ]).or(where(status_code: 200..399)) }

  ACTIVITY_FILTERS = %w[all success errors].freeze

  def self.for_activity_filter(filter)
    case filter
    when "success" then activity_success
    when "errors" then with_errors
    else all
    end
  end
  validates :used_at, :endpoint, presence: true

  def self.error_rate_for(scope)
    total = scope.count
    return 0 if total.zero?

    errors = scope.where("status_code >= ?", 400).count
    ((errors.to_f / total) * 100).round(2)
  end
end
