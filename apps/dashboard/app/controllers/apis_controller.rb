# frozen_string_literal: true

class ApisController < ApplicationController
  include ApisHelper

  def index
    @categories = api_categories
    @popular_apis = popular_apis
    @apis_by_category = apis_grouped_by_category
    @categories_with_apis = categories_with_apis
    @search_query = params[:q] # Capture search query from URL
  end

  def show
    @api = find_api(params[:id])

    if @api.nil?
      redirect_to apis_path, alert: t("apis.flash.not_found")
      return
    end

    @category = find_category(Array(@api["categories"]).first)
    @categories = api_categories
    @documentation = api_documentation(params[:id])

    if @documentation.nil?
      redirect_to apis_path, alert: t("apis.flash.no_documentation")
      nil
    end
  end
end
