# frozen_string_literal: true

require "ipaddr"

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

  def timezone
    city = params[:city].to_s.strip
    lat  = params[:lat].to_s.strip
    lon  = params[:lon].to_s.strip

    has_city   = city.present?
    has_coords = lat.present? && lon.present?

    unless has_city || has_coords
      return render_demo_error("timezone", t("tools.timezone.demo.error_empty"))
    end

    if !has_city && !valid_coordinates?(lat, lon)
      return render_demo_error("timezone", t("tools.timezone.demo.error_invalid"))
    end

    query_params = has_city ? { city: city } : { lat: lat, lon: lon }
    result = api_call(endpoint: "/v1/places/timezone", method: "GET", params: query_params)

    if result.status_code == 429
      return render_demo_error("timezone", t("tools.timezone.demo.error_rate_limit"))
    end

    if result.status_code == 404
      return render_demo_error("timezone", t("tools.timezone.demo.error_no_data"))
    end

    unless result.status_code == 200
      return render_demo_error("timezone", t("tools.timezone.demo.error_generic"))
    end

    data = result.data&.dig("data", "data") || result.data&.dig("data")
    return render_demo_error("timezone", t("tools.timezone.demo.error_no_data")) if data.nil?

    label = has_city ? city : "#{lat}, #{lon}"
    render "tool_demos/timezone", locals: { data: data, label: label }
  end

  def trivia
    category = params[:category].to_s.strip
    difficulty = params[:difficulty].to_s.strip

    query = {}
    query[:category] = category if category.present?
    query[:difficulty] = difficulty if difficulty.present?

    result = api_call(endpoint: "/v1/entertainment/trivia", method: "GET", params: query)

    return render_demo_error("trivia", t("tools.trivia.demo.error_rate_limit")) if result.status_code == 429
    return render_demo_error("trivia", t("tools.trivia.demo.error_no_data")) if result.status_code == 404
    return render_demo_error("trivia", t("tools.trivia.demo.error_generic")) unless result.status_code == 200

    data = result.data&.dig("data", "data") || result.data&.dig("data")
    return render_demo_error("trivia", t("tools.trivia.demo.error_no_data")) if data.nil?

    render "tool_demos/trivia", locals: { data: data }
  end

  def vpn_detection
    ip = params[:ip].to_s.strip

    if ip.blank?
      return render_demo_error("vpn_detection", t("tools.vpn_detection.demo.error_empty"))
    end

    unless valid_ip?(ip)
      return render_demo_error("vpn_detection", t("tools.vpn_detection.demo.error_invalid"))
    end

    result = api_call(endpoint: "/v1/networking/ip/vpn/#{ip}", method: "GET", params: {})

    if result.status_code == 429
      return render_demo_error("vpn_detection", t("tools.vpn_detection.demo.error_rate_limit"))
    end

    unless result.status_code == 200
      return render_demo_error("vpn_detection", t("tools.vpn_detection.demo.error_generic"))
    end

    data = result.data&.dig("data", "data") || result.data&.dig("data")
    return render_demo_error("vpn_detection", t("tools.vpn_detection.demo.error_no_data")) if data.nil?

    render "tool_demos/vpn_detection", locals: { data: data, ip: ip }
  end

  def thesaurus
    word = params[:word].to_s.strip.downcase

    if word.blank?
      return render_demo_error("thesaurus", t("tools.thesaurus.demo.error_empty"))
    end

    unless word.match?(/\A[\p{L}'-]+\z/)
      return render_demo_error("thesaurus", t("tools.thesaurus.demo.error_invalid"))
    end

    result = api_call(endpoint: "/v1/text/thesaurus/#{ERB::Util.url_encode(word)}", method: "GET", params: {})

    if result.status_code == 429
      return render_demo_error("thesaurus", t("tools.thesaurus.demo.error_rate_limit"))
    end

    if result.status_code == 404
      return render_demo_error("thesaurus", t("tools.thesaurus.demo.error_no_data"))
    end

    unless result.status_code == 200
      return render_demo_error("thesaurus", t("tools.thesaurus.demo.error_generic"))
    end

    data = result.data&.dig("data", "data") || result.data&.dig("data")
    return render_demo_error("thesaurus", t("tools.thesaurus.demo.error_no_data")) if data.nil?

    render "tool_demos/thesaurus", locals: { data: data }
  end

  def spell_check
    text = params[:text].to_s.strip

    if text.blank?
      return render_demo_error("spell_check", t("tools.spell_check.demo.error_empty"))
    end

    result = api_call(endpoint: "/v1/text/spellcheck", method: "POST", params: { text: text })

    if result.status_code == 429
      return render_demo_error("spell_check", t("tools.spell_check.demo.error_rate_limit"))
    end

    unless result.status_code == 200
      return render_demo_error("spell_check", t("tools.spell_check.demo.error_generic"))
    end

    data = result.data&.dig("data", "data") || result.data&.dig("data")
    return render_demo_error("spell_check", t("tools.spell_check.demo.error_no_data")) if data.nil?

    render "tool_demos/spell_check", locals: { data: data, text: text }
  end

  def random_user
    result = api_call(endpoint: "/v1/technology/random-user", method: "GET", params: {})

    return render_demo_error("random_user", t("tools.random_user.demo.error_rate_limit")) if result.status_code == 429
    return render_demo_error("random_user", t("tools.random_user.demo.error_generic")) unless result.status_code == 200

    data = result.data&.dig("data", "data") || result.data&.dig("data")
    return render_demo_error("random_user", t("tools.random_user.demo.error_no_data")) if data.nil?

    render "tool_demos/random_user", locals: { data: data }
  end

  private

  def valid_ip?(ip)
    return false if ip.include?("/")

    IPAddr.new(ip)
    true
  rescue IPAddr::Error
    false
  end

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

  def valid_coordinates?(lat, lon)
    return false if lat.blank? || lon.blank?

    lat_f = Float(lat, exception: false)
    lon_f = Float(lon, exception: false)
    return false if lat_f.nil? || lon_f.nil?

    lat_f.between?(-90, 90) && lon_f.between?(-180, 180)
  end
end
