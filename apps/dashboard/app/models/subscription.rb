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

  def promoted?
    promoted_by_id.present?
  end
end
