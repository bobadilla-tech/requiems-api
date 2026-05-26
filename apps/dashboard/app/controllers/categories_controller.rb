# frozen_string_literal: true

class CategoriesController < ApplicationController
  include ApisHelper

  def show
    @category = find_category(params[:id])

    if @category.nil?
      redirect_to apis_path, alert: t("categories.flash.not_found")
      return
    end

    @apis = apis_by_category(params[:id])
    @categories_with_apis = categories_with_apis
  end
end
