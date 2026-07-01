# frozen_string_literal: true

class SalesMailer < ApplicationMailer
  # Send enterprise inquiry notification to observers
  #
  # @param inquiry [Hash] The inquiry details
  #   @option inquiry [String] :name Contact person name
  #   @option inquiry [String] :email Contact email
  #   @option inquiry [String] :company Company name
  #   @option inquiry [String] :message Optional additional message
  def enterprise_inquiry(inquiry)
    @inquiry = inquiry
    @timestamp = Time.current

    mail(
      to: OBSERVER_EMAILS,
      subject: "Enterprise Inquiry: #{inquiry[:company]}",
      reply_to: inquiry[:email]
    )
  end

  def contact_inquiry(inquiry)
    @inquiry = inquiry
    @timestamp = Time.current

    mail(
      to: OBSERVER_EMAILS,
      subject: "Contact Form: [#{inquiry[:inquiry]}] from #{inquiry[:name]}",
      reply_to: inquiry[:email]
    )
  end
end
