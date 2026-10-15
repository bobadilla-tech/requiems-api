# frozen_string_literal: true

require "test_helper"
require "csv"

class Dashboard::UsageControllerTest < ActionDispatch::IntegrationTest
  XSS_ENDPOINT = %(/v1/"><img src=x onerror=alert(1)>)

  def setup
    @user = create_user(email: "usage-controller@example.com")
    @api_key = @user.api_keys.create!(name: "Main", environment: "test")
    sign_in @user
  end

  def create_log(endpoint:, user: @user, api_key: @api_key, method: "GET", status: 200)
    UsageLog.create!(
      user: user, api_key: api_key, endpoint: endpoint, request_method: method,
      status_code: status, credits_used: 1, used_at: 1.hour.ago
    )
  end

  test "show escapes hostile endpoints in the endpoint usage list" do
    create_log(endpoint: XSS_ENDPOINT)

    get dashboard_usage_path

    assert_response :success
    doc = Nokogiri::HTML(response.body)
    assert_empty doc.css("img[onerror]")
    assert_empty doc.css("#usage-by-endpoint-container img")
    assert_includes doc.at_css("#usage-by-endpoint-container code").text, "<img src=x onerror=alert(1)>"
    assert_not_includes response.body, XSS_ENDPOINT
  end

  test "show does not build endpoint rows with innerHTML in page scripts" do
    get dashboard_usage_path

    assert_not_includes response.body, "usage-by-endpoint-container').innerHTML"
    assert_not_includes response.body, "${endpoint}"
  end

  test "show lists endpoints by count and shows empty state without logs" do
    get dashboard_usage_path
    assert_select "#usage-by-endpoint-container code", count: 0

    2.times { create_log(endpoint: "/v1/a") }
    create_log(endpoint: "/v1/b")
    get dashboard_usage_path

    assert_select "#usage-by-endpoint-container code", count: 2
    assert_equal "/v1/a", css_select("#usage-by-endpoint-container code").first.text
  end

  test "show breaks count ties by endpoint and respects the date range" do
    create_log(endpoint: "/v1/b")
    create_log(endpoint: "/v1/a")
    create_log(endpoint: "/v1/old").update!(used_at: 60.days.ago)

    get dashboard_usage_path

    assert_equal %w[/v1/a /v1/b], css_select("#usage-by-endpoint-container code").map(&:text)

    get dashboard_usage_path(range: "90d")
    assert_includes css_select("#usage-by-endpoint-container code").map(&:text), "/v1/old"
  end

  test "show caps the endpoint list at 100 entries" do
    101.times { |i| create_log(endpoint: "/v1/e#{i}") }

    get dashboard_usage_path

    assert_select "#usage-by-endpoint-container code", count: 100
  end

  test "show does not list other users endpoints" do
    other = create_user(email: "other-usage@example.com")
    other_key = other.api_keys.create!(name: "Theirs", environment: "test")
    create_log(endpoint: "/v1/secret", user: other, api_key: other_key)

    get dashboard_usage_path

    assert_not_includes response.body, "/v1/secret"
  end

  test "export neutralises spreadsheet formula injection" do
    [ "=HYPERLINK(\"http://evil\")", "+1+1", "-2+3", "@SUM(A1)", "\tcmd", "\rcmd" ].each do |payload|
      create_log(endpoint: payload)
    end
    @api_key.update!(name: "=1+1")

    get export_dashboard_usage_path

    assert_response :success
    rows = CSV.parse(response.body, headers: true)
    endpoints = rows.map { |r| r["Endpoint"] }
    assert_equal 6, endpoints.size
    endpoints.each { |e| assert e.start_with?("'"), "#{e.inspect} should be prefixed" }
    assert(rows.all? { |r| r["API Key"] == "'=1+1" })
  end

  test "export neutralises formulas in the method and with leading whitespace" do
    create_log(endpoint: " =1+1", method: "=cmd")
    create_log(endpoint: "\n=1+1")

    get export_dashboard_usage_path

    rows = CSV.parse(response.body, headers: true)
    assert(rows.all? { |r| r["Endpoint"].start_with?("'") })
    assert_includes rows.map { |r| r["Method"] }, "'=cmd"
  end

  test "export is ordered newest first" do
    create_log(endpoint: "/v1/old").update!(used_at: 3.days.ago)
    create_log(endpoint: "/v1/new")

    get export_dashboard_usage_path

    assert_equal %w[/v1/new /v1/old], CSV.parse(response.body, headers: true).map { |r| r["Endpoint"] }
  end

  test "export leaves ordinary cells untouched" do
    create_log(endpoint: "/v1/text/advice")

    get export_dashboard_usage_path

    row = CSV.parse(response.body, headers: true).first
    assert_equal "/v1/text/advice", row["Endpoint"]
    assert_equal "GET", row["Method"]
    assert_equal "Main", row["API Key"]
  end

  test "export only includes the signed-in user's rows" do
    other = create_user(email: "other-export@example.com")
    other_key = other.api_keys.create!(name: "Theirs", environment: "test")
    create_log(endpoint: "/v1/secret", user: other, api_key: other_key)

    get export_dashboard_usage_path

    assert_not_includes response.body, "/v1/secret"
  end
end
