# frozen_string_literal: true

require "active_support/core_ext/integer/time"

Rails.application.configure do
  config.enable_reloading = false
  config.eager_load = true
  config.consider_all_requests_local = false
  config.public_file_server.enabled = true
  config.action_controller.perform_caching = true
  config.public_file_server.headers = { "cache-control" => "public, max-age=#{1.year.to_i}, immutable" }
  config.log_tags = [ :request_id ]
  config.logger   = ActiveSupport::TaggedLogging.logger(STDOUT)
  config.log_level = ENV.fetch("RAILS_LOG_LEVEL", "info")
  config.silence_healthcheck_path = [ "/up", "/favicon.ico" ]
  config.active_support.report_deprecations = false

  config.cache_store = :redis_cache_store, { url: ENV["REDIS_URL"], namespace: "rails_cache" }
  config.active_job.queue_adapter = :sidekiq

  config.action_mailer.raise_delivery_errors = true
  config.action_mailer.perform_caching = false

  config.action_mailer.delivery_method = :smtp

  config.after_initialize do
    ActionMailer::Base.default_url_options = {
      host: AppConfig.mailer_host,
      protocol: "https"
    }

    ActionMailer::Base.smtp_settings = {
      address: AppConfig.smtp_address,
      port: AppConfig.smtp_port,
      domain: AppConfig.smtp_domain,
      user_name: AppConfig.smtp_username,
      password: AppConfig.smtp_password,
      authentication: :plain,
      enable_starttls_auto: true
    }
  end

  config.i18n.fallbacks = true
  config.active_record.dump_schema_after_migration = false
  config.active_record.attributes_for_inspect = [ :id ]
end
