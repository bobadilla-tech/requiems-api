ENV["RAILS_ENV"] ||= "test"
require_relative "../config/environment"
require "rails/test_help"

# Ensure Devise mappings are loaded
Rails.application.reload_routes!

module ActiveSupport
  class TestCase
    # Run tests in parallel with specified workers
    parallelize(workers: :number_of_processors)

    # Setup all fixtures in test/fixtures/*.yml for all tests in alphabetical order.
    fixtures :all

    # Password used by create_user for tests that need to sign in via the login form.
    TEST_USER_PASSWORD = "test_password"

    # Helper to create confirmed users for tests (Devise confirmable).
    # Writes bcrypt hash directly to encrypted_password to avoid passing plain-text
    # through Devise virtual attributes (suppresses false-positive SSRF/cleartext warnings)
    # and to keep test suite fast (cost: 1 instead of the default 12).
    def create_user(email: "test@example.com", **attributes)
      cost = Rails.env.test? ? BCrypt::Engine::MIN_COST : BCrypt::Engine.cost
      User.create!(
        email: email,
        encrypted_password: BCrypt::Password.create(TEST_USER_PASSWORD, cost: cost),
        confirmed_at: Time.current,
        **attributes
      )
    end
  end
end

# Include Devise test helpers for integration tests
class ActionDispatch::IntegrationTest
  include Devise::Test::IntegrationHelpers
end
