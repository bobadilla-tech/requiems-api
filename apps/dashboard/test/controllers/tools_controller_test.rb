# frozen_string_literal: true

require "test_helper"

class ToolsControllerTest < ActionDispatch::IntegrationTest
  test "email validator tool page renders successfully" do
    get "/en/tools/email-validator"
    assert_response :success
  end

  test "unit conversion tool page renders successfully" do
    get "/en/tools/unit-conversion"
    assert_response :success
    assert_select "[data-controller='unit-conversion-demo']"
  end

  test "sentiment analysis tool page renders successfully" do
    get "/en/tools/sentiment-analysis"
    assert_response :success
  end

  test "email normalizer tool page renders successfully" do
    get "/en/tools/email-normalizer"
    assert_response :success
  end

  test "quotes tool page renders successfully" do
    get "/en/tools/quotes"
    assert_response :success
  end

  test "bin lookup tool page renders successfully" do
    get "/en/tools/bin-lookup"
    assert_response :success
  end

  test "unknown tool id redirects to root with alert" do
    get "/en/tools/not-a-real-tool"
    assert_redirected_to root_path
    assert_equal "Tool not found.", flash[:alert]
  end

  test "inflation tool page renders successfully" do
    get "/en/tools/inflation"
    assert_response :success
  end

  test "qr code tool page renders successfully" do
    get "/en/tools/qr-code"
    assert_response :success
  end

  test "profanity filter tool page renders successfully" do
    get "/en/tools/profanity-filter"
    assert_response :success
  end

  test "useragent tool page renders successfully" do
    get "/en/tools/useragent"
    assert_response :success
  end

  test "timezone tool page renders successfully" do
    get "/en/tools/timezone"
    assert_response :success
  end

  test "trivia tool page renders successfully" do
    get "/en/tools/trivia"
    assert_response :success
    assert_select "[data-controller='trivia-demo']"
  end

  test "vpn detection tool page renders successfully" do
    get "/en/tools/vpn-detection"
    assert_response :success
  end

  test "thesaurus tool page renders successfully" do
    get "/en/tools/thesaurus"
    assert_response :success
    assert_select "[data-controller='thesaurus-demo']"
  end

  test "spell check tool page renders successfully" do
    get "/en/tools/spell-check"
    assert_response :success
    assert_select "[data-controller='spell-check-demo']"
  end

  test "random user tool page renders successfully" do
    get "/en/tools/random-user"
    assert_response :success
    assert_select "[data-controller='random-user-demo']"
  end

  test "sudoku tool page renders successfully" do
    get "/en/tools/sudoku"
    assert_response :success
    assert_select "[data-controller='sudoku-demo']"
  end

  test "number base conversion tool page renders successfully" do
    get "/en/tools/number-base-conversion"
    assert_response :success
    assert_select "[data-controller='number-base-conversion-demo']"
  end

  test "mx lookup tool page renders successfully" do
    get "/en/tools/mx-lookup"
    assert_response :success
    assert_select "[data-controller='mx-lookup-demo']"
  end
end
