# frozen_string_literal: true

class ApiKey < ApplicationRecord
  belongs_to :user
  has_many :usage_logs, dependent: :destroy
  has_many :abuse_reports, dependent: :destroy

  # Virtual attribute to store the full key temporarily (only shown once to user)
  attr_accessor :full_key

  # Validations
  validates :key_prefix, presence: true, uniqueness: true
  validates :key_hash, presence: true
  validates :name, presence: true, length: { maximum: 100 }
  validates :environment, inclusion: { in: %w[test live] }, allow_nil: true

  # Scopes
  scope :active_keys, -> { where(active: true, revoked_at: nil) }
  scope :revoked, -> { where.not(revoked_at: nil) }
  scope :for_environment, ->(env) { where(environment: env) }

  # Callbacks
  before_validation :generate_key, on: :create
  # _commit, not plain after_destroy/after_update: revoke! and destroy can run
  # inside a wrapping transaction (e.g. User#ban!). A plain after_update fires
  # before that outer transaction commits — if Go re-caches the key in that
  # window, it reads the still-committed-old row (Postgres isn't done
  # committing yet) and caches it as valid for a full TTL, outliving the
  # revocation it raced with. after_commit only fires once the row change is
  # actually visible to other connections.
  # Two distinctly-named methods, not the same symbol registered twice with
  # different `if:` guards: Rails' callback chain deduplicates by filter name,
  # so `after_destroy_commit :invalidate_go_auth_cache` followed by
  # `after_update_commit :invalidate_go_auth_cache, if: ...` collapses into a
  # single entry that keeps only the second's `if:` guard — silently
  # no-opping on destroy, since `saved_change_to_active?` is never true there.
  # Verified this actually happens (empirically, via `bin/rails runner`, not
  # just in theory) before landing this as two separate methods.
  after_destroy_commit :invalidate_go_auth_cache_on_destroy
  after_update_commit :invalidate_go_auth_cache_on_revoke, if: :saved_change_to_active?

  # Generate and hash a new API key locally — the sole key-generation path in
  # every environment. No Cloudflare/Worker round trip.
  def generate_key
    return if key_prefix.present? # Skip if already generated

    generated_key = ApiKeyGenerator.generate
    self.full_key = generated_key
    self.key_prefix = ApiKeyGenerator.extract_prefix(generated_key)
    self.key_hash = ApiKeyGenerator.hash_key(generated_key)
    self.active = true if active.nil?
  end

  # Verify a key matches this record
  def verify_key(key_to_verify)
    ApiKeyGenerator.verify_key(key_to_verify, key_hash)
  end

  # Revoke the API key
  def revoke!(reason: nil)
    update!(
      active: false,
      revoked_at: Time.current,
      revoked_reason: reason
    )
  end

  # Display-friendly key (shows prefix + masked)
  def masked_key
    "#{key_prefix}••••••••••••"
  end

  private

  # Invalidates the Go auth path's Redis verification cache. Called both on
  # destroy and on revocation (active flipping to false) — either way, the
  # Go-side cached {user_id, plan, revoked} result for this key_prefix must
  # not keep serving a key that no longer exists or is no longer active.
  def invalidate_go_auth_cache_on_destroy
    GoAuthCache.invalidate(key_prefix)
  end

  def invalidate_go_auth_cache_on_revoke
    GoAuthCache.invalidate(key_prefix)
  end
end
