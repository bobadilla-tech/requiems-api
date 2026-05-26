# frozen_string_literal: true

module ApisHelper
  def api_catalog
    @api_catalog ||= YAML.load_file(Rails.root.join("config", "api_catalog.yml"))
  end

  def api_categories
    api_catalog["categories"]
  end

  def all_apis
    api_catalog["apis"]
  end

  def apis_by_category(category_id)
    all_apis.select { |api| Array(api["categories"]).include?(category_id) }
  end

  def find_category(category_id)
    api_categories.find { |cat| cat["id"] == category_id }
  end

  def api_category_display_name(category)
    id = category["id"]
    if id.present? && DivisionSlugs::ALL.include?(id)
      t("divisions.index.cards.#{id}.title", default: category["name"])
    else
      category["name"]
    end
  end

  def find_api(api_id)
    all_apis.find { |api| api["id"] == api_id }
  end

  def live_apis
    all_apis.select { |api| api["status"] == "live" }
  end

  def api_status_variant(status)
    case status
    when "live"
      "success"
    when "beta"
      "warning"
    when "deprecated"
      "danger"
    else
      "default"
    end
  end

  def api_status_text(status)
    case status
    when "live"
      "Live"
    when "beta"
      "Beta"
    when "deprecated"
      "Deprecated"
    else
      status.to_s.titleize
    end
  end

  def api_documentation(api_id)
    doc_path = Rails.root.join("config", "api_docs", "#{api_id}.yml")

    return nil unless File.exist?(doc_path)

    @api_docs ||= {}
    @api_docs[api_id] ||= YAML.load_file(doc_path)
  end

  def popular_apis
    live_apis.select { |api| api["popular"] == true }
  end

  def apis_grouped_by_category
    live_apis.each_with_object({}) do |api, hash|
      Array(api["categories"]).each do |cat|
        (hash[cat] ||= []) << api
      end
    end
  end

  def categories_with_apis
    api_categories.select do |category|
      !category["coming_soon"] && apis_by_category(category["id"]).any?
    end
  end

  def api_documentation_as_text(doc)
    lines = []

    lines << doc["api_name"]
    lines << "=" * doc["api_name"].length
    lines << ""
    lines << doc["description"]
    lines << ""
    lines << "Base URL: #{doc["base_url"]}"
    lines << ""

    if doc["performance"].present?
      p = doc["performance"]
      lines << "Performance (#{p["samples"]} samples, measured #{p["measured_at"]})"
      lines << "  p50 (median): #{p["p50_ms"]} ms"
      lines << "  p95:          #{p["p95_ms"]} ms"
      lines << "  p99:          #{p["p99_ms"]} ms"
      lines << ""
    end

    doc["endpoints"]&.each do |ep|
      lines << "#{ep["method"]} #{ep["path"]}"
      lines << "-" * "#{ep["method"]} #{ep["path"]}".length
      lines << ep["description"] if ep["description"].present?
      lines << ""

      if ep["parameters"].present?
        lines << "Parameters:"
        ep["parameters"].each do |p|
          req = p["required"] ? " (required)" : " (optional)"
          lines << "  #{p["name"]} [#{p["type"]}]#{req} - #{p["description"]}"
        end
        lines << ""
      end

      if ep["response_example"].present?
        lines << "Response example:"
        lines << ep["response_example"].strip
        lines << ""
      end

      if ep["response_fields"].present?
        lines << "Response schema:"
        ep["response_fields"].each do |f|
          lines << "  #{f["name"]} [#{f["type"]}] - #{f["description"]}"
        end
        lines << ""
      end

      if ep["errors"].present?
        lines << "Errors:"
        ep["errors"].each do |e|
          lines << "  #{e["code"]} (#{e["status"]}) - #{e["description"]}"
        end
        lines << ""
      end
    end

    lines.join("\n")
  end

  def api_documentation_as_markdown(doc)
    lines = []
    lines << "# #{doc["api_name"]}"
    lines << ""
    lines << doc["description"]
    lines << ""
    lines << "**Base URL:** `#{doc["base_url"]}`"
    lines << ""

    if doc["performance"].present?
      p = doc["performance"]
      lines << "## Performance"
      lines << ""
      lines << "Measured against production with #{p["samples"]} samples (last updated #{p["measured_at"]})."
      lines << ""
      lines << "| Metric | Value |"
      lines << "|--------|-------|"
      lines << "| p50 (median) | #{p["p50_ms"]} ms |"
      lines << "| p95 | #{p["p95_ms"]} ms |"
      lines << "| p99 | #{p["p99_ms"]} ms |"
      lines << ""
    end

    doc["endpoints"]&.each do |ep|
      lines << "## `#{ep["method"]} #{ep["path"]}`"
      lines << ""
      lines << ep["description"] if ep["description"].present?
      lines << ""

      if ep["parameters"].present?
        lines << "### Parameters"
        lines << ""
        lines << "| Name | Type | Required | Description |"
        lines << "|------|------|----------|-------------|"
        ep["parameters"].each do |p|
          req = p["required"] ? "Yes" : "No"
          lines << "| `#{p["name"]}` | #{p["type"]} | #{req} | #{p["description"]} |"
        end
        lines << ""
      end

      if ep["request_example"].present?
        lines << "### Request"
        lines << ""
        lines << "```json"
        lines << ep["request_example"].strip
        lines << "```"
        lines << ""
      end

      if ep["response_example"].present?
        lines << "### Response Example"
        lines << ""
        lines << "```json"
        lines << ep["response_example"].strip
        lines << "```"
        lines << ""
      end

      if ep["response_fields"].present?
        lines << "### Response Schema"
        lines << ""
        lines << "| Field | Type | Description |"
        lines << "|-------|------|-------------|"
        ep["response_fields"].each do |f|
          lines << "| `#{f["name"]}` | #{f["type"]} | #{f["description"]} |"
        end
        lines << ""
      end

      if ep["errors"].present?
        lines << "### Errors"
        lines << ""
        lines << "| Code | Status | Description |"
        lines << "|------|--------|-------------|"
        ep["errors"].each do |e|
          lines << "| `#{e["code"]}` | #{e["status"]} | #{e["description"]} |"
        end
        lines << ""
      end

      next unless ep["code_examples"].present?

      lines << "### Code Examples"
      lines << ""
      ep["code_examples"].each do |lang, code|
        lines << "**#{lang.capitalize}**"
        lines << ""
        lines << "```#{lang}"
        lines << code.strip
        lines << "```"
        lines << ""
      end
    end

    lines.join("\n")
  end

  CATEGORY_ICON_SVGS = {
    "finance"        => '<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-6 h-6"><path stroke-linecap="round" stroke-linejoin="round" d="M2.25 8.25h19.5M2.25 9h19.5m-16.5 5.25h6m-6 2.25h3m-3.75 3h15a2.25 2.25 0 0 0 2.25-2.25V6.75A2.25 2.25 0 0 0 19.5 4.5h-15a2.25 2.25 0 0 0-2.25 2.25v10.5A2.25 2.25 0 0 0 4.5 19.5Z" /></svg>',
    "validation"     => '<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-6 h-6"><path stroke-linecap="round" stroke-linejoin="round" d="M9 12.75 11.25 15 15 9.75m-3-7.036A11.959 11.959 0 0 1 3.598 6 11.99 11.99 0 0 0 3 9.749c0 5.592 3.824 10.29 9 11.623 5.176-1.332 9-6.03 9-11.622 0-1.31-.21-2.571-.598-3.751h-.152c-3.196 0-6.1-1.248-8.25-3.285Z" /></svg>',
    "networking"     => '<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-6 h-6"><path stroke-linecap="round" stroke-linejoin="round" d="M12 21a9.004 9.004 0 0 0 8.716-6.747M12 21a9.004 9.004 0 0 1-8.716-6.747M12 21c2.485 0 4.5-4.03 4.5-9S14.485 3 12 3m0 18c-2.485 0-4.5-4.03-4.5-9S9.515 3 12 3m0 0a8.997 8.997 0 0 1 7.843 4.582M12 3a8.997 8.997 0 0 0-7.843 4.582m15.686 0A11.953 11.953 0 0 1 12 10.5c-2.998 0-5.74-1.1-7.843-2.918m15.686 0A8.959 8.959 0 0 1 21 12c0 .778-.099 1.533-.284 2.253m0 0A17.919 17.919 0 0 1 12 16.5c-3.162 0-6.133-.815-8.716-2.247m0 0A9.015 9.015 0 0 1 3 12c0-1.605.42-3.113 1.157-4.418" /></svg>',
    "places"         => '<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-6 h-6"><path stroke-linecap="round" stroke-linejoin="round" d="M15 10.5a3 3 0 1 1-6 0 3 3 0 0 1 6 0Z" /><path stroke-linecap="round" stroke-linejoin="round" d="M19.5 10.5c0 7.142-7.5 11.25-7.5 11.25S4.5 17.642 4.5 10.5a7.5 7.5 0 1 1 15 0Z" /></svg>',
    "text"           => '<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-6 h-6"><path stroke-linecap="round" stroke-linejoin="round" d="M19.5 14.25v-2.625a3.375 3.375 0 0 0-3.375-3.375h-1.5A1.125 1.125 0 0 1 13.5 7.125v-1.5a3.375 3.375 0 0 0-3.375-3.375H8.25m0 12.75h7.5m-7.5 3H12M10.5 2.25H5.625c-.621 0-1.125.504-1.125 1.125v17.25c0 .621.504 1.125 1.125 1.125h12.75c.621 0 1.125-.504 1.125-1.125V11.25a9 9 0 0 0-9-9Z" /></svg>',
    "technology"     => '<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-6 h-6"><path stroke-linecap="round" stroke-linejoin="round" d="M17.25 6.75 22.5 12l-5.25 5.25m-10.5 0L1.5 12l5.25-5.25m7.5-3-4.5 16.5" /></svg>',
    "entertainment"  => '<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-6 h-6"><path stroke-linecap="round" stroke-linejoin="round" d="m15.75 10.5 4.72-4.72a.75.75 0 0 1 1.28.53v11.38a.75.75 0 0 1-1.28.53l-4.72-4.72M4.5 18.75h9a2.25 2.25 0 0 0 2.25-2.25v-9a2.25 2.25 0 0 0-2.25-2.25h-9A2.25 2.25 0 0 0 2.25 7.5v9a2.25 2.25 0 0 0 2.25 2.25Z" /></svg>',
    "health"         => '<svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="1.5" stroke="currentColor" class="w-6 h-6"><path stroke-linecap="round" stroke-linejoin="round" d="M21 8.25c0-2.485-2.099-4.5-4.688-4.5-1.935 0-3.597 1.126-4.312 2.733-.715-1.607-2.377-2.733-4.313-2.733C5.1 3.75 3 5.765 3 8.25c0 7.22 9 12 9 12s9-4.78 9-12Z" /></svg>'
  }.freeze

  DIVISION_HERO_GRADIENTS = {
    "finance"       => "bg-gradient-to-br from-gray-950 via-amber-900 to-gray-900",
    "validation"    => "bg-gradient-to-br from-gray-950 via-green-900 to-gray-900",
    "networking"    => "bg-gradient-to-br from-gray-950 via-blue-900 to-slate-900",
    "places"        => "bg-gradient-to-br from-gray-950 via-emerald-900 to-gray-900",
    "text"          => "bg-gradient-to-br from-gray-950 via-purple-900 to-gray-900",
    "technology"    => "bg-gradient-to-br from-gray-950 via-cyan-900 to-slate-900",
    "entertainment" => "bg-gradient-to-br from-gray-950 via-orange-900 to-gray-900",
    "health"        => "bg-gradient-to-br from-gray-950 via-rose-900 to-gray-900"
  }.freeze

  CATEGORY_COLORS = {
    "finance"        => "bg-amber-100 dark:bg-amber-900 text-amber-600 dark:text-amber-300",
    "validation"     => "bg-green-100 dark:bg-green-900 text-green-600 dark:text-green-300",
    "networking"     => "bg-blue-100 dark:bg-blue-900 text-blue-600 dark:text-blue-300",
    "places"         => "bg-emerald-100 dark:bg-emerald-900 text-emerald-600 dark:text-emerald-300",
    "text"           => "bg-purple-100 dark:bg-purple-900 text-purple-600 dark:text-purple-300",
    "technology"     => "bg-cyan-100 dark:bg-cyan-900 text-cyan-600 dark:text-cyan-300",
    "entertainment"  => "bg-orange-100 dark:bg-orange-900 text-orange-500 dark:text-orange-300",
    "health"         => "bg-rose-100 dark:bg-rose-900 text-rose-600 dark:text-rose-300"
  }.freeze

  def category_icon_svg(category_id, size: "w-6 h-6")
    svg = CATEGORY_ICON_SVGS[category_id] || CATEGORY_ICON_SVGS["technology"]
    svg.gsub('class="w-6 h-6"', "class=\"#{size}\"").html_safe
  end

  def category_color_classes(category_id, hover: false)
    base = CATEGORY_COLORS[category_id] || CATEGORY_COLORS["technology"]
    return base unless hover

    hover_map = {
      "finance"        => "group-hover:bg-amber-200 dark:group-hover:bg-amber-800",
      "validation"     => "group-hover:bg-green-200 dark:group-hover:bg-green-800",
      "networking"     => "group-hover:bg-blue-200 dark:group-hover:bg-blue-800",
      "places"         => "group-hover:bg-emerald-200 dark:group-hover:bg-emerald-800",
      "text"           => "group-hover:bg-purple-200 dark:group-hover:bg-purple-800",
      "technology"     => "group-hover:bg-cyan-200 dark:group-hover:bg-cyan-800",
      "entertainment"  => "group-hover:bg-orange-200 dark:group-hover:bg-orange-800",
      "health"         => "group-hover:bg-rose-200 dark:group-hover:bg-rose-800"
    }
    "#{base} #{hover_map[category_id] || hover_map["technology"]}"
  end

  def division_hero_gradient(slug)
    DIVISION_HERO_GRADIENTS[slug] || "bg-gradient-to-br from-slate-900 via-blue-900 to-indigo-900"
  end

  def open_in_claude_url(documentation)
    doc_url = "#{request.base_url}/apis/#{documentation['api_id']}/index.md"
    prompt = "Read this page from the Requiems API docs: #{doc_url} and help me integrate this API into my project."
    "https://claude.ai/new?q=#{ERB::Util.url_encode(prompt)}"
  end

  def open_in_chatgpt_url(documentation)
    doc_url = "#{request.base_url}/apis/#{documentation['api_id']}/index.md"
    prompt = "Read this page from the Requiems API docs: #{doc_url} and help me integrate this API into my project."
    "https://chatgpt.com/?q=#{ERB::Util.url_encode(prompt)}"
  end
end
