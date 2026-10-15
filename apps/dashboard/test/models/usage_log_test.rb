# frozen_string_literal: true

require "test_helper"

class UsageLogTest < ActiveSupport::TestCase
  def setup
    @user = create_user(email: "usage_log_test@example.com")
    @api_key = @user.api_keys.create!(name: "Test Key", environment: "test")
  end

  def create_log(used_at:, status_code:)
    UsageLog.create!(
      user: @user,
      api_key: @api_key,
      endpoint: "/v1/test",
      request_method: "POST",
      status_code: status_code,
      used_at: used_at
    )
  end

  test "in_range returns logs within the date range" do
    inside  = create_log(used_at: 5.days.ago,  status_code: 200)
    outside = create_log(used_at: 40.days.ago, status_code: 200)

    results = UsageLog.in_range(30.days.ago, Time.current)

    assert_includes     results, inside
    assert_not_includes results, outside
  end

  test "successful returns logs with 2xx status codes" do
    ok      = create_log(used_at: 1.day.ago, status_code: 200)
    created = create_log(used_at: 1.day.ago, status_code: 201)
    error   = create_log(used_at: 1.day.ago, status_code: 422)

    results = UsageLog.successful

    assert_includes     results, ok
    assert_includes     results, created
    assert_not_includes results, error
  end

  test "with_errors returns logs with status_code >= 400" do
    ok        = create_log(used_at: 1.day.ago, status_code: 200)
    not_found = create_log(used_at: 1.day.ago, status_code: 404)
    server    = create_log(used_at: 1.day.ago, status_code: 500)

    results = UsageLog.with_errors

    assert_not_includes results, ok
    assert_includes     results, not_found
    assert_includes     results, server
  end

  test "scopes can be chained" do
    inside_ok    = create_log(used_at: 5.days.ago,  status_code: 200)
    inside_error = create_log(used_at: 5.days.ago,  status_code: 500)
    outside_ok   = create_log(used_at: 40.days.ago, status_code: 200)

    results = UsageLog.in_range(30.days.ago, Time.current).with_errors

    assert_includes     results, inside_error
    assert_not_includes results, inside_ok
    assert_not_includes results, outside_ok
  end

  test "activity_success keeps nil, 0 and 200..399 and drops everything else" do
    kept = [ nil, 0, 200, 204, 302, 399 ].map { |code| create_log(used_at: 1.day.ago, status_code: code) }
    dropped = [ 101, 199, 400, 404, 500, 999 ].map { |code| create_log(used_at: 1.day.ago, status_code: code) }

    assert_equal kept.map(&:id).sort, UsageLog.activity_success.pluck(:id).sort
    assert_empty(UsageLog.activity_success.pluck(:id) & dropped.map(&:id))
  end

  test "for_activity_filter maps the allow list and treats anything else as all" do
    ok = create_log(used_at: 1.day.ago, status_code: 200)
    bad = create_log(used_at: 1.day.ago, status_code: 500)

    assert_equal [ ok.id ], UsageLog.for_activity_filter("success").pluck(:id)
    assert_equal [ bad.id ], UsageLog.for_activity_filter("errors").pluck(:id)
    assert_equal [ ok.id, bad.id ].sort, UsageLog.for_activity_filter("all").pluck(:id).sort
    assert_equal [ ok.id, bad.id ].sort, UsageLog.for_activity_filter("anything else").pluck(:id).sort
    assert_equal [ ok.id, bad.id ].sort, UsageLog.for_activity_filter(nil).pluck(:id).sort
  end

  test "recent orders newest first and breaks timestamp ties by id" do
    at = 1.day.ago.change(usec: 0)
    first = UsageLog.create!(user: @user, api_key: @api_key, endpoint: "/a", status_code: 200, used_at: at)
    second = UsageLog.create!(user: @user, api_key: @api_key, endpoint: "/b", status_code: 200, used_at: at)
    newest = UsageLog.create!(user: @user, api_key: @api_key, endpoint: "/c", status_code: 200, used_at: 1.hour.ago)

    assert_equal [ newest.id, second.id, first.id ], UsageLog.recent(10).pluck(:id)
    assert_equal 1, UsageLog.recent(1).size
  end
end
