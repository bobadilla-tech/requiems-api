# frozen_string_literal: true

require "test_helper"

class ApiProxyServiceTest < ActiveSupport::TestCase
  # Minimal stand-in for Net::HTTP, stubbed at the HTTP boundary (no webmock
  # dependency) so ApiProxyService's real header-assembly code runs.
  class FakeHTTP
    attr_reader :last_request

    def initialize(response)
      @response = response
    end

    def use_ssl=(value); end
    def open_timeout=(value); end
    def read_timeout=(value); end

    def request(req)
      @last_request = req
      @response
    end
  end

  def stubbed_ok_response(body: '{"ok":true}')
    response = Net::HTTPOK.new("1.1", "200", "OK")
    response.instance_variable_set(:@body, body)
    response.define_singleton_method(:body) { @body }
    response
  end

  test "sends requiems-api-key alongside X-Backend-Secret to the Go backend" do
    fake_http = FakeHTTP.new(stubbed_ok_response)

    Net::HTTP.stub :new, fake_http do
      ApiProxyService.call(endpoint: "/v1/entertainment/advice", method: "GET", params: {}, forwarded_for: "1.2.3.4")
    end

    assert_equal AppConfig.playground_api_key, fake_http.last_request["requiems-api-key"]
  end

  test "still sends X-Backend-Secret (additive, not a replacement — Caddy still gates on it)" do
    fake_http = FakeHTTP.new(stubbed_ok_response)

    Net::HTTP.stub :new, fake_http do
      ApiProxyService.call(endpoint: "/v1/entertainment/advice", method: "GET", params: {}, forwarded_for: "1.2.3.4")
    end

    assert_equal AppConfig.backend_secret, fake_http.last_request["X-Backend-Secret"]
  end

  test "valid_endpoint? accepts a colon (IPv6 address embedded in the path)" do
    service = ApiProxyService.new("/v1/networking/ip/vpn/2001:db8::1", "GET", {}, "127.0.0.1")
    assert service.send(:valid_endpoint?)
  end

  test "valid_endpoint? accepts a plain alphanumeric path" do
    service = ApiProxyService.new("/v1/finance/bin/424242", "GET", {}, "127.0.0.1")
    assert service.send(:valid_endpoint?)
  end

  test "valid_endpoint? rejects path traversal" do
    service = ApiProxyService.new("/v1/../secrets", "GET", {}, "127.0.0.1")
    assert_not service.send(:valid_endpoint?)
  end

  test "valid_endpoint? rejects a blank endpoint" do
    service = ApiProxyService.new("", "GET", {}, "127.0.0.1")
    assert_not service.send(:valid_endpoint?)
  end

  test "valid_endpoint? rejects disallowed characters" do
    service = ApiProxyService.new("/v1/foo bar", "GET", {}, "127.0.0.1")
    assert_not service.send(:valid_endpoint?)
  end
end
