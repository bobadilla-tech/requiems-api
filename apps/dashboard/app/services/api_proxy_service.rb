# frozen_string_literal: true

require "base64"

# Shared HTTP client for calling the internal Requiems API.
# Used by ApiProxyController (the public playground) and ToolDemosController
# (server-side demo form submissions that render Turbo Frame responses).
class ApiProxyService
  Result = Data.define(:status_code, :data, :error)

  def self.call(endpoint:, method:, params:, forwarded_for:)
    new(endpoint, method.to_s.upcase, params, forwarded_for).call
  end

  def initialize(endpoint, method, params, forwarded_for)
    @endpoint = endpoint
    @method = method
    @params = params.respond_to?(:to_unsafe_h) ? params.to_unsafe_h : params.to_h
    @forwarded_for = forwarded_for
  end

  def call
    raise ArgumentError, "Invalid endpoint: #{@endpoint}" unless valid_endpoint?

    base = URI(::AppConfig.internal_api_url)
    uri_class = base.is_a?(URI::HTTPS) ? URI::HTTPS : URI::HTTP
    uri = uri_class.build(host: base.host, port: base.port, path: @endpoint)

    Rails.logger.debug { "ApiProxyService: #{@method} #{base.host}#{@endpoint} params=#{@params.inspect}" }

    headers = {
      "X-Backend-Secret" => ::AppConfig.backend_secret,
      "Content-Type" => "application/json",
      "User-Agent" => "Requiems-Playground/1.0",
      "X-Forwarded-For" => @forwarded_for
    }

    http = Net::HTTP.new(uri.host, uri.port)
    http.use_ssl = uri.scheme == "https"
    http.open_timeout = 10
    http.read_timeout = 30

    request = build_request(uri)
    headers.each { |k, v| request[k] = v }

    response = http.request(request) # codeql[rb/request-forgery] host/port/scheme fixed from AppConfig; path validated above

    Result.new(
      status_code: response.code.to_i,
      data: parse_body(response),
      error: response.is_a?(Net::HTTPSuccess) ? nil : response.message
    )
  rescue Net::OpenTimeout, Net::ReadTimeout => e
    Result.new(status_code: 504, data: nil, error: "Request timeout: #{e.message}")
  rescue StandardError => e
    Rails.logger.error("ApiProxyService error: #{e.class}: #{e.message}")
    Rails.logger.error(e.backtrace.first(3).join("\n"))
    Result.new(status_code: 500, data: nil, error: "Request failed: #{e.message}")
  end

  private

  def valid_endpoint?
    return false if @endpoint.blank?
    return false if @endpoint.include?("..")

    @endpoint.match?(/\A\/v1\/[a-zA-Z0-9\/\-_.]+\z/)
  end

  def build_request(uri)
    case @method
    when "GET"
      uri.query = URI.encode_www_form(@params) if @params.any?
      Net::HTTP::Get.new(uri)
    when "POST"
      req = Net::HTTP::Post.new(uri)
      req.body = @params.to_json
      req
    else
      raise "Unsupported HTTP method: #{@method}"
    end
  end

  def parse_body(response)
    content_type = response["Content-Type"].to_s
    body = response.body
    return nil if body.blank?

    if content_type.start_with?("image/", "application/octet-stream")
      return { "type" => "image", "content_type" => content_type, "base64" => Base64.strict_encode64(body) }
    end

    JSON.parse(body)
  rescue JSON::ParserError
    body.encode("UTF-8", invalid: :replace, undef: :replace)
  end
end
