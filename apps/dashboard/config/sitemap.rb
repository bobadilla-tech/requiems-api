# frozen_string_literal: true

require "yaml"
require_relative "../lib/division_slugs.rb"

SitemapGenerator::Sitemap.default_host = "https://requiems.xyz"
SitemapGenerator::Sitemap.compress      = false # write sitemap.xml, not sitemap.xml.gz
SitemapGenerator::Sitemap.include_root  = false # all pages added manually
SitemapGenerator::Sitemap.include_index = false # single file, no index needed

catalog   = YAML.load_file(Rails.root.join("config", "api_catalog.yml"))
live_apis = catalog["apis"].select { |api| api["status"] == "live" }

STATIC_PAGES = [
  { path: "/",               changefreq: "weekly",  priority: 1.0 },
  { path: "/apis",           changefreq: "weekly",  priority: 0.9 },
  { path: "/pricing",        changefreq: "monthly", priority: 0.8 },
  { path: "/docs",           changefreq: "monthly", priority: 0.8 },
  { path: "/api_reference",  changefreq: "monthly", priority: 0.7 },
  { path: "/faq",            changefreq: "monthly", priority: 0.6 },
  { path: "/changelog",      changefreq: "weekly",  priority: 0.6 },
  { path: "/blog",           changefreq: "weekly",  priority: 0.6 },
  { path: "/about",          changefreq: "monthly", priority: 0.5 },
  { path: "/team",           changefreq: "monthly", priority: 0.5 },
  { path: "/contact",        changefreq: "monthly", priority: 0.5 },
  { path: "/privacy",        changefreq: "monthly", priority: 0.3 },
  { path: "/terms",          changefreq: "monthly", priority: 0.3 },
  { path: "/glossary",       changefreq: "monthly", priority: 0.5 },
  { path: "/error_codes",    changefreq: "monthly", priority: 0.5 },
  { path: "/status",         changefreq: "always",  priority: 0.4 },
  { path: "/examples",       changefreq: "weekly",  priority: 0.6 },
  { path: "/suggest-an-api", changefreq: "monthly", priority: 0.4 },
  { path: "/talk-to-sales",  changefreq: "monthly", priority: 0.4 }
].freeze

CASE_STUDY_PAGES = [
  { path: "/case-studies", changefreq: "monthly", priority: 0.72 },
  { path: "/case-studies/verigeo", changefreq: "monthly", priority: 0.7 },
  { path: "/case-studies/compilestrength", changefreq: "monthly", priority: 0.7 }
].freeze

DIVISION_MARKETING_PAGES = [
  { path: "/divisions", changefreq: "weekly", priority: 0.75 }
].concat(
  DivisionSlugs::ALL.map do |slug|
    { path: "/#{slug}", changefreq: "weekly", priority: 0.72 }
  end
).freeze

locales = Rails.application.config.i18n.available_locales.map(&:to_s)

SitemapGenerator::Sitemap.create do # rubocop:disable Rails/SaveBang
  STATIC_PAGES.each do |page|
    alts = locales.map { |l| { href: "https://requiems.xyz/#{l}#{page[:path]}", lang: l } }
    locales.each do |locale|
      add "/#{locale}#{page[:path]}",
        changefreq: page[:changefreq],
        priority:   page[:priority],
        alternates: alts
    end
  end

  DIVISION_MARKETING_PAGES.each do |page|
    alts = locales.map { |l| { href: "https://requiems.xyz/#{l}#{page[:path]}", lang: l } }
    locales.each do |locale|
      add "/#{locale}#{page[:path]}",
        changefreq: page[:changefreq],
        priority:   page[:priority],
        alternates: alts
    end
  end

  CASE_STUDY_PAGES.each do |page|
    alts = locales.map { |l| { href: "https://requiems.xyz/#{l}#{page[:path]}", lang: l } }
    locales.each do |locale|
      add "/#{locale}#{page[:path]}",
        changefreq: page[:changefreq],
        priority:   page[:priority],
        alternates: alts
    end
  end

  last_modified = Date.today

  live_apis.each do |api|
    alts = locales.map { |l| { href: "https://requiems.xyz/#{l}/apis/#{api["id"]}", lang: l } }
    locales.each do |locale|
      add "/#{locale}/apis/#{api["id"]}",
        changefreq: "monthly",
        priority:   0.8,
        lastmod:    last_modified,
        alternates: alts
    end
  end
end
