# frozen_string_literal: true

require "json"

module ApiDocs
  # Generates four-language code examples from an endpoint's YAML metadata.
  # Called by ApisHelper#api_documentation when no hand-written code_examples
  # key is present in the endpoint block.
  #
  # Usage:
  #   ApiDocs::SnippetGenerator.new(endpoint, base_url).call
  #   # => { "curl" => "...", "python" => "...", "javascript" => "...", "ruby" => "..." }
  #
  # Design decisions (see docs/plans/2026-08-03-api-docs-codegen-spec.md):
  #   - Prints response.json()["data"] as a whole object rather than cherry-picking
  #     a specific field. Hand-written examples cherry-pick, but replicating that
  #     heuristically requires per-endpoint knowledge the generator shouldn't need.
  #   - Supports one canonical call per endpoint. Endpoints with multiple examples
  #     (e.g. base64.yml showing both standard and url variants) keep their
  #     hand-written code_examples key as a permanent manual override.
  #   - response_kind: binary endpoints get a file-download template instead of
  #     the JSON-parse template.
  class SnippetGenerator
    def initialize(endpoint, base_url)
      @endpoint = endpoint
      @base_url  = base_url
    end

    def call
      {
        "curl"       => CurlSnippet.new(@endpoint, @base_url).call,
        "python"     => PythonSnippet.new(@endpoint, @base_url).call,
        "javascript" => JavascriptSnippet.new(@endpoint, @base_url).call,
        "ruby"       => RubySnippet.new(@endpoint, @base_url).call
      }
    end
  end

  # Shared helpers for all per-language snippet classes.
  class BaseSnippet
    API_KEY_PLACEHOLDER = "YOUR_API_KEY"

    def initialize(endpoint, base_url)
      @endpoint = endpoint
      @base_url  = base_url
    end

    def call
      raise NotImplementedError, "#{self.class}#call not implemented"
    end

    private

    def http_method  = @endpoint["method"]&.upcase || "GET"
    def path         = @endpoint["path"].to_s
    def binary?      = (@endpoint["response_kind"] || "json") == "binary"
    def params       = @endpoint["parameters"] || []
    def path_params  = params.select { |p| p["location"] == "path" }
    def query_params = params.select { |p| p["location"] == "query" }
    def body_params  = params.select { |p| p["location"] == "body" }
    def get?         = http_method == "GET"

    # Full URL with {name} path segments substituted by example values.
    def full_url
      url = "#{@base_url}#{path}"
      path_params.each { |p| url = url.gsub("{#{p["name"]}}", example_str(p)) }
      url
    end

    # Native Ruby value for a parameter's example or a type-appropriate placeholder.
    # YAML example values for array/object params are stored as JSON strings —
    # e.g. example: '["a", "b"]' — so we parse those back to native types.
    def native_example(param)
      raw = param["example"]
      return type_placeholder(param["type"].to_s) if raw.nil?
      return JSON.parse(raw) if raw.is_a?(String) && (raw.start_with?("[") || raw.start_with?("{"))

      raw
    end

    # Stringified version used for URL path/query substitution.
    def example_str(param)
      v = native_example(param)
      v.is_a?(String) ? v : v.to_s
    end

    # Body hash: { field_name => native_value } from location: body params.
    def body_hash
      body_params.each_with_object({}) { |p, h| h[p["name"]] = native_example(p) }
    end

    # Query string built from location: query params.
    def query_string
      return "" if query_params.empty?

      pairs = query_params.map { |p| "#{p["name"]}=#{URI.encode_www_form_component(example_str(p))}" }
      "?#{pairs.join("&")}"
    end

    # Type-appropriate placeholder when no example is given.
    def type_placeholder(type)
      case type
      when "string"        then "example"
      when "integer"       then 1
      when "number"        then 1.0
      when "boolean"       then true
      when "object"        then {}
      when "array<string>" then ["example"]
      when "array<object>" then [{}]
      when "array<number>" then [1.0]
      else "example"
      end
    end
  end

  # curl snippet — one-liner for GET, multi-line with -d for POST/PUT/PATCH.
  class CurlSnippet < BaseSnippet
    def call
      get? ? build_get : build_post
    end

    private

    def build_get
      url = "#{full_url}#{query_string}"
      if binary?
        [
          "curl \"#{url}\" \\",
          "  -H \"requiems-api-key: #{API_KEY_PLACEHOLDER}\" \\",
          "  --output response.png"
        ].join("\n")
      else
        [
          "curl \"#{url}\" \\",
          "  -H \"requiems-api-key: #{API_KEY_PLACEHOLDER}\""
        ].join("\n")
      end
    end

    def build_post
      body_json = JSON.generate(body_hash)
      [
        "curl -X #{http_method} \"#{full_url}\" \\",
        "  -H \"requiems-api-key: #{API_KEY_PLACEHOLDER}\" \\",
        "  -H \"Content-Type: application/json\" \\",
        "  -d '#{body_json}'"
      ].join("\n")
    end
  end

  # Python snippet using the requests library.
  class PythonSnippet < BaseSnippet
    def call
      lines = ["import requests", ""]
      lines << "url = \"#{full_url}\""

      if get?
        build_get(lines)
      else
        build_post(lines)
      end

      lines.join("\n")
    end

    private

    def build_get(lines)
      if query_params.any?
        q_hash = query_params.each_with_object({}) { |p, h| h[p["name"]] = native_example(p) }
        lines << "params = #{to_python(q_hash)}"
        lines << "headers = {\"requiems-api-key\": \"#{API_KEY_PLACEHOLDER}\"}"
        lines << "response = requests.get(url, headers=headers, params=params)"
      else
        lines << "headers = {\"requiems-api-key\": \"#{API_KEY_PLACEHOLDER}\"}"
        lines << "response = requests.get(url, headers=headers)"
      end

      append_response(lines)
    end

    def build_post(lines)
      lines << "headers = {"
      lines << "    \"requiems-api-key\": \"#{API_KEY_PLACEHOLDER}\","
      lines << "    \"Content-Type\": \"application/json\""
      lines << "}"

      if body_hash.any?
        lines << "payload = #{to_python(body_hash)}"
        lines << "response = requests.#{http_method.downcase}(url, headers=headers, json=payload)"
      else
        lines << "response = requests.#{http_method.downcase}(url, headers=headers)"
      end

      append_response(lines)
    end

    def append_response(lines)
      if binary?
        lines << "with open(\"response.png\", \"wb\") as f:"
        lines << "    f.write(response.content)"
      else
        lines << "print(response.json()[\"data\"])"
      end
    end

    def to_python(value)
      case value
      when Hash
        pairs = value.map { |k, v| "\"#{k}\": #{to_python(v)}" }
        "{ #{pairs.join(", ")} }"
      when Array then "[#{value.map { |v| to_python(v) }.join(", ")}]"
      when String then "\"#{value}\""
      when true   then "True"
      when false  then "False"
      when nil    then "None"
      else value.to_s
      end
    end
  end

  # JavaScript snippet using the Fetch API (async/await).
  class JavascriptSnippet < BaseSnippet
    def call
      get? ? build_get : build_post
    end

    private

    def build_get
      url = "#{full_url}#{query_string}"
      lines = [
        "const response = await fetch('#{url}', {",
        "  headers: { 'requiems-api-key': '#{API_KEY_PLACEHOLDER}' }",
        "});"
      ]

      if binary?
        lines << "const blob = await response.blob();"
        lines << "const imgUrl = URL.createObjectURL(blob);"
      else
        lines << "const { data } = await response.json();"
        lines << "console.log(data);"
      end

      lines.join("\n")
    end

    def build_post
      body_js = to_js(body_hash)
      lines = [
        "const response = await fetch('#{full_url}', {",
        "  method: '#{http_method}',",
        "  headers: {",
        "    'requiems-api-key': '#{API_KEY_PLACEHOLDER}',",
        "    'Content-Type': 'application/json'",
        "  },",
        "  body: JSON.stringify(#{body_js})",
        "});"
      ]
      lines << "const { data } = await response.json();"
      lines << "console.log(data);"
      lines.join("\n")
    end

    def to_js(value)
      case value
      when Hash
        pairs = value.map { |k, v| "#{k}: #{to_js(v)}" }
        "{ #{pairs.join(", ")} }"
      when Array then "[#{value.map { |v| to_js(v) }.join(", ")}]"
      when String then "'#{value}'"
      when true   then "true"
      when false  then "false"
      when nil    then "null"
      else value.to_s
      end
    end
  end

  # Ruby snippet using Net::HTTP from the standard library.
  class RubySnippet < BaseSnippet
    def call
      lines = ["require 'net/http'"]
      lines << "require 'json'" unless binary?
      lines << ""
      lines << "uri = URI('#{full_url}')"

      if get? && query_params.any?
        pairs = query_params.map { |p| "#{p["name"]}: '#{example_str(p)}'" }
        lines << "uri.query = URI.encode_www_form(#{pairs.join(", ")})"
      end

      lines << "request = Net::HTTP::#{http_method.capitalize}.new(uri)"
      lines << "request['requiems-api-key'] = '#{API_KEY_PLACEHOLDER}'"

      unless get?
        lines << "request['Content-Type'] = 'application/json'"
        lines << "request.body = #{to_ruby(body_hash)}.to_json" if body_hash.any?
      end

      lines << "response = Net::HTTP.start(uri.hostname, uri.port, use_ssl: true) { |http| http.request(request) }"

      if binary?
        lines << "File.write('response.png', response.body, mode: 'wb')"
      else
        lines << "puts JSON.parse(response.body)['data']"
      end

      lines.join("\n")
    end

    private

    def to_ruby(value)
      case value
      when Hash
        pairs = value.map { |k, v| "#{k}: #{to_ruby(v)}" }
        "{ #{pairs.join(", ")} }"
      when Array then "[#{value.map { |v| to_ruby(v) }.join(", ")}]"
      when String then "'#{value}'"
      when true   then "true"
      when false  then "false"
      when nil    then "nil"
      else value.to_s
      end
    end
  end
end
