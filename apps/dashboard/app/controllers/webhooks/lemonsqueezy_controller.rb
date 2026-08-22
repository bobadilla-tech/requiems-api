# frozen_string_literal: true

class Webhooks::LemonsqueezyController < ApplicationController
  # Webhooks are server-to-server calls; CSRF tokens are not applicable.
  # Authentication is enforced via HMAC-SHA256 signature verification below.
  skip_before_action :verify_authenticity_token # rubocop:disable Rails/LexicallyScopedActionFilter
  before_action :verify_signature

  def create
    event_name = params[:meta][:event_name]

    Rails.logger.info "[LemonSqueezy Webhook] Received: #{event_name}"

    case event_name
    when "subscription_created"
      handle_subscription_created
    when "subscription_updated"
      handle_subscription_updated
    when "subscription_cancelled", "subscription_expired"
      handle_subscription_cancelled
    when "subscription_resumed"
      handle_subscription_resumed
    when "subscription_payment_success"
      handle_payment_success
    else
      Rails.logger.warn "[LemonSqueezy Webhook] Unhandled event: #{event_name}"
    end

    head :ok
  rescue ActiveRecord::RecordInvalid => e
    Rails.logger.error "[LemonSqueezy Webhook] Error: #{e.message}"
    Rails.logger.error e.backtrace.join("\n")
    head :bad_request
  end

  private

  def verify_signature
    signature = request.headers["X-Signature"]

    unless signature
      Rails.logger.error "[LemonSqueezy Webhook] Missing signature"
      head :unauthorized
      return
    end

    secret = AppConfig.lemonsqueezy_signing_secret
    payload = request.body.read

    expected_signature = OpenSSL::HMAC.hexdigest(
      OpenSSL::Digest.new("sha256"),
      secret,
      payload
    )

    unless Rack::Utils.secure_compare(signature, expected_signature)
      Rails.logger.error "[LemonSqueezy Webhook] Invalid signature"
      head :unauthorized
      return
    end

    # Reset request body for controller to read
    request.body.rewind
  end

  def handle_subscription_created
    data = params[:data][:attributes]
    custom_data = params[:meta][:custom_data] || {}

    # Private deployment payment — completely separate product from shared plans
    if custom_data[:private_deployment_request_id].present?
      handle_private_deployment_payment(params[:data][:id], custom_data)
      return
    end

    user = User.find_by(id: custom_data[:user_id])
    unless user
      Rails.logger.error "[LemonSqueezy] User not found: #{custom_data[:user_id]}"
      return
    end

    plan_name = determine_plan_name(data[:variant_id])

    ActiveRecord::Base.transaction do
      subscription = user.subscription || user.build_subscription
      subscription.update!(
        lemonsqueezy_subscription_id: params[:data][:id],
        lemonsqueezy_customer_id: data[:customer_id],
        plan_name: plan_name,
        status: data[:status],
        current_period_start: data[:renews_at] ? Time.zone.parse(data[:renews_at]) - 1.month : Time.current,
        current_period_end: data[:renews_at],
        trial_ends_at: data[:trial_ends_at],
        cancel_at_period_end: false,
        # Clear any active admin promotion — the paid subscription supersedes it
        promoted_by_id: nil,
        promotion_reason: nil,
        promotion_expires_at: nil
      )

      # Mark referral as converted on first paid subscription — idempotent via pending? check
      user.referral_received&.mark_converted!(subscription) if plan_name != "free"
    end

    SubscriptionMailer.upgrade_notification(user, plan_name).deliver_later if plan_name != "free"

    Rails.logger.info "[LemonSqueezy] Subscription created for user #{user.id}: #{plan_name}"
  end

  def handle_private_deployment_payment(lemonsqueezy_subscription_id, custom_data)
    request_id = custom_data[:private_deployment_request_id]
    deployment_request = PrivateDeploymentRequest.find_by(id: request_id)

    unless deployment_request
      Rails.logger.error "[LemonSqueezy] PrivateDeploymentRequest not found: #{request_id}"
      return
    end

    deployment_request.update!(
      status: "pending",
      lemonsqueezy_subscription_id: lemonsqueezy_subscription_id
    )

    begin
      PrivateDeploymentMailer.request_received(deployment_request).deliver_later
      PrivateDeploymentMailer.admin_notification(deployment_request).deliver_later
    rescue StandardError => e
      Rails.logger.error "[LemonSqueezy] Failed to enqueue private deployment emails for request #{request_id}: #{e.message}"
    end

    Rails.logger.info "[LemonSqueezy] Private deployment payment received for request #{request_id}"
  end

  def handle_subscription_updated
    data = params[:data][:attributes]

    subscription = Subscription.find_by(lemonsqueezy_subscription_id: params[:data][:id])
    unless subscription
      Rails.logger.error "[LemonSqueezy] Subscription not found: #{data[:id]}"
      return
    end

    plan_name = determine_plan_name(data[:variant_id])
    previous_plan = subscription.plan_name

    ActiveRecord::Base.transaction do
      subscription.update!(
        status: data[:status],
        plan_name: plan_name,
        current_period_end: data[:renews_at],
        cancel_at_period_end: data[:ends_at].present?
      )
    end

    if plan_name != "free" && plan_name != previous_plan
      SubscriptionMailer.upgrade_notification(subscription.user, plan_name).deliver_later
    end

    Rails.logger.info "[LemonSqueezy] Subscription updated: #{subscription.id}"
  end

  def handle_subscription_cancelled
    data = params[:data][:attributes]

    subscription = Subscription.find_by(lemonsqueezy_subscription_id: params[:data][:id])
    unless subscription
      Rails.logger.error "[LemonSqueezy] Subscription not found: #{data[:id]}"
      return
    end

    ActiveRecord::Base.transaction do
      subscription.update!(
        status: "cancelled",
        cancel_at_period_end: true,
        plan_name: "free"
      )
    end

    Rails.logger.info "[LemonSqueezy] Subscription cancelled: #{subscription.id}"
  end

  def handle_subscription_resumed
    data = params[:data][:attributes]

    subscription = Subscription.find_by(lemonsqueezy_subscription_id: params[:data][:id])
    unless subscription
      Rails.logger.error "[LemonSqueezy] Subscription not found: #{data[:id]}"
      return
    end

    plan_name = determine_plan_name(data[:variant_id])

    ActiveRecord::Base.transaction do
      subscription.update!(
        status: data[:status],
        plan_name: plan_name,
        cancel_at_period_end: false
      )
    end

    Rails.logger.info "[LemonSqueezy] Subscription resumed: #{subscription.id}"
  end

  def handle_payment_success
    data = params[:data][:attributes]

    subscription = Subscription.find_by(lemonsqueezy_subscription_id: data[:subscription_id])
    return unless subscription

    subscription.update!(
      current_period_start: Time.current,
      current_period_end: data[:renews_at]
    )

    Rails.logger.info "[LemonSqueezy] Payment success for subscription: #{subscription.id}"
  end

  def determine_plan_name(variant_id)
    AppConfig.plan_name_for_variant_id(variant_id) || begin
      Rails.logger.warn "[LemonSqueezy] Unknown variant_id: #{variant_id}"
      "free"
    end
  end
end
