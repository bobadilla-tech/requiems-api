# frozen_string_literal: true

require "test_helper"

class SystemsControllerTest < ActionDispatch::IntegrationTest
  test "systems index is successful" do
    get "/en/systems"
    assert_response :success
  end

  test "identity-risk system page is successful" do
    get "/en/systems/identity-risk"
    assert_response :success
  end

  test "payments-intelligence system page is successful" do
    get "/en/systems/payments-intelligence"
    assert_response :success
  end

  test "global-data system page is successful" do
    get "/en/systems/global-data"
    assert_response :success
  end

  test "data-integrity system page is successful" do
    get "/en/systems/data-integrity"
    assert_response :success
  end

  test "developer-utilities system page is successful" do
    get "/en/systems/developer-utilities"
    assert_response :success
  end

  test "unknown system slug returns not found" do
    get "/en/systems/not-a-real-system"
    assert_response :not_found
  end

  test "systems index assigns all five systems" do
    get "/en/systems"
    assert_response :success
    assert_equal 5, SystemsController::SYSTEMS.length
  end

  test "identity-risk page renders engine name" do
    get "/en/systems/identity-risk"
    assert_response :success
    assert_select "h2", text: /Identity & Risk Engine/i
  end

  test "identity-risk page renders API reference section" do
    get "/en/systems/identity-risk"
    assert_response :success
    assert_select "#api-reference"
  end

  test "data-integrity page renders API reference section" do
    get "/en/systems/data-integrity"
    assert_response :success
    assert_select "#api-reference"
  end
end
