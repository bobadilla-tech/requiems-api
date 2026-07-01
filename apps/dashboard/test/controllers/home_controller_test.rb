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

  test "security page is successful" do
    get "/en/security"
    assert_response :success
  end

  test "contact_submit with valid params delivers mail and redirects home" do
    assert_emails 1 do
      post "/en/contact", params: { name: "Alice", email: "alice@example.com", inquiry_type: "general", message: "Hello" }
    end
    assert_redirected_to root_path
  end

  test "contact_submit with missing name redirects with alert" do
    assert_no_emails do
      post "/en/contact", params: { name: "", email: "alice@example.com", message: "Hello" }
    end
    assert_redirected_to contact_path
  end

  test "contact_submit with invalid email redirects with alert" do
    assert_no_emails do
      post "/en/contact", params: { name: "Alice", email: "not-an-email", message: "Hello" }
    end
    assert_redirected_to contact_path
  end

  test "contact_submit with missing message redirects with alert" do
    assert_no_emails do
      post "/en/contact", params: { name: "Alice", email: "alice@example.com", message: "" }
    end
    assert_redirected_to contact_path
  end
end
