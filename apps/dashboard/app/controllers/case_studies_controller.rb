# frozen_string_literal: true

class CaseStudiesController < ApplicationController
  helper DivisionsHelper

  SLUGS = %w[verigeo compilestrength].freeze

  def index
  end

  def show
    @slug = params[:slug].to_s
    unless SLUGS.include?(@slug)
      head :not_found
      return
    end
  end
end
