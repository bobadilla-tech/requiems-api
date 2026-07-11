# frozen_string_literal: true

require "test_helper"

class ApiProxyServiceTest < ActiveSupport::TestCase
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
