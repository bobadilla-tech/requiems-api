# frozen_string_literal: true

require "test_helper"

# Public and settings copy that PR-A corrected: no promises of alerts or weekly
# reports that do not exist, and credits / "billing day" wording.
class UsageCopyTest < ActionDispatch::IntegrationTest
  COMING_SOON = { "en" => /Coming soon/, "es" => /Próximamente/, "fr" => /Bientôt disponible/ }.freeze
  OLD_ALERT_PROMISE = {
    "en" => "send you an alert",
    "es" => "te enviaremos una alerta",
    "fr" => "nous vous envoyons une alerte"
  }.freeze

  %w[en es fr].each do |locale|
    test "settings marks usage alerts and weekly reports as coming soon in #{locale}" do
      sign_in create_user(email: "copy-#{locale}@example.com")

      get dashboard_settings_path(locale: locale)

      assert_response :success
      assert_select "p", text: COMING_SOON[locale], minimum: 2
    end

    test "faq no longer promises alerts and speaks of credits in #{locale}" do
      get "/#{locale}/faq"

      assert_response :success
      assert_not_includes response.body.downcase, OLD_ALERT_PROMISE.fetch(locale)
      assert_includes response.body, ERB::Util.html_escape(I18n.t("home.faq.billing.a1", locale: locale))
      assert_includes response.body, ERB::Util.html_escape(I18n.t("home.faq.billing.a2", locale: locale))
    end

    test "glossary defines the credit as the unit of usage in #{locale}" do
      get "/#{locale}/glossary"

      assert_response :success
      definitions = I18n.t("home.glossary.terms", locale: locale).map { |term| term[:definition] }
      credit = definitions.find { |d| d.match?(/\bcr[eé]dit/i) }

      assert credit, "no glossary definition mentions credits"
      assert_includes response.body, ERB::Util.html_escape(credit)
    end

    test "changelog shows the usage correction entry in #{locale}" do
      get "/#{locale}/changelog"

      assert_response :success
      assert_select "h2", text: I18n.t("home.changelog.usage_cycle_fix.title", locale: locale)
      I18n.t("home.changelog.usage_cycle_fix.items", locale: locale).each do |item|
        assert_includes response.body, ERB::Util.html_escape(item)
      end
    end
  end
end
