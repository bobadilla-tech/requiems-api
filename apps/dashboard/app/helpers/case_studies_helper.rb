# frozen_string_literal: true

module CaseStudiesHelper
  include ApisHelper

  # Same logical paths as home trusted-by / logo_marquee (`app/assets/images/...`).
  def case_study_logo_asset(slug)
    case slug.to_s
    when "verigeo"
      "companies/verigeo.png"
    when "compilestrength"
      "companies/compile-strength.png"
    end
  end

  # Rows from i18n: [{ id:, benefit: }, ...]; merges API display name from catalog when present.
  def case_study_api_rows(i18n_scope)
    raw = I18n.t("#{i18n_scope}.api_rows", default: [])
    return [] unless raw.is_a?(Array)

    raw.filter_map do |row|
      next unless row.is_a?(Hash)

      r = row.transform_keys(&:to_sym)
      id = r[:id].presence&.to_s
      next if id.blank?

      api = find_api(id)
      {
        id: id,
        benefit: r[:benefit].to_s,
        name: api&.dig("name") || id.titleize
      }
    end
  end

  def case_study_json_ld(slug)
    scope = "case_studies.#{slug}"
    headline = I18n.t("#{scope}.meta_title")
    description = I18n.t("#{scope}.meta_description")

    org = {
      "verigeo" => { "@type" => "Organization", "name" => "Verigeo", "url" => "https://verigeo.pe" },
      "compilestrength" => { "@type" => "Organization", "name" => "CompileStrength", "url" => "https://compilestrength.com" }
    }[slug.to_s] || {}

    data = {
      "@context" => "https://schema.org",
      "@type" => "Article",
      "headline" => headline,
      "description" => description,
      "about" => org,
      "publisher" => {
        "@type" => "Organization",
        "name" => "Requiems API",
        "url" => "https://requiemsapi.com"
      }
    }

    data.compact_blank.to_json
  end
end
