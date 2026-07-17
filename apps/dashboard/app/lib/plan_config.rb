# frozen_string_literal: true

class PlanConfig
  PLAN_NAMES = %w[free developer business professional].freeze
  PAID_PLAN_NAMES = %w[developer business professional].freeze

  PLANS = {
    "free" => {
      name: "Free",
      price_monthly: 0,
      price_yearly: 0,
      requests_per_month: 500,
      rate_limit_per_minute: 30,
      features: [
        "500 requests/month",
        "30 requests/minute",
        "Community support",
        "US data centers"
      ]
    },
    "developer" => {
      name: "Developer",
      price_monthly: 30,
      price_yearly: 300,
      requests_per_month: 100_000,
      rate_limit_per_minute: 5_000,
      features: [
        "100,000 requests/month",
        "5,000 requests/minute",
        "Email support",
        "US data centers"
      ]
    },
    "business" => {
      name: "Business",
      price_monthly: 75,
      price_yearly: 750,
      requests_per_month: 1_000_000,
      rate_limit_per_minute: 10_000,
      features: [
        "1M requests/month",
        "10,000 requests/minute",
        "Priority email support",
        "US & EU data centers",
        "99.9% SLA"
      ]
    },
    "professional" => {
      name: "Professional",
      price_monthly: 150,
      price_yearly: 1500,
      requests_per_month: 10_000_000,
      rate_limit_per_minute: 50_000,
      features: [
        "10M requests/month",
        "50,000 requests/minute",
        "24/7 priority support",
        "US & EU data centers",
        "99.99% SLA",
        "Dedicated support engineer"
      ]
    }
  }.freeze

  def self.for(plan_name)
    base = PLANS.fetch(plan_name.to_s, PLANS["free"]).dup
    if plan_name.to_s != "free"
      config = AppConfig.instance
      base[:lemonsqueezy_variant_id_monthly] = config.public_send(:"lemonsqueezy_#{plan_name}_monthly_variant_id")
      base[:lemonsqueezy_variant_id_yearly]  = config.public_send(:"lemonsqueezy_#{plan_name}_yearly_variant_id")
    end
    base
  end

  def self.all
    PLAN_NAMES.map { |name| self.for(name).merge(id: name) }
  end

  def self.requests_per_month(plan_name)
    PLANS.fetch(plan_name.to_s, PLANS["free"])[:requests_per_month]
  end

  def self.paid?(plan_name)
    PAID_PLAN_NAMES.include?(plan_name.to_s)
  end

  # Monthly-equivalent price when billed yearly, e.g. 750 / 12 => 62.5
  def self.price_yearly_monthly(plan_name)
    value = PLANS.fetch(plan_name.to_s, PLANS["free"])[:price_yearly] / 12.0
    value % 1 == 0 ? value.to_i : value
  end

  def self.formatted_requests(plan_name)
    humanize_count(requests_per_month(plan_name))
  end

  def self.formatted_rate_limit(plan_name)
    "#{humanize_count(PLANS.fetch(plan_name.to_s, PLANS["free"])[:rate_limit_per_minute])}/min"
  end

  def self.humanize_count(n)
    return n.to_s if n < 1_000
    return "#{trim_zero(n / 1_000.0)}k" if n < 1_000_000

    "#{trim_zero(n / 1_000_000.0)}M"
  end

  def self.trim_zero(value)
    value % 1 == 0 ? value.to_i : value
  end
  private_class_method :humanize_count, :trim_zero
end
