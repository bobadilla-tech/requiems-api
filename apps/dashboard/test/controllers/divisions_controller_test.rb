# frozen_string_literal: true

require "test_helper"

class DivisionsControllerTest < ActionDispatch::IntegrationTest
  test "divisions index is successful" do
    get "/en/divisions"
    assert_response :success
    assert_select "h1", text: /Ship faster/i
  end

  test "validation division page is successful" do
    get "/en/validation"
    assert_response :success
    assert_select "h1", text: /Validate data before it breaks your app/i
  end

  test "finance division page is successful" do
    get "/en/finance"
    assert_response :success
    assert_select "h1", text: /Move money with cleaner data/i
  end

  test "unknown single segment does not route to divisions" do
    get "/en/not_a_real_division_slug_12345"
    assert_response :not_found
  end

  test "static routes are not captured by division slug route" do
    get "/en/contact"
    assert_response :success
  end
end
