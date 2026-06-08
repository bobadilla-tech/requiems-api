# frozen_string_literal: true

class ComparisonsController < ApplicationController
  COMPETITORS = [
    { slug: "zerobounce",   category: :email_verification, tier: :primary },
    { slug: "neverbounce",  category: :email_verification, tier: :primary },
    { slug: "kickbox",      category: :email_verification, tier: :primary },
    { slug: "mailboxlayer", category: :email_verification, tier: :secondary },
    { slug: "bouncify",     category: :email_verification, tier: :secondary },
    { slug: "ipstack",      category: :ip_geolocation,     tier: :primary },
    { slug: "abstractapi",  category: :multi_purpose,      tier: :primary },
    { slug: "neutrinoapi",  category: :multi_purpose,      tier: :primary },
    { slug: "api-ninjas",   category: :multi_purpose,      tier: :secondary },
    { slug: "numverify",    category: :phone_validation,   tier: :primary }
  ].freeze

  def index
    @competitors = COMPETITORS
  end

  def show
    @slug = params[:slug]
    @competitor = COMPETITORS.find { |c| c[:slug] == @slug }
    head :not_found and return unless @competitor
  end
end
