# frozen_string_literal: true

require "simplecov"
require "simplecov-lcov"

SimpleCov::Formatter::LcovFormatter.config.report_with_single_file = true
SimpleCov.formatter = SimpleCov::Formatter::LcovFormatter
SimpleCov.start "rails"

ENV["RAILS_ENV"] ||= "test"
require_relative "../config/environment"
require "rails/test_help"

Rails.application.reload_routes!

module ActiveSupport
  class TestCase
    parallelize(workers: :number_of_processors)

    parallelize_setup do |worker|
      SimpleCov.command_name "#{SimpleCov.command_name}-#{worker}"
    end

    parallelize_teardown do |_worker|
      SimpleCov.result
    end

    fixtures :all

    # ApplicationController#set_locale mutates the global I18n.locale per
    # request without resetting it. Real Puma requests are safe (i18n's
    # config is thread-local), but parallel test workers run many test
    # methods sequentially on the same thread, so a locale change from one
    # test otherwise leaks into whichever test runs next on that worker.
    teardown do
      I18n.locale = I18n.default_locale
    end

    TEST_USER_PASSWORD = "password123!"

    def create_user(email: "test@example.com", **attributes)
      User.create!(
        email: email,
        password: TEST_USER_PASSWORD,
        password_confirmation: TEST_USER_PASSWORD,
        confirmed_at: Time.current,
        **attributes
      )
    end
  end
end

class ActionDispatch::IntegrationTest
  include Devise::Test::IntegrationHelpers

  setup do
    self.default_url_options = { locale: I18n.default_locale }
  end
end
