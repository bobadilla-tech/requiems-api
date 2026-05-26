# frozen_string_literal: true

class Admin::AbuseReportsController < ApplicationController
  before_action :authenticate_user!
  before_action :require_admin!
  before_action :set_abuse_report, only: [ :show, :resolve, :investigate ]
  layout "admin"

  def index
    @abuse_reports = AbuseReport.includes(:user, :api_key).order(created_at: :desc)

    # Filter by status
    if params[:status].present? && params[:status] != "all"
      @abuse_reports = @abuse_reports.where(status: params[:status])
    end

    # Filter by type
    if params[:report_type].present? && params[:report_type] != "all"
      @abuse_reports = @abuse_reports.where(report_type: params[:report_type])
    end

    # Search by user email or API key
    if params[:search].present?
      search_term = "%#{params[:search]}%"
      @abuse_reports = @abuse_reports.joins(:user)
        .where("users.email ILIKE ? OR abuse_reports.description ILIKE ?", search_term, search_term)
    end

    # Paginate
    @pagy, @abuse_reports = pagy(@abuse_reports, limit: 20)

    # Statistics
    counts_by_status = AbuseReport.group(:status).count
    @total_reports = counts_by_status.values.sum
    @pending_reports = counts_by_status["pending"].to_i
    @investigating_reports = counts_by_status["investigating"].to_i
    @resolved_reports = counts_by_status["resolved"].to_i
  end

  def show
    @user = @abuse_report.user
    @api_key = @abuse_report.api_key
    @resolver = User.find(@abuse_report.resolved_by_id) if @abuse_report.resolved_by_id.present?

    # Get user's other abuse reports
    @other_reports = @user.abuse_reports.where.not(id: @abuse_report.id).order(created_at: :desc).limit(5)

    # Get user's API usage statistics
    @usage_stats = {
      total_requests: @user.usage_logs.count,
      requests_this_month: @user.usage_logs.where("used_at >= ?", Time.current.beginning_of_month).count,
      error_rate: calculate_error_rate(@user),
      last_request_at: @user.usage_logs.maximum(:used_at)
    }
  end

  def investigate
    if @abuse_report.update(status: "investigating")
      redirect_to admin_abuse_report_path(@abuse_report), notice: t("admin.abuse_reports.investigate_success")
    else
      redirect_to admin_abuse_report_path(@abuse_report), alert: t("admin.abuse_reports.investigate_error")
    end
  end

  def resolve
    if @abuse_report.update(
      status: "resolved",
      resolved_at: Time.current,
      resolved_by_id: current_user.id
    )
      redirect_to admin_abuse_report_path(@abuse_report), notice: t("admin.abuse_reports.resolve_success")
    else
      redirect_to admin_abuse_report_path(@abuse_report), alert: t("admin.abuse_reports.resolve_error")
    end
  end

  private

  def require_admin!
    unless current_user.admin?
      redirect_to root_path, alert: t("admin.access_denied")
    end
  end

  def set_abuse_report
    @abuse_report = AbuseReport.find(params[:id])
  rescue ActiveRecord::RecordNotFound
    redirect_to admin_abuse_reports_path, alert: t("admin.abuse_reports.not_found")
  end

  def calculate_error_rate(user)
    UsageLog.error_rate_for(user.usage_logs)
  end
end
