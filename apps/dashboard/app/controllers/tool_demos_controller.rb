# frozen_string_literal: true

# Handles server-side demo form submissions for tool pages.
# Each action calls the internal API via ApiProxyService, then renders a
# Turbo Frame partial — eliminating innerHTML manipulation in JS controllers.
class ToolDemosController < ApplicationController
  layout false

  def unit_conversion
    from  = params[:from].to_s.strip
    to    = params[:to].to_s.strip
    value = params[:value].to_s.strip

    if from.blank? || to.blank? || value.blank?
      return render_demo_error("unit_conversion", t("tools.unit_conversion.demo.error_fill_all_fields"))
    end

    result = api_call(endpoint: "/v1/technology/convert", method: "GET",
                      params: { from: from, to: to, value: value })

    if result.status_code == 429
      return render_demo_error("unit_conversion", t("tools.unit_conversion.demo.error_rate_limit"))
    end

    unless result.status_code == 200
      msg = result.data&.dig("data", "message") ||
            result.data&.dig("message") ||
            t("tools.unit_conversion.demo.error_generic")
      return render_demo_error("unit_conversion", msg)
    end

    data = result.data&.dig("data", "data") || result.data&.dig("data")
    return render_demo_error("unit_conversion", t("tools.unit_conversion.demo.error_no_data")) if data.nil?

    render "tool_demos/unit_conversion", locals: { data: data }
  end

  def sentiment_analysis
    text = params[:text].to_s.strip

    if text.blank?
      return render_demo_error("sentiment_analysis", t("tools.sentiment_analysis.demo.error_empty"))
    end

    result = api_call(endpoint: "/v1/text/sentiment", method: "POST", params: { text: text })

    if result.status_code == 429
      return render_demo_error("sentiment_analysis", t("tools.sentiment_analysis.demo.error_rate_limit"))
    end

    unless result.status_code == 200
      return render_demo_error("sentiment_analysis", t("tools.sentiment_analysis.demo.error_generic"))
    end

    data = result.data&.dig("data", "data") || result.data&.dig("data")
    return render_demo_error("sentiment_analysis", t("tools.sentiment_analysis.demo.error_no_data")) if data.nil?

    render "tool_demos/sentiment_analysis", locals: { data: data }
  end

  def email_normalizer
    email = params[:email].to_s.strip

    if email.blank?
      return render_demo_error("email_normalizer", t("tools.email_normalizer.demo.error_empty"))
    end

    result = api_call(endpoint: "/v1/text/normalize", method: "POST", params: { email: email })

    if result.status_code == 429
      return render_demo_error("email_normalizer", t("tools.email_normalizer.demo.error_rate_limit"))
    end

    unless result.status_code == 200
      return render_demo_error("email_normalizer", t("tools.email_normalizer.demo.error_generic"))
    end

    data = result.data&.dig("data", "data") || result.data&.dig("data")
    return render_demo_error("email_normalizer", t("tools.email_normalizer.demo.error_no_data")) if data.nil?

    render "tool_demos/email_normalizer", locals: { data: data }
  end

  def email_validator
    email = params[:email].to_s.strip

    if email.blank?
      return render_demo_error("email_validator", t("tools.email_validator.demo.error_empty"))
    end

    result = api_call(endpoint: "/v1/validation/email", method: "POST", params: { email: email })

    if result.status_code == 429
      return render_demo_error("email_validator", t("tools.email_validator.demo.error_rate_limit"))
    end

    unless result.status_code == 200
      return render_demo_error("email_validator", t("tools.email_validator.demo.error_generic"))
    end

    data = result.data&.dig("data", "data") || result.data&.dig("data")
    return render_demo_error("email_validator", t("tools.email_validator.demo.error_no_data")) if data.nil?

    render "tool_demos/email_validator", locals: { data: data, email: email }
  end

  private

  def api_call(endpoint:, method:, params:)
    ApiProxyService.call(
      endpoint: endpoint,
      method: method,
      params: params,
      forwarded_for: request.headers["CF-Connecting-IP"] || request.remote_ip
    )
  rescue StandardError => e
    Rails.logger.error("ToolDemosController error: #{e.message}")
    ApiProxyService::Result.new(status_code: 500, data: nil, error: e.message)
  end

  def render_demo_error(tool, message)
    render "tool_demos/demo_error", locals: { tool: tool, message: message }
  end
end
