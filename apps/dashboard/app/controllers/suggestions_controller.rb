# frozen_string_literal: true

class SuggestionsController < ApplicationController
  def new
    # Render form
  end

  def create
    if valid_suggestion?
      # Send email to observers
      SuggestionMailer.new_suggestion(suggestion_params).deliver_now

      redirect_to root_path, notice: t("suggestions.flash.success")
    else
      flash.now[:alert] = t("suggestions.flash.missing_fields")
      render :new, status: :unprocessable_entity
    end
  rescue StandardError => e
    Rails.logger.error "Failed to send suggestion email: #{e.message}"
    flash.now[:alert] = t("suggestions.flash.error")
    render :new, status: :unprocessable_entity
  end

  private

  def suggestion_params
    params.require(:suggestion).permit(:api_name, :description, :use_case, :email)
  end

  def valid_suggestion?
    params[:suggestion].present? &&
      params[:suggestion][:api_name].present? &&
      params[:suggestion][:description].present? &&
      params[:suggestion][:use_case].present?
  end
end
