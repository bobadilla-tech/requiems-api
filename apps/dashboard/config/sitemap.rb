# frozen_string_literal: true

require "yaml"
require_relative "../lib/division_slugs.rb"
require_relative "../lib/comparison_slugs.rb"
require_relative "../lib/industry_slugs.rb"

SitemapGenerator::Sitemap.default_host = "https://requiemsapi.com"
SitemapGenerator::Sitemap.compress      = false # write .xml, not .xml.gz
SitemapGenerator::Sitemap.include_root  = false # all pages added manually
SitemapGenerator::Sitemap.include_index = true  # emit sitemap.xml as index

catalog     = YAML.load_file(Rails.root.join("config", "api_catalog.yml"))
live_apis   = catalog["apis"].select { |api| api["status"] == "live" }
categories  = catalog["categories"].map { |c| c["id"] }
examples    = YAML.load_file(Rails.root.join("config", "examples.yml"))["examples"].map { |e| e["id"] }

SYSTEM_SLUGS = %w[
  identity-risk
  payments-intelligence
  global-data
  data-integrity
  developer-utilities
].freeze

STATIC_PAGES = [
  { path: "",                    changefreq: "weekly",  priority: 1.0 },
  { path: "/apis",               changefreq: "weekly",  priority: 0.9 },
  { path: "/systems",            changefreq: "weekly",  priority: 0.85 },
  { path: "/pricing",            changefreq: "monthly", priority: 0.8 },
  { path: "/api_reference",      changefreq: "monthly", priority: 0.7 },
  { path: "/faq",                changefreq: "monthly", priority: 0.6 },
  { path: "/changelog",          changefreq: "weekly",  priority: 0.6 },
  { path: "/blog",               changefreq: "weekly",  priority: 0.6 },
  { path: "/examples",           changefreq: "weekly",  priority: 0.65 },
  { path: "/domain-checker",     changefreq: "weekly",  priority: 0.55 },
  { path: "/ai",                 changefreq: "monthly", priority: 0.55 },
  { path: "/for-llms",           changefreq: "monthly", priority: 0.5 },
  { path: "/about",              changefreq: "monthly", priority: 0.5 },
  { path: "/security",           changefreq: "monthly", priority: 0.5 },
  { path: "/team",               changefreq: "monthly", priority: 0.5 },
  { path: "/contact",            changefreq: "monthly", priority: 0.5 },
  { path: "/glossary",           changefreq: "monthly", priority: 0.5 },
  { path: "/error_codes",        changefreq: "monthly", priority: 0.5 },
  { path: "/suggest-an-api",     changefreq: "monthly", priority: 0.4 },
  { path: "/talk-to-sales",      changefreq: "monthly", priority: 0.4 },
  { path: "/private-deployment", changefreq: "monthly", priority: 0.4 },
  { path: "/status",             changefreq: "always",  priority: 0.4 },
  { path: "/privacy",            changefreq: "monthly", priority: 0.3 },
  { path: "/terms",              changefreq: "monthly", priority: 0.3 }
].freeze

CASE_STUDY_PAGES = [
  { path: "/case-studies", changefreq: "monthly", priority: 0.72 },
  { path: "/case-studies/verigeo", changefreq: "monthly", priority: 0.7 },
  { path: "/case-studies/compilestrength", changefreq: "monthly", priority: 0.7 }
].freeze

BLOG_POST_PAGES = BlogPost.all.map { |post| { path: "/blog/#{post.slug}", changefreq: "monthly", priority: 0.65 } }.freeze

DIVISION_MARKETING_PAGES = [
  { path: "/divisions", changefreq: "weekly", priority: 0.75 }
].concat(
  DivisionSlugs::ALL.map do |slug|
    { path: "/#{slug}", changefreq: "weekly", priority: 0.72 }
  end
).freeze

COMPARISON_PAGES = [
  { path: "/compare", changefreq: "weekly", priority: 0.75 }
].concat(
  ComparisonSlugs::ALL.map do |slug|
    { path: "/compare/#{slug}", changefreq: "monthly", priority: 0.68 }
  end
).freeze

INDUSTRY_PAGES = [
  { path: "/industries", changefreq: "weekly", priority: 0.75 }
].concat(
  IndustrySlugs::ALL.map do |slug|
    { path: "/industries/#{slug}", changefreq: "monthly", priority: 0.68 }
  end
).freeze

TOOL_PAGES = [
  { path: "/tools", changefreq: "weekly", priority: 0.75 }
].concat(
  ToolsController::SUPPORTED_TOOLS.map do |id|
    { path: "/tools/#{id}", changefreq: "monthly", priority: 0.68 }
  end
).freeze

locales = Rails.application.config.i18n.available_locales.map(&:to_s)

SitemapGenerator::Sitemap.create do # rubocop:disable Rails/SaveBang
  # Static pages, case studies, and division marketing pages
  group(filename: :sitemap_static) do
    STATIC_PAGES.each do |page|
      alts = locales.map { |l| { href: "https://requiemsapi.com/#{l}#{page[:path]}/", lang: l } }
      alts << { href: "https://requiemsapi.com/en#{page[:path]}/", lang: "x-default" }
      locales.each do |locale|
        locale_priority = locale == "en" ? page[:priority] : (page[:priority] * 0.4).round(1).clamp(0.1, 0.5)
        add "/#{locale}#{page[:path]}/",
          changefreq: page[:changefreq],
          priority:   locale_priority,
          alternates: alts
      end
    end

    CASE_STUDY_PAGES.each do |page|
      alts = locales.map { |l| { href: "https://requiemsapi.com/#{l}#{page[:path]}/", lang: l } }
      alts << { href: "https://requiemsapi.com/en#{page[:path]}/", lang: "x-default" }
      locales.each do |locale|
        locale_priority = locale == "en" ? page[:priority] : (page[:priority] * 0.4).round(1).clamp(0.1, 0.5)
        add "/#{locale}#{page[:path]}/",
          changefreq: page[:changefreq],
          priority:   locale_priority,
          alternates: alts
      end
    end

    BLOG_POST_PAGES.each do |page|
      alts = locales.map { |l| { href: "https://requiemsapi.com/#{l}#{page[:path]}/", lang: l } }
      alts << { href: "https://requiemsapi.com/en#{page[:path]}/", lang: "x-default" }
      locales.each do |locale|
        locale_priority = locale == "en" ? page[:priority] : (page[:priority] * 0.4).round(1).clamp(0.1, 0.5)
        add "/#{locale}#{page[:path]}/",
          changefreq: page[:changefreq],
          priority:   locale_priority,
          alternates: alts
      end
    end

    DIVISION_MARKETING_PAGES.each do |page|
      alts = locales.map { |l| { href: "https://requiemsapi.com/#{l}#{page[:path]}/", lang: l } }
      alts << { href: "https://requiemsapi.com/en#{page[:path]}/", lang: "x-default" }
      locales.each do |locale|
        locale_priority = locale == "en" ? page[:priority] : (page[:priority] * 0.4).round(1).clamp(0.1, 0.5)
        add "/#{locale}#{page[:path]}/",
          changefreq: page[:changefreq],
          priority:   locale_priority,
          alternates: alts
      end
    end

    COMPARISON_PAGES.each do |page|
      alts = locales.map { |l| { href: "https://requiemsapi.com/#{l}#{page[:path]}/", lang: l } }
      alts << { href: "https://requiemsapi.com/en#{page[:path]}/", lang: "x-default" }
      locales.each do |locale|
        locale_priority = locale == "en" ? page[:priority] : (page[:priority] * 0.4).round(1).clamp(0.1, 0.5)
        add "/#{locale}#{page[:path]}/",
          changefreq: page[:changefreq],
          priority:   locale_priority,
          alternates: alts
      end
    end

    INDUSTRY_PAGES.each do |page|
      alts = locales.map { |l| { href: "https://requiemsapi.com/#{l}#{page[:path]}/", lang: l } }
      alts << { href: "https://requiemsapi.com/en#{page[:path]}/", lang: "x-default" }
      locales.each do |locale|
        locale_priority = locale == "en" ? page[:priority] : (page[:priority] * 0.4).round(1).clamp(0.1, 0.5)
        add "/#{locale}#{page[:path]}/",
          changefreq: page[:changefreq],
          priority:   locale_priority,
          alternates: alts
      end
    end
  end

  # Tools catalog (index + individual tool pages)
  group(filename: :sitemap_tools) do
    TOOL_PAGES.each do |page|
      alts = locales.map { |l| { href: "https://requiemsapi.com/#{l}#{page[:path]}/", lang: l } }
      alts << { href: "https://requiemsapi.com/en#{page[:path]}/", lang: "x-default" }
      locales.each do |locale|
        locale_priority = locale == "en" ? page[:priority] : (page[:priority] * 0.4).round(1).clamp(0.1, 0.5)
        add "/#{locale}#{page[:path]}/",
          changefreq: page[:changefreq],
          priority:   locale_priority,
          alternates: alts
      end
    end
  end

  # System / engine pages
  group(filename: :sitemap_engines) do
    SYSTEM_SLUGS.each do |slug|
      alts = locales.map { |l| { href: "https://requiemsapi.com/#{l}/systems/#{slug}/", lang: l } }
      alts << { href: "https://requiemsapi.com/en/systems/#{slug}/", lang: "x-default" }
      locales.each do |locale|
        locale_priority = locale == "en" ? 0.8 : 0.3
        add "/#{locale}/systems/#{slug}/",
          changefreq: "monthly",
          priority:   locale_priority,
          alternates: alts
      end
    end
  end

  # Industry / category pages
  group(filename: :sitemap_categories) do
    categories.each do |cat_id|
      alts = locales.map { |l| { href: "https://requiemsapi.com/#{l}/categories/#{cat_id}/", lang: l } }
      alts << { href: "https://requiemsapi.com/en/categories/#{cat_id}/", lang: "x-default" }
      locales.each do |locale|
        locale_priority = locale == "en" ? 0.75 : 0.3
        add "/#{locale}/categories/#{cat_id}/",
          changefreq: "weekly",
          priority:   locale_priority,
          alternates: alts
      end
    end
  end

  # Example pages
  group(filename: :sitemap_examples) do
    examples.each do |example_id|
      alts = locales.map { |l| { href: "https://requiemsapi.com/#{l}/examples/#{example_id}/", lang: l } }
      alts << { href: "https://requiemsapi.com/en/examples/#{example_id}/", lang: "x-default" }
      locales.each do |locale|
        locale_priority = locale == "en" ? 0.6 : 0.2
        add "/#{locale}/examples/#{example_id}/",
          changefreq: "monthly",
          priority:   locale_priority,
          alternates: alts
      end
    end
  end

  # Individual API pages (largest group)
  group(filename: :sitemap_apis) do
    last_modified = Date.today

    live_apis.each do |api|
      alts = locales.map { |l| { href: "https://requiemsapi.com/#{l}/apis/#{api["id"]}/", lang: l } }
      alts << { href: "https://requiemsapi.com/en/apis/#{api["id"]}/", lang: "x-default" }
      locales.each do |locale|
        locale_priority = locale == "en" ? 0.8 : 0.3
        add "/#{locale}/apis/#{api["id"]}/",
          changefreq: "monthly",
          priority:   locale_priority,
          lastmod:    last_modified,
          alternates: alts
      end
    end
  end
end
