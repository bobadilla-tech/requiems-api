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

  BREADCRUMB_NAMES = {
    "apis"          => "APIs",
    "systems"       => "Systems",
    "categories"    => "Categories",
    "examples"      => "Examples",
    "case-studies"  => "Case Studies",
    "pricing"       => "Pricing",
    "faq"           => "FAQ",
    "about"         => "About",
    "team"          => "Team",
    "contact"       => "Contact",
    "blog"          => "Blog",
    "changelog"     => "Changelog",
    "api_reference" => "API Reference",
    "glossary"      => "Glossary",
    "error_codes"   => "Error Codes",
    "status"        => "Status",
    "divisions"     => "Divisions",
    "privacy"       => "Privacy Policy",
    "terms"         => "Terms of Service"
  }.freeze

  def breadcrumb_json_ld
    locale_prefix_re = /\A\/(#{I18n.available_locales.map { |l| Regexp.escape(l.to_s) }.join("|")})(?=\/|\z)/
    base_path = request.path.sub(locale_prefix_re, "").presence || "/"
    return nil if base_path == "/"

    segments = base_path.split("/").reject(&:blank?)
    return nil if segments.empty?

    items = [ { "@type" => "ListItem", "position" => 1, "name" => "Home", "item" => "https://requiems.xyz/#{I18n.locale}/" } ]
    segments.each_with_index do |segment, i|
      items << {
        "@type"    => "ListItem",
        "position" => i + 2,
        "name"     => BREADCRUMB_NAMES[segment] || segment.gsub(/[-_]/, " ").split.map(&:capitalize).join(" "),
        "item"     => "https://requiems.xyz/#{I18n.locale}/#{segments[0..i].join("/")}"
      }
    end

    { "@context" => "https://schema.org", "@type" => "BreadcrumbList", "itemListElement" => items }.to_json
  end

  def faq_json_ld
    rate_answer = t("home.faq.rate_limits.a1") + " " +
      PlanConfig::PLAN_NAMES.map { |p|
        t("home.faq.rate_limits.#{p}", rate_limit: number_with_delimiter(PlanConfig::PLANS[p][:rate_limit_per_minute]))
      }.join(". ") + "."

    pairs = [
      [ t("home.faq.getting_started.q1"), strip_tags(t("home.faq.getting_started.a1_html", docs_link: t("home.faq.getting_started.docs_link"))) ],
      [ t("home.faq.getting_started.q2"), t("home.faq.getting_started.a2") ],
      [ t("home.faq.getting_started.q3"), t("home.faq.getting_started.a3") ],
      [ t("home.faq.authentication.q1"),  "#{t("home.faq.authentication.a1")} Authorization: Bearer YOUR_API_KEY" ],
      [ t("home.faq.authentication.q2"),  strip_tags(t("home.faq.authentication.a2_html")) ],
      [ t("home.faq.authentication.q3"),  t("home.faq.authentication.a3") ],
      [ t("home.faq.billing.q1"),         t("home.faq.billing.a1") ],
      [ t("home.faq.billing.q2"),         t("home.faq.billing.a2") ],
      [ t("home.faq.billing.q3"),         t("home.faq.billing.a3") ],
      [ t("home.faq.billing.q4"),         t("home.faq.billing.a4") ],
      [ t("home.faq.rate_limits.q1"),     rate_answer ],
      [ t("home.faq.rate_limits.q2"),     strip_tags(t("home.faq.rate_limits.a2_html")) ],
      [ t("home.faq.support.q1"),         "#{t("home.faq.support.a1")} #{t("home.faq.support.docs_channel.label")}: #{t("home.faq.support.docs_channel.description")}. #{t("home.faq.support.contact_channel.label")}: #{t("home.faq.support.contact_channel.description")}." ],
      [ t("home.faq.support.q2"),         "#{t("home.faq.support.a2")} #{t("home.faq.support.free_dev")}. #{t("home.faq.support.business_time")}. #{t("home.faq.support.professional_time")}." ],
      [ t("home.faq.support.q3"),         strip_tags(t("home.faq.support.a3_html", sales_link: t("home.faq.support.contact_sales"))) ]
    ]

    {
      "@context"   => "https://schema.org",
      "@type"      => "FAQPage",
      "mainEntity" => pairs.map { |q, a|
        { "@type" => "Question", "name" => q, "acceptedAnswer" => { "@type" => "Answer", "text" => a } }
      }
    }.to_json
  end

  def organization_json_ld
    {
      "@context" => "https://schema.org",
      "@type" => "Organization",
      "name" => "Requiems API",
      "url" => "https://requiems.xyz",
      "logo" => {
        "@type" => "ImageObject",
        "url" => "https://requiems.xyz/logo.png",
        "width" => 512,
        "height" => 512
      },
      "image" => "https://requiems.xyz/og-image.png",
      "description" => "All-in-one backend for SaaS products. Authentication, validation, fraud detection, payments intelligence, and global data through one unified API.",
      "sameAs" => []
    }.to_json
  end

  def website_json_ld
    {
      "@context" => "https://schema.org",
      "@type" => "WebSite",
      "name" => "Requiems API",
      "url" => "https://requiems.xyz",
      "potentialAction" => {
        "@type" => "SearchAction",
        "target" => {
          "@type" => "EntryPoint",
          "urlTemplate" => "https://requiems.xyz/en/apis?search={search_term_string}"
        },
        "query-input" => "required name=search_term_string"
      }
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

  # Renders the <link rel="canonical"> and all <link rel="alternate" hreflang="...">
  # tags for the current page.
  #
  # Strategy:
  #   - Canonical always ends with "/" so /en and /en/ are treated as one URL by Google.
  #   - Canonical is derived from request.path (no query string) so UTM params don't
  #     create duplicate index entries.
  #   - hreflang alternates are built from the locale-stripped base path so every locale
  #     version points at the same content node. Each alternate also ends with "/" for
  #     consistency with the canonical.
  #   - x-default points at the English version (site default locale).
  #   - request.base_url is used instead of a hardcoded host so staging/dev environments
  #     emit correct self-referential URLs automatically.
  def seo_head_tags
    # Regex that matches a leading locale segment, e.g. /en, /es, /fr.
    # The lookahead (?=/|\z) ensures we only strip a full segment, not a prefix of a word.
    locale_re = Regexp.new(
      "\\A/(#{I18n.available_locales.map { |l| Regexp.escape(l.to_s) }.join('|')})(?=/|\\z)"
    )

    # Strip the locale prefix to get the language-neutral path ("/", "/pricing", etc.).
    base_path = request.path.sub(locale_re, "")
    base_path = "/" if base_path.blank?

    # For hreflang hrefs we want no trailing slash on non-root paths ("/pricing", not
    # "/pricing/"), then append "/" after the locale — giving "/en/pricing/" uniformly.
    normalized = base_path == "/" ? "" : base_path.chomp("/")

    # Canonical: always trailing slash, never query params.
    # e.g. /en?utm=x → /en/  |  /en/pricing → /en/pricing/
    canonical_path = request.path.chomp("/") + "/"

    tags = []

    tags << tag.link(rel: "canonical", href: request.base_url + canonical_path)

    I18n.available_locales.each do |loc|
      tags << tag.link(rel: "alternate", hreflang: loc.to_s,
                       href: "#{request.base_url}/#{loc}#{normalized}/")
    end

    # x-default signals the preferred URL for users whose language isn't listed.
    tags << tag.link(rel: "alternate", hreflang: "x-default",
                     href: "#{request.base_url}/#{I18n.default_locale}#{normalized}/")

    safe_join(tags, "\n    ")
  end
end
