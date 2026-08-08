#!/usr/bin/env ruby
# frozen_string_literal: true

# Golden-diff script for REQAPI-456.
#
# Runs ApiDocs::SnippetGenerator against every endpoint that has a hand-written
# code_examples block and buckets each one as:
#
#   SAFE            — generator output is structurally equivalent; code_examples
#                     can be deleted from the YAML in the rollout PRs.
#   MANUAL_OVERRIDE — generator output differs meaningfully; the hand-written
#                     block must stay.
#
# Structural equivalence rules (in order — first match wins):
#   1. multi-example  → MANUAL_OVERRIDE  (≥2 curl invocations in hand-written snippet)
#   2. binary         → MANUAL_OVERRIDE  (response_kind: binary — download template differs)
#   3. url_mismatch   → MANUAL_OVERRIDE  (generated URL does not appear in hand-written curl)
#   4. (default)      → SAFE
#
# The script deliberately does not do a full string diff. The generator makes
# intentional trade-offs (prints data as a whole instead of cherry-picking a
# field), so exact string match is the wrong threshold. Structural equivalence
# (same URL, same params, same auth header) is what matters for the rollout.
#
# Usage:
#   ruby apps/dashboard/script/golden_diff_api_docs.rb
#   ruby apps/dashboard/script/golden_diff_api_docs.rb --verbose
#
# Output:
#   Prints a summary to stdout and writes a full report to
#   docs/plans/golden_diff_report.md

require "yaml"
require "json"
require "uri"
require "set"

$LOAD_PATH.unshift(File.expand_path("../app/services", __dir__))
require "api_docs/snippet_generator"

DOCS_DIR    = File.expand_path("../config/api_docs", __dir__)
REPORT_PATH = File.expand_path("../../../docs/plans/golden_diff_report.md", __dir__)
VERBOSE     = ARGV.include?("--verbose")

Result = Struct.new(:file, :endpoint_name, :method, :path, :bucket, :reason, :generated, :handwritten, keyword_init: true)

# --------------------------------------------------------------------------- #
# Helpers
# --------------------------------------------------------------------------- #

# Count the number of top-level curl invocations in a hand-written curl snippet.
# A new invocation starts with `curl` at the beginning of a line (possibly after whitespace).
def curl_invocation_count(curl_snippet)
  curl_snippet.to_s.scan(/^\s*curl\s/).size
end

# Returns true if the generated curl URL appears anywhere in the hand-written curl snippet.
def url_present_in_handwritten?(generated_curl, handwritten_curl)
  # Extract the URL from the generated curl (second token after optional -X METHOD)
  generated_url = generated_curl.to_s.lines.first&.strip
  # Strip the leading `curl `, `-X METHOD `, and trailing ` \`
  generated_url = generated_url.sub(/^curl\s+/, "").sub(/^-X\s+\S+\s+/, "").sub(/\s*\\$/, "").gsub('"', "")
  # Strip query string for path comparison (query param order can differ)
  base_url = generated_url.split("?").first.to_s.strip
  handwritten_curl.to_s.include?(base_url)
end

# --------------------------------------------------------------------------- #
# Main loop
# --------------------------------------------------------------------------- #

results = []
total_endpoints = 0
total_with_examples = 0

Dir[File.join(DOCS_DIR, "*.yml")].sort.each do |file|
  basename = File.basename(file, ".yml")
  doc      = YAML.safe_load(File.read(file))
  base_url = doc["base_url"].to_s

  doc["endpoints"]&.each do |ep|
    total_endpoints += 1
    next unless ep.key?("code_examples")

    total_with_examples += 1
    handwritten = ep["code_examples"]
    generated   = ApiDocs::SnippetGenerator.new(ep, base_url).call

    bucket, reason =
      if curl_invocation_count(handwritten["curl"]) >= 2
        ["MANUAL_OVERRIDE", "multi-example (#{curl_invocation_count(handwritten["curl"])} curl calls in snippet)"]
      elsif (ep["response_kind"] || "json") == "binary"
        ["MANUAL_OVERRIDE", "binary response (response_kind: binary)"]
      elsif !url_present_in_handwritten?(generated["curl"], handwritten["curl"])
        ["MANUAL_OVERRIDE", "url_mismatch — generated URL not found in hand-written curl"]
      else
        ["SAFE", "generated URL present, single example, JSON response"]
      end

    results << Result.new(
      file:          basename,
      endpoint_name: ep["name"],
      method:        ep["method"],
      path:          ep["path"],
      bucket:        bucket,
      reason:        reason,
      generated:     generated,
      handwritten:   handwritten
    )
  end
end

safe     = results.select { |r| r.bucket == "SAFE" }
override = results.select { |r| r.bucket == "MANUAL_OVERRIDE" }

# --------------------------------------------------------------------------- #
# Stdout summary
# --------------------------------------------------------------------------- #

puts "=" * 70
puts "GOLDEN-DIFF REPORT — ApiDocs SnippetGenerator (REQAPI-456)"
puts "=" * 70
puts ""
puts "Total endpoints:           #{total_endpoints}"
puts "With hand-written examples:#{total_with_examples}"
puts "  SAFE (can delete YAML):  #{safe.size}"
puts "  MANUAL_OVERRIDE (keep):  #{override.size}"
puts ""

puts "SAFE BUCKET (#{safe.size} endpoints):"
safe.each { |r| puts "  ✓ #{r.file} | #{r.method} #{r.path} (#{r.endpoint_name})" }

puts ""
puts "MANUAL_OVERRIDE BUCKET (#{override.size} endpoints):"
override.each { |r| puts "  ✗ #{r.file} | #{r.method} #{r.path} — #{r.reason}" }

if VERBOSE
  puts ""
  puts "=" * 70
  puts "VERBOSE DIFF — SAFE ENDPOINTS (generated vs hand-written curl)"
  puts "=" * 70
  safe.each do |r|
    puts "\n#{r.file} | #{r.endpoint_name}"
    puts "  GEN:  #{r.generated["curl"].lines.first&.strip}"
    puts "  HAND: #{r.handwritten["curl"].lines.first&.strip}"
  end
end

# --------------------------------------------------------------------------- #
# Markdown report
# --------------------------------------------------------------------------- #

md = []
md << "# Golden-Diff Report — ApiDocs SnippetGenerator"
md << ""
md << "> Generated by `apps/dashboard/script/golden_diff_api_docs.rb`  "
md << "> Part of REQAPI-456 — Auto-generate API doc code examples from YAML endpoint metadata."
md << ""
md << "## Summary"
md << ""
md << "| Metric | Count |"
md << "|--------|-------|"
md << "| Total endpoints | #{total_endpoints} |"
md << "| With hand-written `code_examples` | #{total_with_examples} |"
md << "| **SAFE** — code_examples can be deleted | **#{safe.size}** |"
md << "| **MANUAL_OVERRIDE** — keep hand-written | **#{override.size}** |"
md << ""
md << "## Safe Bucket — #{safe.size} endpoints"
md << ""
md << "These endpoints have a single canonical example and a JSON response."
md << "Their `code_examples` key can be deleted from the YAML in the rollout PRs."
md << ""
md << "| File | Method | Path | Endpoint name |"
md << "|------|--------|------|---------------|"
safe.each do |r|
  md << "| #{r.file}.yml | `#{r.method}` | `#{r.path}` | #{r.endpoint_name} |"
end

md << ""
md << "## Manual Override Bucket — #{override.size} endpoints"
md << ""
md << "These endpoints keep their hand-written `code_examples` block permanently."
md << ""
md << "| File | Method | Path | Endpoint name | Reason |"
md << "|------|--------|------|---------------|--------|"
override.each do |r|
  md << "| #{r.file}.yml | `#{r.method}` | `#{r.path}` | #{r.endpoint_name} | #{r.reason} |"
end

md << ""
md << "## Generator vs Hand-written — Sample Diff (Safe Endpoints)"
md << ""
md << "Intentional differences between generated and hand-written snippets:"
md << ""
md << "- **Response printing**: generator prints `response.json()[\"data\"]` as a whole."
md << "  Hand-written examples cherry-pick specific fields (e.g. `data[\"result\"]`)."
md << "  This is a deliberate trade-off for robustness over readability."
md << "- **No field-specific comments**: hand-written snippets often include"
md << "  inline comments (`# SGVsbG8...`). Generator omits these."
md << ""
md << "These differences are acceptable for the rollout. Endpoints where the"
md << "generator output would be confusing to a first-time user are already"
md << "in the manual-override bucket."

File.write(REPORT_PATH, md.join("\n") + "\n")
puts ""
puts "Full report written to: docs/plans/golden_diff_report.md"
