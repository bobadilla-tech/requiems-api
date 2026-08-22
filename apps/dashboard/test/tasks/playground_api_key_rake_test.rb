# frozen_string_literal: true

require "test_helper"

class PlaygroundApiKeyRakeTest < ActiveSupport::TestCase
  parallelize(workers: 1)

  RAW_KEY = "requiem_ratetasktestkey000000001"

  setup do
    Rails.application.load_tasks
  end

  teardown do
    Rake::Task["playground:provision_key"].reenable
    ENV.delete("PLAYGROUND_API_KEY")
  end

  test "raises when PLAYGROUND_API_KEY is not set" do
    ENV.delete("PLAYGROUND_API_KEY")
    assert_raises(RuntimeError) { Rake::Task["playground:provision_key"].invoke }
  end

  test "raises when PLAYGROUND_API_KEY does not match the requiem_ format" do
    ENV["PLAYGROUND_API_KEY"] = "not-a-valid-key"
    assert_raises(RuntimeError) { Rake::Task["playground:provision_key"].invoke }
  end

  test "creates an active, bounded-plan key that authenticates for the presented raw value" do
    ENV["PLAYGROUND_API_KEY"] = RAW_KEY
    Rake::Task["playground:provision_key"].invoke

    user = User.find_by!(email: "playground@internal.requiems.xyz")
    api_key = user.api_keys.sole

    assert api_key.active?
    assert_nil api_key.revoked_at
    assert api_key.verify_key(RAW_KEY)
    assert_equal "developer", user.subscription.plan_name
    assert_equal "active", user.subscription.status
  end

  test "is idempotent — running twice with the same key does not raise or duplicate the key" do
    ENV["PLAYGROUND_API_KEY"] = RAW_KEY
    Rake::Task["playground:provision_key"].invoke
    Rake::Task["playground:provision_key"].reenable
    Rake::Task["playground:provision_key"].invoke

    user = User.find_by!(email: "playground@internal.requiems.xyz")
    assert_equal 1, user.api_keys.count
    assert user.api_keys.sole.active?
  end
end
