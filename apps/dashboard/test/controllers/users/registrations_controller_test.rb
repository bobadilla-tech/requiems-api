# frozen_string_literal: true

require "test_helper"

class Users::RegistrationsControllerTest < ActionDispatch::IntegrationTest
  def setup
    @referrer = create_user(email: "referrer@example.com")
    ActionMailer::Base.deliveries.clear
  end

  test "signup without ref param creates user without referral" do
    assert_difference "User.count", 1 do
      assert_no_difference "Referral.count" do
        post user_registration_path, params: {
          user: {
            email: "newuser@example.com",
            password: TEST_USER_PASSWORD,
            password_confirmation: TEST_USER_PASSWORD
          }
        }
      end
    end
  end

  test "GET new with ref param stores token in session" do
    token = @referrer.referral_token

    get new_user_registration_path(ref: token)

    assert_response :success
    assert_equal token, session[:referral_token]
  end

  test "signup with valid ref token creates referral attribution" do
    token = @referrer.referral_token

    get new_user_registration_path(ref: token)

    assert_difference "Referral.count", 1 do
      post user_registration_path, params: {
        user: {
          email: "referred@example.com",
          password: TEST_USER_PASSWORD,
          password_confirmation: TEST_USER_PASSWORD
        }
      }
    end

    new_user = User.find_by(email: "referred@example.com")
    assert new_user.referral_received.present?
    assert_equal @referrer, new_user.referral_received.referrer
  end

  test "signup with tampered ref token silently ignores referral" do
    get new_user_registration_path(ref: "tampered_token_xyz")

    assert_no_difference "Referral.count" do
      post user_registration_path, params: {
        user: {
          email: "organic@example.com",
          password: TEST_USER_PASSWORD,
          password_confirmation: TEST_USER_PASSWORD
        }
      }
    end
  end

  test "self-referral is silently ignored" do
    # Referrer visits their own referral link and tries to register a second time
    # (simulate: same user ID via token)
    token = @referrer.referral_token

    # We can't re-register the same user; simulate by creating a new user whose
    # referral would be blocked by the self-referral check using a different email
    # but faking the same referrer ID. Instead, test via the Referral model directly.
    # The controller attribute_referral method skips when referrer.id == user.id.
    assert_no_difference "Referral.count" do
      # Build a user with same id as referrer to exercise the self-referral guard
      referral = Referral.new(referrer: @referrer, referred_user: @referrer)
      assert_not referral.valid?
    end
  end

  test "referral token is cleared from session after signup" do
    token = @referrer.referral_token

    get new_user_registration_path(ref: token)
    assert session[:referral_token].present?

    post user_registration_path, params: {
      user: {
        email: "cleartest@example.com",
        password: TEST_USER_PASSWORD,
        password_confirmation: TEST_USER_PASSWORD
      }
    }

    assert_nil session[:referral_token]
  end
end
