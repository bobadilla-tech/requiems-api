# frozen_string_literal: true

class PwaController < ApplicationController
  def manifest
    render layout: false
  end
end
