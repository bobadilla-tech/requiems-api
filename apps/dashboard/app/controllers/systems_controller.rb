# frozen_string_literal: true

class SystemsController < ApplicationController
  include ApisHelper

  SYSTEMS = [
    {
      slug: "identity-risk",
      color: "blue",
      division_slugs: %w[validation networking text]
    },
    {
      slug: "payments-intelligence",
      color: "green",
      division_slugs: %w[finance networking]
    },
    {
      slug: "global-data",
      color: "purple",
      division_slugs: %w[places]
    },
    {
      slug: "data-integrity",
      color: "indigo",
      division_slugs: %w[validation text]
    },
    {
      slug: "developer-utilities",
      color: "gray",
      division_slugs: %w[technology entertainment health]
    }
  ].freeze

  SYSTEM_SLUGS = SYSTEMS.map { |s| s[:slug] }.freeze

  def index
    @systems = SYSTEMS
  end

  def show
    @slug = params[:system_slug]
    @system = SYSTEMS.find { |s| s[:slug] == @slug }
    head :not_found and return unless @system
    @documentation = api_documentation(@slug)
  end
end
