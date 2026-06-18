# frozen_string_literal: true

class Dashboard::ReferralsController < Dashboard::BaseController

  def show
    @referral_link = new_user_registration_url(
      ref: current_user.referral_token,
      locale: I18n.locale
    )
    @total_referred = current_user.referrals_given.count
    @total_converted = current_user.referrals_given.converted.count
    @recent_referrals = current_user.referrals_given.includes(:referred_user).order(created_at: :desc).limit(10)
  end
end
