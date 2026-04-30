# frozen_string_literal: true

module DivisionsHelper
  def cal_com_strategy_url
    ENV.fetch("CAL_COM_STRATEGY_URL", "").presence
  end

  def division_locale_array(scope, key)
    value = I18n.t("#{scope}.#{key}", default: [])
    return [] unless value.is_a?(Array)

    value
  end
end
