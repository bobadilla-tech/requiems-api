# frozen_string_literal: true

class ToolsController < ApplicationController
  SUPPORTED_TOOLS = %w[email-validator sentiment-analysis email-normalizer domain-checker quotes unit-conversion phone-validator inflation].freeze

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
    },
    "domain-checker" => {
      name: "Domain Checker",
      description: "Check domain availability, DNS records (A, MX, NS), and WHOIS data in one API call.",
      icon_classes: "bg-teal-50 dark:bg-teal-900/20 text-teal-600 dark:text-teal-400"
    },
    "phone-validator" => {
      name: "Phone Validator",
      description: "Validate phone numbers with carrier lookup, line type detection, and format normalization in one API call.",
      icon_classes: "bg-green-50 dark:bg-green-900/20 text-green-600 dark:text-green-400"
    },
    "unit-conversion" => {
      name: "Unit Conversion",
      description: "Convert length, weight, volume, temperature and other units instantly.",
      icon_classes: "bg-sky-50 dark:bg-sky-900/20 text-sky-600 dark:text-sky-400"
    },
    "quotes" => {
      name: "Random Quotes",
      description: "Fetch random inspirational quotes with author attribution in a single API call.",
      icon_classes: "bg-amber-50 dark:bg-amber-900/20 text-amber-600 dark:text-amber-400"
    },
    "inflation" => {
      name: "Inflation",
      description: "Historical and current CPI inflation rates for 241 countries, sourced from the World Bank.",
      icon_classes: "bg-orange-50 dark:bg-orange-900/20 text-orange-600 dark:text-orange-400"
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
