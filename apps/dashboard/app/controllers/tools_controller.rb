# frozen_string_literal: true

class ToolsController < ApplicationController
  SUPPORTED_TOOLS = %w[email-validator sentiment-analysis email-normalizer].freeze

  TOOLS_METADATA = {
    "email-validator" => {
      name: "Email Validator",
      description: "Syntax check, MX lookup, disposable detection, and typo correction in one call.",
      icon_classes: "bg-emerald-50 dark:bg-emerald-900/20 text-emerald-600 dark:text-emerald-400"
    },
    "sentiment-analysis" => {
      name: "Sentiment Analysis",
      description: "Classify text as positive, negative, or neutral with a confidence score.",
      icon_classes: "bg-violet-50 dark:bg-violet-900/20 text-violet-600 dark:text-violet-400"
    },
    "email-normalizer" => {
      name: "Email Normalizer",
      description: "Normalize, lowercase, trim, and canonicalize email addresses in one API call.",
      icon_classes: "bg-blue-50 dark:bg-blue-900/20 text-blue-600 dark:text-blue-400"
    }
  }.freeze

  def index
    @tools = SUPPORTED_TOOLS.map { |id| { id: id }.merge(TOOLS_METADATA[id]) }
  end

  def show
    @tool_id = params[:id]

    unless SUPPORTED_TOOLS.include?(@tool_id)
      redirect_to root_path, alert: "Tool not found."
      return
    end

    render "tools/#{@tool_id.tr('-', '_')}/show"
  end
end
