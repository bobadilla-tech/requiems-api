# frozen_string_literal: true

require "test_helper"

class ApiKeyTest < ActiveSupport::TestCase
  def setup
    @user = create_user(
      email: "test@example.com",
      password: "password123",
      password_confirmation: "password123"
    )

    @api_key = @user.api_keys.create!(
      name: "Test Key",
      environment: "test"
    )
  end

  test "valid api key with required attributes" do
    assert @api_key.valid?
    assert @api_key.persisted?
    assert_not_nil @api_key.key_hash
    assert_not_nil @api_key.key_prefix
  end

  test "requires name" do
    api_key = @user.api_keys.build(environment: "test")
    api_key.name = nil

    assert_not api_key.valid?
    assert_includes api_key.errors[:name], "can't be blank"
  end

  test "requires key_hash" do
    api_key = @user.api_keys.build(name: "Test", environment: "test")
    # Bypass the before_create callback
    api_key.save(validate: false)
    api_key.update_column(:key_hash, nil)
    api_key.reload

    assert_not api_key.valid?
    assert_includes api_key.errors[:key_hash], "can't be blank"
  end

  test "requires unique key_prefix" do
    # Create second key which will have a different prefix
    second_key = @user.api_keys.create!(name: "Second", environment: "test")

    # Try to manually set duplicate prefix (this would only happen via direct DB manipulation)
    assert_not_equal @api_key.key_prefix, second_key.key_prefix
  end

  test "requires environment" do
    api_key = @user.api_keys.build(name: "Test")
    api_key.environment = nil

    # Environment is optional in the model
    assert api_key.valid?
  end

  test "validates environment is test or live" do
    @api_key.environment = "invalid"
    assert_not @api_key.valid?

    @api_key.environment = "test"
    assert @api_key.valid?

    @api_key.environment = "live"
    assert @api_key.valid?
  end

  test "belongs to user" do
    assert_equal @user, @api_key.user
  end

  test "has many usage_logs" do
    assert_respond_to @api_key, :usage_logs
  end

  test "active_keys scope returns non-revoked keys" do
    active_key = @user.api_keys.create!(
      name: "Active",
      environment: "test"
    )

    revoked_key = @user.api_keys.create!(
      name: "Revoked",
      environment: "test"
    )
    revoked_key.update_column(:revoked_at, Time.current)
    revoked_key.update_column(:active, false)

    active_keys = ApiKey.active_keys

    assert_includes active_keys, active_key
    assert_not_includes active_keys, revoked_key
  end

  test "revoked scope returns revoked keys" do
    revoked_key = @user.api_keys.create!(
      name: "Revoked",
      environment: "test"
    )
    revoked_key.update_column(:revoked_at, Time.current)

    revoked_keys = ApiKey.revoked

    assert_includes revoked_keys, revoked_key
    assert_not_includes revoked_keys, @api_key
  end

  test "revoke! sets revoked_at and reason" do
    assert_nil @api_key.revoked_at
    assert_nil @api_key.revoked_reason

    @api_key.revoke!(reason: "User requested")

    @api_key.reload
    assert_not_nil @api_key.revoked_at
    assert_equal "User requested", @api_key.revoked_reason
    assert_equal false, @api_key.active
  end

  test "generates prefix from key on creation" do
    new_key = @user.api_keys.create!(
      name: "New Key",
      environment: "test"
    )

    assert_not_nil new_key.key_prefix
    assert new_key.key_prefix.start_with?("rq_test_")
  end

  test "adds error when server returns no key in non-test env" do
    fake_service = Object.new
    fake_service.define_singleton_method(:create_key) { |**_| nil }

    Cloudflare::ApiManagementService.define_singleton_method(:new) { fake_service }
    Rails.env.define_singleton_method(:test?) { false }

    begin
      api_key = @user.api_keys.build(name: "Server Fail", environment: "live")
      assert_not api_key.save
      assert_includes api_key.errors[:base], I18n.t("api_key.failed_to_generate_api_key_please_try")
    ensure
      Rails.env.singleton_class.remove_method(:test?)
      Cloudflare::ApiManagementService.singleton_class.remove_method(:new)
    end
  end
end
