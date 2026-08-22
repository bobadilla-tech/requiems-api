# frozen_string_literal: true

require "bcrypt"

class ApiKeyGenerator
  # Matches apps/api/platform/middleware/apikeyauth.go's keyPrefixLength (12)
  # and the validator both Go and the (retired) Worker enforce:
  # ^requiem_[0-9a-zA-Z]{24}$
  ALPHABET = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

  MAX_GENERATION_ATTEMPTS = 5

  class CollisionError < StandardError; end

  # Generate a new, prefix-unique API key in the format: requiem_<24_random_chars>
  # Returns the full key (which should be shown to the user once). Retries on
  # a key_prefix collision (extremely unlikely with nanoid, but the same
  # good-practice check the Cloudflare-backed path used to make) up to
  # MAX_GENERATION_ATTEMPTS before raising.
  def self.generate
    MAX_GENERATION_ATTEMPTS.times do
      candidate = generate_candidate
      return candidate unless ApiKey.exists?(key_prefix: extract_prefix(candidate))
    end

    raise CollisionError, "Could not generate a unique API key after #{MAX_GENERATION_ATTEMPTS} attempts"
  end

  # Extract the prefix for display (first 12 characters)
  def self.extract_prefix(full_key)
    full_key[0..11] if full_key
  end

  # Hash the full key for secure storage
  def self.hash_key(full_key)
    BCrypt::Password.create(full_key)
  end

  # Verify a key against a hash
  def self.verify_key(full_key, key_hash)
    BCrypt::Password.new(key_hash) == full_key
  rescue BCrypt::Errors::InvalidHash
    false
  end

  def self.generate_candidate
    "requiem_#{Nanoid.generate(size: 24, alphabet: ALPHABET)}"
  end
  private_class_method :generate_candidate
end
