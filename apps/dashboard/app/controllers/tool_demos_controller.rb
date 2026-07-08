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

  def domain_checker
    domain = params[:domain].to_s.strip.downcase
    domain = domain.sub(/\Ahttps?:\/\//, "")  # strip protocol
    domain = domain.split("/", 2).first.to_s   # strip path
    domain = domain.split("?", 2).first.to_s   # strip query
    domain = domain.split("#", 2).first.to_s   # strip fragment
    domain = domain.split(":", 2).first.to_s   # strip port
    domain = domain.strip
    return render_demo_error("domain_checker", t("tools.domain_checker.demo.error_empty")) if domain.blank?

    unless domain.match?(/\A[a-z0-9]([a-z0-9\-]{0,61}[a-z0-9])?(\.[a-z0-9]([a-z0-9\-]{0,61}[a-z0-9])?)+\z/)
      return render_demo_error("domain_checker", t("tools.domain_checker.demo.error_invalid"))
    end

    result = api_call(endpoint: "/v1/networking/domain/#{domain}", method: "GET", params: {})

    return render_demo_error("domain_checker", t("tools.domain_checker.demo.error_rate_limit")) if result.status_code == 429
    return render_demo_error("domain_checker", t("tools.domain_checker.demo.error_generic")) unless result.status_code == 200

    data = result.data&.dig("data", "data") || result.data&.dig("data")
    return render_demo_error("domain_checker", t("tools.domain_checker.demo.error_no_data")) if data.nil?

    render "tool_demos/domain_checker", locals: { data: data }
  end

  def phone_validator
    number = params[:phone].to_s.strip

    if number.blank?
      return render_demo_error("phone_validator", t("tools.phone_validator.demo.error_empty"))
    end

    result = api_call(endpoint: "/v1/validation/phone", method: "GET", params: { number: number })

    return render_demo_error("phone_validator", t("tools.phone_validator.demo.error_rate_limit")) if result.status_code == 429
    return render_demo_error("phone_validator", t("tools.phone_validator.demo.error_generic")) unless result.status_code == 200

    data = result.data&.dig("data", "data") || result.data&.dig("data")
    return render_demo_error("phone_validator", t("tools.phone_validator.demo.error_no_data")) if data.nil?

    render "tool_demos/phone_validator", locals: { data: data, number: number }
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

  def inflation
    country = params[:country].to_s.strip.upcase

    if country.blank?
      return render_demo_error("inflation", t("tools.inflation.demo.error_empty"))
    end

    unless country.match?(/\A[A-Z]{2}\z/)
      return render_demo_error("inflation", t("tools.inflation.demo.error_invalid"))
    end

    result = api_call(endpoint: "/v1/finance/inflation", method: "GET", params: { country: country })

    if result.status_code == 429
      return render_demo_error("inflation", t("tools.inflation.demo.error_rate_limit"))
    end

    if result.status_code == 400
      return render_demo_error("inflation", t("tools.inflation.demo.error_invalid"))
    end

    if result.status_code == 404
      return render_demo_error("inflation", t("tools.inflation.demo.error_no_data"))
    end

    unless result.status_code == 200
      return render_demo_error("inflation", t("tools.inflation.demo.error_generic"))
    end

    data = result.data&.dig("data", "data") || result.data&.dig("data")
    return render_demo_error("inflation", t("tools.inflation.demo.error_no_data")) if data.nil?

    render "tool_demos/inflation", locals: { data: data }
  end

  def qr_code
    data = params[:data].to_s.strip

    if data.blank?
      return render_demo_error("qr_code", t("tools.qr_code.demo.error_empty"))
    end

    result = api_call(endpoint: "/v1/technology/qr/base64", method: "GET", params: { data: data })

    if result.status_code == 429
      return render_demo_error("qr_code", t("tools.qr_code.demo.error_rate_limit"))
    end

    unless result.status_code == 200
      return render_demo_error("qr_code", t("tools.qr_code.demo.error_generic"))
    end

    qr_data = result.data&.dig("data", "data") || result.data&.dig("data")
    return render_demo_error("qr_code", t("tools.qr_code.demo.error_no_data")) if qr_data.nil?

    render "tool_demos/qr_code", locals: { data: qr_data }
  end

  def bin_lookup
    bin = params[:bin].to_s.strip

    if bin.blank?
      return render_demo_error("bin_lookup", t("tools.bin_lookup.demo.error_empty"))
    end

    unless bin.match?(/\A\d{6,8}\z/)
      return render_demo_error("bin_lookup", t("tools.bin_lookup.demo.error_invalid"))
    end

    result = api_call(endpoint: "/v1/finance/bin/#{bin}", method: "GET", params: {})

    if result.status_code == 429
      return render_demo_error("bin_lookup", t("tools.bin_lookup.demo.error_rate_limit"))
    end

    if result.status_code == 404
      return render_demo_error("bin_lookup", t("tools.bin_lookup.demo.error_no_data"))
    end

    unless result.status_code == 200
      return render_demo_error("bin_lookup", t("tools.bin_lookup.demo.error_generic"))
    end

    data = result.data&.dig("data", "data") || result.data&.dig("data")
    return render_demo_error("bin_lookup", t("tools.bin_lookup.demo.error_no_data")) if data.nil?

    render "tool_demos/bin_lookup", locals: { data: data, bin: bin }
  end

  def profanity_filter
    text = params[:text].to_s.strip

    if text.blank?
      return render_demo_error("profanity_filter", t("tools.profanity_filter.demo.error_empty"))
    end

    result = api_call(endpoint: "/v1/validation/profanity", method: "POST", params: { text: text })

    if result.status_code == 429
      return render_demo_error("profanity_filter", t("tools.profanity_filter.demo.error_rate_limit"))
    end

    unless result.status_code == 200
      return render_demo_error("profanity_filter", t("tools.profanity_filter.demo.error_generic"))
    end

    data = result.data&.dig("data", "data") || result.data&.dig("data")
    return render_demo_error("profanity_filter", t("tools.profanity_filter.demo.error_no_data")) if data.nil?

    render "tool_demos/profanity_filter", locals: { data: data, text: text }
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
