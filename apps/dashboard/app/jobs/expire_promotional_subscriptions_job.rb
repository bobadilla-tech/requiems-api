# frozen_string_literal: true

# Background job to expire admin-granted promotional plan upgrades.
#
# Runs hourly via Sidekiq Cron. Finds all subscriptions where
# a promotion has passed its expiry date and downgrades them back to the free
# plan, writing an audit log entry.
#
class ExpirePromotionalSubscriptionsJob < ApplicationJob
  queue_as :default

  def perform
    expired = Subscription
      .where.not(promoted_by_id: nil)
      .where("promotion_expires_at <= ?", Time.current)
      .where.not(plan_name: "free")

    count = 0

    expired.find_each do |subscription|
      ActiveRecord::Base.transaction do
        previous_plan = subscription.plan_name

        subscription.update!(
          plan_name: "free",
          status: "active",
          promoted_by_id: nil,
          promotion_reason: nil,
          promotion_expires_at: nil,
          current_period_end: nil
        )

        AuditLog.create!(
          user: subscription.user,
          action: "promotion_expired",
          details: { previous_plan: previous_plan }.to_json
        )
      end

      count += 1
      Rails.logger.info "[ExpirePromotionalSubscriptionsJob] Expired promotion for user #{subscription.user_id} (was #{subscription.plan_name})"
    rescue StandardError => e
      Rails.logger.error "[ExpirePromotionalSubscriptionsJob] Failed to expire subscription #{subscription.id}: #{e.message}"
    end

    Rails.logger.info "[ExpirePromotionalSubscriptionsJob] Processed #{count} expired promotions"
  end
end
