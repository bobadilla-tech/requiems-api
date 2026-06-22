# frozen_string_literal: true

require "test_helper"

class HomeControllerTest < ActionDispatch::IntegrationTest
  test "domain checker page is successful" do
    get "/en/domain-checker"
    assert_response :success
  end

  test "domain checker page renders hero headline" do
    get "/en/domain-checker"
    assert_select "h1", text: /Check Any Domain/i
  end

  test "domain checker page renders stimulus demo controller" do
    get "/en/domain-checker"
    assert_select "[data-controller='domain-checker-demo']"
  end

  test "domain checker page renders FAQ section" do
    get "/en/domain-checker"
    assert_select "[data-controller='faq-accordion']"
  end
end
