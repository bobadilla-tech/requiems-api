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

  # Plan the Go API enforces: the active/trialing subscription, else free.
  def current_plan
    Subscription.effective_for(self)&.plan_name || "free"
  end

  def usage_summary(now: Time.current)
    UsageSummary.new(self, now: now)
  end

  # Credits used in the current billing cycle (see UsageSummary).
  def usage_this_month
    usage_summary.used
  end

  def total_requests
    usage_logs.count
  end

  # nil when the plan has no known limit.
  def requests_remaining
    usage_summary.remaining
  end

  def avg_response_time_ms
    logs = usage_logs.last_7_days.with_response_time
    return 0 if logs.empty?

    (logs.average(:response_time_ms) || 0).round
  end

  def recent_activity(limit = 10, filter: "all")
    usage_logs.for_activity_filter(filter).recent(limit)
  end

  def ban!(reason:, admin_user:)
    key_prefixes_to_invalidate = transaction do
      update!(
        status: "banned",
        banned_at: Time.current,
        banned_reason: reason,
        active: false
      )

      # update_all bypasses ActiveRecord callbacks entirely, so ApiKey's
      # after_update_commit callback never fires for these rows. Collect the
      # rows affected by this update inside the transaction instead of relying
      # on an active_keys snapshot taken before it.
      revoked_at = Time.current
      api_keys.update_all(active: false, revoked_at: revoked_at)

      AuditLog.create!(
        user: self,
        admin_user: admin_user,
        action: "ban_user",
        details: { reason: reason }.to_json
      )
      api_keys.where(revoked_at: revoked_at).pluck(:key_prefix)
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
