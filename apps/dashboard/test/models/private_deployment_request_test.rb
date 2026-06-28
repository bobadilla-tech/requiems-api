# frozen_string_literal: true

require "test_helper"

class PrivateDeploymentRequestTest < ActiveSupport::TestCase
  def setup
    @user = create_user(email: "deployments@example.com")
  end

  test "requires at least one selected service" do
    request = build_request(selected_services: [])

    assert_not request.valid?
    assert_includes request.errors[:selected_services], "must include at least one service"
  end

  test "rejects invalid selected services" do
    request = build_request(selected_services: %w[email invalid_service])

    assert_not request.valid?
    assert_includes request.errors[:selected_services], "contains invalid services: invalid_service"
  end

  test "requires a long tenant secret only when active" do
    active_request = build_request(status: "active", tenant_secret: "short-secret")
    pending_request = build_request(status: "pending", tenant_secret: "short-secret")

    assert_not active_request.valid?
    assert_includes active_request.errors[:tenant_secret], "is too short (minimum is 32 characters)"

    assert pending_request.valid?
  end

  test "monthly_price_dollars uses stored monthly price when present" do
    request = build_request(monthly_price_cents: 42_500, server_tier: "scale")

    assert_equal 425.0, request.monthly_price_dollars
  end

  test "monthly_price_dollars falls back to tier price when monthly price is nil" do
    request = build_request(monthly_price_cents: nil, server_tier: "growth")

    assert_equal 300.0, request.monthly_price_dollars
  end

  test "total_price_dollars uses yearly pricing table for yearly billing" do
    request = build_request(server_tier: "enterprise", billing_cycle: "yearly")

    assert_equal 4896.0, request.total_price_dollars
  end

  test "live_url returns tenant url when subdomain is present" do
    request = build_request(subdomain_slug: "acme-team")

    assert_equal "https://acme-team.requiems.xyz", request.live_url
  end

  test "rejects invalid subdomain slug format" do
    request = build_request(subdomain_slug: "UPPER_CASE!")

    assert_not request.valid?
    assert_includes request.errors[:subdomain_slug],
                    I18n.t("private_deployment_request.must_be_lowercase_letters_numbers_and")
  end

  test "accepts valid subdomain slug" do
    request = build_request(subdomain_slug: "acme-team-42")

    assert request.valid?
  end

  test "tier_specs returns configured specs for tier" do
    request = build_request(server_tier: "starter")

    assert_equal(
      { hetzner: "CPX21", vcpu: 3, ram: "4 GB", ssd: "80 GB" },
      request.tier_specs
    )
  end

  private

  def build_request(**attributes)
    PrivateDeploymentRequest.new(
      {
        user: @user,
        company: "Acme",
        contact_name: "Taylor",
        contact_email: "deployments@example.com",
        server_tier: "starter",
        billing_cycle: "monthly",
        status: "pending",
        selected_services: %w[email text],
        monthly_price_cents: 20_000
      }.merge(attributes)
    )
  end
end
