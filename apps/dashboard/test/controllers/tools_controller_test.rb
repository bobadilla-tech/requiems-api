# frozen_string_literal: true

require "test_helper"

class ToolsControllerTest < ActionDispatch::IntegrationTest
  test "email validator tool page renders successfully" do
    get "/en/tools/email-validator"
    assert_response :success
  end

  test "sentiment analysis tool page renders successfully" do
    get "/en/tools/sentiment-analysis"
    assert_response :success
  end

  test "email normalizer tool page renders successfully" do
    get "/en/tools/email-normalizer"
    assert_response :success
  end

  test "unknown tool id redirects to root with alert" do
    get "/en/tools/not-a-real-tool"
    assert_redirected_to root_path
    assert_equal "Tool not found.", flash[:alert]
  end
end
