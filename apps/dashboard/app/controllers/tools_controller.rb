# frozen_string_literal: true

class ToolsController < ApplicationController
  SUPPORTED_TOOLS = %w[email-validator].freeze

  def show
    @tool_id = params[:id]

    unless SUPPORTED_TOOLS.include?(@tool_id)
      redirect_to root_path, alert: "Tool not found."
      return
    end

    render "tools/#{@tool_id.tr('-', '_')}/show"
  end
end
