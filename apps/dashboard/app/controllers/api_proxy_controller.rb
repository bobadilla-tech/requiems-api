# frozen_string_literal: true

class ApiProxyController < ApplicationController
  # Rate limiting handled by Rack::Attack (anonymous: 10/min, authenticated: 30/min)

  def create
    endpoint = params[:endpoint]
    method = params[:method]&.upcase || "GET"
    request_params = params[:params] || {}

    unless valid_endpoint?(endpoint)
      return render json: {
        error: "Invalid endpoint",
        message: t('api_proxy_controller.the_endpoint_must_start_with_v1')
      }, status: :bad_request
    end

    start_time = Time.current
    result = ApiProxyService.call(
      endpoint: endpoint,
      method: method,
      params: request_params,
      forwarded_for: request.headers["CF-Connecting-IP"] || request.remote_ip
    )
    response_time = ((Time.current - start_time) * 1000).round

    render json: {
      status_code: result.status_code,
      response_time_ms: response_time,
      data: result.data,
      error: result.error
    }, status: result.status_code
  rescue StandardError => e
    Rails.logger.error("API Proxy Error: #{e.message}")
    Rails.logger.error(e.backtrace.join("\n"))

    render json: {
      error: "Proxy error",
      message: "Failed to connect to API: #{e.message}"
    }, status: :internal_server_error
  end

  private

  def valid_endpoint?(endpoint)
    return false if endpoint.blank?
    return false if endpoint.include?("..")

    endpoint.match?(/\A\/v1\/[a-zA-Z0-9\/\-_.]+\z/)
  end
end
