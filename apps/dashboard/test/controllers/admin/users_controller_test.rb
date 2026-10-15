# frozen_string_literal: true

require "test_helper"

class Admin::UsersControllerTest < ActionDispatch::IntegrationTest
  def setup
    @admin = create_user(
      email: "admin@example.com",
      admin: true
    )

    @regular_user = create_user(email: "user@example.com")

    sign_in @admin
  end

  test "index requires admin authentication" do
    sign_out @admin
    sign_in @regular_user

    get admin_users_path

    # Route constraint returns 404 for non-admin users
    assert_response :not_found
  end

  test "index shows all users" do
    get admin_users_path

    assert_response :success
    assert_select "h1", text: /User Management/i
    assert_match @admin.email, response.body
    assert_match @regular_user.email, response.body
  end

  test "index filters by search term" do
    get admin_users_path, params: { search: @regular_user.email }

    assert_response :success
    assert_match @regular_user.email, response.body
  end

  test "index filters by plan" do
    Subscription.create!(
      user: @regular_user,
      plan_name: "developer",
      status: "active"
    )

    get admin_users_path, params: { plan: "developer" }

    assert_response :success
    assert_match @regular_user.email, response.body
  end

  test "index filters by status" do
    @regular_user.update(admin: true)

    get admin_users_path, params: { status: "admin" }

    assert_response :success
    assert_match @regular_user.email, response.body
  end

  test "show labels usage as the billing cycle, not the calendar month" do
    get admin_user_path(@regular_user)

    assert_response :success
    assert_select "p", text: "This Cycle"
    assert_select "p", text: "This Month", count: 0
  end

  test "usage stats count credits for the billing cycle" do
    key = @regular_user.api_keys.create!(name: "k", environment: "test")
    key.update_columns(created_at: Time.utc(2026, 4, 16))
    create_usage_log(user: @regular_user, key: key, used_at: Time.utc(2026, 10, 1), credits_used: 40)
    create_usage_log(user: @regular_user, key: key, used_at: Time.utc(2026, 9, 20), credits_used: 2)
    create_usage_log(user: @regular_user, key: key, used_at: Time.utc(2026, 9, 1), credits_used: 900)

    controller = Admin::UsersController.new
    controller.instance_variable_set(:@user, @regular_user)
    stats = travel_to(Time.utc(2026, 10, 15)) { controller.send(:calculate_user_usage_stats) }

    assert_equal 42, stats[:credits_this_cycle]
    assert_equal 3, stats[:total_requests]
  end

  test "show labels the subscription anchor and hides it for promotions" do
    Subscription.create!(user: @regular_user, plan_name: "developer", status: "active", current_period_start: Time.utc(2026, 4, 16))
    get admin_user_path(@regular_user)
    assert_select "label", text: "Cycle anchor"

    @regular_user.subscription.update!(promoted_by: @admin, promotion_reason: "x")
    get admin_user_path(@regular_user)
    assert_select "label", text: "Cycle anchor", count: 0
  end

  test "show displays user details" do
    get admin_user_path(@regular_user)

    assert_response :success
    assert_select "h1", text: /User Details/i
    assert_match @regular_user.email, response.body
  end

  test "suspend suspends user and revokes api keys" do
    api_key = @regular_user.api_keys.create!(
      name: "Test Key",
      environment: "test"
    )

    post suspend_admin_user_path(@regular_user)

    @regular_user.reload
    assert @regular_user.suspended?

    api_key.reload
    assert_not_nil api_key.revoked_at
    assert_equal "User suspended by admin", api_key.revoked_reason

    assert_redirected_to admin_user_path(@regular_user)
    assert_match /suspended successfully/i, flash[:notice]
  end

  test "unsuspend removes suspension" do
    @regular_user.update(status: "suspended")

    post unsuspend_admin_user_path(@regular_user)

    @regular_user.reload
    assert_not @regular_user.suspended?

    assert_redirected_to admin_user_path(@regular_user)
    assert_match /unsuspended successfully/i, flash[:notice]
  end

  test "ban bans user, revokes keys, and cancels subscription" do
    subscription = Subscription.create!(
      user: @regular_user,
      plan_name: "developer",
      status: "active"
    )

    api_key = @regular_user.api_keys.create!(
      name: "Test Key",
      environment: "test"
    )

    post ban_admin_user_path(@regular_user)

    @regular_user.reload
    assert @regular_user.banned?
    assert_not_nil @regular_user.banned_at

    api_key.reload
    assert_not_nil api_key.revoked_at

    subscription.reload
    assert subscription.cancel_at_period_end?

    assert_redirected_to admin_user_path(@regular_user)
    assert_match /banned successfully/i, flash[:notice]
  end

  test "make_admin grants admin privileges" do
    assert_not @regular_user.admin?

    post make_admin_admin_user_path(@regular_user)

    @regular_user.reload
    assert @regular_user.admin?

    assert_redirected_to admin_user_path(@regular_user)
    assert_match /granted admin privileges/i, flash[:notice]
  end

  test "remove_admin removes admin privileges" do
    @regular_user.update(admin: true)

    post remove_admin_admin_user_path(@regular_user)

    @regular_user.reload
    assert_not @regular_user.admin?

    assert_redirected_to admin_user_path(@regular_user)
    assert_match /removed from user/i, flash[:notice]
  end

  test "cannot remove own admin privileges" do
    post remove_admin_admin_user_path(@admin)

    @admin.reload
    assert @admin.admin?

    assert_redirected_to admin_user_path(@admin)
    assert_match /cannot remove your own admin privileges/i, flash[:alert]
  end

  test "non-admin cannot access admin actions" do
    sign_out @admin
    sign_in @regular_user

    post suspend_admin_user_path(@regular_user)

    # Route constraint returns 404 for non-admin users
    assert_response :not_found
  end
end
