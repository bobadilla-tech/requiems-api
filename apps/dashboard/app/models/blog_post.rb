# frozen_string_literal: true

# File-backed blog post: content/blog/*.md, YAML frontmatter + Markdown body.
class BlogPost
  CONTENT_DIR = Rails.root.join("content", "blog")
  FRONTMATTER_RE = /\A---\s*\n(.*?)\n---\s*\n(.*)\z/m

  attr_reader :slug, :title, :date, :author, :description, :body

  def initialize(slug:, title:, date:, author:, description:, body:)
    @slug = slug
    @title = title
    @date = date
    @author = author
    @description = description
    @body = body
  end

  class << self
    def all
      @all ||= Dir.glob(CONTENT_DIR.join("*.md")).filter_map { |path| load_file(path) }
        .sort_by(&:date)
        .reverse
    end

    def find(slug)
      all.find { |post| post.slug == slug.to_s }
    end

    private

    def load_file(path)
      raw = File.read(path)
      match = FRONTMATTER_RE.match(raw)
      return nil unless match

      meta = YAML.safe_load(match[1], permitted_classes: [ Date ])
      new(
        slug: meta["slug"] || File.basename(path, ".md").sub(/\A\d{4}-\d{2}-\d{2}-/, ""),
        title: meta["title"],
        date: meta["date"].is_a?(Date) ? meta["date"] : Date.parse(meta["date"].to_s),
        author: meta["author"],
        description: meta["description"],
        body: match[2]
      )
    end
  end

  def to_html
    @to_html ||= Kramdown::Document.new(body, input: "GFM", hard_wrap: false).to_html
  end

  def reading_time_minutes
    [ (body.split.size / 200.0).ceil, 1 ].max
  end
end
