# frozen_string_literal: true

module DivisionsHelper
  DEFAULT_CAL_COM_STRATEGY_URL = "https://cal.com/alexandra-flores/book-a-strategy-call"

  def cal_com_strategy_url
    ENV.fetch("CAL_COM_STRATEGY_URL", "").presence || DEFAULT_CAL_COM_STRATEGY_URL
  end

  # I18n may return arrays of hashes (string keys) or plain strings (e.g. industries list).
  def division_locale_array(scope, key)
    value = I18n.t("#{scope}.#{key}", default: [])
    return [] unless value.is_a?(Array)

    value.filter_map do |entry|
      case entry
      when Hash
        entry.transform_keys(&:to_s)
      when String, Symbol
        entry.to_s
      end
    end
  end
end
