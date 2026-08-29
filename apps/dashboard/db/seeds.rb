# frozen_string_literal: true

# This file should ensure the existence of records required to run the application in every environment (production,
# development, test). The code here should be idempotent so that it can be executed at any point in every environment.
# The data can then be loaded with the bin/rails db:seed command (or created alongside the database with db:setup).

# Development: Create test user if it doesn't exist
if Rails.env.development?
  test_user = User.find_or_initialize_by(email: "eliaz.bobadilladeva@gmail.com")
  if test_user.new_record?
    test_user.password = SecureRandom.hex(16)
    test_user.admin = true
    test_user.save!
    puts "✓ Created test user: #{test_user.email} (admin: true)"
  else
    puts "✓ Test user already exists: #{test_user.email}"
  end

  # Create additional test user
  user2 = User.find_or_initialize_by(email: "test@example.com")
  if user2.new_record?
    user2.password = "password123!"
    user2.save!
    puts "✓ Created additional test user: #{user2.email}"
  end

  # Dev API key for exercising the Go auth path directly
  # (apps/api/platform/middleware/apikeyauth.go). The configured raw value is
  # reconciled on every development seed so an existing Docker Postgres
  # volume cannot leave the stack using an old, unrecoverable key. The raw key
  # is accepted only from an operator-provided environment variable and is
  # never printed by the seed process.
  raw_key = ENV["LOCAL_DEV_API_KEY"]
  raise "LOCAL_DEV_API_KEY must be set for the local development API key" if raw_key.blank?
  unless raw_key.match?(/\Arequiem_[0-9A-Za-z]{24}\z/)
    raise "LOCAL_DEV_API_KEY must match requiem_<24 alphanumeric characters>"
  end

  key_prefix = ApiKeyGenerator.extract_prefix(raw_key)
  dev_key = test_user.api_keys.find_by(name: "Local Dev Key")

  if dev_key&.active? && dev_key.revoked_at.nil? && dev_key.verify_key(raw_key)
    puts "✓ Local development API key already matches LOCAL_DEV_API_KEY " \
         "(key_prefix: #{dev_key.key_prefix})"
  else
    prefix_owner = ApiKey.find_by(key_prefix: key_prefix)
    if prefix_owner && prefix_owner != dev_key
      raise "LOCAL_DEV_API_KEY prefix is already used by another API key"
    end

    if dev_key
      dev_key.update!(
        key_prefix: key_prefix,
        key_hash: ApiKeyGenerator.hash_key(raw_key),
        full_key: raw_key,
        active: true,
        revoked_at: nil,
        revoked_reason: nil
      )
      puts "✓ Updated local development API key for #{test_user.email} " \
           "(key_prefix: #{dev_key.key_prefix})"
    else
      dev_key = test_user.api_keys.create!(
        name: "Local Dev Key",
        key_prefix: key_prefix,
        key_hash: ApiKeyGenerator.hash_key(raw_key),
        full_key: raw_key,
        active: true,
        revoked_at: nil
      )

      puts "✓ Created API key for #{test_user.email} (key_prefix: #{dev_key.key_prefix})"
    end
  end
end
