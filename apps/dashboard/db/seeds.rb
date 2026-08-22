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

  # Postgres-side dev API key for exercising the Go auth path directly
  # (apps/api/platform/middleware/apikeyauth.go), independent of the
  # Cloudflare-Worker/KV-seeded keys from
  # apps/workers/auth-gateway/scripts/seed-dev.ts.
  #
  # NOTE: ApiKey.create! is deliberately NOT used with a bare `name:`/`user:`
  # here — ApiKey#request_key_from_server only takes the local-generation
  # branch under Rails.env.test?; in development it would call out to
  # Cloudflare::ApiManagementService over HTTP. Calling ApiKeyGenerator
  # directly and passing the resulting key_prefix/key_hash in makes
  # request_key_from_server's before_validation early-return (it only
  # generates when key_prefix is blank), so no network call happens.
  if test_user.api_keys.active_keys.none?
    full_key = ApiKeyGenerator.generate(environment: :live)

    ApiKey.create!(
      user: test_user,
      name: "Local Dev Key",
      key_prefix: ApiKeyGenerator.extract_prefix(full_key),
      key_hash: ApiKeyGenerator.hash_key(full_key),
      active: true
    )

    puts "✓ Created Postgres-seeded API key for #{test_user.email}: #{full_key}"
    puts "  (queryable by key_prefix; use this with the Go auth path, not the Worker)"
  else
    puts "✓ Postgres-seeded API key already exists for #{test_user.email} " \
         "(key_prefix: #{test_user.api_keys.active_keys.first.key_prefix}, raw key not recoverable)"
  end
end
