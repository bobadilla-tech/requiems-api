# frozen_string_literal: true

class SystemsController < ApplicationController
  SYSTEMS = [
    {
      slug: "identity-risk",
      icon: "shield",
      color: "blue",
      division_slugs: %w[validation networking text]
    },
    {
      slug: "payments-intelligence",
      icon: "credit-card",
      color: "green",
      division_slugs: %w[finance networking]
    },
    {
      slug: "global-data",
      icon: "globe",
      color: "purple",
      division_slugs: %w[places]
    },
    {
      slug: "data-integrity",
      icon: "check-circle",
      color: "indigo",
      division_slugs: %w[validation text]
    },
    {
      slug: "developer-utilities",
      icon: "wrench",
      color: "gray",
      division_slugs: %w[technology entertainment health]
    }
  ].freeze

  def index
    @systems = SYSTEMS
  end
end
