# frozen_string_literal: true

class Referral < ApplicationRecord
  belongs_to :referrer, class_name: "User"
  belongs_to :referred_user, class_name: "User"
  belongs_to :converting_subscription, class_name: "Subscription", optional: true

  validates :status, presence: true, inclusion: { in: %w[pending converted] }
  validate :no_self_referral

  scope :converted, -> { where(status: "converted") }
  scope :pending, -> { where(status: "pending") }

  def pending?
    status == "pending"
  end

  def converted?
    status == "converted"
  end

  def mark_converted!(subscription)
    with_lock do
      return if converted?

      update!(
        status: "converted",
        converted_at: Time.current,
        converting_subscription: subscription
      )
    end
  end

  private

  def no_self_referral
    return if referrer_id.blank? || referred_user_id.blank?

    errors.add(:referred_user, "cannot be the same as referrer") if referrer_id == referred_user_id
  end
end
