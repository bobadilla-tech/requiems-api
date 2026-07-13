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
    assert_match I18n.t("tools.sudoku.demo.error_empty"), response.body
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
end
