# frozen_string_literal: true

class Admin::PromotionsController < ApplicationController
  before_action :authenticate_user!
  before_action :require_admin!
  before_action :set_user
  layout "admin"

  ALLOWED_PLANS = %w[developer business professional].freeze

  def create
    plan_name = promotion_params[:plan_name]
    expires_at = parse_expires_at(promotion_params[:expires_at])
    reason = promotion_params[:reason]

    unless ALLOWED_PLANS.include?(plan_name)
      redirect_to admin_user_path(@user), alert: t("admin.promotions.invalid_plan") and return
    end

    if expires_at.nil? || expires_at <= Time.current
      redirect_to admin_user_path(@user), alert: t("admin.promotions.expiry_past") and return
    end

    if reason.blank?
      redirect_to admin_user_path(@user), alert: t("admin.promotions.reason_required") and return
    end

    ActiveRecord::Base.transaction do
      subscription = @user.subscription || @user.build_subscription
      subscription.update!(
        plan_name: plan_name,
        status: "active",
        promoted_by: current_user,
        promotion_reason: reason,
        promotion_expires_at: expires_at,
        current_period_start: Time.current,
        current_period_end: expires_at,
        lemonsqueezy_subscription_id: nil,
        lemonsqueezy_customer_id: nil,
        cancel_at_period_end: false
      )

      AuditLog.create!(
        user: @user,
        admin_user_id: current_user.id,
        action: "promotion_granted",
        details: { plan_name: plan_name, expires_at: expires_at.iso8601, reason: reason }.to_json
      )
    end

    begin
      PromotionMailer.upgrade_notification(@user, plan_name, expires_at, reason).deliver_later
    rescue StandardError => e
      Rails.logger.error "[PromotionsController] Failed to enqueue upgrade email for user #{@user.id}: #{e.message}"
    end

    redirect_to admin_user_path(@user),
                notice: t("admin.promotions.upgrade_success", email: @user.email, plan: plan_name.titleize, date: expires_at.strftime("%B %d, %Y"))
  rescue ActiveRecord::RecordInvalid => e
    redirect_to admin_user_path(@user), alert: t("admin.promotions.upgrade_error", message: e.message)
  end

  def destroy
    subscription = @user.subscription
    unless subscription&.promoted?
      redirect_to admin_user_path(@user), alert: t("admin.promotions.no_active_promotion") and return
    end

    previous_plan = subscription.plan_name

    ActiveRecord::Base.transaction do
      subscription.update!(
        plan_name: "free",
        status: "active",
        promoted_by: nil,
        promotion_reason: nil,
        promotion_expires_at: nil,
        current_period_end: nil
      )

      AuditLog.create!(
        user: @user,
        admin_user_id: current_user.id,
        action: "promotion_revoked",
        details: { previous_plan: previous_plan }.to_json
      )
    end

    redirect_to admin_user_path(@user),
                notice: t("admin.promotions.revoke_success", email: @user.email)
  rescue ActiveRecord::RecordInvalid => e
    redirect_to admin_user_path(@user), alert: t("admin.promotions.revoke_error", message: e.message)
  end

  private

  def require_admin!
    unless current_user.admin?
      redirect_to root_path, alert: t("admin.access_denied")
    end
  end

  def set_user
    @user = User.find(params[:user_id])
  rescue ActiveRecord::RecordNotFound
    redirect_to admin_users_path, alert: t("admin.users.not_found")
  end

  def promotion_params
    params.require(:promotion).permit(:plan_name, :expires_at, :reason)
  end

  def parse_expires_at(value)
    return nil if value.blank?
    Time.zone.parse(value)
  rescue ArgumentError, TypeError
    nil
  end
end
