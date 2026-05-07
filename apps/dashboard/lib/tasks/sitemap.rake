# frozen_string_literal: true

# Pretty-print sitemap.xml after sitemap_generator writes (it minifies to one line).
Rake::Task["sitemap:refresh"].enhance do
  require "rexml/document"

  path = Rails.root.join("public", "sitemap.xml")
  next unless path.exist?

  doc = REXML::Document.new(path.read)
  fmt = REXML::Formatters::Pretty.new(2)
  fmt.compact = true
  out = +""
  fmt.write(doc, out)

  path.write("#{out}\n")
  puts "sitemap: generated #{path}"
end
