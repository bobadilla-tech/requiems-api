# frozen_string_literal: true

require "test_helper"
require Rails.root.join("script/golden_diff_api_docs").to_s

class GoldenDiffApiDocsTest < ActiveSupport::TestCase
  test "compare_curls matches equivalent query, JSON, and multiline formatting" do
    generated = <<~CURL
      curl -X POST "https://api.requiems.xyz/v1/test?b=two&a=one" \
        -H "Content-Type: application/json" \
        -H "requiems-api-key: YOUR_API_KEY" \
        -d '{"a":1,"nested":{"enabled":true}}'
    CURL
    handwritten = <<~CURL
      curl -X POST 'https://api.requiems.xyz/v1/test?a=one&b=two' \
        -H 'requiems-api-key: YOUR_API_KEY' \
        -H 'Content-Type: application/json' \
        -d '{"nested":{"enabled":true},"a":1}'
    CURL

    assert_equal :match, compare_curls(generated, handwritten)
    assert_equal "SAFE", classify_curl_result(
      :match, invocation_count: 1, response_kind: "json"
    ).first
  end

  test "compare_curls reports structural mismatches for manual override" do
    generated = 'curl -X POST "https://api.requiems.xyz/v1/test" -d \'{"value":1}\''
    handwritten = 'curl -X PUT "https://api.requiems.xyz/v1/test" -d \'{"value":1}\''

    assert_equal :mismatch, compare_curls(generated, handwritten)
    assert_equal "MANUAL_OVERRIDE", classify_curl_result(
      :mismatch, invocation_count: 1, response_kind: "json"
    ).first
  end

  test "compare_curls reports unparseable snippets for needs review" do
    generated = "curl"
    handwritten = 'curl "https://api.requiems.xyz/v1/test"'

    assert_equal :parse_failure, compare_curls(generated, handwritten)
    assert_equal "NEEDS_REVIEW", classify_curl_result(
      :parse_failure, invocation_count: 1, response_kind: "json"
    ).first
  end
end
