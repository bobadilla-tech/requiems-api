# frozen_string_literal: true

class ApplicationController < ActionController::Base
  include Pagy::Method
  include HttpAcceptLanguage::EasyAccess

  # Navbar and shared layouts use category SVGs (same as home / API directory).
  helper ApisHelper

  allow_browser versions: :modern

  stale_when_importmap_changes

  NON_LOCALE_CONTROLLERS = %w[rails/health sitemap api_proxy locale webhooks/lemonsqueezy pwa].freeze

  before_action :set_locale

  private

  def set_locale
    requested = params[:locale].presence
    validated = requested if I18n.available_locales.map(&:to_s).include?(requested)
    I18n.locale = validated ||
                  current_user_locale ||
                  http_accept_language.compatible_language_from(I18n.available_locales) ||
                  I18n.default_locale

    # Skip redirect for non-locale-scoped controllers (sitemap, health check, webhooks, etc.).
    # Cannot use request.path_parameters.key?(:locale) because Rails does not add optional
    # route segments to path_parameters when they are absent from the URL.
    return if NON_LOCALE_CONTROLLERS.include?(controller_path)

    # Redirect all un-prefixed locale-scoped URLs to their explicit /{locale}/... equivalent
    # so every public page has a single canonical URL (e.g. / → /en/, /pricing → /en/pricing).
    if params[:locale].blank? && (request.get? || request.head?)
      target = "/#{I18n.locale}#{request.path}"
      target += "?#{request.query_string}" if request.query_string.present?
      redirect_to target, status: :moved_permanently, allow_other_host: false
    end
  end

  def current_user_locale
    current_user&.locale&.presence&.to_sym
  end

  def default_url_options
    { locale: I18n.locale }
  end
end
