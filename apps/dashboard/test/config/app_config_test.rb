# frozen_string_literal: true

require "test_helper"

class AppConfigTest < ActiveSupport::TestCase
  test "private deployment checkout uuid lookup returns configured uuid" do
    assert_equal(
      "00000000-0000-0000-0000-000000000021",
      AppConfig.private_deployment_checkout_uuid_for(tier: "starter", billing_cycle: "monthly")
    )
    assert_equal(
      "00000000-0000-0000-0000-000000000028",
      AppConfig.private_deployment_checkout_uuid_for(tier: "enterprise", billing_cycle: "yearly")
    )
  end

  test "validate_config raises when lemonsqueezy store id is non-numeric" do
    config = AppConfig.instance
    original = config.instance_variable_get(:@lemonsqueezy_store_id)

    config.instance_variable_set(:@lemonsqueezy_store_id, "not-a-number")

    error = assert_raises(AppConfig::InvalidConfigError) do
      config.send(:validate_config)
    end
    assert_equal I18n.t("app.lib.app_config.lemonsqueezy_store_id_must_be_numeric"), error.message
  ensure
    config&.instance_variable_set(:@lemonsqueezy_store_id, original)
  end

  test "private deployment checkout uuid lookup raises when config is missing" do
    config = nil
    original = nil
    config = AppConfig.instance
    original = config.instance_variable_get(:@lemonsqueezy_private_starter_monthly_checkout_uuid)

    config.instance_variable_set(:@lemonsqueezy_private_starter_monthly_checkout_uuid, nil)

    assert_raises(AppConfig::InvalidConfigError) do
      config.private_deployment_checkout_uuid_for(tier: "starter", billing_cycle: "monthly")
    end
  ensure
    config&.instance_variable_set(:@lemonsqueezy_private_starter_monthly_checkout_uuid, original)
  end
end
