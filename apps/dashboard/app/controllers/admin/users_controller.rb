# frozen_string_literal: true

class Admin::UsersController < ApplicationController
  before_action :authenticate_user!
  before_action :require_admin!
  before_action :set_user, only: [ :show, :suspend, :unsuspend, :ban, :make_admin, :remove_admin ]
  layout "admin"

  def index
    @users = User.all
    @users = @users.search_by(params[:search]) if params[:search].present?
    @users = @users.with_plan(params[:plan]) if params[:plan].present? && params[:plan] != "all"
    @users = @users.with_status(params[:status]) if params[:status].present?
    @users = @users.sorted_by(params[:sort])
    @pagy, @users = pagy(@users, limit: 20)
  end

  def show
    @api_keys = @user.api_keys.order(created_at: :desc)
    @usage_stats = calculate_user_usage_stats
    @recent_activity = @user.usage_logs.order(used_at: :desc).limit(20)
  end

  def suspend
    if @user.update(status: "suspended", active: false)
      # Revoke all API keys
      @user.api_keys.active_keys.each do |key|
        key.revoke!(reason: "User suspended by admin")
      end

      redirect_to admin_user_path(@user), notice: "User suspended successfully. All API keys have been revoked."
    else
      redirect_to admin_user_path(@user), alert: "Failed to suspend user."
    end
  end

  def unsuspend
    if @user.update(status: "active", active: true)
      redirect_to admin_user_path(@user), notice: "User unsuspended successfully."
    else
      redirect_to admin_user_path(@user), alert: "Failed to unsuspend user."
    end
  end

  def ban
    if @user.update(status: "banned", banned_at: Time.current, active: false)
      # Revoke all API keys
      @user.api_keys.each do |key|
        key.revoke!(reason: "User banned by admin")
      end

      # Cancel subscription
      if @user.subscription
        @user.subscription.update!(
          cancel_at_period_end: true,
          canceled_at: Time.current
        )
      end

      redirect_to admin_user_path(@user), notice: "User banned successfully. All API keys revoked and subscription canceled."
    else
      redirect_to admin_user_path(@user), alert: "Failed to ban user."
    end
  end

  def make_admin
    if @user.update(admin: true)
      redirect_to admin_user_path(@user), notice: "User granted admin privileges."
    else
      redirect_to admin_user_path(@user), alert: "Failed to grant admin privileges."
    end
  end

  def remove_admin
    if @user.id == current_user.id
      redirect_to admin_user_path(@user), alert: "You cannot remove your own admin privileges."
      return
    end

    if @user.update(admin: false)
      redirect_to admin_user_path(@user), notice: "Admin privileges removed from user."
    else
      redirect_to admin_user_path(@user), alert: "Failed to remove admin privileges."
    end
  end

  private

  def require_admin!
    unless current_user.admin?
      redirect_to root_path, alert: t("admin.access_denied")
    end
  end

  def set_user
    @user = User.find(params[:id])
  rescue ActiveRecord::RecordNotFound
    redirect_to admin_users_path, alert: "User not found."
  end

  def calculate_user_usage_stats
    {
      total_requests: @user.usage_logs.count,
      requests_this_month: @user.usage_logs.where("used_at >= ?", Time.current.beginning_of_month).count,
      total_requests_used: @user.usage_logs.sum(:credits_used),
      avg_response_time: @user.usage_logs.where.not(response_time_ms: nil).average(:response_time_ms)&.round || 0,
      error_rate: calculate_user_error_rate
    }
  end

  def calculate_user_error_rate
    UsageLog.error_rate_for(@user.usage_logs)
  end
end
