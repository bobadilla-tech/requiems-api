# frozen_string_literal: true

class Subscription < ApplicationRecord
  belongs_to :user
  belongs_to :promoted_by, class_name: "User", optional: true
  has_one :referral, foreign_key: "converting_subscription_id", dependent: :nullify, inverse_of: :converting_subscription

  # Validations
  validates :plan_name, presence: true, inclusion: { in: %w[free developer business professional] }
  validates :lemonsqueezy_subscription_id, uniqueness: true, allow_nil: true
  validates :status, presence: true

  # Scopes
  scope :active, -> { where(status: %w[active trialing]) }
  scope :cancelled, -> { where(status: "cancelled") }
  scope :paid, -> { where.not(plan_name: "free") }
  scope :paying, -> { paid.where(cancel_at_period_end: [ false, nil ]) }
  scope :promotional, -> { where.not(promoted_by_id: nil) }

  # The subscription the Go API bills against: active before trialing, then
  # the most recently updated. Several rows per user are possible because
  # subscriptions.user_id is not unique.
  def self.effective_for(user)
    active.where(user_id: user.id)
      .order(
        Arel.sql("CASE status WHEN 'active' THEN 0 ELSE 1 END"),
        Arel.sql("COALESCE(updated_at, created_at) DESC"),
        id: :desc
      )
      .first
  end

  def promoted?
    promoted_by_id.present?
  end
end
