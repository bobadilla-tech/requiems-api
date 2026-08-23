# frozen_string_literal: true

require "test_helper"

class Webhooks::LemonsqueezyControllerTest < ActionDispatch::IntegrationTest
  include ActionMailer::TestHelper
  include ActiveJob::TestHelper

  def setup
    @user = create_user(email: "payer@example.com")
    @pdr = PrivateDeploymentRequest.create!(
      user: @user,
      company: "Acme Corp",
      contact_name: "Sam",
      contact_email: @user.email,
      server_tier: "starter",
      billing_cycle: "monthly",
      status: "pending_payment",
      selected_services: %w[email text],
      monthly_price_cents: 20_000
    )
    clear_enqueued_jobs
    ActionMailer::Base.deliveries.clear
  end

  test "rejects webhook without signature" do
    post webhooks_lemonsqueezy_path, params: webhook_payload.to_json, headers: {
      "CONTENT_TYPE" => "application/json"
    }

    assert_response :unauthorized
  end

  test "marks private deployment request pending after subscription is created" do
    assert_enqueued_emails 2 do
      post webhooks_lemonsqueezy_path, params: webhook_payload.to_json, headers: signed_headers(webhook_payload)
    end

    assert_response :ok

    @pdr.reload
    assert_equal "pending", @pdr.status
    assert_equal "sub_123", @pdr.lemonsqueezy_subscription_id
  end

  test "returns ok for unknown private deployment request id" do
    payload = webhook_payload(private_deployment_request_id: "999999")

    assert_no_enqueued_emails do
      post webhooks_lemonsqueezy_path, params: payload.to_json, headers: signed_headers(payload)
    end

    assert_response :ok
  end

  test "subscription_created marks referred user's referral as converted" do
    referrer = create_user(email: "referrer@example.com")
    referred = create_user(email: "referred@example.com")
    referral = Referral.create!(referrer: referrer, referred_user: referred)

    payload = subscription_created_payload(user_id: referred.id)
    post webhooks_lemonsqueezy_path, params: payload.to_json, headers: signed_headers(payload)

    assert_response :ok
    referral.reload
    assert referral.converted?
    assert_not_nil referral.converted_at
  end

  test "subscription_created is idempotent — repeated webhook does not alter converted_at" do
    referrer = create_user(email: "ref2@example.com")
    referred = create_user(email: "ref2user@example.com")
    Referral.create!(referrer: referrer, referred_user: referred)

    payload = subscription_created_payload(user_id: referred.id)
    post webhooks_lemonsqueezy_path, params: payload.to_json, headers: signed_headers(payload)
    first_converted_at = referred.referral_received.reload.converted_at

    # Second delivery — subscription update is idempotent; referral stays converted
    post webhooks_lemonsqueezy_path, params: payload.to_json, headers: signed_headers(payload)
    assert_equal first_converted_at, referred.referral_received.reload.converted_at
  end

  test "subscription_created for non-referred user does not create referral" do
    payload = subscription_created_payload(user_id: @user.id)
    post webhooks_lemonsqueezy_path, params: payload.to_json, headers: signed_headers(payload)

    assert_response :ok
    assert_nil @user.referral_received
  end

  test "subscription_created sets plan (billing cycle) from the monthly variant_id" do
    payload = subscription_created_payload(user_id: @user.id, variant_id: AppConfig.lemonsqueezy_developer_monthly_variant_id)
    post webhooks_lemonsqueezy_path, params: payload.to_json, headers: signed_headers(payload)

    assert_response :ok
    assert_equal "monthly", @user.subscription.reload.plan
  end

  test "subscription_created sets plan (billing cycle) from the yearly variant_id" do
    payload = subscription_created_payload(user_id: @user.id, variant_id: AppConfig.lemonsqueezy_developer_yearly_variant_id)
    post webhooks_lemonsqueezy_path, params: payload.to_json, headers: signed_headers(payload)

    assert_response :ok
    assert_equal "yearly", @user.subscription.reload.plan
  end

  test "subscription_updated updates plan (billing cycle) when the variant_id switches from monthly to yearly" do
    created = subscription_created_payload(user_id: @user.id, variant_id: AppConfig.lemonsqueezy_developer_monthly_variant_id)
    post webhooks_lemonsqueezy_path, params: created.to_json, headers: signed_headers(created)
    assert_equal "monthly", @user.subscription.reload.plan

    updated = {
      meta: { event_name: "subscription_updated" },
      data: {
        id: created[:data][:id],
        attributes: {
          status: "active",
          variant_id: AppConfig.lemonsqueezy_developer_yearly_variant_id,
          renews_at: 1.year.from_now.iso8601
        }
      }
    }
    post webhooks_lemonsqueezy_path, params: updated.to_json, headers: signed_headers(updated)

    assert_response :ok
    assert_equal "yearly", @user.subscription.reload.plan
  end

  test "subscription_resumed updates plan (billing cycle) from the variant_id" do
    created = subscription_created_payload(user_id: @user.id, variant_id: AppConfig.lemonsqueezy_developer_monthly_variant_id)
    post webhooks_lemonsqueezy_path, params: created.to_json, headers: signed_headers(created)
    assert_equal "monthly", @user.subscription.reload.plan

    resumed = {
      meta: { event_name: "subscription_resumed" },
      data: {
        id: created[:data][:id],
        attributes: {
          status: "active",
          variant_id: AppConfig.lemonsqueezy_developer_yearly_variant_id
        }
      }
    }
    post webhooks_lemonsqueezy_path, params: resumed.to_json, headers: signed_headers(resumed)

    assert_response :ok
    assert_equal "yearly", @user.subscription.reload.plan
  end

  test "subscription_created falls back to monthly billing cycle for an unrecognized variant_id" do
    payload = subscription_created_payload(user_id: @user.id, variant_id: "unrecognized_variant")
    post webhooks_lemonsqueezy_path, params: payload.to_json, headers: signed_headers(payload)

    assert_response :ok
    assert_equal "monthly", @user.subscription.reload.plan
  end

  private

  def subscription_created_payload(user_id:, variant_id: AppConfig.lemonsqueezy_developer_monthly_variant_id)
    {
      meta: {
        event_name: "subscription_created",
        custom_data: {
          user_id: user_id.to_s
        }
      },
      data: {
        id: "sub_#{SecureRandom.hex(4)}",
        attributes: {
          customer_id: "cust_#{SecureRandom.hex(4)}",
          status: "active",
          variant_id: variant_id
        }
      }
    }
  end

  def webhook_payload(private_deployment_request_id: @pdr.id.to_s)
    {
      meta: {
        event_name: "subscription_created",
        custom_data: {
          private_deployment_request_id: private_deployment_request_id,
          user_id: @user.id.to_s
        }
      },
      data: {
        id: "sub_123",
        attributes: {
          customer_id: "cust_123",
          status: "active",
          variant_id: AppConfig.lemonsqueezy_developer_monthly_variant_id
        }
      }
    }
  end

  def signed_headers(payload)
    raw_payload = payload.to_json
    signature = OpenSSL::HMAC.hexdigest(
      OpenSSL::Digest.new("sha256"),
      AppConfig.lemonsqueezy_signing_secret,
      raw_payload
    )

    {
      "CONTENT_TYPE" => "application/json",
      "X-Signature" => signature
    }
  end
end
