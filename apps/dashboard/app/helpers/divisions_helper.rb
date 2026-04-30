# frozen_string_literal: true

module DivisionsHelper
  DEFAULT_CAL_COM_STRATEGY_URL = "https://cal.com/alexandra-flores/book-a-strategy-call"

  def cal_com_strategy_url
    ENV.fetch("CAL_COM_STRATEGY_URL", "").presence || DEFAULT_CAL_COM_STRATEGY_URL
  end

  # I18n may return hashes with symbol or string keys; normalize for ERB.
  def division_locale_array(scope, key)
    value = I18n.t("#{scope}.#{key}", default: [])
    return [] unless value.is_a?(Array)

    value.map do |entry|
      next entry unless entry.is_a?(Hash)

      entry.transform_keys(&:to_s)
    end
  end
end
