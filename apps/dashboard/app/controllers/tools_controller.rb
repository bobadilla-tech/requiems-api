# frozen_string_literal: true

class ToolsController < ApplicationController
  SUPPORTED_TOOLS = %w[email-validator sentiment-analysis email-normalizer domain-checker quotes unit-conversion phone-validator bin-lookup inflation qr-code profanity-filter useragent timezone trivia vpn-detection thesaurus spell-check random-user sudoku number-base-conversion mx-lookup mortgage].freeze

  TOOLS_METADATA = {
    "email-validator" => {
      name: "Email Validator",
      description: "Syntax check, MX lookup, disposable detection, and typo correction in one call.",
      icon_classes: "bg-emerald-50 dark:bg-emerald-900/20 text-emerald-600 dark:text-emerald-400"
    },
    "bin-lookup" => {
      name: "BIN Lookup",
      description: "Card network, type, issuing bank, and country from the first 6–8 digits of any payment card.",
      icon_classes: "bg-blue-50 dark:bg-blue-900/20 text-blue-600 dark:text-blue-400"
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
    },
    "qr-code" => {
      name: "QR Code Generator",
      description: "Generate scannable QR codes from any text or URL — raw PNG or base64 JSON, configurable size and error correction.",
      icon_classes: "bg-purple-50 dark:bg-purple-900/20 text-purple-600 dark:text-purple-400"
    },
    "profanity-filter" => {
      name: "Profanity Filter",
      description: "Detect and censor offensive language in any text. Returns flagged words and a clean censored version.",
      icon_classes: "bg-red-50 dark:bg-red-900/20 text-red-600 dark:text-red-400"
    },
    "useragent" => {
      name: "User Agent Parser",
      description: "Parse user agent strings to extract browser, OS, device type, and bot detection.",
      icon_classes: "bg-teal-50 dark:bg-teal-900/20 text-teal-600 dark:text-teal-400"
    },
    "timezone" => {
      name: "Timezone",
      description: "IANA timezone, UTC offset, current time, and DST status for any city or coordinates.",
      icon_classes: "bg-cyan-50 dark:bg-cyan-900/20 text-cyan-600 dark:text-cyan-400"
    },
    "trivia" => {
      name: "Trivia",
      description: "Random trivia questions with multiple-choice answers, filterable by category and difficulty.",
      icon_classes: "bg-pink-50 dark:bg-pink-900/20 text-pink-600 dark:text-pink-400"
    },
    "vpn-detection" => {
      name: "VPN & Proxy Detection",
      description: "Detect if an IP address belongs to a VPN, proxy, Tor exit node, or hosting provider, with threat and fraud scoring.",
      icon_classes: "bg-indigo-50 dark:bg-indigo-900/20 text-indigo-600 dark:text-indigo-400"
    },
    "thesaurus" => {
      name: "Thesaurus",
      description: "Find synonyms and antonyms for any word to enhance vocabulary and writing.",
      icon_classes: "bg-indigo-50 dark:bg-indigo-900/20 text-indigo-600 dark:text-indigo-400"
    },
    "spell-check" => {
      name: "Spell Check",
      description: "Check spelling and get correction suggestions for misspelled words in any text.",
      icon_classes: "bg-yellow-50 dark:bg-yellow-900/20 text-yellow-600 dark:text-yellow-400"
    },
    "random-user" => {
      name: "Random User",
      description: "Generate random fake user profiles for testing and prototyping — names, emails, phone numbers, addresses, and avatars.",
      icon_classes: "bg-cyan-50 dark:bg-cyan-900/20 text-cyan-600 dark:text-cyan-400"
    },
    "sudoku" => {
      name: "Sudoku",
      description: "Generate Sudoku puzzles with solutions across multiple difficulty levels.",
      icon_classes: "bg-cyan-50 dark:bg-cyan-900/20 text-cyan-600 dark:text-cyan-400"
    },
    "number-base-conversion" => {
      name: "Number Base Conversion",
      description: "Convert integers between binary, octal, decimal, and hexadecimal, with optional 0x/0b/0o prefixes.",
      icon_classes: "bg-slate-50 dark:bg-slate-900/20 text-slate-600 dark:text-slate-400"
    },
    "mx-lookup" => {
      name: "MX Lookup",
      description: "Look up MX records for any domain — mail server hostnames and priorities, sorted by delivery preference.",
      icon_classes: "bg-sky-50 dark:bg-sky-900/20 text-sky-600 dark:text-sky-400"
    },
    "mortgage" => {
      name: "Mortgage Calculator",
      description: "Calculate monthly payments, total interest, and full amortization schedules for any fixed-rate mortgage.",
      icon_classes: "bg-green-50 dark:bg-green-900/20 text-green-600 dark:text-green-400"
    }
  }.freeze

  def index
    @tools = SUPPORTED_TOOLS.map { |id| { id: id }.merge(TOOLS_METADATA[id]) }
  end

  def show
    @tool_id = params[:id]

    unless SUPPORTED_TOOLS.include?(@tool_id)
      redirect_to root_path, alert: t("tools_controller.tool_not_found")
      return
    end

    render "tools/#{@tool_id.tr('-', '_')}/show"
  end
end
