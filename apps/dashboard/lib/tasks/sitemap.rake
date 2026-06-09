# frozen_string_literal: true

# sitemap:refresh generates sitemap.xml (index) + sitemap_static.xml,
# sitemap_engines.xml, sitemap_categories.xml, sitemap_examples.xml, sitemap_apis.xml.
# No post-processing needed; core-sitemap.xml is retired.
Rake::Task["sitemap:refresh"].enhance do
end
