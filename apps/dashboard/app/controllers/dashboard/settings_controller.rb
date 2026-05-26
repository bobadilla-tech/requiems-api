# frozen_string_literal: true

class Dashboard::SettingsController < ApplicationController
  before_action :authenticate_user!
  layout "dashboard"

  def show
    # Settings page - view current settings
  end

  def update
    # Update account settings
    if current_user.update(user_params)
      # If email was changed and Devise confirmable is enabled, send confirmation
      if user_params[:email].present? && current_user.email != current_user.email_was
        flash[:notice] = t("dashboard.settings.flash.email_confirm")
      else
        flash[:notice] = t("dashboard.settings.flash.updated")
      end

      redirect_to dashboard_settings_path
    else
      flash.now[:alert] = t("dashboard.settings.flash.error")
      render :show, status: :unprocessable_entity
    end
  end

  def request_deletion
    reason = params[:deletion_reason].to_s.strip

    if reason.length < 10
      redirect_to dashboard_settings_path, alert: t("dashboard.settings.flash.reason_too_short")
      return
    end

    current_user.request_account_deletion!(reason)
    AccountDeletionMailer.confirmation(current_user).deliver_later

    redirect_to dashboard_settings_path,
      notice: t("dashboard.settings.flash.deletion_email_sent")
  end

  def confirm_deletion
    token = params[:token].to_s

    unless current_user.deletion_token_valid?(token)
      redirect_to dashboard_settings_path, alert: t("dashboard.settings.flash.invalid_token")
      return
    end

    @token = token
    @reason = current_user.deletion_reason
  end

  def execute_deletion
    token = params[:token].to_s

    unless current_user.deletion_token_valid?(token)
      redirect_to dashboard_settings_path, alert: t("dashboard.settings.flash.invalid_token")
      return
    end

    # Revoke all API keys
    current_user.api_keys.each do |key|
      key.revoke!(reason: "Account deleted by user")
    end

    # Cancel subscription if exists
    if current_user.subscription
      current_user.subscription.update!(
        cancel_at_period_end: true,
        canceled_at: Time.current
      )
    end

    current_user.destroy!
    sign_out current_user

    redirect_to root_path, notice: t("dashboard.settings.flash.deleted")
  end

  private

  def user_params
    p = params.require(:user).permit(:email, :name, :company, :locale, :email_notifications, :usage_alerts, :weekly_reports)
    p[:locale] = p[:locale].presence # convert "" (auto-detect) to nil
    p
  end
end
