# frozen_string_literal: true

class ToolsController < ApplicationController
  SUPPORTED_TOOLS = %w[email-validator sentiment-analysis email-normalizer domain-checker quotes unit-conversion phone-validator bin-lookup inflation qr-code profanity-filter useragent timezone trivia vpn-detection thesaurus spell-check random-user sudoku number-base-conversion mx-lookup mortgage markdown barcode advice base64 whois lorem-ipsum working-days world-time words].freeze

  # name/description come from I18n (tools.catalog.<id>.name / .description) —
  # only presentation data that isn't locale-dependent lives here.
  TOOLS_METADATA = {
    "email-validator" => {
      icon_classes: "bg-emerald-50 dark:bg-emerald-900/20 text-emerald-600 dark:text-emerald-400"
    },
    "bin-lookup" => {
      icon_classes: "bg-blue-50 dark:bg-blue-900/20 text-blue-600 dark:text-blue-400"
    },
    "sentiment-analysis" => {
      icon_classes: "bg-violet-50 dark:bg-violet-900/20 text-violet-600 dark:text-violet-400"
    },
    "email-normalizer" => {
      icon_classes: "bg-blue-50 dark:bg-blue-900/20 text-blue-600 dark:text-blue-400"
    },
    "domain-checker" => {
      icon_classes: "bg-teal-50 dark:bg-teal-900/20 text-teal-600 dark:text-teal-400"
    },
    "phone-validator" => {
      icon_classes: "bg-green-50 dark:bg-green-900/20 text-green-600 dark:text-green-400"
    },
    "unit-conversion" => {
      icon_classes: "bg-sky-50 dark:bg-sky-900/20 text-sky-600 dark:text-sky-400"
    },
    "quotes" => {
      icon_classes: "bg-amber-50 dark:bg-amber-900/20 text-amber-600 dark:text-amber-400"
    },
    "inflation" => {
      icon_classes: "bg-orange-50 dark:bg-orange-900/20 text-orange-600 dark:text-orange-400"
    },
    "qr-code" => {
      icon_classes: "bg-purple-50 dark:bg-purple-900/20 text-purple-600 dark:text-purple-400"
    },
    "profanity-filter" => {
      icon_classes: "bg-red-50 dark:bg-red-900/20 text-red-600 dark:text-red-400"
    },
    "useragent" => {
      icon_classes: "bg-teal-50 dark:bg-teal-900/20 text-teal-600 dark:text-teal-400"
    },
    "timezone" => {
      icon_classes: "bg-cyan-50 dark:bg-cyan-900/20 text-cyan-600 dark:text-cyan-400"
    },
    "trivia" => {
      icon_classes: "bg-pink-50 dark:bg-pink-900/20 text-pink-600 dark:text-pink-400"
    },
    "vpn-detection" => {
      icon_classes: "bg-indigo-50 dark:bg-indigo-900/20 text-indigo-600 dark:text-indigo-400"
    },
    "thesaurus" => {
      icon_classes: "bg-indigo-50 dark:bg-indigo-900/20 text-indigo-600 dark:text-indigo-400"
    },
    "spell-check" => {
      icon_classes: "bg-yellow-50 dark:bg-yellow-900/20 text-yellow-600 dark:text-yellow-400"
    },
    "random-user" => {
      icon_classes: "bg-cyan-50 dark:bg-cyan-900/20 text-cyan-600 dark:text-cyan-400"
    },
    "sudoku" => {
      icon_classes: "bg-cyan-50 dark:bg-cyan-900/20 text-cyan-600 dark:text-cyan-400"
    },
    "number-base-conversion" => {
      icon_classes: "bg-slate-50 dark:bg-slate-900/20 text-slate-600 dark:text-slate-400"
    },
    "mx-lookup" => {
      icon_classes: "bg-sky-50 dark:bg-sky-900/20 text-sky-600 dark:text-sky-400"
    },
    "mortgage" => {
      icon_classes: "bg-green-50 dark:bg-green-900/20 text-green-600 dark:text-green-400"
    },
    "markdown" => {
      icon_classes: "bg-blue-50 dark:bg-blue-900/20 text-blue-600 dark:text-blue-400"
    },
    "barcode" => {
      icon_classes: "bg-fuchsia-50 dark:bg-fuchsia-900/20 text-fuchsia-600 dark:text-fuchsia-400"
    },
    "advice" => {
      icon_classes: "bg-rose-50 dark:bg-rose-900/20 text-rose-600 dark:text-rose-400"
    },
    "base64" => {
      icon_classes: "bg-lime-50 dark:bg-lime-900/20 text-lime-600 dark:text-lime-400"
    },
    "whois" => {
      icon_classes: "bg-violet-50 dark:bg-violet-900/20 text-violet-600 dark:text-violet-400"
    },
    "lorem-ipsum" => {
      icon_classes: "bg-amber-50 dark:bg-amber-900/20 text-amber-600 dark:text-amber-400"
    },
    "working-days" => {
      icon_classes: "bg-emerald-50 dark:bg-emerald-900/20 text-emerald-600 dark:text-emerald-400"
    },
    "world-time" => {
      icon_classes: "bg-cyan-50 dark:bg-cyan-900/20 text-cyan-600 dark:text-cyan-400"
    },
    "words" => {
      icon_classes: "bg-pink-50 dark:bg-pink-900/20 text-pink-600 dark:text-pink-400"
    }
  }.freeze

  # icon_category maps to the category ids in ApisHelper::CATEGORY_ICON_SVGS / CATEGORY_COLORS,
  # reused here so tool categories carry the same iconography as the API directory.
  CATEGORIES = [
    {
      key: "validation",
      icon_category: "validation",
      tools: %w[email-validator email-normalizer phone-validator domain-checker]
    },
    {
      key: "text",
      icon_category: "text",
      tools: %w[sentiment-analysis profanity-filter thesaurus spell-check words]
    },
    {
      key: "network",
      icon_category: "networking",
      tools: %w[useragent vpn-detection timezone mx-lookup whois world-time]
    },
    {
      key: "finance",
      icon_category: "finance",
      tools: %w[bin-lookup inflation mortgage]
    },
    {
      key: "entertainment",
      icon_category: "entertainment",
      tools: %w[quotes trivia sudoku advice]
    },
    {
      key: "dev",
      icon_category: "technology",
      tools: %w[unit-conversion qr-code random-user number-base-conversion markdown barcode base64 lorem-ipsum working-days]
    }
  ].freeze

  def index
    @tools = SUPPORTED_TOOLS.map { |id| tool_data(id) }
    @categories = CATEGORIES.map do |cat|
      {
        key: cat[:key],
        icon_category: cat[:icon_category],
        tools: cat[:tools].filter_map { |id| tool_data(id) if TOOLS_METADATA[id] }
      }
    end
  end

  def show
    @tool_id = params[:id]

    unless SUPPORTED_TOOLS.include?(@tool_id)
      redirect_to root_path, alert: t("tools_controller.tool_not_found")
      return
    end

    render "tools/#{@tool_id.tr('-', '_')}/show"
  end

  private

  def tool_data(id)
    { id: id }.merge(TOOLS_METADATA[id]).merge(
      name: t("tools.catalog.#{id.tr('-', '_')}.name"),
      description: t("tools.catalog.#{id.tr('-', '_')}.description")
    )
  end
end
