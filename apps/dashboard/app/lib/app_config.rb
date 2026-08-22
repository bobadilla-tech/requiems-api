# frozen_string_literal: true

class AppConfig
  class MissingConfigError < StandardError; end
  class InvalidConfigError < StandardError; end

  include Singleton

  attr_reader :api_management_api_key

  attr_reader :lemonsqueezy_store_id,
              :lemonsqueezy_store_slug,
              :lemonsqueezy_signing_secret,
              :lemonsqueezy_test_mode

  attr_reader :lemonsqueezy_developer_monthly_variant_id,
              :lemonsqueezy_developer_yearly_variant_id,
              :lemonsqueezy_business_monthly_variant_id,
              :lemonsqueezy_business_yearly_variant_id,
              :lemonsqueezy_professional_monthly_variant_id,
              :lemonsqueezy_professional_yearly_variant_id

  attr_reader :lemonsqueezy_developer_monthly_checkout_uuid,
              :lemonsqueezy_developer_yearly_checkout_uuid,
              :lemonsqueezy_business_monthly_checkout_uuid,
              :lemonsqueezy_business_yearly_checkout_uuid,
              :lemonsqueezy_professional_monthly_checkout_uuid,
              :lemonsqueezy_professional_yearly_checkout_uuid

  # Private deployment tier checkout UUIDs (monthly + yearly per tier)
  attr_reader :lemonsqueezy_private_starter_monthly_checkout_uuid,
              :lemonsqueezy_private_starter_yearly_checkout_uuid,
              :lemonsqueezy_private_growth_monthly_checkout_uuid,
              :lemonsqueezy_private_growth_yearly_checkout_uuid,
              :lemonsqueezy_private_scale_monthly_checkout_uuid,
              :lemonsqueezy_private_scale_yearly_checkout_uuid,
              :lemonsqueezy_private_enterprise_monthly_checkout_uuid,
              :lemonsqueezy_private_enterprise_yearly_checkout_uuid

  attr_reader :api_base_url,
              :playground_api_key,
              :internal_api_url,
              :backend_secret

  attr_reader :smtp_address,
              :smtp_port,
              :smtp_domain,
              :smtp_username,
              :smtp_password,
              :mailer_host

  def initialize
    load_config
    validate_config unless ENV["SECRET_KEY_BASE_DUMMY"].present?
  end

  def self.instance
    @instance ||= new
  end

  def self.private_deployment_checkout_uuid_for(tier:, billing_cycle:)
    instance.private_deployment_checkout_uuid_for(tier:, billing_cycle:)
  end

  def self.method_missing(method_name, *args, **kwargs, &block)
    if instance.respond_to?(method_name)
      instance.public_send(method_name, *args, **kwargs, &block)
    else
      super
    end
  end

  def self.respond_to_missing?(method_name, include_private = false)
    instance.respond_to?(method_name) || super
  end

  def variant_id_for(plan:, billing_cycle:)
    case plan.to_s.downcase
    when "developer"
      billing_cycle == "monthly" ? lemonsqueezy_developer_monthly_variant_id : lemonsqueezy_developer_yearly_variant_id
    when "business"
      billing_cycle == "monthly" ? lemonsqueezy_business_monthly_variant_id : lemonsqueezy_business_yearly_variant_id
    when "professional"
      billing_cycle == "monthly" ? lemonsqueezy_professional_monthly_variant_id : lemonsqueezy_professional_yearly_variant_id
    else
      raise InvalidConfigError, "Unknown plan: #{plan}"
    end
  end

  def plan_name_for_variant_id(variant_id)
    variant_id = variant_id.to_s
    case variant_id
    when lemonsqueezy_developer_monthly_variant_id, lemonsqueezy_developer_yearly_variant_id
      "developer"
    when lemonsqueezy_business_monthly_variant_id, lemonsqueezy_business_yearly_variant_id
      "business"
    when lemonsqueezy_professional_monthly_variant_id, lemonsqueezy_professional_yearly_variant_id
      "professional"
    else
      nil
    end
  end

  # The webhook payload's variant_id already encodes monthly-vs-yearly: each
  # paid plan has two distinct configured variant IDs (see the
  # lemonsqueezy_*_monthly/yearly_variant_id attrs above), one per billing
  # cycle. No separate LemonSqueezy API call or payload field is needed.
  def billing_cycle_for_variant_id(variant_id)
    variant_id = variant_id.to_s
    case variant_id
    when lemonsqueezy_developer_monthly_variant_id, lemonsqueezy_business_monthly_variant_id,
         lemonsqueezy_professional_monthly_variant_id
      "monthly"
    when lemonsqueezy_developer_yearly_variant_id, lemonsqueezy_business_yearly_variant_id,
         lemonsqueezy_professional_yearly_variant_id
      "yearly"
    else
      nil
    end
  end

  def checkout_uuid_for(plan:, billing_cycle:)
    case plan.to_s.downcase
    when "developer"
      billing_cycle == "monthly" ? lemonsqueezy_developer_monthly_checkout_uuid : lemonsqueezy_developer_yearly_checkout_uuid
    when "business"
      billing_cycle == "monthly" ? lemonsqueezy_business_monthly_checkout_uuid : lemonsqueezy_business_yearly_checkout_uuid
    when "professional"
      billing_cycle == "monthly" ? lemonsqueezy_professional_monthly_checkout_uuid : lemonsqueezy_professional_yearly_checkout_uuid
    else
      raise InvalidConfigError, "Unknown plan: #{plan}"
    end
  end

  def private_deployment_checkout_uuid_for(tier:, billing_cycle:)
    key = "lemonsqueezy_private_#{tier}_#{billing_cycle}_checkout_uuid"
    value = public_send(key)
    raise InvalidConfigError, "Unknown private deployment tier/cycle: #{tier}/#{billing_cycle}" if value.nil?
    value
  end

  def smtp_configured?
    smtp_address.present? && smtp_username.present? && smtp_password.present?
  end

  private

  def load_config
    @api_management_api_key = require_env("API_MANAGEMENT_API_KEY")

    @lemonsqueezy_store_id = require_env("LEMONSQUEEZY_STORE_ID")
    @lemonsqueezy_store_slug = optional_env("LEMONSQUEEZY_STORE_SLUG", default: "requiems")
    @lemonsqueezy_test_mode = optional_env("LEMONSQUEEZY_TEST_MODE", default: "false") == "true"
    suffix = @lemonsqueezy_test_mode ? "_TEST" : ""

    @lemonsqueezy_signing_secret = require_env("LEMONSQUEEZY_SIGNING_SECRET#{suffix}")

    @lemonsqueezy_developer_monthly_variant_id = require_env("LEMONSQUEEZY_DEVELOPER_MONTHLY_VARIANT_ID#{suffix}")
    @lemonsqueezy_developer_yearly_variant_id = require_env("LEMONSQUEEZY_DEVELOPER_YEARLY_VARIANT_ID#{suffix}")
    @lemonsqueezy_business_monthly_variant_id = require_env("LEMONSQUEEZY_BUSINESS_MONTHLY_VARIANT_ID#{suffix}")
    @lemonsqueezy_business_yearly_variant_id = require_env("LEMONSQUEEZY_BUSINESS_YEARLY_VARIANT_ID#{suffix}")
    @lemonsqueezy_professional_monthly_variant_id = require_env("LEMONSQUEEZY_PROFESSIONAL_MONTHLY_VARIANT_ID#{suffix}")
    @lemonsqueezy_professional_yearly_variant_id = require_env("LEMONSQUEEZY_PROFESSIONAL_YEARLY_VARIANT_ID#{suffix}")

    @lemonsqueezy_developer_monthly_checkout_uuid = require_env("LEMONSQUEEZY_DEVELOPER_MONTHLY_CHECKOUT_UUID#{suffix}")
    @lemonsqueezy_developer_yearly_checkout_uuid = require_env("LEMONSQUEEZY_DEVELOPER_YEARLY_CHECKOUT_UUID#{suffix}")
    @lemonsqueezy_business_monthly_checkout_uuid = require_env("LEMONSQUEEZY_BUSINESS_MONTHLY_CHECKOUT_UUID#{suffix}")
    @lemonsqueezy_business_yearly_checkout_uuid = require_env("LEMONSQUEEZY_BUSINESS_YEARLY_CHECKOUT_UUID#{suffix}")
    @lemonsqueezy_professional_monthly_checkout_uuid = require_env("LEMONSQUEEZY_PROFESSIONAL_MONTHLY_CHECKOUT_UUID#{suffix}")
    @lemonsqueezy_professional_yearly_checkout_uuid = require_env("LEMONSQUEEZY_PROFESSIONAL_YEARLY_CHECKOUT_UUID#{suffix}")

    # Private deployment tier checkout UUIDs
    @lemonsqueezy_private_starter_monthly_checkout_uuid    = optional_env("LEMONSQUEEZY_PRIVATE_STARTER_MONTHLY_CHECKOUT_UUID#{suffix}")
    @lemonsqueezy_private_starter_yearly_checkout_uuid     = optional_env("LEMONSQUEEZY_PRIVATE_STARTER_YEARLY_CHECKOUT_UUID#{suffix}")
    @lemonsqueezy_private_growth_monthly_checkout_uuid     = optional_env("LEMONSQUEEZY_PRIVATE_GROWTH_MONTHLY_CHECKOUT_UUID#{suffix}")
    @lemonsqueezy_private_growth_yearly_checkout_uuid      = optional_env("LEMONSQUEEZY_PRIVATE_GROWTH_YEARLY_CHECKOUT_UUID#{suffix}")
    @lemonsqueezy_private_scale_monthly_checkout_uuid      = optional_env("LEMONSQUEEZY_PRIVATE_SCALE_MONTHLY_CHECKOUT_UUID#{suffix}")
    @lemonsqueezy_private_scale_yearly_checkout_uuid       = optional_env("LEMONSQUEEZY_PRIVATE_SCALE_YEARLY_CHECKOUT_UUID#{suffix}")
    @lemonsqueezy_private_enterprise_monthly_checkout_uuid = optional_env("LEMONSQUEEZY_PRIVATE_ENTERPRISE_MONTHLY_CHECKOUT_UUID#{suffix}")
    @lemonsqueezy_private_enterprise_yearly_checkout_uuid  = optional_env("LEMONSQUEEZY_PRIVATE_ENTERPRISE_YEARLY_CHECKOUT_UUID#{suffix}")

    @api_base_url = optional_env("API_BASE_URL", default: "https://api.requiems.xyz")
    @playground_api_key = optional_env("PLAYGROUND_API_KEY", default: "rq_test_playground_demo_key")
    @internal_api_url = optional_env("INTERNAL_API_URL", default: "http://localhost:8080")
    @backend_secret = optional_env("BACKEND_SECRET", default: "dev_backend_secret")

    @smtp_address = optional_env("SMTP_ADDRESS")
    @smtp_port = optional_env("SMTP_PORT", default: "587").to_i
    @smtp_domain = optional_env("SMTP_DOMAIN")
    @smtp_username = optional_env("SMTP_USERNAME")
    @smtp_password = optional_env("SMTP_PASSWORD")
    @mailer_host = optional_env("MAILER_HOST", default: "requiems.xyz")
  end

  def validate_config
    validate_url(@api_base_url, "API_BASE_URL")
    validate_url(@internal_api_url, "INTERNAL_API_URL")

    unless @lemonsqueezy_store_id.match?(/^\d+$/)
      raise InvalidConfigError, I18n.t("app.lib.app_config.lemonsqueezy_store_id_must_be_numeric")
    end

    validate_variant_id(@lemonsqueezy_developer_monthly_variant_id, "LEMONSQUEEZY_DEVELOPER_MONTHLY_VARIANT_ID")
    validate_variant_id(@lemonsqueezy_developer_yearly_variant_id, "LEMONSQUEEZY_DEVELOPER_YEARLY_VARIANT_ID")
    validate_variant_id(@lemonsqueezy_business_monthly_variant_id, "LEMONSQUEEZY_BUSINESS_MONTHLY_VARIANT_ID")
    validate_variant_id(@lemonsqueezy_business_yearly_variant_id, "LEMONSQUEEZY_BUSINESS_YEARLY_VARIANT_ID")
    validate_variant_id(@lemonsqueezy_professional_monthly_variant_id, "LEMONSQUEEZY_PROFESSIONAL_MONTHLY_VARIANT_ID")
    validate_variant_id(@lemonsqueezy_professional_yearly_variant_id, "LEMONSQUEEZY_PROFESSIONAL_YEARLY_VARIANT_ID")

    if Rails.env.production? && !smtp_configured?
      Rails.logger.warn("SMTP not fully configured in production environment")
    end
  end

  def require_env(key)
    value = ENV.fetch(key, nil)
    return value if value.present?
    if Rails.env.test?
      test_defaults[key]
    elsif ENV["SECRET_KEY_BASE_DUMMY"].present?
      "dummy"
    else
      raise MissingConfigError, "Missing required environment variable: #{key}"
    end
  end

  def optional_env(key, default: nil)
    value = ENV.fetch(key, nil)
    return value unless value.nil?
    Rails.env.test? ? (test_defaults[key] || default) : default
  end

  def test_defaults
    {
      "API_MANAGEMENT_API_KEY" => "test_api_management_key",
      "LEMONSQUEEZY_STORE_ID" => "12345",
      "LEMONSQUEEZY_SIGNING_SECRET" => "test_signing_secret",
      "LEMONSQUEEZY_DEVELOPER_MONTHLY_VARIANT_ID" => "123456",
      "LEMONSQUEEZY_DEVELOPER_YEARLY_VARIANT_ID" => "123457",
      "LEMONSQUEEZY_BUSINESS_MONTHLY_VARIANT_ID" => "123458",
      "LEMONSQUEEZY_BUSINESS_YEARLY_VARIANT_ID" => "123459",
      "LEMONSQUEEZY_PROFESSIONAL_MONTHLY_VARIANT_ID" => "123460",
      "LEMONSQUEEZY_PROFESSIONAL_YEARLY_VARIANT_ID" => "123461",
      "LEMONSQUEEZY_DEVELOPER_MONTHLY_CHECKOUT_UUID" => "00000000-0000-0000-0000-000000000001",
      "LEMONSQUEEZY_DEVELOPER_YEARLY_CHECKOUT_UUID" => "00000000-0000-0000-0000-000000000002",
      "LEMONSQUEEZY_BUSINESS_MONTHLY_CHECKOUT_UUID" => "00000000-0000-0000-0000-000000000003",
      "LEMONSQUEEZY_BUSINESS_YEARLY_CHECKOUT_UUID" => "00000000-0000-0000-0000-000000000004",
      "LEMONSQUEEZY_PROFESSIONAL_MONTHLY_CHECKOUT_UUID" => "00000000-0000-0000-0000-000000000005",
      "LEMONSQUEEZY_PROFESSIONAL_YEARLY_CHECKOUT_UUID" => "00000000-0000-0000-0000-000000000006",
      "LEMONSQUEEZY_PRIVATE_STARTER_MONTHLY_CHECKOUT_UUID"         => "00000000-0000-0000-0000-000000000021",
      "LEMONSQUEEZY_PRIVATE_STARTER_YEARLY_CHECKOUT_UUID"          => "00000000-0000-0000-0000-000000000022",
      "LEMONSQUEEZY_PRIVATE_GROWTH_MONTHLY_CHECKOUT_UUID"          => "00000000-0000-0000-0000-000000000023",
      "LEMONSQUEEZY_PRIVATE_GROWTH_YEARLY_CHECKOUT_UUID"           => "00000000-0000-0000-0000-000000000024",
      "LEMONSQUEEZY_PRIVATE_SCALE_MONTHLY_CHECKOUT_UUID"           => "00000000-0000-0000-0000-000000000025",
      "LEMONSQUEEZY_PRIVATE_SCALE_YEARLY_CHECKOUT_UUID"            => "00000000-0000-0000-0000-000000000026",
      "LEMONSQUEEZY_PRIVATE_ENTERPRISE_MONTHLY_CHECKOUT_UUID"      => "00000000-0000-0000-0000-000000000027",
      "LEMONSQUEEZY_PRIVATE_ENTERPRISE_YEARLY_CHECKOUT_UUID"       => "00000000-0000-0000-0000-000000000028",
      "LEMONSQUEEZY_PRIVATE_STARTER_MONTHLY_CHECKOUT_UUID_TEST"    => "00000000-0000-0000-0000-000000000021",
      "LEMONSQUEEZY_PRIVATE_STARTER_YEARLY_CHECKOUT_UUID_TEST"     => "00000000-0000-0000-0000-000000000022",
      "LEMONSQUEEZY_PRIVATE_GROWTH_MONTHLY_CHECKOUT_UUID_TEST"     => "00000000-0000-0000-0000-000000000023",
      "LEMONSQUEEZY_PRIVATE_GROWTH_YEARLY_CHECKOUT_UUID_TEST"      => "00000000-0000-0000-0000-000000000024",
      "LEMONSQUEEZY_PRIVATE_SCALE_MONTHLY_CHECKOUT_UUID_TEST"      => "00000000-0000-0000-0000-000000000025",
      "LEMONSQUEEZY_PRIVATE_SCALE_YEARLY_CHECKOUT_UUID_TEST"       => "00000000-0000-0000-0000-000000000026",
      "LEMONSQUEEZY_PRIVATE_ENTERPRISE_MONTHLY_CHECKOUT_UUID_TEST" => "00000000-0000-0000-0000-000000000027",
      "LEMONSQUEEZY_PRIVATE_ENTERPRISE_YEARLY_CHECKOUT_UUID_TEST"  => "00000000-0000-0000-0000-000000000028",
      "BACKEND_SECRET" => "test_backend_secret",
      "LEMONSQUEEZY_SIGNING_SECRET_TEST" => "test_signing_secret_test",
      "LEMONSQUEEZY_DEVELOPER_MONTHLY_VARIANT_ID_TEST" => "223456",
      "LEMONSQUEEZY_DEVELOPER_YEARLY_VARIANT_ID_TEST" => "223457",
      "LEMONSQUEEZY_BUSINESS_MONTHLY_VARIANT_ID_TEST" => "223458",
      "LEMONSQUEEZY_BUSINESS_YEARLY_VARIANT_ID_TEST" => "223459",
      "LEMONSQUEEZY_PROFESSIONAL_MONTHLY_VARIANT_ID_TEST" => "223460",
      "LEMONSQUEEZY_PROFESSIONAL_YEARLY_VARIANT_ID_TEST" => "223461",
      "LEMONSQUEEZY_DEVELOPER_MONTHLY_CHECKOUT_UUID_TEST" => "00000000-0000-0000-0000-000000000011",
      "LEMONSQUEEZY_DEVELOPER_YEARLY_CHECKOUT_UUID_TEST" => "00000000-0000-0000-0000-000000000012",
      "LEMONSQUEEZY_BUSINESS_MONTHLY_CHECKOUT_UUID_TEST" => "00000000-0000-0000-0000-000000000013",
      "LEMONSQUEEZY_BUSINESS_YEARLY_CHECKOUT_UUID_TEST" => "00000000-0000-0000-0000-000000000014",
      "LEMONSQUEEZY_PROFESSIONAL_MONTHLY_CHECKOUT_UUID_TEST" => "00000000-0000-0000-0000-000000000015",
      "LEMONSQUEEZY_PROFESSIONAL_YEARLY_CHECKOUT_UUID_TEST" => "00000000-0000-0000-0000-000000000016"
    }
  end

  def validate_url(url, key)
    uri = URI.parse(url)

    unless uri.is_a?(URI::HTTP) || uri.is_a?(URI::HTTPS)
      raise InvalidConfigError, "#{key} must be a valid HTTP/HTTPS URL"
    end

  rescue URI::InvalidURIError
    raise InvalidConfigError, "#{key} is not a valid URL: #{url}"
  end

  def validate_variant_id(value, key)
    unless value.present? && value.match?(/^\w+$/)
      raise InvalidConfigError, "#{key} must be present and contain valid characters"
    end
  end
end
