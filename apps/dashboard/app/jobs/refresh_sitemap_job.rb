# frozen_string_literal: true

# Regenerates sitemap.xml and its per-section files (see lib/tasks/sitemap.rake).
#
# Runs nightly via Sidekiq Cron so new blog posts, APIs, and case studies are
# picked up automatically instead of requiring a manual `rake sitemap:refresh`.
#
class RefreshSitemapJob < ApplicationJob
  queue_as :default

  def perform
    Rails.application.load_tasks unless Rake::Task.task_defined?("sitemap:refresh")
    Rake::Task["sitemap:refresh"].execute
    Rails.logger.info "[RefreshSitemapJob] Sitemap regenerated"
  rescue StandardError => e
    Rails.logger.error "[RefreshSitemapJob] Failed to regenerate sitemap: #{e.message}"
    raise
  end
end
