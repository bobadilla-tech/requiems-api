# frozen_string_literal: true

class Users::RegistrationsController < Devise::RegistrationsController
  layout :resolve_layout

  before_action :configure_sign_up_params, only: [ :create ]
  before_action :configure_account_update_params, only: [ :update ]
  before_action :capture_referral_token, only: [ :new, :create ]

  # GET /resource/sign_up
  def new
    super
  end

  # POST /resource
  def create
    super do |resource|
      attribute_referral(resource, session.delete(:referral_token)) if resource.persisted?
    end
  end

  # GET /resource/edit
  # def edit
  #   super
  # end

  # PUT /resource
  # def update
  #   super
  # end

  # DELETE /resource
  # def destroy
  #   super
  # end

  # GET /resource/cancel
  # Forces the session data which is usually expired after sign
  # in to be expired now. This is useful if the user wants to
  # cancel oauth signing in/up in the middle of the process,
  # removing all OAuth session data.
  # def cancel
  #   super
  # end

  protected

  # If you have extra params to permit, append them to the sanitizer.
  def configure_sign_up_params
    devise_parameter_sanitizer.permit(:sign_up, keys: [ :name, :company ])
  end

  # If you have extra params to permit, append them to the sanitizer.
  def configure_account_update_params
    devise_parameter_sanitizer.permit(:account_update, keys: [ :name, :company ])
  end

  # The path used after sign up.
  def after_sign_up_path_for(resource)
    dashboard_root_path
  end

  # The path used after sign up for inactive accounts.
  def after_inactive_sign_up_path_for(resource)
    root_path
  end

  private

  def capture_referral_token
    session[:referral_token] = params[:ref] if params[:ref].present?
  end

  def attribute_referral(user, token)
    return if token.blank?

    referral_code = Rails.application.message_verifier("referral").verified(token)
    return if referral_code.blank?

    referrer = User.find_by(referral_code: referral_code)
    return if referrer.nil?
    return if referrer.id == user.id

    Referral.create!(referrer: referrer, referred_user: user)
  rescue ActiveRecord::RecordInvalid, ActiveRecord::RecordNotUnique
    # Silently ignore: duplicate referral (unique index) or validation failure
  end

  def resolve_layout
    case action_name
    when "edit", "update"
      "dashboard"
    else
      "application"
    end
  end
end
