# frozen_string_literal: true

require "test_helper"

class Devise::PasswordsUiTest < ActionDispatch::IntegrationTest
  test "new password page renders styled form" do
    get new_user_password_path

    assert_response :success
    assert_select "div.min-h-screen"
    assert_select "form.mt-8.space-y-6"
    assert_select "input[type='email'][name='user[email]']"
  end

  test "edit password page renders styled form with reset token" do
    user = create_user(email: "reset-ui@example.com")
    reset_token = user.send_reset_password_instructions

    get edit_user_password_path(reset_password_token: reset_token)

    assert_response :success
    assert_select "div.min-h-screen"
    assert_select "form.mt-8.space-y-6"
    assert_select "input[type='hidden'][name='user[reset_password_token]'][value=?]", reset_token
    assert_select "input[type='password'][name='user[password]']"
    assert_select "input[type='password'][name='user[password_confirmation]']"
  end

  test "requesting reset instructions renders notice toast" do
    user = create_user(email: "toast-reset@example.com")

    post user_password_path, params: { user: { email: user.email } }
    follow_redirect!

    assert_response :success
    assert_select "div.fixed.top-4.right-4.w-full.max-w-md"
    assert_select "div[data-controller='flash']"
    assert_includes response.body, I18n.t("devise.passwords.send_instructions")
  end
end
