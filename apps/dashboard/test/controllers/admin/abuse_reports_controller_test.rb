# frozen_string_literal: true

require "test_helper"

class Admin::AbuseReportsControllerTest < ActionDispatch::IntegrationTest
  def setup
    @admin = create_user(email: "abuse-admin@example.com", admin: true)
    @reported = create_user(email: "reported-user@example.com")
    @key = @reported.api_keys.create!(name: "k", environment: "test")
    @key.update_columns(created_at: Time.utc(2026, 4, 16))
    @report = AbuseReport.create!(user: @reported, api_key: @key, report_type: "spam", status: "pending")

    sign_in @admin
  end

  test "show builds usage stats from credits for the billing cycle, not the calendar month" do
    create_usage_log(user: @reported, key: @key, used_at: Time.utc(2026, 10, 1), credits_used: 40)
    create_usage_log(user: @reported, key: @key, used_at: Time.utc(2026, 9, 20), credits_used: 2)
    create_usage_log(user: @reported, key: @key, used_at: Time.utc(2026, 9, 1), credits_used: 900)

    # The show view itself raises today (ApiKey#prefix does not exist, an
    # out-of-scope bug noted in the plan), so read the stats the action builds.
    controller = Admin::AbuseReportsController.new
    controller.instance_variable_set(:@abuse_report, @report)
    travel_to(Time.utc(2026, 10, 15)) { controller.show }
    stats = controller.instance_variable_get(:@usage_stats)

    assert_equal 42, stats[:credits_this_cycle]
    assert_equal 3, stats[:total_requests]
  end
end
