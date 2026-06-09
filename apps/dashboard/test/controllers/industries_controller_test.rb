# frozen_string_literal: true

require "test_helper"

class IndustriesControllerTest < ActionDispatch::IntegrationTest
  test "industries index is successful" do
    get "/en/industries"
    assert_response :success
  end

  test "valid industry slug page is successful" do
    get "/en/industries/fintech"
    assert_response :success
  end

  test "another valid industry slug is successful" do
    get "/en/industries/healthcare"
    assert_response :success
  end

  test "unknown industry slug returns not found" do
    get "/en/industries/not-a-real-industry-slug-12345"
    assert_response :not_found
  end

  test "static routes are not captured by industry slug route" do
    get "/en/contact"
    assert_response :success
  end
end
