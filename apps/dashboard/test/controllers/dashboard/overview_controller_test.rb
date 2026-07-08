# frozen_string_literal: true

require "test_helper"

class Dashboard::OverviewControllerTest < ActionDispatch::IntegrationTest
  def setup
    @user = create_user(
      email: "test@example.com",
      password: "password123",
      password_confirmation: "password123"
    )

    sign_in @user
  end

  test "index requires authentication" do
    sign_out @user
    get dashboard_root_path

    assert_redirected_to new_user_session_path
  end

  test "index renders successfully" do
    get dashboard_root_path

    assert_response :success
  end

  %w[en es fr].each do |locale|
    test "index renders successfully in #{locale} locale" do
      get dashboard_root_path(locale: locale)

      assert_response :success
    end
  end
end
