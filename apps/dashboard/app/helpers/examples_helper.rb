# frozen_string_literal: true

module ExamplesHelper
  CATEGORY_KEYS = {
    "Web Apps" => :web_apps,
    "Mobile Apps" => :mobile_apps,
    "CLI Tools" => :cli_tools,
    "Bots & Automation" => :bots_automation,
    "Backend" => :backend
  }.freeze

  DIFFICULTY_KEYS = {
    "Beginner" => :beginner,
    "Intermediate" => :intermediate,
    "Advanced" => :advanced
  }.freeze

  def example_localized(example, field)
    id = example["id"]
    translated = t("examples.items.#{id}.#{field}", default: nil)
    translated.presence || example[field]
  end

  def example_category_label(name)
    key = CATEGORY_KEYS[name]
    return name unless key

    t("examples.categories.#{key}", default: name)
  end

  def example_difficulty_label(difficulty)
    key = DIFFICULTY_KEYS[difficulty]
    return difficulty unless key

    t("examples.difficulties.#{key}", default: difficulty)
  end

  def example_difficulty_variant(difficulty)
    case difficulty
    when "Beginner" then "success"
    when "Intermediate" then "warning"
    when "Advanced" then "danger"
    else "default"
    end
  end
end
