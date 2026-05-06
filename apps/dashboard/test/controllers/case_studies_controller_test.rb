# frozen_string_literal: true

require "test_helper"

class CaseStudiesControllerTest < ActionDispatch::IntegrationTest
  test "case studies hub is successful" do
    get "/en/case-studies"
    assert_response :success
    assert_select "h1", text: /Built with teams/i
  end

  test "verigeo case study is successful" do
    get "/en/case-studies/verigeo"
    assert_response :success
    assert_select "h1", text: /Verigeo scales geomarketing/i
  end

  test "compilestrength case study is successful" do
    get "/en/case-studies/compilestrength"
    assert_response :success
    assert_select "h1", text: /CompileStrength trains smarter/i
  end

  test "unknown case study slug returns not found" do
    get "/en/case-studies/not-a-partner"
    assert_response :not_found
  end

  test "case-studies path is not captured by division slug route" do
    get "/en/case-studies"
    assert_response :success
  end
end
