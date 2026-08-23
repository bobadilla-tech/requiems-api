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
    domain = normalize_domain(params[:domain])
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

  def useragent
    ua = params[:ua].to_s.strip
    return render_demo_error("useragent", t("tools.useragent.demo.error_empty")) if ua.blank?

    result = api_call(endpoint: "/v1/technology/useragent", method: "GET", params: { ua: ua })

    return render_demo_error("useragent", t("tools.useragent.demo.error_rate_limit")) if result.status_code == 429
    return render_demo_error("useragent", t("tools.useragent.demo.error_generic")) unless result.status_code == 200

    data = result.data&.dig("data", "data") || result.data&.dig("data")
    return render_demo_error("useragent", t("tools.useragent.demo.error_no_data")) if data.nil?

    render "tool_demos/useragent", locals: { data: data }
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

  def sudoku
    difficulty = params[:difficulty].to_s.strip
    difficulty = "medium" if difficulty.blank?

    unless %w[easy medium hard].include?(difficulty)
      return render_demo_error("sudoku", t("tools.sudoku.demo.error_invalid_difficulty"))
    end

    result = api_call(endpoint: "/v1/entertainment/sudoku", method: "GET", params: { difficulty: difficulty })

    return render_demo_error("sudoku", t("tools.sudoku.demo.error_rate_limit")) if result.status_code == 429
    return render_demo_error("sudoku", t("tools.sudoku.demo.error_generic")) unless result.status_code == 200

    data = result.data&.dig("data", "data") || result.data&.dig("data")
    return render_demo_error("sudoku", t("tools.sudoku.demo.error_no_data")) if data.nil?

    render "tool_demos/sudoku", locals: { data: data }
  end

  def number_base_conversion
    from  = params[:from].to_s.strip
    to    = params[:to].to_s.strip
    value = params[:value].to_s.strip

    if from.blank? || to.blank? || value.blank?
      return render_demo_error("number_base_conversion", t("tools.number_base_conversion.demo.error_fill_all_fields"))
    end

    result = api_call(endpoint: "/v1/technology/base", method: "GET",
                      params: { from: from, to: to, value: value })

    if result.status_code == 429
      return render_demo_error("number_base_conversion", t("tools.number_base_conversion.demo.error_rate_limit"))
    end

    unless result.status_code == 200
      msg = result.data&.dig("data", "message") ||
            result.data&.dig("message") ||
            t("tools.number_base_conversion.demo.error_generic")
      return render_demo_error("number_base_conversion", msg)
    end

    data = result.data&.dig("data", "data") || result.data&.dig("data")
    return render_demo_error("number_base_conversion", t("tools.number_base_conversion.demo.error_no_data")) if data.nil?

    render "tool_demos/number_base_conversion", locals: { data: data }
  end

  def mx_lookup
    domain = normalize_domain(params[:domain])

    if domain.blank?
      return render_demo_error("mx_lookup", t("tools.mx_lookup.demo.error_empty"))
    end

    unless valid_domain?(domain)
      return render_demo_error("mx_lookup", t("tools.mx_lookup.demo.error_invalid"))
    end

    result = api_call(endpoint: "/v1/networking/mx/#{domain}", method: "GET", params: {})

    if result.status_code == 429
      return render_demo_error("mx_lookup", t("tools.mx_lookup.demo.error_rate_limit"))
    end

    if result.status_code == 404
      return render_demo_error("mx_lookup", t("tools.mx_lookup.demo.error_not_found"))
    end

    unless result.status_code == 200
      return render_demo_error("mx_lookup", t("tools.mx_lookup.demo.error_generic"))
    end

    data = result.data&.dig("data", "data") || result.data&.dig("data")
    return render_demo_error("mx_lookup", t("tools.mx_lookup.demo.error_no_data")) if data.nil?

    render "tool_demos/mx_lookup", locals: { data: data }
  end

  def mortgage
    principal = params[:principal].to_s.strip
    rate      = params[:rate].to_s.strip
    years     = params[:years].to_s.strip

    if principal.blank? || rate.blank? || years.blank?
      return render_demo_error("mortgage", t("tools.mortgage.demo.error_empty"))
    end

    principal_f = Float(principal, exception: false)
    rate_f      = Float(rate, exception: false)
    years_i     = Integer(years, exception: false)

    if principal_f.nil? || principal_f <= 0 ||
       rate_f.nil? || rate_f <= 0 ||
       years_i.nil? || years_i < 1 || years_i > 50
      return render_demo_error("mortgage", t("tools.mortgage.demo.error_invalid"))
    end

    result = api_call(
      endpoint: "/v1/finance/mortgage",
      method: "GET",
      params: { principal: principal_f, rate: rate_f, years: years_i }
    )

    if result.status_code == 429
      return render_demo_error("mortgage", t("tools.mortgage.demo.error_rate_limit"))
    end
    unless result.status_code == 200
      return render_demo_error("mortgage", t("tools.mortgage.demo.error_generic"))
    end

    data = result.data&.dig("data", "data") || result.data&.dig("data")
    return render_demo_error("mortgage", t("tools.mortgage.demo.error_no_data")) if data.nil?

    render "tool_demos/mortgage", locals: { data: data }
  end

  def markdown
    markdown_text = params[:markdown].to_s.strip
    return render_demo_error("markdown", t("tools.markdown.demo.error_empty")) if markdown_text.blank?

    result = api_call(endpoint: "/v1/technology/markdown", method: "POST",
                      params: { markdown: markdown_text, sanitize: true })
    return render_demo_error("markdown", t("tools.markdown.demo.error_rate_limit")) if result.status_code == 429
    unless result.status_code == 200
      return render_demo_error("markdown", t("tools.markdown.demo.error_generic"))
    end

    data = result.data&.dig("data", "data") || result.data&.dig("data")
    return render_demo_error("markdown", t("tools.markdown.demo.error_no_data")) if data.nil?

    render "tool_demos/markdown", locals: { data: data }
  end

  def barcode
    data = params[:data].to_s.strip
    type = params[:type].to_s.strip

    if data.blank? || type.blank?
      return render_demo_error("barcode", t("tools.barcode.demo.error_empty"))
    end

    unless %w[code128 code93 code39 ean8 ean13].include?(type)
      return render_demo_error("barcode", t("tools.barcode.demo.error_invalid_type"))
    end

    result = api_call(endpoint: "/v1/technology/barcode/base64", method: "GET",
                      params: { data: data, type: type })

    if result.status_code == 429
      return render_demo_error("barcode", t("tools.barcode.demo.error_rate_limit"))
    end

    unless result.status_code == 200
      msg = result.data&.dig("data", "message") ||
            result.data&.dig("message") ||
            t("tools.barcode.demo.error_generic")
      return render_demo_error("barcode", msg)
    end

    data_out = result.data&.dig("data", "data") || result.data&.dig("data")
    return render_demo_error("barcode", t("tools.barcode.demo.error_no_data")) if data_out.nil?

    render "tool_demos/barcode", locals: { data: data_out }
  end

  def advice
    result = api_call(endpoint: "/v1/entertainment/advice", method: "GET", params: {})

    return render_demo_error("advice", t("tools.advice.demo.error_rate_limit")) if result.status_code == 429
    return render_demo_error("advice", t("tools.advice.demo.error_generic")) unless result.status_code == 200

    data = result.data&.dig("data", "data") || result.data&.dig("data")
    return render_demo_error("advice", t("tools.advice.demo.error_no_data")) if data.nil?

    render "tool_demos/advice", locals: { data: data }
  end

  def base64
    mode    = params[:mode].to_s.strip
    value   = params[:value].to_s.strip
    variant = params[:variant].to_s.strip
    variant = "standard" unless %w[standard url].include?(variant)

    return render_demo_error("base64", t("tools.base64.demo.error_empty")) if value.blank?

    unless %w[encode decode].include?(mode)
      return render_demo_error("base64", t("tools.base64.demo.error_invalid_mode"))
    end

    endpoint = mode == "encode" ? "/v1/technology/base64/encode" : "/v1/technology/base64/decode"
    result = api_call(endpoint: endpoint, method: "POST", params: { value: value, variant: variant })

    return render_demo_error("base64", t("tools.base64.demo.error_rate_limit")) if result.status_code == 429

    if result.status_code == 422
      return render_demo_error("base64", t("tools.base64.demo.error_invalid_base64"))
    end

    unless result.status_code == 200
      return render_demo_error("base64", t("tools.base64.demo.error_generic"))
    end

    data = result.data&.dig("data", "data") || result.data&.dig("data")
    return render_demo_error("base64", t("tools.base64.demo.error_no_data")) if data.nil?

    render "tool_demos/base64", locals: { data: data, mode: mode }
  end

  def whois
    domain = normalize_domain(params[:domain])
    return render_demo_error("whois", t("tools.whois.demo.error_empty")) if domain.blank?

    unless valid_domain?(domain)
      return render_demo_error("whois", t("tools.whois.demo.error_invalid"))
    end

    result = api_call(endpoint: "/v1/networking/whois/#{domain}", method: "GET", params: {})

    return render_demo_error("whois", t("tools.whois.demo.error_rate_limit")) if result.status_code == 429
    return render_demo_error("whois", t("tools.whois.demo.error_not_found")) if result.status_code == 404
    return render_demo_error("whois", t("tools.whois.demo.error_generic")) unless result.status_code == 200

    data = result.data&.dig("data", "data") || result.data&.dig("data")
    return render_demo_error("whois", t("tools.whois.demo.error_no_data")) if data.nil?

    render "tool_demos/whois", locals: { data: data }
  end

  def lorem_ipsum
    paragraphs = params[:paragraphs].to_s.strip
    sentences  = params[:sentences].to_s.strip

    paragraphs_i = paragraphs.blank? ? 1 : Integer(paragraphs, exception: false)
    sentences_i  = sentences.blank? ? 5 : Integer(sentences, exception: false)

    if paragraphs_i.nil? || !paragraphs_i.between?(1, 20) ||
       sentences_i.nil? || !sentences_i.between?(1, 20)
      return render_demo_error("lorem_ipsum", t("tools.lorem_ipsum.demo.error_invalid"))
    end

    result = api_call(endpoint: "/v1/text/lorem", method: "GET",
                      params: { paragraphs: paragraphs_i, sentences: sentences_i })

    return render_demo_error("lorem_ipsum", t("tools.lorem_ipsum.demo.error_rate_limit")) if result.status_code == 429
    return render_demo_error("lorem_ipsum", t("tools.lorem_ipsum.demo.error_generic")) unless result.status_code == 200

    data = result.data&.dig("data", "data") || result.data&.dig("data")
    return render_demo_error("lorem_ipsum", t("tools.lorem_ipsum.demo.error_no_data")) if data.nil?

    render "tool_demos/lorem_ipsum", locals: { data: data }
  end

  def working_days
    from = params[:from].to_s.strip
    to   = params[:to].to_s.strip
    country = params[:country].to_s.strip.upcase
    subdivision = params[:subdivision].to_s.strip.upcase

    if from.blank? || to.blank?
      return render_demo_error("working_days", t("tools.working_days.demo.error_empty"))
    end

    from_date = parse_iso_date(from)
    to_date   = parse_iso_date(to)

    if from_date.nil? || to_date.nil? || to_date < from_date
      return render_demo_error("working_days", t("tools.working_days.demo.error_invalid"))
    end

    query = { from: from, to: to }
    query[:country] = country if country.present?
    query[:subdivision] = subdivision if subdivision.present?

    result = api_call(endpoint: "/v1/places/working-days", method: "GET", params: query)

    return render_demo_error("working_days", t("tools.working_days.demo.error_rate_limit")) if result.status_code == 429
    return render_demo_error("working_days", t("tools.working_days.demo.error_invalid")) if result.status_code == 400
    return render_demo_error("working_days", t("tools.working_days.demo.error_generic")) unless result.status_code == 200

    data = result.data&.dig("data", "data") || result.data&.dig("data")
    return render_demo_error("working_days", t("tools.working_days.demo.error_no_data")) if data.nil?

    render "tool_demos/working_days", locals: { data: data }
  end

  def world_time
    timezone_name = params[:timezone].to_s.strip
    return render_demo_error("world_time", t("tools.world_time.demo.error_empty")) if timezone_name.blank?

    result = api_call(endpoint: "/v1/places/time/#{timezone_name}", method: "GET", params: {})

    return render_demo_error("world_time", t("tools.world_time.demo.error_rate_limit")) if result.status_code == 429
    return render_demo_error("world_time", t("tools.world_time.demo.error_not_found")) if result.status_code == 404
    return render_demo_error("world_time", t("tools.world_time.demo.error_generic")) unless result.status_code == 200

    data = result.data&.dig("data", "data") || result.data&.dig("data")
    return render_demo_error("world_time", t("tools.world_time.demo.error_no_data")) if data.nil?

    render "tool_demos/world_time", locals: { data: data, timezone: timezone_name }
  end

  def words
    result = api_call(endpoint: "/v1/text/words/random", method: "GET", params: {})

    return render_demo_error("words", t("tools.words.demo.error_rate_limit")) if result.status_code == 429
    return render_demo_error("words", t("tools.words.demo.error_generic")) unless result.status_code == 200

    data = result.data&.dig("data", "data") || result.data&.dig("data")
    return render_demo_error("words", t("tools.words.demo.error_no_data")) if data.nil?

    render "tool_demos/words", locals: { data: data }
  end

  private

  def valid_ip?(ip)
    return false if ip.include?("/")

    IPAddr.new(ip)
    true
  rescue IPAddr::Error
    false
  end

  def valid_domain?(domain)
    domain.match?(/\A[a-z0-9]([a-z0-9\-]{0,61}[a-z0-9])?(\.[a-z0-9]([a-z0-9\-]{0,61}[a-z0-9])?)+\z/i)
  end

  def parse_iso_date(str)
    Date.iso8601(str)
  rescue ArgumentError, TypeError
    nil
  end

  def normalize_domain(domain)
    domain = domain.to_s.strip.downcase
    domain = domain.sub(/\Ahttps?:\/\//, "")  # strip protocol
    domain = domain.split("/", 2).first.to_s   # strip path
    domain = domain.split("?", 2).first.to_s   # strip query
    domain = domain.split("#", 2).first.to_s   # strip fragment
    domain = domain.split(":", 2).first.to_s   # strip port
    domain.strip
  end

  def api_call(endpoint:, method:, params:)
    ApiProxyService.call(
      endpoint: endpoint,
      method: method,
      params: params,
      forwarded_for: TrustedProxy.client_ip(request)
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
