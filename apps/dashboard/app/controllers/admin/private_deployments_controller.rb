# frozen_string_literal: true

class Admin::PrivateDeploymentsController < Admin::BaseController
  before_action :set_deployment_request, only: [ :show, :activate, :cancel ]

  def index
    @deployment_requests = PrivateDeploymentRequest.includes(:user)
                                                   .order(created_at: :desc)

    @deployment_requests = case params[:status]
    when "pending"   then @deployment_requests.pending
    when "deploying" then @deployment_requests.deploying
    when "active"    then @deployment_requests.active
    when "cancelled" then @deployment_requests.cancelled
    else @deployment_requests
    end
  end

  def show
  end

  def activate
    p = deployment_activation_params
    subdomain_slug = p[:subdomain_slug].to_s.strip.downcase
    tenant_secret  = p[:tenant_secret].to_s.strip
    admin_notes    = p[:admin_notes].to_s.strip

    if subdomain_slug.blank?
      redirect_to admin_private_deployment_path(@deployment_request), alert: t("admin.private_deployments.subdomain_required") and return
    end

    if tenant_secret.blank?
      redirect_to admin_private_deployment_path(@deployment_request), alert: t("admin.private_deployments.secret_required") and return
    end

    merged_notes = [ @deployment_request.admin_notes, admin_notes.presence ].compact.join("\n\n---\n\n")

    @deployment_request.update!(
      subdomain_slug: subdomain_slug,
      tenant_secret: tenant_secret,
      admin_notes: merged_notes.presence,
      status: "active",
      deployed_at: Time.current
    )

    begin
      PrivateDeploymentMailer.deployment_ready(@deployment_request).deliver_later
    rescue StandardError => e
      Rails.logger.error "[Admin::PrivateDeploymentsController] Failed to enqueue deployment_ready email for request #{@deployment_request.id}: #{e.message}"
    end

    redirect_to admin_private_deployment_path(@deployment_request),
      notice: t("admin.private_deployments.activate_success", email: @deployment_request.contact_email)
  rescue ActiveRecord::RecordInvalid => e
    Rails.logger.error "[Admin::PrivateDeploymentsController#activate] #{e.class}: #{e.message}"
    redirect_to admin_private_deployment_path(@deployment_request), alert: t("admin.private_deployments.activate_error")
  end

  def cancel
    @deployment_request.update!(status: "cancelled")
    redirect_to admin_private_deployments_path, notice: t("admin.private_deployments.cancel_success")
  rescue ActiveRecord::RecordInvalid => e
    Rails.logger.error "[Admin::PrivateDeploymentsController#cancel] #{e.class}: #{e.message}"
    redirect_to admin_private_deployment_path(@deployment_request), alert: t("admin.private_deployments.cancel_error")
  end

  private

  def deployment_activation_params
    params.permit(:subdomain_slug, :tenant_secret, :admin_notes)
  end

  def set_deployment_request
    @deployment_request = PrivateDeploymentRequest.find(params[:id])
  rescue ActiveRecord::RecordNotFound
    redirect_to admin_private_deployments_path, alert: t("admin.private_deployments.not_found")
  end
end
