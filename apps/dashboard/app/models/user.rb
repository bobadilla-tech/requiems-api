# frozen_string_literal: true

class User < ApplicationRecord
  devise :database_authenticatable, :registerable,
         :recoverable, :rememberable, :validatable,
         :confirmable, :trackable

  has_many :api_keys, dependent: :destroy
  has_one :subscription, dependent: :destroy
  has_many :usage_logs, dependent: :destroy
  has_many :daily_usage_summaries, dependent: :destroy
  has_many :credit_adjustments, dependent: :destroy
  has_many :audit_logs, dependent: :destroy
  has_many :abuse_reports, dependent: :destroy
  has_many :private_deployment_requests, dependent: :destroy
  has_many :referrals_given, class_name: "Referral", foreign_key: "referrer_id", dependent: :destroy, inverse_of: :referrer
  has_one :referral_received, class_name: "Referral", foreign_key: "referred_user_id", dependent: :destroy, inverse_of: :referred_user

  PLAN_LIMITS = PlanConfig::PLANS.transform_values { |v| v[:requests_per_month] }.freeze
  SUPPORTED_LOCALES = Rails.application.config.i18n.available_locales.map(&:to_s).freeze

  validates :locale, inclusion: { in: SUPPORTED_LOCALES }, allow_nil: true
  validates :referral_code, uniqueness: true, allow_nil: true

  before_create :ensure_referral_code

  scope :admins, -> { where(admin: true) }
  scope :active_users, -> { where(status: "active") }
  scope :suspended, -> { where(status: "suspended") }
  scope :banned, -> { where(status: "banned") }

  scope :search_by, ->(query) {
    term = "%#{query}%"
    where("email ILIKE ? OR name ILIKE ? OR company ILIKE ?", term, term, term)
  }

  scope :with_plan, ->(plan) {
    if plan == "free"
      where.missing(:subscription)
        .or(left_joins(:subscription).where(subscriptions: { plan_name: "free" }))
    else
      joins(:subscription).where(subscriptions: { plan_name: plan })
    end
  }

  scope :with_status, ->(status) {
    case status
    when "admin" then where(admin: true)
    when "suspended" then where(status: "suspended")
    when "active" then where(status: "active")
    else all
    end
  }

  scope :sorted_by, ->(sort) {
    case sort
    when "oldest" then order(created_at: :asc)
    when "name" then order(name: :asc)
    else order(created_at: :desc)
    end
  }

  def admin?
    admin == true
  end

  def active_user?
    status == "active" && !banned_at
  end

  def suspended?
    status == "suspended"
  end

  def banned?
    status == "banned" || banned_at.present?
  end

  def display_name
    name.presence || email.split("@").first.titleize
  end

  def current_plan
    subscription&.plan_name || "free"
  end

  def usage_this_month
    period_start = subscription&.current_period_start || Time.current.beginning_of_month
    usage_logs.where("used_at >= ?", period_start).count
  end

  def total_requests
    usage_logs.count
  end

  def requests_remaining
    limit = PlanConfig.requests_per_month(current_plan)
    [ limit - usage_this_month, 0 ].max
  end

  def avg_response_time_ms
    logs = usage_logs.last_7_days.with_response_time
    return 0 if logs.empty?

    (logs.average(:response_time_ms) || 0).round
  end

  def recent_activity(limit = 10)
    usage_logs.recent(limit)
  end

  def ban!(reason:, admin_user:)
    key_prefixes_to_invalidate = api_keys.active_keys.pluck(:key_prefix)

    transaction do
      update!(
        status: "banned",
        banned_at: Time.current,
        banned_reason: reason,
        active: false
      )

      # update_all bypasses ActiveRecord callbacks entirely, so ApiKey's
      # after_update_commit :invalidate_go_auth_cache never fires for these
      # rows — invalidate explicitly below instead.
      api_keys.update_all(active: false, revoked_at: Time.current)

      AuditLog.create!(
        user: self,
        admin_user: admin_user,
        action: "ban_user",
        details: { reason: reason }.to_json
      )
    end

    # Runs after the transaction above has committed (not inside a Rails
    # callback), so this can't race a Go-side cache re-warm the way an
    # after_update callback fired mid-transaction could.
    key_prefixes_to_invalidate.each { |prefix| GoAuthCache.invalidate(prefix) }
  end

  def suspend!(admin_user:)
    update!(status: "suspended", active: false)
    AuditLog.create!(user: self, admin_user: admin_user, action: "suspend_user")
  end

  def unsuspend!(admin_user:)
    update!(status: "active", active: true)
    AuditLog.create!(user: self, admin_user: admin_user, action: "unsuspend_user")
  end

  def request_account_deletion!(reason)
    update!(
      deletion_token: SecureRandom.urlsafe_base64(32),
      deletion_token_sent_at: Time.current,
      deletion_reason: reason
    )
  end

  def deletion_token_valid?(token)
    deletion_token.present? &&
      deletion_token_sent_at.present? &&
      deletion_token_sent_at > 1.hour.ago &&
      ActiveSupport::SecurityUtils.secure_compare(deletion_token, token)
  end

  def clear_deletion_token!
    update_columns(deletion_token: nil, deletion_token_sent_at: nil, deletion_reason: nil)
  end

  def referral_token
    Rails.application.message_verifier("referral").generate(referral_code, expires_in: 30.days)
  end

  private

  def ensure_referral_code
    return if referral_code.present?

    loop do
      self.referral_code = SecureRandom.alphanumeric(8).upcase
      break unless User.exists?(referral_code: referral_code)
    end
  end
end
