# frozen_string_literal: true

class UsageLog < ApplicationRecord
  belongs_to :user
  belongs_to :api_key

  scope :in_range, ->(start_date, end_date) { where(used_at: start_date..end_date) }
  scope :successful, -> { where(status_code: 200..299) }
  scope :with_errors, -> { where("status_code >= ?", 400) }
  scope :this_month, -> { where("used_at >= ?", Time.current.beginning_of_month) }
  scope :last_7_days, -> { where("used_at >= ?", 7.days.ago) }
  scope :with_response_time, -> { where.not(response_time_ms: nil) }
  scope :recent, ->(limit = 10) { order(used_at: :desc).limit(limit).includes(:api_key) }
  validates :user_id, :used_at, :endpoint, presence: true
  validates :user_id, :used_at, :endpoint, presence: true

  before_create { self.api_key_name ||= api_key&.name }

  def self.error_rate_for(scope)
    total = scope.count
    return 0 if total.zero?

    errors = scope.where("status_code >= ?", 400).count
    ((errors.to_f / total) * 100).round(2)
  end
end
