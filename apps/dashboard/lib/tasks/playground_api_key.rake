# frozen_string_literal: true

namespace :playground do
  desc "Provision (or update) the dedicated system API key backing the public playground and demo forms"
  task provision_key: :environment do
    raw_key = ENV["PLAYGROUND_API_KEY"]
    raise "PLAYGROUND_API_KEY must be set to the raw key value before running this task" if raw_key.blank?
    unless raw_key.match?(/\Arequiem_[0-9a-zA-Z]{24}\z/)
      raise "PLAYGROUND_API_KEY must match the requiem_<24-char-alnum> format ApiKeyGenerator produces"
    end

    # No Cloudflare/Worker round trip: this account never authenticates via
    # Devise, only holds a Subscription + ApiKey for Go's auth path.
    user = User.find_or_create_by!(email: "playground@internal.requiems.xyz") do |u|
      u.password = SecureRandom.hex(32)
      u.name = "Playground (internal)"
      u.confirmed_at = Time.current
    end

    # "developer" tier: enough headroom that playground traffic doesn't get
    # rate-limited against a real customer's budget, while staying bounded —
    # Phase 2's quota/rate-limit middleware still catches a demo-abuse spike,
    # unlike an unlimited enterprise-null-limits plan would.
    #
    # This does mean the account shows up in AnalyticsRevenueService's
    # paying/MRR figures like a real $30/mo "developer" subscriber, the same
    # gap admin-promoted (promoted_by_id) subscriptions already have today —
    # not a new problem introduced here, and out of this task's scope to fix.
    subscription = user.subscription || user.build_subscription
    subscription.update!(plan_name: "developer", status: "active", cancel_at_period_end: false)

    key_prefix = ApiKeyGenerator.extract_prefix(raw_key)
    key_hash = ApiKeyGenerator.hash_key(raw_key)

    # Pre-setting key_prefix means ApiKey#generate_key's before_validation
    # early-returns (key_prefix.present?), which also skips its "active
    # defaults to true" line — set active/revoked_at explicitly here rather
    # than relying on that callback.
    api_key = user.api_keys.first_or_initialize
    already_current = api_key.persisted? && api_key.key_prefix == key_prefix && api_key.active? && api_key.revoked_at.nil?

    api_key.assign_attributes(
      name: "Playground (system)",
      key_prefix: key_prefix,
      key_hash: key_hash,
      full_key: raw_key,
      active: true,
      revoked_at: nil
    )
    api_key.save!

    puts already_current ? "Playground API key already matches PLAYGROUND_API_KEY (prefix: #{key_prefix})" :
                            "Provisioned playground API key (prefix: #{key_prefix})"
  end
end
