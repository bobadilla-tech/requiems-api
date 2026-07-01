# frozen_string_literal: true

require "test_helper"

class SalesMailerTest < ActionMailer::TestCase
  test "enterprise_inquiry sends to OBSERVER_EMAILS with company in subject" do
    email = SalesMailer.enterprise_inquiry({ name: "Bob", email: "bob@acme.com", company: "Acme Corp", message: "" })

    assert_equal OBSERVER_EMAILS, email.to
    assert_match "Enterprise Inquiry: Acme Corp", email.subject
    assert_equal [ "bob@acme.com" ], email.reply_to
  end

  test "contact_inquiry sends to OBSERVER_EMAILS with inquiry type and name in subject" do
    email = SalesMailer.contact_inquiry({ name: "Alice", email: "alice@example.com", inquiry: "billing", message: "Need help" })

    assert_equal OBSERVER_EMAILS, email.to
    assert_match "Contact Form: [billing] from Alice", email.subject
    assert_equal [ "alice@example.com" ], email.reply_to
  end
end
