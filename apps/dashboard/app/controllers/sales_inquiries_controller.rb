# frozen_string_literal: true

class SalesInquiriesController < ApplicationController
  def new
    # Render form
  end

  def create
    if valid_inquiry?
      # Send email to observers
      SalesMailer.enterprise_inquiry(inquiry_params).deliver_now

      redirect_to root_path, notice: t("sales_inquiries.flash.success")
    else
      flash.now[:alert] = t("sales_inquiries.flash.missing_fields")
      render :new, status: :unprocessable_entity
    end
  rescue StandardError => e
    Rails.logger.error "Failed to send sales inquiry email: #{e.message}"
    flash.now[:alert] = t("sales_inquiries.flash.error")
    render :new, status: :unprocessable_entity
  end

  private

  def inquiry_params
    params.require(:inquiry).permit(:name, :email, :company, :message)
  end

  def valid_inquiry?
    params[:inquiry].present? &&
      params[:inquiry][:name].present? &&
      params[:inquiry][:email].present? &&
      params[:inquiry][:company].present? &&
      valid_email?(params[:inquiry][:email])
  end

  def valid_email?(email)
    email =~ URI::MailTo::EMAIL_REGEXP
  end
end
