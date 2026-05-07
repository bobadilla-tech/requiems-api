# frozen_string_literal: true

module ApplicationHelper
  include ApisHelper

  LOCALE_NAMES = {
    en: "English",
    es: "Español",
    fr: "Français"
  }.freeze

  def locale_name(locale)
    LOCALE_NAMES[locale.to_sym] || locale.to_s.upcase
  end

  # POST /locale is declared outside the optional `(:locale)` scope. `default_url_options`
  # still merges `locale: I18n.locale`, so `switch_locale_path` becomes "/locale?locale=fr".
  # That duplicates the `locale` param with the select field and breaks language switching.
  def locale_switch_path
    "#{request.script_name}/locale".gsub(%r{/+}, "/")
  end

  def global_search_data
    {
      apis: searchable_apis,
      examples: searchable_examples,
      pages: searchable_pages
    }
  end

  def gravatar_url(email, size: 80)
    hash = Digest::MD5.hexdigest(email.to_s.downcase.strip)
    "https://www.gravatar.com/avatar/#{hash}?s=#{size}&d=mp"
  end

  def status_code_variant(code)
    case code.to_s
    when /\A2/ then "success"
    when /\A4/ then "warning"
    else "danger"
    end
  end

  def compact_number(value)
    number_to_human(
      value.to_i,
      format: "%n%u",
      precision: 3,
      significant: true,
      strip_insignificant_zeros: true,
      units: {
        thousand: "K",
        million: "M",
        billion: "B",
        trillion: "T",
        quadrillion: "Q"
      }
    )
  end

  def organization_json_ld
    {
      "@context" => "https://schema.org",
      "@type" => "Organization",
      "name" => "Requiems API",
      "url" => "https://requiems.xyz",
      "logo" => "https://requiems.xyz/logo.png",
      "description" => "All-in-one backend for SaaS products. Authentication, validation, fraud detection, payments intelligence, and global data through one unified API.",
      "sameAs" => []
    }.to_json
  end

  def api_json_ld(api)
    {
      "@context" => "https://schema.org",
      "@type" => "WebAPI",
      "name" => api["name"],
      "description" => api["description"],
      "documentation" => api["documentation_url"].presence || "https://requiems.xyz/en/apis/#{api["id"]}",
      "provider" => {
        "@type" => "Organization",
        "name" => "Requiems API",
        "url" => "https://requiems.xyz"
      }
    }.compact_blank.to_json
  end

  private

  def searchable_apis
    live_apis.map do |api|
      category = find_category(Array(api["categories"]).first)
      {
        id: api["id"],
        title: api["name"],
        description: api["description"],
        url: api["documentation_url"],
        category: category["name"],
        category_icon: category["icon"],
        type: "api",
        tags: api["tags"] || [],
        endpoints_count: api["endpoints_count"]
      }
    end
  end

  def searchable_examples
    examples_data = YAML.load_file(Rails.root.join("config", "examples.yml"))
    examples_data["examples"].map do |example|
      {
        id: example["id"],
        title: example["title"],
        description: example["description"],
        url: "/examples/#{example['id']}",
        category: example["category"],
        type: "example",
        difficulty: example["difficulty"],
        technologies: example["technologies"]
      }
    end
  end

  def searchable_pages
    pages_data = YAML.load_file(Rails.root.join("config", "searchable_pages.yml"))
    pages_data["pages"].map { |page| page.merge("type" => "page") }
  end
end
