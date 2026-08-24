# frozen_string_literal: true

class PrivateDeploymentMailer < ApplicationMailer
  # Sent to the customer immediately after they submit the form.
  def request_received(deployment_request)
    @deployment_request = deployment_request
    @docs_url = "https://requiems.xyz/docs"

    mail(
      to: @deployment_request.contact_email,
      subject: "Your Private Deployment Request — We're On It"
    )
  end

  # Sent to OBSERVER_EMAILS so the admin knows a new request came in.
  def admin_notification(deployment_request)
    @deployment_request = deployment_request
    @admin_url = admin_private_deployment_url(@deployment_request)

    mail(
      to: OBSERVER_EMAILS,
      reply_to: @deployment_request.contact_email,
      subject: "New Private Deployment Request: #{@deployment_request.company}"
    )
  end

  # Sent to the customer when the admin marks the deployment as active.
  def deployment_ready(deployment_request)
    @deployment_request = deployment_request
    @docs_url = "https://requiems.xyz/docs"

    mail(
      to: @deployment_request.contact_email,
      subject: "Your Private API is Live — #{@deployment_request.live_url}"
    )
  end
end
