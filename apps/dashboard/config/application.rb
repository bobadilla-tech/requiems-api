# frozen_string_literal: true

require_relative "boot"

require "rails"

require "active_model/railtie"
require "active_job/railtie"
require "active_record/railtie"
require "action_controller/railtie"
require "action_mailer/railtie"
require "action_view/railtie"

Bundler.require(*Rails.groups)

module Dashboard
  class Application < Rails::Application
    config.load_defaults 8.1

    config.autoload_lib(ignore: %w[assets tasks])

    config.generators.system_tests = nil

    config.middleware.use Rack::Attack

    config.i18n.available_locales = %i[en es fr]
    config.i18n.default_locale = :en
    config.i18n.fallbacks = true
    config.i18n.load_path += Dir[Rails.root.join("config/locales/**/*.yml")]

    # Exclude Go-managed tables from the Rails schema dump.
    # These tables are created and owned by the Go API (apps/api/migrations/).
    # Rails and Go share one PostgreSQL database, so without this the schema
    # dumper would include all tables on every db:test:load_schema run.
    initializer "schema_dumper.ignore_go_tables" do
      ActiveRecord::SchemaDumper.ignore_tables = %w[
        advice
        bin_data
        commodity_price_history
        commodity_prices
        counters
        exercises
        iban_countries
        inflation_data
        quotes
        schema_migrations
        swift_codes
        words
      ]
    end

    # ActiveRecord Encryption for sensitive columns (tenant_secret on PrivateDeploymentRequest).
    # Keys are loaded from environment variables. Generate with: SecureRandom.base64(32)
    #   AR_ENCRYPTION_PRIMARY_KEY
    #   AR_ENCRYPTION_DETERMINISTIC_KEY
    #   AR_ENCRYPTION_KEY_DERIVATION_SALT
    if ENV["AR_ENCRYPTION_PRIMARY_KEY"].present?
      config.active_record.encryption.primary_key          = ENV["AR_ENCRYPTION_PRIMARY_KEY"]
      config.active_record.encryption.deterministic_key    = ENV["AR_ENCRYPTION_DETERMINISTIC_KEY"]
      config.active_record.encryption.key_derivation_salt  = ENV["AR_ENCRYPTION_KEY_DERIVATION_SALT"]
    end

    # Fixed effective date for legal documents (Privacy Policy, Terms of Service).
    # Update this when the documents are materially revised.
    config.x.legal_effective_date = Date.new(2026, 2, 17)
  end
end
