# frozen_string_literal: true

require "test_helper"

class TrustedProxyTest < ActiveSupport::TestCase
  def fake_request(remote_addr:, x_forwarded_for: nil)
    headers = x_forwarded_for ? { "X-Forwarded-For" => x_forwarded_for } : {}
    ActionDispatch::TestRequest.create("REMOTE_ADDR" => remote_addr).tap do |req|
      headers.each { |k, v| req.set_header("HTTP_#{k.upcase.tr('-', '_')}", v) }
    end
  end

  test "honors X-Forwarded-For when remote_addr is within a Cloudflare range" do
    request = fake_request(remote_addr: "104.16.1.1", x_forwarded_for: "203.0.113.7")
    assert_equal "203.0.113.7", TrustedProxy.client_ip(request)
  end

  test "honors X-Forwarded-For when remote_addr is a private/loopback proxy hop" do
    request = fake_request(remote_addr: "172.20.0.5", x_forwarded_for: "203.0.113.7")
    assert_equal "203.0.113.7", TrustedProxy.client_ip(request)
  end

  test "ignores a spoofed X-Forwarded-For from an untrusted remote_addr" do
    request = fake_request(remote_addr: "203.0.113.55", x_forwarded_for: "6.6.6.6")
    assert_equal "203.0.113.55", TrustedProxy.client_ip(request)
  end

  test "ignores a spoofed X-Forwarded-For from an unconfigured private peer" do
    request = fake_request(remote_addr: "10.0.0.8", x_forwarded_for: "6.6.6.6")
    assert_equal "10.0.0.8", TrustedProxy.client_ip(request)
  end

  test "falls back to remote_addr when no forwarded header is present, even if trusted" do
    request = fake_request(remote_addr: "172.20.0.5")
    assert_equal "172.20.0.5", TrustedProxy.client_ip(request)
  end

  test "takes the first hop when X-Forwarded-For has multiple entries" do
    request = fake_request(remote_addr: "172.20.0.5", x_forwarded_for: "203.0.113.7, 172.20.0.5")
    assert_equal "203.0.113.7", TrustedProxy.client_ip(request)
  end
end
