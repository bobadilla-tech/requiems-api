# frozen_string_literal: true

class Admin::ApiKeysController < Admin::BaseController
  before_action :set_api_key, only: [ :show, :revoke ]

  def index
    @api_keys = ApiKey.includes(:user).order(created_at: :desc)

    if params[:search].present?
      search_term = "%#{params[:search]}%"
      @api_keys = @api_keys.joins(:user)
        .where("api_keys.key_prefix ILIKE ? OR users.email ILIKE ? OR api_keys.name ILIKE ?",
               search_term, search_term, search_term)
    end

    @api_keys = @api_keys.where(active: params[:active] == "true") if params[:active].present?

    @pagy, @api_keys = pagy(@api_keys, limit: 25)
  end

  def show
    @user = @api_key.user
    @recent_usage = @api_key.usage_logs.order(used_at: :desc).limit(20)
  end

  def revoke
    if @api_key.revoke!(reason: "Revoked by admin #{current_user.email}")
      redirect_to admin_api_key_path(@api_key), notice: t("admin.api_keys.revoke_success")
    else
      redirect_to admin_api_key_path(@api_key), alert: t("admin.api_keys.revoke_error")
    end
  rescue StandardError => e
    Rails.logger.error "[Admin::ApiKeysController#revoke] #{e.class}: #{e.message}\n#{e.backtrace.first(5).join("\n")}"
    redirect_to admin_api_key_path(@api_key), alert: t("admin.api_keys.revoke_error")
  end

  private

  def set_api_key
    @api_key = ApiKey.find(params[:id])
  rescue ActiveRecord::RecordNotFound
    redirect_to admin_api_keys_path, alert: t("admin.api_keys.not_found")
  end
end
