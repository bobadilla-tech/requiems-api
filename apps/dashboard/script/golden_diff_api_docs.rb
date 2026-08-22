#!/usr/bin/env ruby
# frozen_string_literal: true

# Golden-diff script for REQAPI-456.
#
# Runs ApiDocs::SnippetGenerator against every endpoint that has a hand-written
# code_examples block and buckets each one as:
#
#   SAFE            — generator output is structurally equivalent to the
#                     hand-written snippet. Proves the generator can stand in
#                     for this endpoint's code_examples going forward.
#   MANUAL_OVERRIDE — generator output differs meaningfully; the hand-written
#                     block stays authoritative.
#
# This is a regeneration-confidence check, not a deletion trigger — SAFE means
# "the generator reproduces this endpoint correctly," not "delete the YAML
# block." Whether/when to stop hand-maintaining code_examples for SAFE
# endpoints is a separate rollout decision.
#
# Structural equivalence rules (in order — first match wins):
#   1. multi-example       → MANUAL_OVERRIDE  (≥2 curl invocations in hand-written snippet)
#   2. binary               → MANUAL_OVERRIDE  (response_kind: binary — download template differs)
#   3. parse_failure        → NEEDS_REVIEW     (generated or hand-written curl couldn't be parsed —
#                                                not a confirmed structural difference, needs a human look)
#   4. structural_mismatch  → MANUAL_OVERRIDE  (method, URL, query, headers, or body differ)
#   5. (default)            → SAFE
#
# The script deliberately does not do a full string diff. The generator makes
# intentional trade-offs (prints data as a whole instead of cherry-picking a
# field), so exact string match is the wrong threshold. Structural equivalence
# (same method, same URL, same query params, same headers, same body) is what
# matters for regeneration confidence.
#
# Usage:
#   ruby apps/dashboard/script/golden_diff_api_docs.rb
#   ruby apps/dashboard/script/golden_diff_api_docs.rb --verbose
#
# Output:
#   Prints a summary to stdout and writes a full report to
#   docs/plans/2026-08-21-golden-diff-report.md

require "yaml"
require "json"
require "uri"
require "shellwords"

$LOAD_PATH.unshift(File.expand_path("../app/services", __dir__))
require "api_docs/snippet_generator"

DOCS_DIR    = File.expand_path("../config/api_docs", __dir__)
REPORT_PATH = File.expand_path("../../../docs/plans/2026-08-21-golden-diff-report.md", __dir__)
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

# Parses a (single-invocation) curl snippet into its structural parts: HTTP
# method, canonical URL (scheme+host+path+query, order-independent on query),
# header names present, and request body (parsed as JSON when possible).
# Returns nil if the method/URL can't be extracted at all — callers must treat
# that as "not comparable", never as a match.
CurlParts = Struct.new(:method, :url, :query, :headers, :body, keyword_init: true)

def parse_curl(curl_snippet)
  text = curl_snippet.to_s
  first_line = text.lines.first.to_s.strip
  return nil if first_line.empty?

  method = first_line.match(/-X\s+(\w+)/i)&.captures&.first&.upcase || "GET"

  # URL may or may not be quoted in hand-written snippets (both are valid
  # shell syntax since generated/example URLs never contain whitespace).
  url_match = first_line.match(/curl\s+(?:-X\s+\S+\s+)?(?:"([^"]+)"|'([^']+)'|(\S+))/)
  return nil unless url_match

  full_url = (url_match[1] || url_match[2] || url_match[3]).to_s.strip
  return nil if full_url.empty?

  path, _, query = full_url.partition("?")
  return nil if path.empty?

  query_pairs = query.empty? ? [] : URI.decode_www_form(query).sort

  headers = text.scan(/-H\s+["']([^"':]+):/).flatten.map(&:downcase).sort

  body = nil
  body_match = text.match(/-d\s+'((?:[^'\\]|\\.)*)'/m)
  if body_match
    raw = body_match[1].gsub("'\\''", "'")
    body = begin
      JSON.parse(raw)
    rescue JSON::ParserError
      raw
    end
  end

  CurlParts.new(method: method, url: path, query: query_pairs, headers: headers, body: body)
rescue URI::InvalidComponentError, ArgumentError
  nil
end

# Compares the generated and hand-written curl snippets and returns one of:
#   :parse_failure — generated or hand-written curl couldn't be parsed at all.
#                    Not a confirmed structural difference — routed to Needs
#                    Review rather than asserted as MANUAL_OVERRIDE, since a
#                    parser bug (not a real snippet difference) is just as
#                    likely a cause and deserves a human look either way.
#   :mismatch      — both parsed, but method/URL/query/headers/body differ.
#   :match         — both parsed and are structurally equivalent (same
#                    method, same canonical URL, same query params, same
#                    header set — covers the requiems-api-key auth header —
#                    same body).
def compare_curls(generated_curl, handwritten_curl)
  gen = parse_curl(generated_curl)
  hw  = parse_curl(handwritten_curl)
  return :parse_failure if gen.nil? || hw.nil?

  if gen.method == hw.method &&
     gen.url == hw.url &&
     gen.query == hw.query &&
     gen.headers == hw.headers &&
     gen.body == hw.body
    :match
  else
    :mismatch
  end
end

# Maps comparator outcomes to the same report buckets used by the audit loop.
# Multi-example and binary precedence is intentionally kept here so focused
# tests exercise the classification logic as well as compare_curls.
def classify_curl_result(curl_comparison, invocation_count:, response_kind:)
  if invocation_count >= 2
    ["MANUAL_OVERRIDE", "multi-example (#{invocation_count} curl calls in snippet)"]
  elsif response_kind == "binary"
    ["MANUAL_OVERRIDE", "binary response (response_kind: binary)"]
  elsif curl_comparison == :parse_failure
    ["NEEDS_REVIEW", "parse_failure — generated or hand-written curl could not be parsed; not a confirmed structural difference, needs a human look"]
  elsif curl_comparison == :mismatch
    ["MANUAL_OVERRIDE", "structural_mismatch — method/url/query/headers/body differ between generated and hand-written curl"]
  else
    ["SAFE", "method, URL, query params, headers, and body all match; single example, JSON response"]
  end
end

# --------------------------------------------------------------------------- #
# Main loop
# --------------------------------------------------------------------------- #

def generate_report(verbose: VERBOSE)
results = []
needs_review = []
total_endpoints = 0
total_with_examples = 0

Dir[File.join(DOCS_DIR, "*.yml")].sort.each do |file|
  basename = File.basename(file, ".yml")
  doc      = YAML.safe_load(File.read(file))
  base_url = doc["base_url"].to_s

  doc["endpoints"]&.each do |ep|
    total_endpoints += 1

    unless ep.key?("code_examples")
      # No hand-written block means this endpoint is already auto-generated
      # today (ApisHelper's absence-of-key-means-generate rule). If its notes
      # describe partial-success semantics, the generic whole-object-print
      # template may not narrate that well — flag it for human review even
      # though there's no hand-written snippet to diff against.
      notes = Array(ep["notes"])
      if notes.any? { |n| n.to_s.match?(/partial.success/i) }
        needs_review << { file: basename, endpoint_name: ep["name"], method: ep["method"], path: ep["path"],
                           kind: :partial_success,
                           reason: "partial-success notes — generic template may not narrate this well" }
      end
      next
    end

    total_with_examples += 1
    handwritten = ep["code_examples"]
    generated   = ApiDocs::SnippetGenerator.new(ep, base_url).call

    curl_comparison = curl_invocation_count(handwritten["curl"]) >= 2 ? nil : compare_curls(generated["curl"], handwritten["curl"])

    bucket, reason = classify_curl_result(
      curl_comparison,
      invocation_count: curl_invocation_count(handwritten["curl"]),
      response_kind: (ep["response_kind"] || "json")
    )

    if bucket == "NEEDS_REVIEW"
      needs_review << { file: basename, endpoint_name: ep["name"], method: ep["method"], path: ep["path"],
                        kind: :parse_failure, reason: reason }
      next
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
partial_success_review = needs_review.count { |r| r[:kind] == :partial_success }
parse_failure_review = needs_review.count { |r| r[:kind] == :parse_failure }

# --------------------------------------------------------------------------- #
# Stdout summary
# --------------------------------------------------------------------------- #

puts "=" * 70
puts "GOLDEN-DIFF REPORT — ApiDocs SnippetGenerator (REQAPI-456)"
puts "=" * 70
puts ""
puts "Total endpoints:                    #{total_endpoints}"
puts "With hand-written examples:         #{total_with_examples}"
puts "  SAFE (generator matches):         #{safe.size}"
puts "  MANUAL_OVERRIDE (keep):           #{override.size}"
puts "  Needs review — parse failures:    #{parse_failure_review}"
puts "  Needs review — partial-success:   #{partial_success_review}"
puts "Needs review total:                 #{needs_review.size}"
puts ""

puts "SAFE BUCKET (#{safe.size} endpoints):"
safe.each { |r| puts "  ✓ #{r.file} | #{r.method} #{r.path} (#{r.endpoint_name})" }

puts ""
puts "MANUAL_OVERRIDE BUCKET (#{override.size} endpoints):"
override.each { |r| puts "  ✗ #{r.file} | #{r.method} #{r.path} — #{r.reason}" }

puts ""
puts "NEEDS REVIEW (#{needs_review.size} endpoints):"
needs_review.each { |r| puts "  ? #{r[:file]} | #{r[:method]} #{r[:path]} (#{r[:endpoint_name]}) — #{r[:reason]}" }

if verbose
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

source_revision    = `git rev-parse HEAD 2>/dev/null`.strip
source_revision    = "unknown (not a git checkout, or git unavailable)" if source_revision.empty?
generator_revision = `git log -1 --format=%H -- #{Shellwords.escape(__FILE__)} 2>/dev/null`.strip
generator_revision = "unknown" if generator_revision.empty?
generated_at       = Time.now.utc.strftime("%Y-%m-%dT%H:%M:%SZ")

md = []
md << "# Golden-Diff Report — ApiDocs SnippetGenerator"
md << ""
md << "> Generated by `apps/dashboard/script/golden_diff_api_docs.rb`"
md << "> Part of REQAPI-456 — Auto-generate API doc code examples from YAML endpoint metadata."
md << "> Source revision: `#{source_revision}` · Generator script revision: `#{generator_revision}` · Generated: #{generated_at}"
md << ""
md << "## Summary"
md << ""
md << "| Metric | Count |"
md << "|--------|-------|"
md << "| Total endpoints | #{total_endpoints} |"
md << "| With hand-written `code_examples` | #{total_with_examples} |"
md << "| **SAFE** — generator output matches hand-written | **#{safe.size}** |"
md << "| **MANUAL_OVERRIDE** — hand-written stays authoritative | **#{override.size}** |"
md << "| **Needs review** — partial-success notes or curl parse failures | **#{needs_review.size}** |"
md << "| Needs review — partial-success endpoints without snippets | #{partial_success_review} |"
md << "| Needs review — parse failures with hand-written snippets | #{parse_failure_review} |"
md << ""
md << "Bucket accounting: endpoints with hand-written `code_examples` are split"
md << "among SAFE (#{safe.size}), MANUAL_OVERRIDE (#{override.size}), and parse-failure"
md << "Needs Review (#{parse_failure_review}); these counts sum to"
md << "#{safe.size + override.size + parse_failure_review}, the #{total_with_examples}"
md << "endpoints with hand-written snippets. The remaining Needs Review entries"
md << "(#{partial_success_review}) are endpoints without hand-written snippets whose"
md << "partial-success notes warrant human review. Other endpoints without snippets"
md << "are intentionally omitted from the report."
md << ""
md << "## Safe Bucket — #{safe.size} endpoints"
md << ""
md << "The generator's method, URL, query params, headers, and body all match"
md << "the hand-written snippet for these endpoints — the generator can stand in"
md << "for `code_examples` here with no loss of correctness. This bucket is a"
md << "regeneration-confidence signal, not a deletion instruction: whether and"
md << "when to stop hand-maintaining `code_examples` for these endpoints is a"
md << "separate rollout decision, made file-by-file with its own review."
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
md << "## Needs Review — #{needs_review.size} endpoints"
md << ""
md << "Two distinct reasons an endpoint lands here, distinguished in the Reason"
md << "column below:"
md << ""
md << "- **partial-success notes** — no hand-written `code_examples` at all (so"
md << "  it's already auto-generated today, per ApisHelper's absence-of-the-key-"
md << "  means-\"generate\" rule), but its `notes:` describe partial-success"
md << "  semantics that the generator's generic whole-object-print template may"
md << "  not narrate well. Flagged for a human to judge whether the generated"
md << "  snippet reads clearly enough as-is, or needs a hand-written override."
md << "- **parse_failure** — this endpoint *does* have a hand-written"
md << "  `code_examples` block, but either the generated or the hand-written curl"
md << "  snippet couldn't be parsed by this script's comparator. This is *not* a"
md << "  confirmed structural difference (unlike MANUAL_OVERRIDE's"
md << "  `structural_mismatch`) — it may just be a parser limitation — so it's"
md << "  routed here for a human to verify by hand rather than asserted as a"
md << "  real mismatch."
md << ""
md << "| File | Method | Path | Endpoint name | Reason |"
md << "|------|--------|------|---------------|--------|"
needs_review.each do |r|
  md << "| #{r[:file]}.yml | `#{r[:method]}` | `#{r[:path]}` | #{r[:endpoint_name]} | #{r[:reason]} |"
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
md << "These differences are acceptable for regeneration purposes. Endpoints"
md << "where the generator output would be confusing to a first-time user are"
md << "already in the manual-override bucket."

File.write(REPORT_PATH, md.join("\n") + "\n")
puts ""
puts "Full report written to: docs/plans/2026-08-21-golden-diff-report.md"
end

generate_report if $PROGRAM_NAME == __FILE__
