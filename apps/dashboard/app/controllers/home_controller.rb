# frozen_string_literal: true

class HomeController < ApplicationController
  API_CATALOG = YAML.load_file(Rails.root.join("config", "api_catalog.yml")).freeze

  def index
    categories = API_CATALOG["categories"]

    # Sort categories by priority: live first, then most important coming soon
    priority_order = %w[
      email text finance technology ai_vision data validation
      security conversion health entertainment places transportation
      animals tax misc
    ]

    @categories = categories.sort_by do |cat|
      [ cat["coming_soon"] ? 1 : 0, priority_order.index(cat["id"]) || 999 ]
    end
  end

  def docs
  end

  def pricing
  end

  def blog
  end

  def status
  end

  def glossary
  end

  def error_codes
  end

  def faq
  end

  def team
  end

  def for_llms
  end

  def domain_checker
  end

  def security
  end

  def contact_submit
    name    = params[:name].to_s.strip
    email   = params[:email].to_s.strip
    inquiry = params[:inquiry_type].to_s.strip
    message = params[:message].to_s.strip

    unless name.present? && email.match?(URI::MailTo::EMAIL_REGEXP) && message.present?
      redirect_to contact_path, alert: t("home.contact.form.missing_fields")
      return
    end

    SalesMailer.contact_inquiry({ name:, email:, inquiry:, message: }).deliver_now
    redirect_to root_path, notice: t("home.contact.form.success")
  rescue => e
    Rails.logger.error "Contact form error: #{e.message}"
    redirect_to contact_path, alert: t("home.contact.form.error")
  end
end
