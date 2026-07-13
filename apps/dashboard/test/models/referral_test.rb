# frozen_string_literal: true

require "test_helper"

class ReferralTest < ActiveSupport::TestCase
  def setup
    @referrer = create_user(email: "referrer@example.com")
    @referred = create_user(email: "referred@example.com")
    @referral = Referral.create!(referrer: @referrer, referred_user: @referred)
  end

  test "valid referral is persisted" do
    assert @referral.persisted?
  end

  test "starts with pending status" do
    assert @referral.pending?
    assert_not @referral.converted?
  end

  test "referrer has many referrals_given" do
    assert_includes @referrer.referrals_given, @referral
  end

  test "referred user has one referral_received" do
    assert_equal @referral, @referred.referral_received
  end

  test "blocks self-referral at model validation level" do
    self_referral = Referral.new(referrer: @referrer, referred_user: @referrer)
    assert_not self_referral.valid?
    I18n.with_locale(:en) do
      assert_includes self_referral.errors[:referred_user], "cannot be the same as referrer"
    end
  end

  test "referred_user_id has unique index — second referral for same user is invalid" do
    third_user = create_user(email: "another@example.com")
    assert_raises(ActiveRecord::RecordNotUnique) do
      Referral.create!(referrer: third_user, referred_user: @referred)
    end
  end

  test "mark_converted! transitions status to converted" do
    subscription = @referred.build_subscription(
      plan_name: "developer",
      status: "active",
      lemonsqueezy_subscription_id: "sub_abc"
    )
    subscription.save!

    @referral.mark_converted!(subscription)

    @referral.reload
    assert @referral.converted?
    assert_not_nil @referral.converted_at
    assert_equal subscription, @referral.converting_subscription
  end

  test "mark_converted! is idempotent — second call is a no-op" do
    subscription = @referred.build_subscription(
      plan_name: "developer",
      status: "active",
      lemonsqueezy_subscription_id: "sub_def"
    )
    subscription.save!

    @referral.mark_converted!(subscription)
    first_converted_at = @referral.reload.converted_at

    @referral.mark_converted!(subscription)
    assert_equal first_converted_at, @referral.reload.converted_at
  end

  test "pending? and converted? reflect current status" do
    assert @referral.pending?
    assert_not @referral.converted?

    @referral.update!(status: "converted", converted_at: Time.current)
    assert_not @referral.pending?
    assert @referral.converted?
  end

  test "referrer has a referral_token method" do
    assert_respond_to @referrer, :referral_token
    assert @referrer.referral_token.present?
  end

  test "new user gets a referral_code on create" do
    user = create_user(email: "newuser@example.com")
    assert user.referral_code.present?
    assert_equal 8, user.referral_code.length
  end
end
