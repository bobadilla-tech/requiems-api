# frozen_string_literal: true

class Dashboard::UsageController < Dashboard::BaseController

  def show
    set_date_range

    # Summary statistics
    @total_requests = calculate_total_requests
    @total_requests_used = calculate_total_requests_used
    @avg_response_time = calculate_avg_response_time
    @error_rate = calculate_error_rate

    # Paginated recent requests
    scope = current_user.usage_logs
      .includes(:api_key)
      .where(used_at: @start_date..@end_date)
      .order(used_at: :desc)
    @pagy, @recent_requests = pagy(scope, limit: 25)

    # Data for charts (will be loaded via AJAX)
    # Charts will call by_endpoint, by_date, etc.
  end

  def by_endpoint
    set_date_range

    # Group requests by endpoint and count
    usage_by_endpoint = current_user.usage_logs
      .where(used_at: @start_date..@end_date)
      .group(:endpoint)
      .count

    render json: usage_by_endpoint
  end

  def by_date
    set_date_range

    # Group requests by date
    usage_by_date = current_user.usage_logs
      .where(used_at: @start_date..@end_date)
      .group_by_day(:used_at, time_zone: Time.zone)
      .count

    # Also get error counts by date
    errors_by_date = current_user.usage_logs
      .where(used_at: @start_date..@end_date)
      .where("status_code >= ?", 400)
      .group_by_day(:used_at, time_zone: Time.zone)
      .count

    render json: {
      requests: usage_by_date,
      errors: errors_by_date
    }
  end

  def export
    set_date_range

    # Fetch all usage logs for the date range
    usage_logs = current_user.usage_logs
      .includes(:api_key)
      .where(used_at: @start_date..@end_date)
      .order(used_at: :desc)

    # Generate CSV
    csv_data = generate_csv(usage_logs)

    send_data csv_data,
      filename: "usage-#{@start_date.to_date}-to-#{@end_date.to_date}.csv",
      type: "text/csv"
  end

  private

  def set_date_range
    # Default to last 30 days if no params provided
    case params[:range]
    when "7d"
      @start_date = 7.days.ago
      @end_date = Time.current
      @range_label = "Last 7 Days"
    when "30d"
      @start_date = 30.days.ago
      @end_date = Time.current
      @range_label = "Last 30 Days"
    when "90d"
      @start_date = 90.days.ago
      @end_date = Time.current
      @range_label = "Last 90 Days"
    when "custom"
      @start_date = parse_date_param(params[:start_date])&.beginning_of_day || 30.days.ago
      @end_date = parse_date_param(params[:end_date])&.end_of_day || Time.current
      @range_label = "#{@start_date.to_date} to #{@end_date.to_date}"
    else
      @start_date = 30.days.ago
      @end_date = Time.current
      @range_label = "Last 30 Days"
    end
  end

  def calculate_total_requests
    current_user.usage_logs
      .where(used_at: @start_date..@end_date)
      .count
  end

  def calculate_total_requests_used
    current_user.usage_logs
      .where(used_at: @start_date..@end_date)
      .sum(:credits_used)
  end

  def calculate_avg_response_time
    avg = current_user.usage_logs
      .where(used_at: @start_date..@end_date)
      .where.not(response_time_ms: nil)
      .average(:response_time_ms)

    avg&.round || 0
  end

  def calculate_error_rate
    scope = current_user.usage_logs.where(used_at: @start_date..@end_date)
    UsageLog.error_rate_for(scope)
  end

  DATE_FORMAT = /\A\d{4}-\d{2}-\d{2}\z/

  def parse_date_param(value)
    return nil unless value.is_a?(String) && value.match?(DATE_FORMAT)

    value.to_date
  rescue ArgumentError
    nil
  end

  def generate_csv(usage_logs)
    require "csv"

    CSV.generate(headers: true) do |csv|
      # Header row
      csv << [
        "Timestamp",
        "Endpoint",
        "Method",
        "Status Code",
        "Response Time (ms)",
        "Requests",
        "API Key"
      ]

      # Data rows
      usage_logs.each do |log|
        csv << [
          log.used_at.iso8601,
          log.endpoint,
          log.request_method.presence || "UNKNOWN",
          log.status_code,
          log.response_time_ms,
          log.credits_used || 1,
          log.api_key&.name || "Unknown"
        ]
      end
    end
  end
end
