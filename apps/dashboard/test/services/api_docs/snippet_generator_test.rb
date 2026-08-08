# frozen_string_literal: true

require "test_helper"

class ApiDocs::SnippetGeneratorTest < ActiveSupport::TestCase
  BASE_URL = "https://api.requiems.xyz"

  # ---------------------------------------------------------------------------
  # Helpers
  # ---------------------------------------------------------------------------

  def endpoint(overrides = {})
    {
      "name"        => "Test Endpoint",
      "method"      => "GET",
      "path"        => "/v1/test",
      "description" => "A test endpoint",
      "parameters"  => []
    }.merge(overrides)
  end

  def snippets_for(ep)
    ApiDocs::SnippetGenerator.new(ep, BASE_URL).call
  end

  # ---------------------------------------------------------------------------
  # SnippetGenerator#call — shape
  # ---------------------------------------------------------------------------

  test "call returns all four languages" do
    result = snippets_for(endpoint)
    assert_equal %w[curl python javascript ruby], result.keys
    result.values.each { |v| assert_kind_of String, v }
  end

  # ---------------------------------------------------------------------------
  # GET — no params
  # ---------------------------------------------------------------------------

  test "curl GET with no params includes URL and auth header" do
    curl = snippets_for(endpoint)["curl"]
    assert_includes curl, "#{BASE_URL}/v1/test"
    assert_includes curl, "requiems-api-key: YOUR_API_KEY"
    assert_not_includes curl, "--output"
  end

  test "python GET with no params uses requests.get" do
    py = snippets_for(endpoint)["python"]
    assert_includes py, "import requests"
    assert_includes py, "requests.get"
    assert_includes py, "response.json()[\"data\"]"
  end

  test "javascript GET with no params uses fetch and destructures data" do
    js = snippets_for(endpoint)["javascript"]
    assert_includes js, "await fetch('#{BASE_URL}/v1/test'"
    assert_includes js, "const { data } = await response.json();"
  end

  test "ruby GET with no params uses Net::HTTP::Get" do
    rb = snippets_for(endpoint)["ruby"]
    assert_includes rb, "require 'net/http'"
    assert_includes rb, "Net::HTTP::Get.new(uri)"
    assert_includes rb, "JSON.parse(response.body)['data']"
  end

  # ---------------------------------------------------------------------------
  # GET — path param
  # ---------------------------------------------------------------------------

  test "curl GET substitutes path param with example value" do
    ep = endpoint(
      "path"       => "/v1/text/dictionary/{word}",
      "parameters" => [{ "name" => "word", "type" => "string", "location" => "path",
                         "required" => true, "description" => "The word", "example" => "ephemeral" }]
    )
    curl = snippets_for(ep)["curl"]
    assert_includes curl, "/v1/text/dictionary/ephemeral"
    assert_not_includes curl, "{word}"
  end

  test "path param with no example falls back to type placeholder" do
    ep = endpoint(
      "path"       => "/v1/test/{id}",
      "parameters" => [{ "name" => "id", "type" => "string", "location" => "path",
                         "required" => true, "description" => "ID" }]
    )
    curl = snippets_for(ep)["curl"]
    assert_includes curl, "/v1/test/example"
    assert_not_includes curl, "{id}"
  end

  # ---------------------------------------------------------------------------
  # GET — query params
  # ---------------------------------------------------------------------------

  test "curl GET appends query string from query params" do
    ep = endpoint(
      "parameters" => [
        { "name" => "size", "type" => "integer", "location" => "query",
          "required" => false, "description" => "Size", "example" => 256 }
      ]
    )
    curl = snippets_for(ep)["curl"]
    assert_includes curl, "?size=256"
  end

  test "python GET with query params uses params= kwarg" do
    ep = endpoint(
      "parameters" => [
        { "name" => "q", "type" => "string", "location" => "query",
          "required" => true, "description" => "Query", "example" => "hello" }
      ]
    )
    py = snippets_for(ep)["python"]
    assert_includes py, "params ="
    assert_includes py, "requests.get(url, headers=headers, params=params)"
  end

  test "ruby GET with query params uses URI.encode_www_form" do
    ep = endpoint(
      "parameters" => [
        { "name" => "q", "type" => "string", "location" => "query",
          "required" => true, "description" => "Query", "example" => "hello" }
      ]
    )
    rb = snippets_for(ep)["ruby"]
    assert_includes rb, "uri.query = URI.encode_www_form"
  end

  # ---------------------------------------------------------------------------
  # POST — body params
  # ---------------------------------------------------------------------------

  test "curl POST includes method flag and JSON body" do
    ep = endpoint(
      "method"     => "POST",
      "path"       => "/v1/text/sentiment",
      "parameters" => [
        { "name" => "text", "type" => "string", "location" => "body",
          "required" => true, "description" => "Text", "example" => "I love this!" }
      ]
    )
    curl = snippets_for(ep)["curl"]
    assert_includes curl, "curl -X POST"
    assert_includes curl, "Content-Type: application/json"
    assert_includes curl, "I love this!"
  end

  test "python POST uses json= kwarg and prints data" do
    ep = endpoint(
      "method"     => "POST",
      "path"       => "/v1/text/sentiment",
      "parameters" => [
        { "name" => "text", "type" => "string", "location" => "body",
          "required" => true, "description" => "Text", "example" => "I love this!" }
      ]
    )
    py = snippets_for(ep)["python"]
    assert_includes py, "payload ="
    assert_includes py, "json=payload"
    assert_includes py, "response.json()[\"data\"]"
  end

  test "javascript POST uses JSON.stringify and destructures data" do
    ep = endpoint(
      "method"     => "POST",
      "path"       => "/v1/text/sentiment",
      "parameters" => [
        { "name" => "text", "type" => "string", "location" => "body",
          "required" => true, "description" => "Text", "example" => "I love this!" }
      ]
    )
    js = snippets_for(ep)["javascript"]
    assert_includes js, "method: 'POST'"
    assert_includes js, "JSON.stringify"
    assert_includes js, "const { data } = await response.json();"
  end

  test "ruby POST uses Net::HTTP::Post and sets Content-Type" do
    ep = endpoint(
      "method"     => "POST",
      "path"       => "/v1/text/sentiment",
      "parameters" => [
        { "name" => "text", "type" => "string", "location" => "body",
          "required" => true, "description" => "Text", "example" => "I love this!" }
      ]
    )
    rb = snippets_for(ep)["ruby"]
    assert_includes rb, "Net::HTTP::Post.new(uri)"
    assert_includes rb, "request['Content-Type'] = 'application/json'"
    assert_includes rb, ".to_json"
  end

  # ---------------------------------------------------------------------------
  # POST — array body param (batch pattern)
  # ---------------------------------------------------------------------------

  test "curl POST with array<string> body includes array in JSON" do
    ep = endpoint(
      "method"     => "POST",
      "path"       => "/v1/text/sentiment/batch",
      "parameters" => [
        { "name" => "texts", "type" => "array<string>", "location" => "body",
          "required" => true, "description" => "Texts",
          "example"  => '["I love this!", "Terrible."]' }
      ]
    )
    curl = snippets_for(ep)["curl"]
    assert_includes curl, "I love this!"
    assert_includes curl, "Terrible."
  end

  # ---------------------------------------------------------------------------
  # response_kind: binary
  # ---------------------------------------------------------------------------

  test "curl binary GET adds --output flag" do
    ep = endpoint(
      "path"          => "/v1/technology/barcode",
      "response_kind" => "binary",
      "parameters"    => [
        { "name" => "data", "type" => "string", "location" => "query",
          "required" => true, "description" => "Data", "example" => "12345" }
      ]
    )
    curl = snippets_for(ep)["curl"]
    assert_includes curl, "--output response.png"
    assert_not_includes curl, "response.json"
  end

  test "python binary GET writes response.content to file" do
    ep = endpoint("response_kind" => "binary")
    py = snippets_for(ep)["python"]
    assert_includes py, "response.content"
    assert_includes py, "open(\"response.png\", \"wb\")"
    assert_not_includes py, "response.json"
  end

  test "javascript binary GET creates download link" do
    ep = endpoint("response_kind" => "binary")
    js = snippets_for(ep)["javascript"]
    assert_includes js, "response.blob()"
    assert_includes js, "URL.createObjectURL"
    assert_includes js, "a.download"
    assert_not_includes js, "response.json"
  end

  test "ruby binary GET writes body to file without require json" do
    ep = endpoint("response_kind" => "binary")
    rb = snippets_for(ep)["ruby"]
    assert_includes rb, "File.write('response.png'"
    assert_not_includes rb, "require 'json'"
    assert_not_includes rb, "JSON.parse"
  end

  # ---------------------------------------------------------------------------
  # native_example edge cases
  # ---------------------------------------------------------------------------

  test "native_example parses JSON array string" do
    ep = endpoint(
      "method"     => "POST",
      "parameters" => [
        { "name" => "items", "type" => "array<string>", "location" => "body",
          "required" => true, "description" => "Items",
          "example"  => '["a", "b", "c"]' }
      ]
    )
    curl = snippets_for(ep)["curl"]
    assert_includes curl, '["a","b","c"]'
  end

  test "native_example with malformed JSON falls back to raw string" do
    ep = endpoint(
      "method"     => "POST",
      "parameters" => [
        { "name" => "data", "type" => "string", "location" => "body",
          "required" => true, "description" => "Data",
          "example"  => "[not valid json" }
      ]
    )
    assert_nothing_raised { snippets_for(ep) }
    curl = snippets_for(ep)["curl"]
    assert_includes curl, "not valid json"
  end

  test "native_example uses type placeholder when no example given" do
    ep = endpoint(
      "method"     => "POST",
      "parameters" => [
        { "name" => "count", "type" => "integer", "location" => "body",
          "required" => true, "description" => "Count" }
      ]
    )
    curl = snippets_for(ep)["curl"]
    assert_includes curl, '"count":1'
  end

  test "native_example falls back to default when no example given" do
    ep = endpoint(
      "method"     => "POST",
      "parameters" => [
        { "name" => "variant", "type" => "string", "location" => "body",
          "required" => false, "description" => "Variant", "default" => "standard" }
      ]
    )
    curl = snippets_for(ep)["curl"]
    assert_includes curl, '"variant":"standard"'
  end

  test "body_hash skips nested items[]. params" do
    ep = endpoint(
      "method" => "POST",
      "parameters" => [
        { "name" => "items",      "type" => "array<object>", "location" => "body",
          "required" => true, "description" => "items" },
        { "name" => "items[].from", "type" => "integer", "location" => "body",
          "required" => true, "description" => "from base" },
        { "name" => "items[].to",   "type" => "integer", "location" => "body",
          "required" => true, "description" => "to base" }
      ]
    )
    curl = snippets_for(ep)["curl"]
    assert_not_includes curl, "items[].from"
    assert_not_includes curl, "items[].to"
    assert_includes curl, '"items"'
  end
end
