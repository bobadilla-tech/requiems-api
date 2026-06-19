# frozen_string_literal: true

class ExamplesController < ApplicationController
  EXAMPLES_CONFIG = YAML.load_file(Rails.root.join("config", "examples.yml")).freeze

  def index
    @examples = examples_config["examples"]
    @categories = examples_config["categories"]
    @selected_category = params[:category]

    # Filter by category if specified
    if @selected_category.present?
      @examples = @examples.select { |example| example["category"] == @selected_category }
    end
  end

  def show
    @example = find_example(params[:id])

    if @example.nil?
      redirect_to examples_path, alert: t("examples.flash.not_found")
      return
    end

    @categories = examples_config["categories"]
  end

  private

  def examples_config
    EXAMPLES_CONFIG
  end

  def find_example(id)
    examples_config["examples"].find { |example| example["id"] == id }
  end
end
