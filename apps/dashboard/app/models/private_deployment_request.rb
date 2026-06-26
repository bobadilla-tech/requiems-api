# frozen_string_literal: true

class PrivateDeploymentRequest < ApplicationRecord
  belongs_to :user

  encrypts :tenant_secret

  VALID_STATUSES = %w[pending_payment pending deploying active cancelled].freeze
  VALID_TIERS = %w[starter growth scale enterprise].freeze
  VALID_SERVICES = %w[email text tech places finance entertainment ai convert fitness misc].freeze

  SERVICE_META = {
    "email"         => "Validation, disposable check, normalization",
    "text"          => "Quotes, dictionary, lorem ipsum, advice, profanity, spellcheck",
    "tech"          => "IP/VPN info, phone, QR codes, domain, WHOIS, MX lookup",
    "places"        => "Cities, geocoding, timezone, holidays, postal codes",
    "finance"       => "BIN lookup, crypto, exchange rates, IBAN, SWIFT, mortgage",
    "entertainment" => "Jokes, facts, horoscope, trivia, emoji, sudoku",
    "ai"            => "Similarity, sentiment analysis, language detection",
    "convert"       => "Base64, color, markdown, number base, format conversion",
    "fitness"       => "Exercise database by body part, equipment, muscle",
    "misc"          => "Counter, random user, unit conversion"
  }.freeze

  TIER_PRICES_MONTHLY = {
    "starter"    => 20000,
    "growth"     => 30000,
    "scale"      => 50000,
    "enterprise" => 96000   # $960/mo — $1,000 was too high for semi-annual LemonSqueezy limit
  }.freeze

  # LemonSqueezy per-charge amount for the "yearly" plan (~15% off).
  # Starter + Growth: billed once per year.
  # Scale + Enterprise: billed every 6 months (semi-annual) to stay under LemonSqueezy's $5,000 cap.
  TIER_PRICES_YEARLY = {
    "starter"    => 204000,  # $2,040  — annual charge    ($170/mo effective)
    "growth"     => 306000,  # $3,060  — annual charge    ($255/mo effective)
    "scale"      => 255000,  # $2,550  — semi-annual charge ($425/mo effective)
    "enterprise" => 489600   # $4,896  — semi-annual charge ($816/mo effective)
  }.freeze

  # How many months each "yearly" LemonSqueezy charge covers.
  TIER_YEARLY_BILLING_MONTHS = {
    "starter"    => 12,
    "growth"     => 12,
    "scale"      => 6,
    "enterprise" => 6
  }.freeze

  TIER_SPECS = {
    "starter"    => { hetzner: "CPX21", vcpu: 3, ram: "4 GB",  ssd: "80 GB"  },
    "growth"     => { hetzner: "CPX31", vcpu: 4, ram: "8 GB",  ssd: "160 GB" },
    "scale"      => { hetzner: "CPX41", vcpu: 8, ram: "16 GB", ssd: "240 GB" },
    "enterprise" => { hetzner: "CPX51", vcpu: 16, ram: "32 GB", ssd: "360 GB" }
  }.freeze

  validates :contact_name, :contact_email, :server_tier, :billing_cycle, presence: true
  validates :server_tier, inclusion: { in: VALID_TIERS }
  validates :billing_cycle, inclusion: { in: %w[monthly yearly] }
  validates :status, inclusion: { in: VALID_STATUSES }
  validates :lemonsqueezy_subscription_id, uniqueness: true, allow_nil: true
  validates :tenant_secret, length: { minimum: 32 }, if: -> { status == "active" }
  validate :at_least_one_service_selected
  validates :subdomain_slug, uniqueness: true, allow_nil: true,
            format: { with: /\A[a-z0-9\-]{2,40}\z/, message: "must be lowercase letters, numbers, and hyphens only (2–40 chars)" }
  validates :contact_email, format: { with: URI::MailTo::EMAIL_REGEXP }

  scope :pending,    -> { where(status: "pending") }
  scope :deploying,  -> { where(status: "deploying") }
  scope :active,     -> { where(status: "active") }
  scope :cancelled,  -> { where(status: "cancelled") }

  def monthly_price_dollars
    cents = monthly_price_cents || TIER_PRICES_MONTHLY.fetch(server_tier, 0)
    cents / 100.0
  end

  def total_price_dollars
    price_table = billing_cycle == "yearly" ? TIER_PRICES_YEARLY : TIER_PRICES_MONTHLY
    price_table.fetch(server_tier, 0) / 100.0
  end

  def live_url
    "https://#{subdomain_slug}.requiems.xyz" if subdomain_slug.present?
  end

  def tier_specs
    TIER_SPECS.fetch(server_tier, {})
  end

  private

  def at_least_one_service_selected
    if selected_services.blank? || selected_services.empty?
      errors.add(:selected_services, I18n.t("private_deployment_request.must_include_at_least_one_service"))
      return
    end

    invalid = selected_services.reject { |s| VALID_SERVICES.include?(s) }
    errors.add(:selected_services, "contains invalid services: #{invalid.join(', ')}") if invalid.any?
  end
end
