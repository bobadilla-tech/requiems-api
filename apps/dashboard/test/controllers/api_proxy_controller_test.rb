# frozen_string_literal: true

require "test_helper"

class ApiProxyControllerTest < ActionDispatch::IntegrationTest
  def stub_proxy(status_code, data)
    result = ApiProxyService::Result.new(status_code: status_code, data: data, error: nil)
    original = ApiProxyService.method(:call)
    ApiProxyService.define_singleton_method(:call) { |**_kwargs| result }
    yield
  ensure
    ApiProxyService.define_singleton_method(:call, original)
  end

  test "returns 400 when endpoint is blank" do
    post "/api/proxy", params: { endpoint: "", method: "GET" }
    assert_response :bad_request
    body = JSON.parse(response.body)
    assert_equal "Invalid endpoint", body["error"]
    assert_equal I18n.t("api_proxy_controller.errors.invalid_endpoint"), body["message"]
  end

  test "returns 400 when endpoint does not start with /v1/" do
    post "/api/proxy", params: { endpoint: "/v2/something", method: "GET" }
    assert_response :bad_request
    body = JSON.parse(response.body)
    assert_equal "Invalid endpoint", body["error"]
  end

  test "returns 400 when endpoint contains path traversal" do
    post "/api/proxy", params: { endpoint: "/v1/../secret", method: "GET" }
    assert_response :bad_request
  end

  test "proxies valid endpoint and returns upstream response" do
    stub_proxy(200, { "result" => "ok" }) do
      post "/api/proxy", params: { endpoint: "/v1/email/validate", method: "GET" }
    end
    assert_response :success
    body = JSON.parse(response.body)
    assert_equal 200, body["status_code"]
  end

  test "proxies valid endpoint with POST method" do
    stub_proxy(200, { "data" => "value" }) do
      post "/api/proxy", params: { endpoint: "/v1/text/sentiment", method: "POST", params: { text: "hello" } }
    end
    assert_response :success
  end
end
