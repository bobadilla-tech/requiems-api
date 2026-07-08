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
  before_validation :request_key_from_server, on: :create
  after_destroy :remove_from_cloudflare
  after_update :sync_revocation_to_cloudflare, if: :saved_change_to_active?

  # Request a new API key from the api-management worker.
  # The worker generates the key, stores it in KV + D1, and returns it once.
  def request_key_from_server
    return if key_prefix.present? # Skip if already generated

    if Rails.env.test?
      # In tests, generate locally to avoid external dependencies
      generate_key_locally
    else
      # In all other environments, use the api-management worker so that
      # the key is immediately available in Cloudflare KV (auth) and D1 (usage)
      service = Cloudflare::ApiManagementService.new
      billing_start = user.subscription&.created_at || Time.current

      generated_key = service.create_key(
        user_id: user_id,
        plan: user.current_plan,
        name: name,
        billing_cycle_start: billing_start.iso8601
      )

      if generated_key
        self.full_key = generated_key
        self.key_prefix = ApiKeyGenerator.extract_prefix(generated_key)
        self.key_hash = ApiKeyGenerator.hash_key(generated_key)
        self.active = true if active.nil?
      else
        errors.add(:base, I18n.t("api_key.failed_to_generate_api_key_please_try"))
        throw :abort
      end
    end
  end

  def generate_key_locally
    # Generate a local API key for test/development
    env = environment&.to_sym || :live
    generated_key = ApiKeyGenerator.generate(environment: env)
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

  def remove_from_cloudflare
    return if Rails.env.test?

    Cloudflare::ApiManagementService.new.revoke_key(key_prefix)
  rescue StandardError => e
    Rails.logger.error("[ApiManagement] Failed to revoke API key #{key_prefix}: #{e.message}")
  end

  def sync_revocation_to_cloudflare
    return unless !active && revoked_at

    remove_from_cloudflare
  end
end
