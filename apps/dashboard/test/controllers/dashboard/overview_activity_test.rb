# frozen_string_literal: true

require "test_helper"

class Dashboard::OverviewActivityTest < ActionDispatch::IntegrationTest
  NOW = Time.utc(2026, 10, 15, 12)
  XSS = %(/v1/"><img src=x onerror=alert(1)>)

  def setup
    @user = create_user(email: "activity@example.com")
    @key = @user.api_keys.create!(name: "Main key", environment: "test")
    @key.update_columns(created_at: Time.utc(2026, 4, 16, 8))
    sign_in @user
  end

  def log(status:, credits: 1, at: Time.utc(2026, 10, 14, 9), endpoint: nil, key: @key, user: @user, method: "GET", latency: 42)
    create_usage_log(user: user, key: key, used_at: at, credits_used: credits, status_code: status, endpoint: endpoint).tap do |row|
      row.update_columns(request_method: method, response_time_ms: latency)
    end
  end

  def visit(params = {}, **opts)
    travel_to(NOW) { get dashboard_root_path(params), **opts }
  end

  def rows
    css_select("[data-testid=activity-list] li")
  end

  test "an empty feed with a key points at the docs" do
    visit

    assert_select "#overview-activity [data-testid=empty-activity]"
    assert_select "[data-testid=empty-activity] a[href=?]", apis_path
    assert_select "[data-testid=create-key-cta]", 0
  end

  test "a user without keys is pointed at creating one first" do
    @key.update_columns(active: false, revoked_at: Time.current)
    visit

    assert_select "[data-testid=create-key-cta][href=?]", new_dashboard_api_key_path
  end

  test "the feed lives in the overview-activity turbo frame" do
    log(status: 200)
    visit

    assert_select "turbo-frame#overview-activity [data-testid=activity-list]"
  end

  test "a turbo frame request returns the frame with the filtered rows" do
    log(status: 200)
    log(status: 404, credits: 0)

    visit({ activity: "errors" }, headers: { "Turbo-Frame" => "overview-activity" })

    assert_response :success
    assert_select "turbo-frame#overview-activity"
    assert_select "nav a[autofocus][data-filter=errors]", 1
    assert_equal 1, rows.size
  end

  test "a full page load never autofocuses a filter chip" do
    visit({ activity: "errors" })

    assert_select "nav a[autofocus]", 0
  end

  test "the overview opts into Turbo visits so the filters advance without a reload, other pages do not" do
    visit
    assert_select "meta[name=turbo-visit-control]", 0

    get dashboard_billing_path
    assert_select "meta[name=turbo-visit-control][content=reload]", 1
  end

  test "rows show status chip with icon and screen reader label, method, latency, credits, key and time" do
    log(status: 200, credits: 2, endpoint: "/v1/text/dictionary/hello", latency: 1234)
    visit

    row = rows.first
    chip = row.at_css("span.inline-flex")
    assert_equal "success", row["data-status-kind"]
    assert chip.at_css("svg[aria-hidden=true]")
    assert_equal "Success", chip.at_css(".sr-only").text
    assert_includes chip.text, "200"
    assert_includes row.text, "GET"
    assert_includes row.text, "/v1/text/dictionary/hello"
    assert_includes row.text, "1,234 ms"
    assert_equal "2 credits", row.at_css("[data-testid=credits]").text.strip
    assert_equal "Main key", row.at_css("[data-testid=key-name]").text.strip
    assert_includes row.at_css("time").text, "ago"
  end

  test "status chips cover success, redirect, client and server errors with distinct labels" do
    { 200 => "Success", 302 => "Redirect", 404 => "Client error", 503 => "Server error" }.each_with_index do |(status, _label), i|
      log(status: status, at: Time.utc(2026, 10, 14, 9, i))
    end
    visit

    labels = rows.map { |r| r.at_css(".sr-only").text }
    assert_equal [ "Server error", "Client error", "Redirect", "Success" ], labels
    assert_equal %w[server_error client_error redirect success], rows.map { |r| r["data-status-kind"] }
  end

  test "legacy status 0 and NULL render as successful" do
    log(status: 0)
    log(status: nil, at: Time.utc(2026, 10, 14, 10))
    visit

    assert_equal %w[success success], rows.map { |r| r["data-status-kind"] }
    assert_includes rows.first.text, "OK"
  end

  test "legacy error rows render with their stored credits in all three variants" do
    log(status: 404, credits: 1, at: Time.utc(2026, 10, 14, 9, 1), endpoint: "/v1/text/advice")
    log(status: 404, credits: 0, at: Time.utc(2026, 10, 14, 9, 2), endpoint: "/api/v1/text/advice")
    log(status: 404, credits: nil, at: Time.utc(2026, 10, 14, 9, 3), endpoint: "/v1/text/old")
    visit

    by_endpoint = rows.to_h { |r| [ r.at_css("code").text, r ] }

    assert_equal "1 credit (charged)", by_endpoint["/v1/text/advice"].at_css("[data-testid=credits]").text.strip
    assert_equal "not charged", by_endpoint["/api/v1/text/advice"].at_css("[data-testid=credits]").text.strip
    assert_equal "not recorded", by_endpoint["/v1/text/old"].at_css("[data-testid=credits]").text.strip
    assert(by_endpoint.values.all? { |r| r["data-status-kind"] == "client_error" }, "each carries an error chip")
    assert(by_endpoint.values.all? { |r| r.at_css(".sr-only").text == "Client error" && r.at_css("span.inline-flex").to_s.include?("bg-yellow-100") })
  end

  test "a successful row with NULL credits is not recorded rather than 1 credit" do
    log(status: 200, credits: nil)
    visit

    assert_equal "not recorded", rows.first.at_css("[data-testid=credits]").text.strip
  end

  test "hostile endpoints are escaped" do
    log(status: 200, endpoint: XSS)
    visit

    doc = Nokogiri::HTML(response.body)
    assert_empty doc.css("#overview-activity img[onerror]")
    assert_includes doc.at_css("#overview-activity code").text, "<img src=x onerror=alert(1)>"
    assert_not_includes response.body, XSS
  end

  test "a hostile key name is escaped and a revoked key is labelled instead of named" do
    @key.update_columns(name: %(<b>bold</b>))
    log(status: 200)
    visit
    doc = Nokogiri::HTML(response.body)
    assert_empty doc.css("#overview-activity [data-testid=key-name] b")

    @key.update_columns(revoked_at: Time.current, active: false)
    other = @user.api_keys.create!(name: "Fresh", environment: "test")
    other.update_columns(created_at: Time.utc(2026, 6, 1))
    visit

    name = css_select("[data-testid=key-name]").first.text.strip
    assert_equal "Revoked key", name
  end

  test "the time element carries a machine readable and a localized UTC title" do
    log(status: 200, at: Time.utc(2026, 10, 14, 9, 5, 7))
    visit

    time = css_select("[data-testid=activity-list] time").first
    assert_equal "2026-10-14T09:05:07Z", time["datetime"]
    assert_equal "Oct 14, 2026 09:05:07 UTC", time["title"]

    visit({ locale: "es" })
    assert_equal "14 oct 2026 09:05:07 UTC", css_select("[data-testid=activity-list] time").first["title"]
    visit({ locale: "fr" })
    assert_match(/\A14 oct\.? 2026 09:05:07 UTC\z/, css_select("[data-testid=activity-list] time").first["title"])
  end

  test "the feed shows the ten most recent rows, newest first" do
    12.times { |i| log(status: 200, at: Time.utc(2026, 10, 1, i), endpoint: "/v1/n/#{i}") }
    visit

    assert_equal 10, rows.size
    assert_equal "/v1/n/11", rows.first.at_css("code").text
    assert_equal "/v1/n/2", rows.last.at_css("code").text
  end

  test "an old ledger shows rows plus the stale line with the last request age" do
    log(status: 200, at: Time.utc(2026, 7, 12))
    visit

    assert_equal 1, rows.size
    assert_select "[data-testid=stale-activity]", text: /No requests this cycle · last request 3 months ago/
  end

  test "a current ledger has no stale line" do
    log(status: 200)
    visit

    assert_select "[data-testid=stale-activity]", 0
  end

  test "filters are real links with aria-current on the active one" do
    visit({ activity: "errors" })

    assert_select "nav[aria-label='Filter recent activity'] a", 3
    assert_select "nav a[aria-current=true]", 1
    assert_select "nav a[aria-current=true][data-filter=errors]"
    assert_select "nav a[data-turbo-action=advance]", 3
    assert_select "nav a[data-filter=errors][href=?]", dashboard_root_path(activity: "errors")
    assert_select "nav a[data-filter=all][href=?]", dashboard_root_path
  end

  test "the success filter keeps 2xx, 3xx and legacy rows and drops errors" do
    log(status: 200, at: Time.utc(2026, 10, 14, 1))
    log(status: 302, at: Time.utc(2026, 10, 14, 2))
    log(status: 0, at: Time.utc(2026, 10, 14, 3))
    log(status: 404, at: Time.utc(2026, 10, 14, 4))
    log(status: 500, at: Time.utc(2026, 10, 14, 5))
    visit({ activity: "success" })

    assert_equal %w[success redirect success], rows.map { |r| r["data-status-kind"] }
    assert_equal 3, rows.size
  end

  test "the errors filter keeps 4xx and 5xx only" do
    log(status: 200, at: Time.utc(2026, 10, 14, 1))
    log(status: 404, at: Time.utc(2026, 10, 14, 2))
    log(status: 500, at: Time.utc(2026, 10, 14, 3))
    visit({ activity: "errors" })

    assert_equal %w[server_error client_error], rows.map { |r| r["data-status-kind"] }
  end

  test "a filter with no matches explains itself" do
    log(status: 200)
    visit({ activity: "errors" })

    assert_select "[data-testid=no-matching-activity]"
    assert_select "[data-testid=empty-activity]", 0
  end

  test "unknown or malformed activity params fall back to all and never error" do
    log(status: 200)
    log(status: 404, at: Time.utc(2026, 10, 14, 10))

    [ "bogus", "'; DROP TABLE usage_logs;--", "ERRORS", "", "all" ].each do |value|
      visit({ activity: value })
      assert_response :success
      assert_equal 2, rows.size, value.inspect
    end

    travel_to(NOW) { get "#{dashboard_root_path}?activity[]=errors" }
    assert_response :success
    assert_equal 2, rows.size

    travel_to(NOW) { get "#{dashboard_root_path}?activity[a]=errors" }
    assert_response :success
    assert_equal 2, rows.size
  end

  test "another user's rows never appear, whatever the params" do
    other = create_user(email: "activity-other@example.com")
    other_key = other.api_keys.create!(name: "Other key", environment: "test")
    log(status: 500, user: other, key: other_key, endpoint: "/v1/secret")
    log(status: 200, endpoint: "/v1/mine")

    [ {}, { activity: "errors" }, { activity: "success" }, { user_id: other.id }, { activity: "errors", user_id: other.id, id: other.id } ].each do |params|
      visit(params)

      assert_not_includes response.body, "/v1/secret", params.inspect
    end
  end

  test "the list and filters carry accessible names" do
    log(status: 200)
    visit

    assert_select "div[role=region][tabindex='0'][aria-label='Recent API requests'] ul[data-testid=activity-list]"
    assert_select "nav[aria-label='Filter recent activity']"
  end

  %w[es fr].each do |locale|
    test "the feed renders in #{locale} with translated chips and credits" do
      log(status: 404, credits: 1)
      log(status: 200, credits: 2, at: Time.utc(2026, 10, 14, 10))
      visit({ locale: locale })

      assert_response :success
      assert_no_match(/translation missing/i, response.body)
      credits = css_select("[data-testid=credits]").map { |c| c.text.strip }
      assert_equal 2, credits.size
      assert(credits.all? { |c| c.match?(/crédit/) }, credits.inspect)
    end
  end

  test "there is no horizontal scrolling container in the feed" do
    log(status: 200, endpoint: "/v1/" + ("verylongsegment" * 12))
    visit

    html = css_select("#overview-activity").to_s
    assert_not_includes html, "overflow-x-auto"
    assert_includes html, "break-all"
    assert_includes html, "flex-wrap"
  end

  test "rows with the same timestamp keep a stable newest-id-first order" do
    first = log(status: 200, at: Time.utc(2026, 10, 14, 9), endpoint: "/v1/tie/first")
    second = log(status: 200, at: Time.utc(2026, 10, 14, 9), endpoint: "/v1/tie/second")
    assert_operator second.id, :>, first.id
    visit

    assert_equal %w[/v1/tie/second /v1/tie/first], rows.map { |r| r.at_css("code").text }
  end

  test "missing method and latency render dashes and large latencies get delimiters" do
    row = log(status: 200, latency: 1_234_567)
    row.update_columns(request_method: nil)
    other = log(status: 200, at: Time.utc(2026, 10, 14, 10))
    other.update_columns(response_time_ms: nil)
    visit

    texts = rows.map(&:text)
    assert(texts.any? { |t| t.include?("1,234,567 ms") })
    assert(texts.all? { |t| t.include?("—") }, "each row has a dash for the missing field")
  end

  test "statuses above 599 are server errors, free of the success filter, and kept by the errors filter" do
    log(status: 999)
    visit({ activity: "errors" })
    assert_equal %w[server_error], rows.map { |r| r["data-status-kind"] }

    visit({ activity: "success" })
    assert_select "[data-testid=no-matching-activity]"
  end

  test "an informational status renders an unknown chip under All only" do
    log(status: 101)
    visit
    assert_equal %w[unknown], rows.map { |r| r["data-status-kind"] }
    assert_equal "Unknown status", rows.first.at_css(".sr-only").text

    visit({ activity: "success" })
    assert_equal 0, rows.size
    visit({ activity: "errors" })
    assert_equal 0, rows.size
  end

  test "the success filter with nothing to show explains itself" do
    log(status: 404)
    visit({ activity: "success" })

    assert_select "[data-testid=no-matching-activity]"
  end

  test "the success filter hides another user's successful rows as well" do
    other = create_user(email: "activity-other2@example.com")
    other_key = other.api_keys.create!(name: "Other key", environment: "test")
    log(status: 200, user: other, key: other_key, endpoint: "/v1/theirs/ok")
    log(status: 200, endpoint: "/v1/mine/ok")
    visit({ activity: "success" })

    assert_equal [ "/v1/mine/ok" ], rows.map { |r| r.at_css("code").text }
  end

  %w[es fr].each do |locale|
    test "chip, filters and legacy credit labels are translated in #{locale}" do
      log(status: 404, credits: 0, at: Time.utc(2026, 10, 14, 9, 1))
      log(status: 404, credits: nil, at: Time.utc(2026, 10, 14, 9, 2))
      visit({ locale: locale })

      assert_equal I18n.t("dashboard.overview.activity.status.client_error", locale: locale), rows.first.at_css(".sr-only").text
      assert_equal [ I18n.t("dashboard.overview.activity.credits_unrecorded", locale: locale), I18n.t("dashboard.overview.activity.credits_none", locale: locale) ],
                   css_select("[data-testid=credits]").map { |c| c.text.strip }
      labels = css_select("nav a[data-filter]").map { |a| a.text.strip }
      assert_equal %w[all success errors].map { |k| I18n.t("dashboard.overview.activity.filter_#{k}", locale: locale) }, labels
    end
  end

  test "the response varies on Turbo-Frame" do
    visit

    assert_includes response.headers["Vary"].to_s, "Turbo-Frame"
  end

  test "a polite status summary announces how many rows show" do
    2.times { |i| log(status: 200, at: Time.utc(2026, 10, 14, i)) }
    visit

    assert_select "[data-testid=activity-summary][role=status]", text: /Showing 2 recent requests/

    visit({ activity: "errors" })
    assert_select "[data-testid=activity-summary][role=status]", text: /No matching requests/
  end

  test "credit labels use the singular for one credit and charge wording on 5xx errors" do
    log(status: 200, credits: 1, at: Time.utc(2026, 10, 14, 9, 1), endpoint: "/v1/one")
    log(status: 503, credits: 1, at: Time.utc(2026, 10, 14, 9, 2), endpoint: "/v1/five-hundred")
    log(status: 503, credits: 3, at: Time.utc(2026, 10, 14, 9, 3), endpoint: "/v1/five-hundred-three")
    visit

    labels = rows.to_h { |r| [ r.at_css("code").text, r.at_css("[data-testid=credits]").text.strip ] }
    assert_equal "1 credit", labels["/v1/one"]
    assert_equal "1 credit (charged)", labels["/v1/five-hundred"]
    assert_equal "3 credits (charged)", labels["/v1/five-hundred-three"]
  end
end
