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
  # (apps/api/platform/middleware/apikeyauth.go). Key generation is local-only
  # in every environment (ApiKey#generate_key). The raw key is accepted only
  # from an operator-provided environment variable; it is never generated or
  # printed by the seed process.
  if test_user.api_keys.active_keys.none?
    raw_key = ENV["LOCAL_DEV_API_KEY"]
    raise "LOCAL_DEV_API_KEY must be set when creating the local development API key" if raw_key.blank?
    unless raw_key.match?(/\Arequiem_[0-9A-Za-z]{24}\z/)
      raise "LOCAL_DEV_API_KEY must match requiem_<24 alphanumeric characters>"
    end

    dev_key = test_user.api_keys.create!(
      name: "Local Dev Key",
      key_prefix: ApiKeyGenerator.extract_prefix(raw_key),
      key_hash: ApiKeyGenerator.hash_key(raw_key),
      full_key: raw_key
    )

    puts "✓ Created API key for #{test_user.email} (key_prefix: #{dev_key.key_prefix})"
  else
    puts "✓ Postgres-seeded API key already exists for #{test_user.email} " \
         "(key_prefix: #{test_user.api_keys.active_keys.first.key_prefix}, raw key not recoverable)"
  end
end
