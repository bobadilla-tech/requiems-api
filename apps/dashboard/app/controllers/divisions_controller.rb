# frozen_string_literal: true

class DivisionsController < ApplicationController
  include ApisHelper

  def index
    @division_slugs = DivisionSlugs::ALL
    @categories = @division_slugs.filter_map { |id| find_category(id) }
  end

  def show
    @slug = params[:division_slug].to_s
    unless DivisionSlugs::ALL.include?(@slug)
      head :not_found
      return
    end

    @category = find_category(@slug)
  end
end
