# frozen_string_literal: true

require "test_helper"

class GoAuthCacheTest < ActiveSupport::TestCase
  def setup
    @redis = Redis.new(url: ENV.fetch("REDIS_URL", "redis://localhost:6379"))
    @redis.ping
  rescue StandardError => e
    skip "Redis unavailable for GoAuthCache tests: #{e.message}"
  end

  test "deletes the raw, unnamespaced key Go reads/writes" do
    prefix = "requiem_test#{SecureRandom.hex(4)}"
    @redis.set("#{GoAuthCache::CACHE_KEY_PREFIX}#{prefix}", '{"user_id":1,"plan":"free","revoked":false}')

    GoAuthCache.invalidate(prefix)

    assert_nil @redis.get("#{GoAuthCache::CACHE_KEY_PREFIX}#{prefix}")
  end

  test "does not touch the Rails.cache-namespaced key" do
    # Regression guard for the exact bug this class exists to avoid: if
    # invalidation ever regresses to Rails.cache.delete(key_prefix), it would
    # delete "rails_cache:{prefix}" and silently leave the plain
    # "apikey:{prefix}" key Go reads untouched.
    prefix = "requiem_test#{SecureRandom.hex(4)}"
    namespaced_key = "rails_cache:#{prefix}"
    @redis.set(namespaced_key, "unrelated")

    GoAuthCache.invalidate(prefix)

    assert_equal "unrelated", @redis.get(namespaced_key)
  ensure
    @redis&.del(namespaced_key)
  end
end
