# frozen_string_literal: true

class HomeController < ApplicationController
  API_CATALOG = YAML.load_file(Rails.root.join("config", "api_catalog.yml")).freeze

  def index
    categories = API_CATALOG["categories"]

    # Sort categories by priority: live first, then most important coming soon
    priority_order = %w[
      email text finance technology ai_vision data validation
      security conversion health entertainment places transportation
      animals tax misc
    ]

    @categories = categories.sort_by do |cat|
      [ cat["coming_soon"] ? 1 : 0, priority_order.index(cat["id"]) || 999 ]
    end
  end

  def docs
  end

  def pricing
  end

  def blog
  end

  def status
  end

  def glossary
  end

  def error_codes
  end

  def faq
  end

  def team
  end

  def for_llms
  end

  def domain_checker
  end
end
