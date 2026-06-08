# frozen_string_literal: true

require "test_helper"

class ComparisonsControllerTest < ActionDispatch::IntegrationTest
  test "comparisons hub is successful" do
    get "/en/compare"
    assert_response :success
  end

  test "zerobounce comparison page is successful" do
    get "/en/compare/zerobounce"
    assert_response :success
  end

  test "neverbounce comparison page is successful" do
    get "/en/compare/neverbounce"
    assert_response :success
  end

  test "api-ninjas comparison page is successful" do
    get "/en/compare/api-ninjas"
    assert_response :success
  end

  test "unknown comparison slug returns not found" do
    get "/en/compare/not-a-real-competitor"
    assert_response :not_found
  end

  test "comparisons hub assigns all ten competitors" do
    get "/en/compare"
    assert_response :success
    assert_equal 10, ComparisonsController::COMPETITORS.length
  end
end
