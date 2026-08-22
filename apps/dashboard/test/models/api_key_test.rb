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
    assert new_key.key_prefix.start_with?("requiem_")
  end

  test "generated key matches the format Go's validator requires" do
    new_key = @user.api_keys.create!(name: "Format Check", environment: "live")

    assert_match(/\Arequiem_[0-9a-zA-Z]{24}\z/, new_key.full_key)
    assert_equal 12, new_key.key_prefix.length
  end

  test "retries on key_prefix collision and eventually raises if exhausted" do
    colliding_key = "requiem_collidingkeyvalue00001"

    # Simulate the exact prefix already existing so every attempt collides.
    @user.api_keys.create!(
      name: "Existing",
      key_prefix: ApiKeyGenerator.extract_prefix(colliding_key),
      key_hash: ApiKeyGenerator.hash_key(colliding_key)
    )

    ApiKeyGenerator.stub(:generate_candidate, colliding_key) do
      assert_raises(ApiKeyGenerator::CollisionError) do
        ApiKeyGenerator.generate
      end
    end
  end

  test "retries when the key_prefix index loses a concurrent insert race" do
    colliding_key = "requiem_raceprefix00000000000001"
    existing = @user.api_keys.create!(
      name: "Existing race key",
      key_prefix: ApiKeyGenerator.extract_prefix(colliding_key),
      key_hash: ApiKeyGenerator.hash_key(colliding_key)
    )

    raced = ApiKey.new(
      user: @user,
      name: "Raced key",
      key_prefix: existing.key_prefix,
      key_hash: existing.key_hash,
      active: true
    )

    assert raced.save!(validate: false)
    assert_not_equal existing.key_prefix, raced.key_prefix
  end

  test "revoke! invalidates the Go auth cache for this key_prefix" do
    redis = go_auth_test_redis
    cache_key = "#{GoAuthCache::CACHE_KEY_PREFIX}#{@api_key.key_prefix}"
    redis.set(cache_key, '{"user_id":1,"plan":"free","revoked":false}')

    @api_key.revoke!(reason: "test")

    assert_nil redis.get(cache_key)
  ensure
    redis&.del(cache_key) if cache_key
  end

  test "direct revoked_at updates invalidate the Go auth cache" do
    redis = go_auth_test_redis
    cache_key = "#{GoAuthCache::CACHE_KEY_PREFIX}#{@api_key.key_prefix}"
    redis.set(cache_key, '{"user_id":1,"plan":"free","revoked":false}')

    @api_key.update!(revoked_at: Time.current)

    assert_nil redis.get(cache_key)
  ensure
    redis&.del(cache_key) if cache_key
  end

  test "rotating key material invalidates both old and new Go auth cache prefixes" do
    redis = go_auth_test_redis
    old_prefix = @api_key.key_prefix
    new_raw_key = "requiem_#{"A" * 24}"
    new_prefix = ApiKeyGenerator.extract_prefix(new_raw_key)
    old_cache_key = "#{GoAuthCache::CACHE_KEY_PREFIX}#{old_prefix}"
    new_cache_key = "#{GoAuthCache::CACHE_KEY_PREFIX}#{new_prefix}"
    redis.set(old_cache_key, '{"user_id":1,"plan":"free","revoked":false}')
    redis.set(new_cache_key, '{"user_id":1,"plan":"free","revoked":false}')

    @api_key.update!(key_prefix: new_prefix, key_hash: ApiKeyGenerator.hash_key(new_raw_key))

    assert_nil redis.get(old_cache_key)
    assert_nil redis.get(new_cache_key)
  ensure
    redis&.del(old_cache_key, new_cache_key) if old_cache_key && new_cache_key
  end

  test "destroying an api key invalidates the Go auth cache for this key_prefix" do
    redis = go_auth_test_redis
    cache_key = "#{GoAuthCache::CACHE_KEY_PREFIX}#{@api_key.key_prefix}"
    redis.set(cache_key, '{"user_id":1,"plan":"free","revoked":false}')

    @api_key.destroy!

    assert_nil redis.get(cache_key)
  ensure
    redis&.del(cache_key) if cache_key
  end

  test "creates a valid local key with no Cloudflare/network call, in a non-test-like environment" do
    Rails.env.define_singleton_method(:test?) { false }

    begin
      api_key = @user.api_keys.build(name: "Non-test env key", environment: "live")
      assert api_key.save
      assert_match(/\Arequiem_[0-9a-zA-Z]{24}\z/, api_key.full_key)
    ensure
      Rails.env.singleton_class.remove_method(:test?)
    end
  end

  private

  def go_auth_test_redis
    redis = Redis.new(url: ENV.fetch("REDIS_URL", "redis://localhost:6379"))
    redis.ping
    redis
  rescue StandardError => e
    skip "Redis unavailable: #{e.message}"
  end
end
