# frozen_string_literal: true

class IndustriesController < ApplicationController
  def index
    @industry_slugs = IndustrySlugs::ALL
  end

  def show
    @slug = params[:industry_slug].to_s
    head :not_found unless IndustrySlugs::ALL.include?(@slug)
  end
end
