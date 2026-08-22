# frozen_string_literal: true

# Invalidates the Go auth path's (apps/api/platform/middleware/apikeyauth.go)
# Redis verification cache on API key revocation/deletion.
#
# This MUST use a raw, unnamespaced Redis connection, not Rails.cache:
# Rails.cache is a :redis_cache_store configured with namespace: "rails_cache"
# (config/environments/{development,production}.rb, what Rack::Attack also
# rides on), so Rails.cache.delete(key_prefix) would actually target
# "rails_cache:{key_prefix}" in Redis — silently missing the plain
# "apikey:{key_prefix}" key Go's raw go-redis client reads/writes. That would
# be a permanent no-op, not a best-effort fallback.
class GoAuthCache
  CACHE_KEY_PREFIX = "apikey:"

  class << self
    def invalidate(key_prefix)
      redis.del("#{CACHE_KEY_PREFIX}#{key_prefix}")
    rescue StandardError => e
      # If this DEL fails, a stale cache entry could keep serving a revoked
      # key as valid until Go's cache TTL expires. At zero scale this is an
      # accepted, explicitly-documented gap (short TTL is the mitigation),
      # not something to build a durable-retry-queue for yet — see
      # docs/plans/2026-08-21-go-auth-foundation-phase-0-1.md, Phase 1 item 3.
      Rails.logger.error("[GoAuthCache] Failed to invalidate key_prefix #{key_prefix}: #{e.message}")
    end

    private

    def redis
      @redis ||= Redis.new(url: ENV.fetch("REDIS_URL", "redis://localhost:6379"))
    end
  end
end
