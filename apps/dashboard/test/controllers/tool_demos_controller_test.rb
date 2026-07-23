# frozen_string_literal: true

require "test_helper"

class ToolDemoControllerTest < ActionDispatch::IntegrationTest
  # ── helpers ─────────────────────────────────────────────────────────────────

  def stub_api(status_code, data)
    result = ApiProxyService::Result.new(status_code: status_code, data: data, error: nil)
    original = ApiProxyService.method(:call)
    ApiProxyService.define_singleton_method(:call) { |*_args, **_kwargs| result }
    yield
  ensure
    ApiProxyService.define_singleton_method(:call, original)
  end

  def success_data(payload)
    { "data" => { "data" => payload } }
  end

  # ── email_normalizer ─────────────────────────────────────────────────────────

  test "email_normalizer renders result on success" do
    payload = { "original" => "Test@Gmail.com", "normalized" => "test@gmail.com",
                "local" => "test", "domain" => "gmail.com", "changes" => [ "lowercased" ] }
    stub_api(200, success_data(payload)) do
      post "/tools/demos/email-normalizer", params: { email: "Test@Gmail.com" }
    end
    assert_response :success
    assert_match "test@gmail.com", response.body
  end

  test "email_normalizer renders error when email is blank" do
    post "/tools/demos/email-normalizer", params: { email: "" }
    assert_response :success
    assert_match I18n.t("tools.email_normalizer.demo.error_empty"), response.body
  end

  test "email_normalizer renders error on 429" do
    stub_api(429, nil) do
      post "/tools/demos/email-normalizer", params: { email: "test@example.com" }
    end
    assert_response :success
    assert_match I18n.t("tools.email_normalizer.demo.error_rate_limit"), response.body
  end

  test "email_normalizer renders error on non-200 response" do
    stub_api(500, nil) do
      post "/tools/demos/email-normalizer", params: { email: "test@example.com" }
    end
    assert_response :success
    assert_match I18n.t("tools.email_normalizer.demo.error_generic"), response.body
  end

  test "email_normalizer renders error when data is nil" do
    stub_api(200, nil) do
      post "/tools/demos/email-normalizer", params: { email: "test@example.com" }
    end
    assert_response :success
    assert_match I18n.t("tools.email_normalizer.demo.error_no_data"), response.body
  end

  # ── email_validator ──────────────────────────────────────────────────────────

  test "email_validator renders result on success" do
    payload = { "valid" => true, "normalized" => "test@example.com", "domain" => "example.com",
                "syntax" => true, "mx_record" => true, "disposable" => false }
    stub_api(200, success_data(payload)) do
      post "/tools/demos/email-validator", params: { email: "test@example.com" }
    end
    assert_response :success
    assert_match "test@example.com", response.body
  end

  test "email_validator renders error when email is blank" do
    post "/tools/demos/email-validator", params: { email: "" }
    assert_response :success
    assert_match I18n.t("tools.email_validator.demo.error_empty"), response.body
  end

  test "email_validator renders error on 429" do
    stub_api(429, nil) do
      post "/tools/demos/email-validator", params: { email: "test@example.com" }
    end
    assert_response :success
    assert_match I18n.t("tools.email_validator.demo.error_rate_limit"), response.body
  end

  test "email_validator renders error on non-200 response" do
    stub_api(500, nil) do
      post "/tools/demos/email-validator", params: { email: "test@example.com" }
    end
    assert_response :success
    assert_match I18n.t("tools.email_validator.demo.error_generic"), response.body
  end

  test "email_validator renders error when data is nil" do
    stub_api(200, nil) do
      post "/tools/demos/email-validator", params: { email: "test@example.com" }
    end
    assert_response :success
    assert_match I18n.t("tools.email_validator.demo.error_no_data"), response.body
  end

  # ── sentiment_analysis ───────────────────────────────────────────────────────

  test "sentiment_analysis renders result on success" do
    payload = { "sentiment" => "positive", "score" => 0.92, "comparative" => 0.46,
                "positive" => [ "great" ], "negative" => [] }
    stub_api(200, success_data(payload)) do
      post "/tools/demos/sentiment-analysis", params: { text: "This is great!" }
    end
    assert_response :success
    assert_match(/positive/i, response.body)
  end

  test "sentiment_analysis renders error when text is blank" do
    post "/tools/demos/sentiment-analysis", params: { text: "" }
    assert_response :success
    assert_match I18n.t("tools.sentiment_analysis.demo.error_empty"), response.body
  end

  test "sentiment_analysis renders error on 429" do
    stub_api(429, nil) do
      post "/tools/demos/sentiment-analysis", params: { text: "hello" }
    end
    assert_response :success
    assert_match I18n.t("tools.sentiment_analysis.demo.error_rate_limit"), response.body
  end

  test "sentiment_analysis renders error on non-200 response" do
    stub_api(500, nil) do
      post "/tools/demos/sentiment-analysis", params: { text: "hello" }
    end
    assert_response :success
    assert_match I18n.t("tools.sentiment_analysis.demo.error_generic"), response.body
  end

  test "sentiment_analysis renders error when data is nil" do
    stub_api(200, nil) do
      post "/tools/demos/sentiment-analysis", params: { text: "hello" }
    end
    assert_response :success
    assert_match I18n.t("tools.sentiment_analysis.demo.error_no_data"), response.body
  end

  # ── domain_checker ───────────────────────────────────────────────────────────

  test "domain_checker renders result on success" do
    payload = { "domain" => "example.com", "available" => false,
                "dns" => { "a" => [ "93.184.216.34" ], "mx" => [], "ns" => [ "a.iana-servers.net" ] } }
    stub_api(200, success_data(payload)) do
      post "/tools/demos/domain-checker", params: { domain: "example.com" }
    end
    assert_response :success
    assert_match "example.com", response.body
  end

  test "domain_checker renders error when domain is blank" do
    post "/tools/demos/domain-checker", params: { domain: "" }
    assert_response :success
    assert_match I18n.t("tools.domain_checker.demo.error_empty"), response.body
  end

  test "domain_checker normalizes URL input" do
    payload = { "domain" => "example.com", "available" => false,
                "dns" => { "a" => [ "93.184.216.34" ], "mx" => [], "ns" => [ "a.iana-servers.net" ] } }
    stub_api(200, success_data(payload)) do
      post "/tools/demos/domain-checker", params: { domain: "https://example.com/some/path?q=1" }
    end
    assert_response :success
    assert_match "example.com", response.body
  end

  test "domain_checker renders error when domain format is invalid" do
    post "/tools/demos/domain-checker", params: { domain: "not-a-domain" }
    assert_response :success
    assert_match I18n.t("tools.domain_checker.demo.error_invalid"), response.body
  end

  test "domain_checker renders error on 429" do
    stub_api(429, nil) do
      post "/tools/demos/domain-checker", params: { domain: "example.com" }
    end
    assert_response :success
    assert_match I18n.t("tools.domain_checker.demo.error_rate_limit"), response.body
  end

  test "domain_checker renders error on non-200 response" do
    stub_api(500, nil) do
      post "/tools/demos/domain-checker", params: { domain: "example.com" }
    end
    assert_response :success
    assert_match I18n.t("tools.domain_checker.demo.error_generic"), response.body
  end

  test "domain_checker renders error when data is nil" do
    stub_api(200, nil) do
      post "/tools/demos/domain-checker", params: { domain: "example.com" }
    end
    assert_response :success
    assert_match I18n.t("tools.domain_checker.demo.error_no_data"), response.body
  end

  # ── phone_validator ──────────────────────────────────────────────────────────

  test "phone_validator renders result on success" do
    payload = { "number" => "+14155552671", "valid" => true, "country" => "US",
                "type" => "mobile", "formatted" => "+1 415-555-2671",
                "carrier" => { "name" => "T-Mobile", "source" => "metadata" },
                "risk" => { "is_voip" => false, "is_virtual" => false } }
    stub_api(200, success_data(payload)) do
      post "/tools/demos/phone-validator", params: { phone: "+14155552671" }
    end
    assert_response :success
    assert_match "+14155552671", response.body
    assert_match "Valid", response.body
  end

  test "phone_validator renders error when phone is blank" do
    post "/tools/demos/phone-validator", params: { phone: "" }
    assert_response :success
    assert_match "Enter a phone number to validate.", response.body
  end

  test "phone_validator renders error on 429" do
    stub_api(429, nil) do
      post "/tools/demos/phone-validator", params: { phone: "+14155552671" }
    end
    assert_response :success
    assert_match "Rate limit reached. Wait a moment and try again.", response.body
  end

  test "phone_validator renders error on non-200 response" do
    stub_api(500, nil) do
      post "/tools/demos/phone-validator", params: { phone: "+14155552671" }
    end
    assert_response :success
    assert_match "Something went wrong. Try again.", response.body
  end

  test "phone_validator renders error when data is nil" do
    stub_api(200, nil) do
      post "/tools/demos/phone-validator", params: { phone: "+14155552671" }
    end
    assert_response :success
    assert_match "No data returned.", response.body
  end

  # ── unit_conversion ──────────────────────────────────────────────────────────

  test "unit_conversion renders result on success" do
    payload = { "result" => 100.0, "from" => "meter", "to" => "centimeter",
                "value" => 1.0, "formula" => "× 100" }
    stub_api(200, success_data(payload)) do
      post "/tools/demos/unit-conversion", params: { from: "meter", to: "centimeter", value: "1" }
    end
    assert_response :success
    assert_match "100", response.body
  end

  test "unit_conversion renders error when any field is blank" do
    post "/tools/demos/unit-conversion", params: { from: "meter", to: "", value: "1" }
    assert_response :success
    assert_match I18n.t("tools.unit_conversion.demo.error_fill_all_fields"), response.body
  end

  test "unit_conversion renders error on 429" do
    stub_api(429, nil) do
      post "/tools/demos/unit-conversion", params: { from: "meter", to: "centimeter", value: "1" }
    end
    assert_response :success
    assert_match I18n.t("tools.unit_conversion.demo.error_rate_limit"), response.body
  end

  test "unit_conversion renders error on non-200 response" do
    stub_api(500, nil) do
      post "/tools/demos/unit-conversion", params: { from: "meter", to: "centimeter", value: "1" }
    end
    assert_response :success
    assert_match I18n.t("tools.unit_conversion.demo.error_generic"), response.body
  end

  test "unit_conversion renders error when data is nil" do
    stub_api(200, nil) do
      post "/tools/demos/unit-conversion", params: { from: "meter", to: "centimeter", value: "1" }
    end
    assert_response :success
    assert_match I18n.t("tools.unit_conversion.demo.error_no_data"), response.body
  end

  # ── inflation ────────────────────────────────────────────────────────────────

  test "inflation renders result on success" do
    payload = {
      "country" => "US", "rate" => 2.9495, "period" => "2024",
      "historical" => [
        { "period" => "2023", "rate" => 4.1163 },
        { "period" => "2022", "rate" => 8.0028 }
      ]
    }
    stub_api(200, success_data(payload)) do
      post "/tools/demos/inflation", params: { country: "US" }
    end
    assert_response :success
    assert_match "2.95", response.body
    assert_match "2024", response.body
  end

  test "inflation renders error when country is blank" do
    post "/tools/demos/inflation", params: { country: "" }
    assert_response :success
    assert_match I18n.t("tools.inflation.demo.error_empty"), response.body
  end

  test "inflation renders error when country code is invalid" do
    post "/tools/demos/inflation", params: { country: "123" }
    assert_response :success
    assert_match I18n.t("tools.inflation.demo.error_invalid"), response.body
  end

  test "inflation renders error on 400 bad request" do
    stub_api(400, nil) do
      post "/tools/demos/inflation", params: { country: "ZZ" }
    end
    assert_response :success
    assert_match I18n.t("tools.inflation.demo.error_invalid"), response.body
  end

  test "inflation renders error on 404 not found" do
    stub_api(404, nil) do
      post "/tools/demos/inflation", params: { country: "ZZ" }
    end
    assert_response :success
    assert_match I18n.t("tools.inflation.demo.error_no_data"), response.body
  end

  test "inflation renders error on 429 rate limit" do
    stub_api(429, nil) do
      post "/tools/demos/inflation", params: { country: "US" }
    end
    assert_response :success
    assert_match I18n.t("tools.inflation.demo.error_rate_limit"), response.body
  end

  test "inflation renders error on non-200 response" do
    stub_api(500, nil) do
      post "/tools/demos/inflation", params: { country: "US" }
    end
    assert_response :success
    assert_match I18n.t("tools.inflation.demo.error_generic"), response.body
  end

  test "inflation renders error when data is nil" do
    stub_api(200, nil) do
      post "/tools/demos/inflation", params: { country: "US" }
    end
    assert_response :success
    assert_match I18n.t("tools.inflation.demo.error_no_data"), response.body
  end

  test "inflation normalizes lowercase country code" do
    payload = {
      "country" => "US", "rate" => 2.9495, "period" => "2024",
      "historical" => []
    }
    stub_api(200, success_data(payload)) do
      post "/tools/demos/inflation", params: { country: "us" }
    end
    assert_response :success
    assert_match "2.95", response.body
  end

  # ── bin_lookup ───────────────────────────────────────────────────────────────

  test "bin_lookup renders result on success" do
    data = {
      "bin" => "424242", "scheme" => "visa", "card_type" => "credit",
      "card_level" => "classic", "issuer_name" => "Chase",
      "issuer_url" => "www.chase.com", "issuer_phone" => "+18002324000",
      "country_code" => "US", "country_name" => "United States",
      "prepaid" => false, "luhn" => true, "confidence" => 0.92
    }
    stub_api(200, success_data(data)) do
      post "/tools/demos/bin-lookup", params: { bin: "424242" }
    end
    assert_response :success
    assert_match "424242", response.body
    assert_match "visa", response.body
  end

  test "bin_lookup renders error when bin is blank" do
    post "/tools/demos/bin-lookup", params: { bin: "" }
    assert_response :success
    assert_match I18n.t("tools.bin_lookup.demo.error_empty"), response.body
  end

  test "bin_lookup renders error when bin is whitespace" do
    post "/tools/demos/bin-lookup", params: { bin: "   " }
    assert_response :success
    assert_match I18n.t("tools.bin_lookup.demo.error_empty"), response.body
  end

  test "bin_lookup renders error when bin has invalid format" do
    post "/tools/demos/bin-lookup", params: { bin: "ABCDEF" }
    assert_response :success
    assert_match I18n.t("tools.bin_lookup.demo.error_invalid"), response.body
  end

  test "bin_lookup renders error on 404 not found" do
    stub_api(404, nil) do
      post "/tools/demos/bin-lookup", params: { bin: "000000" }
    end
    assert_response :success
    assert_match I18n.t("tools.bin_lookup.demo.error_no_data"), response.body
  end

  test "bin_lookup renders error on 429 rate limit" do
    stub_api(429, nil) do
      post "/tools/demos/bin-lookup", params: { bin: "424242" }
    end
    assert_response :success
    assert_match I18n.t("tools.bin_lookup.demo.error_rate_limit"), response.body
  end

  test "bin_lookup renders error on non-200 response" do
    stub_api(500, nil) do
      post "/tools/demos/bin-lookup", params: { bin: "424242" }
    end
    assert_response :success
    assert_match I18n.t("tools.bin_lookup.demo.error_generic"), response.body
  end

  test "bin_lookup renders error when data is nil" do
    stub_api(200, nil) do
      post "/tools/demos/bin-lookup", params: { bin: "424242" }
    end
    assert_response :success
    assert_match I18n.t("tools.bin_lookup.demo.error_no_data"), response.body
  end

  # ── qr_code ──────────────────────────────────────────────────────────────────

  test "qr_code renders result on success" do
    payload = { "image" => "iVBORw0KGgoAAAANSUhEUgAA", "width" => 256, "height" => 256 }
    stub_api(200, success_data(payload)) do
      post "/tools/demos/qr-code", params: { data: "https://example.com" }
    end
    assert_response :success
    assert_match "256", response.body
  end

  test "qr_code renders error when data is blank" do
    post "/tools/demos/qr-code", params: { data: "" }
    assert_response :success
    assert_match I18n.t("tools.qr_code.demo.error_empty"), response.body
  end

  test "qr_code renders error on 429" do
    stub_api(429, nil) do
      post "/tools/demos/qr-code", params: { data: "https://example.com" }
    end
    assert_response :success
    assert_match I18n.t("tools.qr_code.demo.error_rate_limit"), response.body
  end

  test "qr_code renders error on non-200 response" do
    stub_api(500, nil) do
      post "/tools/demos/qr-code", params: { data: "https://example.com" }
    end
    assert_response :success
    assert_match I18n.t("tools.qr_code.demo.error_generic"), response.body
  end

  test "qr_code renders error when data is nil" do
    stub_api(200, nil) do
      post "/tools/demos/qr-code", params: { data: "https://example.com" }
    end
    assert_response :success
    assert_match I18n.t("tools.qr_code.demo.error_no_data"), response.body
  end

  # ── profanity_filter ─────────────────────────────────────────────────────────

  test "profanity_filter renders result on success" do
    payload = { "has_profanity" => true, "censored" => "This is a **** day",
                "flagged_words" => [ "damn" ] }
    stub_api(200, success_data(payload)) do
      post "/tools/demos/profanity-filter", params: { text: "This is a damn day" }
    end
    assert_response :success
    assert_match "damn", response.body
    assert_match "This is a **** day", response.body
  end

  test "profanity_filter renders error when text is blank" do
    post "/tools/demos/profanity-filter", params: { text: "" }
    assert_response :success
    assert_match I18n.t("tools.profanity_filter.demo.error_empty"), response.body
  end

  test "profanity_filter renders error on 429" do
    stub_api(429, nil) do
      post "/tools/demos/profanity-filter", params: { text: "hello" }
    end
    assert_response :success
    assert_match I18n.t("tools.profanity_filter.demo.error_rate_limit"), response.body
  end

  test "profanity_filter renders error on non-200 response" do
    stub_api(500, nil) do
      post "/tools/demos/profanity-filter", params: { text: "hello" }
    end
    assert_response :success
    assert_match I18n.t("tools.profanity_filter.demo.error_generic"), response.body
  end

  test "profanity_filter renders error when data is nil" do
    stub_api(200, nil) do
      post "/tools/demos/profanity-filter", params: { text: "hello" }
    end
    assert_response :success
    assert_match I18n.t("tools.profanity_filter.demo.error_no_data"), response.body
  end

  # ── useragent ─────────────────────────────────────────────────────────────────

  test "useragent renders result on success" do
    payload = { "browser" => "Chrome", "browser_version" => "120.0",
                "os" => "Windows", "os_version" => "10/11",
                "device" => "desktop", "is_bot" => false }
    stub_api(200, success_data(payload)) do
      post "/tools/demos/useragent", params: { ua: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) Chrome/120.0" }
    end
    assert_response :success
    assert_match "Chrome", response.body
    assert_match "Windows", response.body
  end

  test "useragent renders error when ua is blank" do
    post "/tools/demos/useragent", params: { ua: "" }
    assert_response :success
    assert_match I18n.t("tools.useragent.demo.error_empty"), response.body
  end

  test "useragent renders error on 429" do
    stub_api(429, nil) do
      post "/tools/demos/useragent", params: { ua: "Mozilla/5.0" }
    end
    assert_response :success
    assert_match I18n.t("tools.useragent.demo.error_rate_limit"), response.body
  end

  test "useragent renders error on non-200 response" do
    stub_api(500, nil) do
      post "/tools/demos/useragent", params: { ua: "Mozilla/5.0" }
    end
    assert_response :success
    assert_match I18n.t("tools.useragent.demo.error_generic"), response.body
  end

  test "useragent renders error when data is nil" do
    stub_api(200, nil) do
      post "/tools/demos/useragent", params: { ua: "Mozilla/5.0" }
    end
    assert_response :success
    assert_match I18n.t("tools.useragent.demo.error_no_data"), response.body
  end

  # ── timezone ─────────────────────────────────────────────────────────────────

  test "timezone renders result on success by city" do
    payload = { "timezone" => "Asia/Tokyo", "offset" => "+09:00",
                "current_time" => "2026-07-11T23:00:00Z", "is_dst" => false }
    stub_api(200, success_data(payload)) do
      post "/tools/demos/timezone", params: { city: "Tokyo" }
    end
    assert_response :success
    assert_match "Asia/Tokyo", response.body
  end

  test "timezone renders result on success by coordinates" do
    payload = { "timezone" => "Europe/London", "offset" => "+00:00",
                "current_time" => "2026-07-11T23:00:00Z", "is_dst" => false }
    stub_api(200, success_data(payload)) do
      post "/tools/demos/timezone", params: { lat: "51.5", lon: "-0.1" }
    end
    assert_response :success
    assert_match "Europe/London", response.body
  end

  test "timezone renders error when city and coordinates are all blank" do
    post "/tools/demos/timezone", params: { city: "" }
    assert_response :success
    assert_match I18n.t("tools.timezone.demo.error_empty"), response.body
  end

  test "timezone renders error when only one coordinate is given" do
    post "/tools/demos/timezone", params: { lat: "51.5" }
    assert_response :success
    assert_match I18n.t("tools.timezone.demo.error_empty"), response.body
  end

  test "timezone renders error when coordinates are out of bounds" do
    post "/tools/demos/timezone", params: { lat: "999", lon: "0" }
    assert_response :success
    assert_match I18n.t("tools.timezone.demo.error_invalid"), response.body
  end

  test "timezone renders error when coordinates are non-numeric" do
    post "/tools/demos/timezone", params: { lat: "abc", lon: "def" }
    assert_response :success
    assert_match I18n.t("tools.timezone.demo.error_invalid"), response.body
  end

  test "timezone renders error on 404 not found" do
    stub_api(404, nil) do
      post "/tools/demos/timezone", params: { city: "Nowhereville" }
    end
    assert_response :success
    assert_match I18n.t("tools.timezone.demo.error_no_data"), response.body
  end

  test "timezone renders error on 429 rate limit" do
    stub_api(429, nil) do
      post "/tools/demos/timezone", params: { city: "Tokyo" }
    end
    assert_response :success
    assert_match I18n.t("tools.timezone.demo.error_rate_limit"), response.body
  end

  test "timezone renders error on non-200 response" do
    stub_api(500, nil) do
      post "/tools/demos/timezone", params: { city: "Tokyo" }
    end
    assert_response :success
    assert_match I18n.t("tools.timezone.demo.error_generic"), response.body
  end

  test "timezone renders error when data is nil" do
    stub_api(200, nil) do
      post "/tools/demos/timezone", params: { city: "Tokyo" }
    end
    assert_response :success
    assert_match I18n.t("tools.timezone.demo.error_no_data"), response.body
  end

  # ── trivia ───────────────────────────────────────────────────────────────────

  test "trivia renders result on success with filters" do
    payload = {
      "question" => "What is the largest planet in our solar system?",
      "options" => [ "Earth", "Jupiter", "Saturn", "Mars" ],
      "answer" => "Jupiter", "category" => "science", "difficulty" => "easy"
    }
    stub_api(200, success_data(payload)) do
      post "/tools/demos/trivia", params: { category: "science", difficulty: "easy" }
    end
    assert_response :success
    assert_match "Jupiter", response.body
    assert_match "largest planet", response.body
  end

  test "trivia renders result on success with no filters" do
    payload = {
      "question" => "What is H2O?", "options" => [ "Water", "Salt", "Sugar", "Oil" ],
      "answer" => "Water", "category" => "science", "difficulty" => "medium"
    }
    stub_api(200, success_data(payload)) do
      post "/tools/demos/trivia", params: {}
    end
    assert_response :success
    assert_match "Water", response.body
  end

  test "trivia renders error on 404 no matching questions" do
    stub_api(404, nil) do
      post "/tools/demos/trivia", params: { category: "science", difficulty: "hard" }
    end
    assert_response :success
    assert_match I18n.t("tools.trivia.demo.error_no_data"), response.body
  end

  test "trivia renders error on 429 rate limit" do
    stub_api(429, nil) do
      post "/tools/demos/trivia", params: {}
    end
    assert_response :success
    assert_match I18n.t("tools.trivia.demo.error_rate_limit"), response.body
  end

  test "trivia renders error on non-200 response" do
    stub_api(500, nil) do
      post "/tools/demos/trivia", params: {}
    end
    assert_response :success
    assert_match I18n.t("tools.trivia.demo.error_generic"), response.body
  end

  test "trivia renders error when data is nil" do
    stub_api(200, nil) do
      post "/tools/demos/trivia", params: {}
    end
    assert_response :success
    assert_match I18n.t("tools.trivia.demo.error_no_data"), response.body
  end

  # ── vpn_detection ────────────────────────────────────────────────────────────

  test "vpn_detection renders result on success" do
    payload = { "ip" => "8.8.8.8", "is_vpn" => false, "is_proxy" => false,
                "is_tor" => false, "is_hosting" => true, "score" => 1,
                "threat" => 1, "fraud_score" => 0, "asn_org" => "GOOGLE-ASN" }
    stub_api(200, success_data(payload)) do
      post "/tools/demos/vpn-detection", params: { ip: "8.8.8.8" }
    end
    assert_response :success
    assert_match "8.8.8.8", response.body
    assert_match "GOOGLE-ASN", response.body
  end

  test "vpn_detection renders error when ip is blank" do
    post "/tools/demos/vpn-detection", params: { ip: "" }
    assert_response :success
    assert_match I18n.t("tools.vpn_detection.demo.error_empty"), response.body
  end

  test "vpn_detection renders error when ip is whitespace" do
    post "/tools/demos/vpn-detection", params: { ip: "   " }
    assert_response :success
    assert_match I18n.t("tools.vpn_detection.demo.error_empty"), response.body
  end

  test "vpn_detection renders error when ip has invalid format" do
    post "/tools/demos/vpn-detection", params: { ip: "not-an-ip" }
    assert_response :success
    assert_match I18n.t("tools.vpn_detection.demo.error_invalid"), response.body
  end

  test "vpn_detection renders error when ip is a CIDR range" do
    post "/tools/demos/vpn-detection", params: { ip: "192.168.1.0/24" }
    assert_response :success
    assert_match I18n.t("tools.vpn_detection.demo.error_invalid"), response.body
  end

  test "vpn_detection accepts a valid IPv6 address" do
    payload = { "ip" => "2001:db8::1", "is_vpn" => false, "is_proxy" => false,
                "is_tor" => false, "is_hosting" => false, "score" => 0,
                "threat" => 0, "fraud_score" => 0, "asn_org" => "" }
    stub_api(200, success_data(payload)) do
      post "/tools/demos/vpn-detection", params: { ip: "2001:db8::1" }
    end
    assert_response :success
    assert_match "2001:db8::1", response.body
  end

  test "vpn_detection renders error on 429 rate limit" do
    stub_api(429, nil) do
      post "/tools/demos/vpn-detection", params: { ip: "8.8.8.8" }
    end
    assert_response :success
    assert_match I18n.t("tools.vpn_detection.demo.error_rate_limit"), response.body
  end

  test "vpn_detection renders error on non-200 response" do
    stub_api(500, nil) do
      post "/tools/demos/vpn-detection", params: { ip: "8.8.8.8" }
    end
    assert_response :success
    assert_match I18n.t("tools.vpn_detection.demo.error_generic"), response.body
  end

  test "vpn_detection renders error when data is nil" do
    stub_api(200, nil) do
      post "/tools/demos/vpn-detection", params: { ip: "8.8.8.8" }
    end
    assert_response :success
    assert_match I18n.t("tools.vpn_detection.demo.error_no_data"), response.body
  end

  # ── thesaurus ────────────────────────────────────────────────────────────────

  test "thesaurus renders result on success" do
    payload = { "word" => "happy", "synonyms" => [ "joyful", "cheerful" ], "antonyms" => [ "sad", "unhappy" ] }
    stub_api(200, success_data(payload)) do
      post "/tools/demos/thesaurus", params: { word: "happy" }
    end
    assert_response :success
    assert_match "joyful", response.body
    assert_match "sad", response.body
  end

  test "thesaurus renders error when word is blank" do
    post "/tools/demos/thesaurus", params: { word: "" }
    assert_response :success
    assert_match I18n.t("tools.thesaurus.demo.error_empty"), response.body
  end

  test "thesaurus renders error when word is whitespace" do
    post "/tools/demos/thesaurus", params: { word: "   " }
    assert_response :success
    assert_match I18n.t("tools.thesaurus.demo.error_empty"), response.body
  end

  test "thesaurus renders error when word has invalid format" do
    post "/tools/demos/thesaurus", params: { word: "happy123" }
    assert_response :success
    assert_match I18n.t("tools.thesaurus.demo.error_invalid"), response.body
  end

  test "thesaurus renders error on 404 not found" do
    stub_api(404, nil) do
      post "/tools/demos/thesaurus", params: { word: "notaword" }
    end
    assert_response :success
    assert_match I18n.t("tools.thesaurus.demo.error_no_data"), response.body
  end

  test "thesaurus renders error on 429 rate limit" do
    stub_api(429, nil) do
      post "/tools/demos/thesaurus", params: { word: "happy" }
    end
    assert_response :success
    assert_match I18n.t("tools.thesaurus.demo.error_rate_limit"), response.body
  end

  test "thesaurus renders error on non-200 response" do
    stub_api(500, nil) do
      post "/tools/demos/thesaurus", params: { word: "happy" }
    end
    assert_response :success
    assert_match I18n.t("tools.thesaurus.demo.error_generic"), response.body
  end

  test "thesaurus renders error when data is nil" do
    stub_api(200, nil) do
      post "/tools/demos/thesaurus", params: { word: "happy" }
    end
    assert_response :success
    assert_match I18n.t("tools.thesaurus.demo.error_no_data"), response.body
  end

  # ── spell_check ───────────────────────────────────────────────────────────────

  test "spell_check renders result on success with corrections" do
    payload = {
      "corrected" => "This is a simple test with errors",
      "corrections" => [
        { "original" => "simiple", "suggested" => "simple", "suggestions" => [ "sample" ] },
        { "original" => "erors", "suggested" => "errors", "suggestions" => [ "errs" ] }
      ]
    }
    stub_api(200, success_data(payload)) do
      post "/tools/demos/spell-check", params: { text: "This is a simiple test with erors" }
    end
    assert_response :success
    assert_match "simple", response.body
    assert_match "errors", response.body
    assert_match "simiple", response.body
    assert_match "erors", response.body
  end

  test "spell_check renders result on success with no errors" do
    payload = { "corrected" => "Clean text here", "corrections" => [] }
    stub_api(200, success_data(payload)) do
      post "/tools/demos/spell-check", params: { text: "Clean text here" }
    end
    assert_response :success
    assert_match I18n.t("tools.spell_check.demo.badge_clean"), response.body
  end

  test "spell_check renders error when text is blank" do
    post "/tools/demos/spell-check", params: { text: "" }
    assert_response :success
    assert_match I18n.t("tools.spell_check.demo.error_empty"), response.body
  end

  test "spell_check renders error on 429" do
    stub_api(429, nil) do
      post "/tools/demos/spell-check", params: { text: "hello" }
    end
    assert_response :success
    assert_match I18n.t("tools.spell_check.demo.error_rate_limit"), response.body
  end

  test "spell_check renders error on non-200 response" do
    stub_api(500, nil) do
      post "/tools/demos/spell-check", params: { text: "hello" }
    end
    assert_response :success
    assert_match I18n.t("tools.spell_check.demo.error_generic"), response.body
  end

  test "spell_check renders error when data is nil" do
    stub_api(200, nil) do
      post "/tools/demos/spell-check", params: { text: "hello" }
    end
    assert_response :success
    assert_match I18n.t("tools.spell_check.demo.error_no_data"), response.body
  end

  # ── random_user ────────────────────────────────────────────────────────────────

  test "random_user renders result on success" do
    payload = {
      "name" => "Alice Johnson", "email" => "alice@example.org",
      "phone" => "+1234567890",
      "address" => {
        "street" => "123 Main St",
        "city" => "Springfield",
        "state" => "IL",
        "zip" => "62701",
        "country" => "US"
      },
      "avatar" => "https://api.dicebear.com/7.x/identicon/svg/seed=Alice"
    }
    stub_api(200, success_data(payload)) do
      post "/tools/demos/random-user"
    end
    assert_response :success
    assert_match "Alice Johnson", response.body
    assert_match "alice@example.org", response.body
    assert_match "123 Main St", response.body
  end

  test "random_user renders error on 429" do
    stub_api(429, nil) do
      post "/tools/demos/random-user"
    end
    assert_response :success
    assert_match I18n.t("tools.random_user.demo.error_rate_limit"), response.body
  end

  test "random_user renders error on non-200 response" do
    stub_api(500, nil) do
      post "/tools/demos/random-user"
    end
    assert_response :success
    assert_match I18n.t("tools.random_user.demo.error_generic"), response.body
  end

  test "random_user renders error when data is nil" do
    stub_api(200, nil) do
      post "/tools/demos/random-user"
    end
    assert_response :success
    assert_match I18n.t("tools.random_user.demo.error_no_data"), response.body
  end

  # ── sudoku ─────────────────────────────────────────────────────────────────────

  test "sudoku renders result on success with default difficulty" do
    payload = {
      "difficulty" => "medium",
      "puzzle" => [
        [5, 3, 0, 0, 7, 0, 0, 0, 0],
        [6, 0, 0, 1, 9, 5, 0, 0, 0],
        [0, 9, 8, 0, 0, 0, 0, 6, 0],
        [8, 0, 0, 0, 6, 0, 0, 0, 3],
        [4, 0, 0, 8, 0, 3, 0, 0, 1],
        [7, 0, 0, 0, 2, 0, 0, 0, 6],
        [0, 6, 0, 0, 0, 0, 2, 8, 0],
        [0, 0, 0, 4, 1, 9, 0, 0, 5],
        [0, 0, 0, 0, 8, 0, 0, 7, 9]
      ],
      "solution" => [
        [5, 3, 4, 6, 7, 8, 9, 1, 2],
        [6, 7, 2, 1, 9, 5, 3, 4, 8],
        [1, 9, 8, 3, 4, 2, 5, 6, 7],
        [8, 5, 9, 7, 6, 1, 4, 2, 3],
        [4, 2, 6, 8, 5, 3, 7, 9, 1],
        [7, 1, 3, 9, 2, 4, 8, 5, 6],
        [9, 6, 1, 5, 3, 7, 2, 8, 4],
        [2, 8, 7, 4, 1, 9, 6, 3, 5],
        [3, 4, 5, 2, 8, 6, 1, 7, 9]
      ]
    }
    stub_api(200, success_data(payload)) do
      post "/tools/demos/sudoku", params: {}
    end
    assert_response :success
    assert_match "medium", response.body
    assert_match "5", response.body
    assert_match I18n.t("tools.sudoku.demo.label_solution"), response.body
  end

  test "sudoku renders result on success with easy difficulty" do
    payload = {
      "difficulty" => "easy",
      "puzzle" => [
        [0, 0, 4, 0, 5, 0, 0, 0, 0],
        [9, 0, 0, 7, 3, 4, 6, 0, 0],
        [0, 0, 3, 0, 2, 1, 0, 4, 9],
        [0, 3, 5, 0, 9, 0, 4, 8, 0],
        [6, 0, 0, 0, 0, 0, 0, 0, 7],
        [0, 7, 9, 0, 8, 0, 1, 3, 0],
        [4, 5, 0, 9, 1, 0, 7, 0, 0],
        [0, 0, 2, 4, 7, 8, 0, 0, 3],
        [0, 0, 0, 0, 4, 0, 9, 0, 0]
      ],
      "solution" => [
        [2, 6, 4, 8, 5, 9, 3, 7, 1],
        [9, 1, 8, 7, 3, 4, 6, 5, 2],
        [7, 0, 3, 6, 2, 1, 8, 4, 9],
        [1, 3, 5, 2, 9, 7, 4, 8, 6],
        [6, 8, 0, 3, 0, 5, 2, 0, 7],
        [0, 7, 9, 4, 8, 6, 1, 3, 5],
        [4, 5, 6, 9, 1, 3, 7, 2, 8],
        [0, 9, 2, 4, 7, 8, 5, 6, 3],
        [3, 0, 7, 5, 4, 2, 9, 1, 0]
      ]
    }
    stub_api(200, success_data(payload)) do
      post "/tools/demos/sudoku", params: { difficulty: "easy" }
    end
    assert_response :success
    assert_match I18n.t("tools.sudoku.demo.badge_easy"), response.body
  end

  test "sudoku renders result on success with hard difficulty" do
    payload = {
      "difficulty" => "hard",
      "puzzle" => [
        [0, 0, 0, 0, 0, 0, 0, 0, 0],
        [0, 0, 0, 0, 0, 3, 0, 8, 5],
        [0, 0, 1, 0, 2, 0, 0, 0, 0],
        [0, 0, 0, 5, 0, 7, 0, 0, 0],
        [0, 0, 4, 0, 0, 0, 1, 0, 0],
        [0, 9, 0, 0, 0, 0, 0, 0, 0],
        [5, 0, 0, 0, 0, 0, 0, 7, 3],
        [0, 0, 2, 0, 1, 0, 0, 0, 0],
        [0, 0, 0, 0, 4, 0, 0, 0, 9]
      ],
      "solution" => [
        [9, 8, 7, 6, 5, 4, 3, 2, 1],
        [2, 4, 6, 1, 7, 3, 9, 8, 5],
        [3, 5, 1, 9, 2, 8, 7, 4, 6],
        [1, 2, 3, 5, 9, 7, 4, 6, 8],
        [7, 6, 4, 2, 8, 0, 1, 9, 0],
        [8, 9, 5, 4, 6, 1, 2, 3, 7],
        [5, 1, 9, 8, 0, 6, 0, 7, 3],
        [4, 7, 2, 3, 1, 9, 8, 5, 0],
        [6, 3, 8, 7, 4, 5, 0, 1, 9]
      ]
    }
    stub_api(200, success_data(payload)) do
      post "/tools/demos/sudoku", params: { difficulty: "hard" }
    end
    assert_response :success
    assert_match I18n.t("tools.sudoku.demo.badge_hard"), response.body
  end

  test "sudoku renders error when difficulty is invalid" do
    post "/tools/demos/sudoku", params: { difficulty: "extreme" }
    assert_response :success
    assert_match I18n.t("tools.sudoku.demo.error_invalid_difficulty"), response.body
  end

  test "sudoku renders error on 429" do
    stub_api(429, nil) do
      post "/tools/demos/sudoku", params: { difficulty: "medium" }
    end
    assert_response :success
    assert_match I18n.t("tools.sudoku.demo.error_rate_limit"), response.body
  end

  test "sudoku renders error on non-200 response" do
    stub_api(500, nil) do
      post "/tools/demos/sudoku", params: { difficulty: "medium" }
    end
    assert_response :success
    assert_match I18n.t("tools.sudoku.demo.error_generic"), response.body
  end

  test "sudoku renders error when data is nil" do
    stub_api(200, nil) do
      post "/tools/demos/sudoku", params: { difficulty: "medium" }
    end
    assert_response :success
    assert_match I18n.t("tools.sudoku.demo.error_no_data"), response.body
  end

  # ── number_base_conversion ───────────────────────────────────────────────────

  test "number_base_conversion renders result on success" do
    payload = { "input" => "255", "from" => 10, "to" => 16, "result" => "ff" }
    stub_api(200, success_data(payload)) do
      post "/tools/demos/number-base-conversion", params: { from: "10", to: "16", value: "255" }
    end
    assert_response :success
    assert_match "255", response.body
    assert_match "ff", response.body
  end

  test "number_base_conversion renders error when any field is blank" do
    post "/tools/demos/number-base-conversion", params: { from: "10", to: "", value: "255" }
    assert_response :success
    assert_match I18n.t("tools.number_base_conversion.demo.error_fill_all_fields"), response.body
  end

  test "number_base_conversion renders error on 429" do
    stub_api(429, nil) do
      post "/tools/demos/number-base-conversion", params: { from: "10", to: "16", value: "255" }
    end
    assert_response :success
    assert_match I18n.t("tools.number_base_conversion.demo.error_rate_limit"), response.body
  end

  test "number_base_conversion renders error on non-200 response" do
    stub_api(500, nil) do
      post "/tools/demos/number-base-conversion", params: { from: "10", to: "16", value: "255" }
    end
    assert_response :success
    assert_match I18n.t("tools.number_base_conversion.demo.error_generic"), response.body
  end

  test "number_base_conversion renders error when data is nil" do
    stub_api(200, nil) do
      post "/tools/demos/number-base-conversion", params: { from: "10", to: "16", value: "255" }
    end
    assert_response :success
    assert_match I18n.t("tools.number_base_conversion.demo.error_no_data"), response.body
  end

  # ── mx_lookup ────────────────────────────────────────────────────────────────

  test "mx_lookup renders result on success" do
    payload = { "domain" => "gmail.com", "records" => [
      { "host" => "gmail-smtp-in.l.google.com.", "priority" => 5 },
      { "host" => "alt1.gmail-smtp-in.l.google.com.", "priority" => 10 }
    ] }
    stub_api(200, success_data(payload)) do
      post "/tools/demos/mx-lookup", params: { domain: "gmail.com" }
    end
    assert_response :success
    assert_match "gmail.com", response.body
    assert_match "gmail-smtp-in.l.google.com.", response.body
  end

  test "mx_lookup renders error when domain is blank" do
    post "/tools/demos/mx-lookup", params: { domain: "" }
    assert_response :success
    assert_match I18n.t("tools.mx_lookup.demo.error_empty"), response.body
  end

  test "mx_lookup renders error when domain is whitespace" do
    post "/tools/demos/mx-lookup", params: { domain: "   " }
    assert_response :success
    assert_match I18n.t("tools.mx_lookup.demo.error_empty"), response.body
  end

  test "mx_lookup renders error when domain has invalid format" do
    post "/tools/demos/mx-lookup", params: { domain: "not a domain" }
    assert_response :success
    assert_match I18n.t("tools.mx_lookup.demo.error_invalid"), response.body
  end

  test "mx_lookup renders error on 404 no mx records" do
    stub_api(404, nil) do
      post "/tools/demos/mx-lookup", params: { domain: "no-mx.example.com" }
    end
    assert_response :success
    assert_match I18n.t("tools.mx_lookup.demo.error_not_found"), response.body
  end

  test "mx_lookup renders error on 429 rate limit" do
    stub_api(429, nil) do
      post "/tools/demos/mx-lookup", params: { domain: "gmail.com" }
    end
    assert_response :success
    assert_match I18n.t("tools.mx_lookup.demo.error_rate_limit"), response.body
  end

  test "mx_lookup renders error on non-200 response" do
    stub_api(500, nil) do
      post "/tools/demos/mx-lookup", params: { domain: "gmail.com" }
    end
    assert_response :success
    assert_match I18n.t("tools.mx_lookup.demo.error_generic"), response.body
  end

  test "mx_lookup renders error when data is nil" do
    stub_api(200, nil) do
      post "/tools/demos/mx-lookup", params: { domain: "gmail.com" }
    end
    assert_response :success
    assert_match I18n.t("tools.mx_lookup.demo.error_no_data"), response.body
  end

  # ── mortgage ─────────────────────────────────────────────────────────────────

  test "mortgage renders result on success" do
    payload = {
      "principal" => 300_000, "rate" => 6.5, "years" => 30,
      "monthly_payment" => 1896.20, "total_payment" => 682_632.00,
      "total_interest" => 382_632.00,
      "schedule" => [
        { "month" => 1, "payment" => 1896.20,
          "principal" => 271.20, "interest" => 1625.00, "balance" => 299_728.80 }
      ]
    }
    stub_api(200, success_data(payload)) do
      post "/tools/demos/mortgage",
           params: { principal: "300000", rate: "6.5", years: "30" }
    end
    assert_response :success
    assert_match "1,896.20", response.body
    assert_match "682,632.00", response.body
  end

  test "mortgage renders error when principal is blank" do
    post "/tools/demos/mortgage", params: { principal: "", rate: "6.5", years: "30" }
    assert_response :success
    assert_match I18n.t("tools.mortgage.demo.error_empty"), response.body
  end

  test "mortgage renders error when rate is blank" do
    post "/tools/demos/mortgage", params: { principal: "300000", rate: "", years: "30" }
    assert_response :success
    assert_match I18n.t("tools.mortgage.demo.error_empty"), response.body
  end

  test "mortgage renders error when years is blank" do
    post "/tools/demos/mortgage", params: { principal: "300000", rate: "6.5", years: "" }
    assert_response :success
    assert_match I18n.t("tools.mortgage.demo.error_empty"), response.body
  end

  test "mortgage renders error when params are invalid" do
    post "/tools/demos/mortgage", params: { principal: "-1", rate: "0", years: "100" }
    assert_response :success
    # Match a substring without HTML-sensitive characters (> becomes &gt; in the response body)
    assert_match "years between 1 and 50", response.body
  end

  test "mortgage renders error on 429 rate limit" do
    stub_api(429, nil) do
      post "/tools/demos/mortgage",
           params: { principal: "300000", rate: "6.5", years: "30" }
    end
    assert_response :success
    assert_match I18n.t("tools.mortgage.demo.error_rate_limit"), response.body
  end

  test "mortgage renders error on non-200 response" do
    stub_api(500, nil) do
      post "/tools/demos/mortgage",
           params: { principal: "300000", rate: "6.5", years: "30" }
    end
    assert_response :success
    assert_match I18n.t("tools.mortgage.demo.error_generic"), response.body
  end

  test "mortgage renders error when data is nil" do
    stub_api(200, nil) do
      post "/tools/demos/mortgage",
           params: { principal: "300000", rate: "6.5", years: "30" }
    end
    assert_response :success
    assert_match I18n.t("tools.mortgage.demo.error_no_data"), response.body
  end

  # ── markdown ─────────────────────────────────────────────────────────────────
  test "markdown renders result on success" do
    payload = { "html" => "<h1>Hello</h1>\n<p>This is <strong>bold</strong> text.</p>" }
    stub_api(200, success_data(payload)) do
      post "/tools/demos/markdown", params: { markdown: "# Hello\n\nThis is **bold** text." }
    end
    assert_response :success
    assert_match(/<h1>Hello<\/h1>/i, response.body)
  end
  test "markdown renders error when markdown is blank" do
    post "/tools/demos/markdown", params: { markdown: "" }
    assert_response :success
    assert_match I18n.t("tools.markdown.demo.error_empty"), response.body
  end
  test "markdown renders error on 429" do
    stub_api(429, nil) do
      post "/tools/demos/markdown", params: { markdown: "# Hello" }
    end
    assert_response :success
    assert_match I18n.t("tools.markdown.demo.error_rate_limit"), response.body
  end
  test "markdown renders error on non-200 response" do
    stub_api(500, nil) do
      post "/tools/demos/markdown", params: { markdown: "# Hello" }
    end
    assert_response :success
    assert_match I18n.t("tools.markdown.demo.error_generic"), response.body
  end
  test "markdown renders error when data is nil" do
    stub_api(200, nil) do
      post "/tools/demos/markdown", params: { markdown: "# Hello" }
    end
    assert_response :success
    assert_match I18n.t("tools.markdown.demo.error_no_data"), response.body
  end

  # ── barcode ──────────────────────────────────────────────────────────────────

  test "barcode renders result on success" do
    payload = { "image" => "iVBORw0KGgoAAAANSUhEUgAA", "type" => "code128", "width" => 300, "height" => 100 }
    stub_api(200, success_data(payload)) do
      post "/tools/demos/barcode", params: { data: "123456789", type: "code128" }
    end
    assert_response :success
    assert_match "iVBORw0KGgoAAAANSUhEUgAA", response.body
  end

  test "barcode renders error when data is blank" do
    post "/tools/demos/barcode", params: { data: "", type: "code128" }
    assert_response :success
    assert_match I18n.t("tools.barcode.demo.error_empty"), response.body
  end

  test "barcode renders error when type is invalid" do
    post "/tools/demos/barcode", params: { data: "123456789", type: "not-a-format" }
    assert_response :success
    assert_match I18n.t("tools.barcode.demo.error_invalid_type"), response.body
  end

  test "barcode renders error on 429 rate limit" do
    stub_api(429, nil) do
      post "/tools/demos/barcode", params: { data: "123456789", type: "code128" }
    end
    assert_response :success
    assert_match I18n.t("tools.barcode.demo.error_rate_limit"), response.body
  end

  test "barcode renders error on non-200 response" do
    stub_api(500, nil) do
      post "/tools/demos/barcode", params: { data: "123456789", type: "code128" }
    end
    assert_response :success
    assert_match I18n.t("tools.barcode.demo.error_generic"), response.body
  end

  test "barcode renders error when data is nil" do
    stub_api(200, nil) do
      post "/tools/demos/barcode", params: { data: "123456789", type: "code128" }
    end
    assert_response :success
    assert_match I18n.t("tools.barcode.demo.error_no_data"), response.body
  end

  # ── advice ───────────────────────────────────────────────────────────────────

  test "advice renders result on success" do
    payload = { "id" => 42, "advice" => "Don't compare yourself to others." }
    stub_api(200, success_data(payload)) do
      post "/tools/demos/advice"
    end
    assert_response :success
    assert_match "compare yourself to others", response.body
  end

  test "advice renders error on 429 rate limit" do
    stub_api(429, nil) do
      post "/tools/demos/advice"
    end
    assert_response :success
    assert_match I18n.t("tools.advice.demo.error_rate_limit"), response.body
  end

  test "advice renders error on non-200 response" do
    stub_api(500, nil) do
      post "/tools/demos/advice"
    end
    assert_response :success
    assert_match I18n.t("tools.advice.demo.error_generic"), response.body
  end

  test "advice renders error when data is nil" do
    stub_api(200, nil) do
      post "/tools/demos/advice"
    end
    assert_response :success
    assert_match I18n.t("tools.advice.demo.error_no_data"), response.body
  end

  # ── base64 ───────────────────────────────────────────────────────────────────

  test "base64 renders result on success for encode" do
    payload = { "result" => "SGVsbG8sIHdvcmxkIQ==" }
    stub_api(200, success_data(payload)) do
      post "/tools/demos/base64", params: { mode: "encode", value: "Hello, world!" }
    end
    assert_response :success
    assert_match "SGVsbG8sIHdvcmxkIQ==", response.body
  end

  test "base64 renders result on success for decode" do
    payload = { "result" => "Hello, world!" }
    stub_api(200, success_data(payload)) do
      post "/tools/demos/base64", params: { mode: "decode", value: "SGVsbG8sIHdvcmxkIQ==" }
    end
    assert_response :success
    assert_match "Hello, world!", response.body
  end

  test "base64 renders error when value is blank" do
    post "/tools/demos/base64", params: { mode: "encode", value: "" }
    assert_response :success
    assert_match I18n.t("tools.base64.demo.error_empty"), response.body
  end

  test "base64 renders error when mode is invalid" do
    post "/tools/demos/base64", params: { mode: "reverse", value: "Hello" }
    assert_response :success
    assert_match I18n.t("tools.base64.demo.error_invalid_mode"), response.body
  end

  test "base64 renders error on 422 invalid base64" do
    stub_api(422, nil) do
      post "/tools/demos/base64", params: { mode: "decode", value: "not-valid-base64!!" }
    end
    assert_response :success
    # Match a substring without HTML-sensitive characters (' becomes &#39; in the response body)
    assert_match "Check the padding and characters", response.body
  end

  test "base64 renders error on 429 rate limit" do
    stub_api(429, nil) do
      post "/tools/demos/base64", params: { mode: "encode", value: "Hello" }
    end
    assert_response :success
    assert_match I18n.t("tools.base64.demo.error_rate_limit"), response.body
  end

  test "base64 renders error on non-200 response" do
    stub_api(500, nil) do
      post "/tools/demos/base64", params: { mode: "encode", value: "Hello" }
    end
    assert_response :success
    assert_match I18n.t("tools.base64.demo.error_generic"), response.body
  end

  test "base64 renders error when data is nil" do
    stub_api(200, nil) do
      post "/tools/demos/base64", params: { mode: "encode", value: "Hello" }
    end
    assert_response :success
    assert_match I18n.t("tools.base64.demo.error_no_data"), response.body
  end

  # ── whois ────────────────────────────────────────────────────────────────────

  test "whois renders result on success" do
    payload = { "domain" => "example.com", "registrar" => "RESERVED-Internet Assigned Numbers Authority",
                "name_servers" => [ "A.IANA-SERVERS.NET" ], "status" => [ "clientTransferProhibited" ],
                "created_date" => "1995-08-14T04:00:00Z", "updated_date" => "2023-08-14T07:01:38Z",
                "expiry_date" => "2024-08-13T04:00:00Z", "dnssec" => true }
    stub_api(200, success_data(payload)) do
      post "/tools/demos/whois", params: { domain: "example.com" }
    end
    assert_response :success
    assert_match "RESERVED-Internet Assigned Numbers Authority", response.body
  end

  test "whois renders error when domain is blank" do
    post "/tools/demos/whois", params: { domain: "" }
    assert_response :success
    assert_match I18n.t("tools.whois.demo.error_empty"), response.body
  end

  test "whois renders error when domain has invalid format" do
    post "/tools/demos/whois", params: { domain: "not a domain" }
    assert_response :success
    assert_match I18n.t("tools.whois.demo.error_invalid"), response.body
  end

  test "whois renders error on 404 not found" do
    stub_api(404, nil) do
      post "/tools/demos/whois", params: { domain: "doesnotexist.example" }
    end
    assert_response :success
    assert_match I18n.t("tools.whois.demo.error_not_found"), response.body
  end

  test "whois renders error on 429 rate limit" do
    stub_api(429, nil) do
      post "/tools/demos/whois", params: { domain: "example.com" }
    end
    assert_response :success
    assert_match I18n.t("tools.whois.demo.error_rate_limit"), response.body
  end

  test "whois renders error on non-200 response" do
    stub_api(500, nil) do
      post "/tools/demos/whois", params: { domain: "example.com" }
    end
    assert_response :success
    assert_match I18n.t("tools.whois.demo.error_generic"), response.body
  end

  test "whois renders error when data is nil" do
    stub_api(200, nil) do
      post "/tools/demos/whois", params: { domain: "example.com" }
    end
    assert_response :success
    assert_match I18n.t("tools.whois.demo.error_no_data"), response.body
  end

  # ── lorem_ipsum ──────────────────────────────────────────────────────────────

  test "lorem_ipsum renders result on success" do
    payload = { "text" => "Lorem ipsum dolor sit amet.", "paragraphs" => 1, "word_count" => 5 }
    stub_api(200, success_data(payload)) do
      post "/tools/demos/lorem-ipsum", params: { paragraphs: "1", sentences: "5" }
    end
    assert_response :success
    assert_match "Lorem ipsum dolor sit amet.", response.body
  end

  test "lorem_ipsum uses defaults when params are blank" do
    payload = { "text" => "Lorem ipsum dolor sit amet.", "paragraphs" => 1, "word_count" => 5 }
    stub_api(200, success_data(payload)) do
      post "/tools/demos/lorem-ipsum", params: {}
    end
    assert_response :success
    assert_match "Lorem ipsum dolor sit amet.", response.body
  end

  test "lorem_ipsum renders error when paragraphs is out of range" do
    post "/tools/demos/lorem-ipsum", params: { paragraphs: "21", sentences: "5" }
    assert_response :success
    assert_match I18n.t("tools.lorem_ipsum.demo.error_invalid"), response.body
  end

  test "lorem_ipsum renders error when sentences is out of range" do
    post "/tools/demos/lorem-ipsum", params: { paragraphs: "1", sentences: "0" }
    assert_response :success
    assert_match I18n.t("tools.lorem_ipsum.demo.error_invalid"), response.body
  end

  test "lorem_ipsum renders error on 429 rate limit" do
    stub_api(429, nil) do
      post "/tools/demos/lorem-ipsum", params: { paragraphs: "1", sentences: "5" }
    end
    assert_response :success
    assert_match I18n.t("tools.lorem_ipsum.demo.error_rate_limit"), response.body
  end

  test "lorem_ipsum renders error on non-200 response" do
    stub_api(500, nil) do
      post "/tools/demos/lorem-ipsum", params: { paragraphs: "1", sentences: "5" }
    end
    assert_response :success
    assert_match I18n.t("tools.lorem_ipsum.demo.error_generic"), response.body
  end

  test "lorem_ipsum renders error when data is nil" do
    stub_api(200, nil) do
      post "/tools/demos/lorem-ipsum", params: { paragraphs: "1", sentences: "5" }
    end
    assert_response :success
    assert_match I18n.t("tools.lorem_ipsum.demo.error_no_data"), response.body
  end

  # ── working_days ─────────────────────────────────────────────────────────────

  test "working_days renders result on success" do
    payload = { "working_days" => 4, "from" => "2024-02-23", "to" => "2024-02-28",
                "country" => "US", "subdivision" => "" }
    stub_api(200, success_data(payload)) do
      post "/tools/demos/working-days", params: { from: "2024-02-23", to: "2024-02-28", country: "US" }
    end
    assert_response :success
    assert_match "2024-02-23", response.body
  end

  test "working_days renders error when from or to is blank" do
    post "/tools/demos/working-days", params: { from: "", to: "2024-02-28" }
    assert_response :success
    assert_match I18n.t("tools.working_days.demo.error_empty"), response.body
  end

  test "working_days renders error when to is before from" do
    post "/tools/demos/working-days", params: { from: "2024-02-28", to: "2024-02-23" }
    assert_response :success
    assert_match I18n.t("tools.working_days.demo.error_invalid"), response.body
  end

  test "working_days renders error when dates are malformed" do
    post "/tools/demos/working-days", params: { from: "not-a-date", to: "2024-02-28" }
    assert_response :success
    assert_match I18n.t("tools.working_days.demo.error_invalid"), response.body
  end

  test "working_days renders error on 429 rate limit" do
    stub_api(429, nil) do
      post "/tools/demos/working-days", params: { from: "2024-02-23", to: "2024-02-28" }
    end
    assert_response :success
    assert_match I18n.t("tools.working_days.demo.error_rate_limit"), response.body
  end

  test "working_days renders error on non-200 response" do
    stub_api(500, nil) do
      post "/tools/demos/working-days", params: { from: "2024-02-23", to: "2024-02-28" }
    end
    assert_response :success
    assert_match I18n.t("tools.working_days.demo.error_generic"), response.body
  end

  test "working_days renders error when data is nil" do
    stub_api(200, nil) do
      post "/tools/demos/working-days", params: { from: "2024-02-23", to: "2024-02-28" }
    end
    assert_response :success
    assert_match I18n.t("tools.working_days.demo.error_no_data"), response.body
  end
end
