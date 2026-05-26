# frozen_string_literal: true

class PwaController < ApplicationController
  skip_before_action :authenticate_user!, raise: false

  def manifest
    render layout: false
  end
end
