# frozen_string_literal: true

require "test_helper"

class UserTest < ActiveSupport::TestCase
  def setup
    @user = create_user(
      email: "test@example.com",
      password: "password123",
      password_confirmation: "password123"
    )
  end

  test "valid user with email and password" do
    assert @user.valid?
    assert @user.persisted?
  end

  test "requires email" do
    user = User.new(password: "password123")
    assert_not user.valid?
    assert_includes user.errors[:email], "can't be blank"
  end

  test "requires valid email format" do
    user = User.new(email: "invalid", password: "password123")
    assert_not user.valid?
  end

  test "requires unique email" do
    duplicate_user = User.new(
      email: @user.email,
      password: "password123"
    )
    assert_not duplicate_user.valid?
    assert_includes duplicate_user.errors[:email], "has already been taken"
  end

  test "admin? returns admin status" do
    assert_not @user.admin?

    @user.update(admin: true)
    assert @user.admin?
  end

  test "suspended? returns suspension status" do
    assert_not @user.suspended?

    @user.update(status: "suspended")
    assert @user.suspended?
  end

  test "banned? returns ban status" do
    assert_not @user.banned?

    @user.update(status: "banned", banned_at: Time.current)
    assert @user.banned?
  end

  test "ban! invalidates the Go auth cache for all active api keys, even though update_all skips ApiKey callbacks" do
    redis = Redis.new(url: ENV.fetch("REDIS_URL", "redis://localhost:6379"))
    redis.ping
  rescue StandardError => e
    skip "Redis unavailable: #{e.message}"
  else
    admin = create_user(email: "admin@example.com", password: "password123", password_confirmation: "password123")
    key_one = @user.api_keys.create!(name: "One", environment: "test")
    key_two = @user.api_keys.create!(name: "Two", environment: "test")

    cache_key_one = "#{GoAuthCache::CACHE_KEY_PREFIX}#{key_one.key_prefix}"
    cache_key_two = "#{GoAuthCache::CACHE_KEY_PREFIX}#{key_two.key_prefix}"
    redis.set(cache_key_one, "warm")
    redis.set(cache_key_two, "warm")

    @user.ban!(reason: "test", admin_user: admin)

    assert_nil redis.get(cache_key_one)
    assert_nil redis.get(cache_key_two)
  ensure
    redis&.del(cache_key_one, cache_key_two) if cache_key_one
  end

  test "has many api_keys" do
    assert_respond_to @user, :api_keys

    api_key = @user.api_keys.create!(
      name: "Test Key",
      environment: "test"
    )

    assert_includes @user.api_keys, api_key
  end

  test "has one subscription" do
    assert_respond_to @user, :subscription
  end

  test "has many usage_logs" do
    assert_respond_to @user, :usage_logs
  end

  test "destroys dependent api_keys when destroyed" do
    api_key = @user.api_keys.create!(
      name: "Test Key",
      environment: "test"
    )

    assert_difference "ApiKey.count", -1 do
      @user.destroy
    end
  end

  test "PLAN_LIMITS matches workers-shared/plan-limits.json" do
    json_path = Rails.root.join("../workers/shared/plan-limits.json")
    shared_limits = JSON.parse(File.read(json_path))

    shared_limits.each do |plan, limit|
      assert_equal limit, User::PLAN_LIMITS[plan],
        "User::PLAN_LIMITS[#{plan.inspect}] (#{User::PLAN_LIMITS[plan]}) " \
        "does not match workers-shared/plan-limits.json (#{limit}). " \
        "Update both files together to keep them in sync."
    end

    assert_equal shared_limits.keys.sort, User::PLAN_LIMITS.keys.sort,
      "User::PLAN_LIMITS has different plan keys than workers-shared/plan-limits.json"
  end
end
