# frozen_string_literal: true

class CaseStudiesController < ApplicationController
  helper DivisionsHelper

  SLUGS = %w[verigeo compilestrength].freeze

  def index
  end

  def show
    @slug = params[:slug].to_s
    head :not_found unless SLUGS.include?(@slug)
  end
end
