# frozen_string_literal: true

Rake::Task["sitemap:refresh"].enhance do
  src = Rails.root.join("public", "sitemap.xml")
  dst = Rails.root.join("public", "core-sitemap.xml")
  FileUtils.cp(src, dst)
end
