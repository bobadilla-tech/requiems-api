# frozen_string_literal: true

class BlogController < ApplicationController
  def index
    @posts = BlogPost.all
  end

  def show
    @post = BlogPost.find(params[:slug])
    head :not_found unless @post
  end
end
